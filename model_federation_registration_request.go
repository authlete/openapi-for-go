/*
Authlete API

Authlete API Document.

API version: 2.3.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

// checks if the FederationRegistrationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FederationRegistrationRequest{}

// FederationRegistrationRequest struct for FederationRegistrationRequest
type FederationRegistrationRequest struct {
	// The entity configuration of a relying party.
	EntityConfiguration *string `json:"entityConfiguration,omitempty"`
	// The trust chain of a relying party.
	TrustChain *string `json:"trustChain,omitempty"`
}

// NewFederationRegistrationRequest instantiates a new FederationRegistrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederationRegistrationRequest() *FederationRegistrationRequest {
	this := FederationRegistrationRequest{}
	return &this
}

// NewFederationRegistrationRequestWithDefaults instantiates a new FederationRegistrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederationRegistrationRequestWithDefaults() *FederationRegistrationRequest {
	this := FederationRegistrationRequest{}
	return &this
}

// GetEntityConfiguration returns the EntityConfiguration field value if set, zero value otherwise.
func (o *FederationRegistrationRequest) GetEntityConfiguration() string {
	if o == nil || isNil(o.EntityConfiguration) {
		var ret string
		return ret
	}
	return *o.EntityConfiguration
}

// GetEntityConfigurationOk returns a tuple with the EntityConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationRegistrationRequest) GetEntityConfigurationOk() (*string, bool) {
	if o == nil || isNil(o.EntityConfiguration) {
		return nil, false
	}
	return o.EntityConfiguration, true
}

// HasEntityConfiguration returns a boolean if a field has been set.
func (o *FederationRegistrationRequest) HasEntityConfiguration() bool {
	if o != nil && !isNil(o.EntityConfiguration) {
		return true
	}

	return false
}

// SetEntityConfiguration gets a reference to the given string and assigns it to the EntityConfiguration field.
func (o *FederationRegistrationRequest) SetEntityConfiguration(v string) {
	o.EntityConfiguration = &v
}

// GetTrustChain returns the TrustChain field value if set, zero value otherwise.
func (o *FederationRegistrationRequest) GetTrustChain() string {
	if o == nil || isNil(o.TrustChain) {
		var ret string
		return ret
	}
	return *o.TrustChain
}

// GetTrustChainOk returns a tuple with the TrustChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationRegistrationRequest) GetTrustChainOk() (*string, bool) {
	if o == nil || isNil(o.TrustChain) {
		return nil, false
	}
	return o.TrustChain, true
}

// HasTrustChain returns a boolean if a field has been set.
func (o *FederationRegistrationRequest) HasTrustChain() bool {
	if o != nil && !isNil(o.TrustChain) {
		return true
	}

	return false
}

// SetTrustChain gets a reference to the given string and assigns it to the TrustChain field.
func (o *FederationRegistrationRequest) SetTrustChain(v string) {
	o.TrustChain = &v
}

func (o FederationRegistrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FederationRegistrationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EntityConfiguration) {
		toSerialize["entityConfiguration"] = o.EntityConfiguration
	}
	if !isNil(o.TrustChain) {
		toSerialize["trustChain"] = o.TrustChain
	}
	return toSerialize, nil
}

type NullableFederationRegistrationRequest struct {
	value *FederationRegistrationRequest
	isSet bool
}

func (v NullableFederationRegistrationRequest) Get() *FederationRegistrationRequest {
	return v.value
}

func (v *NullableFederationRegistrationRequest) Set(val *FederationRegistrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFederationRegistrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFederationRegistrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederationRegistrationRequest(val *FederationRegistrationRequest) *NullableFederationRegistrationRequest {
	return &NullableFederationRegistrationRequest{value: val, isSet: true}
}

func (v NullableFederationRegistrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederationRegistrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
