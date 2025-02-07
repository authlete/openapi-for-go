# \DeviceFlowAPI

All URIs are relative to *https://us.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAuthorizationApi**](DeviceFlowAPI.md#DeviceAuthorizationApi) | **Post** /api/{serviceId}/device/authorization | Process Device Authorization Request
[**DeviceCompleteApi**](DeviceFlowAPI.md#DeviceCompleteApi) | **Post** /api/{serviceId}/device/complete | Complete Device Authorization
[**DeviceVerificationApi**](DeviceFlowAPI.md#DeviceVerificationApi) | **Post** /api/{serviceId}/device/verification | Process Device Verification Request



## DeviceAuthorizationApi

> DeviceAuthorizationResponse DeviceAuthorizationApi(ctx, serviceId).DeviceAuthorizationRequest(deviceAuthorizationRequest).Execute()

Process Device Authorization Request



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
	deviceAuthorizationRequest := *openapiclient.NewDeviceAuthorizationRequest("Parameters_example") // DeviceAuthorizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceFlowAPI.DeviceAuthorizationApi(context.Background(), serviceId).DeviceAuthorizationRequest(deviceAuthorizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceFlowAPI.DeviceAuthorizationApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceAuthorizationApi`: DeviceAuthorizationResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceFlowAPI.DeviceAuthorizationApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAuthorizationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceAuthorizationRequest** | [**DeviceAuthorizationRequest**](DeviceAuthorizationRequest.md) |  | 

### Return type

[**DeviceAuthorizationResponse**](DeviceAuthorizationResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceCompleteApi

> DeviceCompleteResponse DeviceCompleteApi(ctx, serviceId).DeviceCompleteRequest(deviceCompleteRequest).Execute()

Complete Device Authorization



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
	deviceCompleteRequest := *openapiclient.NewDeviceCompleteRequest("UserCode_example", "Result_example", "Subject_example") // DeviceCompleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceFlowAPI.DeviceCompleteApi(context.Background(), serviceId).DeviceCompleteRequest(deviceCompleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceFlowAPI.DeviceCompleteApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceCompleteApi`: DeviceCompleteResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceFlowAPI.DeviceCompleteApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceCompleteApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceCompleteRequest** | [**DeviceCompleteRequest**](DeviceCompleteRequest.md) |  | 

### Return type

[**DeviceCompleteResponse**](DeviceCompleteResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceVerificationApi

> DeviceVerificationResponse DeviceVerificationApi(ctx, serviceId).DeviceVerificationRequest(deviceVerificationRequest).Execute()

Process Device Verification Request



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
	deviceVerificationRequest := *openapiclient.NewDeviceVerificationRequest("UserCode_example") // DeviceVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceFlowAPI.DeviceVerificationApi(context.Background(), serviceId).DeviceVerificationRequest(deviceVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceFlowAPI.DeviceVerificationApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceVerificationApi`: DeviceVerificationResponse
	fmt.Fprintf(os.Stdout, "Response from `DeviceFlowAPI.DeviceVerificationApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeviceVerificationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceVerificationRequest** | [**DeviceVerificationRequest**](DeviceVerificationRequest.md) |  | 

### Return type

[**DeviceVerificationResponse**](DeviceVerificationResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

