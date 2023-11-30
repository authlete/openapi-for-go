# HskGetListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | Result of the API call | [optional] 
**Hsks** | Pointer to [**[]Hsk**](Hsk.md) | List of HSK | [optional] 

## Methods

### NewHskGetListResponse

`func NewHskGetListResponse() *HskGetListResponse`

NewHskGetListResponse instantiates a new HskGetListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHskGetListResponseWithDefaults

`func NewHskGetListResponseWithDefaults() *HskGetListResponse`

NewHskGetListResponseWithDefaults instantiates a new HskGetListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *HskGetListResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *HskGetListResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *HskGetListResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *HskGetListResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *HskGetListResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *HskGetListResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *HskGetListResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *HskGetListResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *HskGetListResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HskGetListResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HskGetListResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *HskGetListResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetHsks

`func (o *HskGetListResponse) GetHsks() []Hsk`

GetHsks returns the Hsks field if non-nil, zero value otherwise.

### GetHsksOk

`func (o *HskGetListResponse) GetHsksOk() (*[]Hsk, bool)`

GetHsksOk returns a tuple with the Hsks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsks

`func (o *HskGetListResponse) SetHsks(v []Hsk)`

SetHsks sets Hsks field to given value.

### HasHsks

`func (o *HskGetListResponse) HasHsks() bool`

HasHsks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


