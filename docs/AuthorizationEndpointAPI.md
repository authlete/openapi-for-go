# \AuthorizationEndpointAPI

All URIs are relative to *https://us.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServiceIdAuthAuthorizationTicketInfoGet**](AuthorizationEndpointAPI.md#ApiServiceIdAuthAuthorizationTicketInfoGet) | **Get** /api/{serviceId}/auth/authorization/ticket/info | Get Ticket Information
[**ApiServiceIdAuthAuthorizationTicketUpdatePost**](AuthorizationEndpointAPI.md#ApiServiceIdAuthAuthorizationTicketUpdatePost) | **Post** /api/{serviceId}/auth/authorization/ticket/update | Update Ticket Information
[**AuthAuthorizationApi**](AuthorizationEndpointAPI.md#AuthAuthorizationApi) | **Post** /api/{serviceId}/auth/authorization | Process Authorization Request
[**AuthAuthorizationFailApi**](AuthorizationEndpointAPI.md#AuthAuthorizationFailApi) | **Post** /api/{serviceId}/auth/authorization/fail | Fail Authorization Request
[**AuthAuthorizationIssueApi**](AuthorizationEndpointAPI.md#AuthAuthorizationIssueApi) | **Post** /api/{serviceId}/auth/authorization/issue | Issue Authorization Response



## ApiServiceIdAuthAuthorizationTicketInfoGet

> AuthorizationTicketInfoResponse ApiServiceIdAuthAuthorizationTicketInfoGet(ctx).AuthorizationTicketInfoRequest(authorizationTicketInfoRequest).Execute()

Get Ticket Information

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
	authorizationTicketInfoRequest := *openapiclient.NewAuthorizationTicketInfoRequest("Ticket_example") // AuthorizationTicketInfoRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationEndpointAPI.ApiServiceIdAuthAuthorizationTicketInfoGet(context.Background()).AuthorizationTicketInfoRequest(authorizationTicketInfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationEndpointAPI.ApiServiceIdAuthAuthorizationTicketInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiServiceIdAuthAuthorizationTicketInfoGet`: AuthorizationTicketInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationEndpointAPI.ApiServiceIdAuthAuthorizationTicketInfoGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceIdAuthAuthorizationTicketInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationTicketInfoRequest** | [**AuthorizationTicketInfoRequest**](AuthorizationTicketInfoRequest.md) |  | 

### Return type

[**AuthorizationTicketInfoResponse**](AuthorizationTicketInfoResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiServiceIdAuthAuthorizationTicketUpdatePost

> AuthorizationTicketUpdateResponse ApiServiceIdAuthAuthorizationTicketUpdatePost(ctx).AuthorizationTicketUpdateRequest(authorizationTicketUpdateRequest).Execute()

Update Ticket Information

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
	authorizationTicketUpdateRequest := *openapiclient.NewAuthorizationTicketUpdateRequest("Ticket_example", "Info_example") // AuthorizationTicketUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationEndpointAPI.ApiServiceIdAuthAuthorizationTicketUpdatePost(context.Background()).AuthorizationTicketUpdateRequest(authorizationTicketUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationEndpointAPI.ApiServiceIdAuthAuthorizationTicketUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiServiceIdAuthAuthorizationTicketUpdatePost`: AuthorizationTicketUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationEndpointAPI.ApiServiceIdAuthAuthorizationTicketUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceIdAuthAuthorizationTicketUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationTicketUpdateRequest** | [**AuthorizationTicketUpdateRequest**](AuthorizationTicketUpdateRequest.md) |  | 

### Return type

[**AuthorizationTicketUpdateResponse**](AuthorizationTicketUpdateResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthAuthorizationApi

> AuthorizationResponse AuthAuthorizationApi(ctx, serviceId).AuthorizationRequest(authorizationRequest).Execute()

Process Authorization Request



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
	authorizationRequest := *openapiclient.NewAuthorizationRequest("Parameters_example") // AuthorizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationEndpointAPI.AuthAuthorizationApi(context.Background(), serviceId).AuthorizationRequest(authorizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationEndpointAPI.AuthAuthorizationApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthAuthorizationApi`: AuthorizationResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationEndpointAPI.AuthAuthorizationApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthAuthorizationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizationRequest** | [**AuthorizationRequest**](AuthorizationRequest.md) |  | 

### Return type

[**AuthorizationResponse**](AuthorizationResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthAuthorizationFailApi

> AuthorizationFailResponse AuthAuthorizationFailApi(ctx, serviceId).AuthorizationFailRequest(authorizationFailRequest).Execute()

Fail Authorization Request



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
	authorizationFailRequest := *openapiclient.NewAuthorizationFailRequest("Ticket_example", "Reason_example") // AuthorizationFailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationEndpointAPI.AuthAuthorizationFailApi(context.Background(), serviceId).AuthorizationFailRequest(authorizationFailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationEndpointAPI.AuthAuthorizationFailApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthAuthorizationFailApi`: AuthorizationFailResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationEndpointAPI.AuthAuthorizationFailApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthAuthorizationFailApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizationFailRequest** | [**AuthorizationFailRequest**](AuthorizationFailRequest.md) |  | 

### Return type

[**AuthorizationFailResponse**](AuthorizationFailResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthAuthorizationIssueApi

> AuthorizationIssueResponse AuthAuthorizationIssueApi(ctx, serviceId).AuthorizationIssueRequest(authorizationIssueRequest).Execute()

Issue Authorization Response



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
	authorizationIssueRequest := *openapiclient.NewAuthorizationIssueRequest("Ticket_example", "Subject_example") // AuthorizationIssueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorizationEndpointAPI.AuthAuthorizationIssueApi(context.Background(), serviceId).AuthorizationIssueRequest(authorizationIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationEndpointAPI.AuthAuthorizationIssueApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthAuthorizationIssueApi`: AuthorizationIssueResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorizationEndpointAPI.AuthAuthorizationIssueApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthAuthorizationIssueApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizationIssueRequest** | [**AuthorizationIssueRequest**](AuthorizationIssueRequest.md) |  | 

### Return type

[**AuthorizationIssueResponse**](AuthorizationIssueResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

