# \HardwareSecurityKeyAPI

All URIs are relative to *https://us.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HskCreateApi**](HardwareSecurityKeyAPI.md#HskCreateApi) | **Post** /api/{serviceId}/hsk/create | Create Security Key
[**HskDeleteApi**](HardwareSecurityKeyAPI.md#HskDeleteApi) | **Delete** /api/{serviceId}/hsk/delete/{handle} | Delete Security Key
[**HskGetApi**](HardwareSecurityKeyAPI.md#HskGetApi) | **Get** /api/{serviceId}/hsk/get/{handle} | Get Security Key
[**HskGetListApi**](HardwareSecurityKeyAPI.md#HskGetListApi) | **Get** /api/{serviceId}/hsk/get/list | List Security Keys



## HskCreateApi

> HskCreateResponse HskCreateApi(ctx, serviceId).HskCreateRequest(hskCreateRequest).Execute()

Create Security Key

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
	hskCreateRequest := *openapiclient.NewHskCreateRequest() // HskCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HardwareSecurityKeyAPI.HskCreateApi(context.Background(), serviceId).HskCreateRequest(hskCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HardwareSecurityKeyAPI.HskCreateApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HskCreateApi`: HskCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `HardwareSecurityKeyAPI.HskCreateApi`: %v\n", resp)
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

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HskDeleteApi

> HskDeleteResponse HskDeleteApi(ctx, serviceId, handle).Execute()

Delete Security Key

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
	handle := "handle_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HardwareSecurityKeyAPI.HskDeleteApi(context.Background(), serviceId, handle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HardwareSecurityKeyAPI.HskDeleteApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HskDeleteApi`: HskDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `HardwareSecurityKeyAPI.HskDeleteApi`: %v\n", resp)
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

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HskGetApi

> HskGetResponse HskGetApi(ctx, serviceId, handle).Execute()

Get Security Key

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
	handle := "handle_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HardwareSecurityKeyAPI.HskGetApi(context.Background(), serviceId, handle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HardwareSecurityKeyAPI.HskGetApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HskGetApi`: HskGetResponse
	fmt.Fprintf(os.Stdout, "Response from `HardwareSecurityKeyAPI.HskGetApi`: %v\n", resp)
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

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HskGetListApi

> HskGetListResponse HskGetListApi(ctx, serviceId).Execute()

List Security Keys

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
	resp, r, err := apiClient.HardwareSecurityKeyAPI.HskGetListApi(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HardwareSecurityKeyAPI.HskGetListApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HskGetListApi`: HskGetListResponse
	fmt.Fprintf(os.Stdout, "Response from `HardwareSecurityKeyAPI.HskGetListApi`: %v\n", resp)
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

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

