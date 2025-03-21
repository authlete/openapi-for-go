/*
Authlete API

Authlete API Document.

API version: 2.3.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
	"fmt"
)

// Display The display mode which the client application requests by `display` request parameter. When the authorization request does not have `display` request parameter, `PAGE` is set as the default value.  It is ensured that the value of `display` is one of the supported display modes which are specified by `supportedDisplays` configuration parameter of the service. If the display mode specified by the authorization request is not supported, an error is raised.  Values for this property correspond to the values listed in \"[OpenID Connect Core 1.0, 3.1.2.1. Authentication Request](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest), display\".
type Display string

// List of Display
const (
	DISPLAY_PAGE  Display = "PAGE"
	DISPLAY_POPUP Display = "POPUP"
	DISPLAY_TOUCH Display = "TOUCH"
	DISPLAY_WAP   Display = "WAP"
)

// All allowed values of Display enum
var AllowedDisplayEnumValues = []Display{
	"PAGE",
	"POPUP",
	"TOUCH",
	"WAP",
}

func (v *Display) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Display(value)
	for _, existing := range AllowedDisplayEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Display", value)
}

// NewDisplayFromValue returns a pointer to a valid Display
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDisplayFromValue(v string) (*Display, error) {
	ev := Display(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Display: valid values are %v", v, AllowedDisplayEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Display) IsValid() bool {
	for _, existing := range AllowedDisplayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Display value
func (v Display) Ptr() *Display {
	return &v
}

type NullableDisplay struct {
	value *Display
	isSet bool
}

func (v NullableDisplay) Get() *Display {
	return v.value
}

func (v *NullableDisplay) Set(val *Display) {
	v.value = val
	v.isSet = true
}

func (v NullableDisplay) IsSet() bool {
	return v.isSet
}

func (v *NullableDisplay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisplay(val *Display) *NullableDisplay {
	return &NullableDisplay{value: val, isSet: true}
}

func (v NullableDisplay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisplay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
