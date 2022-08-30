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

const JWT = "myJWT"
const PathInVault = "/v1/kubernetes/info?namespace=default"
const AccessKey = "myAccessKey"
const SecretKey = "mySecretKey"
const Token = "vaultToken"
const ContentType = "Content-Type"
const ApplicationJSON = "application/json"

var logger zerolog.Logger
var jwtFile *os.File
var mockVaultServer *httptest.Server

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
				t.Error(errors.New("unexpected request format"))
				return
			}

			response := map[string]interface{}{
				"request_id": "00000000-0000-0000-0000-000000000000",
				"auth": map[string]interface{}{
					"client_token": Token,
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

			w.Header().Add(ContentType, ApplicationJSON)
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
