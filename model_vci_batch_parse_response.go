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

// checks if the VciBatchParseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VciBatchParseResponse{}

// VciBatchParseResponse struct for VciBatchParseResponse
type VciBatchParseResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the batch credential endpoint should take.
	Action *string `json:"action,omitempty"`
	// The content of the response to the request sender.
	ResponseContent *string `json:"responseContent,omitempty"`
	// Information about the credential requests in the batch credential request. 
	Info []CredentialRequestInfo `json:"info,omitempty"`
}

// NewVciBatchParseResponse instantiates a new VciBatchParseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVciBatchParseResponse() *VciBatchParseResponse {
	this := VciBatchParseResponse{}
	return &this
}

// NewVciBatchParseResponseWithDefaults instantiates a new VciBatchParseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVciBatchParseResponseWithDefaults() *VciBatchParseResponse {
	this := VciBatchParseResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *VciBatchParseResponse) GetResultCode() string {
	if o == nil || isNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciBatchParseResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *VciBatchParseResponse) HasResultCode() bool {
	if o != nil && !isNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *VciBatchParseResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *VciBatchParseResponse) GetResultMessage() string {
	if o == nil || isNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciBatchParseResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || isNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *VciBatchParseResponse) HasResultMessage() bool {
	if o != nil && !isNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *VciBatchParseResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *VciBatchParseResponse) GetAction() string {
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciBatchParseResponse) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *VciBatchParseResponse) HasAction() bool {
	if o != nil && !isNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *VciBatchParseResponse) SetAction(v string) {
	o.Action = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *VciBatchParseResponse) GetResponseContent() string {
	if o == nil || isNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciBatchParseResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || isNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *VciBatchParseResponse) HasResponseContent() bool {
	if o != nil && !isNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *VciBatchParseResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *VciBatchParseResponse) GetInfo() []CredentialRequestInfo {
	if o == nil || isNil(o.Info) {
		var ret []CredentialRequestInfo
		return ret
	}
	return o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciBatchParseResponse) GetInfoOk() ([]CredentialRequestInfo, bool) {
	if o == nil || isNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *VciBatchParseResponse) HasInfo() bool {
	if o != nil && !isNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given []CredentialRequestInfo and assigns it to the Info field.
func (o *VciBatchParseResponse) SetInfo(v []CredentialRequestInfo) {
	o.Info = v
}

func (o VciBatchParseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VciBatchParseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResultCode) {
		toSerialize["resultCode"] = o.ResultCode
	}
	if !isNil(o.ResultMessage) {
		toSerialize["resultMessage"] = o.ResultMessage
	}
	if !isNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !isNil(o.ResponseContent) {
		toSerialize["responseContent"] = o.ResponseContent
	}
	if !isNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	return toSerialize, nil
}

type NullableVciBatchParseResponse struct {
	value *VciBatchParseResponse
	isSet bool
}

func (v NullableVciBatchParseResponse) Get() *VciBatchParseResponse {
	return v.value
}

func (v *NullableVciBatchParseResponse) Set(val *VciBatchParseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVciBatchParseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVciBatchParseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVciBatchParseResponse(val *VciBatchParseResponse) *NullableVciBatchParseResponse {
	return &NullableVciBatchParseResponse{value: val, isSet: true}
}

func (v NullableVciBatchParseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVciBatchParseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


