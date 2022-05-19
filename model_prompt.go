/*
Authlete API

Authlete API Document. 

API version: 2.2.25
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
	"fmt"
)

// Prompt The prompt that the UI displayed to the end-user must satisfy as the minimum level. This value comes from `prompt` request parameter.  When the authorization request does not contain `prompt` request parameter, `CONSENT` is used as the default value.  See \"[OpenID Connect Core 1.0, 3.1.2.1. Authentication Request](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest), prompt\" for `prompt` request parameter. 
type Prompt string

// List of Prompt
const (
	PROMPT_NONE Prompt = "NONE"
	PROMPT_LOGIN Prompt = "LOGIN"
	PROMPT_CONSENT Prompt = "CONSENT"
	PROMPT_SELECT_ACCOUNT Prompt = "SELECT_ACCOUNT"
)

// All allowed values of Prompt enum
var AllowedPromptEnumValues = []Prompt{
	"NONE",
	"LOGIN",
	"CONSENT",
	"SELECT_ACCOUNT",
}

func (v *Prompt) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Prompt(value)
	for _, existing := range AllowedPromptEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Prompt", value)
}

// NewPromptFromValue returns a pointer to a valid Prompt
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromptFromValue(v string) (*Prompt, error) {
	ev := Prompt(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Prompt: valid values are %v", v, AllowedPromptEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Prompt) IsValid() bool {
	for _, existing := range AllowedPromptEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Prompt value
func (v Prompt) Ptr() *Prompt {
	return &v
}

type NullablePrompt struct {
	value *Prompt
	isSet bool
}

func (v NullablePrompt) Get() *Prompt {
	return v.value
}

func (v *NullablePrompt) Set(val *Prompt) {
	v.value = val
	v.isSet = true
}

func (v NullablePrompt) IsSet() bool {
	return v.isSet
}

func (v *NullablePrompt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrompt(val *Prompt) *NullablePrompt {
	return &NullablePrompt{value: val, isSet: true}
}

func (v NullablePrompt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrompt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

