# HskGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | Result of the API call | [optional] 
**Hsk** | Pointer to [**Hsk**](Hsk.md) |  | [optional] 

## Methods

### NewHskGetResponse

`func NewHskGetResponse() *HskGetResponse`

NewHskGetResponse instantiates a new HskGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHskGetResponseWithDefaults

`func NewHskGetResponseWithDefaults() *HskGetResponse`

NewHskGetResponseWithDefaults instantiates a new HskGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *HskGetResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *HskGetResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *HskGetResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *HskGetResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *HskGetResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *HskGetResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *HskGetResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *HskGetResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *HskGetResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HskGetResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HskGetResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *HskGetResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetHsk

`func (o *HskGetResponse) GetHsk() Hsk`

GetHsk returns the Hsk field if non-nil, zero value otherwise.

### GetHskOk

`func (o *HskGetResponse) GetHskOk() (*Hsk, bool)`

GetHskOk returns a tuple with the Hsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsk

`func (o *HskGetResponse) SetHsk(v Hsk)`

SetHsk sets Hsk field to given value.

### HasHsk

`func (o *HskGetResponse) HasHsk() bool`

HasHsk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


