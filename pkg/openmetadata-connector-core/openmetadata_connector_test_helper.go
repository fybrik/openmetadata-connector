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

	structs "github.com/fatih/structs"
	zerolog "github.com/rs/zerolog"

	client "fybrik.io/openmetadata-connector/datacatalog-go-client"
	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
	utils "fybrik.io/openmetadata-connector/pkg/utils"
)

const TestAccessKey = "myAccessKey"
const TestAssetName = "Fake data"
const TestAuthPath = "kubernetes"
const TestBucket = "fakeBucket"
const TestConnectorCredentials = "fake-credentials"
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
const TestVersion = 0.1

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
const TablesURI = "/v1/tables"
const TagFQNStr = "tagFQN"
const ZeroUUID = "00000000-0000-0000-0000-000000000000"

var logger zerolog.Logger
var jwtFile *os.File
var mockVaultServer *httptest.Server
var mockOMServer *httptest.Server
var mockDataCatalog = make(map[string]interface{})
var vaultConf map[interface{}]interface{}

func getCreateAssetRequestAndReponse() (*models.CreateAssetRequest, *models.GetAssetResponse) {
	var destinationAssetID = TestDestinationAssetID
	var name = TestAssetName
	var geography = TestGeography
	var dataFormat = TestDataFormat
	var credentials = TestPathInVault

	const Name = "name"
	const Age = "age"
	const BuildingNumber = "building_number"
	const Street = "street"
	const City = "city"
	const Postcode = "postcode"
	const Financial = "financial"

	details := models.ResourceDetails{
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
	}

	resourceMetadata := models.ResourceMetadata{
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
		Tags: map[string]interface{}{Financial: "true"},
	}

	createAssetRequest := &models.CreateAssetRequest{
		Credentials:          &credentials,
		DestinationCatalogID: TestDestinationCatalogID,
		DestinationAssetID:   &destinationAssetID,
		Details:              details,
		ResourceMetadata:     resourceMetadata,
	}

	getAssetResponse := &models.GetAssetResponse{
		Credentials:      credentials,
		Details:          details,
		ResourceMetadata: resourceMetadata,
	}

	return createAssetRequest, getAssetResponse
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

func constructDataBaseServiceStruct(serviceInfo map[string]interface{}) client.DatabaseService {
	connectionInfo, ok := utils.InterfaceToMap(serviceInfo["connection"])
	if !ok {
		return client.DatabaseService{}
	}
	connectionConfig, ok := utils.InterfaceToMap(connectionInfo["config"])
	if !ok {
		return client.DatabaseService{}
	}

	connection := client.DatabaseConnection{Config: connectionConfig}
	return client.DatabaseService{
		Name:        serviceInfo[Name].(string),
		ServiceType: serviceInfo["serviceType"].(string),
		Connection:  connection,
	}
}

func constructTableStruct(assetInfo map[string]interface{}) client.Table { //nolint
	var extension interface{}
	var columns interface{}
	var ok bool
	var columnsWithTags []client.Column
	var assetTags []client.TagLabel

	extensionValueMap := make(map[string]interface{})

	var tags interface{}
	if tags, ok = assetInfo["/tags"]; ok {
		tagsMap, ok1 := utils.InterfaceToMap(tags)
		tagsValue := tagsMap[Value]
		tagsArr, ok2 := utils.InterfaceToArray(tagsValue)
		if !ok1 || !ok2 {
			return client.Table{}
		}
		for j := range tagsArr {
			columnTagMap, ok2 := utils.InterfaceToMap(tagsArr[j])
			if !ok2 {
				return client.Table{}
			}
			assetTags = append(assetTags, client.TagLabel{TagFQN: columnTagMap[TagFQNStr].(string)})
		}
	}

	if extension, ok = assetInfo["/extension"]; ok {
		extensionMap, ok1 := utils.InterfaceToMap(extension)
		if !ok1 {
			return client.Table{}
		}
		extensionValue := extensionMap[Value]
		extensionValueMap, ok1 = utils.InterfaceToMap(extensionValue)
		if !ok1 {
			return client.Table{}
		}
	}

	if columns, ok = assetInfo["/columns"]; ok {
		columnsMap, ok := utils.InterfaceToMap(columns)
		if !ok {
			return client.Table{}
		}

		columnValue := columnsMap[Value]
		columnValueArr, ok := utils.InterfaceToArray(columnValue)
		if !ok {
			return client.Table{}
		}

		for i := range columnValueArr {
			c, ok := utils.InterfaceToMap(columnValueArr[i])
			if !ok {
				return client.Table{}
			}
			var tags []client.TagLabel
			tagsArr, ok := utils.InterfaceToArray(c["tags"])
			if !ok {
				return client.Table{}
			}
			for j := range tagsArr {
				columnTagMap, ok := utils.InterfaceToMap(tagsArr[j])
				if !ok {
					return client.Table{}
				}
				tags = append(tags, client.TagLabel{TagFQN: columnTagMap[TagFQNStr].(string)})
			}
			columnsWithTags = append(columnsWithTags, client.Column{Name: c[Name].(string), Tags: tags})
		}
	}

	version := TestVersion
	return client.Table{
		Id:        ZeroUUID,
		Version:   &version,
		Extension: extensionValueMap,
		Columns:   columnsWithTags,
		Service:   &client.EntityReference{Id: ZeroUUID},
		Tags:      assetTags,
		Name:      TestObjectName,
	}
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
		fmt.Sprintf(TablesURI+"/name/%s.%s.%s.%s", TestDatabaseService, TestDatabase,
			TestBucket, TestObjectName)) ||
		strings.HasPrefix(r.RequestURI, TablesURI+"/"+ZeroUUID) {
		assetInfo, ok := mockDataCatalog[ZeroUUID]
		if !ok {
			return nil, http.StatusNotFound
		}
		assetInfoMap, ok := utils.InterfaceToMap(assetInfo)
		if !ok {
			return nil, http.StatusNotFound
		}
		table := constructTableStruct(assetInfoMap)
		return structs.Map(table), 0
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
	if r.RequestURI == "/v1/services/databaseServices/"+ZeroUUID {
		serviceInfo, ok := mockDataCatalog[DatabaseService]
		if !ok {
			return nil, http.StatusNotFound
		}
		serviceInfoMap, ok := utils.InterfaceToMap(serviceInfo)
		if !ok {
			return nil, http.StatusNotFound
		}
		service := constructDataBaseServiceStruct(serviceInfoMap)
		return structs.Map(service), 0
	}

	t.Fatalf("unrecognized GET request")
	return nil, http.StatusInternalServerError
}

func handlePostMockOMServer(t *testing.T, r *http.Request,
	requestMap map[string]interface{}) (map[string]interface{}, int) {
	if r.RequestURI == "/v1/tags" {
		return map[string]interface{}{}, 0
	}

	if r.RequestURI == DatabaseServicesURI {
		// keep the connection information in the mock data catalog
		mockDataCatalog[DatabaseService] = requestMap
	}

	if r.RequestURI == TablesURI {
		mockDataCatalog[ZeroUUID] = requestMap
		requestColumns, ok := utils.InterfaceToArray(requestMap["columns"])
		if !ok {
			return nil, http.StatusInternalServerError
		}
		var columns []client.Column
		for i := range requestColumns {
			c, ok := utils.InterfaceToMap(requestColumns[i])
			if !ok {
				return nil, http.StatusInternalServerError
			}
			columns = append(columns, client.Column{Name: c[Name].(string)})
		}
		table := client.Table{Columns: columns, Id: ZeroUUID}
		return structs.Map(table), 0
	}

	if r.RequestURI == DatabaseServicesURI ||
		r.RequestURI == "/v1/services/ingestionPipelines" ||
		r.RequestURI == "/v1/databases" ||
		r.RequestURI == "/v1/databaseSchemas" ||
		r.RequestURI == "/v1/tags/Fybrik" {
		return map[string]interface{}{
			ID:                 ZeroUUID,
			FullyQualifiedName: TestDatabaseService}, 0
	}

	t.Fatalf("unrecognized POST request")
	return nil, http.StatusInternalServerError
}

func createMockOMServer(t *testing.T) *httptest.Server { //nolint
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
			response, statusCode = handlePostMockOMServer(t, r, requestMap)
			if statusCode != 0 {
				w.WriteHeader(statusCode)
				return
			}
		} else if r.Method == http.MethodPut && r.RequestURI == "/v1/metadata/types/1" {
			response = map[string]interface{}{}
		} else if r.Method == http.MethodPatch && r.RequestURI == TablesURI+"/"+ZeroUUID {
			assetMap, ok := utils.InterfaceToMap(mockDataCatalog[ZeroUUID])
			if !ok {
				t.Fatalf("error transforming asset information to map")
			}
			for i := range requestArr {
				patchMap, ok := utils.InterfaceToMap(requestArr[i])
				if !ok {
					t.Fatalf("error transforming patch to map")
				}
				patchName := patchMap[Path].(string)
				assetMap[patchName] = requestArr[i]
			}
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
