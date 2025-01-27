# AuthorizationTicketInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to **string** | Information about the ticket. | [optional] 
**Action** | Pointer to **string** | The result of the &#x60;/auth/authorization/ticket/info&#x60; API call. | [optional] 
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 

## Methods

### NewAuthorizationTicketInfoResponse

`func NewAuthorizationTicketInfoResponse() *AuthorizationTicketInfoResponse`

NewAuthorizationTicketInfoResponse instantiates a new AuthorizationTicketInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationTicketInfoResponseWithDefaults

`func NewAuthorizationTicketInfoResponseWithDefaults() *AuthorizationTicketInfoResponse`

NewAuthorizationTicketInfoResponseWithDefaults instantiates a new AuthorizationTicketInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *AuthorizationTicketInfoResponse) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AuthorizationTicketInfoResponse) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AuthorizationTicketInfoResponse) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AuthorizationTicketInfoResponse) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetAction

`func (o *AuthorizationTicketInfoResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuthorizationTicketInfoResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuthorizationTicketInfoResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AuthorizationTicketInfoResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResultCode

`func (o *AuthorizationTicketInfoResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *AuthorizationTicketInfoResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *AuthorizationTicketInfoResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *AuthorizationTicketInfoResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *AuthorizationTicketInfoResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *AuthorizationTicketInfoResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *AuthorizationTicketInfoResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *AuthorizationTicketInfoResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


