# \TokenOperationsApi

All URIs are relative to *https://beta.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthTokenCreateApi**](TokenOperationsApi.md#AuthTokenCreateApi) | **Post** /api/{serviceId}/auth/token/create | /api/{serviceId}/auth/token/create API
[**AuthTokenCreateBatchApi**](TokenOperationsApi.md#AuthTokenCreateBatchApi) | **Post** /api/{serviceId}/auth/token/create/batch | /api/{serviceId}/auth/token/create/batch API
[**AuthTokenCreateBatchStatusApi**](TokenOperationsApi.md#AuthTokenCreateBatchStatusApi) | **Post** /api/{serviceId}/auth/token/create/batch/status/{requestId} | /api/{serviceId}/auth/token/create/batch/status/{requestId} API
[**AuthTokenDeleteApi**](TokenOperationsApi.md#AuthTokenDeleteApi) | **Delete** /api/{serviceId}/auth/token/delete/{accessTokenIdentifier} | /api/{serviceId}/auth/token/delete API
[**AuthTokenGetListApi**](TokenOperationsApi.md#AuthTokenGetListApi) | **Get** /api/{serviceId}/auth/token/get/list | /api/{serviceId}/auth/token/get/list API
[**AuthTokenRevokeApi**](TokenOperationsApi.md#AuthTokenRevokeApi) | **Post** /api/{serviceId}/auth/token/revoke | /api/{serviceId}/auth/token/revoke API
[**AuthTokenUpdateApi**](TokenOperationsApi.md#AuthTokenUpdateApi) | **Post** /api/{serviceId}/auth/token/update | /api/{serviceId}/auth/token/update API



## AuthTokenCreateApi

> TokenCreateResponse AuthTokenCreateApi(ctx, serviceId).TokenCreateRequest(tokenCreateRequest).Execute()

/api/{serviceId}/auth/token/create API



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
    serviceId := "serviceId_example" // string | A service ID.
    tokenCreateRequest := *openapiclient.NewTokenCreateRequest(openapiclient.grant_type("AUTHORIZATION_CODE"), int64(123)) // TokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenOperationsApi.AuthTokenCreateApi(context.Background(), serviceId).TokenCreateRequest(tokenCreateRequest).Execute()
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
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenCreateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenCreateRequest** | [**TokenCreateRequest**](TokenCreateRequest.md) |  | 

### Return type

[**TokenCreateResponse**](TokenCreateResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenCreateBatchApi

> TokenCreateBatchResponse AuthTokenCreateBatchApi(ctx, serviceId).TokenCreateBatchRequestInner(tokenCreateBatchRequestInner).Execute()

/api/{serviceId}/auth/token/create/batch API



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
    serviceId := "serviceId_example" // string | A service ID.
    tokenCreateBatchRequestInner := []openapiclient.TokenCreateBatchRequestInner{*openapiclient.NewTokenCreateBatchRequestInner(openapiclient.grant_type("AUTHORIZATION_CODE"), int64(123), "AccessToken_example")} // []TokenCreateBatchRequestInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenOperationsApi.AuthTokenCreateBatchApi(context.Background(), serviceId).TokenCreateBatchRequestInner(tokenCreateBatchRequestInner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenOperationsApi.AuthTokenCreateBatchApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenCreateBatchApi`: TokenCreateBatchResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenOperationsApi.AuthTokenCreateBatchApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenCreateBatchApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenCreateBatchRequestInner** | [**[]TokenCreateBatchRequestInner**](TokenCreateBatchRequestInner.md) |  | 

### Return type

[**TokenCreateBatchResponse**](TokenCreateBatchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenCreateBatchStatusApi

> TokenCreateBatchStatusResponse AuthTokenCreateBatchStatusApi(ctx, serviceId, requestId).Execute()

/api/{serviceId}/auth/token/create/batch/status/{requestId} API



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
    serviceId := "serviceId_example" // string | A service ID.
    requestId := "requestId_example" // string | A request ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenOperationsApi.AuthTokenCreateBatchStatusApi(context.Background(), serviceId, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokenOperationsApi.AuthTokenCreateBatchStatusApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthTokenCreateBatchStatusApi`: TokenCreateBatchStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TokenOperationsApi.AuthTokenCreateBatchStatusApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 
**requestId** | **string** | A request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenCreateBatchStatusApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TokenCreateBatchStatusResponse**](TokenCreateBatchStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenDeleteApi

> AuthTokenDeleteApi(ctx, serviceId, accessTokenIdentifier).Execute()

/api/{serviceId}/auth/token/delete API



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
    serviceId := "serviceId_example" // string | A service ID.
    accessTokenIdentifier := "accessTokenIdentifier_example" // string | The identifier of an existing access token. The identifier is the value of the access token or the value of the hash of the access token. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenOperationsApi.AuthTokenDeleteApi(context.Background(), serviceId, accessTokenIdentifier).Execute()
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
**serviceId** | **string** | A service ID. | 
**accessTokenIdentifier** | **string** | The identifier of an existing access token. The identifier is the value of the access token or the value of the hash of the access token.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenDeleteApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenGetListApi

> TokenGetListResponse AuthTokenGetListApi(ctx, serviceId).ClientIdentifier(clientIdentifier).Subject(subject).Start(start).End(end).Execute()

/api/{serviceId}/auth/token/get/list API



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
    serviceId := "serviceId_example" // string | A service ID.
    clientIdentifier := "clientIdentifier_example" // string | Client Identifier (client ID or client ID alias).  (optional)
    subject := "subject_example" // string | Unique user ID.  (optional)
    start := int32(56) // int32 | Start index of search results (inclusive). The default value is 0. (optional)
    end := int32(56) // int32 | End index of search results (exclusive). The default value is 5.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenOperationsApi.AuthTokenGetListApi(context.Background(), serviceId).ClientIdentifier(clientIdentifier).Subject(subject).Start(start).End(end).Execute()
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
**serviceId** | **string** | A service ID. | 

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

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenRevokeApi

> TokenRevokeResponse AuthTokenRevokeApi(ctx, serviceId).TokenRevokeRequest(tokenRevokeRequest).Execute()

/api/{serviceId}/auth/token/revoke API



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
    serviceId := "serviceId_example" // string | A service ID.
    tokenRevokeRequest := *openapiclient.NewTokenRevokeRequest() // TokenRevokeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenOperationsApi.AuthTokenRevokeApi(context.Background(), serviceId).TokenRevokeRequest(tokenRevokeRequest).Execute()
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
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenRevokeApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenRevokeRequest** | [**TokenRevokeRequest**](TokenRevokeRequest.md) |  | 

### Return type

[**TokenRevokeResponse**](TokenRevokeResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenUpdateApi

> TokenUpdateResponse AuthTokenUpdateApi(ctx, serviceId).TokenUpdateRequest(tokenUpdateRequest).Execute()

/api/{serviceId}/auth/token/update API



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
    serviceId := "serviceId_example" // string | A service ID.
    tokenUpdateRequest := *openapiclient.NewTokenUpdateRequest("AccessToken_example") // TokenUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokenOperationsApi.AuthTokenUpdateApi(context.Background(), serviceId).TokenUpdateRequest(tokenUpdateRequest).Execute()
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
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenUpdateRequest** | [**TokenUpdateRequest**](TokenUpdateRequest.md) |  | 

### Return type

[**TokenUpdateResponse**](TokenUpdateResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

