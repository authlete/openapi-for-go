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

// checks if the UserinfoIssueRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserinfoIssueRequest{}

// UserinfoIssueRequest struct for UserinfoIssueRequest
type UserinfoIssueRequest struct {
	// The access token that has been passed to the userinfo endpoint by the client application. In other words, the access token which was contained in the userinfo request.
	Token string `json:"token"`
	// Claims in JSON format. As for the format, see [OpenID Connect Core 1.0, 5.1. Standard Claims](https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims).
	Claims *string `json:"claims,omitempty"`
	// The value of the `sub` claim. If the value of this request parameter is not empty, it is used as the value of the `sub` claim. Otherwise, the value of the subject associated with the access token is used.
	Sub *string `json:"sub,omitempty"`
	// Claim key-value pairs that are used to compute transformed claims.
	ClaimsForTx *string `json:"claimsForTx,omitempty"`
	// The Signature header value from the request.
	RequestSignature *string `json:"requestSignature,omitempty"`
	// HTTP headers to be included in processing the signature. If this is a signed request, this must include the  Signature and Signature-Input headers, as well as any additional headers covered by the signature.
	Headers []Pair `json:"headers,omitempty"`
}

type _UserinfoIssueRequest UserinfoIssueRequest

// NewUserinfoIssueRequest instantiates a new UserinfoIssueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserinfoIssueRequest(token string) *UserinfoIssueRequest {
	this := UserinfoIssueRequest{}
	this.Token = token
	return &this
}

// NewUserinfoIssueRequestWithDefaults instantiates a new UserinfoIssueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserinfoIssueRequestWithDefaults() *UserinfoIssueRequest {
	this := UserinfoIssueRequest{}
	return &this
}

// GetToken returns the Token field value
func (o *UserinfoIssueRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UserinfoIssueRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UserinfoIssueRequest) SetToken(v string) {
	o.Token = v
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *UserinfoIssueRequest) GetClaims() string {
	if o == nil || IsNil(o.Claims) {
		var ret string
		return ret
	}
	return *o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueRequest) GetClaimsOk() (*string, bool) {
	if o == nil || IsNil(o.Claims) {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *UserinfoIssueRequest) HasClaims() bool {
	if o != nil && !IsNil(o.Claims) {
		return true
	}

	return false
}

// SetClaims gets a reference to the given string and assigns it to the Claims field.
func (o *UserinfoIssueRequest) SetClaims(v string) {
	o.Claims = &v
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *UserinfoIssueRequest) GetSub() string {
	if o == nil || IsNil(o.Sub) {
		var ret string
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueRequest) GetSubOk() (*string, bool) {
	if o == nil || IsNil(o.Sub) {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *UserinfoIssueRequest) HasSub() bool {
	if o != nil && !IsNil(o.Sub) {
		return true
	}

	return false
}

// SetSub gets a reference to the given string and assigns it to the Sub field.
func (o *UserinfoIssueRequest) SetSub(v string) {
	o.Sub = &v
}

// GetClaimsForTx returns the ClaimsForTx field value if set, zero value otherwise.
func (o *UserinfoIssueRequest) GetClaimsForTx() string {
	if o == nil || IsNil(o.ClaimsForTx) {
		var ret string
		return ret
	}
	return *o.ClaimsForTx
}

// GetClaimsForTxOk returns a tuple with the ClaimsForTx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueRequest) GetClaimsForTxOk() (*string, bool) {
	if o == nil || IsNil(o.ClaimsForTx) {
		return nil, false
	}
	return o.ClaimsForTx, true
}

// HasClaimsForTx returns a boolean if a field has been set.
func (o *UserinfoIssueRequest) HasClaimsForTx() bool {
	if o != nil && !IsNil(o.ClaimsForTx) {
		return true
	}

	return false
}

// SetClaimsForTx gets a reference to the given string and assigns it to the ClaimsForTx field.
func (o *UserinfoIssueRequest) SetClaimsForTx(v string) {
	o.ClaimsForTx = &v
}

// GetRequestSignature returns the RequestSignature field value if set, zero value otherwise.
func (o *UserinfoIssueRequest) GetRequestSignature() string {
	if o == nil || IsNil(o.RequestSignature) {
		var ret string
		return ret
	}
	return *o.RequestSignature
}

// GetRequestSignatureOk returns a tuple with the RequestSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueRequest) GetRequestSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.RequestSignature) {
		return nil, false
	}
	return o.RequestSignature, true
}

// HasRequestSignature returns a boolean if a field has been set.
func (o *UserinfoIssueRequest) HasRequestSignature() bool {
	if o != nil && !IsNil(o.RequestSignature) {
		return true
	}

	return false
}

// SetRequestSignature gets a reference to the given string and assigns it to the RequestSignature field.
func (o *UserinfoIssueRequest) SetRequestSignature(v string) {
	o.RequestSignature = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *UserinfoIssueRequest) GetHeaders() []Pair {
	if o == nil || IsNil(o.Headers) {
		var ret []Pair
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoIssueRequest) GetHeadersOk() ([]Pair, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *UserinfoIssueRequest) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []Pair and assigns it to the Headers field.
func (o *UserinfoIssueRequest) SetHeaders(v []Pair) {
	o.Headers = v
}

func (o UserinfoIssueRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserinfoIssueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	if !IsNil(o.Claims) {
		toSerialize["claims"] = o.Claims
	}
	if !IsNil(o.Sub) {
		toSerialize["sub"] = o.Sub
	}
	if !IsNil(o.ClaimsForTx) {
		toSerialize["claimsForTx"] = o.ClaimsForTx
	}
	if !IsNil(o.RequestSignature) {
		toSerialize["requestSignature"] = o.RequestSignature
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	return toSerialize, nil
}

func (o *UserinfoIssueRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
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

	varUserinfoIssueRequest := _UserinfoIssueRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserinfoIssueRequest)

	if err != nil {
		return err
	}

	*o = UserinfoIssueRequest(varUserinfoIssueRequest)

	return err
}

type NullableUserinfoIssueRequest struct {
	value *UserinfoIssueRequest
	isSet bool
}

func (v NullableUserinfoIssueRequest) Get() *UserinfoIssueRequest {
	return v.value
}

func (v *NullableUserinfoIssueRequest) Set(val *UserinfoIssueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserinfoIssueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserinfoIssueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserinfoIssueRequest(val *UserinfoIssueRequest) *NullableUserinfoIssueRequest {
	return &NullableUserinfoIssueRequest{value: val, isSet: true}
}

func (v NullableUserinfoIssueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserinfoIssueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
