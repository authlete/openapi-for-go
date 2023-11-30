# TokenRevokeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Count** | Pointer to **int32** | The number of tokens revoked | [optional] 

## Methods

### NewTokenRevokeResponse

`func NewTokenRevokeResponse() *TokenRevokeResponse`

NewTokenRevokeResponse instantiates a new TokenRevokeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRevokeResponseWithDefaults

`func NewTokenRevokeResponseWithDefaults() *TokenRevokeResponse`

NewTokenRevokeResponseWithDefaults instantiates a new TokenRevokeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *TokenRevokeResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *TokenRevokeResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *TokenRevokeResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *TokenRevokeResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *TokenRevokeResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *TokenRevokeResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *TokenRevokeResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *TokenRevokeResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetCount

`func (o *TokenRevokeResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TokenRevokeResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TokenRevokeResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TokenRevokeResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


