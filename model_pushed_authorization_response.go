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

// checks if the PushedAuthorizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushedAuthorizationResponse{}

// PushedAuthorizationResponse struct for PushedAuthorizationResponse
type PushedAuthorizationResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take. Any other value other than \"CREATED\" should be handled as unsuccessful result.
	Action *string `json:"action,omitempty"`
	// The request_uri created to the client to be used as request_uri on the authorize call.
	RequestUri *string `json:"requestUri,omitempty"`
	// The content that the authorization server implementation is to return to the client application.
	ResponseContent *string `json:"responseContent,omitempty"`
}

// NewPushedAuthorizationResponse instantiates a new PushedAuthorizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushedAuthorizationResponse() *PushedAuthorizationResponse {
	this := PushedAuthorizationResponse{}
	return &this
}

// NewPushedAuthorizationResponseWithDefaults instantiates a new PushedAuthorizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushedAuthorizationResponseWithDefaults() *PushedAuthorizationResponse {
	this := PushedAuthorizationResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *PushedAuthorizationResponse) GetResultCode() string {
	if o == nil || isNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *PushedAuthorizationResponse) HasResultCode() bool {
	if o != nil && !isNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *PushedAuthorizationResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *PushedAuthorizationResponse) GetResultMessage() string {
	if o == nil || isNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || isNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *PushedAuthorizationResponse) HasResultMessage() bool {
	if o != nil && !isNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *PushedAuthorizationResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PushedAuthorizationResponse) GetAction() string {
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationResponse) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PushedAuthorizationResponse) HasAction() bool {
	if o != nil && !isNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PushedAuthorizationResponse) SetAction(v string) {
	o.Action = &v
}

// GetRequestUri returns the RequestUri field value if set, zero value otherwise.
func (o *PushedAuthorizationResponse) GetRequestUri() string {
	if o == nil || isNil(o.RequestUri) {
		var ret string
		return ret
	}
	return *o.RequestUri
}

// GetRequestUriOk returns a tuple with the RequestUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationResponse) GetRequestUriOk() (*string, bool) {
	if o == nil || isNil(o.RequestUri) {
		return nil, false
	}
	return o.RequestUri, true
}

// HasRequestUri returns a boolean if a field has been set.
func (o *PushedAuthorizationResponse) HasRequestUri() bool {
	if o != nil && !isNil(o.RequestUri) {
		return true
	}

	return false
}

// SetRequestUri gets a reference to the given string and assigns it to the RequestUri field.
func (o *PushedAuthorizationResponse) SetRequestUri(v string) {
	o.RequestUri = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *PushedAuthorizationResponse) GetResponseContent() string {
	if o == nil || isNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || isNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *PushedAuthorizationResponse) HasResponseContent() bool {
	if o != nil && !isNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *PushedAuthorizationResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

func (o PushedAuthorizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushedAuthorizationResponse) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.RequestUri) {
		toSerialize["requestUri"] = o.RequestUri
	}
	if !isNil(o.ResponseContent) {
		toSerialize["responseContent"] = o.ResponseContent
	}
	return toSerialize, nil
}

type NullablePushedAuthorizationResponse struct {
	value *PushedAuthorizationResponse
	isSet bool
}

func (v NullablePushedAuthorizationResponse) Get() *PushedAuthorizationResponse {
	return v.value
}

func (v *NullablePushedAuthorizationResponse) Set(val *PushedAuthorizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePushedAuthorizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePushedAuthorizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushedAuthorizationResponse(val *PushedAuthorizationResponse) *NullablePushedAuthorizationResponse {
	return &NullablePushedAuthorizationResponse{value: val, isSet: true}
}

func (v NullablePushedAuthorizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushedAuthorizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
