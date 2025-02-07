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

// checks if the HskCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HskCreateRequest{}

// HskCreateRequest struct for HskCreateRequest
type HskCreateRequest struct {
	// The key type (EC or RSA)
	Kty *string `json:"kty,omitempty"`
	// The key on the HSM.  When the key use is \"sig\" (signature), the private key on the HSM is used to sign data and the corresponding public key is used to verify the signature. When the key use is \"enc\" (encryption), the private key on the HSM is used to decrypt encrypted data which have been encrypted with the corresponding public key
	Use *string `json:"use,omitempty"`
	// Key ID for the key on the HSM.
	Kid *string `json:"kid,omitempty"`
	// The name of the HSM. The identifier for the HSM that sits behind the Authlete server. For example, \"google\".
	HsmName *string `json:"hsmName,omitempty"`
	// The handle for the key on the HSM. A handle is a base64url-encoded 256-bit random value (43 letters) which is assigned by Authlete on the call of the /api/hsk/create API
	Handle *string `json:"handle,omitempty"`
	// The public key that corresponds to the key on the HSM.
	PublicKey *string `json:"publicKey,omitempty"`
}

// NewHskCreateRequest instantiates a new HskCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHskCreateRequest() *HskCreateRequest {
	this := HskCreateRequest{}
	return &this
}

// NewHskCreateRequestWithDefaults instantiates a new HskCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHskCreateRequestWithDefaults() *HskCreateRequest {
	this := HskCreateRequest{}
	return &this
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *HskCreateRequest) GetKty() string {
	if o == nil || IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HskCreateRequest) GetKtyOk() (*string, bool) {
	if o == nil || IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *HskCreateRequest) HasKty() bool {
	if o != nil && !IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *HskCreateRequest) SetKty(v string) {
	o.Kty = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *HskCreateRequest) GetUse() string {
	if o == nil || IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HskCreateRequest) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *HskCreateRequest) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *HskCreateRequest) SetUse(v string) {
	o.Use = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *HskCreateRequest) GetKid() string {
	if o == nil || IsNil(o.Kid) {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HskCreateRequest) GetKidOk() (*string, bool) {
	if o == nil || IsNil(o.Kid) {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *HskCreateRequest) HasKid() bool {
	if o != nil && !IsNil(o.Kid) {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *HskCreateRequest) SetKid(v string) {
	o.Kid = &v
}

// GetHsmName returns the HsmName field value if set, zero value otherwise.
func (o *HskCreateRequest) GetHsmName() string {
	if o == nil || IsNil(o.HsmName) {
		var ret string
		return ret
	}
	return *o.HsmName
}

// GetHsmNameOk returns a tuple with the HsmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HskCreateRequest) GetHsmNameOk() (*string, bool) {
	if o == nil || IsNil(o.HsmName) {
		return nil, false
	}
	return o.HsmName, true
}

// HasHsmName returns a boolean if a field has been set.
func (o *HskCreateRequest) HasHsmName() bool {
	if o != nil && !IsNil(o.HsmName) {
		return true
	}

	return false
}

// SetHsmName gets a reference to the given string and assigns it to the HsmName field.
func (o *HskCreateRequest) SetHsmName(v string) {
	o.HsmName = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *HskCreateRequest) GetHandle() string {
	if o == nil || IsNil(o.Handle) {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HskCreateRequest) GetHandleOk() (*string, bool) {
	if o == nil || IsNil(o.Handle) {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *HskCreateRequest) HasHandle() bool {
	if o != nil && !IsNil(o.Handle) {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *HskCreateRequest) SetHandle(v string) {
	o.Handle = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *HskCreateRequest) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HskCreateRequest) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *HskCreateRequest) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *HskCreateRequest) SetPublicKey(v string) {
	o.PublicKey = &v
}

func (o HskCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HskCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kty) {
		toSerialize["kty"] = o.Kty
	}
	if !IsNil(o.Use) {
		toSerialize["use"] = o.Use
	}
	if !IsNil(o.Kid) {
		toSerialize["kid"] = o.Kid
	}
	if !IsNil(o.HsmName) {
		toSerialize["hsmName"] = o.HsmName
	}
	if !IsNil(o.Handle) {
		toSerialize["handle"] = o.Handle
	}
	if !IsNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	return toSerialize, nil
}

type NullableHskCreateRequest struct {
	value *HskCreateRequest
	isSet bool
}

func (v NullableHskCreateRequest) Get() *HskCreateRequest {
	return v.value
}

func (v *NullableHskCreateRequest) Set(val *HskCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHskCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHskCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHskCreateRequest(val *HskCreateRequest) *NullableHskCreateRequest {
	return &NullableHskCreateRequest{value: val, isSet: true}
}

func (v NullableHskCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHskCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
