# \IntrospectionEndpointAPI

All URIs are relative to *https://us.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthIntrospectionApi**](IntrospectionEndpointAPI.md#AuthIntrospectionApi) | **Post** /api/{serviceId}/auth/introspection | Process Introspection Request
[**AuthIntrospectionStandardApi**](IntrospectionEndpointAPI.md#AuthIntrospectionStandardApi) | **Post** /api/{serviceId}/auth/introspection/standard | Process OAuth 2.0 Introspection Request



## AuthIntrospectionApi

> IntrospectionResponse AuthIntrospectionApi(ctx, serviceId).IntrospectionRequest(introspectionRequest).Execute()

Process Introspection Request



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
	introspectionRequest := *openapiclient.NewIntrospectionRequest("Token_example") // IntrospectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntrospectionEndpointAPI.AuthIntrospectionApi(context.Background(), serviceId).IntrospectionRequest(introspectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntrospectionEndpointAPI.AuthIntrospectionApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthIntrospectionApi`: IntrospectionResponse
	fmt.Fprintf(os.Stdout, "Response from `IntrospectionEndpointAPI.AuthIntrospectionApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthIntrospectionApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **introspectionRequest** | [**IntrospectionRequest**](IntrospectionRequest.md) |  | 

### Return type

[**IntrospectionResponse**](IntrospectionResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthIntrospectionStandardApi

> StandardIntrospectionResponse AuthIntrospectionStandardApi(ctx, serviceId).StandardIntrospectionRequest(standardIntrospectionRequest).Execute()

Process OAuth 2.0 Introspection Request



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
	standardIntrospectionRequest := *openapiclient.NewStandardIntrospectionRequest("Parameters_example") // StandardIntrospectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntrospectionEndpointAPI.AuthIntrospectionStandardApi(context.Background(), serviceId).StandardIntrospectionRequest(standardIntrospectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntrospectionEndpointAPI.AuthIntrospectionStandardApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthIntrospectionStandardApi`: StandardIntrospectionResponse
	fmt.Fprintf(os.Stdout, "Response from `IntrospectionEndpointAPI.AuthIntrospectionStandardApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthIntrospectionStandardApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **standardIntrospectionRequest** | [**StandardIntrospectionRequest**](StandardIntrospectionRequest.md) |  | 

### Return type

[**StandardIntrospectionResponse**](StandardIntrospectionResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

