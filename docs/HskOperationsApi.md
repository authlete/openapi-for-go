# \HskOperationsApi

All URIs are relative to *https://beta.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HskCreateApi**](HskOperationsApi.md#HskCreateApi) | **Post** /api/{serviceId}/hsk/create | /api/{serviceId}/hsk/create API
[**HskDeleteApi**](HskOperationsApi.md#HskDeleteApi) | **Delete** /api/{serviceId}/hsk/delete/{handle} | /api/{serviceId}/hsk/delete/{handle} API
[**HskGetApi**](HskOperationsApi.md#HskGetApi) | **Get** /api/{serviceId}/hsk/get/{handle} | /api/{serviceId}/hsk/get/{handle} API
[**HskGetListApi**](HskOperationsApi.md#HskGetListApi) | **Get** /api/{serviceId}/hsk/get/list | /api/{serviceId}/hsk/get/list API



## HskCreateApi

> HskCreateResponse HskCreateApi(ctx, serviceId).HskCreateRequest(hskCreateRequest).Execute()

/api/{serviceId}/hsk/create API

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
    hskCreateRequest := *openapiclient.NewHskCreateRequest() // HskCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HskOperationsApi.HskCreateApi(context.Background(), serviceId).HskCreateRequest(hskCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HskOperationsApi.HskCreateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HskCreateApi`: HskCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `HskOperationsApi.HskCreateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHskCreateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hskCreateRequest** | [**HskCreateRequest**](HskCreateRequest.md) |  | 

### Return type

[**HskCreateResponse**](HskCreateResponse.md)

### Authorization

[ModifyService](../README.md#ModifyService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HskDeleteApi

> HskDeleteResponse HskDeleteApi(ctx, serviceId, handle).Execute()

/api/{serviceId}/hsk/delete/{handle} API

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
    handle := "handle_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HskOperationsApi.HskDeleteApi(context.Background(), serviceId, handle).Execute()
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
**serviceId** | **string** | A service ID. | 
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHskDeleteApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HskDeleteResponse**](HskDeleteResponse.md)

### Authorization

[ModifyService](../README.md#ModifyService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HskGetApi

> HskGetResponse HskGetApi(ctx, serviceId, handle).Execute()

/api/{serviceId}/hsk/get/{handle} API

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
    handle := "handle_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HskOperationsApi.HskGetApi(context.Background(), serviceId, handle).Execute()
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
**serviceId** | **string** | A service ID. | 
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHskGetApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HskGetResponse**](HskGetResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HskGetListApi

> HskGetListResponse HskGetListApi(ctx, serviceId).Execute()

/api/{serviceId}/hsk/get/list API

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HskOperationsApi.HskGetListApi(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HskOperationsApi.HskGetListApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HskGetListApi`: HskGetListResponse
    fmt.Fprintf(os.Stdout, "Response from `HskOperationsApi.HskGetListApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHskGetListApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HskGetListResponse**](HskGetListResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

