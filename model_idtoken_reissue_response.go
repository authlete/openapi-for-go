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

// checks if the IdtokenReissueResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdtokenReissueResponse{}

// IdtokenReissueResponse struct for IdtokenReissueResponse
type IdtokenReissueResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the implementation of the token endpoint should take.
	Action *string `json:"action,omitempty"`
	// The response content that can be used as the message body of the token response that should be returned from the token endpoint. 
	ResponseContent *string `json:"responseContent,omitempty"`
	// The reissued ID token
	IdToken *string `json:"idToken,omitempty"`
}

// NewIdtokenReissueResponse instantiates a new IdtokenReissueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdtokenReissueResponse() *IdtokenReissueResponse {
	this := IdtokenReissueResponse{}
	return &this
}

// NewIdtokenReissueResponseWithDefaults instantiates a new IdtokenReissueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdtokenReissueResponseWithDefaults() *IdtokenReissueResponse {
	this := IdtokenReissueResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *IdtokenReissueResponse) GetResultCode() string {
	if o == nil || isNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdtokenReissueResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *IdtokenReissueResponse) HasResultCode() bool {
	if o != nil && !isNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *IdtokenReissueResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *IdtokenReissueResponse) GetResultMessage() string {
	if o == nil || isNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdtokenReissueResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || isNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *IdtokenReissueResponse) HasResultMessage() bool {
	if o != nil && !isNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *IdtokenReissueResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *IdtokenReissueResponse) GetAction() string {
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdtokenReissueResponse) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *IdtokenReissueResponse) HasAction() bool {
	if o != nil && !isNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *IdtokenReissueResponse) SetAction(v string) {
	o.Action = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *IdtokenReissueResponse) GetResponseContent() string {
	if o == nil || isNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdtokenReissueResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || isNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *IdtokenReissueResponse) HasResponseContent() bool {
	if o != nil && !isNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *IdtokenReissueResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise.
func (o *IdtokenReissueResponse) GetIdToken() string {
	if o == nil || isNil(o.IdToken) {
		var ret string
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdtokenReissueResponse) GetIdTokenOk() (*string, bool) {
	if o == nil || isNil(o.IdToken) {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *IdtokenReissueResponse) HasIdToken() bool {
	if o != nil && !isNil(o.IdToken) {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given string and assigns it to the IdToken field.
func (o *IdtokenReissueResponse) SetIdToken(v string) {
	o.IdToken = &v
}

func (o IdtokenReissueResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdtokenReissueResponse) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.IdToken) {
		toSerialize["idToken"] = o.IdToken
	}
	return toSerialize, nil
}

type NullableIdtokenReissueResponse struct {
	value *IdtokenReissueResponse
	isSet bool
}

func (v NullableIdtokenReissueResponse) Get() *IdtokenReissueResponse {
	return v.value
}

func (v *NullableIdtokenReissueResponse) Set(val *IdtokenReissueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIdtokenReissueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIdtokenReissueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdtokenReissueResponse(val *IdtokenReissueResponse) *NullableIdtokenReissueResponse {
	return &NullableIdtokenReissueResponse{value: val, isSet: true}
}

func (v NullableIdtokenReissueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdtokenReissueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

