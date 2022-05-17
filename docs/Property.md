# Property

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The key part. | [optional] 
**Value** | Pointer to **string** | The value part. | [optional] 
**Hidden** | Pointer to **bool** | The flag to indicate whether this property hidden from or visible to client applications. If &#x60;true&#x60;, this property is hidden from client applications. Otherwise, this property is visible to client applications. | [optional] 

## Methods

### NewProperty

`func NewProperty() *Property`

NewProperty instantiates a new Property object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyWithDefaults

`func NewPropertyWithDefaults() *Property`

NewPropertyWithDefaults instantiates a new Property object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Property) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Property) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Property) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Property) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *Property) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Property) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Property) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Property) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetHidden

`func (o *Property) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Property) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Property) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Property) HasHidden() bool`

HasHidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


