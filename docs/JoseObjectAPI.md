# \JoseObjectAPI

All URIs are relative to *https://us.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JoseVerifyApi**](JoseObjectAPI.md#JoseVerifyApi) | **Post** /api/{serviceId}/jose/verify | Verify JOSE



## JoseVerifyApi

> JoseVerifyResponse JoseVerifyApi(ctx, serviceId).JoseVerifyRequest(joseVerifyRequest).Execute()

Verify JOSE



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/authlete/openapi-for-go/v3"
)

func main() {
	serviceId := "serviceId_example" // string | A service ID.
	joseVerifyRequest := *openapiclient.NewJoseVerifyRequest("Jose_example") // JoseVerifyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JoseObjectAPI.JoseVerifyApi(context.Background(), serviceId).JoseVerifyRequest(joseVerifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JoseObjectAPI.JoseVerifyApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JoseVerifyApi`: JoseVerifyResponse
	fmt.Fprintf(os.Stdout, "Response from `JoseObjectAPI.JoseVerifyApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJoseVerifyApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **joseVerifyRequest** | [**JoseVerifyRequest**](JoseVerifyRequest.md) |  | 

### Return type

[**JoseVerifyResponse**](JoseVerifyResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

