// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package openapiconnectorcore

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"fybrik.io/fybrik/pkg/logging"
	"github.com/fatih/structs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog"

	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
	api "fybrik.io/openmetadata-connector/datacatalog-go/go"
	"fybrik.io/openmetadata-connector/pkg/utils"
	"fybrik.io/openmetadata-connector/pkg/vault"
)

func TestOpenApiConnectorCore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "openApiConnectorCore Suite")
}

var _ = Describe("Vault and OM connector", Ordered, func() {
	var logger zerolog.Logger
	var err error
	var n int
	utils.InitHTTPClient(&logger)
	BeforeEach(func() {
		logger = logging.LogInit(logging.CONNECTOR, "OpenMetadata Connector Tests")
		logger.Trace().Msg("Setting up OM Connector test suite")
		jwtFile, err = ioutil.TempFile("/tmp", "jwt")
		Expect(err).ToNot(HaveOccurred())
		n, err = jwtFile.WriteString(TestJWT)
		Expect(err).ToNot(HaveOccurred())
		Expect(n).To(Equal(len(TestJWT)))
		mockVaultServer = createMockVaultServer()
		mockOMServer = createMockOMServer()

		// populate vault configuration
		vaultConf = make(map[interface{}]interface{})
		vaultConf[Address] = mockVaultServer.URL
		vaultConf[JwtFilePath] = jwtFile.Name()
		vaultConf[AuthPath] = TestAuthPath
		vaultConf[Role] = FybrikLowerCase
		DeferCleanup(func() {
			err = os.Remove(jwtFile.Name())
			Expect(err).ToNot(HaveOccurred())
		})
	})
	Describe("Vault credentials retrieval flow", Ordered, func() {
		var vaultClient *vault.VaultClient
		It("should receive valid access and secret keys", func() {
			vaultClient = vault.NewVaultClient(vaultConf, &logger, utils.HTTPClient)
			testPathInVault := TestPathInVault
			secrets, err1 := vaultClient.GetSecretMap(&testPathInVault)
			Expect(err1).ToNot(HaveOccurred())
			accessKey := secrets[AccessKey].(string)
			secretKey := secrets[SecretKey].(string)
			Expect(accessKey).To(Equal(TestAccessKey))
			Expect(secretKey).To(Equal(TestSecretKey))
			Expect(err).ToNot(HaveOccurred())
		})
	})
	Describe("Connector creating assets and getting their info", Ordered, func() {
		var getAssetResponse *models.GetAssetResponse
		var createAssetRequest *models.CreateAssetRequest
		var response api.ImplResponse
		var servicer OpenMetadataAPIServicer
		var assetID string
		var ctx context.Context
		It("should create an asset and receive a valid response", func() {
			conf := make(map[string]interface{})
			conf["openmetadata_endpoint"] = mockOMServer.URL
			conf["vault"] = vaultConf
			servicer = NewOpenMetadataAPIService(conf, nil, &logger)
			ctx = context.Background()
			createAssetRequest, getAssetResponse = getCreateAssetRequestAndResponse()
			response, err = servicer.CreateAsset(ctx, TestConnectorCredentials, createAssetRequest)
			Expect(err).ToNot(HaveOccurred())
			assetID = fmt.Sprintf("%s.%s.%s.%s", TestDatabaseService, TestDatabase, TestBucket, TestObjectName)
			// check the asset ID in the response
			responseStr := string(mustAsJSON(response.Body))
			Expect(responseStr).To(Equal(fmt.Sprintf("{\"assetID\":%q}", assetID)))
		})
		It("should get the asset info", func() {
			response, err = servicer.GetAssetInfo(ctx, TestConnectorCredentials, &api.GetAssetRequest{AssetID: assetID, OperationType: "read"})
			Expect(err).ToNot(HaveOccurred())
			responseMap := structs.Map(response.Body)
			expectedMap := structs.Map(getAssetResponse)
			Expect(responseMap).To(Equal(expectedMap))
		})
		It("should create an asset with no columns specified and receive a valid response", func() {
			clearMockDataCatalog()
			// remove columns from createAssetRequest
			createAssetRequest.ResourceMetadata.Columns =
				createAssetRequest.ResourceMetadata.Columns[:0]
			getAssetResponse.ResourceMetadata.Columns =
				getAssetResponse.ResourceMetadata.Columns[:0]
			response, err = servicer.CreateAsset(ctx, TestConnectorCredentials, createAssetRequest)
			Expect(err).ToNot(HaveOccurred())
			assetID = fmt.Sprintf("%s.%s.%s.%s", TestDatabaseService, TestDatabase, TestBucket, TestObjectName)
			// check the asset ID in the response
			responseStr := string(mustAsJSON(response.Body))
			Expect(responseStr).To(Equal(fmt.Sprintf("{\"assetID\":%q}", assetID)))
		})
		It("should get the asset info", func() {
			response, err = servicer.GetAssetInfo(ctx, TestConnectorCredentials, &api.GetAssetRequest{AssetID: assetID, OperationType: "read"})
			Expect(err).ToNot(HaveOccurred())
			responseMap := structs.Map(response.Body)
			expectedMap := structs.Map(getAssetResponse)
			Expect(responseMap).To(Equal(expectedMap))
		})
	})
})
