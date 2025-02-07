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

// checks if the BackchannelAuthenticationIssueResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackchannelAuthenticationIssueResponse{}

// BackchannelAuthenticationIssueResponse struct for BackchannelAuthenticationIssueResponse
type BackchannelAuthenticationIssueResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take.
	Action *string `json:"action,omitempty"`
	// The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of `action` parameter.
	ResponseContent *string `json:"responseContent,omitempty"`
	// The newly issued authentication request ID.
	AuthReqId *string `json:"authReqId,omitempty"`
	// The duration of the issued authentication request ID in seconds.
	ExpiresIn *int32 `json:"expiresIn,omitempty"`
	// The minimum amount of time in seconds that the client must wait for between polling requests to the token endpoint.
	Interval *int32 `json:"interval,omitempty"`
}

// NewBackchannelAuthenticationIssueResponse instantiates a new BackchannelAuthenticationIssueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackchannelAuthenticationIssueResponse() *BackchannelAuthenticationIssueResponse {
	this := BackchannelAuthenticationIssueResponse{}
	return &this
}

// NewBackchannelAuthenticationIssueResponseWithDefaults instantiates a new BackchannelAuthenticationIssueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackchannelAuthenticationIssueResponseWithDefaults() *BackchannelAuthenticationIssueResponse {
	this := BackchannelAuthenticationIssueResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *BackchannelAuthenticationIssueResponse) GetResultCode() string {
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationIssueResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *BackchannelAuthenticationIssueResponse) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *BackchannelAuthenticationIssueResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *BackchannelAuthenticationIssueResponse) GetResultMessage() string {
	if o == nil || IsNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationIssueResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *BackchannelAuthenticationIssueResponse) HasResultMessage() bool {
	if o != nil && !IsNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *BackchannelAuthenticationIssueResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BackchannelAuthenticationIssueResponse) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationIssueResponse) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BackchannelAuthenticationIssueResponse) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BackchannelAuthenticationIssueResponse) SetAction(v string) {
	o.Action = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *BackchannelAuthenticationIssueResponse) GetResponseContent() string {
	if o == nil || IsNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationIssueResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *BackchannelAuthenticationIssueResponse) HasResponseContent() bool {
	if o != nil && !IsNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *BackchannelAuthenticationIssueResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetAuthReqId returns the AuthReqId field value if set, zero value otherwise.
func (o *BackchannelAuthenticationIssueResponse) GetAuthReqId() string {
	if o == nil || IsNil(o.AuthReqId) {
		var ret string
		return ret
	}
	return *o.AuthReqId
}

// GetAuthReqIdOk returns a tuple with the AuthReqId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationIssueResponse) GetAuthReqIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuthReqId) {
		return nil, false
	}
	return o.AuthReqId, true
}

// HasAuthReqId returns a boolean if a field has been set.
func (o *BackchannelAuthenticationIssueResponse) HasAuthReqId() bool {
	if o != nil && !IsNil(o.AuthReqId) {
		return true
	}

	return false
}

// SetAuthReqId gets a reference to the given string and assigns it to the AuthReqId field.
func (o *BackchannelAuthenticationIssueResponse) SetAuthReqId(v string) {
	o.AuthReqId = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *BackchannelAuthenticationIssueResponse) GetExpiresIn() int32 {
	if o == nil || IsNil(o.ExpiresIn) {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationIssueResponse) GetExpiresInOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *BackchannelAuthenticationIssueResponse) HasExpiresIn() bool {
	if o != nil && !IsNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *BackchannelAuthenticationIssueResponse) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *BackchannelAuthenticationIssueResponse) GetInterval() int32 {
	if o == nil || IsNil(o.Interval) {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackchannelAuthenticationIssueResponse) GetIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *BackchannelAuthenticationIssueResponse) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *BackchannelAuthenticationIssueResponse) SetInterval(v int32) {
	o.Interval = &v
}

func (o BackchannelAuthenticationIssueResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackchannelAuthenticationIssueResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultCode) {
		toSerialize["resultCode"] = o.ResultCode
	}
	if !IsNil(o.ResultMessage) {
		toSerialize["resultMessage"] = o.ResultMessage
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.ResponseContent) {
		toSerialize["responseContent"] = o.ResponseContent
	}
	if !IsNil(o.AuthReqId) {
		toSerialize["authReqId"] = o.AuthReqId
	}
	if !IsNil(o.ExpiresIn) {
		toSerialize["expiresIn"] = o.ExpiresIn
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	return toSerialize, nil
}

type NullableBackchannelAuthenticationIssueResponse struct {
	value *BackchannelAuthenticationIssueResponse
	isSet bool
}

func (v NullableBackchannelAuthenticationIssueResponse) Get() *BackchannelAuthenticationIssueResponse {
	return v.value
}

func (v *NullableBackchannelAuthenticationIssueResponse) Set(val *BackchannelAuthenticationIssueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBackchannelAuthenticationIssueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBackchannelAuthenticationIssueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackchannelAuthenticationIssueResponse(val *BackchannelAuthenticationIssueResponse) *NullableBackchannelAuthenticationIssueResponse {
	return &NullableBackchannelAuthenticationIssueResponse{value: val, isSet: true}
}

func (v NullableBackchannelAuthenticationIssueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackchannelAuthenticationIssueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
