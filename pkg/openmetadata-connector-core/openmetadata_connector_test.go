package openapiconnectorcore

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"fybrik.io/fybrik/pkg/logging"
	"github.com/fatih/structs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	api "fybrik.io/openmetadata-connector/datacatalog-go/go"
	vault "fybrik.io/openmetadata-connector/pkg/vault"
)

func TestOpenApiConnectorCore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "openApiConnectorCore Suite")
}

var _ = Describe("Vault and OM connector", func() {
	BeforeEach(func() {
		teardownSuite := setupSuite()
		DeferCleanup(func() {
			teardownSuite()
		})
	})
	Describe("Vault credentials retrieval flow", Ordered, func() {
		var token string
		var err error
		var secret []byte
		vaultClient := vault.NewVaultClient(vaultConf, &logger)
		It("should generate a vaild token", func() {
			token, err = vaultClient.GetToken()
			Expect(token).To(Equal(TestToken))
			Expect(err).ToNot(HaveOccurred())
		})
		It("should generate a vaild secret", func() {
			secret, err = vaultClient.GetSecret(token, TestPathInVault)
			Expect(err).ToNot(HaveOccurred())
		})
		It("should generate valid access and secret keys", func() {
			accessKey, secretKey, err := vaultClient.ExtractS3CredentialsFromSecret(secret)
			Expect(accessKey).To(Equal(TestAccessKey))
			Expect(secretKey).To(Equal(TestSecretKey))
			Expect(err).ToNot(HaveOccurred())
		})
	})
	Describe("Connector creating and getting assesets", Ordered, func() {
		conf := make(map[interface{}]interface{})
		conf["openmetadata_endpoint"] = mockOMServer.URL
		conf["vault"] = vaultConf
		servicer := NewOpenMetadataAPIService(conf, &logger)
		ctx := context.Background()
		createAssetRequest, getAssetResponse := getCreateAssetRequestAndReponse()
		response, err := servicer.CreateAsset(ctx, TestConnectorCredentials, createAssetRequest)
		Expect(err).ToNot(HaveOccurred())

		assetID := fmt.Sprintf("%s.%s.%s.%s", TestDatabaseService, TestDatabase, TestBucket, TestObjectName)
		// check the asset ID in the response
		responseStr := string(mustAsJSON(t, response.Body))

		if responseStr != fmt.Sprintf("{\"assetID\":%q}", assetID) {
			t.Error(errors.New("unexpected response from asset creation"))
		}

		response, err = servicer.GetAssetInfo(ctx, TestConnectorCredentials, &api.GetAssetRequest{AssetID: assetID, OperationType: "read"})
		if err != nil {
			t.Error(err)
		}

		responseMap := structs.Map(response.Body)
		expectedMap := structs.Map(getAssetResponse)

		if !reflect.DeepEqual(responseMap, expectedMap) {
			t.Error("Test failed. getAssetInfo response not as expected")
		}
	})
})

func setupSuite() func() {
	logger = logging.LogInit(logging.CONNECTOR, "OpenMetadata Connector Tests")
	logger.Trace().Msg("Setting up OM Connector test suite")

	var err error
	jwtFile, err = ioutil.TempFile("/tmp", "jwt")
	if err != nil {
		logger.Fatal().Msg("failed to create temporary file")
	}
	n, err := jwtFile.WriteString(TestJWT)
	if err != nil || n != len(TestJWT) {
		logger.Fatal().Msg("failed to write to temporary file")
	}

	mockVaultServer = createMockVaultServer()
	mockOMServer = createMockOMServer()

	// populate vault configuration
	vaultConf = make(map[interface{}]interface{})
	vaultConf[Address] = mockVaultServer.URL
	vaultConf[JwtFilePath] = jwtFile.Name()
	vaultConf[AuthPath] = TestAuthPath
	vaultConf[Role] = FybrikLowerCase

	// Return a function to teardown the test
	return func() {
		logger.Trace().Msg("Tearing down OM Connector test suite")
		os.Remove(jwtFile.Name())
	}
}

func TestVault(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite()

	vaultClient := vault.NewVaultClient(vaultConf, &logger)

	token, err := vaultClient.GetToken()
	if err != nil || token != TestToken {
		t.Error(err)
	}

	secret, err := vaultClient.GetSecret(token, TestPathInVault)
	if err != nil {
		t.Error(err)
	}

	accessKey, secretKey, err := vaultClient.ExtractS3CredentialsFromSecret(secret)
	if err != nil || accessKey != TestAccessKey || secretKey != TestSecretKey {
		t.Error(err)
	}
}

func TestCreateAndGetAsset(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite()

	conf := make(map[interface{}]interface{})
	conf["openmetadata_endpoint"] = mockOMServer.URL
	conf["vault"] = vaultConf
	servicer := NewOpenMetadataAPIService(conf, &logger)

	ctx := context.Background()
	createAssetRequest, getAssetResponse := getCreateAssetRequestAndReponse()
	response, err := servicer.CreateAsset(ctx, TestConnectorCredentials, createAssetRequest)
	if err != nil {
		t.Error(err)
	}

	assetID := fmt.Sprintf("%s.%s.%s.%s", TestDatabaseService, TestDatabase, TestBucket, TestObjectName)

	// check the asset ID in the response
	responseStr := string(mustAsJSON(t, response.Body))
	if responseStr != fmt.Sprintf("{\"assetID\":%q}", assetID) {
		t.Error(errors.New("unexpected response from asset creation"))
	}

	response, err = servicer.GetAssetInfo(ctx, TestConnectorCredentials, &api.GetAssetRequest{AssetID: assetID, OperationType: "read"})
	if err != nil {
		t.Error(err)
	}

	responseMap := structs.Map(response.Body)
	expectedMap := structs.Map(getAssetResponse)

	if !reflect.DeepEqual(responseMap, expectedMap) {
		t.Error("Test failed. getAssetInfo response not as expected")
	}
}
