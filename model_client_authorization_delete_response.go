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

// checks if the ClientAuthorizationDeleteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientAuthorizationDeleteResponse{}

// ClientAuthorizationDeleteResponse struct for ClientAuthorizationDeleteResponse
type ClientAuthorizationDeleteResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// A short message which explains the result of the API call.
	ServiceApiKey *int64 `json:"serviceApiKey,omitempty"`
	// Get the client ID.
	ClientId *int64 `json:"clientId,omitempty"`
	// Get the subject (= unique identifier) of the user who has granted authorization to the client.
	Subject *string `json:"subject,omitempty"`
	// Get the scopes granted to the client application by the last authorization process by the user (who is identified by the subject).
	LatestGrantedScopes []string `json:"latestGrantedScopes,omitempty"`
	// Get the scopes granted to the client application by all the past authorization processes. Note that revoked scopes are not included.
	MergedGrantedScopes []string `json:"mergedGrantedScopes,omitempty"`
	// Get the timestamp in milliseconds since Unix epoch at which this record was modified.
	ModifiedAt *int64 `json:"modifiedAt,omitempty"`
}

// NewClientAuthorizationDeleteResponse instantiates a new ClientAuthorizationDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientAuthorizationDeleteResponse() *ClientAuthorizationDeleteResponse {
	this := ClientAuthorizationDeleteResponse{}
	return &this
}

// NewClientAuthorizationDeleteResponseWithDefaults instantiates a new ClientAuthorizationDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientAuthorizationDeleteResponseWithDefaults() *ClientAuthorizationDeleteResponse {
	this := ClientAuthorizationDeleteResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *ClientAuthorizationDeleteResponse) GetResultCode() string {
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuthorizationDeleteResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *ClientAuthorizationDeleteResponse) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *ClientAuthorizationDeleteResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *ClientAuthorizationDeleteResponse) GetResultMessage() string {
	if o == nil || IsNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuthorizationDeleteResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *ClientAuthorizationDeleteResponse) HasResultMessage() bool {
	if o != nil && !IsNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *ClientAuthorizationDeleteResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetServiceApiKey returns the ServiceApiKey field value if set, zero value otherwise.
func (o *ClientAuthorizationDeleteResponse) GetServiceApiKey() int64 {
	if o == nil || IsNil(o.ServiceApiKey) {
		var ret int64
		return ret
	}
	return *o.ServiceApiKey
}

// GetServiceApiKeyOk returns a tuple with the ServiceApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuthorizationDeleteResponse) GetServiceApiKeyOk() (*int64, bool) {
	if o == nil || IsNil(o.ServiceApiKey) {
		return nil, false
	}
	return o.ServiceApiKey, true
}

// HasServiceApiKey returns a boolean if a field has been set.
func (o *ClientAuthorizationDeleteResponse) HasServiceApiKey() bool {
	if o != nil && !IsNil(o.ServiceApiKey) {
		return true
	}

	return false
}

// SetServiceApiKey gets a reference to the given int64 and assigns it to the ServiceApiKey field.
func (o *ClientAuthorizationDeleteResponse) SetServiceApiKey(v int64) {
	o.ServiceApiKey = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ClientAuthorizationDeleteResponse) GetClientId() int64 {
	if o == nil || IsNil(o.ClientId) {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuthorizationDeleteResponse) GetClientIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ClientAuthorizationDeleteResponse) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *ClientAuthorizationDeleteResponse) SetClientId(v int64) {
	o.ClientId = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *ClientAuthorizationDeleteResponse) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuthorizationDeleteResponse) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *ClientAuthorizationDeleteResponse) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *ClientAuthorizationDeleteResponse) SetSubject(v string) {
	o.Subject = &v
}

// GetLatestGrantedScopes returns the LatestGrantedScopes field value if set, zero value otherwise.
func (o *ClientAuthorizationDeleteResponse) GetLatestGrantedScopes() []string {
	if o == nil || IsNil(o.LatestGrantedScopes) {
		var ret []string
		return ret
	}
	return o.LatestGrantedScopes
}

// GetLatestGrantedScopesOk returns a tuple with the LatestGrantedScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuthorizationDeleteResponse) GetLatestGrantedScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.LatestGrantedScopes) {
		return nil, false
	}
	return o.LatestGrantedScopes, true
}

// HasLatestGrantedScopes returns a boolean if a field has been set.
func (o *ClientAuthorizationDeleteResponse) HasLatestGrantedScopes() bool {
	if o != nil && !IsNil(o.LatestGrantedScopes) {
		return true
	}

	return false
}

// SetLatestGrantedScopes gets a reference to the given []string and assigns it to the LatestGrantedScopes field.
func (o *ClientAuthorizationDeleteResponse) SetLatestGrantedScopes(v []string) {
	o.LatestGrantedScopes = v
}

// GetMergedGrantedScopes returns the MergedGrantedScopes field value if set, zero value otherwise.
func (o *ClientAuthorizationDeleteResponse) GetMergedGrantedScopes() []string {
	if o == nil || IsNil(o.MergedGrantedScopes) {
		var ret []string
		return ret
	}
	return o.MergedGrantedScopes
}

// GetMergedGrantedScopesOk returns a tuple with the MergedGrantedScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuthorizationDeleteResponse) GetMergedGrantedScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.MergedGrantedScopes) {
		return nil, false
	}
	return o.MergedGrantedScopes, true
}

// HasMergedGrantedScopes returns a boolean if a field has been set.
func (o *ClientAuthorizationDeleteResponse) HasMergedGrantedScopes() bool {
	if o != nil && !IsNil(o.MergedGrantedScopes) {
		return true
	}

	return false
}

// SetMergedGrantedScopes gets a reference to the given []string and assigns it to the MergedGrantedScopes field.
func (o *ClientAuthorizationDeleteResponse) SetMergedGrantedScopes(v []string) {
	o.MergedGrantedScopes = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *ClientAuthorizationDeleteResponse) GetModifiedAt() int64 {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret int64
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuthorizationDeleteResponse) GetModifiedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *ClientAuthorizationDeleteResponse) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given int64 and assigns it to the ModifiedAt field.
func (o *ClientAuthorizationDeleteResponse) SetModifiedAt(v int64) {
	o.ModifiedAt = &v
}

func (o ClientAuthorizationDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientAuthorizationDeleteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultCode) {
		toSerialize["resultCode"] = o.ResultCode
	}
	if !IsNil(o.ResultMessage) {
		toSerialize["resultMessage"] = o.ResultMessage
	}
	if !IsNil(o.ServiceApiKey) {
		toSerialize["serviceApiKey"] = o.ServiceApiKey
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.LatestGrantedScopes) {
		toSerialize["latestGrantedScopes"] = o.LatestGrantedScopes
	}
	if !IsNil(o.MergedGrantedScopes) {
		toSerialize["mergedGrantedScopes"] = o.MergedGrantedScopes
	}
	if !IsNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	return toSerialize, nil
}

type NullableClientAuthorizationDeleteResponse struct {
	value *ClientAuthorizationDeleteResponse
	isSet bool
}

func (v NullableClientAuthorizationDeleteResponse) Get() *ClientAuthorizationDeleteResponse {
	return v.value
}

func (v *NullableClientAuthorizationDeleteResponse) Set(val *ClientAuthorizationDeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClientAuthorizationDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClientAuthorizationDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientAuthorizationDeleteResponse(val *ClientAuthorizationDeleteResponse) *NullableClientAuthorizationDeleteResponse {
	return &NullableClientAuthorizationDeleteResponse{value: val, isSet: true}
}

func (v NullableClientAuthorizationDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientAuthorizationDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
