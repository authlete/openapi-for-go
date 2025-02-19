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

// checks if the ClientFlagUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientFlagUpdateRequest{}

// ClientFlagUpdateRequest struct for ClientFlagUpdateRequest
type ClientFlagUpdateRequest struct {
	// The flag value to be set
	ClientLocked bool `json:"clientLocked"`
}

// NewClientFlagUpdateRequest instantiates a new ClientFlagUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientFlagUpdateRequest(clientLocked bool) *ClientFlagUpdateRequest {
	this := ClientFlagUpdateRequest{}
	this.ClientLocked = clientLocked
	return &this
}

// NewClientFlagUpdateRequestWithDefaults instantiates a new ClientFlagUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientFlagUpdateRequestWithDefaults() *ClientFlagUpdateRequest {
	this := ClientFlagUpdateRequest{}
	return &this
}

// GetClientLocked returns the ClientLocked field value
func (o *ClientFlagUpdateRequest) GetClientLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ClientLocked
}

// GetClientLockedOk returns a tuple with the ClientLocked field value
// and a boolean to check if the value has been set.
func (o *ClientFlagUpdateRequest) GetClientLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientLocked, true
}

// SetClientLocked sets field value
func (o *ClientFlagUpdateRequest) SetClientLocked(v bool) {
	o.ClientLocked = v
}

func (o ClientFlagUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientFlagUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientLocked"] = o.ClientLocked
	return toSerialize, nil
}

type NullableClientFlagUpdateRequest struct {
	value *ClientFlagUpdateRequest
	isSet bool
}

func (v NullableClientFlagUpdateRequest) Get() *ClientFlagUpdateRequest {
	return v.value
}

func (v *NullableClientFlagUpdateRequest) Set(val *ClientFlagUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClientFlagUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClientFlagUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientFlagUpdateRequest(val *ClientFlagUpdateRequest) *NullableClientFlagUpdateRequest {
	return &NullableClientFlagUpdateRequest{value: val, isSet: true}
}

func (v NullableClientFlagUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientFlagUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
