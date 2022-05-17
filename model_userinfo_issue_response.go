/*
Authlete API

Authlete API Document. 

API version: 2.2.19
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

// UserinfoIssueResponse 
type UserinfoIssueResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take.
	Action *string `json:"action,omitempty"`
	// The content that the authorization server implementation can use as the value of `WWW-Authenticate` header on errors. 
	ResponseContent *string `json:"responseContent,omitempty"`
}

// NewUserinfoIssueResponse instantiates a new UserinfoIssueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserinfoIssueResponse() *UserinfoIssueResponse {
	this := UserinfoIssueResponse{}
	return &this
}

// NewUserinfoIssueResponseWithDefaults instantiates a new UserinfoIssueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserinfoIssueResponseWithDefaults() *UserinfoIssueResponse {
	this := UserinfoIssueResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *UserinfoIssueResponse) GetResultCode() string {
	if o == nil || o.ResultCode == nil {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || o.ResultCode == nil {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *UserinfoIssueResponse) HasResultCode() bool {
	if o != nil && o.ResultCode != nil {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *UserinfoIssueResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *UserinfoIssueResponse) GetResultMessage() string {
	if o == nil || o.ResultMessage == nil {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || o.ResultMessage == nil {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *UserinfoIssueResponse) HasResultMessage() bool {
	if o != nil && o.ResultMessage != nil {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *UserinfoIssueResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *UserinfoIssueResponse) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueResponse) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *UserinfoIssueResponse) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *UserinfoIssueResponse) SetAction(v string) {
	o.Action = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *UserinfoIssueResponse) GetResponseContent() string {
	if o == nil || o.ResponseContent == nil {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || o.ResponseContent == nil {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *UserinfoIssueResponse) HasResponseContent() bool {
	if o != nil && o.ResponseContent != nil {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *UserinfoIssueResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

func (o UserinfoIssueResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResultCode != nil {
		toSerialize["resultCode"] = o.ResultCode
	}
	if o.ResultMessage != nil {
		toSerialize["resultMessage"] = o.ResultMessage
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.ResponseContent != nil {
		toSerialize["responseContent"] = o.ResponseContent
	}
	return json.Marshal(toSerialize)
}

type NullableUserinfoIssueResponse struct {
	value *UserinfoIssueResponse
	isSet bool
}

func (v NullableUserinfoIssueResponse) Get() *UserinfoIssueResponse {
	return v.value
}

func (v *NullableUserinfoIssueResponse) Set(val *UserinfoIssueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserinfoIssueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserinfoIssueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserinfoIssueResponse(val *UserinfoIssueResponse) *NullableUserinfoIssueResponse {
	return &NullableUserinfoIssueResponse{value: val, isSet: true}
}

func (v NullableUserinfoIssueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserinfoIssueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


