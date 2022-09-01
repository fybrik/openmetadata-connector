# \ConfigApi

All URIs are relative to *http://localhost:8585/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAirflowConfiguration**](ConfigApi.md#GetAirflowConfiguration) | **Get** /v1/config/airflow | Get airflow configuration
[**GetAuthConfiguration**](ConfigApi.md#GetAuthConfiguration) | **Get** /v1/config/auth | Get auth configuration
[**GetAuthorizerConfig**](ConfigApi.md#GetAuthorizerConfig) | **Get** /v1/config/authorizer | Get authorizer configuration
[**GetJWKSResponse**](ConfigApi.md#GetJWKSResponse) | **Get** /v1/config/jwks | Get JWKS public key
[**GetSandboxConfiguration**](ConfigApi.md#GetSandboxConfiguration) | **Get** /v1/config/sandbox | Get sandbox mode



## GetAirflowConfiguration

> AirflowConfigurationForAPI GetAirflowConfiguration(ctx).Execute()

Get airflow configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.GetAirflowConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetAirflowConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAirflowConfiguration`: AirflowConfigurationForAPI
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetAirflowConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAirflowConfigurationRequest struct via the builder pattern


### Return type

[**AirflowConfigurationForAPI**](AirflowConfigurationForAPI.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthConfiguration

> AuthenticationConfiguration GetAuthConfiguration(ctx).Execute()

Get auth configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.GetAuthConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetAuthConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthConfiguration`: AuthenticationConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetAuthConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthConfigurationRequest struct via the builder pattern


### Return type

[**AuthenticationConfiguration**](AuthenticationConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorizerConfig

> AuthorizerConfiguration GetAuthorizerConfig(ctx).Execute()

Get authorizer configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.GetAuthorizerConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetAuthorizerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorizerConfig`: AuthorizerConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetAuthorizerConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizerConfigRequest struct via the builder pattern


### Return type

[**AuthorizerConfiguration**](AuthorizerConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJWKSResponse

> JWKSResponse GetJWKSResponse(ctx).Execute()

Get JWKS public key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.GetJWKSResponse(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetJWKSResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJWKSResponse`: JWKSResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetJWKSResponse`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetJWKSResponseRequest struct via the builder pattern


### Return type

[**JWKSResponse**](JWKSResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSandboxConfiguration

> SandboxConfiguration GetSandboxConfiguration(ctx).Execute()

Get sandbox mode

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigApi.GetSandboxConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigApi.GetSandboxConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSandboxConfiguration`: SandboxConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ConfigApi.GetSandboxConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSandboxConfigurationRequest struct via the builder pattern


### Return type

[**SandboxConfiguration**](SandboxConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

