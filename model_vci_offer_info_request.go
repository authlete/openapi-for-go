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

// checks if the VciOfferInfoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VciOfferInfoRequest{}

// VciOfferInfoRequest struct for VciOfferInfoRequest
type VciOfferInfoRequest struct {
	// The identifier of the credential offer.
	Identifier *string `json:"identifier,omitempty"`
}

// NewVciOfferInfoRequest instantiates a new VciOfferInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVciOfferInfoRequest() *VciOfferInfoRequest {
	this := VciOfferInfoRequest{}
	return &this
}

// NewVciOfferInfoRequestWithDefaults instantiates a new VciOfferInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVciOfferInfoRequestWithDefaults() *VciOfferInfoRequest {
	this := VciOfferInfoRequest{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *VciOfferInfoRequest) GetIdentifier() string {
	if o == nil || isNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciOfferInfoRequest) GetIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *VciOfferInfoRequest) HasIdentifier() bool {
	if o != nil && !isNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *VciOfferInfoRequest) SetIdentifier(v string) {
	o.Identifier = &v
}

func (o VciOfferInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VciOfferInfoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	return toSerialize, nil
}

type NullableVciOfferInfoRequest struct {
	value *VciOfferInfoRequest
	isSet bool
}

func (v NullableVciOfferInfoRequest) Get() *VciOfferInfoRequest {
	return v.value
}

func (v *NullableVciOfferInfoRequest) Set(val *VciOfferInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVciOfferInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVciOfferInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVciOfferInfoRequest(val *VciOfferInfoRequest) *NullableVciOfferInfoRequest {
	return &NullableVciOfferInfoRequest{value: val, isSet: true}
}

func (v NullableVciOfferInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVciOfferInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

