# \TokenEndpointApi

All URIs are relative to *https://api.authlete.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthTokenApi**](TokenEndpointApi.md#AuthTokenApi) | **Post** /auth/token | /auth/token API
[**AuthTokenFailApi**](TokenEndpointApi.md#AuthTokenFailApi) | **Post** /auth/token/fail | /auth/token/fail API
[**AuthTokenIssueApi**](TokenEndpointApi.md#AuthTokenIssueApi) | **Post** /auth/token/issue | /auth/token/issue API



## AuthTokenApi

> TokenResponse AuthTokenApi(ctx).TokenRequest(tokenRequest).Execute()

/auth/token API



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
    tokenRequest := *openapiclient.NewTokenRequest("Parameters_example") // TokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenEndpointApi.AuthTokenApi(context.Background()).TokenRequest(tokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenEndpointApi.AuthTokenApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenApi`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenEndpointApi.AuthTokenApi`: %v\n", resp)
}
```

### Path Parameters



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

> TokenFailResponse AuthTokenFailApi(ctx).TokenFailRequest(tokenFailRequest).Execute()

/auth/token/fail API



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
    tokenFailRequest := *openapiclient.NewTokenFailRequest("Ticket_example", "Reason_example") // TokenFailRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenEndpointApi.AuthTokenFailApi(context.Background()).TokenFailRequest(tokenFailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenEndpointApi.AuthTokenFailApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenFailApi`: TokenFailResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenEndpointApi.AuthTokenFailApi`: %v\n", resp)
}
```

### Path Parameters



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

> TokenIssueResponse AuthTokenIssueApi(ctx).TokenIssueRequest(tokenIssueRequest).Execute()

/auth/token/issue API



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
    tokenIssueRequest := *openapiclient.NewTokenIssueRequest("Ticket_example", "Subject_example") // TokenIssueRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenEndpointApi.AuthTokenIssueApi(context.Background()).TokenIssueRequest(tokenIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenEndpointApi.AuthTokenIssueApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenIssueApi`: TokenIssueResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenEndpointApi.AuthTokenIssueApi`: %v\n", resp)
}
```

### Path Parameters



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

