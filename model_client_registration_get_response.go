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

// checks if the ClientRegistrationGetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientRegistrationGetResponse{}

// ClientRegistrationGetResponse struct for ClientRegistrationGetResponse
type ClientRegistrationGetResponse struct {
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

// NewClientRegistrationGetResponse instantiates a new ClientRegistrationGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientRegistrationGetResponse() *ClientRegistrationGetResponse {
	this := ClientRegistrationGetResponse{}
	return &this
}

// NewClientRegistrationGetResponseWithDefaults instantiates a new ClientRegistrationGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientRegistrationGetResponseWithDefaults() *ClientRegistrationGetResponse {
	this := ClientRegistrationGetResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *ClientRegistrationGetResponse) GetResultCode() string {
	if o == nil || isNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRegistrationGetResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *ClientRegistrationGetResponse) HasResultCode() bool {
	if o != nil && !isNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *ClientRegistrationGetResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *ClientRegistrationGetResponse) GetResultMessage() string {
	if o == nil || isNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRegistrationGetResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || isNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *ClientRegistrationGetResponse) HasResultMessage() bool {
	if o != nil && !isNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *ClientRegistrationGetResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ClientRegistrationGetResponse) GetAction() string {
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRegistrationGetResponse) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ClientRegistrationGetResponse) HasAction() bool {
	if o != nil && !isNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ClientRegistrationGetResponse) SetAction(v string) {
	o.Action = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *ClientRegistrationGetResponse) GetResponseContent() string {
	if o == nil || isNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRegistrationGetResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || isNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *ClientRegistrationGetResponse) HasResponseContent() bool {
	if o != nil && !isNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *ClientRegistrationGetResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *ClientRegistrationGetResponse) GetClient() Client {
	if o == nil || isNil(o.Client) {
		var ret Client
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRegistrationGetResponse) GetClientOk() (*Client, bool) {
	if o == nil || isNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *ClientRegistrationGetResponse) HasClient() bool {
	if o != nil && !isNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given Client and assigns it to the Client field.
func (o *ClientRegistrationGetResponse) SetClient(v Client) {
	o.Client = &v
}

func (o ClientRegistrationGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientRegistrationGetResponse) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	return toSerialize, nil
}

type NullableClientRegistrationGetResponse struct {
	value *ClientRegistrationGetResponse
	isSet bool
}

func (v NullableClientRegistrationGetResponse) Get() *ClientRegistrationGetResponse {
	return v.value
}

func (v *NullableClientRegistrationGetResponse) Set(val *ClientRegistrationGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClientRegistrationGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClientRegistrationGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientRegistrationGetResponse(val *ClientRegistrationGetResponse) *NullableClientRegistrationGetResponse {
	return &NullableClientRegistrationGetResponse{value: val, isSet: true}
}

func (v NullableClientRegistrationGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientRegistrationGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


