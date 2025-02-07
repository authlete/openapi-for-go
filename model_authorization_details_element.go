/*
Authlete API Explorer

<div class=\"min-h-screen bg-gray-100 dark:bg-gray-900 text-gray-900 dark:text-gray-100 p-6\">   <div class=\"flex justify-end mb-4\">     <label for=\"theme-toggle\" class=\"flex items-center cursor-pointer\">       <div class=\"relative\">Dark mode:         <input type=\"checkbox\" id=\"theme-toggle\" class=\"sr-only\" onchange=\"toggleTheme()\">         <div class=\"block bg-gray-600 w-14 h-8 rounded-full\"></div>         <div class=\"dot absolute left-1 top-1 bg-white w-6 h-6 rounded-full transition\"></div>       </div>     </label>   </div>   <header class=\"bg-green-500 dark:bg-green-700 p-4 rounded-lg text-white text-center\">     <p>       Welcome to the <strong>Authlete API documentation</strong>. Authlete is an <strong>API-first service</strong>       where every aspect of the platform is configurable via API. This explorer provides a convenient way to       authenticate and interact with the API, allowing you to see Authlete in action quickly. 🚀     </p>     <p>       At a high level, the Authlete API is grouped into two categories:     </p>     <ul class=\"list-disc list-inside\">       <li><strong>Management APIs</strong>: Enable you to manage services and clients. 🔧</li>       <li><strong>Runtime APIs</strong>: Allow you to build your own Authorization Servers or Verifiable Credential (VC)         issuers. 🔐</li>     </ul>     <p>All API endpoints are secured using access tokens issued by Authlete's Identity Provider (IdP). If you already       have an Authlete account, simply use the <em>Get Token</em> option on the Authentication page to log in and obtain       an access token for API usage. If you don't have an account yet, <a href=\"https://console.authlete.com/register\">sign up         here</a> to get started.</p>   </header>   <main>     <section id=\"api-servers\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🌐 API Servers</h2>       <p>Authlete is a global service with clusters available in multiple regions across the world.</p>       <p>Currently, our service is available in the following regions:</p>       <div class=\"grid grid-cols-2 gap-4\">         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇺🇸 US</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇯🇵 JP</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇪🇺 EU</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇧🇷 Brazil</p>         </div>       </div>       <p>Our customers can host their data in the region that best meets their requirements.</p>       <a href=\"#servers\" class=\"block mt-4 text-green-500 dark:text-green-300 hover:underline text-center\">Select your         preferred server</a>     </section>     <section id=\"authentication\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🔑 Authentication</h2>       <p>The API Explorer requires an access token to call the API.</p>       <p>You can create the access token from the <a href=\"https://console.authlete.com\">Authlete Management Console</a> and set it in the HTTP Bearer section of Authentication page.</p>       <p>Alternatively, if you have an Authlete account, the API Explorer can log you in with your Authlete account and         automatically acquire the required access token.</p>       <div class=\"theme-admonition theme-admonition-warning admonition_o5H7 alert alert--warning\">         <div class=\"admonitionContent_Knsx\">           <p>⚠️ <strong>Important Note:</strong> When the API Explorer acquires the token after login, the access tokens             will have the same permissions as the user who logs in as part of this flow.</p>         </div>       </div>       <a href=\"#auth\" class=\"block mt-4 text-green-500 dark:text-green-300 hover:underline text-center\">Setup your         access token</a>     </section>     <section id=\"tutorials\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🎓 Tutorials</h2>       <p>If you have successfully tested the API from the API Console and want to take the next step of integrating the         API into your application, or if you want to see a sample using Authlete APIs, follow the links below. These         resources will help you understand key concepts and how to integrate Authlete API into your applications.</p>       <div class=\"mt-4\">         <a href=\"https://www.authlete.com/developers/getting_started/\"           class=\"block text-green-500 dark:text-green-300 font-bold hover:underline mb-2\">🚀 Getting Started with           Authlete</a>           </br>         <a href=\"https://www.authlete.com/developers/tutorial/signup/\"           class=\"block text-green-500 dark:text-green-300 font-bold hover:underline\">🔑 From Sign-Up to the First API           Request</a>       </div>     </section>     <section id=\"support\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🛠 Contact Us</h2>       <p>If you have any questions or need assistance, our team is here to help.</p>       <a href=\"https://www.authlete.com/contact/\"         class=\"block mt-4 text-green-500 dark:text-green-300 font-bold hover:underline\">Contact Page</a>     </section>   </main> </div>

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AuthorizationDetailsElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationDetailsElement{}

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
	// The RAR request in the JSON format excluding the pre-defined attributes such as `type` and `locations`. The content and semantics are specific to the deployment and the use case implemented.
	OtherFields *string `json:"otherFields,omitempty"`
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
	if o == nil || IsNil(o.Locations) {
		var ret []string
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailsElement) GetLocationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Locations) {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *AuthorizationDetailsElement) HasLocations() bool {
	if o != nil && !IsNil(o.Locations) {
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
	if o == nil || IsNil(o.Actions) {
		var ret []string
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailsElement) GetActionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *AuthorizationDetailsElement) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
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
	if o == nil || IsNil(o.DataTypes) {
		var ret []string
		return ret
	}
	return o.DataTypes
}

// GetDataTypesOk returns a tuple with the DataTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailsElement) GetDataTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.DataTypes) {
		return nil, false
	}
	return o.DataTypes, true
}

// HasDataTypes returns a boolean if a field has been set.
func (o *AuthorizationDetailsElement) HasDataTypes() bool {
	if o != nil && !IsNil(o.DataTypes) {
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
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailsElement) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *AuthorizationDetailsElement) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
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
	if o == nil || IsNil(o.Privileges) {
		var ret []string
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailsElement) GetPrivilegesOk() ([]string, bool) {
	if o == nil || IsNil(o.Privileges) {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *AuthorizationDetailsElement) HasPrivileges() bool {
	if o != nil && !IsNil(o.Privileges) {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []string and assigns it to the Privileges field.
func (o *AuthorizationDetailsElement) SetPrivileges(v []string) {
	o.Privileges = v
}

// GetOtherFields returns the OtherFields field value if set, zero value otherwise.
func (o *AuthorizationDetailsElement) GetOtherFields() string {
	if o == nil || IsNil(o.OtherFields) {
		var ret string
		return ret
	}
	return *o.OtherFields
}

// GetOtherFieldsOk returns a tuple with the OtherFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailsElement) GetOtherFieldsOk() (*string, bool) {
	if o == nil || IsNil(o.OtherFields) {
		return nil, false
	}
	return o.OtherFields, true
}

// HasOtherFields returns a boolean if a field has been set.
func (o *AuthorizationDetailsElement) HasOtherFields() bool {
	if o != nil && !IsNil(o.OtherFields) {
		return true
	}

	return false
}

// SetOtherFields gets a reference to the given string and assigns it to the OtherFields field.
func (o *AuthorizationDetailsElement) SetOtherFields(v string) {
	o.OtherFields = &v
}

func (o AuthorizationDetailsElement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationDetailsElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Locations) {
		toSerialize["locations"] = o.Locations
	}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.DataTypes) {
		toSerialize["dataTypes"] = o.DataTypes
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Privileges) {
		toSerialize["privileges"] = o.Privileges
	}
	if !IsNil(o.OtherFields) {
		toSerialize["otherFields"] = o.OtherFields
	}
	return toSerialize, nil
}

func (o *AuthorizationDetailsElement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAuthorizationDetailsElement := _AuthorizationDetailsElement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthorizationDetailsElement)

	if err != nil {
		return err
	}

	*o = AuthorizationDetailsElement(varAuthorizationDetailsElement)

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
