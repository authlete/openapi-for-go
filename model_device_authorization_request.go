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

// checks if the DeviceAuthorizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAuthorizationRequest{}

// DeviceAuthorizationRequest struct for DeviceAuthorizationRequest
type DeviceAuthorizationRequest struct {
	// Parameters of a device authorization request which are the request parameters that the device authorization endpoint of the authorization server implementation received from the client application.  The value of `parameters` is the entire entity body (which is formatted in `application/x-www-form-urlencoded`) of the request from the client application.
	Parameters string `json:"parameters"`
	// The client ID extracted from Authorization header of the device authorization request from the client application.  If the device authorization endpoint of the authorization server implementation supports Basic `Authentication` as a means of client authentication, and the request from the client application contained its client ID in `Authorization` header, the value should be extracted and set to this parameter.
	ClientId *string `json:"clientId,omitempty"`
	// The client secret extracted from `Authorization` header of the device authorization request from the client application.  If the device authorization endpoint of the authorization server implementation supports Basic Authentication as a means of client authentication, and the request from the client application contained its client secret in `Authorization` header, the value should be extracted and set to this parameter.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// The client certificate used in the TLS connection between the client application and the device authorization endpoint of the authorization server.
	ClientCertificate *string `json:"clientCertificate,omitempty"`
	// The client certificate path presented by the client during client authentication. Each element is a string in PEM format.
	ClientCertificatePath *string `json:"clientCertificatePath,omitempty"`
}

type _DeviceAuthorizationRequest DeviceAuthorizationRequest

// NewDeviceAuthorizationRequest instantiates a new DeviceAuthorizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthorizationRequest(parameters string) *DeviceAuthorizationRequest {
	this := DeviceAuthorizationRequest{}
	this.Parameters = parameters
	return &this
}

// NewDeviceAuthorizationRequestWithDefaults instantiates a new DeviceAuthorizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthorizationRequestWithDefaults() *DeviceAuthorizationRequest {
	this := DeviceAuthorizationRequest{}
	return &this
}

// GetParameters returns the Parameters field value
func (o *DeviceAuthorizationRequest) GetParameters() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationRequest) GetParametersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *DeviceAuthorizationRequest) SetParameters(v string) {
	o.Parameters = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *DeviceAuthorizationRequest) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationRequest) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *DeviceAuthorizationRequest) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *DeviceAuthorizationRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *DeviceAuthorizationRequest) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *DeviceAuthorizationRequest) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *DeviceAuthorizationRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetClientCertificate returns the ClientCertificate field value if set, zero value otherwise.
func (o *DeviceAuthorizationRequest) GetClientCertificate() string {
	if o == nil || IsNil(o.ClientCertificate) {
		var ret string
		return ret
	}
	return *o.ClientCertificate
}

// GetClientCertificateOk returns a tuple with the ClientCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationRequest) GetClientCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.ClientCertificate) {
		return nil, false
	}
	return o.ClientCertificate, true
}

// HasClientCertificate returns a boolean if a field has been set.
func (o *DeviceAuthorizationRequest) HasClientCertificate() bool {
	if o != nil && !IsNil(o.ClientCertificate) {
		return true
	}

	return false
}

// SetClientCertificate gets a reference to the given string and assigns it to the ClientCertificate field.
func (o *DeviceAuthorizationRequest) SetClientCertificate(v string) {
	o.ClientCertificate = &v
}

// GetClientCertificatePath returns the ClientCertificatePath field value if set, zero value otherwise.
func (o *DeviceAuthorizationRequest) GetClientCertificatePath() string {
	if o == nil || IsNil(o.ClientCertificatePath) {
		var ret string
		return ret
	}
	return *o.ClientCertificatePath
}

// GetClientCertificatePathOk returns a tuple with the ClientCertificatePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationRequest) GetClientCertificatePathOk() (*string, bool) {
	if o == nil || IsNil(o.ClientCertificatePath) {
		return nil, false
	}
	return o.ClientCertificatePath, true
}

// HasClientCertificatePath returns a boolean if a field has been set.
func (o *DeviceAuthorizationRequest) HasClientCertificatePath() bool {
	if o != nil && !IsNil(o.ClientCertificatePath) {
		return true
	}

	return false
}

// SetClientCertificatePath gets a reference to the given string and assigns it to the ClientCertificatePath field.
func (o *DeviceAuthorizationRequest) SetClientCertificatePath(v string) {
	o.ClientCertificatePath = &v
}

func (o DeviceAuthorizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAuthorizationRequest) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

func (o *DeviceAuthorizationRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDeviceAuthorizationRequest := _DeviceAuthorizationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeviceAuthorizationRequest)

	if err != nil {
		return err
	}

	*o = DeviceAuthorizationRequest(varDeviceAuthorizationRequest)

	return err
}

type NullableDeviceAuthorizationRequest struct {
	value *DeviceAuthorizationRequest
	isSet bool
}

func (v NullableDeviceAuthorizationRequest) Get() *DeviceAuthorizationRequest {
	return v.value
}

func (v *NullableDeviceAuthorizationRequest) Set(val *DeviceAuthorizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthorizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthorizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthorizationRequest(val *DeviceAuthorizationRequest) *NullableDeviceAuthorizationRequest {
	return &NullableDeviceAuthorizationRequest{value: val, isSet: true}
}

func (v NullableDeviceAuthorizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthorizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
