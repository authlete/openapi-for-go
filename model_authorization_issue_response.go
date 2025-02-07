/*
Authlete API Explorer

<div class=\"min-h-screen bg-gray-100 dark:bg-gray-900 text-gray-900 dark:text-gray-100 p-6\">   <div class=\"flex justify-end mb-4\">     <label for=\"theme-toggle\" class=\"flex items-center cursor-pointer\">       <div class=\"relative\">Dark mode:         <input type=\"checkbox\" id=\"theme-toggle\" class=\"sr-only\" onchange=\"toggleTheme()\">         <div class=\"block bg-gray-600 w-14 h-8 rounded-full\"></div>         <div class=\"dot absolute left-1 top-1 bg-white w-6 h-6 rounded-full transition\"></div>       </div>     </label>   </div>   <header class=\"bg-green-500 dark:bg-green-700 p-4 rounded-lg text-white text-center\">     <p>       Welcome to the <strong>Authlete API documentation</strong>. Authlete is an <strong>API-first service</strong>       where every aspect of the platform is configurable via API. This explorer provides a convenient way to       authenticate and interact with the API, allowing you to see Authlete in action quickly. 🚀     </p>     <p>       At a high level, the Authlete API is grouped into two categories:     </p>     <ul class=\"list-disc list-inside\">       <li><strong>Management APIs</strong>: Enable you to manage services and clients. 🔧</li>       <li><strong>Runtime APIs</strong>: Allow you to build your own Authorization Servers or Verifiable Credential (VC)         issuers. 🔐</li>     </ul>     <p>All API endpoints are secured using access tokens issued by Authlete's Identity Provider (IdP). If you already       have an Authlete account, simply use the <em>Get Token</em> option on the Authentication page to log in and obtain       an access token for API usage. If you don't have an account yet, <a href=\"https://console.authlete.com/register\">sign up         here</a> to get started.</p>   </header>   <main>     <section id=\"api-servers\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🌐 API Servers</h2>       <p>Authlete is a global service with clusters available in multiple regions across the world.</p>       <p>Currently, our service is available in the following regions:</p>       <div class=\"grid grid-cols-2 gap-4\">         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇺🇸 US</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇯🇵 JP</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇪🇺 EU</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇧🇷 Brazil</p>         </div>       </div>       <p>Our customers can host their data in the region that best meets their requirements.</p>       <a href=\"#servers\" class=\"block mt-4 text-green-500 dark:text-green-300 hover:underline text-center\">Select your         preferred server</a>     </section>     <section id=\"authentication\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🔑 Authentication</h2>       <p>The API Explorer requires an access token to call the API.</p>       <p>You can create the access token from the <a href=\"https://console.authlete.com\">Authlete Management Console</a> and set it in the HTTP Bearer section of Authentication page.</p>       <p>Alternatively, if you have an Authlete account, the API Explorer can log you in with your Authlete account and         automatically acquire the required access token.</p>       <div class=\"theme-admonition theme-admonition-warning admonition_o5H7 alert alert--warning\">         <div class=\"admonitionContent_Knsx\">           <p>⚠️ <strong>Important Note:</strong> When the API Explorer acquires the token after login, the access tokens             will have the same permissions as the user who logs in as part of this flow.</p>         </div>       </div>       <a href=\"#auth\" class=\"block mt-4 text-green-500 dark:text-green-300 hover:underline text-center\">Setup your         access token</a>     </section>     <section id=\"tutorials\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🎓 Tutorials</h2>       <p>If you have successfully tested the API from the API Console and want to take the next step of integrating the         API into your application, or if you want to see a sample using Authlete APIs, follow the links below. These         resources will help you understand key concepts and how to integrate Authlete API into your applications.</p>       <div class=\"mt-4\">         <a href=\"https://www.authlete.com/developers/getting_started/\"           class=\"block text-green-500 dark:text-green-300 font-bold hover:underline mb-2\">🚀 Getting Started with           Authlete</a>           </br>         <a href=\"https://www.authlete.com/developers/tutorial/signup/\"           class=\"block text-green-500 dark:text-green-300 font-bold hover:underline\">🔑 From Sign-Up to the First API           Request</a>       </div>     </section>     <section id=\"support\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🛠 Contact Us</h2>       <p>If you have any questions or need assistance, our team is here to help.</p>       <a href=\"https://www.authlete.com/contact/\"         class=\"block mt-4 text-green-500 dark:text-green-300 font-bold hover:underline\">Contact Page</a>     </section>   </main> </div>

API version: 3.0.0
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
	// The information about the ticket.
	TicketInfo *string `json:"ticketInfo,omitempty"`
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
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
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
	if o == nil || IsNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasResultMessage() bool {
	if o != nil && !IsNil(o.ResultMessage) {
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
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
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
	if o == nil || IsNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasResponseContent() bool {
	if o != nil && !IsNil(o.ResponseContent) {
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
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
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
	if o == nil || IsNil(o.AccessTokenExpiresAt) {
		var ret int64
		return ret
	}
	return *o.AccessTokenExpiresAt
}

// GetAccessTokenExpiresAtOk returns a tuple with the AccessTokenExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetAccessTokenExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.AccessTokenExpiresAt) {
		return nil, false
	}
	return o.AccessTokenExpiresAt, true
}

// HasAccessTokenExpiresAt returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasAccessTokenExpiresAt() bool {
	if o != nil && !IsNil(o.AccessTokenExpiresAt) {
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
	if o == nil || IsNil(o.AccessTokenDuration) {
		var ret int64
		return ret
	}
	return *o.AccessTokenDuration
}

// GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetAccessTokenDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.AccessTokenDuration) {
		return nil, false
	}
	return o.AccessTokenDuration, true
}

// HasAccessTokenDuration returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasAccessTokenDuration() bool {
	if o != nil && !IsNil(o.AccessTokenDuration) {
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
	if o == nil || IsNil(o.IdToken) {
		var ret string
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetIdTokenOk() (*string, bool) {
	if o == nil || IsNil(o.IdToken) {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasIdToken() bool {
	if o != nil && !IsNil(o.IdToken) {
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
	if o == nil || IsNil(o.AuthorizationCode) {
		var ret string
		return ret
	}
	return *o.AuthorizationCode
}

// GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetAuthorizationCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationCode) {
		return nil, false
	}
	return o.AuthorizationCode, true
}

// HasAuthorizationCode returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasAuthorizationCode() bool {
	if o != nil && !IsNil(o.AuthorizationCode) {
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
	if o == nil || IsNil(o.JwtAccessToken) {
		var ret string
		return ret
	}
	return *o.JwtAccessToken
}

// GetJwtAccessTokenOk returns a tuple with the JwtAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetJwtAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.JwtAccessToken) {
		return nil, false
	}
	return o.JwtAccessToken, true
}

// HasJwtAccessToken returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasJwtAccessToken() bool {
	if o != nil && !IsNil(o.JwtAccessToken) {
		return true
	}

	return false
}

// SetJwtAccessToken gets a reference to the given string and assigns it to the JwtAccessToken field.
func (o *AuthorizationIssueResponse) SetJwtAccessToken(v string) {
	o.JwtAccessToken = &v
}

// GetTicketInfo returns the TicketInfo field value if set, zero value otherwise.
func (o *AuthorizationIssueResponse) GetTicketInfo() string {
	if o == nil || IsNil(o.TicketInfo) {
		var ret string
		return ret
	}
	return *o.TicketInfo
}

// GetTicketInfoOk returns a tuple with the TicketInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueResponse) GetTicketInfoOk() (*string, bool) {
	if o == nil || IsNil(o.TicketInfo) {
		return nil, false
	}
	return o.TicketInfo, true
}

// HasTicketInfo returns a boolean if a field has been set.
func (o *AuthorizationIssueResponse) HasTicketInfo() bool {
	if o != nil && !IsNil(o.TicketInfo) {
		return true
	}

	return false
}

// SetTicketInfo gets a reference to the given string and assigns it to the TicketInfo field.
func (o *AuthorizationIssueResponse) SetTicketInfo(v string) {
	o.TicketInfo = &v
}

func (o AuthorizationIssueResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationIssueResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultCode) {
		toSerialize["resultCode"] = o.ResultCode
	}
	if !IsNil(o.ResultMessage) {
		toSerialize["resultMessage"] = o.ResultMessage
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.ResponseContent) {
		toSerialize["responseContent"] = o.ResponseContent
	}
	if !IsNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if !IsNil(o.AccessTokenExpiresAt) {
		toSerialize["accessTokenExpiresAt"] = o.AccessTokenExpiresAt
	}
	if !IsNil(o.AccessTokenDuration) {
		toSerialize["accessTokenDuration"] = o.AccessTokenDuration
	}
	if !IsNil(o.IdToken) {
		toSerialize["idToken"] = o.IdToken
	}
	if !IsNil(o.AuthorizationCode) {
		toSerialize["authorizationCode"] = o.AuthorizationCode
	}
	if !IsNil(o.JwtAccessToken) {
		toSerialize["jwtAccessToken"] = o.JwtAccessToken
	}
	if !IsNil(o.TicketInfo) {
		toSerialize["ticketInfo"] = o.TicketInfo
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
