package openapiconnectorcore

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"fybrik.io/fybrik/pkg/logging"

	models "fybrik.io/openmetadata-connector/datacatalog-go-models"
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

	destinationAssetID := "assetID"
	name := "Fake data"
	geography := "Hogwarts"
	dataFormat := "csv"
	createAssetRequest := models.CreateAssetRequest{
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
	response, err := servicer.CreateAsset(context.Background(), "fake-credentials", &createAssetRequest)
	if err != nil {
		t.Error(err)
	}

	responseStr := string(mustAsJSON(t, response.Body))
	if responseStr != "{\"assetID\":\"openmetadata-s3.default.fakeBucket.csvAsset\"}" {
		t.Error(errors.New("unexpected response from asset creation"))
	}
}
