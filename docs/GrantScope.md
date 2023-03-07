# GrantScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to **string** | Space-delimited scopes.  | [optional] 
**Resource** | Pointer to **[]string** | List of resource indicators.  | [optional] 

## Methods

### NewGrantScope

`func NewGrantScope() *GrantScope`

NewGrantScope instantiates a new GrantScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantScopeWithDefaults

`func NewGrantScopeWithDefaults() *GrantScope`

NewGrantScopeWithDefaults instantiates a new GrantScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *GrantScope) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GrantScope) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GrantScope) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GrantScope) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetResource

`func (o *GrantScope) GetResource() []string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *GrantScope) GetResourceOk() (*[]string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *GrantScope) SetResource(v []string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *GrantScope) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


