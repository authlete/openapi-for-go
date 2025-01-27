# \DeviceFlowApi

All URIs are relative to *https://beta.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAuthorizationApi**](DeviceFlowApi.md#DeviceAuthorizationApi) | **Post** /api/{serviceId}/device/authorization | /api/{serviceId}/device/authorization API
[**DeviceCompleteApi**](DeviceFlowApi.md#DeviceCompleteApi) | **Post** /api/{serviceId}/device/complete | /api/{serviceId}/device/complete API
[**DeviceVerificationApi**](DeviceFlowApi.md#DeviceVerificationApi) | **Post** /api/{serviceId}/device/verification | /api/{serviceId}/device/verification API



## DeviceAuthorizationApi

> DeviceAuthorizationResponse DeviceAuthorizationApi(ctx, serviceId).DeviceAuthorizationRequest(deviceAuthorizationRequest).Execute()

/api/{serviceId}/device/authorization API



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
    deviceAuthorizationRequest := *openapiclient.NewDeviceAuthorizationRequest("Parameters_example") // DeviceAuthorizationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceFlowApi.DeviceAuthorizationApi(context.Background(), serviceId).DeviceAuthorizationRequest(deviceAuthorizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceFlowApi.DeviceAuthorizationApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAuthorizationApi`: DeviceAuthorizationResponse
    fmt.Fprintf(os.Stdout, "Response from `DeviceFlowApi.DeviceAuthorizationApi`: %v\n", resp)
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

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceCompleteApi

> DeviceCompleteResponse DeviceCompleteApi(ctx, serviceId).DeviceCompleteRequest(deviceCompleteRequest).Execute()

/api/{serviceId}/device/complete API



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
    deviceCompleteRequest := *openapiclient.NewDeviceCompleteRequest("UserCode_example", "Result_example", "Subject_example") // DeviceCompleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceFlowApi.DeviceCompleteApi(context.Background(), serviceId).DeviceCompleteRequest(deviceCompleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceFlowApi.DeviceCompleteApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceCompleteApi`: DeviceCompleteResponse
    fmt.Fprintf(os.Stdout, "Response from `DeviceFlowApi.DeviceCompleteApi`: %v\n", resp)
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

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceVerificationApi

> DeviceVerificationResponse DeviceVerificationApi(ctx, serviceId).DeviceVerificationRequest(deviceVerificationRequest).Execute()

/api/{serviceId}/device/verification API



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
    deviceVerificationRequest := *openapiclient.NewDeviceVerificationRequest("UserCode_example") // DeviceVerificationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceFlowApi.DeviceVerificationApi(context.Background(), serviceId).DeviceVerificationRequest(deviceVerificationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceFlowApi.DeviceVerificationApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceVerificationApi`: DeviceVerificationResponse
    fmt.Fprintf(os.Stdout, "Response from `DeviceFlowApi.DeviceVerificationApi`: %v\n", resp)
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

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

