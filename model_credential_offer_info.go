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

// checks if the CredentialOfferInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialOfferInfo{}

// CredentialOfferInfo struct for CredentialOfferInfo
type CredentialOfferInfo struct {
	// The identifier of the credential offer.
	Identifier *string `json:"identifier,omitempty"`
	// The credential offer in the JSON format.
	CredentialOffer *string `json:"credentialOffer,omitempty"`
	// The identifier of the credential issuer.
	CredentialIssuer *string `json:"credentialIssuer,omitempty"`
	// The value of the `credentials` object in the JSON format.
	Credentials *string `json:"credentials,omitempty"`
	// The flag indicating whether the `authorization_code` object is included in the `grants` object.
	AuthorizationCodeGrantIncluded *bool `json:"authorizationCodeGrantIncluded,omitempty"`
	// The flag indicating whether the `issuer_state` property is included in the `authorization_code` object in the `grants` object.
	IssuerStateIncluded *bool `json:"issuerStateIncluded,omitempty"`
	// The value of the `issuer_state` property in the `authorization_code` object in the `grants` object.
	IssuerState *string `json:"issuerState,omitempty"`
	// The flag indicating whether the `urn:ietf:params:oauth:grant-type:pre-authorized_code` object is included in the `grants` object.
	PreAuthorizedCodeGrantIncluded *bool `json:"preAuthorizedCodeGrantIncluded,omitempty"`
	// The value of the `pre-authorized_code` property in the `urn:ietf:params:oauth:grant-type:pre-authorized_code` object in the `grants` object.
	PreAuthorizedCode *string `json:"preAuthorizedCode,omitempty"`
	// The value of the `user_pin_required` property in the `urn:ietf:params:oauth:grant-type:pre-authorized_code` object in the `grants` object.
	UserPinRequired *bool `json:"userPinRequired,omitempty"`
	// The value of the user PIN associated with the credential offer.
	UserPin *string `json:"userPin,omitempty"`
	// The subject associated with the credential offer.
	Subject *string `json:"subject,omitempty"`
	// The time at which the credential offer will expire.
	ExpiresAt *int64 `json:"expiresAt,omitempty"`
	// The general-purpose arbitrary string.
	Context *string `json:"context,omitempty"`
	// Extra properties to associate with the credential offer.
	Properties []Property `json:"properties,omitempty"`
	// Additional claims that are added to the payload part of the JWT access token.
	JwtAtClaims *string `json:"jwtAtClaims,omitempty"`
	// The time at which the user authentication was performed during the course of issuing the credential offer.
	AuthTime *int64 `json:"authTime,omitempty"`
	// The Authentication Context Class Reference of the user authentication performed during the course of issuing the credential offer.
	Acr *string `json:"acr,omitempty"`
}

// NewCredentialOfferInfo instantiates a new CredentialOfferInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialOfferInfo() *CredentialOfferInfo {
	this := CredentialOfferInfo{}
	return &this
}

// NewCredentialOfferInfoWithDefaults instantiates a new CredentialOfferInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialOfferInfoWithDefaults() *CredentialOfferInfo {
	this := CredentialOfferInfo{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *CredentialOfferInfo) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetCredentialOffer returns the CredentialOffer field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetCredentialOffer() string {
	if o == nil || IsNil(o.CredentialOffer) {
		var ret string
		return ret
	}
	return *o.CredentialOffer
}

// GetCredentialOfferOk returns a tuple with the CredentialOffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetCredentialOfferOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialOffer) {
		return nil, false
	}
	return o.CredentialOffer, true
}

// HasCredentialOffer returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasCredentialOffer() bool {
	if o != nil && !IsNil(o.CredentialOffer) {
		return true
	}

	return false
}

// SetCredentialOffer gets a reference to the given string and assigns it to the CredentialOffer field.
func (o *CredentialOfferInfo) SetCredentialOffer(v string) {
	o.CredentialOffer = &v
}

// GetCredentialIssuer returns the CredentialIssuer field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetCredentialIssuer() string {
	if o == nil || IsNil(o.CredentialIssuer) {
		var ret string
		return ret
	}
	return *o.CredentialIssuer
}

// GetCredentialIssuerOk returns a tuple with the CredentialIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetCredentialIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialIssuer) {
		return nil, false
	}
	return o.CredentialIssuer, true
}

// HasCredentialIssuer returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasCredentialIssuer() bool {
	if o != nil && !IsNil(o.CredentialIssuer) {
		return true
	}

	return false
}

// SetCredentialIssuer gets a reference to the given string and assigns it to the CredentialIssuer field.
func (o *CredentialOfferInfo) SetCredentialIssuer(v string) {
	o.CredentialIssuer = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetCredentials() string {
	if o == nil || IsNil(o.Credentials) {
		var ret string
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetCredentialsOk() (*string, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given string and assigns it to the Credentials field.
func (o *CredentialOfferInfo) SetCredentials(v string) {
	o.Credentials = &v
}

// GetAuthorizationCodeGrantIncluded returns the AuthorizationCodeGrantIncluded field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetAuthorizationCodeGrantIncluded() bool {
	if o == nil || IsNil(o.AuthorizationCodeGrantIncluded) {
		var ret bool
		return ret
	}
	return *o.AuthorizationCodeGrantIncluded
}

// GetAuthorizationCodeGrantIncludedOk returns a tuple with the AuthorizationCodeGrantIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetAuthorizationCodeGrantIncludedOk() (*bool, bool) {
	if o == nil || IsNil(o.AuthorizationCodeGrantIncluded) {
		return nil, false
	}
	return o.AuthorizationCodeGrantIncluded, true
}

// HasAuthorizationCodeGrantIncluded returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasAuthorizationCodeGrantIncluded() bool {
	if o != nil && !IsNil(o.AuthorizationCodeGrantIncluded) {
		return true
	}

	return false
}

// SetAuthorizationCodeGrantIncluded gets a reference to the given bool and assigns it to the AuthorizationCodeGrantIncluded field.
func (o *CredentialOfferInfo) SetAuthorizationCodeGrantIncluded(v bool) {
	o.AuthorizationCodeGrantIncluded = &v
}

// GetIssuerStateIncluded returns the IssuerStateIncluded field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetIssuerStateIncluded() bool {
	if o == nil || IsNil(o.IssuerStateIncluded) {
		var ret bool
		return ret
	}
	return *o.IssuerStateIncluded
}

// GetIssuerStateIncludedOk returns a tuple with the IssuerStateIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetIssuerStateIncludedOk() (*bool, bool) {
	if o == nil || IsNil(o.IssuerStateIncluded) {
		return nil, false
	}
	return o.IssuerStateIncluded, true
}

// HasIssuerStateIncluded returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasIssuerStateIncluded() bool {
	if o != nil && !IsNil(o.IssuerStateIncluded) {
		return true
	}

	return false
}

// SetIssuerStateIncluded gets a reference to the given bool and assigns it to the IssuerStateIncluded field.
func (o *CredentialOfferInfo) SetIssuerStateIncluded(v bool) {
	o.IssuerStateIncluded = &v
}

// GetIssuerState returns the IssuerState field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetIssuerState() string {
	if o == nil || IsNil(o.IssuerState) {
		var ret string
		return ret
	}
	return *o.IssuerState
}

// GetIssuerStateOk returns a tuple with the IssuerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetIssuerStateOk() (*string, bool) {
	if o == nil || IsNil(o.IssuerState) {
		return nil, false
	}
	return o.IssuerState, true
}

// HasIssuerState returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasIssuerState() bool {
	if o != nil && !IsNil(o.IssuerState) {
		return true
	}

	return false
}

// SetIssuerState gets a reference to the given string and assigns it to the IssuerState field.
func (o *CredentialOfferInfo) SetIssuerState(v string) {
	o.IssuerState = &v
}

// GetPreAuthorizedCodeGrantIncluded returns the PreAuthorizedCodeGrantIncluded field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetPreAuthorizedCodeGrantIncluded() bool {
	if o == nil || IsNil(o.PreAuthorizedCodeGrantIncluded) {
		var ret bool
		return ret
	}
	return *o.PreAuthorizedCodeGrantIncluded
}

// GetPreAuthorizedCodeGrantIncludedOk returns a tuple with the PreAuthorizedCodeGrantIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetPreAuthorizedCodeGrantIncludedOk() (*bool, bool) {
	if o == nil || IsNil(o.PreAuthorizedCodeGrantIncluded) {
		return nil, false
	}
	return o.PreAuthorizedCodeGrantIncluded, true
}

// HasPreAuthorizedCodeGrantIncluded returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasPreAuthorizedCodeGrantIncluded() bool {
	if o != nil && !IsNil(o.PreAuthorizedCodeGrantIncluded) {
		return true
	}

	return false
}

// SetPreAuthorizedCodeGrantIncluded gets a reference to the given bool and assigns it to the PreAuthorizedCodeGrantIncluded field.
func (o *CredentialOfferInfo) SetPreAuthorizedCodeGrantIncluded(v bool) {
	o.PreAuthorizedCodeGrantIncluded = &v
}

// GetPreAuthorizedCode returns the PreAuthorizedCode field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetPreAuthorizedCode() string {
	if o == nil || IsNil(o.PreAuthorizedCode) {
		var ret string
		return ret
	}
	return *o.PreAuthorizedCode
}

// GetPreAuthorizedCodeOk returns a tuple with the PreAuthorizedCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetPreAuthorizedCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PreAuthorizedCode) {
		return nil, false
	}
	return o.PreAuthorizedCode, true
}

// HasPreAuthorizedCode returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasPreAuthorizedCode() bool {
	if o != nil && !IsNil(o.PreAuthorizedCode) {
		return true
	}

	return false
}

// SetPreAuthorizedCode gets a reference to the given string and assigns it to the PreAuthorizedCode field.
func (o *CredentialOfferInfo) SetPreAuthorizedCode(v string) {
	o.PreAuthorizedCode = &v
}

// GetUserPinRequired returns the UserPinRequired field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetUserPinRequired() bool {
	if o == nil || IsNil(o.UserPinRequired) {
		var ret bool
		return ret
	}
	return *o.UserPinRequired
}

// GetUserPinRequiredOk returns a tuple with the UserPinRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetUserPinRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.UserPinRequired) {
		return nil, false
	}
	return o.UserPinRequired, true
}

// HasUserPinRequired returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasUserPinRequired() bool {
	if o != nil && !IsNil(o.UserPinRequired) {
		return true
	}

	return false
}

// SetUserPinRequired gets a reference to the given bool and assigns it to the UserPinRequired field.
func (o *CredentialOfferInfo) SetUserPinRequired(v bool) {
	o.UserPinRequired = &v
}

// GetUserPin returns the UserPin field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetUserPin() string {
	if o == nil || IsNil(o.UserPin) {
		var ret string
		return ret
	}
	return *o.UserPin
}

// GetUserPinOk returns a tuple with the UserPin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetUserPinOk() (*string, bool) {
	if o == nil || IsNil(o.UserPin) {
		return nil, false
	}
	return o.UserPin, true
}

// HasUserPin returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasUserPin() bool {
	if o != nil && !IsNil(o.UserPin) {
		return true
	}

	return false
}

// SetUserPin gets a reference to the given string and assigns it to the UserPin field.
func (o *CredentialOfferInfo) SetUserPin(v string) {
	o.UserPin = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *CredentialOfferInfo) SetSubject(v string) {
	o.Subject = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *CredentialOfferInfo) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *CredentialOfferInfo) SetContext(v string) {
	o.Context = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetProperties() []Property {
	if o == nil || IsNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetPropertiesOk() ([]Property, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *CredentialOfferInfo) SetProperties(v []Property) {
	o.Properties = v
}

// GetJwtAtClaims returns the JwtAtClaims field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetJwtAtClaims() string {
	if o == nil || IsNil(o.JwtAtClaims) {
		var ret string
		return ret
	}
	return *o.JwtAtClaims
}

// GetJwtAtClaimsOk returns a tuple with the JwtAtClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetJwtAtClaimsOk() (*string, bool) {
	if o == nil || IsNil(o.JwtAtClaims) {
		return nil, false
	}
	return o.JwtAtClaims, true
}

// HasJwtAtClaims returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasJwtAtClaims() bool {
	if o != nil && !IsNil(o.JwtAtClaims) {
		return true
	}

	return false
}

// SetJwtAtClaims gets a reference to the given string and assigns it to the JwtAtClaims field.
func (o *CredentialOfferInfo) SetJwtAtClaims(v string) {
	o.JwtAtClaims = &v
}

// GetAuthTime returns the AuthTime field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetAuthTime() int64 {
	if o == nil || IsNil(o.AuthTime) {
		var ret int64
		return ret
	}
	return *o.AuthTime
}

// GetAuthTimeOk returns a tuple with the AuthTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetAuthTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.AuthTime) {
		return nil, false
	}
	return o.AuthTime, true
}

// HasAuthTime returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasAuthTime() bool {
	if o != nil && !IsNil(o.AuthTime) {
		return true
	}

	return false
}

// SetAuthTime gets a reference to the given int64 and assigns it to the AuthTime field.
func (o *CredentialOfferInfo) SetAuthTime(v int64) {
	o.AuthTime = &v
}

// GetAcr returns the Acr field value if set, zero value otherwise.
func (o *CredentialOfferInfo) GetAcr() string {
	if o == nil || IsNil(o.Acr) {
		var ret string
		return ret
	}
	return *o.Acr
}

// GetAcrOk returns a tuple with the Acr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialOfferInfo) GetAcrOk() (*string, bool) {
	if o == nil || IsNil(o.Acr) {
		return nil, false
	}
	return o.Acr, true
}

// HasAcr returns a boolean if a field has been set.
func (o *CredentialOfferInfo) HasAcr() bool {
	if o != nil && !IsNil(o.Acr) {
		return true
	}

	return false
}

// SetAcr gets a reference to the given string and assigns it to the Acr field.
func (o *CredentialOfferInfo) SetAcr(v string) {
	o.Acr = &v
}

func (o CredentialOfferInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialOfferInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.CredentialOffer) {
		toSerialize["credentialOffer"] = o.CredentialOffer
	}
	if !IsNil(o.CredentialIssuer) {
		toSerialize["credentialIssuer"] = o.CredentialIssuer
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.AuthorizationCodeGrantIncluded) {
		toSerialize["authorizationCodeGrantIncluded"] = o.AuthorizationCodeGrantIncluded
	}
	if !IsNil(o.IssuerStateIncluded) {
		toSerialize["issuerStateIncluded"] = o.IssuerStateIncluded
	}
	if !IsNil(o.IssuerState) {
		toSerialize["issuerState"] = o.IssuerState
	}
	if !IsNil(o.PreAuthorizedCodeGrantIncluded) {
		toSerialize["preAuthorizedCodeGrantIncluded"] = o.PreAuthorizedCodeGrantIncluded
	}
	if !IsNil(o.PreAuthorizedCode) {
		toSerialize["preAuthorizedCode"] = o.PreAuthorizedCode
	}
	if !IsNil(o.UserPinRequired) {
		toSerialize["userPinRequired"] = o.UserPinRequired
	}
	if !IsNil(o.UserPin) {
		toSerialize["userPin"] = o.UserPin
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.JwtAtClaims) {
		toSerialize["jwtAtClaims"] = o.JwtAtClaims
	}
	if !IsNil(o.AuthTime) {
		toSerialize["authTime"] = o.AuthTime
	}
	if !IsNil(o.Acr) {
		toSerialize["acr"] = o.Acr
	}
	return toSerialize, nil
}

type NullableCredentialOfferInfo struct {
	value *CredentialOfferInfo
	isSet bool
}

func (v NullableCredentialOfferInfo) Get() *CredentialOfferInfo {
	return v.value
}

func (v *NullableCredentialOfferInfo) Set(val *CredentialOfferInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialOfferInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialOfferInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialOfferInfo(val *CredentialOfferInfo) *NullableCredentialOfferInfo {
	return &NullableCredentialOfferInfo{value: val, isSet: true}
}

func (v NullableCredentialOfferInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialOfferInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
