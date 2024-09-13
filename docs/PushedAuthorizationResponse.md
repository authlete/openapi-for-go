# PushedAuthorizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take. Any other value other than \&quot;CREATED\&quot; should be handled as unsuccessful result. | [optional] 
**RequestUri** | Pointer to **string** | The request_uri created to the client to be used as request_uri on the authorize call.  | [optional] 
**ResponseContent** | Pointer to **string** | The content that the authorization server implementation is to return to the client application.  | [optional] 
**ClientAuthMethod** | Pointer to **string** | The client authentication method that the client application declares that it uses at the token endpoint. This property corresponds to &#x60;token_endpoint_auth_method&#x60; in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**DpopNonce** | Pointer to **string** | Get the expected nonce value for DPoP proof JWT, which should be used as the value of the &#x60;DPoP-Nonce&#x60; HTTP header. | [optional] 

## Methods

### NewPushedAuthorizationResponse

`func NewPushedAuthorizationResponse() *PushedAuthorizationResponse`

NewPushedAuthorizationResponse instantiates a new PushedAuthorizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushedAuthorizationResponseWithDefaults

`func NewPushedAuthorizationResponseWithDefaults() *PushedAuthorizationResponse`

NewPushedAuthorizationResponseWithDefaults instantiates a new PushedAuthorizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *PushedAuthorizationResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *PushedAuthorizationResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *PushedAuthorizationResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *PushedAuthorizationResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *PushedAuthorizationResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *PushedAuthorizationResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *PushedAuthorizationResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *PushedAuthorizationResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *PushedAuthorizationResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PushedAuthorizationResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PushedAuthorizationResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PushedAuthorizationResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRequestUri

`func (o *PushedAuthorizationResponse) GetRequestUri() string`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *PushedAuthorizationResponse) GetRequestUriOk() (*string, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUri

`func (o *PushedAuthorizationResponse) SetRequestUri(v string)`

SetRequestUri sets RequestUri field to given value.

### HasRequestUri

`func (o *PushedAuthorizationResponse) HasRequestUri() bool`

HasRequestUri returns a boolean if a field has been set.

### GetResponseContent

`func (o *PushedAuthorizationResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *PushedAuthorizationResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *PushedAuthorizationResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *PushedAuthorizationResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetClientAuthMethod

`func (o *PushedAuthorizationResponse) GetClientAuthMethod() string`

GetClientAuthMethod returns the ClientAuthMethod field if non-nil, zero value otherwise.

### GetClientAuthMethodOk

`func (o *PushedAuthorizationResponse) GetClientAuthMethodOk() (*string, bool)`

GetClientAuthMethodOk returns a tuple with the ClientAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthMethod

`func (o *PushedAuthorizationResponse) SetClientAuthMethod(v string)`

SetClientAuthMethod sets ClientAuthMethod field to given value.

### HasClientAuthMethod

`func (o *PushedAuthorizationResponse) HasClientAuthMethod() bool`

HasClientAuthMethod returns a boolean if a field has been set.

### GetDpopNonce

`func (o *PushedAuthorizationResponse) GetDpopNonce() string`

GetDpopNonce returns the DpopNonce field if non-nil, zero value otherwise.

### GetDpopNonceOk

`func (o *PushedAuthorizationResponse) GetDpopNonceOk() (*string, bool)`

GetDpopNonceOk returns a tuple with the DpopNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpopNonce

`func (o *PushedAuthorizationResponse) SetDpopNonce(v string)`

SetDpopNonce sets DpopNonce field to given value.

### HasDpopNonce

`func (o *PushedAuthorizationResponse) HasDpopNonce() bool`

HasDpopNonce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


