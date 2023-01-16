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

// checks if the BackchannelAuthenticationIssueRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackchannelAuthenticationIssueRequest{}

// BackchannelAuthenticationIssueRequest struct for BackchannelAuthenticationIssueRequest
type BackchannelAuthenticationIssueRequest struct {
	// The ticket issued from Authlete's `/backchannel/authentication` API.
	Ticket string `json:"ticket"`
}

// NewBackchannelAuthenticationIssueRequest instantiates a new BackchannelAuthenticationIssueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackchannelAuthenticationIssueRequest(ticket string) *BackchannelAuthenticationIssueRequest {
	this := BackchannelAuthenticationIssueRequest{}
	this.Ticket = ticket
	return &this
}

// NewBackchannelAuthenticationIssueRequestWithDefaults instantiates a new BackchannelAuthenticationIssueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackchannelAuthenticationIssueRequestWithDefaults() *BackchannelAuthenticationIssueRequest {
	this := BackchannelAuthenticationIssueRequest{}
	return &this
}

// GetTicket returns the Ticket field value
func (o *BackchannelAuthenticationIssueRequest) GetTicket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ticket
}

// GetTicketOk returns a tuple with the Ticket field value
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationIssueRequest) GetTicketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ticket, true
}

// SetTicket sets field value
func (o *BackchannelAuthenticationIssueRequest) SetTicket(v string) {
	o.Ticket = v
}

func (o BackchannelAuthenticationIssueRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackchannelAuthenticationIssueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ticket"] = o.Ticket
	return toSerialize, nil
}

type NullableBackchannelAuthenticationIssueRequest struct {
	value *BackchannelAuthenticationIssueRequest
	isSet bool
}

func (v NullableBackchannelAuthenticationIssueRequest) Get() *BackchannelAuthenticationIssueRequest {
	return v.value
}

func (v *NullableBackchannelAuthenticationIssueRequest) Set(val *BackchannelAuthenticationIssueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBackchannelAuthenticationIssueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBackchannelAuthenticationIssueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackchannelAuthenticationIssueRequest(val *BackchannelAuthenticationIssueRequest) *NullableBackchannelAuthenticationIssueRequest {
	return &NullableBackchannelAuthenticationIssueRequest{value: val, isSet: true}
}

func (v NullableBackchannelAuthenticationIssueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackchannelAuthenticationIssueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


