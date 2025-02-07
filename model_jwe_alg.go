/*
Authlete API Explorer

<div class=\"min-h-screen bg-gray-100 dark:bg-gray-900 text-gray-900 dark:text-gray-100 p-6\">   <div class=\"flex justify-end mb-4\">     <label for=\"theme-toggle\" class=\"flex items-center cursor-pointer\">       <div class=\"relative\">Dark mode:         <input type=\"checkbox\" id=\"theme-toggle\" class=\"sr-only\" onchange=\"toggleTheme()\">         <div class=\"block bg-gray-600 w-14 h-8 rounded-full\"></div>         <div class=\"dot absolute left-1 top-1 bg-white w-6 h-6 rounded-full transition\"></div>       </div>     </label>   </div>   <header class=\"bg-green-500 dark:bg-green-700 p-4 rounded-lg text-white text-center\">     <p>       Welcome to the <strong>Authlete API documentation</strong>. Authlete is an <strong>API-first service</strong>       where every aspect of the platform is configurable via API. This explorer provides a convenient way to       authenticate and interact with the API, allowing you to see Authlete in action quickly. 🚀     </p>     <p>       At a high level, the Authlete API is grouped into two categories:     </p>     <ul class=\"list-disc list-inside\">       <li><strong>Management APIs</strong>: Enable you to manage services and clients. 🔧</li>       <li><strong>Runtime APIs</strong>: Allow you to build your own Authorization Servers or Verifiable Credential (VC)         issuers. 🔐</li>     </ul>     <p>All API endpoints are secured using access tokens issued by Authlete's Identity Provider (IdP). If you already       have an Authlete account, simply use the <em>Get Token</em> option on the Authentication page to log in and obtain       an access token for API usage. If you don't have an account yet, <a href=\"https://console.authlete.com/register\">sign up         here</a> to get started.</p>   </header>   <main>     <section id=\"api-servers\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🌐 API Servers</h2>       <p>Authlete is a global service with clusters available in multiple regions across the world.</p>       <p>Currently, our service is available in the following regions:</p>       <div class=\"grid grid-cols-2 gap-4\">         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇺🇸 US</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇯🇵 JP</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇪🇺 EU</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇧🇷 Brazil</p>         </div>       </div>       <p>Our customers can host their data in the region that best meets their requirements.</p>       <a href=\"#servers\" class=\"block mt-4 text-green-500 dark:text-green-300 hover:underline text-center\">Select your         preferred server</a>     </section>     <section id=\"authentication\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🔑 Authentication</h2>       <p>The API Explorer requires an access token to call the API.</p>       <p>You can create the access token from the <a href=\"https://console.authlete.com\">Authlete Management Console</a> and set it in the HTTP Bearer section of Authentication page.</p>       <p>Alternatively, if you have an Authlete account, the API Explorer can log you in with your Authlete account and         automatically acquire the required access token.</p>       <div class=\"theme-admonition theme-admonition-warning admonition_o5H7 alert alert--warning\">         <div class=\"admonitionContent_Knsx\">           <p>⚠️ <strong>Important Note:</strong> When the API Explorer acquires the token after login, the access tokens             will have the same permissions as the user who logs in as part of this flow.</p>         </div>       </div>       <a href=\"#auth\" class=\"block mt-4 text-green-500 dark:text-green-300 hover:underline text-center\">Setup your         access token</a>     </section>     <section id=\"tutorials\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🎓 Tutorials</h2>       <p>If you have successfully tested the API from the API Console and want to take the next step of integrating the         API into your application, or if you want to see a sample using Authlete APIs, follow the links below. These         resources will help you understand key concepts and how to integrate Authlete API into your applications.</p>       <div class=\"mt-4\">         <a href=\"https://www.authlete.com/developers/getting_started/\"           class=\"block text-green-500 dark:text-green-300 font-bold hover:underline mb-2\">🚀 Getting Started with           Authlete</a>           </br>         <a href=\"https://www.authlete.com/developers/tutorial/signup/\"           class=\"block text-green-500 dark:text-green-300 font-bold hover:underline\">🔑 From Sign-Up to the First API           Request</a>       </div>     </section>     <section id=\"support\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🛠 Contact Us</h2>       <p>If you have any questions or need assistance, our team is here to help.</p>       <a href=\"https://www.authlete.com/contact/\"         class=\"block mt-4 text-green-500 dark:text-green-300 font-bold hover:underline\">Contact Page</a>     </section>   </main> </div>

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
	"fmt"
)

// JweAlg this is the 'alg' header value for encrypted JWT tokens. Depending upon the context, this refers to key transport scheme to be used by the client and by the server. For instance: - as `authorizationEncryptionAlg` value, it refers to the encoding algorithm used by server for transporting they keys on JARM objects - as `requestEncryptionAlg` value, it refers to the expected key transport encoding algorithm that server expect from client when encrypting a Request Object - as `idTokenEncryptionAlg` value, it refers to the algorithm used by the server to key transport of id_tokens  **Please note that some of the algorithms are more secure than others, some are not supported very well cross platforms and some (like RSA1_5) is known to be weak**.
type JweAlg string

// List of jwe_alg
const (
	JWEALG_RSA1_5              JweAlg = "RSA1_5"
	JWEALG_RSA_OAEP            JweAlg = "RSA_OAEP"
	JWEALG_RSA_OAEP_256        JweAlg = "RSA_OAEP_256"
	JWEALG_A128_KW             JweAlg = "A128KW"
	JWEALG_A192_KW             JweAlg = "A192KW"
	JWEALG_A256_KW             JweAlg = "A256KW"
	JWEALG_DIR                 JweAlg = "DIR"
	JWEALG_ECDH_ES             JweAlg = "ECDH_ES"
	JWEALG_ECDH_ES_A128_KW     JweAlg = "ECDH_ES_A128KW"
	JWEALG_ECDH_ES_A192_KW     JweAlg = "ECDH_ES_A192KW"
	JWEALG_ECDH_ES_A256_KW     JweAlg = "ECDH_ES_A256KW"
	JWEALG_A128_GCMKW          JweAlg = "A128GCMKW"
	JWEALG_A192_GCMKW          JweAlg = "A192GCMKW"
	JWEALG_A256_GCMKW          JweAlg = "A256GCMKW"
	JWEALG_PBES2_HS256_A128_KW JweAlg = "PBES2_HS256_A128KW"
	JWEALG_PBES2_HS384_A192_KW JweAlg = "PBES2_HS384_A192KW"
	JWEALG_PBES2_HS512_A256_KW JweAlg = "PBES2_HS512_A256KW"
)

// All allowed values of JweAlg enum
var AllowedJweAlgEnumValues = []JweAlg{
	"RSA1_5",
	"RSA_OAEP",
	"RSA_OAEP_256",
	"A128KW",
	"A192KW",
	"A256KW",
	"DIR",
	"ECDH_ES",
	"ECDH_ES_A128KW",
	"ECDH_ES_A192KW",
	"ECDH_ES_A256KW",
	"A128GCMKW",
	"A192GCMKW",
	"A256GCMKW",
	"PBES2_HS256_A128KW",
	"PBES2_HS384_A192KW",
	"PBES2_HS512_A256KW",
}

func (v *JweAlg) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JweAlg(value)
	for _, existing := range AllowedJweAlgEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JweAlg", value)
}

// NewJweAlgFromValue returns a pointer to a valid JweAlg
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJweAlgFromValue(v string) (*JweAlg, error) {
	ev := JweAlg(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JweAlg: valid values are %v", v, AllowedJweAlgEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JweAlg) IsValid() bool {
	for _, existing := range AllowedJweAlgEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to jwe_alg value
func (v JweAlg) Ptr() *JweAlg {
	return &v
}

type NullableJweAlg struct {
	value *JweAlg
	isSet bool
}

func (v NullableJweAlg) Get() *JweAlg {
	return v.value
}

func (v *NullableJweAlg) Set(val *JweAlg) {
	v.value = val
	v.isSet = true
}

func (v NullableJweAlg) IsSet() bool {
	return v.isSet
}

func (v *NullableJweAlg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJweAlg(val *JweAlg) *NullableJweAlg {
	return &NullableJweAlg{value: val, isSet: true}
}

func (v NullableJweAlg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJweAlg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
