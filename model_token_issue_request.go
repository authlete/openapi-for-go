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

// checks if the TokenIssueRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenIssueRequest{}

// TokenIssueRequest struct for TokenIssueRequest
type TokenIssueRequest struct {
	// The ticket issued from Authlete `/auth/token` API. 
	Ticket string `json:"ticket"`
	// The subject (= unique identifier) of the authenticated user. 
	Subject string `json:"subject"`
	// Extra properties to associate with a newly created access token. Note that properties parameter is accepted only when `Content-Type` of the request is `application/json`, so don't use `application/x-www-form-urlencoded` if you want to specify properties. 
	Properties []Property `json:"properties,omitempty"`
	// The representation of an access token that may be issued as a result of the Authlete API call.
	AccessToken *string `json:"accessToken,omitempty"`
}

// NewTokenIssueRequest instantiates a new TokenIssueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenIssueRequest(ticket string, subject string) *TokenIssueRequest {
	this := TokenIssueRequest{}
	this.Ticket = ticket
	this.Subject = subject
	return &this
}

// NewTokenIssueRequestWithDefaults instantiates a new TokenIssueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenIssueRequestWithDefaults() *TokenIssueRequest {
	this := TokenIssueRequest{}
	return &this
}

// GetTicket returns the Ticket field value
func (o *TokenIssueRequest) GetTicket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ticket
}

// GetTicketOk returns a tuple with the Ticket field value
// and a boolean to check if the value has been set.
func (o *TokenIssueRequest) GetTicketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ticket, true
}

// SetTicket sets field value
func (o *TokenIssueRequest) SetTicket(v string) {
	o.Ticket = v
}

// GetSubject returns the Subject field value
func (o *TokenIssueRequest) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *TokenIssueRequest) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *TokenIssueRequest) SetSubject(v string) {
	o.Subject = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *TokenIssueRequest) GetProperties() []Property {
	if o == nil || isNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueRequest) GetPropertiesOk() ([]Property, bool) {
	if o == nil || isNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *TokenIssueRequest) HasProperties() bool {
	if o != nil && !isNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *TokenIssueRequest) SetProperties(v []Property) {
	o.Properties = v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *TokenIssueRequest) GetAccessToken() string {
	if o == nil || isNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenIssueRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil || isNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *TokenIssueRequest) HasAccessToken() bool {
	if o != nil && !isNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *TokenIssueRequest) SetAccessToken(v string) {
	o.AccessToken = &v
}

func (o TokenIssueRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenIssueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ticket"] = o.Ticket
	toSerialize["subject"] = o.Subject
	if !isNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !isNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	return toSerialize, nil
}

type NullableTokenIssueRequest struct {
	value *TokenIssueRequest
	isSet bool
}

func (v NullableTokenIssueRequest) Get() *TokenIssueRequest {
	return v.value
}

func (v *NullableTokenIssueRequest) Set(val *TokenIssueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenIssueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenIssueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenIssueRequest(val *TokenIssueRequest) *NullableTokenIssueRequest {
	return &NullableTokenIssueRequest{value: val, isSet: true}
}

func (v NullableTokenIssueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenIssueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


