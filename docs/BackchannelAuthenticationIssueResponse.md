# BackchannelAuthenticationIssueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take. | [optional] 
**ResponseContent** | Pointer to **string** | The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of &#x60;action&#x60; parameter.  | [optional] 
**AuthReqId** | Pointer to **string** | The newly issued authentication request ID.  | [optional] 
**ExpiresIn** | Pointer to **int32** | The duration of the issued authentication request ID in seconds.  | [optional] 
**Interval** | Pointer to **int32** | The minimum amount of time in seconds that the client must wait for between polling requests to the token endpoint.  | [optional] 

## Methods

### NewBackchannelAuthenticationIssueResponse

`func NewBackchannelAuthenticationIssueResponse() *BackchannelAuthenticationIssueResponse`

NewBackchannelAuthenticationIssueResponse instantiates a new BackchannelAuthenticationIssueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackchannelAuthenticationIssueResponseWithDefaults

`func NewBackchannelAuthenticationIssueResponseWithDefaults() *BackchannelAuthenticationIssueResponse`

NewBackchannelAuthenticationIssueResponseWithDefaults instantiates a new BackchannelAuthenticationIssueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *BackchannelAuthenticationIssueResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *BackchannelAuthenticationIssueResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *BackchannelAuthenticationIssueResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *BackchannelAuthenticationIssueResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *BackchannelAuthenticationIssueResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *BackchannelAuthenticationIssueResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *BackchannelAuthenticationIssueResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *BackchannelAuthenticationIssueResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *BackchannelAuthenticationIssueResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BackchannelAuthenticationIssueResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BackchannelAuthenticationIssueResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BackchannelAuthenticationIssueResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResponseContent

`func (o *BackchannelAuthenticationIssueResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *BackchannelAuthenticationIssueResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *BackchannelAuthenticationIssueResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *BackchannelAuthenticationIssueResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetAuthReqId

`func (o *BackchannelAuthenticationIssueResponse) GetAuthReqId() string`

GetAuthReqId returns the AuthReqId field if non-nil, zero value otherwise.

### GetAuthReqIdOk

`func (o *BackchannelAuthenticationIssueResponse) GetAuthReqIdOk() (*string, bool)`

GetAuthReqIdOk returns a tuple with the AuthReqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthReqId

`func (o *BackchannelAuthenticationIssueResponse) SetAuthReqId(v string)`

SetAuthReqId sets AuthReqId field to given value.

### HasAuthReqId

`func (o *BackchannelAuthenticationIssueResponse) HasAuthReqId() bool`

HasAuthReqId returns a boolean if a field has been set.

### GetExpiresIn

`func (o *BackchannelAuthenticationIssueResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *BackchannelAuthenticationIssueResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *BackchannelAuthenticationIssueResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *BackchannelAuthenticationIssueResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetInterval

`func (o *BackchannelAuthenticationIssueResponse) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *BackchannelAuthenticationIssueResponse) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *BackchannelAuthenticationIssueResponse) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *BackchannelAuthenticationIssueResponse) HasInterval() bool`

HasInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


