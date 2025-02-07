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

// checks if the UserinfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserinfoResponse{}

// UserinfoResponse struct for UserinfoResponse
type UserinfoResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take.
	Action *string `json:"action,omitempty"`
	// The list of claims that the client application requests to be embedded in the ID token.
	Claims []string `json:"claims,omitempty"`
	// The ID of the client application which is associated with the access token.
	ClientId *int64 `json:"clientId,omitempty"`
	// The client ID alias when the authorization request for the access token was made.
	ClientIdAlias *string `json:"clientIdAlias,omitempty"`
	// The flag which indicates whether the client ID alias was used when the authorization request for the access token was made.
	ClientIdAliasUsed *bool `json:"clientIdAliasUsed,omitempty"`
	// The content that the authorization server implementation can use as the value of `WWW-Authenticate` header on errors.
	ResponseContent *string `json:"responseContent,omitempty"`
	// The scopes covered by the access token.
	Scopes []string `json:"scopes,omitempty"`
	// The subject (= resource owner's ID).
	Subject *string `json:"subject,omitempty"`
	// The access token that came along with the userinfo request.
	Token *string `json:"token,omitempty"`
	// The extra properties associated with the access token.
	Properties []Property `json:"properties,omitempty"`
	// The value of the `userinfo` property in the `claims` request parameter or in the `claims` property in an authorization request object.  A client application may request certain claims be embedded in an ID token or in a response from the userInfo endpoint. There are several ways. Including the `claims` request parameter and including the `claims` property in a request object are such examples. In both cases, the value of the `claims` parameter/property is JSON. Its format is described in [5.5. Requesting Claims using the \"claims\" Request Parameter](https://openid.net/specs/openid-connect-core-1_0.html#ClaimsParameter).  The following is an excerpt from the specification. You can find `userinfo` and `id_token` are top-level properties.  ```json {   \"userinfo\":   {     \"given_name\": { \"essential\": true },     \"nickname\": null,     \"email\": { \"essential\": true },     \"email_verified\": { \"essential\": true },     \"picture\": null,     \"http://example.info/claims/groups\": null   },   \"id_token\":   {     \"auth_time\": { \"essential\": true },     \"acr\": { \"values\": [ \"urn:mace:incommon:iap:silver\" ] }   } } ````  The value of this property is the value of the `userinfo` property in JSON format. For example, if the JSON above is included in an authorization request, this property holds JSON equivalent to the following.  ```json {   \"given_name\": { \"essential\": true },   \"nickname\": null,   \"email\": { \"essential\": true },   \"email_verified\": { \"essential\": true },   \"picture\": null,   \"http://example.info/claims/groups\": null } ```  Note that if a request object is given and it contains the `claims` property and if the `claims` request parameter is also given, the value of this property holds the former value.
	UserInfoClaims *string `json:"userInfoClaims,omitempty"`
	// The attributes of this service that the client application belongs to.
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`
	// The attributes of the client.
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`
	// the claims that the user has consented for the client application to know.
	ConsentedClaims []string `json:"consentedClaims,omitempty"`
	// Get names of claims that are requested indirectly by <i>\"transformed claims\"</i>.  <p> A client application can request <i>\"transformed claims\"</i> by adding names of transformed claims in the `claims` request parameter. The following is an example of the `claims` request parameter that requests a predefined transformed claim named `18_or_over` and a transformed claim named `nationality_usa` to be embedded in the response from the userinfo endpoint. </p>  ```json {   \"transformed_claims\": {     \"nationality_usa\": {       \"claim\": \"nationalities\",       \"fn\": [         [ \"eq\", \"USA\" ],         \"any\"       ]     }   },   \"userinfo\": {     \"::18_or_over\": null,     \":nationality_usa\": null   } } ```  The example above assumes that a transformed claim named `18_or_over` is predefined by the authorization server like below.  ```json {   \"18_or_over\": {     \"claim\": \"birthdate\",     \"fn\": [       \"years_ago\",       [ \"gte\", 18 ]     ]   } } ```  In the example, the `nationalities` claim is requested indirectly by the `nationality_usa` transformed claim. Likewise, the `birthdate` claim is requested indirectly by the `18_or_over` transformed claim.  When the `claims` request parameter of an authorization request is like the example above, this `requestedClaimsForTx` property will hold the following value.  ```json [ \"birthdate\", \"nationalities\" ] ```  It is expected that the authorization server implementation prepares values of the listed claims and passes them as the value of the `claimsForTx` request parameter when it calls the `/api/auth/userinfo/issue` API. The following is an example of the value of the `claimsForTx` request parameter.  ```json {   \"birthdate\": \"1970-01-23\",   \"nationalities\": [ \"DEU\", \"USA\" ] } ```
	RequestedClaimsForTx []string `json:"requestedClaimsForTx,omitempty"`
	// Names of verified claims that will be referenced when transformed claims are computed.
	RequestedVerifiedClaimsForTx [][]string `json:"requestedVerifiedClaimsForTx,omitempty"`
	// the value of the `transformed_claims` property in the `claims` request parameter of an authorization request or in the `claims` property in a request object.
	TransformedClaims *string `json:"transformedClaims,omitempty"`
	// The entity ID of the client.
	ClientEntityId *string `json:"clientEntityId,omitempty"`
	// Flag which indicates whether the entity ID of the client was used when the request for the access token was made.
	ClientEntityIdUsed *bool `json:"clientEntityIdUsed,omitempty"`
	// The expected nonce value for DPoP proof JWT, which should be used as the value of the `DPoP-Nonce` HTTP header.
	DpopNonce *string `json:"dpopNonce,omitempty"`
}

// NewUserinfoResponse instantiates a new UserinfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserinfoResponse() *UserinfoResponse {
	this := UserinfoResponse{}
	return &this
}

// NewUserinfoResponseWithDefaults instantiates a new UserinfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserinfoResponseWithDefaults() *UserinfoResponse {
	this := UserinfoResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *UserinfoResponse) GetResultCode() string {
	if o == nil || IsNil(o.ResultCode) {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *UserinfoResponse) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *UserinfoResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *UserinfoResponse) GetResultMessage() string {
	if o == nil || IsNil(o.ResultMessage) {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ResultMessage) {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *UserinfoResponse) HasResultMessage() bool {
	if o != nil && !IsNil(o.ResultMessage) {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *UserinfoResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *UserinfoResponse) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *UserinfoResponse) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *UserinfoResponse) SetAction(v string) {
	o.Action = &v
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *UserinfoResponse) GetClaims() []string {
	if o == nil || IsNil(o.Claims) {
		var ret []string
		return ret
	}
	return o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetClaimsOk() ([]string, bool) {
	if o == nil || IsNil(o.Claims) {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *UserinfoResponse) HasClaims() bool {
	if o != nil && !IsNil(o.Claims) {
		return true
	}

	return false
}

// SetClaims gets a reference to the given []string and assigns it to the Claims field.
func (o *UserinfoResponse) SetClaims(v []string) {
	o.Claims = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *UserinfoResponse) GetClientId() int64 {
	if o == nil || IsNil(o.ClientId) {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetClientIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *UserinfoResponse) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *UserinfoResponse) SetClientId(v int64) {
	o.ClientId = &v
}

// GetClientIdAlias returns the ClientIdAlias field value if set, zero value otherwise.
func (o *UserinfoResponse) GetClientIdAlias() string {
	if o == nil || IsNil(o.ClientIdAlias) {
		var ret string
		return ret
	}
	return *o.ClientIdAlias
}

// GetClientIdAliasOk returns a tuple with the ClientIdAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetClientIdAliasOk() (*string, bool) {
	if o == nil || IsNil(o.ClientIdAlias) {
		return nil, false
	}
	return o.ClientIdAlias, true
}

// HasClientIdAlias returns a boolean if a field has been set.
func (o *UserinfoResponse) HasClientIdAlias() bool {
	if o != nil && !IsNil(o.ClientIdAlias) {
		return true
	}

	return false
}

// SetClientIdAlias gets a reference to the given string and assigns it to the ClientIdAlias field.
func (o *UserinfoResponse) SetClientIdAlias(v string) {
	o.ClientIdAlias = &v
}

// GetClientIdAliasUsed returns the ClientIdAliasUsed field value if set, zero value otherwise.
func (o *UserinfoResponse) GetClientIdAliasUsed() bool {
	if o == nil || IsNil(o.ClientIdAliasUsed) {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasUsed
}

// GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetClientIdAliasUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientIdAliasUsed) {
		return nil, false
	}
	return o.ClientIdAliasUsed, true
}

// HasClientIdAliasUsed returns a boolean if a field has been set.
func (o *UserinfoResponse) HasClientIdAliasUsed() bool {
	if o != nil && !IsNil(o.ClientIdAliasUsed) {
		return true
	}

	return false
}

// SetClientIdAliasUsed gets a reference to the given bool and assigns it to the ClientIdAliasUsed field.
func (o *UserinfoResponse) SetClientIdAliasUsed(v bool) {
	o.ClientIdAliasUsed = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *UserinfoResponse) GetResponseContent() string {
	if o == nil || IsNil(o.ResponseContent) {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseContent) {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *UserinfoResponse) HasResponseContent() bool {
	if o != nil && !IsNil(o.ResponseContent) {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *UserinfoResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *UserinfoResponse) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *UserinfoResponse) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *UserinfoResponse) SetScopes(v []string) {
	o.Scopes = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *UserinfoResponse) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *UserinfoResponse) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *UserinfoResponse) SetSubject(v string) {
	o.Subject = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UserinfoResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UserinfoResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UserinfoResponse) SetToken(v string) {
	o.Token = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *UserinfoResponse) GetProperties() []Property {
	if o == nil || IsNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetPropertiesOk() ([]Property, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *UserinfoResponse) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *UserinfoResponse) SetProperties(v []Property) {
	o.Properties = v
}

// GetUserInfoClaims returns the UserInfoClaims field value if set, zero value otherwise.
func (o *UserinfoResponse) GetUserInfoClaims() string {
	if o == nil || IsNil(o.UserInfoClaims) {
		var ret string
		return ret
	}
	return *o.UserInfoClaims
}

// GetUserInfoClaimsOk returns a tuple with the UserInfoClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetUserInfoClaimsOk() (*string, bool) {
	if o == nil || IsNil(o.UserInfoClaims) {
		return nil, false
	}
	return o.UserInfoClaims, true
}

// HasUserInfoClaims returns a boolean if a field has been set.
func (o *UserinfoResponse) HasUserInfoClaims() bool {
	if o != nil && !IsNil(o.UserInfoClaims) {
		return true
	}

	return false
}

// SetUserInfoClaims gets a reference to the given string and assigns it to the UserInfoClaims field.
func (o *UserinfoResponse) SetUserInfoClaims(v string) {
	o.UserInfoClaims = &v
}

// GetServiceAttributes returns the ServiceAttributes field value if set, zero value otherwise.
func (o *UserinfoResponse) GetServiceAttributes() []Pair {
	if o == nil || IsNil(o.ServiceAttributes) {
		var ret []Pair
		return ret
	}
	return o.ServiceAttributes
}

// GetServiceAttributesOk returns a tuple with the ServiceAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetServiceAttributesOk() ([]Pair, bool) {
	if o == nil || IsNil(o.ServiceAttributes) {
		return nil, false
	}
	return o.ServiceAttributes, true
}

// HasServiceAttributes returns a boolean if a field has been set.
func (o *UserinfoResponse) HasServiceAttributes() bool {
	if o != nil && !IsNil(o.ServiceAttributes) {
		return true
	}

	return false
}

// SetServiceAttributes gets a reference to the given []Pair and assigns it to the ServiceAttributes field.
func (o *UserinfoResponse) SetServiceAttributes(v []Pair) {
	o.ServiceAttributes = v
}

// GetClientAttributes returns the ClientAttributes field value if set, zero value otherwise.
func (o *UserinfoResponse) GetClientAttributes() []Pair {
	if o == nil || IsNil(o.ClientAttributes) {
		var ret []Pair
		return ret
	}
	return o.ClientAttributes
}

// GetClientAttributesOk returns a tuple with the ClientAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetClientAttributesOk() ([]Pair, bool) {
	if o == nil || IsNil(o.ClientAttributes) {
		return nil, false
	}
	return o.ClientAttributes, true
}

// HasClientAttributes returns a boolean if a field has been set.
func (o *UserinfoResponse) HasClientAttributes() bool {
	if o != nil && !IsNil(o.ClientAttributes) {
		return true
	}

	return false
}

// SetClientAttributes gets a reference to the given []Pair and assigns it to the ClientAttributes field.
func (o *UserinfoResponse) SetClientAttributes(v []Pair) {
	o.ClientAttributes = v
}

// GetConsentedClaims returns the ConsentedClaims field value if set, zero value otherwise.
func (o *UserinfoResponse) GetConsentedClaims() []string {
	if o == nil || IsNil(o.ConsentedClaims) {
		var ret []string
		return ret
	}
	return o.ConsentedClaims
}

// GetConsentedClaimsOk returns a tuple with the ConsentedClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetConsentedClaimsOk() ([]string, bool) {
	if o == nil || IsNil(o.ConsentedClaims) {
		return nil, false
	}
	return o.ConsentedClaims, true
}

// HasConsentedClaims returns a boolean if a field has been set.
func (o *UserinfoResponse) HasConsentedClaims() bool {
	if o != nil && !IsNil(o.ConsentedClaims) {
		return true
	}

	return false
}

// SetConsentedClaims gets a reference to the given []string and assigns it to the ConsentedClaims field.
func (o *UserinfoResponse) SetConsentedClaims(v []string) {
	o.ConsentedClaims = v
}

// GetRequestedClaimsForTx returns the RequestedClaimsForTx field value if set, zero value otherwise.
func (o *UserinfoResponse) GetRequestedClaimsForTx() []string {
	if o == nil || IsNil(o.RequestedClaimsForTx) {
		var ret []string
		return ret
	}
	return o.RequestedClaimsForTx
}

// GetRequestedClaimsForTxOk returns a tuple with the RequestedClaimsForTx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetRequestedClaimsForTxOk() ([]string, bool) {
	if o == nil || IsNil(o.RequestedClaimsForTx) {
		return nil, false
	}
	return o.RequestedClaimsForTx, true
}

// HasRequestedClaimsForTx returns a boolean if a field has been set.
func (o *UserinfoResponse) HasRequestedClaimsForTx() bool {
	if o != nil && !IsNil(o.RequestedClaimsForTx) {
		return true
	}

	return false
}

// SetRequestedClaimsForTx gets a reference to the given []string and assigns it to the RequestedClaimsForTx field.
func (o *UserinfoResponse) SetRequestedClaimsForTx(v []string) {
	o.RequestedClaimsForTx = v
}

// GetRequestedVerifiedClaimsForTx returns the RequestedVerifiedClaimsForTx field value if set, zero value otherwise.
func (o *UserinfoResponse) GetRequestedVerifiedClaimsForTx() [][]string {
	if o == nil || IsNil(o.RequestedVerifiedClaimsForTx) {
		var ret [][]string
		return ret
	}
	return o.RequestedVerifiedClaimsForTx
}

// GetRequestedVerifiedClaimsForTxOk returns a tuple with the RequestedVerifiedClaimsForTx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetRequestedVerifiedClaimsForTxOk() ([][]string, bool) {
	if o == nil || IsNil(o.RequestedVerifiedClaimsForTx) {
		return nil, false
	}
	return o.RequestedVerifiedClaimsForTx, true
}

// HasRequestedVerifiedClaimsForTx returns a boolean if a field has been set.
func (o *UserinfoResponse) HasRequestedVerifiedClaimsForTx() bool {
	if o != nil && !IsNil(o.RequestedVerifiedClaimsForTx) {
		return true
	}

	return false
}

// SetRequestedVerifiedClaimsForTx gets a reference to the given [][]string and assigns it to the RequestedVerifiedClaimsForTx field.
func (o *UserinfoResponse) SetRequestedVerifiedClaimsForTx(v [][]string) {
	o.RequestedVerifiedClaimsForTx = v
}

// GetTransformedClaims returns the TransformedClaims field value if set, zero value otherwise.
func (o *UserinfoResponse) GetTransformedClaims() string {
	if o == nil || IsNil(o.TransformedClaims) {
		var ret string
		return ret
	}
	return *o.TransformedClaims
}

// GetTransformedClaimsOk returns a tuple with the TransformedClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetTransformedClaimsOk() (*string, bool) {
	if o == nil || IsNil(o.TransformedClaims) {
		return nil, false
	}
	return o.TransformedClaims, true
}

// HasTransformedClaims returns a boolean if a field has been set.
func (o *UserinfoResponse) HasTransformedClaims() bool {
	if o != nil && !IsNil(o.TransformedClaims) {
		return true
	}

	return false
}

// SetTransformedClaims gets a reference to the given string and assigns it to the TransformedClaims field.
func (o *UserinfoResponse) SetTransformedClaims(v string) {
	o.TransformedClaims = &v
}

// GetClientEntityId returns the ClientEntityId field value if set, zero value otherwise.
func (o *UserinfoResponse) GetClientEntityId() string {
	if o == nil || IsNil(o.ClientEntityId) {
		var ret string
		return ret
	}
	return *o.ClientEntityId
}

// GetClientEntityIdOk returns a tuple with the ClientEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetClientEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientEntityId) {
		return nil, false
	}
	return o.ClientEntityId, true
}

// HasClientEntityId returns a boolean if a field has been set.
func (o *UserinfoResponse) HasClientEntityId() bool {
	if o != nil && !IsNil(o.ClientEntityId) {
		return true
	}

	return false
}

// SetClientEntityId gets a reference to the given string and assigns it to the ClientEntityId field.
func (o *UserinfoResponse) SetClientEntityId(v string) {
	o.ClientEntityId = &v
}

// GetClientEntityIdUsed returns the ClientEntityIdUsed field value if set, zero value otherwise.
func (o *UserinfoResponse) GetClientEntityIdUsed() bool {
	if o == nil || IsNil(o.ClientEntityIdUsed) {
		var ret bool
		return ret
	}
	return *o.ClientEntityIdUsed
}

// GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetClientEntityIdUsedOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientEntityIdUsed) {
		return nil, false
	}
	return o.ClientEntityIdUsed, true
}

// HasClientEntityIdUsed returns a boolean if a field has been set.
func (o *UserinfoResponse) HasClientEntityIdUsed() bool {
	if o != nil && !IsNil(o.ClientEntityIdUsed) {
		return true
	}

	return false
}

// SetClientEntityIdUsed gets a reference to the given bool and assigns it to the ClientEntityIdUsed field.
func (o *UserinfoResponse) SetClientEntityIdUsed(v bool) {
	o.ClientEntityIdUsed = &v
}

// GetDpopNonce returns the DpopNonce field value if set, zero value otherwise.
func (o *UserinfoResponse) GetDpopNonce() string {
	if o == nil || IsNil(o.DpopNonce) {
		var ret string
		return ret
	}
	return *o.DpopNonce
}

// GetDpopNonceOk returns a tuple with the DpopNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetDpopNonceOk() (*string, bool) {
	if o == nil || IsNil(o.DpopNonce) {
		return nil, false
	}
	return o.DpopNonce, true
}

// HasDpopNonce returns a boolean if a field has been set.
func (o *UserinfoResponse) HasDpopNonce() bool {
	if o != nil && !IsNil(o.DpopNonce) {
		return true
	}

	return false
}

// SetDpopNonce gets a reference to the given string and assigns it to the DpopNonce field.
func (o *UserinfoResponse) SetDpopNonce(v string) {
	o.DpopNonce = &v
}

func (o UserinfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserinfoResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Claims) {
		toSerialize["claims"] = o.Claims
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.ClientIdAlias) {
		toSerialize["clientIdAlias"] = o.ClientIdAlias
	}
	if !IsNil(o.ClientIdAliasUsed) {
		toSerialize["clientIdAliasUsed"] = o.ClientIdAliasUsed
	}
	if !IsNil(o.ResponseContent) {
		toSerialize["responseContent"] = o.ResponseContent
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.UserInfoClaims) {
		toSerialize["userInfoClaims"] = o.UserInfoClaims
	}
	if !IsNil(o.ServiceAttributes) {
		toSerialize["serviceAttributes"] = o.ServiceAttributes
	}
	if !IsNil(o.ClientAttributes) {
		toSerialize["clientAttributes"] = o.ClientAttributes
	}
	if !IsNil(o.ConsentedClaims) {
		toSerialize["consentedClaims"] = o.ConsentedClaims
	}
	if !IsNil(o.RequestedClaimsForTx) {
		toSerialize["requestedClaimsForTx"] = o.RequestedClaimsForTx
	}
	if !IsNil(o.RequestedVerifiedClaimsForTx) {
		toSerialize["requestedVerifiedClaimsForTx"] = o.RequestedVerifiedClaimsForTx
	}
	if !IsNil(o.TransformedClaims) {
		toSerialize["transformedClaims"] = o.TransformedClaims
	}
	if !IsNil(o.ClientEntityId) {
		toSerialize["clientEntityId"] = o.ClientEntityId
	}
	if !IsNil(o.ClientEntityIdUsed) {
		toSerialize["clientEntityIdUsed"] = o.ClientEntityIdUsed
	}
	if !IsNil(o.DpopNonce) {
		toSerialize["dpopNonce"] = o.DpopNonce
	}
	return toSerialize, nil
}

type NullableUserinfoResponse struct {
	value *UserinfoResponse
	isSet bool
}

func (v NullableUserinfoResponse) Get() *UserinfoResponse {
	return v.value
}

func (v *NullableUserinfoResponse) Set(val *UserinfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserinfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserinfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserinfoResponse(val *UserinfoResponse) *NullableUserinfoResponse {
	return &NullableUserinfoResponse{value: val, isSet: true}
}

func (v NullableUserinfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserinfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
