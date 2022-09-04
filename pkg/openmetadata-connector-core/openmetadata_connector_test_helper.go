package openapiconnectorcore

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	zerolog "github.com/rs/zerolog"
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
var vaultConf map[interface{}]interface{}

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
