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

// checks if the Service type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Service{}

// Service struct for Service
type Service struct {
	// The sequential number of the service. The value of this property is assigned by Authlete.
	Number *int32 `json:"number,omitempty"`
	// The name of this service.
	ServiceName *string `json:"serviceName,omitempty"`
	// The issuer identifier of the service.  A URL that starts with  https:// and has no query or fragment component.  The value of this property is used as `iss` claim in an [ID token](https://openid.net/specs/openid-connect-core-1_0.html#IDToken) and `issuer` property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	Issuer *string `json:"issuer,omitempty"`
	// The description about the service.
	Description *string `json:"description,omitempty"`
	// The service ID used in Authlete API calls. The value of this property is assigned by Authlete.
	ApiKey *int64 `json:"apiKey,omitempty"`
	// Deprecated. Always `true`.
	ClientIdAliasEnabled *bool `json:"clientIdAliasEnabled,omitempty"`
	// The `metadata` of the service. The content of the returned array depends on contexts. The predefined service metadata is listed in the following table.    | Key | Description |   | --- | --- |   | `clientCount` | The number of client applications which belong to this service.  |
	Metadata []Pair `json:"metadata,omitempty"`
	// The time at which this service was created. The value is represented as milliseconds since the UNIX epoch (`1970-01-01`).
	CreatedAt *int64 `json:"createdAt,omitempty"`
	// The time at which this service was last modified. The value is represented as milliseconds since the UNIX epoch (1970-01-01).
	ModifiedAt *int64 `json:"modifiedAt,omitempty"`
	// A Web API endpoint for user authentication which is to be prepared on the service side.  The endpoint must be implemented if you do not implement the UI at the authorization endpoint but use the one provided by Authlete.  The user authentication at the authorization endpoint provided by Authlete is performed by making a `POST` request to this endpoint.
	AuthenticationCallbackEndpoint *string `json:"authenticationCallbackEndpoint,omitempty"`
	// API key for basic authentication at the authentication callback endpoint.  If the value is not empty, Authlete generates Authorization header for Basic authentication when making a request to the authentication callback endpoint.
	AuthenticationCallbackApiKey *string `json:"authenticationCallbackApiKey,omitempty"`
	// API secret for `basic` authentication at the authentication callback endpoint.
	AuthenticationCallbackApiSecret *string `json:"authenticationCallbackApiSecret,omitempty"`
	// Values of acrs (authentication context class references) that the service supports.  The value of this property is used as `acr_values_supported` property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	SupportedAcrs []string `json:"supportedAcrs,omitempty"`
	// Values of `grant_type` request parameter that the service supports.  The value of this property is used as `grant_types_supported property` in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	SupportedGrantTypes []GrantType `json:"supportedGrantTypes,omitempty"`
	// Values of `response_type` request parameter that the service supports. Valid values are listed in Response Type.  The value of this property is used as `response_types_supported` property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	SupportedResponseTypes []ResponseType `json:"supportedResponseTypes,omitempty"`
	// The supported data types that can be used as values of the type field in `authorization_details`.  This property corresponds to the `authorization_details_types_supported` metadata. See \"OAuth 2.0 Rich Authorization Requests\" (RAR) for details.
	SupportedAuthorizationDetailsTypes []string `json:"supportedAuthorizationDetailsTypes,omitempty"`
	// The profiles that this service supports.
	SupportedServiceProfiles []ServiceProfile `json:"supportedServiceProfiles,omitempty"`
	// The flag to indicate whether the `error_description` response parameter is omitted.  According to [RFC 6749](https://tools.ietf.org/html/rfc6749), an authorization server may include the `error_description` response parameter in error responses.  If `true`, Authlete does not embed the `error_description` response parameter in error responses.
	ErrorDescriptionOmitted *bool `json:"errorDescriptionOmitted,omitempty"`
	// The flag to indicate whether the `error_uri` response parameter is omitted.  According to [RFC 6749](https://tools.ietf.org/html/rfc6749), an authorization server may include the `error_uri` response parameter in error responses.  If `true`, Authlete does not embed the `error_uri` response parameter in error responses.
	ErrorUriOmitted *bool `json:"errorUriOmitted,omitempty"`
	// The authorization endpoint of the service.  A URL that starts with `https://` and has no fragment component. For example, `https://example.com/auth/authorization`.  The value of this property is used as `authorization_endpoint` property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty"`
	// The flag to indicate whether the direct authorization endpoint is enabled or not.  The path of the endpoint is `/api/auth/authorization/direct/service-api-key`.
	DirectAuthorizationEndpointEnabled *bool `json:"directAuthorizationEndpointEnabled,omitempty"`
	// UI locales that the service supports.  Each element is a language tag defined in [RFC 5646](https://tools.ietf.org/html/rfc5646). For example, `en-US` and `ja-JP`.  The value of this property is used as `ui_locales_supported` property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	SupportedUiLocales []string `json:"supportedUiLocales,omitempty"`
	// Values of `display` request parameter that service supports.  The value of this property is used as `display_values_supported` property in the Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	SupportedDisplays []Display `json:"supportedDisplays,omitempty"`
	// The flag to indicate whether the use of Proof Key for Code Exchange (PKCE) is always required for authorization requests by Authorization Code Flow.  If `true`, `code_challenge` request parameter is always required for authorization requests using Authorization Code Flow.  See [RFC 7636](https://tools.ietf.org/html/rfc7636) (Proof Key for Code Exchange by OAuth Public Clients) for details about `code_challenge` request parameter.
	PkceRequired *bool `json:"pkceRequired,omitempty"`
	// The flag to indicate whether `S256` is always required as the code challenge method whenever [PKCE (RFC 7636)](https://tools.ietf.org/html/rfc7636) is used.  If this flag is set to `true`, `code_challenge_method=S256` must be included in the authorization request whenever it includes the `code_challenge` request parameter. Neither omission of the `code_challenge_method` request parameter nor use of plain (`code_challenge_method=plain`) is allowed.
	PkceS256Required *bool `json:"pkceS256Required,omitempty"`
	// The duration of authorization response JWTs in seconds.  [Financial-grade API: JWT Secured Authorization Response Mode for OAuth 2.0 (JARM)](https://openid.net/specs/openid-financial-api-jarm.html) defines new values for the `response_mode` request parameter. They are `query.jwt`, `fragment.jwt`, `form_post.jwt` and `jwt`. If one of them is specified as the response mode, response parameters from the authorization endpoint will be packed into a JWT. This property is used to compute the value of the `exp` claim of the JWT.
	AuthorizationResponseDuration *int64 `json:"authorizationResponseDuration,omitempty"`
	// The [token endpoint](https://tools.ietf.org/html/rfc6749#section-3.2) of the service.  A URL that starts with `https://` and has not fragment component. For example, `https://example.com/auth/token`.  The value of this property is used as `token_endpoint` property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	TokenEndpoint *string `json:"tokenEndpoint,omitempty"`
	// The flag to indicate whether the direct token endpoint is enabled or not. The path of the endpoint is `/api/auth/token/direct/service-api-key`.
	DirectTokenEndpointEnabled *bool `json:"directTokenEndpointEnabled,omitempty"`
	// Client authentication methods supported by the token endpoint of the service.  The value of this property is used as `token_endpoint_auth_methods_supports` property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	SupportedTokenAuthMethods []ClientAuthMethod `json:"supportedTokenAuthMethods,omitempty"`
	// The flag to indicate token requests from public clients without the `client_id` request parameter are allowed when the client can be guessed from `authorization_code` or `refresh_token`.  This flag should not be set unless you have special reasons.
	MissingClientIdAllowed *bool `json:"missingClientIdAllowed,omitempty"`
	// The [revocation endpoint](https://tools.ietf.org/html/rfc7009) of the service.  A URL that starts with `https://`. For example, `https://example.com/auth/revocation`.
	RevocationEndpoint *string `json:"revocationEndpoint,omitempty"`
	// The flag to indicate whether the direct revocation endpoint is enabled or not. The URL of the endpoint is `/api/auth/revocation/direct/service-api-key`.
	DirectRevocationEndpointEnabled *bool `json:"directRevocationEndpointEnabled,omitempty"`
	// Client authentication methods supported at the revocation endpoint.
	SupportedRevocationAuthMethods []ClientAuthMethod `json:"supportedRevocationAuthMethods,omitempty"`
	// The URI of the introspection endpoint.
	IntrospectionEndpoint *string `json:"introspectionEndpoint,omitempty"`
	// The flag to indicate whether the direct userinfo endpoint is enabled or not. The path of the endpoint is `/api/auth/userinfo/direct/{serviceApiKey}`.
	DirectIntrospectionEndpointEnabled *bool `json:"directIntrospectionEndpointEnabled,omitempty"`
	// Client authentication methods supported at the introspection endpoint.
	SupportedIntrospectionAuthMethods []ClientAuthMethod `json:"supportedIntrospectionAuthMethods,omitempty"`
	// The URI of the pushed authorization request endpoint.  This property corresponds to the `pushed_authorization_request_endpoint` metadata defined in \"[5. Authorization Server Metadata](https://tools.ietf.org/html/draft-lodderstedt-oauth-par#section-5)\" of OAuth 2.0 Pushed Authorization Requests.
	PushedAuthReqEndpoint *string `json:"pushedAuthReqEndpoint,omitempty"`
	// The duration of pushed authorization requests in seconds.  [OAuth 2.0 Pushed Authorization Requests](https://tools.ietf.org/html/draft-lodderstedt-oauth-par) defines an endpoint (called \"pushed authorization request endpoint\") which client applications can register authorization requests into and get corresponding URIs (called \"request URIs\") from. The issued URIs represent the registered authorization requests. The client applications can use the URIs as the value of the `request_uri` request parameter in an authorization request.  The property represents the duration of registered authorization requests and is used as the value of the `expires_in` parameter in responses from the pushed authorization request endpoint.
	PushedAuthReqDuration *int64 `json:"pushedAuthReqDuration,omitempty"`
	// The flag to indicate whether this service requires that clients use the pushed authorization request endpoint.  This property corresponds to the `require_pushed_authorization_requests` server metadata defined in [OAuth 2.0 Pushed Authorization Requests](https://tools.ietf.org/html/draft-lodderstedt-oauth-par).
	ParRequired *bool `json:"parRequired,omitempty"`
	// The flag to indicate whether this service requires that authorization requests always utilize a request object by using either request or `request_uri` request parameter.  If this flag is set to `true` and the value of `traditionalRequestObjectProcessingApplied` is `false`, the value of `require_signed_request_object` server metadata of this service is reported as `true` in the discovery document. The metadata is defined in JAR (JWT Secured Authorization Request). That `require_signed_request_object` is `true` means that authorization requests which don't conform to the JAR specification are rejected.
	RequestObjectRequired *bool `json:"requestObjectRequired,omitempty"`
	// The flag to indicate whether a request object is processed based on rules defined in [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html) or JAR (JWT Secured Authorization Request).  Differences between rules in OpenID Connect Core 1.0 and ones in JAR are as follows.   - JAR requires that a request object be always -signed.   - JAR does not allow request parameters outside a request object to be referred to.   - OIDC Core 1.0 requires that response_type request parameter exist outside a request object even if the request object includes the request parameter.   - OIDC Core 1.0 requires that scope request parameter exist outside a request object if the authorization request is an   - OIDC request even if the request object includes the request parameter.  If this flag is set to `false` and the value of `requestObjectRequired` is `true`, the value of `require_signed_request_object` server metadata of this service is reported as `true` in the discovery document. The metadata is defined in JAR (JWT Secured Authorization Request). That `require_signed_request_object` is `true` means that authorization requests which don't conform to the JAR specification are rejected.
	TraditionalRequestObjectProcessingApplied *bool `json:"traditionalRequestObjectProcessingApplied,omitempty"`
	// The flag to indicate whether this service validates certificate chains during PKI-based client mutual TLS authentication.
	MutualTlsValidatePkiCertChain *bool `json:"mutualTlsValidatePkiCertChain,omitempty"`
	// The list of root certificates trusted by this service for PKI-based client mutual TLS authentication.
	TrustedRootCertificates []string `json:"trustedRootCertificates,omitempty"`
	// The MTLS endpoint aliases.  This property corresponds to the mtls_endpoint_aliases metadata defined in \"5. Metadata for Mutual TLS Endpoint Aliases\" of [OAuth 2.0 Mutual TLS Client Authentication and Certificate-Bound Access Tokens](https://datatracker.ietf.org/doc/rfc8705/).  The aliases will be embedded in the response from the discovery endpoint like the following.  ```json {   ......,   \"mtls_endpoint_aliases\": {     \"token_endpoint\":         \"https://mtls.example.com/token\",     \"revocation_endpoint\":    \"https://mtls.example.com/revo\",     \"introspection_endpoint\": \"https://mtls.example.com/introspect\"   } } ```
	MtlsEndpointAliases []NamedUri `json:"mtlsEndpointAliases,omitempty"`
	// The access token type.  This value is used as the value of `token_type` property in access token responses. If this service complies with [RFC 6750](https://tools.ietf.org/html/rfc6750), the value of this property should be `Bearer`.  See [RFC 6749 (OAuth 2.0), 7.1. Access Token Types](https://tools.ietf.org/html/rfc6749#section-7.1) for details.
	AccessTokenType *string `json:"accessTokenType,omitempty"`
	// The flag to indicate whether this service supports issuing TLS client certificate bound access tokens.
	TlsClientCertificateBoundAccessTokens *bool `json:"tlsClientCertificateBoundAccessTokens,omitempty"`
	// The duration of access tokens in seconds. This value is used as the value of `expires_in` property in access token responses. `expires_in` is defined [RFC 6749, 5.1. Successful Response](https://tools.ietf.org/html/rfc6749#section-5.1).
	AccessTokenDuration *int64 `json:"accessTokenDuration,omitempty"`
	// The flag to indicate whether the number of access tokens per subject (and per client) is at most one or can be more.  If `true`, an attempt to issue a new access token invalidates existing access tokens that are associated with the same subject and the same client.  Note that, however, attempts by [Client Credentials Flow](https://tools.ietf.org/html/rfc6749#section-4.4) do not invalidate existing access tokens because access tokens issued by Client Credentials Flow are not associated with any end-user's subject. Also note that an attempt by [Refresh Token Flow](https://tools.ietf.org/html/rfc6749#section-6) invalidates the coupled access token only and this invalidation is always performed regardless of whether the value of this setting item is `true` or `false`.
	SingleAccessTokenPerSubject *bool   `json:"singleAccessTokenPerSubject,omitempty"`
	AccessTokenSignAlg          *JwsAlg `json:"accessTokenSignAlg,omitempty"`
	// The key ID to identify a JWK used for signing access tokens.  A JWK Set can be registered as a property of a service. A JWK Set can contain 0 or more JWKs. Authlete Server has to pick up one JWK for signing from the JWK Set when it generates a JWT-based access token. Authlete Server searches the registered JWK Set for a JWK which satisfies conditions for access token signature. If the number of JWK candidates which satisfy the conditions is 1, there is no problem. On the other hand, if there exist multiple candidates, a Key ID is needed to be specified so that Authlete Server can pick up one JWK from among the JWK candidates.
	AccessTokenSignatureKeyId *string `json:"accessTokenSignatureKeyId,omitempty"`
	// The duration of refresh tokens in seconds. The related specifications have no requirements on refresh token duration, but Authlete sets expiration for refresh tokens.
	RefreshTokenDuration *int64 `json:"refreshTokenDuration,omitempty"`
	// The flag to indicate whether the remaining duration of the used refresh token is taken over to the newly issued refresh token.
	RefreshTokenDurationKept *bool `json:"refreshTokenDurationKept,omitempty"`
	// The flag which indicates whether duration of refresh tokens are reset when they are used even if the `refreshTokenKept` property of this service set to is `true` (= even if \"Refresh Token Continuous Use\" is \"Kept\").  This flag has no effect when the `refreshTokenKept` property is set to `false`. In other words, if this service issues a new refresh token on every refresh token request, the refresh token will have fresh duration (unless `refreshTokenDurationKept` is set to `true`) and this `refreshTokenDurationReset` property is not referenced.
	RefreshTokenDurationReset *bool `json:"refreshTokenDurationReset,omitempty"`
	// The flag to indicate whether a refresh token remains unchanged or gets renewed after its use.  If `true`, a refresh token used to get a new access token remains valid after its use. Otherwise, if `false`, a refresh token is invalidated after its use and a new refresh token is issued.  See [RFC 6749 6. Refreshing an Access Token](https://tools.ietf.org/html/rfc6749#section-6), as to how to get a new access token using a refresh token.
	RefreshTokenKept *bool `json:"refreshTokenKept,omitempty"`
	// Scopes supported by the service.  Authlete strongly recommends that the service register at least the following scopes.  | Name | Description | | --- | --- | | openid | A permission to get an ID token of an end-user. The `openid` scope appears in [OpenID Connect Core 1.0, 3.1.2.1. Authentication Request, scope](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest). Without this scope, Authlete does not allow `response_type` request parameter to have values other than code and token. | | profile | A permission to get information about `name`, `family_name`, `given_name`, `middle_name`, `nickname`, `preferred_username`, `profile`, `picture`, `website`, `gender`, `birthdate`, `zoneinfo`, `locale` and `updated_at` from the user info endpoint. See [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) for details. | | email | A permission to get information about `email` and `email_verified` from the user info endpoint. See [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) for details. | | address | A permission to get information about address from the user info endpoint. See [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) and [5.1.1. Address Claim](https://openid.net/specs/openid-connect-core-1_0.html#AddressClaim) for details. | | phone | A permission to get information about `phone_number` and `phone_number_verified` from the user info endpoint. See [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) for details. | | offline_access | A permission to get information from the user info endpoint even when the end-user is not present. See [OpenID Connect Core 1.0, 11. Offline Access](https://openid.net/specs/openid-connect-core-1_0.html#OfflineAccess) for details. |  The value of this property is used as `scopes_supported` property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	SupportedScopes []Scope `json:"supportedScopes,omitempty"`
	// The flag to indicate whether requests that request no scope are rejected or not.  When a request has no explicit `scope` parameter and the service's pre-defined default scope set is empty, the authorization server regards the request requests no scope. When this flag is set to `true`, requests that request no scope are rejected.  The requirement below excerpted from [RFC 6749 Section 3.3](https://tools.ietf.org/html/rfc6749#section-3.3) does not explicitly mention the case where the default scope set is empty.  > If the client omits the scope parameter when requesting authorization, the authorization server MUST either process the request using a pre-defined default value or fail the request indicating an invalid scope.  However, if you interpret *\"the default scope set exists but is empty\"* as *\"the default scope set does not exist\"* and want to strictly conform to the requirement above, this flag has to be `true`.
	ScopeRequired *bool `json:"scopeRequired,omitempty"`
	// 'The duration of [ID token](https://openid.net/specs/openid-connect-core-1_0.html#IDToken)s in seconds. This value is used to calculate the value of `exp` claim in an ID token.'
	IdTokenDuration *int64 `json:"idTokenDuration,omitempty"`
	// The allowable clock skew between the server and clients in seconds.  The clock skew is taken into consideration when time-related claims in a JWT (e.g. `exp`, `iat`, `nbf`) are verified.
	AllowableClockSkew *int32 `json:"allowableClockSkew,omitempty"`
	// Claim types supported by the service. Valid values are listed in Claim Type. Note that Authlete currently doesn't provide any API to help implementations for `AGGREGATED` and `DISTRIBUTED`.  The value of this property is used as `claim_types_supported` property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	SupportedClaimTypes []ClaimType `json:"supportedClaimTypes,omitempty"`
	// Claim locales that the service supports. Each element is a language tag defined in [RFC 5646](https://tools.ietf.org/html/rfc5646). For example, `en-US` and `ja-JP`. See [OpenID Connect Core 1.0, 5.2. Languages and Scripts](https://openid.net/specs/openid-connect-core-1_0.html#ClaimsLanguagesAndScripts) for details.  The value of this property is used as `claims_locales_supported` property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	SupportedClaimLocales []string `json:"supportedClaimLocales,omitempty"`
	// Claim names that the service supports. The standard claim names listed in [OpenID Connect Core 1.0, 5.1. Standard Claim](https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims) should be supported. The following is the list of standard claims.  - `sub` - `name` - `given_name` - `family_name` - `middle_name` - `nickname` - `preferred_username` - `profile` - `picture` - `website` - `email` - `email_verified` - `gender` - `birthdate` - `zoneinfo` - `locale` - `phone_number` - `phone_number_verified` - `address` - `updated_at`  The value of this property is used as `claims_supported` property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  The service may support its original claim names. See [OpenID Connect Core 1.0, 5.1.2. Additional Claims](https://openid.net/specs/openid-connect-core-1_0.html#AdditionalClaims).
	SupportedClaims []string `json:"supportedClaims,omitempty"`
	// The flag indicating whether claims specified by shortcut scopes (e.g. `profile`) are included in the issued ID token only when no access token is issued.  To strictly conform to the description below excerpted from [OpenID Connect Core 1.0 Section 5.4](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims), this flag has to be `true`.  > The Claims requested by the profile, email, address, and phone scope values are returned from the UserInfo Endpoint, as described in Section 5.3.2, when a response_type value is used that results in an Access Token being issued. However, when no Access Token is issued (which is the case for the response_type value id_token), the resulting Claims are returned in the ID Token.
	ClaimShortcutRestrictive *bool `json:"claimShortcutRestrictive,omitempty"`
	// The URL of the service's [JSON Web Key Set](https://tools.ietf.org/html/rfc7517) document. For example, `http://example.com/auth/jwks`.  Client applications accesses this URL (1) to get the public key of the service to validate the signature of an ID token issued by the service and (2) to get the public key of the service to encrypt an request object of the client application. See [OpenID Connect Core 1.0, 10. Signatures and Encryption](https://openid.net/specs/openid-connect-core-1_0.html#SigEnc) for details.  The value of this property is used as `jwks_uri` property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	JwksUri *string `json:"jwksUri,omitempty"`
	// 'The flag to indicate whether the direct jwks endpoint is enabled or not. The path of the endpoint is `/api/service/jwks/get/direct/service-api-key`. '
	DirectJwksEndpointEnabled *bool `json:"directJwksEndpointEnabled,omitempty"`
	// The content of the service's [JSON Web Key Set](https://tools.ietf.org/html/rfc7517) document.  If this property is not `null` in a `/service/create` request or a `/service/update` request, Authlete hosts the content in the database. This property must not be `null` and must contain pairs of public/private keys if the service wants to support asymmetric signatures for ID tokens and asymmetric encryption for request objects. See [OpenID Connect Core 1.0, 10. Signatures and Encryption](https://openid.net/specs/openid-connect-core-1_0.html#SigEnc) for details.
	Jwks *string `json:"jwks,omitempty"`
	// The key ID to identify a JWK used for ID token signature using an asymmetric key.  A JWK Set can be registered as a property of a Service. A JWK Set can contain 0 or more JWKs (See [RFC 7517](https://tools.ietf.org/html/rfc7517) for details about JWK). Authlete Server has to pick up one JWK for signature from the JWK Set when it generates an ID token and signature using an asymmetric key is required. Authlete Server searches the registered JWK Set for a JWK which satisfies conditions for ID token signature. If the number of JWK candidates which satisfy the conditions is 1, there is no problem. On the other hand, if there exist multiple candidates, a [Key ID](https://tools.ietf.org/html/rfc7517#section-4.5) is needed to be specified so that Authlete Server can pick up one JWK from among the JWK candidates.  This `idTokenSignatureKeyId` property exists for the purpose described above. For key rotation (OpenID Connect Core 1.0, [10.1.1. Rotation of Asymmetric Signing Keys](http://openid.net/specs/openid-connect-core-1_0.html#RotateSigKeys)), this mechanism is needed.
	IdTokenSignatureKeyId *string `json:"idTokenSignatureKeyId,omitempty"`
	// The key ID to identify a JWK used for user info signature using an asymmetric key.  A JWK Set can be registered as a property of a Service. A JWK Set can contain 0 or more JWKs (See [RFC 7517](https://tools.ietf.org/html/rfc7517) for details about JWK). Authlete Server has to pick up one JWK for signature from the JWK Set when it is required to sign user info (which is returned from [userinfo endpoint](http://openid.net/specs/openid-connect-core-1_0.html#UserInfo)) using an asymmetric key. Authlete Server searches the registered JWK Set for a JWK which satisfies conditions for user info signature. If the number of JWK candidates which satisfy the conditions is 1, there is no problem. On the other hand, if there exist multiple candidates, a [Key ID](https://tools.ietf.org/html/rfc7517#section-4.5) is needed to be specified so that Authlete Server can pick up one JWK from among the JWK candidates.  This `userInfoSignatureKeyId` property exists for the purpose described above. For key rotation (OpenID Connect Core 1.0, [10.1.1. Rotation of Asymmetric Signing Keys](http://openid.net/specs/openid-connect-core-1_0.html#RotateSigKeys)), this mechanism is needed.
	UserInfoSignatureKeyId *string `json:"userInfoSignatureKeyId,omitempty"`
	// The key ID to identify a JWK used for signing authorization responses using an asymmetric key.  [Financial-grade API: JWT Secured Authorization Response Mode for OAuth 2.0 (JARM)](https://openid.net/specs/openid-financial-api-jarm.html) defines new values for the `response_mode` request parameter. They are `query.jwt`, `fragment.jwt`, `form_post.jwt` and `jwt`. If one of them is specified as the response mode, response parameters from the authorization endpoint will be packed into a JWT. This property is used to compute the value of the `exp` claim of the JWT.  Authlete Server searches the JWK Set for a JWK which satisfies conditions for authorization response signature. If the number of JWK candidates which satisfy the conditions is 1, there is no problem. On the other hand, if there exist multiple candidates, a Key ID is needed to be specified so that Authlete Server can pick up one JWK from among the JWK candidates. This property exists to specify the key ID.
	AuthorizationSignatureKeyId *string `json:"authorizationSignatureKeyId,omitempty"`
	// The [user info endpoint](http://openid.net/specs/openid-connect-core-1_0.html#UserInfo) of the service. A URL that starts with `https://`. For example, `https://example.com/auth/userinfo`.  The value of this property is used as `userinfo_endpoint` property in the [OpenID Provider Metadata](http://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	UserInfoEndpoint *string `json:"userInfoEndpoint,omitempty"`
	// The flag to indicate whether the direct userinfo endpoint is enabled or not. The path of the endpoint is `/api/auth/userinfo/direct/service-api-key`.
	DirectUserInfoEndpointEnabled *bool `json:"directUserInfoEndpointEnabled,omitempty"`
	// The boolean flag which indicates whether the [OAuth 2.0 Dynamic Client Registration Protocol](https://tools.ietf.org/html/rfc7591) is supported.
	DynamicRegistrationSupported *bool `json:"dynamicRegistrationSupported,omitempty"`
	// The [registration endpoint](http://openid.net/specs/openid-connect-registration-1_0.html#ClientRegistration) of the service. A URL that starts with `https://`. For example, `https://example.com/auth/registration`.  The value of this property is used as `registration_endpoint` property in the [OpenID Provider Metadata](http://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	RegistrationEndpoint *string `json:"registrationEndpoint,omitempty"`
	// The URI of the registration management endpoint. If dynamic client registration is supported, and this is set, this URI will be used as the basis of the client's management endpoint by appending `/clientid}/` to it as a path element. If this is unset, the value of `registrationEndpoint` will be used as the URI base instead.
	RegistrationManagementEndpoint *string `json:"registrationManagementEndpoint,omitempty"`
	// The URL of the \"Policy\" of the service.  The value of this property is used as `op_policy_uri` property in the [OpenID Provider Metadata](http://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	PolicyUri *string `json:"policyUri,omitempty"`
	// The URL of the \"Terms Of Service\" of the service.  The value of this property is used as `op_tos_uri` property in the [OpenID Provider Metadata](http://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	TosUri *string `json:"tosUri,omitempty"`
	// The URL of a page where documents for developers can be found.  The value of this property is used as `service_documentation` property in the [OpenID Provider Metadata](http://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	ServiceDocumentation *string `json:"serviceDocumentation,omitempty"`
	// The URI of backchannel authentication endpoint, which is defined in the specification of [CIBA (Client Initiated Backchannel Authentication)](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html).
	BackchannelAuthenticationEndpoint *string `json:"backchannelAuthenticationEndpoint,omitempty"`
	// The supported backchannel token delivery modes. This property corresponds to the `backchannel_token_delivery_modes_supported` metadata.  Backchannel token delivery modes are defined in the specification of [CIBA (Client Initiated Backchannel Authentication)](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html).
	SupportedBackchannelTokenDeliveryModes []DeliveryMode `json:"supportedBackchannelTokenDeliveryModes,omitempty"`
	// The duration of backchannel authentication request IDs issued from the backchannel authentication endpoint in seconds. This is used as the value of the `expires_in` property in responses from the backchannel authentication endpoint.
	BackchannelAuthReqIdDuration *int32 `json:"backchannelAuthReqIdDuration,omitempty"`
	// The minimum interval between polling requests to the token endpoint from client applications in seconds. This is used as the value of the `interval` property in responses from the backchannel authentication endpoint.
	BackchannelPollingInterval *int32 `json:"backchannelPollingInterval,omitempty"`
	// The boolean flag which indicates whether the `user_code` request parameter is supported at the backchannel authentication endpoint. This property corresponds to the `backchannel_user_code_parameter_supported` metadata.
	BackchannelUserCodeParameterSupported *bool `json:"backchannelUserCodeParameterSupported,omitempty"`
	// The flag to indicate whether the `binding_message` request parameter is always required whenever a backchannel authentication request is judged as a request for Financial-grade API.  The FAPI-CIBA profile requires that the authorization server _\"shall ensure unique authorization context exists in the authorization request or require a `binding_message` in the authorization request\"_ (FAPI-CIBA, 5.2.2, 2). The simplest way to fulfill this requirement is to set this property to `true`.  If this property is set to `false`, the `binding_message` request parameter remains optional even in FAPI context, but in exchange, your authorization server must implement a custom mechanism that ensures each backchannel authentication request has unique context.
	BackchannelBindingMessageRequiredInFapi *bool `json:"backchannelBindingMessageRequiredInFapi,omitempty"`
	// The URI of the device authorization endpoint.  Device authorization endpoint is defined in the specification of OAuth 2.0 Device Authorization Grant.
	DeviceAuthorizationEndpoint *string `json:"deviceAuthorizationEndpoint,omitempty"`
	// The verification URI for the device flow. This URI is used as the value of the `verification_uri` parameter in responses from the device authorization endpoint.
	DeviceVerificationUri *string `json:"deviceVerificationUri,omitempty"`
	// The verification URI for the device flow with a placeholder for a user code. This URI is used to build the value of the `verification_uri_complete` parameter in responses from the device authorization endpoint.  It is expected that the URI contains a fixed string `USER_CODE` somewhere as a placeholder for a user code. For example, like the following.  `https://example.com/device?user\\_code=USER\\_CODE`  The fixed string is replaced with an actual user code when Authlete builds a verification URI with a user code for the `verification_uri_complete` parameter.  If this URI is not set, the `verification_uri_complete` parameter won't appear in device authorization responses.
	DeviceVerificationUriComplete *string `json:"deviceVerificationUriComplete,omitempty"`
	// The duration of device verification codes and end-user verification codes issued from the device authorization endpoint in seconds. This is used as the value of the `expires_in` property in responses from the device authorization endpoint.
	DeviceFlowCodeDuration *int32 `json:"deviceFlowCodeDuration,omitempty"`
	// The minimum interval between polling requests to the token endpoint from client applications in seconds in device flow. This is used as the value of the `interval` property in responses from the device authorization endpoint.
	DeviceFlowPollingInterval *int32           `json:"deviceFlowPollingInterval,omitempty"`
	UserCodeCharset           *UserCodeCharset `json:"userCodeCharset,omitempty"`
	// The length of end-user verification codes (`user_code`) for Device Flow.
	UserCodeLength *int32 `json:"userCodeLength,omitempty"`
	// Trust frameworks supported by this service. This corresponds to the `trust_frameworks_supported` [metadata](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html#rfc.section.7).
	SupportedTrustFrameworks []string `json:"supportedTrustFrameworks,omitempty"`
	// Evidence supported by this service. This corresponds to the `evidence_supported` [metadata](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html#rfc.section.7).
	SupportedEvidence []string `json:"supportedEvidence,omitempty"`
	// Identity documents supported by this service. This corresponds to the `id_documents_supported` [metadata](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html#rfc.section.7).
	SupportedIdentityDocuments []string `json:"supportedIdentityDocuments,omitempty"`
	// Verification methods supported by this service. This corresponds to the `id_documents_verification_methods_supported` [metadata](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html#rfc.section.7).
	SupportedVerificationMethods []string `json:"supportedVerificationMethods,omitempty"`
	// Verified claims supported by this service. This corresponds to the `claims_in_verified_claims_supported` [metadata](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html#rfc.section.7).
	SupportedVerifiedClaims []string `json:"supportedVerifiedClaims,omitempty"`
	// OIDC4IDA / verifiedClaimsValidationSchemaSet
	VerifiedClaimsValidationSchemaSet *VerifiedClaimsValidationSchema `json:"verifiedClaimsValidationSchemaSet,omitempty"`
	// The attributes of this service.
	Attributes []Pair `json:"attributes,omitempty"`
	// The flag indicating whether the nbf claim in the request object is optional even when the authorization request is regarded as a FAPI-Part2 request.  The final version of Financial-grade API was approved in January, 2021. The Part 2 of the final version has new requirements on lifetime of request objects. They require that request objects contain an `nbf` claim and the lifetime computed by `exp` - `nbf` be no longer than 60 minutes.  Therefore, when an authorization request is regarded as a FAPI-Part2 request, the request object used in the authorization request must contain an nbf claim. Otherwise, the authorization server rejects the authorization request.  When this flag is `true`, the `nbf` claim is treated as an optional claim even when the authorization request is regarded as a FAPI-Part2 request. That is, the authorization server does not perform the validation on lifetime of the request object.  Skipping the validation is a violation of the FAPI specification. The reason why this flag has been prepared nevertheless is that the new requirements (which do not exist in the Implementer's Draft 2 released in October, 2018) have big impacts on deployed implementations of client applications and Authlete thinks there should be a mechanism whereby to make the migration from ID2 to Final smooth without breaking live systems.
	NbfOptional *bool `json:"nbfOptional,omitempty"`
	// The flag indicating whether generation of the iss response parameter is suppressed.  \"OAuth 2.0 Authorization Server Issuer Identifier in Authorization Response\" has defined a new authorization response parameter, `iss`, as a countermeasure for a certain type of mix-up attacks.  The specification requires that the `iss` response parameter always be included in authorization responses unless JARM (JWT Secured Authorization Response Mode) is used.  When this flag is `true`, the authorization server does not include the `iss` response parameter in authorization responses. By turning this flag on and off, developers of client applications can experiment the mix-up attack and the effect of the `iss` response parameter.  Note that this flag should not be `true` in production environment unless there are special reasons for it.
	IssSuppressed *bool `json:"issSuppressed,omitempty"`
	// custom client metadata supported by this service.  Standard specifications define client metadata as necessary. The following are such examples.  * [OpenID Connect Dynamic Client Registration 1.0](https://openid.net/specs/openid-connect-registration-1_0.html) * [RFC 7591 OAuth 2.0 Dynamic Client Registration Protocol](https://www.rfc-editor.org/rfc/rfc7591.html) * [RFC 8705 OAuth 2.0 Mutual-TLS Client Authentication and Certificate-Bound Access Tokens](https://www.rfc-editor.org/rfc/rfc8705.html) * [OpenID Connect Client-Initiated Backchannel Authentication Flow - Core 1.0](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html) * [The OAuth 2.0 Authorization Framework: JWT Secured Authorization Request (JAR)](https://datatracker.ietf.org/doc/draft-ietf-oauth-jwsreq/) * [Financial-grade API: JWT Secured Authorization Response Mode for OAuth 2.0 (JARM)](https://openid.net/specs/openid-financial-api-jarm.html) * [OAuth 2.0 Pushed Authorization Requests (PAR)](https://datatracker.ietf.org/doc/rfc9126/) * [OAuth 2.0 Rich Authorization Requests (RAR)](https://datatracker.ietf.org/doc/draft-ietf-oauth-rar/)  Standard client metadata included in Client Registration Request and Client Update Request (cf. [OIDC DynReg](https://openid.net/specs/openid-connect-registration-1_0.html), [RFC 7591](https://www.rfc-editor.org/rfc/rfc7591.html) and [RFC 7592](https://www.rfc-editor.org/rfc/rfc7592.html)) are, if supported by Authlete, stored into Authlete database. On the other hand, unrecognized client metadata are discarded.  By listing up custom client metadata in advance by using this property (`supportedCustomClientMetadata`), Authlete can recognize them and stores their values into the database. The stored custom client metadata values can be referenced by `customMetadata`.
	SupportedCustomClientMetadata []string `json:"supportedCustomClientMetadata,omitempty"`
	// The flag indicating whether the expiration date of an access token never exceeds that of the corresponding refresh token.  When a new access token is issued by a refresh token request (= a token request with `grant_type=refresh_token`), the expiration date of the access token may exceed the expiration date of the corresponding refresh token. This behavior itself is not wrong and may happen when `refreshTokenKept` is `true` and/or when `refreshTokenDurationKept` is `true`.  When this flag is `true`, the expiration date of an access token never exceeds that of the corresponding refresh token regardless of the calculated duration based on other settings such as `accessTokenDuration`, `accessTokenDuration` in `extension` and `access_token.duration` scope attribute.  It is technically possible to set a value which is bigger than the duration of refresh tokens as the duration of access tokens although it is strange. In the case, the duration of an access token becomes longer than the duration of the refresh token which is issued together with the access token. Even if the duration values are configured so, if this flag is `true`, the expiration date of the access token does not exceed that of the refresh token. That is, the duration of the access token will be shortened, and as a result, the access token and the refresh token will have the same expiration date.
	TokenExpirationLinked *bool `json:"tokenExpirationLinked,omitempty"`
	// The flag indicating whether encryption of request object is required when the request object is passed through the front channel.  This flag does not affect the processing of request objects at the Pushed Authorization Request Endpoint, which is defined in [OAuth 2.0 Pushed Authorization Requests](https://datatracker.ietf.org/doc/rfc9126/). Unecrypted request objects are accepted at the endpoint even if this flag is `true`.  This flag does not indicate whether a request object is always required. There is a different flag, `requestObjectRequired`, for the purpose. See the description of `requestObjectRequired` for details.  Even if this flag is `false`, encryption of request object is required if the `frontChannelRequestObjectEncryptionRequired` flag of the client is `true`.
	FrontChannelRequestObjectEncryptionRequired *bool `json:"frontChannelRequestObjectEncryptionRequired,omitempty"`
	// The flag indicating whether the JWE alg of encrypted request object must match the `request_object_encryption_alg` client metadata of the client that has sent the request object.  The request_object_encryption_alg client metadata itself is defined in [OpenID Connect Dynamic Client Registration 1.0](https://openid.net/specs/openid-connect-registration-1_0.html) as follows.  > request_object_encryption_alg > > OPTIONAL. JWE [JWE] alg algorithm [JWA] the RP is declaring that it may use for encrypting Request Objects sent to the OP. This parameter SHOULD be included when symmetric encryption will be used, since this signals to the OP that a client_secret value needs to be returned from which the symmetric key will be derived, that might not otherwise be returned. The RP MAY still use other supported encryption algorithms or send unencrypted Request Objects, even when this parameter is present. If both signing and encryption are requested, the Request Object will be signed then encrypted, with the result being a Nested JWT, as defined in [JWT]. The default, if omitted, is that the RP is not declaring whether it might encrypt any Request Objects.  The point here is \"The RP MAY still use other supported encryption algorithms or send unencrypted Request Objects, even when this parameter is present.\"  The Client's property that represents the client metadata is `requestEncryptionAlg`. See the description of `requestEncryptionAlg` for details.  Even if this flag is `false`, the match is required if the `requestObjectEncryptionAlgMatchRequired` flag of the client is `true`.
	RequestObjectEncryptionAlgMatchRequired *bool `json:"requestObjectEncryptionAlgMatchRequired,omitempty"`
	// The flag indicating whether the JWE `enc` of encrypted request object must match the `request_object_encryption_enc` client metadata of the client that has sent the request object.  The `request_object_encryption_enc` client metadata itself is defined in [OpenID Connect Dynamic Client Registration 1.0](https://openid.net/specs/openid-connect-registration-1_0.html) as follows.  > request_object_encryption_enc > > OPTIONAL. JWE enc algorithm [JWA] the RP is declaring that it may use for encrypting Request Objects sent to the OP. If request_object_encryption_alg is specified, the default for this value is A128CBC-HS256. When request_object_encryption_enc is included, request_object_encryption_alg MUST also be provided.  The Client's property that represents the client metadata is `requestEncryptionEnc`. See the description of `requestEncryptionEnc` for details.  Even if this flag is false, the match is required if the `requestObjectEncryptionEncMatchRequired` flag is `true`.
	RequestObjectEncryptionEncMatchRequired *bool `json:"requestObjectEncryptionEncMatchRequired,omitempty"`
	// The flag indicating whether HSM (Hardware Security Module) support is enabled for this service.  When this flag is `false`, keys managed in HSMs are not used even if they exist. In addition, `/api/hsk/_*` APIs reject all requests.  Even if this flag is `true`, HSM-related features do not work if the configuration of the Authlete server you are using does not support HSM.
	HsmEnabled *bool `json:"hsmEnabled,omitempty"`
	// The information about keys managed on HSMs (Hardware Security Modules).  This `hsks` property is output only, meaning that `hsks` in requests to `/api/service/create` API and `/api/service/update` API do not have any effect. The contents of this property is controlled only by `/api/hsk/_*` APIs.
	Hsks []Hsk `json:"hsks,omitempty"`
	// The URL of the grant management endpoint.
	GrantManagementEndpoint *string `json:"grantManagementEndpoint,omitempty"`
	// The flag indicating whether every authorization request (and any request serving as an authorization request such as CIBA backchannel authentication request and device authorization request) must include the `grant_management_action` request parameter.  This property corresponds to the `grant_management_action_required` server metadata defined in [Grant Management for OAuth 2.0](https://openid.net/specs/fapi-grant-management.html).  Note that setting true to this property will result in blocking all public clients because the specification requires that grant management be usable only by confidential clients for security reasons.
	GrantManagementActionRequired *bool `json:"grantManagementActionRequired,omitempty"`
	// The flag indicating whether Authlete's `/api/client/registration` API uses `UNAUTHORIZED` as a value of the `action` response parameter when appropriate.  The `UNAUTHORIZED` enum value was initially not defined as a possible value of the `action` parameter in an `/api/client/registration` API response. This means that implementations of client `configuration` endpoint were not able to conform to [RFC 7592](https://www.rfc-editor.org/rfc/rfc7592.html) strictly.  For backward compatibility (to avoid breaking running systems), Authlete's `/api/client/registration` API does not return the `UNAUTHORIZED` enum value if this flag is not turned on.  The steps an existing implementation of client configuration endpoint has to do in order to conform to the requirement related to \"401 Unauthorized\" are as follows.  1. Update the Authlete library (e.g. authlete-java-common) your system is using. 2. Update your implementation of client configuration endpoint so that it can handle the `UNAUTHORIZED` action. 3. Turn on this `unauthorizedOnClientConfigSupported` flag.
	UnauthorizedOnClientConfigSupported *bool `json:"unauthorizedOnClientConfigSupported,omitempty"`
	// The flag indicating whether the `scope` request parameter in dynamic client registration and update requests (RFC 7591 and RFC 7592) is used as scopes that the client can request.  Limiting the range of scopes that a client can request is achieved by listing scopes in the `client.extension.requestableScopes` property and setting the `client.extension.requestableScopesEnabled` property to `true`. This feature is called \"requestable scopes\".  This property affects behaviors of `/api/client/registration` and other family APIs.
	DcrScopeUsedAsRequestable *bool `json:"dcrScopeUsedAsRequestable,omitempty"`
	// The endpoint for clients ending the sessions.  A URL that starts with `https://` and has no fragment component. For example, `https://example.com/auth/endSession`.  The value of this property is used as `end_session_endpoint` property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	EndSessionEndpoint *string `json:"endSessionEndpoint,omitempty"`
	// The flag indicating whether the port number component of redirection URIs can be variable when the host component indicates loopback.  When this flag is `true`, if the host component of a redirection URI specified in an authorization request indicates loopback (to be precise, when the host component is localhost, `127.0.0.1` or `::1`), the port number component is ignored when the specified redirection URI is compared to pre-registered ones. This behavior is described in [7.3. Loopback Interface Redirection]( https://www.rfc-editor.org/rfc/rfc8252.html#section-7.3) of [RFC 8252 OAuth 2.0](https://www.rfc-editor.org/rfc/rfc8252.html) for Native Apps.  [3.1.2.3. Dynamic Configuration](https://www.rfc-editor.org/rfc/rfc6749.html#section-3.1.2.3) of [RFC 6749](https://www.rfc-editor.org/rfc/rfc6749.html) states _\"If the client registration included the full redirection URI, the authorization server MUST compare the two URIs using simple string comparison as defined in [RFC3986] Section 6.2.1.\"_ Also, the description of `redirect_uri` in [3.1.2.1. Authentication Request](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest) of [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html) states _\"This URI MUST exactly match one of the Redirection URI values for the Client pre-registered at the OpenID Provider, with the matching performed as described in Section 6.2.1 of [RFC3986] (**Simple String Comparison**).\"_ These \"Simple String Comparison\" requirements are preceded by this flag. That is, even when the conditions described in RFC 6749 and OpenID Connect Core 1.0 are satisfied, the port number component of loopback redirection URIs can be variable when this flag is `true`.  [8.3. Loopback Redirect Considerations](https://www.rfc-editor.org/rfc/rfc8252.html#section-8.3) of [RFC 8252](https://www.rfc-editor.org/rfc/rfc8252.html) states as follows.  > While redirect URIs using localhost (i.e., `\"http://localhost:{port}/{path}\"`) function similarly to loopback IP redirects described in Section 7.3, the use of localhost is NOT RECOMMENDED. Specifying a redirect URI with the loopback IP literal rather than localhost avoids inadvertently listening on network interfaces other than the loopback interface. It is also less susceptible to client-side firewalls and misconfigured host name resolution on the user's device.  However, Authlete allows the port number component to be variable in the case of `localhost`, too. It is left to client applications whether they use `localhost` or a literal loopback IP address (`127.0.0.1` for IPv4 or `::1` for IPv6).  Section 7.3 and Section 8.3 of [RFC 8252](https://www.rfc-editor.org/rfc/rfc8252.html) state that loopback redirection URIs use the `\"http\"` scheme, but Authlete allows the port number component to be variable in other cases (e.g. in the case of the `\"https\"` scheme), too.
	LoopbackRedirectionUriVariable *bool `json:"loopbackRedirectionUriVariable,omitempty"`
	// The flag indicating whether Authlete checks whether the `aud` claim of request objects matches the issuer identifier of this service.  [Section 6.1. Passing a Request Object by Value](https://openid.net/specs/openid-connect-core-1_0.html#JWTRequests) of [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html) has the following statement.  > The `aud` value SHOULD be or include the OP's Issuer Identifier URL.  Likewise, [Section 4. Request Object](https://www.rfc-editor.org/rfc/rfc9101.html#section-4) of [RFC 9101](https://www.rfc-editor.org/rfc/rfc9101.html) (The OAuth 2.0 Authorization Framework: JWT-Secured Authorization Request (JAR)) has the following statement.  > The value of aud should be the value of the authorization server (AS) issuer, as defined in [RFC 8414](https://www.rfc-editor.org/rfc/rfc8414.html).  As excerpted above, validation on the `aud` claim of request objects is optional. However, if this flag is turned on, Authlete checks whether the `aud` claim of request objects matches the issuer identifier of this service and raises an error if they are different.
	RequestObjectAudienceChecked *bool `json:"requestObjectAudienceChecked,omitempty"`
	// The flag indicating whether Authlete generates access tokens for external attachments and embeds them in ID tokens and userinfo responses.
	AccessTokenForExternalAttachmentEmbedded *bool `json:"accessTokenForExternalAttachmentEmbedded,omitempty"`
	// Identifiers of entities that can issue entity statements for this service. This property corresponds to the `authority_hints` property that appears in a self-signed entity statement that is defined in OpenID Connect Federation 1.0.
	AuthorityHints []string `json:"authorityHints,omitempty"`
	// flag indicating whether this service supports OpenID Connect Federation 1
	FederationEnabled *bool `json:"federationEnabled,omitempty"`
	// JWK Set document containing keys that are used to sign (1) self-signed entity statement of this service and (2) the response from `signed_jwks_uri`.
	FederationJwks *string `json:"federationJwks,omitempty"`
	// A key ID to identify a JWK used to sign the entity configuration and the signed JWK Set.
	FederationSignatureKeyId *string `json:"federationSignatureKeyId,omitempty"`
	// The duration of the entity configuration in seconds.
	FederationConfigurationDuration *int32 `json:"federationConfigurationDuration,omitempty"`
	// The URI of the federation registration endpoint. This property corresponds to the `federation_registration_endpoint` server metadata that is defined in OpenID Connect Federation 1.0.
	FederationRegistrationEndpoint *string `json:"federationRegistrationEndpoint,omitempty"`
	// The human-readable name representing the organization that operates this service. This property corresponds to the `organization_name` server metadata that is defined in OpenID Connect Federation 1.0.
	OrganizationName *string `json:"organizationName,omitempty"`
	// The transformed claims predefined by this service in JSON format. This property corresponds to the `transformed_claims_predefined` server metadata.
	PredefinedTransformedClaims *string `json:"predefinedTransformedClaims,omitempty"`
	// flag indicating whether refresh token requests with the same refresh token can be made multiple times in quick succession and they can obtain the same renewed refresh token within the short period.
	RefreshTokenIdempotent *bool `json:"refreshTokenIdempotent,omitempty"`
	// The URI of the endpoint that returns this service's JWK Set document in the JWT format. This property corresponds to the `signed_jwks_uri` server metadata defined in OpenID Connect Federation 1.0.
	SignedJwksUri *string `json:"signedJwksUri,omitempty"`
	// Supported attachment types. This property corresponds to the {@code attachments_supported} server metadata which was added by the third implementer's draft of OpenID Connect for Identity Assurance 1.0.
	SupportedAttachments []AttachmentType `json:"supportedAttachments,omitempty"`
	// Supported algorithms used to compute digest values of external attachments. This property corresponds to the `digest_algorithms_supported` server metadata which was added by the third implementer's draft of OpenID Connect for Identity Assurance 1.0.
	SupportedDigestAlgorithms []string `json:"supportedDigestAlgorithms,omitempty"`
	// Document types supported by this service. This property corresponds to the `documents_supported` server metadata.
	SupportedDocuments []string `json:"supportedDocuments,omitempty"`
	// validation and verification processes supported by this service. This property corresponds to the `documents_methods_supported` server metadata.  The third implementer's draft of [OpenID Connect for Identity Assurance 1.0](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html) renamed the `id_documents_verification_methods_supported` server metadata to `documents_methods_supported`.
	SupportedDocumentsMethods []string `json:"supportedDocumentsMethods,omitempty"`
	// Document validation methods supported by this service. This property corresponds to the `documents_validation_methods_supported` server metadata which was added by the third implementer's draft of <a href= [OpenID Connect for Identity Assurance 1.0](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html)
	SupportedDocumentsValidationMethods []string `json:"supportedDocumentsValidationMethods,omitempty"`
	// Document verification methods supported by this service. This property corresponds to the `documents_verification_methods_supported` server metadata which was added by the third implementer's draft of [OpenID Connect for Identity Assurance 1.0](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html)
	SupportedDocumentsVerificationMethods []string `json:"supportedDocumentsVerificationMethods,omitempty"`
	// Electronic record types supported by this service. This property corresponds to the `electronic_records_supported` server metadata which was added by the third implementer's draft of [OpenID Connect for Identity Assurance 1.0](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html)
	SupportedElectronicRecords       []string                 `json:"supportedElectronicRecords,omitempty"`
	SupportedClientRegistrationTypes []ClientRegistrationType `json:"supportedClientRegistrationTypes,omitempty"`
	// The flag indicating whether to prohibit unidentifiable clients from making token exchange requests.
	TokenExchangeByIdentifiableClientsOnly *bool `json:"tokenExchangeByIdentifiableClientsOnly,omitempty"`
	// The flag indicating whether to prohibit public clients from making token exchange requests.
	TokenExchangeByConfidentialClientsOnly *bool `json:"tokenExchangeByConfidentialClientsOnly,omitempty"`
	// The flag indicating whether to prohibit clients that have no explicit permission from making token exchange requests.
	TokenExchangeByPermittedClientsOnly *bool `json:"tokenExchangeByPermittedClientsOnly,omitempty"`
	// The flag indicating whether to reject token exchange requests which use encrypted JWTs as input tokens.
	TokenExchangeEncryptedJwtRejected *bool `json:"tokenExchangeEncryptedJwtRejected,omitempty"`
	// The flag indicating whether to reject token exchange requests which use unsigned JWTs as input tokens.
	TokenExchangeUnsignedJwtRejected *bool `json:"tokenExchangeUnsignedJwtRejected,omitempty"`
	// The flag indicating whether to prohibit unidentifiable clients from using the grant type \"urn:ietf:params:oauth:grant-type:jwt-bearer\".
	JwtGrantByIdentifiableClientsOnly *bool `json:"jwtGrantByIdentifiableClientsOnly,omitempty"`
	// The flag indicating whether to reject token requests that use an encrypted JWT as an authorization grant with the grant type \"urn:ietf:params:oauth:grant-type:jwt-bearer\".
	JwtGrantEncryptedJwtRejected *bool `json:"jwtGrantEncryptedJwtRejected,omitempty"`
	// The flag indicating whether to reject token requests that use an unsigned JWT as an authorization grant with the grant type \"urn:ietf:params:oauth:grant-type:jwt-bearer\".
	JwtGrantUnsignedJwtRejected *bool `json:"jwtGrantUnsignedJwtRejected,omitempty"`
	// The flag indicating whether to block DCR (Dynamic Client Registration) requests whose \"software_id\" has already been used previously.
	DcrDuplicateSoftwareIdBlocked *bool `json:"dcrDuplicateSoftwareIdBlocked,omitempty"`
	// The trust anchors that are referenced when this service resolves trust chains of relying parties.  If this property is empty, client registration fails regardless of whether its type is `automatic` or `explicit`. It means that OpenID Connect Federation 1.0 does not work.
	TrustAnchors []TrustAnchor `json:"trustAnchors,omitempty"`
	// The flag indicating whether the openid scope should be dropped from scopes list assigned to access token issued when a refresh token grant is used.
	OpenidDroppedOnRefreshWithoutOfflineAccess *bool `json:"openidDroppedOnRefreshWithoutOfflineAccess,omitempty"`
	// Supported document check methods. This property corresponds to the `documents_check_methods_supported` server metadata which was added by the fourth implementer's draft of OpenID Connect for Identity Assurance 1.0.
	SupportedDocumentsCheckMethods []string `json:"supportedDocumentsCheckMethods,omitempty"`
	// The flag indicating whether this service signs responses from the resource server.
	RsResponseSigned *bool `json:"rsResponseSigned,omitempty"`
	// The duration of `c_nonce`.
	CnonceDuration *int64 `json:"cnonceDuration,omitempty"`
	// Whether to require DPoP proof JWTs to include the `nonce` claim whenever they are presented.
	DpopNonceRequired *bool `json:"dpopNonceRequired,omitempty"`
	// Get the flag indicating whether the feature of Verifiable Credentials for this service is enabled or not.
	VerifiableCredentialsEnabled *bool `json:"verifiableCredentialsEnabled,omitempty"`
	// The URL at which the JWK Set document of the credential issuer is exposed.
	CredentialJwksUri *string `json:"credentialJwksUri,omitempty"`
	// The default duration of credential offers in seconds.
	CredentialOfferDuration *int64 `json:"credentialOfferDuration,omitempty"`
	// The duration of nonce values for DPoP proof JWTs in seconds.
	DpopNonceDuration *int64 `json:"dpopNonceDuration,omitempty"`
	// The flag indicating whether token requests using the pre-authorized code grant flow by unidentifiable clients are allowed.
	PreAuthorizedGrantAnonymousAccessSupported *bool `json:"preAuthorizedGrantAnonymousAccessSupported,omitempty"`
	// The duration of transaction ID in seconds that may be issued as a result of a credential request or a batch credential request.
	CredentialTransactionDuration *int64 `json:"credentialTransactionDuration,omitempty"`
	// The key ID of the key for signing introspection responses.
	IntrospectionSignatureKeyId *string `json:"introspectionSignatureKeyId,omitempty"`
	// The key ID of the key for signing introspection responses.
	ResourceSignatureKeyId *string `json:"resourceSignatureKeyId,omitempty"`
	// The default length of user PINs.
	UserPinLength *int32 `json:"userPinLength,omitempty"`
	// The supported `prompt` values.
	SupportedPromptValues []Prompt `json:"supportedPromptValues,omitempty"`
	// The flag indicating whether to enable the feature of ID token reissuance in the refresh token flow.
	IdTokenReissuable *bool `json:"idTokenReissuable,omitempty"`
	// The JWK Set document containing private keys that are used to sign verifiable credentials.
	CredentialJwks *string `json:"credentialJwks,omitempty"`
	// FAPI modes for this service.  When the value of this property is not `null`, Authlete always processes requests to this service based on the specified FAPI modes if the FAPI feature is enabled in Authlete and the FAPI profile is supported by this service.  For instance, when this property is set to an array containing `FAPI1_ADVANCED` only, Authlete always processes requests to this service based on \"Financial-grade API Security Profile 1.0 - Part 2: Advanced\" if the FAPI feature is enabled in Authlete and the FAPI profile is supported by this service.
	FapiModes []FapiMode `json:"fapiModes,omitempty"`
	// The default duration of verifiable credentials in seconds.
	CredentialDuration       *int64                    `json:"credentialDuration,omitempty"`
	CredentialIssuerMetadata *CredentialIssuerMetadata `json:"credentialIssuerMetadata,omitempty"`
	// The type of the `aud` claim in ID tokens.
	IdTokenAudType *string `json:"idTokenAudType,omitempty"`
}

// NewService instantiates a new Service object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewService() *Service {
	this := Service{}
	return &this
}

// NewServiceWithDefaults instantiates a new Service object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWithDefaults() *Service {
	this := Service{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *Service) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *Service) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *Service) SetNumber(v int32) {
	o.Number = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *Service) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *Service) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *Service) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *Service) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *Service) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *Service) SetIssuer(v string) {
	o.Issuer = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Service) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Service) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Service) SetDescription(v string) {
	o.Description = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *Service) GetApiKey() int64 {
	if o == nil || IsNil(o.ApiKey) {
		var ret int64
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetApiKeyOk() (*int64, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *Service) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given int64 and assigns it to the ApiKey field.
func (o *Service) SetApiKey(v int64) {
	o.ApiKey = &v
}

// GetClientIdAliasEnabled returns the ClientIdAliasEnabled field value if set, zero value otherwise.
func (o *Service) GetClientIdAliasEnabled() bool {
	if o == nil || IsNil(o.ClientIdAliasEnabled) {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasEnabled
}

// GetClientIdAliasEnabledOk returns a tuple with the ClientIdAliasEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetClientIdAliasEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientIdAliasEnabled) {
		return nil, false
	}
	return o.ClientIdAliasEnabled, true
}

// HasClientIdAliasEnabled returns a boolean if a field has been set.
func (o *Service) HasClientIdAliasEnabled() bool {
	if o != nil && !IsNil(o.ClientIdAliasEnabled) {
		return true
	}

	return false
}

// SetClientIdAliasEnabled gets a reference to the given bool and assigns it to the ClientIdAliasEnabled field.
func (o *Service) SetClientIdAliasEnabled(v bool) {
	o.ClientIdAliasEnabled = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Service) GetMetadata() []Pair {
	if o == nil || IsNil(o.Metadata) {
		var ret []Pair
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetMetadataOk() ([]Pair, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Service) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []Pair and assigns it to the Metadata field.
func (o *Service) SetMetadata(v []Pair) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Service) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Service) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *Service) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *Service) GetModifiedAt() int64 {
	if o == nil || IsNil(o.ModifiedAt) {
		var ret int64
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetModifiedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ModifiedAt) {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *Service) HasModifiedAt() bool {
	if o != nil && !IsNil(o.ModifiedAt) {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given int64 and assigns it to the ModifiedAt field.
func (o *Service) SetModifiedAt(v int64) {
	o.ModifiedAt = &v
}

// GetAuthenticationCallbackEndpoint returns the AuthenticationCallbackEndpoint field value if set, zero value otherwise.
func (o *Service) GetAuthenticationCallbackEndpoint() string {
	if o == nil || IsNil(o.AuthenticationCallbackEndpoint) {
		var ret string
		return ret
	}
	return *o.AuthenticationCallbackEndpoint
}

// GetAuthenticationCallbackEndpointOk returns a tuple with the AuthenticationCallbackEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAuthenticationCallbackEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationCallbackEndpoint) {
		return nil, false
	}
	return o.AuthenticationCallbackEndpoint, true
}

// HasAuthenticationCallbackEndpoint returns a boolean if a field has been set.
func (o *Service) HasAuthenticationCallbackEndpoint() bool {
	if o != nil && !IsNil(o.AuthenticationCallbackEndpoint) {
		return true
	}

	return false
}

// SetAuthenticationCallbackEndpoint gets a reference to the given string and assigns it to the AuthenticationCallbackEndpoint field.
func (o *Service) SetAuthenticationCallbackEndpoint(v string) {
	o.AuthenticationCallbackEndpoint = &v
}

// GetAuthenticationCallbackApiKey returns the AuthenticationCallbackApiKey field value if set, zero value otherwise.
func (o *Service) GetAuthenticationCallbackApiKey() string {
	if o == nil || IsNil(o.AuthenticationCallbackApiKey) {
		var ret string
		return ret
	}
	return *o.AuthenticationCallbackApiKey
}

// GetAuthenticationCallbackApiKeyOk returns a tuple with the AuthenticationCallbackApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAuthenticationCallbackApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationCallbackApiKey) {
		return nil, false
	}
	return o.AuthenticationCallbackApiKey, true
}

// HasAuthenticationCallbackApiKey returns a boolean if a field has been set.
func (o *Service) HasAuthenticationCallbackApiKey() bool {
	if o != nil && !IsNil(o.AuthenticationCallbackApiKey) {
		return true
	}

	return false
}

// SetAuthenticationCallbackApiKey gets a reference to the given string and assigns it to the AuthenticationCallbackApiKey field.
func (o *Service) SetAuthenticationCallbackApiKey(v string) {
	o.AuthenticationCallbackApiKey = &v
}

// GetAuthenticationCallbackApiSecret returns the AuthenticationCallbackApiSecret field value if set, zero value otherwise.
func (o *Service) GetAuthenticationCallbackApiSecret() string {
	if o == nil || IsNil(o.AuthenticationCallbackApiSecret) {
		var ret string
		return ret
	}
	return *o.AuthenticationCallbackApiSecret
}

// GetAuthenticationCallbackApiSecretOk returns a tuple with the AuthenticationCallbackApiSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAuthenticationCallbackApiSecretOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationCallbackApiSecret) {
		return nil, false
	}
	return o.AuthenticationCallbackApiSecret, true
}

// HasAuthenticationCallbackApiSecret returns a boolean if a field has been set.
func (o *Service) HasAuthenticationCallbackApiSecret() bool {
	if o != nil && !IsNil(o.AuthenticationCallbackApiSecret) {
		return true
	}

	return false
}

// SetAuthenticationCallbackApiSecret gets a reference to the given string and assigns it to the AuthenticationCallbackApiSecret field.
func (o *Service) SetAuthenticationCallbackApiSecret(v string) {
	o.AuthenticationCallbackApiSecret = &v
}

// GetSupportedAcrs returns the SupportedAcrs field value if set, zero value otherwise.
func (o *Service) GetSupportedAcrs() []string {
	if o == nil || IsNil(o.SupportedAcrs) {
		var ret []string
		return ret
	}
	return o.SupportedAcrs
}

// GetSupportedAcrsOk returns a tuple with the SupportedAcrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedAcrsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedAcrs) {
		return nil, false
	}
	return o.SupportedAcrs, true
}

// HasSupportedAcrs returns a boolean if a field has been set.
func (o *Service) HasSupportedAcrs() bool {
	if o != nil && !IsNil(o.SupportedAcrs) {
		return true
	}

	return false
}

// SetSupportedAcrs gets a reference to the given []string and assigns it to the SupportedAcrs field.
func (o *Service) SetSupportedAcrs(v []string) {
	o.SupportedAcrs = v
}

// GetSupportedGrantTypes returns the SupportedGrantTypes field value if set, zero value otherwise.
func (o *Service) GetSupportedGrantTypes() []GrantType {
	if o == nil || IsNil(o.SupportedGrantTypes) {
		var ret []GrantType
		return ret
	}
	return o.SupportedGrantTypes
}

// GetSupportedGrantTypesOk returns a tuple with the SupportedGrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedGrantTypesOk() ([]GrantType, bool) {
	if o == nil || IsNil(o.SupportedGrantTypes) {
		return nil, false
	}
	return o.SupportedGrantTypes, true
}

// HasSupportedGrantTypes returns a boolean if a field has been set.
func (o *Service) HasSupportedGrantTypes() bool {
	if o != nil && !IsNil(o.SupportedGrantTypes) {
		return true
	}

	return false
}

// SetSupportedGrantTypes gets a reference to the given []GrantType and assigns it to the SupportedGrantTypes field.
func (o *Service) SetSupportedGrantTypes(v []GrantType) {
	o.SupportedGrantTypes = v
}

// GetSupportedResponseTypes returns the SupportedResponseTypes field value if set, zero value otherwise.
func (o *Service) GetSupportedResponseTypes() []ResponseType {
	if o == nil || IsNil(o.SupportedResponseTypes) {
		var ret []ResponseType
		return ret
	}
	return o.SupportedResponseTypes
}

// GetSupportedResponseTypesOk returns a tuple with the SupportedResponseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedResponseTypesOk() ([]ResponseType, bool) {
	if o == nil || IsNil(o.SupportedResponseTypes) {
		return nil, false
	}
	return o.SupportedResponseTypes, true
}

// HasSupportedResponseTypes returns a boolean if a field has been set.
func (o *Service) HasSupportedResponseTypes() bool {
	if o != nil && !IsNil(o.SupportedResponseTypes) {
		return true
	}

	return false
}

// SetSupportedResponseTypes gets a reference to the given []ResponseType and assigns it to the SupportedResponseTypes field.
func (o *Service) SetSupportedResponseTypes(v []ResponseType) {
	o.SupportedResponseTypes = v
}

// GetSupportedAuthorizationDetailsTypes returns the SupportedAuthorizationDetailsTypes field value if set, zero value otherwise.
func (o *Service) GetSupportedAuthorizationDetailsTypes() []string {
	if o == nil || IsNil(o.SupportedAuthorizationDetailsTypes) {
		var ret []string
		return ret
	}
	return o.SupportedAuthorizationDetailsTypes
}

// GetSupportedAuthorizationDetailsTypesOk returns a tuple with the SupportedAuthorizationDetailsTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedAuthorizationDetailsTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedAuthorizationDetailsTypes) {
		return nil, false
	}
	return o.SupportedAuthorizationDetailsTypes, true
}

// HasSupportedAuthorizationDetailsTypes returns a boolean if a field has been set.
func (o *Service) HasSupportedAuthorizationDetailsTypes() bool {
	if o != nil && !IsNil(o.SupportedAuthorizationDetailsTypes) {
		return true
	}

	return false
}

// SetSupportedAuthorizationDetailsTypes gets a reference to the given []string and assigns it to the SupportedAuthorizationDetailsTypes field.
func (o *Service) SetSupportedAuthorizationDetailsTypes(v []string) {
	o.SupportedAuthorizationDetailsTypes = v
}

// GetSupportedServiceProfiles returns the SupportedServiceProfiles field value if set, zero value otherwise.
func (o *Service) GetSupportedServiceProfiles() []ServiceProfile {
	if o == nil || IsNil(o.SupportedServiceProfiles) {
		var ret []ServiceProfile
		return ret
	}
	return o.SupportedServiceProfiles
}

// GetSupportedServiceProfilesOk returns a tuple with the SupportedServiceProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedServiceProfilesOk() ([]ServiceProfile, bool) {
	if o == nil || IsNil(o.SupportedServiceProfiles) {
		return nil, false
	}
	return o.SupportedServiceProfiles, true
}

// HasSupportedServiceProfiles returns a boolean if a field has been set.
func (o *Service) HasSupportedServiceProfiles() bool {
	if o != nil && !IsNil(o.SupportedServiceProfiles) {
		return true
	}

	return false
}

// SetSupportedServiceProfiles gets a reference to the given []ServiceProfile and assigns it to the SupportedServiceProfiles field.
func (o *Service) SetSupportedServiceProfiles(v []ServiceProfile) {
	o.SupportedServiceProfiles = v
}

// GetErrorDescriptionOmitted returns the ErrorDescriptionOmitted field value if set, zero value otherwise.
func (o *Service) GetErrorDescriptionOmitted() bool {
	if o == nil || IsNil(o.ErrorDescriptionOmitted) {
		var ret bool
		return ret
	}
	return *o.ErrorDescriptionOmitted
}

// GetErrorDescriptionOmittedOk returns a tuple with the ErrorDescriptionOmitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetErrorDescriptionOmittedOk() (*bool, bool) {
	if o == nil || IsNil(o.ErrorDescriptionOmitted) {
		return nil, false
	}
	return o.ErrorDescriptionOmitted, true
}

// HasErrorDescriptionOmitted returns a boolean if a field has been set.
func (o *Service) HasErrorDescriptionOmitted() bool {
	if o != nil && !IsNil(o.ErrorDescriptionOmitted) {
		return true
	}

	return false
}

// SetErrorDescriptionOmitted gets a reference to the given bool and assigns it to the ErrorDescriptionOmitted field.
func (o *Service) SetErrorDescriptionOmitted(v bool) {
	o.ErrorDescriptionOmitted = &v
}

// GetErrorUriOmitted returns the ErrorUriOmitted field value if set, zero value otherwise.
func (o *Service) GetErrorUriOmitted() bool {
	if o == nil || IsNil(o.ErrorUriOmitted) {
		var ret bool
		return ret
	}
	return *o.ErrorUriOmitted
}

// GetErrorUriOmittedOk returns a tuple with the ErrorUriOmitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetErrorUriOmittedOk() (*bool, bool) {
	if o == nil || IsNil(o.ErrorUriOmitted) {
		return nil, false
	}
	return o.ErrorUriOmitted, true
}

// HasErrorUriOmitted returns a boolean if a field has been set.
func (o *Service) HasErrorUriOmitted() bool {
	if o != nil && !IsNil(o.ErrorUriOmitted) {
		return true
	}

	return false
}

// SetErrorUriOmitted gets a reference to the given bool and assigns it to the ErrorUriOmitted field.
func (o *Service) SetErrorUriOmitted(v bool) {
	o.ErrorUriOmitted = &v
}

// GetAuthorizationEndpoint returns the AuthorizationEndpoint field value if set, zero value otherwise.
func (o *Service) GetAuthorizationEndpoint() string {
	if o == nil || IsNil(o.AuthorizationEndpoint) {
		var ret string
		return ret
	}
	return *o.AuthorizationEndpoint
}

// GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAuthorizationEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationEndpoint) {
		return nil, false
	}
	return o.AuthorizationEndpoint, true
}

// HasAuthorizationEndpoint returns a boolean if a field has been set.
func (o *Service) HasAuthorizationEndpoint() bool {
	if o != nil && !IsNil(o.AuthorizationEndpoint) {
		return true
	}

	return false
}

// SetAuthorizationEndpoint gets a reference to the given string and assigns it to the AuthorizationEndpoint field.
func (o *Service) SetAuthorizationEndpoint(v string) {
	o.AuthorizationEndpoint = &v
}

// GetDirectAuthorizationEndpointEnabled returns the DirectAuthorizationEndpointEnabled field value if set, zero value otherwise.
func (o *Service) GetDirectAuthorizationEndpointEnabled() bool {
	if o == nil || IsNil(o.DirectAuthorizationEndpointEnabled) {
		var ret bool
		return ret
	}
	return *o.DirectAuthorizationEndpointEnabled
}

// GetDirectAuthorizationEndpointEnabledOk returns a tuple with the DirectAuthorizationEndpointEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDirectAuthorizationEndpointEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DirectAuthorizationEndpointEnabled) {
		return nil, false
	}
	return o.DirectAuthorizationEndpointEnabled, true
}

// HasDirectAuthorizationEndpointEnabled returns a boolean if a field has been set.
func (o *Service) HasDirectAuthorizationEndpointEnabled() bool {
	if o != nil && !IsNil(o.DirectAuthorizationEndpointEnabled) {
		return true
	}

	return false
}

// SetDirectAuthorizationEndpointEnabled gets a reference to the given bool and assigns it to the DirectAuthorizationEndpointEnabled field.
func (o *Service) SetDirectAuthorizationEndpointEnabled(v bool) {
	o.DirectAuthorizationEndpointEnabled = &v
}

// GetSupportedUiLocales returns the SupportedUiLocales field value if set, zero value otherwise.
func (o *Service) GetSupportedUiLocales() []string {
	if o == nil || IsNil(o.SupportedUiLocales) {
		var ret []string
		return ret
	}
	return o.SupportedUiLocales
}

// GetSupportedUiLocalesOk returns a tuple with the SupportedUiLocales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedUiLocalesOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedUiLocales) {
		return nil, false
	}
	return o.SupportedUiLocales, true
}

// HasSupportedUiLocales returns a boolean if a field has been set.
func (o *Service) HasSupportedUiLocales() bool {
	if o != nil && !IsNil(o.SupportedUiLocales) {
		return true
	}

	return false
}

// SetSupportedUiLocales gets a reference to the given []string and assigns it to the SupportedUiLocales field.
func (o *Service) SetSupportedUiLocales(v []string) {
	o.SupportedUiLocales = v
}

// GetSupportedDisplays returns the SupportedDisplays field value if set, zero value otherwise.
func (o *Service) GetSupportedDisplays() []Display {
	if o == nil || IsNil(o.SupportedDisplays) {
		var ret []Display
		return ret
	}
	return o.SupportedDisplays
}

// GetSupportedDisplaysOk returns a tuple with the SupportedDisplays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedDisplaysOk() ([]Display, bool) {
	if o == nil || IsNil(o.SupportedDisplays) {
		return nil, false
	}
	return o.SupportedDisplays, true
}

// HasSupportedDisplays returns a boolean if a field has been set.
func (o *Service) HasSupportedDisplays() bool {
	if o != nil && !IsNil(o.SupportedDisplays) {
		return true
	}

	return false
}

// SetSupportedDisplays gets a reference to the given []Display and assigns it to the SupportedDisplays field.
func (o *Service) SetSupportedDisplays(v []Display) {
	o.SupportedDisplays = v
}

// GetPkceRequired returns the PkceRequired field value if set, zero value otherwise.
func (o *Service) GetPkceRequired() bool {
	if o == nil || IsNil(o.PkceRequired) {
		var ret bool
		return ret
	}
	return *o.PkceRequired
}

// GetPkceRequiredOk returns a tuple with the PkceRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetPkceRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.PkceRequired) {
		return nil, false
	}
	return o.PkceRequired, true
}

// HasPkceRequired returns a boolean if a field has been set.
func (o *Service) HasPkceRequired() bool {
	if o != nil && !IsNil(o.PkceRequired) {
		return true
	}

	return false
}

// SetPkceRequired gets a reference to the given bool and assigns it to the PkceRequired field.
func (o *Service) SetPkceRequired(v bool) {
	o.PkceRequired = &v
}

// GetPkceS256Required returns the PkceS256Required field value if set, zero value otherwise.
func (o *Service) GetPkceS256Required() bool {
	if o == nil || IsNil(o.PkceS256Required) {
		var ret bool
		return ret
	}
	return *o.PkceS256Required
}

// GetPkceS256RequiredOk returns a tuple with the PkceS256Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetPkceS256RequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.PkceS256Required) {
		return nil, false
	}
	return o.PkceS256Required, true
}

// HasPkceS256Required returns a boolean if a field has been set.
func (o *Service) HasPkceS256Required() bool {
	if o != nil && !IsNil(o.PkceS256Required) {
		return true
	}

	return false
}

// SetPkceS256Required gets a reference to the given bool and assigns it to the PkceS256Required field.
func (o *Service) SetPkceS256Required(v bool) {
	o.PkceS256Required = &v
}

// GetAuthorizationResponseDuration returns the AuthorizationResponseDuration field value if set, zero value otherwise.
func (o *Service) GetAuthorizationResponseDuration() int64 {
	if o == nil || IsNil(o.AuthorizationResponseDuration) {
		var ret int64
		return ret
	}
	return *o.AuthorizationResponseDuration
}

// GetAuthorizationResponseDurationOk returns a tuple with the AuthorizationResponseDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAuthorizationResponseDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.AuthorizationResponseDuration) {
		return nil, false
	}
	return o.AuthorizationResponseDuration, true
}

// HasAuthorizationResponseDuration returns a boolean if a field has been set.
func (o *Service) HasAuthorizationResponseDuration() bool {
	if o != nil && !IsNil(o.AuthorizationResponseDuration) {
		return true
	}

	return false
}

// SetAuthorizationResponseDuration gets a reference to the given int64 and assigns it to the AuthorizationResponseDuration field.
func (o *Service) SetAuthorizationResponseDuration(v int64) {
	o.AuthorizationResponseDuration = &v
}

// GetTokenEndpoint returns the TokenEndpoint field value if set, zero value otherwise.
func (o *Service) GetTokenEndpoint() string {
	if o == nil || IsNil(o.TokenEndpoint) {
		var ret string
		return ret
	}
	return *o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetTokenEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.TokenEndpoint) {
		return nil, false
	}
	return o.TokenEndpoint, true
}

// HasTokenEndpoint returns a boolean if a field has been set.
func (o *Service) HasTokenEndpoint() bool {
	if o != nil && !IsNil(o.TokenEndpoint) {
		return true
	}

	return false
}

// SetTokenEndpoint gets a reference to the given string and assigns it to the TokenEndpoint field.
func (o *Service) SetTokenEndpoint(v string) {
	o.TokenEndpoint = &v
}

// GetDirectTokenEndpointEnabled returns the DirectTokenEndpointEnabled field value if set, zero value otherwise.
func (o *Service) GetDirectTokenEndpointEnabled() bool {
	if o == nil || IsNil(o.DirectTokenEndpointEnabled) {
		var ret bool
		return ret
	}
	return *o.DirectTokenEndpointEnabled
}

// GetDirectTokenEndpointEnabledOk returns a tuple with the DirectTokenEndpointEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDirectTokenEndpointEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DirectTokenEndpointEnabled) {
		return nil, false
	}
	return o.DirectTokenEndpointEnabled, true
}

// HasDirectTokenEndpointEnabled returns a boolean if a field has been set.
func (o *Service) HasDirectTokenEndpointEnabled() bool {
	if o != nil && !IsNil(o.DirectTokenEndpointEnabled) {
		return true
	}

	return false
}

// SetDirectTokenEndpointEnabled gets a reference to the given bool and assigns it to the DirectTokenEndpointEnabled field.
func (o *Service) SetDirectTokenEndpointEnabled(v bool) {
	o.DirectTokenEndpointEnabled = &v
}

// GetSupportedTokenAuthMethods returns the SupportedTokenAuthMethods field value if set, zero value otherwise.
func (o *Service) GetSupportedTokenAuthMethods() []ClientAuthMethod {
	if o == nil || IsNil(o.SupportedTokenAuthMethods) {
		var ret []ClientAuthMethod
		return ret
	}
	return o.SupportedTokenAuthMethods
}

// GetSupportedTokenAuthMethodsOk returns a tuple with the SupportedTokenAuthMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedTokenAuthMethodsOk() ([]ClientAuthMethod, bool) {
	if o == nil || IsNil(o.SupportedTokenAuthMethods) {
		return nil, false
	}
	return o.SupportedTokenAuthMethods, true
}

// HasSupportedTokenAuthMethods returns a boolean if a field has been set.
func (o *Service) HasSupportedTokenAuthMethods() bool {
	if o != nil && !IsNil(o.SupportedTokenAuthMethods) {
		return true
	}

	return false
}

// SetSupportedTokenAuthMethods gets a reference to the given []ClientAuthMethod and assigns it to the SupportedTokenAuthMethods field.
func (o *Service) SetSupportedTokenAuthMethods(v []ClientAuthMethod) {
	o.SupportedTokenAuthMethods = v
}

// GetMissingClientIdAllowed returns the MissingClientIdAllowed field value if set, zero value otherwise.
func (o *Service) GetMissingClientIdAllowed() bool {
	if o == nil || IsNil(o.MissingClientIdAllowed) {
		var ret bool
		return ret
	}
	return *o.MissingClientIdAllowed
}

// GetMissingClientIdAllowedOk returns a tuple with the MissingClientIdAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetMissingClientIdAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.MissingClientIdAllowed) {
		return nil, false
	}
	return o.MissingClientIdAllowed, true
}

// HasMissingClientIdAllowed returns a boolean if a field has been set.
func (o *Service) HasMissingClientIdAllowed() bool {
	if o != nil && !IsNil(o.MissingClientIdAllowed) {
		return true
	}

	return false
}

// SetMissingClientIdAllowed gets a reference to the given bool and assigns it to the MissingClientIdAllowed field.
func (o *Service) SetMissingClientIdAllowed(v bool) {
	o.MissingClientIdAllowed = &v
}

// GetRevocationEndpoint returns the RevocationEndpoint field value if set, zero value otherwise.
func (o *Service) GetRevocationEndpoint() string {
	if o == nil || IsNil(o.RevocationEndpoint) {
		var ret string
		return ret
	}
	return *o.RevocationEndpoint
}

// GetRevocationEndpointOk returns a tuple with the RevocationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetRevocationEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.RevocationEndpoint) {
		return nil, false
	}
	return o.RevocationEndpoint, true
}

// HasRevocationEndpoint returns a boolean if a field has been set.
func (o *Service) HasRevocationEndpoint() bool {
	if o != nil && !IsNil(o.RevocationEndpoint) {
		return true
	}

	return false
}

// SetRevocationEndpoint gets a reference to the given string and assigns it to the RevocationEndpoint field.
func (o *Service) SetRevocationEndpoint(v string) {
	o.RevocationEndpoint = &v
}

// GetDirectRevocationEndpointEnabled returns the DirectRevocationEndpointEnabled field value if set, zero value otherwise.
func (o *Service) GetDirectRevocationEndpointEnabled() bool {
	if o == nil || IsNil(o.DirectRevocationEndpointEnabled) {
		var ret bool
		return ret
	}
	return *o.DirectRevocationEndpointEnabled
}

// GetDirectRevocationEndpointEnabledOk returns a tuple with the DirectRevocationEndpointEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDirectRevocationEndpointEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DirectRevocationEndpointEnabled) {
		return nil, false
	}
	return o.DirectRevocationEndpointEnabled, true
}

// HasDirectRevocationEndpointEnabled returns a boolean if a field has been set.
func (o *Service) HasDirectRevocationEndpointEnabled() bool {
	if o != nil && !IsNil(o.DirectRevocationEndpointEnabled) {
		return true
	}

	return false
}

// SetDirectRevocationEndpointEnabled gets a reference to the given bool and assigns it to the DirectRevocationEndpointEnabled field.
func (o *Service) SetDirectRevocationEndpointEnabled(v bool) {
	o.DirectRevocationEndpointEnabled = &v
}

// GetSupportedRevocationAuthMethods returns the SupportedRevocationAuthMethods field value if set, zero value otherwise.
func (o *Service) GetSupportedRevocationAuthMethods() []ClientAuthMethod {
	if o == nil || IsNil(o.SupportedRevocationAuthMethods) {
		var ret []ClientAuthMethod
		return ret
	}
	return o.SupportedRevocationAuthMethods
}

// GetSupportedRevocationAuthMethodsOk returns a tuple with the SupportedRevocationAuthMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedRevocationAuthMethodsOk() ([]ClientAuthMethod, bool) {
	if o == nil || IsNil(o.SupportedRevocationAuthMethods) {
		return nil, false
	}
	return o.SupportedRevocationAuthMethods, true
}

// HasSupportedRevocationAuthMethods returns a boolean if a field has been set.
func (o *Service) HasSupportedRevocationAuthMethods() bool {
	if o != nil && !IsNil(o.SupportedRevocationAuthMethods) {
		return true
	}

	return false
}

// SetSupportedRevocationAuthMethods gets a reference to the given []ClientAuthMethod and assigns it to the SupportedRevocationAuthMethods field.
func (o *Service) SetSupportedRevocationAuthMethods(v []ClientAuthMethod) {
	o.SupportedRevocationAuthMethods = v
}

// GetIntrospectionEndpoint returns the IntrospectionEndpoint field value if set, zero value otherwise.
func (o *Service) GetIntrospectionEndpoint() string {
	if o == nil || IsNil(o.IntrospectionEndpoint) {
		var ret string
		return ret
	}
	return *o.IntrospectionEndpoint
}

// GetIntrospectionEndpointOk returns a tuple with the IntrospectionEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetIntrospectionEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.IntrospectionEndpoint) {
		return nil, false
	}
	return o.IntrospectionEndpoint, true
}

// HasIntrospectionEndpoint returns a boolean if a field has been set.
func (o *Service) HasIntrospectionEndpoint() bool {
	if o != nil && !IsNil(o.IntrospectionEndpoint) {
		return true
	}

	return false
}

// SetIntrospectionEndpoint gets a reference to the given string and assigns it to the IntrospectionEndpoint field.
func (o *Service) SetIntrospectionEndpoint(v string) {
	o.IntrospectionEndpoint = &v
}

// GetDirectIntrospectionEndpointEnabled returns the DirectIntrospectionEndpointEnabled field value if set, zero value otherwise.
func (o *Service) GetDirectIntrospectionEndpointEnabled() bool {
	if o == nil || IsNil(o.DirectIntrospectionEndpointEnabled) {
		var ret bool
		return ret
	}
	return *o.DirectIntrospectionEndpointEnabled
}

// GetDirectIntrospectionEndpointEnabledOk returns a tuple with the DirectIntrospectionEndpointEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDirectIntrospectionEndpointEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DirectIntrospectionEndpointEnabled) {
		return nil, false
	}
	return o.DirectIntrospectionEndpointEnabled, true
}

// HasDirectIntrospectionEndpointEnabled returns a boolean if a field has been set.
func (o *Service) HasDirectIntrospectionEndpointEnabled() bool {
	if o != nil && !IsNil(o.DirectIntrospectionEndpointEnabled) {
		return true
	}

	return false
}

// SetDirectIntrospectionEndpointEnabled gets a reference to the given bool and assigns it to the DirectIntrospectionEndpointEnabled field.
func (o *Service) SetDirectIntrospectionEndpointEnabled(v bool) {
	o.DirectIntrospectionEndpointEnabled = &v
}

// GetSupportedIntrospectionAuthMethods returns the SupportedIntrospectionAuthMethods field value if set, zero value otherwise.
func (o *Service) GetSupportedIntrospectionAuthMethods() []ClientAuthMethod {
	if o == nil || IsNil(o.SupportedIntrospectionAuthMethods) {
		var ret []ClientAuthMethod
		return ret
	}
	return o.SupportedIntrospectionAuthMethods
}

// GetSupportedIntrospectionAuthMethodsOk returns a tuple with the SupportedIntrospectionAuthMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedIntrospectionAuthMethodsOk() ([]ClientAuthMethod, bool) {
	if o == nil || IsNil(o.SupportedIntrospectionAuthMethods) {
		return nil, false
	}
	return o.SupportedIntrospectionAuthMethods, true
}

// HasSupportedIntrospectionAuthMethods returns a boolean if a field has been set.
func (o *Service) HasSupportedIntrospectionAuthMethods() bool {
	if o != nil && !IsNil(o.SupportedIntrospectionAuthMethods) {
		return true
	}

	return false
}

// SetSupportedIntrospectionAuthMethods gets a reference to the given []ClientAuthMethod and assigns it to the SupportedIntrospectionAuthMethods field.
func (o *Service) SetSupportedIntrospectionAuthMethods(v []ClientAuthMethod) {
	o.SupportedIntrospectionAuthMethods = v
}

// GetPushedAuthReqEndpoint returns the PushedAuthReqEndpoint field value if set, zero value otherwise.
func (o *Service) GetPushedAuthReqEndpoint() string {
	if o == nil || IsNil(o.PushedAuthReqEndpoint) {
		var ret string
		return ret
	}
	return *o.PushedAuthReqEndpoint
}

// GetPushedAuthReqEndpointOk returns a tuple with the PushedAuthReqEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetPushedAuthReqEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.PushedAuthReqEndpoint) {
		return nil, false
	}
	return o.PushedAuthReqEndpoint, true
}

// HasPushedAuthReqEndpoint returns a boolean if a field has been set.
func (o *Service) HasPushedAuthReqEndpoint() bool {
	if o != nil && !IsNil(o.PushedAuthReqEndpoint) {
		return true
	}

	return false
}

// SetPushedAuthReqEndpoint gets a reference to the given string and assigns it to the PushedAuthReqEndpoint field.
func (o *Service) SetPushedAuthReqEndpoint(v string) {
	o.PushedAuthReqEndpoint = &v
}

// GetPushedAuthReqDuration returns the PushedAuthReqDuration field value if set, zero value otherwise.
func (o *Service) GetPushedAuthReqDuration() int64 {
	if o == nil || IsNil(o.PushedAuthReqDuration) {
		var ret int64
		return ret
	}
	return *o.PushedAuthReqDuration
}

// GetPushedAuthReqDurationOk returns a tuple with the PushedAuthReqDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetPushedAuthReqDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.PushedAuthReqDuration) {
		return nil, false
	}
	return o.PushedAuthReqDuration, true
}

// HasPushedAuthReqDuration returns a boolean if a field has been set.
func (o *Service) HasPushedAuthReqDuration() bool {
	if o != nil && !IsNil(o.PushedAuthReqDuration) {
		return true
	}

	return false
}

// SetPushedAuthReqDuration gets a reference to the given int64 and assigns it to the PushedAuthReqDuration field.
func (o *Service) SetPushedAuthReqDuration(v int64) {
	o.PushedAuthReqDuration = &v
}

// GetParRequired returns the ParRequired field value if set, zero value otherwise.
func (o *Service) GetParRequired() bool {
	if o == nil || IsNil(o.ParRequired) {
		var ret bool
		return ret
	}
	return *o.ParRequired
}

// GetParRequiredOk returns a tuple with the ParRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetParRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ParRequired) {
		return nil, false
	}
	return o.ParRequired, true
}

// HasParRequired returns a boolean if a field has been set.
func (o *Service) HasParRequired() bool {
	if o != nil && !IsNil(o.ParRequired) {
		return true
	}

	return false
}

// SetParRequired gets a reference to the given bool and assigns it to the ParRequired field.
func (o *Service) SetParRequired(v bool) {
	o.ParRequired = &v
}

// GetRequestObjectRequired returns the RequestObjectRequired field value if set, zero value otherwise.
func (o *Service) GetRequestObjectRequired() bool {
	if o == nil || IsNil(o.RequestObjectRequired) {
		var ret bool
		return ret
	}
	return *o.RequestObjectRequired
}

// GetRequestObjectRequiredOk returns a tuple with the RequestObjectRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetRequestObjectRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestObjectRequired) {
		return nil, false
	}
	return o.RequestObjectRequired, true
}

// HasRequestObjectRequired returns a boolean if a field has been set.
func (o *Service) HasRequestObjectRequired() bool {
	if o != nil && !IsNil(o.RequestObjectRequired) {
		return true
	}

	return false
}

// SetRequestObjectRequired gets a reference to the given bool and assigns it to the RequestObjectRequired field.
func (o *Service) SetRequestObjectRequired(v bool) {
	o.RequestObjectRequired = &v
}

// GetTraditionalRequestObjectProcessingApplied returns the TraditionalRequestObjectProcessingApplied field value if set, zero value otherwise.
func (o *Service) GetTraditionalRequestObjectProcessingApplied() bool {
	if o == nil || IsNil(o.TraditionalRequestObjectProcessingApplied) {
		var ret bool
		return ret
	}
	return *o.TraditionalRequestObjectProcessingApplied
}

// GetTraditionalRequestObjectProcessingAppliedOk returns a tuple with the TraditionalRequestObjectProcessingApplied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetTraditionalRequestObjectProcessingAppliedOk() (*bool, bool) {
	if o == nil || IsNil(o.TraditionalRequestObjectProcessingApplied) {
		return nil, false
	}
	return o.TraditionalRequestObjectProcessingApplied, true
}

// HasTraditionalRequestObjectProcessingApplied returns a boolean if a field has been set.
func (o *Service) HasTraditionalRequestObjectProcessingApplied() bool {
	if o != nil && !IsNil(o.TraditionalRequestObjectProcessingApplied) {
		return true
	}

	return false
}

// SetTraditionalRequestObjectProcessingApplied gets a reference to the given bool and assigns it to the TraditionalRequestObjectProcessingApplied field.
func (o *Service) SetTraditionalRequestObjectProcessingApplied(v bool) {
	o.TraditionalRequestObjectProcessingApplied = &v
}

// GetMutualTlsValidatePkiCertChain returns the MutualTlsValidatePkiCertChain field value if set, zero value otherwise.
func (o *Service) GetMutualTlsValidatePkiCertChain() bool {
	if o == nil || IsNil(o.MutualTlsValidatePkiCertChain) {
		var ret bool
		return ret
	}
	return *o.MutualTlsValidatePkiCertChain
}

// GetMutualTlsValidatePkiCertChainOk returns a tuple with the MutualTlsValidatePkiCertChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetMutualTlsValidatePkiCertChainOk() (*bool, bool) {
	if o == nil || IsNil(o.MutualTlsValidatePkiCertChain) {
		return nil, false
	}
	return o.MutualTlsValidatePkiCertChain, true
}

// HasMutualTlsValidatePkiCertChain returns a boolean if a field has been set.
func (o *Service) HasMutualTlsValidatePkiCertChain() bool {
	if o != nil && !IsNil(o.MutualTlsValidatePkiCertChain) {
		return true
	}

	return false
}

// SetMutualTlsValidatePkiCertChain gets a reference to the given bool and assigns it to the MutualTlsValidatePkiCertChain field.
func (o *Service) SetMutualTlsValidatePkiCertChain(v bool) {
	o.MutualTlsValidatePkiCertChain = &v
}

// GetTrustedRootCertificates returns the TrustedRootCertificates field value if set, zero value otherwise.
func (o *Service) GetTrustedRootCertificates() []string {
	if o == nil || IsNil(o.TrustedRootCertificates) {
		var ret []string
		return ret
	}
	return o.TrustedRootCertificates
}

// GetTrustedRootCertificatesOk returns a tuple with the TrustedRootCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetTrustedRootCertificatesOk() ([]string, bool) {
	if o == nil || IsNil(o.TrustedRootCertificates) {
		return nil, false
	}
	return o.TrustedRootCertificates, true
}

// HasTrustedRootCertificates returns a boolean if a field has been set.
func (o *Service) HasTrustedRootCertificates() bool {
	if o != nil && !IsNil(o.TrustedRootCertificates) {
		return true
	}

	return false
}

// SetTrustedRootCertificates gets a reference to the given []string and assigns it to the TrustedRootCertificates field.
func (o *Service) SetTrustedRootCertificates(v []string) {
	o.TrustedRootCertificates = v
}

// GetMtlsEndpointAliases returns the MtlsEndpointAliases field value if set, zero value otherwise.
func (o *Service) GetMtlsEndpointAliases() []NamedUri {
	if o == nil || IsNil(o.MtlsEndpointAliases) {
		var ret []NamedUri
		return ret
	}
	return o.MtlsEndpointAliases
}

// GetMtlsEndpointAliasesOk returns a tuple with the MtlsEndpointAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetMtlsEndpointAliasesOk() ([]NamedUri, bool) {
	if o == nil || IsNil(o.MtlsEndpointAliases) {
		return nil, false
	}
	return o.MtlsEndpointAliases, true
}

// HasMtlsEndpointAliases returns a boolean if a field has been set.
func (o *Service) HasMtlsEndpointAliases() bool {
	if o != nil && !IsNil(o.MtlsEndpointAliases) {
		return true
	}

	return false
}

// SetMtlsEndpointAliases gets a reference to the given []NamedUri and assigns it to the MtlsEndpointAliases field.
func (o *Service) SetMtlsEndpointAliases(v []NamedUri) {
	o.MtlsEndpointAliases = v
}

// GetAccessTokenType returns the AccessTokenType field value if set, zero value otherwise.
func (o *Service) GetAccessTokenType() string {
	if o == nil || IsNil(o.AccessTokenType) {
		var ret string
		return ret
	}
	return *o.AccessTokenType
}

// GetAccessTokenTypeOk returns a tuple with the AccessTokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAccessTokenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessTokenType) {
		return nil, false
	}
	return o.AccessTokenType, true
}

// HasAccessTokenType returns a boolean if a field has been set.
func (o *Service) HasAccessTokenType() bool {
	if o != nil && !IsNil(o.AccessTokenType) {
		return true
	}

	return false
}

// SetAccessTokenType gets a reference to the given string and assigns it to the AccessTokenType field.
func (o *Service) SetAccessTokenType(v string) {
	o.AccessTokenType = &v
}

// GetTlsClientCertificateBoundAccessTokens returns the TlsClientCertificateBoundAccessTokens field value if set, zero value otherwise.
func (o *Service) GetTlsClientCertificateBoundAccessTokens() bool {
	if o == nil || IsNil(o.TlsClientCertificateBoundAccessTokens) {
		var ret bool
		return ret
	}
	return *o.TlsClientCertificateBoundAccessTokens
}

// GetTlsClientCertificateBoundAccessTokensOk returns a tuple with the TlsClientCertificateBoundAccessTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetTlsClientCertificateBoundAccessTokensOk() (*bool, bool) {
	if o == nil || IsNil(o.TlsClientCertificateBoundAccessTokens) {
		return nil, false
	}
	return o.TlsClientCertificateBoundAccessTokens, true
}

// HasTlsClientCertificateBoundAccessTokens returns a boolean if a field has been set.
func (o *Service) HasTlsClientCertificateBoundAccessTokens() bool {
	if o != nil && !IsNil(o.TlsClientCertificateBoundAccessTokens) {
		return true
	}

	return false
}

// SetTlsClientCertificateBoundAccessTokens gets a reference to the given bool and assigns it to the TlsClientCertificateBoundAccessTokens field.
func (o *Service) SetTlsClientCertificateBoundAccessTokens(v bool) {
	o.TlsClientCertificateBoundAccessTokens = &v
}

// GetAccessTokenDuration returns the AccessTokenDuration field value if set, zero value otherwise.
func (o *Service) GetAccessTokenDuration() int64 {
	if o == nil || IsNil(o.AccessTokenDuration) {
		var ret int64
		return ret
	}
	return *o.AccessTokenDuration
}

// GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAccessTokenDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.AccessTokenDuration) {
		return nil, false
	}
	return o.AccessTokenDuration, true
}

// HasAccessTokenDuration returns a boolean if a field has been set.
func (o *Service) HasAccessTokenDuration() bool {
	if o != nil && !IsNil(o.AccessTokenDuration) {
		return true
	}

	return false
}

// SetAccessTokenDuration gets a reference to the given int64 and assigns it to the AccessTokenDuration field.
func (o *Service) SetAccessTokenDuration(v int64) {
	o.AccessTokenDuration = &v
}

// GetSingleAccessTokenPerSubject returns the SingleAccessTokenPerSubject field value if set, zero value otherwise.
func (o *Service) GetSingleAccessTokenPerSubject() bool {
	if o == nil || IsNil(o.SingleAccessTokenPerSubject) {
		var ret bool
		return ret
	}
	return *o.SingleAccessTokenPerSubject
}

// GetSingleAccessTokenPerSubjectOk returns a tuple with the SingleAccessTokenPerSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSingleAccessTokenPerSubjectOk() (*bool, bool) {
	if o == nil || IsNil(o.SingleAccessTokenPerSubject) {
		return nil, false
	}
	return o.SingleAccessTokenPerSubject, true
}

// HasSingleAccessTokenPerSubject returns a boolean if a field has been set.
func (o *Service) HasSingleAccessTokenPerSubject() bool {
	if o != nil && !IsNil(o.SingleAccessTokenPerSubject) {
		return true
	}

	return false
}

// SetSingleAccessTokenPerSubject gets a reference to the given bool and assigns it to the SingleAccessTokenPerSubject field.
func (o *Service) SetSingleAccessTokenPerSubject(v bool) {
	o.SingleAccessTokenPerSubject = &v
}

// GetAccessTokenSignAlg returns the AccessTokenSignAlg field value if set, zero value otherwise.
func (o *Service) GetAccessTokenSignAlg() JwsAlg {
	if o == nil || IsNil(o.AccessTokenSignAlg) {
		var ret JwsAlg
		return ret
	}
	return *o.AccessTokenSignAlg
}

// GetAccessTokenSignAlgOk returns a tuple with the AccessTokenSignAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAccessTokenSignAlgOk() (*JwsAlg, bool) {
	if o == nil || IsNil(o.AccessTokenSignAlg) {
		return nil, false
	}
	return o.AccessTokenSignAlg, true
}

// HasAccessTokenSignAlg returns a boolean if a field has been set.
func (o *Service) HasAccessTokenSignAlg() bool {
	if o != nil && !IsNil(o.AccessTokenSignAlg) {
		return true
	}

	return false
}

// SetAccessTokenSignAlg gets a reference to the given JwsAlg and assigns it to the AccessTokenSignAlg field.
func (o *Service) SetAccessTokenSignAlg(v JwsAlg) {
	o.AccessTokenSignAlg = &v
}

// GetAccessTokenSignatureKeyId returns the AccessTokenSignatureKeyId field value if set, zero value otherwise.
func (o *Service) GetAccessTokenSignatureKeyId() string {
	if o == nil || IsNil(o.AccessTokenSignatureKeyId) {
		var ret string
		return ret
	}
	return *o.AccessTokenSignatureKeyId
}

// GetAccessTokenSignatureKeyIdOk returns a tuple with the AccessTokenSignatureKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAccessTokenSignatureKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessTokenSignatureKeyId) {
		return nil, false
	}
	return o.AccessTokenSignatureKeyId, true
}

// HasAccessTokenSignatureKeyId returns a boolean if a field has been set.
func (o *Service) HasAccessTokenSignatureKeyId() bool {
	if o != nil && !IsNil(o.AccessTokenSignatureKeyId) {
		return true
	}

	return false
}

// SetAccessTokenSignatureKeyId gets a reference to the given string and assigns it to the AccessTokenSignatureKeyId field.
func (o *Service) SetAccessTokenSignatureKeyId(v string) {
	o.AccessTokenSignatureKeyId = &v
}

// GetRefreshTokenDuration returns the RefreshTokenDuration field value if set, zero value otherwise.
func (o *Service) GetRefreshTokenDuration() int64 {
	if o == nil || IsNil(o.RefreshTokenDuration) {
		var ret int64
		return ret
	}
	return *o.RefreshTokenDuration
}

// GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetRefreshTokenDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.RefreshTokenDuration) {
		return nil, false
	}
	return o.RefreshTokenDuration, true
}

// HasRefreshTokenDuration returns a boolean if a field has been set.
func (o *Service) HasRefreshTokenDuration() bool {
	if o != nil && !IsNil(o.RefreshTokenDuration) {
		return true
	}

	return false
}

// SetRefreshTokenDuration gets a reference to the given int64 and assigns it to the RefreshTokenDuration field.
func (o *Service) SetRefreshTokenDuration(v int64) {
	o.RefreshTokenDuration = &v
}

// GetRefreshTokenDurationKept returns the RefreshTokenDurationKept field value if set, zero value otherwise.
func (o *Service) GetRefreshTokenDurationKept() bool {
	if o == nil || IsNil(o.RefreshTokenDurationKept) {
		var ret bool
		return ret
	}
	return *o.RefreshTokenDurationKept
}

// GetRefreshTokenDurationKeptOk returns a tuple with the RefreshTokenDurationKept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetRefreshTokenDurationKeptOk() (*bool, bool) {
	if o == nil || IsNil(o.RefreshTokenDurationKept) {
		return nil, false
	}
	return o.RefreshTokenDurationKept, true
}

// HasRefreshTokenDurationKept returns a boolean if a field has been set.
func (o *Service) HasRefreshTokenDurationKept() bool {
	if o != nil && !IsNil(o.RefreshTokenDurationKept) {
		return true
	}

	return false
}

// SetRefreshTokenDurationKept gets a reference to the given bool and assigns it to the RefreshTokenDurationKept field.
func (o *Service) SetRefreshTokenDurationKept(v bool) {
	o.RefreshTokenDurationKept = &v
}

// GetRefreshTokenDurationReset returns the RefreshTokenDurationReset field value if set, zero value otherwise.
func (o *Service) GetRefreshTokenDurationReset() bool {
	if o == nil || IsNil(o.RefreshTokenDurationReset) {
		var ret bool
		return ret
	}
	return *o.RefreshTokenDurationReset
}

// GetRefreshTokenDurationResetOk returns a tuple with the RefreshTokenDurationReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetRefreshTokenDurationResetOk() (*bool, bool) {
	if o == nil || IsNil(o.RefreshTokenDurationReset) {
		return nil, false
	}
	return o.RefreshTokenDurationReset, true
}

// HasRefreshTokenDurationReset returns a boolean if a field has been set.
func (o *Service) HasRefreshTokenDurationReset() bool {
	if o != nil && !IsNil(o.RefreshTokenDurationReset) {
		return true
	}

	return false
}

// SetRefreshTokenDurationReset gets a reference to the given bool and assigns it to the RefreshTokenDurationReset field.
func (o *Service) SetRefreshTokenDurationReset(v bool) {
	o.RefreshTokenDurationReset = &v
}

// GetRefreshTokenKept returns the RefreshTokenKept field value if set, zero value otherwise.
func (o *Service) GetRefreshTokenKept() bool {
	if o == nil || IsNil(o.RefreshTokenKept) {
		var ret bool
		return ret
	}
	return *o.RefreshTokenKept
}

// GetRefreshTokenKeptOk returns a tuple with the RefreshTokenKept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetRefreshTokenKeptOk() (*bool, bool) {
	if o == nil || IsNil(o.RefreshTokenKept) {
		return nil, false
	}
	return o.RefreshTokenKept, true
}

// HasRefreshTokenKept returns a boolean if a field has been set.
func (o *Service) HasRefreshTokenKept() bool {
	if o != nil && !IsNil(o.RefreshTokenKept) {
		return true
	}

	return false
}

// SetRefreshTokenKept gets a reference to the given bool and assigns it to the RefreshTokenKept field.
func (o *Service) SetRefreshTokenKept(v bool) {
	o.RefreshTokenKept = &v
}

// GetSupportedScopes returns the SupportedScopes field value if set, zero value otherwise.
func (o *Service) GetSupportedScopes() []Scope {
	if o == nil || IsNil(o.SupportedScopes) {
		var ret []Scope
		return ret
	}
	return o.SupportedScopes
}

// GetSupportedScopesOk returns a tuple with the SupportedScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedScopesOk() ([]Scope, bool) {
	if o == nil || IsNil(o.SupportedScopes) {
		return nil, false
	}
	return o.SupportedScopes, true
}

// HasSupportedScopes returns a boolean if a field has been set.
func (o *Service) HasSupportedScopes() bool {
	if o != nil && !IsNil(o.SupportedScopes) {
		return true
	}

	return false
}

// SetSupportedScopes gets a reference to the given []Scope and assigns it to the SupportedScopes field.
func (o *Service) SetSupportedScopes(v []Scope) {
	o.SupportedScopes = v
}

// GetScopeRequired returns the ScopeRequired field value if set, zero value otherwise.
func (o *Service) GetScopeRequired() bool {
	if o == nil || IsNil(o.ScopeRequired) {
		var ret bool
		return ret
	}
	return *o.ScopeRequired
}

// GetScopeRequiredOk returns a tuple with the ScopeRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetScopeRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ScopeRequired) {
		return nil, false
	}
	return o.ScopeRequired, true
}

// HasScopeRequired returns a boolean if a field has been set.
func (o *Service) HasScopeRequired() bool {
	if o != nil && !IsNil(o.ScopeRequired) {
		return true
	}

	return false
}

// SetScopeRequired gets a reference to the given bool and assigns it to the ScopeRequired field.
func (o *Service) SetScopeRequired(v bool) {
	o.ScopeRequired = &v
}

// GetIdTokenDuration returns the IdTokenDuration field value if set, zero value otherwise.
func (o *Service) GetIdTokenDuration() int64 {
	if o == nil || IsNil(o.IdTokenDuration) {
		var ret int64
		return ret
	}
	return *o.IdTokenDuration
}

// GetIdTokenDurationOk returns a tuple with the IdTokenDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetIdTokenDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.IdTokenDuration) {
		return nil, false
	}
	return o.IdTokenDuration, true
}

// HasIdTokenDuration returns a boolean if a field has been set.
func (o *Service) HasIdTokenDuration() bool {
	if o != nil && !IsNil(o.IdTokenDuration) {
		return true
	}

	return false
}

// SetIdTokenDuration gets a reference to the given int64 and assigns it to the IdTokenDuration field.
func (o *Service) SetIdTokenDuration(v int64) {
	o.IdTokenDuration = &v
}

// GetAllowableClockSkew returns the AllowableClockSkew field value if set, zero value otherwise.
func (o *Service) GetAllowableClockSkew() int32 {
	if o == nil || IsNil(o.AllowableClockSkew) {
		var ret int32
		return ret
	}
	return *o.AllowableClockSkew
}

// GetAllowableClockSkewOk returns a tuple with the AllowableClockSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAllowableClockSkewOk() (*int32, bool) {
	if o == nil || IsNil(o.AllowableClockSkew) {
		return nil, false
	}
	return o.AllowableClockSkew, true
}

// HasAllowableClockSkew returns a boolean if a field has been set.
func (o *Service) HasAllowableClockSkew() bool {
	if o != nil && !IsNil(o.AllowableClockSkew) {
		return true
	}

	return false
}

// SetAllowableClockSkew gets a reference to the given int32 and assigns it to the AllowableClockSkew field.
func (o *Service) SetAllowableClockSkew(v int32) {
	o.AllowableClockSkew = &v
}

// GetSupportedClaimTypes returns the SupportedClaimTypes field value if set, zero value otherwise.
func (o *Service) GetSupportedClaimTypes() []ClaimType {
	if o == nil || IsNil(o.SupportedClaimTypes) {
		var ret []ClaimType
		return ret
	}
	return o.SupportedClaimTypes
}

// GetSupportedClaimTypesOk returns a tuple with the SupportedClaimTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedClaimTypesOk() ([]ClaimType, bool) {
	if o == nil || IsNil(o.SupportedClaimTypes) {
		return nil, false
	}
	return o.SupportedClaimTypes, true
}

// HasSupportedClaimTypes returns a boolean if a field has been set.
func (o *Service) HasSupportedClaimTypes() bool {
	if o != nil && !IsNil(o.SupportedClaimTypes) {
		return true
	}

	return false
}

// SetSupportedClaimTypes gets a reference to the given []ClaimType and assigns it to the SupportedClaimTypes field.
func (o *Service) SetSupportedClaimTypes(v []ClaimType) {
	o.SupportedClaimTypes = v
}

// GetSupportedClaimLocales returns the SupportedClaimLocales field value if set, zero value otherwise.
func (o *Service) GetSupportedClaimLocales() []string {
	if o == nil || IsNil(o.SupportedClaimLocales) {
		var ret []string
		return ret
	}
	return o.SupportedClaimLocales
}

// GetSupportedClaimLocalesOk returns a tuple with the SupportedClaimLocales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedClaimLocalesOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedClaimLocales) {
		return nil, false
	}
	return o.SupportedClaimLocales, true
}

// HasSupportedClaimLocales returns a boolean if a field has been set.
func (o *Service) HasSupportedClaimLocales() bool {
	if o != nil && !IsNil(o.SupportedClaimLocales) {
		return true
	}

	return false
}

// SetSupportedClaimLocales gets a reference to the given []string and assigns it to the SupportedClaimLocales field.
func (o *Service) SetSupportedClaimLocales(v []string) {
	o.SupportedClaimLocales = v
}

// GetSupportedClaims returns the SupportedClaims field value if set, zero value otherwise.
func (o *Service) GetSupportedClaims() []string {
	if o == nil || IsNil(o.SupportedClaims) {
		var ret []string
		return ret
	}
	return o.SupportedClaims
}

// GetSupportedClaimsOk returns a tuple with the SupportedClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedClaimsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedClaims) {
		return nil, false
	}
	return o.SupportedClaims, true
}

// HasSupportedClaims returns a boolean if a field has been set.
func (o *Service) HasSupportedClaims() bool {
	if o != nil && !IsNil(o.SupportedClaims) {
		return true
	}

	return false
}

// SetSupportedClaims gets a reference to the given []string and assigns it to the SupportedClaims field.
func (o *Service) SetSupportedClaims(v []string) {
	o.SupportedClaims = v
}

// GetClaimShortcutRestrictive returns the ClaimShortcutRestrictive field value if set, zero value otherwise.
func (o *Service) GetClaimShortcutRestrictive() bool {
	if o == nil || IsNil(o.ClaimShortcutRestrictive) {
		var ret bool
		return ret
	}
	return *o.ClaimShortcutRestrictive
}

// GetClaimShortcutRestrictiveOk returns a tuple with the ClaimShortcutRestrictive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetClaimShortcutRestrictiveOk() (*bool, bool) {
	if o == nil || IsNil(o.ClaimShortcutRestrictive) {
		return nil, false
	}
	return o.ClaimShortcutRestrictive, true
}

// HasClaimShortcutRestrictive returns a boolean if a field has been set.
func (o *Service) HasClaimShortcutRestrictive() bool {
	if o != nil && !IsNil(o.ClaimShortcutRestrictive) {
		return true
	}

	return false
}

// SetClaimShortcutRestrictive gets a reference to the given bool and assigns it to the ClaimShortcutRestrictive field.
func (o *Service) SetClaimShortcutRestrictive(v bool) {
	o.ClaimShortcutRestrictive = &v
}

// GetJwksUri returns the JwksUri field value if set, zero value otherwise.
func (o *Service) GetJwksUri() string {
	if o == nil || IsNil(o.JwksUri) {
		var ret string
		return ret
	}
	return *o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetJwksUriOk() (*string, bool) {
	if o == nil || IsNil(o.JwksUri) {
		return nil, false
	}
	return o.JwksUri, true
}

// HasJwksUri returns a boolean if a field has been set.
func (o *Service) HasJwksUri() bool {
	if o != nil && !IsNil(o.JwksUri) {
		return true
	}

	return false
}

// SetJwksUri gets a reference to the given string and assigns it to the JwksUri field.
func (o *Service) SetJwksUri(v string) {
	o.JwksUri = &v
}

// GetDirectJwksEndpointEnabled returns the DirectJwksEndpointEnabled field value if set, zero value otherwise.
func (o *Service) GetDirectJwksEndpointEnabled() bool {
	if o == nil || IsNil(o.DirectJwksEndpointEnabled) {
		var ret bool
		return ret
	}
	return *o.DirectJwksEndpointEnabled
}

// GetDirectJwksEndpointEnabledOk returns a tuple with the DirectJwksEndpointEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDirectJwksEndpointEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DirectJwksEndpointEnabled) {
		return nil, false
	}
	return o.DirectJwksEndpointEnabled, true
}

// HasDirectJwksEndpointEnabled returns a boolean if a field has been set.
func (o *Service) HasDirectJwksEndpointEnabled() bool {
	if o != nil && !IsNil(o.DirectJwksEndpointEnabled) {
		return true
	}

	return false
}

// SetDirectJwksEndpointEnabled gets a reference to the given bool and assigns it to the DirectJwksEndpointEnabled field.
func (o *Service) SetDirectJwksEndpointEnabled(v bool) {
	o.DirectJwksEndpointEnabled = &v
}

// GetJwks returns the Jwks field value if set, zero value otherwise.
func (o *Service) GetJwks() string {
	if o == nil || IsNil(o.Jwks) {
		var ret string
		return ret
	}
	return *o.Jwks
}

// GetJwksOk returns a tuple with the Jwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetJwksOk() (*string, bool) {
	if o == nil || IsNil(o.Jwks) {
		return nil, false
	}
	return o.Jwks, true
}

// HasJwks returns a boolean if a field has been set.
func (o *Service) HasJwks() bool {
	if o != nil && !IsNil(o.Jwks) {
		return true
	}

	return false
}

// SetJwks gets a reference to the given string and assigns it to the Jwks field.
func (o *Service) SetJwks(v string) {
	o.Jwks = &v
}

// GetIdTokenSignatureKeyId returns the IdTokenSignatureKeyId field value if set, zero value otherwise.
func (o *Service) GetIdTokenSignatureKeyId() string {
	if o == nil || IsNil(o.IdTokenSignatureKeyId) {
		var ret string
		return ret
	}
	return *o.IdTokenSignatureKeyId
}

// GetIdTokenSignatureKeyIdOk returns a tuple with the IdTokenSignatureKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetIdTokenSignatureKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdTokenSignatureKeyId) {
		return nil, false
	}
	return o.IdTokenSignatureKeyId, true
}

// HasIdTokenSignatureKeyId returns a boolean if a field has been set.
func (o *Service) HasIdTokenSignatureKeyId() bool {
	if o != nil && !IsNil(o.IdTokenSignatureKeyId) {
		return true
	}

	return false
}

// SetIdTokenSignatureKeyId gets a reference to the given string and assigns it to the IdTokenSignatureKeyId field.
func (o *Service) SetIdTokenSignatureKeyId(v string) {
	o.IdTokenSignatureKeyId = &v
}

// GetUserInfoSignatureKeyId returns the UserInfoSignatureKeyId field value if set, zero value otherwise.
func (o *Service) GetUserInfoSignatureKeyId() string {
	if o == nil || IsNil(o.UserInfoSignatureKeyId) {
		var ret string
		return ret
	}
	return *o.UserInfoSignatureKeyId
}

// GetUserInfoSignatureKeyIdOk returns a tuple with the UserInfoSignatureKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetUserInfoSignatureKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserInfoSignatureKeyId) {
		return nil, false
	}
	return o.UserInfoSignatureKeyId, true
}

// HasUserInfoSignatureKeyId returns a boolean if a field has been set.
func (o *Service) HasUserInfoSignatureKeyId() bool {
	if o != nil && !IsNil(o.UserInfoSignatureKeyId) {
		return true
	}

	return false
}

// SetUserInfoSignatureKeyId gets a reference to the given string and assigns it to the UserInfoSignatureKeyId field.
func (o *Service) SetUserInfoSignatureKeyId(v string) {
	o.UserInfoSignatureKeyId = &v
}

// GetAuthorizationSignatureKeyId returns the AuthorizationSignatureKeyId field value if set, zero value otherwise.
func (o *Service) GetAuthorizationSignatureKeyId() string {
	if o == nil || IsNil(o.AuthorizationSignatureKeyId) {
		var ret string
		return ret
	}
	return *o.AuthorizationSignatureKeyId
}

// GetAuthorizationSignatureKeyIdOk returns a tuple with the AuthorizationSignatureKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAuthorizationSignatureKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationSignatureKeyId) {
		return nil, false
	}
	return o.AuthorizationSignatureKeyId, true
}

// HasAuthorizationSignatureKeyId returns a boolean if a field has been set.
func (o *Service) HasAuthorizationSignatureKeyId() bool {
	if o != nil && !IsNil(o.AuthorizationSignatureKeyId) {
		return true
	}

	return false
}

// SetAuthorizationSignatureKeyId gets a reference to the given string and assigns it to the AuthorizationSignatureKeyId field.
func (o *Service) SetAuthorizationSignatureKeyId(v string) {
	o.AuthorizationSignatureKeyId = &v
}

// GetUserInfoEndpoint returns the UserInfoEndpoint field value if set, zero value otherwise.
func (o *Service) GetUserInfoEndpoint() string {
	if o == nil || IsNil(o.UserInfoEndpoint) {
		var ret string
		return ret
	}
	return *o.UserInfoEndpoint
}

// GetUserInfoEndpointOk returns a tuple with the UserInfoEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetUserInfoEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.UserInfoEndpoint) {
		return nil, false
	}
	return o.UserInfoEndpoint, true
}

// HasUserInfoEndpoint returns a boolean if a field has been set.
func (o *Service) HasUserInfoEndpoint() bool {
	if o != nil && !IsNil(o.UserInfoEndpoint) {
		return true
	}

	return false
}

// SetUserInfoEndpoint gets a reference to the given string and assigns it to the UserInfoEndpoint field.
func (o *Service) SetUserInfoEndpoint(v string) {
	o.UserInfoEndpoint = &v
}

// GetDirectUserInfoEndpointEnabled returns the DirectUserInfoEndpointEnabled field value if set, zero value otherwise.
func (o *Service) GetDirectUserInfoEndpointEnabled() bool {
	if o == nil || IsNil(o.DirectUserInfoEndpointEnabled) {
		var ret bool
		return ret
	}
	return *o.DirectUserInfoEndpointEnabled
}

// GetDirectUserInfoEndpointEnabledOk returns a tuple with the DirectUserInfoEndpointEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDirectUserInfoEndpointEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DirectUserInfoEndpointEnabled) {
		return nil, false
	}
	return o.DirectUserInfoEndpointEnabled, true
}

// HasDirectUserInfoEndpointEnabled returns a boolean if a field has been set.
func (o *Service) HasDirectUserInfoEndpointEnabled() bool {
	if o != nil && !IsNil(o.DirectUserInfoEndpointEnabled) {
		return true
	}

	return false
}

// SetDirectUserInfoEndpointEnabled gets a reference to the given bool and assigns it to the DirectUserInfoEndpointEnabled field.
func (o *Service) SetDirectUserInfoEndpointEnabled(v bool) {
	o.DirectUserInfoEndpointEnabled = &v
}

// GetDynamicRegistrationSupported returns the DynamicRegistrationSupported field value if set, zero value otherwise.
func (o *Service) GetDynamicRegistrationSupported() bool {
	if o == nil || IsNil(o.DynamicRegistrationSupported) {
		var ret bool
		return ret
	}
	return *o.DynamicRegistrationSupported
}

// GetDynamicRegistrationSupportedOk returns a tuple with the DynamicRegistrationSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDynamicRegistrationSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.DynamicRegistrationSupported) {
		return nil, false
	}
	return o.DynamicRegistrationSupported, true
}

// HasDynamicRegistrationSupported returns a boolean if a field has been set.
func (o *Service) HasDynamicRegistrationSupported() bool {
	if o != nil && !IsNil(o.DynamicRegistrationSupported) {
		return true
	}

	return false
}

// SetDynamicRegistrationSupported gets a reference to the given bool and assigns it to the DynamicRegistrationSupported field.
func (o *Service) SetDynamicRegistrationSupported(v bool) {
	o.DynamicRegistrationSupported = &v
}

// GetRegistrationEndpoint returns the RegistrationEndpoint field value if set, zero value otherwise.
func (o *Service) GetRegistrationEndpoint() string {
	if o == nil || IsNil(o.RegistrationEndpoint) {
		var ret string
		return ret
	}
	return *o.RegistrationEndpoint
}

// GetRegistrationEndpointOk returns a tuple with the RegistrationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetRegistrationEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.RegistrationEndpoint) {
		return nil, false
	}
	return o.RegistrationEndpoint, true
}

// HasRegistrationEndpoint returns a boolean if a field has been set.
func (o *Service) HasRegistrationEndpoint() bool {
	if o != nil && !IsNil(o.RegistrationEndpoint) {
		return true
	}

	return false
}

// SetRegistrationEndpoint gets a reference to the given string and assigns it to the RegistrationEndpoint field.
func (o *Service) SetRegistrationEndpoint(v string) {
	o.RegistrationEndpoint = &v
}

// GetRegistrationManagementEndpoint returns the RegistrationManagementEndpoint field value if set, zero value otherwise.
func (o *Service) GetRegistrationManagementEndpoint() string {
	if o == nil || IsNil(o.RegistrationManagementEndpoint) {
		var ret string
		return ret
	}
	return *o.RegistrationManagementEndpoint
}

// GetRegistrationManagementEndpointOk returns a tuple with the RegistrationManagementEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetRegistrationManagementEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.RegistrationManagementEndpoint) {
		return nil, false
	}
	return o.RegistrationManagementEndpoint, true
}

// HasRegistrationManagementEndpoint returns a boolean if a field has been set.
func (o *Service) HasRegistrationManagementEndpoint() bool {
	if o != nil && !IsNil(o.RegistrationManagementEndpoint) {
		return true
	}

	return false
}

// SetRegistrationManagementEndpoint gets a reference to the given string and assigns it to the RegistrationManagementEndpoint field.
func (o *Service) SetRegistrationManagementEndpoint(v string) {
	o.RegistrationManagementEndpoint = &v
}

// GetPolicyUri returns the PolicyUri field value if set, zero value otherwise.
func (o *Service) GetPolicyUri() string {
	if o == nil || IsNil(o.PolicyUri) {
		var ret string
		return ret
	}
	return *o.PolicyUri
}

// GetPolicyUriOk returns a tuple with the PolicyUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetPolicyUriOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyUri) {
		return nil, false
	}
	return o.PolicyUri, true
}

// HasPolicyUri returns a boolean if a field has been set.
func (o *Service) HasPolicyUri() bool {
	if o != nil && !IsNil(o.PolicyUri) {
		return true
	}

	return false
}

// SetPolicyUri gets a reference to the given string and assigns it to the PolicyUri field.
func (o *Service) SetPolicyUri(v string) {
	o.PolicyUri = &v
}

// GetTosUri returns the TosUri field value if set, zero value otherwise.
func (o *Service) GetTosUri() string {
	if o == nil || IsNil(o.TosUri) {
		var ret string
		return ret
	}
	return *o.TosUri
}

// GetTosUriOk returns a tuple with the TosUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetTosUriOk() (*string, bool) {
	if o == nil || IsNil(o.TosUri) {
		return nil, false
	}
	return o.TosUri, true
}

// HasTosUri returns a boolean if a field has been set.
func (o *Service) HasTosUri() bool {
	if o != nil && !IsNil(o.TosUri) {
		return true
	}

	return false
}

// SetTosUri gets a reference to the given string and assigns it to the TosUri field.
func (o *Service) SetTosUri(v string) {
	o.TosUri = &v
}

// GetServiceDocumentation returns the ServiceDocumentation field value if set, zero value otherwise.
func (o *Service) GetServiceDocumentation() string {
	if o == nil || IsNil(o.ServiceDocumentation) {
		var ret string
		return ret
	}
	return *o.ServiceDocumentation
}

// GetServiceDocumentationOk returns a tuple with the ServiceDocumentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetServiceDocumentationOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceDocumentation) {
		return nil, false
	}
	return o.ServiceDocumentation, true
}

// HasServiceDocumentation returns a boolean if a field has been set.
func (o *Service) HasServiceDocumentation() bool {
	if o != nil && !IsNil(o.ServiceDocumentation) {
		return true
	}

	return false
}

// SetServiceDocumentation gets a reference to the given string and assigns it to the ServiceDocumentation field.
func (o *Service) SetServiceDocumentation(v string) {
	o.ServiceDocumentation = &v
}

// GetBackchannelAuthenticationEndpoint returns the BackchannelAuthenticationEndpoint field value if set, zero value otherwise.
func (o *Service) GetBackchannelAuthenticationEndpoint() string {
	if o == nil || IsNil(o.BackchannelAuthenticationEndpoint) {
		var ret string
		return ret
	}
	return *o.BackchannelAuthenticationEndpoint
}

// GetBackchannelAuthenticationEndpointOk returns a tuple with the BackchannelAuthenticationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetBackchannelAuthenticationEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.BackchannelAuthenticationEndpoint) {
		return nil, false
	}
	return o.BackchannelAuthenticationEndpoint, true
}

// HasBackchannelAuthenticationEndpoint returns a boolean if a field has been set.
func (o *Service) HasBackchannelAuthenticationEndpoint() bool {
	if o != nil && !IsNil(o.BackchannelAuthenticationEndpoint) {
		return true
	}

	return false
}

// SetBackchannelAuthenticationEndpoint gets a reference to the given string and assigns it to the BackchannelAuthenticationEndpoint field.
func (o *Service) SetBackchannelAuthenticationEndpoint(v string) {
	o.BackchannelAuthenticationEndpoint = &v
}

// GetSupportedBackchannelTokenDeliveryModes returns the SupportedBackchannelTokenDeliveryModes field value if set, zero value otherwise.
func (o *Service) GetSupportedBackchannelTokenDeliveryModes() []DeliveryMode {
	if o == nil || IsNil(o.SupportedBackchannelTokenDeliveryModes) {
		var ret []DeliveryMode
		return ret
	}
	return o.SupportedBackchannelTokenDeliveryModes
}

// GetSupportedBackchannelTokenDeliveryModesOk returns a tuple with the SupportedBackchannelTokenDeliveryModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedBackchannelTokenDeliveryModesOk() ([]DeliveryMode, bool) {
	if o == nil || IsNil(o.SupportedBackchannelTokenDeliveryModes) {
		return nil, false
	}
	return o.SupportedBackchannelTokenDeliveryModes, true
}

// HasSupportedBackchannelTokenDeliveryModes returns a boolean if a field has been set.
func (o *Service) HasSupportedBackchannelTokenDeliveryModes() bool {
	if o != nil && !IsNil(o.SupportedBackchannelTokenDeliveryModes) {
		return true
	}

	return false
}

// SetSupportedBackchannelTokenDeliveryModes gets a reference to the given []DeliveryMode and assigns it to the SupportedBackchannelTokenDeliveryModes field.
func (o *Service) SetSupportedBackchannelTokenDeliveryModes(v []DeliveryMode) {
	o.SupportedBackchannelTokenDeliveryModes = v
}

// GetBackchannelAuthReqIdDuration returns the BackchannelAuthReqIdDuration field value if set, zero value otherwise.
func (o *Service) GetBackchannelAuthReqIdDuration() int32 {
	if o == nil || IsNil(o.BackchannelAuthReqIdDuration) {
		var ret int32
		return ret
	}
	return *o.BackchannelAuthReqIdDuration
}

// GetBackchannelAuthReqIdDurationOk returns a tuple with the BackchannelAuthReqIdDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetBackchannelAuthReqIdDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.BackchannelAuthReqIdDuration) {
		return nil, false
	}
	return o.BackchannelAuthReqIdDuration, true
}

// HasBackchannelAuthReqIdDuration returns a boolean if a field has been set.
func (o *Service) HasBackchannelAuthReqIdDuration() bool {
	if o != nil && !IsNil(o.BackchannelAuthReqIdDuration) {
		return true
	}

	return false
}

// SetBackchannelAuthReqIdDuration gets a reference to the given int32 and assigns it to the BackchannelAuthReqIdDuration field.
func (o *Service) SetBackchannelAuthReqIdDuration(v int32) {
	o.BackchannelAuthReqIdDuration = &v
}

// GetBackchannelPollingInterval returns the BackchannelPollingInterval field value if set, zero value otherwise.
func (o *Service) GetBackchannelPollingInterval() int32 {
	if o == nil || IsNil(o.BackchannelPollingInterval) {
		var ret int32
		return ret
	}
	return *o.BackchannelPollingInterval
}

// GetBackchannelPollingIntervalOk returns a tuple with the BackchannelPollingInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetBackchannelPollingIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.BackchannelPollingInterval) {
		return nil, false
	}
	return o.BackchannelPollingInterval, true
}

// HasBackchannelPollingInterval returns a boolean if a field has been set.
func (o *Service) HasBackchannelPollingInterval() bool {
	if o != nil && !IsNil(o.BackchannelPollingInterval) {
		return true
	}

	return false
}

// SetBackchannelPollingInterval gets a reference to the given int32 and assigns it to the BackchannelPollingInterval field.
func (o *Service) SetBackchannelPollingInterval(v int32) {
	o.BackchannelPollingInterval = &v
}

// GetBackchannelUserCodeParameterSupported returns the BackchannelUserCodeParameterSupported field value if set, zero value otherwise.
func (o *Service) GetBackchannelUserCodeParameterSupported() bool {
	if o == nil || IsNil(o.BackchannelUserCodeParameterSupported) {
		var ret bool
		return ret
	}
	return *o.BackchannelUserCodeParameterSupported
}

// GetBackchannelUserCodeParameterSupportedOk returns a tuple with the BackchannelUserCodeParameterSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetBackchannelUserCodeParameterSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.BackchannelUserCodeParameterSupported) {
		return nil, false
	}
	return o.BackchannelUserCodeParameterSupported, true
}

// HasBackchannelUserCodeParameterSupported returns a boolean if a field has been set.
func (o *Service) HasBackchannelUserCodeParameterSupported() bool {
	if o != nil && !IsNil(o.BackchannelUserCodeParameterSupported) {
		return true
	}

	return false
}

// SetBackchannelUserCodeParameterSupported gets a reference to the given bool and assigns it to the BackchannelUserCodeParameterSupported field.
func (o *Service) SetBackchannelUserCodeParameterSupported(v bool) {
	o.BackchannelUserCodeParameterSupported = &v
}

// GetBackchannelBindingMessageRequiredInFapi returns the BackchannelBindingMessageRequiredInFapi field value if set, zero value otherwise.
func (o *Service) GetBackchannelBindingMessageRequiredInFapi() bool {
	if o == nil || IsNil(o.BackchannelBindingMessageRequiredInFapi) {
		var ret bool
		return ret
	}
	return *o.BackchannelBindingMessageRequiredInFapi
}

// GetBackchannelBindingMessageRequiredInFapiOk returns a tuple with the BackchannelBindingMessageRequiredInFapi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetBackchannelBindingMessageRequiredInFapiOk() (*bool, bool) {
	if o == nil || IsNil(o.BackchannelBindingMessageRequiredInFapi) {
		return nil, false
	}
	return o.BackchannelBindingMessageRequiredInFapi, true
}

// HasBackchannelBindingMessageRequiredInFapi returns a boolean if a field has been set.
func (o *Service) HasBackchannelBindingMessageRequiredInFapi() bool {
	if o != nil && !IsNil(o.BackchannelBindingMessageRequiredInFapi) {
		return true
	}

	return false
}

// SetBackchannelBindingMessageRequiredInFapi gets a reference to the given bool and assigns it to the BackchannelBindingMessageRequiredInFapi field.
func (o *Service) SetBackchannelBindingMessageRequiredInFapi(v bool) {
	o.BackchannelBindingMessageRequiredInFapi = &v
}

// GetDeviceAuthorizationEndpoint returns the DeviceAuthorizationEndpoint field value if set, zero value otherwise.
func (o *Service) GetDeviceAuthorizationEndpoint() string {
	if o == nil || IsNil(o.DeviceAuthorizationEndpoint) {
		var ret string
		return ret
	}
	return *o.DeviceAuthorizationEndpoint
}

// GetDeviceAuthorizationEndpointOk returns a tuple with the DeviceAuthorizationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDeviceAuthorizationEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceAuthorizationEndpoint) {
		return nil, false
	}
	return o.DeviceAuthorizationEndpoint, true
}

// HasDeviceAuthorizationEndpoint returns a boolean if a field has been set.
func (o *Service) HasDeviceAuthorizationEndpoint() bool {
	if o != nil && !IsNil(o.DeviceAuthorizationEndpoint) {
		return true
	}

	return false
}

// SetDeviceAuthorizationEndpoint gets a reference to the given string and assigns it to the DeviceAuthorizationEndpoint field.
func (o *Service) SetDeviceAuthorizationEndpoint(v string) {
	o.DeviceAuthorizationEndpoint = &v
}

// GetDeviceVerificationUri returns the DeviceVerificationUri field value if set, zero value otherwise.
func (o *Service) GetDeviceVerificationUri() string {
	if o == nil || IsNil(o.DeviceVerificationUri) {
		var ret string
		return ret
	}
	return *o.DeviceVerificationUri
}

// GetDeviceVerificationUriOk returns a tuple with the DeviceVerificationUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDeviceVerificationUriOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceVerificationUri) {
		return nil, false
	}
	return o.DeviceVerificationUri, true
}

// HasDeviceVerificationUri returns a boolean if a field has been set.
func (o *Service) HasDeviceVerificationUri() bool {
	if o != nil && !IsNil(o.DeviceVerificationUri) {
		return true
	}

	return false
}

// SetDeviceVerificationUri gets a reference to the given string and assigns it to the DeviceVerificationUri field.
func (o *Service) SetDeviceVerificationUri(v string) {
	o.DeviceVerificationUri = &v
}

// GetDeviceVerificationUriComplete returns the DeviceVerificationUriComplete field value if set, zero value otherwise.
func (o *Service) GetDeviceVerificationUriComplete() string {
	if o == nil || IsNil(o.DeviceVerificationUriComplete) {
		var ret string
		return ret
	}
	return *o.DeviceVerificationUriComplete
}

// GetDeviceVerificationUriCompleteOk returns a tuple with the DeviceVerificationUriComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDeviceVerificationUriCompleteOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceVerificationUriComplete) {
		return nil, false
	}
	return o.DeviceVerificationUriComplete, true
}

// HasDeviceVerificationUriComplete returns a boolean if a field has been set.
func (o *Service) HasDeviceVerificationUriComplete() bool {
	if o != nil && !IsNil(o.DeviceVerificationUriComplete) {
		return true
	}

	return false
}

// SetDeviceVerificationUriComplete gets a reference to the given string and assigns it to the DeviceVerificationUriComplete field.
func (o *Service) SetDeviceVerificationUriComplete(v string) {
	o.DeviceVerificationUriComplete = &v
}

// GetDeviceFlowCodeDuration returns the DeviceFlowCodeDuration field value if set, zero value otherwise.
func (o *Service) GetDeviceFlowCodeDuration() int32 {
	if o == nil || IsNil(o.DeviceFlowCodeDuration) {
		var ret int32
		return ret
	}
	return *o.DeviceFlowCodeDuration
}

// GetDeviceFlowCodeDurationOk returns a tuple with the DeviceFlowCodeDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDeviceFlowCodeDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.DeviceFlowCodeDuration) {
		return nil, false
	}
	return o.DeviceFlowCodeDuration, true
}

// HasDeviceFlowCodeDuration returns a boolean if a field has been set.
func (o *Service) HasDeviceFlowCodeDuration() bool {
	if o != nil && !IsNil(o.DeviceFlowCodeDuration) {
		return true
	}

	return false
}

// SetDeviceFlowCodeDuration gets a reference to the given int32 and assigns it to the DeviceFlowCodeDuration field.
func (o *Service) SetDeviceFlowCodeDuration(v int32) {
	o.DeviceFlowCodeDuration = &v
}

// GetDeviceFlowPollingInterval returns the DeviceFlowPollingInterval field value if set, zero value otherwise.
func (o *Service) GetDeviceFlowPollingInterval() int32 {
	if o == nil || IsNil(o.DeviceFlowPollingInterval) {
		var ret int32
		return ret
	}
	return *o.DeviceFlowPollingInterval
}

// GetDeviceFlowPollingIntervalOk returns a tuple with the DeviceFlowPollingInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDeviceFlowPollingIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.DeviceFlowPollingInterval) {
		return nil, false
	}
	return o.DeviceFlowPollingInterval, true
}

// HasDeviceFlowPollingInterval returns a boolean if a field has been set.
func (o *Service) HasDeviceFlowPollingInterval() bool {
	if o != nil && !IsNil(o.DeviceFlowPollingInterval) {
		return true
	}

	return false
}

// SetDeviceFlowPollingInterval gets a reference to the given int32 and assigns it to the DeviceFlowPollingInterval field.
func (o *Service) SetDeviceFlowPollingInterval(v int32) {
	o.DeviceFlowPollingInterval = &v
}

// GetUserCodeCharset returns the UserCodeCharset field value if set, zero value otherwise.
func (o *Service) GetUserCodeCharset() UserCodeCharset {
	if o == nil || IsNil(o.UserCodeCharset) {
		var ret UserCodeCharset
		return ret
	}
	return *o.UserCodeCharset
}

// GetUserCodeCharsetOk returns a tuple with the UserCodeCharset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetUserCodeCharsetOk() (*UserCodeCharset, bool) {
	if o == nil || IsNil(o.UserCodeCharset) {
		return nil, false
	}
	return o.UserCodeCharset, true
}

// HasUserCodeCharset returns a boolean if a field has been set.
func (o *Service) HasUserCodeCharset() bool {
	if o != nil && !IsNil(o.UserCodeCharset) {
		return true
	}

	return false
}

// SetUserCodeCharset gets a reference to the given UserCodeCharset and assigns it to the UserCodeCharset field.
func (o *Service) SetUserCodeCharset(v UserCodeCharset) {
	o.UserCodeCharset = &v
}

// GetUserCodeLength returns the UserCodeLength field value if set, zero value otherwise.
func (o *Service) GetUserCodeLength() int32 {
	if o == nil || IsNil(o.UserCodeLength) {
		var ret int32
		return ret
	}
	return *o.UserCodeLength
}

// GetUserCodeLengthOk returns a tuple with the UserCodeLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetUserCodeLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.UserCodeLength) {
		return nil, false
	}
	return o.UserCodeLength, true
}

// HasUserCodeLength returns a boolean if a field has been set.
func (o *Service) HasUserCodeLength() bool {
	if o != nil && !IsNil(o.UserCodeLength) {
		return true
	}

	return false
}

// SetUserCodeLength gets a reference to the given int32 and assigns it to the UserCodeLength field.
func (o *Service) SetUserCodeLength(v int32) {
	o.UserCodeLength = &v
}

// GetSupportedTrustFrameworks returns the SupportedTrustFrameworks field value if set, zero value otherwise.
func (o *Service) GetSupportedTrustFrameworks() []string {
	if o == nil || IsNil(o.SupportedTrustFrameworks) {
		var ret []string
		return ret
	}
	return o.SupportedTrustFrameworks
}

// GetSupportedTrustFrameworksOk returns a tuple with the SupportedTrustFrameworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedTrustFrameworksOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedTrustFrameworks) {
		return nil, false
	}
	return o.SupportedTrustFrameworks, true
}

// HasSupportedTrustFrameworks returns a boolean if a field has been set.
func (o *Service) HasSupportedTrustFrameworks() bool {
	if o != nil && !IsNil(o.SupportedTrustFrameworks) {
		return true
	}

	return false
}

// SetSupportedTrustFrameworks gets a reference to the given []string and assigns it to the SupportedTrustFrameworks field.
func (o *Service) SetSupportedTrustFrameworks(v []string) {
	o.SupportedTrustFrameworks = v
}

// GetSupportedEvidence returns the SupportedEvidence field value if set, zero value otherwise.
func (o *Service) GetSupportedEvidence() []string {
	if o == nil || IsNil(o.SupportedEvidence) {
		var ret []string
		return ret
	}
	return o.SupportedEvidence
}

// GetSupportedEvidenceOk returns a tuple with the SupportedEvidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedEvidenceOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedEvidence) {
		return nil, false
	}
	return o.SupportedEvidence, true
}

// HasSupportedEvidence returns a boolean if a field has been set.
func (o *Service) HasSupportedEvidence() bool {
	if o != nil && !IsNil(o.SupportedEvidence) {
		return true
	}

	return false
}

// SetSupportedEvidence gets a reference to the given []string and assigns it to the SupportedEvidence field.
func (o *Service) SetSupportedEvidence(v []string) {
	o.SupportedEvidence = v
}

// GetSupportedIdentityDocuments returns the SupportedIdentityDocuments field value if set, zero value otherwise.
func (o *Service) GetSupportedIdentityDocuments() []string {
	if o == nil || IsNil(o.SupportedIdentityDocuments) {
		var ret []string
		return ret
	}
	return o.SupportedIdentityDocuments
}

// GetSupportedIdentityDocumentsOk returns a tuple with the SupportedIdentityDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedIdentityDocumentsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedIdentityDocuments) {
		return nil, false
	}
	return o.SupportedIdentityDocuments, true
}

// HasSupportedIdentityDocuments returns a boolean if a field has been set.
func (o *Service) HasSupportedIdentityDocuments() bool {
	if o != nil && !IsNil(o.SupportedIdentityDocuments) {
		return true
	}

	return false
}

// SetSupportedIdentityDocuments gets a reference to the given []string and assigns it to the SupportedIdentityDocuments field.
func (o *Service) SetSupportedIdentityDocuments(v []string) {
	o.SupportedIdentityDocuments = v
}

// GetSupportedVerificationMethods returns the SupportedVerificationMethods field value if set, zero value otherwise.
func (o *Service) GetSupportedVerificationMethods() []string {
	if o == nil || IsNil(o.SupportedVerificationMethods) {
		var ret []string
		return ret
	}
	return o.SupportedVerificationMethods
}

// GetSupportedVerificationMethodsOk returns a tuple with the SupportedVerificationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedVerificationMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedVerificationMethods) {
		return nil, false
	}
	return o.SupportedVerificationMethods, true
}

// HasSupportedVerificationMethods returns a boolean if a field has been set.
func (o *Service) HasSupportedVerificationMethods() bool {
	if o != nil && !IsNil(o.SupportedVerificationMethods) {
		return true
	}

	return false
}

// SetSupportedVerificationMethods gets a reference to the given []string and assigns it to the SupportedVerificationMethods field.
func (o *Service) SetSupportedVerificationMethods(v []string) {
	o.SupportedVerificationMethods = v
}

// GetSupportedVerifiedClaims returns the SupportedVerifiedClaims field value if set, zero value otherwise.
func (o *Service) GetSupportedVerifiedClaims() []string {
	if o == nil || IsNil(o.SupportedVerifiedClaims) {
		var ret []string
		return ret
	}
	return o.SupportedVerifiedClaims
}

// GetSupportedVerifiedClaimsOk returns a tuple with the SupportedVerifiedClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedVerifiedClaimsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedVerifiedClaims) {
		return nil, false
	}
	return o.SupportedVerifiedClaims, true
}

// HasSupportedVerifiedClaims returns a boolean if a field has been set.
func (o *Service) HasSupportedVerifiedClaims() bool {
	if o != nil && !IsNil(o.SupportedVerifiedClaims) {
		return true
	}

	return false
}

// SetSupportedVerifiedClaims gets a reference to the given []string and assigns it to the SupportedVerifiedClaims field.
func (o *Service) SetSupportedVerifiedClaims(v []string) {
	o.SupportedVerifiedClaims = v
}

// GetVerifiedClaimsValidationSchemaSet returns the VerifiedClaimsValidationSchemaSet field value if set, zero value otherwise.
func (o *Service) GetVerifiedClaimsValidationSchemaSet() VerifiedClaimsValidationSchema {
	if o == nil || IsNil(o.VerifiedClaimsValidationSchemaSet) {
		var ret VerifiedClaimsValidationSchema
		return ret
	}
	return *o.VerifiedClaimsValidationSchemaSet
}

// GetVerifiedClaimsValidationSchemaSetOk returns a tuple with the VerifiedClaimsValidationSchemaSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetVerifiedClaimsValidationSchemaSetOk() (*VerifiedClaimsValidationSchema, bool) {
	if o == nil || IsNil(o.VerifiedClaimsValidationSchemaSet) {
		return nil, false
	}
	return o.VerifiedClaimsValidationSchemaSet, true
}

// HasVerifiedClaimsValidationSchemaSet returns a boolean if a field has been set.
func (o *Service) HasVerifiedClaimsValidationSchemaSet() bool {
	if o != nil && !IsNil(o.VerifiedClaimsValidationSchemaSet) {
		return true
	}

	return false
}

// SetVerifiedClaimsValidationSchemaSet gets a reference to the given VerifiedClaimsValidationSchema and assigns it to the VerifiedClaimsValidationSchemaSet field.
func (o *Service) SetVerifiedClaimsValidationSchemaSet(v VerifiedClaimsValidationSchema) {
	o.VerifiedClaimsValidationSchemaSet = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Service) GetAttributes() []Pair {
	if o == nil || IsNil(o.Attributes) {
		var ret []Pair
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAttributesOk() ([]Pair, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Service) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []Pair and assigns it to the Attributes field.
func (o *Service) SetAttributes(v []Pair) {
	o.Attributes = v
}

// GetNbfOptional returns the NbfOptional field value if set, zero value otherwise.
func (o *Service) GetNbfOptional() bool {
	if o == nil || IsNil(o.NbfOptional) {
		var ret bool
		return ret
	}
	return *o.NbfOptional
}

// GetNbfOptionalOk returns a tuple with the NbfOptional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetNbfOptionalOk() (*bool, bool) {
	if o == nil || IsNil(o.NbfOptional) {
		return nil, false
	}
	return o.NbfOptional, true
}

// HasNbfOptional returns a boolean if a field has been set.
func (o *Service) HasNbfOptional() bool {
	if o != nil && !IsNil(o.NbfOptional) {
		return true
	}

	return false
}

// SetNbfOptional gets a reference to the given bool and assigns it to the NbfOptional field.
func (o *Service) SetNbfOptional(v bool) {
	o.NbfOptional = &v
}

// GetIssSuppressed returns the IssSuppressed field value if set, zero value otherwise.
func (o *Service) GetIssSuppressed() bool {
	if o == nil || IsNil(o.IssSuppressed) {
		var ret bool
		return ret
	}
	return *o.IssSuppressed
}

// GetIssSuppressedOk returns a tuple with the IssSuppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetIssSuppressedOk() (*bool, bool) {
	if o == nil || IsNil(o.IssSuppressed) {
		return nil, false
	}
	return o.IssSuppressed, true
}

// HasIssSuppressed returns a boolean if a field has been set.
func (o *Service) HasIssSuppressed() bool {
	if o != nil && !IsNil(o.IssSuppressed) {
		return true
	}

	return false
}

// SetIssSuppressed gets a reference to the given bool and assigns it to the IssSuppressed field.
func (o *Service) SetIssSuppressed(v bool) {
	o.IssSuppressed = &v
}

// GetSupportedCustomClientMetadata returns the SupportedCustomClientMetadata field value if set, zero value otherwise.
func (o *Service) GetSupportedCustomClientMetadata() []string {
	if o == nil || IsNil(o.SupportedCustomClientMetadata) {
		var ret []string
		return ret
	}
	return o.SupportedCustomClientMetadata
}

// GetSupportedCustomClientMetadataOk returns a tuple with the SupportedCustomClientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedCustomClientMetadataOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedCustomClientMetadata) {
		return nil, false
	}
	return o.SupportedCustomClientMetadata, true
}

// HasSupportedCustomClientMetadata returns a boolean if a field has been set.
func (o *Service) HasSupportedCustomClientMetadata() bool {
	if o != nil && !IsNil(o.SupportedCustomClientMetadata) {
		return true
	}

	return false
}

// SetSupportedCustomClientMetadata gets a reference to the given []string and assigns it to the SupportedCustomClientMetadata field.
func (o *Service) SetSupportedCustomClientMetadata(v []string) {
	o.SupportedCustomClientMetadata = v
}

// GetTokenExpirationLinked returns the TokenExpirationLinked field value if set, zero value otherwise.
func (o *Service) GetTokenExpirationLinked() bool {
	if o == nil || IsNil(o.TokenExpirationLinked) {
		var ret bool
		return ret
	}
	return *o.TokenExpirationLinked
}

// GetTokenExpirationLinkedOk returns a tuple with the TokenExpirationLinked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetTokenExpirationLinkedOk() (*bool, bool) {
	if o == nil || IsNil(o.TokenExpirationLinked) {
		return nil, false
	}
	return o.TokenExpirationLinked, true
}

// HasTokenExpirationLinked returns a boolean if a field has been set.
func (o *Service) HasTokenExpirationLinked() bool {
	if o != nil && !IsNil(o.TokenExpirationLinked) {
		return true
	}

	return false
}

// SetTokenExpirationLinked gets a reference to the given bool and assigns it to the TokenExpirationLinked field.
func (o *Service) SetTokenExpirationLinked(v bool) {
	o.TokenExpirationLinked = &v
}

// GetFrontChannelRequestObjectEncryptionRequired returns the FrontChannelRequestObjectEncryptionRequired field value if set, zero value otherwise.
func (o *Service) GetFrontChannelRequestObjectEncryptionRequired() bool {
	if o == nil || IsNil(o.FrontChannelRequestObjectEncryptionRequired) {
		var ret bool
		return ret
	}
	return *o.FrontChannelRequestObjectEncryptionRequired
}

// GetFrontChannelRequestObjectEncryptionRequiredOk returns a tuple with the FrontChannelRequestObjectEncryptionRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetFrontChannelRequestObjectEncryptionRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.FrontChannelRequestObjectEncryptionRequired) {
		return nil, false
	}
	return o.FrontChannelRequestObjectEncryptionRequired, true
}

// HasFrontChannelRequestObjectEncryptionRequired returns a boolean if a field has been set.
func (o *Service) HasFrontChannelRequestObjectEncryptionRequired() bool {
	if o != nil && !IsNil(o.FrontChannelRequestObjectEncryptionRequired) {
		return true
	}

	return false
}

// SetFrontChannelRequestObjectEncryptionRequired gets a reference to the given bool and assigns it to the FrontChannelRequestObjectEncryptionRequired field.
func (o *Service) SetFrontChannelRequestObjectEncryptionRequired(v bool) {
	o.FrontChannelRequestObjectEncryptionRequired = &v
}

// GetRequestObjectEncryptionAlgMatchRequired returns the RequestObjectEncryptionAlgMatchRequired field value if set, zero value otherwise.
func (o *Service) GetRequestObjectEncryptionAlgMatchRequired() bool {
	if o == nil || IsNil(o.RequestObjectEncryptionAlgMatchRequired) {
		var ret bool
		return ret
	}
	return *o.RequestObjectEncryptionAlgMatchRequired
}

// GetRequestObjectEncryptionAlgMatchRequiredOk returns a tuple with the RequestObjectEncryptionAlgMatchRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetRequestObjectEncryptionAlgMatchRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestObjectEncryptionAlgMatchRequired) {
		return nil, false
	}
	return o.RequestObjectEncryptionAlgMatchRequired, true
}

// HasRequestObjectEncryptionAlgMatchRequired returns a boolean if a field has been set.
func (o *Service) HasRequestObjectEncryptionAlgMatchRequired() bool {
	if o != nil && !IsNil(o.RequestObjectEncryptionAlgMatchRequired) {
		return true
	}

	return false
}

// SetRequestObjectEncryptionAlgMatchRequired gets a reference to the given bool and assigns it to the RequestObjectEncryptionAlgMatchRequired field.
func (o *Service) SetRequestObjectEncryptionAlgMatchRequired(v bool) {
	o.RequestObjectEncryptionAlgMatchRequired = &v
}

// GetRequestObjectEncryptionEncMatchRequired returns the RequestObjectEncryptionEncMatchRequired field value if set, zero value otherwise.
func (o *Service) GetRequestObjectEncryptionEncMatchRequired() bool {
	if o == nil || IsNil(o.RequestObjectEncryptionEncMatchRequired) {
		var ret bool
		return ret
	}
	return *o.RequestObjectEncryptionEncMatchRequired
}

// GetRequestObjectEncryptionEncMatchRequiredOk returns a tuple with the RequestObjectEncryptionEncMatchRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetRequestObjectEncryptionEncMatchRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestObjectEncryptionEncMatchRequired) {
		return nil, false
	}
	return o.RequestObjectEncryptionEncMatchRequired, true
}

// HasRequestObjectEncryptionEncMatchRequired returns a boolean if a field has been set.
func (o *Service) HasRequestObjectEncryptionEncMatchRequired() bool {
	if o != nil && !IsNil(o.RequestObjectEncryptionEncMatchRequired) {
		return true
	}

	return false
}

// SetRequestObjectEncryptionEncMatchRequired gets a reference to the given bool and assigns it to the RequestObjectEncryptionEncMatchRequired field.
func (o *Service) SetRequestObjectEncryptionEncMatchRequired(v bool) {
	o.RequestObjectEncryptionEncMatchRequired = &v
}

// GetHsmEnabled returns the HsmEnabled field value if set, zero value otherwise.
func (o *Service) GetHsmEnabled() bool {
	if o == nil || IsNil(o.HsmEnabled) {
		var ret bool
		return ret
	}
	return *o.HsmEnabled
}

// GetHsmEnabledOk returns a tuple with the HsmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetHsmEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.HsmEnabled) {
		return nil, false
	}
	return o.HsmEnabled, true
}

// HasHsmEnabled returns a boolean if a field has been set.
func (o *Service) HasHsmEnabled() bool {
	if o != nil && !IsNil(o.HsmEnabled) {
		return true
	}

	return false
}

// SetHsmEnabled gets a reference to the given bool and assigns it to the HsmEnabled field.
func (o *Service) SetHsmEnabled(v bool) {
	o.HsmEnabled = &v
}

// GetHsks returns the Hsks field value if set, zero value otherwise.
func (o *Service) GetHsks() []Hsk {
	if o == nil || IsNil(o.Hsks) {
		var ret []Hsk
		return ret
	}
	return o.Hsks
}

// GetHsksOk returns a tuple with the Hsks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetHsksOk() ([]Hsk, bool) {
	if o == nil || IsNil(o.Hsks) {
		return nil, false
	}
	return o.Hsks, true
}

// HasHsks returns a boolean if a field has been set.
func (o *Service) HasHsks() bool {
	if o != nil && !IsNil(o.Hsks) {
		return true
	}

	return false
}

// SetHsks gets a reference to the given []Hsk and assigns it to the Hsks field.
func (o *Service) SetHsks(v []Hsk) {
	o.Hsks = v
}

// GetGrantManagementEndpoint returns the GrantManagementEndpoint field value if set, zero value otherwise.
func (o *Service) GetGrantManagementEndpoint() string {
	if o == nil || IsNil(o.GrantManagementEndpoint) {
		var ret string
		return ret
	}
	return *o.GrantManagementEndpoint
}

// GetGrantManagementEndpointOk returns a tuple with the GrantManagementEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetGrantManagementEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.GrantManagementEndpoint) {
		return nil, false
	}
	return o.GrantManagementEndpoint, true
}

// HasGrantManagementEndpoint returns a boolean if a field has been set.
func (o *Service) HasGrantManagementEndpoint() bool {
	if o != nil && !IsNil(o.GrantManagementEndpoint) {
		return true
	}

	return false
}

// SetGrantManagementEndpoint gets a reference to the given string and assigns it to the GrantManagementEndpoint field.
func (o *Service) SetGrantManagementEndpoint(v string) {
	o.GrantManagementEndpoint = &v
}

// GetGrantManagementActionRequired returns the GrantManagementActionRequired field value if set, zero value otherwise.
func (o *Service) GetGrantManagementActionRequired() bool {
	if o == nil || IsNil(o.GrantManagementActionRequired) {
		var ret bool
		return ret
	}
	return *o.GrantManagementActionRequired
}

// GetGrantManagementActionRequiredOk returns a tuple with the GrantManagementActionRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetGrantManagementActionRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.GrantManagementActionRequired) {
		return nil, false
	}
	return o.GrantManagementActionRequired, true
}

// HasGrantManagementActionRequired returns a boolean if a field has been set.
func (o *Service) HasGrantManagementActionRequired() bool {
	if o != nil && !IsNil(o.GrantManagementActionRequired) {
		return true
	}

	return false
}

// SetGrantManagementActionRequired gets a reference to the given bool and assigns it to the GrantManagementActionRequired field.
func (o *Service) SetGrantManagementActionRequired(v bool) {
	o.GrantManagementActionRequired = &v
}

// GetUnauthorizedOnClientConfigSupported returns the UnauthorizedOnClientConfigSupported field value if set, zero value otherwise.
func (o *Service) GetUnauthorizedOnClientConfigSupported() bool {
	if o == nil || IsNil(o.UnauthorizedOnClientConfigSupported) {
		var ret bool
		return ret
	}
	return *o.UnauthorizedOnClientConfigSupported
}

// GetUnauthorizedOnClientConfigSupportedOk returns a tuple with the UnauthorizedOnClientConfigSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetUnauthorizedOnClientConfigSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.UnauthorizedOnClientConfigSupported) {
		return nil, false
	}
	return o.UnauthorizedOnClientConfigSupported, true
}

// HasUnauthorizedOnClientConfigSupported returns a boolean if a field has been set.
func (o *Service) HasUnauthorizedOnClientConfigSupported() bool {
	if o != nil && !IsNil(o.UnauthorizedOnClientConfigSupported) {
		return true
	}

	return false
}

// SetUnauthorizedOnClientConfigSupported gets a reference to the given bool and assigns it to the UnauthorizedOnClientConfigSupported field.
func (o *Service) SetUnauthorizedOnClientConfigSupported(v bool) {
	o.UnauthorizedOnClientConfigSupported = &v
}

// GetDcrScopeUsedAsRequestable returns the DcrScopeUsedAsRequestable field value if set, zero value otherwise.
func (o *Service) GetDcrScopeUsedAsRequestable() bool {
	if o == nil || IsNil(o.DcrScopeUsedAsRequestable) {
		var ret bool
		return ret
	}
	return *o.DcrScopeUsedAsRequestable
}

// GetDcrScopeUsedAsRequestableOk returns a tuple with the DcrScopeUsedAsRequestable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDcrScopeUsedAsRequestableOk() (*bool, bool) {
	if o == nil || IsNil(o.DcrScopeUsedAsRequestable) {
		return nil, false
	}
	return o.DcrScopeUsedAsRequestable, true
}

// HasDcrScopeUsedAsRequestable returns a boolean if a field has been set.
func (o *Service) HasDcrScopeUsedAsRequestable() bool {
	if o != nil && !IsNil(o.DcrScopeUsedAsRequestable) {
		return true
	}

	return false
}

// SetDcrScopeUsedAsRequestable gets a reference to the given bool and assigns it to the DcrScopeUsedAsRequestable field.
func (o *Service) SetDcrScopeUsedAsRequestable(v bool) {
	o.DcrScopeUsedAsRequestable = &v
}

// GetEndSessionEndpoint returns the EndSessionEndpoint field value if set, zero value otherwise.
func (o *Service) GetEndSessionEndpoint() string {
	if o == nil || IsNil(o.EndSessionEndpoint) {
		var ret string
		return ret
	}
	return *o.EndSessionEndpoint
}

// GetEndSessionEndpointOk returns a tuple with the EndSessionEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetEndSessionEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.EndSessionEndpoint) {
		return nil, false
	}
	return o.EndSessionEndpoint, true
}

// HasEndSessionEndpoint returns a boolean if a field has been set.
func (o *Service) HasEndSessionEndpoint() bool {
	if o != nil && !IsNil(o.EndSessionEndpoint) {
		return true
	}

	return false
}

// SetEndSessionEndpoint gets a reference to the given string and assigns it to the EndSessionEndpoint field.
func (o *Service) SetEndSessionEndpoint(v string) {
	o.EndSessionEndpoint = &v
}

// GetLoopbackRedirectionUriVariable returns the LoopbackRedirectionUriVariable field value if set, zero value otherwise.
func (o *Service) GetLoopbackRedirectionUriVariable() bool {
	if o == nil || IsNil(o.LoopbackRedirectionUriVariable) {
		var ret bool
		return ret
	}
	return *o.LoopbackRedirectionUriVariable
}

// GetLoopbackRedirectionUriVariableOk returns a tuple with the LoopbackRedirectionUriVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetLoopbackRedirectionUriVariableOk() (*bool, bool) {
	if o == nil || IsNil(o.LoopbackRedirectionUriVariable) {
		return nil, false
	}
	return o.LoopbackRedirectionUriVariable, true
}

// HasLoopbackRedirectionUriVariable returns a boolean if a field has been set.
func (o *Service) HasLoopbackRedirectionUriVariable() bool {
	if o != nil && !IsNil(o.LoopbackRedirectionUriVariable) {
		return true
	}

	return false
}

// SetLoopbackRedirectionUriVariable gets a reference to the given bool and assigns it to the LoopbackRedirectionUriVariable field.
func (o *Service) SetLoopbackRedirectionUriVariable(v bool) {
	o.LoopbackRedirectionUriVariable = &v
}

// GetRequestObjectAudienceChecked returns the RequestObjectAudienceChecked field value if set, zero value otherwise.
func (o *Service) GetRequestObjectAudienceChecked() bool {
	if o == nil || IsNil(o.RequestObjectAudienceChecked) {
		var ret bool
		return ret
	}
	return *o.RequestObjectAudienceChecked
}

// GetRequestObjectAudienceCheckedOk returns a tuple with the RequestObjectAudienceChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetRequestObjectAudienceCheckedOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestObjectAudienceChecked) {
		return nil, false
	}
	return o.RequestObjectAudienceChecked, true
}

// HasRequestObjectAudienceChecked returns a boolean if a field has been set.
func (o *Service) HasRequestObjectAudienceChecked() bool {
	if o != nil && !IsNil(o.RequestObjectAudienceChecked) {
		return true
	}

	return false
}

// SetRequestObjectAudienceChecked gets a reference to the given bool and assigns it to the RequestObjectAudienceChecked field.
func (o *Service) SetRequestObjectAudienceChecked(v bool) {
	o.RequestObjectAudienceChecked = &v
}

// GetAccessTokenForExternalAttachmentEmbedded returns the AccessTokenForExternalAttachmentEmbedded field value if set, zero value otherwise.
func (o *Service) GetAccessTokenForExternalAttachmentEmbedded() bool {
	if o == nil || IsNil(o.AccessTokenForExternalAttachmentEmbedded) {
		var ret bool
		return ret
	}
	return *o.AccessTokenForExternalAttachmentEmbedded
}

// GetAccessTokenForExternalAttachmentEmbeddedOk returns a tuple with the AccessTokenForExternalAttachmentEmbedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAccessTokenForExternalAttachmentEmbeddedOk() (*bool, bool) {
	if o == nil || IsNil(o.AccessTokenForExternalAttachmentEmbedded) {
		return nil, false
	}
	return o.AccessTokenForExternalAttachmentEmbedded, true
}

// HasAccessTokenForExternalAttachmentEmbedded returns a boolean if a field has been set.
func (o *Service) HasAccessTokenForExternalAttachmentEmbedded() bool {
	if o != nil && !IsNil(o.AccessTokenForExternalAttachmentEmbedded) {
		return true
	}

	return false
}

// SetAccessTokenForExternalAttachmentEmbedded gets a reference to the given bool and assigns it to the AccessTokenForExternalAttachmentEmbedded field.
func (o *Service) SetAccessTokenForExternalAttachmentEmbedded(v bool) {
	o.AccessTokenForExternalAttachmentEmbedded = &v
}

// GetAuthorityHints returns the AuthorityHints field value if set, zero value otherwise.
func (o *Service) GetAuthorityHints() []string {
	if o == nil || IsNil(o.AuthorityHints) {
		var ret []string
		return ret
	}
	return o.AuthorityHints
}

// GetAuthorityHintsOk returns a tuple with the AuthorityHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetAuthorityHintsOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthorityHints) {
		return nil, false
	}
	return o.AuthorityHints, true
}

// HasAuthorityHints returns a boolean if a field has been set.
func (o *Service) HasAuthorityHints() bool {
	if o != nil && !IsNil(o.AuthorityHints) {
		return true
	}

	return false
}

// SetAuthorityHints gets a reference to the given []string and assigns it to the AuthorityHints field.
func (o *Service) SetAuthorityHints(v []string) {
	o.AuthorityHints = v
}

// GetFederationEnabled returns the FederationEnabled field value if set, zero value otherwise.
func (o *Service) GetFederationEnabled() bool {
	if o == nil || IsNil(o.FederationEnabled) {
		var ret bool
		return ret
	}
	return *o.FederationEnabled
}

// GetFederationEnabledOk returns a tuple with the FederationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetFederationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FederationEnabled) {
		return nil, false
	}
	return o.FederationEnabled, true
}

// HasFederationEnabled returns a boolean if a field has been set.
func (o *Service) HasFederationEnabled() bool {
	if o != nil && !IsNil(o.FederationEnabled) {
		return true
	}

	return false
}

// SetFederationEnabled gets a reference to the given bool and assigns it to the FederationEnabled field.
func (o *Service) SetFederationEnabled(v bool) {
	o.FederationEnabled = &v
}

// GetFederationJwks returns the FederationJwks field value if set, zero value otherwise.
func (o *Service) GetFederationJwks() string {
	if o == nil || IsNil(o.FederationJwks) {
		var ret string
		return ret
	}
	return *o.FederationJwks
}

// GetFederationJwksOk returns a tuple with the FederationJwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetFederationJwksOk() (*string, bool) {
	if o == nil || IsNil(o.FederationJwks) {
		return nil, false
	}
	return o.FederationJwks, true
}

// HasFederationJwks returns a boolean if a field has been set.
func (o *Service) HasFederationJwks() bool {
	if o != nil && !IsNil(o.FederationJwks) {
		return true
	}

	return false
}

// SetFederationJwks gets a reference to the given string and assigns it to the FederationJwks field.
func (o *Service) SetFederationJwks(v string) {
	o.FederationJwks = &v
}

// GetFederationSignatureKeyId returns the FederationSignatureKeyId field value if set, zero value otherwise.
func (o *Service) GetFederationSignatureKeyId() string {
	if o == nil || IsNil(o.FederationSignatureKeyId) {
		var ret string
		return ret
	}
	return *o.FederationSignatureKeyId
}

// GetFederationSignatureKeyIdOk returns a tuple with the FederationSignatureKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetFederationSignatureKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.FederationSignatureKeyId) {
		return nil, false
	}
	return o.FederationSignatureKeyId, true
}

// HasFederationSignatureKeyId returns a boolean if a field has been set.
func (o *Service) HasFederationSignatureKeyId() bool {
	if o != nil && !IsNil(o.FederationSignatureKeyId) {
		return true
	}

	return false
}

// SetFederationSignatureKeyId gets a reference to the given string and assigns it to the FederationSignatureKeyId field.
func (o *Service) SetFederationSignatureKeyId(v string) {
	o.FederationSignatureKeyId = &v
}

// GetFederationConfigurationDuration returns the FederationConfigurationDuration field value if set, zero value otherwise.
func (o *Service) GetFederationConfigurationDuration() int32 {
	if o == nil || IsNil(o.FederationConfigurationDuration) {
		var ret int32
		return ret
	}
	return *o.FederationConfigurationDuration
}

// GetFederationConfigurationDurationOk returns a tuple with the FederationConfigurationDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetFederationConfigurationDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.FederationConfigurationDuration) {
		return nil, false
	}
	return o.FederationConfigurationDuration, true
}

// HasFederationConfigurationDuration returns a boolean if a field has been set.
func (o *Service) HasFederationConfigurationDuration() bool {
	if o != nil && !IsNil(o.FederationConfigurationDuration) {
		return true
	}

	return false
}

// SetFederationConfigurationDuration gets a reference to the given int32 and assigns it to the FederationConfigurationDuration field.
func (o *Service) SetFederationConfigurationDuration(v int32) {
	o.FederationConfigurationDuration = &v
}

// GetFederationRegistrationEndpoint returns the FederationRegistrationEndpoint field value if set, zero value otherwise.
func (o *Service) GetFederationRegistrationEndpoint() string {
	if o == nil || IsNil(o.FederationRegistrationEndpoint) {
		var ret string
		return ret
	}
	return *o.FederationRegistrationEndpoint
}

// GetFederationRegistrationEndpointOk returns a tuple with the FederationRegistrationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetFederationRegistrationEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.FederationRegistrationEndpoint) {
		return nil, false
	}
	return o.FederationRegistrationEndpoint, true
}

// HasFederationRegistrationEndpoint returns a boolean if a field has been set.
func (o *Service) HasFederationRegistrationEndpoint() bool {
	if o != nil && !IsNil(o.FederationRegistrationEndpoint) {
		return true
	}

	return false
}

// SetFederationRegistrationEndpoint gets a reference to the given string and assigns it to the FederationRegistrationEndpoint field.
func (o *Service) SetFederationRegistrationEndpoint(v string) {
	o.FederationRegistrationEndpoint = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *Service) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationName) {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *Service) HasOrganizationName() bool {
	if o != nil && !IsNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *Service) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetPredefinedTransformedClaims returns the PredefinedTransformedClaims field value if set, zero value otherwise.
func (o *Service) GetPredefinedTransformedClaims() string {
	if o == nil || IsNil(o.PredefinedTransformedClaims) {
		var ret string
		return ret
	}
	return *o.PredefinedTransformedClaims
}

// GetPredefinedTransformedClaimsOk returns a tuple with the PredefinedTransformedClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetPredefinedTransformedClaimsOk() (*string, bool) {
	if o == nil || IsNil(o.PredefinedTransformedClaims) {
		return nil, false
	}
	return o.PredefinedTransformedClaims, true
}

// HasPredefinedTransformedClaims returns a boolean if a field has been set.
func (o *Service) HasPredefinedTransformedClaims() bool {
	if o != nil && !IsNil(o.PredefinedTransformedClaims) {
		return true
	}

	return false
}

// SetPredefinedTransformedClaims gets a reference to the given string and assigns it to the PredefinedTransformedClaims field.
func (o *Service) SetPredefinedTransformedClaims(v string) {
	o.PredefinedTransformedClaims = &v
}

// GetRefreshTokenIdempotent returns the RefreshTokenIdempotent field value if set, zero value otherwise.
func (o *Service) GetRefreshTokenIdempotent() bool {
	if o == nil || IsNil(o.RefreshTokenIdempotent) {
		var ret bool
		return ret
	}
	return *o.RefreshTokenIdempotent
}

// GetRefreshTokenIdempotentOk returns a tuple with the RefreshTokenIdempotent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetRefreshTokenIdempotentOk() (*bool, bool) {
	if o == nil || IsNil(o.RefreshTokenIdempotent) {
		return nil, false
	}
	return o.RefreshTokenIdempotent, true
}

// HasRefreshTokenIdempotent returns a boolean if a field has been set.
func (o *Service) HasRefreshTokenIdempotent() bool {
	if o != nil && !IsNil(o.RefreshTokenIdempotent) {
		return true
	}

	return false
}

// SetRefreshTokenIdempotent gets a reference to the given bool and assigns it to the RefreshTokenIdempotent field.
func (o *Service) SetRefreshTokenIdempotent(v bool) {
	o.RefreshTokenIdempotent = &v
}

// GetSignedJwksUri returns the SignedJwksUri field value if set, zero value otherwise.
func (o *Service) GetSignedJwksUri() string {
	if o == nil || IsNil(o.SignedJwksUri) {
		var ret string
		return ret
	}
	return *o.SignedJwksUri
}

// GetSignedJwksUriOk returns a tuple with the SignedJwksUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSignedJwksUriOk() (*string, bool) {
	if o == nil || IsNil(o.SignedJwksUri) {
		return nil, false
	}
	return o.SignedJwksUri, true
}

// HasSignedJwksUri returns a boolean if a field has been set.
func (o *Service) HasSignedJwksUri() bool {
	if o != nil && !IsNil(o.SignedJwksUri) {
		return true
	}

	return false
}

// SetSignedJwksUri gets a reference to the given string and assigns it to the SignedJwksUri field.
func (o *Service) SetSignedJwksUri(v string) {
	o.SignedJwksUri = &v
}

// GetSupportedAttachments returns the SupportedAttachments field value if set, zero value otherwise.
func (o *Service) GetSupportedAttachments() []AttachmentType {
	if o == nil || IsNil(o.SupportedAttachments) {
		var ret []AttachmentType
		return ret
	}
	return o.SupportedAttachments
}

// GetSupportedAttachmentsOk returns a tuple with the SupportedAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedAttachmentsOk() ([]AttachmentType, bool) {
	if o == nil || IsNil(o.SupportedAttachments) {
		return nil, false
	}
	return o.SupportedAttachments, true
}

// HasSupportedAttachments returns a boolean if a field has been set.
func (o *Service) HasSupportedAttachments() bool {
	if o != nil && !IsNil(o.SupportedAttachments) {
		return true
	}

	return false
}

// SetSupportedAttachments gets a reference to the given []AttachmentType and assigns it to the SupportedAttachments field.
func (o *Service) SetSupportedAttachments(v []AttachmentType) {
	o.SupportedAttachments = v
}

// GetSupportedDigestAlgorithms returns the SupportedDigestAlgorithms field value if set, zero value otherwise.
func (o *Service) GetSupportedDigestAlgorithms() []string {
	if o == nil || IsNil(o.SupportedDigestAlgorithms) {
		var ret []string
		return ret
	}
	return o.SupportedDigestAlgorithms
}

// GetSupportedDigestAlgorithmsOk returns a tuple with the SupportedDigestAlgorithms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedDigestAlgorithmsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedDigestAlgorithms) {
		return nil, false
	}
	return o.SupportedDigestAlgorithms, true
}

// HasSupportedDigestAlgorithms returns a boolean if a field has been set.
func (o *Service) HasSupportedDigestAlgorithms() bool {
	if o != nil && !IsNil(o.SupportedDigestAlgorithms) {
		return true
	}

	return false
}

// SetSupportedDigestAlgorithms gets a reference to the given []string and assigns it to the SupportedDigestAlgorithms field.
func (o *Service) SetSupportedDigestAlgorithms(v []string) {
	o.SupportedDigestAlgorithms = v
}

// GetSupportedDocuments returns the SupportedDocuments field value if set, zero value otherwise.
func (o *Service) GetSupportedDocuments() []string {
	if o == nil || IsNil(o.SupportedDocuments) {
		var ret []string
		return ret
	}
	return o.SupportedDocuments
}

// GetSupportedDocumentsOk returns a tuple with the SupportedDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedDocumentsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedDocuments) {
		return nil, false
	}
	return o.SupportedDocuments, true
}

// HasSupportedDocuments returns a boolean if a field has been set.
func (o *Service) HasSupportedDocuments() bool {
	if o != nil && !IsNil(o.SupportedDocuments) {
		return true
	}

	return false
}

// SetSupportedDocuments gets a reference to the given []string and assigns it to the SupportedDocuments field.
func (o *Service) SetSupportedDocuments(v []string) {
	o.SupportedDocuments = v
}

// GetSupportedDocumentsMethods returns the SupportedDocumentsMethods field value if set, zero value otherwise.
func (o *Service) GetSupportedDocumentsMethods() []string {
	if o == nil || IsNil(o.SupportedDocumentsMethods) {
		var ret []string
		return ret
	}
	return o.SupportedDocumentsMethods
}

// GetSupportedDocumentsMethodsOk returns a tuple with the SupportedDocumentsMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedDocumentsMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedDocumentsMethods) {
		return nil, false
	}
	return o.SupportedDocumentsMethods, true
}

// HasSupportedDocumentsMethods returns a boolean if a field has been set.
func (o *Service) HasSupportedDocumentsMethods() bool {
	if o != nil && !IsNil(o.SupportedDocumentsMethods) {
		return true
	}

	return false
}

// SetSupportedDocumentsMethods gets a reference to the given []string and assigns it to the SupportedDocumentsMethods field.
func (o *Service) SetSupportedDocumentsMethods(v []string) {
	o.SupportedDocumentsMethods = v
}

// GetSupportedDocumentsValidationMethods returns the SupportedDocumentsValidationMethods field value if set, zero value otherwise.
func (o *Service) GetSupportedDocumentsValidationMethods() []string {
	if o == nil || IsNil(o.SupportedDocumentsValidationMethods) {
		var ret []string
		return ret
	}
	return o.SupportedDocumentsValidationMethods
}

// GetSupportedDocumentsValidationMethodsOk returns a tuple with the SupportedDocumentsValidationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedDocumentsValidationMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedDocumentsValidationMethods) {
		return nil, false
	}
	return o.SupportedDocumentsValidationMethods, true
}

// HasSupportedDocumentsValidationMethods returns a boolean if a field has been set.
func (o *Service) HasSupportedDocumentsValidationMethods() bool {
	if o != nil && !IsNil(o.SupportedDocumentsValidationMethods) {
		return true
	}

	return false
}

// SetSupportedDocumentsValidationMethods gets a reference to the given []string and assigns it to the SupportedDocumentsValidationMethods field.
func (o *Service) SetSupportedDocumentsValidationMethods(v []string) {
	o.SupportedDocumentsValidationMethods = v
}

// GetSupportedDocumentsVerificationMethods returns the SupportedDocumentsVerificationMethods field value if set, zero value otherwise.
func (o *Service) GetSupportedDocumentsVerificationMethods() []string {
	if o == nil || IsNil(o.SupportedDocumentsVerificationMethods) {
		var ret []string
		return ret
	}
	return o.SupportedDocumentsVerificationMethods
}

// GetSupportedDocumentsVerificationMethodsOk returns a tuple with the SupportedDocumentsVerificationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedDocumentsVerificationMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedDocumentsVerificationMethods) {
		return nil, false
	}
	return o.SupportedDocumentsVerificationMethods, true
}

// HasSupportedDocumentsVerificationMethods returns a boolean if a field has been set.
func (o *Service) HasSupportedDocumentsVerificationMethods() bool {
	if o != nil && !IsNil(o.SupportedDocumentsVerificationMethods) {
		return true
	}

	return false
}

// SetSupportedDocumentsVerificationMethods gets a reference to the given []string and assigns it to the SupportedDocumentsVerificationMethods field.
func (o *Service) SetSupportedDocumentsVerificationMethods(v []string) {
	o.SupportedDocumentsVerificationMethods = v
}

// GetSupportedElectronicRecords returns the SupportedElectronicRecords field value if set, zero value otherwise.
func (o *Service) GetSupportedElectronicRecords() []string {
	if o == nil || IsNil(o.SupportedElectronicRecords) {
		var ret []string
		return ret
	}
	return o.SupportedElectronicRecords
}

// GetSupportedElectronicRecordsOk returns a tuple with the SupportedElectronicRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedElectronicRecordsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedElectronicRecords) {
		return nil, false
	}
	return o.SupportedElectronicRecords, true
}

// HasSupportedElectronicRecords returns a boolean if a field has been set.
func (o *Service) HasSupportedElectronicRecords() bool {
	if o != nil && !IsNil(o.SupportedElectronicRecords) {
		return true
	}

	return false
}

// SetSupportedElectronicRecords gets a reference to the given []string and assigns it to the SupportedElectronicRecords field.
func (o *Service) SetSupportedElectronicRecords(v []string) {
	o.SupportedElectronicRecords = v
}

// GetSupportedClientRegistrationTypes returns the SupportedClientRegistrationTypes field value if set, zero value otherwise.
func (o *Service) GetSupportedClientRegistrationTypes() []ClientRegistrationType {
	if o == nil || IsNil(o.SupportedClientRegistrationTypes) {
		var ret []ClientRegistrationType
		return ret
	}
	return o.SupportedClientRegistrationTypes
}

// GetSupportedClientRegistrationTypesOk returns a tuple with the SupportedClientRegistrationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedClientRegistrationTypesOk() ([]ClientRegistrationType, bool) {
	if o == nil || IsNil(o.SupportedClientRegistrationTypes) {
		return nil, false
	}
	return o.SupportedClientRegistrationTypes, true
}

// HasSupportedClientRegistrationTypes returns a boolean if a field has been set.
func (o *Service) HasSupportedClientRegistrationTypes() bool {
	if o != nil && !IsNil(o.SupportedClientRegistrationTypes) {
		return true
	}

	return false
}

// SetSupportedClientRegistrationTypes gets a reference to the given []ClientRegistrationType and assigns it to the SupportedClientRegistrationTypes field.
func (o *Service) SetSupportedClientRegistrationTypes(v []ClientRegistrationType) {
	o.SupportedClientRegistrationTypes = v
}

// GetTokenExchangeByIdentifiableClientsOnly returns the TokenExchangeByIdentifiableClientsOnly field value if set, zero value otherwise.
func (o *Service) GetTokenExchangeByIdentifiableClientsOnly() bool {
	if o == nil || IsNil(o.TokenExchangeByIdentifiableClientsOnly) {
		var ret bool
		return ret
	}
	return *o.TokenExchangeByIdentifiableClientsOnly
}

// GetTokenExchangeByIdentifiableClientsOnlyOk returns a tuple with the TokenExchangeByIdentifiableClientsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetTokenExchangeByIdentifiableClientsOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.TokenExchangeByIdentifiableClientsOnly) {
		return nil, false
	}
	return o.TokenExchangeByIdentifiableClientsOnly, true
}

// HasTokenExchangeByIdentifiableClientsOnly returns a boolean if a field has been set.
func (o *Service) HasTokenExchangeByIdentifiableClientsOnly() bool {
	if o != nil && !IsNil(o.TokenExchangeByIdentifiableClientsOnly) {
		return true
	}

	return false
}

// SetTokenExchangeByIdentifiableClientsOnly gets a reference to the given bool and assigns it to the TokenExchangeByIdentifiableClientsOnly field.
func (o *Service) SetTokenExchangeByIdentifiableClientsOnly(v bool) {
	o.TokenExchangeByIdentifiableClientsOnly = &v
}

// GetTokenExchangeByConfidentialClientsOnly returns the TokenExchangeByConfidentialClientsOnly field value if set, zero value otherwise.
func (o *Service) GetTokenExchangeByConfidentialClientsOnly() bool {
	if o == nil || IsNil(o.TokenExchangeByConfidentialClientsOnly) {
		var ret bool
		return ret
	}
	return *o.TokenExchangeByConfidentialClientsOnly
}

// GetTokenExchangeByConfidentialClientsOnlyOk returns a tuple with the TokenExchangeByConfidentialClientsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetTokenExchangeByConfidentialClientsOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.TokenExchangeByConfidentialClientsOnly) {
		return nil, false
	}
	return o.TokenExchangeByConfidentialClientsOnly, true
}

// HasTokenExchangeByConfidentialClientsOnly returns a boolean if a field has been set.
func (o *Service) HasTokenExchangeByConfidentialClientsOnly() bool {
	if o != nil && !IsNil(o.TokenExchangeByConfidentialClientsOnly) {
		return true
	}

	return false
}

// SetTokenExchangeByConfidentialClientsOnly gets a reference to the given bool and assigns it to the TokenExchangeByConfidentialClientsOnly field.
func (o *Service) SetTokenExchangeByConfidentialClientsOnly(v bool) {
	o.TokenExchangeByConfidentialClientsOnly = &v
}

// GetTokenExchangeByPermittedClientsOnly returns the TokenExchangeByPermittedClientsOnly field value if set, zero value otherwise.
func (o *Service) GetTokenExchangeByPermittedClientsOnly() bool {
	if o == nil || IsNil(o.TokenExchangeByPermittedClientsOnly) {
		var ret bool
		return ret
	}
	return *o.TokenExchangeByPermittedClientsOnly
}

// GetTokenExchangeByPermittedClientsOnlyOk returns a tuple with the TokenExchangeByPermittedClientsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetTokenExchangeByPermittedClientsOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.TokenExchangeByPermittedClientsOnly) {
		return nil, false
	}
	return o.TokenExchangeByPermittedClientsOnly, true
}

// HasTokenExchangeByPermittedClientsOnly returns a boolean if a field has been set.
func (o *Service) HasTokenExchangeByPermittedClientsOnly() bool {
	if o != nil && !IsNil(o.TokenExchangeByPermittedClientsOnly) {
		return true
	}

	return false
}

// SetTokenExchangeByPermittedClientsOnly gets a reference to the given bool and assigns it to the TokenExchangeByPermittedClientsOnly field.
func (o *Service) SetTokenExchangeByPermittedClientsOnly(v bool) {
	o.TokenExchangeByPermittedClientsOnly = &v
}

// GetTokenExchangeEncryptedJwtRejected returns the TokenExchangeEncryptedJwtRejected field value if set, zero value otherwise.
func (o *Service) GetTokenExchangeEncryptedJwtRejected() bool {
	if o == nil || IsNil(o.TokenExchangeEncryptedJwtRejected) {
		var ret bool
		return ret
	}
	return *o.TokenExchangeEncryptedJwtRejected
}

// GetTokenExchangeEncryptedJwtRejectedOk returns a tuple with the TokenExchangeEncryptedJwtRejected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetTokenExchangeEncryptedJwtRejectedOk() (*bool, bool) {
	if o == nil || IsNil(o.TokenExchangeEncryptedJwtRejected) {
		return nil, false
	}
	return o.TokenExchangeEncryptedJwtRejected, true
}

// HasTokenExchangeEncryptedJwtRejected returns a boolean if a field has been set.
func (o *Service) HasTokenExchangeEncryptedJwtRejected() bool {
	if o != nil && !IsNil(o.TokenExchangeEncryptedJwtRejected) {
		return true
	}

	return false
}

// SetTokenExchangeEncryptedJwtRejected gets a reference to the given bool and assigns it to the TokenExchangeEncryptedJwtRejected field.
func (o *Service) SetTokenExchangeEncryptedJwtRejected(v bool) {
	o.TokenExchangeEncryptedJwtRejected = &v
}

// GetTokenExchangeUnsignedJwtRejected returns the TokenExchangeUnsignedJwtRejected field value if set, zero value otherwise.
func (o *Service) GetTokenExchangeUnsignedJwtRejected() bool {
	if o == nil || IsNil(o.TokenExchangeUnsignedJwtRejected) {
		var ret bool
		return ret
	}
	return *o.TokenExchangeUnsignedJwtRejected
}

// GetTokenExchangeUnsignedJwtRejectedOk returns a tuple with the TokenExchangeUnsignedJwtRejected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetTokenExchangeUnsignedJwtRejectedOk() (*bool, bool) {
	if o == nil || IsNil(o.TokenExchangeUnsignedJwtRejected) {
		return nil, false
	}
	return o.TokenExchangeUnsignedJwtRejected, true
}

// HasTokenExchangeUnsignedJwtRejected returns a boolean if a field has been set.
func (o *Service) HasTokenExchangeUnsignedJwtRejected() bool {
	if o != nil && !IsNil(o.TokenExchangeUnsignedJwtRejected) {
		return true
	}

	return false
}

// SetTokenExchangeUnsignedJwtRejected gets a reference to the given bool and assigns it to the TokenExchangeUnsignedJwtRejected field.
func (o *Service) SetTokenExchangeUnsignedJwtRejected(v bool) {
	o.TokenExchangeUnsignedJwtRejected = &v
}

// GetJwtGrantByIdentifiableClientsOnly returns the JwtGrantByIdentifiableClientsOnly field value if set, zero value otherwise.
func (o *Service) GetJwtGrantByIdentifiableClientsOnly() bool {
	if o == nil || IsNil(o.JwtGrantByIdentifiableClientsOnly) {
		var ret bool
		return ret
	}
	return *o.JwtGrantByIdentifiableClientsOnly
}

// GetJwtGrantByIdentifiableClientsOnlyOk returns a tuple with the JwtGrantByIdentifiableClientsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetJwtGrantByIdentifiableClientsOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.JwtGrantByIdentifiableClientsOnly) {
		return nil, false
	}
	return o.JwtGrantByIdentifiableClientsOnly, true
}

// HasJwtGrantByIdentifiableClientsOnly returns a boolean if a field has been set.
func (o *Service) HasJwtGrantByIdentifiableClientsOnly() bool {
	if o != nil && !IsNil(o.JwtGrantByIdentifiableClientsOnly) {
		return true
	}

	return false
}

// SetJwtGrantByIdentifiableClientsOnly gets a reference to the given bool and assigns it to the JwtGrantByIdentifiableClientsOnly field.
func (o *Service) SetJwtGrantByIdentifiableClientsOnly(v bool) {
	o.JwtGrantByIdentifiableClientsOnly = &v
}

// GetJwtGrantEncryptedJwtRejected returns the JwtGrantEncryptedJwtRejected field value if set, zero value otherwise.
func (o *Service) GetJwtGrantEncryptedJwtRejected() bool {
	if o == nil || IsNil(o.JwtGrantEncryptedJwtRejected) {
		var ret bool
		return ret
	}
	return *o.JwtGrantEncryptedJwtRejected
}

// GetJwtGrantEncryptedJwtRejectedOk returns a tuple with the JwtGrantEncryptedJwtRejected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetJwtGrantEncryptedJwtRejectedOk() (*bool, bool) {
	if o == nil || IsNil(o.JwtGrantEncryptedJwtRejected) {
		return nil, false
	}
	return o.JwtGrantEncryptedJwtRejected, true
}

// HasJwtGrantEncryptedJwtRejected returns a boolean if a field has been set.
func (o *Service) HasJwtGrantEncryptedJwtRejected() bool {
	if o != nil && !IsNil(o.JwtGrantEncryptedJwtRejected) {
		return true
	}

	return false
}

// SetJwtGrantEncryptedJwtRejected gets a reference to the given bool and assigns it to the JwtGrantEncryptedJwtRejected field.
func (o *Service) SetJwtGrantEncryptedJwtRejected(v bool) {
	o.JwtGrantEncryptedJwtRejected = &v
}

// GetJwtGrantUnsignedJwtRejected returns the JwtGrantUnsignedJwtRejected field value if set, zero value otherwise.
func (o *Service) GetJwtGrantUnsignedJwtRejected() bool {
	if o == nil || IsNil(o.JwtGrantUnsignedJwtRejected) {
		var ret bool
		return ret
	}
	return *o.JwtGrantUnsignedJwtRejected
}

// GetJwtGrantUnsignedJwtRejectedOk returns a tuple with the JwtGrantUnsignedJwtRejected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetJwtGrantUnsignedJwtRejectedOk() (*bool, bool) {
	if o == nil || IsNil(o.JwtGrantUnsignedJwtRejected) {
		return nil, false
	}
	return o.JwtGrantUnsignedJwtRejected, true
}

// HasJwtGrantUnsignedJwtRejected returns a boolean if a field has been set.
func (o *Service) HasJwtGrantUnsignedJwtRejected() bool {
	if o != nil && !IsNil(o.JwtGrantUnsignedJwtRejected) {
		return true
	}

	return false
}

// SetJwtGrantUnsignedJwtRejected gets a reference to the given bool and assigns it to the JwtGrantUnsignedJwtRejected field.
func (o *Service) SetJwtGrantUnsignedJwtRejected(v bool) {
	o.JwtGrantUnsignedJwtRejected = &v
}

// GetDcrDuplicateSoftwareIdBlocked returns the DcrDuplicateSoftwareIdBlocked field value if set, zero value otherwise.
func (o *Service) GetDcrDuplicateSoftwareIdBlocked() bool {
	if o == nil || IsNil(o.DcrDuplicateSoftwareIdBlocked) {
		var ret bool
		return ret
	}
	return *o.DcrDuplicateSoftwareIdBlocked
}

// GetDcrDuplicateSoftwareIdBlockedOk returns a tuple with the DcrDuplicateSoftwareIdBlocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDcrDuplicateSoftwareIdBlockedOk() (*bool, bool) {
	if o == nil || IsNil(o.DcrDuplicateSoftwareIdBlocked) {
		return nil, false
	}
	return o.DcrDuplicateSoftwareIdBlocked, true
}

// HasDcrDuplicateSoftwareIdBlocked returns a boolean if a field has been set.
func (o *Service) HasDcrDuplicateSoftwareIdBlocked() bool {
	if o != nil && !IsNil(o.DcrDuplicateSoftwareIdBlocked) {
		return true
	}

	return false
}

// SetDcrDuplicateSoftwareIdBlocked gets a reference to the given bool and assigns it to the DcrDuplicateSoftwareIdBlocked field.
func (o *Service) SetDcrDuplicateSoftwareIdBlocked(v bool) {
	o.DcrDuplicateSoftwareIdBlocked = &v
}

// GetTrustAnchors returns the TrustAnchors field value if set, zero value otherwise.
func (o *Service) GetTrustAnchors() []TrustAnchor {
	if o == nil || IsNil(o.TrustAnchors) {
		var ret []TrustAnchor
		return ret
	}
	return o.TrustAnchors
}

// GetTrustAnchorsOk returns a tuple with the TrustAnchors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetTrustAnchorsOk() ([]TrustAnchor, bool) {
	if o == nil || IsNil(o.TrustAnchors) {
		return nil, false
	}
	return o.TrustAnchors, true
}

// HasTrustAnchors returns a boolean if a field has been set.
func (o *Service) HasTrustAnchors() bool {
	if o != nil && !IsNil(o.TrustAnchors) {
		return true
	}

	return false
}

// SetTrustAnchors gets a reference to the given []TrustAnchor and assigns it to the TrustAnchors field.
func (o *Service) SetTrustAnchors(v []TrustAnchor) {
	o.TrustAnchors = v
}

// GetOpenidDroppedOnRefreshWithoutOfflineAccess returns the OpenidDroppedOnRefreshWithoutOfflineAccess field value if set, zero value otherwise.
func (o *Service) GetOpenidDroppedOnRefreshWithoutOfflineAccess() bool {
	if o == nil || IsNil(o.OpenidDroppedOnRefreshWithoutOfflineAccess) {
		var ret bool
		return ret
	}
	return *o.OpenidDroppedOnRefreshWithoutOfflineAccess
}

// GetOpenidDroppedOnRefreshWithoutOfflineAccessOk returns a tuple with the OpenidDroppedOnRefreshWithoutOfflineAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetOpenidDroppedOnRefreshWithoutOfflineAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.OpenidDroppedOnRefreshWithoutOfflineAccess) {
		return nil, false
	}
	return o.OpenidDroppedOnRefreshWithoutOfflineAccess, true
}

// HasOpenidDroppedOnRefreshWithoutOfflineAccess returns a boolean if a field has been set.
func (o *Service) HasOpenidDroppedOnRefreshWithoutOfflineAccess() bool {
	if o != nil && !IsNil(o.OpenidDroppedOnRefreshWithoutOfflineAccess) {
		return true
	}

	return false
}

// SetOpenidDroppedOnRefreshWithoutOfflineAccess gets a reference to the given bool and assigns it to the OpenidDroppedOnRefreshWithoutOfflineAccess field.
func (o *Service) SetOpenidDroppedOnRefreshWithoutOfflineAccess(v bool) {
	o.OpenidDroppedOnRefreshWithoutOfflineAccess = &v
}

// GetSupportedDocumentsCheckMethods returns the SupportedDocumentsCheckMethods field value if set, zero value otherwise.
func (o *Service) GetSupportedDocumentsCheckMethods() []string {
	if o == nil || IsNil(o.SupportedDocumentsCheckMethods) {
		var ret []string
		return ret
	}
	return o.SupportedDocumentsCheckMethods
}

// GetSupportedDocumentsCheckMethodsOk returns a tuple with the SupportedDocumentsCheckMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedDocumentsCheckMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedDocumentsCheckMethods) {
		return nil, false
	}
	return o.SupportedDocumentsCheckMethods, true
}

// HasSupportedDocumentsCheckMethods returns a boolean if a field has been set.
func (o *Service) HasSupportedDocumentsCheckMethods() bool {
	if o != nil && !IsNil(o.SupportedDocumentsCheckMethods) {
		return true
	}

	return false
}

// SetSupportedDocumentsCheckMethods gets a reference to the given []string and assigns it to the SupportedDocumentsCheckMethods field.
func (o *Service) SetSupportedDocumentsCheckMethods(v []string) {
	o.SupportedDocumentsCheckMethods = v
}

// GetRsResponseSigned returns the RsResponseSigned field value if set, zero value otherwise.
func (o *Service) GetRsResponseSigned() bool {
	if o == nil || IsNil(o.RsResponseSigned) {
		var ret bool
		return ret
	}
	return *o.RsResponseSigned
}

// GetRsResponseSignedOk returns a tuple with the RsResponseSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetRsResponseSignedOk() (*bool, bool) {
	if o == nil || IsNil(o.RsResponseSigned) {
		return nil, false
	}
	return o.RsResponseSigned, true
}

// HasRsResponseSigned returns a boolean if a field has been set.
func (o *Service) HasRsResponseSigned() bool {
	if o != nil && !IsNil(o.RsResponseSigned) {
		return true
	}

	return false
}

// SetRsResponseSigned gets a reference to the given bool and assigns it to the RsResponseSigned field.
func (o *Service) SetRsResponseSigned(v bool) {
	o.RsResponseSigned = &v
}

// GetCnonceDuration returns the CnonceDuration field value if set, zero value otherwise.
func (o *Service) GetCnonceDuration() int64 {
	if o == nil || IsNil(o.CnonceDuration) {
		var ret int64
		return ret
	}
	return *o.CnonceDuration
}

// GetCnonceDurationOk returns a tuple with the CnonceDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetCnonceDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.CnonceDuration) {
		return nil, false
	}
	return o.CnonceDuration, true
}

// HasCnonceDuration returns a boolean if a field has been set.
func (o *Service) HasCnonceDuration() bool {
	if o != nil && !IsNil(o.CnonceDuration) {
		return true
	}

	return false
}

// SetCnonceDuration gets a reference to the given int64 and assigns it to the CnonceDuration field.
func (o *Service) SetCnonceDuration(v int64) {
	o.CnonceDuration = &v
}

// GetDpopNonceRequired returns the DpopNonceRequired field value if set, zero value otherwise.
func (o *Service) GetDpopNonceRequired() bool {
	if o == nil || IsNil(o.DpopNonceRequired) {
		var ret bool
		return ret
	}
	return *o.DpopNonceRequired
}

// GetDpopNonceRequiredOk returns a tuple with the DpopNonceRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDpopNonceRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.DpopNonceRequired) {
		return nil, false
	}
	return o.DpopNonceRequired, true
}

// HasDpopNonceRequired returns a boolean if a field has been set.
func (o *Service) HasDpopNonceRequired() bool {
	if o != nil && !IsNil(o.DpopNonceRequired) {
		return true
	}

	return false
}

// SetDpopNonceRequired gets a reference to the given bool and assigns it to the DpopNonceRequired field.
func (o *Service) SetDpopNonceRequired(v bool) {
	o.DpopNonceRequired = &v
}

// GetVerifiableCredentialsEnabled returns the VerifiableCredentialsEnabled field value if set, zero value otherwise.
func (o *Service) GetVerifiableCredentialsEnabled() bool {
	if o == nil || IsNil(o.VerifiableCredentialsEnabled) {
		var ret bool
		return ret
	}
	return *o.VerifiableCredentialsEnabled
}

// GetVerifiableCredentialsEnabledOk returns a tuple with the VerifiableCredentialsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetVerifiableCredentialsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifiableCredentialsEnabled) {
		return nil, false
	}
	return o.VerifiableCredentialsEnabled, true
}

// HasVerifiableCredentialsEnabled returns a boolean if a field has been set.
func (o *Service) HasVerifiableCredentialsEnabled() bool {
	if o != nil && !IsNil(o.VerifiableCredentialsEnabled) {
		return true
	}

	return false
}

// SetVerifiableCredentialsEnabled gets a reference to the given bool and assigns it to the VerifiableCredentialsEnabled field.
func (o *Service) SetVerifiableCredentialsEnabled(v bool) {
	o.VerifiableCredentialsEnabled = &v
}

// GetCredentialJwksUri returns the CredentialJwksUri field value if set, zero value otherwise.
func (o *Service) GetCredentialJwksUri() string {
	if o == nil || IsNil(o.CredentialJwksUri) {
		var ret string
		return ret
	}
	return *o.CredentialJwksUri
}

// GetCredentialJwksUriOk returns a tuple with the CredentialJwksUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetCredentialJwksUriOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialJwksUri) {
		return nil, false
	}
	return o.CredentialJwksUri, true
}

// HasCredentialJwksUri returns a boolean if a field has been set.
func (o *Service) HasCredentialJwksUri() bool {
	if o != nil && !IsNil(o.CredentialJwksUri) {
		return true
	}

	return false
}

// SetCredentialJwksUri gets a reference to the given string and assigns it to the CredentialJwksUri field.
func (o *Service) SetCredentialJwksUri(v string) {
	o.CredentialJwksUri = &v
}

// GetCredentialOfferDuration returns the CredentialOfferDuration field value if set, zero value otherwise.
func (o *Service) GetCredentialOfferDuration() int64 {
	if o == nil || IsNil(o.CredentialOfferDuration) {
		var ret int64
		return ret
	}
	return *o.CredentialOfferDuration
}

// GetCredentialOfferDurationOk returns a tuple with the CredentialOfferDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetCredentialOfferDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.CredentialOfferDuration) {
		return nil, false
	}
	return o.CredentialOfferDuration, true
}

// HasCredentialOfferDuration returns a boolean if a field has been set.
func (o *Service) HasCredentialOfferDuration() bool {
	if o != nil && !IsNil(o.CredentialOfferDuration) {
		return true
	}

	return false
}

// SetCredentialOfferDuration gets a reference to the given int64 and assigns it to the CredentialOfferDuration field.
func (o *Service) SetCredentialOfferDuration(v int64) {
	o.CredentialOfferDuration = &v
}

// GetDpopNonceDuration returns the DpopNonceDuration field value if set, zero value otherwise.
func (o *Service) GetDpopNonceDuration() int64 {
	if o == nil || IsNil(o.DpopNonceDuration) {
		var ret int64
		return ret
	}
	return *o.DpopNonceDuration
}

// GetDpopNonceDurationOk returns a tuple with the DpopNonceDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDpopNonceDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.DpopNonceDuration) {
		return nil, false
	}
	return o.DpopNonceDuration, true
}

// HasDpopNonceDuration returns a boolean if a field has been set.
func (o *Service) HasDpopNonceDuration() bool {
	if o != nil && !IsNil(o.DpopNonceDuration) {
		return true
	}

	return false
}

// SetDpopNonceDuration gets a reference to the given int64 and assigns it to the DpopNonceDuration field.
func (o *Service) SetDpopNonceDuration(v int64) {
	o.DpopNonceDuration = &v
}

// GetPreAuthorizedGrantAnonymousAccessSupported returns the PreAuthorizedGrantAnonymousAccessSupported field value if set, zero value otherwise.
func (o *Service) GetPreAuthorizedGrantAnonymousAccessSupported() bool {
	if o == nil || IsNil(o.PreAuthorizedGrantAnonymousAccessSupported) {
		var ret bool
		return ret
	}
	return *o.PreAuthorizedGrantAnonymousAccessSupported
}

// GetPreAuthorizedGrantAnonymousAccessSupportedOk returns a tuple with the PreAuthorizedGrantAnonymousAccessSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetPreAuthorizedGrantAnonymousAccessSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.PreAuthorizedGrantAnonymousAccessSupported) {
		return nil, false
	}
	return o.PreAuthorizedGrantAnonymousAccessSupported, true
}

// HasPreAuthorizedGrantAnonymousAccessSupported returns a boolean if a field has been set.
func (o *Service) HasPreAuthorizedGrantAnonymousAccessSupported() bool {
	if o != nil && !IsNil(o.PreAuthorizedGrantAnonymousAccessSupported) {
		return true
	}

	return false
}

// SetPreAuthorizedGrantAnonymousAccessSupported gets a reference to the given bool and assigns it to the PreAuthorizedGrantAnonymousAccessSupported field.
func (o *Service) SetPreAuthorizedGrantAnonymousAccessSupported(v bool) {
	o.PreAuthorizedGrantAnonymousAccessSupported = &v
}

// GetCredentialTransactionDuration returns the CredentialTransactionDuration field value if set, zero value otherwise.
func (o *Service) GetCredentialTransactionDuration() int64 {
	if o == nil || IsNil(o.CredentialTransactionDuration) {
		var ret int64
		return ret
	}
	return *o.CredentialTransactionDuration
}

// GetCredentialTransactionDurationOk returns a tuple with the CredentialTransactionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetCredentialTransactionDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.CredentialTransactionDuration) {
		return nil, false
	}
	return o.CredentialTransactionDuration, true
}

// HasCredentialTransactionDuration returns a boolean if a field has been set.
func (o *Service) HasCredentialTransactionDuration() bool {
	if o != nil && !IsNil(o.CredentialTransactionDuration) {
		return true
	}

	return false
}

// SetCredentialTransactionDuration gets a reference to the given int64 and assigns it to the CredentialTransactionDuration field.
func (o *Service) SetCredentialTransactionDuration(v int64) {
	o.CredentialTransactionDuration = &v
}

// GetIntrospectionSignatureKeyId returns the IntrospectionSignatureKeyId field value if set, zero value otherwise.
func (o *Service) GetIntrospectionSignatureKeyId() string {
	if o == nil || IsNil(o.IntrospectionSignatureKeyId) {
		var ret string
		return ret
	}
	return *o.IntrospectionSignatureKeyId
}

// GetIntrospectionSignatureKeyIdOk returns a tuple with the IntrospectionSignatureKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetIntrospectionSignatureKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.IntrospectionSignatureKeyId) {
		return nil, false
	}
	return o.IntrospectionSignatureKeyId, true
}

// HasIntrospectionSignatureKeyId returns a boolean if a field has been set.
func (o *Service) HasIntrospectionSignatureKeyId() bool {
	if o != nil && !IsNil(o.IntrospectionSignatureKeyId) {
		return true
	}

	return false
}

// SetIntrospectionSignatureKeyId gets a reference to the given string and assigns it to the IntrospectionSignatureKeyId field.
func (o *Service) SetIntrospectionSignatureKeyId(v string) {
	o.IntrospectionSignatureKeyId = &v
}

// GetResourceSignatureKeyId returns the ResourceSignatureKeyId field value if set, zero value otherwise.
func (o *Service) GetResourceSignatureKeyId() string {
	if o == nil || IsNil(o.ResourceSignatureKeyId) {
		var ret string
		return ret
	}
	return *o.ResourceSignatureKeyId
}

// GetResourceSignatureKeyIdOk returns a tuple with the ResourceSignatureKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetResourceSignatureKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceSignatureKeyId) {
		return nil, false
	}
	return o.ResourceSignatureKeyId, true
}

// HasResourceSignatureKeyId returns a boolean if a field has been set.
func (o *Service) HasResourceSignatureKeyId() bool {
	if o != nil && !IsNil(o.ResourceSignatureKeyId) {
		return true
	}

	return false
}

// SetResourceSignatureKeyId gets a reference to the given string and assigns it to the ResourceSignatureKeyId field.
func (o *Service) SetResourceSignatureKeyId(v string) {
	o.ResourceSignatureKeyId = &v
}

// GetUserPinLength returns the UserPinLength field value if set, zero value otherwise.
func (o *Service) GetUserPinLength() int32 {
	if o == nil || IsNil(o.UserPinLength) {
		var ret int32
		return ret
	}
	return *o.UserPinLength
}

// GetUserPinLengthOk returns a tuple with the UserPinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetUserPinLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.UserPinLength) {
		return nil, false
	}
	return o.UserPinLength, true
}

// HasUserPinLength returns a boolean if a field has been set.
func (o *Service) HasUserPinLength() bool {
	if o != nil && !IsNil(o.UserPinLength) {
		return true
	}

	return false
}

// SetUserPinLength gets a reference to the given int32 and assigns it to the UserPinLength field.
func (o *Service) SetUserPinLength(v int32) {
	o.UserPinLength = &v
}

// GetSupportedPromptValues returns the SupportedPromptValues field value if set, zero value otherwise.
func (o *Service) GetSupportedPromptValues() []Prompt {
	if o == nil || IsNil(o.SupportedPromptValues) {
		var ret []Prompt
		return ret
	}
	return o.SupportedPromptValues
}

// GetSupportedPromptValuesOk returns a tuple with the SupportedPromptValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetSupportedPromptValuesOk() ([]Prompt, bool) {
	if o == nil || IsNil(o.SupportedPromptValues) {
		return nil, false
	}
	return o.SupportedPromptValues, true
}

// HasSupportedPromptValues returns a boolean if a field has been set.
func (o *Service) HasSupportedPromptValues() bool {
	if o != nil && !IsNil(o.SupportedPromptValues) {
		return true
	}

	return false
}

// SetSupportedPromptValues gets a reference to the given []Prompt and assigns it to the SupportedPromptValues field.
func (o *Service) SetSupportedPromptValues(v []Prompt) {
	o.SupportedPromptValues = v
}

// GetIdTokenReissuable returns the IdTokenReissuable field value if set, zero value otherwise.
func (o *Service) GetIdTokenReissuable() bool {
	if o == nil || IsNil(o.IdTokenReissuable) {
		var ret bool
		return ret
	}
	return *o.IdTokenReissuable
}

// GetIdTokenReissuableOk returns a tuple with the IdTokenReissuable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetIdTokenReissuableOk() (*bool, bool) {
	if o == nil || IsNil(o.IdTokenReissuable) {
		return nil, false
	}
	return o.IdTokenReissuable, true
}

// HasIdTokenReissuable returns a boolean if a field has been set.
func (o *Service) HasIdTokenReissuable() bool {
	if o != nil && !IsNil(o.IdTokenReissuable) {
		return true
	}

	return false
}

// SetIdTokenReissuable gets a reference to the given bool and assigns it to the IdTokenReissuable field.
func (o *Service) SetIdTokenReissuable(v bool) {
	o.IdTokenReissuable = &v
}

// GetCredentialJwks returns the CredentialJwks field value if set, zero value otherwise.
func (o *Service) GetCredentialJwks() string {
	if o == nil || IsNil(o.CredentialJwks) {
		var ret string
		return ret
	}
	return *o.CredentialJwks
}

// GetCredentialJwksOk returns a tuple with the CredentialJwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetCredentialJwksOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialJwks) {
		return nil, false
	}
	return o.CredentialJwks, true
}

// HasCredentialJwks returns a boolean if a field has been set.
func (o *Service) HasCredentialJwks() bool {
	if o != nil && !IsNil(o.CredentialJwks) {
		return true
	}

	return false
}

// SetCredentialJwks gets a reference to the given string and assigns it to the CredentialJwks field.
func (o *Service) SetCredentialJwks(v string) {
	o.CredentialJwks = &v
}

// GetFapiModes returns the FapiModes field value if set, zero value otherwise.
func (o *Service) GetFapiModes() []FapiMode {
	if o == nil || IsNil(o.FapiModes) {
		var ret []FapiMode
		return ret
	}
	return o.FapiModes
}

// GetFapiModesOk returns a tuple with the FapiModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetFapiModesOk() ([]FapiMode, bool) {
	if o == nil || IsNil(o.FapiModes) {
		return nil, false
	}
	return o.FapiModes, true
}

// HasFapiModes returns a boolean if a field has been set.
func (o *Service) HasFapiModes() bool {
	if o != nil && !IsNil(o.FapiModes) {
		return true
	}

	return false
}

// SetFapiModes gets a reference to the given []FapiMode and assigns it to the FapiModes field.
func (o *Service) SetFapiModes(v []FapiMode) {
	o.FapiModes = v
}

// GetCredentialDuration returns the CredentialDuration field value if set, zero value otherwise.
func (o *Service) GetCredentialDuration() int64 {
	if o == nil || IsNil(o.CredentialDuration) {
		var ret int64
		return ret
	}
	return *o.CredentialDuration
}

// GetCredentialDurationOk returns a tuple with the CredentialDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetCredentialDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.CredentialDuration) {
		return nil, false
	}
	return o.CredentialDuration, true
}

// HasCredentialDuration returns a boolean if a field has been set.
func (o *Service) HasCredentialDuration() bool {
	if o != nil && !IsNil(o.CredentialDuration) {
		return true
	}

	return false
}

// SetCredentialDuration gets a reference to the given int64 and assigns it to the CredentialDuration field.
func (o *Service) SetCredentialDuration(v int64) {
	o.CredentialDuration = &v
}

// GetCredentialIssuerMetadata returns the CredentialIssuerMetadata field value if set, zero value otherwise.
func (o *Service) GetCredentialIssuerMetadata() CredentialIssuerMetadata {
	if o == nil || IsNil(o.CredentialIssuerMetadata) {
		var ret CredentialIssuerMetadata
		return ret
	}
	return *o.CredentialIssuerMetadata
}

// GetCredentialIssuerMetadataOk returns a tuple with the CredentialIssuerMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetCredentialIssuerMetadataOk() (*CredentialIssuerMetadata, bool) {
	if o == nil || IsNil(o.CredentialIssuerMetadata) {
		return nil, false
	}
	return o.CredentialIssuerMetadata, true
}

// HasCredentialIssuerMetadata returns a boolean if a field has been set.
func (o *Service) HasCredentialIssuerMetadata() bool {
	if o != nil && !IsNil(o.CredentialIssuerMetadata) {
		return true
	}

	return false
}

// SetCredentialIssuerMetadata gets a reference to the given CredentialIssuerMetadata and assigns it to the CredentialIssuerMetadata field.
func (o *Service) SetCredentialIssuerMetadata(v CredentialIssuerMetadata) {
	o.CredentialIssuerMetadata = &v
}

// GetIdTokenAudType returns the IdTokenAudType field value if set, zero value otherwise.
func (o *Service) GetIdTokenAudType() string {
	if o == nil || IsNil(o.IdTokenAudType) {
		var ret string
		return ret
	}
	return *o.IdTokenAudType
}

// GetIdTokenAudTypeOk returns a tuple with the IdTokenAudType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetIdTokenAudTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IdTokenAudType) {
		return nil, false
	}
	return o.IdTokenAudType, true
}

// HasIdTokenAudType returns a boolean if a field has been set.
func (o *Service) HasIdTokenAudType() bool {
	if o != nil && !IsNil(o.IdTokenAudType) {
		return true
	}

	return false
}

// SetIdTokenAudType gets a reference to the given string and assigns it to the IdTokenAudType field.
func (o *Service) SetIdTokenAudType(v string) {
	o.IdTokenAudType = &v
}

func (o Service) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Service) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.ServiceName) {
		toSerialize["serviceName"] = o.ServiceName
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	if !IsNil(o.ClientIdAliasEnabled) {
		toSerialize["clientIdAliasEnabled"] = o.ClientIdAliasEnabled
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.ModifiedAt) {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if !IsNil(o.AuthenticationCallbackEndpoint) {
		toSerialize["authenticationCallbackEndpoint"] = o.AuthenticationCallbackEndpoint
	}
	if !IsNil(o.AuthenticationCallbackApiKey) {
		toSerialize["authenticationCallbackApiKey"] = o.AuthenticationCallbackApiKey
	}
	if !IsNil(o.AuthenticationCallbackApiSecret) {
		toSerialize["authenticationCallbackApiSecret"] = o.AuthenticationCallbackApiSecret
	}
	if !IsNil(o.SupportedAcrs) {
		toSerialize["supportedAcrs"] = o.SupportedAcrs
	}
	if !IsNil(o.SupportedGrantTypes) {
		toSerialize["supportedGrantTypes"] = o.SupportedGrantTypes
	}
	if !IsNil(o.SupportedResponseTypes) {
		toSerialize["supportedResponseTypes"] = o.SupportedResponseTypes
	}
	if !IsNil(o.SupportedAuthorizationDetailsTypes) {
		toSerialize["supportedAuthorizationDetailsTypes"] = o.SupportedAuthorizationDetailsTypes
	}
	if !IsNil(o.SupportedServiceProfiles) {
		toSerialize["supportedServiceProfiles"] = o.SupportedServiceProfiles
	}
	if !IsNil(o.ErrorDescriptionOmitted) {
		toSerialize["errorDescriptionOmitted"] = o.ErrorDescriptionOmitted
	}
	if !IsNil(o.ErrorUriOmitted) {
		toSerialize["errorUriOmitted"] = o.ErrorUriOmitted
	}
	if !IsNil(o.AuthorizationEndpoint) {
		toSerialize["authorizationEndpoint"] = o.AuthorizationEndpoint
	}
	if !IsNil(o.DirectAuthorizationEndpointEnabled) {
		toSerialize["directAuthorizationEndpointEnabled"] = o.DirectAuthorizationEndpointEnabled
	}
	if !IsNil(o.SupportedUiLocales) {
		toSerialize["supportedUiLocales"] = o.SupportedUiLocales
	}
	if !IsNil(o.SupportedDisplays) {
		toSerialize["supportedDisplays"] = o.SupportedDisplays
	}
	if !IsNil(o.PkceRequired) {
		toSerialize["pkceRequired"] = o.PkceRequired
	}
	if !IsNil(o.PkceS256Required) {
		toSerialize["pkceS256Required"] = o.PkceS256Required
	}
	if !IsNil(o.AuthorizationResponseDuration) {
		toSerialize["authorizationResponseDuration"] = o.AuthorizationResponseDuration
	}
	if !IsNil(o.TokenEndpoint) {
		toSerialize["tokenEndpoint"] = o.TokenEndpoint
	}
	if !IsNil(o.DirectTokenEndpointEnabled) {
		toSerialize["directTokenEndpointEnabled"] = o.DirectTokenEndpointEnabled
	}
	if !IsNil(o.SupportedTokenAuthMethods) {
		toSerialize["supportedTokenAuthMethods"] = o.SupportedTokenAuthMethods
	}
	if !IsNil(o.MissingClientIdAllowed) {
		toSerialize["missingClientIdAllowed"] = o.MissingClientIdAllowed
	}
	if !IsNil(o.RevocationEndpoint) {
		toSerialize["revocationEndpoint"] = o.RevocationEndpoint
	}
	if !IsNil(o.DirectRevocationEndpointEnabled) {
		toSerialize["directRevocationEndpointEnabled"] = o.DirectRevocationEndpointEnabled
	}
	if !IsNil(o.SupportedRevocationAuthMethods) {
		toSerialize["supportedRevocationAuthMethods"] = o.SupportedRevocationAuthMethods
	}
	if !IsNil(o.IntrospectionEndpoint) {
		toSerialize["introspectionEndpoint"] = o.IntrospectionEndpoint
	}
	if !IsNil(o.DirectIntrospectionEndpointEnabled) {
		toSerialize["directIntrospectionEndpointEnabled"] = o.DirectIntrospectionEndpointEnabled
	}
	if !IsNil(o.SupportedIntrospectionAuthMethods) {
		toSerialize["supportedIntrospectionAuthMethods"] = o.SupportedIntrospectionAuthMethods
	}
	if !IsNil(o.PushedAuthReqEndpoint) {
		toSerialize["pushedAuthReqEndpoint"] = o.PushedAuthReqEndpoint
	}
	if !IsNil(o.PushedAuthReqDuration) {
		toSerialize["pushedAuthReqDuration"] = o.PushedAuthReqDuration
	}
	if !IsNil(o.ParRequired) {
		toSerialize["parRequired"] = o.ParRequired
	}
	if !IsNil(o.RequestObjectRequired) {
		toSerialize["requestObjectRequired"] = o.RequestObjectRequired
	}
	if !IsNil(o.TraditionalRequestObjectProcessingApplied) {
		toSerialize["traditionalRequestObjectProcessingApplied"] = o.TraditionalRequestObjectProcessingApplied
	}
	if !IsNil(o.MutualTlsValidatePkiCertChain) {
		toSerialize["mutualTlsValidatePkiCertChain"] = o.MutualTlsValidatePkiCertChain
	}
	if !IsNil(o.TrustedRootCertificates) {
		toSerialize["trustedRootCertificates"] = o.TrustedRootCertificates
	}
	if !IsNil(o.MtlsEndpointAliases) {
		toSerialize["mtlsEndpointAliases"] = o.MtlsEndpointAliases
	}
	if !IsNil(o.AccessTokenType) {
		toSerialize["accessTokenType"] = o.AccessTokenType
	}
	if !IsNil(o.TlsClientCertificateBoundAccessTokens) {
		toSerialize["tlsClientCertificateBoundAccessTokens"] = o.TlsClientCertificateBoundAccessTokens
	}
	if !IsNil(o.AccessTokenDuration) {
		toSerialize["accessTokenDuration"] = o.AccessTokenDuration
	}
	if !IsNil(o.SingleAccessTokenPerSubject) {
		toSerialize["singleAccessTokenPerSubject"] = o.SingleAccessTokenPerSubject
	}
	if !IsNil(o.AccessTokenSignAlg) {
		toSerialize["accessTokenSignAlg"] = o.AccessTokenSignAlg
	}
	if !IsNil(o.AccessTokenSignatureKeyId) {
		toSerialize["accessTokenSignatureKeyId"] = o.AccessTokenSignatureKeyId
	}
	if !IsNil(o.RefreshTokenDuration) {
		toSerialize["refreshTokenDuration"] = o.RefreshTokenDuration
	}
	if !IsNil(o.RefreshTokenDurationKept) {
		toSerialize["refreshTokenDurationKept"] = o.RefreshTokenDurationKept
	}
	if !IsNil(o.RefreshTokenDurationReset) {
		toSerialize["refreshTokenDurationReset"] = o.RefreshTokenDurationReset
	}
	if !IsNil(o.RefreshTokenKept) {
		toSerialize["refreshTokenKept"] = o.RefreshTokenKept
	}
	if !IsNil(o.SupportedScopes) {
		toSerialize["supportedScopes"] = o.SupportedScopes
	}
	if !IsNil(o.ScopeRequired) {
		toSerialize["scopeRequired"] = o.ScopeRequired
	}
	if !IsNil(o.IdTokenDuration) {
		toSerialize["idTokenDuration"] = o.IdTokenDuration
	}
	if !IsNil(o.AllowableClockSkew) {
		toSerialize["allowableClockSkew"] = o.AllowableClockSkew
	}
	if !IsNil(o.SupportedClaimTypes) {
		toSerialize["supportedClaimTypes"] = o.SupportedClaimTypes
	}
	if !IsNil(o.SupportedClaimLocales) {
		toSerialize["supportedClaimLocales"] = o.SupportedClaimLocales
	}
	if !IsNil(o.SupportedClaims) {
		toSerialize["supportedClaims"] = o.SupportedClaims
	}
	if !IsNil(o.ClaimShortcutRestrictive) {
		toSerialize["claimShortcutRestrictive"] = o.ClaimShortcutRestrictive
	}
	if !IsNil(o.JwksUri) {
		toSerialize["jwksUri"] = o.JwksUri
	}
	if !IsNil(o.DirectJwksEndpointEnabled) {
		toSerialize["directJwksEndpointEnabled"] = o.DirectJwksEndpointEnabled
	}
	if !IsNil(o.Jwks) {
		toSerialize["jwks"] = o.Jwks
	}
	if !IsNil(o.IdTokenSignatureKeyId) {
		toSerialize["idTokenSignatureKeyId"] = o.IdTokenSignatureKeyId
	}
	if !IsNil(o.UserInfoSignatureKeyId) {
		toSerialize["userInfoSignatureKeyId"] = o.UserInfoSignatureKeyId
	}
	if !IsNil(o.AuthorizationSignatureKeyId) {
		toSerialize["authorizationSignatureKeyId"] = o.AuthorizationSignatureKeyId
	}
	if !IsNil(o.UserInfoEndpoint) {
		toSerialize["userInfoEndpoint"] = o.UserInfoEndpoint
	}
	if !IsNil(o.DirectUserInfoEndpointEnabled) {
		toSerialize["directUserInfoEndpointEnabled"] = o.DirectUserInfoEndpointEnabled
	}
	if !IsNil(o.DynamicRegistrationSupported) {
		toSerialize["dynamicRegistrationSupported"] = o.DynamicRegistrationSupported
	}
	if !IsNil(o.RegistrationEndpoint) {
		toSerialize["registrationEndpoint"] = o.RegistrationEndpoint
	}
	if !IsNil(o.RegistrationManagementEndpoint) {
		toSerialize["registrationManagementEndpoint"] = o.RegistrationManagementEndpoint
	}
	if !IsNil(o.PolicyUri) {
		toSerialize["policyUri"] = o.PolicyUri
	}
	if !IsNil(o.TosUri) {
		toSerialize["tosUri"] = o.TosUri
	}
	if !IsNil(o.ServiceDocumentation) {
		toSerialize["serviceDocumentation"] = o.ServiceDocumentation
	}
	if !IsNil(o.BackchannelAuthenticationEndpoint) {
		toSerialize["backchannelAuthenticationEndpoint"] = o.BackchannelAuthenticationEndpoint
	}
	if !IsNil(o.SupportedBackchannelTokenDeliveryModes) {
		toSerialize["supportedBackchannelTokenDeliveryModes"] = o.SupportedBackchannelTokenDeliveryModes
	}
	if !IsNil(o.BackchannelAuthReqIdDuration) {
		toSerialize["backchannelAuthReqIdDuration"] = o.BackchannelAuthReqIdDuration
	}
	if !IsNil(o.BackchannelPollingInterval) {
		toSerialize["backchannelPollingInterval"] = o.BackchannelPollingInterval
	}
	if !IsNil(o.BackchannelUserCodeParameterSupported) {
		toSerialize["backchannelUserCodeParameterSupported"] = o.BackchannelUserCodeParameterSupported
	}
	if !IsNil(o.BackchannelBindingMessageRequiredInFapi) {
		toSerialize["backchannelBindingMessageRequiredInFapi"] = o.BackchannelBindingMessageRequiredInFapi
	}
	if !IsNil(o.DeviceAuthorizationEndpoint) {
		toSerialize["deviceAuthorizationEndpoint"] = o.DeviceAuthorizationEndpoint
	}
	if !IsNil(o.DeviceVerificationUri) {
		toSerialize["deviceVerificationUri"] = o.DeviceVerificationUri
	}
	if !IsNil(o.DeviceVerificationUriComplete) {
		toSerialize["deviceVerificationUriComplete"] = o.DeviceVerificationUriComplete
	}
	if !IsNil(o.DeviceFlowCodeDuration) {
		toSerialize["deviceFlowCodeDuration"] = o.DeviceFlowCodeDuration
	}
	if !IsNil(o.DeviceFlowPollingInterval) {
		toSerialize["deviceFlowPollingInterval"] = o.DeviceFlowPollingInterval
	}
	if !IsNil(o.UserCodeCharset) {
		toSerialize["userCodeCharset"] = o.UserCodeCharset
	}
	if !IsNil(o.UserCodeLength) {
		toSerialize["userCodeLength"] = o.UserCodeLength
	}
	if !IsNil(o.SupportedTrustFrameworks) {
		toSerialize["supportedTrustFrameworks"] = o.SupportedTrustFrameworks
	}
	if !IsNil(o.SupportedEvidence) {
		toSerialize["supportedEvidence"] = o.SupportedEvidence
	}
	if !IsNil(o.SupportedIdentityDocuments) {
		toSerialize["supportedIdentityDocuments"] = o.SupportedIdentityDocuments
	}
	if !IsNil(o.SupportedVerificationMethods) {
		toSerialize["supportedVerificationMethods"] = o.SupportedVerificationMethods
	}
	if !IsNil(o.SupportedVerifiedClaims) {
		toSerialize["supportedVerifiedClaims"] = o.SupportedVerifiedClaims
	}
	if !IsNil(o.VerifiedClaimsValidationSchemaSet) {
		toSerialize["verifiedClaimsValidationSchemaSet"] = o.VerifiedClaimsValidationSchemaSet
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.NbfOptional) {
		toSerialize["nbfOptional"] = o.NbfOptional
	}
	if !IsNil(o.IssSuppressed) {
		toSerialize["issSuppressed"] = o.IssSuppressed
	}
	if !IsNil(o.SupportedCustomClientMetadata) {
		toSerialize["supportedCustomClientMetadata"] = o.SupportedCustomClientMetadata
	}
	if !IsNil(o.TokenExpirationLinked) {
		toSerialize["tokenExpirationLinked"] = o.TokenExpirationLinked
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
	if !IsNil(o.HsmEnabled) {
		toSerialize["hsmEnabled"] = o.HsmEnabled
	}
	if !IsNil(o.Hsks) {
		toSerialize["hsks"] = o.Hsks
	}
	if !IsNil(o.GrantManagementEndpoint) {
		toSerialize["grantManagementEndpoint"] = o.GrantManagementEndpoint
	}
	if !IsNil(o.GrantManagementActionRequired) {
		toSerialize["grantManagementActionRequired"] = o.GrantManagementActionRequired
	}
	if !IsNil(o.UnauthorizedOnClientConfigSupported) {
		toSerialize["unauthorizedOnClientConfigSupported"] = o.UnauthorizedOnClientConfigSupported
	}
	if !IsNil(o.DcrScopeUsedAsRequestable) {
		toSerialize["dcrScopeUsedAsRequestable"] = o.DcrScopeUsedAsRequestable
	}
	if !IsNil(o.EndSessionEndpoint) {
		toSerialize["endSessionEndpoint"] = o.EndSessionEndpoint
	}
	if !IsNil(o.LoopbackRedirectionUriVariable) {
		toSerialize["loopbackRedirectionUriVariable"] = o.LoopbackRedirectionUriVariable
	}
	if !IsNil(o.RequestObjectAudienceChecked) {
		toSerialize["requestObjectAudienceChecked"] = o.RequestObjectAudienceChecked
	}
	if !IsNil(o.AccessTokenForExternalAttachmentEmbedded) {
		toSerialize["accessTokenForExternalAttachmentEmbedded"] = o.AccessTokenForExternalAttachmentEmbedded
	}
	if !IsNil(o.AuthorityHints) {
		toSerialize["authorityHints"] = o.AuthorityHints
	}
	if !IsNil(o.FederationEnabled) {
		toSerialize["federationEnabled"] = o.FederationEnabled
	}
	if !IsNil(o.FederationJwks) {
		toSerialize["federationJwks"] = o.FederationJwks
	}
	if !IsNil(o.FederationSignatureKeyId) {
		toSerialize["federationSignatureKeyId"] = o.FederationSignatureKeyId
	}
	if !IsNil(o.FederationConfigurationDuration) {
		toSerialize["federationConfigurationDuration"] = o.FederationConfigurationDuration
	}
	if !IsNil(o.FederationRegistrationEndpoint) {
		toSerialize["federationRegistrationEndpoint"] = o.FederationRegistrationEndpoint
	}
	if !IsNil(o.OrganizationName) {
		toSerialize["organizationName"] = o.OrganizationName
	}
	if !IsNil(o.PredefinedTransformedClaims) {
		toSerialize["predefinedTransformedClaims"] = o.PredefinedTransformedClaims
	}
	if !IsNil(o.RefreshTokenIdempotent) {
		toSerialize["refreshTokenIdempotent"] = o.RefreshTokenIdempotent
	}
	if !IsNil(o.SignedJwksUri) {
		toSerialize["signedJwksUri"] = o.SignedJwksUri
	}
	if !IsNil(o.SupportedAttachments) {
		toSerialize["supportedAttachments"] = o.SupportedAttachments
	}
	if !IsNil(o.SupportedDigestAlgorithms) {
		toSerialize["supportedDigestAlgorithms"] = o.SupportedDigestAlgorithms
	}
	if !IsNil(o.SupportedDocuments) {
		toSerialize["supportedDocuments"] = o.SupportedDocuments
	}
	if !IsNil(o.SupportedDocumentsMethods) {
		toSerialize["supportedDocumentsMethods"] = o.SupportedDocumentsMethods
	}
	if !IsNil(o.SupportedDocumentsValidationMethods) {
		toSerialize["supportedDocumentsValidationMethods"] = o.SupportedDocumentsValidationMethods
	}
	if !IsNil(o.SupportedDocumentsVerificationMethods) {
		toSerialize["supportedDocumentsVerificationMethods"] = o.SupportedDocumentsVerificationMethods
	}
	if !IsNil(o.SupportedElectronicRecords) {
		toSerialize["supportedElectronicRecords"] = o.SupportedElectronicRecords
	}
	if !IsNil(o.SupportedClientRegistrationTypes) {
		toSerialize["supportedClientRegistrationTypes"] = o.SupportedClientRegistrationTypes
	}
	if !IsNil(o.TokenExchangeByIdentifiableClientsOnly) {
		toSerialize["tokenExchangeByIdentifiableClientsOnly"] = o.TokenExchangeByIdentifiableClientsOnly
	}
	if !IsNil(o.TokenExchangeByConfidentialClientsOnly) {
		toSerialize["tokenExchangeByConfidentialClientsOnly"] = o.TokenExchangeByConfidentialClientsOnly
	}
	if !IsNil(o.TokenExchangeByPermittedClientsOnly) {
		toSerialize["tokenExchangeByPermittedClientsOnly"] = o.TokenExchangeByPermittedClientsOnly
	}
	if !IsNil(o.TokenExchangeEncryptedJwtRejected) {
		toSerialize["tokenExchangeEncryptedJwtRejected"] = o.TokenExchangeEncryptedJwtRejected
	}
	if !IsNil(o.TokenExchangeUnsignedJwtRejected) {
		toSerialize["tokenExchangeUnsignedJwtRejected"] = o.TokenExchangeUnsignedJwtRejected
	}
	if !IsNil(o.JwtGrantByIdentifiableClientsOnly) {
		toSerialize["jwtGrantByIdentifiableClientsOnly"] = o.JwtGrantByIdentifiableClientsOnly
	}
	if !IsNil(o.JwtGrantEncryptedJwtRejected) {
		toSerialize["jwtGrantEncryptedJwtRejected"] = o.JwtGrantEncryptedJwtRejected
	}
	if !IsNil(o.JwtGrantUnsignedJwtRejected) {
		toSerialize["jwtGrantUnsignedJwtRejected"] = o.JwtGrantUnsignedJwtRejected
	}
	if !IsNil(o.DcrDuplicateSoftwareIdBlocked) {
		toSerialize["dcrDuplicateSoftwareIdBlocked"] = o.DcrDuplicateSoftwareIdBlocked
	}
	if !IsNil(o.TrustAnchors) {
		toSerialize["trustAnchors"] = o.TrustAnchors
	}
	if !IsNil(o.OpenidDroppedOnRefreshWithoutOfflineAccess) {
		toSerialize["openidDroppedOnRefreshWithoutOfflineAccess"] = o.OpenidDroppedOnRefreshWithoutOfflineAccess
	}
	if !IsNil(o.SupportedDocumentsCheckMethods) {
		toSerialize["supportedDocumentsCheckMethods"] = o.SupportedDocumentsCheckMethods
	}
	if !IsNil(o.RsResponseSigned) {
		toSerialize["rsResponseSigned"] = o.RsResponseSigned
	}
	if !IsNil(o.CnonceDuration) {
		toSerialize["cnonceDuration"] = o.CnonceDuration
	}
	if !IsNil(o.DpopNonceRequired) {
		toSerialize["dpopNonceRequired"] = o.DpopNonceRequired
	}
	if !IsNil(o.VerifiableCredentialsEnabled) {
		toSerialize["verifiableCredentialsEnabled"] = o.VerifiableCredentialsEnabled
	}
	if !IsNil(o.CredentialJwksUri) {
		toSerialize["credentialJwksUri"] = o.CredentialJwksUri
	}
	if !IsNil(o.CredentialOfferDuration) {
		toSerialize["credentialOfferDuration"] = o.CredentialOfferDuration
	}
	if !IsNil(o.DpopNonceDuration) {
		toSerialize["dpopNonceDuration"] = o.DpopNonceDuration
	}
	if !IsNil(o.PreAuthorizedGrantAnonymousAccessSupported) {
		toSerialize["preAuthorizedGrantAnonymousAccessSupported"] = o.PreAuthorizedGrantAnonymousAccessSupported
	}
	if !IsNil(o.CredentialTransactionDuration) {
		toSerialize["credentialTransactionDuration"] = o.CredentialTransactionDuration
	}
	if !IsNil(o.IntrospectionSignatureKeyId) {
		toSerialize["introspectionSignatureKeyId"] = o.IntrospectionSignatureKeyId
	}
	if !IsNil(o.ResourceSignatureKeyId) {
		toSerialize["resourceSignatureKeyId"] = o.ResourceSignatureKeyId
	}
	if !IsNil(o.UserPinLength) {
		toSerialize["userPinLength"] = o.UserPinLength
	}
	if !IsNil(o.SupportedPromptValues) {
		toSerialize["supportedPromptValues"] = o.SupportedPromptValues
	}
	if !IsNil(o.IdTokenReissuable) {
		toSerialize["idTokenReissuable"] = o.IdTokenReissuable
	}
	if !IsNil(o.CredentialJwks) {
		toSerialize["credentialJwks"] = o.CredentialJwks
	}
	if !IsNil(o.FapiModes) {
		toSerialize["fapiModes"] = o.FapiModes
	}
	if !IsNil(o.CredentialDuration) {
		toSerialize["credentialDuration"] = o.CredentialDuration
	}
	if !IsNil(o.CredentialIssuerMetadata) {
		toSerialize["credentialIssuerMetadata"] = o.CredentialIssuerMetadata
	}
	if !IsNil(o.IdTokenAudType) {
		toSerialize["idTokenAudType"] = o.IdTokenAudType
	}
	return toSerialize, nil
}

type NullableService struct {
	value *Service
	isSet bool
}

func (v NullableService) Get() *Service {
	return v.value
}

func (v *NullableService) Set(val *Service) {
	v.value = val
	v.isSet = true
}

func (v NullableService) IsSet() bool {
	return v.isSet
}

func (v *NullableService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableService(val *Service) *NullableService {
	return &NullableService{value: val, isSet: true}
}

func (v NullableService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
