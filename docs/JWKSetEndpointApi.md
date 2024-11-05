# \JWKSetEndpointApi

All URIs are relative to *https://beta.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceJwksGetApi**](JWKSetEndpointApi.md#ServiceJwksGetApi) | **Get** /api/{serviceId}/service/jwks/get | /api/{serviceId}/service/jwks/get API



## ServiceJwksGetApi

> ServiceJwksGetResponse ServiceJwksGetApi(ctx, serviceId).IncludePrivateKeys(includePrivateKeys).Pretty(pretty).Execute()

/api/{serviceId}/service/jwks/get API



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
    includePrivateKeys := true // bool | The boolean value that indicates whether the response should include the private keys associated with the service or not. If `true`, the private keys are included in the response. The default value is `false`. (optional)
    pretty := true // bool | This boolean value indicates whether the JSON in the response should be formatted or not. If `true`, the JSON in the response is pretty-formatted. The default value is `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JWKSetEndpointApi.ServiceJwksGetApi(context.Background(), serviceId).IncludePrivateKeys(includePrivateKeys).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JWKSetEndpointApi.ServiceJwksGetApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceJwksGetApi`: ServiceJwksGetResponse
    fmt.Fprintf(os.Stdout, "Response from `JWKSetEndpointApi.ServiceJwksGetApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceJwksGetApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includePrivateKeys** | **bool** | The boolean value that indicates whether the response should include the private keys associated with the service or not. If &#x60;true&#x60;, the private keys are included in the response. The default value is &#x60;false&#x60;. | 
 **pretty** | **bool** | This boolean value indicates whether the JSON in the response should be formatted or not. If &#x60;true&#x60;, the JSON in the response is pretty-formatted. The default value is &#x60;false&#x60;. | 

### Return type

[**ServiceJwksGetResponse**](ServiceJwksGetResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

