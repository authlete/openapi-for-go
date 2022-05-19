/*
Authlete API

Authlete API Document. 

API version: 2.2.25
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

// Scope struct for Scope
type Scope struct {
	// The name of the scope.
	Name *string `json:"name,omitempty"`
	// `true` to mark the scope as default. Scopes marked as default are regarded as requested when an authorization request from a client application does not contain scope request parameter. 
	DefaultEntry *bool `json:"defaultEntry,omitempty"`
	// The description about the scope.
	Description *string `json:"description,omitempty"`
	// The descriptions about this scope in multiple languages.
	Descriptions []TaggedValue `json:"descriptions,omitempty"`
	// The attributes of the scope.
	Attributes []Pair `json:"attributes,omitempty"`
}

// NewScope instantiates a new Scope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScope() *Scope {
	this := Scope{}
	return &this
}

// NewScopeWithDefaults instantiates a new Scope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopeWithDefaults() *Scope {
	this := Scope{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Scope) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scope) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Scope) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Scope) SetName(v string) {
	o.Name = &v
}

// GetDefaultEntry returns the DefaultEntry field value if set, zero value otherwise.
func (o *Scope) GetDefaultEntry() bool {
	if o == nil || o.DefaultEntry == nil {
		var ret bool
		return ret
	}
	return *o.DefaultEntry
}

// GetDefaultEntryOk returns a tuple with the DefaultEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scope) GetDefaultEntryOk() (*bool, bool) {
	if o == nil || o.DefaultEntry == nil {
		return nil, false
	}
	return o.DefaultEntry, true
}

// HasDefaultEntry returns a boolean if a field has been set.
func (o *Scope) HasDefaultEntry() bool {
	if o != nil && o.DefaultEntry != nil {
		return true
	}

	return false
}

// SetDefaultEntry gets a reference to the given bool and assigns it to the DefaultEntry field.
func (o *Scope) SetDefaultEntry(v bool) {
	o.DefaultEntry = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Scope) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scope) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Scope) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Scope) SetDescription(v string) {
	o.Description = &v
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *Scope) GetDescriptions() []TaggedValue {
	if o == nil || o.Descriptions == nil {
		var ret []TaggedValue
		return ret
	}
	return o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scope) GetDescriptionsOk() ([]TaggedValue, bool) {
	if o == nil || o.Descriptions == nil {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *Scope) HasDescriptions() bool {
	if o != nil && o.Descriptions != nil {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []TaggedValue and assigns it to the Descriptions field.
func (o *Scope) SetDescriptions(v []TaggedValue) {
	o.Descriptions = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Scope) GetAttributes() []Pair {
	if o == nil || o.Attributes == nil {
		var ret []Pair
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scope) GetAttributesOk() ([]Pair, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Scope) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []Pair and assigns it to the Attributes field.
func (o *Scope) SetAttributes(v []Pair) {
	o.Attributes = v
}

func (o Scope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DefaultEntry != nil {
		toSerialize["defaultEntry"] = o.DefaultEntry
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Descriptions != nil {
		toSerialize["descriptions"] = o.Descriptions
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableScope struct {
	value *Scope
	isSet bool
}

func (v NullableScope) Get() *Scope {
	return v.value
}

func (v *NullableScope) Set(val *Scope) {
	v.value = val
	v.isSet = true
}

func (v NullableScope) IsSet() bool {
	return v.isSet
}

func (v *NullableScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScope(val *Scope) *NullableScope {
	return &NullableScope{value: val, isSet: true}
}

func (v NullableScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


