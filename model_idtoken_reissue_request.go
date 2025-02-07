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

// checks if the IdtokenReissueRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdtokenReissueRequest{}

// IdtokenReissueRequest struct for IdtokenReissueRequest
type IdtokenReissueRequest struct {
	// <p> The value of this parameter should be (a) the value of the \"`jwtAccessToken`\" parameter in a response from the `/auth/token` API when the value is available, or (b) the value of the \"`accessToken`\" parameter in the response from the `/auth/token` API when the value of the \"`jwtAccessToken`\" parameter is not available. </p>
	AccessToken string `json:"accessToken"`
	// <p> The value of this parameter should be the value of the \"`refreshToken`\" parameter in a response from the `/auth/token` API. </p>
	RefreshToken string `json:"refreshToken"`
	// The value that should be used as the value of the \"`sub`\" claim of the ID token.  <p> This parameter is optional. When omitted, the value of the subject associated with the access token is used. </p>
	Sub *string `json:"sub,omitempty"`
	// Additional claims that should be embedded in the payload part of the ID token. The format is a JSON object.  <p> This parameter is optional. </p>
	Claims *string `json:"claims,omitempty"`
	// Additional parameters that should be embedded in the JWS header of the ID token. The format is a JSON object.  <p> This parameter is optional. </p>
	IdtHeaderParams *string `json:"idtHeaderParams,omitempty"`
	// The type of the \"`aud`\" claim of the ID token being issued.  <p> Valid values of this parameter are as follows. </p>  <blockquote> <table border=\"1\" cellpadding=\"5\" style=\"border-collapse: collapse;\">   <tr bgcolor=\"orange\">     <th>Value</th>     <th>Description</th>   </tr>   <tr>     <td>\"`array`\"</td>     <td>The type of the `aud` claim becomes an array of strings.</td>   </tr>   <tr>     <td>\"`string`\"</td>     <td>The type of the `aud` claim becomes a single string.</td>   </tr> </table> </blockquote>  <p> This parameter is optional, and the default value on omission is \"`array`\". </p>  <p> This parameter takes precedence over the `idTokenAudType` property of {@link Service} (cf. {@link Service#getIdTokenAudType()}). </p>
	IdTokenAudType *string `json:"idTokenAudType,omitempty"`
}

type _IdtokenReissueRequest IdtokenReissueRequest

// NewIdtokenReissueRequest instantiates a new IdtokenReissueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdtokenReissueRequest(accessToken string, refreshToken string) *IdtokenReissueRequest {
	this := IdtokenReissueRequest{}
	this.AccessToken = accessToken
	this.RefreshToken = refreshToken
	return &this
}

// NewIdtokenReissueRequestWithDefaults instantiates a new IdtokenReissueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdtokenReissueRequestWithDefaults() *IdtokenReissueRequest {
	this := IdtokenReissueRequest{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *IdtokenReissueRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *IdtokenReissueRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *IdtokenReissueRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetRefreshToken returns the RefreshToken field value
func (o *IdtokenReissueRequest) GetRefreshToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value
// and a boolean to check if the value has been set.
func (o *IdtokenReissueRequest) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshToken, true
}

// SetRefreshToken sets field value
func (o *IdtokenReissueRequest) SetRefreshToken(v string) {
	o.RefreshToken = v
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *IdtokenReissueRequest) GetSub() string {
	if o == nil || IsNil(o.Sub) {
		var ret string
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdtokenReissueRequest) GetSubOk() (*string, bool) {
	if o == nil || IsNil(o.Sub) {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *IdtokenReissueRequest) HasSub() bool {
	if o != nil && !IsNil(o.Sub) {
		return true
	}

	return false
}

// SetSub gets a reference to the given string and assigns it to the Sub field.
func (o *IdtokenReissueRequest) SetSub(v string) {
	o.Sub = &v
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *IdtokenReissueRequest) GetClaims() string {
	if o == nil || IsNil(o.Claims) {
		var ret string
		return ret
	}
	return *o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdtokenReissueRequest) GetClaimsOk() (*string, bool) {
	if o == nil || IsNil(o.Claims) {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *IdtokenReissueRequest) HasClaims() bool {
	if o != nil && !IsNil(o.Claims) {
		return true
	}

	return false
}

// SetClaims gets a reference to the given string and assigns it to the Claims field.
func (o *IdtokenReissueRequest) SetClaims(v string) {
	o.Claims = &v
}

// GetIdtHeaderParams returns the IdtHeaderParams field value if set, zero value otherwise.
func (o *IdtokenReissueRequest) GetIdtHeaderParams() string {
	if o == nil || IsNil(o.IdtHeaderParams) {
		var ret string
		return ret
	}
	return *o.IdtHeaderParams
}

// GetIdtHeaderParamsOk returns a tuple with the IdtHeaderParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdtokenReissueRequest) GetIdtHeaderParamsOk() (*string, bool) {
	if o == nil || IsNil(o.IdtHeaderParams) {
		return nil, false
	}
	return o.IdtHeaderParams, true
}

// HasIdtHeaderParams returns a boolean if a field has been set.
func (o *IdtokenReissueRequest) HasIdtHeaderParams() bool {
	if o != nil && !IsNil(o.IdtHeaderParams) {
		return true
	}

	return false
}

// SetIdtHeaderParams gets a reference to the given string and assigns it to the IdtHeaderParams field.
func (o *IdtokenReissueRequest) SetIdtHeaderParams(v string) {
	o.IdtHeaderParams = &v
}

// GetIdTokenAudType returns the IdTokenAudType field value if set, zero value otherwise.
func (o *IdtokenReissueRequest) GetIdTokenAudType() string {
	if o == nil || IsNil(o.IdTokenAudType) {
		var ret string
		return ret
	}
	return *o.IdTokenAudType
}

// GetIdTokenAudTypeOk returns a tuple with the IdTokenAudType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdtokenReissueRequest) GetIdTokenAudTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IdTokenAudType) {
		return nil, false
	}
	return o.IdTokenAudType, true
}

// HasIdTokenAudType returns a boolean if a field has been set.
func (o *IdtokenReissueRequest) HasIdTokenAudType() bool {
	if o != nil && !IsNil(o.IdTokenAudType) {
		return true
	}

	return false
}

// SetIdTokenAudType gets a reference to the given string and assigns it to the IdTokenAudType field.
func (o *IdtokenReissueRequest) SetIdTokenAudType(v string) {
	o.IdTokenAudType = &v
}

func (o IdtokenReissueRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdtokenReissueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accessToken"] = o.AccessToken
	toSerialize["refreshToken"] = o.RefreshToken
	if !IsNil(o.Sub) {
		toSerialize["sub"] = o.Sub
	}
	if !IsNil(o.Claims) {
		toSerialize["claims"] = o.Claims
	}
	if !IsNil(o.IdtHeaderParams) {
		toSerialize["idtHeaderParams"] = o.IdtHeaderParams
	}
	if !IsNil(o.IdTokenAudType) {
		toSerialize["idTokenAudType"] = o.IdTokenAudType
	}
	return toSerialize, nil
}

func (o *IdtokenReissueRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accessToken",
		"refreshToken",
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

	varIdtokenReissueRequest := _IdtokenReissueRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIdtokenReissueRequest)

	if err != nil {
		return err
	}

	*o = IdtokenReissueRequest(varIdtokenReissueRequest)

	return err
}

type NullableIdtokenReissueRequest struct {
	value *IdtokenReissueRequest
	isSet bool
}

func (v NullableIdtokenReissueRequest) Get() *IdtokenReissueRequest {
	return v.value
}

func (v *NullableIdtokenReissueRequest) Set(val *IdtokenReissueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdtokenReissueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdtokenReissueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdtokenReissueRequest(val *IdtokenReissueRequest) *NullableIdtokenReissueRequest {
	return &NullableIdtokenReissueRequest{value: val, isSet: true}
}

func (v NullableIdtokenReissueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdtokenReissueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
