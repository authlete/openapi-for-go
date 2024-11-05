# ServiceGetListLimitedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **int32** | Start index (inclusive) of the result set. The default value is 0. Must not be a negative number.  | [optional] 
**End** | Pointer to **int32** | Start index (inclusive) of the result set. The default value is 0. Must not be a negative number.  | [optional] 
**TotalCount** | Pointer to **int32** | Total number of services owned by the service owner. This doesn&#39;t mean the number of services contained in the response.  | [optional] 
**Services** | Pointer to [**[]ServiceLimited**](ServiceLimited.md) | An array of services.  | [optional] 

## Methods

### NewServiceGetListLimitedResponse

`func NewServiceGetListLimitedResponse() *ServiceGetListLimitedResponse`

NewServiceGetListLimitedResponse instantiates a new ServiceGetListLimitedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceGetListLimitedResponseWithDefaults

`func NewServiceGetListLimitedResponseWithDefaults() *ServiceGetListLimitedResponse`

NewServiceGetListLimitedResponseWithDefaults instantiates a new ServiceGetListLimitedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *ServiceGetListLimitedResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ServiceGetListLimitedResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ServiceGetListLimitedResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ServiceGetListLimitedResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *ServiceGetListLimitedResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ServiceGetListLimitedResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ServiceGetListLimitedResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ServiceGetListLimitedResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetTotalCount

`func (o *ServiceGetListLimitedResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ServiceGetListLimitedResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ServiceGetListLimitedResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ServiceGetListLimitedResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetServices

`func (o *ServiceGetListLimitedResponse) GetServices() []ServiceLimited`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ServiceGetListLimitedResponse) GetServicesOk() (*[]ServiceLimited, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ServiceGetListLimitedResponse) SetServices(v []ServiceLimited)`

SetServices sets Services field to given value.

### HasServices

`func (o *ServiceGetListLimitedResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


