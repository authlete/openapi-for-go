# \IntrospectionEndpointApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthIntrospectionApi**](IntrospectionEndpointApi.md#AuthIntrospectionApi) | **Post** /api/auth/introspection | /api/auth/introspection API
[**AuthIntrospectionStandardApi**](IntrospectionEndpointApi.md#AuthIntrospectionStandardApi) | **Post** /api/auth/introspection/standard | /api/auth/intraspection/standard API



## AuthIntrospectionApi

> IntrospectionResponse AuthIntrospectionApi(ctx).IntrospectionRequest(introspectionRequest).Execute()

/api/auth/introspection API



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
    introspectionRequest := *openapiclient.NewIntrospectionRequest("Token_example") // IntrospectionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntrospectionEndpointApi.AuthIntrospectionApi(context.Background()).IntrospectionRequest(introspectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntrospectionEndpointApi.AuthIntrospectionApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthIntrospectionApi`: IntrospectionResponse
    fmt.Fprintf(os.Stdout, "Response from `IntrospectionEndpointApi.AuthIntrospectionApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthIntrospectionApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **introspectionRequest** | [**IntrospectionRequest**](IntrospectionRequest.md) |  | 

### Return type

[**IntrospectionResponse**](IntrospectionResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthIntrospectionStandardApi

> StandardIntrospectionResponse AuthIntrospectionStandardApi(ctx).StandardIntrospectionRequest(standardIntrospectionRequest).Execute()

/api/auth/intraspection/standard API



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
    standardIntrospectionRequest := *openapiclient.NewStandardIntrospectionRequest("Parameters_example") // StandardIntrospectionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntrospectionEndpointApi.AuthIntrospectionStandardApi(context.Background()).StandardIntrospectionRequest(standardIntrospectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntrospectionEndpointApi.AuthIntrospectionStandardApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthIntrospectionStandardApi`: StandardIntrospectionResponse
    fmt.Fprintf(os.Stdout, "Response from `IntrospectionEndpointApi.AuthIntrospectionStandardApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthIntrospectionStandardApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **standardIntrospectionRequest** | [**StandardIntrospectionRequest**](StandardIntrospectionRequest.md) |  | 

### Return type

[**StandardIntrospectionResponse**](StandardIntrospectionResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

