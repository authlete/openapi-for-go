# ServiceGetListApi200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **int32** | Start index (inclusive) of the result set. The default value is 0. Must not be a negative number.  | [optional] 
**End** | Pointer to **int32** | Start index (inclusive) of the result set. The default value is 0. Must not be a negative number.  | [optional] 
**TotalCount** | Pointer to **int32** | Total number of services owned by the service owner. This doesn&#39;t mean the number of services contained in the response.  | [optional] 
**Services** | Pointer to [**[]ServiceLimited**](ServiceLimited.md) | An array of services.  | [optional] 

## Methods

### NewServiceGetListApi200Response

`func NewServiceGetListApi200Response() *ServiceGetListApi200Response`

NewServiceGetListApi200Response instantiates a new ServiceGetListApi200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceGetListApi200ResponseWithDefaults

`func NewServiceGetListApi200ResponseWithDefaults() *ServiceGetListApi200Response`

NewServiceGetListApi200ResponseWithDefaults instantiates a new ServiceGetListApi200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *ServiceGetListApi200Response) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ServiceGetListApi200Response) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ServiceGetListApi200Response) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ServiceGetListApi200Response) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *ServiceGetListApi200Response) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ServiceGetListApi200Response) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ServiceGetListApi200Response) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ServiceGetListApi200Response) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetTotalCount

`func (o *ServiceGetListApi200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ServiceGetListApi200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ServiceGetListApi200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ServiceGetListApi200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetServices

`func (o *ServiceGetListApi200Response) GetServices() []ServiceLimited`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ServiceGetListApi200Response) GetServicesOk() (*[]ServiceLimited, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ServiceGetListApi200Response) SetServices(v []ServiceLimited)`

SetServices sets Services field to given value.

### HasServices

`func (o *ServiceGetListApi200Response) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


