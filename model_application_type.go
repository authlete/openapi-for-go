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

// ApplicationType The application type. The value of this property affects the validation steps for a redirect URI. See the description about `redirectUris` property for more details.
type ApplicationType string

// List of application_type
const (
	APPLICATIONTYPE_WEB    ApplicationType = "WEB"
	APPLICATIONTYPE_NATIVE ApplicationType = "NATIVE"
	APPLICATIONTYPE_NULL   ApplicationType = "null"
)

// All allowed values of ApplicationType enum
var AllowedApplicationTypeEnumValues = []ApplicationType{
	"WEB",
	"NATIVE",
	"null",
}

func (v *ApplicationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApplicationType(value)
	for _, existing := range AllowedApplicationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApplicationType", value)
}

// NewApplicationTypeFromValue returns a pointer to a valid ApplicationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApplicationTypeFromValue(v string) (*ApplicationType, error) {
	ev := ApplicationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApplicationType: valid values are %v", v, AllowedApplicationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApplicationType) IsValid() bool {
	for _, existing := range AllowedApplicationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to application_type value
func (v ApplicationType) Ptr() *ApplicationType {
	return &v
}

type NullableApplicationType struct {
	value *ApplicationType
	isSet bool
}

func (v NullableApplicationType) Get() *ApplicationType {
	return v.value
}

func (v *NullableApplicationType) Set(val *ApplicationType) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationType) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationType(val *ApplicationType) *NullableApplicationType {
	return &NullableApplicationType{value: val, isSet: true}
}

func (v NullableApplicationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
