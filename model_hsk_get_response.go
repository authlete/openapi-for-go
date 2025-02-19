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

// checks if the HskGetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HskGetResponse{}

// HskGetResponse struct for HskGetResponse
type HskGetResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// Result of the API call
	Action *string `json:"action,omitempty"`
	Hsk    *Hsk    `json:"hsk,omitempty"`
}

// NewHskGetResponse instantiates a new HskGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHskGetResponse() *HskGetResponse {
	this := HskGetResponse{}
	return &this
}

// NewHskGetResponseWithDefaults instantiates a new HskGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHskGetResponseWithDefaults() *HskGetResponse {
	this := HskGetResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *HskGetResponse) GetResultCode() string {
	if o == nil || isNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HskGetResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *HskGetResponse) HasResultCode() bool {
	if o != nil && !isNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *HskGetResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *HskGetResponse) GetResultMessage() string {
	if o == nil || isNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HskGetResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || isNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *HskGetResponse) HasResultMessage() bool {
	if o != nil && !isNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *HskGetResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *HskGetResponse) GetAction() string {
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HskGetResponse) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *HskGetResponse) HasAction() bool {
	if o != nil && !isNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *HskGetResponse) SetAction(v string) {
	o.Action = &v
}

// GetHsk returns the Hsk field value if set, zero value otherwise.
func (o *HskGetResponse) GetHsk() Hsk {
	if o == nil || isNil(o.Hsk) {
		var ret Hsk
		return ret
	}
	return *o.Hsk
}

// GetHskOk returns a tuple with the Hsk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HskGetResponse) GetHskOk() (*Hsk, bool) {
	if o == nil || isNil(o.Hsk) {
		return nil, false
	}
	return o.Hsk, true
}

// HasHsk returns a boolean if a field has been set.
func (o *HskGetResponse) HasHsk() bool {
	if o != nil && !isNil(o.Hsk) {
		return true
	}

	return false
}

// SetHsk gets a reference to the given Hsk and assigns it to the Hsk field.
func (o *HskGetResponse) SetHsk(v Hsk) {
	o.Hsk = &v
}

func (o HskGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HskGetResponse) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.Hsk) {
		toSerialize["hsk"] = o.Hsk
	}
	return toSerialize, nil
}

type NullableHskGetResponse struct {
	value *HskGetResponse
	isSet bool
}

func (v NullableHskGetResponse) Get() *HskGetResponse {
	return v.value
}

func (v *NullableHskGetResponse) Set(val *HskGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHskGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHskGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHskGetResponse(val *HskGetResponse) *NullableHskGetResponse {
	return &NullableHskGetResponse{value: val, isSet: true}
}

func (v NullableHskGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHskGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
