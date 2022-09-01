package openapiconnectorcore

import (
	"io/ioutil"
	"os"
	"testing"

	"fybrik.io/fybrik/pkg/logging"

	"fybrik.io/openmetadata-connector/vault"
)

func setupSuite(t *testing.T) func() {
	logger = logging.LogInit(logging.CONNECTOR, "OpenMetadata Connector Tests")
	logger.Trace().Msg("Setting up OM Connector test suite")

	var err error
	jwtFile, err = ioutil.TempFile("/tmp", "jwt")
	if err != nil {
		logger.Fatal().Msg("failed to create temporary file")
	}
	n, err := jwtFile.WriteString(JWT)
	if err != nil || n != len(JWT) {
		logger.Fatal().Msg("failed to write to temporary file")
	}

	mockVaultServer = createMockVaultServer(t)
	mockOMServer = createMockOMServer(t)

	// populate vault configuration
	vaultConf = make(map[interface{}]interface{})
	vaultConf["address"] = mockVaultServer.URL
	vaultConf["jwt_file_path"] = jwtFile.Name()

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
	if err != nil || token != Token {
		t.Error(err)
	}

	secret, err := vaultClient.GetSecret(token, PathInVault)
	if err != nil {
		t.Error(err)
	}

	accessKey, secretKey, err := vaultClient.ExtractS3CredentialsFromSecret(secret)
	if err != nil || accessKey != AccessKey || secretKey != SecretKey {
		t.Error(err)
	}
}

func TestCreateAsset(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite()

	conf := make(map[interface{}]interface{})
	conf["openmetadata_endpoint"] = mockOMServer.URL
	conf["vault"] = vaultConf
	NewOpenMetadataAPIService(conf, &logger)
}
