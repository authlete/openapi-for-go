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

// checks if the CredentialIssuanceOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialIssuanceOrder{}

// CredentialIssuanceOrder struct for CredentialIssuanceOrder
type CredentialIssuanceOrder struct {
	// The identifier of a credential request.
	RequestIdentifier *string `json:"requestIdentifier,omitempty"`
	// The additional payload that will be added into a credential to be issued.
	CredentialPayload *string `json:"credentialPayload,omitempty"`
	// The flag indicating whether to defer credential issuance.
	IssuanceDeferred *bool `json:"issuanceDeferred,omitempty"`
	// The duration of a credential to be issued.
	CredentialDuration *int64 `json:"credentialDuration,omitempty"`
	// The key ID of a private key that should be used for signing a credential to be issued.
	SigningKeyId *string `json:"signingKeyId,omitempty"`
}

// NewCredentialIssuanceOrder instantiates a new CredentialIssuanceOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialIssuanceOrder() *CredentialIssuanceOrder {
	this := CredentialIssuanceOrder{}
	return &this
}

// NewCredentialIssuanceOrderWithDefaults instantiates a new CredentialIssuanceOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialIssuanceOrderWithDefaults() *CredentialIssuanceOrder {
	this := CredentialIssuanceOrder{}
	return &this
}

// GetRequestIdentifier returns the RequestIdentifier field value if set, zero value otherwise.
func (o *CredentialIssuanceOrder) GetRequestIdentifier() string {
	if o == nil || IsNil(o.RequestIdentifier) {
		var ret string
		return ret
	}
	return *o.RequestIdentifier
}

// GetRequestIdentifierOk returns a tuple with the RequestIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceOrder) GetRequestIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.RequestIdentifier) {
		return nil, false
	}
	return o.RequestIdentifier, true
}

// HasRequestIdentifier returns a boolean if a field has been set.
func (o *CredentialIssuanceOrder) HasRequestIdentifier() bool {
	if o != nil && !IsNil(o.RequestIdentifier) {
		return true
	}

	return false
}

// SetRequestIdentifier gets a reference to the given string and assigns it to the RequestIdentifier field.
func (o *CredentialIssuanceOrder) SetRequestIdentifier(v string) {
	o.RequestIdentifier = &v
}

// GetCredentialPayload returns the CredentialPayload field value if set, zero value otherwise.
func (o *CredentialIssuanceOrder) GetCredentialPayload() string {
	if o == nil || IsNil(o.CredentialPayload) {
		var ret string
		return ret
	}
	return *o.CredentialPayload
}

// GetCredentialPayloadOk returns a tuple with the CredentialPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceOrder) GetCredentialPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialPayload) {
		return nil, false
	}
	return o.CredentialPayload, true
}

// HasCredentialPayload returns a boolean if a field has been set.
func (o *CredentialIssuanceOrder) HasCredentialPayload() bool {
	if o != nil && !IsNil(o.CredentialPayload) {
		return true
	}

	return false
}

// SetCredentialPayload gets a reference to the given string and assigns it to the CredentialPayload field.
func (o *CredentialIssuanceOrder) SetCredentialPayload(v string) {
	o.CredentialPayload = &v
}

// GetIssuanceDeferred returns the IssuanceDeferred field value if set, zero value otherwise.
func (o *CredentialIssuanceOrder) GetIssuanceDeferred() bool {
	if o == nil || IsNil(o.IssuanceDeferred) {
		var ret bool
		return ret
	}
	return *o.IssuanceDeferred
}

// GetIssuanceDeferredOk returns a tuple with the IssuanceDeferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceOrder) GetIssuanceDeferredOk() (*bool, bool) {
	if o == nil || IsNil(o.IssuanceDeferred) {
		return nil, false
	}
	return o.IssuanceDeferred, true
}

// HasIssuanceDeferred returns a boolean if a field has been set.
func (o *CredentialIssuanceOrder) HasIssuanceDeferred() bool {
	if o != nil && !IsNil(o.IssuanceDeferred) {
		return true
	}

	return false
}

// SetIssuanceDeferred gets a reference to the given bool and assigns it to the IssuanceDeferred field.
func (o *CredentialIssuanceOrder) SetIssuanceDeferred(v bool) {
	o.IssuanceDeferred = &v
}

// GetCredentialDuration returns the CredentialDuration field value if set, zero value otherwise.
func (o *CredentialIssuanceOrder) GetCredentialDuration() int64 {
	if o == nil || IsNil(o.CredentialDuration) {
		var ret int64
		return ret
	}
	return *o.CredentialDuration
}

// GetCredentialDurationOk returns a tuple with the CredentialDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceOrder) GetCredentialDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.CredentialDuration) {
		return nil, false
	}
	return o.CredentialDuration, true
}

// HasCredentialDuration returns a boolean if a field has been set.
func (o *CredentialIssuanceOrder) HasCredentialDuration() bool {
	if o != nil && !IsNil(o.CredentialDuration) {
		return true
	}

	return false
}

// SetCredentialDuration gets a reference to the given int64 and assigns it to the CredentialDuration field.
func (o *CredentialIssuanceOrder) SetCredentialDuration(v int64) {
	o.CredentialDuration = &v
}

// GetSigningKeyId returns the SigningKeyId field value if set, zero value otherwise.
func (o *CredentialIssuanceOrder) GetSigningKeyId() string {
	if o == nil || IsNil(o.SigningKeyId) {
		var ret string
		return ret
	}
	return *o.SigningKeyId
}

// GetSigningKeyIdOk returns a tuple with the SigningKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceOrder) GetSigningKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.SigningKeyId) {
		return nil, false
	}
	return o.SigningKeyId, true
}

// HasSigningKeyId returns a boolean if a field has been set.
func (o *CredentialIssuanceOrder) HasSigningKeyId() bool {
	if o != nil && !IsNil(o.SigningKeyId) {
		return true
	}

	return false
}

// SetSigningKeyId gets a reference to the given string and assigns it to the SigningKeyId field.
func (o *CredentialIssuanceOrder) SetSigningKeyId(v string) {
	o.SigningKeyId = &v
}

func (o CredentialIssuanceOrder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialIssuanceOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestIdentifier) {
		toSerialize["requestIdentifier"] = o.RequestIdentifier
	}
	if !IsNil(o.CredentialPayload) {
		toSerialize["credentialPayload"] = o.CredentialPayload
	}
	if !IsNil(o.IssuanceDeferred) {
		toSerialize["issuanceDeferred"] = o.IssuanceDeferred
	}
	if !IsNil(o.CredentialDuration) {
		toSerialize["credentialDuration"] = o.CredentialDuration
	}
	if !IsNil(o.SigningKeyId) {
		toSerialize["signingKeyId"] = o.SigningKeyId
	}
	return toSerialize, nil
}

type NullableCredentialIssuanceOrder struct {
	value *CredentialIssuanceOrder
	isSet bool
}

func (v NullableCredentialIssuanceOrder) Get() *CredentialIssuanceOrder {
	return v.value
}

func (v *NullableCredentialIssuanceOrder) Set(val *CredentialIssuanceOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialIssuanceOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialIssuanceOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialIssuanceOrder(val *CredentialIssuanceOrder) *NullableCredentialIssuanceOrder {
	return &NullableCredentialIssuanceOrder{value: val, isSet: true}
}

func (v NullableCredentialIssuanceOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialIssuanceOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
