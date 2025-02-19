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

// checks if the BackchannelAuthenticationIssueResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackchannelAuthenticationIssueResponse{}

// BackchannelAuthenticationIssueResponse struct for BackchannelAuthenticationIssueResponse
type BackchannelAuthenticationIssueResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take.
	Action *string `json:"action,omitempty"`
	// The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of `action` parameter.
	ResponseContent *string `json:"responseContent,omitempty"`
	// The newly issued authentication request ID.
	AuthReqId *string `json:"authReqId,omitempty"`
	// The duration of the issued authentication request ID in seconds.
	ExpiresIn *int32 `json:"expiresIn,omitempty"`
	// The minimum amount of time in seconds that the client must wait for between polling requests to the token endpoint.
	Interval *int32 `json:"interval,omitempty"`
}

// NewBackchannelAuthenticationIssueResponse instantiates a new BackchannelAuthenticationIssueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackchannelAuthenticationIssueResponse() *BackchannelAuthenticationIssueResponse {
	this := BackchannelAuthenticationIssueResponse{}
	return &this
}

// NewBackchannelAuthenticationIssueResponseWithDefaults instantiates a new BackchannelAuthenticationIssueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackchannelAuthenticationIssueResponseWithDefaults() *BackchannelAuthenticationIssueResponse {
	this := BackchannelAuthenticationIssueResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *BackchannelAuthenticationIssueResponse) GetResultCode() string {
	if o == nil || isNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationIssueResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *BackchannelAuthenticationIssueResponse) HasResultCode() bool {
	if o != nil && !isNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *BackchannelAuthenticationIssueResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *BackchannelAuthenticationIssueResponse) GetResultMessage() string {
	if o == nil || isNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationIssueResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || isNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *BackchannelAuthenticationIssueResponse) HasResultMessage() bool {
	if o != nil && !isNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *BackchannelAuthenticationIssueResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BackchannelAuthenticationIssueResponse) GetAction() string {
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationIssueResponse) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BackchannelAuthenticationIssueResponse) HasAction() bool {
	if o != nil && !isNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BackchannelAuthenticationIssueResponse) SetAction(v string) {
	o.Action = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *BackchannelAuthenticationIssueResponse) GetResponseContent() string {
	if o == nil || isNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationIssueResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || isNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *BackchannelAuthenticationIssueResponse) HasResponseContent() bool {
	if o != nil && !isNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *BackchannelAuthenticationIssueResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetAuthReqId returns the AuthReqId field value if set, zero value otherwise.
func (o *BackchannelAuthenticationIssueResponse) GetAuthReqId() string {
	if o == nil || isNil(o.AuthReqId) {
		var ret string
		return ret
	}
	return *o.AuthReqId
}

// GetAuthReqIdOk returns a tuple with the AuthReqId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationIssueResponse) GetAuthReqIdOk() (*string, bool) {
	if o == nil || isNil(o.AuthReqId) {
		return nil, false
	}
	return o.AuthReqId, true
}

// HasAuthReqId returns a boolean if a field has been set.
func (o *BackchannelAuthenticationIssueResponse) HasAuthReqId() bool {
	if o != nil && !isNil(o.AuthReqId) {
		return true
	}

	return false
}

// SetAuthReqId gets a reference to the given string and assigns it to the AuthReqId field.
func (o *BackchannelAuthenticationIssueResponse) SetAuthReqId(v string) {
	o.AuthReqId = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *BackchannelAuthenticationIssueResponse) GetExpiresIn() int32 {
	if o == nil || isNil(o.ExpiresIn) {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationIssueResponse) GetExpiresInOk() (*int32, bool) {
	if o == nil || isNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *BackchannelAuthenticationIssueResponse) HasExpiresIn() bool {
	if o != nil && !isNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *BackchannelAuthenticationIssueResponse) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *BackchannelAuthenticationIssueResponse) GetInterval() int32 {
	if o == nil || isNil(o.Interval) {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationIssueResponse) GetIntervalOk() (*int32, bool) {
	if o == nil || isNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *BackchannelAuthenticationIssueResponse) HasInterval() bool {
	if o != nil && !isNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *BackchannelAuthenticationIssueResponse) SetInterval(v int32) {
	o.Interval = &v
}

func (o BackchannelAuthenticationIssueResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackchannelAuthenticationIssueResponse) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.AuthReqId) {
		toSerialize["authReqId"] = o.AuthReqId
	}
	if !isNil(o.ExpiresIn) {
		toSerialize["expiresIn"] = o.ExpiresIn
	}
	if !isNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	return toSerialize, nil
}

type NullableBackchannelAuthenticationIssueResponse struct {
	value *BackchannelAuthenticationIssueResponse
	isSet bool
}

func (v NullableBackchannelAuthenticationIssueResponse) Get() *BackchannelAuthenticationIssueResponse {
	return v.value
}

func (v *NullableBackchannelAuthenticationIssueResponse) Set(val *BackchannelAuthenticationIssueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBackchannelAuthenticationIssueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBackchannelAuthenticationIssueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackchannelAuthenticationIssueResponse(val *BackchannelAuthenticationIssueResponse) *NullableBackchannelAuthenticationIssueResponse {
	return &NullableBackchannelAuthenticationIssueResponse{value: val, isSet: true}
}

func (v NullableBackchannelAuthenticationIssueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackchannelAuthenticationIssueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
