package openapiconnectorcore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	zerolog "github.com/rs/zerolog"

	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
)

const TestAccessKey = "myAccessKey"
const TestAssetName = "Fake data"
const TestAuthPath = "kubernetes"
const TestBucket = "fakeBucket"
const TestDatabase = "default"
const TestDatabaseService = "openmetadata-s3"
const TestDataFormat = "csv"
const TestDestinationAssetID = "assetID"
const TestDestinationCatalogID = "openmetadata"
const TestGeography = "Hogwarts"
const TestJWT = "myJWT"
const TestObjectName = "csvAsset"
const TestPathInVault = "/v1/" + TestAuthPath + "/info?namespace=default"
const TestSecretKey = "mySecretKey"
const TestToken = "vaultToken"

const AccessKey = "access_key"
const Address = "address"
const ApplicationJSON = "application/json"
const Auth = "auth"
const AuthPath = "authPath"
const ClientToken = "client_token"
const ContentType = "Content-Type"
const Data = "data"
const DatabaseServicesURI = "/v1/services/databaseServices"
const FullyQualifiedName = "fullyQualifiedName"
const FybrikLowerCase = "fybrik"
const ID = "id"
const JwtFilePath = "jwt_file_path"
const JWT = "jwt"
const RequestID = "request_id"
const Role = "role"
const SecretKey = "secret_key"
const ZeroUUID = "00000000-0000-0000-0000-000000000000"

var logger zerolog.Logger
var jwtFile *os.File
var mockVaultServer *httptest.Server
var mockOMServer *httptest.Server
var vaultConf map[interface{}]interface{}

func getCreateAssetRequest() *models.CreateAssetRequest {
	var destinationAssetID = TestDestinationAssetID
	var name = TestAssetName
	var geography = TestGeography
	var dataFormat = TestDataFormat

	const Name = "name"
	const Age = "age"
	const BuildingNumber = "building_number"
	const Street = "street"
	const City = "city"
	const Postcode = "postcode"

	return &models.CreateAssetRequest{
		DestinationCatalogID: TestDestinationCatalogID,
		DestinationAssetID:   &destinationAssetID,
		Details: models.ResourceDetails{
			DataFormat: &dataFormat,
			Connection: models.Connection{
				Name: S3,
				AdditionalProperties: map[string]interface{}{
					S3: map[string]interface{}{
						"endpoint":   "https://s3.eu-de.cloud-object-storage.appdomain.cloud",
						"region":     "eu-de",
						"bucket":     TestBucket,
						"object_key": TestObjectName,
					},
				},
			},
		},
		ResourceMetadata: models.ResourceMetadata{
			Name:      &name,
			Geography: &geography,
			Columns: []models.ResourceColumn{
				{Name: Name, Tags: map[string]interface{}{Name: "true"}},
				{Name: Age, Tags: map[string]interface{}{Age: "true"}},
				{Name: BuildingNumber, Tags: map[string]interface{}{BuildingNumber: "true"}},
				{Name: Street, Tags: map[string]interface{}{Street: "true"}},
				{Name: City, Tags: map[string]interface{}{City: "true"}},
				{Name: Postcode, Tags: map[string]interface{}{Postcode: "true"}},
			},
		},
	}
}

func mustAsJSON(t *testing.T, in interface{}) []byte {
	result, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func createMockVaultServer(t *testing.T) *httptest.Server {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestBytes, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}

		if r.Method == http.MethodPost && r.RequestURI == "/v1/auth/"+TestAuthPath+"/login" {
			requestMap := make(map[string]interface{})
			err = json.Unmarshal(requestBytes, &requestMap)
			if err != nil {
				t.Error(err)
				return
			}
			if (requestMap[Role].(string) != FybrikLowerCase) ||
				(requestMap[JWT].(string) != TestJWT) {
				t.Error(errors.New("unexpected request format"))
				return
			}

			response := map[string]interface{}{
				RequestID: ZeroUUID,
				Auth: map[string]interface{}{
					ClientToken: TestToken,
				},
			}

			responseBytes := mustAsJSON(t, response)

			w.Header().Add(ContentType, ApplicationJSON)
			_, err = w.Write(responseBytes)
			if err != nil {
				t.Error(err)
			}

			return
		}

		if r.Method == http.MethodGet {
			if r.Header.Get("X-Vault-Token") != TestToken {
				t.Error(err)
			}
			logger.Trace().Msg("About to mock response secret request")
			response := map[string]interface{}{
				Data: map[string]interface{}{
					AccessKey: TestAccessKey,
					SecretKey: TestSecretKey,
				},
			}

			responseBytes := mustAsJSON(t, response)

			w.Header().Add(ContentType, ApplicationJSON)
			_, err = w.Write(responseBytes)
			if err != nil {
				t.Error(err)
			}

			return
		}
		t.Fatal("unexpected request")
	}))

	return svr
}

func handleGetMockOMServer(t *testing.T, r *http.Request) (map[string]interface{}, int) {
	if strings.HasPrefix(r.RequestURI, "/v1/metadata/types?category=entity") {
		var types []interface{}
		types = append(types, map[string]interface{}{FullyQualifiedName: "table", ID: "1"})
		return map[string]interface{}{Data: types}, 0
	}
	if r.RequestURI == DatabaseServicesURI {
		var databaseServices []interface{}
		return map[string]interface{}{Data: databaseServices}, 0
	}
	if strings.HasPrefix(r.RequestURI, "/v1/metadata/types?category=field") {
		var types []interface{}
		types = append(types, map[string]interface{}{FullyQualifiedName: "string", ID: "2"})
		return map[string]interface{}{Data: types}, 0
	}
	if strings.HasPrefix(r.RequestURI,
		fmt.Sprintf("/v1/tables/name/%s.%s.%s.%s", TestDatabaseService, TestDatabase,
			TestBucket, TestObjectName)) {
		return nil, http.StatusNotFound
	}
	if r.RequestURI ==
		fmt.Sprintf("/v1/services/ingestionPipelines/name/%s.%%22pipeline-%s.%s%%22",
			TestDatabaseService, TestDestinationCatalogID, TestDestinationAssetID) {
		return nil, http.StatusNotFound
	}
	if r.RequestURI == fmt.Sprintf("/v1/databases/name/%s.%s", TestDatabaseService, TestDatabase) {
		return nil, http.StatusNotFound
	}
	if r.RequestURI ==
		fmt.Sprintf("/v1/databaseSchemas/name/%s.%s.%s", TestDatabaseService, TestDatabase, TestBucket) {
		return nil, http.StatusNotFound
	}

	t.Fatalf("unrecognized GET request")
	return nil, http.StatusInternalServerError
}

func handlePostMockOMServer(t *testing.T, r *http.Request) (map[string]interface{}, int) {
	if r.RequestURI == "/v1/tags" {
		return map[string]interface{}{}, 0
	}
	if r.RequestURI == DatabaseServicesURI ||
		r.RequestURI == "/v1/services/ingestionPipelines" ||
		r.RequestURI == "/v1/databases" ||
		r.RequestURI == "/v1/databaseSchemas" ||
		r.RequestURI == "/v1/tables" {
		return map[string]interface{}{
			ID:                 ZeroUUID,
			FullyQualifiedName: TestDatabaseService}, 0
	}

	t.Fatalf("unrecognized POST request")
	return nil, http.StatusInternalServerError
}

func createMockOMServer(t *testing.T) *httptest.Server {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var response map[string]interface{}
		var statusCode int

		requestBytes, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
			return
		}

		requestMap := make(map[string]interface{})
		var requestArr []interface{}
		if len(requestBytes) > 0 {
			err = json.Unmarshal(requestBytes, &requestMap)
			if err != nil {
				err = json.Unmarshal(requestBytes, &requestArr)
				if err != nil {
					t.Fatalf(string(requestBytes))
					return
				}
			}
		}

		if r.Method == http.MethodGet {
			response, statusCode = handleGetMockOMServer(t, r)
			if statusCode != 0 {
				w.WriteHeader(statusCode)
				return
			}
		} else if r.Method == http.MethodPost {
			response, statusCode = handlePostMockOMServer(t, r)
			if statusCode != 0 {
				w.WriteHeader(statusCode)
				return
			}
		} else if r.Method == http.MethodPut && r.RequestURI == "/v1/metadata/types/1" {
			response = map[string]interface{}{}
		} else if r.Method == http.MethodPatch && r.RequestURI == "/v1/tables/00000000-0000-0000-0000-000000000000" {
			response = map[string]interface{}{}
		} else {
			t.Fatalf("unexpected request: %s -- %s", r.Method, r.RequestURI)
		}

		responseBytes := mustAsJSON(t, response)

		w.Header().Add(ContentType, ApplicationJSON)
		_, err = w.Write(responseBytes)
		if err != nil {
			t.Error(err)
		}
	}))

	return svr
}
