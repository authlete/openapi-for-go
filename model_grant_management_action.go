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

// GrantManagementAction The grant management action of the device authorization request.  The `grant_management_action` request parameter is defined in [Grant Management for OAuth 2.0](https://openid.net/specs/fapi-grant-management.html).
type GrantManagementAction string

// List of grant_management_action
const (
	GRANTMANAGEMENTACTION_CREATE  GrantManagementAction = "CREATE"
	GRANTMANAGEMENTACTION_QUERY   GrantManagementAction = "QUERY"
	GRANTMANAGEMENTACTION_REPLACE GrantManagementAction = "REPLACE"
	GRANTMANAGEMENTACTION_REVOKE  GrantManagementAction = "REVOKE"
	GRANTMANAGEMENTACTION_MERGE   GrantManagementAction = "MERGE"
)

// All allowed values of GrantManagementAction enum
var AllowedGrantManagementActionEnumValues = []GrantManagementAction{
	"CREATE",
	"QUERY",
	"REPLACE",
	"REVOKE",
	"MERGE",
}

func (v *GrantManagementAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GrantManagementAction(value)
	for _, existing := range AllowedGrantManagementActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GrantManagementAction", value)
}

// NewGrantManagementActionFromValue returns a pointer to a valid GrantManagementAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGrantManagementActionFromValue(v string) (*GrantManagementAction, error) {
	ev := GrantManagementAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GrantManagementAction: valid values are %v", v, AllowedGrantManagementActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GrantManagementAction) IsValid() bool {
	for _, existing := range AllowedGrantManagementActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to grant_management_action value
func (v GrantManagementAction) Ptr() *GrantManagementAction {
	return &v
}

type NullableGrantManagementAction struct {
	value *GrantManagementAction
	isSet bool
}

func (v NullableGrantManagementAction) Get() *GrantManagementAction {
	return v.value
}

func (v *NullableGrantManagementAction) Set(val *GrantManagementAction) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantManagementAction) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantManagementAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantManagementAction(val *GrantManagementAction) *NullableGrantManagementAction {
	return &NullableGrantManagementAction{value: val, isSet: true}
}

func (v NullableGrantManagementAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantManagementAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
