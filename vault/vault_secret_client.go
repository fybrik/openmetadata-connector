package vault

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type VaultClient struct {
	address       string
	authPath      string
	role          string
	jwt_file_path string
	logger        zerolog.Logger
}

// return structure for Vault Client, based on configuration
func NewVaultClient(conf map[interface{}]interface{}, logger zerolog.Logger) VaultClient {
	address := "http://vault.fybrik-system:8200"
	authPath := "/v1/auth/kubernetes/login"
	role := "fybrik"
	jwt_file_path := "/var/run/secrets/kubernetes.io/serviceaccount/token"
	if conf != nil {
		if addressConf, ok := conf["address"]; ok {
			address = addressConf.(string)
		}
		if authPathConf, ok := conf["authPath"]; ok {
			authPath = authPathConf.(string)
		}
		if roleConf, ok := conf["role"]; ok {
			role = roleConf.(string)
		}
		if jwtFilePathConf, ok := conf["jwt_file_path"]; ok {
			jwt_file_path = jwtFilePathConf.(string)
		}
	}
	return VaultClient{address: address, authPath: authPath, role: role, jwt_file_path: jwt_file_path, logger: logger}
}

func (v *VaultClient) GetToken() (string, error) {
	// read JWT file. JWT is used to authenticate against Vault
	jwt, err := os.ReadFile(v.jwt_file_path)
	if err != nil {
		v.logger.Error().Msg("Failed to read JWT file")
		return "", err
	}

	j := make(map[string]string)
	j["jwt"] = string(jwt)
	j["role"] = v.role

	full_auth_path := v.address + v.authPath
	jsonStr, err := json.Marshal(j)
	if err != nil {
		v.logger.Error().Msg("Failed to transform map to JSON string")
		return "", err
	}

	// request token from vault
	requestBody := strings.NewReader(string(jsonStr))
	resp, _ := http.Post(full_auth_path, "encoding/json", requestBody)
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		v.logger.Error().Msg("Failed to get token from vault")
		return "", err
	}

	// parse response
	responseMap := make(map[string]interface{})
	err = json.Unmarshal(responseBody, &responseMap)
	if err != nil {
		v.logger.Error().Msg("malformed response from vault")
		return "", err
	}

	var token string
	if value, ok := responseMap["auth"]; ok {
		token = value.(map[string]interface{})["client_token"].(string)
		v.logger.Info().Msg("Successfully obtained token from Vault")
		return token, nil
	} else {
		v.logger.Error().Msg("malformed response from vault")
		return "", errors.New("malformed response from vault")
	}
}

func (v *VaultClient) GetSecret(token string, secretPath string) ([]byte, error) {
	req, err := http.NewRequest("GET", v.address+secretPath, nil)
	if err != nil {
		v.logger.Error().Msg("Failed to prepare Vault secret request")
		return nil, err
	}

	req.Header.Set("X-Vault-Token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		v.logger.Error().Msg("Failed to obtain secret from Vault")
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		v.logger.Error().Msg("Failed to read Vault secret")
		return nil, err
	}

	v.logger.Info().Msg("Successfully read Vault secret")
	return responseBody, nil
}

func (v *VaultClient) ExtractS3CredentialsFromSecret(secret []byte) (string, string) {
	secretMap := make(map[string]interface{})
	err := json.Unmarshal(secret, &secretMap)
	if err != nil {
		v.logger.Error().Msg("malformed secret response from vault")
		return "", ""
	}

	if value, ok := secretMap["data"]; ok {
		data := value.(map[string]interface{})
		v.logger.Info().Msg("Successfully extracted S3 credentials from Vault secret")
		return data["access_key"].(string), data["secret_key"].(string)
	} else {
		v.logger.Error().Msg("Failed to extract S3 credentials from Vault secret")
		return "", ""
	}
}
