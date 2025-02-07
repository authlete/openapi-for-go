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

// checks if the ClientLimited type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientLimited{}

// ClientLimited struct for ClientLimited
type ClientLimited struct {
	// The sequential number of the client. The value of this property is assigned by Authlete.
	Number *int32 `json:"number,omitempty"`
	// The name of the client application. This property corresponds to `client_name` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).
	ClientName *string `json:"clientName,omitempty"`
	// Client names with language tags. If the client application has different names for different languages, this property can be used to register the names.
	ClientNames []TaggedValue `json:"clientNames,omitempty"`
	// The description about the client application.
	Description *string `json:"description,omitempty"`
	// Descriptions about the client application with language tags. If the client application has different descriptions for different languages, this property can be used to register the descriptions.
	Descriptions []TaggedValue `json:"descriptions,omitempty"`
	// The client identifier used in Authlete API calls. The value of this property is assigned by Authlete.
	ClientId *int64 `json:"clientId,omitempty"`
	// The value of the client's `client_id` property used in OAuth and OpenID Connect calls. By default, this is a string version of the `clientId` property.
	ClientIdAlias *string `json:"clientIdAlias,omitempty"`
	// Deprecated. Always set to `true`.
	ClientIdAliasEnabled *bool       `json:"clientIdAliasEnabled,omitempty"`
	ClientType           *ClientType `json:"clientType,omitempty"`
}

// NewClientLimited instantiates a new ClientLimited object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientLimited() *ClientLimited {
	this := ClientLimited{}
	return &this
}

// NewClientLimitedWithDefaults instantiates a new ClientLimited object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientLimitedWithDefaults() *ClientLimited {
	this := ClientLimited{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *ClientLimited) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimited) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *ClientLimited) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *ClientLimited) SetNumber(v int32) {
	o.Number = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *ClientLimited) GetClientName() string {
	if o == nil || IsNil(o.ClientName) {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimited) GetClientNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClientName) {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *ClientLimited) HasClientName() bool {
	if o != nil && !IsNil(o.ClientName) {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *ClientLimited) SetClientName(v string) {
	o.ClientName = &v
}

// GetClientNames returns the ClientNames field value if set, zero value otherwise.
func (o *ClientLimited) GetClientNames() []TaggedValue {
	if o == nil || IsNil(o.ClientNames) {
		var ret []TaggedValue
		return ret
	}
	return o.ClientNames
}

// GetClientNamesOk returns a tuple with the ClientNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimited) GetClientNamesOk() ([]TaggedValue, bool) {
	if o == nil || IsNil(o.ClientNames) {
		return nil, false
	}
	return o.ClientNames, true
}

// HasClientNames returns a boolean if a field has been set.
func (o *ClientLimited) HasClientNames() bool {
	if o != nil && !IsNil(o.ClientNames) {
		return true
	}

	return false
}

// SetClientNames gets a reference to the given []TaggedValue and assigns it to the ClientNames field.
func (o *ClientLimited) SetClientNames(v []TaggedValue) {
	o.ClientNames = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClientLimited) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimited) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClientLimited) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClientLimited) SetDescription(v string) {
	o.Description = &v
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *ClientLimited) GetDescriptions() []TaggedValue {
	if o == nil || IsNil(o.Descriptions) {
		var ret []TaggedValue
		return ret
	}
	return o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimited) GetDescriptionsOk() ([]TaggedValue, bool) {
	if o == nil || IsNil(o.Descriptions) {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *ClientLimited) HasDescriptions() bool {
	if o != nil && !IsNil(o.Descriptions) {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []TaggedValue and assigns it to the Descriptions field.
func (o *ClientLimited) SetDescriptions(v []TaggedValue) {
	o.Descriptions = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ClientLimited) GetClientId() int64 {
	if o == nil || IsNil(o.ClientId) {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimited) GetClientIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ClientLimited) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *ClientLimited) SetClientId(v int64) {
	o.ClientId = &v
}

// GetClientIdAlias returns the ClientIdAlias field value if set, zero value otherwise.
func (o *ClientLimited) GetClientIdAlias() string {
	if o == nil || IsNil(o.ClientIdAlias) {
		var ret string
		return ret
	}
	return *o.ClientIdAlias
}

// GetClientIdAliasOk returns a tuple with the ClientIdAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimited) GetClientIdAliasOk() (*string, bool) {
	if o == nil || IsNil(o.ClientIdAlias) {
		return nil, false
	}
	return o.ClientIdAlias, true
}

// HasClientIdAlias returns a boolean if a field has been set.
func (o *ClientLimited) HasClientIdAlias() bool {
	if o != nil && !IsNil(o.ClientIdAlias) {
		return true
	}

	return false
}

// SetClientIdAlias gets a reference to the given string and assigns it to the ClientIdAlias field.
func (o *ClientLimited) SetClientIdAlias(v string) {
	o.ClientIdAlias = &v
}

// GetClientIdAliasEnabled returns the ClientIdAliasEnabled field value if set, zero value otherwise.
func (o *ClientLimited) GetClientIdAliasEnabled() bool {
	if o == nil || IsNil(o.ClientIdAliasEnabled) {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasEnabled
}

// GetClientIdAliasEnabledOk returns a tuple with the ClientIdAliasEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimited) GetClientIdAliasEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientIdAliasEnabled) {
		return nil, false
	}
	return o.ClientIdAliasEnabled, true
}

// HasClientIdAliasEnabled returns a boolean if a field has been set.
func (o *ClientLimited) HasClientIdAliasEnabled() bool {
	if o != nil && !IsNil(o.ClientIdAliasEnabled) {
		return true
	}

	return false
}

// SetClientIdAliasEnabled gets a reference to the given bool and assigns it to the ClientIdAliasEnabled field.
func (o *ClientLimited) SetClientIdAliasEnabled(v bool) {
	o.ClientIdAliasEnabled = &v
}

// GetClientType returns the ClientType field value if set, zero value otherwise.
func (o *ClientLimited) GetClientType() ClientType {
	if o == nil || IsNil(o.ClientType) {
		var ret ClientType
		return ret
	}
	return *o.ClientType
}

// GetClientTypeOk returns a tuple with the ClientType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientLimited) GetClientTypeOk() (*ClientType, bool) {
	if o == nil || IsNil(o.ClientType) {
		return nil, false
	}
	return o.ClientType, true
}

// HasClientType returns a boolean if a field has been set.
func (o *ClientLimited) HasClientType() bool {
	if o != nil && !IsNil(o.ClientType) {
		return true
	}

	return false
}

// SetClientType gets a reference to the given ClientType and assigns it to the ClientType field.
func (o *ClientLimited) SetClientType(v ClientType) {
	o.ClientType = &v
}

func (o ClientLimited) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientLimited) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.ClientName) {
		toSerialize["clientName"] = o.ClientName
	}
	if !IsNil(o.ClientNames) {
		toSerialize["clientNames"] = o.ClientNames
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Descriptions) {
		toSerialize["descriptions"] = o.Descriptions
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.ClientIdAlias) {
		toSerialize["clientIdAlias"] = o.ClientIdAlias
	}
	if !IsNil(o.ClientIdAliasEnabled) {
		toSerialize["clientIdAliasEnabled"] = o.ClientIdAliasEnabled
	}
	if !IsNil(o.ClientType) {
		toSerialize["clientType"] = o.ClientType
	}
	return toSerialize, nil
}

type NullableClientLimited struct {
	value *ClientLimited
	isSet bool
}

func (v NullableClientLimited) Get() *ClientLimited {
	return v.value
}

func (v *NullableClientLimited) Set(val *ClientLimited) {
	v.value = val
	v.isSet = true
}

func (v NullableClientLimited) IsSet() bool {
	return v.isSet
}

func (v *NullableClientLimited) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientLimited(val *ClientLimited) *NullableClientLimited {
	return &NullableClientLimited{value: val, isSet: true}
}

func (v NullableClientLimited) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientLimited) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
