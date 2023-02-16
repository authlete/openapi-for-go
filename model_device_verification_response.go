/*
Authlete API

Authlete API Document. 

API version: 2.2.30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

// checks if the DeviceVerificationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceVerificationResponse{}

// DeviceVerificationResponse struct for DeviceVerificationResponse
type DeviceVerificationResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take.
	Action *string `json:"action,omitempty"`
	// The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of `action` parameter. 
	ResponseContent *string `json:"responseContent,omitempty"`
	// The client ID of the client application to which the user code has been issued. 
	ClientId *int64 `json:"clientId,omitempty"`
	// The client ID alias of the client application to which the user code has been issued. 
	ClientIdAlias *string `json:"clientIdAlias,omitempty"`
	// `true` if the value of the `client_id` request parameter included in the device authorization request is the client ID alias. `false` if the value is the original numeric client ID. 
	ClientIdAliasUsed *bool `json:"clientIdAliasUsed,omitempty"`
	// The name of the client application to which the user code has been issued. 
	ClientName *string `json:"clientName,omitempty"`
	// The scopes requested by the device authorization request.  Note that `description` property and `descriptions` property of each scope object in the array contained in this property is always null even if descriptions of the scopes are registered. 
	Scopes []Scope `json:"scopes,omitempty"`
	// The names of the claims which were requested indirectly via some special scopes. See [5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) in OpenID Connect Core 1.0 for details.  This property is always `null` if the `scope` request parameter of the device authorization request does not include the `openid` scope even if special scopes (such as `profile`) are included in the request (unless the openid scope is included in the default set of scopes which is used when the `scope` request parameter is omitted). 
	ClaimNames []string `json:"claimNames,omitempty"`
	// The list of ACR values requested by the device authorization request. 
	Acrs []string `json:"acrs,omitempty"`
	// The resources specified by the `resource` request parameters or by the `resource` property in the request object. If both are given, the values in the request object should be set. See \"Resource Indicators for OAuth 2.0\" for details. 
	Resources []string `json:"resources,omitempty"`
	AuthorizationDetails *AuthorizationDetails `json:"authorizationDetails,omitempty"`
	// The attributes of this service that the client application belongs to. 
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`
	// The attributes of the client. 
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`
	// The dynamic scopes which the client application requested by the scope request parameter. 
	DynamicScopes []DynamicScope `json:"dynamicScopes,omitempty"`
	// Get the date in milliseconds since the Unix epoch (1970-01-01) at which the user code will expire. 
	ExpiresAt *int64 `json:"expiresAt,omitempty"`
}

// NewDeviceVerificationResponse instantiates a new DeviceVerificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceVerificationResponse() *DeviceVerificationResponse {
	this := DeviceVerificationResponse{}
	return &this
}

// NewDeviceVerificationResponseWithDefaults instantiates a new DeviceVerificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceVerificationResponseWithDefaults() *DeviceVerificationResponse {
	this := DeviceVerificationResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetResultCode() string {
	if o == nil || isNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasResultCode() bool {
	if o != nil && !isNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *DeviceVerificationResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetResultMessage() string {
	if o == nil || isNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || isNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasResultMessage() bool {
	if o != nil && !isNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *DeviceVerificationResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetAction() string {
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasAction() bool {
	if o != nil && !isNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *DeviceVerificationResponse) SetAction(v string) {
	o.Action = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetResponseContent() string {
	if o == nil || isNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || isNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasResponseContent() bool {
	if o != nil && !isNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *DeviceVerificationResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetClientId() int64 {
	if o == nil || isNil(o.ClientId) {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetClientIdOk() (*int64, bool) {
	if o == nil || isNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *DeviceVerificationResponse) SetClientId(v int64) {
	o.ClientId = &v
}

// GetClientIdAlias returns the ClientIdAlias field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetClientIdAlias() string {
	if o == nil || isNil(o.ClientIdAlias) {
		var ret string
		return ret
	}
	return *o.ClientIdAlias
}

// GetClientIdAliasOk returns a tuple with the ClientIdAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetClientIdAliasOk() (*string, bool) {
	if o == nil || isNil(o.ClientIdAlias) {
		return nil, false
	}
	return o.ClientIdAlias, true
}

// HasClientIdAlias returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasClientIdAlias() bool {
	if o != nil && !isNil(o.ClientIdAlias) {
		return true
	}

	return false
}

// SetClientIdAlias gets a reference to the given string and assigns it to the ClientIdAlias field.
func (o *DeviceVerificationResponse) SetClientIdAlias(v string) {
	o.ClientIdAlias = &v
}

// GetClientIdAliasUsed returns the ClientIdAliasUsed field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetClientIdAliasUsed() bool {
	if o == nil || isNil(o.ClientIdAliasUsed) {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasUsed
}

// GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetClientIdAliasUsedOk() (*bool, bool) {
	if o == nil || isNil(o.ClientIdAliasUsed) {
		return nil, false
	}
	return o.ClientIdAliasUsed, true
}

// HasClientIdAliasUsed returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasClientIdAliasUsed() bool {
	if o != nil && !isNil(o.ClientIdAliasUsed) {
		return true
	}

	return false
}

// SetClientIdAliasUsed gets a reference to the given bool and assigns it to the ClientIdAliasUsed field.
func (o *DeviceVerificationResponse) SetClientIdAliasUsed(v bool) {
	o.ClientIdAliasUsed = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetClientName() string {
	if o == nil || isNil(o.ClientName) {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetClientNameOk() (*string, bool) {
	if o == nil || isNil(o.ClientName) {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasClientName() bool {
	if o != nil && !isNil(o.ClientName) {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *DeviceVerificationResponse) SetClientName(v string) {
	o.ClientName = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetScopes() []Scope {
	if o == nil || isNil(o.Scopes) {
		var ret []Scope
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetScopesOk() ([]Scope, bool) {
	if o == nil || isNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasScopes() bool {
	if o != nil && !isNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []Scope and assigns it to the Scopes field.
func (o *DeviceVerificationResponse) SetScopes(v []Scope) {
	o.Scopes = v
}

// GetClaimNames returns the ClaimNames field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetClaimNames() []string {
	if o == nil || isNil(o.ClaimNames) {
		var ret []string
		return ret
	}
	return o.ClaimNames
}

// GetClaimNamesOk returns a tuple with the ClaimNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetClaimNamesOk() ([]string, bool) {
	if o == nil || isNil(o.ClaimNames) {
		return nil, false
	}
	return o.ClaimNames, true
}

// HasClaimNames returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasClaimNames() bool {
	if o != nil && !isNil(o.ClaimNames) {
		return true
	}

	return false
}

// SetClaimNames gets a reference to the given []string and assigns it to the ClaimNames field.
func (o *DeviceVerificationResponse) SetClaimNames(v []string) {
	o.ClaimNames = v
}

// GetAcrs returns the Acrs field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetAcrs() []string {
	if o == nil || isNil(o.Acrs) {
		var ret []string
		return ret
	}
	return o.Acrs
}

// GetAcrsOk returns a tuple with the Acrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetAcrsOk() ([]string, bool) {
	if o == nil || isNil(o.Acrs) {
		return nil, false
	}
	return o.Acrs, true
}

// HasAcrs returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasAcrs() bool {
	if o != nil && !isNil(o.Acrs) {
		return true
	}

	return false
}

// SetAcrs gets a reference to the given []string and assigns it to the Acrs field.
func (o *DeviceVerificationResponse) SetAcrs(v []string) {
	o.Acrs = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetResources() []string {
	if o == nil || isNil(o.Resources) {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetResourcesOk() ([]string, bool) {
	if o == nil || isNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasResources() bool {
	if o != nil && !isNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *DeviceVerificationResponse) SetResources(v []string) {
	o.Resources = v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetAuthorizationDetails() AuthorizationDetails {
	if o == nil || isNil(o.AuthorizationDetails) {
		var ret AuthorizationDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetAuthorizationDetailsOk() (*AuthorizationDetails, bool) {
	if o == nil || isNil(o.AuthorizationDetails) {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasAuthorizationDetails() bool {
	if o != nil && !isNil(o.AuthorizationDetails) {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthorizationDetails and assigns it to the AuthorizationDetails field.
func (o *DeviceVerificationResponse) SetAuthorizationDetails(v AuthorizationDetails) {
	o.AuthorizationDetails = &v
}

// GetServiceAttributes returns the ServiceAttributes field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetServiceAttributes() []Pair {
	if o == nil || isNil(o.ServiceAttributes) {
		var ret []Pair
		return ret
	}
	return o.ServiceAttributes
}

// GetServiceAttributesOk returns a tuple with the ServiceAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetServiceAttributesOk() ([]Pair, bool) {
	if o == nil || isNil(o.ServiceAttributes) {
		return nil, false
	}
	return o.ServiceAttributes, true
}

// HasServiceAttributes returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasServiceAttributes() bool {
	if o != nil && !isNil(o.ServiceAttributes) {
		return true
	}

	return false
}

// SetServiceAttributes gets a reference to the given []Pair and assigns it to the ServiceAttributes field.
func (o *DeviceVerificationResponse) SetServiceAttributes(v []Pair) {
	o.ServiceAttributes = v
}

// GetClientAttributes returns the ClientAttributes field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetClientAttributes() []Pair {
	if o == nil || isNil(o.ClientAttributes) {
		var ret []Pair
		return ret
	}
	return o.ClientAttributes
}

// GetClientAttributesOk returns a tuple with the ClientAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetClientAttributesOk() ([]Pair, bool) {
	if o == nil || isNil(o.ClientAttributes) {
		return nil, false
	}
	return o.ClientAttributes, true
}

// HasClientAttributes returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasClientAttributes() bool {
	if o != nil && !isNil(o.ClientAttributes) {
		return true
	}

	return false
}

// SetClientAttributes gets a reference to the given []Pair and assigns it to the ClientAttributes field.
func (o *DeviceVerificationResponse) SetClientAttributes(v []Pair) {
	o.ClientAttributes = v
}

// GetDynamicScopes returns the DynamicScopes field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetDynamicScopes() []DynamicScope {
	if o == nil || isNil(o.DynamicScopes) {
		var ret []DynamicScope
		return ret
	}
	return o.DynamicScopes
}

// GetDynamicScopesOk returns a tuple with the DynamicScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetDynamicScopesOk() ([]DynamicScope, bool) {
	if o == nil || isNil(o.DynamicScopes) {
		return nil, false
	}
	return o.DynamicScopes, true
}

// HasDynamicScopes returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasDynamicScopes() bool {
	if o != nil && !isNil(o.DynamicScopes) {
		return true
	}

	return false
}

// SetDynamicScopes gets a reference to the given []DynamicScope and assigns it to the DynamicScopes field.
func (o *DeviceVerificationResponse) SetDynamicScopes(v []DynamicScope) {
	o.DynamicScopes = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *DeviceVerificationResponse) GetExpiresAt() int64 {
	if o == nil || isNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceVerificationResponse) GetExpiresAtOk() (*int64, bool) {
	if o == nil || isNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *DeviceVerificationResponse) HasExpiresAt() bool {
	if o != nil && !isNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *DeviceVerificationResponse) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

func (o DeviceVerificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceVerificationResponse) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !isNil(o.ClientIdAlias) {
		toSerialize["clientIdAlias"] = o.ClientIdAlias
	}
	if !isNil(o.ClientIdAliasUsed) {
		toSerialize["clientIdAliasUsed"] = o.ClientIdAliasUsed
	}
	if !isNil(o.ClientName) {
		toSerialize["clientName"] = o.ClientName
	}
	if !isNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !isNil(o.ClaimNames) {
		toSerialize["claimNames"] = o.ClaimNames
	}
	if !isNil(o.Acrs) {
		toSerialize["acrs"] = o.Acrs
	}
	if !isNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !isNil(o.AuthorizationDetails) {
		toSerialize["authorizationDetails"] = o.AuthorizationDetails
	}
	if !isNil(o.ServiceAttributes) {
		toSerialize["serviceAttributes"] = o.ServiceAttributes
	}
	if !isNil(o.ClientAttributes) {
		toSerialize["clientAttributes"] = o.ClientAttributes
	}
	if !isNil(o.DynamicScopes) {
		toSerialize["dynamicScopes"] = o.DynamicScopes
	}
	if !isNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullableDeviceVerificationResponse struct {
	value *DeviceVerificationResponse
	isSet bool
}

func (v NullableDeviceVerificationResponse) Get() *DeviceVerificationResponse {
	return v.value
}

func (v *NullableDeviceVerificationResponse) Set(val *DeviceVerificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceVerificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceVerificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceVerificationResponse(val *DeviceVerificationResponse) *NullableDeviceVerificationResponse {
	return &NullableDeviceVerificationResponse{value: val, isSet: true}
}

func (v NullableDeviceVerificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceVerificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


