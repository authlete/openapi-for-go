# \RevocationEndpointApi

All URIs are relative to *https://api.authlete.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthRevocationApi**](RevocationEndpointApi.md#AuthRevocationApi) | **Post** /auth/revocation | /auth/revocation API



## AuthRevocationApi

> RevocationResponse AuthRevocationApi(ctx).RevocationRequest(revocationRequest).Execute()

/auth/revocation API



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
    revocationRequest := *openapiclient.NewRevocationRequest("Parameters_example") // RevocationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RevocationEndpointApi.AuthRevocationApi(context.Background()).RevocationRequest(revocationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RevocationEndpointApi.AuthRevocationApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthRevocationApi`: RevocationResponse
    fmt.Fprintf(os.Stdout, "Response from `RevocationEndpointApi.AuthRevocationApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthRevocationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revocationRequest** | [**RevocationRequest**](RevocationRequest.md) |  | 

### Return type

[**RevocationResponse**](RevocationResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

