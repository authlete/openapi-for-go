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

// checks if the TokenFailRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenFailRequest{}

// TokenFailRequest struct for TokenFailRequest
type TokenFailRequest struct {
	// The ticket issued from Authlete `/auth/token` API. 
	Ticket string `json:"ticket"`
	// The reason of the failure of the token request.
	Reason string `json:"reason"`
}

// NewTokenFailRequest instantiates a new TokenFailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenFailRequest(ticket string, reason string) *TokenFailRequest {
	this := TokenFailRequest{}
	this.Ticket = ticket
	this.Reason = reason
	return &this
}

// NewTokenFailRequestWithDefaults instantiates a new TokenFailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenFailRequestWithDefaults() *TokenFailRequest {
	this := TokenFailRequest{}
	return &this
}

// GetTicket returns the Ticket field value
func (o *TokenFailRequest) GetTicket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ticket
}

// GetTicketOk returns a tuple with the Ticket field value
// and a boolean to check if the value has been set.
func (o *TokenFailRequest) GetTicketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ticket, true
}

// SetTicket sets field value
func (o *TokenFailRequest) SetTicket(v string) {
	o.Ticket = v
}

// GetReason returns the Reason field value
func (o *TokenFailRequest) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *TokenFailRequest) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *TokenFailRequest) SetReason(v string) {
	o.Reason = v
}

func (o TokenFailRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenFailRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ticket"] = o.Ticket
	toSerialize["reason"] = o.Reason
	return toSerialize, nil
}

type NullableTokenFailRequest struct {
	value *TokenFailRequest
	isSet bool
}

func (v NullableTokenFailRequest) Get() *TokenFailRequest {
	return v.value
}

func (v *NullableTokenFailRequest) Set(val *TokenFailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenFailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenFailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenFailRequest(val *TokenFailRequest) *NullableTokenFailRequest {
	return &NullableTokenFailRequest{value: val, isSet: true}
}

func (v NullableTokenFailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenFailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


