# VciOfferCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The result of the &#x60;/vci/offer/create&#x60; API call. | [optional] 
**Info** | Pointer to [**CredentialOfferInfo**](CredentialOfferInfo.md) |  | [optional] 

## Methods

### NewVciOfferCreateResponse

`func NewVciOfferCreateResponse() *VciOfferCreateResponse`

NewVciOfferCreateResponse instantiates a new VciOfferCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVciOfferCreateResponseWithDefaults

`func NewVciOfferCreateResponseWithDefaults() *VciOfferCreateResponse`

NewVciOfferCreateResponseWithDefaults instantiates a new VciOfferCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *VciOfferCreateResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *VciOfferCreateResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *VciOfferCreateResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *VciOfferCreateResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *VciOfferCreateResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *VciOfferCreateResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *VciOfferCreateResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *VciOfferCreateResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *VciOfferCreateResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VciOfferCreateResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VciOfferCreateResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *VciOfferCreateResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetInfo

`func (o *VciOfferCreateResponse) GetInfo() CredentialOfferInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *VciOfferCreateResponse) GetInfoOk() (*CredentialOfferInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *VciOfferCreateResponse) SetInfo(v CredentialOfferInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *VciOfferCreateResponse) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


