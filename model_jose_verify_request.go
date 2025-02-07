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

// checks if the JoseVerifyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JoseVerifyRequest{}

// JoseVerifyRequest struct for JoseVerifyRequest
type JoseVerifyRequest struct {
	// A JOSE object.
	Jose string `json:"jose"`
	// Mandatory claims that are required to be included in the JOSE object.
	MandatoryClaims *string `json:"mandatoryClaims,omitempty"`
	// Allowable clock skew in seconds.
	ClockSkew *int32 `json:"clockSkew,omitempty"`
	// The identifier of the client application whose keys are required for verification of the JOSE object.
	ClientIdentifier *string `json:"clientIdentifier,omitempty"`
	// The flag which indicates whether the signature of the JOSE object has been signed by a client application with the client's private key or a shared symmetric key.
	SignedByClient *bool `json:"signedByClient,omitempty"`
}

type _JoseVerifyRequest JoseVerifyRequest

// NewJoseVerifyRequest instantiates a new JoseVerifyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJoseVerifyRequest(jose string) *JoseVerifyRequest {
	this := JoseVerifyRequest{}
	this.Jose = jose
	return &this
}

// NewJoseVerifyRequestWithDefaults instantiates a new JoseVerifyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJoseVerifyRequestWithDefaults() *JoseVerifyRequest {
	this := JoseVerifyRequest{}
	return &this
}

// GetJose returns the Jose field value
func (o *JoseVerifyRequest) GetJose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Jose
}

// GetJoseOk returns a tuple with the Jose field value
// and a boolean to check if the value has been set.
func (o *JoseVerifyRequest) GetJoseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jose, true
}

// SetJose sets field value
func (o *JoseVerifyRequest) SetJose(v string) {
	o.Jose = v
}

// GetMandatoryClaims returns the MandatoryClaims field value if set, zero value otherwise.
func (o *JoseVerifyRequest) GetMandatoryClaims() string {
	if o == nil || IsNil(o.MandatoryClaims) {
		var ret string
		return ret
	}
	return *o.MandatoryClaims
}

// GetMandatoryClaimsOk returns a tuple with the MandatoryClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoseVerifyRequest) GetMandatoryClaimsOk() (*string, bool) {
	if o == nil || IsNil(o.MandatoryClaims) {
		return nil, false
	}
	return o.MandatoryClaims, true
}

// HasMandatoryClaims returns a boolean if a field has been set.
func (o *JoseVerifyRequest) HasMandatoryClaims() bool {
	if o != nil && !IsNil(o.MandatoryClaims) {
		return true
	}

	return false
}

// SetMandatoryClaims gets a reference to the given string and assigns it to the MandatoryClaims field.
func (o *JoseVerifyRequest) SetMandatoryClaims(v string) {
	o.MandatoryClaims = &v
}

// GetClockSkew returns the ClockSkew field value if set, zero value otherwise.
func (o *JoseVerifyRequest) GetClockSkew() int32 {
	if o == nil || IsNil(o.ClockSkew) {
		var ret int32
		return ret
	}
	return *o.ClockSkew
}

// GetClockSkewOk returns a tuple with the ClockSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoseVerifyRequest) GetClockSkewOk() (*int32, bool) {
	if o == nil || IsNil(o.ClockSkew) {
		return nil, false
	}
	return o.ClockSkew, true
}

// HasClockSkew returns a boolean if a field has been set.
func (o *JoseVerifyRequest) HasClockSkew() bool {
	if o != nil && !IsNil(o.ClockSkew) {
		return true
	}

	return false
}

// SetClockSkew gets a reference to the given int32 and assigns it to the ClockSkew field.
func (o *JoseVerifyRequest) SetClockSkew(v int32) {
	o.ClockSkew = &v
}

// GetClientIdentifier returns the ClientIdentifier field value if set, zero value otherwise.
func (o *JoseVerifyRequest) GetClientIdentifier() string {
	if o == nil || IsNil(o.ClientIdentifier) {
		var ret string
		return ret
	}
	return *o.ClientIdentifier
}

// GetClientIdentifierOk returns a tuple with the ClientIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoseVerifyRequest) GetClientIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.ClientIdentifier) {
		return nil, false
	}
	return o.ClientIdentifier, true
}

// HasClientIdentifier returns a boolean if a field has been set.
func (o *JoseVerifyRequest) HasClientIdentifier() bool {
	if o != nil && !IsNil(o.ClientIdentifier) {
		return true
	}

	return false
}

// SetClientIdentifier gets a reference to the given string and assigns it to the ClientIdentifier field.
func (o *JoseVerifyRequest) SetClientIdentifier(v string) {
	o.ClientIdentifier = &v
}

// GetSignedByClient returns the SignedByClient field value if set, zero value otherwise.
func (o *JoseVerifyRequest) GetSignedByClient() bool {
	if o == nil || IsNil(o.SignedByClient) {
		var ret bool
		return ret
	}
	return *o.SignedByClient
}

// GetSignedByClientOk returns a tuple with the SignedByClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JoseVerifyRequest) GetSignedByClientOk() (*bool, bool) {
	if o == nil || IsNil(o.SignedByClient) {
		return nil, false
	}
	return o.SignedByClient, true
}

// HasSignedByClient returns a boolean if a field has been set.
func (o *JoseVerifyRequest) HasSignedByClient() bool {
	if o != nil && !IsNil(o.SignedByClient) {
		return true
	}

	return false
}

// SetSignedByClient gets a reference to the given bool and assigns it to the SignedByClient field.
func (o *JoseVerifyRequest) SetSignedByClient(v bool) {
	o.SignedByClient = &v
}

func (o JoseVerifyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JoseVerifyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jose"] = o.Jose
	if !IsNil(o.MandatoryClaims) {
		toSerialize["mandatoryClaims"] = o.MandatoryClaims
	}
	if !IsNil(o.ClockSkew) {
		toSerialize["clockSkew"] = o.ClockSkew
	}
	if !IsNil(o.ClientIdentifier) {
		toSerialize["clientIdentifier"] = o.ClientIdentifier
	}
	if !IsNil(o.SignedByClient) {
		toSerialize["signedByClient"] = o.SignedByClient
	}
	return toSerialize, nil
}

func (o *JoseVerifyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"jose",
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

	varJoseVerifyRequest := _JoseVerifyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJoseVerifyRequest)

	if err != nil {
		return err
	}

	*o = JoseVerifyRequest(varJoseVerifyRequest)

	return err
}

type NullableJoseVerifyRequest struct {
	value *JoseVerifyRequest
	isSet bool
}

func (v NullableJoseVerifyRequest) Get() *JoseVerifyRequest {
	return v.value
}

func (v *NullableJoseVerifyRequest) Set(val *JoseVerifyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableJoseVerifyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableJoseVerifyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJoseVerifyRequest(val *JoseVerifyRequest) *NullableJoseVerifyRequest {
	return &NullableJoseVerifyRequest{value: val, isSet: true}
}

func (v NullableJoseVerifyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJoseVerifyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
