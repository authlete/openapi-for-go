/*
Authlete API

Authlete API Document. 

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

// checks if the AuthzDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthzDetails{}

// AuthzDetails The authorization details. This represents the value of the `authorization_details` request parameter in the preceding device authorization request which is defined in \"OAuth 2.0 Rich Authorization Requests\". 
type AuthzDetails struct {
	// Elements of this authorization details. 
	Elements []AuthorizationDetailsElement `json:"elements,omitempty"`
}

// NewAuthzDetails instantiates a new AuthzDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthzDetails() *AuthzDetails {
	this := AuthzDetails{}
	return &this
}

// NewAuthzDetailsWithDefaults instantiates a new AuthzDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthzDetailsWithDefaults() *AuthzDetails {
	this := AuthzDetails{}
	return &this
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *AuthzDetails) GetElements() []AuthorizationDetailsElement {
	if o == nil || isNil(o.Elements) {
		var ret []AuthorizationDetailsElement
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthzDetails) GetElementsOk() ([]AuthorizationDetailsElement, bool) {
	if o == nil || isNil(o.Elements) {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *AuthzDetails) HasElements() bool {
	if o != nil && !isNil(o.Elements) {
		return true
	}

	return false
}

// SetElements gets a reference to the given []AuthorizationDetailsElement and assigns it to the Elements field.
func (o *AuthzDetails) SetElements(v []AuthorizationDetailsElement) {
	o.Elements = v
}

func (o AuthzDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthzDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Elements) {
		toSerialize["elements"] = o.Elements
	}
	return toSerialize, nil
}

type NullableAuthzDetails struct {
	value *AuthzDetails
	isSet bool
}

func (v NullableAuthzDetails) Get() *AuthzDetails {
	return v.value
}

func (v *NullableAuthzDetails) Set(val *AuthzDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthzDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthzDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthzDetails(val *AuthzDetails) *NullableAuthzDetails {
	return &NullableAuthzDetails{value: val, isSet: true}
}

func (v NullableAuthzDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthzDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


