# \RevocationEndpointApi

All URIs are relative to *https://beta.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthRevocationApi**](RevocationEndpointApi.md#AuthRevocationApi) | **Post** /api/{serviceId}/auth/revocation | /api/{serviceId}/auth/revocation API



## AuthRevocationApi

> RevocationResponse AuthRevocationApi(ctx, serviceId).RevocationRequest(revocationRequest).Execute()

/api/{serviceId}/auth/revocation API



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
    revocationRequest := *openapiclient.NewRevocationRequest("Parameters_example") // RevocationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RevocationEndpointApi.AuthRevocationApi(context.Background(), serviceId).RevocationRequest(revocationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevocationEndpointApi.AuthRevocationApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthRevocationApi`: RevocationResponse
    fmt.Fprintf(os.Stdout, "Response from `RevocationEndpointApi.AuthRevocationApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthRevocationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revocationRequest** | [**RevocationRequest**](RevocationRequest.md) |  | 

### Return type

[**RevocationResponse**](RevocationResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

