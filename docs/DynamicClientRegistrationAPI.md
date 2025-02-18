# \DynamicClientRegistrationAPI

All URIs are relative to *https://us.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClientRegistrationApi**](DynamicClientRegistrationAPI.md#ClientRegistrationApi) | **Post** /api/{serviceId}/client/registration | Register Client
[**ClientRegistrationDeleteApi**](DynamicClientRegistrationAPI.md#ClientRegistrationDeleteApi) | **Post** /api/{serviceId}/client/registration/delete | Delete Client
[**ClientRegistrationGetApi**](DynamicClientRegistrationAPI.md#ClientRegistrationGetApi) | **Post** /api/{serviceId}/client/registration/get | Get Client
[**ClientRegistrationUpdateApi**](DynamicClientRegistrationAPI.md#ClientRegistrationUpdateApi) | **Post** /api/{serviceId}/client/registration/update | Update Client



## ClientRegistrationApi

> ClientRegistrationResponse ClientRegistrationApi(ctx, serviceId).ClientRegistrationRequest(clientRegistrationRequest).Execute()

Register Client



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
	clientRegistrationRequest := *openapiclient.NewClientRegistrationRequest("Json_example") // ClientRegistrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicClientRegistrationAPI.ClientRegistrationApi(context.Background(), serviceId).ClientRegistrationRequest(clientRegistrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicClientRegistrationAPI.ClientRegistrationApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientRegistrationApi`: ClientRegistrationResponse
	fmt.Fprintf(os.Stdout, "Response from `DynamicClientRegistrationAPI.ClientRegistrationApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientRegistrationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRegistrationRequest** | [**ClientRegistrationRequest**](ClientRegistrationRequest.md) |  | 

### Return type

[**ClientRegistrationResponse**](ClientRegistrationResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientRegistrationDeleteApi

> ClientRegistrationDeleteResponse ClientRegistrationDeleteApi(ctx, serviceId).ClientRegistrationDeleteRequest(clientRegistrationDeleteRequest).Execute()

Delete Client



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
	clientRegistrationDeleteRequest := *openapiclient.NewClientRegistrationDeleteRequest("ClientId_example", "Token_example") // ClientRegistrationDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicClientRegistrationAPI.ClientRegistrationDeleteApi(context.Background(), serviceId).ClientRegistrationDeleteRequest(clientRegistrationDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicClientRegistrationAPI.ClientRegistrationDeleteApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientRegistrationDeleteApi`: ClientRegistrationDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `DynamicClientRegistrationAPI.ClientRegistrationDeleteApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientRegistrationDeleteApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRegistrationDeleteRequest** | [**ClientRegistrationDeleteRequest**](ClientRegistrationDeleteRequest.md) |  | 

### Return type

[**ClientRegistrationDeleteResponse**](ClientRegistrationDeleteResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientRegistrationGetApi

> ClientRegistrationResponse ClientRegistrationGetApi(ctx, serviceId).ClientRegistrationRequest(clientRegistrationRequest).Execute()

Get Client



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
	clientRegistrationRequest := *openapiclient.NewClientRegistrationRequest("Json_example") // ClientRegistrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicClientRegistrationAPI.ClientRegistrationGetApi(context.Background(), serviceId).ClientRegistrationRequest(clientRegistrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicClientRegistrationAPI.ClientRegistrationGetApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientRegistrationGetApi`: ClientRegistrationResponse
	fmt.Fprintf(os.Stdout, "Response from `DynamicClientRegistrationAPI.ClientRegistrationGetApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientRegistrationGetApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRegistrationRequest** | [**ClientRegistrationRequest**](ClientRegistrationRequest.md) |  | 

### Return type

[**ClientRegistrationResponse**](ClientRegistrationResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientRegistrationUpdateApi

> ClientRegistrationUpdateResponse ClientRegistrationUpdateApi(ctx, serviceId).ClientRegistrationUpdateRequest(clientRegistrationUpdateRequest).Execute()

Update Client



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
	clientRegistrationUpdateRequest := *openapiclient.NewClientRegistrationUpdateRequest("ClientId_example", "Token_example", "Json_example") // ClientRegistrationUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicClientRegistrationAPI.ClientRegistrationUpdateApi(context.Background(), serviceId).ClientRegistrationUpdateRequest(clientRegistrationUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicClientRegistrationAPI.ClientRegistrationUpdateApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientRegistrationUpdateApi`: ClientRegistrationUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `DynamicClientRegistrationAPI.ClientRegistrationUpdateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientRegistrationUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRegistrationUpdateRequest** | [**ClientRegistrationUpdateRequest**](ClientRegistrationUpdateRequest.md) |  | 

### Return type

[**ClientRegistrationUpdateResponse**](ClientRegistrationUpdateResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

