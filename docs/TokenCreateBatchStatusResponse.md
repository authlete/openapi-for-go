# TokenCreateBatchStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Status** | Pointer to [**TokenBatchStatus**](TokenBatchStatus.md) |  | [optional] 

## Methods

### NewTokenCreateBatchStatusResponse

`func NewTokenCreateBatchStatusResponse() *TokenCreateBatchStatusResponse`

NewTokenCreateBatchStatusResponse instantiates a new TokenCreateBatchStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCreateBatchStatusResponseWithDefaults

`func NewTokenCreateBatchStatusResponseWithDefaults() *TokenCreateBatchStatusResponse`

NewTokenCreateBatchStatusResponseWithDefaults instantiates a new TokenCreateBatchStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *TokenCreateBatchStatusResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *TokenCreateBatchStatusResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *TokenCreateBatchStatusResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *TokenCreateBatchStatusResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *TokenCreateBatchStatusResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *TokenCreateBatchStatusResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *TokenCreateBatchStatusResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *TokenCreateBatchStatusResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetStatus

`func (o *TokenCreateBatchStatusResponse) GetStatus() TokenBatchStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TokenCreateBatchStatusResponse) GetStatusOk() (*TokenBatchStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TokenCreateBatchStatusResponse) SetStatus(v TokenBatchStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TokenCreateBatchStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


