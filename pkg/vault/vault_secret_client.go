// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package vault

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type VaultClient struct {
	address     string
	authPath    string
	role        string
	jwtFilePath string
	logger      *zerolog.Logger
}

const ROLE = "role"
const JWT = "jwt"

func getFullAuthPath(authPath string) string {
	if authPath == "" {
		return ""
	}
	fullAuthPath := fmt.Sprintf("/v1/auth/%s/login", authPath)
	return fullAuthPath
}

// return structure for Vault Client, based on configuration
func NewVaultClient(conf map[interface{}]interface{}, logger *zerolog.Logger) VaultClient {
	var address, authPath, role, jwtFilePath string
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
			jwtFilePath = jwtFilePathConf.(string)
		}
	}
	return VaultClient{address: address, authPath: getFullAuthPath(authPath),
		role: role, jwtFilePath: jwtFilePath, logger: logger}
}

func (v *VaultClient) GetToken() (string, error) {
	// read JWT file. JWT is used to authenticate against Vault
	jwt, err := os.ReadFile(v.jwtFilePath)
	if err != nil {
		v.logger.Error().Msg("Failed to read JWT file")
		return "", err
	}

	j := make(map[string]string)
	j[JWT] = string(jwt)
	j[ROLE] = v.role

	fullAuthPath := v.address + v.authPath
	jsonStr, err := json.Marshal(j)
	if err != nil {
		v.logger.Error().Msg("Failed to transform map to JSON string")
		return "", err
	}

	// request token from vault
	requestBody := strings.NewReader(string(jsonStr))
	resp, err := http.Post(fullAuthPath, "encoding/json", requestBody) //nolint
	if err != nil {
		v.logger.Error().Msg("vault POST request failed")
		return "", err
	}
	responseBody, err := io.ReadAll(resp.Body)
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
	}
	const MalformedVaultResposeMessage = "malformed response from vault"
	v.logger.Error().Msg(MalformedVaultResposeMessage)
	return "", errors.New(MalformedVaultResposeMessage)
}

func (v *VaultClient) GetSecret(token, secretPath string) ([]byte, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", v.address+secretPath, http.NoBody)
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
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		v.logger.Error().Msg("Failed to read Vault secret")
		return nil, err
	}

	v.logger.Info().Msg("Successfully read Vault secret")
	return responseBody, nil
}

const MalformedSecretResponseFromVault = "malformed secret response from vault"
const FailedToExtractS3CredentialsFromVaultSecret = "failed to extract S3 credentials from Vault secret"

func (v *VaultClient) ExtractS3CredentialsFromSecret(secret []byte) (string, string, error) {
	secretMap := make(map[string]interface{})
	err := json.Unmarshal(secret, &secretMap)
	if err != nil {
		v.logger.Error().Msg(MalformedSecretResponseFromVault)
		return "", "", errors.New(MalformedSecretResponseFromVault)
	}

	if value, ok := secretMap["data"]; ok {
		data := value.(map[string]interface{})
		v.logger.Info().Msg("Successfully extracted S3 credentials from Vault secret")
		return data["access_key"].(string), data["secret_key"].(string), nil
	}
	v.logger.Error().Msg(FailedToExtractS3CredentialsFromVaultSecret)
	return "", "", errors.New(FailedToExtractS3CredentialsFromVaultSecret)
}
