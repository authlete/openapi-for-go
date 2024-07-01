# \ServiceManagementApi

All URIs are relative to *https://beta.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceCreateApi**](ServiceManagementApi.md#ServiceCreateApi) | **Post** /api/service/create | /api/service/create API
[**ServiceDeleteApi**](ServiceManagementApi.md#ServiceDeleteApi) | **Delete** /api/{serviceId}/service/delete | /api/{serviceId}/service/delete API
[**ServiceGetApi**](ServiceManagementApi.md#ServiceGetApi) | **Get** /api/{serviceId}/service/get | /api/{serviceId}/service/get API
[**ServiceGetListApi**](ServiceManagementApi.md#ServiceGetListApi) | **Get** /api/service/get/list | /api/service/get/list API
[**ServiceUpdateApi**](ServiceManagementApi.md#ServiceUpdateApi) | **Post** /api/{serviceId}/service/update | /api/{serviceId}/service/update API



## ServiceCreateApi

> Service ServiceCreateApi(ctx).Service(service).Execute()

/api/service/create API



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
    service := *openapiclient.NewService() // Service |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceManagementApi.ServiceCreateApi(context.Background()).Service(service).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceManagementApi.ServiceCreateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceCreateApi`: Service
    fmt.Fprintf(os.Stdout, "Response from `ServiceManagementApi.ServiceCreateApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceCreateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | [**Service**](Service.md) |  | 

### Return type

[**Service**](Service.md)

### Authorization

[CreateService](../README.md#CreateService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceDeleteApi

> ServiceDeleteApi(ctx, serviceId).Execute()

/api/{serviceId}/service/delete API



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
    resp, r, err := apiClient.ServiceManagementApi.ServiceDeleteApi(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceManagementApi.ServiceDeleteApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceDeleteApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[DeleteService](../README.md#DeleteService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceGetApi

> ServiceGetApi200Response ServiceGetApi(ctx, serviceId).Execute()

/api/{serviceId}/service/get API



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
    resp, r, err := apiClient.ServiceManagementApi.ServiceGetApi(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceManagementApi.ServiceGetApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceGetApi`: ServiceGetApi200Response
    fmt.Fprintf(os.Stdout, "Response from `ServiceManagementApi.ServiceGetApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceGetApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceGetApi200Response**](ServiceGetApi200Response.md)

### Authorization

[ViewService](../README.md#ViewService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceGetListApi

> ServiceGetListApi200Response ServiceGetListApi(ctx).Start(start).End(end).Execute()

/api/service/get/list API



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
    start := int32(56) // int32 | Start index (inclusive) of the result set. The default value is 0. Must not be a negative number. (optional)
    end := int32(56) // int32 | End index (exclusive) of the result set. The default value is 5. Must not be a negative number. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceManagementApi.ServiceGetListApi(context.Background()).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceManagementApi.ServiceGetListApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceGetListApi`: ServiceGetListApi200Response
    fmt.Fprintf(os.Stdout, "Response from `ServiceManagementApi.ServiceGetListApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceGetListApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Start index (inclusive) of the result set. The default value is 0. Must not be a negative number. | 
 **end** | **int32** | End index (exclusive) of the result set. The default value is 5. Must not be a negative number. | 

### Return type

[**ServiceGetListApi200Response**](ServiceGetListApi200Response.md)

### Authorization

[Admin](../README.md#Admin), [ViewService](../README.md#ViewService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceUpdateApi

> Service ServiceUpdateApi(ctx, serviceId).Service(service).Execute()

/api/{serviceId}/service/update API



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
    service := *openapiclient.NewService() // Service |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceManagementApi.ServiceUpdateApi(context.Background(), serviceId).Service(service).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceManagementApi.ServiceUpdateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceUpdateApi`: Service
    fmt.Fprintf(os.Stdout, "Response from `ServiceManagementApi.ServiceUpdateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **service** | [**Service**](Service.md) |  | 

### Return type

[**Service**](Service.md)

### Authorization

[ModifyService](../README.md#ModifyService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

