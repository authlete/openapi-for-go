# \ClientManagementAPI

All URIs are relative to *https://us.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClientAuthorizationDeleteApi**](ClientManagementAPI.md#ClientAuthorizationDeleteApi) | **Delete** /api/{serviceId}/client/authorization/delete/{clientId} | Delete Client Tokens
[**ClientAuthorizationGetListApi**](ClientManagementAPI.md#ClientAuthorizationGetListApi) | **Get** /api/{serviceId}/client/authorization/get/list | Get Authorized Applications
[**ClientAuthorizationUpdateApi**](ClientManagementAPI.md#ClientAuthorizationUpdateApi) | **Post** /api/{serviceId}/client/authorization/update/{clientId} | Update Client Tokens
[**ClientCreateApi**](ClientManagementAPI.md#ClientCreateApi) | **Post** /api/{serviceId}/client/create | Create Client
[**ClientDeleteApi**](ClientManagementAPI.md#ClientDeleteApi) | **Delete** /api/{serviceId}/client/delete/{clientId} | Delete Client ⚡
[**ClientExtensionRequestablesScopesDeleteApi**](ClientManagementAPI.md#ClientExtensionRequestablesScopesDeleteApi) | **Delete** /api/{serviceId}/client/extension/requestable_scopes/delete/{clientId} | Delete Requestable Scopes
[**ClientExtensionRequestablesScopesGetApi**](ClientManagementAPI.md#ClientExtensionRequestablesScopesGetApi) | **Get** /api/{serviceId}/client/extension/requestable_scopes/get/{clientId} | Get Requestable Scopes
[**ClientExtensionRequestablesScopesUpdateApi**](ClientManagementAPI.md#ClientExtensionRequestablesScopesUpdateApi) | **Put** /api/{serviceId}/client/extension/requestable_scopes/update/{clientId} | Update Requestable Scopes
[**ClientFlagUpdateApi**](ClientManagementAPI.md#ClientFlagUpdateApi) | **Post** /api/{serviceId}/client/lock_flag/update/{clientIdentifier} | Update Client Lock
[**ClientGetApi**](ClientManagementAPI.md#ClientGetApi) | **Get** /api/{serviceId}/client/get/{clientId} | Get Client
[**ClientGetListApi**](ClientManagementAPI.md#ClientGetListApi) | **Get** /api/{serviceId}/client/get/list | List Clients
[**ClientGrantedScopesDeleteApi**](ClientManagementAPI.md#ClientGrantedScopesDeleteApi) | **Delete** /api/{serviceId}/client/granted_scopes/delete/{clientId} | Delete Granted Scopes
[**ClientGrantedScopesGetApi**](ClientManagementAPI.md#ClientGrantedScopesGetApi) | **Get** /api/{serviceId}/client/granted_scopes/get/{clientId} | Get Granted Scopes
[**ClientSecretRefreshApi**](ClientManagementAPI.md#ClientSecretRefreshApi) | **Get** /api/{serviceId}/client/secret/refresh/{clientIdentifier} | Rotate Client Secret
[**ClientSecretUpdateApi**](ClientManagementAPI.md#ClientSecretUpdateApi) | **Post** /api/{serviceId}/client/secret/update/{clientIdentifier} | Update Client Secret
[**ClientUpdateApi**](ClientManagementAPI.md#ClientUpdateApi) | **Post** /api/{serviceId}/client/update/{clientId} | Update Client



## ClientAuthorizationDeleteApi

> ClientAuthorizationDeleteResponse ClientAuthorizationDeleteApi(ctx, serviceId, clientId, subject).Subject2(subject2).Execute()

Delete Client Tokens



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
	clientId := "clientId_example" // string | A client ID. 
	subject := "subject_example" // string | Unique user ID of an end-user. 
	subject2 := "subject_example" // string | Unique user ID of an end-user. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.ClientAuthorizationDeleteApi(context.Background(), serviceId, clientId, subject).Subject2(subject2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ClientAuthorizationDeleteApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientAuthorizationDeleteApi`: ClientAuthorizationDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.ClientAuthorizationDeleteApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 
**clientId** | **string** | A client ID.  | 
**subject** | **string** | Unique user ID of an end-user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientAuthorizationDeleteApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **subject2** | **string** | Unique user ID of an end-user.  | 

### Return type

[**ClientAuthorizationDeleteResponse**](ClientAuthorizationDeleteResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientAuthorizationGetListApi

> ClientAuthorizationGetListResponse ClientAuthorizationGetListApi(ctx, serviceId, subject).Subject2(subject2).Developer(developer).Start(start).End(end).Execute()

Get Authorized Applications



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
	subject := "subject_example" // string | Unique user ID of an end-user. 
	subject2 := "subject_example" // string | Unique user ID of an end-user. 
	developer := "developer_example" // string | Unique ID of a client developer.  (optional)
	start := int32(56) // int32 | Start index of search results (inclusive). The default value is 0. (optional)
	end := int32(56) // int32 | End index of search results (exclusive). The default value is 5.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.ClientAuthorizationGetListApi(context.Background(), serviceId, subject).Subject2(subject2).Developer(developer).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ClientAuthorizationGetListApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientAuthorizationGetListApi`: ClientAuthorizationGetListResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.ClientAuthorizationGetListApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 
**subject** | **string** | Unique user ID of an end-user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientAuthorizationGetListApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **subject2** | **string** | Unique user ID of an end-user.  | 
 **developer** | **string** | Unique ID of a client developer.  | 
 **start** | **int32** | Start index of search results (inclusive). The default value is 0. | 
 **end** | **int32** | End index of search results (exclusive). The default value is 5.  | 

### Return type

[**ClientAuthorizationGetListResponse**](ClientAuthorizationGetListResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientAuthorizationUpdateApi

> ClientAuthorizationUpdateResponse ClientAuthorizationUpdateApi(ctx, serviceId, clientId).ClientAuthorizationUpdateRequest(clientAuthorizationUpdateRequest).Execute()

Update Client Tokens



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
	clientId := "clientId_example" // string | A client ID. 
	clientAuthorizationUpdateRequest := *openapiclient.NewClientAuthorizationUpdateRequest("Subject_example") // ClientAuthorizationUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.ClientAuthorizationUpdateApi(context.Background(), serviceId, clientId).ClientAuthorizationUpdateRequest(clientAuthorizationUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ClientAuthorizationUpdateApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientAuthorizationUpdateApi`: ClientAuthorizationUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.ClientAuthorizationUpdateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 
**clientId** | **string** | A client ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientAuthorizationUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientAuthorizationUpdateRequest** | [**ClientAuthorizationUpdateRequest**](ClientAuthorizationUpdateRequest.md) |  | 

### Return type

[**ClientAuthorizationUpdateResponse**](ClientAuthorizationUpdateResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientCreateApi

> Client ClientCreateApi(ctx, serviceId).Client(client).Execute()

Create Client



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
	client := *openapiclient.NewClient() // Client |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.ClientCreateApi(context.Background(), serviceId).Client(client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ClientCreateApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientCreateApi`: Client
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.ClientCreateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientCreateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **client** | [**Client**](Client.md) |  | 

### Return type

[**Client**](Client.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientDeleteApi

> ClientDeleteApi(ctx, serviceId, clientId).Execute()

Delete Client ⚡



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
	clientId := "clientId_example" // string | The client ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientManagementAPI.ClientDeleteApi(context.Background(), serviceId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ClientDeleteApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 
**clientId** | **string** | The client ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientDeleteApiRequest struct via the builder pattern


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


## ClientExtensionRequestablesScopesDeleteApi

> ClientExtensionRequestablesScopesDeleteApi(ctx, serviceId, clientId).Execute()

Delete Requestable Scopes



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
	clientId := "clientId_example" // string | A client ID. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientManagementAPI.ClientExtensionRequestablesScopesDeleteApi(context.Background(), serviceId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ClientExtensionRequestablesScopesDeleteApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 
**clientId** | **string** | A client ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientExtensionRequestablesScopesDeleteApiRequest struct via the builder pattern


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


## ClientExtensionRequestablesScopesGetApi

> ClientExtensionRequestableScopesGetResponse ClientExtensionRequestablesScopesGetApi(ctx, serviceId, clientId).Execute()

Get Requestable Scopes



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
	clientId := "clientId_example" // string | A client ID. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.ClientExtensionRequestablesScopesGetApi(context.Background(), serviceId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ClientExtensionRequestablesScopesGetApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientExtensionRequestablesScopesGetApi`: ClientExtensionRequestableScopesGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.ClientExtensionRequestablesScopesGetApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 
**clientId** | **string** | A client ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientExtensionRequestablesScopesGetApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClientExtensionRequestableScopesGetResponse**](ClientExtensionRequestableScopesGetResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientExtensionRequestablesScopesUpdateApi

> ClientExtensionRequestableScopesUpdateResponse ClientExtensionRequestablesScopesUpdateApi(ctx, serviceId, clientId).ClientExtensionRequestableScopesUpdateRequest(clientExtensionRequestableScopesUpdateRequest).Execute()

Update Requestable Scopes



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
	clientId := "clientId_example" // string | A client ID. 
	clientExtensionRequestableScopesUpdateRequest := *openapiclient.NewClientExtensionRequestableScopesUpdateRequest() // ClientExtensionRequestableScopesUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.ClientExtensionRequestablesScopesUpdateApi(context.Background(), serviceId, clientId).ClientExtensionRequestableScopesUpdateRequest(clientExtensionRequestableScopesUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ClientExtensionRequestablesScopesUpdateApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientExtensionRequestablesScopesUpdateApi`: ClientExtensionRequestableScopesUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.ClientExtensionRequestablesScopesUpdateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 
**clientId** | **string** | A client ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientExtensionRequestablesScopesUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientExtensionRequestableScopesUpdateRequest** | [**ClientExtensionRequestableScopesUpdateRequest**](ClientExtensionRequestableScopesUpdateRequest.md) |  | 

### Return type

[**ClientExtensionRequestableScopesUpdateResponse**](ClientExtensionRequestableScopesUpdateResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientFlagUpdateApi

> ClientFlagUpdateResponse ClientFlagUpdateApi(ctx, serviceId, clientIdentifier).ClientFlagUpdateRequest(clientFlagUpdateRequest).Execute()

Update Client Lock



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
	clientIdentifier := "clientIdentifier_example" // string | A client ID.
	clientFlagUpdateRequest := *openapiclient.NewClientFlagUpdateRequest(false) // ClientFlagUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.ClientFlagUpdateApi(context.Background(), serviceId, clientIdentifier).ClientFlagUpdateRequest(clientFlagUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ClientFlagUpdateApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientFlagUpdateApi`: ClientFlagUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.ClientFlagUpdateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 
**clientIdentifier** | **string** | A client ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientFlagUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientFlagUpdateRequest** | [**ClientFlagUpdateRequest**](ClientFlagUpdateRequest.md) |  | 

### Return type

[**ClientFlagUpdateResponse**](ClientFlagUpdateResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientGetApi

> Client ClientGetApi(ctx, serviceId, clientId).Execute()

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
	clientId := "clientId_example" // string | A client ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.ClientGetApi(context.Background(), serviceId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ClientGetApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientGetApi`: Client
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.ClientGetApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 
**clientId** | **string** | A client ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientGetApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Client**](Client.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientGetListApi

> ClientGetListApi200Response ClientGetListApi(ctx, serviceId).Developer(developer).Start(start).End(end).Execute()

List Clients



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
	developer := "developer_example" // string | The developer of client applications. The default value is null. If this parameter is not set to `null`, client application of the specified developer are returned. Otherwise, all client applications that belong to the service are returned.  (optional)
	start := int32(56) // int32 | Start index (inclusive) of the result set. The default value is 0. Must not be a negative number. (optional)
	end := int32(56) // int32 | End index (exclusive) of the result set. The default value is 5. Must not be a negative number. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.ClientGetListApi(context.Background(), serviceId).Developer(developer).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ClientGetListApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientGetListApi`: ClientGetListApi200Response
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.ClientGetListApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientGetListApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **developer** | **string** | The developer of client applications. The default value is null. If this parameter is not set to &#x60;null&#x60;, client application of the specified developer are returned. Otherwise, all client applications that belong to the service are returned.  | 
 **start** | **int32** | Start index (inclusive) of the result set. The default value is 0. Must not be a negative number. | 
 **end** | **int32** | End index (exclusive) of the result set. The default value is 5. Must not be a negative number. | 

### Return type

[**ClientGetListApi200Response**](ClientGetListApi200Response.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientGrantedScopesDeleteApi

> ClientGrantedScopesDeleteResponse ClientGrantedScopesDeleteApi(ctx, serviceId, clientId, subject).Subject2(subject2).Execute()

Delete Granted Scopes



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
	clientId := "clientId_example" // string | A client ID. 
	subject := "subject_example" // string | Unique user ID of an end-user. 
	subject2 := "subject_example" // string | Unique user ID of an end-user. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.ClientGrantedScopesDeleteApi(context.Background(), serviceId, clientId, subject).Subject2(subject2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ClientGrantedScopesDeleteApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientGrantedScopesDeleteApi`: ClientGrantedScopesDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.ClientGrantedScopesDeleteApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 
**clientId** | **string** | A client ID.  | 
**subject** | **string** | Unique user ID of an end-user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientGrantedScopesDeleteApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **subject2** | **string** | Unique user ID of an end-user.  | 

### Return type

[**ClientGrantedScopesDeleteResponse**](ClientGrantedScopesDeleteResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientGrantedScopesGetApi

> ClientAuthorizationDeleteResponse ClientGrantedScopesGetApi(ctx, serviceId, clientId, subject).Subject2(subject2).Execute()

Get Granted Scopes



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
	clientId := "clientId_example" // string | A client ID. 
	subject := "subject_example" // string | Unique user ID of an end-user. 
	subject2 := "subject_example" // string | Unique user ID of an end-user. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.ClientGrantedScopesGetApi(context.Background(), serviceId, clientId, subject).Subject2(subject2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ClientGrantedScopesGetApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientGrantedScopesGetApi`: ClientAuthorizationDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.ClientGrantedScopesGetApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 
**clientId** | **string** | A client ID.  | 
**subject** | **string** | Unique user ID of an end-user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientGrantedScopesGetApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **subject2** | **string** | Unique user ID of an end-user.  | 

### Return type

[**ClientAuthorizationDeleteResponse**](ClientAuthorizationDeleteResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientSecretRefreshApi

> ClientSecretRefreshResponse ClientSecretRefreshApi(ctx, serviceId, clientIdentifier).Execute()

Rotate Client Secret



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
	clientIdentifier := "clientIdentifier_example" // string | The client ID or the client ID alias of a client. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.ClientSecretRefreshApi(context.Background(), serviceId, clientIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ClientSecretRefreshApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientSecretRefreshApi`: ClientSecretRefreshResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.ClientSecretRefreshApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 
**clientIdentifier** | **string** | The client ID or the client ID alias of a client.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientSecretRefreshApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClientSecretRefreshResponse**](ClientSecretRefreshResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientSecretUpdateApi

> ClientSecretUpdateResponse ClientSecretUpdateApi(ctx, serviceId, clientIdentifier).ClientSecretUpdateRequest(clientSecretUpdateRequest).Execute()

Update Client Secret



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
	clientIdentifier := "clientIdentifier_example" // string | The client ID or the client ID alias of a client. 
	clientSecretUpdateRequest := *openapiclient.NewClientSecretUpdateRequest("ClientSecret_example") // ClientSecretUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.ClientSecretUpdateApi(context.Background(), serviceId, clientIdentifier).ClientSecretUpdateRequest(clientSecretUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ClientSecretUpdateApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientSecretUpdateApi`: ClientSecretUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.ClientSecretUpdateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 
**clientIdentifier** | **string** | The client ID or the client ID alias of a client.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientSecretUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientSecretUpdateRequest** | [**ClientSecretUpdateRequest**](ClientSecretUpdateRequest.md) |  | 

### Return type

[**ClientSecretUpdateResponse**](ClientSecretUpdateResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientUpdateApi

> Client ClientUpdateApi(ctx, serviceId, clientId).Client(client).Execute()

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
	clientId := "clientId_example" // string | A client ID.
	client := *openapiclient.NewClient() // Client |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientManagementAPI.ClientUpdateApi(context.Background(), serviceId, clientId).Client(client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementAPI.ClientUpdateApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientUpdateApi`: Client
	fmt.Fprintf(os.Stdout, "Response from `ClientManagementAPI.ClientUpdateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 
**clientId** | **string** | A client ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **client** | [**Client**](Client.md) |  | 

### Return type

[**Client**](Client.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

