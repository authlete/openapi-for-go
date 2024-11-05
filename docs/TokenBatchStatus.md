# TokenBatchStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchKind** | Pointer to **string** | The kind of the batch. | [optional] 
**RequestId** | Pointer to **string** | The request ID associated with the status. | [optional] 
**Result** | Pointer to [**TokenBatchStatusResult**](TokenBatchStatusResult.md) |  | [optional] 
**TokenCount** | Pointer to **int64** | The number of access tokens processed by the batch. | [optional] 
**ErrorCode** | Pointer to **string** | The error code for the batch operation | [optional] 
**ErrorDescription** | Pointer to **string** | The error description for the batch operation | [optional] 
**CreatedAt** | Pointer to **int64** | The time when this status was created | [optional] 
**ModifiedAt** | Pointer to **int64** | The time when this status was modified | [optional] 

## Methods

### NewTokenBatchStatus

`func NewTokenBatchStatus() *TokenBatchStatus`

NewTokenBatchStatus instantiates a new TokenBatchStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenBatchStatusWithDefaults

`func NewTokenBatchStatusWithDefaults() *TokenBatchStatus`

NewTokenBatchStatusWithDefaults instantiates a new TokenBatchStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchKind

`func (o *TokenBatchStatus) GetBatchKind() string`

GetBatchKind returns the BatchKind field if non-nil, zero value otherwise.

### GetBatchKindOk

`func (o *TokenBatchStatus) GetBatchKindOk() (*string, bool)`

GetBatchKindOk returns a tuple with the BatchKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchKind

`func (o *TokenBatchStatus) SetBatchKind(v string)`

SetBatchKind sets BatchKind field to given value.

### HasBatchKind

`func (o *TokenBatchStatus) HasBatchKind() bool`

HasBatchKind returns a boolean if a field has been set.

### GetRequestId

`func (o *TokenBatchStatus) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TokenBatchStatus) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TokenBatchStatus) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *TokenBatchStatus) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetResult

`func (o *TokenBatchStatus) GetResult() TokenBatchStatusResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TokenBatchStatus) GetResultOk() (*TokenBatchStatusResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TokenBatchStatus) SetResult(v TokenBatchStatusResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *TokenBatchStatus) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetTokenCount

`func (o *TokenBatchStatus) GetTokenCount() int64`

GetTokenCount returns the TokenCount field if non-nil, zero value otherwise.

### GetTokenCountOk

`func (o *TokenBatchStatus) GetTokenCountOk() (*int64, bool)`

GetTokenCountOk returns a tuple with the TokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCount

`func (o *TokenBatchStatus) SetTokenCount(v int64)`

SetTokenCount sets TokenCount field to given value.

### HasTokenCount

`func (o *TokenBatchStatus) HasTokenCount() bool`

HasTokenCount returns a boolean if a field has been set.

### GetErrorCode

`func (o *TokenBatchStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *TokenBatchStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *TokenBatchStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *TokenBatchStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *TokenBatchStatus) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *TokenBatchStatus) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *TokenBatchStatus) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *TokenBatchStatus) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TokenBatchStatus) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TokenBatchStatus) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TokenBatchStatus) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TokenBatchStatus) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *TokenBatchStatus) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *TokenBatchStatus) GetModifiedAtOk() (*int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *TokenBatchStatus) SetModifiedAt(v int64)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *TokenBatchStatus) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


