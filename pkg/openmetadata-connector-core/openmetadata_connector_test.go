package openapiconnectorcore

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"fybrik.io/fybrik/pkg/logging"

	api "fybrik.io/openmetadata-connector/datacatalog-go/go"
	"fybrik.io/openmetadata-connector/pkg/vault"
)

func setupSuite(t *testing.T) func() {
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

	mockVaultServer = createMockVaultServer(t)
	mockOMServer = createMockOMServer(t)

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
	createAssetRequest := getCreateAssetRequest()
	response, err := servicer.CreateAsset(ctx, "fake-credentials", createAssetRequest)
	if err != nil {
		t.Error(err)
	}

	assetID := fmt.Sprintf("%s.%s.%s.%s", TestDatabaseService, TestDatabase, TestBucket, TestObjectName)

	// check the asset ID in the response
	responseStr := string(mustAsJSON(t, response.Body))
	if responseStr != fmt.Sprintf("{\"assetID\":%q}", assetID) {
		t.Error(errors.New("unexpected response from asset creation"))
	}

	response, err = servicer.GetAssetInfo(ctx, "fake-credentials", &api.GetAssetRequest{AssetID: assetID, OperationType: "read"})
	if err != nil {
		t.Error(err)
	}
}
