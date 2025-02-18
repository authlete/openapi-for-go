# \CIBAAPI

All URIs are relative to *https://us.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackchannelAuthenticationApi**](CIBAAPI.md#BackchannelAuthenticationApi) | **Post** /api/{serviceId}/backchannel/authentication | Process Backchannel Authentication Request
[**BackchannelAuthenticationCompleteApi**](CIBAAPI.md#BackchannelAuthenticationCompleteApi) | **Post** /api/{serviceId}/backchannel/authentication/complete | Complete Backchannel Authentication
[**BackchannelAuthenticationFailApi**](CIBAAPI.md#BackchannelAuthenticationFailApi) | **Post** /api/{serviceId}/backchannel/authentication/fail | Fail Backchannel Authentication Request
[**BackchannelAuthenticationIssueApi**](CIBAAPI.md#BackchannelAuthenticationIssueApi) | **Post** /api/{serviceId}/backchannel/authentication/issue | Issue Backchannel Authentication Response



## BackchannelAuthenticationApi

> BackchannelAuthenticationResponse BackchannelAuthenticationApi(ctx, serviceId).BackchannelAuthenticationRequest(backchannelAuthenticationRequest).Execute()

Process Backchannel Authentication Request



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
	backchannelAuthenticationRequest := *openapiclient.NewBackchannelAuthenticationRequest("Parameters_example") // BackchannelAuthenticationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CIBAAPI.BackchannelAuthenticationApi(context.Background(), serviceId).BackchannelAuthenticationRequest(backchannelAuthenticationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIBAAPI.BackchannelAuthenticationApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackchannelAuthenticationApi`: BackchannelAuthenticationResponse
	fmt.Fprintf(os.Stdout, "Response from `CIBAAPI.BackchannelAuthenticationApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackchannelAuthenticationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backchannelAuthenticationRequest** | [**BackchannelAuthenticationRequest**](BackchannelAuthenticationRequest.md) |  | 

### Return type

[**BackchannelAuthenticationResponse**](BackchannelAuthenticationResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackchannelAuthenticationCompleteApi

> BackchannelAuthenticationCompleteResponse BackchannelAuthenticationCompleteApi(ctx, serviceId).BackchannelAuthenticationCompleteRequest(backchannelAuthenticationCompleteRequest).Execute()

Complete Backchannel Authentication



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
	backchannelAuthenticationCompleteRequest := *openapiclient.NewBackchannelAuthenticationCompleteRequest("Ticket_example", "Result_example", "Subject_example") // BackchannelAuthenticationCompleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CIBAAPI.BackchannelAuthenticationCompleteApi(context.Background(), serviceId).BackchannelAuthenticationCompleteRequest(backchannelAuthenticationCompleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIBAAPI.BackchannelAuthenticationCompleteApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackchannelAuthenticationCompleteApi`: BackchannelAuthenticationCompleteResponse
	fmt.Fprintf(os.Stdout, "Response from `CIBAAPI.BackchannelAuthenticationCompleteApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackchannelAuthenticationCompleteApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backchannelAuthenticationCompleteRequest** | [**BackchannelAuthenticationCompleteRequest**](BackchannelAuthenticationCompleteRequest.md) |  | 

### Return type

[**BackchannelAuthenticationCompleteResponse**](BackchannelAuthenticationCompleteResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackchannelAuthenticationFailApi

> BackchannelAuthenticationFailResponse BackchannelAuthenticationFailApi(ctx, serviceId).BackchannelAuthenticationFailRequest(backchannelAuthenticationFailRequest).Execute()

Fail Backchannel Authentication Request



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
	backchannelAuthenticationFailRequest := *openapiclient.NewBackchannelAuthenticationFailRequest("Ticket_example", "Reason_example") // BackchannelAuthenticationFailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CIBAAPI.BackchannelAuthenticationFailApi(context.Background(), serviceId).BackchannelAuthenticationFailRequest(backchannelAuthenticationFailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIBAAPI.BackchannelAuthenticationFailApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackchannelAuthenticationFailApi`: BackchannelAuthenticationFailResponse
	fmt.Fprintf(os.Stdout, "Response from `CIBAAPI.BackchannelAuthenticationFailApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackchannelAuthenticationFailApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backchannelAuthenticationFailRequest** | [**BackchannelAuthenticationFailRequest**](BackchannelAuthenticationFailRequest.md) |  | 

### Return type

[**BackchannelAuthenticationFailResponse**](BackchannelAuthenticationFailResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackchannelAuthenticationIssueApi

> BackchannelAuthenticationIssueResponse BackchannelAuthenticationIssueApi(ctx, serviceId).BackchannelAuthenticationIssueRequest(backchannelAuthenticationIssueRequest).Execute()

Issue Backchannel Authentication Response



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
	backchannelAuthenticationIssueRequest := *openapiclient.NewBackchannelAuthenticationIssueRequest("Ticket_example") // BackchannelAuthenticationIssueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CIBAAPI.BackchannelAuthenticationIssueApi(context.Background(), serviceId).BackchannelAuthenticationIssueRequest(backchannelAuthenticationIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIBAAPI.BackchannelAuthenticationIssueApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackchannelAuthenticationIssueApi`: BackchannelAuthenticationIssueResponse
	fmt.Fprintf(os.Stdout, "Response from `CIBAAPI.BackchannelAuthenticationIssueApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBackchannelAuthenticationIssueApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **backchannelAuthenticationIssueRequest** | [**BackchannelAuthenticationIssueRequest**](BackchannelAuthenticationIssueRequest.md) |  | 

### Return type

[**BackchannelAuthenticationIssueResponse**](BackchannelAuthenticationIssueResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

