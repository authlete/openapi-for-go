# \DeviceFlowApi

All URIs are relative to *https://api.authlete.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceAuthorizationApi**](DeviceFlowApi.md#DeviceAuthorizationApi) | **Post** /device/authorization | /device/authorization API
[**DeviceCompleteApi**](DeviceFlowApi.md#DeviceCompleteApi) | **Post** /device/complete | /device/complete API
[**DeviceVerificationApi**](DeviceFlowApi.md#DeviceVerificationApi) | **Post** /device/verification | /device/verification API



## DeviceAuthorizationApi

> DeviceAuthorizationResponse DeviceAuthorizationApi(ctx).DeviceAuthorizationRequest(deviceAuthorizationRequest).Execute()

/device/authorization API



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
    deviceAuthorizationRequest := *openapiclient.NewDeviceAuthorizationRequest("Parameters_example") // DeviceAuthorizationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceFlowApi.DeviceAuthorizationApi(context.Background()).DeviceAuthorizationRequest(deviceAuthorizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceFlowApi.DeviceAuthorizationApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceAuthorizationApi`: DeviceAuthorizationResponse
    fmt.Fprintf(os.Stdout, "Response from `DeviceFlowApi.DeviceAuthorizationApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceAuthorizationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceAuthorizationRequest** | [**DeviceAuthorizationRequest**](DeviceAuthorizationRequest.md) |  | 

### Return type

[**DeviceAuthorizationResponse**](DeviceAuthorizationResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceCompleteApi

> DeviceCompleteResponse DeviceCompleteApi(ctx).DeviceCompleteRequest(deviceCompleteRequest).Execute()

/device/complete API



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
    deviceCompleteRequest := *openapiclient.NewDeviceCompleteRequest("UserCode_example", "Result_example", "Subject_example") // DeviceCompleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceFlowApi.DeviceCompleteApi(context.Background()).DeviceCompleteRequest(deviceCompleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceFlowApi.DeviceCompleteApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceCompleteApi`: DeviceCompleteResponse
    fmt.Fprintf(os.Stdout, "Response from `DeviceFlowApi.DeviceCompleteApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceCompleteApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceCompleteRequest** | [**DeviceCompleteRequest**](DeviceCompleteRequest.md) |  | 

### Return type

[**DeviceCompleteResponse**](DeviceCompleteResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceVerificationApi

> DeviceVerificationResponse DeviceVerificationApi(ctx).DeviceVerificationRequest(deviceVerificationRequest).Execute()

/device/verification API



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
    deviceVerificationRequest := *openapiclient.NewDeviceVerificationRequest("UserCode_example") // DeviceVerificationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceFlowApi.DeviceVerificationApi(context.Background()).DeviceVerificationRequest(deviceVerificationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceFlowApi.DeviceVerificationApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceVerificationApi`: DeviceVerificationResponse
    fmt.Fprintf(os.Stdout, "Response from `DeviceFlowApi.DeviceVerificationApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceVerificationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceVerificationRequest** | [**DeviceVerificationRequest**](DeviceVerificationRequest.md) |  | 

### Return type

[**DeviceVerificationResponse**](DeviceVerificationResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

