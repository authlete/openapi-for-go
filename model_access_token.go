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

// checks if the AccessToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessToken{}

// AccessToken struct for AccessToken
type AccessToken struct {
	// The hash of the access token.
	AccessTokenHash *string `json:"accessTokenHash,omitempty"`
	// The timestamp at which the access token will expire.
	AccessTokenExpiresAt *int64 `json:"accessTokenExpiresAt,omitempty"`
	// The hash of the refresh token.
	RefreshTokenHash *string `json:"refreshTokenHash,omitempty"`
	// The timestamp at which the refresh token will expire.
	RefreshTokenExpiresAt *int64 `json:"refreshTokenExpiresAt,omitempty"`
	// The timestamp at which the access token was first created.
	CreatedAt *int64 `json:"createdAt,omitempty"`
	// The timestamp at which the access token was last refreshed using the refresh token.
	LastRefreshedAt *int64 `json:"lastRefreshedAt,omitempty"`
	// The ID of the client associated with the access token.
	ClientId *int64 `json:"clientId,omitempty"`
	// The subject (= unique user ID) associated with the access token.
	Subject   *string    `json:"subject,omitempty"`
	GrantType *GrantType `json:"grantType,omitempty"`
	// The scopes associated with the access token.
	Scopes []string `json:"scopes,omitempty"`
	// The properties associated with the access token.
	Properties []Property `json:"properties,omitempty"`
}

// NewAccessToken instantiates a new AccessToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessToken() *AccessToken {
	this := AccessToken{}
	return &this
}

// NewAccessTokenWithDefaults instantiates a new AccessToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTokenWithDefaults() *AccessToken {
	this := AccessToken{}
	return &this
}

// GetAccessTokenHash returns the AccessTokenHash field value if set, zero value otherwise.
func (o *AccessToken) GetAccessTokenHash() string {
	if o == nil || IsNil(o.AccessTokenHash) {
		var ret string
		return ret
	}
	return *o.AccessTokenHash
}

// GetAccessTokenHashOk returns a tuple with the AccessTokenHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetAccessTokenHashOk() (*string, bool) {
	if o == nil || IsNil(o.AccessTokenHash) {
		return nil, false
	}
	return o.AccessTokenHash, true
}

// HasAccessTokenHash returns a boolean if a field has been set.
func (o *AccessToken) HasAccessTokenHash() bool {
	if o != nil && !IsNil(o.AccessTokenHash) {
		return true
	}

	return false
}

// SetAccessTokenHash gets a reference to the given string and assigns it to the AccessTokenHash field.
func (o *AccessToken) SetAccessTokenHash(v string) {
	o.AccessTokenHash = &v
}

// GetAccessTokenExpiresAt returns the AccessTokenExpiresAt field value if set, zero value otherwise.
func (o *AccessToken) GetAccessTokenExpiresAt() int64 {
	if o == nil || IsNil(o.AccessTokenExpiresAt) {
		var ret int64
		return ret
	}
	return *o.AccessTokenExpiresAt
}

// GetAccessTokenExpiresAtOk returns a tuple with the AccessTokenExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetAccessTokenExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.AccessTokenExpiresAt) {
		return nil, false
	}
	return o.AccessTokenExpiresAt, true
}

// HasAccessTokenExpiresAt returns a boolean if a field has been set.
func (o *AccessToken) HasAccessTokenExpiresAt() bool {
	if o != nil && !IsNil(o.AccessTokenExpiresAt) {
		return true
	}

	return false
}

// SetAccessTokenExpiresAt gets a reference to the given int64 and assigns it to the AccessTokenExpiresAt field.
func (o *AccessToken) SetAccessTokenExpiresAt(v int64) {
	o.AccessTokenExpiresAt = &v
}

// GetRefreshTokenHash returns the RefreshTokenHash field value if set, zero value otherwise.
func (o *AccessToken) GetRefreshTokenHash() string {
	if o == nil || IsNil(o.RefreshTokenHash) {
		var ret string
		return ret
	}
	return *o.RefreshTokenHash
}

// GetRefreshTokenHashOk returns a tuple with the RefreshTokenHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetRefreshTokenHashOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshTokenHash) {
		return nil, false
	}
	return o.RefreshTokenHash, true
}

// HasRefreshTokenHash returns a boolean if a field has been set.
func (o *AccessToken) HasRefreshTokenHash() bool {
	if o != nil && !IsNil(o.RefreshTokenHash) {
		return true
	}

	return false
}

// SetRefreshTokenHash gets a reference to the given string and assigns it to the RefreshTokenHash field.
func (o *AccessToken) SetRefreshTokenHash(v string) {
	o.RefreshTokenHash = &v
}

// GetRefreshTokenExpiresAt returns the RefreshTokenExpiresAt field value if set, zero value otherwise.
func (o *AccessToken) GetRefreshTokenExpiresAt() int64 {
	if o == nil || IsNil(o.RefreshTokenExpiresAt) {
		var ret int64
		return ret
	}
	return *o.RefreshTokenExpiresAt
}

// GetRefreshTokenExpiresAtOk returns a tuple with the RefreshTokenExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetRefreshTokenExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.RefreshTokenExpiresAt) {
		return nil, false
	}
	return o.RefreshTokenExpiresAt, true
}

// HasRefreshTokenExpiresAt returns a boolean if a field has been set.
func (o *AccessToken) HasRefreshTokenExpiresAt() bool {
	if o != nil && !IsNil(o.RefreshTokenExpiresAt) {
		return true
	}

	return false
}

// SetRefreshTokenExpiresAt gets a reference to the given int64 and assigns it to the RefreshTokenExpiresAt field.
func (o *AccessToken) SetRefreshTokenExpiresAt(v int64) {
	o.RefreshTokenExpiresAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AccessToken) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AccessToken) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *AccessToken) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetLastRefreshedAt returns the LastRefreshedAt field value if set, zero value otherwise.
func (o *AccessToken) GetLastRefreshedAt() int64 {
	if o == nil || IsNil(o.LastRefreshedAt) {
		var ret int64
		return ret
	}
	return *o.LastRefreshedAt
}

// GetLastRefreshedAtOk returns a tuple with the LastRefreshedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetLastRefreshedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.LastRefreshedAt) {
		return nil, false
	}
	return o.LastRefreshedAt, true
}

// HasLastRefreshedAt returns a boolean if a field has been set.
func (o *AccessToken) HasLastRefreshedAt() bool {
	if o != nil && !IsNil(o.LastRefreshedAt) {
		return true
	}

	return false
}

// SetLastRefreshedAt gets a reference to the given int64 and assigns it to the LastRefreshedAt field.
func (o *AccessToken) SetLastRefreshedAt(v int64) {
	o.LastRefreshedAt = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AccessToken) GetClientId() int64 {
	if o == nil || IsNil(o.ClientId) {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetClientIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AccessToken) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *AccessToken) SetClientId(v int64) {
	o.ClientId = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *AccessToken) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *AccessToken) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *AccessToken) SetSubject(v string) {
	o.Subject = &v
}

// GetGrantType returns the GrantType field value if set, zero value otherwise.
func (o *AccessToken) GetGrantType() GrantType {
	if o == nil || IsNil(o.GrantType) {
		var ret GrantType
		return ret
	}
	return *o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetGrantTypeOk() (*GrantType, bool) {
	if o == nil || IsNil(o.GrantType) {
		return nil, false
	}
	return o.GrantType, true
}

// HasGrantType returns a boolean if a field has been set.
func (o *AccessToken) HasGrantType() bool {
	if o != nil && !IsNil(o.GrantType) {
		return true
	}

	return false
}

// SetGrantType gets a reference to the given GrantType and assigns it to the GrantType field.
func (o *AccessToken) SetGrantType(v GrantType) {
	o.GrantType = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *AccessToken) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *AccessToken) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *AccessToken) SetScopes(v []string) {
	o.Scopes = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *AccessToken) GetProperties() []Property {
	if o == nil || IsNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessToken) GetPropertiesOk() ([]Property, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *AccessToken) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *AccessToken) SetProperties(v []Property) {
	o.Properties = v
}

func (o AccessToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessTokenHash) {
		toSerialize["accessTokenHash"] = o.AccessTokenHash
	}
	if !IsNil(o.AccessTokenExpiresAt) {
		toSerialize["accessTokenExpiresAt"] = o.AccessTokenExpiresAt
	}
	if !IsNil(o.RefreshTokenHash) {
		toSerialize["refreshTokenHash"] = o.RefreshTokenHash
	}
	if !IsNil(o.RefreshTokenExpiresAt) {
		toSerialize["refreshTokenExpiresAt"] = o.RefreshTokenExpiresAt
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.LastRefreshedAt) {
		toSerialize["lastRefreshedAt"] = o.LastRefreshedAt
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.GrantType) {
		toSerialize["grantType"] = o.GrantType
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableAccessToken struct {
	value *AccessToken
	isSet bool
}

func (v NullableAccessToken) Get() *AccessToken {
	return v.value
}

func (v *NullableAccessToken) Set(val *AccessToken) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessToken) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessToken(val *AccessToken) *NullableAccessToken {
	return &NullableAccessToken{value: val, isSet: true}
}

func (v NullableAccessToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
