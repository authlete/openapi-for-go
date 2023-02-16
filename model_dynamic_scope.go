/*
Authlete API

Authlete API Document. 

API version: 2.2.30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

// checks if the DynamicScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DynamicScope{}

// DynamicScope struct for DynamicScope
type DynamicScope struct {
	// The scope name.
	Name *string `json:"name,omitempty"`
	// The scope value.
	Value *string `json:"value,omitempty"`
}

// NewDynamicScope instantiates a new DynamicScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicScope() *DynamicScope {
	this := DynamicScope{}
	return &this
}

// NewDynamicScopeWithDefaults instantiates a new DynamicScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicScopeWithDefaults() *DynamicScope {
	this := DynamicScope{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DynamicScope) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicScope) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DynamicScope) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DynamicScope) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DynamicScope) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicScope) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DynamicScope) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DynamicScope) SetValue(v string) {
	o.Value = &v
}

func (o DynamicScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DynamicScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableDynamicScope struct {
	value *DynamicScope
	isSet bool
}

func (v NullableDynamicScope) Get() *DynamicScope {
	return v.value
}

func (v *NullableDynamicScope) Set(val *DynamicScope) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicScope) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicScope(val *DynamicScope) *NullableDynamicScope {
	return &NullableDynamicScope{value: val, isSet: true}
}

func (v NullableDynamicScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


