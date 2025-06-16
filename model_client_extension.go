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

// checks if the ClientExtension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientExtension{}

// ClientExtension struct for ClientExtension
type ClientExtension struct {
	// The set of scopes that the client application is allowed to request. This paramter will be one of the following.    - `null`   - an empty set   - a set with at least one element  When the value of this parameter is `null`, it means that the set of scopes that the client application is allowed to request is the set of the scopes that the service supports. When the value of this parameter is an empty set, it means that the client application is not allowed to request any scopes. When the value of this parameter is a set with at least one element, it means that the set is the set of scopes that the client application is allowed to request.
	RequestableScopes []string `json:"requestableScopes,omitempty"`
	// The flag to indicate whether \"Requestable Scopes per Client\" is enabled or not. If `true`, you can define the set of scopes which this client application can request. If `false`, this client application can request any scope which is supported by the authorization server.
	RequestableScopesEnabled *bool `json:"requestableScopesEnabled,omitempty"`
	// The value of the duration of access tokens per client in seconds. In normal cases, the value of the service's `accessTokenDuration` property is used as the duration of access tokens issued by the service. However, if this `accessTokenDuration` property holds a non-zero positive number and its value is less than the duration configured by the service, the value is used as the duration of access tokens issued to the client application.  Note that the duration of access tokens can be controlled by the scope attribute `access_token.duration`, too. Authlete chooses the minimum value among the candidates.
	AccessTokenDuration *int64 `json:"accessTokenDuration,omitempty"`
	// The value of the duration of refresh tokens per client in seconds. In normal cases, the value of the service's `refreshTokenDuration` property is used as the duration of refresh tokens issued by the service. However, if this `refreshTokenDuration` property holds a non-zero positive number and its value is less than the duration configured by the service, the value is used as the duration of refresh tokens issued to the client application.  Note that the duration of refresh tokens can be controlled by the scope attribute `refresh_token.duration`, too. Authlete chooses the minimum value among the candidates.
	RefreshTokenDuration *int64 `json:"refreshTokenDuration,omitempty"`
	// The value of the duration of ID tokens per client in seconds. In normal cases, the value of the service's `idTokenDuration` property is used as the duration of ID tokens issued by the service. However, if this `idTokenDuration` property holds a non-zero positive number and its value is less than the duration configured by the service, the value is used as the duration of ID tokens issued to the client application.  Note that the duration of refresh tokens can be controlled by the scope attribute `id_token.duration`, too. Authlete chooses the minimum value among the candidates.
	IdTokenDuration *int64 `json:"idTokenDuration,omitempty"`
	// Get the flag indicating whether the client is explicitly given a permission to make token exchange requests ([RFC 8693][https://www.rfc-editor.org/rfc/rfc8693.html])
	TokenExchangePermitted *bool `json:"tokenExchangePermitted,omitempty"`
}

// NewClientExtension instantiates a new ClientExtension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientExtension() *ClientExtension {
	this := ClientExtension{}
	return &this
}

// NewClientExtensionWithDefaults instantiates a new ClientExtension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientExtensionWithDefaults() *ClientExtension {
	this := ClientExtension{}
	return &this
}

// GetRequestableScopes returns the RequestableScopes field value if set, zero value otherwise.
func (o *ClientExtension) GetRequestableScopes() []string {
	if o == nil || IsNil(o.RequestableScopes) {
		var ret []string
		return ret
	}
	return o.RequestableScopes
}

// GetRequestableScopesOk returns a tuple with the RequestableScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientExtension) GetRequestableScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.RequestableScopes) {
		return nil, false
	}
	return o.RequestableScopes, true
}

// HasRequestableScopes returns a boolean if a field has been set.
func (o *ClientExtension) HasRequestableScopes() bool {
	if o != nil && !IsNil(o.RequestableScopes) {
		return true
	}

	return false
}

// SetRequestableScopes gets a reference to the given []string and assigns it to the RequestableScopes field.
func (o *ClientExtension) SetRequestableScopes(v []string) {
	o.RequestableScopes = v
}

// GetRequestableScopesEnabled returns the RequestableScopesEnabled field value if set, zero value otherwise.
func (o *ClientExtension) GetRequestableScopesEnabled() bool {
	if o == nil || IsNil(o.RequestableScopesEnabled) {
		var ret bool
		return ret
	}
	return *o.RequestableScopesEnabled
}

// GetRequestableScopesEnabledOk returns a tuple with the RequestableScopesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientExtension) GetRequestableScopesEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestableScopesEnabled) {
		return nil, false
	}
	return o.RequestableScopesEnabled, true
}

// HasRequestableScopesEnabled returns a boolean if a field has been set.
func (o *ClientExtension) HasRequestableScopesEnabled() bool {
	if o != nil && !IsNil(o.RequestableScopesEnabled) {
		return true
	}

	return false
}

// SetRequestableScopesEnabled gets a reference to the given bool and assigns it to the RequestableScopesEnabled field.
func (o *ClientExtension) SetRequestableScopesEnabled(v bool) {
	o.RequestableScopesEnabled = &v
}

// GetAccessTokenDuration returns the AccessTokenDuration field value if set, zero value otherwise.
func (o *ClientExtension) GetAccessTokenDuration() int64 {
	if o == nil || IsNil(o.AccessTokenDuration) {
		var ret int64
		return ret
	}
	return *o.AccessTokenDuration
}

// GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientExtension) GetAccessTokenDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.AccessTokenDuration) {
		return nil, false
	}
	return o.AccessTokenDuration, true
}

// HasAccessTokenDuration returns a boolean if a field has been set.
func (o *ClientExtension) HasAccessTokenDuration() bool {
	if o != nil && !IsNil(o.AccessTokenDuration) {
		return true
	}

	return false
}

// SetAccessTokenDuration gets a reference to the given int64 and assigns it to the AccessTokenDuration field.
func (o *ClientExtension) SetAccessTokenDuration(v int64) {
	o.AccessTokenDuration = &v
}

// GetRefreshTokenDuration returns the RefreshTokenDuration field value if set, zero value otherwise.
func (o *ClientExtension) GetRefreshTokenDuration() int64 {
	if o == nil || IsNil(o.RefreshTokenDuration) {
		var ret int64
		return ret
	}
	return *o.RefreshTokenDuration
}

// GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientExtension) GetRefreshTokenDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.RefreshTokenDuration) {
		return nil, false
	}
	return o.RefreshTokenDuration, true
}

// HasRefreshTokenDuration returns a boolean if a field has been set.
func (o *ClientExtension) HasRefreshTokenDuration() bool {
	if o != nil && !IsNil(o.RefreshTokenDuration) {
		return true
	}

	return false
}

// SetRefreshTokenDuration gets a reference to the given int64 and assigns it to the RefreshTokenDuration field.
func (o *ClientExtension) SetRefreshTokenDuration(v int64) {
	o.RefreshTokenDuration = &v
}

// GetIdTokenDuration returns the IdTokenDuration field value if set, zero value otherwise.
func (o *ClientExtension) GetIdTokenDuration() int64 {
	if o == nil || IsNil(o.IdTokenDuration) {
		var ret int64
		return ret
	}
	return *o.IdTokenDuration
}

// GetIdTokenDurationOk returns a tuple with the IdTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientExtension) GetIdTokenDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.IdTokenDuration) {
		return nil, false
	}
	return o.IdTokenDuration, true
}

// HasIdTokenDuration returns a boolean if a field has been set.
func (o *ClientExtension) HasIdTokenDuration() bool {
	if o != nil && !IsNil(o.IdTokenDuration) {
		return true
	}

	return false
}

// SetIdTokenDuration gets a reference to the given int64 and assigns it to the IdTokenDuration field.
func (o *ClientExtension) SetIdTokenDuration(v int64) {
	o.IdTokenDuration = &v
}

// GetTokenExchangePermitted returns the TokenExchangePermitted field value if set, zero value otherwise.
func (o *ClientExtension) GetTokenExchangePermitted() bool {
	if o == nil || IsNil(o.TokenExchangePermitted) {
		var ret bool
		return ret
	}
	return *o.TokenExchangePermitted
}

// GetTokenExchangePermittedOk returns a tuple with the TokenExchangePermitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientExtension) GetTokenExchangePermittedOk() (*bool, bool) {
	if o == nil || IsNil(o.TokenExchangePermitted) {
		return nil, false
	}
	return o.TokenExchangePermitted, true
}

// HasTokenExchangePermitted returns a boolean if a field has been set.
func (o *ClientExtension) HasTokenExchangePermitted() bool {
	if o != nil && !IsNil(o.TokenExchangePermitted) {
		return true
	}

	return false
}

// SetTokenExchangePermitted gets a reference to the given bool and assigns it to the TokenExchangePermitted field.
func (o *ClientExtension) SetTokenExchangePermitted(v bool) {
	o.TokenExchangePermitted = &v
}

func (o ClientExtension) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientExtension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestableScopes) {
		toSerialize["requestableScopes"] = o.RequestableScopes
	}
	if !IsNil(o.RequestableScopesEnabled) {
		toSerialize["requestableScopesEnabled"] = o.RequestableScopesEnabled
	}
	if !IsNil(o.AccessTokenDuration) {
		toSerialize["accessTokenDuration"] = o.AccessTokenDuration
	}
	if !IsNil(o.RefreshTokenDuration) {
		toSerialize["refreshTokenDuration"] = o.RefreshTokenDuration
	}
	if !IsNil(o.IdTokenDuration) {
		toSerialize["idTokenDuration"] = o.IdTokenDuration
	}
	if !IsNil(o.TokenExchangePermitted) {
		toSerialize["tokenExchangePermitted"] = o.TokenExchangePermitted
	}
	return toSerialize, nil
}

type NullableClientExtension struct {
	value *ClientExtension
	isSet bool
}

func (v NullableClientExtension) Get() *ClientExtension {
	return v.value
}

func (v *NullableClientExtension) Set(val *ClientExtension) {
	v.value = val
	v.isSet = true
}

func (v NullableClientExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableClientExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientExtension(val *ClientExtension) *NullableClientExtension {
	return &NullableClientExtension{value: val, isSet: true}
}

func (v NullableClientExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
