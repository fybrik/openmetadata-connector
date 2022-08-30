package openapiconnectorcore

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"fybrik.io/fybrik/pkg/logging"
	"github.com/rs/zerolog"

	"fybrik.io/openmetadata-connector/vault"
)

const Token = "vaultToken"
const JWT = "myJWT"
const PathInVault = "/v1/kubernetes/info?namespace=default"
const AccessKey = "myAccessKey"
const SecretKey = "mySecretKey"

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

		if r.Method == "POST" && r.RequestURI == "/v1/auth/kubernetes/login" {
			requestMap := make(map[string]interface{})
			err = json.Unmarshal(requestBytes, &requestMap)
			if err != nil {
				t.Error(err)
				return
			}
			if (requestMap["role"].(string) != "fybrik") ||
				(requestMap["jwt"].(string) != JWT) {
				t.Error(errors.New("Unexpected request format"))
				return
			}

			response := map[string]interface{}{
				"request_id": "00000000-0000-0000-0000-000000000000",
				"auth": map[string]interface{}{
					"client_token": Token,
				},
			}

			responseBytes := mustAsJSON(t, response)

			w.Header().Add("Content-Type", "application/json")
			_, err = w.Write(responseBytes)
			if err != nil {
				t.Error(err)
			}

			return
		}

		if r.Method == "GET" {
			if r.Header.Get("X-Vault-Token") != Token {
				t.Error(err)
			}
			logger.Trace().Msg("About to mock response secret request")
			response := map[string]interface{}{
				"data": map[string]interface{}{
					"access_key": AccessKey,
					"secret_key": SecretKey,
				},
			}

			responseBytes := mustAsJSON(t, response)

			w.Header().Add("Content-Type", "application/json")
			_, err = w.Write(responseBytes)
			if err != nil {
				t.Error(err)
			}

			return
		}
		t.Error(errors.New(""))
	}))

	return svr
}

var logger zerolog.Logger
var jwtFile *os.File

func setupSuite() func() {
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

	// Return a function to teardown the test
	return func() {
		logger.Trace().Msg("Tearing down OM Connector test suite")
		os.Remove(jwtFile.Name())
	}
}

func TestVault(t *testing.T) {
	teardownSuite := setupSuite()
	defer teardownSuite()

	vaultServer := createMockVaultServer(t)
	conf := make(map[interface{}]interface{})
	conf["address"] = vaultServer.URL
	conf["jwt_file_path"] = jwtFile.Name()
	vaultClient := vault.NewVaultClient(conf, &logger)

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
