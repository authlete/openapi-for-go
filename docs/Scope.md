# Scope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the scope. | [optional] 
**DefaultEntry** | Pointer to **bool** | &#x60;true&#x60; to mark the scope as default. Scopes marked as default are regarded as requested when an authorization request from a client application does not contain scope request parameter.  | [optional] 
**Description** | Pointer to **string** | The description about the scope. | [optional] 
**Descriptions** | Pointer to [**[]TaggedValue**](TaggedValue.md) | The descriptions about this scope in multiple languages. | [optional] 
**Attributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of the scope. | [optional] 

## Methods

### NewScope

`func NewScope() *Scope`

NewScope instantiates a new Scope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeWithDefaults

`func NewScopeWithDefaults() *Scope`

NewScopeWithDefaults instantiates a new Scope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Scope) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Scope) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Scope) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Scope) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultEntry

`func (o *Scope) GetDefaultEntry() bool`

GetDefaultEntry returns the DefaultEntry field if non-nil, zero value otherwise.

### GetDefaultEntryOk

`func (o *Scope) GetDefaultEntryOk() (*bool, bool)`

GetDefaultEntryOk returns a tuple with the DefaultEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEntry

`func (o *Scope) SetDefaultEntry(v bool)`

SetDefaultEntry sets DefaultEntry field to given value.

### HasDefaultEntry

`func (o *Scope) HasDefaultEntry() bool`

HasDefaultEntry returns a boolean if a field has been set.

### GetDescription

`func (o *Scope) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Scope) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Scope) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Scope) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptions

`func (o *Scope) GetDescriptions() []TaggedValue`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *Scope) GetDescriptionsOk() (*[]TaggedValue, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *Scope) SetDescriptions(v []TaggedValue)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *Scope) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetAttributes

`func (o *Scope) GetAttributes() []Pair`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Scope) GetAttributesOk() (*[]Pair, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Scope) SetAttributes(v []Pair)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Scope) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


