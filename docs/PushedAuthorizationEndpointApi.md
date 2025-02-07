# \PushedAuthorizationEndpointAPI

All URIs are relative to *https://us.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PushedAuthReqApi**](PushedAuthorizationEndpointAPI.md#PushedAuthReqApi) | **Post** /api/{serviceId}/pushed_auth_req | Process Pushed Authorization Request



## PushedAuthReqApi

> PushedAuthorizationResponse PushedAuthReqApi(ctx, serviceId).PushedAuthorizationRequest(pushedAuthorizationRequest).Execute()

Process Pushed Authorization Request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/openapi-for-go"
)

func main() {
	serviceId := "serviceId_example" // string | A service ID.
	pushedAuthorizationRequest := *openapiclient.NewPushedAuthorizationRequest("Parameters_example") // PushedAuthorizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PushedAuthorizationEndpointAPI.PushedAuthReqApi(context.Background(), serviceId).PushedAuthorizationRequest(pushedAuthorizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PushedAuthorizationEndpointAPI.PushedAuthReqApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PushedAuthReqApi`: PushedAuthorizationResponse
	fmt.Fprintf(os.Stdout, "Response from `PushedAuthorizationEndpointAPI.PushedAuthReqApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPushedAuthReqApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pushedAuthorizationRequest** | [**PushedAuthorizationRequest**](PushedAuthorizationRequest.md) |  | 

### Return type

[**PushedAuthorizationResponse**](PushedAuthorizationResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

