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

// checks if the ClientAuthorizationDeleteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientAuthorizationDeleteResponse{}

// ClientAuthorizationDeleteResponse struct for ClientAuthorizationDeleteResponse
type ClientAuthorizationDeleteResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
}

// NewClientAuthorizationDeleteResponse instantiates a new ClientAuthorizationDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientAuthorizationDeleteResponse() *ClientAuthorizationDeleteResponse {
	this := ClientAuthorizationDeleteResponse{}
	return &this
}

// NewClientAuthorizationDeleteResponseWithDefaults instantiates a new ClientAuthorizationDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientAuthorizationDeleteResponseWithDefaults() *ClientAuthorizationDeleteResponse {
	this := ClientAuthorizationDeleteResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *ClientAuthorizationDeleteResponse) GetResultCode() string {
	if o == nil || isNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuthorizationDeleteResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *ClientAuthorizationDeleteResponse) HasResultCode() bool {
	if o != nil && !isNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *ClientAuthorizationDeleteResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *ClientAuthorizationDeleteResponse) GetResultMessage() string {
	if o == nil || isNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuthorizationDeleteResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || isNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *ClientAuthorizationDeleteResponse) HasResultMessage() bool {
	if o != nil && !isNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *ClientAuthorizationDeleteResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

func (o ClientAuthorizationDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientAuthorizationDeleteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResultCode) {
		toSerialize["resultCode"] = o.ResultCode
	}
	if !isNil(o.ResultMessage) {
		toSerialize["resultMessage"] = o.ResultMessage
	}
	return toSerialize, nil
}

type NullableClientAuthorizationDeleteResponse struct {
	value *ClientAuthorizationDeleteResponse
	isSet bool
}

func (v NullableClientAuthorizationDeleteResponse) Get() *ClientAuthorizationDeleteResponse {
	return v.value
}

func (v *NullableClientAuthorizationDeleteResponse) Set(val *ClientAuthorizationDeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClientAuthorizationDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClientAuthorizationDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientAuthorizationDeleteResponse(val *ClientAuthorizationDeleteResponse) *NullableClientAuthorizationDeleteResponse {
	return &NullableClientAuthorizationDeleteResponse{value: val, isSet: true}
}

func (v NullableClientAuthorizationDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientAuthorizationDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
