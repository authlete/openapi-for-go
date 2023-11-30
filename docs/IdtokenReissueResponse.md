# IdtokenReissueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the implementation of the token endpoint should take. | [optional] 
**ResponseContent** | Pointer to **string** | The response content that can be used as the message body of the token response that should be returned from the token endpoint.  | [optional] 
**IdToken** | Pointer to **string** | The reissued ID token | [optional] 

## Methods

### NewIdtokenReissueResponse

`func NewIdtokenReissueResponse() *IdtokenReissueResponse`

NewIdtokenReissueResponse instantiates a new IdtokenReissueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdtokenReissueResponseWithDefaults

`func NewIdtokenReissueResponseWithDefaults() *IdtokenReissueResponse`

NewIdtokenReissueResponseWithDefaults instantiates a new IdtokenReissueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *IdtokenReissueResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *IdtokenReissueResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *IdtokenReissueResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *IdtokenReissueResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *IdtokenReissueResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *IdtokenReissueResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *IdtokenReissueResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *IdtokenReissueResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *IdtokenReissueResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IdtokenReissueResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IdtokenReissueResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *IdtokenReissueResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResponseContent

`func (o *IdtokenReissueResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *IdtokenReissueResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *IdtokenReissueResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *IdtokenReissueResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetIdToken

`func (o *IdtokenReissueResponse) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *IdtokenReissueResponse) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *IdtokenReissueResponse) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *IdtokenReissueResponse) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


