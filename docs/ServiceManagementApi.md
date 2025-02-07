# \ServiceManagementAPI

All URIs are relative to *https://us.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceConfigurationApi**](ServiceManagementAPI.md#ServiceConfigurationApi) | **Get** /api/{serviceId}/service/configuration | Get Service Configuration
[**ServiceCreateApi**](ServiceManagementAPI.md#ServiceCreateApi) | **Post** /api/service/create | Create Service
[**ServiceDeleteApi**](ServiceManagementAPI.md#ServiceDeleteApi) | **Delete** /api/{serviceId}/service/delete | Delete Service ⚡
[**ServiceGetApi**](ServiceManagementAPI.md#ServiceGetApi) | **Get** /api/{serviceId}/service/get | Get Service
[**ServiceGetListApi**](ServiceManagementAPI.md#ServiceGetListApi) | **Get** /api/service/get/list | List Services
[**ServiceUpdateApi**](ServiceManagementAPI.md#ServiceUpdateApi) | **Post** /api/{serviceId}/service/update | Update Service



## ServiceConfigurationApi

> map[string]interface{} ServiceConfigurationApi(ctx, serviceId).Pretty(pretty).Patch(patch).Execute()

Get Service Configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/openapi-for-go"
)

func main() {
	serviceId := "serviceId_example" // string | A service ID.
	pretty := true // bool | This boolean value indicates whether the JSON in the response should be formatted or not. If `true`, the JSON in the response is pretty-formatted. The default value is `false`. (optional)
	patch := "patch_example" // string | Get the JSON Patch [RFC 6902 JavaScript Object Notation (JSON) Patch](https://www.rfc-editor.org/rfc/rfc6902) to be applied. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceManagementAPI.ServiceConfigurationApi(context.Background(), serviceId).Pretty(pretty).Patch(patch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceManagementAPI.ServiceConfigurationApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceConfigurationApi`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ServiceManagementAPI.ServiceConfigurationApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceConfigurationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | This boolean value indicates whether the JSON in the response should be formatted or not. If &#x60;true&#x60;, the JSON in the response is pretty-formatted. The default value is &#x60;false&#x60;. | 
 **patch** | **string** | Get the JSON Patch [RFC 6902 JavaScript Object Notation (JSON) Patch](https://www.rfc-editor.org/rfc/rfc6902) to be applied. | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceCreateApi

> Service ServiceCreateApi(ctx).Service(service).Execute()

Create Service



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/openapi-for-go"
)

func main() {
	service := *openapiclient.NewService() // Service |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceManagementAPI.ServiceCreateApi(context.Background()).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceManagementAPI.ServiceCreateApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceCreateApi`: Service
	fmt.Fprintf(os.Stdout, "Response from `ServiceManagementAPI.ServiceCreateApi`: %v\n", resp)
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

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceDeleteApi

> ServiceDeleteApi(ctx, serviceId).Execute()

Delete Service ⚡



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/openapi-for-go"
)

func main() {
	serviceId := "serviceId_example" // string | A service ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceManagementAPI.ServiceDeleteApi(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceManagementAPI.ServiceDeleteApi``: %v\n", err)
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

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceGetApi

> ServiceGetApi200Response ServiceGetApi(ctx, serviceId).Execute()

Get Service



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/openapi-for-go"
)

func main() {
	serviceId := "serviceId_example" // string | A service ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceManagementAPI.ServiceGetApi(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceManagementAPI.ServiceGetApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceGetApi`: ServiceGetApi200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceManagementAPI.ServiceGetApi`: %v\n", resp)
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

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceGetListApi

> ServiceGetListApi200Response ServiceGetListApi(ctx).Start(start).End(end).Execute()

List Services



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/openapi-for-go"
)

func main() {
	start := int32(56) // int32 | Start index (inclusive) of the result set. The default value is 0. Must not be a negative number. (optional)
	end := int32(56) // int32 | End index (exclusive) of the result set. The default value is 5. Must not be a negative number. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceManagementAPI.ServiceGetListApi(context.Background()).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceManagementAPI.ServiceGetListApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceGetListApi`: ServiceGetListApi200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceManagementAPI.ServiceGetListApi`: %v\n", resp)
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

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceUpdateApi

> Service ServiceUpdateApi(ctx, serviceId).Service(service).Execute()

Update Service



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/openapi-for-go"
)

func main() {
	serviceId := "serviceId_example" // string | A service ID.
	service := *openapiclient.NewService() // Service |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceManagementAPI.ServiceUpdateApi(context.Background(), serviceId).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceManagementAPI.ServiceUpdateApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceUpdateApi`: Service
	fmt.Fprintf(os.Stdout, "Response from `ServiceManagementAPI.ServiceUpdateApi`: %v\n", resp)
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

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

