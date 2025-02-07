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

// checks if the PushedAuthorizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushedAuthorizationRequest{}

// PushedAuthorizationRequest struct for PushedAuthorizationRequest
type PushedAuthorizationRequest struct {
	// The pushed authorization request body received from the client application.  The value of parameters is the entire entity body (which is formatted in `application/x-www-form-urlencoded`) of the request from the client application.
	Parameters string `json:"parameters"`
	// The client ID extracted from `Authorization` header of the pushed request from the client application.
	ClientId *string `json:"clientId,omitempty"`
	// The client secret extracted from `Authorization` header of the pushed authorization request from the client application.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// The client certificate from the MTLS connection to pushed authorization endpoint from the client application.
	ClientCertificate *string `json:"clientCertificate,omitempty"`
	// The certificate path presented by the client during client authentication. These certificates are strings in PEM format.
	ClientCertificatePath *string `json:"clientCertificatePath,omitempty"`
	// DPoP Header
	Dpop *string `json:"dpop,omitempty"`
	// HTTP Method (for DPoP validation).
	Htm *string `json:"htm,omitempty"`
	// HTTP URL base (for DPoP validation).
	Htu *string `json:"htu,omitempty"`
}

type _PushedAuthorizationRequest PushedAuthorizationRequest

// NewPushedAuthorizationRequest instantiates a new PushedAuthorizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushedAuthorizationRequest(parameters string) *PushedAuthorizationRequest {
	this := PushedAuthorizationRequest{}
	this.Parameters = parameters
	return &this
}

// NewPushedAuthorizationRequestWithDefaults instantiates a new PushedAuthorizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushedAuthorizationRequestWithDefaults() *PushedAuthorizationRequest {
	this := PushedAuthorizationRequest{}
	return &this
}

// GetParameters returns the Parameters field value
func (o *PushedAuthorizationRequest) GetParameters() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationRequest) GetParametersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *PushedAuthorizationRequest) SetParameters(v string) {
	o.Parameters = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PushedAuthorizationRequest) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationRequest) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PushedAuthorizationRequest) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PushedAuthorizationRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *PushedAuthorizationRequest) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *PushedAuthorizationRequest) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *PushedAuthorizationRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetClientCertificate returns the ClientCertificate field value if set, zero value otherwise.
func (o *PushedAuthorizationRequest) GetClientCertificate() string {
	if o == nil || IsNil(o.ClientCertificate) {
		var ret string
		return ret
	}
	return *o.ClientCertificate
}

// GetClientCertificateOk returns a tuple with the ClientCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationRequest) GetClientCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.ClientCertificate) {
		return nil, false
	}
	return o.ClientCertificate, true
}

// HasClientCertificate returns a boolean if a field has been set.
func (o *PushedAuthorizationRequest) HasClientCertificate() bool {
	if o != nil && !IsNil(o.ClientCertificate) {
		return true
	}

	return false
}

// SetClientCertificate gets a reference to the given string and assigns it to the ClientCertificate field.
func (o *PushedAuthorizationRequest) SetClientCertificate(v string) {
	o.ClientCertificate = &v
}

// GetClientCertificatePath returns the ClientCertificatePath field value if set, zero value otherwise.
func (o *PushedAuthorizationRequest) GetClientCertificatePath() string {
	if o == nil || IsNil(o.ClientCertificatePath) {
		var ret string
		return ret
	}
	return *o.ClientCertificatePath
}

// GetClientCertificatePathOk returns a tuple with the ClientCertificatePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationRequest) GetClientCertificatePathOk() (*string, bool) {
	if o == nil || IsNil(o.ClientCertificatePath) {
		return nil, false
	}
	return o.ClientCertificatePath, true
}

// HasClientCertificatePath returns a boolean if a field has been set.
func (o *PushedAuthorizationRequest) HasClientCertificatePath() bool {
	if o != nil && !IsNil(o.ClientCertificatePath) {
		return true
	}

	return false
}

// SetClientCertificatePath gets a reference to the given string and assigns it to the ClientCertificatePath field.
func (o *PushedAuthorizationRequest) SetClientCertificatePath(v string) {
	o.ClientCertificatePath = &v
}

// GetDpop returns the Dpop field value if set, zero value otherwise.
func (o *PushedAuthorizationRequest) GetDpop() string {
	if o == nil || IsNil(o.Dpop) {
		var ret string
		return ret
	}
	return *o.Dpop
}

// GetDpopOk returns a tuple with the Dpop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationRequest) GetDpopOk() (*string, bool) {
	if o == nil || IsNil(o.Dpop) {
		return nil, false
	}
	return o.Dpop, true
}

// HasDpop returns a boolean if a field has been set.
func (o *PushedAuthorizationRequest) HasDpop() bool {
	if o != nil && !IsNil(o.Dpop) {
		return true
	}

	return false
}

// SetDpop gets a reference to the given string and assigns it to the Dpop field.
func (o *PushedAuthorizationRequest) SetDpop(v string) {
	o.Dpop = &v
}

// GetHtm returns the Htm field value if set, zero value otherwise.
func (o *PushedAuthorizationRequest) GetHtm() string {
	if o == nil || IsNil(o.Htm) {
		var ret string
		return ret
	}
	return *o.Htm
}

// GetHtmOk returns a tuple with the Htm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationRequest) GetHtmOk() (*string, bool) {
	if o == nil || IsNil(o.Htm) {
		return nil, false
	}
	return o.Htm, true
}

// HasHtm returns a boolean if a field has been set.
func (o *PushedAuthorizationRequest) HasHtm() bool {
	if o != nil && !IsNil(o.Htm) {
		return true
	}

	return false
}

// SetHtm gets a reference to the given string and assigns it to the Htm field.
func (o *PushedAuthorizationRequest) SetHtm(v string) {
	o.Htm = &v
}

// GetHtu returns the Htu field value if set, zero value otherwise.
func (o *PushedAuthorizationRequest) GetHtu() string {
	if o == nil || IsNil(o.Htu) {
		var ret string
		return ret
	}
	return *o.Htu
}

// GetHtuOk returns a tuple with the Htu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushedAuthorizationRequest) GetHtuOk() (*string, bool) {
	if o == nil || IsNil(o.Htu) {
		return nil, false
	}
	return o.Htu, true
}

// HasHtu returns a boolean if a field has been set.
func (o *PushedAuthorizationRequest) HasHtu() bool {
	if o != nil && !IsNil(o.Htu) {
		return true
	}

	return false
}

// SetHtu gets a reference to the given string and assigns it to the Htu field.
func (o *PushedAuthorizationRequest) SetHtu(v string) {
	o.Htu = &v
}

func (o PushedAuthorizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushedAuthorizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parameters"] = o.Parameters
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !IsNil(o.ClientCertificate) {
		toSerialize["clientCertificate"] = o.ClientCertificate
	}
	if !IsNil(o.ClientCertificatePath) {
		toSerialize["clientCertificatePath"] = o.ClientCertificatePath
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
	return toSerialize, nil
}

func (o *PushedAuthorizationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parameters",
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

	varPushedAuthorizationRequest := _PushedAuthorizationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPushedAuthorizationRequest)

	if err != nil {
		return err
	}

	*o = PushedAuthorizationRequest(varPushedAuthorizationRequest)

	return err
}

type NullablePushedAuthorizationRequest struct {
	value *PushedAuthorizationRequest
	isSet bool
}

func (v NullablePushedAuthorizationRequest) Get() *PushedAuthorizationRequest {
	return v.value
}

func (v *NullablePushedAuthorizationRequest) Set(val *PushedAuthorizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePushedAuthorizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePushedAuthorizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushedAuthorizationRequest(val *PushedAuthorizationRequest) *NullablePushedAuthorizationRequest {
	return &NullablePushedAuthorizationRequest{value: val, isSet: true}
}

func (v NullablePushedAuthorizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushedAuthorizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
