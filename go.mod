module fybrik.io/openmetadata-connector

go 1.17

require (
	fybrik.io/fybrik v1.1.0
	fybrik.io/openmetadata-connector/datacatalog-go v0.0.0
	fybrik.io/openmetadata-connector/datacatalog-go-client v0.0.0
	fybrik.io/openmetadata-connector/datacatalog-go-models v0.0.0
	github.com/rs/zerolog v1.26.0
	github.com/spf13/cobra v1.5.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/fatih/structs v1.1.0
	github.com/hashicorp/go-retryablehttp v0.7.1
	github.com/onsi/ginkgo/v2 v2.1.4
	github.com/onsi/gomega v1.19.0
)

require (
	github.com/go-logr/logr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.23.5 // indirect
	k8s.io/apimachinery v0.23.5 // indirect
	k8s.io/klog/v2 v2.30.0 // indirect
	k8s.io/utils v0.0.0-20211116205334-6203023598ed // indirect
	sigs.k8s.io/json v0.0.0-20211020170558-c049b76a60c6 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
)

replace fybrik.io/openmetadata-connector/datacatalog-go => ./auto-generated/api

replace fybrik.io/openmetadata-connector/datacatalog-go-models => ./auto-generated/models

replace fybrik.io/openmetadata-connector/datacatalog-go-client => ./auto-generated/client
