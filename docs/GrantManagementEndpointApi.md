# \GrantManagementEndpointApi

All URIs are relative to *https://beta.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GrantMApi**](GrantManagementEndpointApi.md#GrantMApi) | **Post** /api/{serviceId}/gm | /api/{serviceId}/gm API



## GrantMApi

> GMResponse GrantMApi(ctx, serviceId).GMRequest(gMRequest).Execute()

/api/{serviceId}/gm API



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
    gMRequest := *openapiclient.NewGMRequest() // GMRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GrantManagementEndpointApi.GrantMApi(context.Background(), serviceId).GMRequest(gMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GrantManagementEndpointApi.GrantMApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GrantMApi`: GMResponse
    fmt.Fprintf(os.Stdout, "Response from `GrantManagementEndpointApi.GrantMApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantMApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gMRequest** | [**GMRequest**](GMRequest.md) |  | 

### Return type

[**GMResponse**](GMResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

