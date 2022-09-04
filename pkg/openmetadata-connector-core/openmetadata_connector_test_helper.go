package openapiconnectorcore

import (
	"encoding/json"
	"errors"
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
const TestJWT = "myJWT"
const TestPathInVault = "/v1/kubernetes/info?namespace=default"
const TestSecretKey = "mySecretKey"
const TestToken = "vaultToken"

const AccessKey = "access_key"
const Address = "address"
const ApplicationJSON = "application/json"
const Auth = "auth"
const AuthPath = "authPath"
const ClientToken = "client_token"
const ContentType = "Content-Type"
const FybrikLowerCase = "fybrik"
const JwtFilePath = "jwt_file_path"
const JWT = "jwt"
const Kubernetes = "kubernetes"
const RequestID = "request_id"
const Role = "role"
const SecretKey = "secret_key"

var logger zerolog.Logger
var jwtFile *os.File
var mockVaultServer *httptest.Server
var mockOMServer *httptest.Server
var vaultConf map[interface{}]interface{}

func getCreateAssetRequest() *models.CreateAssetRequest {
	var destinationAssetID = "assetID"
	var name = "Fake data"
	var geography = "Hogwarts"
	var dataFormat = "csv"

	return &models.CreateAssetRequest{
		DestinationCatalogID: "openmetadata",
		DestinationAssetID:   &destinationAssetID,
		Details: models.ResourceDetails{
			DataFormat: &dataFormat,
			Connection: models.Connection{
				Name: "s3",
				AdditionalProperties: map[string]interface{}{
					"s3": map[string]interface{}{
						"endpoint":   "https://s3.eu-de.cloud-object-storage.appdomain.cloud",
						"region":     "eu-de",
						"bucket":     "fakeBucket",
						"object_key": "csvAsset",
					},
				},
			},
		},
		ResourceMetadata: models.ResourceMetadata{
			Name:      &name,
			Geography: &geography,
			Columns: []models.ResourceColumn{
				{Name: "name", Tags: map[string]interface{}{"name": "true"}},
				{Name: "age", Tags: map[string]interface{}{"age": "true"}},
				{Name: "building_number", Tags: map[string]interface{}{"building_number": "true"}},
				{Name: "street", Tags: map[string]interface{}{"street": "true"}},
				{Name: "city", Tags: map[string]interface{}{"city": "true"}},
				{Name: "postcode", Tags: map[string]interface{}{"postcode": "true"}},
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

		if r.Method == http.MethodPost && r.RequestURI == "/v1/auth/kubernetes/login" {
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
				RequestID: "00000000-0000-0000-0000-000000000000",
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
				"data": map[string]interface{}{
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

func handleGetMockOMServer(r *http.Request) (map[string]interface{}, int) {
	if r.RequestURI == "/v1/metadata/types?category=entity&limit=100" {
		var types []interface{}
		types = append(types, map[string]interface{}{"fullyQualifiedName": "table", "id": "1"})
		return map[string]interface{}{"data": types}, 0
	}
	if r.RequestURI == "/v1/services/databaseServices" {
		var databaseServices []interface{}
		return map[string]interface{}{"data": databaseServices}, 0
	}
	if r.RequestURI == "/v1/metadata/types?category=field&limit=100" {
		var types []interface{}
		types = append(types, map[string]interface{}{"fullyQualifiedName": "string", "id": "2"})
		return map[string]interface{}{"data": types}, 0
	}
	if strings.HasPrefix(r.RequestURI, "/v1/tables/name/openmetadata-s3.default.fakeBucket.csvAsset") {
		return nil, http.StatusNotFound
	}
	if r.RequestURI == "/v1/services/ingestionPipelines/name/openmetadata-s3.%22pipeline-openmetadata.assetID%22" {
		return nil, http.StatusNotFound
	}
	if r.RequestURI == "/v1/databases/name/openmetadata-s3.default" {
		return nil, http.StatusNotFound
	}
	if r.RequestURI == "/v1/databaseSchemas/name/openmetadata-s3.default.fakeBucket" {
		return nil, http.StatusNotFound
	}

	return nil, http.StatusInternalServerError
}

func handlePostMockOMServer(r *http.Request) (map[string]interface{}, int) {
	if r.RequestURI == "/v1/tags" {
		return map[string]interface{}{}, 0
	}
	if r.RequestURI == "/v1/services/databaseServices" ||
		r.RequestURI == "/v1/services/ingestionPipelines" ||
		r.RequestURI == "/v1/databases" ||
		r.RequestURI == "/v1/databaseSchemas" ||
		r.RequestURI == "/v1/tables" {
		return map[string]interface{}{
			"id":                 "00000000-0000-0000-0000-000000000000",
			"fullyQualifiedName": "openmetadata-s3"}, 0
	}

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
			response, statusCode = handleGetMockOMServer(r)
			if statusCode != 0 {
				w.WriteHeader(statusCode)
				return
			}
		} else if r.Method == http.MethodPost {
			response, statusCode = handlePostMockOMServer(r)
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
