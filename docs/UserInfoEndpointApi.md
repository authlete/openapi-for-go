# \UserInfoEndpointApi

All URIs are relative to *https://api.authlete.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthUserinfoApi**](UserInfoEndpointApi.md#AuthUserinfoApi) | **Post** /auth/userinfo | /auth/userinfo API
[**AuthUserinfoIssueApi**](UserInfoEndpointApi.md#AuthUserinfoIssueApi) | **Post** /auth/userinfo/issue | /auth/userinfo/issue API



## AuthUserinfoApi

> UserinfoResponse AuthUserinfoApi(ctx).UserinfoRequest(userinfoRequest).Execute()

/auth/userinfo API



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
    userinfoRequest := *openapiclient.NewUserinfoRequest("Token_example") // UserinfoRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserInfoEndpointApi.AuthUserinfoApi(context.Background()).UserinfoRequest(userinfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserInfoEndpointApi.AuthUserinfoApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthUserinfoApi`: UserinfoResponse
    fmt.Fprintf(os.Stdout, "Response from `UserInfoEndpointApi.AuthUserinfoApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthUserinfoApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userinfoRequest** | [**UserinfoRequest**](UserinfoRequest.md) |  | 

### Return type

[**UserinfoResponse**](UserinfoResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthUserinfoIssueApi

> UserinfoIssueResponse AuthUserinfoIssueApi(ctx).UserinfoIssueRequest(userinfoIssueRequest).Execute()

/auth/userinfo/issue API



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
    userinfoIssueRequest := *openapiclient.NewUserinfoIssueRequest("Token_example") // UserinfoIssueRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserInfoEndpointApi.AuthUserinfoIssueApi(context.Background()).UserinfoIssueRequest(userinfoIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserInfoEndpointApi.AuthUserinfoIssueApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthUserinfoIssueApi`: UserinfoIssueResponse
    fmt.Fprintf(os.Stdout, "Response from `UserInfoEndpointApi.AuthUserinfoIssueApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthUserinfoIssueApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userinfoIssueRequest** | [**UserinfoIssueRequest**](UserinfoIssueRequest.md) |  | 

### Return type

[**UserinfoIssueResponse**](UserinfoIssueResponse.md)

### Authorization

[ServiceCredentials](../README.md#ServiceCredentials)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

