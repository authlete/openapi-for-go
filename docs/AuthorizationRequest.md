# AuthorizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | **string** | OAuth 2.0 authorization request parameters which are the request parameters that the OAuth 2.0 authorization endpoint of the authorization server implementation received from the client application.  The value of parameters is either (1) the entire query string when the HTTP method of the request from the client application is &#x60;GET&#x60; or (2) the entire entity body (which is formatted in &#x60;application/x-www-form-urlencoded&#x60;) when the HTTP method of the request from the client application is &#x60;POST&#x60;. | 

## Methods

### NewAuthorizationRequest

`func NewAuthorizationRequest(parameters string, ) *AuthorizationRequest`

NewAuthorizationRequest instantiates a new AuthorizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationRequestWithDefaults

`func NewAuthorizationRequestWithDefaults() *AuthorizationRequest`

NewAuthorizationRequestWithDefaults instantiates a new AuthorizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *AuthorizationRequest) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AuthorizationRequest) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AuthorizationRequest) SetParameters(v string)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


