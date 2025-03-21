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

// checks if the ClientExtensionRequestableScopesUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientExtensionRequestableScopesUpdateRequest{}

// ClientExtensionRequestableScopesUpdateRequest struct for ClientExtensionRequestableScopesUpdateRequest
type ClientExtensionRequestableScopesUpdateRequest struct {
	// The set of scopes that the client application is allowed to request. This parameter will be one of the following. Details are described in the description.   - an empty set - a set with at least one element  If this parameter contains scopes that the service does not support, those scopes are just ignored. Also, if this parameter is `null` or is not included in the request, it is equivalent to calling `/client/extension/requestable_scopes/delete` API.
	RequestableScopes []string `json:"requestableScopes,omitempty"`
}

// NewClientExtensionRequestableScopesUpdateRequest instantiates a new ClientExtensionRequestableScopesUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientExtensionRequestableScopesUpdateRequest() *ClientExtensionRequestableScopesUpdateRequest {
	this := ClientExtensionRequestableScopesUpdateRequest{}
	return &this
}

// NewClientExtensionRequestableScopesUpdateRequestWithDefaults instantiates a new ClientExtensionRequestableScopesUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientExtensionRequestableScopesUpdateRequestWithDefaults() *ClientExtensionRequestableScopesUpdateRequest {
	this := ClientExtensionRequestableScopesUpdateRequest{}
	return &this
}

// GetRequestableScopes returns the RequestableScopes field value if set, zero value otherwise.
func (o *ClientExtensionRequestableScopesUpdateRequest) GetRequestableScopes() []string {
	if o == nil || isNil(o.RequestableScopes) {
		var ret []string
		return ret
	}
	return o.RequestableScopes
}

// GetRequestableScopesOk returns a tuple with the RequestableScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientExtensionRequestableScopesUpdateRequest) GetRequestableScopesOk() ([]string, bool) {
	if o == nil || isNil(o.RequestableScopes) {
		return nil, false
	}
	return o.RequestableScopes, true
}

// HasRequestableScopes returns a boolean if a field has been set.
func (o *ClientExtensionRequestableScopesUpdateRequest) HasRequestableScopes() bool {
	if o != nil && !isNil(o.RequestableScopes) {
		return true
	}

	return false
}

// SetRequestableScopes gets a reference to the given []string and assigns it to the RequestableScopes field.
func (o *ClientExtensionRequestableScopesUpdateRequest) SetRequestableScopes(v []string) {
	o.RequestableScopes = v
}

func (o ClientExtensionRequestableScopesUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientExtensionRequestableScopesUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RequestableScopes) {
		toSerialize["requestableScopes"] = o.RequestableScopes
	}
	return toSerialize, nil
}

type NullableClientExtensionRequestableScopesUpdateRequest struct {
	value *ClientExtensionRequestableScopesUpdateRequest
	isSet bool
}

func (v NullableClientExtensionRequestableScopesUpdateRequest) Get() *ClientExtensionRequestableScopesUpdateRequest {
	return v.value
}

func (v *NullableClientExtensionRequestableScopesUpdateRequest) Set(val *ClientExtensionRequestableScopesUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClientExtensionRequestableScopesUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClientExtensionRequestableScopesUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientExtensionRequestableScopesUpdateRequest(val *ClientExtensionRequestableScopesUpdateRequest) *NullableClientExtensionRequestableScopesUpdateRequest {
	return &NullableClientExtensionRequestableScopesUpdateRequest{value: val, isSet: true}
}

func (v NullableClientExtensionRequestableScopesUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientExtensionRequestableScopesUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
