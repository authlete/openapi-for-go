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

// AuthorizationDetails The authorization details. This represents the value of the `authorization_details` request parameter in the preceding device authorization request which is defined in \"OAuth 2.0 Rich Authorization Requests\". 
type AuthorizationDetails struct {
	// Elements of this authorization details. 
	Elements []AuthorizationDetailsElement `json:"elements,omitempty"`
}

// NewAuthorizationDetails instantiates a new AuthorizationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationDetails() *AuthorizationDetails {
	this := AuthorizationDetails{}
	return &this
}

// NewAuthorizationDetailsWithDefaults instantiates a new AuthorizationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationDetailsWithDefaults() *AuthorizationDetails {
	this := AuthorizationDetails{}
	return &this
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *AuthorizationDetails) GetElements() []AuthorizationDetailsElement {
	if o == nil || o.Elements == nil {
		var ret []AuthorizationDetailsElement
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDetails) GetElementsOk() ([]AuthorizationDetailsElement, bool) {
	if o == nil || o.Elements == nil {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *AuthorizationDetails) HasElements() bool {
	if o != nil && o.Elements != nil {
		return true
	}

	return false
}

// SetElements gets a reference to the given []AuthorizationDetailsElement and assigns it to the Elements field.
func (o *AuthorizationDetails) SetElements(v []AuthorizationDetailsElement) {
	o.Elements = v
}

func (o AuthorizationDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Elements != nil {
		toSerialize["elements"] = o.Elements
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationDetails struct {
	value *AuthorizationDetails
	isSet bool
}

func (v NullableAuthorizationDetails) Get() *AuthorizationDetails {
	return v.value
}

func (v *NullableAuthorizationDetails) Set(val *AuthorizationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationDetails(val *AuthorizationDetails) *NullableAuthorizationDetails {
	return &NullableAuthorizationDetails{value: val, isSet: true}
}

func (v NullableAuthorizationDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

