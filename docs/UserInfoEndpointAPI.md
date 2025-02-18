# \UserInfoEndpointAPI

All URIs are relative to *https://us.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthUserinfoApi**](UserInfoEndpointAPI.md#AuthUserinfoApi) | **Post** /api/{serviceId}/auth/userinfo | Process UserInfo Request
[**AuthUserinfoIssueApi**](UserInfoEndpointAPI.md#AuthUserinfoIssueApi) | **Post** /api/{serviceId}/auth/userinfo/issue | Issue UserInfo Response



## AuthUserinfoApi

> UserinfoResponse AuthUserinfoApi(ctx, serviceId).UserinfoRequest(userinfoRequest).Execute()

Process UserInfo Request



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
	userinfoRequest := *openapiclient.NewUserinfoRequest("Token_example") // UserinfoRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserInfoEndpointAPI.AuthUserinfoApi(context.Background(), serviceId).UserinfoRequest(userinfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserInfoEndpointAPI.AuthUserinfoApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthUserinfoApi`: UserinfoResponse
	fmt.Fprintf(os.Stdout, "Response from `UserInfoEndpointAPI.AuthUserinfoApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthUserinfoApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userinfoRequest** | [**UserinfoRequest**](UserinfoRequest.md) |  | 

### Return type

[**UserinfoResponse**](UserinfoResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthUserinfoIssueApi

> UserinfoIssueResponse AuthUserinfoIssueApi(ctx, serviceId).UserinfoIssueRequest(userinfoIssueRequest).Execute()

Issue UserInfo Response



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
	userinfoIssueRequest := *openapiclient.NewUserinfoIssueRequest("Token_example") // UserinfoIssueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserInfoEndpointAPI.AuthUserinfoIssueApi(context.Background(), serviceId).UserinfoIssueRequest(userinfoIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserInfoEndpointAPI.AuthUserinfoIssueApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthUserinfoIssueApi`: UserinfoIssueResponse
	fmt.Fprintf(os.Stdout, "Response from `UserInfoEndpointAPI.AuthUserinfoIssueApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthUserinfoIssueApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userinfoIssueRequest** | [**UserinfoIssueRequest**](UserinfoIssueRequest.md) |  | 

### Return type

[**UserinfoIssueResponse**](UserinfoIssueResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

