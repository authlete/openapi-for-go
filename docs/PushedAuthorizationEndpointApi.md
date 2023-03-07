# \PushedAuthorizationEndpointApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PushedAuthApi**](PushedAuthorizationEndpointApi.md#PushedAuthApi) | **Post** /api/pushed_auth_req | /api/pushed_auth_req API



## PushedAuthApi

> PushedAuthorizationResponse PushedAuthApi(ctx).PushedAuthorizationRequest(pushedAuthorizationRequest).Execute()

/api/pushed_auth_req API



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
    pushedAuthorizationRequest := *openapiclient.NewPushedAuthorizationRequest("Parameters_example") // PushedAuthorizationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PushedAuthorizationEndpointApi.PushedAuthApi(context.Background()).PushedAuthorizationRequest(pushedAuthorizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushedAuthorizationEndpointApi.PushedAuthApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PushedAuthApi`: PushedAuthorizationResponse
    fmt.Fprintf(os.Stdout, "Response from `PushedAuthorizationEndpointApi.PushedAuthApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPushedAuthApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pushedAuthorizationRequest** | [**PushedAuthorizationRequest**](PushedAuthorizationRequest.md) |  | 

### Return type

[**PushedAuthorizationResponse**](PushedAuthorizationResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

