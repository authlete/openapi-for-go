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

// ClientRegistrationResponse struct for ClientRegistrationResponse
type ClientRegistrationResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take. 
	Action *string `json:"action,omitempty"`
	// The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of `action` parameter. 
	ResponseContent *string `json:"responseContent,omitempty"`
	Client *Client `json:"client,omitempty"`
}

// NewClientRegistrationResponse instantiates a new ClientRegistrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientRegistrationResponse() *ClientRegistrationResponse {
	this := ClientRegistrationResponse{}
	return &this
}

// NewClientRegistrationResponseWithDefaults instantiates a new ClientRegistrationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientRegistrationResponseWithDefaults() *ClientRegistrationResponse {
	this := ClientRegistrationResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *ClientRegistrationResponse) GetResultCode() string {
	if o == nil || o.ResultCode == nil {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRegistrationResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || o.ResultCode == nil {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *ClientRegistrationResponse) HasResultCode() bool {
	if o != nil && o.ResultCode != nil {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *ClientRegistrationResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *ClientRegistrationResponse) GetResultMessage() string {
	if o == nil || o.ResultMessage == nil {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRegistrationResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || o.ResultMessage == nil {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *ClientRegistrationResponse) HasResultMessage() bool {
	if o != nil && o.ResultMessage != nil {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *ClientRegistrationResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ClientRegistrationResponse) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRegistrationResponse) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ClientRegistrationResponse) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ClientRegistrationResponse) SetAction(v string) {
	o.Action = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *ClientRegistrationResponse) GetResponseContent() string {
	if o == nil || o.ResponseContent == nil {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRegistrationResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || o.ResponseContent == nil {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *ClientRegistrationResponse) HasResponseContent() bool {
	if o != nil && o.ResponseContent != nil {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *ClientRegistrationResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *ClientRegistrationResponse) GetClient() Client {
	if o == nil || o.Client == nil {
		var ret Client
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRegistrationResponse) GetClientOk() (*Client, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *ClientRegistrationResponse) HasClient() bool {
	if o != nil && o.Client != nil {
		return true
	}

	return false
}

// SetClient gets a reference to the given Client and assigns it to the Client field.
func (o *ClientRegistrationResponse) SetClient(v Client) {
	o.Client = &v
}

func (o ClientRegistrationResponse) MarshalJSON() ([]byte, error) {
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
	if o.Client != nil {
		toSerialize["client"] = o.Client
	}
	return json.Marshal(toSerialize)
}

type NullableClientRegistrationResponse struct {
	value *ClientRegistrationResponse
	isSet bool
}

func (v NullableClientRegistrationResponse) Get() *ClientRegistrationResponse {
	return v.value
}

func (v *NullableClientRegistrationResponse) Set(val *ClientRegistrationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClientRegistrationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClientRegistrationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientRegistrationResponse(val *ClientRegistrationResponse) *NullableClientRegistrationResponse {
	return &NullableClientRegistrationResponse{value: val, isSet: true}
}

func (v NullableClientRegistrationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientRegistrationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


