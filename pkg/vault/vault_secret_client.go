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
	address      string
	authPath     string
	pluginPrefix string
	role         string
	jwtFilePath  string
	logger       *zerolog.Logger
	httpClient   *http.Client
}

const Data = "data"
const EmptyString = ""
const JWT = "jwt"
const ROLE = "role"

const FailedToExtractCredentialsFromVaultSecret = "failed to extract credentials from Vault secret"
const GetSecretFailed = "getSecret failed"
const GetTokenFailed = "getToken failed"
const MalformedSecretResponseFromVault = "malformed secret response from vault"

func getFullAuthPath(authPath string) string {
	if authPath == EmptyString {
		return EmptyString
	}
	fullAuthPath := fmt.Sprintf("/v1/auth/%s/login", authPath)
	return fullAuthPath
}

// return structure for Vault Client, based on configuration
func NewVaultClient(conf map[interface{}]interface{}, logger *zerolog.Logger, client *http.Client) *VaultClient {
	var address, authPath, pluginPrefix, role, jwtFilePath string
	if conf != nil {
		if addressConf, ok := conf["address"]; ok {
			address = addressConf.(string)
		}
		if authPathConf, ok := conf["authPath"]; ok {
			authPath = authPathConf.(string)
		}
		if pluginPrefixConf, ok := conf["pluginPrefix"]; ok {
			pluginPrefix = pluginPrefixConf.(string)
		}
		if roleConf, ok := conf["role"]; ok {
			role = roleConf.(string)
		}
		if jwtFilePathConf, ok := conf["jwt_file_path"]; ok {
			jwtFilePath = jwtFilePathConf.(string)
		}
	}

	return &VaultClient{address: address, authPath: getFullAuthPath(authPath),
		pluginPrefix: pluginPrefix, role: role, jwtFilePath: jwtFilePath, logger: logger, httpClient: client}
}

func GetFullSecretPath(pluginPrefix, secret string) string {
	return fmt.Sprintf("/v1/%s/%s", pluginPrefix, secret)
}

func (v *VaultClient) getToken() (string, error) {
	// read JWT file. JWT is used to authenticate against Vault
	jwt, err := os.ReadFile(v.jwtFilePath)
	if err != nil {
		v.logger.Error().Msg("Failed to read JWT file")
		return EmptyString, err
	}

	j := make(map[string]string)
	j[JWT] = string(jwt)
	j[ROLE] = v.role

	fullAuthPath := v.address + v.authPath
	jsonStr, err := json.Marshal(j)
	if err != nil {
		v.logger.Error().Msg("Failed to transform map to JSON string")
		return EmptyString, err
	}

	// request token from vault
	requestBody := strings.NewReader(string(jsonStr))
	resp, err := v.httpClient.Post(fullAuthPath, "encoding/json", requestBody) //nolint
	if err != nil {
		v.logger.Error().Msg("vault POST request failed")
		return EmptyString, err
	}
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		v.logger.Error().Msg("Failed to get token from vault")
		return EmptyString, err
	}

	// parse response
	responseMap := make(map[string]interface{})
	err = json.Unmarshal(responseBody, &responseMap)
	if err != nil {
		v.logger.Error().Msg("malformed response from vault")
		return EmptyString, err
	}

	var token string
	if value, ok := responseMap["auth"]; ok {
		token = value.(map[string]interface{})["client_token"].(string)
		v.logger.Info().Msg("Successfully obtained token from Vault")
		return token, nil
	}
	const MalformedVaultResposeMessage = "malformed response from vault"
	v.logger.Error().Msg(MalformedVaultResposeMessage)
	return EmptyString, errors.New(MalformedVaultResposeMessage)
}

func (v *VaultClient) getSecret(token, secretPath string) ([]byte, error) {
	req, err := http.NewRequestWithContext(context.Background(), "GET", v.address+secretPath, http.NoBody)
	if err != nil {
		v.logger.Error().Msg("Failed to prepare Vault secret request")
		return nil, err
	}

	req.Header.Set("X-Vault-Token", token)
	resp, err := v.httpClient.Do(req)
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

func (v *VaultClient) GetSecretMap(credentialsPath *string) (map[string]interface{}, error) {
	// if we fail, we log Warn messages instead of Error messages, for scenarios where the connector
	// is given all necessary credentials, and does not need credentials hidden in a secret.
	token, err := v.getToken()
	if err != nil {
		v.logger.Warn().Msg(GetTokenFailed)
		return nil, errors.New(GetTokenFailed)
	}
	secret, err := v.getSecret(token, *credentialsPath)
	if err != nil {
		v.logger.Warn().Msg(GetSecretFailed)
		return nil, errors.New(GetSecretFailed)
	}

	secretMap := make(map[string]interface{})
	err = json.Unmarshal(secret, &secretMap)
	if err != nil {
		v.logger.Warn().Err(err).Msg(MalformedSecretResponseFromVault)
	}
	if value, ok := secretMap[Data]; ok {
		data := value.(map[string]interface{})
		v.logger.Info().Msg("Successfully extracted credentials from Vault secret")
		return data, nil
	}
	v.logger.Warn().Msg(FailedToExtractCredentialsFromVaultSecret)
	return nil, errors.New(FailedToExtractCredentialsFromVaultSecret)
}
