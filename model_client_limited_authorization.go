/*
Authlete API

Authlete API Document. 

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

// checks if the ClientLimitedAuthorization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientLimitedAuthorization{}

// ClientLimitedAuthorization struct for ClientLimitedAuthorization
type ClientLimitedAuthorization struct {
	// The sequential number of the client. The value of this property is assigned by Authlete. 
	Number *int32 `json:"number,omitempty"`
	// The name of the client application. This property corresponds to `client_name` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata). 
	ClientName *string `json:"clientName,omitempty"`
	// Client names with language tags. If the client application has different names for different languages, this property can be used to register the names. 
	ClientNames []TaggedValue `json:"clientNames,omitempty"`
	// The description about the client application.
	Description *string `json:"description,omitempty"`
	// Descriptions about the client application with language tags. If the client application has different descriptions for different languages, this property can be used to register the descriptions. 
	Descriptions []TaggedValue `json:"descriptions,omitempty"`
	// The client identifier used in Authlete API calls. The value of this property is assigned by Authlete.
	ClientId *int64 `json:"clientId,omitempty"`
	// The value of the client's `client_id` property used in OAuth and OpenID Connect calls. By default, this is a string version of the `clientId` property. 
	ClientIdAlias *string `json:"clientIdAlias,omitempty"`
	// Deprecated. Always set to `true`.
	ClientIdAliasEnabled *bool `json:"clientIdAliasEnabled,omitempty"`
	ClientType *ClientType `json:"clientType,omitempty"`
	// The URL pointing to the logo image of the client application.  This property corresponds to `logo_uri` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata). 
	LogoUri *string `json:"logoUri,omitempty"`
	// Logo image URLs with language tags. If the client application has different logo images for different languages, this property can be used to register URLs of the images. 
	LogoUris []TaggedValue `json:"logoUris,omitempty"`
	// The URL pointing to the \"Terms Of Service\" page.  This property corresponds to `tos_uri` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata). 
	TosUri *string `json:"tosUri,omitempty"`
	// URLs of \"Terms Of Service\" pages with language tags.  If the client application has different \"Terms Of Service\" pages for different languages, this property can be used to register the URLs. 
	TosUris []TaggedValue `json:"tosUris,omitempty"`
	// The URL pointing to the page which describes the policy as to how end-user's profile data is used.  This property corresponds to `policy_uri` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata). 
	PolicyUri *string `json:"policyUri,omitempty"`
	// URLs of policy pages with language tags. If the client application has different policy pages for different languages, this property can be used to register the URLs. 
	PolicyUris []TaggedValue `json:"policyUris,omitempty"`
}

// NewClientLimitedAuthorization instantiates a new ClientLimitedAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientLimitedAuthorization() *ClientLimitedAuthorization {
	this := ClientLimitedAuthorization{}
	return &this
}

// NewClientLimitedAuthorizationWithDefaults instantiates a new ClientLimitedAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientLimitedAuthorizationWithDefaults() *ClientLimitedAuthorization {
	this := ClientLimitedAuthorization{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *ClientLimitedAuthorization) GetNumber() int32 {
	if o == nil || isNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimitedAuthorization) GetNumberOk() (*int32, bool) {
	if o == nil || isNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *ClientLimitedAuthorization) HasNumber() bool {
	if o != nil && !isNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *ClientLimitedAuthorization) SetNumber(v int32) {
	o.Number = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *ClientLimitedAuthorization) GetClientName() string {
	if o == nil || isNil(o.ClientName) {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimitedAuthorization) GetClientNameOk() (*string, bool) {
	if o == nil || isNil(o.ClientName) {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *ClientLimitedAuthorization) HasClientName() bool {
	if o != nil && !isNil(o.ClientName) {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *ClientLimitedAuthorization) SetClientName(v string) {
	o.ClientName = &v
}

// GetClientNames returns the ClientNames field value if set, zero value otherwise.
func (o *ClientLimitedAuthorization) GetClientNames() []TaggedValue {
	if o == nil || isNil(o.ClientNames) {
		var ret []TaggedValue
		return ret
	}
	return o.ClientNames
}

// GetClientNamesOk returns a tuple with the ClientNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimitedAuthorization) GetClientNamesOk() ([]TaggedValue, bool) {
	if o == nil || isNil(o.ClientNames) {
		return nil, false
	}
	return o.ClientNames, true
}

// HasClientNames returns a boolean if a field has been set.
func (o *ClientLimitedAuthorization) HasClientNames() bool {
	if o != nil && !isNil(o.ClientNames) {
		return true
	}

	return false
}

// SetClientNames gets a reference to the given []TaggedValue and assigns it to the ClientNames field.
func (o *ClientLimitedAuthorization) SetClientNames(v []TaggedValue) {
	o.ClientNames = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClientLimitedAuthorization) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimitedAuthorization) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClientLimitedAuthorization) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClientLimitedAuthorization) SetDescription(v string) {
	o.Description = &v
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *ClientLimitedAuthorization) GetDescriptions() []TaggedValue {
	if o == nil || isNil(o.Descriptions) {
		var ret []TaggedValue
		return ret
	}
	return o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimitedAuthorization) GetDescriptionsOk() ([]TaggedValue, bool) {
	if o == nil || isNil(o.Descriptions) {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *ClientLimitedAuthorization) HasDescriptions() bool {
	if o != nil && !isNil(o.Descriptions) {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []TaggedValue and assigns it to the Descriptions field.
func (o *ClientLimitedAuthorization) SetDescriptions(v []TaggedValue) {
	o.Descriptions = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ClientLimitedAuthorization) GetClientId() int64 {
	if o == nil || isNil(o.ClientId) {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimitedAuthorization) GetClientIdOk() (*int64, bool) {
	if o == nil || isNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ClientLimitedAuthorization) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *ClientLimitedAuthorization) SetClientId(v int64) {
	o.ClientId = &v
}

// GetClientIdAlias returns the ClientIdAlias field value if set, zero value otherwise.
func (o *ClientLimitedAuthorization) GetClientIdAlias() string {
	if o == nil || isNil(o.ClientIdAlias) {
		var ret string
		return ret
	}
	return *o.ClientIdAlias
}

// GetClientIdAliasOk returns a tuple with the ClientIdAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimitedAuthorization) GetClientIdAliasOk() (*string, bool) {
	if o == nil || isNil(o.ClientIdAlias) {
		return nil, false
	}
	return o.ClientIdAlias, true
}

// HasClientIdAlias returns a boolean if a field has been set.
func (o *ClientLimitedAuthorization) HasClientIdAlias() bool {
	if o != nil && !isNil(o.ClientIdAlias) {
		return true
	}

	return false
}

// SetClientIdAlias gets a reference to the given string and assigns it to the ClientIdAlias field.
func (o *ClientLimitedAuthorization) SetClientIdAlias(v string) {
	o.ClientIdAlias = &v
}

// GetClientIdAliasEnabled returns the ClientIdAliasEnabled field value if set, zero value otherwise.
func (o *ClientLimitedAuthorization) GetClientIdAliasEnabled() bool {
	if o == nil || isNil(o.ClientIdAliasEnabled) {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasEnabled
}

// GetClientIdAliasEnabledOk returns a tuple with the ClientIdAliasEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimitedAuthorization) GetClientIdAliasEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.ClientIdAliasEnabled) {
		return nil, false
	}
	return o.ClientIdAliasEnabled, true
}

// HasClientIdAliasEnabled returns a boolean if a field has been set.
func (o *ClientLimitedAuthorization) HasClientIdAliasEnabled() bool {
	if o != nil && !isNil(o.ClientIdAliasEnabled) {
		return true
	}

	return false
}

// SetClientIdAliasEnabled gets a reference to the given bool and assigns it to the ClientIdAliasEnabled field.
func (o *ClientLimitedAuthorization) SetClientIdAliasEnabled(v bool) {
	o.ClientIdAliasEnabled = &v
}

// GetClientType returns the ClientType field value if set, zero value otherwise.
func (o *ClientLimitedAuthorization) GetClientType() ClientType {
	if o == nil || isNil(o.ClientType) {
		var ret ClientType
		return ret
	}
	return *o.ClientType
}

// GetClientTypeOk returns a tuple with the ClientType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimitedAuthorization) GetClientTypeOk() (*ClientType, bool) {
	if o == nil || isNil(o.ClientType) {
		return nil, false
	}
	return o.ClientType, true
}

// HasClientType returns a boolean if a field has been set.
func (o *ClientLimitedAuthorization) HasClientType() bool {
	if o != nil && !isNil(o.ClientType) {
		return true
	}

	return false
}

// SetClientType gets a reference to the given ClientType and assigns it to the ClientType field.
func (o *ClientLimitedAuthorization) SetClientType(v ClientType) {
	o.ClientType = &v
}

// GetLogoUri returns the LogoUri field value if set, zero value otherwise.
func (o *ClientLimitedAuthorization) GetLogoUri() string {
	if o == nil || isNil(o.LogoUri) {
		var ret string
		return ret
	}
	return *o.LogoUri
}

// GetLogoUriOk returns a tuple with the LogoUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimitedAuthorization) GetLogoUriOk() (*string, bool) {
	if o == nil || isNil(o.LogoUri) {
		return nil, false
	}
	return o.LogoUri, true
}

// HasLogoUri returns a boolean if a field has been set.
func (o *ClientLimitedAuthorization) HasLogoUri() bool {
	if o != nil && !isNil(o.LogoUri) {
		return true
	}

	return false
}

// SetLogoUri gets a reference to the given string and assigns it to the LogoUri field.
func (o *ClientLimitedAuthorization) SetLogoUri(v string) {
	o.LogoUri = &v
}

// GetLogoUris returns the LogoUris field value if set, zero value otherwise.
func (o *ClientLimitedAuthorization) GetLogoUris() []TaggedValue {
	if o == nil || isNil(o.LogoUris) {
		var ret []TaggedValue
		return ret
	}
	return o.LogoUris
}

// GetLogoUrisOk returns a tuple with the LogoUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimitedAuthorization) GetLogoUrisOk() ([]TaggedValue, bool) {
	if o == nil || isNil(o.LogoUris) {
		return nil, false
	}
	return o.LogoUris, true
}

// HasLogoUris returns a boolean if a field has been set.
func (o *ClientLimitedAuthorization) HasLogoUris() bool {
	if o != nil && !isNil(o.LogoUris) {
		return true
	}

	return false
}

// SetLogoUris gets a reference to the given []TaggedValue and assigns it to the LogoUris field.
func (o *ClientLimitedAuthorization) SetLogoUris(v []TaggedValue) {
	o.LogoUris = v
}

// GetTosUri returns the TosUri field value if set, zero value otherwise.
func (o *ClientLimitedAuthorization) GetTosUri() string {
	if o == nil || isNil(o.TosUri) {
		var ret string
		return ret
	}
	return *o.TosUri
}

// GetTosUriOk returns a tuple with the TosUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimitedAuthorization) GetTosUriOk() (*string, bool) {
	if o == nil || isNil(o.TosUri) {
		return nil, false
	}
	return o.TosUri, true
}

// HasTosUri returns a boolean if a field has been set.
func (o *ClientLimitedAuthorization) HasTosUri() bool {
	if o != nil && !isNil(o.TosUri) {
		return true
	}

	return false
}

// SetTosUri gets a reference to the given string and assigns it to the TosUri field.
func (o *ClientLimitedAuthorization) SetTosUri(v string) {
	o.TosUri = &v
}

// GetTosUris returns the TosUris field value if set, zero value otherwise.
func (o *ClientLimitedAuthorization) GetTosUris() []TaggedValue {
	if o == nil || isNil(o.TosUris) {
		var ret []TaggedValue
		return ret
	}
	return o.TosUris
}

// GetTosUrisOk returns a tuple with the TosUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimitedAuthorization) GetTosUrisOk() ([]TaggedValue, bool) {
	if o == nil || isNil(o.TosUris) {
		return nil, false
	}
	return o.TosUris, true
}

// HasTosUris returns a boolean if a field has been set.
func (o *ClientLimitedAuthorization) HasTosUris() bool {
	if o != nil && !isNil(o.TosUris) {
		return true
	}

	return false
}

// SetTosUris gets a reference to the given []TaggedValue and assigns it to the TosUris field.
func (o *ClientLimitedAuthorization) SetTosUris(v []TaggedValue) {
	o.TosUris = v
}

// GetPolicyUri returns the PolicyUri field value if set, zero value otherwise.
func (o *ClientLimitedAuthorization) GetPolicyUri() string {
	if o == nil || isNil(o.PolicyUri) {
		var ret string
		return ret
	}
	return *o.PolicyUri
}

// GetPolicyUriOk returns a tuple with the PolicyUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimitedAuthorization) GetPolicyUriOk() (*string, bool) {
	if o == nil || isNil(o.PolicyUri) {
		return nil, false
	}
	return o.PolicyUri, true
}

// HasPolicyUri returns a boolean if a field has been set.
func (o *ClientLimitedAuthorization) HasPolicyUri() bool {
	if o != nil && !isNil(o.PolicyUri) {
		return true
	}

	return false
}

// SetPolicyUri gets a reference to the given string and assigns it to the PolicyUri field.
func (o *ClientLimitedAuthorization) SetPolicyUri(v string) {
	o.PolicyUri = &v
}

// GetPolicyUris returns the PolicyUris field value if set, zero value otherwise.
func (o *ClientLimitedAuthorization) GetPolicyUris() []TaggedValue {
	if o == nil || isNil(o.PolicyUris) {
		var ret []TaggedValue
		return ret
	}
	return o.PolicyUris
}

// GetPolicyUrisOk returns a tuple with the PolicyUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimitedAuthorization) GetPolicyUrisOk() ([]TaggedValue, bool) {
	if o == nil || isNil(o.PolicyUris) {
		return nil, false
	}
	return o.PolicyUris, true
}

// HasPolicyUris returns a boolean if a field has been set.
func (o *ClientLimitedAuthorization) HasPolicyUris() bool {
	if o != nil && !isNil(o.PolicyUris) {
		return true
	}

	return false
}

// SetPolicyUris gets a reference to the given []TaggedValue and assigns it to the PolicyUris field.
func (o *ClientLimitedAuthorization) SetPolicyUris(v []TaggedValue) {
	o.PolicyUris = v
}

func (o ClientLimitedAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientLimitedAuthorization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !isNil(o.ClientName) {
		toSerialize["clientName"] = o.ClientName
	}
	if !isNil(o.ClientNames) {
		toSerialize["clientNames"] = o.ClientNames
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Descriptions) {
		toSerialize["descriptions"] = o.Descriptions
	}
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !isNil(o.ClientIdAlias) {
		toSerialize["clientIdAlias"] = o.ClientIdAlias
	}
	if !isNil(o.ClientIdAliasEnabled) {
		toSerialize["clientIdAliasEnabled"] = o.ClientIdAliasEnabled
	}
	if !isNil(o.ClientType) {
		toSerialize["clientType"] = o.ClientType
	}
	if !isNil(o.LogoUri) {
		toSerialize["logoUri"] = o.LogoUri
	}
	if !isNil(o.LogoUris) {
		toSerialize["logoUris"] = o.LogoUris
	}
	if !isNil(o.TosUri) {
		toSerialize["tosUri"] = o.TosUri
	}
	if !isNil(o.TosUris) {
		toSerialize["tosUris"] = o.TosUris
	}
	if !isNil(o.PolicyUri) {
		toSerialize["policyUri"] = o.PolicyUri
	}
	if !isNil(o.PolicyUris) {
		toSerialize["policyUris"] = o.PolicyUris
	}
	return toSerialize, nil
}

type NullableClientLimitedAuthorization struct {
	value *ClientLimitedAuthorization
	isSet bool
}

func (v NullableClientLimitedAuthorization) Get() *ClientLimitedAuthorization {
	return v.value
}

func (v *NullableClientLimitedAuthorization) Set(val *ClientLimitedAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableClientLimitedAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableClientLimitedAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientLimitedAuthorization(val *ClientLimitedAuthorization) *NullableClientLimitedAuthorization {
	return &NullableClientLimitedAuthorization{value: val, isSet: true}
}

func (v NullableClientLimitedAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientLimitedAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


