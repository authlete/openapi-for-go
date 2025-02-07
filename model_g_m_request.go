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

// checks if the GMRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GMRequest{}

// GMRequest struct for GMRequest
type GMRequest struct {
	// An access token to introspect.
	AccessToken *string `json:"accessToken,omitempty"`
	// A string array listing names of scopes which the caller (= a protected resource endpoint of the service) requires. When the content type of the request from the service is `application/x-www-form-urlencoded`, the format of `scopes` is a space-separated list of scope names.  If this parameter is a non-empty array and if it contains a scope which is not covered by the access token,`action=FORBIDDEN` with `error=insufficient_scope` is returned from Authlete.
	Scopes []string `json:"scopes,omitempty"`
	// A subject (= a user account managed by the service) whom the caller (= a protected resource endpoint of the service) requires.  If this parameter is not `null` and if the value does not match the subject who is associated with the access token, `action=FORBIDDEN` with `error=invalid_request` is returned from Authlete.
	Subject *string `json:"subject,omitempty"`
	// Client certificate in PEM format, used to validate binding against access tokens using the TLS client certificate confirmation method.
	ClientCertificate *string `json:"clientCertificate,omitempty"`
	// `DPoP` header presented by the client during the request to the resource server.  The header contains a signed JWT which includes the public key that is paired with the private key used to sign the JWT. See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details.
	Dpop *string `json:"dpop,omitempty"`
	// HTTP method of the request from the client to the protected resource endpoint. This field is used to validate the `DPoP` header.  See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details.
	Htm *string `json:"htm,omitempty"`
	// URL of the protected resource endpoint. This field is used to validate the `DPoP` header.  See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details.
	Htu *string `json:"htu,omitempty"`
	// The resources specified by the `resource` request parameters in the token request. See \"Resource Indicators for OAuth 2.0\" for details.
	Resources []string               `json:"resources,omitempty"`
	GmAction  *GrantManagementAction `json:"gmAction,omitempty"`
	// The value of the `grant_id` request parameter of the device authorization request.  The `grant_id` request parameter is defined in [Grant Management for OAuth 2.0](https://openid.net/specs/fapi-grant-management.html) , which is supported by Authlete 2.3 and newer versions.
	GrantId *string `json:"grantId,omitempty"`
}

// NewGMRequest instantiates a new GMRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGMRequest() *GMRequest {
	this := GMRequest{}
	return &this
}

// NewGMRequestWithDefaults instantiates a new GMRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGMRequestWithDefaults() *GMRequest {
	this := GMRequest{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *GMRequest) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *GMRequest) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *GMRequest) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *GMRequest) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMRequest) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *GMRequest) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *GMRequest) SetScopes(v []string) {
	o.Scopes = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *GMRequest) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMRequest) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *GMRequest) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *GMRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetClientCertificate returns the ClientCertificate field value if set, zero value otherwise.
func (o *GMRequest) GetClientCertificate() string {
	if o == nil || IsNil(o.ClientCertificate) {
		var ret string
		return ret
	}
	return *o.ClientCertificate
}

// GetClientCertificateOk returns a tuple with the ClientCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMRequest) GetClientCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.ClientCertificate) {
		return nil, false
	}
	return o.ClientCertificate, true
}

// HasClientCertificate returns a boolean if a field has been set.
func (o *GMRequest) HasClientCertificate() bool {
	if o != nil && !IsNil(o.ClientCertificate) {
		return true
	}

	return false
}

// SetClientCertificate gets a reference to the given string and assigns it to the ClientCertificate field.
func (o *GMRequest) SetClientCertificate(v string) {
	o.ClientCertificate = &v
}

// GetDpop returns the Dpop field value if set, zero value otherwise.
func (o *GMRequest) GetDpop() string {
	if o == nil || IsNil(o.Dpop) {
		var ret string
		return ret
	}
	return *o.Dpop
}

// GetDpopOk returns a tuple with the Dpop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMRequest) GetDpopOk() (*string, bool) {
	if o == nil || IsNil(o.Dpop) {
		return nil, false
	}
	return o.Dpop, true
}

// HasDpop returns a boolean if a field has been set.
func (o *GMRequest) HasDpop() bool {
	if o != nil && !IsNil(o.Dpop) {
		return true
	}

	return false
}

// SetDpop gets a reference to the given string and assigns it to the Dpop field.
func (o *GMRequest) SetDpop(v string) {
	o.Dpop = &v
}

// GetHtm returns the Htm field value if set, zero value otherwise.
func (o *GMRequest) GetHtm() string {
	if o == nil || IsNil(o.Htm) {
		var ret string
		return ret
	}
	return *o.Htm
}

// GetHtmOk returns a tuple with the Htm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMRequest) GetHtmOk() (*string, bool) {
	if o == nil || IsNil(o.Htm) {
		return nil, false
	}
	return o.Htm, true
}

// HasHtm returns a boolean if a field has been set.
func (o *GMRequest) HasHtm() bool {
	if o != nil && !IsNil(o.Htm) {
		return true
	}

	return false
}

// SetHtm gets a reference to the given string and assigns it to the Htm field.
func (o *GMRequest) SetHtm(v string) {
	o.Htm = &v
}

// GetHtu returns the Htu field value if set, zero value otherwise.
func (o *GMRequest) GetHtu() string {
	if o == nil || IsNil(o.Htu) {
		var ret string
		return ret
	}
	return *o.Htu
}

// GetHtuOk returns a tuple with the Htu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMRequest) GetHtuOk() (*string, bool) {
	if o == nil || IsNil(o.Htu) {
		return nil, false
	}
	return o.Htu, true
}

// HasHtu returns a boolean if a field has been set.
func (o *GMRequest) HasHtu() bool {
	if o != nil && !IsNil(o.Htu) {
		return true
	}

	return false
}

// SetHtu gets a reference to the given string and assigns it to the Htu field.
func (o *GMRequest) SetHtu(v string) {
	o.Htu = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *GMRequest) GetResources() []string {
	if o == nil || IsNil(o.Resources) {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMRequest) GetResourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *GMRequest) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *GMRequest) SetResources(v []string) {
	o.Resources = v
}

// GetGmAction returns the GmAction field value if set, zero value otherwise.
func (o *GMRequest) GetGmAction() GrantManagementAction {
	if o == nil || IsNil(o.GmAction) {
		var ret GrantManagementAction
		return ret
	}
	return *o.GmAction
}

// GetGmActionOk returns a tuple with the GmAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMRequest) GetGmActionOk() (*GrantManagementAction, bool) {
	if o == nil || IsNil(o.GmAction) {
		return nil, false
	}
	return o.GmAction, true
}

// HasGmAction returns a boolean if a field has been set.
func (o *GMRequest) HasGmAction() bool {
	if o != nil && !IsNil(o.GmAction) {
		return true
	}

	return false
}

// SetGmAction gets a reference to the given GrantManagementAction and assigns it to the GmAction field.
func (o *GMRequest) SetGmAction(v GrantManagementAction) {
	o.GmAction = &v
}

// GetGrantId returns the GrantId field value if set, zero value otherwise.
func (o *GMRequest) GetGrantId() string {
	if o == nil || IsNil(o.GrantId) {
		var ret string
		return ret
	}
	return *o.GrantId
}

// GetGrantIdOk returns a tuple with the GrantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GMRequest) GetGrantIdOk() (*string, bool) {
	if o == nil || IsNil(o.GrantId) {
		return nil, false
	}
	return o.GrantId, true
}

// HasGrantId returns a boolean if a field has been set.
func (o *GMRequest) HasGrantId() bool {
	if o != nil && !IsNil(o.GrantId) {
		return true
	}

	return false
}

// SetGrantId gets a reference to the given string and assigns it to the GrantId field.
func (o *GMRequest) SetGrantId(v string) {
	o.GrantId = &v
}

func (o GMRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GMRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.ClientCertificate) {
		toSerialize["clientCertificate"] = o.ClientCertificate
	}
	if !IsNil(o.Dpop) {
		toSerialize["dpop"] = o.Dpop
	}
	if !IsNil(o.Htm) {
		toSerialize["htm"] = o.Htm
	}
	if !IsNil(o.Htu) {
		toSerialize["htu"] = o.Htu
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.GmAction) {
		toSerialize["gmAction"] = o.GmAction
	}
	if !IsNil(o.GrantId) {
		toSerialize["grantId"] = o.GrantId
	}
	return toSerialize, nil
}

type NullableGMRequest struct {
	value *GMRequest
	isSet bool
}

func (v NullableGMRequest) Get() *GMRequest {
	return v.value
}

func (v *NullableGMRequest) Set(val *GMRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGMRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGMRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGMRequest(val *GMRequest) *NullableGMRequest {
	return &NullableGMRequest{value: val, isSet: true}
}

func (v NullableGMRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGMRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
