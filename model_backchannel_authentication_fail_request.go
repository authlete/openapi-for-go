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

// checks if the BackchannelAuthenticationFailRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackchannelAuthenticationFailRequest{}

// BackchannelAuthenticationFailRequest struct for BackchannelAuthenticationFailRequest
type BackchannelAuthenticationFailRequest struct {
	// The ticket which should be deleted on a call of Authlete's `/backchannel/authentication/fail` API. This request parameter is not mandatory but optional. If this request parameter is given and the ticket belongs to the service, the specified ticket is deleted from the database. Giving this parameter is recommended to clean up the storage area for the service. 
	Ticket string `json:"ticket"`
	// The reason of the failure of the backchannel authentication request. This request parameter is not mandatory but optional. However, giving this parameter is recommended. If omitted, `SERVER_ERROR` is used as a reason. 
	Reason string `json:"reason"`
	// The description of the error. This corresponds to the `error_description` property in the response to the client. 
	ErrorDescription *string `json:"errorDescription,omitempty"`
	// The URI of a document which describes the error in detail. If this optional request parameter is given, its value is used as the value of the `error_uri` property.
	ErrorUri *string `json:"errorUri,omitempty"`
}

// NewBackchannelAuthenticationFailRequest instantiates a new BackchannelAuthenticationFailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackchannelAuthenticationFailRequest(ticket string, reason string) *BackchannelAuthenticationFailRequest {
	this := BackchannelAuthenticationFailRequest{}
	this.Ticket = ticket
	this.Reason = reason
	return &this
}

// NewBackchannelAuthenticationFailRequestWithDefaults instantiates a new BackchannelAuthenticationFailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackchannelAuthenticationFailRequestWithDefaults() *BackchannelAuthenticationFailRequest {
	this := BackchannelAuthenticationFailRequest{}
	return &this
}

// GetTicket returns the Ticket field value
func (o *BackchannelAuthenticationFailRequest) GetTicket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ticket
}

// GetTicketOk returns a tuple with the Ticket field value
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationFailRequest) GetTicketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ticket, true
}

// SetTicket sets field value
func (o *BackchannelAuthenticationFailRequest) SetTicket(v string) {
	o.Ticket = v
}

// GetReason returns the Reason field value
func (o *BackchannelAuthenticationFailRequest) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationFailRequest) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *BackchannelAuthenticationFailRequest) SetReason(v string) {
	o.Reason = v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *BackchannelAuthenticationFailRequest) GetErrorDescription() string {
	if o == nil || isNil(o.ErrorDescription) {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationFailRequest) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.ErrorDescription) {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *BackchannelAuthenticationFailRequest) HasErrorDescription() bool {
	if o != nil && !isNil(o.ErrorDescription) {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *BackchannelAuthenticationFailRequest) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorUri returns the ErrorUri field value if set, zero value otherwise.
func (o *BackchannelAuthenticationFailRequest) GetErrorUri() string {
	if o == nil || isNil(o.ErrorUri) {
		var ret string
		return ret
	}
	return *o.ErrorUri
}

// GetErrorUriOk returns a tuple with the ErrorUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationFailRequest) GetErrorUriOk() (*string, bool) {
	if o == nil || isNil(o.ErrorUri) {
		return nil, false
	}
	return o.ErrorUri, true
}

// HasErrorUri returns a boolean if a field has been set.
func (o *BackchannelAuthenticationFailRequest) HasErrorUri() bool {
	if o != nil && !isNil(o.ErrorUri) {
		return true
	}

	return false
}

// SetErrorUri gets a reference to the given string and assigns it to the ErrorUri field.
func (o *BackchannelAuthenticationFailRequest) SetErrorUri(v string) {
	o.ErrorUri = &v
}

func (o BackchannelAuthenticationFailRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackchannelAuthenticationFailRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ticket"] = o.Ticket
	toSerialize["reason"] = o.Reason
	if !isNil(o.ErrorDescription) {
		toSerialize["errorDescription"] = o.ErrorDescription
	}
	if !isNil(o.ErrorUri) {
		toSerialize["errorUri"] = o.ErrorUri
	}
	return toSerialize, nil
}

type NullableBackchannelAuthenticationFailRequest struct {
	value *BackchannelAuthenticationFailRequest
	isSet bool
}

func (v NullableBackchannelAuthenticationFailRequest) Get() *BackchannelAuthenticationFailRequest {
	return v.value
}

func (v *NullableBackchannelAuthenticationFailRequest) Set(val *BackchannelAuthenticationFailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBackchannelAuthenticationFailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBackchannelAuthenticationFailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackchannelAuthenticationFailRequest(val *BackchannelAuthenticationFailRequest) *NullableBackchannelAuthenticationFailRequest {
	return &NullableBackchannelAuthenticationFailRequest{value: val, isSet: true}
}

func (v NullableBackchannelAuthenticationFailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackchannelAuthenticationFailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


