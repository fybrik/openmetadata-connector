// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package openapiconnectorcore

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/fatih/structs"
	. "github.com/onsi/ginkgo/v2" //nolint
	. "github.com/onsi/gomega"    //nolint
	"github.com/rs/zerolog"

	client "fybrik.io/openmetadata-connector/datacatalog-go-client"
	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
	"fybrik.io/openmetadata-connector/pkg/utils"
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
const TestPathInVault = "/v1/omd-secrets/" + TestDatabaseService
const TestSecretKey = "mySecretKey"
const TestToken = "vaultToken"
const TestVersion = 0.1

const AccessKey = "access_key"
const Address = "address"
const ApplicationJSON = "application/json"
const Auth = "auth"
const AuthPath = "authPath"
const ClientToken = "client_token"
const Columns = "columns"
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
const ServiceType = "serviceType"
const TablesURI = "/v1/tables"
const TagFQNStr = "tagFQN"
const ZeroUUID = "00000000-0000-0000-0000-000000000000"

var logger zerolog.Logger
var jwtFile *os.File
var mockVaultServer *httptest.Server
var mockOMServer *httptest.Server
var mockDataCatalog = make(map[string]interface{})
var vaultConf map[interface{}]interface{}

// returns two similar structures: The structure used for requesting an asset creation,
// and a structure representing the response to a getAssetInfo request regarding the
// same asset
func getCreateAssetRequestAndResponse() (*models.CreateAssetRequest, *models.GetAssetResponse) {
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
	const PII = "PII"
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
			{Name: Name, Tags: map[string]interface{}{PII: True}},
			{Name: Age, Tags: map[string]interface{}{PII: True}},
			{Name: BuildingNumber, Tags: map[string]interface{}{PII: True}},
			{Name: Street},
			{Name: City},
			{Name: Postcode},
		},
		Tags: map[string]interface{}{Financial: True},
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

func mustAsJSON(in interface{}) []byte {
	result, err := json.Marshal(in)
	Expect(err).ToNot(HaveOccurred())
	return result
}

// Code for mock vault server. Handles both a request for a token, and a request for a secret
func createMockVaultServer() *httptest.Server {
	By("opreating a mock Vault server")
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestBytes, err := io.ReadAll(r.Body)
		Expect(err).ToNot(HaveOccurred())
		Expect(r.Method).To(SatisfyAny(Equal(http.MethodPost), Equal(http.MethodGet)))
		if r.Method == http.MethodPost {
			By("receiving a Post request")
			Expect(r.RequestURI).To(Equal("/v1/auth/" + TestAuthPath + "/login"))
			requestMap := make(map[string]interface{})
			err = json.Unmarshal(requestBytes, &requestMap)
			Expect(err).ToNot(HaveOccurred())
			Expect(requestMap[Role].(string)).To(Equal(FybrikLowerCase))
			Expect(requestMap[JWT].(string)).To(Equal(TestJWT))
			response := map[string]interface{}{
				RequestID: ZeroUUID,
				Auth: map[string]interface{}{
					ClientToken: TestToken,
				},
			}
			responseBytes := mustAsJSON(response)
			w.Header().Add(ContentType, ApplicationJSON)
			_, err = w.Write(responseBytes)
			Expect(err).ToNot(HaveOccurred())
			return
		}
		if r.Method == http.MethodGet {
			By("receiving a Get request")
			Expect(r.Header.Get("X-Vault-Token")).To(Equal(TestToken))
			logger.Trace().Msg("About to mock response secret request")
			response := map[string]interface{}{
				Data: map[string]interface{}{
					AccessKey: TestAccessKey,
					SecretKey: TestSecretKey,
				},
			}

			responseBytes := mustAsJSON(response)

			w.Header().Add(ContentType, ApplicationJSON)
			_, err = w.Write(responseBytes)
			Expect(err).ToNot(HaveOccurred())
			return
		}
	}))

	return svr
}

// creates a client.DatabaseConnection structure, based on a request sent to the mock
// OM server
func constructDataBaseServiceStruct(serviceInfo map[string]interface{}) *client.DatabaseService {
	connectionInfo, ok := utils.InterfaceToMap(serviceInfo["connection"], &logger)
	if !ok {
		return nil
	}
	connectionConfig, ok := utils.InterfaceToMap(connectionInfo["config"], &logger)
	if !ok {
		return nil
	}

	connection := client.DatabaseConnection{Config: connectionConfig}
	serviceName := serviceInfo[Name].(string)
	return &client.DatabaseService{
		Name:               serviceName,
		FullyQualifiedName: &serviceName,
		ServiceType:        serviceInfo[ServiceType].(string),
		Connection:         connection,
	}
}

// returns a table structure, based on a table creation request sent to the
// mock OM server. This table is later enriched using the patchTable() method
func constructTableStruct(assetInfo map[string]interface{}) (*client.Table, bool) {
	version := TestVersion
	requestColumns, ok := utils.InterfaceToArray(assetInfo[Columns], &logger)
	if !ok {
		return nil, false
	}
	var columns []client.Column
	for i := range requestColumns {
		c, ok := utils.InterfaceToMap(requestColumns[i], &logger)
		if !ok {
			return nil, false
		}
		columns = append(columns, client.Column{Name: c[Name].(string)})
	}

	return &client.Table{
		Id:      ZeroUUID,
		Columns: columns,
		Version: &version,
		Service: &client.EntityReference{Id: ZeroUUID},
		Name:    TestObjectName,
	}, true
}

// enrich table structure using tags, custom properties, and column tags
func patchTable(table *client.Table, patchArray []interface{}) bool { //nolint
	for i := range patchArray {
		patchMap, ok := utils.InterfaceToMap(patchArray[i], &logger)
		if !ok {
			return false
		}
		patchName := patchMap[Path].(string)
		switch {
		case patchName == TagsPath:
			var assetTags []client.TagLabel
			tagsArr, ok := utils.InterfaceToArray(patchMap[Value], &logger)
			if !ok {
				return false
			}
			for j := range tagsArr {
				columnTagMap, ok2 := utils.InterfaceToMap(tagsArr[j], &logger)
				if !ok2 {
					return false
				}
				assetTags = append(assetTags, client.TagLabel{TagFQN: columnTagMap[TagFQNStr].(string)})
			}
			table.Tags = assetTags
		case patchName == ExtensionPath:
			extensionValueMap, ok1 := utils.InterfaceToMap(patchMap[Value], &logger)
			if !ok1 {
				return false
			}
			table.Extension = extensionValueMap
		case patchName == ColumnsPath:
			var columnsWithTags []client.Column
			columnValueArr, ok := utils.InterfaceToArray(patchMap[Value], &logger)
			if !ok {
				return false
			}
			for i := range columnValueArr {
				c, ok := utils.InterfaceToMap(columnValueArr[i], &logger)
				if !ok {
					return false
				}
				var tags []client.TagLabel
				tagsArr, ok := utils.InterfaceToArray(c["tags"], &logger)
				if !ok {
					return false
				}
				for j := range tagsArr {
					columnTagMap, ok := utils.InterfaceToMap(tagsArr[j], &logger)
					if !ok {
						return false
					}
					tags = append(tags, client.TagLabel{TagFQN: columnTagMap[TagFQNStr].(string)})
				}
				columnsWithTags = append(columnsWithTags, client.Column{Name: c[Name].(string), Tags: tags})
			}
			table.Columns = columnsWithTags
		}
	}
	return true
}

// handle mock OM server GET requests
func handleGetMockOMServer(r *http.Request) (map[string]interface{}, int) {
	if strings.HasPrefix(r.RequestURI, "/v1/metadata/types") {
		var types []interface{}
		types = append(types, map[string]interface{}{FullyQualifiedName: Table, ID: "1"},
			map[string]interface{}{FullyQualifiedName: String, ID: "2"})
		return map[string]interface{}{Data: types}, 0
	}
	if r.RequestURI == DatabaseServicesURI {
		var databaseServices []interface{}
		return map[string]interface{}{Data: databaseServices}, 0
	}
	if strings.HasPrefix(r.RequestURI,
		fmt.Sprintf(TablesURI+"/name/%s.%s.%s.%s", TestDatabaseService, TestDatabase,
			TestBucket, TestObjectName)) ||
		strings.HasPrefix(r.RequestURI, TablesURI+"/"+ZeroUUID) {
		table, ok := mockDataCatalog[ZeroUUID].(*client.Table)
		if !ok {
			return nil, http.StatusNotFound
		}
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
		service, ok := mockDataCatalog[DatabaseService]
		if !ok {
			return nil, http.StatusNotFound
		}
		return structs.Map(service), 0
	}

	return nil, http.StatusInternalServerError
}

// handle mock OM server POST requests
func handlePostMockOMServer(r *http.Request,
	requestMap map[string]interface{}) (map[string]interface{}, int) {
	if r.RequestURI == "/v1/users/login" {
		return map[string]interface{}{
			"tokenType":   "Bearer",
			"accessToken": "abc"}, 0
	}
	if r.RequestURI == "/v1/tags" {
		return map[string]interface{}{}, 0
	}
	if r.RequestURI == DatabaseServicesURI {
		// keep the connection information in the mock data catalog
		service := constructDataBaseServiceStruct(requestMap)
		if service == nil {
			return nil, http.StatusInternalServerError
		}
		mockDataCatalog[DatabaseService] = service
	}
	if r.RequestURI == TablesURI {
		table, ok := constructTableStruct(requestMap)
		if !ok {
			return nil, http.StatusInternalServerError
		}
		mockDataCatalog[ZeroUUID] = table
		return structs.Map(table), 0
	}
	if r.RequestURI == DatabaseServicesURI ||
		r.RequestURI == "/v1/services/ingestionPipelines" ||
		r.RequestURI == "/v1/databases" ||
		r.RequestURI == "/v1/databaseSchemas" ||
		r.RequestURI == "/v1/tags/GenericTags" {
		return map[string]interface{}{
			ID:                 ZeroUUID,
			FullyQualifiedName: TestDatabaseService}, 0
	}
	return nil, http.StatusInternalServerError
}

// returns a mock OM server, which handles select GET, POST, PUT, and PATCH requests
func createMockOMServer() *httptest.Server {
	By("opreating a mock OM server")
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var response map[string]interface{}
		var statusCode int
		requestBytes, err := io.ReadAll(r.Body)
		Expect(err).ToNot(HaveOccurred())
		requestMap := make(map[string]interface{})
		var requestArr []interface{}
		if len(requestBytes) > 0 {
			err = json.Unmarshal(requestBytes, &requestMap)
			if err != nil {
				err = json.Unmarshal(requestBytes, &requestArr)
				Expect(err).ToNot(HaveOccurred())
			}
		}
		Expect(r.Method).To(SatisfyAny(Equal(http.MethodGet), Equal(http.MethodPost), Equal(http.MethodPut), Equal(http.MethodPatch)))
		if r.Method == http.MethodGet {
			By("receiving a Get request to the OM server")
			response, statusCode = handleGetMockOMServer(r)
			Expect(statusCode).ToNot(Equal(http.StatusInternalServerError))
			if statusCode != 0 {
				w.WriteHeader(statusCode)
				return
			}
		} else if r.Method == http.MethodPost {
			By("receiving a Post request to the OM server")
			response, statusCode = handlePostMockOMServer(r, requestMap)
			Expect(statusCode).ToNot(Equal(http.StatusInternalServerError))
			if statusCode != 0 {
				w.WriteHeader(statusCode)
				return
			}
		} else if r.Method == http.MethodPut {
			By("receiving a Put request to the OM server")
			Expect(r.RequestURI).To(Equal("/v1/metadata/types/1"))
			response = map[string]interface{}{}
		} else if r.Method == http.MethodPatch {
			By("receiving a Patch request to the OM server")
			Expect(r.RequestURI).To(Equal(TablesURI + "/" + ZeroUUID))
			table, ok := mockDataCatalog[ZeroUUID].(*client.Table)
			Expect(ok).To(BeTrue())
			ok = patchTable(table, requestArr)
			Expect(ok).To(BeTrue())
			mockDataCatalog[ZeroUUID] = table
			response = map[string]interface{}{}
		}
		responseBytes := mustAsJSON(response)
		w.Header().Add(ContentType, ApplicationJSON)
		_, err = w.Write(responseBytes)
		Expect(err).ToNot(HaveOccurred())
	}))

	return svr
}

func clearMockDataCatalog() {
	for k := range mockDataCatalog {
		delete(mockDataCatalog, k)
	}
}
