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

// AuthorizationDetailsElement struct for AuthorizationDetailsElement
type AuthorizationDetailsElement struct {
	// The type of this element.  From _\"OAuth 2.0 Rich Authorization Requests\"_: _\"The type of authorization data as a string. This field MAY define which other elements are allowed in the request. This element is REQUIRED.\"_  This property is always NOT `null`. 
	Type string `json:"type"`
	// The resources and/or resource servers. This property may be `null`.  From _\"OAuth 2.0 Rich Authorization Requests\"_: _\"An array of strings representing the location of the resource or resource server. This is typically composed of URIs.\"_  This property may be `null`. 
	Locations []string `json:"locations,omitempty"`
	// The actions.  From _\"OAuth 2.0 Rich Authorization Requests\"_: _\"An array of strings representing the kinds of actions to be taken at the resource. The values of the strings are determined by the API being protected.\"_  This property may be `null`. 
	Actions []string `json:"actions,omitempty"`
	// From _\"OAuth 2.0 Rich Authorization Requests\"_: _\"An array of strings representing the kinds of data being requested from the resource.\"_  This property may be `null`. 
	DataTypes []string `json:"dataTypes,omitempty"`
	// The identifier of a specific resource. From _\"OAuth 2.0 Rich Authorization Requests\"_: _\"A string identifier indicating a specific resource available at the API.\"_  This property may be `null`. 
	Identifier *string `json:"identifier,omitempty"`
	// The types or levels of privilege. From \"OAuth 2.0 Rich Authorization Requests\": _\"An array of strings representing the types or levels of privilege being requested at the resource.\"_  This property may be `null`. 
	Privileges []string `json:"privileges,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationDetailsElement AuthorizationDetailsElement

// NewAuthorizationDetailsElement instantiates a new AuthorizationDetailsElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationDetailsElement(type_ string) *AuthorizationDetailsElement {
	this := AuthorizationDetailsElement{}
	this.Type = type_
	return &this
}

// NewAuthorizationDetailsElementWithDefaults instantiates a new AuthorizationDetailsElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationDetailsElementWithDefaults() *AuthorizationDetailsElement {
	this := AuthorizationDetailsElement{}
	return &this
}

// GetType returns the Type field value
func (o *AuthorizationDetailsElement) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailsElement) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthorizationDetailsElement) SetType(v string) {
	o.Type = v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *AuthorizationDetailsElement) GetLocations() []string {
	if o == nil || o.Locations == nil {
		var ret []string
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailsElement) GetLocationsOk() ([]string, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *AuthorizationDetailsElement) HasLocations() bool {
	if o != nil && o.Locations != nil {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []string and assigns it to the Locations field.
func (o *AuthorizationDetailsElement) SetLocations(v []string) {
	o.Locations = v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *AuthorizationDetailsElement) GetActions() []string {
	if o == nil || o.Actions == nil {
		var ret []string
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailsElement) GetActionsOk() ([]string, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *AuthorizationDetailsElement) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []string and assigns it to the Actions field.
func (o *AuthorizationDetailsElement) SetActions(v []string) {
	o.Actions = v
}

// GetDataTypes returns the DataTypes field value if set, zero value otherwise.
func (o *AuthorizationDetailsElement) GetDataTypes() []string {
	if o == nil || o.DataTypes == nil {
		var ret []string
		return ret
	}
	return o.DataTypes
}

// GetDataTypesOk returns a tuple with the DataTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailsElement) GetDataTypesOk() ([]string, bool) {
	if o == nil || o.DataTypes == nil {
		return nil, false
	}
	return o.DataTypes, true
}

// HasDataTypes returns a boolean if a field has been set.
func (o *AuthorizationDetailsElement) HasDataTypes() bool {
	if o != nil && o.DataTypes != nil {
		return true
	}

	return false
}

// SetDataTypes gets a reference to the given []string and assigns it to the DataTypes field.
func (o *AuthorizationDetailsElement) SetDataTypes(v []string) {
	o.DataTypes = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *AuthorizationDetailsElement) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailsElement) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *AuthorizationDetailsElement) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *AuthorizationDetailsElement) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *AuthorizationDetailsElement) GetPrivileges() []string {
	if o == nil || o.Privileges == nil {
		var ret []string
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailsElement) GetPrivilegesOk() ([]string, bool) {
	if o == nil || o.Privileges == nil {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *AuthorizationDetailsElement) HasPrivileges() bool {
	if o != nil && o.Privileges != nil {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []string and assigns it to the Privileges field.
func (o *AuthorizationDetailsElement) SetPrivileges(v []string) {
	o.Privileges = v
}

func (o AuthorizationDetailsElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.DataTypes != nil {
		toSerialize["dataTypes"] = o.DataTypes
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.Privileges != nil {
		toSerialize["privileges"] = o.Privileges
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthorizationDetailsElement) UnmarshalJSON(bytes []byte) (err error) {
	varAuthorizationDetailsElement := _AuthorizationDetailsElement{}

	if err = json.Unmarshal(bytes, &varAuthorizationDetailsElement); err == nil {
		*o = AuthorizationDetailsElement(varAuthorizationDetailsElement)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "locations")
		delete(additionalProperties, "actions")
		delete(additionalProperties, "dataTypes")
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "privileges")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthorizationDetailsElement struct {
	value *AuthorizationDetailsElement
	isSet bool
}

func (v NullableAuthorizationDetailsElement) Get() *AuthorizationDetailsElement {
	return v.value
}

func (v *NullableAuthorizationDetailsElement) Set(val *AuthorizationDetailsElement) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationDetailsElement) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationDetailsElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationDetailsElement(val *AuthorizationDetailsElement) *NullableAuthorizationDetailsElement {
	return &NullableAuthorizationDetailsElement{value: val, isSet: true}
}

func (v NullableAuthorizationDetailsElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationDetailsElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


