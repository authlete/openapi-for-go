# \PushedAuthorizationEndpointApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PushedAuthReqApi**](PushedAuthorizationEndpointApi.md#PushedAuthReqApi) | **Post** /api/pushed_auth_req | /api/pushed_auth_req API



## PushedAuthReqApi

> PushedAuthorizationResponse PushedAuthReqApi(ctx).PushedAuthorizationRequest(pushedAuthorizationRequest).Execute()

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
    resp, r, err := apiClient.PushedAuthorizationEndpointApi.PushedAuthReqApi(context.Background()).PushedAuthorizationRequest(pushedAuthorizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushedAuthorizationEndpointApi.PushedAuthReqApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PushedAuthReqApi`: PushedAuthorizationResponse
    fmt.Fprintf(os.Stdout, "Response from `PushedAuthorizationEndpointApi.PushedAuthReqApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPushedAuthReqApiRequest struct via the builder pattern


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

