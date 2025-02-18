# \FederationEndpointAPI

All URIs are relative to *https://us.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FederationConfigurationApi**](FederationEndpointAPI.md#FederationConfigurationApi) | **Post** /api/{serviceId}/federation/configuration | Process Entity Configuration Request
[**FederationRegistrationApi**](FederationEndpointAPI.md#FederationRegistrationApi) | **Post** /api/{serviceId}/federation/registration | Process Federation Registration Request



## FederationConfigurationApi

> FederationConfigurationResponse FederationConfigurationApi(ctx, serviceId).Body(body).Execute()

Process Entity Configuration Request



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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FederationEndpointAPI.FederationConfigurationApi(context.Background(), serviceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FederationEndpointAPI.FederationConfigurationApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FederationConfigurationApi`: FederationConfigurationResponse
	fmt.Fprintf(os.Stdout, "Response from `FederationEndpointAPI.FederationConfigurationApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFederationConfigurationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**FederationConfigurationResponse**](FederationConfigurationResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FederationRegistrationApi

> FederationRegistrationResponse FederationRegistrationApi(ctx, serviceId).FederationRegistrationRequest(federationRegistrationRequest).Execute()

Process Federation Registration Request



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
	federationRegistrationRequest := *openapiclient.NewFederationRegistrationRequest() // FederationRegistrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FederationEndpointAPI.FederationRegistrationApi(context.Background(), serviceId).FederationRegistrationRequest(federationRegistrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FederationEndpointAPI.FederationRegistrationApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FederationRegistrationApi`: FederationRegistrationResponse
	fmt.Fprintf(os.Stdout, "Response from `FederationEndpointAPI.FederationRegistrationApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFederationRegistrationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **federationRegistrationRequest** | [**FederationRegistrationRequest**](FederationRegistrationRequest.md) |  | 

### Return type

[**FederationRegistrationResponse**](FederationRegistrationResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

