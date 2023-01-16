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

// checks if the StandardIntrospectionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardIntrospectionRequest{}

// StandardIntrospectionRequest struct for StandardIntrospectionRequest
type StandardIntrospectionRequest struct {
	// Request parameters which comply with the introspection request defined in \"[2.1. Introspection Request](https://datatracker.ietf.org/doc/html/rfc7662#section-2.1)\" in RFC 7662.  The implementation of the introspection endpoint of your authorization server will receive an HTTP POST [[RFC 7231](https://datatracker.ietf.org/doc/html/rfc7231)] request with parameters in the `application/x-www-form-urlencoded` format. It is the entity body of the request that Authlete's  `/api/auth/introspection/standard` API expects as the value of `parameters`. 
	Parameters string `json:"parameters"`
	// Flag indicating whether to include hidden properties in the output.  Authlete has a mechanism whereby to associate arbitrary key-value pairs with an access token. Each key-value pair has a hidden attribute. By default, key-value pairs whose hidden attribute is set to `true` are not embedded in the standard introspection output.  If the `withHiddenProperties` request parameter is given and its value is `true`, `/api/auth/introspection/standard API includes all the associated key-value pairs into the output regardless of the value of the hidden attribute.
	WithHiddenProperties *string `json:"withHiddenProperties,omitempty"`
}

// NewStandardIntrospectionRequest instantiates a new StandardIntrospectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardIntrospectionRequest(parameters string) *StandardIntrospectionRequest {
	this := StandardIntrospectionRequest{}
	this.Parameters = parameters
	return &this
}

// NewStandardIntrospectionRequestWithDefaults instantiates a new StandardIntrospectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardIntrospectionRequestWithDefaults() *StandardIntrospectionRequest {
	this := StandardIntrospectionRequest{}
	return &this
}

// GetParameters returns the Parameters field value
func (o *StandardIntrospectionRequest) GetParameters() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *StandardIntrospectionRequest) GetParametersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *StandardIntrospectionRequest) SetParameters(v string) {
	o.Parameters = v
}

// GetWithHiddenProperties returns the WithHiddenProperties field value if set, zero value otherwise.
func (o *StandardIntrospectionRequest) GetWithHiddenProperties() string {
	if o == nil || isNil(o.WithHiddenProperties) {
		var ret string
		return ret
	}
	return *o.WithHiddenProperties
}

// GetWithHiddenPropertiesOk returns a tuple with the WithHiddenProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardIntrospectionRequest) GetWithHiddenPropertiesOk() (*string, bool) {
	if o == nil || isNil(o.WithHiddenProperties) {
		return nil, false
	}
	return o.WithHiddenProperties, true
}

// HasWithHiddenProperties returns a boolean if a field has been set.
func (o *StandardIntrospectionRequest) HasWithHiddenProperties() bool {
	if o != nil && !isNil(o.WithHiddenProperties) {
		return true
	}

	return false
}

// SetWithHiddenProperties gets a reference to the given string and assigns it to the WithHiddenProperties field.
func (o *StandardIntrospectionRequest) SetWithHiddenProperties(v string) {
	o.WithHiddenProperties = &v
}

func (o StandardIntrospectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandardIntrospectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parameters"] = o.Parameters
	if !isNil(o.WithHiddenProperties) {
		toSerialize["withHiddenProperties"] = o.WithHiddenProperties
	}
	return toSerialize, nil
}

type NullableStandardIntrospectionRequest struct {
	value *StandardIntrospectionRequest
	isSet bool
}

func (v NullableStandardIntrospectionRequest) Get() *StandardIntrospectionRequest {
	return v.value
}

func (v *NullableStandardIntrospectionRequest) Set(val *StandardIntrospectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardIntrospectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardIntrospectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardIntrospectionRequest(val *StandardIntrospectionRequest) *NullableStandardIntrospectionRequest {
	return &NullableStandardIntrospectionRequest{value: val, isSet: true}
}

func (v NullableStandardIntrospectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardIntrospectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


