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

// checks if the BackchannelAuthenticationCompleteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackchannelAuthenticationCompleteResponse{}

// BackchannelAuthenticationCompleteResponse struct for BackchannelAuthenticationCompleteResponse
type BackchannelAuthenticationCompleteResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take. 
	Action *string `json:"action,omitempty"`
	// The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of `action` parameter. 
	ResponseContent *string `json:"responseContent,omitempty"`
	// The client ID of the client application that has made the backchannel authentication request. 
	ClientId *int64 `json:"clientId,omitempty"`
	// The client ID alias of the client application that has made the backchannel authentication request. 
	ClientIdAlias *string `json:"clientIdAlias,omitempty"`
	// `true` if the value of the client_id request parameter included in the backchannel authentication request is the client ID alias. `false` if the value is the original numeric client ID. 
	ClientIdAliasUsed *bool `json:"clientIdAliasUsed,omitempty"`
	// The name of the client application which has made the backchannel authentication request. 
	ClientName *string `json:"clientName,omitempty"`
	DeliveryMode *DeliveryMode `json:"deliveryMode,omitempty"`
	// The client notification endpoint to which a notification needs to be sent. This corresponds to the `client_notification_endpoint` metadata of the client application. 
	ClientNotificationEndpoint *string `json:"clientNotificationEndpoint,omitempty"`
	// The client notification token which needs to be embedded as a Bearer token in the Authorization header in the notification. This is the value of the `client_notification_token` request parameter included in the backchannel authentication request. 
	ClientNotificationToken *string `json:"clientNotificationToken,omitempty"`
	// The newly issued authentication request ID. 
	AuthReqId *string `json:"authReqId,omitempty"`
	// The issued access token. 
	AccessToken *string `json:"accessToken,omitempty"`
	// The issued refresh token. 
	RefreshToken *string `json:"refreshToken,omitempty"`
	// The issued ID token. 
	IdToken *string `json:"idToken,omitempty"`
	// The duration of the access token in seconds. 
	AccessTokenDuration *int64 `json:"accessTokenDuration,omitempty"`
	// The duration of the refresh token in seconds. 
	RefreshTokenDuration *int64 `json:"refreshTokenDuration,omitempty"`
	// The duration of the ID token in seconds. 
	IdTokenDuration *int64 `json:"idTokenDuration,omitempty"`
	// The issued access token in JWT format. 
	JwtAccessToken *string `json:"jwtAccessToken,omitempty"`
	// The resources specified by the `resource` request parameters or by the `resource` property in the request object. If both are given, the values in the request object should be set. See \"Resource Indicators for OAuth 2.0\" for details. 
	Resources []string `json:"resources,omitempty"`
	AuthorizationDetails *AuthzDetails `json:"authorizationDetails,omitempty"`
	// The attributes of this service that the client application belongs to. 
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`
	// The attributes of the client. 
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`
	// the value of the `grant_id` request parameter of the device authorization request.  The `grant_id` request parameter is defined in [Grant Management for OAuth 2.0](https://openid.net/specs/fapi-grant-management.html) , which is supported by Authlete 2.3 and newer versions. 
	GrantId *string `json:"grantId,omitempty"`
	// The entity ID of the client. 
	ClientEntityId *string `json:"clientEntityId,omitempty"`
	// Flag which indicates whether the entity ID of the client was used when the request for the access token was made.
	ClientEntityIdUsed *bool `json:"clientEntityIdUsed,omitempty"`
}

// NewBackchannelAuthenticationCompleteResponse instantiates a new BackchannelAuthenticationCompleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackchannelAuthenticationCompleteResponse() *BackchannelAuthenticationCompleteResponse {
	this := BackchannelAuthenticationCompleteResponse{}
	return &this
}

// NewBackchannelAuthenticationCompleteResponseWithDefaults instantiates a new BackchannelAuthenticationCompleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackchannelAuthenticationCompleteResponseWithDefaults() *BackchannelAuthenticationCompleteResponse {
	this := BackchannelAuthenticationCompleteResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetResultCode() string {
	if o == nil || isNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasResultCode() bool {
	if o != nil && !isNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *BackchannelAuthenticationCompleteResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetResultMessage() string {
	if o == nil || isNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || isNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasResultMessage() bool {
	if o != nil && !isNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *BackchannelAuthenticationCompleteResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetAction() string {
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasAction() bool {
	if o != nil && !isNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BackchannelAuthenticationCompleteResponse) SetAction(v string) {
	o.Action = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetResponseContent() string {
	if o == nil || isNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || isNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasResponseContent() bool {
	if o != nil && !isNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *BackchannelAuthenticationCompleteResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetClientId() int64 {
	if o == nil || isNil(o.ClientId) {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetClientIdOk() (*int64, bool) {
	if o == nil || isNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *BackchannelAuthenticationCompleteResponse) SetClientId(v int64) {
	o.ClientId = &v
}

// GetClientIdAlias returns the ClientIdAlias field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetClientIdAlias() string {
	if o == nil || isNil(o.ClientIdAlias) {
		var ret string
		return ret
	}
	return *o.ClientIdAlias
}

// GetClientIdAliasOk returns a tuple with the ClientIdAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetClientIdAliasOk() (*string, bool) {
	if o == nil || isNil(o.ClientIdAlias) {
		return nil, false
	}
	return o.ClientIdAlias, true
}

// HasClientIdAlias returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasClientIdAlias() bool {
	if o != nil && !isNil(o.ClientIdAlias) {
		return true
	}

	return false
}

// SetClientIdAlias gets a reference to the given string and assigns it to the ClientIdAlias field.
func (o *BackchannelAuthenticationCompleteResponse) SetClientIdAlias(v string) {
	o.ClientIdAlias = &v
}

// GetClientIdAliasUsed returns the ClientIdAliasUsed field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetClientIdAliasUsed() bool {
	if o == nil || isNil(o.ClientIdAliasUsed) {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasUsed
}

// GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetClientIdAliasUsedOk() (*bool, bool) {
	if o == nil || isNil(o.ClientIdAliasUsed) {
		return nil, false
	}
	return o.ClientIdAliasUsed, true
}

// HasClientIdAliasUsed returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasClientIdAliasUsed() bool {
	if o != nil && !isNil(o.ClientIdAliasUsed) {
		return true
	}

	return false
}

// SetClientIdAliasUsed gets a reference to the given bool and assigns it to the ClientIdAliasUsed field.
func (o *BackchannelAuthenticationCompleteResponse) SetClientIdAliasUsed(v bool) {
	o.ClientIdAliasUsed = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetClientName() string {
	if o == nil || isNil(o.ClientName) {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetClientNameOk() (*string, bool) {
	if o == nil || isNil(o.ClientName) {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasClientName() bool {
	if o != nil && !isNil(o.ClientName) {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *BackchannelAuthenticationCompleteResponse) SetClientName(v string) {
	o.ClientName = &v
}

// GetDeliveryMode returns the DeliveryMode field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetDeliveryMode() DeliveryMode {
	if o == nil || isNil(o.DeliveryMode) {
		var ret DeliveryMode
		return ret
	}
	return *o.DeliveryMode
}

// GetDeliveryModeOk returns a tuple with the DeliveryMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetDeliveryModeOk() (*DeliveryMode, bool) {
	if o == nil || isNil(o.DeliveryMode) {
		return nil, false
	}
	return o.DeliveryMode, true
}

// HasDeliveryMode returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasDeliveryMode() bool {
	if o != nil && !isNil(o.DeliveryMode) {
		return true
	}

	return false
}

// SetDeliveryMode gets a reference to the given DeliveryMode and assigns it to the DeliveryMode field.
func (o *BackchannelAuthenticationCompleteResponse) SetDeliveryMode(v DeliveryMode) {
	o.DeliveryMode = &v
}

// GetClientNotificationEndpoint returns the ClientNotificationEndpoint field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetClientNotificationEndpoint() string {
	if o == nil || isNil(o.ClientNotificationEndpoint) {
		var ret string
		return ret
	}
	return *o.ClientNotificationEndpoint
}

// GetClientNotificationEndpointOk returns a tuple with the ClientNotificationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetClientNotificationEndpointOk() (*string, bool) {
	if o == nil || isNil(o.ClientNotificationEndpoint) {
		return nil, false
	}
	return o.ClientNotificationEndpoint, true
}

// HasClientNotificationEndpoint returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasClientNotificationEndpoint() bool {
	if o != nil && !isNil(o.ClientNotificationEndpoint) {
		return true
	}

	return false
}

// SetClientNotificationEndpoint gets a reference to the given string and assigns it to the ClientNotificationEndpoint field.
func (o *BackchannelAuthenticationCompleteResponse) SetClientNotificationEndpoint(v string) {
	o.ClientNotificationEndpoint = &v
}

// GetClientNotificationToken returns the ClientNotificationToken field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetClientNotificationToken() string {
	if o == nil || isNil(o.ClientNotificationToken) {
		var ret string
		return ret
	}
	return *o.ClientNotificationToken
}

// GetClientNotificationTokenOk returns a tuple with the ClientNotificationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetClientNotificationTokenOk() (*string, bool) {
	if o == nil || isNil(o.ClientNotificationToken) {
		return nil, false
	}
	return o.ClientNotificationToken, true
}

// HasClientNotificationToken returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasClientNotificationToken() bool {
	if o != nil && !isNil(o.ClientNotificationToken) {
		return true
	}

	return false
}

// SetClientNotificationToken gets a reference to the given string and assigns it to the ClientNotificationToken field.
func (o *BackchannelAuthenticationCompleteResponse) SetClientNotificationToken(v string) {
	o.ClientNotificationToken = &v
}

// GetAuthReqId returns the AuthReqId field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetAuthReqId() string {
	if o == nil || isNil(o.AuthReqId) {
		var ret string
		return ret
	}
	return *o.AuthReqId
}

// GetAuthReqIdOk returns a tuple with the AuthReqId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetAuthReqIdOk() (*string, bool) {
	if o == nil || isNil(o.AuthReqId) {
		return nil, false
	}
	return o.AuthReqId, true
}

// HasAuthReqId returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasAuthReqId() bool {
	if o != nil && !isNil(o.AuthReqId) {
		return true
	}

	return false
}

// SetAuthReqId gets a reference to the given string and assigns it to the AuthReqId field.
func (o *BackchannelAuthenticationCompleteResponse) SetAuthReqId(v string) {
	o.AuthReqId = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetAccessToken() string {
	if o == nil || isNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil || isNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasAccessToken() bool {
	if o != nil && !isNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *BackchannelAuthenticationCompleteResponse) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetRefreshToken() string {
	if o == nil || isNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || isNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasRefreshToken() bool {
	if o != nil && !isNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *BackchannelAuthenticationCompleteResponse) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetIdToken() string {
	if o == nil || isNil(o.IdToken) {
		var ret string
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetIdTokenOk() (*string, bool) {
	if o == nil || isNil(o.IdToken) {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasIdToken() bool {
	if o != nil && !isNil(o.IdToken) {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given string and assigns it to the IdToken field.
func (o *BackchannelAuthenticationCompleteResponse) SetIdToken(v string) {
	o.IdToken = &v
}

// GetAccessTokenDuration returns the AccessTokenDuration field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetAccessTokenDuration() int64 {
	if o == nil || isNil(o.AccessTokenDuration) {
		var ret int64
		return ret
	}
	return *o.AccessTokenDuration
}

// GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetAccessTokenDurationOk() (*int64, bool) {
	if o == nil || isNil(o.AccessTokenDuration) {
		return nil, false
	}
	return o.AccessTokenDuration, true
}

// HasAccessTokenDuration returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasAccessTokenDuration() bool {
	if o != nil && !isNil(o.AccessTokenDuration) {
		return true
	}

	return false
}

// SetAccessTokenDuration gets a reference to the given int64 and assigns it to the AccessTokenDuration field.
func (o *BackchannelAuthenticationCompleteResponse) SetAccessTokenDuration(v int64) {
	o.AccessTokenDuration = &v
}

// GetRefreshTokenDuration returns the RefreshTokenDuration field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetRefreshTokenDuration() int64 {
	if o == nil || isNil(o.RefreshTokenDuration) {
		var ret int64
		return ret
	}
	return *o.RefreshTokenDuration
}

// GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetRefreshTokenDurationOk() (*int64, bool) {
	if o == nil || isNil(o.RefreshTokenDuration) {
		return nil, false
	}
	return o.RefreshTokenDuration, true
}

// HasRefreshTokenDuration returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasRefreshTokenDuration() bool {
	if o != nil && !isNil(o.RefreshTokenDuration) {
		return true
	}

	return false
}

// SetRefreshTokenDuration gets a reference to the given int64 and assigns it to the RefreshTokenDuration field.
func (o *BackchannelAuthenticationCompleteResponse) SetRefreshTokenDuration(v int64) {
	o.RefreshTokenDuration = &v
}

// GetIdTokenDuration returns the IdTokenDuration field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetIdTokenDuration() int64 {
	if o == nil || isNil(o.IdTokenDuration) {
		var ret int64
		return ret
	}
	return *o.IdTokenDuration
}

// GetIdTokenDurationOk returns a tuple with the IdTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetIdTokenDurationOk() (*int64, bool) {
	if o == nil || isNil(o.IdTokenDuration) {
		return nil, false
	}
	return o.IdTokenDuration, true
}

// HasIdTokenDuration returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasIdTokenDuration() bool {
	if o != nil && !isNil(o.IdTokenDuration) {
		return true
	}

	return false
}

// SetIdTokenDuration gets a reference to the given int64 and assigns it to the IdTokenDuration field.
func (o *BackchannelAuthenticationCompleteResponse) SetIdTokenDuration(v int64) {
	o.IdTokenDuration = &v
}

// GetJwtAccessToken returns the JwtAccessToken field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetJwtAccessToken() string {
	if o == nil || isNil(o.JwtAccessToken) {
		var ret string
		return ret
	}
	return *o.JwtAccessToken
}

// GetJwtAccessTokenOk returns a tuple with the JwtAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetJwtAccessTokenOk() (*string, bool) {
	if o == nil || isNil(o.JwtAccessToken) {
		return nil, false
	}
	return o.JwtAccessToken, true
}

// HasJwtAccessToken returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasJwtAccessToken() bool {
	if o != nil && !isNil(o.JwtAccessToken) {
		return true
	}

	return false
}

// SetJwtAccessToken gets a reference to the given string and assigns it to the JwtAccessToken field.
func (o *BackchannelAuthenticationCompleteResponse) SetJwtAccessToken(v string) {
	o.JwtAccessToken = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetResources() []string {
	if o == nil || isNil(o.Resources) {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetResourcesOk() ([]string, bool) {
	if o == nil || isNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasResources() bool {
	if o != nil && !isNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *BackchannelAuthenticationCompleteResponse) SetResources(v []string) {
	o.Resources = v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetAuthorizationDetails() AuthzDetails {
	if o == nil || isNil(o.AuthorizationDetails) {
		var ret AuthzDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetAuthorizationDetailsOk() (*AuthzDetails, bool) {
	if o == nil || isNil(o.AuthorizationDetails) {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasAuthorizationDetails() bool {
	if o != nil && !isNil(o.AuthorizationDetails) {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthzDetails and assigns it to the AuthorizationDetails field.
func (o *BackchannelAuthenticationCompleteResponse) SetAuthorizationDetails(v AuthzDetails) {
	o.AuthorizationDetails = &v
}

// GetServiceAttributes returns the ServiceAttributes field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetServiceAttributes() []Pair {
	if o == nil || isNil(o.ServiceAttributes) {
		var ret []Pair
		return ret
	}
	return o.ServiceAttributes
}

// GetServiceAttributesOk returns a tuple with the ServiceAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetServiceAttributesOk() ([]Pair, bool) {
	if o == nil || isNil(o.ServiceAttributes) {
		return nil, false
	}
	return o.ServiceAttributes, true
}

// HasServiceAttributes returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasServiceAttributes() bool {
	if o != nil && !isNil(o.ServiceAttributes) {
		return true
	}

	return false
}

// SetServiceAttributes gets a reference to the given []Pair and assigns it to the ServiceAttributes field.
func (o *BackchannelAuthenticationCompleteResponse) SetServiceAttributes(v []Pair) {
	o.ServiceAttributes = v
}

// GetClientAttributes returns the ClientAttributes field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetClientAttributes() []Pair {
	if o == nil || isNil(o.ClientAttributes) {
		var ret []Pair
		return ret
	}
	return o.ClientAttributes
}

// GetClientAttributesOk returns a tuple with the ClientAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetClientAttributesOk() ([]Pair, bool) {
	if o == nil || isNil(o.ClientAttributes) {
		return nil, false
	}
	return o.ClientAttributes, true
}

// HasClientAttributes returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasClientAttributes() bool {
	if o != nil && !isNil(o.ClientAttributes) {
		return true
	}

	return false
}

// SetClientAttributes gets a reference to the given []Pair and assigns it to the ClientAttributes field.
func (o *BackchannelAuthenticationCompleteResponse) SetClientAttributes(v []Pair) {
	o.ClientAttributes = v
}

// GetGrantId returns the GrantId field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetGrantId() string {
	if o == nil || isNil(o.GrantId) {
		var ret string
		return ret
	}
	return *o.GrantId
}

// GetGrantIdOk returns a tuple with the GrantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetGrantIdOk() (*string, bool) {
	if o == nil || isNil(o.GrantId) {
		return nil, false
	}
	return o.GrantId, true
}

// HasGrantId returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasGrantId() bool {
	if o != nil && !isNil(o.GrantId) {
		return true
	}

	return false
}

// SetGrantId gets a reference to the given string and assigns it to the GrantId field.
func (o *BackchannelAuthenticationCompleteResponse) SetGrantId(v string) {
	o.GrantId = &v
}

// GetClientEntityId returns the ClientEntityId field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetClientEntityId() string {
	if o == nil || isNil(o.ClientEntityId) {
		var ret string
		return ret
	}
	return *o.ClientEntityId
}

// GetClientEntityIdOk returns a tuple with the ClientEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetClientEntityIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientEntityId) {
		return nil, false
	}
	return o.ClientEntityId, true
}

// HasClientEntityId returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasClientEntityId() bool {
	if o != nil && !isNil(o.ClientEntityId) {
		return true
	}

	return false
}

// SetClientEntityId gets a reference to the given string and assigns it to the ClientEntityId field.
func (o *BackchannelAuthenticationCompleteResponse) SetClientEntityId(v string) {
	o.ClientEntityId = &v
}

// GetClientEntityIdUsed returns the ClientEntityIdUsed field value if set, zero value otherwise.
func (o *BackchannelAuthenticationCompleteResponse) GetClientEntityIdUsed() bool {
	if o == nil || isNil(o.ClientEntityIdUsed) {
		var ret bool
		return ret
	}
	return *o.ClientEntityIdUsed
}

// GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationCompleteResponse) GetClientEntityIdUsedOk() (*bool, bool) {
	if o == nil || isNil(o.ClientEntityIdUsed) {
		return nil, false
	}
	return o.ClientEntityIdUsed, true
}

// HasClientEntityIdUsed returns a boolean if a field has been set.
func (o *BackchannelAuthenticationCompleteResponse) HasClientEntityIdUsed() bool {
	if o != nil && !isNil(o.ClientEntityIdUsed) {
		return true
	}

	return false
}

// SetClientEntityIdUsed gets a reference to the given bool and assigns it to the ClientEntityIdUsed field.
func (o *BackchannelAuthenticationCompleteResponse) SetClientEntityIdUsed(v bool) {
	o.ClientEntityIdUsed = &v
}

func (o BackchannelAuthenticationCompleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackchannelAuthenticationCompleteResponse) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.DeliveryMode) {
		toSerialize["deliveryMode"] = o.DeliveryMode
	}
	if !isNil(o.ClientNotificationEndpoint) {
		toSerialize["clientNotificationEndpoint"] = o.ClientNotificationEndpoint
	}
	if !isNil(o.ClientNotificationToken) {
		toSerialize["clientNotificationToken"] = o.ClientNotificationToken
	}
	if !isNil(o.AuthReqId) {
		toSerialize["authReqId"] = o.AuthReqId
	}
	if !isNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if !isNil(o.RefreshToken) {
		toSerialize["refreshToken"] = o.RefreshToken
	}
	if !isNil(o.IdToken) {
		toSerialize["idToken"] = o.IdToken
	}
	if !isNil(o.AccessTokenDuration) {
		toSerialize["accessTokenDuration"] = o.AccessTokenDuration
	}
	if !isNil(o.RefreshTokenDuration) {
		toSerialize["refreshTokenDuration"] = o.RefreshTokenDuration
	}
	if !isNil(o.IdTokenDuration) {
		toSerialize["idTokenDuration"] = o.IdTokenDuration
	}
	if !isNil(o.JwtAccessToken) {
		toSerialize["jwtAccessToken"] = o.JwtAccessToken
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
	if !isNil(o.GrantId) {
		toSerialize["grantId"] = o.GrantId
	}
	if !isNil(o.ClientEntityId) {
		toSerialize["clientEntityId"] = o.ClientEntityId
	}
	if !isNil(o.ClientEntityIdUsed) {
		toSerialize["clientEntityIdUsed"] = o.ClientEntityIdUsed
	}
	return toSerialize, nil
}

type NullableBackchannelAuthenticationCompleteResponse struct {
	value *BackchannelAuthenticationCompleteResponse
	isSet bool
}

func (v NullableBackchannelAuthenticationCompleteResponse) Get() *BackchannelAuthenticationCompleteResponse {
	return v.value
}

func (v *NullableBackchannelAuthenticationCompleteResponse) Set(val *BackchannelAuthenticationCompleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBackchannelAuthenticationCompleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBackchannelAuthenticationCompleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackchannelAuthenticationCompleteResponse(val *BackchannelAuthenticationCompleteResponse) *NullableBackchannelAuthenticationCompleteResponse {
	return &NullableBackchannelAuthenticationCompleteResponse{value: val, isSet: true}
}

func (v NullableBackchannelAuthenticationCompleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackchannelAuthenticationCompleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


