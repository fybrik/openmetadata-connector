// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"fybrik.io/fybrik/pkg/logging"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	api "fybrik.io/openmetadata-connector/datacatalog-go/go"
	occ "fybrik.io/openmetadata-connector/pkg/openmetadata-connector-core"
)

const DefaultListeningPort = 8081
const DefaultConfigFile = "/etc/conf/conf.yaml"
const DefaultCustomizationFile = "./customization.yaml"

// RunCmd defines the command for running the connector
func RunCmd() *cobra.Command {
	logger := logging.LogInit(logging.CONNECTOR, "OpenMetadata Connector")
	configFile := DefaultConfigFile
	customizationFile := DefaultCustomizationFile
	port := DefaultListeningPort
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run the connector",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO - add logging level

			configFileBytes, err := os.ReadFile(configFile)
			if err != nil {
				logger.Error().Msg(fmt.Sprintf("failure to read config file: %v", err))
				return err
			}

			customizationFileBytes, err := os.ReadFile(customizationFile)
			if err != nil {
				logger.Error().Msg(fmt.Sprintf("failure to read customization file: %v", err))
				return err
			}

			conf := make(map[string]interface{})
			customization := make(map[string]interface{})

			err = yaml.Unmarshal(configFileBytes, &conf)
			if err != nil {
				logger.Error().Msg(fmt.Sprintf("failure to parse config file: %v", err))
				return err
			}
			err = yaml.Unmarshal(customizationFileBytes, &customization)
			if err != nil {
				logger.Error().Msg(fmt.Sprintf("failure to parse customization file: %v", err))
				return err
			}

			logger.Info().Msg("Server started")

			DefaultAPIService := occ.NewOpenMetadataAPIService(conf, customization, &logger)
			DefaultAPIController := occ.NewOpenMetadataAPIController(DefaultAPIService)

			router := api.NewRouter(DefaultAPIController)

			http.ListenAndServe(":"+strconv.Itoa(port), router) //nolint

			return nil
		},
	}

	cmd.Flags().StringVar(&configFile, "config", configFile, "Configuration file")
	cmd.Flags().IntVar(&port, "port", port, "Listening port")
	cmd.Flags().StringVar(&customizationFile, "customization", customizationFile,
		"File containing tags and custom properties needed for working with Fybrik")
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
	return cmd
}

func main() {
	// Run the cli
	if err := RootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
