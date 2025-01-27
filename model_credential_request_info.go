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

// checks if the CredentialRequestInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialRequestInfo{}

// CredentialRequestInfo struct for CredentialRequestInfo
type CredentialRequestInfo struct {
	// The identifier of the credential offer.
	Identifier *string `json:"identifier,omitempty"`
	// The value of the format parameter in the credential request.
	Format *string `json:"format,omitempty"`
	// The binding key specified by the proof in the credential request.
	BindingKey *string `json:"bindingKey,omitempty"`
	// The details about the credential request.
	Details *string `json:"details,omitempty"`
}

// NewCredentialRequestInfo instantiates a new CredentialRequestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialRequestInfo() *CredentialRequestInfo {
	this := CredentialRequestInfo{}
	return &this
}

// NewCredentialRequestInfoWithDefaults instantiates a new CredentialRequestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialRequestInfoWithDefaults() *CredentialRequestInfo {
	this := CredentialRequestInfo{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *CredentialRequestInfo) GetIdentifier() string {
	if o == nil || isNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialRequestInfo) GetIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *CredentialRequestInfo) HasIdentifier() bool {
	if o != nil && !isNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *CredentialRequestInfo) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *CredentialRequestInfo) GetFormat() string {
	if o == nil || isNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialRequestInfo) GetFormatOk() (*string, bool) {
	if o == nil || isNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *CredentialRequestInfo) HasFormat() bool {
	if o != nil && !isNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *CredentialRequestInfo) SetFormat(v string) {
	o.Format = &v
}

// GetBindingKey returns the BindingKey field value if set, zero value otherwise.
func (o *CredentialRequestInfo) GetBindingKey() string {
	if o == nil || isNil(o.BindingKey) {
		var ret string
		return ret
	}
	return *o.BindingKey
}

// GetBindingKeyOk returns a tuple with the BindingKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialRequestInfo) GetBindingKeyOk() (*string, bool) {
	if o == nil || isNil(o.BindingKey) {
		return nil, false
	}
	return o.BindingKey, true
}

// HasBindingKey returns a boolean if a field has been set.
func (o *CredentialRequestInfo) HasBindingKey() bool {
	if o != nil && !isNil(o.BindingKey) {
		return true
	}

	return false
}

// SetBindingKey gets a reference to the given string and assigns it to the BindingKey field.
func (o *CredentialRequestInfo) SetBindingKey(v string) {
	o.BindingKey = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *CredentialRequestInfo) GetDetails() string {
	if o == nil || isNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialRequestInfo) GetDetailsOk() (*string, bool) {
	if o == nil || isNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *CredentialRequestInfo) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *CredentialRequestInfo) SetDetails(v string) {
	o.Details = &v
}

func (o CredentialRequestInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialRequestInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !isNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !isNil(o.BindingKey) {
		toSerialize["bindingKey"] = o.BindingKey
	}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableCredentialRequestInfo struct {
	value *CredentialRequestInfo
	isSet bool
}

func (v NullableCredentialRequestInfo) Get() *CredentialRequestInfo {
	return v.value
}

func (v *NullableCredentialRequestInfo) Set(val *CredentialRequestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialRequestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialRequestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialRequestInfo(val *CredentialRequestInfo) *NullableCredentialRequestInfo {
	return &NullableCredentialRequestInfo{value: val, isSet: true}
}

func (v NullableCredentialRequestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialRequestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


