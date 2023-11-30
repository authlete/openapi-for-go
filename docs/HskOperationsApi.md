# \HskOperationsApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HskCreateApi**](HskOperationsApi.md#HskCreateApi) | **Post** /api/hsk/create | /api/hsk/create API
[**HskDeleteApi**](HskOperationsApi.md#HskDeleteApi) | **Delete** /api/hsk/delete/{handle} | /api/hsk/delete/{handle} API
[**HskGetApi**](HskOperationsApi.md#HskGetApi) | **Get** /api/hsk/get/{handle} | /api/hsk/get/{handle} API
[**HskGetListApi**](HskOperationsApi.md#HskGetListApi) | **Get** /api/hsk/get/list | /api/hsk/get/list API



## HskCreateApi

> HskCreateResponse HskCreateApi(ctx).HskCreateRequest(hskCreateRequest).Execute()

/api/hsk/create API

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
    hskCreateRequest := *openapiclient.NewHskCreateRequest() // HskCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HskOperationsApi.HskCreateApi(context.Background()).HskCreateRequest(hskCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HskOperationsApi.HskCreateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HskCreateApi`: HskCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `HskOperationsApi.HskCreateApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHskCreateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hskCreateRequest** | [**HskCreateRequest**](HskCreateRequest.md) |  | 

### Return type

[**HskCreateResponse**](HskCreateResponse.md)

### Authorization

[ServiceOwnerCredentials](../README.md#ServiceOwnerCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HskDeleteApi

> HskDeleteResponse HskDeleteApi(ctx, handle).Execute()

/api/hsk/delete/{handle} API

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
    handle := "handle_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HskOperationsApi.HskDeleteApi(context.Background(), handle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HskOperationsApi.HskDeleteApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HskDeleteApi`: HskDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `HskOperationsApi.HskDeleteApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHskDeleteApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HskDeleteResponse**](HskDeleteResponse.md)

### Authorization

[ServiceOwnerCredentials](../README.md#ServiceOwnerCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HskGetApi

> HskGetResponse HskGetApi(ctx, handle).Execute()

/api/hsk/get/{handle} API

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
    handle := "handle_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HskOperationsApi.HskGetApi(context.Background(), handle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HskOperationsApi.HskGetApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HskGetApi`: HskGetResponse
    fmt.Fprintf(os.Stdout, "Response from `HskOperationsApi.HskGetApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHskGetApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HskGetResponse**](HskGetResponse.md)

### Authorization

[ServiceOwnerCredentials](../README.md#ServiceOwnerCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HskGetListApi

> HskGetListResponse HskGetListApi(ctx).Execute()

/api/hsk/get/list API

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
    resp, r, err := apiClient.HskOperationsApi.HskGetListApi(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HskOperationsApi.HskGetListApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HskGetListApi`: HskGetListResponse
    fmt.Fprintf(os.Stdout, "Response from `HskOperationsApi.HskGetListApi`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHskGetListApiRequest struct via the builder pattern


### Return type

[**HskGetListResponse**](HskGetListResponse.md)

### Authorization

[ServiceOwnerCredentials](../README.md#ServiceOwnerCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

