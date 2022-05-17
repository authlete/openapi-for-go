# \AuthorizationEndpointApi

All URIs are relative to *https://api.authlete.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthAuthorizationApi**](AuthorizationEndpointApi.md#AuthAuthorizationApi) | **Post** /auth/authorization | /auth/authorization API
[**AuthAuthorizationFailApi**](AuthorizationEndpointApi.md#AuthAuthorizationFailApi) | **Post** /auth/authorization/fail | /auth/authorization/fail API
[**AuthAuthorizationIssueApi**](AuthorizationEndpointApi.md#AuthAuthorizationIssueApi) | **Post** /auth/authorization/issue | /auth/authorization/issue API



## AuthAuthorizationApi

> AuthorizationResponse AuthAuthorizationApi(ctx).AuthorizationRequest(authorizationRequest).Execute()

/auth/authorization API



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
    authorizationRequest := *openapiclient.NewAuthorizationRequest("Parameters_example") // AuthorizationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationEndpointApi.AuthAuthorizationApi(context.Background()).AuthorizationRequest(authorizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationEndpointApi.AuthAuthorizationApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthAuthorizationApi`: AuthorizationResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationEndpointApi.AuthAuthorizationApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthAuthorizationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationRequest** | [**AuthorizationRequest**](AuthorizationRequest.md) |  | 

### Return type

[**AuthorizationResponse**](AuthorizationResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthAuthorizationFailApi

> AuthorizationFailResponse AuthAuthorizationFailApi(ctx).AuthorizationFailRequest(authorizationFailRequest).Execute()

/auth/authorization/fail API



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
    authorizationFailRequest := *openapiclient.NewAuthorizationFailRequest("Ticket_example", "Reason_example") // AuthorizationFailRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationEndpointApi.AuthAuthorizationFailApi(context.Background()).AuthorizationFailRequest(authorizationFailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationEndpointApi.AuthAuthorizationFailApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthAuthorizationFailApi`: AuthorizationFailResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationEndpointApi.AuthAuthorizationFailApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthAuthorizationFailApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationFailRequest** | [**AuthorizationFailRequest**](AuthorizationFailRequest.md) |  | 

### Return type

[**AuthorizationFailResponse**](AuthorizationFailResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthAuthorizationIssueApi

> AuthorizationIssueResponse AuthAuthorizationIssueApi(ctx).AuthorizationIssueRequest(authorizationIssueRequest).Execute()

/auth/authorization/issue API



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
    authorizationIssueRequest := *openapiclient.NewAuthorizationIssueRequest("Ticket_example", "Subject_example") // AuthorizationIssueRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationEndpointApi.AuthAuthorizationIssueApi(context.Background()).AuthorizationIssueRequest(authorizationIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationEndpointApi.AuthAuthorizationIssueApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthAuthorizationIssueApi`: AuthorizationIssueResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationEndpointApi.AuthAuthorizationIssueApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthAuthorizationIssueApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationIssueRequest** | [**AuthorizationIssueRequest**](AuthorizationIssueRequest.md) |  | 

### Return type

[**AuthorizationIssueResponse**](AuthorizationIssueResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

