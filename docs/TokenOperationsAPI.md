# \TokenOperationsAPI

All URIs are relative to *https://us.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthTokenCreateApi**](TokenOperationsAPI.md#AuthTokenCreateApi) | **Post** /api/{serviceId}/auth/token/create | Create Access Token
[**AuthTokenDeleteApi**](TokenOperationsAPI.md#AuthTokenDeleteApi) | **Delete** /api/{serviceId}/auth/token/delete/{accessTokenIdentifier} | Delete Access Token
[**AuthTokenGetListApi**](TokenOperationsAPI.md#AuthTokenGetListApi) | **Get** /api/{serviceId}/auth/token/get/list | List Issued Tokens
[**AuthTokenRevokeApi**](TokenOperationsAPI.md#AuthTokenRevokeApi) | **Post** /api/{serviceId}/auth/token/revoke | Revoke Access Token
[**AuthTokenUpdateApi**](TokenOperationsAPI.md#AuthTokenUpdateApi) | **Post** /api/{serviceId}/auth/token/update | Update Access Token



## AuthTokenCreateApi

> TokenCreateResponse AuthTokenCreateApi(ctx, serviceId).TokenCreateRequest(tokenCreateRequest).Execute()

Create Access Token



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
	tokenCreateRequest := *openapiclient.NewTokenCreateRequest(openapiclient.grant_type("AUTHORIZATION_CODE"), int64(123)) // TokenCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenOperationsAPI.AuthTokenCreateApi(context.Background(), serviceId).TokenCreateRequest(tokenCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenOperationsAPI.AuthTokenCreateApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthTokenCreateApi`: TokenCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenOperationsAPI.AuthTokenCreateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenCreateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenCreateRequest** | [**TokenCreateRequest**](TokenCreateRequest.md) |  | 

### Return type

[**TokenCreateResponse**](TokenCreateResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenDeleteApi

> AuthTokenDeleteApi(ctx, serviceId, accessTokenIdentifier).Execute()

Delete Access Token



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
	accessTokenIdentifier := "accessTokenIdentifier_example" // string | The identifier of an existing access token. The identifier is the value of the access token or the value of the hash of the access token. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TokenOperationsAPI.AuthTokenDeleteApi(context.Background(), serviceId, accessTokenIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenOperationsAPI.AuthTokenDeleteApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 
**accessTokenIdentifier** | **string** | The identifier of an existing access token. The identifier is the value of the access token or the value of the hash of the access token.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenDeleteApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenGetListApi

> TokenGetListResponse AuthTokenGetListApi(ctx, serviceId).ClientIdentifier(clientIdentifier).Subject(subject).Start(start).End(end).Execute()

List Issued Tokens



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
	clientIdentifier := "clientIdentifier_example" // string | Client Identifier (client ID or client ID alias).  (optional)
	subject := "subject_example" // string | Unique user ID.  (optional)
	start := int32(56) // int32 | Start index of search results (inclusive). The default value is 0. (optional)
	end := int32(56) // int32 | End index of search results (exclusive). The default value is 5.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenOperationsAPI.AuthTokenGetListApi(context.Background(), serviceId).ClientIdentifier(clientIdentifier).Subject(subject).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenOperationsAPI.AuthTokenGetListApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthTokenGetListApi`: TokenGetListResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenOperationsAPI.AuthTokenGetListApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenGetListApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientIdentifier** | **string** | Client Identifier (client ID or client ID alias).  | 
 **subject** | **string** | Unique user ID.  | 
 **start** | **int32** | Start index of search results (inclusive). The default value is 0. | 
 **end** | **int32** | End index of search results (exclusive). The default value is 5.  | 

### Return type

[**TokenGetListResponse**](TokenGetListResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenRevokeApi

> TokenRevokeResponse AuthTokenRevokeApi(ctx, serviceId).TokenRevokeRequest(tokenRevokeRequest).Execute()

Revoke Access Token



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
	tokenRevokeRequest := *openapiclient.NewTokenRevokeRequest() // TokenRevokeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenOperationsAPI.AuthTokenRevokeApi(context.Background(), serviceId).TokenRevokeRequest(tokenRevokeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenOperationsAPI.AuthTokenRevokeApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthTokenRevokeApi`: TokenRevokeResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenOperationsAPI.AuthTokenRevokeApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenRevokeApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenRevokeRequest** | [**TokenRevokeRequest**](TokenRevokeRequest.md) |  | 

### Return type

[**TokenRevokeResponse**](TokenRevokeResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthTokenUpdateApi

> TokenUpdateResponse AuthTokenUpdateApi(ctx, serviceId).TokenUpdateRequest(tokenUpdateRequest).Execute()

Update Access Token



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
	tokenUpdateRequest := *openapiclient.NewTokenUpdateRequest("AccessToken_example") // TokenUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokenOperationsAPI.AuthTokenUpdateApi(context.Background(), serviceId).TokenUpdateRequest(tokenUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokenOperationsAPI.AuthTokenUpdateApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthTokenUpdateApi`: TokenUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `TokenOperationsAPI.AuthTokenUpdateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenUpdateRequest** | [**TokenUpdateRequest**](TokenUpdateRequest.md) |  | 

### Return type

[**TokenUpdateResponse**](TokenUpdateResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

