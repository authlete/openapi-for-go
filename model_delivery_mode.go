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

// DeliveryMode the model 'DeliveryMode'
type DeliveryMode string

// List of delivery_mode
const (
	DELIVERYMODE_PING DeliveryMode = "PING"
	DELIVERYMODE_POLL DeliveryMode = "POLL"
	DELIVERYMODE_PUSH DeliveryMode = "PUSH"
)

// All allowed values of DeliveryMode enum
var AllowedDeliveryModeEnumValues = []DeliveryMode{
	"PING",
	"POLL",
	"PUSH",
}

func (v *DeliveryMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeliveryMode(value)
	for _, existing := range AllowedDeliveryModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeliveryMode", value)
}

// NewDeliveryModeFromValue returns a pointer to a valid DeliveryMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeliveryModeFromValue(v string) (*DeliveryMode, error) {
	ev := DeliveryMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeliveryMode: valid values are %v", v, AllowedDeliveryModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeliveryMode) IsValid() bool {
	for _, existing := range AllowedDeliveryModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to delivery_mode value
func (v DeliveryMode) Ptr() *DeliveryMode {
	return &v
}

type NullableDeliveryMode struct {
	value *DeliveryMode
	isSet bool
}

func (v NullableDeliveryMode) Get() *DeliveryMode {
	return v.value
}

func (v *NullableDeliveryMode) Set(val *DeliveryMode) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryMode) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryMode(val *DeliveryMode) *NullableDeliveryMode {
	return &NullableDeliveryMode{value: val, isSet: true}
}

func (v NullableDeliveryMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
