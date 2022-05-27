# \CIBAApi

All URIs are relative to *https://api.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackchannelAuthenticationApi**](CIBAApi.md#BackchannelAuthenticationApi) | **Post** /api/backchannel/authentication | /api/backchannel/authentication API
[**BackchannelAuthenticationCompleteApi**](CIBAApi.md#BackchannelAuthenticationCompleteApi) | **Post** /api/backchannel/authentication/complete | /api/backchannel/authentication/complete API
[**BackchannelAuthenticationFailApi**](CIBAApi.md#BackchannelAuthenticationFailApi) | **Post** /api/backchannel/authentication/fail | /api/backchannel/authentication/fail API
[**BackchannelAuthenticationIssueApi**](CIBAApi.md#BackchannelAuthenticationIssueApi) | **Post** /api/backchannel/authentication/issue | /api/backchannel/authentication/issue API



## BackchannelAuthenticationApi

> BackchannelAuthenticationResponse BackchannelAuthenticationApi(ctx).BackchannelAuthenticationRequest(backchannelAuthenticationRequest).Execute()

/api/backchannel/authentication API



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
    backchannelAuthenticationRequest := *openapiclient.NewBackchannelAuthenticationRequest("Parameters_example") // BackchannelAuthenticationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CIBAApi.BackchannelAuthenticationApi(context.Background()).BackchannelAuthenticationRequest(backchannelAuthenticationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CIBAApi.BackchannelAuthenticationApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackchannelAuthenticationApi`: BackchannelAuthenticationResponse
    fmt.Fprintf(os.Stdout, "Response from `CIBAApi.BackchannelAuthenticationApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackchannelAuthenticationApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backchannelAuthenticationRequest** | [**BackchannelAuthenticationRequest**](BackchannelAuthenticationRequest.md) |  | 

### Return type

[**BackchannelAuthenticationResponse**](BackchannelAuthenticationResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackchannelAuthenticationCompleteApi

> BackchannelAuthenticationCompleteResponse BackchannelAuthenticationCompleteApi(ctx).BackchannelAuthenticationCompleteRequest(backchannelAuthenticationCompleteRequest).Execute()

/api/backchannel/authentication/complete API



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
    backchannelAuthenticationCompleteRequest := *openapiclient.NewBackchannelAuthenticationCompleteRequest("Ticket_example", "Result_example", "Subject_example") // BackchannelAuthenticationCompleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CIBAApi.BackchannelAuthenticationCompleteApi(context.Background()).BackchannelAuthenticationCompleteRequest(backchannelAuthenticationCompleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CIBAApi.BackchannelAuthenticationCompleteApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackchannelAuthenticationCompleteApi`: BackchannelAuthenticationCompleteResponse
    fmt.Fprintf(os.Stdout, "Response from `CIBAApi.BackchannelAuthenticationCompleteApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackchannelAuthenticationCompleteApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backchannelAuthenticationCompleteRequest** | [**BackchannelAuthenticationCompleteRequest**](BackchannelAuthenticationCompleteRequest.md) |  | 

### Return type

[**BackchannelAuthenticationCompleteResponse**](BackchannelAuthenticationCompleteResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackchannelAuthenticationFailApi

> BackchannelAuthenticationFailResponse BackchannelAuthenticationFailApi(ctx).BackchannelAuthenticationFailRequest(backchannelAuthenticationFailRequest).Execute()

/api/backchannel/authentication/fail API



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
    backchannelAuthenticationFailRequest := *openapiclient.NewBackchannelAuthenticationFailRequest("Ticket_example", "Reason_example") // BackchannelAuthenticationFailRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CIBAApi.BackchannelAuthenticationFailApi(context.Background()).BackchannelAuthenticationFailRequest(backchannelAuthenticationFailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CIBAApi.BackchannelAuthenticationFailApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackchannelAuthenticationFailApi`: BackchannelAuthenticationFailResponse
    fmt.Fprintf(os.Stdout, "Response from `CIBAApi.BackchannelAuthenticationFailApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackchannelAuthenticationFailApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backchannelAuthenticationFailRequest** | [**BackchannelAuthenticationFailRequest**](BackchannelAuthenticationFailRequest.md) |  | 

### Return type

[**BackchannelAuthenticationFailResponse**](BackchannelAuthenticationFailResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackchannelAuthenticationIssueApi

> BackchannelAuthenticationIssueResponse BackchannelAuthenticationIssueApi(ctx).BackchannelAuthenticationIssueRequest(backchannelAuthenticationIssueRequest).Execute()

/api/backchannel/authentication/issue API



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
    backchannelAuthenticationIssueRequest := *openapiclient.NewBackchannelAuthenticationIssueRequest("Ticket_example") // BackchannelAuthenticationIssueRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CIBAApi.BackchannelAuthenticationIssueApi(context.Background()).BackchannelAuthenticationIssueRequest(backchannelAuthenticationIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CIBAApi.BackchannelAuthenticationIssueApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BackchannelAuthenticationIssueApi`: BackchannelAuthenticationIssueResponse
    fmt.Fprintf(os.Stdout, "Response from `CIBAApi.BackchannelAuthenticationIssueApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackchannelAuthenticationIssueApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backchannelAuthenticationIssueRequest** | [**BackchannelAuthenticationIssueRequest**](BackchannelAuthenticationIssueRequest.md) |  | 

### Return type

[**BackchannelAuthenticationIssueResponse**](BackchannelAuthenticationIssueResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

