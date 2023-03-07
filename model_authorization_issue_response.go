/*
Authlete API

Authlete API Document. 

API version: 2.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

// checks if the AuthorizationIssueResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationIssueResponse{}

// AuthorizationIssueResponse struct for AuthorizationIssueResponse
type AuthorizationIssueResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take.
	Action *string `json:"action,omitempty"`
	// The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of `action` parameter. 
	ResponseContent *string `json:"responseContent,omitempty"`
	// The newly issued access token. Note that an access token is issued from an authorization endpoint only when `response_type` contains token. 
	AccessToken *string `json:"accessToken,omitempty"`
	// The datetime at which the newly issued access token will expire. The value is represented in milliseconds since the Unix epoch (1970-01-01). 
	AccessTokenExpiresAt *int64 `json:"accessTokenExpiresAt,omitempty"`
	// The duration of the newly issued access token in seconds. 
	AccessTokenDuration *int64 `json:"accessTokenDuration,omitempty"`
	// The newly issued ID token. Note that an ID token is issued from an authorization endpoint only when `response_type` contains `id_token`. 
	IdToken *string `json:"idToken,omitempty"`
	// The newly issued authorization code. Note that an authorization code is issued only when `response_type` contains code. 
	AuthorizationCode *string `json:"authorizationCode,omitempty"`
	// The newly issued access token in JWT format. If the service is not configured to issue JWT-based access tokens, this property is always set to `null`. 
	JwtAccessToken *string `json:"jwtAccessToken,omitempty"`
}

// NewAuthorizationIssueResponse instantiates a new AuthorizationIssueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationIssueResponse() *AuthorizationIssueResponse {
	this := AuthorizationIssueResponse{}
	return &this
}

// NewAuthorizationIssueResponseWithDefaults instantiates a new AuthorizationIssueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationIssueResponseWithDefaults() *AuthorizationIssueResponse {
	this := AuthorizationIssueResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *AuthorizationIssueResponse) GetResultCode() string {
	if o == nil || isNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasResultCode() bool {
	if o != nil && !isNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *AuthorizationIssueResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *AuthorizationIssueResponse) GetResultMessage() string {
	if o == nil || isNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || isNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasResultMessage() bool {
	if o != nil && !isNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *AuthorizationIssueResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AuthorizationIssueResponse) GetAction() string {
	if o == nil || isNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetActionOk() (*string, bool) {
	if o == nil || isNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasAction() bool {
	if o != nil && !isNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *AuthorizationIssueResponse) SetAction(v string) {
	o.Action = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *AuthorizationIssueResponse) GetResponseContent() string {
	if o == nil || isNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || isNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasResponseContent() bool {
	if o != nil && !isNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *AuthorizationIssueResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *AuthorizationIssueResponse) GetAccessToken() string {
	if o == nil || isNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil || isNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasAccessToken() bool {
	if o != nil && !isNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *AuthorizationIssueResponse) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetAccessTokenExpiresAt returns the AccessTokenExpiresAt field value if set, zero value otherwise.
func (o *AuthorizationIssueResponse) GetAccessTokenExpiresAt() int64 {
	if o == nil || isNil(o.AccessTokenExpiresAt) {
		var ret int64
		return ret
	}
	return *o.AccessTokenExpiresAt
}

// GetAccessTokenExpiresAtOk returns a tuple with the AccessTokenExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetAccessTokenExpiresAtOk() (*int64, bool) {
	if o == nil || isNil(o.AccessTokenExpiresAt) {
		return nil, false
	}
	return o.AccessTokenExpiresAt, true
}

// HasAccessTokenExpiresAt returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasAccessTokenExpiresAt() bool {
	if o != nil && !isNil(o.AccessTokenExpiresAt) {
		return true
	}

	return false
}

// SetAccessTokenExpiresAt gets a reference to the given int64 and assigns it to the AccessTokenExpiresAt field.
func (o *AuthorizationIssueResponse) SetAccessTokenExpiresAt(v int64) {
	o.AccessTokenExpiresAt = &v
}

// GetAccessTokenDuration returns the AccessTokenDuration field value if set, zero value otherwise.
func (o *AuthorizationIssueResponse) GetAccessTokenDuration() int64 {
	if o == nil || isNil(o.AccessTokenDuration) {
		var ret int64
		return ret
	}
	return *o.AccessTokenDuration
}

// GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetAccessTokenDurationOk() (*int64, bool) {
	if o == nil || isNil(o.AccessTokenDuration) {
		return nil, false
	}
	return o.AccessTokenDuration, true
}

// HasAccessTokenDuration returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasAccessTokenDuration() bool {
	if o != nil && !isNil(o.AccessTokenDuration) {
		return true
	}

	return false
}

// SetAccessTokenDuration gets a reference to the given int64 and assigns it to the AccessTokenDuration field.
func (o *AuthorizationIssueResponse) SetAccessTokenDuration(v int64) {
	o.AccessTokenDuration = &v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise.
func (o *AuthorizationIssueResponse) GetIdToken() string {
	if o == nil || isNil(o.IdToken) {
		var ret string
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetIdTokenOk() (*string, bool) {
	if o == nil || isNil(o.IdToken) {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasIdToken() bool {
	if o != nil && !isNil(o.IdToken) {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given string and assigns it to the IdToken field.
func (o *AuthorizationIssueResponse) SetIdToken(v string) {
	o.IdToken = &v
}

// GetAuthorizationCode returns the AuthorizationCode field value if set, zero value otherwise.
func (o *AuthorizationIssueResponse) GetAuthorizationCode() string {
	if o == nil || isNil(o.AuthorizationCode) {
		var ret string
		return ret
	}
	return *o.AuthorizationCode
}

// GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetAuthorizationCodeOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizationCode) {
		return nil, false
	}
	return o.AuthorizationCode, true
}

// HasAuthorizationCode returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasAuthorizationCode() bool {
	if o != nil && !isNil(o.AuthorizationCode) {
		return true
	}

	return false
}

// SetAuthorizationCode gets a reference to the given string and assigns it to the AuthorizationCode field.
func (o *AuthorizationIssueResponse) SetAuthorizationCode(v string) {
	o.AuthorizationCode = &v
}

// GetJwtAccessToken returns the JwtAccessToken field value if set, zero value otherwise.
func (o *AuthorizationIssueResponse) GetJwtAccessToken() string {
	if o == nil || isNil(o.JwtAccessToken) {
		var ret string
		return ret
	}
	return *o.JwtAccessToken
}

// GetJwtAccessTokenOk returns a tuple with the JwtAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetJwtAccessTokenOk() (*string, bool) {
	if o == nil || isNil(o.JwtAccessToken) {
		return nil, false
	}
	return o.JwtAccessToken, true
}

// HasJwtAccessToken returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasJwtAccessToken() bool {
	if o != nil && !isNil(o.JwtAccessToken) {
		return true
	}

	return false
}

// SetJwtAccessToken gets a reference to the given string and assigns it to the JwtAccessToken field.
func (o *AuthorizationIssueResponse) SetJwtAccessToken(v string) {
	o.JwtAccessToken = &v
}

func (o AuthorizationIssueResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationIssueResponse) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if !isNil(o.AccessTokenExpiresAt) {
		toSerialize["accessTokenExpiresAt"] = o.AccessTokenExpiresAt
	}
	if !isNil(o.AccessTokenDuration) {
		toSerialize["accessTokenDuration"] = o.AccessTokenDuration
	}
	if !isNil(o.IdToken) {
		toSerialize["idToken"] = o.IdToken
	}
	if !isNil(o.AuthorizationCode) {
		toSerialize["authorizationCode"] = o.AuthorizationCode
	}
	if !isNil(o.JwtAccessToken) {
		toSerialize["jwtAccessToken"] = o.JwtAccessToken
	}
	return toSerialize, nil
}

type NullableAuthorizationIssueResponse struct {
	value *AuthorizationIssueResponse
	isSet bool
}

func (v NullableAuthorizationIssueResponse) Get() *AuthorizationIssueResponse {
	return v.value
}

func (v *NullableAuthorizationIssueResponse) Set(val *AuthorizationIssueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationIssueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationIssueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationIssueResponse(val *AuthorizationIssueResponse) *NullableAuthorizationIssueResponse {
	return &NullableAuthorizationIssueResponse{value: val, isSet: true}
}

func (v NullableAuthorizationIssueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationIssueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


