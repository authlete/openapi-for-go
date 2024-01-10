# \IDTokenEndpointApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdtokenReissueApi**](IDTokenEndpointApi.md#IdtokenReissueApi) | **Post** /api/idtoken/reissue | /api/idtoken/reissue API



## IdtokenReissueApi

> IdtokenReissueResponse IdtokenReissueApi(ctx).IdtokenReissueRequest(idtokenReissueRequest).Execute()

/api/idtoken/reissue API



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
    idtokenReissueRequest := *openapiclient.NewIdtokenReissueRequest("AccessToken_example", "RefreshToken_example") // IdtokenReissueRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IDTokenEndpointApi.IdtokenReissueApi(context.Background()).IdtokenReissueRequest(idtokenReissueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IDTokenEndpointApi.IdtokenReissueApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdtokenReissueApi`: IdtokenReissueResponse
    fmt.Fprintf(os.Stdout, "Response from `IDTokenEndpointApi.IdtokenReissueApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdtokenReissueApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idtokenReissueRequest** | [**IdtokenReissueRequest**](IdtokenReissueRequest.md) |  | 

### Return type

[**IdtokenReissueResponse**](IdtokenReissueResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

