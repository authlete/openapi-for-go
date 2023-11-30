# HskDeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | Result of the API call | [optional] 
**Hsk** | Pointer to [**Hsk**](Hsk.md) |  | [optional] 

## Methods

### NewHskDeleteResponse

`func NewHskDeleteResponse() *HskDeleteResponse`

NewHskDeleteResponse instantiates a new HskDeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHskDeleteResponseWithDefaults

`func NewHskDeleteResponseWithDefaults() *HskDeleteResponse`

NewHskDeleteResponseWithDefaults instantiates a new HskDeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *HskDeleteResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *HskDeleteResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *HskDeleteResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *HskDeleteResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *HskDeleteResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *HskDeleteResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *HskDeleteResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *HskDeleteResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *HskDeleteResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HskDeleteResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HskDeleteResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *HskDeleteResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetHsk

`func (o *HskDeleteResponse) GetHsk() Hsk`

GetHsk returns the Hsk field if non-nil, zero value otherwise.

### GetHskOk

`func (o *HskDeleteResponse) GetHskOk() (*Hsk, bool)`

GetHskOk returns a tuple with the Hsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsk

`func (o *HskDeleteResponse) SetHsk(v Hsk)`

SetHsk sets Hsk field to given value.

### HasHsk

`func (o *HskDeleteResponse) HasHsk() bool`

HasHsk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


