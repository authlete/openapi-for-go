# \TokenOperationsApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthTokenCreateApi**](TokenOperationsApi.md#AuthTokenCreateApi) | **Post** /api/auth/token/create | /api/auth/token/create API
[**AuthTokenDeleteApi**](TokenOperationsApi.md#AuthTokenDeleteApi) | **Delete** /api/auth/token/delete/{accessTokenIdentifier} | /api/auth/token/delete API
[**AuthTokenGetListApi**](TokenOperationsApi.md#AuthTokenGetListApi) | **Get** /api/auth/token/get/list | /api/auth/token/get/list API
[**AuthTokenRevokeApi**](TokenOperationsApi.md#AuthTokenRevokeApi) | **Post** /api/auth/token/revoke | /api/auth/token/revoke API
[**AuthTokenUpdateApi**](TokenOperationsApi.md#AuthTokenUpdateApi) | **Post** /api/auth/token/update | /api/auth/token/update API



## AuthTokenCreateApi

> TokenCreateResponse AuthTokenCreateApi(ctx).TokenCreateRequest(tokenCreateRequest).Execute()

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
    tokenCreateRequest := *openapiclient.NewTokenCreateRequest(openapiclient.grant_type("AUTHORIZATION_CODE"), int64(123)) // TokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenOperationsApi.AuthTokenCreateApi(context.Background()).TokenCreateRequest(tokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenOperationsApi.AuthTokenCreateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenCreateApi`: TokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenOperationsApi.AuthTokenCreateApi`: %v\n", resp)
}
```

### Path Parameters



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

> AuthTokenDeleteApi(ctx, accessTokenIdentifier).Execute()

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
    accessTokenIdentifier := "accessTokenIdentifier_example" // string | The identifier of an existing access token. The identifier is the value of the access token or the value of the hash of the access token. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenOperationsApi.AuthTokenDeleteApi(context.Background(), accessTokenIdentifier).Execute()
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

> TokenGetListResponse AuthTokenGetListApi(ctx).ClientIdentifier(clientIdentifier).Subject(subject).Start(start).End(end).Execute()

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
    clientIdentifier := "clientIdentifier_example" // string | Client Identifier (client ID or client ID alias).  (optional)
    subject := "subject_example" // string | Unique user ID.  (optional)
    start := int32(56) // int32 | Start index of search results (inclusive). The default value is 0. (optional)
    end := int32(56) // int32 | End index of search results (exclusive). The default value is 5.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenOperationsApi.AuthTokenGetListApi(context.Background()).ClientIdentifier(clientIdentifier).Subject(subject).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenOperationsApi.AuthTokenGetListApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenGetListApi`: TokenGetListResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenOperationsApi.AuthTokenGetListApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenGetListApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientIdentifier** | **string** | Client Identifier (client ID or client ID alias).  | 
 **subject** | **string** | Unique user ID.  | 
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

> TokenRevokeResponse AuthTokenRevokeApi(ctx).TokenRevokeRequest(tokenRevokeRequest).Execute()

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
    tokenRevokeRequest := *openapiclient.NewTokenRevokeRequest() // TokenRevokeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenOperationsApi.AuthTokenRevokeApi(context.Background()).TokenRevokeRequest(tokenRevokeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenOperationsApi.AuthTokenRevokeApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenRevokeApi`: TokenRevokeResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenOperationsApi.AuthTokenRevokeApi`: %v\n", resp)
}
```

### Path Parameters



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

> TokenUpdateResponse AuthTokenUpdateApi(ctx).TokenUpdateRequest(tokenUpdateRequest).Execute()

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
    tokenUpdateRequest := *openapiclient.NewTokenUpdateRequest("AccessToken_example") // TokenUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenOperationsApi.AuthTokenUpdateApi(context.Background()).TokenUpdateRequest(tokenUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenOperationsApi.AuthTokenUpdateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenUpdateApi`: TokenUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenOperationsApi.AuthTokenUpdateApi`: %v\n", resp)
}
```

### Path Parameters



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

