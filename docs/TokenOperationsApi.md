# \TokenOperationsApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthTokenCreateApi**](TokenOperationsApi.md#AuthTokenCreateApi) | **Post** /api/{serviceApiKey}/auth/token/create | /api/auth/token/create API
[**AuthTokenDeleteApi**](TokenOperationsApi.md#AuthTokenDeleteApi) | **Delete** /api/{serviceApiKey}/auth/token/delete/{accessTokenIdentifier} | /api/auth/token/delete API
[**AuthTokenGetListApi**](TokenOperationsApi.md#AuthTokenGetListApi) | **Get** /api/{serviceApiKey}/auth/token/get/list | /api/auth/token/get/list API
[**AuthTokenRevokeApi**](TokenOperationsApi.md#AuthTokenRevokeApi) | **Post** /api/{serviceApiKey}/auth/token/revoke | /api/auth/token/revoke API
[**AuthTokenUpdateApi**](TokenOperationsApi.md#AuthTokenUpdateApi) | **Post** /api/{serviceApiKey}/auth/token/update | /api/auth/token/update API



## AuthTokenCreateApi

> TokenCreateResponse AuthTokenCreateApi(ctx, serviceApiKey).TokenCreateRequest(tokenCreateRequest).Execute()

/api/auth/token/create API



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
    tokenCreateRequest := *openapiclient.NewTokenCreateRequest(openapiclient.grant_type("AUTHORIZATION_CODE"), int64(123)) // TokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenOperationsApi.AuthTokenCreateApi(context.Background(), serviceApiKey).TokenCreateRequest(tokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenOperationsApi.AuthTokenCreateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenCreateApi`: TokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenOperationsApi.AuthTokenCreateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceApiKey** | **string** | serviceApiKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenCreateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenCreateRequest** | [**TokenCreateRequest**](TokenCreateRequest.md) |  | 

### Return type

[**TokenCreateResponse**](TokenCreateResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenDeleteApi

> AuthTokenDeleteApi(ctx, serviceApiKey, accessTokenIdentifier).Execute()

/api/auth/token/delete API



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
    accessTokenIdentifier := "accessTokenIdentifier_example" // string | The identifier of an existing access token. The identifier is the value of the access token or the value of the hash of the access token. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenOperationsApi.AuthTokenDeleteApi(context.Background(), serviceApiKey, accessTokenIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenOperationsApi.AuthTokenDeleteApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceApiKey** | **string** | serviceApiKey | 
**accessTokenIdentifier** | **string** | The identifier of an existing access token. The identifier is the value of the access token or the value of the hash of the access token.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenDeleteApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenGetListApi

> TokenGetListResponse AuthTokenGetListApi(ctx, serviceApiKey).Subject(subject).ClientIdentifier(clientIdentifier).Start(start).End(end).Execute()

/api/auth/token/get/list API



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
    subject := "subject_example" // string | Unique user ID. 
    clientIdentifier := "clientIdentifier_example" // string | Client Identifier (client ID or client ID alias).  (optional)
    start := int32(56) // int32 | Start index of search results (inclusive). The default value is 0. (optional)
    end := int32(56) // int32 | End index of search results (exclusive). The default value is 5.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenOperationsApi.AuthTokenGetListApi(context.Background(), serviceApiKey).Subject(subject).ClientIdentifier(clientIdentifier).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenOperationsApi.AuthTokenGetListApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenGetListApi`: TokenGetListResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenOperationsApi.AuthTokenGetListApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceApiKey** | **string** | serviceApiKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenGetListApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subject** | **string** | Unique user ID.  | 
 **clientIdentifier** | **string** | Client Identifier (client ID or client ID alias).  | 
 **start** | **int32** | Start index of search results (inclusive). The default value is 0. | 
 **end** | **int32** | End index of search results (exclusive). The default value is 5.  | 

### Return type

[**TokenGetListResponse**](TokenGetListResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenRevokeApi

> TokenRevokeResponse AuthTokenRevokeApi(ctx, serviceApiKey).TokenRevokeRequest(tokenRevokeRequest).Execute()

/api/auth/token/revoke API



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
    tokenRevokeRequest := *openapiclient.NewTokenRevokeRequest("AccessTokenIdentifier_example") // TokenRevokeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenOperationsApi.AuthTokenRevokeApi(context.Background(), serviceApiKey).TokenRevokeRequest(tokenRevokeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenOperationsApi.AuthTokenRevokeApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenRevokeApi`: TokenRevokeResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenOperationsApi.AuthTokenRevokeApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceApiKey** | **string** | serviceApiKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenRevokeApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenRevokeRequest** | [**TokenRevokeRequest**](TokenRevokeRequest.md) |  | 

### Return type

[**TokenRevokeResponse**](TokenRevokeResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenUpdateApi

> TokenUpdateResponse AuthTokenUpdateApi(ctx, serviceApiKey).TokenUpdateRequest(tokenUpdateRequest).Execute()

/api/auth/token/update API



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
    tokenUpdateRequest := *openapiclient.NewTokenUpdateRequest("AccessToken_example") // TokenUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenOperationsApi.AuthTokenUpdateApi(context.Background(), serviceApiKey).TokenUpdateRequest(tokenUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenOperationsApi.AuthTokenUpdateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenUpdateApi`: TokenUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenOperationsApi.AuthTokenUpdateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceApiKey** | **string** | serviceApiKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenUpdateRequest** | [**TokenUpdateRequest**](TokenUpdateRequest.md) |  | 

### Return type

[**TokenUpdateResponse**](TokenUpdateResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

