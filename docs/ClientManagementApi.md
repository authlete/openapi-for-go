# \ClientManagementApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClientAuthorizationDeleteApi**](ClientManagementApi.md#ClientAuthorizationDeleteApi) | **Delete** /api/client/authorization/delete/{clientId} | /api/client/authorization/delete/{clientId}/{subject} API
[**ClientAuthorizationGetListApi**](ClientManagementApi.md#ClientAuthorizationGetListApi) | **Get** /api/client/authorization/get/list | /api/client/authorization/get/list/{subject} API
[**ClientAuthorizationUpdateApi**](ClientManagementApi.md#ClientAuthorizationUpdateApi) | **Post** /api/client/authorization/update/{clientId} | /api/client/authorization/update/{clientId} API
[**ClientCreateApi**](ClientManagementApi.md#ClientCreateApi) | **Post** /api/client/create | /api/client/create API
[**ClientDeleteApi**](ClientManagementApi.md#ClientDeleteApi) | **Delete** /api/client/delete/{clientId} | /api/client/delete/{clientId} API
[**ClientFlagUpdateApi**](ClientManagementApi.md#ClientFlagUpdateApi) | **Post** /api/client/lock_flag/update/{clientIdentifier} | /api/client/lock_flag/update/{clientIdentifier} API
[**ClientGetApi**](ClientManagementApi.md#ClientGetApi) | **Get** /api/client/get/{clientId} | /api/client/get/{clientId} API
[**ClientGetListApi**](ClientManagementApi.md#ClientGetListApi) | **Get** /api/client/get/list | /api/client/get/list API
[**ClientGrantedScopesDeleteApi**](ClientManagementApi.md#ClientGrantedScopesDeleteApi) | **Delete** /api/client/granted_scopes/delete/{clientId} | /api/client/granted_scopes/delete/{clientId}/{subject} API
[**ClientGrantedScopesGetApi**](ClientManagementApi.md#ClientGrantedScopesGetApi) | **Get** /api/client/granted_scopes/get/{clientId} | /api/client/granted_scopes/get/{clientId}/{subject} API
[**ClientSecretRefreshApi**](ClientManagementApi.md#ClientSecretRefreshApi) | **Get** /api/client/secret/refresh/{clientIdentifier} | /api/client/secret/refresh API
[**ClientSecretUpdateApi**](ClientManagementApi.md#ClientSecretUpdateApi) | **Post** /api/client/secret/update/{clientIdentifier} | /api/client/secret/update API
[**ClientUpdateApi**](ClientManagementApi.md#ClientUpdateApi) | **Post** /api/client/update/{clientId} | /api/client/update/{clientId} API



## ClientAuthorizationDeleteApi

> ClientAuthorizationDeleteResponse ClientAuthorizationDeleteApi(ctx, clientId, subject).Subject2(subject2).Execute()

/api/client/authorization/delete/{clientId}/{subject} API



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clientId := "clientId_example" // string | A client ID. 
    subject := "subject_example" // string | Unique user ID of an end-user. 
    subject2 := "subject_example" // string | Unique user ID of an end-user. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientManagementApi.ClientAuthorizationDeleteApi(context.Background(), clientId, subject).Subject2(subject2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementApi.ClientAuthorizationDeleteApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientAuthorizationDeleteApi`: ClientAuthorizationDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientManagementApi.ClientAuthorizationDeleteApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientAuthorizationGetListApi

> ClientAuthorizationGetListResponse ClientAuthorizationGetListApi(ctx, subject).Subject2(subject2).Developer(developer).Start(start).End(end).Execute()

/api/client/authorization/get/list/{subject} API



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subject := "subject_example" // string | Unique user ID of an end-user. 
    subject2 := "subject_example" // string | Unique user ID of an end-user. 
    developer := "developer_example" // string | Unique ID of a client developer.  (optional)
    start := int32(56) // int32 | Start index of search results (inclusive). The default value is 0. (optional)
    end := int32(56) // int32 | End index of search results (exclusive). The default value is 5.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientManagementApi.ClientAuthorizationGetListApi(context.Background(), subject).Subject2(subject2).Developer(developer).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementApi.ClientAuthorizationGetListApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientAuthorizationGetListApi`: ClientAuthorizationGetListResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientManagementApi.ClientAuthorizationGetListApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientAuthorizationUpdateApi

> ClientAuthorizationUpdateResponse ClientAuthorizationUpdateApi(ctx, clientId).ClientAuthorizationUpdateRequest(clientAuthorizationUpdateRequest).Execute()

/api/client/authorization/update/{clientId} API



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clientId := "clientId_example" // string | A client ID. 
    clientAuthorizationUpdateRequest := *openapiclient.NewClientAuthorizationUpdateRequest("Subject_example") // ClientAuthorizationUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientManagementApi.ClientAuthorizationUpdateApi(context.Background(), clientId).ClientAuthorizationUpdateRequest(clientAuthorizationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementApi.ClientAuthorizationUpdateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientAuthorizationUpdateApi`: ClientAuthorizationUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientManagementApi.ClientAuthorizationUpdateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | A client ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientAuthorizationUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientAuthorizationUpdateRequest** | [**ClientAuthorizationUpdateRequest**](ClientAuthorizationUpdateRequest.md) |  | 

### Return type

[**ClientAuthorizationUpdateResponse**](ClientAuthorizationUpdateResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientCreateApi

> Client ClientCreateApi(ctx).Client(client).Execute()

/api/client/create API



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    client := *openapiclient.NewClient() // Client |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientManagementApi.ClientCreateApi(context.Background()).Client(client).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementApi.ClientCreateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientCreateApi`: Client
    fmt.Fprintf(os.Stdout, "Response from `ClientManagementApi.ClientCreateApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClientCreateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **client** | [**Client**](Client.md) |  | 

### Return type

[**Client**](Client.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientDeleteApi

> ClientDeleteApi(ctx, clientId).Execute()

/api/client/delete/{clientId} API



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clientId := "clientId_example" // string | The client ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientManagementApi.ClientDeleteApi(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementApi.ClientDeleteApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | The client ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientDeleteApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientFlagUpdateApi

> ClientFlagUpdateResponse ClientFlagUpdateApi(ctx, clientIdentifier).ClientFlagUpdateRequest(clientFlagUpdateRequest).Execute()

/api/client/lock_flag/update/{clientIdentifier} API



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clientIdentifier := "clientIdentifier_example" // string | A client ID.
    clientFlagUpdateRequest := *openapiclient.NewClientFlagUpdateRequest(false) // ClientFlagUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientManagementApi.ClientFlagUpdateApi(context.Background(), clientIdentifier).ClientFlagUpdateRequest(clientFlagUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementApi.ClientFlagUpdateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientFlagUpdateApi`: ClientFlagUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientManagementApi.ClientFlagUpdateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientIdentifier** | **string** | A client ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientFlagUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientFlagUpdateRequest** | [**ClientFlagUpdateRequest**](ClientFlagUpdateRequest.md) |  | 

### Return type

[**ClientFlagUpdateResponse**](ClientFlagUpdateResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientGetApi

> Client ClientGetApi(ctx, clientId).Execute()

/api/client/get/{clientId} API



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clientId := "clientId_example" // string | A client ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientManagementApi.ClientGetApi(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementApi.ClientGetApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientGetApi`: Client
    fmt.Fprintf(os.Stdout, "Response from `ClientManagementApi.ClientGetApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | A client ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientGetApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Client**](Client.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientGetListApi

> ClientGetListResponse ClientGetListApi(ctx).Developer(developer).Start(start).End(end).Execute()

/api/client/get/list API



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    developer := "developer_example" // string | The developer of client applications. The default value is null. If this parameter is not set to `null`, client application of the specified developer are returned. Otherwise, all client applications that belong to the service are returned.  (optional)
    start := int32(56) // int32 | Start index (inclusive) of the result set. The default value is 0. Must not be a negative number. (optional)
    end := int32(56) // int32 | End index (exclusive) of the result set. The default value is 5. Must not be a negative number. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientManagementApi.ClientGetListApi(context.Background()).Developer(developer).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementApi.ClientGetListApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientGetListApi`: ClientGetListResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientManagementApi.ClientGetListApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClientGetListApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **developer** | **string** | The developer of client applications. The default value is null. If this parameter is not set to &#x60;null&#x60;, client application of the specified developer are returned. Otherwise, all client applications that belong to the service are returned.  | 
 **start** | **int32** | Start index (inclusive) of the result set. The default value is 0. Must not be a negative number. | 
 **end** | **int32** | End index (exclusive) of the result set. The default value is 5. Must not be a negative number. | 

### Return type

[**ClientGetListResponse**](ClientGetListResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientGrantedScopesDeleteApi

> ClientGrantedScopesDeleteResponse ClientGrantedScopesDeleteApi(ctx, clientId, subject).Subject2(subject2).Execute()

/api/client/granted_scopes/delete/{clientId}/{subject} API



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clientId := "clientId_example" // string | A client ID. 
    subject := "subject_example" // string | Unique user ID of an end-user. 
    subject2 := "subject_example" // string | Unique user ID of an end-user. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientManagementApi.ClientGrantedScopesDeleteApi(context.Background(), clientId, subject).Subject2(subject2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementApi.ClientGrantedScopesDeleteApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientGrantedScopesDeleteApi`: ClientGrantedScopesDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientManagementApi.ClientGrantedScopesDeleteApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientGrantedScopesGetApi

> ClientGrantedScopesGetResponse ClientGrantedScopesGetApi(ctx, clientId, subject).Subject2(subject2).Execute()

/api/client/granted_scopes/get/{clientId}/{subject} API



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clientId := "clientId_example" // string | A client ID. 
    subject := "subject_example" // string | Unique user ID of an end-user. 
    subject2 := "subject_example" // string | Unique user ID of an end-user. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientManagementApi.ClientGrantedScopesGetApi(context.Background(), clientId, subject).Subject2(subject2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementApi.ClientGrantedScopesGetApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientGrantedScopesGetApi`: ClientGrantedScopesGetResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientManagementApi.ClientGrantedScopesGetApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | A client ID.  | 
**subject** | **string** | Unique user ID of an end-user.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientGrantedScopesGetApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **subject2** | **string** | Unique user ID of an end-user.  | 

### Return type

[**ClientGrantedScopesGetResponse**](ClientGrantedScopesGetResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientSecretRefreshApi

> ClientSecretRefreshResponse ClientSecretRefreshApi(ctx, clientIdentifier).Execute()

/api/client/secret/refresh API



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clientIdentifier := "clientIdentifier_example" // string | The client ID or the client ID alias of a client. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientManagementApi.ClientSecretRefreshApi(context.Background(), clientIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementApi.ClientSecretRefreshApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientSecretRefreshApi`: ClientSecretRefreshResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientManagementApi.ClientSecretRefreshApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientIdentifier** | **string** | The client ID or the client ID alias of a client.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientSecretRefreshApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientSecretRefreshResponse**](ClientSecretRefreshResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientSecretUpdateApi

> ClientSecretUpdateResponse ClientSecretUpdateApi(ctx, clientIdentifier).ClientSecretUpdateRequest(clientSecretUpdateRequest).Execute()

/api/client/secret/update API



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clientIdentifier := "clientIdentifier_example" // string | The client ID or the client ID alias of a client. 
    clientSecretUpdateRequest := *openapiclient.NewClientSecretUpdateRequest("ClientSecret_example") // ClientSecretUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientManagementApi.ClientSecretUpdateApi(context.Background(), clientIdentifier).ClientSecretUpdateRequest(clientSecretUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementApi.ClientSecretUpdateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientSecretUpdateApi`: ClientSecretUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientManagementApi.ClientSecretUpdateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientIdentifier** | **string** | The client ID or the client ID alias of a client.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientSecretUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientSecretUpdateRequest** | [**ClientSecretUpdateRequest**](ClientSecretUpdateRequest.md) |  | 

### Return type

[**ClientSecretUpdateResponse**](ClientSecretUpdateResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientUpdateApi

> Client ClientUpdateApi(ctx, clientId).Client(client).Execute()

/api/client/update/{clientId} API



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    clientId := "clientId_example" // string | A client ID.
    client := *openapiclient.NewClient() // Client |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientManagementApi.ClientUpdateApi(context.Background(), clientId).Client(client).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientManagementApi.ClientUpdateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientUpdateApi`: Client
    fmt.Fprintf(os.Stdout, "Response from `ClientManagementApi.ClientUpdateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | A client ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **client** | [**Client**](Client.md) |  | 

### Return type

[**Client**](Client.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

