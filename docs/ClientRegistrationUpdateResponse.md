# ClientRegistrationUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take.  | [optional] 
**ResponseContent** | Pointer to **string** | The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of &#x60;action&#x60; parameter.  | [optional] 
**Client** | Pointer to [**Client**](Client.md) |  | [optional] 

## Methods

### NewClientRegistrationUpdateResponse

`func NewClientRegistrationUpdateResponse() *ClientRegistrationUpdateResponse`

NewClientRegistrationUpdateResponse instantiates a new ClientRegistrationUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientRegistrationUpdateResponseWithDefaults

`func NewClientRegistrationUpdateResponseWithDefaults() *ClientRegistrationUpdateResponse`

NewClientRegistrationUpdateResponseWithDefaults instantiates a new ClientRegistrationUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *ClientRegistrationUpdateResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *ClientRegistrationUpdateResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *ClientRegistrationUpdateResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *ClientRegistrationUpdateResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *ClientRegistrationUpdateResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *ClientRegistrationUpdateResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *ClientRegistrationUpdateResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *ClientRegistrationUpdateResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *ClientRegistrationUpdateResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ClientRegistrationUpdateResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ClientRegistrationUpdateResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ClientRegistrationUpdateResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResponseContent

`func (o *ClientRegistrationUpdateResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *ClientRegistrationUpdateResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *ClientRegistrationUpdateResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *ClientRegistrationUpdateResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetClient

`func (o *ClientRegistrationUpdateResponse) GetClient() Client`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *ClientRegistrationUpdateResponse) GetClientOk() (*Client, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *ClientRegistrationUpdateResponse) SetClient(v Client)`

SetClient sets Client field to given value.

### HasClient

`func (o *ClientRegistrationUpdateResponse) HasClient() bool`

HasClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


