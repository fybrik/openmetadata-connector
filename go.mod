module fybrik.io/openmetadata-connector

go 1.17

require (
	fybrik.io/fybrik v1.0.0
	fybrik.io/openmetadata-connector/datacatalog-go v0.0.0
	fybrik.io/openmetadata-connector/datacatalog-go-client v0.0.0
	fybrik.io/openmetadata-connector/datacatalog-go-models v0.0.0
	github.com/rs/zerolog v1.26.0
	github.com/spf13/cobra v1.5.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/fatih/structs v1.1.0
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace fybrik.io/openmetadata-connector/datacatalog-go => ./auto-generated/api

replace fybrik.io/openmetadata-connector/datacatalog-go-models => ./auto-generated/models

replace fybrik.io/openmetadata-connector/datacatalog-go-client => ./auto-generated/client
