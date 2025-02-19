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

// checks if the UserinfoIssueResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserinfoIssueResponse{}

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
	// The signature header of the response message.
	Signature *string `json:"signature,omitempty"`
	// The signature-input header of the response message
	SignatureInput *string `json:"signatureInput,omitempty"`
	// The content-digest header of the response message
	ContentDigest *string `json:"contentDigest,omitempty"`
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
	if o == nil || isNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *UserinfoIssueResponse) HasResultCode() bool {
	if o != nil && !isNil(o.ResultCode) {
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
	if o == nil || isNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || isNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *UserinfoIssueResponse) HasResultMessage() bool {
	if o != nil && !isNil(o.ResultMessage) {
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
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueResponse) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *UserinfoIssueResponse) HasAction() bool {
	if o != nil && !isNil(o.Action) {
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
	if o == nil || isNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || isNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *UserinfoIssueResponse) HasResponseContent() bool {
	if o != nil && !isNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *UserinfoIssueResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *UserinfoIssueResponse) GetSignature() string {
	if o == nil || isNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueResponse) GetSignatureOk() (*string, bool) {
	if o == nil || isNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *UserinfoIssueResponse) HasSignature() bool {
	if o != nil && !isNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *UserinfoIssueResponse) SetSignature(v string) {
	o.Signature = &v
}

// GetSignatureInput returns the SignatureInput field value if set, zero value otherwise.
func (o *UserinfoIssueResponse) GetSignatureInput() string {
	if o == nil || isNil(o.SignatureInput) {
		var ret string
		return ret
	}
	return *o.SignatureInput
}

// GetSignatureInputOk returns a tuple with the SignatureInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueResponse) GetSignatureInputOk() (*string, bool) {
	if o == nil || isNil(o.SignatureInput) {
		return nil, false
	}
	return o.SignatureInput, true
}

// HasSignatureInput returns a boolean if a field has been set.
func (o *UserinfoIssueResponse) HasSignatureInput() bool {
	if o != nil && !isNil(o.SignatureInput) {
		return true
	}

	return false
}

// SetSignatureInput gets a reference to the given string and assigns it to the SignatureInput field.
func (o *UserinfoIssueResponse) SetSignatureInput(v string) {
	o.SignatureInput = &v
}

// GetContentDigest returns the ContentDigest field value if set, zero value otherwise.
func (o *UserinfoIssueResponse) GetContentDigest() string {
	if o == nil || isNil(o.ContentDigest) {
		var ret string
		return ret
	}
	return *o.ContentDigest
}

// GetContentDigestOk returns a tuple with the ContentDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueResponse) GetContentDigestOk() (*string, bool) {
	if o == nil || isNil(o.ContentDigest) {
		return nil, false
	}
	return o.ContentDigest, true
}

// HasContentDigest returns a boolean if a field has been set.
func (o *UserinfoIssueResponse) HasContentDigest() bool {
	if o != nil && !isNil(o.ContentDigest) {
		return true
	}

	return false
}

// SetContentDigest gets a reference to the given string and assigns it to the ContentDigest field.
func (o *UserinfoIssueResponse) SetContentDigest(v string) {
	o.ContentDigest = &v
}

func (o UserinfoIssueResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserinfoIssueResponse) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if !isNil(o.SignatureInput) {
		toSerialize["signatureInput"] = o.SignatureInput
	}
	if !isNil(o.ContentDigest) {
		toSerialize["contentDigest"] = o.ContentDigest
	}
	return toSerialize, nil
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
