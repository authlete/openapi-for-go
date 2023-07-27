# \JoseObjectApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JoseVerifyApi**](JoseObjectApi.md#JoseVerifyApi) | **Post** /api/{serviceApiKey}/jose/verify | /api/jose/verify API



## JoseVerifyApi

> JoseVerifyResponse JoseVerifyApi(ctx, serviceApiKey).JoseVerifyRequest(joseVerifyRequest).Execute()

/api/jose/verify API



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
    joseVerifyRequest := *openapiclient.NewJoseVerifyRequest("Jose_example") // JoseVerifyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JoseObjectApi.JoseVerifyApi(context.Background(), serviceApiKey).JoseVerifyRequest(joseVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JoseObjectApi.JoseVerifyApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JoseVerifyApi`: JoseVerifyResponse
    fmt.Fprintf(os.Stdout, "Response from `JoseObjectApi.JoseVerifyApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceApiKey** | **string** | serviceApiKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiJoseVerifyApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **joseVerifyRequest** | [**JoseVerifyRequest**](JoseVerifyRequest.md) |  | 

### Return type

[**JoseVerifyResponse**](JoseVerifyResponse.md)

### Authorization

[ServiceOwnerCredentials](../README.md#ServiceOwnerCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

