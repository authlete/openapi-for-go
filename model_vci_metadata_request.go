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

// checks if the VciMetadataRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VciMetadataRequest{}

// VciMetadataRequest struct for VciMetadataRequest
type VciMetadataRequest struct {
	// The flag indicating whether the metadata is written in the pretty format or not.
	Pretty bool `json:"pretty"`
}

// NewVciMetadataRequest instantiates a new VciMetadataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVciMetadataRequest(pretty bool) *VciMetadataRequest {
	this := VciMetadataRequest{}
	this.Pretty = pretty
	return &this
}

// NewVciMetadataRequestWithDefaults instantiates a new VciMetadataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVciMetadataRequestWithDefaults() *VciMetadataRequest {
	this := VciMetadataRequest{}
	return &this
}

// GetPretty returns the Pretty field value
func (o *VciMetadataRequest) GetPretty() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Pretty
}

// GetPrettyOk returns a tuple with the Pretty field value
// and a boolean to check if the value has been set.
func (o *VciMetadataRequest) GetPrettyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pretty, true
}

// SetPretty sets field value
func (o *VciMetadataRequest) SetPretty(v bool) {
	o.Pretty = v
}

func (o VciMetadataRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VciMetadataRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pretty"] = o.Pretty
	return toSerialize, nil
}

type NullableVciMetadataRequest struct {
	value *VciMetadataRequest
	isSet bool
}

func (v NullableVciMetadataRequest) Get() *VciMetadataRequest {
	return v.value
}

func (v *NullableVciMetadataRequest) Set(val *VciMetadataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVciMetadataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVciMetadataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVciMetadataRequest(val *VciMetadataRequest) *NullableVciMetadataRequest {
	return &NullableVciMetadataRequest{value: val, isSet: true}
}

func (v NullableVciMetadataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVciMetadataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


