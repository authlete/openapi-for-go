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

// checks if the DeviceCompleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceCompleteRequest{}

// DeviceCompleteRequest struct for DeviceCompleteRequest
type DeviceCompleteRequest struct {
	// A user code.
	UserCode string `json:"userCode"`
	// The result of the end-user authentication and authorization. One of the following. Details are described in the description.
	Result string `json:"result"`
	// The subject (= unique identifier) of the end-user.
	Subject string `json:"subject"`
	// The value of the sub claim that should be used in the ID token.
	Sub *string `json:"sub,omitempty"`
	// The time at which the end-user was authenticated. Its value is the number of seconds from `1970-01-01`.
	AuthTime *int64 `json:"authTime,omitempty"`
	// The reference of the authentication context class which the end-user authentication satisfied.
	Acr *string `json:"acr,omitempty"`
	// Additional claims which will be embedded in the ID token.
	Claims *string `json:"claims,omitempty"`
	// The extra properties associated with the access token.
	Properties []Property `json:"properties,omitempty"`
	// Scopes to replace the scopes specified in the original device authorization request with. When nothing is specified for this parameter, replacement is not performed.
	Scopes []string `json:"scopes,omitempty"`
	// The description of the error. If this optional request parameter is given, its value is used as the value of the `error_description` property, but it is used only when the result is not `AUTHORIZED`. To comply with the specification strictly, the description must not include characters outside the set `%x20-21 / %x23-5B / %x5D-7E`.
	ErrorDescription *string `json:"errorDescription,omitempty"`
	// The URI of a document which describes the error in detail. This corresponds to the `error_uri` property in the response to the client.
	ErrorUri *string `json:"errorUri,omitempty"`
	// JSON that represents additional JWS header parameters for ID tokens.
	IdtHeaderParams *string `json:"idtHeaderParams,omitempty"`
	// the claims that the user has consented for the client application to know.
	ConsentedClaims []string `json:"consentedClaims,omitempty"`
	// Additional claims that are added to the payload part of the JWT access token.
	JwtAtClaims *string `json:"jwtAtClaims,omitempty"`
}

type _DeviceCompleteRequest DeviceCompleteRequest

// NewDeviceCompleteRequest instantiates a new DeviceCompleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceCompleteRequest(userCode string, result string, subject string) *DeviceCompleteRequest {
	this := DeviceCompleteRequest{}
	this.UserCode = userCode
	this.Result = result
	this.Subject = subject
	return &this
}

// NewDeviceCompleteRequestWithDefaults instantiates a new DeviceCompleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceCompleteRequestWithDefaults() *DeviceCompleteRequest {
	this := DeviceCompleteRequest{}
	return &this
}

// GetUserCode returns the UserCode field value
func (o *DeviceCompleteRequest) GetUserCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserCode
}

// GetUserCodeOk returns a tuple with the UserCode field value
// and a boolean to check if the value has been set.
func (o *DeviceCompleteRequest) GetUserCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserCode, true
}

// SetUserCode sets field value
func (o *DeviceCompleteRequest) SetUserCode(v string) {
	o.UserCode = v
}

// GetResult returns the Result field value
func (o *DeviceCompleteRequest) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *DeviceCompleteRequest) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *DeviceCompleteRequest) SetResult(v string) {
	o.Result = v
}

// GetSubject returns the Subject field value
func (o *DeviceCompleteRequest) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *DeviceCompleteRequest) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *DeviceCompleteRequest) SetSubject(v string) {
	o.Subject = v
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *DeviceCompleteRequest) GetSub() string {
	if o == nil || IsNil(o.Sub) {
		var ret string
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCompleteRequest) GetSubOk() (*string, bool) {
	if o == nil || IsNil(o.Sub) {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *DeviceCompleteRequest) HasSub() bool {
	if o != nil && !IsNil(o.Sub) {
		return true
	}

	return false
}

// SetSub gets a reference to the given string and assigns it to the Sub field.
func (o *DeviceCompleteRequest) SetSub(v string) {
	o.Sub = &v
}

// GetAuthTime returns the AuthTime field value if set, zero value otherwise.
func (o *DeviceCompleteRequest) GetAuthTime() int64 {
	if o == nil || IsNil(o.AuthTime) {
		var ret int64
		return ret
	}
	return *o.AuthTime
}

// GetAuthTimeOk returns a tuple with the AuthTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCompleteRequest) GetAuthTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.AuthTime) {
		return nil, false
	}
	return o.AuthTime, true
}

// HasAuthTime returns a boolean if a field has been set.
func (o *DeviceCompleteRequest) HasAuthTime() bool {
	if o != nil && !IsNil(o.AuthTime) {
		return true
	}

	return false
}

// SetAuthTime gets a reference to the given int64 and assigns it to the AuthTime field.
func (o *DeviceCompleteRequest) SetAuthTime(v int64) {
	o.AuthTime = &v
}

// GetAcr returns the Acr field value if set, zero value otherwise.
func (o *DeviceCompleteRequest) GetAcr() string {
	if o == nil || IsNil(o.Acr) {
		var ret string
		return ret
	}
	return *o.Acr
}

// GetAcrOk returns a tuple with the Acr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCompleteRequest) GetAcrOk() (*string, bool) {
	if o == nil || IsNil(o.Acr) {
		return nil, false
	}
	return o.Acr, true
}

// HasAcr returns a boolean if a field has been set.
func (o *DeviceCompleteRequest) HasAcr() bool {
	if o != nil && !IsNil(o.Acr) {
		return true
	}

	return false
}

// SetAcr gets a reference to the given string and assigns it to the Acr field.
func (o *DeviceCompleteRequest) SetAcr(v string) {
	o.Acr = &v
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *DeviceCompleteRequest) GetClaims() string {
	if o == nil || IsNil(o.Claims) {
		var ret string
		return ret
	}
	return *o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCompleteRequest) GetClaimsOk() (*string, bool) {
	if o == nil || IsNil(o.Claims) {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *DeviceCompleteRequest) HasClaims() bool {
	if o != nil && !IsNil(o.Claims) {
		return true
	}

	return false
}

// SetClaims gets a reference to the given string and assigns it to the Claims field.
func (o *DeviceCompleteRequest) SetClaims(v string) {
	o.Claims = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *DeviceCompleteRequest) GetProperties() []Property {
	if o == nil || IsNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCompleteRequest) GetPropertiesOk() ([]Property, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *DeviceCompleteRequest) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *DeviceCompleteRequest) SetProperties(v []Property) {
	o.Properties = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *DeviceCompleteRequest) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCompleteRequest) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *DeviceCompleteRequest) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *DeviceCompleteRequest) SetScopes(v []string) {
	o.Scopes = v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *DeviceCompleteRequest) GetErrorDescription() string {
	if o == nil || IsNil(o.ErrorDescription) {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCompleteRequest) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorDescription) {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *DeviceCompleteRequest) HasErrorDescription() bool {
	if o != nil && !IsNil(o.ErrorDescription) {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *DeviceCompleteRequest) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorUri returns the ErrorUri field value if set, zero value otherwise.
func (o *DeviceCompleteRequest) GetErrorUri() string {
	if o == nil || IsNil(o.ErrorUri) {
		var ret string
		return ret
	}
	return *o.ErrorUri
}

// GetErrorUriOk returns a tuple with the ErrorUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCompleteRequest) GetErrorUriOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorUri) {
		return nil, false
	}
	return o.ErrorUri, true
}

// HasErrorUri returns a boolean if a field has been set.
func (o *DeviceCompleteRequest) HasErrorUri() bool {
	if o != nil && !IsNil(o.ErrorUri) {
		return true
	}

	return false
}

// SetErrorUri gets a reference to the given string and assigns it to the ErrorUri field.
func (o *DeviceCompleteRequest) SetErrorUri(v string) {
	o.ErrorUri = &v
}

// GetIdtHeaderParams returns the IdtHeaderParams field value if set, zero value otherwise.
func (o *DeviceCompleteRequest) GetIdtHeaderParams() string {
	if o == nil || IsNil(o.IdtHeaderParams) {
		var ret string
		return ret
	}
	return *o.IdtHeaderParams
}

// GetIdtHeaderParamsOk returns a tuple with the IdtHeaderParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCompleteRequest) GetIdtHeaderParamsOk() (*string, bool) {
	if o == nil || IsNil(o.IdtHeaderParams) {
		return nil, false
	}
	return o.IdtHeaderParams, true
}

// HasIdtHeaderParams returns a boolean if a field has been set.
func (o *DeviceCompleteRequest) HasIdtHeaderParams() bool {
	if o != nil && !IsNil(o.IdtHeaderParams) {
		return true
	}

	return false
}

// SetIdtHeaderParams gets a reference to the given string and assigns it to the IdtHeaderParams field.
func (o *DeviceCompleteRequest) SetIdtHeaderParams(v string) {
	o.IdtHeaderParams = &v
}

// GetConsentedClaims returns the ConsentedClaims field value if set, zero value otherwise.
func (o *DeviceCompleteRequest) GetConsentedClaims() []string {
	if o == nil || IsNil(o.ConsentedClaims) {
		var ret []string
		return ret
	}
	return o.ConsentedClaims
}

// GetConsentedClaimsOk returns a tuple with the ConsentedClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCompleteRequest) GetConsentedClaimsOk() ([]string, bool) {
	if o == nil || IsNil(o.ConsentedClaims) {
		return nil, false
	}
	return o.ConsentedClaims, true
}

// HasConsentedClaims returns a boolean if a field has been set.
func (o *DeviceCompleteRequest) HasConsentedClaims() bool {
	if o != nil && !IsNil(o.ConsentedClaims) {
		return true
	}

	return false
}

// SetConsentedClaims gets a reference to the given []string and assigns it to the ConsentedClaims field.
func (o *DeviceCompleteRequest) SetConsentedClaims(v []string) {
	o.ConsentedClaims = v
}

// GetJwtAtClaims returns the JwtAtClaims field value if set, zero value otherwise.
func (o *DeviceCompleteRequest) GetJwtAtClaims() string {
	if o == nil || IsNil(o.JwtAtClaims) {
		var ret string
		return ret
	}
	return *o.JwtAtClaims
}

// GetJwtAtClaimsOk returns a tuple with the JwtAtClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCompleteRequest) GetJwtAtClaimsOk() (*string, bool) {
	if o == nil || IsNil(o.JwtAtClaims) {
		return nil, false
	}
	return o.JwtAtClaims, true
}

// HasJwtAtClaims returns a boolean if a field has been set.
func (o *DeviceCompleteRequest) HasJwtAtClaims() bool {
	if o != nil && !IsNil(o.JwtAtClaims) {
		return true
	}

	return false
}

// SetJwtAtClaims gets a reference to the given string and assigns it to the JwtAtClaims field.
func (o *DeviceCompleteRequest) SetJwtAtClaims(v string) {
	o.JwtAtClaims = &v
}

func (o DeviceCompleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceCompleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userCode"] = o.UserCode
	toSerialize["result"] = o.Result
	toSerialize["subject"] = o.Subject
	if !IsNil(o.Sub) {
		toSerialize["sub"] = o.Sub
	}
	if !IsNil(o.AuthTime) {
		toSerialize["authTime"] = o.AuthTime
	}
	if !IsNil(o.Acr) {
		toSerialize["acr"] = o.Acr
	}
	if !IsNil(o.Claims) {
		toSerialize["claims"] = o.Claims
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.ErrorDescription) {
		toSerialize["errorDescription"] = o.ErrorDescription
	}
	if !IsNil(o.ErrorUri) {
		toSerialize["errorUri"] = o.ErrorUri
	}
	if !IsNil(o.IdtHeaderParams) {
		toSerialize["idtHeaderParams"] = o.IdtHeaderParams
	}
	if !IsNil(o.ConsentedClaims) {
		toSerialize["consentedClaims"] = o.ConsentedClaims
	}
	if !IsNil(o.JwtAtClaims) {
		toSerialize["jwtAtClaims"] = o.JwtAtClaims
	}
	return toSerialize, nil
}

func (o *DeviceCompleteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userCode",
		"result",
		"subject",
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

	varDeviceCompleteRequest := _DeviceCompleteRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeviceCompleteRequest)

	if err != nil {
		return err
	}

	*o = DeviceCompleteRequest(varDeviceCompleteRequest)

	return err
}

type NullableDeviceCompleteRequest struct {
	value *DeviceCompleteRequest
	isSet bool
}

func (v NullableDeviceCompleteRequest) Get() *DeviceCompleteRequest {
	return v.value
}

func (v *NullableDeviceCompleteRequest) Set(val *DeviceCompleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceCompleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceCompleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceCompleteRequest(val *DeviceCompleteRequest) *NullableDeviceCompleteRequest {
	return &NullableDeviceCompleteRequest{value: val, isSet: true}
}

func (v NullableDeviceCompleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceCompleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
