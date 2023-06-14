# \TokenEndpointApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthTokenApi**](TokenEndpointApi.md#AuthTokenApi) | **Post** /api/{serviceApiKey}/auth/token | /api/auth/token API
[**AuthTokenFailApi**](TokenEndpointApi.md#AuthTokenFailApi) | **Post** /api/{serviceApiKey}/auth/token/fail | /api/auth/token/fail API
[**AuthTokenIssueApi**](TokenEndpointApi.md#AuthTokenIssueApi) | **Post** /api/{serviceApiKey}/auth/token/issue | /api/auth/token/issue API



## AuthTokenApi

> TokenResponse AuthTokenApi(ctx, serviceApiKey).TokenRequest(tokenRequest).Execute()

/api/auth/token API



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
    serviceApiKey := "serviceApiKey_example" // string | serviceApiKey
    tokenRequest := *openapiclient.NewTokenRequest("Parameters_example") // TokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenEndpointApi.AuthTokenApi(context.Background(), serviceApiKey).TokenRequest(tokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenEndpointApi.AuthTokenApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenApi`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenEndpointApi.AuthTokenApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceApiKey** | **string** | serviceApiKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenFailApi

> TokenFailResponse AuthTokenFailApi(ctx, serviceApiKey).TokenFailRequest(tokenFailRequest).Execute()

/api/auth/token/fail API



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
    serviceApiKey := "serviceApiKey_example" // string | serviceApiKey
    tokenFailRequest := *openapiclient.NewTokenFailRequest("Ticket_example", "Reason_example") // TokenFailRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenEndpointApi.AuthTokenFailApi(context.Background(), serviceApiKey).TokenFailRequest(tokenFailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenEndpointApi.AuthTokenFailApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenFailApi`: TokenFailResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenEndpointApi.AuthTokenFailApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceApiKey** | **string** | serviceApiKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenFailApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenFailRequest** | [**TokenFailRequest**](TokenFailRequest.md) |  | 

### Return type

[**TokenFailResponse**](TokenFailResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenIssueApi

> TokenIssueResponse AuthTokenIssueApi(ctx, serviceApiKey).TokenIssueRequest(tokenIssueRequest).Execute()

/api/auth/token/issue API



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
    serviceApiKey := "serviceApiKey_example" // string | serviceApiKey
    tokenIssueRequest := *openapiclient.NewTokenIssueRequest("Ticket_example", "Subject_example") // TokenIssueRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenEndpointApi.AuthTokenIssueApi(context.Background(), serviceApiKey).TokenIssueRequest(tokenIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenEndpointApi.AuthTokenIssueApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenIssueApi`: TokenIssueResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenEndpointApi.AuthTokenIssueApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceApiKey** | **string** | serviceApiKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenIssueApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenIssueRequest** | [**TokenIssueRequest**](TokenIssueRequest.md) |  | 

### Return type

[**TokenIssueResponse**](TokenIssueResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

