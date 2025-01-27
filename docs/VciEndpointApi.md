# \VciEndpointApi

All URIs are relative to *https://beta.authlete.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VciBatchIssueApi**](VciEndpointApi.md#VciBatchIssueApi) | **Post** /api/{serviceId}/vci/batch/issue | /api/{serviceId}/vci/batch/issue API
[**VciBatchParseApi**](VciEndpointApi.md#VciBatchParseApi) | **Post** /api/{serviceId}/vci/batch/parse | /api/{serviceId}/vci/batch/parse API
[**VciDeferredIssueApi**](VciEndpointApi.md#VciDeferredIssueApi) | **Post** /api/{serviceId}/vci/deferred/issue | /api/{serviceId}/vci/deferred/issue API
[**VciDeferredParseApi**](VciEndpointApi.md#VciDeferredParseApi) | **Post** /api/{serviceId}/vci/deferred/parse | /api/{serviceId}/vci/deferred/parse API
[**VciJwksApi**](VciEndpointApi.md#VciJwksApi) | **Post** /api/{serviceId}/vci/jwks | /api/{serviceId}/vci/jwks API
[**VciJwtissuerApi**](VciEndpointApi.md#VciJwtissuerApi) | **Post** /api/{serviceId}/vci/jwtissuer | /api/{serviceId}/vci/jwtissuer API
[**VciMetadataApi**](VciEndpointApi.md#VciMetadataApi) | **Post** /api/{serviceId}/vci/metadata | /api/{serviceId}/vci/metadata API
[**VciOfferCreateApi**](VciEndpointApi.md#VciOfferCreateApi) | **Post** /api/{serviceId}/vci/offer/create | /api/{serviceId}/vci/offer/create API
[**VciOfferInfoApi**](VciEndpointApi.md#VciOfferInfoApi) | **Post** /api/{serviceId}/vci/offer/info | /api/{serviceId}/vci/offer/info API
[**VciSingleIssueApi**](VciEndpointApi.md#VciSingleIssueApi) | **Post** /api/{serviceId}/vci/single/issue | /api/{serviceId}/vci/single/issue API
[**VciSingleParseApi**](VciEndpointApi.md#VciSingleParseApi) | **Post** /api/{serviceId}/vci/single/parse | /api/{serviceId}/vci/single/parse API



## VciBatchIssueApi

> VciBatchIssueResponse VciBatchIssueApi(ctx, serviceId).VciBatchIssueRequest(vciBatchIssueRequest).Execute()

/api/{serviceId}/vci/batch/issue API

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
    serviceId := "serviceId_example" // string | A service ID.
    vciBatchIssueRequest := *openapiclient.NewVciBatchIssueRequest() // VciBatchIssueRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VciEndpointApi.VciBatchIssueApi(context.Background(), serviceId).VciBatchIssueRequest(vciBatchIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VciEndpointApi.VciBatchIssueApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VciBatchIssueApi`: VciBatchIssueResponse
    fmt.Fprintf(os.Stdout, "Response from `VciEndpointApi.VciBatchIssueApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVciBatchIssueApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vciBatchIssueRequest** | [**VciBatchIssueRequest**](VciBatchIssueRequest.md) |  | 

### Return type

[**VciBatchIssueResponse**](VciBatchIssueResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VciBatchParseApi

> VciBatchParseResponse VciBatchParseApi(ctx, serviceId).VciBatchParseRequest(vciBatchParseRequest).Execute()

/api/{serviceId}/vci/batch/parse API

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
    serviceId := "serviceId_example" // string | A service ID.
    vciBatchParseRequest := *openapiclient.NewVciBatchParseRequest() // VciBatchParseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VciEndpointApi.VciBatchParseApi(context.Background(), serviceId).VciBatchParseRequest(vciBatchParseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VciEndpointApi.VciBatchParseApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VciBatchParseApi`: VciBatchParseResponse
    fmt.Fprintf(os.Stdout, "Response from `VciEndpointApi.VciBatchParseApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVciBatchParseApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vciBatchParseRequest** | [**VciBatchParseRequest**](VciBatchParseRequest.md) |  | 

### Return type

[**VciBatchParseResponse**](VciBatchParseResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VciDeferredIssueApi

> VciDeferredIssueResponse VciDeferredIssueApi(ctx, serviceId).VciDeferredIssueRequest(vciDeferredIssueRequest).Execute()

/api/{serviceId}/vci/deferred/issue API

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
    serviceId := "serviceId_example" // string | A service ID.
    vciDeferredIssueRequest := *openapiclient.NewVciDeferredIssueRequest() // VciDeferredIssueRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VciEndpointApi.VciDeferredIssueApi(context.Background(), serviceId).VciDeferredIssueRequest(vciDeferredIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VciEndpointApi.VciDeferredIssueApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VciDeferredIssueApi`: VciDeferredIssueResponse
    fmt.Fprintf(os.Stdout, "Response from `VciEndpointApi.VciDeferredIssueApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVciDeferredIssueApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vciDeferredIssueRequest** | [**VciDeferredIssueRequest**](VciDeferredIssueRequest.md) |  | 

### Return type

[**VciDeferredIssueResponse**](VciDeferredIssueResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VciDeferredParseApi

> VciDeferredParseResponse VciDeferredParseApi(ctx, serviceId).VciDeferredParseRequest(vciDeferredParseRequest).Execute()

/api/{serviceId}/vci/deferred/parse API

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
    serviceId := "serviceId_example" // string | A service ID.
    vciDeferredParseRequest := *openapiclient.NewVciDeferredParseRequest() // VciDeferredParseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VciEndpointApi.VciDeferredParseApi(context.Background(), serviceId).VciDeferredParseRequest(vciDeferredParseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VciEndpointApi.VciDeferredParseApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VciDeferredParseApi`: VciDeferredParseResponse
    fmt.Fprintf(os.Stdout, "Response from `VciEndpointApi.VciDeferredParseApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVciDeferredParseApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vciDeferredParseRequest** | [**VciDeferredParseRequest**](VciDeferredParseRequest.md) |  | 

### Return type

[**VciDeferredParseResponse**](VciDeferredParseResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VciJwksApi

> VciJwksResponse VciJwksApi(ctx, serviceId).VciJwksRequest(vciJwksRequest).Execute()

/api/{serviceId}/vci/jwks API

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
    serviceId := "serviceId_example" // string | A service ID.
    vciJwksRequest := *openapiclient.NewVciJwksRequest(false) // VciJwksRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VciEndpointApi.VciJwksApi(context.Background(), serviceId).VciJwksRequest(vciJwksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VciEndpointApi.VciJwksApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VciJwksApi`: VciJwksResponse
    fmt.Fprintf(os.Stdout, "Response from `VciEndpointApi.VciJwksApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVciJwksApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vciJwksRequest** | [**VciJwksRequest**](VciJwksRequest.md) |  | 

### Return type

[**VciJwksResponse**](VciJwksResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VciJwtissuerApi

> VciJwtissuerResponse VciJwtissuerApi(ctx, serviceId).VciJwtissuerRequest(vciJwtissuerRequest).Execute()

/api/{serviceId}/vci/jwtissuer API

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
    serviceId := "serviceId_example" // string | A service ID.
    vciJwtissuerRequest := *openapiclient.NewVciJwtissuerRequest(false) // VciJwtissuerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VciEndpointApi.VciJwtissuerApi(context.Background(), serviceId).VciJwtissuerRequest(vciJwtissuerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VciEndpointApi.VciJwtissuerApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VciJwtissuerApi`: VciJwtissuerResponse
    fmt.Fprintf(os.Stdout, "Response from `VciEndpointApi.VciJwtissuerApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVciJwtissuerApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vciJwtissuerRequest** | [**VciJwtissuerRequest**](VciJwtissuerRequest.md) |  | 

### Return type

[**VciJwtissuerResponse**](VciJwtissuerResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VciMetadataApi

> VciMetadataResponse VciMetadataApi(ctx, serviceId).VciMetadataRequest(vciMetadataRequest).Execute()

/api/{serviceId}/vci/metadata API

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
    serviceId := "serviceId_example" // string | A service ID.
    vciMetadataRequest := *openapiclient.NewVciMetadataRequest(false) // VciMetadataRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VciEndpointApi.VciMetadataApi(context.Background(), serviceId).VciMetadataRequest(vciMetadataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VciEndpointApi.VciMetadataApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VciMetadataApi`: VciMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `VciEndpointApi.VciMetadataApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVciMetadataApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vciMetadataRequest** | [**VciMetadataRequest**](VciMetadataRequest.md) |  | 

### Return type

[**VciMetadataResponse**](VciMetadataResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VciOfferCreateApi

> VciOfferCreateResponse VciOfferCreateApi(ctx, serviceId).VciOfferCreateRequest(vciOfferCreateRequest).Execute()

/api/{serviceId}/vci/offer/create API

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
    serviceId := "serviceId_example" // string | A service ID.
    vciOfferCreateRequest := *openapiclient.NewVciOfferCreateRequest() // VciOfferCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VciEndpointApi.VciOfferCreateApi(context.Background(), serviceId).VciOfferCreateRequest(vciOfferCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VciEndpointApi.VciOfferCreateApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VciOfferCreateApi`: VciOfferCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `VciEndpointApi.VciOfferCreateApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVciOfferCreateApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vciOfferCreateRequest** | [**VciOfferCreateRequest**](VciOfferCreateRequest.md) |  | 

### Return type

[**VciOfferCreateResponse**](VciOfferCreateResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VciOfferInfoApi

> VciOfferInfoResponse VciOfferInfoApi(ctx, serviceId).VciOfferInfoRequest(vciOfferInfoRequest).Execute()

/api/{serviceId}/vci/offer/info API

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
    serviceId := "serviceId_example" // string | A service ID.
    vciOfferInfoRequest := *openapiclient.NewVciOfferInfoRequest() // VciOfferInfoRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VciEndpointApi.VciOfferInfoApi(context.Background(), serviceId).VciOfferInfoRequest(vciOfferInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VciEndpointApi.VciOfferInfoApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VciOfferInfoApi`: VciOfferInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `VciEndpointApi.VciOfferInfoApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVciOfferInfoApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vciOfferInfoRequest** | [**VciOfferInfoRequest**](VciOfferInfoRequest.md) |  | 

### Return type

[**VciOfferInfoResponse**](VciOfferInfoResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VciSingleIssueApi

> VciSingleIssueResponse VciSingleIssueApi(ctx, serviceId).VciSingleIssueRequest(vciSingleIssueRequest).Execute()

/api/{serviceId}/vci/single/issue API

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
    serviceId := "serviceId_example" // string | A service ID.
    vciSingleIssueRequest := *openapiclient.NewVciSingleIssueRequest() // VciSingleIssueRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VciEndpointApi.VciSingleIssueApi(context.Background(), serviceId).VciSingleIssueRequest(vciSingleIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VciEndpointApi.VciSingleIssueApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VciSingleIssueApi`: VciSingleIssueResponse
    fmt.Fprintf(os.Stdout, "Response from `VciEndpointApi.VciSingleIssueApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVciSingleIssueApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vciSingleIssueRequest** | [**VciSingleIssueRequest**](VciSingleIssueRequest.md) |  | 

### Return type

[**VciSingleIssueResponse**](VciSingleIssueResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VciSingleParseApi

> VciSingleParseResponse VciSingleParseApi(ctx, serviceId).VciSingleParseRequest(vciSingleParseRequest).Execute()

/api/{serviceId}/vci/single/parse API

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
    serviceId := "serviceId_example" // string | A service ID.
    vciSingleParseRequest := *openapiclient.NewVciSingleParseRequest() // VciSingleParseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VciEndpointApi.VciSingleParseApi(context.Background(), serviceId).VciSingleParseRequest(vciSingleParseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VciEndpointApi.VciSingleParseApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VciSingleParseApi`: VciSingleParseResponse
    fmt.Fprintf(os.Stdout, "Response from `VciEndpointApi.VciSingleParseApi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | A service ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVciSingleParseApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vciSingleParseRequest** | [**VciSingleParseRequest**](VciSingleParseRequest.md) |  | 

### Return type

[**VciSingleParseResponse**](VciSingleParseResponse.md)

### Authorization

[UseService](../README.md#UseService)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

