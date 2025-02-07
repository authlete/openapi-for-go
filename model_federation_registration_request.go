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

// checks if the FederationRegistrationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FederationRegistrationRequest{}

// FederationRegistrationRequest struct for FederationRegistrationRequest
type FederationRegistrationRequest struct {
	// The entity configuration of a relying party.
	EntityConfiguration *string `json:"entityConfiguration,omitempty"`
	// The trust chain of a relying party.
	TrustChain *string `json:"trustChain,omitempty"`
}

// NewFederationRegistrationRequest instantiates a new FederationRegistrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederationRegistrationRequest() *FederationRegistrationRequest {
	this := FederationRegistrationRequest{}
	return &this
}

// NewFederationRegistrationRequestWithDefaults instantiates a new FederationRegistrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederationRegistrationRequestWithDefaults() *FederationRegistrationRequest {
	this := FederationRegistrationRequest{}
	return &this
}

// GetEntityConfiguration returns the EntityConfiguration field value if set, zero value otherwise.
func (o *FederationRegistrationRequest) GetEntityConfiguration() string {
	if o == nil || IsNil(o.EntityConfiguration) {
		var ret string
		return ret
	}
	return *o.EntityConfiguration
}

// GetEntityConfigurationOk returns a tuple with the EntityConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationRegistrationRequest) GetEntityConfigurationOk() (*string, bool) {
	if o == nil || IsNil(o.EntityConfiguration) {
		return nil, false
	}
	return o.EntityConfiguration, true
}

// HasEntityConfiguration returns a boolean if a field has been set.
func (o *FederationRegistrationRequest) HasEntityConfiguration() bool {
	if o != nil && !IsNil(o.EntityConfiguration) {
		return true
	}

	return false
}

// SetEntityConfiguration gets a reference to the given string and assigns it to the EntityConfiguration field.
func (o *FederationRegistrationRequest) SetEntityConfiguration(v string) {
	o.EntityConfiguration = &v
}

// GetTrustChain returns the TrustChain field value if set, zero value otherwise.
func (o *FederationRegistrationRequest) GetTrustChain() string {
	if o == nil || IsNil(o.TrustChain) {
		var ret string
		return ret
	}
	return *o.TrustChain
}

// GetTrustChainOk returns a tuple with the TrustChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationRegistrationRequest) GetTrustChainOk() (*string, bool) {
	if o == nil || IsNil(o.TrustChain) {
		return nil, false
	}
	return o.TrustChain, true
}

// HasTrustChain returns a boolean if a field has been set.
func (o *FederationRegistrationRequest) HasTrustChain() bool {
	if o != nil && !IsNil(o.TrustChain) {
		return true
	}

	return false
}

// SetTrustChain gets a reference to the given string and assigns it to the TrustChain field.
func (o *FederationRegistrationRequest) SetTrustChain(v string) {
	o.TrustChain = &v
}

func (o FederationRegistrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FederationRegistrationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityConfiguration) {
		toSerialize["entityConfiguration"] = o.EntityConfiguration
	}
	if !IsNil(o.TrustChain) {
		toSerialize["trustChain"] = o.TrustChain
	}
	return toSerialize, nil
}

type NullableFederationRegistrationRequest struct {
	value *FederationRegistrationRequest
	isSet bool
}

func (v NullableFederationRegistrationRequest) Get() *FederationRegistrationRequest {
	return v.value
}

func (v *NullableFederationRegistrationRequest) Set(val *FederationRegistrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFederationRegistrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFederationRegistrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederationRegistrationRequest(val *FederationRegistrationRequest) *NullableFederationRegistrationRequest {
	return &NullableFederationRegistrationRequest{value: val, isSet: true}
}

func (v NullableFederationRegistrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederationRegistrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
