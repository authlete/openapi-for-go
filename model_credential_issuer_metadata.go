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

// checks if the CredentialIssuerMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialIssuerMetadata{}

// CredentialIssuerMetadata struct for CredentialIssuerMetadata
type CredentialIssuerMetadata struct {
	// The identifier of a credential request.
	CredentialIssuer *string `json:"credentialIssuer,omitempty"`
	// The identifier of the authorization server that the credential issuer relies on for authorization.
	AuthorizationServer *string `json:"authorizationServer,omitempty"`
	// The URL of the credential endpoint of the credential issuer.
	CredentialEndpoint *bool `json:"credentialEndpoint,omitempty"`
	// The URL of the batch credential endpoint of the credential issuer.
	BatchCredentialEndpoint *int64 `json:"batchCredentialEndpoint,omitempty"`
	// The URL of the deferred credential endpoint of the credential issuer.
	DeferredCredentialEndpoint *string `json:"deferredCredentialEndpoint,omitempty"`
	// A JSON array describing supported credentials.
	CredentialsSupported *bool `json:"credentialsSupported,omitempty"`
}

// NewCredentialIssuerMetadata instantiates a new CredentialIssuerMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialIssuerMetadata() *CredentialIssuerMetadata {
	this := CredentialIssuerMetadata{}
	return &this
}

// NewCredentialIssuerMetadataWithDefaults instantiates a new CredentialIssuerMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialIssuerMetadataWithDefaults() *CredentialIssuerMetadata {
	this := CredentialIssuerMetadata{}
	return &this
}

// GetCredentialIssuer returns the CredentialIssuer field value if set, zero value otherwise.
func (o *CredentialIssuerMetadata) GetCredentialIssuer() string {
	if o == nil || IsNil(o.CredentialIssuer) {
		var ret string
		return ret
	}
	return *o.CredentialIssuer
}

// GetCredentialIssuerOk returns a tuple with the CredentialIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuerMetadata) GetCredentialIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialIssuer) {
		return nil, false
	}
	return o.CredentialIssuer, true
}

// HasCredentialIssuer returns a boolean if a field has been set.
func (o *CredentialIssuerMetadata) HasCredentialIssuer() bool {
	if o != nil && !IsNil(o.CredentialIssuer) {
		return true
	}

	return false
}

// SetCredentialIssuer gets a reference to the given string and assigns it to the CredentialIssuer field.
func (o *CredentialIssuerMetadata) SetCredentialIssuer(v string) {
	o.CredentialIssuer = &v
}

// GetAuthorizationServer returns the AuthorizationServer field value if set, zero value otherwise.
func (o *CredentialIssuerMetadata) GetAuthorizationServer() string {
	if o == nil || IsNil(o.AuthorizationServer) {
		var ret string
		return ret
	}
	return *o.AuthorizationServer
}

// GetAuthorizationServerOk returns a tuple with the AuthorizationServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuerMetadata) GetAuthorizationServerOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationServer) {
		return nil, false
	}
	return o.AuthorizationServer, true
}

// HasAuthorizationServer returns a boolean if a field has been set.
func (o *CredentialIssuerMetadata) HasAuthorizationServer() bool {
	if o != nil && !IsNil(o.AuthorizationServer) {
		return true
	}

	return false
}

// SetAuthorizationServer gets a reference to the given string and assigns it to the AuthorizationServer field.
func (o *CredentialIssuerMetadata) SetAuthorizationServer(v string) {
	o.AuthorizationServer = &v
}

// GetCredentialEndpoint returns the CredentialEndpoint field value if set, zero value otherwise.
func (o *CredentialIssuerMetadata) GetCredentialEndpoint() bool {
	if o == nil || IsNil(o.CredentialEndpoint) {
		var ret bool
		return ret
	}
	return *o.CredentialEndpoint
}

// GetCredentialEndpointOk returns a tuple with the CredentialEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuerMetadata) GetCredentialEndpointOk() (*bool, bool) {
	if o == nil || IsNil(o.CredentialEndpoint) {
		return nil, false
	}
	return o.CredentialEndpoint, true
}

// HasCredentialEndpoint returns a boolean if a field has been set.
func (o *CredentialIssuerMetadata) HasCredentialEndpoint() bool {
	if o != nil && !IsNil(o.CredentialEndpoint) {
		return true
	}

	return false
}

// SetCredentialEndpoint gets a reference to the given bool and assigns it to the CredentialEndpoint field.
func (o *CredentialIssuerMetadata) SetCredentialEndpoint(v bool) {
	o.CredentialEndpoint = &v
}

// GetBatchCredentialEndpoint returns the BatchCredentialEndpoint field value if set, zero value otherwise.
func (o *CredentialIssuerMetadata) GetBatchCredentialEndpoint() int64 {
	if o == nil || IsNil(o.BatchCredentialEndpoint) {
		var ret int64
		return ret
	}
	return *o.BatchCredentialEndpoint
}

// GetBatchCredentialEndpointOk returns a tuple with the BatchCredentialEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuerMetadata) GetBatchCredentialEndpointOk() (*int64, bool) {
	if o == nil || IsNil(o.BatchCredentialEndpoint) {
		return nil, false
	}
	return o.BatchCredentialEndpoint, true
}

// HasBatchCredentialEndpoint returns a boolean if a field has been set.
func (o *CredentialIssuerMetadata) HasBatchCredentialEndpoint() bool {
	if o != nil && !IsNil(o.BatchCredentialEndpoint) {
		return true
	}

	return false
}

// SetBatchCredentialEndpoint gets a reference to the given int64 and assigns it to the BatchCredentialEndpoint field.
func (o *CredentialIssuerMetadata) SetBatchCredentialEndpoint(v int64) {
	o.BatchCredentialEndpoint = &v
}

// GetDeferredCredentialEndpoint returns the DeferredCredentialEndpoint field value if set, zero value otherwise.
func (o *CredentialIssuerMetadata) GetDeferredCredentialEndpoint() string {
	if o == nil || IsNil(o.DeferredCredentialEndpoint) {
		var ret string
		return ret
	}
	return *o.DeferredCredentialEndpoint
}

// GetDeferredCredentialEndpointOk returns a tuple with the DeferredCredentialEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuerMetadata) GetDeferredCredentialEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.DeferredCredentialEndpoint) {
		return nil, false
	}
	return o.DeferredCredentialEndpoint, true
}

// HasDeferredCredentialEndpoint returns a boolean if a field has been set.
func (o *CredentialIssuerMetadata) HasDeferredCredentialEndpoint() bool {
	if o != nil && !IsNil(o.DeferredCredentialEndpoint) {
		return true
	}

	return false
}

// SetDeferredCredentialEndpoint gets a reference to the given string and assigns it to the DeferredCredentialEndpoint field.
func (o *CredentialIssuerMetadata) SetDeferredCredentialEndpoint(v string) {
	o.DeferredCredentialEndpoint = &v
}

// GetCredentialsSupported returns the CredentialsSupported field value if set, zero value otherwise.
func (o *CredentialIssuerMetadata) GetCredentialsSupported() bool {
	if o == nil || IsNil(o.CredentialsSupported) {
		var ret bool
		return ret
	}
	return *o.CredentialsSupported
}

// GetCredentialsSupportedOk returns a tuple with the CredentialsSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuerMetadata) GetCredentialsSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.CredentialsSupported) {
		return nil, false
	}
	return o.CredentialsSupported, true
}

// HasCredentialsSupported returns a boolean if a field has been set.
func (o *CredentialIssuerMetadata) HasCredentialsSupported() bool {
	if o != nil && !IsNil(o.CredentialsSupported) {
		return true
	}

	return false
}

// SetCredentialsSupported gets a reference to the given bool and assigns it to the CredentialsSupported field.
func (o *CredentialIssuerMetadata) SetCredentialsSupported(v bool) {
	o.CredentialsSupported = &v
}

func (o CredentialIssuerMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialIssuerMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CredentialIssuer) {
		toSerialize["credentialIssuer"] = o.CredentialIssuer
	}
	if !IsNil(o.AuthorizationServer) {
		toSerialize["authorizationServer"] = o.AuthorizationServer
	}
	if !IsNil(o.CredentialEndpoint) {
		toSerialize["credentialEndpoint"] = o.CredentialEndpoint
	}
	if !IsNil(o.BatchCredentialEndpoint) {
		toSerialize["batchCredentialEndpoint"] = o.BatchCredentialEndpoint
	}
	if !IsNil(o.DeferredCredentialEndpoint) {
		toSerialize["deferredCredentialEndpoint"] = o.DeferredCredentialEndpoint
	}
	if !IsNil(o.CredentialsSupported) {
		toSerialize["credentialsSupported"] = o.CredentialsSupported
	}
	return toSerialize, nil
}

type NullableCredentialIssuerMetadata struct {
	value *CredentialIssuerMetadata
	isSet bool
}

func (v NullableCredentialIssuerMetadata) Get() *CredentialIssuerMetadata {
	return v.value
}

func (v *NullableCredentialIssuerMetadata) Set(val *CredentialIssuerMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialIssuerMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialIssuerMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialIssuerMetadata(val *CredentialIssuerMetadata) *NullableCredentialIssuerMetadata {
	return &NullableCredentialIssuerMetadata{value: val, isSet: true}
}

func (v NullableCredentialIssuerMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialIssuerMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
