# \ConfigurationEndpointApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceConfigurationApi**](ConfigurationEndpointApi.md#ServiceConfigurationApi) | **Get** /api/{serviceApiKey}/service/configuration | /api/service/configuration API



## ServiceConfigurationApi

> map[string]interface{} ServiceConfigurationApi(ctx, serviceApiKey).Pretty(pretty).Execute()

/api/service/configuration API



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
    serviceApiKey := "serviceApiKey_example" // string | serviceApiKey
    pretty := true // bool | This boolean value indicates whether the JSON in the response should be formatted or not. If `true`, the JSON in the response is pretty-formatted. The default value is `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationEndpointApi.ServiceConfigurationApi(context.Background(), serviceApiKey).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationEndpointApi.ServiceConfigurationApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceConfigurationApi`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationEndpointApi.ServiceConfigurationApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceApiKey** | **string** | serviceApiKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceConfigurationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | This boolean value indicates whether the JSON in the response should be formatted or not. If &#x60;true&#x60;, the JSON in the response is pretty-formatted. The default value is &#x60;false&#x60;. | 

### Return type

**map[string]interface{}**

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

