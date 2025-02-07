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

// checks if the TokenUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenUpdateRequest{}

// TokenUpdateRequest struct for TokenUpdateRequest
type TokenUpdateRequest struct {
	// An access token.
	AccessToken string `json:"accessToken"`
	// A new date at which the access token will expire in milliseconds since the Unix epoch (1970-01-01). If the `accessTokenExpiresAt` request parameter is not included in a request or its value is 0 (or negative), the expiration date of the access token is not changed.
	AccessTokenExpiresAt *int64 `json:"accessTokenExpiresAt,omitempty"`
	// A new set of scopes assigned to the access token. Scopes that are not supported by the service and those that the client application associated with the access token is not allowed to request are ignored on the server side. If the `scopes` request parameter is not included in a request or its value is `null`, the scopes of the access token are not changed. Note that `properties` parameter is accepted only when `Content-Type` of the request is `application/json`, so don't use `application/x-www-form-urlencoded` if you want to specify `properties`.
	Scopes []string `json:"scopes,omitempty"`
	// A new set of properties assigned to the access token. If the `properties` request parameter is not included in a request or its value is null, the properties of the access token are not changed.
	Properties []Property `json:"properties,omitempty"`
	// A boolean request parameter which indicates whether the API attempts to update the expiration date of the access token when the scopes linked to the access token are changed by this request.
	AccessTokenExpiresAtUpdatedOnScopeUpdate *bool `json:"accessTokenExpiresAtUpdatedOnScopeUpdate,omitempty"`
	// The hash of the access token value. Used when the hash of the token is known (perhaps from lookup) but the value of the token itself is not. The value of the `accessToken` parameter takes precedence.
	AccessTokenHash *string `json:"accessTokenHash,omitempty"`
	// A boolean request parameter which indicates whether to update the value of the access token in the data store. If this parameter is set to `true` then a new access token value is generated by the server and returned in the response.
	AccessTokenValueUpdated *bool `json:"accessTokenValueUpdated,omitempty"`
	// The flag which indicates whether the access token expires or not. By default, all access tokens expire after a period of time determined by their service. If this request parameter is `true` then the access token will not automatically expire and must be revoked or deleted manually at the service.  If this request parameter is `true`, the `accessTokenExpiresAt` request parameter is ignored. If this request parameter is `false`, the `accessTokenExpiresAt` request parameter is processed normally.
	AccessTokenPersistent *bool `json:"accessTokenPersistent,omitempty"`
	// The thumbprint of the MTLS certificate bound to this token. If this property is set, a certificate with the corresponding value MUST be presented with the access token when it is used by a client. The value of this property must be a SHA256 certificate thumbprint, base64url encoded.
	CertificateThumbprint *string `json:"certificateThumbprint,omitempty"`
	// The thumbprint of the public key used for DPoP presentation of this token. If this property is set, a DPoP proof signed with the corresponding private key MUST be presented with the access token when it is used by a client. Additionally, the token's `token_type` will be set to 'DPoP'.
	DpopKeyThumbprint    *string       `json:"dpopKeyThumbprint,omitempty"`
	AuthorizationDetails *AuthzDetails `json:"authorizationDetails,omitempty"`
	// the flag which indicates whether the access token is for an external attachment.
	ForExternalAttachment *bool `json:"forExternalAttachment,omitempty"`
}

type _TokenUpdateRequest TokenUpdateRequest

// NewTokenUpdateRequest instantiates a new TokenUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenUpdateRequest(accessToken string) *TokenUpdateRequest {
	this := TokenUpdateRequest{}
	this.AccessToken = accessToken
	return &this
}

// NewTokenUpdateRequestWithDefaults instantiates a new TokenUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenUpdateRequestWithDefaults() *TokenUpdateRequest {
	this := TokenUpdateRequest{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *TokenUpdateRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TokenUpdateRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetAccessTokenExpiresAt returns the AccessTokenExpiresAt field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetAccessTokenExpiresAt() int64 {
	if o == nil || IsNil(o.AccessTokenExpiresAt) {
		var ret int64
		return ret
	}
	return *o.AccessTokenExpiresAt
}

// GetAccessTokenExpiresAtOk returns a tuple with the AccessTokenExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetAccessTokenExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.AccessTokenExpiresAt) {
		return nil, false
	}
	return o.AccessTokenExpiresAt, true
}

// HasAccessTokenExpiresAt returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasAccessTokenExpiresAt() bool {
	if o != nil && !IsNil(o.AccessTokenExpiresAt) {
		return true
	}

	return false
}

// SetAccessTokenExpiresAt gets a reference to the given int64 and assigns it to the AccessTokenExpiresAt field.
func (o *TokenUpdateRequest) SetAccessTokenExpiresAt(v int64) {
	o.AccessTokenExpiresAt = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *TokenUpdateRequest) SetScopes(v []string) {
	o.Scopes = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetProperties() []Property {
	if o == nil || IsNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetPropertiesOk() ([]Property, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *TokenUpdateRequest) SetProperties(v []Property) {
	o.Properties = v
}

// GetAccessTokenExpiresAtUpdatedOnScopeUpdate returns the AccessTokenExpiresAtUpdatedOnScopeUpdate field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetAccessTokenExpiresAtUpdatedOnScopeUpdate() bool {
	if o == nil || IsNil(o.AccessTokenExpiresAtUpdatedOnScopeUpdate) {
		var ret bool
		return ret
	}
	return *o.AccessTokenExpiresAtUpdatedOnScopeUpdate
}

// GetAccessTokenExpiresAtUpdatedOnScopeUpdateOk returns a tuple with the AccessTokenExpiresAtUpdatedOnScopeUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetAccessTokenExpiresAtUpdatedOnScopeUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.AccessTokenExpiresAtUpdatedOnScopeUpdate) {
		return nil, false
	}
	return o.AccessTokenExpiresAtUpdatedOnScopeUpdate, true
}

// HasAccessTokenExpiresAtUpdatedOnScopeUpdate returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasAccessTokenExpiresAtUpdatedOnScopeUpdate() bool {
	if o != nil && !IsNil(o.AccessTokenExpiresAtUpdatedOnScopeUpdate) {
		return true
	}

	return false
}

// SetAccessTokenExpiresAtUpdatedOnScopeUpdate gets a reference to the given bool and assigns it to the AccessTokenExpiresAtUpdatedOnScopeUpdate field.
func (o *TokenUpdateRequest) SetAccessTokenExpiresAtUpdatedOnScopeUpdate(v bool) {
	o.AccessTokenExpiresAtUpdatedOnScopeUpdate = &v
}

// GetAccessTokenHash returns the AccessTokenHash field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetAccessTokenHash() string {
	if o == nil || IsNil(o.AccessTokenHash) {
		var ret string
		return ret
	}
	return *o.AccessTokenHash
}

// GetAccessTokenHashOk returns a tuple with the AccessTokenHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetAccessTokenHashOk() (*string, bool) {
	if o == nil || IsNil(o.AccessTokenHash) {
		return nil, false
	}
	return o.AccessTokenHash, true
}

// HasAccessTokenHash returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasAccessTokenHash() bool {
	if o != nil && !IsNil(o.AccessTokenHash) {
		return true
	}

	return false
}

// SetAccessTokenHash gets a reference to the given string and assigns it to the AccessTokenHash field.
func (o *TokenUpdateRequest) SetAccessTokenHash(v string) {
	o.AccessTokenHash = &v
}

// GetAccessTokenValueUpdated returns the AccessTokenValueUpdated field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetAccessTokenValueUpdated() bool {
	if o == nil || IsNil(o.AccessTokenValueUpdated) {
		var ret bool
		return ret
	}
	return *o.AccessTokenValueUpdated
}

// GetAccessTokenValueUpdatedOk returns a tuple with the AccessTokenValueUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetAccessTokenValueUpdatedOk() (*bool, bool) {
	if o == nil || IsNil(o.AccessTokenValueUpdated) {
		return nil, false
	}
	return o.AccessTokenValueUpdated, true
}

// HasAccessTokenValueUpdated returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasAccessTokenValueUpdated() bool {
	if o != nil && !IsNil(o.AccessTokenValueUpdated) {
		return true
	}

	return false
}

// SetAccessTokenValueUpdated gets a reference to the given bool and assigns it to the AccessTokenValueUpdated field.
func (o *TokenUpdateRequest) SetAccessTokenValueUpdated(v bool) {
	o.AccessTokenValueUpdated = &v
}

// GetAccessTokenPersistent returns the AccessTokenPersistent field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetAccessTokenPersistent() bool {
	if o == nil || IsNil(o.AccessTokenPersistent) {
		var ret bool
		return ret
	}
	return *o.AccessTokenPersistent
}

// GetAccessTokenPersistentOk returns a tuple with the AccessTokenPersistent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetAccessTokenPersistentOk() (*bool, bool) {
	if o == nil || IsNil(o.AccessTokenPersistent) {
		return nil, false
	}
	return o.AccessTokenPersistent, true
}

// HasAccessTokenPersistent returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasAccessTokenPersistent() bool {
	if o != nil && !IsNil(o.AccessTokenPersistent) {
		return true
	}

	return false
}

// SetAccessTokenPersistent gets a reference to the given bool and assigns it to the AccessTokenPersistent field.
func (o *TokenUpdateRequest) SetAccessTokenPersistent(v bool) {
	o.AccessTokenPersistent = &v
}

// GetCertificateThumbprint returns the CertificateThumbprint field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetCertificateThumbprint() string {
	if o == nil || IsNil(o.CertificateThumbprint) {
		var ret string
		return ret
	}
	return *o.CertificateThumbprint
}

// GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetCertificateThumbprintOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateThumbprint) {
		return nil, false
	}
	return o.CertificateThumbprint, true
}

// HasCertificateThumbprint returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasCertificateThumbprint() bool {
	if o != nil && !IsNil(o.CertificateThumbprint) {
		return true
	}

	return false
}

// SetCertificateThumbprint gets a reference to the given string and assigns it to the CertificateThumbprint field.
func (o *TokenUpdateRequest) SetCertificateThumbprint(v string) {
	o.CertificateThumbprint = &v
}

// GetDpopKeyThumbprint returns the DpopKeyThumbprint field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetDpopKeyThumbprint() string {
	if o == nil || IsNil(o.DpopKeyThumbprint) {
		var ret string
		return ret
	}
	return *o.DpopKeyThumbprint
}

// GetDpopKeyThumbprintOk returns a tuple with the DpopKeyThumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetDpopKeyThumbprintOk() (*string, bool) {
	if o == nil || IsNil(o.DpopKeyThumbprint) {
		return nil, false
	}
	return o.DpopKeyThumbprint, true
}

// HasDpopKeyThumbprint returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasDpopKeyThumbprint() bool {
	if o != nil && !IsNil(o.DpopKeyThumbprint) {
		return true
	}

	return false
}

// SetDpopKeyThumbprint gets a reference to the given string and assigns it to the DpopKeyThumbprint field.
func (o *TokenUpdateRequest) SetDpopKeyThumbprint(v string) {
	o.DpopKeyThumbprint = &v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetAuthorizationDetails() AuthzDetails {
	if o == nil || IsNil(o.AuthorizationDetails) {
		var ret AuthzDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetAuthorizationDetailsOk() (*AuthzDetails, bool) {
	if o == nil || IsNil(o.AuthorizationDetails) {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasAuthorizationDetails() bool {
	if o != nil && !IsNil(o.AuthorizationDetails) {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthzDetails and assigns it to the AuthorizationDetails field.
func (o *TokenUpdateRequest) SetAuthorizationDetails(v AuthzDetails) {
	o.AuthorizationDetails = &v
}

// GetForExternalAttachment returns the ForExternalAttachment field value if set, zero value otherwise.
func (o *TokenUpdateRequest) GetForExternalAttachment() bool {
	if o == nil || IsNil(o.ForExternalAttachment) {
		var ret bool
		return ret
	}
	return *o.ForExternalAttachment
}

// GetForExternalAttachmentOk returns a tuple with the ForExternalAttachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenUpdateRequest) GetForExternalAttachmentOk() (*bool, bool) {
	if o == nil || IsNil(o.ForExternalAttachment) {
		return nil, false
	}
	return o.ForExternalAttachment, true
}

// HasForExternalAttachment returns a boolean if a field has been set.
func (o *TokenUpdateRequest) HasForExternalAttachment() bool {
	if o != nil && !IsNil(o.ForExternalAttachment) {
		return true
	}

	return false
}

// SetForExternalAttachment gets a reference to the given bool and assigns it to the ForExternalAttachment field.
func (o *TokenUpdateRequest) SetForExternalAttachment(v bool) {
	o.ForExternalAttachment = &v
}

func (o TokenUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accessToken"] = o.AccessToken
	if !IsNil(o.AccessTokenExpiresAt) {
		toSerialize["accessTokenExpiresAt"] = o.AccessTokenExpiresAt
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.AccessTokenExpiresAtUpdatedOnScopeUpdate) {
		toSerialize["accessTokenExpiresAtUpdatedOnScopeUpdate"] = o.AccessTokenExpiresAtUpdatedOnScopeUpdate
	}
	if !IsNil(o.AccessTokenHash) {
		toSerialize["accessTokenHash"] = o.AccessTokenHash
	}
	if !IsNil(o.AccessTokenValueUpdated) {
		toSerialize["accessTokenValueUpdated"] = o.AccessTokenValueUpdated
	}
	if !IsNil(o.AccessTokenPersistent) {
		toSerialize["accessTokenPersistent"] = o.AccessTokenPersistent
	}
	if !IsNil(o.CertificateThumbprint) {
		toSerialize["certificateThumbprint"] = o.CertificateThumbprint
	}
	if !IsNil(o.DpopKeyThumbprint) {
		toSerialize["dpopKeyThumbprint"] = o.DpopKeyThumbprint
	}
	if !IsNil(o.AuthorizationDetails) {
		toSerialize["authorizationDetails"] = o.AuthorizationDetails
	}
	if !IsNil(o.ForExternalAttachment) {
		toSerialize["forExternalAttachment"] = o.ForExternalAttachment
	}
	return toSerialize, nil
}

func (o *TokenUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accessToken",
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

	varTokenUpdateRequest := _TokenUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenUpdateRequest)

	if err != nil {
		return err
	}

	*o = TokenUpdateRequest(varTokenUpdateRequest)

	return err
}

type NullableTokenUpdateRequest struct {
	value *TokenUpdateRequest
	isSet bool
}

func (v NullableTokenUpdateRequest) Get() *TokenUpdateRequest {
	return v.value
}

func (v *NullableTokenUpdateRequest) Set(val *TokenUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenUpdateRequest(val *TokenUpdateRequest) *NullableTokenUpdateRequest {
	return &NullableTokenUpdateRequest{value: val, isSet: true}
}

func (v NullableTokenUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
