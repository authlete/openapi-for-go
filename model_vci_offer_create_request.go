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

// checks if the VciOfferCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VciOfferCreateRequest{}

// VciOfferCreateRequest struct for VciOfferCreateRequest
type VciOfferCreateRequest struct {
	// The value of the `credentials` object in the JSON format.
	Credentials *string `json:"credentials,omitempty"`
	// The flag indicating whether the `authorization_code` object is included in the `grants` object.
	AuthorizationCodeGrantIncluded *bool `json:"authorizationCodeGrantIncluded,omitempty"`
	// The flag indicating whether the `issuer_state` property is included in the `authorization_code` object in the `grants` object.
	IssuerStateIncluded *bool `json:"issuerStateIncluded,omitempty"`
	// The flag to include the `urn:ietf:params:oauth:grant-type:pre-authorized_code` object in the `grants` object.
	PreAuthorizedCodeGrantIncluded *bool `json:"preAuthorizedCodeGrantIncluded,omitempty"`
	// The value of the `user_pin_required` property in the `urn:ietf:params:oauth:grant-type:pre-authorized_code` object in the `grants` object.
	UserPinRequired *bool `json:"userPinRequired,omitempty"`
	// The length of the user PIN to generate.
	UserPinLength *int32 `json:"userPinLength,omitempty"`
	// The subject associated with the credential offer.
	Subject *string `json:"subject,omitempty"`
	// The duration of the credential offer.
	Duration *int64 `json:"duration,omitempty"`
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

// NewVciOfferCreateRequest instantiates a new VciOfferCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVciOfferCreateRequest() *VciOfferCreateRequest {
	this := VciOfferCreateRequest{}
	return &this
}

// NewVciOfferCreateRequestWithDefaults instantiates a new VciOfferCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVciOfferCreateRequestWithDefaults() *VciOfferCreateRequest {
	this := VciOfferCreateRequest{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *VciOfferCreateRequest) GetCredentials() string {
	if o == nil || IsNil(o.Credentials) {
		var ret string
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciOfferCreateRequest) GetCredentialsOk() (*string, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *VciOfferCreateRequest) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given string and assigns it to the Credentials field.
func (o *VciOfferCreateRequest) SetCredentials(v string) {
	o.Credentials = &v
}

// GetAuthorizationCodeGrantIncluded returns the AuthorizationCodeGrantIncluded field value if set, zero value otherwise.
func (o *VciOfferCreateRequest) GetAuthorizationCodeGrantIncluded() bool {
	if o == nil || IsNil(o.AuthorizationCodeGrantIncluded) {
		var ret bool
		return ret
	}
	return *o.AuthorizationCodeGrantIncluded
}

// GetAuthorizationCodeGrantIncludedOk returns a tuple with the AuthorizationCodeGrantIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciOfferCreateRequest) GetAuthorizationCodeGrantIncludedOk() (*bool, bool) {
	if o == nil || IsNil(o.AuthorizationCodeGrantIncluded) {
		return nil, false
	}
	return o.AuthorizationCodeGrantIncluded, true
}

// HasAuthorizationCodeGrantIncluded returns a boolean if a field has been set.
func (o *VciOfferCreateRequest) HasAuthorizationCodeGrantIncluded() bool {
	if o != nil && !IsNil(o.AuthorizationCodeGrantIncluded) {
		return true
	}

	return false
}

// SetAuthorizationCodeGrantIncluded gets a reference to the given bool and assigns it to the AuthorizationCodeGrantIncluded field.
func (o *VciOfferCreateRequest) SetAuthorizationCodeGrantIncluded(v bool) {
	o.AuthorizationCodeGrantIncluded = &v
}

// GetIssuerStateIncluded returns the IssuerStateIncluded field value if set, zero value otherwise.
func (o *VciOfferCreateRequest) GetIssuerStateIncluded() bool {
	if o == nil || IsNil(o.IssuerStateIncluded) {
		var ret bool
		return ret
	}
	return *o.IssuerStateIncluded
}

// GetIssuerStateIncludedOk returns a tuple with the IssuerStateIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciOfferCreateRequest) GetIssuerStateIncludedOk() (*bool, bool) {
	if o == nil || IsNil(o.IssuerStateIncluded) {
		return nil, false
	}
	return o.IssuerStateIncluded, true
}

// HasIssuerStateIncluded returns a boolean if a field has been set.
func (o *VciOfferCreateRequest) HasIssuerStateIncluded() bool {
	if o != nil && !IsNil(o.IssuerStateIncluded) {
		return true
	}

	return false
}

// SetIssuerStateIncluded gets a reference to the given bool and assigns it to the IssuerStateIncluded field.
func (o *VciOfferCreateRequest) SetIssuerStateIncluded(v bool) {
	o.IssuerStateIncluded = &v
}

// GetPreAuthorizedCodeGrantIncluded returns the PreAuthorizedCodeGrantIncluded field value if set, zero value otherwise.
func (o *VciOfferCreateRequest) GetPreAuthorizedCodeGrantIncluded() bool {
	if o == nil || IsNil(o.PreAuthorizedCodeGrantIncluded) {
		var ret bool
		return ret
	}
	return *o.PreAuthorizedCodeGrantIncluded
}

// GetPreAuthorizedCodeGrantIncludedOk returns a tuple with the PreAuthorizedCodeGrantIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciOfferCreateRequest) GetPreAuthorizedCodeGrantIncludedOk() (*bool, bool) {
	if o == nil || IsNil(o.PreAuthorizedCodeGrantIncluded) {
		return nil, false
	}
	return o.PreAuthorizedCodeGrantIncluded, true
}

// HasPreAuthorizedCodeGrantIncluded returns a boolean if a field has been set.
func (o *VciOfferCreateRequest) HasPreAuthorizedCodeGrantIncluded() bool {
	if o != nil && !IsNil(o.PreAuthorizedCodeGrantIncluded) {
		return true
	}

	return false
}

// SetPreAuthorizedCodeGrantIncluded gets a reference to the given bool and assigns it to the PreAuthorizedCodeGrantIncluded field.
func (o *VciOfferCreateRequest) SetPreAuthorizedCodeGrantIncluded(v bool) {
	o.PreAuthorizedCodeGrantIncluded = &v
}

// GetUserPinRequired returns the UserPinRequired field value if set, zero value otherwise.
func (o *VciOfferCreateRequest) GetUserPinRequired() bool {
	if o == nil || IsNil(o.UserPinRequired) {
		var ret bool
		return ret
	}
	return *o.UserPinRequired
}

// GetUserPinRequiredOk returns a tuple with the UserPinRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciOfferCreateRequest) GetUserPinRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.UserPinRequired) {
		return nil, false
	}
	return o.UserPinRequired, true
}

// HasUserPinRequired returns a boolean if a field has been set.
func (o *VciOfferCreateRequest) HasUserPinRequired() bool {
	if o != nil && !IsNil(o.UserPinRequired) {
		return true
	}

	return false
}

// SetUserPinRequired gets a reference to the given bool and assigns it to the UserPinRequired field.
func (o *VciOfferCreateRequest) SetUserPinRequired(v bool) {
	o.UserPinRequired = &v
}

// GetUserPinLength returns the UserPinLength field value if set, zero value otherwise.
func (o *VciOfferCreateRequest) GetUserPinLength() int32 {
	if o == nil || IsNil(o.UserPinLength) {
		var ret int32
		return ret
	}
	return *o.UserPinLength
}

// GetUserPinLengthOk returns a tuple with the UserPinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciOfferCreateRequest) GetUserPinLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.UserPinLength) {
		return nil, false
	}
	return o.UserPinLength, true
}

// HasUserPinLength returns a boolean if a field has been set.
func (o *VciOfferCreateRequest) HasUserPinLength() bool {
	if o != nil && !IsNil(o.UserPinLength) {
		return true
	}

	return false
}

// SetUserPinLength gets a reference to the given int32 and assigns it to the UserPinLength field.
func (o *VciOfferCreateRequest) SetUserPinLength(v int32) {
	o.UserPinLength = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *VciOfferCreateRequest) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciOfferCreateRequest) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *VciOfferCreateRequest) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *VciOfferCreateRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *VciOfferCreateRequest) GetDuration() int64 {
	if o == nil || IsNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciOfferCreateRequest) GetDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *VciOfferCreateRequest) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *VciOfferCreateRequest) SetDuration(v int64) {
	o.Duration = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *VciOfferCreateRequest) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciOfferCreateRequest) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *VciOfferCreateRequest) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *VciOfferCreateRequest) SetContext(v string) {
	o.Context = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *VciOfferCreateRequest) GetProperties() []Property {
	if o == nil || IsNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciOfferCreateRequest) GetPropertiesOk() ([]Property, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *VciOfferCreateRequest) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *VciOfferCreateRequest) SetProperties(v []Property) {
	o.Properties = v
}

// GetJwtAtClaims returns the JwtAtClaims field value if set, zero value otherwise.
func (o *VciOfferCreateRequest) GetJwtAtClaims() string {
	if o == nil || IsNil(o.JwtAtClaims) {
		var ret string
		return ret
	}
	return *o.JwtAtClaims
}

// GetJwtAtClaimsOk returns a tuple with the JwtAtClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciOfferCreateRequest) GetJwtAtClaimsOk() (*string, bool) {
	if o == nil || IsNil(o.JwtAtClaims) {
		return nil, false
	}
	return o.JwtAtClaims, true
}

// HasJwtAtClaims returns a boolean if a field has been set.
func (o *VciOfferCreateRequest) HasJwtAtClaims() bool {
	if o != nil && !IsNil(o.JwtAtClaims) {
		return true
	}

	return false
}

// SetJwtAtClaims gets a reference to the given string and assigns it to the JwtAtClaims field.
func (o *VciOfferCreateRequest) SetJwtAtClaims(v string) {
	o.JwtAtClaims = &v
}

// GetAuthTime returns the AuthTime field value if set, zero value otherwise.
func (o *VciOfferCreateRequest) GetAuthTime() int64 {
	if o == nil || IsNil(o.AuthTime) {
		var ret int64
		return ret
	}
	return *o.AuthTime
}

// GetAuthTimeOk returns a tuple with the AuthTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciOfferCreateRequest) GetAuthTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.AuthTime) {
		return nil, false
	}
	return o.AuthTime, true
}

// HasAuthTime returns a boolean if a field has been set.
func (o *VciOfferCreateRequest) HasAuthTime() bool {
	if o != nil && !IsNil(o.AuthTime) {
		return true
	}

	return false
}

// SetAuthTime gets a reference to the given int64 and assigns it to the AuthTime field.
func (o *VciOfferCreateRequest) SetAuthTime(v int64) {
	o.AuthTime = &v
}

// GetAcr returns the Acr field value if set, zero value otherwise.
func (o *VciOfferCreateRequest) GetAcr() string {
	if o == nil || IsNil(o.Acr) {
		var ret string
		return ret
	}
	return *o.Acr
}

// GetAcrOk returns a tuple with the Acr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciOfferCreateRequest) GetAcrOk() (*string, bool) {
	if o == nil || IsNil(o.Acr) {
		return nil, false
	}
	return o.Acr, true
}

// HasAcr returns a boolean if a field has been set.
func (o *VciOfferCreateRequest) HasAcr() bool {
	if o != nil && !IsNil(o.Acr) {
		return true
	}

	return false
}

// SetAcr gets a reference to the given string and assigns it to the Acr field.
func (o *VciOfferCreateRequest) SetAcr(v string) {
	o.Acr = &v
}

func (o VciOfferCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VciOfferCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.AuthorizationCodeGrantIncluded) {
		toSerialize["authorizationCodeGrantIncluded"] = o.AuthorizationCodeGrantIncluded
	}
	if !IsNil(o.IssuerStateIncluded) {
		toSerialize["issuerStateIncluded"] = o.IssuerStateIncluded
	}
	if !IsNil(o.PreAuthorizedCodeGrantIncluded) {
		toSerialize["preAuthorizedCodeGrantIncluded"] = o.PreAuthorizedCodeGrantIncluded
	}
	if !IsNil(o.UserPinRequired) {
		toSerialize["userPinRequired"] = o.UserPinRequired
	}
	if !IsNil(o.UserPinLength) {
		toSerialize["userPinLength"] = o.UserPinLength
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
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

type NullableVciOfferCreateRequest struct {
	value *VciOfferCreateRequest
	isSet bool
}

func (v NullableVciOfferCreateRequest) Get() *VciOfferCreateRequest {
	return v.value
}

func (v *NullableVciOfferCreateRequest) Set(val *VciOfferCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVciOfferCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVciOfferCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVciOfferCreateRequest(val *VciOfferCreateRequest) *NullableVciOfferCreateRequest {
	return &NullableVciOfferCreateRequest{value: val, isSet: true}
}

func (v NullableVciOfferCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVciOfferCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
