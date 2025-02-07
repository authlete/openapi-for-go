# \TokenEndpointAPI

All URIs are relative to *https://us.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthTokenApi**](TokenEndpointAPI.md#AuthTokenApi) | **Post** /api/{serviceId}/auth/token | Process Token Request
[**AuthTokenFailApi**](TokenEndpointAPI.md#AuthTokenFailApi) | **Post** /api/{serviceId}/auth/token/fail | Fail Token Request
[**AuthTokenIssueApi**](TokenEndpointAPI.md#AuthTokenIssueApi) | **Post** /api/{serviceId}/auth/token/issue | Issue Token Response
[**IdtokenReissueApi**](TokenEndpointAPI.md#IdtokenReissueApi) | **Post** /api/{serviceId}/idtoken/reissue | Reissue ID Token



## AuthTokenApi

> TokenResponse AuthTokenApi(ctx, serviceId).TokenRequest(tokenRequest).Execute()

Process Token Request



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
	tokenRequest := *openapiclient.NewTokenRequest("Parameters_example") // TokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenEndpointAPI.AuthTokenApi(context.Background(), serviceId).TokenRequest(tokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenEndpointAPI.AuthTokenApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthTokenApi`: TokenResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenEndpointAPI.AuthTokenApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenFailApi

> TokenFailResponse AuthTokenFailApi(ctx, serviceId).TokenFailRequest(tokenFailRequest).Execute()

Fail Token Request



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
	tokenFailRequest := *openapiclient.NewTokenFailRequest("Ticket_example", "Reason_example") // TokenFailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenEndpointAPI.AuthTokenFailApi(context.Background(), serviceId).TokenFailRequest(tokenFailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenEndpointAPI.AuthTokenFailApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthTokenFailApi`: TokenFailResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenEndpointAPI.AuthTokenFailApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenFailApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenFailRequest** | [**TokenFailRequest**](TokenFailRequest.md) |  | 

### Return type

[**TokenFailResponse**](TokenFailResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenIssueApi

> TokenIssueResponse AuthTokenIssueApi(ctx, serviceId).TokenIssueRequest(tokenIssueRequest).Execute()

Issue Token Response



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
	tokenIssueRequest := *openapiclient.NewTokenIssueRequest("Ticket_example", "Subject_example") // TokenIssueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenEndpointAPI.AuthTokenIssueApi(context.Background(), serviceId).TokenIssueRequest(tokenIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenEndpointAPI.AuthTokenIssueApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthTokenIssueApi`: TokenIssueResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenEndpointAPI.AuthTokenIssueApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenIssueApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenIssueRequest** | [**TokenIssueRequest**](TokenIssueRequest.md) |  | 

### Return type

[**TokenIssueResponse**](TokenIssueResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdtokenReissueApi

> IdtokenReissueResponse IdtokenReissueApi(ctx, serviceId).IdtokenReissueRequest(idtokenReissueRequest).Execute()

Reissue ID Token



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
	idtokenReissueRequest := *openapiclient.NewIdtokenReissueRequest("AccessToken_example", "RefreshToken_example") // IdtokenReissueRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenEndpointAPI.IdtokenReissueApi(context.Background(), serviceId).IdtokenReissueRequest(idtokenReissueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenEndpointAPI.IdtokenReissueApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IdtokenReissueApi`: IdtokenReissueResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenEndpointAPI.IdtokenReissueApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIdtokenReissueApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idtokenReissueRequest** | [**IdtokenReissueRequest**](IdtokenReissueRequest.md) |  | 

### Return type

[**IdtokenReissueResponse**](IdtokenReissueResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

