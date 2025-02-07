# \RevocationEndpointAPI

All URIs are relative to *https://us.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthRevocationApi**](RevocationEndpointAPI.md#AuthRevocationApi) | **Post** /api/{serviceId}/auth/revocation | Process Revocation Request



## AuthRevocationApi

> RevocationResponse AuthRevocationApi(ctx, serviceId).RevocationRequest(revocationRequest).Execute()

Process Revocation Request



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
	revocationRequest := *openapiclient.NewRevocationRequest("Parameters_example") // RevocationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevocationEndpointAPI.AuthRevocationApi(context.Background(), serviceId).RevocationRequest(revocationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevocationEndpointAPI.AuthRevocationApi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthRevocationApi`: RevocationResponse
	fmt.Fprintf(os.Stdout, "Response from `RevocationEndpointAPI.AuthRevocationApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthRevocationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revocationRequest** | [**RevocationRequest**](RevocationRequest.md) |  | 

### Return type

[**RevocationResponse**](RevocationResponse.md)

### Authorization

[bearer](../README.md#bearer), [authlete](../README.md#authlete)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

