# \PushedAuthorizationEndpointApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PushedAuthReqApi**](PushedAuthorizationEndpointApi.md#PushedAuthReqApi) | **Post** /api/pushed_auth_req | /api/pushed_auth_req API



## PushedAuthReqApi

> PushedAuthReqResponse PushedAuthReqApi(ctx).PushedAuthReqRequest(pushedAuthReqRequest).Execute()

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
    pushedAuthReqRequest := *openapiclient.NewPushedAuthReqRequest("Parameters_example") // PushedAuthReqRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PushedAuthorizationEndpointApi.PushedAuthReqApi(context.Background()).PushedAuthReqRequest(pushedAuthReqRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushedAuthorizationEndpointApi.PushedAuthReqApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PushedAuthReqApi`: PushedAuthReqResponse
    fmt.Fprintf(os.Stdout, "Response from `PushedAuthorizationEndpointApi.PushedAuthReqApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPushedAuthReqApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pushedAuthReqRequest** | [**PushedAuthReqRequest**](PushedAuthReqRequest.md) |  | 

### Return type

[**PushedAuthReqResponse**](PushedAuthReqResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

