# \DynamicClientRegistrationApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClientRegistrationApi**](DynamicClientRegistrationApi.md#ClientRegistrationApi) | **Post** /api/client/registration | /api/client/registration API
[**ClientRegistrationDeleteApi**](DynamicClientRegistrationApi.md#ClientRegistrationDeleteApi) | **Post** /api/client/registration/delete | /api/client/registration/delete API
[**ClientRegistrationGetApi**](DynamicClientRegistrationApi.md#ClientRegistrationGetApi) | **Post** /api/client/registration/get | /api/client/registration/get API
[**ClientRegistrationUpdateApi**](DynamicClientRegistrationApi.md#ClientRegistrationUpdateApi) | **Post** /api/client/registration/update | /api/client/registration/update API



## ClientRegistrationApi

> ClientRegistrationResponse ClientRegistrationApi(ctx).ClientRegistrationRequest(clientRegistrationRequest).Execute()

/api/client/registration API



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
    clientRegistrationRequest := *openapiclient.NewClientRegistrationRequest("Json_example") // ClientRegistrationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DynamicClientRegistrationApi.ClientRegistrationApi(context.Background()).ClientRegistrationRequest(clientRegistrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DynamicClientRegistrationApi.ClientRegistrationApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientRegistrationApi`: ClientRegistrationResponse
    fmt.Fprintf(os.Stdout, "Response from `DynamicClientRegistrationApi.ClientRegistrationApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClientRegistrationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientRegistrationRequest** | [**ClientRegistrationRequest**](ClientRegistrationRequest.md) |  | 

### Return type

[**ClientRegistrationResponse**](ClientRegistrationResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientRegistrationDeleteApi

> ClientRegistrationDeleteResponse ClientRegistrationDeleteApi(ctx).ClientRegistrationDeleteRequest(clientRegistrationDeleteRequest).Execute()

/api/client/registration/delete API



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
    clientRegistrationDeleteRequest := *openapiclient.NewClientRegistrationDeleteRequest("ClientId_example", "Token_example") // ClientRegistrationDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DynamicClientRegistrationApi.ClientRegistrationDeleteApi(context.Background()).ClientRegistrationDeleteRequest(clientRegistrationDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DynamicClientRegistrationApi.ClientRegistrationDeleteApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientRegistrationDeleteApi`: ClientRegistrationDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `DynamicClientRegistrationApi.ClientRegistrationDeleteApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClientRegistrationDeleteApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientRegistrationDeleteRequest** | [**ClientRegistrationDeleteRequest**](ClientRegistrationDeleteRequest.md) |  | 

### Return type

[**ClientRegistrationDeleteResponse**](ClientRegistrationDeleteResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientRegistrationGetApi

> ClientRegistrationResponse ClientRegistrationGetApi(ctx).ClientRegistrationRequest(clientRegistrationRequest).Execute()

/api/client/registration/get API



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
    clientRegistrationRequest := *openapiclient.NewClientRegistrationRequest("Json_example") // ClientRegistrationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DynamicClientRegistrationApi.ClientRegistrationGetApi(context.Background()).ClientRegistrationRequest(clientRegistrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DynamicClientRegistrationApi.ClientRegistrationGetApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientRegistrationGetApi`: ClientRegistrationResponse
    fmt.Fprintf(os.Stdout, "Response from `DynamicClientRegistrationApi.ClientRegistrationGetApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClientRegistrationGetApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientRegistrationRequest** | [**ClientRegistrationRequest**](ClientRegistrationRequest.md) |  | 

### Return type

[**ClientRegistrationResponse**](ClientRegistrationResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientRegistrationUpdateApi

> ClientRegistrationUpdateResponse ClientRegistrationUpdateApi(ctx).ClientRegistrationUpdateRequest(clientRegistrationUpdateRequest).Execute()

/api/client/registration/update API



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
    clientRegistrationUpdateRequest := *openapiclient.NewClientRegistrationUpdateRequest("ClientId_example", "Token_example", "Json_example") // ClientRegistrationUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DynamicClientRegistrationApi.ClientRegistrationUpdateApi(context.Background()).ClientRegistrationUpdateRequest(clientRegistrationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DynamicClientRegistrationApi.ClientRegistrationUpdateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClientRegistrationUpdateApi`: ClientRegistrationUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `DynamicClientRegistrationApi.ClientRegistrationUpdateApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClientRegistrationUpdateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientRegistrationUpdateRequest** | [**ClientRegistrationUpdateRequest**](ClientRegistrationUpdateRequest.md) |  | 

### Return type

[**ClientRegistrationUpdateResponse**](ClientRegistrationUpdateResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

