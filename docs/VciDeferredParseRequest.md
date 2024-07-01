# VciDeferredParseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | The access token that came along with the deferred credential request. | [optional] 
**RequestContent** | Pointer to **string** | The message body of the deferred credential request. | [optional] 

## Methods

### NewVciDeferredParseRequest

`func NewVciDeferredParseRequest() *VciDeferredParseRequest`

NewVciDeferredParseRequest instantiates a new VciDeferredParseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVciDeferredParseRequestWithDefaults

`func NewVciDeferredParseRequestWithDefaults() *VciDeferredParseRequest`

NewVciDeferredParseRequestWithDefaults instantiates a new VciDeferredParseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *VciDeferredParseRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *VciDeferredParseRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *VciDeferredParseRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *VciDeferredParseRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRequestContent

`func (o *VciDeferredParseRequest) GetRequestContent() string`

GetRequestContent returns the RequestContent field if non-nil, zero value otherwise.

### GetRequestContentOk

`func (o *VciDeferredParseRequest) GetRequestContentOk() (*string, bool)`

GetRequestContentOk returns a tuple with the RequestContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContent

`func (o *VciDeferredParseRequest) SetRequestContent(v string)`

SetRequestContent sets RequestContent field to given value.

### HasRequestContent

`func (o *VciDeferredParseRequest) HasRequestContent() bool`

HasRequestContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


