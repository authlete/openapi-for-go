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

// checks if the Client type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Client{}

// Client struct for Client
type Client struct {
	// The sequential number of the client. The value of this property is assigned by Authlete.
	Number *int32 `json:"number,omitempty"`
	// The sequential number of the service of the client application. The value of this property is assigned by Authlete.
	ServiceNumber *int32 `json:"serviceNumber,omitempty"`
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
	// The client secret. A random 512-bit value encoded by base64url (86 letters). The value of this property is assigned by Authlete.  Note that Authlete issues a client secret even to a \"public\" client application, but the client application should not use the client secret unless it changes its client type to \"confidential\". That is, a public client application should behave as if it had not been issued a client secret. To be specific, a token request from a public client of Authlete should not come along with a client secret although [RFC 6749, 3.2.1. Client Authentication](https://datatracker.ietf.org/doc/html/rfc6749#section-3.2.1) says as follows.  > Confidential clients or other clients issued client credentials MUST authenticate with the authorization server as described in Section 2.3 when making requests to the token endpoint.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// The value of the client's `client_id` property used in OAuth and OpenID Connect calls. By default, this is a string version of the `clientId` property.
	ClientIdAlias *string `json:"clientIdAlias,omitempty"`
	// Deprecated. Always set to `true`.
	ClientIdAliasEnabled *bool                   `json:"clientIdAliasEnabled,omitempty"`
	ClientType           *ClientType             `json:"clientType,omitempty"`
	ApplicationType      NullableApplicationType `json:"applicationType,omitempty"`
	// The URL pointing to the logo image of the client application.  This property corresponds to `logo_uri` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).
	LogoUri *string `json:"logoUri,omitempty"`
	// Logo image URLs with language tags. If the client application has different logo images for different languages, this property can be used to register URLs of the images.
	LogoUris []TaggedValue `json:"logoUris,omitempty"`
	// An array of email addresses of people responsible for the client application.  This property corresponds to contacts in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).
	Contacts []string `json:"contacts,omitempty"`
	// The flag to indicate whether this client use TLS client certificate bound access tokens.
	TlsClientCertificateBoundAccessTokens *bool `json:"tlsClientCertificateBoundAccessTokens,omitempty"`
	// The flag to indicate whether this client has been registered dynamically. For more details, see [RFC 7591](https://datatracker.ietf.org/doc/html/rfc7591).
	DynamicallyRegistered *bool `json:"dynamicallyRegistered,omitempty"`
	// The unique identifier string assigned by the client developer or software publisher used by registration endpoints to identify the client software to be dynamically registered.  This property corresponds to the `software_id metadata` defined in [2. Client Metadata](https://datatracker.ietf.org/doc/html/rfc7591#section-2) of [RFC 7591](https://datatracker.ietf.org/doc/html/rfc7591).
	SoftwareId *string `json:"softwareId,omitempty"`
	// The version identifier string for the client software identified by the software ID.  This property corresponds to the software_version metadata defined in [2. Client Metadata](https://datatracker.ietf.org/doc/html/rfc7591#section-2) of [RFC 7591](https://datatracker.ietf.org/doc/html/rfc7591).
	SoftwareVersion *string `json:"softwareVersion,omitempty"`
	// The hash of the registration access token for this client.
	RegistrationAccessTokenHash *string `json:"registrationAccessTokenHash,omitempty"`
	// The time at which this client was created. The value is represented as milliseconds since the UNIX epoch (1970-01-01).
	CreatedAt *int64 `json:"createdAt,omitempty"`
	// The time at which this client was last modified. The value is represented as milliseconds since the UNIX epoch (1970-01-01).
	ModifiedAt *int64 `json:"modifiedAt,omitempty"`
	// A string array of grant types which the client application declares that it will restrict itself to using. This property corresponds to `grant_types` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).
	GrantTypes []GrantType `json:"grantTypes,omitempty"`
	// A string array of response types which the client application declares that it will restrict itself to using. This property corresponds to `response_types` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).
	ResponseTypes []ResponseType `json:"responseTypes,omitempty"`
	// Redirect URIs that the client application uses to receive a response from the authorization endpoint. Requirements for a redirect URI are as follows.  **Requirements by RFC 6749** (From [RFC 6749, 3.1.2. Redirection Endpoint](https://datatracker.ietf.org/doc/html/rfc6749#section-3.1.2))  - Must be an absolute URI. - Must not have a fragment component.  **Requirements by OpenID Connect** (From \"[OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata), application_type\")  - The scheme of the redirect URI used for Implicit Grant by a client application whose application is `web` must be `https`. This is checked at runtime by Authlete. - The hostname of the redirect URI used for Implicit Grant by a client application whose application type is `web` must not be `localhost`. This is checked at runtime by Authlete. - The scheme of the redirect URI used by a client application whose application type is `native` must be either (1) a custom scheme or (2) `http`, which is allowed only when the hostname part is `localhost`. This is checked at runtime by Authlete.  **Requirements by Authlete**  - Must consist of printable ASCII letters only. - Must not exceed 200 letters.  Note that Authlete allows the application type to be `null`. In other words, a client application does not have to choose `web` or `native` as its application type. If the application type is `null`, the requirements by OpenID Connect are not checked at runtime.  An authorization request from a client application which has not registered any redirect URI fails unless at least all the following conditions are satisfied.  - The client type of the client application is `confidential`. - The value of `response_type` request parameter is `code`. - The authorization request has the `redirect_uri` request parameter. - The value of `scope` request parameter does not contain `openid`.  RFC 6749 allows partial match of redirect URI under some conditions (see [RFC 6749, 3.1.2.2. Registration Requirements](https://datatracker.ietf.org/doc/html/rfc6749#section-3.1.2.2) for details), but OpenID Connect requires exact match.
	RedirectUris               []string          `json:"redirectUris,omitempty"`
	AuthorizationSignAlg       *JwsAlg           `json:"authorizationSignAlg,omitempty"`
	AuthorizationEncryptionAlg *JweAlg           `json:"authorizationEncryptionAlg,omitempty"`
	AuthorizationEncryptionEnc *JweEnc           `json:"authorizationEncryptionEnc,omitempty"`
	TokenAuthMethod            *ClientAuthMethod `json:"tokenAuthMethod,omitempty"`
	TokenAuthSignAlg           *JwsAlg           `json:"tokenAuthSignAlg,omitempty"`
	// The key ID of a JWK containing a self-signed certificate of this client.
	SelfSignedCertificateKeyId *string `json:"selfSignedCertificateKeyId,omitempty"`
	// The string representation of the expected subject distinguished name of the certificate this client will use in mutual TLS authentication.  See `tls_client_auth_subject_dn` in \"Mutual TLS Profiles for OAuth Clients, 2.3. Dynamic Client Registration\" for details.
	TlsClientAuthSubjectDn *string `json:"tlsClientAuthSubjectDn,omitempty"`
	// The string representation of the expected DNS subject alternative name of the certificate this client will use in mutual TLS authentication.  See `tls_client_auth_san_dns` in \"Mutual TLS Profiles for OAuth Clients, 2.3. Dynamic Client Registration\" for details.
	TlsClientAuthSanDns *string `json:"tlsClientAuthSanDns,omitempty"`
	// The string representation of the expected URI subject alternative name of the certificate this client will use in mutual TLS authentication.  See `tls_client_auth_san_uri` in \"Mutual TLS Profiles for OAuth Clients, 2.3. Dynamic Client Registration\" for details.
	TlsClientAuthSanUri *string `json:"tlsClientAuthSanUri,omitempty"`
	// The string representation of the expected IP address subject alternative name of the certificate this client will use in mutual TLS authentication.  See `tls_client_auth_san_ip` in \"Mutual TLS Profiles for OAuth Clients, 2.3. Dynamic Client Registration\" for details.
	TlsClientAuthSanIp *string `json:"tlsClientAuthSanIp,omitempty"`
	// The string representation of the expected email address subject alternative name of the certificate this client will use in mutual TLS authentication.  See `tls_client_auth_san_email` in \"Mutual TLS Profiles for OAuth Clients, 2.3. Dynamic Client Registration\" for details.
	TlsClientAuthSanEmail *string `json:"tlsClientAuthSanEmail,omitempty"`
	// The flag to indicate whether this client is required to use the pushed authorization request endpoint. This property corresponds to the `require_pushed_authorization_requests` client metadata defined in \"OAuth 2.0 Pushed Authorization Requests\".
	ParRequired *bool `json:"parRequired,omitempty"`
	// The flag to indicate whether authorization requests from this client are always required to utilize a request object by using either `request` or `request_uri` request parameter.  If this flag is set to `true` and the service's `traditionalRequestObjectProcessingApplied` is set to `false`, authorization requests from this client are processed as if `require_signed_request_object` client metadata of this client is `true`. The metadata is defined in \"JAR (JWT Secured Authorization Request)\".
	RequestObjectRequired *bool   `json:"requestObjectRequired,omitempty"`
	RequestSignAlg        *JwsAlg `json:"requestSignAlg,omitempty"`
	RequestEncryptionAlg  *JweAlg `json:"requestEncryptionAlg,omitempty"`
	RequestEncryptionEnc  *JweEnc `json:"requestEncryptionEnc,omitempty"`
	// An array of URLs each of which points to a request object.  Authlete requires that URLs used as values for `request_uri` request parameter be pre-registered. This property is used for the pre-registration. See [OpenID Connect Core 1.0, 6.2. Passing a Request Object by Reference](https://openid.net/specs/openid-connect-core-1_0.html#RequestUriParameter) for details.
	RequestUris []string `json:"requestUris,omitempty"`
	// The default maximum authentication age in seconds. This value is used when an authorization request from the client application does not have `max_age` request parameter.  This property corresponds to `default_max_age` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).
	DefaultMaxAge *int32 `json:"defaultMaxAge,omitempty"`
	// The default ACRs (Authentication Context Class References). This value is used when an authorization request from the client application has neither `acr_values` request parameter nor `acr` claim in claims request parameter.
	DefaultAcrs          []string `json:"defaultAcrs,omitempty"`
	IdTokenSignAlg       *JwsAlg  `json:"idTokenSignAlg,omitempty"`
	IdTokenEncryptionAlg *JweAlg  `json:"idTokenEncryptionAlg,omitempty"`
	IdTokenEncryptionEnc *JweEnc  `json:"idTokenEncryptionEnc,omitempty"`
	// The flag to indicate whether this client requires `auth_time` claim to be embedded in the ID token.  This property corresponds to `require_auth_time` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).
	AuthTimeRequired *bool        `json:"authTimeRequired,omitempty"`
	SubjectType      *SubjectType `json:"subjectType,omitempty"`
	// The value of the sector identifier URI. This represents the `sector_identifier_uri` client metadata which is defined in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata)
	SectorIdentifierUri *string `json:"sectorIdentifierUri,omitempty"`
	// The sector identifier host component as derived from either the `sector_identifier_uri` or the registered redirect URI. If no `sector_identifier_uri` is registered and multiple redirect URIs are also registered, the value of this property is `null`.
	DerivedSectorIdentifier *string `json:"derivedSectorIdentifier,omitempty"`
	// The URL pointing to the JWK Set of the client application. The content pointed to by the URL is JSON which complies with the format described in [JSON Web Key (JWK), 5. JWK Set Format](https://datatracker.ietf.org/doc/html/rfc7517#section-5). The JWK Set must not include private keys of the client application.  If the client application requests encryption for ID tokens (from the authorization/token/userinfo endpoints) and/or signs request objects, it must make available its JWK Set containing public keys for the encryption and/or the signature at the URL of `jwksUri`. The service (Authlete) fetches the JWK Set from the URL as necessary.  [OpenID Connect Dynamic Client Registration 1.0](https://openid.net/specs/openid-connect-registration-1_0.html) says that `jwks` must not be used when the client can use `jwks_uri`, but Authlete allows both properties to be registered at the same time. However, Authlete does not use the content of `jwks` when `jwksUri` is registered.  This property corresponds to `jwks_uri` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).
	JwksUri *string `json:"jwksUri,omitempty"`
	// The content of the JWK Set of the client application. The format is described in [JSON Web Key (JWK), 5. JWK Set Format](https://datatracker.ietf.org/doc/html/rfc7517#section-5). The JWK Set must not include private keys of the client application.  [OpenID Connect Dynamic Client Registration 1.0](https://openid.net/specs/openid-connect-registration-1_0.html) says that `jwks` must not be used when the client can use `jwks_uri`, but Authlete allows both properties to be registered at the same time. However, Authlete does not use the content of `jwks` when `jwksUri` is registered.  This property corresponds to `jwks_uri` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).
	Jwks                  *string `json:"jwks,omitempty"`
	UserInfoSignAlg       *JwsAlg `json:"userInfoSignAlg,omitempty"`
	UserInfoEncryptionAlg *JweAlg `json:"userInfoEncryptionAlg,omitempty"`
	UserInfoEncryptionEnc *JweEnc `json:"userInfoEncryptionEnc,omitempty"`
	// The URL which a third party can use to initiate a login by the client application.  This property corresponds to `initiate_login_uri` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).
	LoginUri *string `json:"loginUri,omitempty"`
	// The URL pointing to the \"Terms Of Service\" page.  This property corresponds to `tos_uri` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).
	TosUri *string `json:"tosUri,omitempty"`
	// URLs of \"Terms Of Service\" pages with language tags.  If the client application has different \"Terms Of Service\" pages for different languages, this property can be used to register the URLs.
	TosUris []TaggedValue `json:"tosUris,omitempty"`
	// The URL pointing to the page which describes the policy as to how end-user's profile data is used.  This property corresponds to `policy_uri` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).
	PolicyUri *string `json:"policyUri,omitempty"`
	// URLs of policy pages with language tags. If the client application has different policy pages for different languages, this property can be used to register the URLs.
	PolicyUris []TaggedValue `json:"policyUris,omitempty"`
	// The URL pointing to the home page of the client application.  This property corresponds to `client_uri` in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).
	ClientUri *string `json:"clientUri,omitempty"`
	// Home page URLs with language tags. If the client application has different home pages for different languages, this property can be used to register the URLs.
	ClientUris []TaggedValue `json:"clientUris,omitempty"`
	// The backchannel token delivery mode.  This property corresponds to the `backchannel_token_delivery_mode` metadata. The backchannel token delivery mode is defined in the specification of \"CIBA (Client Initiated Backchannel Authentication)\".
	BcDeliveryMode *string `json:"bcDeliveryMode,omitempty"`
	// The backchannel client notification endpoint.  This property corresponds to the `backchannel_client_notification_endpoint` metadata. The backchannel token delivery mode is defined in the specification of \"CIBA (Client Initiated Backchannel Authentication)\".
	BcNotificationEndpoint *string `json:"bcNotificationEndpoint,omitempty"`
	BcRequestSignAlg       *JwsAlg `json:"bcRequestSignAlg,omitempty"`
	// The boolean flag to indicate whether a user code is required when this client makes a backchannel authentication request.  This property corresponds to the `backchannel_user_code_parameter` metadata.
	BcUserCodeRequired *bool `json:"bcUserCodeRequired,omitempty"`
	// The attributes of this client.
	Attributes []Pair           `json:"attributes,omitempty"`
	Extension  *ClientExtension `json:"extension,omitempty"`
	// The authorization details types that this client may use as values of the `type` field in `authorization_details`.  This property corresponds to the `authorization_details_types` metadata. See [OAuth 2.0 Rich Authorization Requests (RAR)](https://datatracker.ietf.org/doc/draft-ietf-oauth-rar/) for details.  Note that the property name was renamed from authorizationDataTypes to authorizationDetailsTypes to align with the change made by the 5th draft of the RAR specification.
	AuthorizationDetailsTypes []string `json:"authorizationDetailsTypes,omitempty"`
	// The custom client metadata in JSON format.  Standard specifications define client metadata as necessary. The following are such examples.  * [OpenID Connect Dynamic Client Registration 1.0](https://openid.net/specs/openid-connect-registration-1_0.html) * [RFC 7591 OAuth 2.0 Dynamic Client Registration Protocol](https://www.rfc-editor.org/rfc/rfc7591.html) * [RFC 8705 OAuth 2.0 Mutual-TLS Client Authentication and Certificate-Bound Access Tokens](https://www.rfc-editor.org/rfc/rfc8705.html) * [OpenID Connect Client-Initiated Backchannel Authentication Flow - Core 1.0](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html) * [The OAuth 2.0 Authorization Framework: JWT Secured Authorization Request (JAR)](https://datatracker.ietf.org/doc/draft-ietf-oauth-jwsreq/) * [Financial-grade API: JWT Secured Authorization Response Mode for OAuth 2.0 (JARM)](https://openid.net/specs/openid-financial-api-jarm.html) * [OAuth 2.0 Pushed Authorization Requests (PAR)](https://datatracker.ietf.org/doc/rfc9126/) * [OAuth 2.0 Rich Authorization Requests (RAR)](https://datatracker.ietf.org/doc/draft-ietf-oauth-rar/)  Standard client metadata included in Client Registration Request and Client Update Request (cf. [OIDC DynReg](https://openid.net/specs/openid-connect-registration-1_0.html), [RFC 7591](https://www.rfc-editor.org/rfc/rfc7591.html) and [RFC 7592](https://www.rfc-editor.org/rfc/rfc7592.html)) are, if supported by Authlete, set to corresponding properties of the client application. For example, the value of the `client_name` client metadata in Client Registration/Update Request is set to the clientName property. On the other hand, unrecognized client metadata are discarded.  By listing up custom client metadata in advance by using the `supportedCustomClientMetadata` property of Service, Authlete can recognize them and stores their values into the database. The stored custom client metadata values can be referenced by this property.
	CustomMetadata *string `json:"customMetadata,omitempty"`
	// The flag indicating whether encryption of request object is required when the request object is passed through the front channel.  This flag does not affect the processing of request objects at the Pushed Authorization Request Endpoint, which is defined in [OAuth 2.0 Pushed Authorization Requests](https://datatracker.ietf.org/doc/rfc9126/). Unecrypted request objects are accepted at the endpoint even if this flag is `true`.  This flag does not indicate whether a request object is always required. There is a different flag, `requestObjectRequired`, for the purpose.  Even if this flag is `false`, encryption of request object is required if the `frontChannelRequestObjectEncryptionRequired` flag of the service is `true`.
	FrontChannelRequestObjectEncryptionRequired *bool `json:"frontChannelRequestObjectEncryptionRequired,omitempty"`
	// The flag indicating whether the JWE alg of encrypted request object must match the `request_object_encryption_alg` client metadata.  The `request_object_encryption_alg` client metadata itself is defined in [OpenID Connect Dynamic Client Registration 1.0](https://openid.net/specs/openid-connect-registration-1_0.html) as follows.  > request_object_encryption_alg > > OPTIONAL. JWE [JWE] alg algorithm [JWA] the RP is declaring that it may use for encrypting Request   Objects sent to the OP. This parameter SHOULD be included when symmetric encryption will be used,   since this signals to the OP that a client_secret value needs to be returned from which the   symmetric key will be derived, that might not otherwise be returned. The RP MAY still use other   supported encryption algorithms or send unencrypted Request Objects, even when this parameter   is present. If both signing and encryption are requested, the Request Object will be signed   then encrypted, with the result being a Nested JWT, as defined in [JWT]. The default, if omitted,   is that the RP is not declaring whether it might encrypt any Request Objects.  The point here is \"The RP MAY still use other supported encryption algorithms or send unencrypted Request Objects, even when this parameter is present.\"  The property that represents the client metadata is `requestEncryptionAlg`. See the description of `requestEncryptionAlg` for details.  Even if this flag is `false`, the match is required if the `requestObjectEncryptionAlgMatchRequired` flag of the service is `true`.
	RequestObjectEncryptionAlgMatchRequired *bool `json:"requestObjectEncryptionAlgMatchRequired,omitempty"`
	// The flag indicating whether the JWE enc of encrypted request object must match the `request_object_encryption_enc` client metadata.  The `request_object_encryption_enc` client metadata itself is defined in [OpenID Connect Dynamic Client Registration 1.0](https://openid.net/specs/openid-connect-registration-1_0.html) as follows.  > request_object_encryption_enc > > OPTIONAL. JWE enc algorithm [JWA] the RP is declaring that it may use for encrypting Request   Objects sent to the OP. If request_object_encryption_alg is specified, the default for this   value is A128CBC-HS256. When request_object_encryption_enc is included, request_object_encryption_alg   MUST also be provided.  The property that represents the client metadata is `requestEncryptionEnc`. See the description of `requestEncryptionEnc`  for details.  Even if this flag is `false`, the match is required if the `requestObjectEncryptionEncMatchRequired` flag of the service is `true`.
	RequestObjectEncryptionEncMatchRequired *bool `json:"requestObjectEncryptionEncMatchRequired,omitempty"`
	// The digest algorithm that this client requests the server to use when it computes digest values of <a href= \"https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html#name-external-attachments\" >external attachments</a>, which may be referenced from within ID tokens or userinfo responses (or any place that can have the `verified_claims` claim).  Possible values are listed in the <a href= \"https://www.iana.org/assignments/named-information/named-information.xhtml#hash-alg\" >Hash Algorithm Registry</a> of IANA (Internet Assigned Numbers Authority), but the server does not necessarily support all the values there. When this property is omitted, `sha-256` is used as the default algorithm.  This property corresponds to the `digest_algorithm` client metadata which was defined by the third implementer's draft of [OpenID Connect for Identity Assurance 1.0](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html).
	DigestAlgorithm *string `json:"digestAlgorithm,omitempty"`
	// If `Enabled` is selected, an attempt to issue a new access token invalidates existing access tokens that are associated with the same combination of subject and client.  Note that, however, attempts by Client Credentials Flow do not invalidate existing access tokens because access tokens issued by Client Credentials Flow are not associated with any end-user's subject.  Even if `Disabled` is selected here, single access token per subject is effective if `singleAccessTokenPerSubject` of the `Service` this client belongs to is Enabled.
	SingleAccessTokenPerSubject *bool `json:"singleAccessTokenPerSubject,omitempty"`
	// The flag to indicate whether the use of Proof Key for Code Exchange (PKCE) is always required for authorization requests by Authorization Code Flow.  If `true`, `code_challenge` request parameter is always required for authorization requests using Authorization Code Flow.  See [RFC 7636](https://tools.ietf.org/html/rfc7636) (Proof Key for Code Exchange by OAuth Public Clients) for details about `code_challenge` request parameter.
	PkceRequired *bool `json:"pkceRequired,omitempty"`
	// The flag to indicate whether `S256` is always required as the code challenge method whenever [PKCE (RFC 7636)](https://tools.ietf.org/html/rfc7636) is used.  If this flag is set to `true`, `code_challenge_method=S256` must be included in the authorization request whenever it includes the `code_challenge` request parameter. Neither omission of the `code_challenge_method` request parameter nor use of plain (`code_challenge_method=plain`) is allowed.
	PkceS256Required *bool `json:"pkceS256Required,omitempty"`
	// If the DPoP is required for this client
	DpopRequired *bool `json:"dpopRequired,omitempty"`
	// The flag indicating whether this client was registered by the \"automatic\" client registration of OIDC Federation.
	AutomaticallyRegistered *bool `json:"automaticallyRegistered,omitempty"`
	// The flag indicating whether this client was registered by the \"explicit\" client registration of OIDC Federation.
	ExplicitlyRegistered *bool `json:"explicitlyRegistered,omitempty"`
	// The flag indicating whether this service signs responses from the resource server.
	RsRequestSigned *bool `json:"rsRequestSigned,omitempty"`
	// The key ID of a JWK containing the public key used by this client to sign requests to the resource server.
	RsSignedRequestKeyId *string `json:"rsSignedRequestKeyId,omitempty"`
	// The client registration types that the client has declared it may use.
	ClientRegistrationTypes []ClientRegistrationType `json:"clientRegistrationTypes,omitempty"`
	// The human-readable name representing the organization that manages this client. This property corresponds to the organization_name client metadata that is defined in OpenID Connect Federation 1.0.
	OrganizationName *string `json:"organizationName,omitempty"`
	// The URI of the endpoint that returns this client's JWK Set document in the JWT format. This property corresponds to the `signed_jwks_uri` client metadata defined in OpenID Connect Federation 1.0.
	SignedJwksUri *string `json:"signedJwksUri,omitempty"`
	// the entity ID of this client.
	EntityId *string `json:"entityId,omitempty"`
	// The entity ID of the trust anchor of the trust chain that was used when this client was registered or updated by the mechanism defined in OpenID Connect Federation 1.0
	TrustAnchorId *string `json:"trustAnchorId,omitempty"`
	// The trust chain that was used when this client was registered or updated by the mechanism defined in OpenID Connect Federation 1.0
	TrustChain []string `json:"trustChain,omitempty"`
	// the expiration time of the trust chain that was used when this client was registered or updated by the mechanism defined in OpenID Connect Federation 1.0. The value is represented as milliseconds elapsed since the Unix epoch (1970-01-01).
	TrustChainExpiresAt *int64 `json:"trustChainExpiresAt,omitempty"`
	// the time at which the trust chain was updated by the mechanism defined in OpenID Connect Federation 1.0
	TrustChainUpdatedAt *int64 `json:"trustChainUpdatedAt,omitempty"`
	// The flag which indicates whether this client is locked.
	Locked *bool `json:"locked,omitempty"`
	// The URL of the credential offer endpoint at which this client (wallet) receives a credential offer from the credential issuer.
	CredentialOfferEndpoint *string `json:"credentialOfferEndpoint,omitempty"`
	// The FAPI modes for this client.  When the value of this property is not `null`, Authlete always processes requests from this client based on the specified FAPI modes if the FAPI feature is enabled in Authlete, the FAPI profile is supported by the service, and the FAPI modes for the service are set to `null`.  For instance, when this property is set to an array containing `FAPI1_ADVANCED` only, Authlete always processes requests from this client based on \"Financial-grade API Security Profile 1.0 - Part 2: Advanced\" if the FAPI feature is enabled in Authlete, the FAPI profile is supported by the service, and the FAPI modes for the service are set to `null`.
	FapiModes []FapiMode `json:"fapiModes,omitempty"`
	// The response modes that this client may use.
	ResponseModes []string `json:"responseModes,omitempty"`
	// True if credential responses to this client must be always encrypted.
	CredentialResponseEncryptionRequired *bool `json:"credentialResponseEncryptionRequired,omitempty"`
}

// NewClient instantiates a new Client object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClient() *Client {
	this := Client{}
	return &this
}

// NewClientWithDefaults instantiates a new Client object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientWithDefaults() *Client {
	this := Client{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *Client) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *Client) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *Client) SetNumber(v int32) {
	o.Number = &v
}

// GetServiceNumber returns the ServiceNumber field value if set, zero value otherwise.
func (o *Client) GetServiceNumber() int32 {
	if o == nil || IsNil(o.ServiceNumber) {
		var ret int32
		return ret
	}
	return *o.ServiceNumber
}

// GetServiceNumberOk returns a tuple with the ServiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetServiceNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.ServiceNumber) {
		return nil, false
	}
	return o.ServiceNumber, true
}

// HasServiceNumber returns a boolean if a field has been set.
func (o *Client) HasServiceNumber() bool {
	if o != nil && !IsNil(o.ServiceNumber) {
		return true
	}

	return false
}

// SetServiceNumber gets a reference to the given int32 and assigns it to the ServiceNumber field.
func (o *Client) SetServiceNumber(v int32) {
	o.ServiceNumber = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *Client) GetClientName() string {
	if o == nil || IsNil(o.ClientName) {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetClientNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClientName) {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *Client) HasClientName() bool {
	if o != nil && !IsNil(o.ClientName) {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *Client) SetClientName(v string) {
	o.ClientName = &v
}

// GetClientNames returns the ClientNames field value if set, zero value otherwise.
func (o *Client) GetClientNames() []TaggedValue {
	if o == nil || IsNil(o.ClientNames) {
		var ret []TaggedValue
		return ret
	}
	return o.ClientNames
}

// GetClientNamesOk returns a tuple with the ClientNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetClientNamesOk() ([]TaggedValue, bool) {
	if o == nil || IsNil(o.ClientNames) {
		return nil, false
	}
	return o.ClientNames, true
}

// HasClientNames returns a boolean if a field has been set.
func (o *Client) HasClientNames() bool {
	if o != nil && !IsNil(o.ClientNames) {
		return true
	}

	return false
}

// SetClientNames gets a reference to the given []TaggedValue and assigns it to the ClientNames field.
func (o *Client) SetClientNames(v []TaggedValue) {
	o.ClientNames = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Client) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Client) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Client) SetDescription(v string) {
	o.Description = &v
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *Client) GetDescriptions() []TaggedValue {
	if o == nil || IsNil(o.Descriptions) {
		var ret []TaggedValue
		return ret
	}
	return o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetDescriptionsOk() ([]TaggedValue, bool) {
	if o == nil || IsNil(o.Descriptions) {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *Client) HasDescriptions() bool {
	if o != nil && !IsNil(o.Descriptions) {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []TaggedValue and assigns it to the Descriptions field.
func (o *Client) SetDescriptions(v []TaggedValue) {
	o.Descriptions = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *Client) GetClientId() int64 {
	if o == nil || IsNil(o.ClientId) {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetClientIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *Client) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *Client) SetClientId(v int64) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *Client) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *Client) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *Client) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetClientIdAlias returns the ClientIdAlias field value if set, zero value otherwise.
func (o *Client) GetClientIdAlias() string {
	if o == nil || IsNil(o.ClientIdAlias) {
		var ret string
		return ret
	}
	return *o.ClientIdAlias
}

// GetClientIdAliasOk returns a tuple with the ClientIdAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetClientIdAliasOk() (*string, bool) {
	if o == nil || IsNil(o.ClientIdAlias) {
		return nil, false
	}
	return o.ClientIdAlias, true
}

// HasClientIdAlias returns a boolean if a field has been set.
func (o *Client) HasClientIdAlias() bool {
	if o != nil && !IsNil(o.ClientIdAlias) {
		return true
	}

	return false
}

// SetClientIdAlias gets a reference to the given string and assigns it to the ClientIdAlias field.
func (o *Client) SetClientIdAlias(v string) {
	o.ClientIdAlias = &v
}

// GetClientIdAliasEnabled returns the ClientIdAliasEnabled field value if set, zero value otherwise.
func (o *Client) GetClientIdAliasEnabled() bool {
	if o == nil || IsNil(o.ClientIdAliasEnabled) {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasEnabled
}

// GetClientIdAliasEnabledOk returns a tuple with the ClientIdAliasEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetClientIdAliasEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientIdAliasEnabled) {
		return nil, false
	}
	return o.ClientIdAliasEnabled, true
}

// HasClientIdAliasEnabled returns a boolean if a field has been set.
func (o *Client) HasClientIdAliasEnabled() bool {
	if o != nil && !IsNil(o.ClientIdAliasEnabled) {
		return true
	}

	return false
}

// SetClientIdAliasEnabled gets a reference to the given bool and assigns it to the ClientIdAliasEnabled field.
func (o *Client) SetClientIdAliasEnabled(v bool) {
	o.ClientIdAliasEnabled = &v
}

// GetClientType returns the ClientType field value if set, zero value otherwise.
func (o *Client) GetClientType() ClientType {
	if o == nil || IsNil(o.ClientType) {
		var ret ClientType
		return ret
	}
	return *o.ClientType
}

// GetClientTypeOk returns a tuple with the ClientType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetClientTypeOk() (*ClientType, bool) {
	if o == nil || IsNil(o.ClientType) {
		return nil, false
	}
	return o.ClientType, true
}

// HasClientType returns a boolean if a field has been set.
func (o *Client) HasClientType() bool {
	if o != nil && !IsNil(o.ClientType) {
		return true
	}

	return false
}

// SetClientType gets a reference to the given ClientType and assigns it to the ClientType field.
func (o *Client) SetClientType(v ClientType) {
	o.ClientType = &v
}

// GetApplicationType returns the ApplicationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Client) GetApplicationType() ApplicationType {
	if o == nil || IsNil(o.ApplicationType.Get()) {
		var ret ApplicationType
		return ret
	}
	return *o.ApplicationType.Get()
}

// GetApplicationTypeOk returns a tuple with the ApplicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Client) GetApplicationTypeOk() (*ApplicationType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationType.Get(), o.ApplicationType.IsSet()
}

// HasApplicationType returns a boolean if a field has been set.
func (o *Client) HasApplicationType() bool {
	if o != nil && o.ApplicationType.IsSet() {
		return true
	}

	return false
}

// SetApplicationType gets a reference to the given NullableApplicationType and assigns it to the ApplicationType field.
func (o *Client) SetApplicationType(v ApplicationType) {
	o.ApplicationType.Set(&v)
}

// SetApplicationTypeNil sets the value for ApplicationType to be an explicit nil
func (o *Client) SetApplicationTypeNil() {
	o.ApplicationType.Set(nil)
}

// UnsetApplicationType ensures that no value is present for ApplicationType, not even an explicit nil
func (o *Client) UnsetApplicationType() {
	o.ApplicationType.Unset()
}

// GetLogoUri returns the LogoUri field value if set, zero value otherwise.
func (o *Client) GetLogoUri() string {
	if o == nil || IsNil(o.LogoUri) {
		var ret string
		return ret
	}
	return *o.LogoUri
}

// GetLogoUriOk returns a tuple with the LogoUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetLogoUriOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUri) {
		return nil, false
	}
	return o.LogoUri, true
}

// HasLogoUri returns a boolean if a field has been set.
func (o *Client) HasLogoUri() bool {
	if o != nil && !IsNil(o.LogoUri) {
		return true
	}

	return false
}

// SetLogoUri gets a reference to the given string and assigns it to the LogoUri field.
func (o *Client) SetLogoUri(v string) {
	o.LogoUri = &v
}

// GetLogoUris returns the LogoUris field value if set, zero value otherwise.
func (o *Client) GetLogoUris() []TaggedValue {
	if o == nil || IsNil(o.LogoUris) {
		var ret []TaggedValue
		return ret
	}
	return o.LogoUris
}

// GetLogoUrisOk returns a tuple with the LogoUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetLogoUrisOk() ([]TaggedValue, bool) {
	if o == nil || IsNil(o.LogoUris) {
		return nil, false
	}
	return o.LogoUris, true
}

// HasLogoUris returns a boolean if a field has been set.
func (o *Client) HasLogoUris() bool {
	if o != nil && !IsNil(o.LogoUris) {
		return true
	}

	return false
}

// SetLogoUris gets a reference to the given []TaggedValue and assigns it to the LogoUris field.
func (o *Client) SetLogoUris(v []TaggedValue) {
	o.LogoUris = v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *Client) GetContacts() []string {
	if o == nil || IsNil(o.Contacts) {
		var ret []string
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetContactsOk() ([]string, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *Client) HasContacts() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []string and assigns it to the Contacts field.
func (o *Client) SetContacts(v []string) {
	o.Contacts = v
}

// GetTlsClientCertificateBoundAccessTokens returns the TlsClientCertificateBoundAccessTokens field value if set, zero value otherwise.
func (o *Client) GetTlsClientCertificateBoundAccessTokens() bool {
	if o == nil || IsNil(o.TlsClientCertificateBoundAccessTokens) {
		var ret bool
		return ret
	}
	return *o.TlsClientCertificateBoundAccessTokens
}

// GetTlsClientCertificateBoundAccessTokensOk returns a tuple with the TlsClientCertificateBoundAccessTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetTlsClientCertificateBoundAccessTokensOk() (*bool, bool) {
	if o == nil || IsNil(o.TlsClientCertificateBoundAccessTokens) {
		return nil, false
	}
	return o.TlsClientCertificateBoundAccessTokens, true
}

// HasTlsClientCertificateBoundAccessTokens returns a boolean if a field has been set.
func (o *Client) HasTlsClientCertificateBoundAccessTokens() bool {
	if o != nil && !IsNil(o.TlsClientCertificateBoundAccessTokens) {
		return true
	}

	return false
}

// SetTlsClientCertificateBoundAccessTokens gets a reference to the given bool and assigns it to the TlsClientCertificateBoundAccessTokens field.
func (o *Client) SetTlsClientCertificateBoundAccessTokens(v bool) {
	o.TlsClientCertificateBoundAccessTokens = &v
}

// GetDynamicallyRegistered returns the DynamicallyRegistered field value if set, zero value otherwise.
func (o *Client) GetDynamicallyRegistered() bool {
	if o == nil || IsNil(o.DynamicallyRegistered) {
		var ret bool
		return ret
	}
	return *o.DynamicallyRegistered
}

// GetDynamicallyRegisteredOk returns a tuple with the DynamicallyRegistered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetDynamicallyRegisteredOk() (*bool, bool) {
	if o == nil || IsNil(o.DynamicallyRegistered) {
		return nil, false
	}
	return o.DynamicallyRegistered, true
}

// HasDynamicallyRegistered returns a boolean if a field has been set.
func (o *Client) HasDynamicallyRegistered() bool {
	if o != nil && !IsNil(o.DynamicallyRegistered) {
		return true
	}

	return false
}

// SetDynamicallyRegistered gets a reference to the given bool and assigns it to the DynamicallyRegistered field.
func (o *Client) SetDynamicallyRegistered(v bool) {
	o.DynamicallyRegistered = &v
}

// GetSoftwareId returns the SoftwareId field value if set, zero value otherwise.
func (o *Client) GetSoftwareId() string {
	if o == nil || IsNil(o.SoftwareId) {
		var ret string
		return ret
	}
	return *o.SoftwareId
}

// GetSoftwareIdOk returns a tuple with the SoftwareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetSoftwareIdOk() (*string, bool) {
	if o == nil || IsNil(o.SoftwareId) {
		return nil, false
	}
	return o.SoftwareId, true
}

// HasSoftwareId returns a boolean if a field has been set.
func (o *Client) HasSoftwareId() bool {
	if o != nil && !IsNil(o.SoftwareId) {
		return true
	}

	return false
}

// SetSoftwareId gets a reference to the given string and assigns it to the SoftwareId field.
func (o *Client) SetSoftwareId(v string) {
	o.SoftwareId = &v
}

// GetSoftwareVersion returns the SoftwareVersion field value if set, zero value otherwise.
func (o *Client) GetSoftwareVersion() string {
	if o == nil || IsNil(o.SoftwareVersion) {
		var ret string
		return ret
	}
	return *o.SoftwareVersion
}

// GetSoftwareVersionOk returns a tuple with the SoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetSoftwareVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SoftwareVersion) {
		return nil, false
	}
	return o.SoftwareVersion, true
}

// HasSoftwareVersion returns a boolean if a field has been set.
func (o *Client) HasSoftwareVersion() bool {
	if o != nil && !IsNil(o.SoftwareVersion) {
		return true
	}

	return false
}

// SetSoftwareVersion gets a reference to the given string and assigns it to the SoftwareVersion field.
func (o *Client) SetSoftwareVersion(v string) {
	o.SoftwareVersion = &v
}

// GetRegistrationAccessTokenHash returns the RegistrationAccessTokenHash field value if set, zero value otherwise.
func (o *Client) GetRegistrationAccessTokenHash() string {
	if o == nil || IsNil(o.RegistrationAccessTokenHash) {
		var ret string
		return ret
	}
	return *o.RegistrationAccessTokenHash
}

// GetRegistrationAccessTokenHashOk returns a tuple with the RegistrationAccessTokenHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetRegistrationAccessTokenHashOk() (*string, bool) {
	if o == nil || IsNil(o.RegistrationAccessTokenHash) {
		return nil, false
	}
	return o.RegistrationAccessTokenHash, true
}

// HasRegistrationAccessTokenHash returns a boolean if a field has been set.
func (o *Client) HasRegistrationAccessTokenHash() bool {
	if o != nil && !IsNil(o.RegistrationAccessTokenHash) {
		return true
	}

	return false
}

// SetRegistrationAccessTokenHash gets a reference to the given string and assigns it to the RegistrationAccessTokenHash field.
func (o *Client) SetRegistrationAccessTokenHash(v string) {
	o.RegistrationAccessTokenHash = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Client) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Client) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *Client) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *Client) GetModifiedAt() int64 {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret int64
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetModifiedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *Client) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given int64 and assigns it to the ModifiedAt field.
func (o *Client) SetModifiedAt(v int64) {
	o.ModifiedAt = &v
}

// GetGrantTypes returns the GrantTypes field value if set, zero value otherwise.
func (o *Client) GetGrantTypes() []GrantType {
	if o == nil || IsNil(o.GrantTypes) {
		var ret []GrantType
		return ret
	}
	return o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetGrantTypesOk() ([]GrantType, bool) {
	if o == nil || IsNil(o.GrantTypes) {
		return nil, false
	}
	return o.GrantTypes, true
}

// HasGrantTypes returns a boolean if a field has been set.
func (o *Client) HasGrantTypes() bool {
	if o != nil && !IsNil(o.GrantTypes) {
		return true
	}

	return false
}

// SetGrantTypes gets a reference to the given []GrantType and assigns it to the GrantTypes field.
func (o *Client) SetGrantTypes(v []GrantType) {
	o.GrantTypes = v
}

// GetResponseTypes returns the ResponseTypes field value if set, zero value otherwise.
func (o *Client) GetResponseTypes() []ResponseType {
	if o == nil || IsNil(o.ResponseTypes) {
		var ret []ResponseType
		return ret
	}
	return o.ResponseTypes
}

// GetResponseTypesOk returns a tuple with the ResponseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetResponseTypesOk() ([]ResponseType, bool) {
	if o == nil || IsNil(o.ResponseTypes) {
		return nil, false
	}
	return o.ResponseTypes, true
}

// HasResponseTypes returns a boolean if a field has been set.
func (o *Client) HasResponseTypes() bool {
	if o != nil && !IsNil(o.ResponseTypes) {
		return true
	}

	return false
}

// SetResponseTypes gets a reference to the given []ResponseType and assigns it to the ResponseTypes field.
func (o *Client) SetResponseTypes(v []ResponseType) {
	o.ResponseTypes = v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *Client) GetRedirectUris() []string {
	if o == nil || IsNil(o.RedirectUris) {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetRedirectUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.RedirectUris) {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *Client) HasRedirectUris() bool {
	if o != nil && !IsNil(o.RedirectUris) {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *Client) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetAuthorizationSignAlg returns the AuthorizationSignAlg field value if set, zero value otherwise.
func (o *Client) GetAuthorizationSignAlg() JwsAlg {
	if o == nil || IsNil(o.AuthorizationSignAlg) {
		var ret JwsAlg
		return ret
	}
	return *o.AuthorizationSignAlg
}

// GetAuthorizationSignAlgOk returns a tuple with the AuthorizationSignAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetAuthorizationSignAlgOk() (*JwsAlg, bool) {
	if o == nil || IsNil(o.AuthorizationSignAlg) {
		return nil, false
	}
	return o.AuthorizationSignAlg, true
}

// HasAuthorizationSignAlg returns a boolean if a field has been set.
func (o *Client) HasAuthorizationSignAlg() bool {
	if o != nil && !IsNil(o.AuthorizationSignAlg) {
		return true
	}

	return false
}

// SetAuthorizationSignAlg gets a reference to the given JwsAlg and assigns it to the AuthorizationSignAlg field.
func (o *Client) SetAuthorizationSignAlg(v JwsAlg) {
	o.AuthorizationSignAlg = &v
}

// GetAuthorizationEncryptionAlg returns the AuthorizationEncryptionAlg field value if set, zero value otherwise.
func (o *Client) GetAuthorizationEncryptionAlg() JweAlg {
	if o == nil || IsNil(o.AuthorizationEncryptionAlg) {
		var ret JweAlg
		return ret
	}
	return *o.AuthorizationEncryptionAlg
}

// GetAuthorizationEncryptionAlgOk returns a tuple with the AuthorizationEncryptionAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetAuthorizationEncryptionAlgOk() (*JweAlg, bool) {
	if o == nil || IsNil(o.AuthorizationEncryptionAlg) {
		return nil, false
	}
	return o.AuthorizationEncryptionAlg, true
}

// HasAuthorizationEncryptionAlg returns a boolean if a field has been set.
func (o *Client) HasAuthorizationEncryptionAlg() bool {
	if o != nil && !IsNil(o.AuthorizationEncryptionAlg) {
		return true
	}

	return false
}

// SetAuthorizationEncryptionAlg gets a reference to the given JweAlg and assigns it to the AuthorizationEncryptionAlg field.
func (o *Client) SetAuthorizationEncryptionAlg(v JweAlg) {
	o.AuthorizationEncryptionAlg = &v
}

// GetAuthorizationEncryptionEnc returns the AuthorizationEncryptionEnc field value if set, zero value otherwise.
func (o *Client) GetAuthorizationEncryptionEnc() JweEnc {
	if o == nil || IsNil(o.AuthorizationEncryptionEnc) {
		var ret JweEnc
		return ret
	}
	return *o.AuthorizationEncryptionEnc
}

// GetAuthorizationEncryptionEncOk returns a tuple with the AuthorizationEncryptionEnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetAuthorizationEncryptionEncOk() (*JweEnc, bool) {
	if o == nil || IsNil(o.AuthorizationEncryptionEnc) {
		return nil, false
	}
	return o.AuthorizationEncryptionEnc, true
}

// HasAuthorizationEncryptionEnc returns a boolean if a field has been set.
func (o *Client) HasAuthorizationEncryptionEnc() bool {
	if o != nil && !IsNil(o.AuthorizationEncryptionEnc) {
		return true
	}

	return false
}

// SetAuthorizationEncryptionEnc gets a reference to the given JweEnc and assigns it to the AuthorizationEncryptionEnc field.
func (o *Client) SetAuthorizationEncryptionEnc(v JweEnc) {
	o.AuthorizationEncryptionEnc = &v
}

// GetTokenAuthMethod returns the TokenAuthMethod field value if set, zero value otherwise.
func (o *Client) GetTokenAuthMethod() ClientAuthMethod {
	if o == nil || IsNil(o.TokenAuthMethod) {
		var ret ClientAuthMethod
		return ret
	}
	return *o.TokenAuthMethod
}

// GetTokenAuthMethodOk returns a tuple with the TokenAuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetTokenAuthMethodOk() (*ClientAuthMethod, bool) {
	if o == nil || IsNil(o.TokenAuthMethod) {
		return nil, false
	}
	return o.TokenAuthMethod, true
}

// HasTokenAuthMethod returns a boolean if a field has been set.
func (o *Client) HasTokenAuthMethod() bool {
	if o != nil && !IsNil(o.TokenAuthMethod) {
		return true
	}

	return false
}

// SetTokenAuthMethod gets a reference to the given ClientAuthMethod and assigns it to the TokenAuthMethod field.
func (o *Client) SetTokenAuthMethod(v ClientAuthMethod) {
	o.TokenAuthMethod = &v
}

// GetTokenAuthSignAlg returns the TokenAuthSignAlg field value if set, zero value otherwise.
func (o *Client) GetTokenAuthSignAlg() JwsAlg {
	if o == nil || IsNil(o.TokenAuthSignAlg) {
		var ret JwsAlg
		return ret
	}
	return *o.TokenAuthSignAlg
}

// GetTokenAuthSignAlgOk returns a tuple with the TokenAuthSignAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetTokenAuthSignAlgOk() (*JwsAlg, bool) {
	if o == nil || IsNil(o.TokenAuthSignAlg) {
		return nil, false
	}
	return o.TokenAuthSignAlg, true
}

// HasTokenAuthSignAlg returns a boolean if a field has been set.
func (o *Client) HasTokenAuthSignAlg() bool {
	if o != nil && !IsNil(o.TokenAuthSignAlg) {
		return true
	}

	return false
}

// SetTokenAuthSignAlg gets a reference to the given JwsAlg and assigns it to the TokenAuthSignAlg field.
func (o *Client) SetTokenAuthSignAlg(v JwsAlg) {
	o.TokenAuthSignAlg = &v
}

// GetSelfSignedCertificateKeyId returns the SelfSignedCertificateKeyId field value if set, zero value otherwise.
func (o *Client) GetSelfSignedCertificateKeyId() string {
	if o == nil || IsNil(o.SelfSignedCertificateKeyId) {
		var ret string
		return ret
	}
	return *o.SelfSignedCertificateKeyId
}

// GetSelfSignedCertificateKeyIdOk returns a tuple with the SelfSignedCertificateKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetSelfSignedCertificateKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.SelfSignedCertificateKeyId) {
		return nil, false
	}
	return o.SelfSignedCertificateKeyId, true
}

// HasSelfSignedCertificateKeyId returns a boolean if a field has been set.
func (o *Client) HasSelfSignedCertificateKeyId() bool {
	if o != nil && !IsNil(o.SelfSignedCertificateKeyId) {
		return true
	}

	return false
}

// SetSelfSignedCertificateKeyId gets a reference to the given string and assigns it to the SelfSignedCertificateKeyId field.
func (o *Client) SetSelfSignedCertificateKeyId(v string) {
	o.SelfSignedCertificateKeyId = &v
}

// GetTlsClientAuthSubjectDn returns the TlsClientAuthSubjectDn field value if set, zero value otherwise.
func (o *Client) GetTlsClientAuthSubjectDn() string {
	if o == nil || IsNil(o.TlsClientAuthSubjectDn) {
		var ret string
		return ret
	}
	return *o.TlsClientAuthSubjectDn
}

// GetTlsClientAuthSubjectDnOk returns a tuple with the TlsClientAuthSubjectDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetTlsClientAuthSubjectDnOk() (*string, bool) {
	if o == nil || IsNil(o.TlsClientAuthSubjectDn) {
		return nil, false
	}
	return o.TlsClientAuthSubjectDn, true
}

// HasTlsClientAuthSubjectDn returns a boolean if a field has been set.
func (o *Client) HasTlsClientAuthSubjectDn() bool {
	if o != nil && !IsNil(o.TlsClientAuthSubjectDn) {
		return true
	}

	return false
}

// SetTlsClientAuthSubjectDn gets a reference to the given string and assigns it to the TlsClientAuthSubjectDn field.
func (o *Client) SetTlsClientAuthSubjectDn(v string) {
	o.TlsClientAuthSubjectDn = &v
}

// GetTlsClientAuthSanDns returns the TlsClientAuthSanDns field value if set, zero value otherwise.
func (o *Client) GetTlsClientAuthSanDns() string {
	if o == nil || IsNil(o.TlsClientAuthSanDns) {
		var ret string
		return ret
	}
	return *o.TlsClientAuthSanDns
}

// GetTlsClientAuthSanDnsOk returns a tuple with the TlsClientAuthSanDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetTlsClientAuthSanDnsOk() (*string, bool) {
	if o == nil || IsNil(o.TlsClientAuthSanDns) {
		return nil, false
	}
	return o.TlsClientAuthSanDns, true
}

// HasTlsClientAuthSanDns returns a boolean if a field has been set.
func (o *Client) HasTlsClientAuthSanDns() bool {
	if o != nil && !IsNil(o.TlsClientAuthSanDns) {
		return true
	}

	return false
}

// SetTlsClientAuthSanDns gets a reference to the given string and assigns it to the TlsClientAuthSanDns field.
func (o *Client) SetTlsClientAuthSanDns(v string) {
	o.TlsClientAuthSanDns = &v
}

// GetTlsClientAuthSanUri returns the TlsClientAuthSanUri field value if set, zero value otherwise.
func (o *Client) GetTlsClientAuthSanUri() string {
	if o == nil || IsNil(o.TlsClientAuthSanUri) {
		var ret string
		return ret
	}
	return *o.TlsClientAuthSanUri
}

// GetTlsClientAuthSanUriOk returns a tuple with the TlsClientAuthSanUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetTlsClientAuthSanUriOk() (*string, bool) {
	if o == nil || IsNil(o.TlsClientAuthSanUri) {
		return nil, false
	}
	return o.TlsClientAuthSanUri, true
}

// HasTlsClientAuthSanUri returns a boolean if a field has been set.
func (o *Client) HasTlsClientAuthSanUri() bool {
	if o != nil && !IsNil(o.TlsClientAuthSanUri) {
		return true
	}

	return false
}

// SetTlsClientAuthSanUri gets a reference to the given string and assigns it to the TlsClientAuthSanUri field.
func (o *Client) SetTlsClientAuthSanUri(v string) {
	o.TlsClientAuthSanUri = &v
}

// GetTlsClientAuthSanIp returns the TlsClientAuthSanIp field value if set, zero value otherwise.
func (o *Client) GetTlsClientAuthSanIp() string {
	if o == nil || IsNil(o.TlsClientAuthSanIp) {
		var ret string
		return ret
	}
	return *o.TlsClientAuthSanIp
}

// GetTlsClientAuthSanIpOk returns a tuple with the TlsClientAuthSanIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetTlsClientAuthSanIpOk() (*string, bool) {
	if o == nil || IsNil(o.TlsClientAuthSanIp) {
		return nil, false
	}
	return o.TlsClientAuthSanIp, true
}

// HasTlsClientAuthSanIp returns a boolean if a field has been set.
func (o *Client) HasTlsClientAuthSanIp() bool {
	if o != nil && !IsNil(o.TlsClientAuthSanIp) {
		return true
	}

	return false
}

// SetTlsClientAuthSanIp gets a reference to the given string and assigns it to the TlsClientAuthSanIp field.
func (o *Client) SetTlsClientAuthSanIp(v string) {
	o.TlsClientAuthSanIp = &v
}

// GetTlsClientAuthSanEmail returns the TlsClientAuthSanEmail field value if set, zero value otherwise.
func (o *Client) GetTlsClientAuthSanEmail() string {
	if o == nil || IsNil(o.TlsClientAuthSanEmail) {
		var ret string
		return ret
	}
	return *o.TlsClientAuthSanEmail
}

// GetTlsClientAuthSanEmailOk returns a tuple with the TlsClientAuthSanEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetTlsClientAuthSanEmailOk() (*string, bool) {
	if o == nil || IsNil(o.TlsClientAuthSanEmail) {
		return nil, false
	}
	return o.TlsClientAuthSanEmail, true
}

// HasTlsClientAuthSanEmail returns a boolean if a field has been set.
func (o *Client) HasTlsClientAuthSanEmail() bool {
	if o != nil && !IsNil(o.TlsClientAuthSanEmail) {
		return true
	}

	return false
}

// SetTlsClientAuthSanEmail gets a reference to the given string and assigns it to the TlsClientAuthSanEmail field.
func (o *Client) SetTlsClientAuthSanEmail(v string) {
	o.TlsClientAuthSanEmail = &v
}

// GetParRequired returns the ParRequired field value if set, zero value otherwise.
func (o *Client) GetParRequired() bool {
	if o == nil || IsNil(o.ParRequired) {
		var ret bool
		return ret
	}
	return *o.ParRequired
}

// GetParRequiredOk returns a tuple with the ParRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetParRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ParRequired) {
		return nil, false
	}
	return o.ParRequired, true
}

// HasParRequired returns a boolean if a field has been set.
func (o *Client) HasParRequired() bool {
	if o != nil && !IsNil(o.ParRequired) {
		return true
	}

	return false
}

// SetParRequired gets a reference to the given bool and assigns it to the ParRequired field.
func (o *Client) SetParRequired(v bool) {
	o.ParRequired = &v
}

// GetRequestObjectRequired returns the RequestObjectRequired field value if set, zero value otherwise.
func (o *Client) GetRequestObjectRequired() bool {
	if o == nil || IsNil(o.RequestObjectRequired) {
		var ret bool
		return ret
	}
	return *o.RequestObjectRequired
}

// GetRequestObjectRequiredOk returns a tuple with the RequestObjectRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetRequestObjectRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestObjectRequired) {
		return nil, false
	}
	return o.RequestObjectRequired, true
}

// HasRequestObjectRequired returns a boolean if a field has been set.
func (o *Client) HasRequestObjectRequired() bool {
	if o != nil && !IsNil(o.RequestObjectRequired) {
		return true
	}

	return false
}

// SetRequestObjectRequired gets a reference to the given bool and assigns it to the RequestObjectRequired field.
func (o *Client) SetRequestObjectRequired(v bool) {
	o.RequestObjectRequired = &v
}

// GetRequestSignAlg returns the RequestSignAlg field value if set, zero value otherwise.
func (o *Client) GetRequestSignAlg() JwsAlg {
	if o == nil || IsNil(o.RequestSignAlg) {
		var ret JwsAlg
		return ret
	}
	return *o.RequestSignAlg
}

// GetRequestSignAlgOk returns a tuple with the RequestSignAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetRequestSignAlgOk() (*JwsAlg, bool) {
	if o == nil || IsNil(o.RequestSignAlg) {
		return nil, false
	}
	return o.RequestSignAlg, true
}

// HasRequestSignAlg returns a boolean if a field has been set.
func (o *Client) HasRequestSignAlg() bool {
	if o != nil && !IsNil(o.RequestSignAlg) {
		return true
	}

	return false
}

// SetRequestSignAlg gets a reference to the given JwsAlg and assigns it to the RequestSignAlg field.
func (o *Client) SetRequestSignAlg(v JwsAlg) {
	o.RequestSignAlg = &v
}

// GetRequestEncryptionAlg returns the RequestEncryptionAlg field value if set, zero value otherwise.
func (o *Client) GetRequestEncryptionAlg() JweAlg {
	if o == nil || IsNil(o.RequestEncryptionAlg) {
		var ret JweAlg
		return ret
	}
	return *o.RequestEncryptionAlg
}

// GetRequestEncryptionAlgOk returns a tuple with the RequestEncryptionAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetRequestEncryptionAlgOk() (*JweAlg, bool) {
	if o == nil || IsNil(o.RequestEncryptionAlg) {
		return nil, false
	}
	return o.RequestEncryptionAlg, true
}

// HasRequestEncryptionAlg returns a boolean if a field has been set.
func (o *Client) HasRequestEncryptionAlg() bool {
	if o != nil && !IsNil(o.RequestEncryptionAlg) {
		return true
	}

	return false
}

// SetRequestEncryptionAlg gets a reference to the given JweAlg and assigns it to the RequestEncryptionAlg field.
func (o *Client) SetRequestEncryptionAlg(v JweAlg) {
	o.RequestEncryptionAlg = &v
}

// GetRequestEncryptionEnc returns the RequestEncryptionEnc field value if set, zero value otherwise.
func (o *Client) GetRequestEncryptionEnc() JweEnc {
	if o == nil || IsNil(o.RequestEncryptionEnc) {
		var ret JweEnc
		return ret
	}
	return *o.RequestEncryptionEnc
}

// GetRequestEncryptionEncOk returns a tuple with the RequestEncryptionEnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetRequestEncryptionEncOk() (*JweEnc, bool) {
	if o == nil || IsNil(o.RequestEncryptionEnc) {
		return nil, false
	}
	return o.RequestEncryptionEnc, true
}

// HasRequestEncryptionEnc returns a boolean if a field has been set.
func (o *Client) HasRequestEncryptionEnc() bool {
	if o != nil && !IsNil(o.RequestEncryptionEnc) {
		return true
	}

	return false
}

// SetRequestEncryptionEnc gets a reference to the given JweEnc and assigns it to the RequestEncryptionEnc field.
func (o *Client) SetRequestEncryptionEnc(v JweEnc) {
	o.RequestEncryptionEnc = &v
}

// GetRequestUris returns the RequestUris field value if set, zero value otherwise.
func (o *Client) GetRequestUris() []string {
	if o == nil || IsNil(o.RequestUris) {
		var ret []string
		return ret
	}
	return o.RequestUris
}

// GetRequestUrisOk returns a tuple with the RequestUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetRequestUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.RequestUris) {
		return nil, false
	}
	return o.RequestUris, true
}

// HasRequestUris returns a boolean if a field has been set.
func (o *Client) HasRequestUris() bool {
	if o != nil && !IsNil(o.RequestUris) {
		return true
	}

	return false
}

// SetRequestUris gets a reference to the given []string and assigns it to the RequestUris field.
func (o *Client) SetRequestUris(v []string) {
	o.RequestUris = v
}

// GetDefaultMaxAge returns the DefaultMaxAge field value if set, zero value otherwise.
func (o *Client) GetDefaultMaxAge() int32 {
	if o == nil || IsNil(o.DefaultMaxAge) {
		var ret int32
		return ret
	}
	return *o.DefaultMaxAge
}

// GetDefaultMaxAgeOk returns a tuple with the DefaultMaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetDefaultMaxAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultMaxAge) {
		return nil, false
	}
	return o.DefaultMaxAge, true
}

// HasDefaultMaxAge returns a boolean if a field has been set.
func (o *Client) HasDefaultMaxAge() bool {
	if o != nil && !IsNil(o.DefaultMaxAge) {
		return true
	}

	return false
}

// SetDefaultMaxAge gets a reference to the given int32 and assigns it to the DefaultMaxAge field.
func (o *Client) SetDefaultMaxAge(v int32) {
	o.DefaultMaxAge = &v
}

// GetDefaultAcrs returns the DefaultAcrs field value if set, zero value otherwise.
func (o *Client) GetDefaultAcrs() []string {
	if o == nil || IsNil(o.DefaultAcrs) {
		var ret []string
		return ret
	}
	return o.DefaultAcrs
}

// GetDefaultAcrsOk returns a tuple with the DefaultAcrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetDefaultAcrsOk() ([]string, bool) {
	if o == nil || IsNil(o.DefaultAcrs) {
		return nil, false
	}
	return o.DefaultAcrs, true
}

// HasDefaultAcrs returns a boolean if a field has been set.
func (o *Client) HasDefaultAcrs() bool {
	if o != nil && !IsNil(o.DefaultAcrs) {
		return true
	}

	return false
}

// SetDefaultAcrs gets a reference to the given []string and assigns it to the DefaultAcrs field.
func (o *Client) SetDefaultAcrs(v []string) {
	o.DefaultAcrs = v
}

// GetIdTokenSignAlg returns the IdTokenSignAlg field value if set, zero value otherwise.
func (o *Client) GetIdTokenSignAlg() JwsAlg {
	if o == nil || IsNil(o.IdTokenSignAlg) {
		var ret JwsAlg
		return ret
	}
	return *o.IdTokenSignAlg
}

// GetIdTokenSignAlgOk returns a tuple with the IdTokenSignAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetIdTokenSignAlgOk() (*JwsAlg, bool) {
	if o == nil || IsNil(o.IdTokenSignAlg) {
		return nil, false
	}
	return o.IdTokenSignAlg, true
}

// HasIdTokenSignAlg returns a boolean if a field has been set.
func (o *Client) HasIdTokenSignAlg() bool {
	if o != nil && !IsNil(o.IdTokenSignAlg) {
		return true
	}

	return false
}

// SetIdTokenSignAlg gets a reference to the given JwsAlg and assigns it to the IdTokenSignAlg field.
func (o *Client) SetIdTokenSignAlg(v JwsAlg) {
	o.IdTokenSignAlg = &v
}

// GetIdTokenEncryptionAlg returns the IdTokenEncryptionAlg field value if set, zero value otherwise.
func (o *Client) GetIdTokenEncryptionAlg() JweAlg {
	if o == nil || IsNil(o.IdTokenEncryptionAlg) {
		var ret JweAlg
		return ret
	}
	return *o.IdTokenEncryptionAlg
}

// GetIdTokenEncryptionAlgOk returns a tuple with the IdTokenEncryptionAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetIdTokenEncryptionAlgOk() (*JweAlg, bool) {
	if o == nil || IsNil(o.IdTokenEncryptionAlg) {
		return nil, false
	}
	return o.IdTokenEncryptionAlg, true
}

// HasIdTokenEncryptionAlg returns a boolean if a field has been set.
func (o *Client) HasIdTokenEncryptionAlg() bool {
	if o != nil && !IsNil(o.IdTokenEncryptionAlg) {
		return true
	}

	return false
}

// SetIdTokenEncryptionAlg gets a reference to the given JweAlg and assigns it to the IdTokenEncryptionAlg field.
func (o *Client) SetIdTokenEncryptionAlg(v JweAlg) {
	o.IdTokenEncryptionAlg = &v
}

// GetIdTokenEncryptionEnc returns the IdTokenEncryptionEnc field value if set, zero value otherwise.
func (o *Client) GetIdTokenEncryptionEnc() JweEnc {
	if o == nil || IsNil(o.IdTokenEncryptionEnc) {
		var ret JweEnc
		return ret
	}
	return *o.IdTokenEncryptionEnc
}

// GetIdTokenEncryptionEncOk returns a tuple with the IdTokenEncryptionEnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetIdTokenEncryptionEncOk() (*JweEnc, bool) {
	if o == nil || IsNil(o.IdTokenEncryptionEnc) {
		return nil, false
	}
	return o.IdTokenEncryptionEnc, true
}

// HasIdTokenEncryptionEnc returns a boolean if a field has been set.
func (o *Client) HasIdTokenEncryptionEnc() bool {
	if o != nil && !IsNil(o.IdTokenEncryptionEnc) {
		return true
	}

	return false
}

// SetIdTokenEncryptionEnc gets a reference to the given JweEnc and assigns it to the IdTokenEncryptionEnc field.
func (o *Client) SetIdTokenEncryptionEnc(v JweEnc) {
	o.IdTokenEncryptionEnc = &v
}

// GetAuthTimeRequired returns the AuthTimeRequired field value if set, zero value otherwise.
func (o *Client) GetAuthTimeRequired() bool {
	if o == nil || IsNil(o.AuthTimeRequired) {
		var ret bool
		return ret
	}
	return *o.AuthTimeRequired
}

// GetAuthTimeRequiredOk returns a tuple with the AuthTimeRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetAuthTimeRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.AuthTimeRequired) {
		return nil, false
	}
	return o.AuthTimeRequired, true
}

// HasAuthTimeRequired returns a boolean if a field has been set.
func (o *Client) HasAuthTimeRequired() bool {
	if o != nil && !IsNil(o.AuthTimeRequired) {
		return true
	}

	return false
}

// SetAuthTimeRequired gets a reference to the given bool and assigns it to the AuthTimeRequired field.
func (o *Client) SetAuthTimeRequired(v bool) {
	o.AuthTimeRequired = &v
}

// GetSubjectType returns the SubjectType field value if set, zero value otherwise.
func (o *Client) GetSubjectType() SubjectType {
	if o == nil || IsNil(o.SubjectType) {
		var ret SubjectType
		return ret
	}
	return *o.SubjectType
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetSubjectTypeOk() (*SubjectType, bool) {
	if o == nil || IsNil(o.SubjectType) {
		return nil, false
	}
	return o.SubjectType, true
}

// HasSubjectType returns a boolean if a field has been set.
func (o *Client) HasSubjectType() bool {
	if o != nil && !IsNil(o.SubjectType) {
		return true
	}

	return false
}

// SetSubjectType gets a reference to the given SubjectType and assigns it to the SubjectType field.
func (o *Client) SetSubjectType(v SubjectType) {
	o.SubjectType = &v
}

// GetSectorIdentifierUri returns the SectorIdentifierUri field value if set, zero value otherwise.
func (o *Client) GetSectorIdentifierUri() string {
	if o == nil || IsNil(o.SectorIdentifierUri) {
		var ret string
		return ret
	}
	return *o.SectorIdentifierUri
}

// GetSectorIdentifierUriOk returns a tuple with the SectorIdentifierUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetSectorIdentifierUriOk() (*string, bool) {
	if o == nil || IsNil(o.SectorIdentifierUri) {
		return nil, false
	}
	return o.SectorIdentifierUri, true
}

// HasSectorIdentifierUri returns a boolean if a field has been set.
func (o *Client) HasSectorIdentifierUri() bool {
	if o != nil && !IsNil(o.SectorIdentifierUri) {
		return true
	}

	return false
}

// SetSectorIdentifierUri gets a reference to the given string and assigns it to the SectorIdentifierUri field.
func (o *Client) SetSectorIdentifierUri(v string) {
	o.SectorIdentifierUri = &v
}

// GetDerivedSectorIdentifier returns the DerivedSectorIdentifier field value if set, zero value otherwise.
func (o *Client) GetDerivedSectorIdentifier() string {
	if o == nil || IsNil(o.DerivedSectorIdentifier) {
		var ret string
		return ret
	}
	return *o.DerivedSectorIdentifier
}

// GetDerivedSectorIdentifierOk returns a tuple with the DerivedSectorIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetDerivedSectorIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.DerivedSectorIdentifier) {
		return nil, false
	}
	return o.DerivedSectorIdentifier, true
}

// HasDerivedSectorIdentifier returns a boolean if a field has been set.
func (o *Client) HasDerivedSectorIdentifier() bool {
	if o != nil && !IsNil(o.DerivedSectorIdentifier) {
		return true
	}

	return false
}

// SetDerivedSectorIdentifier gets a reference to the given string and assigns it to the DerivedSectorIdentifier field.
func (o *Client) SetDerivedSectorIdentifier(v string) {
	o.DerivedSectorIdentifier = &v
}

// GetJwksUri returns the JwksUri field value if set, zero value otherwise.
func (o *Client) GetJwksUri() string {
	if o == nil || IsNil(o.JwksUri) {
		var ret string
		return ret
	}
	return *o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetJwksUriOk() (*string, bool) {
	if o == nil || IsNil(o.JwksUri) {
		return nil, false
	}
	return o.JwksUri, true
}

// HasJwksUri returns a boolean if a field has been set.
func (o *Client) HasJwksUri() bool {
	if o != nil && !IsNil(o.JwksUri) {
		return true
	}

	return false
}

// SetJwksUri gets a reference to the given string and assigns it to the JwksUri field.
func (o *Client) SetJwksUri(v string) {
	o.JwksUri = &v
}

// GetJwks returns the Jwks field value if set, zero value otherwise.
func (o *Client) GetJwks() string {
	if o == nil || IsNil(o.Jwks) {
		var ret string
		return ret
	}
	return *o.Jwks
}

// GetJwksOk returns a tuple with the Jwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetJwksOk() (*string, bool) {
	if o == nil || IsNil(o.Jwks) {
		return nil, false
	}
	return o.Jwks, true
}

// HasJwks returns a boolean if a field has been set.
func (o *Client) HasJwks() bool {
	if o != nil && !IsNil(o.Jwks) {
		return true
	}

	return false
}

// SetJwks gets a reference to the given string and assigns it to the Jwks field.
func (o *Client) SetJwks(v string) {
	o.Jwks = &v
}

// GetUserInfoSignAlg returns the UserInfoSignAlg field value if set, zero value otherwise.
func (o *Client) GetUserInfoSignAlg() JwsAlg {
	if o == nil || IsNil(o.UserInfoSignAlg) {
		var ret JwsAlg
		return ret
	}
	return *o.UserInfoSignAlg
}

// GetUserInfoSignAlgOk returns a tuple with the UserInfoSignAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetUserInfoSignAlgOk() (*JwsAlg, bool) {
	if o == nil || IsNil(o.UserInfoSignAlg) {
		return nil, false
	}
	return o.UserInfoSignAlg, true
}

// HasUserInfoSignAlg returns a boolean if a field has been set.
func (o *Client) HasUserInfoSignAlg() bool {
	if o != nil && !IsNil(o.UserInfoSignAlg) {
		return true
	}

	return false
}

// SetUserInfoSignAlg gets a reference to the given JwsAlg and assigns it to the UserInfoSignAlg field.
func (o *Client) SetUserInfoSignAlg(v JwsAlg) {
	o.UserInfoSignAlg = &v
}

// GetUserInfoEncryptionAlg returns the UserInfoEncryptionAlg field value if set, zero value otherwise.
func (o *Client) GetUserInfoEncryptionAlg() JweAlg {
	if o == nil || IsNil(o.UserInfoEncryptionAlg) {
		var ret JweAlg
		return ret
	}
	return *o.UserInfoEncryptionAlg
}

// GetUserInfoEncryptionAlgOk returns a tuple with the UserInfoEncryptionAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetUserInfoEncryptionAlgOk() (*JweAlg, bool) {
	if o == nil || IsNil(o.UserInfoEncryptionAlg) {
		return nil, false
	}
	return o.UserInfoEncryptionAlg, true
}

// HasUserInfoEncryptionAlg returns a boolean if a field has been set.
func (o *Client) HasUserInfoEncryptionAlg() bool {
	if o != nil && !IsNil(o.UserInfoEncryptionAlg) {
		return true
	}

	return false
}

// SetUserInfoEncryptionAlg gets a reference to the given JweAlg and assigns it to the UserInfoEncryptionAlg field.
func (o *Client) SetUserInfoEncryptionAlg(v JweAlg) {
	o.UserInfoEncryptionAlg = &v
}

// GetUserInfoEncryptionEnc returns the UserInfoEncryptionEnc field value if set, zero value otherwise.
func (o *Client) GetUserInfoEncryptionEnc() JweEnc {
	if o == nil || IsNil(o.UserInfoEncryptionEnc) {
		var ret JweEnc
		return ret
	}
	return *o.UserInfoEncryptionEnc
}

// GetUserInfoEncryptionEncOk returns a tuple with the UserInfoEncryptionEnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetUserInfoEncryptionEncOk() (*JweEnc, bool) {
	if o == nil || IsNil(o.UserInfoEncryptionEnc) {
		return nil, false
	}
	return o.UserInfoEncryptionEnc, true
}

// HasUserInfoEncryptionEnc returns a boolean if a field has been set.
func (o *Client) HasUserInfoEncryptionEnc() bool {
	if o != nil && !IsNil(o.UserInfoEncryptionEnc) {
		return true
	}

	return false
}

// SetUserInfoEncryptionEnc gets a reference to the given JweEnc and assigns it to the UserInfoEncryptionEnc field.
func (o *Client) SetUserInfoEncryptionEnc(v JweEnc) {
	o.UserInfoEncryptionEnc = &v
}

// GetLoginUri returns the LoginUri field value if set, zero value otherwise.
func (o *Client) GetLoginUri() string {
	if o == nil || IsNil(o.LoginUri) {
		var ret string
		return ret
	}
	return *o.LoginUri
}

// GetLoginUriOk returns a tuple with the LoginUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetLoginUriOk() (*string, bool) {
	if o == nil || IsNil(o.LoginUri) {
		return nil, false
	}
	return o.LoginUri, true
}

// HasLoginUri returns a boolean if a field has been set.
func (o *Client) HasLoginUri() bool {
	if o != nil && !IsNil(o.LoginUri) {
		return true
	}

	return false
}

// SetLoginUri gets a reference to the given string and assigns it to the LoginUri field.
func (o *Client) SetLoginUri(v string) {
	o.LoginUri = &v
}

// GetTosUri returns the TosUri field value if set, zero value otherwise.
func (o *Client) GetTosUri() string {
	if o == nil || IsNil(o.TosUri) {
		var ret string
		return ret
	}
	return *o.TosUri
}

// GetTosUriOk returns a tuple with the TosUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetTosUriOk() (*string, bool) {
	if o == nil || IsNil(o.TosUri) {
		return nil, false
	}
	return o.TosUri, true
}

// HasTosUri returns a boolean if a field has been set.
func (o *Client) HasTosUri() bool {
	if o != nil && !IsNil(o.TosUri) {
		return true
	}

	return false
}

// SetTosUri gets a reference to the given string and assigns it to the TosUri field.
func (o *Client) SetTosUri(v string) {
	o.TosUri = &v
}

// GetTosUris returns the TosUris field value if set, zero value otherwise.
func (o *Client) GetTosUris() []TaggedValue {
	if o == nil || IsNil(o.TosUris) {
		var ret []TaggedValue
		return ret
	}
	return o.TosUris
}

// GetTosUrisOk returns a tuple with the TosUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetTosUrisOk() ([]TaggedValue, bool) {
	if o == nil || IsNil(o.TosUris) {
		return nil, false
	}
	return o.TosUris, true
}

// HasTosUris returns a boolean if a field has been set.
func (o *Client) HasTosUris() bool {
	if o != nil && !IsNil(o.TosUris) {
		return true
	}

	return false
}

// SetTosUris gets a reference to the given []TaggedValue and assigns it to the TosUris field.
func (o *Client) SetTosUris(v []TaggedValue) {
	o.TosUris = v
}

// GetPolicyUri returns the PolicyUri field value if set, zero value otherwise.
func (o *Client) GetPolicyUri() string {
	if o == nil || IsNil(o.PolicyUri) {
		var ret string
		return ret
	}
	return *o.PolicyUri
}

// GetPolicyUriOk returns a tuple with the PolicyUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetPolicyUriOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyUri) {
		return nil, false
	}
	return o.PolicyUri, true
}

// HasPolicyUri returns a boolean if a field has been set.
func (o *Client) HasPolicyUri() bool {
	if o != nil && !IsNil(o.PolicyUri) {
		return true
	}

	return false
}

// SetPolicyUri gets a reference to the given string and assigns it to the PolicyUri field.
func (o *Client) SetPolicyUri(v string) {
	o.PolicyUri = &v
}

// GetPolicyUris returns the PolicyUris field value if set, zero value otherwise.
func (o *Client) GetPolicyUris() []TaggedValue {
	if o == nil || IsNil(o.PolicyUris) {
		var ret []TaggedValue
		return ret
	}
	return o.PolicyUris
}

// GetPolicyUrisOk returns a tuple with the PolicyUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetPolicyUrisOk() ([]TaggedValue, bool) {
	if o == nil || IsNil(o.PolicyUris) {
		return nil, false
	}
	return o.PolicyUris, true
}

// HasPolicyUris returns a boolean if a field has been set.
func (o *Client) HasPolicyUris() bool {
	if o != nil && !IsNil(o.PolicyUris) {
		return true
	}

	return false
}

// SetPolicyUris gets a reference to the given []TaggedValue and assigns it to the PolicyUris field.
func (o *Client) SetPolicyUris(v []TaggedValue) {
	o.PolicyUris = v
}

// GetClientUri returns the ClientUri field value if set, zero value otherwise.
func (o *Client) GetClientUri() string {
	if o == nil || IsNil(o.ClientUri) {
		var ret string
		return ret
	}
	return *o.ClientUri
}

// GetClientUriOk returns a tuple with the ClientUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetClientUriOk() (*string, bool) {
	if o == nil || IsNil(o.ClientUri) {
		return nil, false
	}
	return o.ClientUri, true
}

// HasClientUri returns a boolean if a field has been set.
func (o *Client) HasClientUri() bool {
	if o != nil && !IsNil(o.ClientUri) {
		return true
	}

	return false
}

// SetClientUri gets a reference to the given string and assigns it to the ClientUri field.
func (o *Client) SetClientUri(v string) {
	o.ClientUri = &v
}

// GetClientUris returns the ClientUris field value if set, zero value otherwise.
func (o *Client) GetClientUris() []TaggedValue {
	if o == nil || IsNil(o.ClientUris) {
		var ret []TaggedValue
		return ret
	}
	return o.ClientUris
}

// GetClientUrisOk returns a tuple with the ClientUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetClientUrisOk() ([]TaggedValue, bool) {
	if o == nil || IsNil(o.ClientUris) {
		return nil, false
	}
	return o.ClientUris, true
}

// HasClientUris returns a boolean if a field has been set.
func (o *Client) HasClientUris() bool {
	if o != nil && !IsNil(o.ClientUris) {
		return true
	}

	return false
}

// SetClientUris gets a reference to the given []TaggedValue and assigns it to the ClientUris field.
func (o *Client) SetClientUris(v []TaggedValue) {
	o.ClientUris = v
}

// GetBcDeliveryMode returns the BcDeliveryMode field value if set, zero value otherwise.
func (o *Client) GetBcDeliveryMode() string {
	if o == nil || IsNil(o.BcDeliveryMode) {
		var ret string
		return ret
	}
	return *o.BcDeliveryMode
}

// GetBcDeliveryModeOk returns a tuple with the BcDeliveryMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetBcDeliveryModeOk() (*string, bool) {
	if o == nil || IsNil(o.BcDeliveryMode) {
		return nil, false
	}
	return o.BcDeliveryMode, true
}

// HasBcDeliveryMode returns a boolean if a field has been set.
func (o *Client) HasBcDeliveryMode() bool {
	if o != nil && !IsNil(o.BcDeliveryMode) {
		return true
	}

	return false
}

// SetBcDeliveryMode gets a reference to the given string and assigns it to the BcDeliveryMode field.
func (o *Client) SetBcDeliveryMode(v string) {
	o.BcDeliveryMode = &v
}

// GetBcNotificationEndpoint returns the BcNotificationEndpoint field value if set, zero value otherwise.
func (o *Client) GetBcNotificationEndpoint() string {
	if o == nil || IsNil(o.BcNotificationEndpoint) {
		var ret string
		return ret
	}
	return *o.BcNotificationEndpoint
}

// GetBcNotificationEndpointOk returns a tuple with the BcNotificationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetBcNotificationEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.BcNotificationEndpoint) {
		return nil, false
	}
	return o.BcNotificationEndpoint, true
}

// HasBcNotificationEndpoint returns a boolean if a field has been set.
func (o *Client) HasBcNotificationEndpoint() bool {
	if o != nil && !IsNil(o.BcNotificationEndpoint) {
		return true
	}

	return false
}

// SetBcNotificationEndpoint gets a reference to the given string and assigns it to the BcNotificationEndpoint field.
func (o *Client) SetBcNotificationEndpoint(v string) {
	o.BcNotificationEndpoint = &v
}

// GetBcRequestSignAlg returns the BcRequestSignAlg field value if set, zero value otherwise.
func (o *Client) GetBcRequestSignAlg() JwsAlg {
	if o == nil || IsNil(o.BcRequestSignAlg) {
		var ret JwsAlg
		return ret
	}
	return *o.BcRequestSignAlg
}

// GetBcRequestSignAlgOk returns a tuple with the BcRequestSignAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetBcRequestSignAlgOk() (*JwsAlg, bool) {
	if o == nil || IsNil(o.BcRequestSignAlg) {
		return nil, false
	}
	return o.BcRequestSignAlg, true
}

// HasBcRequestSignAlg returns a boolean if a field has been set.
func (o *Client) HasBcRequestSignAlg() bool {
	if o != nil && !IsNil(o.BcRequestSignAlg) {
		return true
	}

	return false
}

// SetBcRequestSignAlg gets a reference to the given JwsAlg and assigns it to the BcRequestSignAlg field.
func (o *Client) SetBcRequestSignAlg(v JwsAlg) {
	o.BcRequestSignAlg = &v
}

// GetBcUserCodeRequired returns the BcUserCodeRequired field value if set, zero value otherwise.
func (o *Client) GetBcUserCodeRequired() bool {
	if o == nil || IsNil(o.BcUserCodeRequired) {
		var ret bool
		return ret
	}
	return *o.BcUserCodeRequired
}

// GetBcUserCodeRequiredOk returns a tuple with the BcUserCodeRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetBcUserCodeRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.BcUserCodeRequired) {
		return nil, false
	}
	return o.BcUserCodeRequired, true
}

// HasBcUserCodeRequired returns a boolean if a field has been set.
func (o *Client) HasBcUserCodeRequired() bool {
	if o != nil && !IsNil(o.BcUserCodeRequired) {
		return true
	}

	return false
}

// SetBcUserCodeRequired gets a reference to the given bool and assigns it to the BcUserCodeRequired field.
func (o *Client) SetBcUserCodeRequired(v bool) {
	o.BcUserCodeRequired = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Client) GetAttributes() []Pair {
	if o == nil || IsNil(o.Attributes) {
		var ret []Pair
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetAttributesOk() ([]Pair, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Client) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []Pair and assigns it to the Attributes field.
func (o *Client) SetAttributes(v []Pair) {
	o.Attributes = v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *Client) GetExtension() ClientExtension {
	if o == nil || IsNil(o.Extension) {
		var ret ClientExtension
		return ret
	}
	return *o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetExtensionOk() (*ClientExtension, bool) {
	if o == nil || IsNil(o.Extension) {
		return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *Client) HasExtension() bool {
	if o != nil && !IsNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given ClientExtension and assigns it to the Extension field.
func (o *Client) SetExtension(v ClientExtension) {
	o.Extension = &v
}

// GetAuthorizationDetailsTypes returns the AuthorizationDetailsTypes field value if set, zero value otherwise.
func (o *Client) GetAuthorizationDetailsTypes() []string {
	if o == nil || IsNil(o.AuthorizationDetailsTypes) {
		var ret []string
		return ret
	}
	return o.AuthorizationDetailsTypes
}

// GetAuthorizationDetailsTypesOk returns a tuple with the AuthorizationDetailsTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetAuthorizationDetailsTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthorizationDetailsTypes) {
		return nil, false
	}
	return o.AuthorizationDetailsTypes, true
}

// HasAuthorizationDetailsTypes returns a boolean if a field has been set.
func (o *Client) HasAuthorizationDetailsTypes() bool {
	if o != nil && !IsNil(o.AuthorizationDetailsTypes) {
		return true
	}

	return false
}

// SetAuthorizationDetailsTypes gets a reference to the given []string and assigns it to the AuthorizationDetailsTypes field.
func (o *Client) SetAuthorizationDetailsTypes(v []string) {
	o.AuthorizationDetailsTypes = v
}

// GetCustomMetadata returns the CustomMetadata field value if set, zero value otherwise.
func (o *Client) GetCustomMetadata() string {
	if o == nil || IsNil(o.CustomMetadata) {
		var ret string
		return ret
	}
	return *o.CustomMetadata
}

// GetCustomMetadataOk returns a tuple with the CustomMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetCustomMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.CustomMetadata) {
		return nil, false
	}
	return o.CustomMetadata, true
}

// HasCustomMetadata returns a boolean if a field has been set.
func (o *Client) HasCustomMetadata() bool {
	if o != nil && !IsNil(o.CustomMetadata) {
		return true
	}

	return false
}

// SetCustomMetadata gets a reference to the given string and assigns it to the CustomMetadata field.
func (o *Client) SetCustomMetadata(v string) {
	o.CustomMetadata = &v
}

// GetFrontChannelRequestObjectEncryptionRequired returns the FrontChannelRequestObjectEncryptionRequired field value if set, zero value otherwise.
func (o *Client) GetFrontChannelRequestObjectEncryptionRequired() bool {
	if o == nil || IsNil(o.FrontChannelRequestObjectEncryptionRequired) {
		var ret bool
		return ret
	}
	return *o.FrontChannelRequestObjectEncryptionRequired
}

// GetFrontChannelRequestObjectEncryptionRequiredOk returns a tuple with the FrontChannelRequestObjectEncryptionRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetFrontChannelRequestObjectEncryptionRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.FrontChannelRequestObjectEncryptionRequired) {
		return nil, false
	}
	return o.FrontChannelRequestObjectEncryptionRequired, true
}

// HasFrontChannelRequestObjectEncryptionRequired returns a boolean if a field has been set.
func (o *Client) HasFrontChannelRequestObjectEncryptionRequired() bool {
	if o != nil && !IsNil(o.FrontChannelRequestObjectEncryptionRequired) {
		return true
	}

	return false
}

// SetFrontChannelRequestObjectEncryptionRequired gets a reference to the given bool and assigns it to the FrontChannelRequestObjectEncryptionRequired field.
func (o *Client) SetFrontChannelRequestObjectEncryptionRequired(v bool) {
	o.FrontChannelRequestObjectEncryptionRequired = &v
}

// GetRequestObjectEncryptionAlgMatchRequired returns the RequestObjectEncryptionAlgMatchRequired field value if set, zero value otherwise.
func (o *Client) GetRequestObjectEncryptionAlgMatchRequired() bool {
	if o == nil || IsNil(o.RequestObjectEncryptionAlgMatchRequired) {
		var ret bool
		return ret
	}
	return *o.RequestObjectEncryptionAlgMatchRequired
}

// GetRequestObjectEncryptionAlgMatchRequiredOk returns a tuple with the RequestObjectEncryptionAlgMatchRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetRequestObjectEncryptionAlgMatchRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestObjectEncryptionAlgMatchRequired) {
		return nil, false
	}
	return o.RequestObjectEncryptionAlgMatchRequired, true
}

// HasRequestObjectEncryptionAlgMatchRequired returns a boolean if a field has been set.
func (o *Client) HasRequestObjectEncryptionAlgMatchRequired() bool {
	if o != nil && !IsNil(o.RequestObjectEncryptionAlgMatchRequired) {
		return true
	}

	return false
}

// SetRequestObjectEncryptionAlgMatchRequired gets a reference to the given bool and assigns it to the RequestObjectEncryptionAlgMatchRequired field.
func (o *Client) SetRequestObjectEncryptionAlgMatchRequired(v bool) {
	o.RequestObjectEncryptionAlgMatchRequired = &v
}

// GetRequestObjectEncryptionEncMatchRequired returns the RequestObjectEncryptionEncMatchRequired field value if set, zero value otherwise.
func (o *Client) GetRequestObjectEncryptionEncMatchRequired() bool {
	if o == nil || IsNil(o.RequestObjectEncryptionEncMatchRequired) {
		var ret bool
		return ret
	}
	return *o.RequestObjectEncryptionEncMatchRequired
}

// GetRequestObjectEncryptionEncMatchRequiredOk returns a tuple with the RequestObjectEncryptionEncMatchRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetRequestObjectEncryptionEncMatchRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestObjectEncryptionEncMatchRequired) {
		return nil, false
	}
	return o.RequestObjectEncryptionEncMatchRequired, true
}

// HasRequestObjectEncryptionEncMatchRequired returns a boolean if a field has been set.
func (o *Client) HasRequestObjectEncryptionEncMatchRequired() bool {
	if o != nil && !IsNil(o.RequestObjectEncryptionEncMatchRequired) {
		return true
	}

	return false
}

// SetRequestObjectEncryptionEncMatchRequired gets a reference to the given bool and assigns it to the RequestObjectEncryptionEncMatchRequired field.
func (o *Client) SetRequestObjectEncryptionEncMatchRequired(v bool) {
	o.RequestObjectEncryptionEncMatchRequired = &v
}

// GetDigestAlgorithm returns the DigestAlgorithm field value if set, zero value otherwise.
func (o *Client) GetDigestAlgorithm() string {
	if o == nil || IsNil(o.DigestAlgorithm) {
		var ret string
		return ret
	}
	return *o.DigestAlgorithm
}

// GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetDigestAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.DigestAlgorithm) {
		return nil, false
	}
	return o.DigestAlgorithm, true
}

// HasDigestAlgorithm returns a boolean if a field has been set.
func (o *Client) HasDigestAlgorithm() bool {
	if o != nil && !IsNil(o.DigestAlgorithm) {
		return true
	}

	return false
}

// SetDigestAlgorithm gets a reference to the given string and assigns it to the DigestAlgorithm field.
func (o *Client) SetDigestAlgorithm(v string) {
	o.DigestAlgorithm = &v
}

// GetSingleAccessTokenPerSubject returns the SingleAccessTokenPerSubject field value if set, zero value otherwise.
func (o *Client) GetSingleAccessTokenPerSubject() bool {
	if o == nil || IsNil(o.SingleAccessTokenPerSubject) {
		var ret bool
		return ret
	}
	return *o.SingleAccessTokenPerSubject
}

// GetSingleAccessTokenPerSubjectOk returns a tuple with the SingleAccessTokenPerSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetSingleAccessTokenPerSubjectOk() (*bool, bool) {
	if o == nil || IsNil(o.SingleAccessTokenPerSubject) {
		return nil, false
	}
	return o.SingleAccessTokenPerSubject, true
}

// HasSingleAccessTokenPerSubject returns a boolean if a field has been set.
func (o *Client) HasSingleAccessTokenPerSubject() bool {
	if o != nil && !IsNil(o.SingleAccessTokenPerSubject) {
		return true
	}

	return false
}

// SetSingleAccessTokenPerSubject gets a reference to the given bool and assigns it to the SingleAccessTokenPerSubject field.
func (o *Client) SetSingleAccessTokenPerSubject(v bool) {
	o.SingleAccessTokenPerSubject = &v
}

// GetPkceRequired returns the PkceRequired field value if set, zero value otherwise.
func (o *Client) GetPkceRequired() bool {
	if o == nil || IsNil(o.PkceRequired) {
		var ret bool
		return ret
	}
	return *o.PkceRequired
}

// GetPkceRequiredOk returns a tuple with the PkceRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetPkceRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.PkceRequired) {
		return nil, false
	}
	return o.PkceRequired, true
}

// HasPkceRequired returns a boolean if a field has been set.
func (o *Client) HasPkceRequired() bool {
	if o != nil && !IsNil(o.PkceRequired) {
		return true
	}

	return false
}

// SetPkceRequired gets a reference to the given bool and assigns it to the PkceRequired field.
func (o *Client) SetPkceRequired(v bool) {
	o.PkceRequired = &v
}

// GetPkceS256Required returns the PkceS256Required field value if set, zero value otherwise.
func (o *Client) GetPkceS256Required() bool {
	if o == nil || IsNil(o.PkceS256Required) {
		var ret bool
		return ret
	}
	return *o.PkceS256Required
}

// GetPkceS256RequiredOk returns a tuple with the PkceS256Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetPkceS256RequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.PkceS256Required) {
		return nil, false
	}
	return o.PkceS256Required, true
}

// HasPkceS256Required returns a boolean if a field has been set.
func (o *Client) HasPkceS256Required() bool {
	if o != nil && !IsNil(o.PkceS256Required) {
		return true
	}

	return false
}

// SetPkceS256Required gets a reference to the given bool and assigns it to the PkceS256Required field.
func (o *Client) SetPkceS256Required(v bool) {
	o.PkceS256Required = &v
}

// GetDpopRequired returns the DpopRequired field value if set, zero value otherwise.
func (o *Client) GetDpopRequired() bool {
	if o == nil || IsNil(o.DpopRequired) {
		var ret bool
		return ret
	}
	return *o.DpopRequired
}

// GetDpopRequiredOk returns a tuple with the DpopRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetDpopRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.DpopRequired) {
		return nil, false
	}
	return o.DpopRequired, true
}

// HasDpopRequired returns a boolean if a field has been set.
func (o *Client) HasDpopRequired() bool {
	if o != nil && !IsNil(o.DpopRequired) {
		return true
	}

	return false
}

// SetDpopRequired gets a reference to the given bool and assigns it to the DpopRequired field.
func (o *Client) SetDpopRequired(v bool) {
	o.DpopRequired = &v
}

// GetAutomaticallyRegistered returns the AutomaticallyRegistered field value if set, zero value otherwise.
func (o *Client) GetAutomaticallyRegistered() bool {
	if o == nil || IsNil(o.AutomaticallyRegistered) {
		var ret bool
		return ret
	}
	return *o.AutomaticallyRegistered
}

// GetAutomaticallyRegisteredOk returns a tuple with the AutomaticallyRegistered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetAutomaticallyRegisteredOk() (*bool, bool) {
	if o == nil || IsNil(o.AutomaticallyRegistered) {
		return nil, false
	}
	return o.AutomaticallyRegistered, true
}

// HasAutomaticallyRegistered returns a boolean if a field has been set.
func (o *Client) HasAutomaticallyRegistered() bool {
	if o != nil && !IsNil(o.AutomaticallyRegistered) {
		return true
	}

	return false
}

// SetAutomaticallyRegistered gets a reference to the given bool and assigns it to the AutomaticallyRegistered field.
func (o *Client) SetAutomaticallyRegistered(v bool) {
	o.AutomaticallyRegistered = &v
}

// GetExplicitlyRegistered returns the ExplicitlyRegistered field value if set, zero value otherwise.
func (o *Client) GetExplicitlyRegistered() bool {
	if o == nil || IsNil(o.ExplicitlyRegistered) {
		var ret bool
		return ret
	}
	return *o.ExplicitlyRegistered
}

// GetExplicitlyRegisteredOk returns a tuple with the ExplicitlyRegistered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetExplicitlyRegisteredOk() (*bool, bool) {
	if o == nil || IsNil(o.ExplicitlyRegistered) {
		return nil, false
	}
	return o.ExplicitlyRegistered, true
}

// HasExplicitlyRegistered returns a boolean if a field has been set.
func (o *Client) HasExplicitlyRegistered() bool {
	if o != nil && !IsNil(o.ExplicitlyRegistered) {
		return true
	}

	return false
}

// SetExplicitlyRegistered gets a reference to the given bool and assigns it to the ExplicitlyRegistered field.
func (o *Client) SetExplicitlyRegistered(v bool) {
	o.ExplicitlyRegistered = &v
}

// GetRsRequestSigned returns the RsRequestSigned field value if set, zero value otherwise.
func (o *Client) GetRsRequestSigned() bool {
	if o == nil || IsNil(o.RsRequestSigned) {
		var ret bool
		return ret
	}
	return *o.RsRequestSigned
}

// GetRsRequestSignedOk returns a tuple with the RsRequestSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetRsRequestSignedOk() (*bool, bool) {
	if o == nil || IsNil(o.RsRequestSigned) {
		return nil, false
	}
	return o.RsRequestSigned, true
}

// HasRsRequestSigned returns a boolean if a field has been set.
func (o *Client) HasRsRequestSigned() bool {
	if o != nil && !IsNil(o.RsRequestSigned) {
		return true
	}

	return false
}

// SetRsRequestSigned gets a reference to the given bool and assigns it to the RsRequestSigned field.
func (o *Client) SetRsRequestSigned(v bool) {
	o.RsRequestSigned = &v
}

// GetRsSignedRequestKeyId returns the RsSignedRequestKeyId field value if set, zero value otherwise.
func (o *Client) GetRsSignedRequestKeyId() string {
	if o == nil || IsNil(o.RsSignedRequestKeyId) {
		var ret string
		return ret
	}
	return *o.RsSignedRequestKeyId
}

// GetRsSignedRequestKeyIdOk returns a tuple with the RsSignedRequestKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetRsSignedRequestKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.RsSignedRequestKeyId) {
		return nil, false
	}
	return o.RsSignedRequestKeyId, true
}

// HasRsSignedRequestKeyId returns a boolean if a field has been set.
func (o *Client) HasRsSignedRequestKeyId() bool {
	if o != nil && !IsNil(o.RsSignedRequestKeyId) {
		return true
	}

	return false
}

// SetRsSignedRequestKeyId gets a reference to the given string and assigns it to the RsSignedRequestKeyId field.
func (o *Client) SetRsSignedRequestKeyId(v string) {
	o.RsSignedRequestKeyId = &v
}

// GetClientRegistrationTypes returns the ClientRegistrationTypes field value if set, zero value otherwise.
func (o *Client) GetClientRegistrationTypes() []ClientRegistrationType {
	if o == nil || IsNil(o.ClientRegistrationTypes) {
		var ret []ClientRegistrationType
		return ret
	}
	return o.ClientRegistrationTypes
}

// GetClientRegistrationTypesOk returns a tuple with the ClientRegistrationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetClientRegistrationTypesOk() ([]ClientRegistrationType, bool) {
	if o == nil || IsNil(o.ClientRegistrationTypes) {
		return nil, false
	}
	return o.ClientRegistrationTypes, true
}

// HasClientRegistrationTypes returns a boolean if a field has been set.
func (o *Client) HasClientRegistrationTypes() bool {
	if o != nil && !IsNil(o.ClientRegistrationTypes) {
		return true
	}

	return false
}

// SetClientRegistrationTypes gets a reference to the given []ClientRegistrationType and assigns it to the ClientRegistrationTypes field.
func (o *Client) SetClientRegistrationTypes(v []ClientRegistrationType) {
	o.ClientRegistrationTypes = v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *Client) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationName) {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *Client) HasOrganizationName() bool {
	if o != nil && !IsNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *Client) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetSignedJwksUri returns the SignedJwksUri field value if set, zero value otherwise.
func (o *Client) GetSignedJwksUri() string {
	if o == nil || IsNil(o.SignedJwksUri) {
		var ret string
		return ret
	}
	return *o.SignedJwksUri
}

// GetSignedJwksUriOk returns a tuple with the SignedJwksUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetSignedJwksUriOk() (*string, bool) {
	if o == nil || IsNil(o.SignedJwksUri) {
		return nil, false
	}
	return o.SignedJwksUri, true
}

// HasSignedJwksUri returns a boolean if a field has been set.
func (o *Client) HasSignedJwksUri() bool {
	if o != nil && !IsNil(o.SignedJwksUri) {
		return true
	}

	return false
}

// SetSignedJwksUri gets a reference to the given string and assigns it to the SignedJwksUri field.
func (o *Client) SetSignedJwksUri(v string) {
	o.SignedJwksUri = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *Client) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *Client) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *Client) SetEntityId(v string) {
	o.EntityId = &v
}

// GetTrustAnchorId returns the TrustAnchorId field value if set, zero value otherwise.
func (o *Client) GetTrustAnchorId() string {
	if o == nil || IsNil(o.TrustAnchorId) {
		var ret string
		return ret
	}
	return *o.TrustAnchorId
}

// GetTrustAnchorIdOk returns a tuple with the TrustAnchorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetTrustAnchorIdOk() (*string, bool) {
	if o == nil || IsNil(o.TrustAnchorId) {
		return nil, false
	}
	return o.TrustAnchorId, true
}

// HasTrustAnchorId returns a boolean if a field has been set.
func (o *Client) HasTrustAnchorId() bool {
	if o != nil && !IsNil(o.TrustAnchorId) {
		return true
	}

	return false
}

// SetTrustAnchorId gets a reference to the given string and assigns it to the TrustAnchorId field.
func (o *Client) SetTrustAnchorId(v string) {
	o.TrustAnchorId = &v
}

// GetTrustChain returns the TrustChain field value if set, zero value otherwise.
func (o *Client) GetTrustChain() []string {
	if o == nil || IsNil(o.TrustChain) {
		var ret []string
		return ret
	}
	return o.TrustChain
}

// GetTrustChainOk returns a tuple with the TrustChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetTrustChainOk() ([]string, bool) {
	if o == nil || IsNil(o.TrustChain) {
		return nil, false
	}
	return o.TrustChain, true
}

// HasTrustChain returns a boolean if a field has been set.
func (o *Client) HasTrustChain() bool {
	if o != nil && !IsNil(o.TrustChain) {
		return true
	}

	return false
}

// SetTrustChain gets a reference to the given []string and assigns it to the TrustChain field.
func (o *Client) SetTrustChain(v []string) {
	o.TrustChain = v
}

// GetTrustChainExpiresAt returns the TrustChainExpiresAt field value if set, zero value otherwise.
func (o *Client) GetTrustChainExpiresAt() int64 {
	if o == nil || IsNil(o.TrustChainExpiresAt) {
		var ret int64
		return ret
	}
	return *o.TrustChainExpiresAt
}

// GetTrustChainExpiresAtOk returns a tuple with the TrustChainExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetTrustChainExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.TrustChainExpiresAt) {
		return nil, false
	}
	return o.TrustChainExpiresAt, true
}

// HasTrustChainExpiresAt returns a boolean if a field has been set.
func (o *Client) HasTrustChainExpiresAt() bool {
	if o != nil && !IsNil(o.TrustChainExpiresAt) {
		return true
	}

	return false
}

// SetTrustChainExpiresAt gets a reference to the given int64 and assigns it to the TrustChainExpiresAt field.
func (o *Client) SetTrustChainExpiresAt(v int64) {
	o.TrustChainExpiresAt = &v
}

// GetTrustChainUpdatedAt returns the TrustChainUpdatedAt field value if set, zero value otherwise.
func (o *Client) GetTrustChainUpdatedAt() int64 {
	if o == nil || IsNil(o.TrustChainUpdatedAt) {
		var ret int64
		return ret
	}
	return *o.TrustChainUpdatedAt
}

// GetTrustChainUpdatedAtOk returns a tuple with the TrustChainUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetTrustChainUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.TrustChainUpdatedAt) {
		return nil, false
	}
	return o.TrustChainUpdatedAt, true
}

// HasTrustChainUpdatedAt returns a boolean if a field has been set.
func (o *Client) HasTrustChainUpdatedAt() bool {
	if o != nil && !IsNil(o.TrustChainUpdatedAt) {
		return true
	}

	return false
}

// SetTrustChainUpdatedAt gets a reference to the given int64 and assigns it to the TrustChainUpdatedAt field.
func (o *Client) SetTrustChainUpdatedAt(v int64) {
	o.TrustChainUpdatedAt = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *Client) GetLocked() bool {
	if o == nil || IsNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *Client) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *Client) SetLocked(v bool) {
	o.Locked = &v
}

// GetCredentialOfferEndpoint returns the CredentialOfferEndpoint field value if set, zero value otherwise.
func (o *Client) GetCredentialOfferEndpoint() string {
	if o == nil || IsNil(o.CredentialOfferEndpoint) {
		var ret string
		return ret
	}
	return *o.CredentialOfferEndpoint
}

// GetCredentialOfferEndpointOk returns a tuple with the CredentialOfferEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetCredentialOfferEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialOfferEndpoint) {
		return nil, false
	}
	return o.CredentialOfferEndpoint, true
}

// HasCredentialOfferEndpoint returns a boolean if a field has been set.
func (o *Client) HasCredentialOfferEndpoint() bool {
	if o != nil && !IsNil(o.CredentialOfferEndpoint) {
		return true
	}

	return false
}

// SetCredentialOfferEndpoint gets a reference to the given string and assigns it to the CredentialOfferEndpoint field.
func (o *Client) SetCredentialOfferEndpoint(v string) {
	o.CredentialOfferEndpoint = &v
}

// GetFapiModes returns the FapiModes field value if set, zero value otherwise.
func (o *Client) GetFapiModes() []FapiMode {
	if o == nil || IsNil(o.FapiModes) {
		var ret []FapiMode
		return ret
	}
	return o.FapiModes
}

// GetFapiModesOk returns a tuple with the FapiModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetFapiModesOk() ([]FapiMode, bool) {
	if o == nil || IsNil(o.FapiModes) {
		return nil, false
	}
	return o.FapiModes, true
}

// HasFapiModes returns a boolean if a field has been set.
func (o *Client) HasFapiModes() bool {
	if o != nil && !IsNil(o.FapiModes) {
		return true
	}

	return false
}

// SetFapiModes gets a reference to the given []FapiMode and assigns it to the FapiModes field.
func (o *Client) SetFapiModes(v []FapiMode) {
	o.FapiModes = v
}

// GetResponseModes returns the ResponseModes field value if set, zero value otherwise.
func (o *Client) GetResponseModes() []string {
	if o == nil || IsNil(o.ResponseModes) {
		var ret []string
		return ret
	}
	return o.ResponseModes
}

// GetResponseModesOk returns a tuple with the ResponseModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetResponseModesOk() ([]string, bool) {
	if o == nil || IsNil(o.ResponseModes) {
		return nil, false
	}
	return o.ResponseModes, true
}

// HasResponseModes returns a boolean if a field has been set.
func (o *Client) HasResponseModes() bool {
	if o != nil && !IsNil(o.ResponseModes) {
		return true
	}

	return false
}

// SetResponseModes gets a reference to the given []string and assigns it to the ResponseModes field.
func (o *Client) SetResponseModes(v []string) {
	o.ResponseModes = v
}

// GetCredentialResponseEncryptionRequired returns the CredentialResponseEncryptionRequired field value if set, zero value otherwise.
func (o *Client) GetCredentialResponseEncryptionRequired() bool {
	if o == nil || IsNil(o.CredentialResponseEncryptionRequired) {
		var ret bool
		return ret
	}
	return *o.CredentialResponseEncryptionRequired
}

// GetCredentialResponseEncryptionRequiredOk returns a tuple with the CredentialResponseEncryptionRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetCredentialResponseEncryptionRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.CredentialResponseEncryptionRequired) {
		return nil, false
	}
	return o.CredentialResponseEncryptionRequired, true
}

// HasCredentialResponseEncryptionRequired returns a boolean if a field has been set.
func (o *Client) HasCredentialResponseEncryptionRequired() bool {
	if o != nil && !IsNil(o.CredentialResponseEncryptionRequired) {
		return true
	}

	return false
}

// SetCredentialResponseEncryptionRequired gets a reference to the given bool and assigns it to the CredentialResponseEncryptionRequired field.
func (o *Client) SetCredentialResponseEncryptionRequired(v bool) {
	o.CredentialResponseEncryptionRequired = &v
}

func (o Client) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Client) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.ServiceNumber) {
		toSerialize["serviceNumber"] = o.ServiceNumber
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
	if !IsNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
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
	if o.ApplicationType.IsSet() {
		toSerialize["applicationType"] = o.ApplicationType.Get()
	}
	if !IsNil(o.LogoUri) {
		toSerialize["logoUri"] = o.LogoUri
	}
	if !IsNil(o.LogoUris) {
		toSerialize["logoUris"] = o.LogoUris
	}
	if !IsNil(o.Contacts) {
		toSerialize["contacts"] = o.Contacts
	}
	if !IsNil(o.TlsClientCertificateBoundAccessTokens) {
		toSerialize["tlsClientCertificateBoundAccessTokens"] = o.TlsClientCertificateBoundAccessTokens
	}
	if !IsNil(o.DynamicallyRegistered) {
		toSerialize["dynamicallyRegistered"] = o.DynamicallyRegistered
	}
	if !IsNil(o.SoftwareId) {
		toSerialize["softwareId"] = o.SoftwareId
	}
	if !IsNil(o.SoftwareVersion) {
		toSerialize["softwareVersion"] = o.SoftwareVersion
	}
	if !IsNil(o.RegistrationAccessTokenHash) {
		toSerialize["registrationAccessTokenHash"] = o.RegistrationAccessTokenHash
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if !IsNil(o.GrantTypes) {
		toSerialize["grantTypes"] = o.GrantTypes
	}
	if !IsNil(o.ResponseTypes) {
		toSerialize["responseTypes"] = o.ResponseTypes
	}
	if !IsNil(o.RedirectUris) {
		toSerialize["redirectUris"] = o.RedirectUris
	}
	if !IsNil(o.AuthorizationSignAlg) {
		toSerialize["authorizationSignAlg"] = o.AuthorizationSignAlg
	}
	if !IsNil(o.AuthorizationEncryptionAlg) {
		toSerialize["authorizationEncryptionAlg"] = o.AuthorizationEncryptionAlg
	}
	if !IsNil(o.AuthorizationEncryptionEnc) {
		toSerialize["authorizationEncryptionEnc"] = o.AuthorizationEncryptionEnc
	}
	if !IsNil(o.TokenAuthMethod) {
		toSerialize["tokenAuthMethod"] = o.TokenAuthMethod
	}
	if !IsNil(o.TokenAuthSignAlg) {
		toSerialize["tokenAuthSignAlg"] = o.TokenAuthSignAlg
	}
	if !IsNil(o.SelfSignedCertificateKeyId) {
		toSerialize["selfSignedCertificateKeyId"] = o.SelfSignedCertificateKeyId
	}
	if !IsNil(o.TlsClientAuthSubjectDn) {
		toSerialize["tlsClientAuthSubjectDn"] = o.TlsClientAuthSubjectDn
	}
	if !IsNil(o.TlsClientAuthSanDns) {
		toSerialize["tlsClientAuthSanDns"] = o.TlsClientAuthSanDns
	}
	if !IsNil(o.TlsClientAuthSanUri) {
		toSerialize["tlsClientAuthSanUri"] = o.TlsClientAuthSanUri
	}
	if !IsNil(o.TlsClientAuthSanIp) {
		toSerialize["tlsClientAuthSanIp"] = o.TlsClientAuthSanIp
	}
	if !IsNil(o.TlsClientAuthSanEmail) {
		toSerialize["tlsClientAuthSanEmail"] = o.TlsClientAuthSanEmail
	}
	if !IsNil(o.ParRequired) {
		toSerialize["parRequired"] = o.ParRequired
	}
	if !IsNil(o.RequestObjectRequired) {
		toSerialize["requestObjectRequired"] = o.RequestObjectRequired
	}
	if !IsNil(o.RequestSignAlg) {
		toSerialize["requestSignAlg"] = o.RequestSignAlg
	}
	if !IsNil(o.RequestEncryptionAlg) {
		toSerialize["requestEncryptionAlg"] = o.RequestEncryptionAlg
	}
	if !IsNil(o.RequestEncryptionEnc) {
		toSerialize["requestEncryptionEnc"] = o.RequestEncryptionEnc
	}
	if !IsNil(o.RequestUris) {
		toSerialize["requestUris"] = o.RequestUris
	}
	if !IsNil(o.DefaultMaxAge) {
		toSerialize["defaultMaxAge"] = o.DefaultMaxAge
	}
	if !IsNil(o.DefaultAcrs) {
		toSerialize["defaultAcrs"] = o.DefaultAcrs
	}
	if !IsNil(o.IdTokenSignAlg) {
		toSerialize["idTokenSignAlg"] = o.IdTokenSignAlg
	}
	if !IsNil(o.IdTokenEncryptionAlg) {
		toSerialize["idTokenEncryptionAlg"] = o.IdTokenEncryptionAlg
	}
	if !IsNil(o.IdTokenEncryptionEnc) {
		toSerialize["idTokenEncryptionEnc"] = o.IdTokenEncryptionEnc
	}
	if !IsNil(o.AuthTimeRequired) {
		toSerialize["authTimeRequired"] = o.AuthTimeRequired
	}
	if !IsNil(o.SubjectType) {
		toSerialize["subjectType"] = o.SubjectType
	}
	if !IsNil(o.SectorIdentifierUri) {
		toSerialize["sectorIdentifierUri"] = o.SectorIdentifierUri
	}
	if !IsNil(o.DerivedSectorIdentifier) {
		toSerialize["derivedSectorIdentifier"] = o.DerivedSectorIdentifier
	}
	if !IsNil(o.JwksUri) {
		toSerialize["jwksUri"] = o.JwksUri
	}
	if !IsNil(o.Jwks) {
		toSerialize["jwks"] = o.Jwks
	}
	if !IsNil(o.UserInfoSignAlg) {
		toSerialize["userInfoSignAlg"] = o.UserInfoSignAlg
	}
	if !IsNil(o.UserInfoEncryptionAlg) {
		toSerialize["userInfoEncryptionAlg"] = o.UserInfoEncryptionAlg
	}
	if !IsNil(o.UserInfoEncryptionEnc) {
		toSerialize["userInfoEncryptionEnc"] = o.UserInfoEncryptionEnc
	}
	if !IsNil(o.LoginUri) {
		toSerialize["loginUri"] = o.LoginUri
	}
	if !IsNil(o.TosUri) {
		toSerialize["tosUri"] = o.TosUri
	}
	if !IsNil(o.TosUris) {
		toSerialize["tosUris"] = o.TosUris
	}
	if !IsNil(o.PolicyUri) {
		toSerialize["policyUri"] = o.PolicyUri
	}
	if !IsNil(o.PolicyUris) {
		toSerialize["policyUris"] = o.PolicyUris
	}
	if !IsNil(o.ClientUri) {
		toSerialize["clientUri"] = o.ClientUri
	}
	if !IsNil(o.ClientUris) {
		toSerialize["clientUris"] = o.ClientUris
	}
	if !IsNil(o.BcDeliveryMode) {
		toSerialize["bcDeliveryMode"] = o.BcDeliveryMode
	}
	if !IsNil(o.BcNotificationEndpoint) {
		toSerialize["bcNotificationEndpoint"] = o.BcNotificationEndpoint
	}
	if !IsNil(o.BcRequestSignAlg) {
		toSerialize["bcRequestSignAlg"] = o.BcRequestSignAlg
	}
	if !IsNil(o.BcUserCodeRequired) {
		toSerialize["bcUserCodeRequired"] = o.BcUserCodeRequired
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Extension) {
		toSerialize["extension"] = o.Extension
	}
	if !IsNil(o.AuthorizationDetailsTypes) {
		toSerialize["authorizationDetailsTypes"] = o.AuthorizationDetailsTypes
	}
	if !IsNil(o.CustomMetadata) {
		toSerialize["customMetadata"] = o.CustomMetadata
	}
	if !IsNil(o.FrontChannelRequestObjectEncryptionRequired) {
		toSerialize["frontChannelRequestObjectEncryptionRequired"] = o.FrontChannelRequestObjectEncryptionRequired
	}
	if !IsNil(o.RequestObjectEncryptionAlgMatchRequired) {
		toSerialize["requestObjectEncryptionAlgMatchRequired"] = o.RequestObjectEncryptionAlgMatchRequired
	}
	if !IsNil(o.RequestObjectEncryptionEncMatchRequired) {
		toSerialize["requestObjectEncryptionEncMatchRequired"] = o.RequestObjectEncryptionEncMatchRequired
	}
	if !IsNil(o.DigestAlgorithm) {
		toSerialize["digestAlgorithm"] = o.DigestAlgorithm
	}
	if !IsNil(o.SingleAccessTokenPerSubject) {
		toSerialize["singleAccessTokenPerSubject"] = o.SingleAccessTokenPerSubject
	}
	if !IsNil(o.PkceRequired) {
		toSerialize["pkceRequired"] = o.PkceRequired
	}
	if !IsNil(o.PkceS256Required) {
		toSerialize["pkceS256Required"] = o.PkceS256Required
	}
	if !IsNil(o.DpopRequired) {
		toSerialize["dpopRequired"] = o.DpopRequired
	}
	if !IsNil(o.AutomaticallyRegistered) {
		toSerialize["automaticallyRegistered"] = o.AutomaticallyRegistered
	}
	if !IsNil(o.ExplicitlyRegistered) {
		toSerialize["explicitlyRegistered"] = o.ExplicitlyRegistered
	}
	if !IsNil(o.RsRequestSigned) {
		toSerialize["rsRequestSigned"] = o.RsRequestSigned
	}
	if !IsNil(o.RsSignedRequestKeyId) {
		toSerialize["rsSignedRequestKeyId"] = o.RsSignedRequestKeyId
	}
	if !IsNil(o.ClientRegistrationTypes) {
		toSerialize["clientRegistrationTypes"] = o.ClientRegistrationTypes
	}
	if !IsNil(o.OrganizationName) {
		toSerialize["organizationName"] = o.OrganizationName
	}
	if !IsNil(o.SignedJwksUri) {
		toSerialize["signedJwksUri"] = o.SignedJwksUri
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.TrustAnchorId) {
		toSerialize["trustAnchorId"] = o.TrustAnchorId
	}
	if !IsNil(o.TrustChain) {
		toSerialize["trustChain"] = o.TrustChain
	}
	if !IsNil(o.TrustChainExpiresAt) {
		toSerialize["trustChainExpiresAt"] = o.TrustChainExpiresAt
	}
	if !IsNil(o.TrustChainUpdatedAt) {
		toSerialize["trustChainUpdatedAt"] = o.TrustChainUpdatedAt
	}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !IsNil(o.CredentialOfferEndpoint) {
		toSerialize["credentialOfferEndpoint"] = o.CredentialOfferEndpoint
	}
	if !IsNil(o.FapiModes) {
		toSerialize["fapiModes"] = o.FapiModes
	}
	if !IsNil(o.ResponseModes) {
		toSerialize["responseModes"] = o.ResponseModes
	}
	if !IsNil(o.CredentialResponseEncryptionRequired) {
		toSerialize["credentialResponseEncryptionRequired"] = o.CredentialResponseEncryptionRequired
	}
	return toSerialize, nil
}

type NullableClient struct {
	value *Client
	isSet bool
}

func (v NullableClient) Get() *Client {
	return v.value
}

func (v *NullableClient) Set(val *Client) {
	v.value = val
	v.isSet = true
}

func (v NullableClient) IsSet() bool {
	return v.isSet
}

func (v *NullableClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClient(val *Client) *NullableClient {
	return &NullableClient{value: val, isSet: true}
}

func (v NullableClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
