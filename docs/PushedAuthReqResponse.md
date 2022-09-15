# PushedAuthReqResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take. Any other value other than \&quot;CREATED\&quot; should be handled as unsuccessful result. | [optional] 
**RequestUri** | Pointer to **string** | The request_uri created to the client to be used as request_uri on the authorize call.  | [optional] 
**ResponseContent** | Pointer to **string** | The content that the authorization server implementation is to return to the client application. | [optional] 

## Methods

### NewPushedAuthReqResponse

`func NewPushedAuthReqResponse() *PushedAuthReqResponse`

NewPushedAuthReqResponse instantiates a new PushedAuthReqResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushedAuthReqResponseWithDefaults

`func NewPushedAuthReqResponseWithDefaults() *PushedAuthReqResponse`

NewPushedAuthReqResponseWithDefaults instantiates a new PushedAuthReqResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *PushedAuthReqResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *PushedAuthReqResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *PushedAuthReqResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *PushedAuthReqResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *PushedAuthReqResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *PushedAuthReqResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *PushedAuthReqResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *PushedAuthReqResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *PushedAuthReqResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PushedAuthReqResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PushedAuthReqResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PushedAuthReqResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRequestUri

`func (o *PushedAuthReqResponse) GetRequestUri() string`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *PushedAuthReqResponse) GetRequestUriOk() (*string, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUri

`func (o *PushedAuthReqResponse) SetRequestUri(v string)`

SetRequestUri sets RequestUri field to given value.

### HasRequestUri

`func (o *PushedAuthReqResponse) HasRequestUri() bool`

HasRequestUri returns a boolean if a field has been set.

### GetResponseContent

`func (o *PushedAuthReqResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *PushedAuthReqResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *PushedAuthReqResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *PushedAuthReqResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


