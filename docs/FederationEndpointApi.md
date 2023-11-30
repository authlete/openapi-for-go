# \FederationEndpointApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FederationConfigurationApi**](FederationEndpointApi.md#FederationConfigurationApi) | **Post** /api/federation/configuration | /api/federation/configuration API
[**FederationRegistrationApi**](FederationEndpointApi.md#FederationRegistrationApi) | **Post** /api/federation/registration | /api/federation/registration API



## FederationConfigurationApi

> FederationConfigurationResponse FederationConfigurationApi(ctx).Body(body).Execute()

/api/federation/configuration API



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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederationEndpointApi.FederationConfigurationApi(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationEndpointApi.FederationConfigurationApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FederationConfigurationApi`: FederationConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `FederationEndpointApi.FederationConfigurationApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFederationConfigurationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**FederationConfigurationResponse**](FederationConfigurationResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FederationRegistrationApi

> FederationRegistrationResponse FederationRegistrationApi(ctx).FederationRegistrationRequest(federationRegistrationRequest).Execute()

/api/federation/registration API



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
    federationRegistrationRequest := *openapiclient.NewFederationRegistrationRequest() // FederationRegistrationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FederationEndpointApi.FederationRegistrationApi(context.Background()).FederationRegistrationRequest(federationRegistrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FederationEndpointApi.FederationRegistrationApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FederationRegistrationApi`: FederationRegistrationResponse
    fmt.Fprintf(os.Stdout, "Response from `FederationEndpointApi.FederationRegistrationApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFederationRegistrationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **federationRegistrationRequest** | [**FederationRegistrationRequest**](FederationRegistrationRequest.md) |  | 

### Return type

[**FederationRegistrationResponse**](FederationRegistrationResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

