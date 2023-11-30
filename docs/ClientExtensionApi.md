# \ClientExtensionApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClientExtensionRequestablesScopesDeleteApi**](ClientExtensionApi.md#ClientExtensionRequestablesScopesDeleteApi) | **Delete** /api/client/extension/requestable_scopes/delete/{clientId} | /api/client/extension/requestable_scopes/delete/{clientId} API
[**ClientExtensionRequestablesScopesGetApi**](ClientExtensionApi.md#ClientExtensionRequestablesScopesGetApi) | **Get** /api/client/extension/requestable_scopes/get/{clientId} | /api/client/extension/requestable_scopes/get/{clientId} API
[**ClientExtensionRequestablesScopesUpdateApi**](ClientExtensionApi.md#ClientExtensionRequestablesScopesUpdateApi) | **Put** /api/client/extension/requestable_scopes/update/{clientId} | /api/client/extension/requestable_scopes/update/{clientId} API



## ClientExtensionRequestablesScopesDeleteApi

> ClientExtensionRequestablesScopesDeleteApi(ctx, clientId).Execute()

/api/client/extension/requestable_scopes/delete/{clientId} API



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
    resp, r, err := apiClient.ClientExtensionApi.ClientExtensionRequestablesScopesDeleteApi(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientExtensionApi.ClientExtensionRequestablesScopesDeleteApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | A client ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientExtensionRequestablesScopesDeleteApiRequest struct via the builder pattern


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


## ClientExtensionRequestablesScopesGetApi

> ClientExtensionRequestableScopesGetResponse ClientExtensionRequestablesScopesGetApi(ctx, clientId).Execute()

/api/client/extension/requestable_scopes/get/{clientId} API



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
    resp, r, err := apiClient.ClientExtensionApi.ClientExtensionRequestablesScopesGetApi(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientExtensionApi.ClientExtensionRequestablesScopesGetApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientExtensionRequestablesScopesGetApi`: ClientExtensionRequestableScopesGetResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientExtensionApi.ClientExtensionRequestablesScopesGetApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | A client ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientExtensionRequestablesScopesGetApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientExtensionRequestableScopesGetResponse**](ClientExtensionRequestableScopesGetResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientExtensionRequestablesScopesUpdateApi

> ClientExtensionRequestableScopesUpdateResponse ClientExtensionRequestablesScopesUpdateApi(ctx, clientId).ClientExtensionRequestableScopesUpdateRequest(clientExtensionRequestableScopesUpdateRequest).Execute()

/api/client/extension/requestable_scopes/update/{clientId} API



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
    clientExtensionRequestableScopesUpdateRequest := *openapiclient.NewClientExtensionRequestableScopesUpdateRequest() // ClientExtensionRequestableScopesUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientExtensionApi.ClientExtensionRequestablesScopesUpdateApi(context.Background(), clientId).ClientExtensionRequestableScopesUpdateRequest(clientExtensionRequestableScopesUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientExtensionApi.ClientExtensionRequestablesScopesUpdateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientExtensionRequestablesScopesUpdateApi`: ClientExtensionRequestableScopesUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `ClientExtensionApi.ClientExtensionRequestablesScopesUpdateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | A client ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClientExtensionRequestablesScopesUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientExtensionRequestableScopesUpdateRequest** | [**ClientExtensionRequestableScopesUpdateRequest**](ClientExtensionRequestableScopesUpdateRequest.md) |  | 

### Return type

[**ClientExtensionRequestableScopesUpdateResponse**](ClientExtensionRequestableScopesUpdateResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

