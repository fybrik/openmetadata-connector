package openapiconnectorcore

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"fybrik.io/fybrik/pkg/logging"

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
	vaultConf[AuthPath] = Kubernetes
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

func TestCreateAsset(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite()

	conf := make(map[interface{}]interface{})
	conf["openmetadata_endpoint"] = mockOMServer.URL
	conf["vault"] = vaultConf
	servicer := NewOpenMetadataAPIService(conf, &logger)

	createAssetRequest := getCreateAssetRequest()
	response, err := servicer.CreateAsset(context.Background(), "fake-credentials", createAssetRequest)
	if err != nil {
		t.Error(err)
	}

	responseStr := string(mustAsJSON(t, response.Body))
	if responseStr != "{\"assetID\":\"openmetadata-s3.default.fakeBucket.csvAsset\"}" {
		t.Error(errors.New("unexpected response from asset creation"))
	}
}
