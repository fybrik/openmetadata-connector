// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	fybrikEnv "fybrik.io/fybrik/pkg/environment"
	"fybrik.io/fybrik/pkg/logging"
	fybrikTLS "fybrik.io/fybrik/pkg/tls"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	api "fybrik.io/openmetadata-connector/datacatalog-go/go"
	occ "fybrik.io/openmetadata-connector/pkg/openmetadata-connector-core"
	"fybrik.io/openmetadata-connector/pkg/utils"
)

const (
	DefaultConfigFile        = "/etc/conf/conf.yaml"
	DefaultCustomizationFile = "./customization.yaml"
)

const (
	Customization                         = "customization"
	FailureToParseCustomizationFile       = "failure to parse customization file"
	FailureToReadConfigFile               = "failure to read configuration file"
	FailureToReadCustomizationFile        = "failure to read customization file"
	FileContainingTagsAndPropertiesNeeded = "File containing tags and custom properties needed for working with Fybrik"
	ParseCustomizationFileFailed          = "parseCustomizationFile() failed. Exiting"
)

func parseCustomizationFile(customizationFile string, logger *zerolog.Logger) (map[string]interface{}, error) {
	customizationFileBytes, err := os.ReadFile(customizationFile)
	if err != nil {
		logger.Error().Err(err).Msg(FailureToReadCustomizationFile)
		return nil, err
	}

	customization := make(map[string]interface{})

	err = yaml.Unmarshal(customizationFileBytes, &customization)
	if err != nil {
		logger.Error().Err(err).Msg(FailureToParseCustomizationFile)
		return nil, fmt.Errorf(FailureToParseCustomizationFile)
	}

	return customization, nil
}

// RunCmd defines the command for running the connector
func RunCmd() *cobra.Command {
	logger := logging.LogInit(logging.CONNECTOR, "OpenMetadata Connector")
	configFile := DefaultConfigFile
	customizationFile := DefaultCustomizationFile
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run the connector",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO - add logging level and pretty logging

			customization, err := parseCustomizationFile(customizationFile, &logger)
			if err != nil {
				logger.Error().Msg(ParseCustomizationFileFailed)
				return
			}

			configFileBytes, err := os.ReadFile(configFile)
			if err != nil {
				logger.Error().Err(err).Msg(FailureToReadConfigFile)
				return
			}

			conf := make(map[string]interface{})
			err = yaml.Unmarshal(configFileBytes, &conf)
			if err != nil {
				logger.Error().Err(err).Msg("failure to parse config file")
				return
			}

			DefaultAPIService := occ.NewOpenMetadataAPIService(conf, customization, &logger)
			DefaultAPIController := occ.NewOpenMetadataAPIController(DefaultAPIService)

			// Init the http client which is used to communicate with Vault and Openmetadata servers.
			utils.InitHTTPClient(&logger)

			router := api.NewRouter(DefaultAPIController)
			addr := ":" + strconv.Itoa(DefaultAPIService.Port)
			if fybrikEnv.IsUsingTLS() {
				tlsConfig, err := fybrikTLS.GetServerConfig(&logger)
				if err != nil {
					logger.Error().Msg("failed to get tls config")
					return
				}
				server := http.Server{Addr: addr, Handler: router, TLSConfig: tlsConfig, ReadHeaderTimeout: occ.ReadHeaderTimeout}
				logger.Info().Msg(occ.ServerStartingMsg)
				err = server.ListenAndServeTLS("", "")
				if err != nil {
					logger.Error().Err(err).Msg("function ListenAndServeTLS returns error")
				}
			} else {
				server := http.Server{Addr: addr, Handler: router, ReadHeaderTimeout: occ.ReadHeaderTimeout}
				logger.Info().Msg(occ.ServerStartingMsg)
				err := server.ListenAndServe()
				if err != nil {
					logger.Error().Err(err).Msg("function ListenAndServe returns error")
				}
			}
		},
	}

	cmd.Flags().StringVar(&configFile, "config", configFile, "Configuration file")
	cmd.Flags().StringVar(&customizationFile, Customization, customizationFile,
		FileContainingTagsAndPropertiesNeeded)
	cmd.CompletionOptions.DisableDefaultCmd = true
	return cmd
}

// PrepareCmd defines the command for preparing OM for Fybrik
// by creating tags and custom properties for Fybrik
func PrepareCmd() *cobra.Command {
	logger := logging.LogInit(logging.CONNECTOR, "Preparing OpenMetadata for Fybrik")
	customizationFile := DefaultCustomizationFile
	cmd := &cobra.Command{
		Use:   "prepare",
		Short: "Prepare OpenMetadata for Fybrik",
		Run: func(cmd *cobra.Command, args []string) {
			customization, err := parseCustomizationFile(customizationFile, &logger)
			if err != nil {
				logger.Error().Msg(ParseCustomizationFileFailed)
				return
			}

			ok, endpoint, user, password := utils.GetEnvironmentVariables()
			if !ok {
				logger.Error().Msg("failed to get environment variables. cannot proceed")
				return
			}
			occ.PrepareOpenMetadataForFybrik(endpoint, user, password, customization, &logger)
		},
	}
	cmd.Flags().StringVar(&customizationFile, Customization, customizationFile, FileContainingTagsAndPropertiesNeeded)
	cmd.CompletionOptions.DisableDefaultCmd = true
	return cmd
}

// RootCmd defines the root cli command
func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "openmetadata-connector",
		Short: "Kubernetes based OpenMetadata data catalog connector for Fybrik",
	}
	cmd.AddCommand(RunCmd())
	cmd.AddCommand(PrepareCmd())
	return cmd
}

func main() {
	// Run the cli
	if err := RootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
