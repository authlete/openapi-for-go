# \GrantManagementEndpointApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GrantMApi**](GrantManagementEndpointApi.md#GrantMApi) | **Post** /api/{serviceApiKey}/gm | /api/gm API



## GrantMApi

> GMResponse GrantMApi(ctx, serviceApiKey).GMRequest(gMRequest).Execute()

/api/gm API



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
    gMRequest := *openapiclient.NewGMRequest() // GMRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GrantManagementEndpointApi.GrantMApi(context.Background(), serviceApiKey).GMRequest(gMRequest).Execute()
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
**serviceApiKey** | **string** | serviceApiKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantMApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gMRequest** | [**GMRequest**](GMRequest.md) |  | 

### Return type

[**GMResponse**](GMResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

