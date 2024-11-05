# VciDeferredParseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the deferred credential endpoint should take. | [optional] 
**ResponseContent** | Pointer to **string** | The content of the response to the request sender. | [optional] 
**Info** | Pointer to [**CredentialRequestInfo**](CredentialRequestInfo.md) |  | [optional] 

## Methods

### NewVciDeferredParseResponse

`func NewVciDeferredParseResponse() *VciDeferredParseResponse`

NewVciDeferredParseResponse instantiates a new VciDeferredParseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVciDeferredParseResponseWithDefaults

`func NewVciDeferredParseResponseWithDefaults() *VciDeferredParseResponse`

NewVciDeferredParseResponseWithDefaults instantiates a new VciDeferredParseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *VciDeferredParseResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *VciDeferredParseResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *VciDeferredParseResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *VciDeferredParseResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *VciDeferredParseResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *VciDeferredParseResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *VciDeferredParseResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *VciDeferredParseResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *VciDeferredParseResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VciDeferredParseResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VciDeferredParseResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *VciDeferredParseResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResponseContent

`func (o *VciDeferredParseResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *VciDeferredParseResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *VciDeferredParseResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *VciDeferredParseResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetInfo

`func (o *VciDeferredParseResponse) GetInfo() CredentialRequestInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *VciDeferredParseResponse) GetInfoOk() (*CredentialRequestInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *VciDeferredParseResponse) SetInfo(v CredentialRequestInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *VciDeferredParseResponse) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


