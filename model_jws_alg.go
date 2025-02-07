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

// JwsAlg The signature algorithm for JWT. This value is represented on 'alg' attribute of the header of JWT.  it's semantics depends upon where is this defined, for instance:   - as service accessTokenSignAlg value, it defines that access token are JWT and the algorithm used to sign it. Check your [KB article](https://kb.authlete.com/en/s/oauth-and-openid-connect/a/jwt-based-access-token).   - as client authorizationSignAlg value, it represents the signature algorithm used when [creating a JARM response](https://kb.authlete.com/en/s/oauth-and-openid-connect/a/enabling-jarm).   - or as client requestSignAlg value, it specifies which is the expected signature used by [client on a Request Object](https://kb.authlete.com/en/s/oauth-and-openid-connect/a/request-objects).
type JwsAlg string

// List of jws_alg
const (
	JWSALG_NONE    JwsAlg = "NONE"
	JWSALG_HS256   JwsAlg = "HS256"
	JWSALG_HS384   JwsAlg = "HS384"
	JWSALG_HS512   JwsAlg = "HS512"
	JWSALG_RS256   JwsAlg = "RS256"
	JWSALG_RS384   JwsAlg = "RS384"
	JWSALG_RS512   JwsAlg = "RS512"
	JWSALG_ES256   JwsAlg = "ES256"
	JWSALG_ES384   JwsAlg = "ES384"
	JWSALG_ES512   JwsAlg = "ES512"
	JWSALG_PS256   JwsAlg = "PS256"
	JWSALG_PS384   JwsAlg = "PS384"
	JWSALG_PS512   JwsAlg = "PS512"
	JWSALG_ES256_K JwsAlg = "ES256K"
	JWSALG_ED_DSA  JwsAlg = "EdDSA"
)

// All allowed values of JwsAlg enum
var AllowedJwsAlgEnumValues = []JwsAlg{
	"NONE",
	"HS256",
	"HS384",
	"HS512",
	"RS256",
	"RS384",
	"RS512",
	"ES256",
	"ES384",
	"ES512",
	"PS256",
	"PS384",
	"PS512",
	"ES256K",
	"EdDSA",
}

func (v *JwsAlg) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JwsAlg(value)
	for _, existing := range AllowedJwsAlgEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JwsAlg", value)
}

// NewJwsAlgFromValue returns a pointer to a valid JwsAlg
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJwsAlgFromValue(v string) (*JwsAlg, error) {
	ev := JwsAlg(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JwsAlg: valid values are %v", v, AllowedJwsAlgEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JwsAlg) IsValid() bool {
	for _, existing := range AllowedJwsAlgEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to jws_alg value
func (v JwsAlg) Ptr() *JwsAlg {
	return &v
}

type NullableJwsAlg struct {
	value *JwsAlg
	isSet bool
}

func (v NullableJwsAlg) Get() *JwsAlg {
	return v.value
}

func (v *NullableJwsAlg) Set(val *JwsAlg) {
	v.value = val
	v.isSet = true
}

func (v NullableJwsAlg) IsSet() bool {
	return v.isSet
}

func (v *NullableJwsAlg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJwsAlg(val *JwsAlg) *NullableJwsAlg {
	return &NullableJwsAlg{value: val, isSet: true}
}

func (v NullableJwsAlg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJwsAlg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
