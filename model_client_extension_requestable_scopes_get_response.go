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

// checks if the ClientExtensionRequestableScopesGetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientExtensionRequestableScopesGetResponse{}

// ClientExtensionRequestableScopesGetResponse struct for ClientExtensionRequestableScopesGetResponse
type ClientExtensionRequestableScopesGetResponse struct {
	RequestableScopes []string `json:"requestableScopes,omitempty"`
}

// NewClientExtensionRequestableScopesGetResponse instantiates a new ClientExtensionRequestableScopesGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientExtensionRequestableScopesGetResponse() *ClientExtensionRequestableScopesGetResponse {
	this := ClientExtensionRequestableScopesGetResponse{}
	return &this
}

// NewClientExtensionRequestableScopesGetResponseWithDefaults instantiates a new ClientExtensionRequestableScopesGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientExtensionRequestableScopesGetResponseWithDefaults() *ClientExtensionRequestableScopesGetResponse {
	this := ClientExtensionRequestableScopesGetResponse{}
	return &this
}

// GetRequestableScopes returns the RequestableScopes field value if set, zero value otherwise.
func (o *ClientExtensionRequestableScopesGetResponse) GetRequestableScopes() []string {
	if o == nil || isNil(o.RequestableScopes) {
		var ret []string
		return ret
	}
	return o.RequestableScopes
}

// GetRequestableScopesOk returns a tuple with the RequestableScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientExtensionRequestableScopesGetResponse) GetRequestableScopesOk() ([]string, bool) {
	if o == nil || isNil(o.RequestableScopes) {
		return nil, false
	}
	return o.RequestableScopes, true
}

// HasRequestableScopes returns a boolean if a field has been set.
func (o *ClientExtensionRequestableScopesGetResponse) HasRequestableScopes() bool {
	if o != nil && !isNil(o.RequestableScopes) {
		return true
	}

	return false
}

// SetRequestableScopes gets a reference to the given []string and assigns it to the RequestableScopes field.
func (o *ClientExtensionRequestableScopesGetResponse) SetRequestableScopes(v []string) {
	o.RequestableScopes = v
}

func (o ClientExtensionRequestableScopesGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientExtensionRequestableScopesGetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RequestableScopes) {
		toSerialize["requestableScopes"] = o.RequestableScopes
	}
	return toSerialize, nil
}

type NullableClientExtensionRequestableScopesGetResponse struct {
	value *ClientExtensionRequestableScopesGetResponse
	isSet bool
}

func (v NullableClientExtensionRequestableScopesGetResponse) Get() *ClientExtensionRequestableScopesGetResponse {
	return v.value
}

func (v *NullableClientExtensionRequestableScopesGetResponse) Set(val *ClientExtensionRequestableScopesGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClientExtensionRequestableScopesGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClientExtensionRequestableScopesGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientExtensionRequestableScopesGetResponse(val *ClientExtensionRequestableScopesGetResponse) *NullableClientExtensionRequestableScopesGetResponse {
	return &NullableClientExtensionRequestableScopesGetResponse{value: val, isSet: true}
}

func (v NullableClientExtensionRequestableScopesGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientExtensionRequestableScopesGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
