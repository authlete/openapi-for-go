# \FederationEndpointApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FederationConfigurationApi**](FederationEndpointApi.md#FederationConfigurationApi) | **Post** /api/{serviceApiKey}/federation/configuration | /api/federation/configuration API



## FederationConfigurationApi

> FederationConfigurationResponse FederationConfigurationApi(ctx, serviceApiKey).Body(body).Execute()

/api/federation/configuration API



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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederationEndpointApi.FederationConfigurationApi(context.Background(), serviceApiKey).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationEndpointApi.FederationConfigurationApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FederationConfigurationApi`: FederationConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `FederationEndpointApi.FederationConfigurationApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceApiKey** | **string** | serviceApiKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiFederationConfigurationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**FederationConfigurationResponse**](FederationConfigurationResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

