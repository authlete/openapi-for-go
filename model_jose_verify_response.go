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

// checks if the JoseVerifyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JoseVerifyResponse{}

// JoseVerifyResponse struct for JoseVerifyResponse
type JoseVerifyResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The result of the verification on the JOSE object.
	Valid *bool `json:"valid,omitempty"`
	// The result of the signature verification.
	SignatureValid *bool `json:"signatureValid,omitempty"`
	// The list of missing claims.
	MissingClaims []string `json:"missingClaims,omitempty"`
	// The list of invalid claims.
	InvalidClaims []string `json:"invalidClaims,omitempty"`
	// The list of error messages.
	ErrorDescriptions []string `json:"errorDescriptions,omitempty"`
}

// NewJoseVerifyResponse instantiates a new JoseVerifyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJoseVerifyResponse() *JoseVerifyResponse {
	this := JoseVerifyResponse{}
	return &this
}

// NewJoseVerifyResponseWithDefaults instantiates a new JoseVerifyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJoseVerifyResponseWithDefaults() *JoseVerifyResponse {
	this := JoseVerifyResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *JoseVerifyResponse) GetResultCode() string {
	if o == nil || isNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoseVerifyResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *JoseVerifyResponse) HasResultCode() bool {
	if o != nil && !isNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *JoseVerifyResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *JoseVerifyResponse) GetResultMessage() string {
	if o == nil || isNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoseVerifyResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || isNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *JoseVerifyResponse) HasResultMessage() bool {
	if o != nil && !isNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *JoseVerifyResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *JoseVerifyResponse) GetValid() bool {
	if o == nil || isNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoseVerifyResponse) GetValidOk() (*bool, bool) {
	if o == nil || isNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *JoseVerifyResponse) HasValid() bool {
	if o != nil && !isNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *JoseVerifyResponse) SetValid(v bool) {
	o.Valid = &v
}

// GetSignatureValid returns the SignatureValid field value if set, zero value otherwise.
func (o *JoseVerifyResponse) GetSignatureValid() bool {
	if o == nil || isNil(o.SignatureValid) {
		var ret bool
		return ret
	}
	return *o.SignatureValid
}

// GetSignatureValidOk returns a tuple with the SignatureValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoseVerifyResponse) GetSignatureValidOk() (*bool, bool) {
	if o == nil || isNil(o.SignatureValid) {
		return nil, false
	}
	return o.SignatureValid, true
}

// HasSignatureValid returns a boolean if a field has been set.
func (o *JoseVerifyResponse) HasSignatureValid() bool {
	if o != nil && !isNil(o.SignatureValid) {
		return true
	}

	return false
}

// SetSignatureValid gets a reference to the given bool and assigns it to the SignatureValid field.
func (o *JoseVerifyResponse) SetSignatureValid(v bool) {
	o.SignatureValid = &v
}

// GetMissingClaims returns the MissingClaims field value if set, zero value otherwise.
func (o *JoseVerifyResponse) GetMissingClaims() []string {
	if o == nil || isNil(o.MissingClaims) {
		var ret []string
		return ret
	}
	return o.MissingClaims
}

// GetMissingClaimsOk returns a tuple with the MissingClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoseVerifyResponse) GetMissingClaimsOk() ([]string, bool) {
	if o == nil || isNil(o.MissingClaims) {
		return nil, false
	}
	return o.MissingClaims, true
}

// HasMissingClaims returns a boolean if a field has been set.
func (o *JoseVerifyResponse) HasMissingClaims() bool {
	if o != nil && !isNil(o.MissingClaims) {
		return true
	}

	return false
}

// SetMissingClaims gets a reference to the given []string and assigns it to the MissingClaims field.
func (o *JoseVerifyResponse) SetMissingClaims(v []string) {
	o.MissingClaims = v
}

// GetInvalidClaims returns the InvalidClaims field value if set, zero value otherwise.
func (o *JoseVerifyResponse) GetInvalidClaims() []string {
	if o == nil || isNil(o.InvalidClaims) {
		var ret []string
		return ret
	}
	return o.InvalidClaims
}

// GetInvalidClaimsOk returns a tuple with the InvalidClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoseVerifyResponse) GetInvalidClaimsOk() ([]string, bool) {
	if o == nil || isNil(o.InvalidClaims) {
		return nil, false
	}
	return o.InvalidClaims, true
}

// HasInvalidClaims returns a boolean if a field has been set.
func (o *JoseVerifyResponse) HasInvalidClaims() bool {
	if o != nil && !isNil(o.InvalidClaims) {
		return true
	}

	return false
}

// SetInvalidClaims gets a reference to the given []string and assigns it to the InvalidClaims field.
func (o *JoseVerifyResponse) SetInvalidClaims(v []string) {
	o.InvalidClaims = v
}

// GetErrorDescriptions returns the ErrorDescriptions field value if set, zero value otherwise.
func (o *JoseVerifyResponse) GetErrorDescriptions() []string {
	if o == nil || isNil(o.ErrorDescriptions) {
		var ret []string
		return ret
	}
	return o.ErrorDescriptions
}

// GetErrorDescriptionsOk returns a tuple with the ErrorDescriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoseVerifyResponse) GetErrorDescriptionsOk() ([]string, bool) {
	if o == nil || isNil(o.ErrorDescriptions) {
		return nil, false
	}
	return o.ErrorDescriptions, true
}

// HasErrorDescriptions returns a boolean if a field has been set.
func (o *JoseVerifyResponse) HasErrorDescriptions() bool {
	if o != nil && !isNil(o.ErrorDescriptions) {
		return true
	}

	return false
}

// SetErrorDescriptions gets a reference to the given []string and assigns it to the ErrorDescriptions field.
func (o *JoseVerifyResponse) SetErrorDescriptions(v []string) {
	o.ErrorDescriptions = v
}

func (o JoseVerifyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JoseVerifyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResultCode) {
		toSerialize["resultCode"] = o.ResultCode
	}
	if !isNil(o.ResultMessage) {
		toSerialize["resultMessage"] = o.ResultMessage
	}
	if !isNil(o.Valid) {
		toSerialize["valid"] = o.Valid
	}
	if !isNil(o.SignatureValid) {
		toSerialize["signatureValid"] = o.SignatureValid
	}
	if !isNil(o.MissingClaims) {
		toSerialize["missingClaims"] = o.MissingClaims
	}
	if !isNil(o.InvalidClaims) {
		toSerialize["invalidClaims"] = o.InvalidClaims
	}
	if !isNil(o.ErrorDescriptions) {
		toSerialize["errorDescriptions"] = o.ErrorDescriptions
	}
	return toSerialize, nil
}

type NullableJoseVerifyResponse struct {
	value *JoseVerifyResponse
	isSet bool
}

func (v NullableJoseVerifyResponse) Get() *JoseVerifyResponse {
	return v.value
}

func (v *NullableJoseVerifyResponse) Set(val *JoseVerifyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJoseVerifyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJoseVerifyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJoseVerifyResponse(val *JoseVerifyResponse) *NullableJoseVerifyResponse {
	return &NullableJoseVerifyResponse{value: val, isSet: true}
}

func (v NullableJoseVerifyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJoseVerifyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
