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

// checks if the VciSingleParseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VciSingleParseResponse{}

// VciSingleParseResponse struct for VciSingleParseResponse
type VciSingleParseResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the credential endpoint should take.
	Action *string `json:"action,omitempty"`
	// The content of the response to the request sender.
	ResponseContent *string `json:"responseContent,omitempty"`
	Info *CredentialRequestInfo `json:"info,omitempty"`
}

// NewVciSingleParseResponse instantiates a new VciSingleParseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVciSingleParseResponse() *VciSingleParseResponse {
	this := VciSingleParseResponse{}
	return &this
}

// NewVciSingleParseResponseWithDefaults instantiates a new VciSingleParseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVciSingleParseResponseWithDefaults() *VciSingleParseResponse {
	this := VciSingleParseResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *VciSingleParseResponse) GetResultCode() string {
	if o == nil || isNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciSingleParseResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *VciSingleParseResponse) HasResultCode() bool {
	if o != nil && !isNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *VciSingleParseResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *VciSingleParseResponse) GetResultMessage() string {
	if o == nil || isNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciSingleParseResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || isNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *VciSingleParseResponse) HasResultMessage() bool {
	if o != nil && !isNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *VciSingleParseResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *VciSingleParseResponse) GetAction() string {
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciSingleParseResponse) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *VciSingleParseResponse) HasAction() bool {
	if o != nil && !isNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *VciSingleParseResponse) SetAction(v string) {
	o.Action = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *VciSingleParseResponse) GetResponseContent() string {
	if o == nil || isNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciSingleParseResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || isNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *VciSingleParseResponse) HasResponseContent() bool {
	if o != nil && !isNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *VciSingleParseResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *VciSingleParseResponse) GetInfo() CredentialRequestInfo {
	if o == nil || isNil(o.Info) {
		var ret CredentialRequestInfo
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciSingleParseResponse) GetInfoOk() (*CredentialRequestInfo, bool) {
	if o == nil || isNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *VciSingleParseResponse) HasInfo() bool {
	if o != nil && !isNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given CredentialRequestInfo and assigns it to the Info field.
func (o *VciSingleParseResponse) SetInfo(v CredentialRequestInfo) {
	o.Info = &v
}

func (o VciSingleParseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VciSingleParseResponse) ToMap() (map[string]interface{}, error) {
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

type NullableVciSingleParseResponse struct {
	value *VciSingleParseResponse
	isSet bool
}

func (v NullableVciSingleParseResponse) Get() *VciSingleParseResponse {
	return v.value
}

func (v *NullableVciSingleParseResponse) Set(val *VciSingleParseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVciSingleParseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVciSingleParseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVciSingleParseResponse(val *VciSingleParseResponse) *NullableVciSingleParseResponse {
	return &NullableVciSingleParseResponse{value: val, isSet: true}
}

func (v NullableVciSingleParseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVciSingleParseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


