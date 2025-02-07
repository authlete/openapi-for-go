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

// checks if the AuthorizationIssueRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationIssueRequest{}

// AuthorizationIssueRequest struct for AuthorizationIssueRequest
type AuthorizationIssueRequest struct {
	// The ticket issued from Authlete `/auth/authorization` API.
	Ticket string `json:"ticket"`
	// The subject (= a user account managed by the service) who has granted authorization to the client application.
	Subject string `json:"subject"`
	// The time when the authentication of the end-user occurred. Its value is the number of seconds from `1970-01-01`.
	AuthTime *int64 `json:"authTime,omitempty"`
	// The Authentication Context Class Reference performed for the end-user authentication.
	Acr *string `json:"acr,omitempty"`
	// The claims of the end-user (= pieces of information about the end-user) in JSON format. See [OpenID Connect Core 1.0, 5.1. Standard Claims](https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims) for details about the format.
	Claims *string `json:"claims,omitempty"`
	// Extra properties to associate with an access token and/or an authorization code.
	Properties []Property `json:"properties,omitempty"`
	// Scopes to associate with an access token and/or an authorization code. If a non-empty string array is given, it replaces the scopes specified by the original authorization request.
	Scopes []string `json:"scopes,omitempty"`
	// The value of the `sub` claim to embed in an ID token. If this request parameter is `null` or empty, the value of the `subject` request parameter is used as the value of the `sub` claim.
	Sub *string `json:"sub,omitempty"`
	// JSON that represents additional JWS header parameters for ID tokens that may be issued based on the authorization request.
	IdtHeaderParams *string `json:"idtHeaderParams,omitempty"`
	// Claim key-value pairs that are used to compute transformed claims.
	ClaimsForTx *string `json:"claimsForTx,omitempty"`
	// the claims that the user has consented for the client application to know.
	ConsentedClaims      []string      `json:"consentedClaims,omitempty"`
	AuthorizationDetails *AuthzDetails `json:"authorizationDetails,omitempty"`
	// Additional claims that are added to the payload part of the JWT access token.
	JwtAtClaims *string `json:"jwtAtClaims,omitempty"`
	// The representation of an access token that may be issued as a result of the Authlete API call.
	AccessToken *string `json:"accessToken,omitempty"`
}

type _AuthorizationIssueRequest AuthorizationIssueRequest

// NewAuthorizationIssueRequest instantiates a new AuthorizationIssueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationIssueRequest(ticket string, subject string) *AuthorizationIssueRequest {
	this := AuthorizationIssueRequest{}
	this.Ticket = ticket
	this.Subject = subject
	return &this
}

// NewAuthorizationIssueRequestWithDefaults instantiates a new AuthorizationIssueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationIssueRequestWithDefaults() *AuthorizationIssueRequest {
	this := AuthorizationIssueRequest{}
	return &this
}

// GetTicket returns the Ticket field value
func (o *AuthorizationIssueRequest) GetTicket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ticket
}

// GetTicketOk returns a tuple with the Ticket field value
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetTicketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ticket, true
}

// SetTicket sets field value
func (o *AuthorizationIssueRequest) SetTicket(v string) {
	o.Ticket = v
}

// GetSubject returns the Subject field value
func (o *AuthorizationIssueRequest) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *AuthorizationIssueRequest) SetSubject(v string) {
	o.Subject = v
}

// GetAuthTime returns the AuthTime field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetAuthTime() int64 {
	if o == nil || IsNil(o.AuthTime) {
		var ret int64
		return ret
	}
	return *o.AuthTime
}

// GetAuthTimeOk returns a tuple with the AuthTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetAuthTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.AuthTime) {
		return nil, false
	}
	return o.AuthTime, true
}

// HasAuthTime returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasAuthTime() bool {
	if o != nil && !IsNil(o.AuthTime) {
		return true
	}

	return false
}

// SetAuthTime gets a reference to the given int64 and assigns it to the AuthTime field.
func (o *AuthorizationIssueRequest) SetAuthTime(v int64) {
	o.AuthTime = &v
}

// GetAcr returns the Acr field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetAcr() string {
	if o == nil || IsNil(o.Acr) {
		var ret string
		return ret
	}
	return *o.Acr
}

// GetAcrOk returns a tuple with the Acr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetAcrOk() (*string, bool) {
	if o == nil || IsNil(o.Acr) {
		return nil, false
	}
	return o.Acr, true
}

// HasAcr returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasAcr() bool {
	if o != nil && !IsNil(o.Acr) {
		return true
	}

	return false
}

// SetAcr gets a reference to the given string and assigns it to the Acr field.
func (o *AuthorizationIssueRequest) SetAcr(v string) {
	o.Acr = &v
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetClaims() string {
	if o == nil || IsNil(o.Claims) {
		var ret string
		return ret
	}
	return *o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetClaimsOk() (*string, bool) {
	if o == nil || IsNil(o.Claims) {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasClaims() bool {
	if o != nil && !IsNil(o.Claims) {
		return true
	}

	return false
}

// SetClaims gets a reference to the given string and assigns it to the Claims field.
func (o *AuthorizationIssueRequest) SetClaims(v string) {
	o.Claims = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetProperties() []Property {
	if o == nil || IsNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetPropertiesOk() ([]Property, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *AuthorizationIssueRequest) SetProperties(v []Property) {
	o.Properties = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *AuthorizationIssueRequest) SetScopes(v []string) {
	o.Scopes = v
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetSub() string {
	if o == nil || IsNil(o.Sub) {
		var ret string
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetSubOk() (*string, bool) {
	if o == nil || IsNil(o.Sub) {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasSub() bool {
	if o != nil && !IsNil(o.Sub) {
		return true
	}

	return false
}

// SetSub gets a reference to the given string and assigns it to the Sub field.
func (o *AuthorizationIssueRequest) SetSub(v string) {
	o.Sub = &v
}

// GetIdtHeaderParams returns the IdtHeaderParams field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetIdtHeaderParams() string {
	if o == nil || IsNil(o.IdtHeaderParams) {
		var ret string
		return ret
	}
	return *o.IdtHeaderParams
}

// GetIdtHeaderParamsOk returns a tuple with the IdtHeaderParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetIdtHeaderParamsOk() (*string, bool) {
	if o == nil || IsNil(o.IdtHeaderParams) {
		return nil, false
	}
	return o.IdtHeaderParams, true
}

// HasIdtHeaderParams returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasIdtHeaderParams() bool {
	if o != nil && !IsNil(o.IdtHeaderParams) {
		return true
	}

	return false
}

// SetIdtHeaderParams gets a reference to the given string and assigns it to the IdtHeaderParams field.
func (o *AuthorizationIssueRequest) SetIdtHeaderParams(v string) {
	o.IdtHeaderParams = &v
}

// GetClaimsForTx returns the ClaimsForTx field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetClaimsForTx() string {
	if o == nil || IsNil(o.ClaimsForTx) {
		var ret string
		return ret
	}
	return *o.ClaimsForTx
}

// GetClaimsForTxOk returns a tuple with the ClaimsForTx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetClaimsForTxOk() (*string, bool) {
	if o == nil || IsNil(o.ClaimsForTx) {
		return nil, false
	}
	return o.ClaimsForTx, true
}

// HasClaimsForTx returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasClaimsForTx() bool {
	if o != nil && !IsNil(o.ClaimsForTx) {
		return true
	}

	return false
}

// SetClaimsForTx gets a reference to the given string and assigns it to the ClaimsForTx field.
func (o *AuthorizationIssueRequest) SetClaimsForTx(v string) {
	o.ClaimsForTx = &v
}

// GetConsentedClaims returns the ConsentedClaims field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetConsentedClaims() []string {
	if o == nil || IsNil(o.ConsentedClaims) {
		var ret []string
		return ret
	}
	return o.ConsentedClaims
}

// GetConsentedClaimsOk returns a tuple with the ConsentedClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetConsentedClaimsOk() ([]string, bool) {
	if o == nil || IsNil(o.ConsentedClaims) {
		return nil, false
	}
	return o.ConsentedClaims, true
}

// HasConsentedClaims returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasConsentedClaims() bool {
	if o != nil && !IsNil(o.ConsentedClaims) {
		return true
	}

	return false
}

// SetConsentedClaims gets a reference to the given []string and assigns it to the ConsentedClaims field.
func (o *AuthorizationIssueRequest) SetConsentedClaims(v []string) {
	o.ConsentedClaims = v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetAuthorizationDetails() AuthzDetails {
	if o == nil || IsNil(o.AuthorizationDetails) {
		var ret AuthzDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetAuthorizationDetailsOk() (*AuthzDetails, bool) {
	if o == nil || IsNil(o.AuthorizationDetails) {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasAuthorizationDetails() bool {
	if o != nil && !IsNil(o.AuthorizationDetails) {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthzDetails and assigns it to the AuthorizationDetails field.
func (o *AuthorizationIssueRequest) SetAuthorizationDetails(v AuthzDetails) {
	o.AuthorizationDetails = &v
}

// GetJwtAtClaims returns the JwtAtClaims field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetJwtAtClaims() string {
	if o == nil || IsNil(o.JwtAtClaims) {
		var ret string
		return ret
	}
	return *o.JwtAtClaims
}

// GetJwtAtClaimsOk returns a tuple with the JwtAtClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetJwtAtClaimsOk() (*string, bool) {
	if o == nil || IsNil(o.JwtAtClaims) {
		return nil, false
	}
	return o.JwtAtClaims, true
}

// HasJwtAtClaims returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasJwtAtClaims() bool {
	if o != nil && !IsNil(o.JwtAtClaims) {
		return true
	}

	return false
}

// SetJwtAtClaims gets a reference to the given string and assigns it to the JwtAtClaims field.
func (o *AuthorizationIssueRequest) SetJwtAtClaims(v string) {
	o.JwtAtClaims = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *AuthorizationIssueRequest) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationIssueRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *AuthorizationIssueRequest) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *AuthorizationIssueRequest) SetAccessToken(v string) {
	o.AccessToken = &v
}

func (o AuthorizationIssueRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationIssueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ticket"] = o.Ticket
	toSerialize["subject"] = o.Subject
	if !IsNil(o.AuthTime) {
		toSerialize["authTime"] = o.AuthTime
	}
	if !IsNil(o.Acr) {
		toSerialize["acr"] = o.Acr
	}
	if !IsNil(o.Claims) {
		toSerialize["claims"] = o.Claims
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Sub) {
		toSerialize["sub"] = o.Sub
	}
	if !IsNil(o.IdtHeaderParams) {
		toSerialize["idtHeaderParams"] = o.IdtHeaderParams
	}
	if !IsNil(o.ClaimsForTx) {
		toSerialize["claimsForTx"] = o.ClaimsForTx
	}
	if !IsNil(o.ConsentedClaims) {
		toSerialize["consentedClaims"] = o.ConsentedClaims
	}
	if !IsNil(o.AuthorizationDetails) {
		toSerialize["authorizationDetails"] = o.AuthorizationDetails
	}
	if !IsNil(o.JwtAtClaims) {
		toSerialize["jwtAtClaims"] = o.JwtAtClaims
	}
	if !IsNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	return toSerialize, nil
}

func (o *AuthorizationIssueRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ticket",
		"subject",
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

	varAuthorizationIssueRequest := _AuthorizationIssueRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthorizationIssueRequest)

	if err != nil {
		return err
	}

	*o = AuthorizationIssueRequest(varAuthorizationIssueRequest)

	return err
}

type NullableAuthorizationIssueRequest struct {
	value *AuthorizationIssueRequest
	isSet bool
}

func (v NullableAuthorizationIssueRequest) Get() *AuthorizationIssueRequest {
	return v.value
}

func (v *NullableAuthorizationIssueRequest) Set(val *AuthorizationIssueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationIssueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationIssueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationIssueRequest(val *AuthorizationIssueRequest) *NullableAuthorizationIssueRequest {
	return &NullableAuthorizationIssueRequest{value: val, isSet: true}
}

func (v NullableAuthorizationIssueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationIssueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
