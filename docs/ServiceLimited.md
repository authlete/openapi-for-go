# ServiceLimited

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | The sequential number of the service. The value of this property is assigned by Authlete. | [optional] [readonly] 
**ServiceName** | Pointer to **string** | The name of this service. | [optional] 
**Description** | Pointer to **string** | The description about the service. | [optional] 
**ApiKey** | Pointer to **int64** | The service ID used in Authlete API calls. The value of this property is assigned by Authlete. | [optional] [readonly] 
**ClientIdAliasEnabled** | Pointer to **bool** | Deprecated. Always &#x60;true&#x60;. | [optional] 

## Methods

### NewServiceLimited

`func NewServiceLimited() *ServiceLimited`

NewServiceLimited instantiates a new ServiceLimited object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceLimitedWithDefaults

`func NewServiceLimitedWithDefaults() *ServiceLimited`

NewServiceLimitedWithDefaults instantiates a new ServiceLimited object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *ServiceLimited) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ServiceLimited) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ServiceLimited) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ServiceLimited) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetServiceName

`func (o *ServiceLimited) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceLimited) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceLimited) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ServiceLimited) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceLimited) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceLimited) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceLimited) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceLimited) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApiKey

`func (o *ServiceLimited) GetApiKey() int64`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ServiceLimited) GetApiKeyOk() (*int64, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ServiceLimited) SetApiKey(v int64)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ServiceLimited) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetClientIdAliasEnabled

`func (o *ServiceLimited) GetClientIdAliasEnabled() bool`

GetClientIdAliasEnabled returns the ClientIdAliasEnabled field if non-nil, zero value otherwise.

### GetClientIdAliasEnabledOk

`func (o *ServiceLimited) GetClientIdAliasEnabledOk() (*bool, bool)`

GetClientIdAliasEnabledOk returns a tuple with the ClientIdAliasEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasEnabled

`func (o *ServiceLimited) SetClientIdAliasEnabled(v bool)`

SetClientIdAliasEnabled sets ClientIdAliasEnabled field to given value.

### HasClientIdAliasEnabled

`func (o *ServiceLimited) HasClientIdAliasEnabled() bool`

HasClientIdAliasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


