# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | The sequential number of the service. The value of this property is assigned by Authlete. | [optional] [readonly] 
**ServiceName** | Pointer to **string** | The name of this service. | [optional] 
**Issuer** | Pointer to **string** | The issuer identifier of the service.  A URL that starts with  https:// and has no query or fragment component.  The value of this property is used as &#x60;iss&#x60; claim in an [ID token](https://openid.net/specs/openid-connect-core-1_0.html#IDToken) and &#x60;issuer&#x60; property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**Description** | Pointer to **string** | The description about the service. | [optional] 
**ApiKey** | Pointer to **int64** | The service ID used in Authlete API calls. The value of this property is assigned by Authlete. | [optional] [readonly] 
**ClientIdAliasEnabled** | Pointer to **bool** | Deprecated. Always &#x60;true&#x60;. | [optional] 
**Metadata** | Pointer to [**[]Pair**](Pair.md) | The &#x60;metadata&#x60; of the service. The content of the returned array depends on contexts. The predefined service metadata is listed in the following table.    | Key | Description |   | --- | --- |   | &#x60;clientCount&#x60; | The number of client applications which belong to this service.  |  | [optional] 
**CreatedAt** | Pointer to **int64** | The time at which this service was created. The value is represented as milliseconds since the UNIX epoch (&#x60;1970-01-01&#x60;).  | [optional] [readonly] 
**ModifiedAt** | Pointer to **int64** | The time at which this service was last modified. The value is represented as milliseconds since the UNIX epoch (1970-01-01).  | [optional] [readonly] 
**AuthenticationCallbackEndpoint** | Pointer to **string** | A Web API endpoint for user authentication which is to be prepared on the service side.  The endpoint must be implemented if you do not implement the UI at the authorization endpoint but use the one provided by Authlete.  The user authentication at the authorization endpoint provided by Authlete is performed by making a &#x60;POST&#x60; request to this endpoint.  | [optional] 
**AuthenticationCallbackApiKey** | Pointer to **string** | API key for basic authentication at the authentication callback endpoint.  If the value is not empty, Authlete generates Authorization header for Basic authentication when making a request to the authentication callback endpoint.  | [optional] 
**AuthenticationCallbackApiSecret** | Pointer to **string** | API secret for &#x60;basic&#x60; authentication at the authentication callback endpoint. | [optional] 
**SupportedAcrs** | Pointer to **[]string** | Values of acrs (authentication context class references) that the service supports.  The value of this property is used as &#x60;acr_values_supported&#x60; property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] [readonly] 
**SupportedGrantTypes** | Pointer to [**[]GrantType**](GrantType.md) | Values of &#x60;grant_type&#x60; request parameter that the service supports.  The value of this property is used as &#x60;grant_types_supported property&#x60; in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**SupportedResponseTypes** | Pointer to [**[]ResponseType**](ResponseType.md) | Values of &#x60;response_type&#x60; request parameter that the service supports. Valid values are listed in Response Type.  The value of this property is used as &#x60;response_types_supported&#x60; property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**SupportedAuthorizationDetailsTypes** | Pointer to **[]string** | The supported data types that can be used as values of the type field in &#x60;authorization_details&#x60;.  This property corresponds to the &#x60;authorization_details_types_supported&#x60; metadata. See \&quot;OAuth 2.0 Rich Authorization Requests\&quot; (RAR) for details.  | [optional] 
**SupportedServiceProfiles** | Pointer to [**[]ServiceProfile**](ServiceProfile.md) | The profiles that this service supports.  | [optional] 
**ErrorDescriptionOmitted** | Pointer to **bool** | The flag to indicate whether the &#x60;error_description&#x60; response parameter is omitted.  According to [RFC 6749](https://tools.ietf.org/html/rfc6749), an authorization server may include the &#x60;error_description&#x60; response parameter in error responses.  If &#x60;true&#x60;, Authlete does not embed the &#x60;error_description&#x60; response parameter in error responses.  | [optional] 
**ErrorUriOmitted** | Pointer to **bool** | The flag to indicate whether the &#x60;error_uri&#x60; response parameter is omitted.  According to [RFC 6749](https://tools.ietf.org/html/rfc6749), an authorization server may include the &#x60;error_uri&#x60; response parameter in error responses.  If &#x60;true&#x60;, Authlete does not embed the &#x60;error_uri&#x60; response parameter in error responses.  | [optional] 
**AuthorizationEndpoint** | Pointer to **string** | The authorization endpoint of the service.  A URL that starts with &#x60;https://&#x60; and has no fragment component. For example, &#x60;https://example.com/auth/authorization&#x60;.  The value of this property is used as &#x60;authorization_endpoint&#x60; property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**DirectAuthorizationEndpointEnabled** | Pointer to **bool** | The flag to indicate whether the direct authorization endpoint is enabled or not.  The path of the endpoint is &#x60;/api/auth/authorization/direct/service-api-key&#x60;.  | [optional] 
**SupportedUiLocales** | Pointer to **[]string** | UI locales that the service supports.  Each element is a language tag defined in [RFC 5646](https://tools.ietf.org/html/rfc5646). For example, &#x60;en-US&#x60; and &#x60;ja-JP&#x60;.  The value of this property is used as &#x60;ui_locales_supported&#x60; property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**SupportedDisplays** | Pointer to [**[]Display**](Display.md) | Values of &#x60;display&#x60; request parameter that service supports.  The value of this property is used as &#x60;display_values_supported&#x60; property in the Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**PkceRequired** | Pointer to **bool** | The flag to indicate whether the use of Proof Key for Code Exchange (PKCE) is always required for authorization requests by Authorization Code Flow.  If &#x60;true&#x60;, &#x60;code_challenge&#x60; request parameter is always required for authorization requests using Authorization Code Flow.  See [RFC 7636](https://tools.ietf.org/html/rfc7636) (Proof Key for Code Exchange by OAuth Public Clients) for details about &#x60;code_challenge&#x60; request parameter.  | [optional] 
**PkceS256Required** | Pointer to **bool** | The flag to indicate whether &#x60;S256&#x60; is always required as the code challenge method whenever [PKCE (RFC 7636)](https://tools.ietf.org/html/rfc7636) is used.  If this flag is set to &#x60;true&#x60;, &#x60;code_challenge_method&#x3D;S256&#x60; must be included in the authorization request whenever it includes the &#x60;code_challenge&#x60; request parameter. Neither omission of the &#x60;code_challenge_method&#x60; request parameter nor use of plain (&#x60;code_challenge_method&#x3D;plain&#x60;) is allowed.  | [optional] 
**AuthorizationResponseDuration** | Pointer to **int64** | The duration of authorization response JWTs in seconds.  [Financial-grade API: JWT Secured Authorization Response Mode for OAuth 2.0 (JARM)](https://openid.net/specs/openid-financial-api-jarm.html) defines new values for the &#x60;response_mode&#x60; request parameter. They are &#x60;query.jwt&#x60;, &#x60;fragment.jwt&#x60;, &#x60;form_post.jwt&#x60; and &#x60;jwt&#x60;. If one of them is specified as the response mode, response parameters from the authorization endpoint will be packed into a JWT. This property is used to compute the value of the &#x60;exp&#x60; claim of the JWT.  | [optional] 
**TokenEndpoint** | Pointer to **string** | The [token endpoint](https://tools.ietf.org/html/rfc6749#section-3.2) of the service.  A URL that starts with &#x60;https://&#x60; and has not fragment component. For example, &#x60;https://example.com/auth/token&#x60;.  The value of this property is used as &#x60;token_endpoint&#x60; property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**DirectTokenEndpointEnabled** | Pointer to **bool** | The flag to indicate whether the direct token endpoint is enabled or not. The path of the endpoint is &#x60;/api/auth/token/direct/service-api-key&#x60;.  | [optional] 
**SupportedTokenAuthMethods** | Pointer to [**[]ClientAuthMethod**](ClientAuthMethod.md) | Client authentication methods supported by the token endpoint of the service.  The value of this property is used as &#x60;token_endpoint_auth_methods_supports&#x60; property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**MissingClientIdAllowed** | Pointer to **bool** | The flag to indicate token requests from public clients without the &#x60;client_id&#x60; request parameter are allowed when the client can be guessed from &#x60;authorization_code&#x60; or &#x60;refresh_token&#x60;.  This flag should not be set unless you have special reasons.  | [optional] 
**RevocationEndpoint** | Pointer to **string** | The [revocation endpoint](https://tools.ietf.org/html/rfc7009) of the service.  A URL that starts with &#x60;https://&#x60;. For example, &#x60;https://example.com/auth/revocation&#x60;.  | [optional] 
**DirectRevocationEndpointEnabled** | Pointer to **bool** | The flag to indicate whether the direct revocation endpoint is enabled or not. The URL of the endpoint is &#x60;/api/auth/revocation/direct/service-api-key&#x60;.  | [optional] 
**SupportedRevocationAuthMethods** | Pointer to [**[]ClientAuthMethod**](ClientAuthMethod.md) | Client authentication methods supported at the revocation endpoint.  | [optional] 
**IntrospectionEndpoint** | Pointer to **string** | The URI of the introspection endpoint. | [optional] 
**DirectIntrospectionEndpointEnabled** | Pointer to **bool** | The flag to indicate whether the direct userinfo endpoint is enabled or not. The path of the endpoint is &#x60;/api/auth/userinfo/direct/{serviceApiKey}&#x60;.  | [optional] 
**SupportedIntrospectionAuthMethods** | Pointer to [**[]ClientAuthMethod**](ClientAuthMethod.md) | Client authentication methods supported at the introspection endpoint.  | [optional] 
**PushedAuthReqEndpoint** | Pointer to **string** | The URI of the pushed authorization request endpoint.  This property corresponds to the &#x60;pushed_authorization_request_endpoint&#x60; metadata defined in \&quot;[5. Authorization Server Metadata](https://tools.ietf.org/html/draft-lodderstedt-oauth-par#section-5)\&quot; of OAuth 2.0 Pushed Authorization Requests.  | [optional] 
**PushedAuthReqDuration** | Pointer to **int64** | The duration of pushed authorization requests in seconds.  [OAuth 2.0 Pushed Authorization Requests](https://tools.ietf.org/html/draft-lodderstedt-oauth-par) defines an endpoint (called \&quot;pushed authorization request endpoint\&quot;) which client applications can register authorization requests into and get corresponding URIs (called \&quot;request URIs\&quot;) from. The issued URIs represent the registered authorization requests. The client applications can use the URIs as the value of the &#x60;request_uri&#x60; request parameter in an authorization request.  The property represents the duration of registered authorization requests and is used as the value of the &#x60;expires_in&#x60; parameter in responses from the pushed authorization request endpoint.  | [optional] 
**ParRequired** | Pointer to **bool** | The flag to indicate whether this service requires that clients use the pushed authorization request endpoint.  This property corresponds to the &#x60;require_pushed_authorization_requests&#x60; server metadata defined in [OAuth 2.0 Pushed Authorization Requests](https://tools.ietf.org/html/draft-lodderstedt-oauth-par).  | [optional] 
**RequestObjectRequired** | Pointer to **bool** | The flag to indicate whether this service requires that authorization requests always utilize a request object by using either request or &#x60;request_uri&#x60; request parameter.  If this flag is set to &#x60;true&#x60; and the value of &#x60;traditionalRequestObjectProcessingApplied&#x60; is &#x60;false&#x60;, the value of &#x60;require_signed_request_object&#x60; server metadata of this service is reported as &#x60;true&#x60; in the discovery document. The metadata is defined in JAR (JWT Secured Authorization Request). That &#x60;require_signed_request_object&#x60; is &#x60;true&#x60; means that authorization requests which don&#39;t conform to the JAR specification are rejected.  | [optional] 
**TraditionalRequestObjectProcessingApplied** | Pointer to **bool** | The flag to indicate whether a request object is processed based on rules defined in [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html) or JAR (JWT Secured Authorization Request).  Differences between rules in OpenID Connect Core 1.0 and ones in JAR are as follows.   - JAR requires that a request object be always -signed.   - JAR does not allow request parameters outside a request object to be referred to.   - OIDC Core 1.0 requires that response_type request parameter exist outside a request object even if the request object includes the request parameter.   - OIDC Core 1.0 requires that scope request parameter exist outside a request object if the authorization request is an   - OIDC request even if the request object includes the request parameter.  If this flag is set to &#x60;false&#x60; and the value of &#x60;requestObjectRequired&#x60; is &#x60;true&#x60;, the value of &#x60;require_signed_request_object&#x60; server metadata of this service is reported as &#x60;true&#x60; in the discovery document. The metadata is defined in JAR (JWT Secured Authorization Request). That &#x60;require_signed_request_object&#x60; is &#x60;true&#x60; means that authorization requests which don&#39;t conform to the JAR specification are rejected.  | [optional] 
**MutualTlsValidatePkiCertChain** | Pointer to **bool** | The flag to indicate whether this service validates certificate chains during PKI-based client mutual TLS authentication.  | [optional] 
**TrustedRootCertificates** | Pointer to **[]string** | The list of root certificates trusted by this service for PKI-based client mutual TLS authentication.  | [optional] 
**MtlsEndpointAliases** | Pointer to [**[]NamedUri**](NamedUri.md) | The MTLS endpoint aliases.  This property corresponds to the mtls_endpoint_aliases metadata defined in \&quot;5. Metadata for Mutual TLS Endpoint Aliases\&quot; of [OAuth 2.0 Mutual TLS Client Authentication and Certificate-Bound Access Tokens](https://datatracker.ietf.org/doc/rfc8705/).  The aliases will be embedded in the response from the discovery endpoint like the following.  &#x60;&#x60;&#x60;json {   ......,   \&quot;mtls_endpoint_aliases\&quot;: {     \&quot;token_endpoint\&quot;:         \&quot;https://mtls.example.com/token\&quot;,     \&quot;revocation_endpoint\&quot;:    \&quot;https://mtls.example.com/revo\&quot;,     \&quot;introspection_endpoint\&quot;: \&quot;https://mtls.example.com/introspect\&quot;   } } &#x60;&#x60;&#x60;  | [optional] 
**AccessTokenType** | Pointer to **string** | The access token type.  This value is used as the value of &#x60;token_type&#x60; property in access token responses. If this service complies with [RFC 6750](https://tools.ietf.org/html/rfc6750), the value of this property should be &#x60;Bearer&#x60;.  See [RFC 6749 (OAuth 2.0), 7.1. Access Token Types](https://tools.ietf.org/html/rfc6749#section-7.1) for details.  | [optional] 
**TlsClientCertificateBoundAccessTokens** | Pointer to **bool** | The flag to indicate whether this service supports issuing TLS client certificate bound access tokens.  | [optional] 
**AccessTokenDuration** | Pointer to **int64** | The duration of access tokens in seconds. This value is used as the value of &#x60;expires_in&#x60; property in access token responses. &#x60;expires_in&#x60; is defined [RFC 6749, 5.1. Successful Response](https://tools.ietf.org/html/rfc6749#section-5.1).  | [optional] 
**SingleAccessTokenPerSubject** | Pointer to **bool** | The flag to indicate whether the number of access tokens per subject (and per client) is at most one or can be more.  If &#x60;true&#x60;, an attempt to issue a new access token invalidates existing access tokens that are associated with the same subject and the same client.  Note that, however, attempts by [Client Credentials Flow](https://tools.ietf.org/html/rfc6749#section-4.4) do not invalidate existing access tokens because access tokens issued by Client Credentials Flow are not associated with any end-user&#39;s subject. Also note that an attempt by [Refresh Token Flow](https://tools.ietf.org/html/rfc6749#section-6) invalidates the coupled access token only and this invalidation is always performed regardless of whether the value of this setting item is &#x60;true&#x60; or &#x60;false&#x60;.  | [optional] 
**AccessTokenSignAlg** | Pointer to [**JwsAlg**](JwsAlg.md) |  | [optional] 
**AccessTokenSignatureKeyId** | Pointer to **string** | The key ID to identify a JWK used for signing access tokens.  A JWK Set can be registered as a property of a service. A JWK Set can contain 0 or more JWKs. Authlete Server has to pick up one JWK for signing from the JWK Set when it generates a JWT-based access token. Authlete Server searches the registered JWK Set for a JWK which satisfies conditions for access token signature. If the number of JWK candidates which satisfy the conditions is 1, there is no problem. On the other hand, if there exist multiple candidates, a Key ID is needed to be specified so that Authlete Server can pick up one JWK from among the JWK candidates.  | [optional] 
**RefreshTokenDuration** | Pointer to **int64** | The duration of refresh tokens in seconds. The related specifications have no requirements on refresh token duration, but Authlete sets expiration for refresh tokens. | [optional] 
**RefreshTokenDurationKept** | Pointer to **bool** | The flag to indicate whether the remaining duration of the used refresh token is taken over to the newly issued refresh token.  | [optional] 
**RefreshTokenDurationReset** | Pointer to **bool** | The flag which indicates whether duration of refresh tokens are reset when they are used even if the &#x60;refreshTokenKept&#x60; property of this service set to is &#x60;true&#x60; (&#x3D; even if \&quot;Refresh Token Continuous Use\&quot; is \&quot;Kept\&quot;).  This flag has no effect when the &#x60;refreshTokenKept&#x60; property is set to &#x60;false&#x60;. In other words, if this service issues a new refresh token on every refresh token request, the refresh token will have fresh duration (unless &#x60;refreshTokenDurationKept&#x60; is set to &#x60;true&#x60;) and this &#x60;refreshTokenDurationReset&#x60; property is not referenced.  | [optional] 
**RefreshTokenKept** | Pointer to **bool** | The flag to indicate whether a refresh token remains unchanged or gets renewed after its use.  If &#x60;true&#x60;, a refresh token used to get a new access token remains valid after its use. Otherwise, if &#x60;false&#x60;, a refresh token is invalidated after its use and a new refresh token is issued.  See [RFC 6749 6. Refreshing an Access Token](https://tools.ietf.org/html/rfc6749#section-6), as to how to get a new access token using a refresh token.  | [optional] 
**SupportedScopes** | Pointer to [**[]Scope**](Scope.md) | Scopes supported by the service.  Authlete strongly recommends that the service register at least the following scopes.  | Name | Description | | --- | --- | | openid | A permission to get an ID token of an end-user. The &#x60;openid&#x60; scope appears in [OpenID Connect Core 1.0, 3.1.2.1. Authentication Request, scope](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest). Without this scope, Authlete does not allow &#x60;response_type&#x60; request parameter to have values other than code and token. | | profile | A permission to get information about &#x60;name&#x60;, &#x60;family_name&#x60;, &#x60;given_name&#x60;, &#x60;middle_name&#x60;, &#x60;nickname&#x60;, &#x60;preferred_username&#x60;, &#x60;profile&#x60;, &#x60;picture&#x60;, &#x60;website&#x60;, &#x60;gender&#x60;, &#x60;birthdate&#x60;, &#x60;zoneinfo&#x60;, &#x60;locale&#x60; and &#x60;updated_at&#x60; from the user info endpoint. See [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) for details. | | email | A permission to get information about &#x60;email&#x60; and &#x60;email_verified&#x60; from the user info endpoint. See [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) for details. | | address | A permission to get information about address from the user info endpoint. See [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) and [5.1.1. Address Claim](https://openid.net/specs/openid-connect-core-1_0.html#AddressClaim) for details. | | phone | A permission to get information about &#x60;phone_number&#x60; and &#x60;phone_number_verified&#x60; from the user info endpoint. See [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) for details. | | offline_access | A permission to get information from the user info endpoint even when the end-user is not present. See [OpenID Connect Core 1.0, 11. Offline Access](https://openid.net/specs/openid-connect-core-1_0.html#OfflineAccess) for details. |  The value of this property is used as &#x60;scopes_supported&#x60; property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**ScopeRequired** | Pointer to **bool** | The flag to indicate whether requests that request no scope are rejected or not.  When a request has no explicit &#x60;scope&#x60; parameter and the service&#39;s pre-defined default scope set is empty, the authorization server regards the request requests no scope. When this flag is set to &#x60;true&#x60;, requests that request no scope are rejected.  The requirement below excerpted from [RFC 6749 Section 3.3](https://tools.ietf.org/html/rfc6749#section-3.3) does not explicitly mention the case where the default scope set is empty.  &gt; If the client omits the scope parameter when requesting authorization, the authorization server MUST either process the request using a pre-defined default value or fail the request indicating an invalid scope.  However, if you interpret *\&quot;the default scope set exists but is empty\&quot;* as *\&quot;the default scope set does not exist\&quot;* and want to strictly conform to the requirement above, this flag has to be &#x60;true&#x60;.  | [optional] 
**IdTokenDuration** | Pointer to **int64** | &#39;The duration of [ID token](https://openid.net/specs/openid-connect-core-1_0.html#IDToken)s in seconds. This value is used to calculate the value of &#x60;exp&#x60; claim in an ID token.&#39;  | [optional] 
**AllowableClockSkew** | Pointer to **int32** | The allowable clock skew between the server and clients in seconds.  The clock skew is taken into consideration when time-related claims in a JWT (e.g. &#x60;exp&#x60;, &#x60;iat&#x60;, &#x60;nbf&#x60;) are verified.  | [optional] 
**SupportedClaimTypes** | Pointer to [**[]ClaimType**](ClaimType.md) | Claim types supported by the service. Valid values are listed in Claim Type. Note that Authlete currently doesn&#39;t provide any API to help implementations for &#x60;AGGREGATED&#x60; and &#x60;DISTRIBUTED&#x60;.  The value of this property is used as &#x60;claim_types_supported&#x60; property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**SupportedClaimLocales** | Pointer to **[]string** | Claim locales that the service supports. Each element is a language tag defined in [RFC 5646](https://tools.ietf.org/html/rfc5646). For example, &#x60;en-US&#x60; and &#x60;ja-JP&#x60;. See [OpenID Connect Core 1.0, 5.2. Languages and Scripts](https://openid.net/specs/openid-connect-core-1_0.html#ClaimsLanguagesAndScripts) for details.  The value of this property is used as &#x60;claims_locales_supported&#x60; property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**SupportedClaims** | Pointer to **[]string** | Claim names that the service supports. The standard claim names listed in [OpenID Connect Core 1.0, 5.1. Standard Claim](https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims) should be supported. The following is the list of standard claims.  - &#x60;sub&#x60; - &#x60;name&#x60; - &#x60;given_name&#x60; - &#x60;family_name&#x60; - &#x60;middle_name&#x60; - &#x60;nickname&#x60; - &#x60;preferred_username&#x60; - &#x60;profile&#x60; - &#x60;picture&#x60; - &#x60;website&#x60; - &#x60;email&#x60; - &#x60;email_verified&#x60; - &#x60;gender&#x60; - &#x60;birthdate&#x60; - &#x60;zoneinfo&#x60; - &#x60;locale&#x60; - &#x60;phone_number&#x60; - &#x60;phone_number_verified&#x60; - &#x60;address&#x60; - &#x60;updated_at&#x60;  The value of this property is used as &#x60;claims_supported&#x60; property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  The service may support its original claim names. See [OpenID Connect Core 1.0, 5.1.2. Additional Claims](https://openid.net/specs/openid-connect-core-1_0.html#AdditionalClaims).  | [optional] 
**ClaimShortcutRestrictive** | Pointer to **bool** | The flag indicating whether claims specified by shortcut scopes (e.g. &#x60;profile&#x60;) are included in the issued ID token only when no access token is issued.  To strictly conform to the description below excerpted from [OpenID Connect Core 1.0 Section 5.4](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims), this flag has to be &#x60;true&#x60;.  &gt; The Claims requested by the profile, email, address, and phone scope values are returned from the UserInfo Endpoint, as described in Section 5.3.2, when a response_type value is used that results in an Access Token being issued. However, when no Access Token is issued (which is the case for the response_type value id_token), the resulting Claims are returned in the ID Token.  | [optional] 
**JwksUri** | Pointer to **string** | The URL of the service&#39;s [JSON Web Key Set](https://tools.ietf.org/html/rfc7517) document. For example, &#x60;http://example.com/auth/jwks&#x60;.  Client applications accesses this URL (1) to get the public key of the service to validate the signature of an ID token issued by the service and (2) to get the public key of the service to encrypt an request object of the client application. See [OpenID Connect Core 1.0, 10. Signatures and Encryption](https://openid.net/specs/openid-connect-core-1_0.html#SigEnc) for details.  The value of this property is used as &#x60;jwks_uri&#x60; property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**DirectJwksEndpointEnabled** | Pointer to **bool** | &#39;The flag to indicate whether the direct jwks endpoint is enabled or not. The path of the endpoint is &#x60;/api/service/jwks/get/direct/service-api-key&#x60;. &#39;  | [optional] 
**Jwks** | Pointer to **string** | The content of the service&#39;s [JSON Web Key Set](https://tools.ietf.org/html/rfc7517) document.  If this property is not &#x60;null&#x60; in a &#x60;/service/create&#x60; request or a &#x60;/service/update&#x60; request, Authlete hosts the content in the database. This property must not be &#x60;null&#x60; and must contain pairs of public/private keys if the service wants to support asymmetric signatures for ID tokens and asymmetric encryption for request objects. See [OpenID Connect Core 1.0, 10. Signatures and Encryption](https://openid.net/specs/openid-connect-core-1_0.html#SigEnc) for details.  | [optional] 
**IdTokenSignatureKeyId** | Pointer to **string** | The key ID to identify a JWK used for ID token signature using an asymmetric key.  A JWK Set can be registered as a property of a Service. A JWK Set can contain 0 or more JWKs (See [RFC 7517](https://tools.ietf.org/html/rfc7517) for details about JWK). Authlete Server has to pick up one JWK for signature from the JWK Set when it generates an ID token and signature using an asymmetric key is required. Authlete Server searches the registered JWK Set for a JWK which satisfies conditions for ID token signature. If the number of JWK candidates which satisfy the conditions is 1, there is no problem. On the other hand, if there exist multiple candidates, a [Key ID](https://tools.ietf.org/html/rfc7517#section-4.5) is needed to be specified so that Authlete Server can pick up one JWK from among the JWK candidates.  This &#x60;idTokenSignatureKeyId&#x60; property exists for the purpose described above. For key rotation (OpenID Connect Core 1.0, [10.1.1. Rotation of Asymmetric Signing Keys](http://openid.net/specs/openid-connect-core-1_0.html#RotateSigKeys)), this mechanism is needed.  | [optional] 
**UserInfoSignatureKeyId** | Pointer to **string** | The key ID to identify a JWK used for user info signature using an asymmetric key.  A JWK Set can be registered as a property of a Service. A JWK Set can contain 0 or more JWKs (See [RFC 7517](https://tools.ietf.org/html/rfc7517) for details about JWK). Authlete Server has to pick up one JWK for signature from the JWK Set when it is required to sign user info (which is returned from [userinfo endpoint](http://openid.net/specs/openid-connect-core-1_0.html#UserInfo)) using an asymmetric key. Authlete Server searches the registered JWK Set for a JWK which satisfies conditions for user info signature. If the number of JWK candidates which satisfy the conditions is 1, there is no problem. On the other hand, if there exist multiple candidates, a [Key ID](https://tools.ietf.org/html/rfc7517#section-4.5) is needed to be specified so that Authlete Server can pick up one JWK from among the JWK candidates.  This &#x60;userInfoSignatureKeyId&#x60; property exists for the purpose described above. For key rotation (OpenID Connect Core 1.0, [10.1.1. Rotation of Asymmetric Signing Keys](http://openid.net/specs/openid-connect-core-1_0.html#RotateSigKeys)), this mechanism is needed.  | [optional] 
**AuthorizationSignatureKeyId** | Pointer to **string** | The key ID to identify a JWK used for signing authorization responses using an asymmetric key.  [Financial-grade API: JWT Secured Authorization Response Mode for OAuth 2.0 (JARM)](https://openid.net/specs/openid-financial-api-jarm.html) defines new values for the &#x60;response_mode&#x60; request parameter. They are &#x60;query.jwt&#x60;, &#x60;fragment.jwt&#x60;, &#x60;form_post.jwt&#x60; and &#x60;jwt&#x60;. If one of them is specified as the response mode, response parameters from the authorization endpoint will be packed into a JWT. This property is used to compute the value of the &#x60;exp&#x60; claim of the JWT.  Authlete Server searches the JWK Set for a JWK which satisfies conditions for authorization response signature. If the number of JWK candidates which satisfy the conditions is 1, there is no problem. On the other hand, if there exist multiple candidates, a Key ID is needed to be specified so that Authlete Server can pick up one JWK from among the JWK candidates. This property exists to specify the key ID.  | [optional] 
**UserInfoEndpoint** | Pointer to **string** | The [user info endpoint](http://openid.net/specs/openid-connect-core-1_0.html#UserInfo) of the service. A URL that starts with &#x60;https://&#x60;. For example, &#x60;https://example.com/auth/userinfo&#x60;.  The value of this property is used as &#x60;userinfo_endpoint&#x60; property in the [OpenID Provider Metadata](http://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**DirectUserInfoEndpointEnabled** | Pointer to **bool** | The flag to indicate whether the direct userinfo endpoint is enabled or not. The path of the endpoint is &#x60;/api/auth/userinfo/direct/service-api-key&#x60;.  | [optional] 
**DynamicRegistrationSupported** | Pointer to **bool** | The boolean flag which indicates whether the [OAuth 2.0 Dynamic Client Registration Protocol](https://tools.ietf.org/html/rfc7591) is supported.  | [optional] 
**RegistrationEndpoint** | Pointer to **string** | The [registration endpoint](http://openid.net/specs/openid-connect-registration-1_0.html#ClientRegistration) of the service. A URL that starts with &#x60;https://&#x60;. For example, &#x60;https://example.com/auth/registration&#x60;.  The value of this property is used as &#x60;registration_endpoint&#x60; property in the [OpenID Provider Metadata](http://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**RegistrationManagementEndpoint** | Pointer to **string** | The URI of the registration management endpoint. If dynamic client registration is supported, and this is set, this URI will be used as the basis of the client&#39;s management endpoint by appending &#x60;/clientid}/&#x60; to it as a path element. If this is unset, the value of &#x60;registrationEndpoint&#x60; will be used as the URI base instead.  | [optional] 
**PolicyUri** | Pointer to **string** | The URL of the \&quot;Policy\&quot; of the service.  The value of this property is used as &#x60;op_policy_uri&#x60; property in the [OpenID Provider Metadata](http://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**TosUri** | Pointer to **string** | The URL of the \&quot;Terms Of Service\&quot; of the service.  The value of this property is used as &#x60;op_tos_uri&#x60; property in the [OpenID Provider Metadata](http://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**ServiceDocumentation** | Pointer to **string** | The URL of a page where documents for developers can be found.  The value of this property is used as &#x60;service_documentation&#x60; property in the [OpenID Provider Metadata](http://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**BackchannelAuthenticationEndpoint** | Pointer to **string** | The URI of backchannel authentication endpoint, which is defined in the specification of [CIBA (Client Initiated Backchannel Authentication)](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html).  | [optional] 
**SupportedBackchannelTokenDeliveryModes** | Pointer to [**[]DeliveryMode**](DeliveryMode.md) | The supported backchannel token delivery modes. This property corresponds to the &#x60;backchannel_token_delivery_modes_supported&#x60; metadata.  Backchannel token delivery modes are defined in the specification of [CIBA (Client Initiated Backchannel Authentication)](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html).  | [optional] 
**BackchannelAuthReqIdDuration** | Pointer to **int32** | The duration of backchannel authentication request IDs issued from the backchannel authentication endpoint in seconds. This is used as the value of the &#x60;expires_in&#x60; property in responses from the backchannel authentication endpoint.  | [optional] 
**BackchannelPollingInterval** | Pointer to **int32** | The minimum interval between polling requests to the token endpoint from client applications in seconds. This is used as the value of the &#x60;interval&#x60; property in responses from the backchannel authentication endpoint.  | [optional] 
**BackchannelUserCodeParameterSupported** | Pointer to **bool** | The boolean flag which indicates whether the &#x60;user_code&#x60; request parameter is supported at the backchannel authentication endpoint. This property corresponds to the &#x60;backchannel_user_code_parameter_supported&#x60; metadata.  | [optional] 
**BackchannelBindingMessageRequiredInFapi** | Pointer to **bool** | The flag to indicate whether the &#x60;binding_message&#x60; request parameter is always required whenever a backchannel authentication request is judged as a request for Financial-grade API.  The FAPI-CIBA profile requires that the authorization server _\&quot;shall ensure unique authorization context exists in the authorization request or require a &#x60;binding_message&#x60; in the authorization request\&quot;_ (FAPI-CIBA, 5.2.2, 2). The simplest way to fulfill this requirement is to set this property to &#x60;true&#x60;.  If this property is set to &#x60;false&#x60;, the &#x60;binding_message&#x60; request parameter remains optional even in FAPI context, but in exchange, your authorization server must implement a custom mechanism that ensures each backchannel authentication request has unique context.  | [optional] 
**DeviceAuthorizationEndpoint** | Pointer to **string** | The URI of the device authorization endpoint.  Device authorization endpoint is defined in the specification of OAuth 2.0 Device Authorization Grant.  | [optional] 
**DeviceVerificationUri** | Pointer to **string** | The verification URI for the device flow. This URI is used as the value of the &#x60;verification_uri&#x60; parameter in responses from the device authorization endpoint.  | [optional] 
**DeviceVerificationUriComplete** | Pointer to **string** | The verification URI for the device flow with a placeholder for a user code. This URI is used to build the value of the &#x60;verification_uri_complete&#x60; parameter in responses from the device authorization endpoint.  It is expected that the URI contains a fixed string &#x60;USER_CODE&#x60; somewhere as a placeholder for a user code. For example, like the following.  &#x60;https://example.com/device?user\\_code&#x3D;USER\\_CODE&#x60;  The fixed string is replaced with an actual user code when Authlete builds a verification URI with a user code for the &#x60;verification_uri_complete&#x60; parameter.  If this URI is not set, the &#x60;verification_uri_complete&#x60; parameter won&#39;t appear in device authorization responses.  | [optional] 
**DeviceFlowCodeDuration** | Pointer to **int32** | The duration of device verification codes and end-user verification codes issued from the device authorization endpoint in seconds. This is used as the value of the &#x60;expires_in&#x60; property in responses from the device authorization endpoint.  | [optional] 
**DeviceFlowPollingInterval** | Pointer to **int32** | The minimum interval between polling requests to the token endpoint from client applications in seconds in device flow. This is used as the value of the &#x60;interval&#x60; property in responses from the device authorization endpoint.  | [optional] 
**UserCodeCharset** | Pointer to [**UserCodeCharset**](UserCodeCharset.md) |  | [optional] 
**UserCodeLength** | Pointer to **int32** | The length of end-user verification codes (&#x60;user_code&#x60;) for Device Flow.  | [optional] 
**SupportedTrustFrameworks** | Pointer to **[]string** | Trust frameworks supported by this service. This corresponds to the &#x60;trust_frameworks_supported&#x60; [metadata](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html#rfc.section.7).  | [optional] 
**SupportedEvidence** | Pointer to **[]string** | Evidence supported by this service. This corresponds to the &#x60;evidence_supported&#x60; [metadata](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html#rfc.section.7).  | [optional] 
**SupportedIdentityDocuments** | Pointer to **[]string** | Identity documents supported by this service. This corresponds to the &#x60;id_documents_supported&#x60; [metadata](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html#rfc.section.7).  | [optional] 
**SupportedVerificationMethods** | Pointer to **[]string** | Verification methods supported by this service. This corresponds to the &#x60;id_documents_verification_methods_supported&#x60; [metadata](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html#rfc.section.7).  | [optional] 
**SupportedVerifiedClaims** | Pointer to **[]string** | Verified claims supported by this service. This corresponds to the &#x60;claims_in_verified_claims_supported&#x60; [metadata](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html#rfc.section.7).  | [optional] 
**VerifiedClaimsValidationSchemaSet** | Pointer to [**NullableVerifiedClaimsValidationSchema**](VerifiedClaimsValidationSchema.md) | OIDC4IDA / verifiedClaimsValidationSchemaSet  | [optional] 
**Attributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of this service.  | [optional] 
**NbfOptional** | Pointer to **bool** | The flag indicating whether the nbf claim in the request object is optional even when the authorization request is regarded as a FAPI-Part2 request.  The final version of Financial-grade API was approved in January, 2021. The Part 2 of the final version has new requirements on lifetime of request objects. They require that request objects contain an &#x60;nbf&#x60; claim and the lifetime computed by &#x60;exp&#x60; - &#x60;nbf&#x60; be no longer than 60 minutes.  Therefore, when an authorization request is regarded as a FAPI-Part2 request, the request object used in the authorization request must contain an nbf claim. Otherwise, the authorization server rejects the authorization request.  When this flag is &#x60;true&#x60;, the &#x60;nbf&#x60; claim is treated as an optional claim even when the authorization request is regarded as a FAPI-Part2 request. That is, the authorization server does not perform the validation on lifetime of the request object.  Skipping the validation is a violation of the FAPI specification. The reason why this flag has been prepared nevertheless is that the new requirements (which do not exist in the Implementer&#39;s Draft 2 released in October, 2018) have big impacts on deployed implementations of client applications and Authlete thinks there should be a mechanism whereby to make the migration from ID2 to Final smooth without breaking live systems.  | [optional] 
**IssSuppressed** | Pointer to **bool** | The flag indicating whether generation of the iss response parameter is suppressed.  \&quot;OAuth 2.0 Authorization Server Issuer Identifier in Authorization Response\&quot; has defined a new authorization response parameter, &#x60;iss&#x60;, as a countermeasure for a certain type of mix-up attacks.  The specification requires that the &#x60;iss&#x60; response parameter always be included in authorization responses unless JARM (JWT Secured Authorization Response Mode) is used.  When this flag is &#x60;true&#x60;, the authorization server does not include the &#x60;iss&#x60; response parameter in authorization responses. By turning this flag on and off, developers of client applications can experiment the mix-up attack and the effect of the &#x60;iss&#x60; response parameter.  Note that this flag should not be &#x60;true&#x60; in production environment unless there are special reasons for it.  | [optional] 
**SupportedCustomClientMetadata** | Pointer to **[]string** | custom client metadata supported by this service.  Standard specifications define client metadata as necessary. The following are such examples.  * [OpenID Connect Dynamic Client Registration 1.0](https://openid.net/specs/openid-connect-registration-1_0.html) * [RFC 7591 OAuth 2.0 Dynamic Client Registration Protocol](https://www.rfc-editor.org/rfc/rfc7591.html) * [RFC 8705 OAuth 2.0 Mutual-TLS Client Authentication and Certificate-Bound Access Tokens](https://www.rfc-editor.org/rfc/rfc8705.html) * [OpenID Connect Client-Initiated Backchannel Authentication Flow - Core 1.0](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html) * [The OAuth 2.0 Authorization Framework: JWT Secured Authorization Request (JAR)](https://datatracker.ietf.org/doc/draft-ietf-oauth-jwsreq/) * [Financial-grade API: JWT Secured Authorization Response Mode for OAuth 2.0 (JARM)](https://openid.net/specs/openid-financial-api-jarm.html) * [OAuth 2.0 Pushed Authorization Requests (PAR)](https://datatracker.ietf.org/doc/rfc9126/) * [OAuth 2.0 Rich Authorization Requests (RAR)](https://datatracker.ietf.org/doc/draft-ietf-oauth-rar/)  Standard client metadata included in Client Registration Request and Client Update Request (cf. [OIDC DynReg](https://openid.net/specs/openid-connect-registration-1_0.html), [RFC 7591](https://www.rfc-editor.org/rfc/rfc7591.html) and [RFC 7592](https://www.rfc-editor.org/rfc/rfc7592.html)) are, if supported by Authlete, stored into Authlete database. On the other hand, unrecognized client metadata are discarded.  By listing up custom client metadata in advance by using this property (&#x60;supportedCustomClientMetadata&#x60;), Authlete can recognize them and stores their values into the database. The stored custom client metadata values can be referenced by &#x60;customMetadata&#x60;.  | [optional] 
**TokenExpirationLinked** | Pointer to **bool** | The flag indicating whether the expiration date of an access token never exceeds that of the corresponding refresh token.  When a new access token is issued by a refresh token request (&#x3D; a token request with &#x60;grant_type&#x3D;refresh_token&#x60;), the expiration date of the access token may exceed the expiration date of the corresponding refresh token. This behavior itself is not wrong and may happen when &#x60;refreshTokenKept&#x60; is &#x60;true&#x60; and/or when &#x60;refreshTokenDurationKept&#x60; is &#x60;true&#x60;.  When this flag is &#x60;true&#x60;, the expiration date of an access token never exceeds that of the corresponding refresh token regardless of the calculated duration based on other settings such as &#x60;accessTokenDuration&#x60;, &#x60;accessTokenDuration&#x60; in &#x60;extension&#x60; and &#x60;access_token.duration&#x60; scope attribute.  It is technically possible to set a value which is bigger than the duration of refresh tokens as the duration of access tokens although it is strange. In the case, the duration of an access token becomes longer than the duration of the refresh token which is issued together with the access token. Even if the duration values are configured so, if this flag is &#x60;true&#x60;, the expiration date of the access token does not exceed that of the refresh token. That is, the duration of the access token will be shortened, and as a result, the access token and the refresh token will have the same expiration date.  | [optional] 
**FrontChannelRequestObjectEncryptionRequired** | Pointer to **bool** | The flag indicating whether encryption of request object is required when the request object is passed through the front channel.  This flag does not affect the processing of request objects at the Pushed Authorization Request Endpoint, which is defined in [OAuth 2.0 Pushed Authorization Requests](https://datatracker.ietf.org/doc/rfc9126/). Unecrypted request objects are accepted at the endpoint even if this flag is &#x60;true&#x60;.  This flag does not indicate whether a request object is always required. There is a different flag, &#x60;requestObjectRequired&#x60;, for the purpose. See the description of &#x60;requestObjectRequired&#x60; for details.  Even if this flag is &#x60;false&#x60;, encryption of request object is required if the &#x60;frontChannelRequestObjectEncryptionRequired&#x60; flag of the client is &#x60;true&#x60;.  | [optional] 
**RequestObjectEncryptionAlgMatchRequired** | Pointer to **bool** | The flag indicating whether the JWE alg of encrypted request object must match the &#x60;request_object_encryption_alg&#x60; client metadata of the client that has sent the request object.  The request_object_encryption_alg client metadata itself is defined in [OpenID Connect Dynamic Client Registration 1.0](https://openid.net/specs/openid-connect-registration-1_0.html) as follows.  &gt; request_object_encryption_alg &gt; &gt; OPTIONAL. JWE [JWE] alg algorithm [JWA] the RP is declaring that it may use for encrypting Request Objects sent to the OP. This parameter SHOULD be included when symmetric encryption will be used, since this signals to the OP that a client_secret value needs to be returned from which the symmetric key will be derived, that might not otherwise be returned. The RP MAY still use other supported encryption algorithms or send unencrypted Request Objects, even when this parameter is present. If both signing and encryption are requested, the Request Object will be signed then encrypted, with the result being a Nested JWT, as defined in [JWT]. The default, if omitted, is that the RP is not declaring whether it might encrypt any Request Objects.  The point here is \&quot;The RP MAY still use other supported encryption algorithms or send unencrypted Request Objects, even when this parameter is present.\&quot;  The Client&#39;s property that represents the client metadata is &#x60;requestEncryptionAlg&#x60;. See the description of &#x60;requestEncryptionAlg&#x60; for details.  Even if this flag is &#x60;false&#x60;, the match is required if the &#x60;requestObjectEncryptionAlgMatchRequired&#x60; flag of the client is &#x60;true&#x60;.  | [optional] 
**RequestObjectEncryptionEncMatchRequired** | Pointer to **bool** | The flag indicating whether the JWE &#x60;enc&#x60; of encrypted request object must match the &#x60;request_object_encryption_enc&#x60; client metadata of the client that has sent the request object.  The &#x60;request_object_encryption_enc&#x60; client metadata itself is defined in [OpenID Connect Dynamic Client Registration 1.0](https://openid.net/specs/openid-connect-registration-1_0.html) as follows.  &gt; request_object_encryption_enc &gt; &gt; OPTIONAL. JWE enc algorithm [JWA] the RP is declaring that it may use for encrypting Request Objects sent to the OP. If request_object_encryption_alg is specified, the default for this value is A128CBC-HS256. When request_object_encryption_enc is included, request_object_encryption_alg MUST also be provided.  The Client&#39;s property that represents the client metadata is &#x60;requestEncryptionEnc&#x60;. See the description of &#x60;requestEncryptionEnc&#x60; for details.  Even if this flag is false, the match is required if the &#x60;requestObjectEncryptionEncMatchRequired&#x60; flag is &#x60;true&#x60;.  | [optional] 
**HsmEnabled** | Pointer to **bool** | The flag indicating whether HSM (Hardware Security Module) support is enabled for this service.  When this flag is &#x60;false&#x60;, keys managed in HSMs are not used even if they exist. In addition, &#x60;/api/hsk/_*&#x60; APIs reject all requests.  Even if this flag is &#x60;true&#x60;, HSM-related features do not work if the configuration of the Authlete server you are using does not support HSM.  | [optional] 
**Hsks** | Pointer to [**[]Hsk**](Hsk.md) | The information about keys managed on HSMs (Hardware Security Modules).  This &#x60;hsks&#x60; property is output only, meaning that &#x60;hsks&#x60; in requests to &#x60;/api/service/create&#x60; API and &#x60;/api/service/update&#x60; API do not have any effect. The contents of this property is controlled only by &#x60;/api/hsk/_*&#x60; APIs.  | [optional] 
**GrantManagementEndpoint** | Pointer to **string** | The URL of the grant management endpoint.  | [optional] 
**GrantManagementActionRequired** | Pointer to **bool** | The flag indicating whether every authorization request (and any request serving as an authorization request such as CIBA backchannel authentication request and device authorization request) must include the &#x60;grant_management_action&#x60; request parameter.  This property corresponds to the &#x60;grant_management_action_required&#x60; server metadata defined in [Grant Management for OAuth 2.0](https://openid.net/specs/fapi-grant-management.html).  Note that setting true to this property will result in blocking all public clients because the specification requires that grant management be usable only by confidential clients for security reasons.  | [optional] 
**UnauthorizedOnClientConfigSupported** | Pointer to **bool** | The flag indicating whether Authlete&#39;s &#x60;/api/client/registration&#x60; API uses &#x60;UNAUTHORIZED&#x60; as a value of the &#x60;action&#x60; response parameter when appropriate.  The &#x60;UNAUTHORIZED&#x60; enum value was initially not defined as a possible value of the &#x60;action&#x60; parameter in an &#x60;/api/client/registration&#x60; API response. This means that implementations of client &#x60;configuration&#x60; endpoint were not able to conform to [RFC 7592](https://www.rfc-editor.org/rfc/rfc7592.html) strictly.  For backward compatibility (to avoid breaking running systems), Authlete&#39;s &#x60;/api/client/registration&#x60; API does not return the &#x60;UNAUTHORIZED&#x60; enum value if this flag is not turned on.  The steps an existing implementation of client configuration endpoint has to do in order to conform to the requirement related to \&quot;401 Unauthorized\&quot; are as follows.  1. Update the Authlete library (e.g. authlete-java-common) your system is using. 2. Update your implementation of client configuration endpoint so that it can handle the &#x60;UNAUTHORIZED&#x60; action. 3. Turn on this &#x60;unauthorizedOnClientConfigSupported&#x60; flag.  | [optional] 
**DcrScopeUsedAsRequestable** | Pointer to **bool** | The flag indicating whether the &#x60;scope&#x60; request parameter in dynamic client registration and update requests (RFC 7591 and RFC 7592) is used as scopes that the client can request.  Limiting the range of scopes that a client can request is achieved by listing scopes in the &#x60;client.extension.requestableScopes&#x60; property and setting the &#x60;client.extension.requestableScopesEnabled&#x60; property to &#x60;true&#x60;. This feature is called \&quot;requestable scopes\&quot;.  This property affects behaviors of &#x60;/api/client/registration&#x60; and other family APIs.  | [optional] 
**EndSessionEndpoint** | Pointer to **string** | The endpoint for clients ending the sessions.  A URL that starts with &#x60;https://&#x60; and has no fragment component. For example, &#x60;https://example.com/auth/endSession&#x60;.  The value of this property is used as &#x60;end_session_endpoint&#x60; property in the [OpenID Provider Metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).  | [optional] 
**LoopbackRedirectionUriVariable** | Pointer to **bool** | The flag indicating whether the port number component of redirection URIs can be variable when the host component indicates loopback.  When this flag is &#x60;true&#x60;, if the host component of a redirection URI specified in an authorization request indicates loopback (to be precise, when the host component is localhost, &#x60;127.0.0.1&#x60; or &#x60;::1&#x60;), the port number component is ignored when the specified redirection URI is compared to pre-registered ones. This behavior is described in [7.3. Loopback Interface Redirection]( https://www.rfc-editor.org/rfc/rfc8252.html#section-7.3) of [RFC 8252 OAuth 2.0](https://www.rfc-editor.org/rfc/rfc8252.html) for Native Apps.  [3.1.2.3. Dynamic Configuration](https://www.rfc-editor.org/rfc/rfc6749.html#section-3.1.2.3) of [RFC 6749](https://www.rfc-editor.org/rfc/rfc6749.html) states _\&quot;If the client registration included the full redirection URI, the authorization server MUST compare the two URIs using simple string comparison as defined in [RFC3986] Section 6.2.1.\&quot;_ Also, the description of &#x60;redirect_uri&#x60; in [3.1.2.1. Authentication Request](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest) of [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html) states _\&quot;This URI MUST exactly match one of the Redirection URI values for the Client pre-registered at the OpenID Provider, with the matching performed as described in Section 6.2.1 of [RFC3986] (**Simple String Comparison**).\&quot;_ These \&quot;Simple String Comparison\&quot; requirements are preceded by this flag. That is, even when the conditions described in RFC 6749 and OpenID Connect Core 1.0 are satisfied, the port number component of loopback redirection URIs can be variable when this flag is &#x60;true&#x60;.  [8.3. Loopback Redirect Considerations](https://www.rfc-editor.org/rfc/rfc8252.html#section-8.3) of [RFC 8252](https://www.rfc-editor.org/rfc/rfc8252.html) states as follows.  &gt; While redirect URIs using localhost (i.e., &#x60;\&quot;http://localhost:{port}/{path}\&quot;&#x60;) function similarly to loopback IP redirects described in Section 7.3, the use of localhost is NOT RECOMMENDED. Specifying a redirect URI with the loopback IP literal rather than localhost avoids inadvertently listening on network interfaces other than the loopback interface. It is also less susceptible to client-side firewalls and misconfigured host name resolution on the user&#39;s device.  However, Authlete allows the port number component to be variable in the case of &#x60;localhost&#x60;, too. It is left to client applications whether they use &#x60;localhost&#x60; or a literal loopback IP address (&#x60;127.0.0.1&#x60; for IPv4 or &#x60;::1&#x60; for IPv6).  Section 7.3 and Section 8.3 of [RFC 8252](https://www.rfc-editor.org/rfc/rfc8252.html) state that loopback redirection URIs use the &#x60;\&quot;http\&quot;&#x60; scheme, but Authlete allows the port number component to be variable in other cases (e.g. in the case of the &#x60;\&quot;https\&quot;&#x60; scheme), too.  | [optional] 
**RequestObjectAudienceChecked** | Pointer to **bool** | The flag indicating whether Authlete checks whether the &#x60;aud&#x60; claim of request objects matches the issuer identifier of this service.  [Section 6.1. Passing a Request Object by Value](https://openid.net/specs/openid-connect-core-1_0.html#JWTRequests) of [OpenID Connect Core 1.0](https://openid.net/specs/openid-connect-core-1_0.html) has the following statement.  &gt; The &#x60;aud&#x60; value SHOULD be or include the OP&#39;s Issuer Identifier URL.  Likewise, [Section 4. Request Object](https://www.rfc-editor.org/rfc/rfc9101.html#section-4) of [RFC 9101](https://www.rfc-editor.org/rfc/rfc9101.html) (The OAuth 2.0 Authorization Framework: JWT-Secured Authorization Request (JAR)) has the following statement.  &gt; The value of aud should be the value of the authorization server (AS) issuer, as defined in [RFC 8414](https://www.rfc-editor.org/rfc/rfc8414.html).  As excerpted above, validation on the &#x60;aud&#x60; claim of request objects is optional. However, if this flag is turned on, Authlete checks whether the &#x60;aud&#x60; claim of request objects matches the issuer identifier of this service and raises an error if they are different.  | [optional] 
**AccessTokenForExternalAttachmentEmbedded** | Pointer to **bool** | The flag indicating whether Authlete generates access tokens for external attachments and embeds them in ID tokens and userinfo responses.  | [optional] 
**AuthorityHints** | Pointer to **[]string** | Identifiers of entities that can issue entity statements for this service. This property corresponds to the &#x60;authority_hints&#x60; property that appears in a self-signed entity statement that is defined in OpenID Connect Federation 1.0.  | [optional] 
**FederationEnabled** | Pointer to **bool** | flag indicating whether this service supports OpenID Connect Federation 1  | [optional] 
**FederationJwks** | Pointer to **string** | JWK Set document containing keys that are used to sign (1) self-signed entity statement of this service and (2) the response from &#x60;signed_jwks_uri&#x60;.  | [optional] 
**FederationSignatureKeyId** | Pointer to **string** | A key ID to identify a JWK used to sign the entity configuration and the signed JWK Set.  | [optional] 
**FederationConfigurationDuration** | Pointer to **int32** | The duration of the entity configuration in seconds.  | [optional] 
**FederationRegistrationEndpoint** | Pointer to **string** | The URI of the federation registration endpoint. This property corresponds to the &#x60;federation_registration_endpoint&#x60; server metadata that is defined in OpenID Connect Federation 1.0.  | [optional] 
**OrganizationName** | Pointer to **string** | The human-readable name representing the organization that operates this service. This property corresponds to the &#x60;organization_name&#x60; server metadata that is defined in OpenID Connect Federation 1.0.  | [optional] 
**PredefinedTransformedClaims** | Pointer to **string** | The transformed claims predefined by this service in JSON format. This property corresponds to the &#x60;transformed_claims_predefined&#x60; server metadata.  | [optional] 
**RefreshTokenIdempotent** | Pointer to **bool** | flag indicating whether refresh token requests with the same refresh token can be made multiple times in quick succession and they can obtain the same renewed refresh token within the short period.  | [optional] 
**SignedJwksUri** | Pointer to **string** | The URI of the endpoint that returns this service&#39;s JWK Set document in the JWT format. This property corresponds to the &#x60;signed_jwks_uri&#x60; server metadata defined in OpenID Connect Federation 1.0.  | [optional] 
**SupportedAttachments** | Pointer to [**[]AttachmentType**](AttachmentType.md) | Supported attachment types. This property corresponds to the {@code attachments_supported} server metadata which was added by the third implementer&#39;s draft of OpenID Connect for Identity Assurance 1.0.  | [optional] 
**SupportedDigestAlgorithms** | Pointer to **[]string** | Supported algorithms used to compute digest values of external attachments. This property corresponds to the &#x60;digest_algorithms_supported&#x60; server metadata which was added by the third implementer&#39;s draft of OpenID Connect for Identity Assurance 1.0.  | [optional] 
**SupportedDocuments** | Pointer to **[]string** | Document types supported by this service. This property corresponds to the &#x60;documents_supported&#x60; server metadata.  | [optional] 
**SupportedDocumentsMethods** | Pointer to **[]string** | validation and verification processes supported by this service. This property corresponds to the &#x60;documents_methods_supported&#x60; server metadata.  The third implementer&#39;s draft of [OpenID Connect for Identity Assurance 1.0](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html) renamed the &#x60;id_documents_verification_methods_supported&#x60; server metadata to &#x60;documents_methods_supported&#x60;.  | [optional] 
**SupportedDocumentsValidationMethods** | Pointer to **[]string** | Document validation methods supported by this service. This property corresponds to the &#x60;documents_validation_methods_supported&#x60; server metadata which was added by the third implementer&#39;s draft of &lt;a href&#x3D; [OpenID Connect for Identity Assurance 1.0](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html)  | [optional] 
**SupportedDocumentsVerificationMethods** | Pointer to **[]string** | Document verification methods supported by this service. This property corresponds to the &#x60;documents_verification_methods_supported&#x60; server metadata which was added by the third implementer&#39;s draft of [OpenID Connect for Identity Assurance 1.0](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html)  | [optional] 
**SupportedElectronicRecords** | Pointer to **[]string** | Electronic record types supported by this service. This property corresponds to the &#x60;electronic_records_supported&#x60; server metadata which was added by the third implementer&#39;s draft of [OpenID Connect for Identity Assurance 1.0](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html)  | [optional] 
**SupportedClientRegistrationTypes** | Pointer to [**[]ClientRegistrationType**](ClientRegistrationType.md) |  | [optional] 
**TokenExchangeByIdentifiableClientsOnly** | Pointer to **bool** | The flag indicating whether to prohibit unidentifiable clients from making token exchange requests.  | [optional] 
**TokenExchangeByConfidentialClientsOnly** | Pointer to **bool** | The flag indicating whether to prohibit public clients from making token exchange requests.  | [optional] 
**TokenExchangeByPermittedClientsOnly** | Pointer to **bool** | The flag indicating whether to prohibit clients that have no explicit permission from making token exchange requests.  | [optional] 
**TokenExchangeEncryptedJwtRejected** | Pointer to **bool** | The flag indicating whether to reject token exchange requests which use encrypted JWTs as input tokens.  | [optional] 
**TokenExchangeUnsignedJwtRejected** | Pointer to **bool** | The flag indicating whether to reject token exchange requests which use unsigned JWTs as input tokens.  | [optional] 
**JwtGrantByIdentifiableClientsOnly** | Pointer to **bool** | The flag indicating whether to prohibit unidentifiable clients from using the grant type \&quot;urn:ietf:params:oauth:grant-type:jwt-bearer\&quot;.  | [optional] 
**JwtGrantEncryptedJwtRejected** | Pointer to **bool** | The flag indicating whether to reject token requests that use an encrypted JWT as an authorization grant with the grant type \&quot;urn:ietf:params:oauth:grant-type:jwt-bearer\&quot;.  | [optional] 
**JwtGrantUnsignedJwtRejected** | Pointer to **bool** | The flag indicating whether to reject token requests that use an unsigned JWT as an authorization grant with the grant type \&quot;urn:ietf:params:oauth:grant-type:jwt-bearer\&quot;.  | [optional] 
**DcrDuplicateSoftwareIdBlocked** | Pointer to **bool** | The flag indicating whether to block DCR (Dynamic Client Registration) requests whose \&quot;software_id\&quot; has already been used previously.  | [optional] 
**TrustAnchors** | Pointer to [**[]TrustAnchor**](TrustAnchor.md) | The trust anchors that are referenced when this service resolves trust chains of relying parties.  If this property is empty, client registration fails regardless of whether its type is &#x60;automatic&#x60; or &#x60;explicit&#x60;. It means that OpenID Connect Federation 1.0 does not work.  | [optional] 
**OpenidDroppedOnRefreshWithoutOfflineAccess** | Pointer to **bool** | The flag indicating whether the openid scope should be dropped from scopes list assigned to access token issued when a refresh token grant is used.  | [optional] 
**SupportedDocumentsCheckMethods** | Pointer to **[]string** | Supported document check methods. This property corresponds to the &#x60;documents_check_methods_supported&#x60; server metadata which was added by the fourth implementer&#39;s draft of OpenID Connect for Identity Assurance 1.0.  | [optional] 
**RsResponseSigned** | Pointer to **bool** | The flag indicating whether this service signs responses from the resource server.  | [optional] 
**CnonceDuration** | Pointer to **int64** | The duration of &#x60;c_nonce&#x60;.  | [optional] 
**DpopNonceRequired** | Pointer to **bool** | Whether to require DPoP proof JWTs to include the &#x60;nonce&#x60; claim whenever they are presented.  | [optional] 
**VerifiableCredentialsEnabled** | Pointer to **bool** | Get the flag indicating whether the feature of Verifiable Credentials for this service is enabled or not.  | [optional] 
**CredentialJwksUri** | Pointer to **string** | The URL at which the JWK Set document of the credential issuer is exposed.  | [optional] 
**CredentialOfferDuration** | Pointer to **int64** | The default duration of credential offers in seconds.  | [optional] 
**DpopNonceDuration** | Pointer to **int64** | The duration of nonce values for DPoP proof JWTs in seconds.  | [optional] 
**PreAuthorizedGrantAnonymousAccessSupported** | Pointer to **bool** | The flag indicating whether token requests using the pre-authorized code grant flow by unidentifiable clients are allowed.  | [optional] 
**CredentialTransactionDuration** | Pointer to **int64** | The duration of transaction ID in seconds that may be issued as a result of a credential request or a batch credential request.  | [optional] 
**IntrospectionSignatureKeyId** | Pointer to **string** | The key ID of the key for signing introspection responses.  | [optional] 
**ResourceSignatureKeyId** | Pointer to **string** | The key ID of the key for signing introspection responses.  | [optional] 
**UserPinLength** | Pointer to **int32** | The default length of user PINs.  | [optional] 
**SupportedPromptValues** | Pointer to [**[]Prompt**](Prompt.md) | The supported &#x60;prompt&#x60; values.  | [optional] 
**IdTokenReissuable** | Pointer to **bool** | The flag indicating whether to enable the feature of ID token reissuance in the refresh token flow.  | [optional] 
**CredentialJwks** | Pointer to **string** | The JWK Set document containing private keys that are used to sign verifiable credentials.  | [optional] 
**FapiModes** | Pointer to [**[]FapiMode**](FapiMode.md) | FAPI modes for this service.  When the value of this property is not &#x60;null&#x60;, Authlete always processes requests to this service based on the specified FAPI modes if the FAPI feature is enabled in Authlete and the FAPI profile is supported by this service.  For instance, when this property is set to an array containing &#x60;FAPI1_ADVANCED&#x60; only, Authlete always processes requests to this service based on \&quot;Financial-grade API Security Profile 1.0 - Part 2: Advanced\&quot; if the FAPI feature is enabled in Authlete and the FAPI profile is supported by this service.  | [optional] 
**CredentialDuration** | Pointer to **int64** | The default duration of verifiable credentials in seconds.  | [optional] 
**CredentialIssuerMetadata** | Pointer to [**CredentialIssuerMetadata**](CredentialIssuerMetadata.md) |  | [optional] 
**IdTokenAudType** | Pointer to **string** | The type of the &#x60;aud&#x60; claim in ID tokens. | [optional] 

## Methods

### NewService

`func NewService() *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *Service) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Service) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Service) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Service) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetServiceName

`func (o *Service) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Service) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Service) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *Service) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetIssuer

`func (o *Service) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Service) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Service) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *Service) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetDescription

`func (o *Service) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Service) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Service) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Service) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApiKey

`func (o *Service) GetApiKey() int64`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *Service) GetApiKeyOk() (*int64, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *Service) SetApiKey(v int64)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *Service) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetClientIdAliasEnabled

`func (o *Service) GetClientIdAliasEnabled() bool`

GetClientIdAliasEnabled returns the ClientIdAliasEnabled field if non-nil, zero value otherwise.

### GetClientIdAliasEnabledOk

`func (o *Service) GetClientIdAliasEnabledOk() (*bool, bool)`

GetClientIdAliasEnabledOk returns a tuple with the ClientIdAliasEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasEnabled

`func (o *Service) SetClientIdAliasEnabled(v bool)`

SetClientIdAliasEnabled sets ClientIdAliasEnabled field to given value.

### HasClientIdAliasEnabled

`func (o *Service) HasClientIdAliasEnabled() bool`

HasClientIdAliasEnabled returns a boolean if a field has been set.

### GetMetadata

`func (o *Service) GetMetadata() []Pair`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Service) GetMetadataOk() (*[]Pair, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Service) SetMetadata(v []Pair)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Service) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Service) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Service) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Service) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Service) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Service) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Service) GetModifiedAtOk() (*int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Service) SetModifiedAt(v int64)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Service) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetAuthenticationCallbackEndpoint

`func (o *Service) GetAuthenticationCallbackEndpoint() string`

GetAuthenticationCallbackEndpoint returns the AuthenticationCallbackEndpoint field if non-nil, zero value otherwise.

### GetAuthenticationCallbackEndpointOk

`func (o *Service) GetAuthenticationCallbackEndpointOk() (*string, bool)`

GetAuthenticationCallbackEndpointOk returns a tuple with the AuthenticationCallbackEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCallbackEndpoint

`func (o *Service) SetAuthenticationCallbackEndpoint(v string)`

SetAuthenticationCallbackEndpoint sets AuthenticationCallbackEndpoint field to given value.

### HasAuthenticationCallbackEndpoint

`func (o *Service) HasAuthenticationCallbackEndpoint() bool`

HasAuthenticationCallbackEndpoint returns a boolean if a field has been set.

### GetAuthenticationCallbackApiKey

`func (o *Service) GetAuthenticationCallbackApiKey() string`

GetAuthenticationCallbackApiKey returns the AuthenticationCallbackApiKey field if non-nil, zero value otherwise.

### GetAuthenticationCallbackApiKeyOk

`func (o *Service) GetAuthenticationCallbackApiKeyOk() (*string, bool)`

GetAuthenticationCallbackApiKeyOk returns a tuple with the AuthenticationCallbackApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCallbackApiKey

`func (o *Service) SetAuthenticationCallbackApiKey(v string)`

SetAuthenticationCallbackApiKey sets AuthenticationCallbackApiKey field to given value.

### HasAuthenticationCallbackApiKey

`func (o *Service) HasAuthenticationCallbackApiKey() bool`

HasAuthenticationCallbackApiKey returns a boolean if a field has been set.

### GetAuthenticationCallbackApiSecret

`func (o *Service) GetAuthenticationCallbackApiSecret() string`

GetAuthenticationCallbackApiSecret returns the AuthenticationCallbackApiSecret field if non-nil, zero value otherwise.

### GetAuthenticationCallbackApiSecretOk

`func (o *Service) GetAuthenticationCallbackApiSecretOk() (*string, bool)`

GetAuthenticationCallbackApiSecretOk returns a tuple with the AuthenticationCallbackApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCallbackApiSecret

`func (o *Service) SetAuthenticationCallbackApiSecret(v string)`

SetAuthenticationCallbackApiSecret sets AuthenticationCallbackApiSecret field to given value.

### HasAuthenticationCallbackApiSecret

`func (o *Service) HasAuthenticationCallbackApiSecret() bool`

HasAuthenticationCallbackApiSecret returns a boolean if a field has been set.

### GetSupportedAcrs

`func (o *Service) GetSupportedAcrs() []string`

GetSupportedAcrs returns the SupportedAcrs field if non-nil, zero value otherwise.

### GetSupportedAcrsOk

`func (o *Service) GetSupportedAcrsOk() (*[]string, bool)`

GetSupportedAcrsOk returns a tuple with the SupportedAcrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAcrs

`func (o *Service) SetSupportedAcrs(v []string)`

SetSupportedAcrs sets SupportedAcrs field to given value.

### HasSupportedAcrs

`func (o *Service) HasSupportedAcrs() bool`

HasSupportedAcrs returns a boolean if a field has been set.

### GetSupportedGrantTypes

`func (o *Service) GetSupportedGrantTypes() []GrantType`

GetSupportedGrantTypes returns the SupportedGrantTypes field if non-nil, zero value otherwise.

### GetSupportedGrantTypesOk

`func (o *Service) GetSupportedGrantTypesOk() (*[]GrantType, bool)`

GetSupportedGrantTypesOk returns a tuple with the SupportedGrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedGrantTypes

`func (o *Service) SetSupportedGrantTypes(v []GrantType)`

SetSupportedGrantTypes sets SupportedGrantTypes field to given value.

### HasSupportedGrantTypes

`func (o *Service) HasSupportedGrantTypes() bool`

HasSupportedGrantTypes returns a boolean if a field has been set.

### GetSupportedResponseTypes

`func (o *Service) GetSupportedResponseTypes() []ResponseType`

GetSupportedResponseTypes returns the SupportedResponseTypes field if non-nil, zero value otherwise.

### GetSupportedResponseTypesOk

`func (o *Service) GetSupportedResponseTypesOk() (*[]ResponseType, bool)`

GetSupportedResponseTypesOk returns a tuple with the SupportedResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedResponseTypes

`func (o *Service) SetSupportedResponseTypes(v []ResponseType)`

SetSupportedResponseTypes sets SupportedResponseTypes field to given value.

### HasSupportedResponseTypes

`func (o *Service) HasSupportedResponseTypes() bool`

HasSupportedResponseTypes returns a boolean if a field has been set.

### GetSupportedAuthorizationDetailsTypes

`func (o *Service) GetSupportedAuthorizationDetailsTypes() []string`

GetSupportedAuthorizationDetailsTypes returns the SupportedAuthorizationDetailsTypes field if non-nil, zero value otherwise.

### GetSupportedAuthorizationDetailsTypesOk

`func (o *Service) GetSupportedAuthorizationDetailsTypesOk() (*[]string, bool)`

GetSupportedAuthorizationDetailsTypesOk returns a tuple with the SupportedAuthorizationDetailsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAuthorizationDetailsTypes

`func (o *Service) SetSupportedAuthorizationDetailsTypes(v []string)`

SetSupportedAuthorizationDetailsTypes sets SupportedAuthorizationDetailsTypes field to given value.

### HasSupportedAuthorizationDetailsTypes

`func (o *Service) HasSupportedAuthorizationDetailsTypes() bool`

HasSupportedAuthorizationDetailsTypes returns a boolean if a field has been set.

### GetSupportedServiceProfiles

`func (o *Service) GetSupportedServiceProfiles() []ServiceProfile`

GetSupportedServiceProfiles returns the SupportedServiceProfiles field if non-nil, zero value otherwise.

### GetSupportedServiceProfilesOk

`func (o *Service) GetSupportedServiceProfilesOk() (*[]ServiceProfile, bool)`

GetSupportedServiceProfilesOk returns a tuple with the SupportedServiceProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedServiceProfiles

`func (o *Service) SetSupportedServiceProfiles(v []ServiceProfile)`

SetSupportedServiceProfiles sets SupportedServiceProfiles field to given value.

### HasSupportedServiceProfiles

`func (o *Service) HasSupportedServiceProfiles() bool`

HasSupportedServiceProfiles returns a boolean if a field has been set.

### GetErrorDescriptionOmitted

`func (o *Service) GetErrorDescriptionOmitted() bool`

GetErrorDescriptionOmitted returns the ErrorDescriptionOmitted field if non-nil, zero value otherwise.

### GetErrorDescriptionOmittedOk

`func (o *Service) GetErrorDescriptionOmittedOk() (*bool, bool)`

GetErrorDescriptionOmittedOk returns a tuple with the ErrorDescriptionOmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescriptionOmitted

`func (o *Service) SetErrorDescriptionOmitted(v bool)`

SetErrorDescriptionOmitted sets ErrorDescriptionOmitted field to given value.

### HasErrorDescriptionOmitted

`func (o *Service) HasErrorDescriptionOmitted() bool`

HasErrorDescriptionOmitted returns a boolean if a field has been set.

### GetErrorUriOmitted

`func (o *Service) GetErrorUriOmitted() bool`

GetErrorUriOmitted returns the ErrorUriOmitted field if non-nil, zero value otherwise.

### GetErrorUriOmittedOk

`func (o *Service) GetErrorUriOmittedOk() (*bool, bool)`

GetErrorUriOmittedOk returns a tuple with the ErrorUriOmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUriOmitted

`func (o *Service) SetErrorUriOmitted(v bool)`

SetErrorUriOmitted sets ErrorUriOmitted field to given value.

### HasErrorUriOmitted

`func (o *Service) HasErrorUriOmitted() bool`

HasErrorUriOmitted returns a boolean if a field has been set.

### GetAuthorizationEndpoint

`func (o *Service) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *Service) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *Service) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.

### HasAuthorizationEndpoint

`func (o *Service) HasAuthorizationEndpoint() bool`

HasAuthorizationEndpoint returns a boolean if a field has been set.

### GetDirectAuthorizationEndpointEnabled

`func (o *Service) GetDirectAuthorizationEndpointEnabled() bool`

GetDirectAuthorizationEndpointEnabled returns the DirectAuthorizationEndpointEnabled field if non-nil, zero value otherwise.

### GetDirectAuthorizationEndpointEnabledOk

`func (o *Service) GetDirectAuthorizationEndpointEnabledOk() (*bool, bool)`

GetDirectAuthorizationEndpointEnabledOk returns a tuple with the DirectAuthorizationEndpointEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectAuthorizationEndpointEnabled

`func (o *Service) SetDirectAuthorizationEndpointEnabled(v bool)`

SetDirectAuthorizationEndpointEnabled sets DirectAuthorizationEndpointEnabled field to given value.

### HasDirectAuthorizationEndpointEnabled

`func (o *Service) HasDirectAuthorizationEndpointEnabled() bool`

HasDirectAuthorizationEndpointEnabled returns a boolean if a field has been set.

### GetSupportedUiLocales

`func (o *Service) GetSupportedUiLocales() []string`

GetSupportedUiLocales returns the SupportedUiLocales field if non-nil, zero value otherwise.

### GetSupportedUiLocalesOk

`func (o *Service) GetSupportedUiLocalesOk() (*[]string, bool)`

GetSupportedUiLocalesOk returns a tuple with the SupportedUiLocales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedUiLocales

`func (o *Service) SetSupportedUiLocales(v []string)`

SetSupportedUiLocales sets SupportedUiLocales field to given value.

### HasSupportedUiLocales

`func (o *Service) HasSupportedUiLocales() bool`

HasSupportedUiLocales returns a boolean if a field has been set.

### GetSupportedDisplays

`func (o *Service) GetSupportedDisplays() []Display`

GetSupportedDisplays returns the SupportedDisplays field if non-nil, zero value otherwise.

### GetSupportedDisplaysOk

`func (o *Service) GetSupportedDisplaysOk() (*[]Display, bool)`

GetSupportedDisplaysOk returns a tuple with the SupportedDisplays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDisplays

`func (o *Service) SetSupportedDisplays(v []Display)`

SetSupportedDisplays sets SupportedDisplays field to given value.

### HasSupportedDisplays

`func (o *Service) HasSupportedDisplays() bool`

HasSupportedDisplays returns a boolean if a field has been set.

### GetPkceRequired

`func (o *Service) GetPkceRequired() bool`

GetPkceRequired returns the PkceRequired field if non-nil, zero value otherwise.

### GetPkceRequiredOk

`func (o *Service) GetPkceRequiredOk() (*bool, bool)`

GetPkceRequiredOk returns a tuple with the PkceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceRequired

`func (o *Service) SetPkceRequired(v bool)`

SetPkceRequired sets PkceRequired field to given value.

### HasPkceRequired

`func (o *Service) HasPkceRequired() bool`

HasPkceRequired returns a boolean if a field has been set.

### GetPkceS256Required

`func (o *Service) GetPkceS256Required() bool`

GetPkceS256Required returns the PkceS256Required field if non-nil, zero value otherwise.

### GetPkceS256RequiredOk

`func (o *Service) GetPkceS256RequiredOk() (*bool, bool)`

GetPkceS256RequiredOk returns a tuple with the PkceS256Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceS256Required

`func (o *Service) SetPkceS256Required(v bool)`

SetPkceS256Required sets PkceS256Required field to given value.

### HasPkceS256Required

`func (o *Service) HasPkceS256Required() bool`

HasPkceS256Required returns a boolean if a field has been set.

### GetAuthorizationResponseDuration

`func (o *Service) GetAuthorizationResponseDuration() int64`

GetAuthorizationResponseDuration returns the AuthorizationResponseDuration field if non-nil, zero value otherwise.

### GetAuthorizationResponseDurationOk

`func (o *Service) GetAuthorizationResponseDurationOk() (*int64, bool)`

GetAuthorizationResponseDurationOk returns a tuple with the AuthorizationResponseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationResponseDuration

`func (o *Service) SetAuthorizationResponseDuration(v int64)`

SetAuthorizationResponseDuration sets AuthorizationResponseDuration field to given value.

### HasAuthorizationResponseDuration

`func (o *Service) HasAuthorizationResponseDuration() bool`

HasAuthorizationResponseDuration returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *Service) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *Service) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *Service) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *Service) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetDirectTokenEndpointEnabled

`func (o *Service) GetDirectTokenEndpointEnabled() bool`

GetDirectTokenEndpointEnabled returns the DirectTokenEndpointEnabled field if non-nil, zero value otherwise.

### GetDirectTokenEndpointEnabledOk

`func (o *Service) GetDirectTokenEndpointEnabledOk() (*bool, bool)`

GetDirectTokenEndpointEnabledOk returns a tuple with the DirectTokenEndpointEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectTokenEndpointEnabled

`func (o *Service) SetDirectTokenEndpointEnabled(v bool)`

SetDirectTokenEndpointEnabled sets DirectTokenEndpointEnabled field to given value.

### HasDirectTokenEndpointEnabled

`func (o *Service) HasDirectTokenEndpointEnabled() bool`

HasDirectTokenEndpointEnabled returns a boolean if a field has been set.

### GetSupportedTokenAuthMethods

`func (o *Service) GetSupportedTokenAuthMethods() []ClientAuthMethod`

GetSupportedTokenAuthMethods returns the SupportedTokenAuthMethods field if non-nil, zero value otherwise.

### GetSupportedTokenAuthMethodsOk

`func (o *Service) GetSupportedTokenAuthMethodsOk() (*[]ClientAuthMethod, bool)`

GetSupportedTokenAuthMethodsOk returns a tuple with the SupportedTokenAuthMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedTokenAuthMethods

`func (o *Service) SetSupportedTokenAuthMethods(v []ClientAuthMethod)`

SetSupportedTokenAuthMethods sets SupportedTokenAuthMethods field to given value.

### HasSupportedTokenAuthMethods

`func (o *Service) HasSupportedTokenAuthMethods() bool`

HasSupportedTokenAuthMethods returns a boolean if a field has been set.

### GetMissingClientIdAllowed

`func (o *Service) GetMissingClientIdAllowed() bool`

GetMissingClientIdAllowed returns the MissingClientIdAllowed field if non-nil, zero value otherwise.

### GetMissingClientIdAllowedOk

`func (o *Service) GetMissingClientIdAllowedOk() (*bool, bool)`

GetMissingClientIdAllowedOk returns a tuple with the MissingClientIdAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingClientIdAllowed

`func (o *Service) SetMissingClientIdAllowed(v bool)`

SetMissingClientIdAllowed sets MissingClientIdAllowed field to given value.

### HasMissingClientIdAllowed

`func (o *Service) HasMissingClientIdAllowed() bool`

HasMissingClientIdAllowed returns a boolean if a field has been set.

### GetRevocationEndpoint

`func (o *Service) GetRevocationEndpoint() string`

GetRevocationEndpoint returns the RevocationEndpoint field if non-nil, zero value otherwise.

### GetRevocationEndpointOk

`func (o *Service) GetRevocationEndpointOk() (*string, bool)`

GetRevocationEndpointOk returns a tuple with the RevocationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEndpoint

`func (o *Service) SetRevocationEndpoint(v string)`

SetRevocationEndpoint sets RevocationEndpoint field to given value.

### HasRevocationEndpoint

`func (o *Service) HasRevocationEndpoint() bool`

HasRevocationEndpoint returns a boolean if a field has been set.

### GetDirectRevocationEndpointEnabled

`func (o *Service) GetDirectRevocationEndpointEnabled() bool`

GetDirectRevocationEndpointEnabled returns the DirectRevocationEndpointEnabled field if non-nil, zero value otherwise.

### GetDirectRevocationEndpointEnabledOk

`func (o *Service) GetDirectRevocationEndpointEnabledOk() (*bool, bool)`

GetDirectRevocationEndpointEnabledOk returns a tuple with the DirectRevocationEndpointEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectRevocationEndpointEnabled

`func (o *Service) SetDirectRevocationEndpointEnabled(v bool)`

SetDirectRevocationEndpointEnabled sets DirectRevocationEndpointEnabled field to given value.

### HasDirectRevocationEndpointEnabled

`func (o *Service) HasDirectRevocationEndpointEnabled() bool`

HasDirectRevocationEndpointEnabled returns a boolean if a field has been set.

### GetSupportedRevocationAuthMethods

`func (o *Service) GetSupportedRevocationAuthMethods() []ClientAuthMethod`

GetSupportedRevocationAuthMethods returns the SupportedRevocationAuthMethods field if non-nil, zero value otherwise.

### GetSupportedRevocationAuthMethodsOk

`func (o *Service) GetSupportedRevocationAuthMethodsOk() (*[]ClientAuthMethod, bool)`

GetSupportedRevocationAuthMethodsOk returns a tuple with the SupportedRevocationAuthMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedRevocationAuthMethods

`func (o *Service) SetSupportedRevocationAuthMethods(v []ClientAuthMethod)`

SetSupportedRevocationAuthMethods sets SupportedRevocationAuthMethods field to given value.

### HasSupportedRevocationAuthMethods

`func (o *Service) HasSupportedRevocationAuthMethods() bool`

HasSupportedRevocationAuthMethods returns a boolean if a field has been set.

### GetIntrospectionEndpoint

`func (o *Service) GetIntrospectionEndpoint() string`

GetIntrospectionEndpoint returns the IntrospectionEndpoint field if non-nil, zero value otherwise.

### GetIntrospectionEndpointOk

`func (o *Service) GetIntrospectionEndpointOk() (*string, bool)`

GetIntrospectionEndpointOk returns a tuple with the IntrospectionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionEndpoint

`func (o *Service) SetIntrospectionEndpoint(v string)`

SetIntrospectionEndpoint sets IntrospectionEndpoint field to given value.

### HasIntrospectionEndpoint

`func (o *Service) HasIntrospectionEndpoint() bool`

HasIntrospectionEndpoint returns a boolean if a field has been set.

### GetDirectIntrospectionEndpointEnabled

`func (o *Service) GetDirectIntrospectionEndpointEnabled() bool`

GetDirectIntrospectionEndpointEnabled returns the DirectIntrospectionEndpointEnabled field if non-nil, zero value otherwise.

### GetDirectIntrospectionEndpointEnabledOk

`func (o *Service) GetDirectIntrospectionEndpointEnabledOk() (*bool, bool)`

GetDirectIntrospectionEndpointEnabledOk returns a tuple with the DirectIntrospectionEndpointEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectIntrospectionEndpointEnabled

`func (o *Service) SetDirectIntrospectionEndpointEnabled(v bool)`

SetDirectIntrospectionEndpointEnabled sets DirectIntrospectionEndpointEnabled field to given value.

### HasDirectIntrospectionEndpointEnabled

`func (o *Service) HasDirectIntrospectionEndpointEnabled() bool`

HasDirectIntrospectionEndpointEnabled returns a boolean if a field has been set.

### GetSupportedIntrospectionAuthMethods

`func (o *Service) GetSupportedIntrospectionAuthMethods() []ClientAuthMethod`

GetSupportedIntrospectionAuthMethods returns the SupportedIntrospectionAuthMethods field if non-nil, zero value otherwise.

### GetSupportedIntrospectionAuthMethodsOk

`func (o *Service) GetSupportedIntrospectionAuthMethodsOk() (*[]ClientAuthMethod, bool)`

GetSupportedIntrospectionAuthMethodsOk returns a tuple with the SupportedIntrospectionAuthMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedIntrospectionAuthMethods

`func (o *Service) SetSupportedIntrospectionAuthMethods(v []ClientAuthMethod)`

SetSupportedIntrospectionAuthMethods sets SupportedIntrospectionAuthMethods field to given value.

### HasSupportedIntrospectionAuthMethods

`func (o *Service) HasSupportedIntrospectionAuthMethods() bool`

HasSupportedIntrospectionAuthMethods returns a boolean if a field has been set.

### GetPushedAuthReqEndpoint

`func (o *Service) GetPushedAuthReqEndpoint() string`

GetPushedAuthReqEndpoint returns the PushedAuthReqEndpoint field if non-nil, zero value otherwise.

### GetPushedAuthReqEndpointOk

`func (o *Service) GetPushedAuthReqEndpointOk() (*string, bool)`

GetPushedAuthReqEndpointOk returns a tuple with the PushedAuthReqEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAuthReqEndpoint

`func (o *Service) SetPushedAuthReqEndpoint(v string)`

SetPushedAuthReqEndpoint sets PushedAuthReqEndpoint field to given value.

### HasPushedAuthReqEndpoint

`func (o *Service) HasPushedAuthReqEndpoint() bool`

HasPushedAuthReqEndpoint returns a boolean if a field has been set.

### GetPushedAuthReqDuration

`func (o *Service) GetPushedAuthReqDuration() int64`

GetPushedAuthReqDuration returns the PushedAuthReqDuration field if non-nil, zero value otherwise.

### GetPushedAuthReqDurationOk

`func (o *Service) GetPushedAuthReqDurationOk() (*int64, bool)`

GetPushedAuthReqDurationOk returns a tuple with the PushedAuthReqDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAuthReqDuration

`func (o *Service) SetPushedAuthReqDuration(v int64)`

SetPushedAuthReqDuration sets PushedAuthReqDuration field to given value.

### HasPushedAuthReqDuration

`func (o *Service) HasPushedAuthReqDuration() bool`

HasPushedAuthReqDuration returns a boolean if a field has been set.

### GetParRequired

`func (o *Service) GetParRequired() bool`

GetParRequired returns the ParRequired field if non-nil, zero value otherwise.

### GetParRequiredOk

`func (o *Service) GetParRequiredOk() (*bool, bool)`

GetParRequiredOk returns a tuple with the ParRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParRequired

`func (o *Service) SetParRequired(v bool)`

SetParRequired sets ParRequired field to given value.

### HasParRequired

`func (o *Service) HasParRequired() bool`

HasParRequired returns a boolean if a field has been set.

### GetRequestObjectRequired

`func (o *Service) GetRequestObjectRequired() bool`

GetRequestObjectRequired returns the RequestObjectRequired field if non-nil, zero value otherwise.

### GetRequestObjectRequiredOk

`func (o *Service) GetRequestObjectRequiredOk() (*bool, bool)`

GetRequestObjectRequiredOk returns a tuple with the RequestObjectRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectRequired

`func (o *Service) SetRequestObjectRequired(v bool)`

SetRequestObjectRequired sets RequestObjectRequired field to given value.

### HasRequestObjectRequired

`func (o *Service) HasRequestObjectRequired() bool`

HasRequestObjectRequired returns a boolean if a field has been set.

### GetTraditionalRequestObjectProcessingApplied

`func (o *Service) GetTraditionalRequestObjectProcessingApplied() bool`

GetTraditionalRequestObjectProcessingApplied returns the TraditionalRequestObjectProcessingApplied field if non-nil, zero value otherwise.

### GetTraditionalRequestObjectProcessingAppliedOk

`func (o *Service) GetTraditionalRequestObjectProcessingAppliedOk() (*bool, bool)`

GetTraditionalRequestObjectProcessingAppliedOk returns a tuple with the TraditionalRequestObjectProcessingApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraditionalRequestObjectProcessingApplied

`func (o *Service) SetTraditionalRequestObjectProcessingApplied(v bool)`

SetTraditionalRequestObjectProcessingApplied sets TraditionalRequestObjectProcessingApplied field to given value.

### HasTraditionalRequestObjectProcessingApplied

`func (o *Service) HasTraditionalRequestObjectProcessingApplied() bool`

HasTraditionalRequestObjectProcessingApplied returns a boolean if a field has been set.

### GetMutualTlsValidatePkiCertChain

`func (o *Service) GetMutualTlsValidatePkiCertChain() bool`

GetMutualTlsValidatePkiCertChain returns the MutualTlsValidatePkiCertChain field if non-nil, zero value otherwise.

### GetMutualTlsValidatePkiCertChainOk

`func (o *Service) GetMutualTlsValidatePkiCertChainOk() (*bool, bool)`

GetMutualTlsValidatePkiCertChainOk returns a tuple with the MutualTlsValidatePkiCertChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualTlsValidatePkiCertChain

`func (o *Service) SetMutualTlsValidatePkiCertChain(v bool)`

SetMutualTlsValidatePkiCertChain sets MutualTlsValidatePkiCertChain field to given value.

### HasMutualTlsValidatePkiCertChain

`func (o *Service) HasMutualTlsValidatePkiCertChain() bool`

HasMutualTlsValidatePkiCertChain returns a boolean if a field has been set.

### GetTrustedRootCertificates

`func (o *Service) GetTrustedRootCertificates() []string`

GetTrustedRootCertificates returns the TrustedRootCertificates field if non-nil, zero value otherwise.

### GetTrustedRootCertificatesOk

`func (o *Service) GetTrustedRootCertificatesOk() (*[]string, bool)`

GetTrustedRootCertificatesOk returns a tuple with the TrustedRootCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedRootCertificates

`func (o *Service) SetTrustedRootCertificates(v []string)`

SetTrustedRootCertificates sets TrustedRootCertificates field to given value.

### HasTrustedRootCertificates

`func (o *Service) HasTrustedRootCertificates() bool`

HasTrustedRootCertificates returns a boolean if a field has been set.

### GetMtlsEndpointAliases

`func (o *Service) GetMtlsEndpointAliases() []NamedUri`

GetMtlsEndpointAliases returns the MtlsEndpointAliases field if non-nil, zero value otherwise.

### GetMtlsEndpointAliasesOk

`func (o *Service) GetMtlsEndpointAliasesOk() (*[]NamedUri, bool)`

GetMtlsEndpointAliasesOk returns a tuple with the MtlsEndpointAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsEndpointAliases

`func (o *Service) SetMtlsEndpointAliases(v []NamedUri)`

SetMtlsEndpointAliases sets MtlsEndpointAliases field to given value.

### HasMtlsEndpointAliases

`func (o *Service) HasMtlsEndpointAliases() bool`

HasMtlsEndpointAliases returns a boolean if a field has been set.

### GetAccessTokenType

`func (o *Service) GetAccessTokenType() string`

GetAccessTokenType returns the AccessTokenType field if non-nil, zero value otherwise.

### GetAccessTokenTypeOk

`func (o *Service) GetAccessTokenTypeOk() (*string, bool)`

GetAccessTokenTypeOk returns a tuple with the AccessTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenType

`func (o *Service) SetAccessTokenType(v string)`

SetAccessTokenType sets AccessTokenType field to given value.

### HasAccessTokenType

`func (o *Service) HasAccessTokenType() bool`

HasAccessTokenType returns a boolean if a field has been set.

### GetTlsClientCertificateBoundAccessTokens

`func (o *Service) GetTlsClientCertificateBoundAccessTokens() bool`

GetTlsClientCertificateBoundAccessTokens returns the TlsClientCertificateBoundAccessTokens field if non-nil, zero value otherwise.

### GetTlsClientCertificateBoundAccessTokensOk

`func (o *Service) GetTlsClientCertificateBoundAccessTokensOk() (*bool, bool)`

GetTlsClientCertificateBoundAccessTokensOk returns a tuple with the TlsClientCertificateBoundAccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientCertificateBoundAccessTokens

`func (o *Service) SetTlsClientCertificateBoundAccessTokens(v bool)`

SetTlsClientCertificateBoundAccessTokens sets TlsClientCertificateBoundAccessTokens field to given value.

### HasTlsClientCertificateBoundAccessTokens

`func (o *Service) HasTlsClientCertificateBoundAccessTokens() bool`

HasTlsClientCertificateBoundAccessTokens returns a boolean if a field has been set.

### GetAccessTokenDuration

`func (o *Service) GetAccessTokenDuration() int64`

GetAccessTokenDuration returns the AccessTokenDuration field if non-nil, zero value otherwise.

### GetAccessTokenDurationOk

`func (o *Service) GetAccessTokenDurationOk() (*int64, bool)`

GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenDuration

`func (o *Service) SetAccessTokenDuration(v int64)`

SetAccessTokenDuration sets AccessTokenDuration field to given value.

### HasAccessTokenDuration

`func (o *Service) HasAccessTokenDuration() bool`

HasAccessTokenDuration returns a boolean if a field has been set.

### GetSingleAccessTokenPerSubject

`func (o *Service) GetSingleAccessTokenPerSubject() bool`

GetSingleAccessTokenPerSubject returns the SingleAccessTokenPerSubject field if non-nil, zero value otherwise.

### GetSingleAccessTokenPerSubjectOk

`func (o *Service) GetSingleAccessTokenPerSubjectOk() (*bool, bool)`

GetSingleAccessTokenPerSubjectOk returns a tuple with the SingleAccessTokenPerSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleAccessTokenPerSubject

`func (o *Service) SetSingleAccessTokenPerSubject(v bool)`

SetSingleAccessTokenPerSubject sets SingleAccessTokenPerSubject field to given value.

### HasSingleAccessTokenPerSubject

`func (o *Service) HasSingleAccessTokenPerSubject() bool`

HasSingleAccessTokenPerSubject returns a boolean if a field has been set.

### GetAccessTokenSignAlg

`func (o *Service) GetAccessTokenSignAlg() JwsAlg`

GetAccessTokenSignAlg returns the AccessTokenSignAlg field if non-nil, zero value otherwise.

### GetAccessTokenSignAlgOk

`func (o *Service) GetAccessTokenSignAlgOk() (*JwsAlg, bool)`

GetAccessTokenSignAlgOk returns a tuple with the AccessTokenSignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenSignAlg

`func (o *Service) SetAccessTokenSignAlg(v JwsAlg)`

SetAccessTokenSignAlg sets AccessTokenSignAlg field to given value.

### HasAccessTokenSignAlg

`func (o *Service) HasAccessTokenSignAlg() bool`

HasAccessTokenSignAlg returns a boolean if a field has been set.

### GetAccessTokenSignatureKeyId

`func (o *Service) GetAccessTokenSignatureKeyId() string`

GetAccessTokenSignatureKeyId returns the AccessTokenSignatureKeyId field if non-nil, zero value otherwise.

### GetAccessTokenSignatureKeyIdOk

`func (o *Service) GetAccessTokenSignatureKeyIdOk() (*string, bool)`

GetAccessTokenSignatureKeyIdOk returns a tuple with the AccessTokenSignatureKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenSignatureKeyId

`func (o *Service) SetAccessTokenSignatureKeyId(v string)`

SetAccessTokenSignatureKeyId sets AccessTokenSignatureKeyId field to given value.

### HasAccessTokenSignatureKeyId

`func (o *Service) HasAccessTokenSignatureKeyId() bool`

HasAccessTokenSignatureKeyId returns a boolean if a field has been set.

### GetRefreshTokenDuration

`func (o *Service) GetRefreshTokenDuration() int64`

GetRefreshTokenDuration returns the RefreshTokenDuration field if non-nil, zero value otherwise.

### GetRefreshTokenDurationOk

`func (o *Service) GetRefreshTokenDurationOk() (*int64, bool)`

GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDuration

`func (o *Service) SetRefreshTokenDuration(v int64)`

SetRefreshTokenDuration sets RefreshTokenDuration field to given value.

### HasRefreshTokenDuration

`func (o *Service) HasRefreshTokenDuration() bool`

HasRefreshTokenDuration returns a boolean if a field has been set.

### GetRefreshTokenDurationKept

`func (o *Service) GetRefreshTokenDurationKept() bool`

GetRefreshTokenDurationKept returns the RefreshTokenDurationKept field if non-nil, zero value otherwise.

### GetRefreshTokenDurationKeptOk

`func (o *Service) GetRefreshTokenDurationKeptOk() (*bool, bool)`

GetRefreshTokenDurationKeptOk returns a tuple with the RefreshTokenDurationKept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDurationKept

`func (o *Service) SetRefreshTokenDurationKept(v bool)`

SetRefreshTokenDurationKept sets RefreshTokenDurationKept field to given value.

### HasRefreshTokenDurationKept

`func (o *Service) HasRefreshTokenDurationKept() bool`

HasRefreshTokenDurationKept returns a boolean if a field has been set.

### GetRefreshTokenDurationReset

`func (o *Service) GetRefreshTokenDurationReset() bool`

GetRefreshTokenDurationReset returns the RefreshTokenDurationReset field if non-nil, zero value otherwise.

### GetRefreshTokenDurationResetOk

`func (o *Service) GetRefreshTokenDurationResetOk() (*bool, bool)`

GetRefreshTokenDurationResetOk returns a tuple with the RefreshTokenDurationReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDurationReset

`func (o *Service) SetRefreshTokenDurationReset(v bool)`

SetRefreshTokenDurationReset sets RefreshTokenDurationReset field to given value.

### HasRefreshTokenDurationReset

`func (o *Service) HasRefreshTokenDurationReset() bool`

HasRefreshTokenDurationReset returns a boolean if a field has been set.

### GetRefreshTokenKept

`func (o *Service) GetRefreshTokenKept() bool`

GetRefreshTokenKept returns the RefreshTokenKept field if non-nil, zero value otherwise.

### GetRefreshTokenKeptOk

`func (o *Service) GetRefreshTokenKeptOk() (*bool, bool)`

GetRefreshTokenKeptOk returns a tuple with the RefreshTokenKept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenKept

`func (o *Service) SetRefreshTokenKept(v bool)`

SetRefreshTokenKept sets RefreshTokenKept field to given value.

### HasRefreshTokenKept

`func (o *Service) HasRefreshTokenKept() bool`

HasRefreshTokenKept returns a boolean if a field has been set.

### GetSupportedScopes

`func (o *Service) GetSupportedScopes() []Scope`

GetSupportedScopes returns the SupportedScopes field if non-nil, zero value otherwise.

### GetSupportedScopesOk

`func (o *Service) GetSupportedScopesOk() (*[]Scope, bool)`

GetSupportedScopesOk returns a tuple with the SupportedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedScopes

`func (o *Service) SetSupportedScopes(v []Scope)`

SetSupportedScopes sets SupportedScopes field to given value.

### HasSupportedScopes

`func (o *Service) HasSupportedScopes() bool`

HasSupportedScopes returns a boolean if a field has been set.

### GetScopeRequired

`func (o *Service) GetScopeRequired() bool`

GetScopeRequired returns the ScopeRequired field if non-nil, zero value otherwise.

### GetScopeRequiredOk

`func (o *Service) GetScopeRequiredOk() (*bool, bool)`

GetScopeRequiredOk returns a tuple with the ScopeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeRequired

`func (o *Service) SetScopeRequired(v bool)`

SetScopeRequired sets ScopeRequired field to given value.

### HasScopeRequired

`func (o *Service) HasScopeRequired() bool`

HasScopeRequired returns a boolean if a field has been set.

### GetIdTokenDuration

`func (o *Service) GetIdTokenDuration() int64`

GetIdTokenDuration returns the IdTokenDuration field if non-nil, zero value otherwise.

### GetIdTokenDurationOk

`func (o *Service) GetIdTokenDurationOk() (*int64, bool)`

GetIdTokenDurationOk returns a tuple with the IdTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenDuration

`func (o *Service) SetIdTokenDuration(v int64)`

SetIdTokenDuration sets IdTokenDuration field to given value.

### HasIdTokenDuration

`func (o *Service) HasIdTokenDuration() bool`

HasIdTokenDuration returns a boolean if a field has been set.

### GetAllowableClockSkew

`func (o *Service) GetAllowableClockSkew() int32`

GetAllowableClockSkew returns the AllowableClockSkew field if non-nil, zero value otherwise.

### GetAllowableClockSkewOk

`func (o *Service) GetAllowableClockSkewOk() (*int32, bool)`

GetAllowableClockSkewOk returns a tuple with the AllowableClockSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableClockSkew

`func (o *Service) SetAllowableClockSkew(v int32)`

SetAllowableClockSkew sets AllowableClockSkew field to given value.

### HasAllowableClockSkew

`func (o *Service) HasAllowableClockSkew() bool`

HasAllowableClockSkew returns a boolean if a field has been set.

### GetSupportedClaimTypes

`func (o *Service) GetSupportedClaimTypes() []ClaimType`

GetSupportedClaimTypes returns the SupportedClaimTypes field if non-nil, zero value otherwise.

### GetSupportedClaimTypesOk

`func (o *Service) GetSupportedClaimTypesOk() (*[]ClaimType, bool)`

GetSupportedClaimTypesOk returns a tuple with the SupportedClaimTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedClaimTypes

`func (o *Service) SetSupportedClaimTypes(v []ClaimType)`

SetSupportedClaimTypes sets SupportedClaimTypes field to given value.

### HasSupportedClaimTypes

`func (o *Service) HasSupportedClaimTypes() bool`

HasSupportedClaimTypes returns a boolean if a field has been set.

### GetSupportedClaimLocales

`func (o *Service) GetSupportedClaimLocales() []string`

GetSupportedClaimLocales returns the SupportedClaimLocales field if non-nil, zero value otherwise.

### GetSupportedClaimLocalesOk

`func (o *Service) GetSupportedClaimLocalesOk() (*[]string, bool)`

GetSupportedClaimLocalesOk returns a tuple with the SupportedClaimLocales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedClaimLocales

`func (o *Service) SetSupportedClaimLocales(v []string)`

SetSupportedClaimLocales sets SupportedClaimLocales field to given value.

### HasSupportedClaimLocales

`func (o *Service) HasSupportedClaimLocales() bool`

HasSupportedClaimLocales returns a boolean if a field has been set.

### GetSupportedClaims

`func (o *Service) GetSupportedClaims() []string`

GetSupportedClaims returns the SupportedClaims field if non-nil, zero value otherwise.

### GetSupportedClaimsOk

`func (o *Service) GetSupportedClaimsOk() (*[]string, bool)`

GetSupportedClaimsOk returns a tuple with the SupportedClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedClaims

`func (o *Service) SetSupportedClaims(v []string)`

SetSupportedClaims sets SupportedClaims field to given value.

### HasSupportedClaims

`func (o *Service) HasSupportedClaims() bool`

HasSupportedClaims returns a boolean if a field has been set.

### GetClaimShortcutRestrictive

`func (o *Service) GetClaimShortcutRestrictive() bool`

GetClaimShortcutRestrictive returns the ClaimShortcutRestrictive field if non-nil, zero value otherwise.

### GetClaimShortcutRestrictiveOk

`func (o *Service) GetClaimShortcutRestrictiveOk() (*bool, bool)`

GetClaimShortcutRestrictiveOk returns a tuple with the ClaimShortcutRestrictive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimShortcutRestrictive

`func (o *Service) SetClaimShortcutRestrictive(v bool)`

SetClaimShortcutRestrictive sets ClaimShortcutRestrictive field to given value.

### HasClaimShortcutRestrictive

`func (o *Service) HasClaimShortcutRestrictive() bool`

HasClaimShortcutRestrictive returns a boolean if a field has been set.

### GetJwksUri

`func (o *Service) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *Service) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *Service) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *Service) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetDirectJwksEndpointEnabled

`func (o *Service) GetDirectJwksEndpointEnabled() bool`

GetDirectJwksEndpointEnabled returns the DirectJwksEndpointEnabled field if non-nil, zero value otherwise.

### GetDirectJwksEndpointEnabledOk

`func (o *Service) GetDirectJwksEndpointEnabledOk() (*bool, bool)`

GetDirectJwksEndpointEnabledOk returns a tuple with the DirectJwksEndpointEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectJwksEndpointEnabled

`func (o *Service) SetDirectJwksEndpointEnabled(v bool)`

SetDirectJwksEndpointEnabled sets DirectJwksEndpointEnabled field to given value.

### HasDirectJwksEndpointEnabled

`func (o *Service) HasDirectJwksEndpointEnabled() bool`

HasDirectJwksEndpointEnabled returns a boolean if a field has been set.

### GetJwks

`func (o *Service) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *Service) GetJwksOk() (*string, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *Service) SetJwks(v string)`

SetJwks sets Jwks field to given value.

### HasJwks

`func (o *Service) HasJwks() bool`

HasJwks returns a boolean if a field has been set.

### GetIdTokenSignatureKeyId

`func (o *Service) GetIdTokenSignatureKeyId() string`

GetIdTokenSignatureKeyId returns the IdTokenSignatureKeyId field if non-nil, zero value otherwise.

### GetIdTokenSignatureKeyIdOk

`func (o *Service) GetIdTokenSignatureKeyIdOk() (*string, bool)`

GetIdTokenSignatureKeyIdOk returns a tuple with the IdTokenSignatureKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenSignatureKeyId

`func (o *Service) SetIdTokenSignatureKeyId(v string)`

SetIdTokenSignatureKeyId sets IdTokenSignatureKeyId field to given value.

### HasIdTokenSignatureKeyId

`func (o *Service) HasIdTokenSignatureKeyId() bool`

HasIdTokenSignatureKeyId returns a boolean if a field has been set.

### GetUserInfoSignatureKeyId

`func (o *Service) GetUserInfoSignatureKeyId() string`

GetUserInfoSignatureKeyId returns the UserInfoSignatureKeyId field if non-nil, zero value otherwise.

### GetUserInfoSignatureKeyIdOk

`func (o *Service) GetUserInfoSignatureKeyIdOk() (*string, bool)`

GetUserInfoSignatureKeyIdOk returns a tuple with the UserInfoSignatureKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoSignatureKeyId

`func (o *Service) SetUserInfoSignatureKeyId(v string)`

SetUserInfoSignatureKeyId sets UserInfoSignatureKeyId field to given value.

### HasUserInfoSignatureKeyId

`func (o *Service) HasUserInfoSignatureKeyId() bool`

HasUserInfoSignatureKeyId returns a boolean if a field has been set.

### GetAuthorizationSignatureKeyId

`func (o *Service) GetAuthorizationSignatureKeyId() string`

GetAuthorizationSignatureKeyId returns the AuthorizationSignatureKeyId field if non-nil, zero value otherwise.

### GetAuthorizationSignatureKeyIdOk

`func (o *Service) GetAuthorizationSignatureKeyIdOk() (*string, bool)`

GetAuthorizationSignatureKeyIdOk returns a tuple with the AuthorizationSignatureKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationSignatureKeyId

`func (o *Service) SetAuthorizationSignatureKeyId(v string)`

SetAuthorizationSignatureKeyId sets AuthorizationSignatureKeyId field to given value.

### HasAuthorizationSignatureKeyId

`func (o *Service) HasAuthorizationSignatureKeyId() bool`

HasAuthorizationSignatureKeyId returns a boolean if a field has been set.

### GetUserInfoEndpoint

`func (o *Service) GetUserInfoEndpoint() string`

GetUserInfoEndpoint returns the UserInfoEndpoint field if non-nil, zero value otherwise.

### GetUserInfoEndpointOk

`func (o *Service) GetUserInfoEndpointOk() (*string, bool)`

GetUserInfoEndpointOk returns a tuple with the UserInfoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoEndpoint

`func (o *Service) SetUserInfoEndpoint(v string)`

SetUserInfoEndpoint sets UserInfoEndpoint field to given value.

### HasUserInfoEndpoint

`func (o *Service) HasUserInfoEndpoint() bool`

HasUserInfoEndpoint returns a boolean if a field has been set.

### GetDirectUserInfoEndpointEnabled

`func (o *Service) GetDirectUserInfoEndpointEnabled() bool`

GetDirectUserInfoEndpointEnabled returns the DirectUserInfoEndpointEnabled field if non-nil, zero value otherwise.

### GetDirectUserInfoEndpointEnabledOk

`func (o *Service) GetDirectUserInfoEndpointEnabledOk() (*bool, bool)`

GetDirectUserInfoEndpointEnabledOk returns a tuple with the DirectUserInfoEndpointEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectUserInfoEndpointEnabled

`func (o *Service) SetDirectUserInfoEndpointEnabled(v bool)`

SetDirectUserInfoEndpointEnabled sets DirectUserInfoEndpointEnabled field to given value.

### HasDirectUserInfoEndpointEnabled

`func (o *Service) HasDirectUserInfoEndpointEnabled() bool`

HasDirectUserInfoEndpointEnabled returns a boolean if a field has been set.

### GetDynamicRegistrationSupported

`func (o *Service) GetDynamicRegistrationSupported() bool`

GetDynamicRegistrationSupported returns the DynamicRegistrationSupported field if non-nil, zero value otherwise.

### GetDynamicRegistrationSupportedOk

`func (o *Service) GetDynamicRegistrationSupportedOk() (*bool, bool)`

GetDynamicRegistrationSupportedOk returns a tuple with the DynamicRegistrationSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicRegistrationSupported

`func (o *Service) SetDynamicRegistrationSupported(v bool)`

SetDynamicRegistrationSupported sets DynamicRegistrationSupported field to given value.

### HasDynamicRegistrationSupported

`func (o *Service) HasDynamicRegistrationSupported() bool`

HasDynamicRegistrationSupported returns a boolean if a field has been set.

### GetRegistrationEndpoint

`func (o *Service) GetRegistrationEndpoint() string`

GetRegistrationEndpoint returns the RegistrationEndpoint field if non-nil, zero value otherwise.

### GetRegistrationEndpointOk

`func (o *Service) GetRegistrationEndpointOk() (*string, bool)`

GetRegistrationEndpointOk returns a tuple with the RegistrationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationEndpoint

`func (o *Service) SetRegistrationEndpoint(v string)`

SetRegistrationEndpoint sets RegistrationEndpoint field to given value.

### HasRegistrationEndpoint

`func (o *Service) HasRegistrationEndpoint() bool`

HasRegistrationEndpoint returns a boolean if a field has been set.

### GetRegistrationManagementEndpoint

`func (o *Service) GetRegistrationManagementEndpoint() string`

GetRegistrationManagementEndpoint returns the RegistrationManagementEndpoint field if non-nil, zero value otherwise.

### GetRegistrationManagementEndpointOk

`func (o *Service) GetRegistrationManagementEndpointOk() (*string, bool)`

GetRegistrationManagementEndpointOk returns a tuple with the RegistrationManagementEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationManagementEndpoint

`func (o *Service) SetRegistrationManagementEndpoint(v string)`

SetRegistrationManagementEndpoint sets RegistrationManagementEndpoint field to given value.

### HasRegistrationManagementEndpoint

`func (o *Service) HasRegistrationManagementEndpoint() bool`

HasRegistrationManagementEndpoint returns a boolean if a field has been set.

### GetPolicyUri

`func (o *Service) GetPolicyUri() string`

GetPolicyUri returns the PolicyUri field if non-nil, zero value otherwise.

### GetPolicyUriOk

`func (o *Service) GetPolicyUriOk() (*string, bool)`

GetPolicyUriOk returns a tuple with the PolicyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUri

`func (o *Service) SetPolicyUri(v string)`

SetPolicyUri sets PolicyUri field to given value.

### HasPolicyUri

`func (o *Service) HasPolicyUri() bool`

HasPolicyUri returns a boolean if a field has been set.

### GetTosUri

`func (o *Service) GetTosUri() string`

GetTosUri returns the TosUri field if non-nil, zero value otherwise.

### GetTosUriOk

`func (o *Service) GetTosUriOk() (*string, bool)`

GetTosUriOk returns a tuple with the TosUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosUri

`func (o *Service) SetTosUri(v string)`

SetTosUri sets TosUri field to given value.

### HasTosUri

`func (o *Service) HasTosUri() bool`

HasTosUri returns a boolean if a field has been set.

### GetServiceDocumentation

`func (o *Service) GetServiceDocumentation() string`

GetServiceDocumentation returns the ServiceDocumentation field if non-nil, zero value otherwise.

### GetServiceDocumentationOk

`func (o *Service) GetServiceDocumentationOk() (*string, bool)`

GetServiceDocumentationOk returns a tuple with the ServiceDocumentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDocumentation

`func (o *Service) SetServiceDocumentation(v string)`

SetServiceDocumentation sets ServiceDocumentation field to given value.

### HasServiceDocumentation

`func (o *Service) HasServiceDocumentation() bool`

HasServiceDocumentation returns a boolean if a field has been set.

### GetBackchannelAuthenticationEndpoint

`func (o *Service) GetBackchannelAuthenticationEndpoint() string`

GetBackchannelAuthenticationEndpoint returns the BackchannelAuthenticationEndpoint field if non-nil, zero value otherwise.

### GetBackchannelAuthenticationEndpointOk

`func (o *Service) GetBackchannelAuthenticationEndpointOk() (*string, bool)`

GetBackchannelAuthenticationEndpointOk returns a tuple with the BackchannelAuthenticationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelAuthenticationEndpoint

`func (o *Service) SetBackchannelAuthenticationEndpoint(v string)`

SetBackchannelAuthenticationEndpoint sets BackchannelAuthenticationEndpoint field to given value.

### HasBackchannelAuthenticationEndpoint

`func (o *Service) HasBackchannelAuthenticationEndpoint() bool`

HasBackchannelAuthenticationEndpoint returns a boolean if a field has been set.

### GetSupportedBackchannelTokenDeliveryModes

`func (o *Service) GetSupportedBackchannelTokenDeliveryModes() []DeliveryMode`

GetSupportedBackchannelTokenDeliveryModes returns the SupportedBackchannelTokenDeliveryModes field if non-nil, zero value otherwise.

### GetSupportedBackchannelTokenDeliveryModesOk

`func (o *Service) GetSupportedBackchannelTokenDeliveryModesOk() (*[]DeliveryMode, bool)`

GetSupportedBackchannelTokenDeliveryModesOk returns a tuple with the SupportedBackchannelTokenDeliveryModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedBackchannelTokenDeliveryModes

`func (o *Service) SetSupportedBackchannelTokenDeliveryModes(v []DeliveryMode)`

SetSupportedBackchannelTokenDeliveryModes sets SupportedBackchannelTokenDeliveryModes field to given value.

### HasSupportedBackchannelTokenDeliveryModes

`func (o *Service) HasSupportedBackchannelTokenDeliveryModes() bool`

HasSupportedBackchannelTokenDeliveryModes returns a boolean if a field has been set.

### GetBackchannelAuthReqIdDuration

`func (o *Service) GetBackchannelAuthReqIdDuration() int32`

GetBackchannelAuthReqIdDuration returns the BackchannelAuthReqIdDuration field if non-nil, zero value otherwise.

### GetBackchannelAuthReqIdDurationOk

`func (o *Service) GetBackchannelAuthReqIdDurationOk() (*int32, bool)`

GetBackchannelAuthReqIdDurationOk returns a tuple with the BackchannelAuthReqIdDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelAuthReqIdDuration

`func (o *Service) SetBackchannelAuthReqIdDuration(v int32)`

SetBackchannelAuthReqIdDuration sets BackchannelAuthReqIdDuration field to given value.

### HasBackchannelAuthReqIdDuration

`func (o *Service) HasBackchannelAuthReqIdDuration() bool`

HasBackchannelAuthReqIdDuration returns a boolean if a field has been set.

### GetBackchannelPollingInterval

`func (o *Service) GetBackchannelPollingInterval() int32`

GetBackchannelPollingInterval returns the BackchannelPollingInterval field if non-nil, zero value otherwise.

### GetBackchannelPollingIntervalOk

`func (o *Service) GetBackchannelPollingIntervalOk() (*int32, bool)`

GetBackchannelPollingIntervalOk returns a tuple with the BackchannelPollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelPollingInterval

`func (o *Service) SetBackchannelPollingInterval(v int32)`

SetBackchannelPollingInterval sets BackchannelPollingInterval field to given value.

### HasBackchannelPollingInterval

`func (o *Service) HasBackchannelPollingInterval() bool`

HasBackchannelPollingInterval returns a boolean if a field has been set.

### GetBackchannelUserCodeParameterSupported

`func (o *Service) GetBackchannelUserCodeParameterSupported() bool`

GetBackchannelUserCodeParameterSupported returns the BackchannelUserCodeParameterSupported field if non-nil, zero value otherwise.

### GetBackchannelUserCodeParameterSupportedOk

`func (o *Service) GetBackchannelUserCodeParameterSupportedOk() (*bool, bool)`

GetBackchannelUserCodeParameterSupportedOk returns a tuple with the BackchannelUserCodeParameterSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelUserCodeParameterSupported

`func (o *Service) SetBackchannelUserCodeParameterSupported(v bool)`

SetBackchannelUserCodeParameterSupported sets BackchannelUserCodeParameterSupported field to given value.

### HasBackchannelUserCodeParameterSupported

`func (o *Service) HasBackchannelUserCodeParameterSupported() bool`

HasBackchannelUserCodeParameterSupported returns a boolean if a field has been set.

### GetBackchannelBindingMessageRequiredInFapi

`func (o *Service) GetBackchannelBindingMessageRequiredInFapi() bool`

GetBackchannelBindingMessageRequiredInFapi returns the BackchannelBindingMessageRequiredInFapi field if non-nil, zero value otherwise.

### GetBackchannelBindingMessageRequiredInFapiOk

`func (o *Service) GetBackchannelBindingMessageRequiredInFapiOk() (*bool, bool)`

GetBackchannelBindingMessageRequiredInFapiOk returns a tuple with the BackchannelBindingMessageRequiredInFapi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelBindingMessageRequiredInFapi

`func (o *Service) SetBackchannelBindingMessageRequiredInFapi(v bool)`

SetBackchannelBindingMessageRequiredInFapi sets BackchannelBindingMessageRequiredInFapi field to given value.

### HasBackchannelBindingMessageRequiredInFapi

`func (o *Service) HasBackchannelBindingMessageRequiredInFapi() bool`

HasBackchannelBindingMessageRequiredInFapi returns a boolean if a field has been set.

### GetDeviceAuthorizationEndpoint

`func (o *Service) GetDeviceAuthorizationEndpoint() string`

GetDeviceAuthorizationEndpoint returns the DeviceAuthorizationEndpoint field if non-nil, zero value otherwise.

### GetDeviceAuthorizationEndpointOk

`func (o *Service) GetDeviceAuthorizationEndpointOk() (*string, bool)`

GetDeviceAuthorizationEndpointOk returns a tuple with the DeviceAuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthorizationEndpoint

`func (o *Service) SetDeviceAuthorizationEndpoint(v string)`

SetDeviceAuthorizationEndpoint sets DeviceAuthorizationEndpoint field to given value.

### HasDeviceAuthorizationEndpoint

`func (o *Service) HasDeviceAuthorizationEndpoint() bool`

HasDeviceAuthorizationEndpoint returns a boolean if a field has been set.

### GetDeviceVerificationUri

`func (o *Service) GetDeviceVerificationUri() string`

GetDeviceVerificationUri returns the DeviceVerificationUri field if non-nil, zero value otherwise.

### GetDeviceVerificationUriOk

`func (o *Service) GetDeviceVerificationUriOk() (*string, bool)`

GetDeviceVerificationUriOk returns a tuple with the DeviceVerificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVerificationUri

`func (o *Service) SetDeviceVerificationUri(v string)`

SetDeviceVerificationUri sets DeviceVerificationUri field to given value.

### HasDeviceVerificationUri

`func (o *Service) HasDeviceVerificationUri() bool`

HasDeviceVerificationUri returns a boolean if a field has been set.

### GetDeviceVerificationUriComplete

`func (o *Service) GetDeviceVerificationUriComplete() string`

GetDeviceVerificationUriComplete returns the DeviceVerificationUriComplete field if non-nil, zero value otherwise.

### GetDeviceVerificationUriCompleteOk

`func (o *Service) GetDeviceVerificationUriCompleteOk() (*string, bool)`

GetDeviceVerificationUriCompleteOk returns a tuple with the DeviceVerificationUriComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVerificationUriComplete

`func (o *Service) SetDeviceVerificationUriComplete(v string)`

SetDeviceVerificationUriComplete sets DeviceVerificationUriComplete field to given value.

### HasDeviceVerificationUriComplete

`func (o *Service) HasDeviceVerificationUriComplete() bool`

HasDeviceVerificationUriComplete returns a boolean if a field has been set.

### GetDeviceFlowCodeDuration

`func (o *Service) GetDeviceFlowCodeDuration() int32`

GetDeviceFlowCodeDuration returns the DeviceFlowCodeDuration field if non-nil, zero value otherwise.

### GetDeviceFlowCodeDurationOk

`func (o *Service) GetDeviceFlowCodeDurationOk() (*int32, bool)`

GetDeviceFlowCodeDurationOk returns a tuple with the DeviceFlowCodeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFlowCodeDuration

`func (o *Service) SetDeviceFlowCodeDuration(v int32)`

SetDeviceFlowCodeDuration sets DeviceFlowCodeDuration field to given value.

### HasDeviceFlowCodeDuration

`func (o *Service) HasDeviceFlowCodeDuration() bool`

HasDeviceFlowCodeDuration returns a boolean if a field has been set.

### GetDeviceFlowPollingInterval

`func (o *Service) GetDeviceFlowPollingInterval() int32`

GetDeviceFlowPollingInterval returns the DeviceFlowPollingInterval field if non-nil, zero value otherwise.

### GetDeviceFlowPollingIntervalOk

`func (o *Service) GetDeviceFlowPollingIntervalOk() (*int32, bool)`

GetDeviceFlowPollingIntervalOk returns a tuple with the DeviceFlowPollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFlowPollingInterval

`func (o *Service) SetDeviceFlowPollingInterval(v int32)`

SetDeviceFlowPollingInterval sets DeviceFlowPollingInterval field to given value.

### HasDeviceFlowPollingInterval

`func (o *Service) HasDeviceFlowPollingInterval() bool`

HasDeviceFlowPollingInterval returns a boolean if a field has been set.

### GetUserCodeCharset

`func (o *Service) GetUserCodeCharset() UserCodeCharset`

GetUserCodeCharset returns the UserCodeCharset field if non-nil, zero value otherwise.

### GetUserCodeCharsetOk

`func (o *Service) GetUserCodeCharsetOk() (*UserCodeCharset, bool)`

GetUserCodeCharsetOk returns a tuple with the UserCodeCharset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCodeCharset

`func (o *Service) SetUserCodeCharset(v UserCodeCharset)`

SetUserCodeCharset sets UserCodeCharset field to given value.

### HasUserCodeCharset

`func (o *Service) HasUserCodeCharset() bool`

HasUserCodeCharset returns a boolean if a field has been set.

### GetUserCodeLength

`func (o *Service) GetUserCodeLength() int32`

GetUserCodeLength returns the UserCodeLength field if non-nil, zero value otherwise.

### GetUserCodeLengthOk

`func (o *Service) GetUserCodeLengthOk() (*int32, bool)`

GetUserCodeLengthOk returns a tuple with the UserCodeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCodeLength

`func (o *Service) SetUserCodeLength(v int32)`

SetUserCodeLength sets UserCodeLength field to given value.

### HasUserCodeLength

`func (o *Service) HasUserCodeLength() bool`

HasUserCodeLength returns a boolean if a field has been set.

### GetSupportedTrustFrameworks

`func (o *Service) GetSupportedTrustFrameworks() []string`

GetSupportedTrustFrameworks returns the SupportedTrustFrameworks field if non-nil, zero value otherwise.

### GetSupportedTrustFrameworksOk

`func (o *Service) GetSupportedTrustFrameworksOk() (*[]string, bool)`

GetSupportedTrustFrameworksOk returns a tuple with the SupportedTrustFrameworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedTrustFrameworks

`func (o *Service) SetSupportedTrustFrameworks(v []string)`

SetSupportedTrustFrameworks sets SupportedTrustFrameworks field to given value.

### HasSupportedTrustFrameworks

`func (o *Service) HasSupportedTrustFrameworks() bool`

HasSupportedTrustFrameworks returns a boolean if a field has been set.

### GetSupportedEvidence

`func (o *Service) GetSupportedEvidence() []string`

GetSupportedEvidence returns the SupportedEvidence field if non-nil, zero value otherwise.

### GetSupportedEvidenceOk

`func (o *Service) GetSupportedEvidenceOk() (*[]string, bool)`

GetSupportedEvidenceOk returns a tuple with the SupportedEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedEvidence

`func (o *Service) SetSupportedEvidence(v []string)`

SetSupportedEvidence sets SupportedEvidence field to given value.

### HasSupportedEvidence

`func (o *Service) HasSupportedEvidence() bool`

HasSupportedEvidence returns a boolean if a field has been set.

### GetSupportedIdentityDocuments

`func (o *Service) GetSupportedIdentityDocuments() []string`

GetSupportedIdentityDocuments returns the SupportedIdentityDocuments field if non-nil, zero value otherwise.

### GetSupportedIdentityDocumentsOk

`func (o *Service) GetSupportedIdentityDocumentsOk() (*[]string, bool)`

GetSupportedIdentityDocumentsOk returns a tuple with the SupportedIdentityDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedIdentityDocuments

`func (o *Service) SetSupportedIdentityDocuments(v []string)`

SetSupportedIdentityDocuments sets SupportedIdentityDocuments field to given value.

### HasSupportedIdentityDocuments

`func (o *Service) HasSupportedIdentityDocuments() bool`

HasSupportedIdentityDocuments returns a boolean if a field has been set.

### GetSupportedVerificationMethods

`func (o *Service) GetSupportedVerificationMethods() []string`

GetSupportedVerificationMethods returns the SupportedVerificationMethods field if non-nil, zero value otherwise.

### GetSupportedVerificationMethodsOk

`func (o *Service) GetSupportedVerificationMethodsOk() (*[]string, bool)`

GetSupportedVerificationMethodsOk returns a tuple with the SupportedVerificationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedVerificationMethods

`func (o *Service) SetSupportedVerificationMethods(v []string)`

SetSupportedVerificationMethods sets SupportedVerificationMethods field to given value.

### HasSupportedVerificationMethods

`func (o *Service) HasSupportedVerificationMethods() bool`

HasSupportedVerificationMethods returns a boolean if a field has been set.

### GetSupportedVerifiedClaims

`func (o *Service) GetSupportedVerifiedClaims() []string`

GetSupportedVerifiedClaims returns the SupportedVerifiedClaims field if non-nil, zero value otherwise.

### GetSupportedVerifiedClaimsOk

`func (o *Service) GetSupportedVerifiedClaimsOk() (*[]string, bool)`

GetSupportedVerifiedClaimsOk returns a tuple with the SupportedVerifiedClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedVerifiedClaims

`func (o *Service) SetSupportedVerifiedClaims(v []string)`

SetSupportedVerifiedClaims sets SupportedVerifiedClaims field to given value.

### HasSupportedVerifiedClaims

`func (o *Service) HasSupportedVerifiedClaims() bool`

HasSupportedVerifiedClaims returns a boolean if a field has been set.

### GetVerifiedClaimsValidationSchemaSet

`func (o *Service) GetVerifiedClaimsValidationSchemaSet() VerifiedClaimsValidationSchema`

GetVerifiedClaimsValidationSchemaSet returns the VerifiedClaimsValidationSchemaSet field if non-nil, zero value otherwise.

### GetVerifiedClaimsValidationSchemaSetOk

`func (o *Service) GetVerifiedClaimsValidationSchemaSetOk() (*VerifiedClaimsValidationSchema, bool)`

GetVerifiedClaimsValidationSchemaSetOk returns a tuple with the VerifiedClaimsValidationSchemaSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedClaimsValidationSchemaSet

`func (o *Service) SetVerifiedClaimsValidationSchemaSet(v VerifiedClaimsValidationSchema)`

SetVerifiedClaimsValidationSchemaSet sets VerifiedClaimsValidationSchemaSet field to given value.

### HasVerifiedClaimsValidationSchemaSet

`func (o *Service) HasVerifiedClaimsValidationSchemaSet() bool`

HasVerifiedClaimsValidationSchemaSet returns a boolean if a field has been set.

### SetVerifiedClaimsValidationSchemaSetNil

`func (o *Service) SetVerifiedClaimsValidationSchemaSetNil(b bool)`

 SetVerifiedClaimsValidationSchemaSetNil sets the value for VerifiedClaimsValidationSchemaSet to be an explicit nil

### UnsetVerifiedClaimsValidationSchemaSet
`func (o *Service) UnsetVerifiedClaimsValidationSchemaSet()`

UnsetVerifiedClaimsValidationSchemaSet ensures that no value is present for VerifiedClaimsValidationSchemaSet, not even an explicit nil
### GetAttributes

`func (o *Service) GetAttributes() []Pair`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Service) GetAttributesOk() (*[]Pair, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Service) SetAttributes(v []Pair)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Service) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetNbfOptional

`func (o *Service) GetNbfOptional() bool`

GetNbfOptional returns the NbfOptional field if non-nil, zero value otherwise.

### GetNbfOptionalOk

`func (o *Service) GetNbfOptionalOk() (*bool, bool)`

GetNbfOptionalOk returns a tuple with the NbfOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbfOptional

`func (o *Service) SetNbfOptional(v bool)`

SetNbfOptional sets NbfOptional field to given value.

### HasNbfOptional

`func (o *Service) HasNbfOptional() bool`

HasNbfOptional returns a boolean if a field has been set.

### GetIssSuppressed

`func (o *Service) GetIssSuppressed() bool`

GetIssSuppressed returns the IssSuppressed field if non-nil, zero value otherwise.

### GetIssSuppressedOk

`func (o *Service) GetIssSuppressedOk() (*bool, bool)`

GetIssSuppressedOk returns a tuple with the IssSuppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssSuppressed

`func (o *Service) SetIssSuppressed(v bool)`

SetIssSuppressed sets IssSuppressed field to given value.

### HasIssSuppressed

`func (o *Service) HasIssSuppressed() bool`

HasIssSuppressed returns a boolean if a field has been set.

### GetSupportedCustomClientMetadata

`func (o *Service) GetSupportedCustomClientMetadata() []string`

GetSupportedCustomClientMetadata returns the SupportedCustomClientMetadata field if non-nil, zero value otherwise.

### GetSupportedCustomClientMetadataOk

`func (o *Service) GetSupportedCustomClientMetadataOk() (*[]string, bool)`

GetSupportedCustomClientMetadataOk returns a tuple with the SupportedCustomClientMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCustomClientMetadata

`func (o *Service) SetSupportedCustomClientMetadata(v []string)`

SetSupportedCustomClientMetadata sets SupportedCustomClientMetadata field to given value.

### HasSupportedCustomClientMetadata

`func (o *Service) HasSupportedCustomClientMetadata() bool`

HasSupportedCustomClientMetadata returns a boolean if a field has been set.

### GetTokenExpirationLinked

`func (o *Service) GetTokenExpirationLinked() bool`

GetTokenExpirationLinked returns the TokenExpirationLinked field if non-nil, zero value otherwise.

### GetTokenExpirationLinkedOk

`func (o *Service) GetTokenExpirationLinkedOk() (*bool, bool)`

GetTokenExpirationLinkedOk returns a tuple with the TokenExpirationLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpirationLinked

`func (o *Service) SetTokenExpirationLinked(v bool)`

SetTokenExpirationLinked sets TokenExpirationLinked field to given value.

### HasTokenExpirationLinked

`func (o *Service) HasTokenExpirationLinked() bool`

HasTokenExpirationLinked returns a boolean if a field has been set.

### GetFrontChannelRequestObjectEncryptionRequired

`func (o *Service) GetFrontChannelRequestObjectEncryptionRequired() bool`

GetFrontChannelRequestObjectEncryptionRequired returns the FrontChannelRequestObjectEncryptionRequired field if non-nil, zero value otherwise.

### GetFrontChannelRequestObjectEncryptionRequiredOk

`func (o *Service) GetFrontChannelRequestObjectEncryptionRequiredOk() (*bool, bool)`

GetFrontChannelRequestObjectEncryptionRequiredOk returns a tuple with the FrontChannelRequestObjectEncryptionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontChannelRequestObjectEncryptionRequired

`func (o *Service) SetFrontChannelRequestObjectEncryptionRequired(v bool)`

SetFrontChannelRequestObjectEncryptionRequired sets FrontChannelRequestObjectEncryptionRequired field to given value.

### HasFrontChannelRequestObjectEncryptionRequired

`func (o *Service) HasFrontChannelRequestObjectEncryptionRequired() bool`

HasFrontChannelRequestObjectEncryptionRequired returns a boolean if a field has been set.

### GetRequestObjectEncryptionAlgMatchRequired

`func (o *Service) GetRequestObjectEncryptionAlgMatchRequired() bool`

GetRequestObjectEncryptionAlgMatchRequired returns the RequestObjectEncryptionAlgMatchRequired field if non-nil, zero value otherwise.

### GetRequestObjectEncryptionAlgMatchRequiredOk

`func (o *Service) GetRequestObjectEncryptionAlgMatchRequiredOk() (*bool, bool)`

GetRequestObjectEncryptionAlgMatchRequiredOk returns a tuple with the RequestObjectEncryptionAlgMatchRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectEncryptionAlgMatchRequired

`func (o *Service) SetRequestObjectEncryptionAlgMatchRequired(v bool)`

SetRequestObjectEncryptionAlgMatchRequired sets RequestObjectEncryptionAlgMatchRequired field to given value.

### HasRequestObjectEncryptionAlgMatchRequired

`func (o *Service) HasRequestObjectEncryptionAlgMatchRequired() bool`

HasRequestObjectEncryptionAlgMatchRequired returns a boolean if a field has been set.

### GetRequestObjectEncryptionEncMatchRequired

`func (o *Service) GetRequestObjectEncryptionEncMatchRequired() bool`

GetRequestObjectEncryptionEncMatchRequired returns the RequestObjectEncryptionEncMatchRequired field if non-nil, zero value otherwise.

### GetRequestObjectEncryptionEncMatchRequiredOk

`func (o *Service) GetRequestObjectEncryptionEncMatchRequiredOk() (*bool, bool)`

GetRequestObjectEncryptionEncMatchRequiredOk returns a tuple with the RequestObjectEncryptionEncMatchRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectEncryptionEncMatchRequired

`func (o *Service) SetRequestObjectEncryptionEncMatchRequired(v bool)`

SetRequestObjectEncryptionEncMatchRequired sets RequestObjectEncryptionEncMatchRequired field to given value.

### HasRequestObjectEncryptionEncMatchRequired

`func (o *Service) HasRequestObjectEncryptionEncMatchRequired() bool`

HasRequestObjectEncryptionEncMatchRequired returns a boolean if a field has been set.

### GetHsmEnabled

`func (o *Service) GetHsmEnabled() bool`

GetHsmEnabled returns the HsmEnabled field if non-nil, zero value otherwise.

### GetHsmEnabledOk

`func (o *Service) GetHsmEnabledOk() (*bool, bool)`

GetHsmEnabledOk returns a tuple with the HsmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmEnabled

`func (o *Service) SetHsmEnabled(v bool)`

SetHsmEnabled sets HsmEnabled field to given value.

### HasHsmEnabled

`func (o *Service) HasHsmEnabled() bool`

HasHsmEnabled returns a boolean if a field has been set.

### GetHsks

`func (o *Service) GetHsks() []Hsk`

GetHsks returns the Hsks field if non-nil, zero value otherwise.

### GetHsksOk

`func (o *Service) GetHsksOk() (*[]Hsk, bool)`

GetHsksOk returns a tuple with the Hsks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsks

`func (o *Service) SetHsks(v []Hsk)`

SetHsks sets Hsks field to given value.

### HasHsks

`func (o *Service) HasHsks() bool`

HasHsks returns a boolean if a field has been set.

### GetGrantManagementEndpoint

`func (o *Service) GetGrantManagementEndpoint() string`

GetGrantManagementEndpoint returns the GrantManagementEndpoint field if non-nil, zero value otherwise.

### GetGrantManagementEndpointOk

`func (o *Service) GetGrantManagementEndpointOk() (*string, bool)`

GetGrantManagementEndpointOk returns a tuple with the GrantManagementEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantManagementEndpoint

`func (o *Service) SetGrantManagementEndpoint(v string)`

SetGrantManagementEndpoint sets GrantManagementEndpoint field to given value.

### HasGrantManagementEndpoint

`func (o *Service) HasGrantManagementEndpoint() bool`

HasGrantManagementEndpoint returns a boolean if a field has been set.

### GetGrantManagementActionRequired

`func (o *Service) GetGrantManagementActionRequired() bool`

GetGrantManagementActionRequired returns the GrantManagementActionRequired field if non-nil, zero value otherwise.

### GetGrantManagementActionRequiredOk

`func (o *Service) GetGrantManagementActionRequiredOk() (*bool, bool)`

GetGrantManagementActionRequiredOk returns a tuple with the GrantManagementActionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantManagementActionRequired

`func (o *Service) SetGrantManagementActionRequired(v bool)`

SetGrantManagementActionRequired sets GrantManagementActionRequired field to given value.

### HasGrantManagementActionRequired

`func (o *Service) HasGrantManagementActionRequired() bool`

HasGrantManagementActionRequired returns a boolean if a field has been set.

### GetUnauthorizedOnClientConfigSupported

`func (o *Service) GetUnauthorizedOnClientConfigSupported() bool`

GetUnauthorizedOnClientConfigSupported returns the UnauthorizedOnClientConfigSupported field if non-nil, zero value otherwise.

### GetUnauthorizedOnClientConfigSupportedOk

`func (o *Service) GetUnauthorizedOnClientConfigSupportedOk() (*bool, bool)`

GetUnauthorizedOnClientConfigSupportedOk returns a tuple with the UnauthorizedOnClientConfigSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthorizedOnClientConfigSupported

`func (o *Service) SetUnauthorizedOnClientConfigSupported(v bool)`

SetUnauthorizedOnClientConfigSupported sets UnauthorizedOnClientConfigSupported field to given value.

### HasUnauthorizedOnClientConfigSupported

`func (o *Service) HasUnauthorizedOnClientConfigSupported() bool`

HasUnauthorizedOnClientConfigSupported returns a boolean if a field has been set.

### GetDcrScopeUsedAsRequestable

`func (o *Service) GetDcrScopeUsedAsRequestable() bool`

GetDcrScopeUsedAsRequestable returns the DcrScopeUsedAsRequestable field if non-nil, zero value otherwise.

### GetDcrScopeUsedAsRequestableOk

`func (o *Service) GetDcrScopeUsedAsRequestableOk() (*bool, bool)`

GetDcrScopeUsedAsRequestableOk returns a tuple with the DcrScopeUsedAsRequestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcrScopeUsedAsRequestable

`func (o *Service) SetDcrScopeUsedAsRequestable(v bool)`

SetDcrScopeUsedAsRequestable sets DcrScopeUsedAsRequestable field to given value.

### HasDcrScopeUsedAsRequestable

`func (o *Service) HasDcrScopeUsedAsRequestable() bool`

HasDcrScopeUsedAsRequestable returns a boolean if a field has been set.

### GetEndSessionEndpoint

`func (o *Service) GetEndSessionEndpoint() string`

GetEndSessionEndpoint returns the EndSessionEndpoint field if non-nil, zero value otherwise.

### GetEndSessionEndpointOk

`func (o *Service) GetEndSessionEndpointOk() (*string, bool)`

GetEndSessionEndpointOk returns a tuple with the EndSessionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSessionEndpoint

`func (o *Service) SetEndSessionEndpoint(v string)`

SetEndSessionEndpoint sets EndSessionEndpoint field to given value.

### HasEndSessionEndpoint

`func (o *Service) HasEndSessionEndpoint() bool`

HasEndSessionEndpoint returns a boolean if a field has been set.

### GetLoopbackRedirectionUriVariable

`func (o *Service) GetLoopbackRedirectionUriVariable() bool`

GetLoopbackRedirectionUriVariable returns the LoopbackRedirectionUriVariable field if non-nil, zero value otherwise.

### GetLoopbackRedirectionUriVariableOk

`func (o *Service) GetLoopbackRedirectionUriVariableOk() (*bool, bool)`

GetLoopbackRedirectionUriVariableOk returns a tuple with the LoopbackRedirectionUriVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackRedirectionUriVariable

`func (o *Service) SetLoopbackRedirectionUriVariable(v bool)`

SetLoopbackRedirectionUriVariable sets LoopbackRedirectionUriVariable field to given value.

### HasLoopbackRedirectionUriVariable

`func (o *Service) HasLoopbackRedirectionUriVariable() bool`

HasLoopbackRedirectionUriVariable returns a boolean if a field has been set.

### GetRequestObjectAudienceChecked

`func (o *Service) GetRequestObjectAudienceChecked() bool`

GetRequestObjectAudienceChecked returns the RequestObjectAudienceChecked field if non-nil, zero value otherwise.

### GetRequestObjectAudienceCheckedOk

`func (o *Service) GetRequestObjectAudienceCheckedOk() (*bool, bool)`

GetRequestObjectAudienceCheckedOk returns a tuple with the RequestObjectAudienceChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectAudienceChecked

`func (o *Service) SetRequestObjectAudienceChecked(v bool)`

SetRequestObjectAudienceChecked sets RequestObjectAudienceChecked field to given value.

### HasRequestObjectAudienceChecked

`func (o *Service) HasRequestObjectAudienceChecked() bool`

HasRequestObjectAudienceChecked returns a boolean if a field has been set.

### GetAccessTokenForExternalAttachmentEmbedded

`func (o *Service) GetAccessTokenForExternalAttachmentEmbedded() bool`

GetAccessTokenForExternalAttachmentEmbedded returns the AccessTokenForExternalAttachmentEmbedded field if non-nil, zero value otherwise.

### GetAccessTokenForExternalAttachmentEmbeddedOk

`func (o *Service) GetAccessTokenForExternalAttachmentEmbeddedOk() (*bool, bool)`

GetAccessTokenForExternalAttachmentEmbeddedOk returns a tuple with the AccessTokenForExternalAttachmentEmbedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenForExternalAttachmentEmbedded

`func (o *Service) SetAccessTokenForExternalAttachmentEmbedded(v bool)`

SetAccessTokenForExternalAttachmentEmbedded sets AccessTokenForExternalAttachmentEmbedded field to given value.

### HasAccessTokenForExternalAttachmentEmbedded

`func (o *Service) HasAccessTokenForExternalAttachmentEmbedded() bool`

HasAccessTokenForExternalAttachmentEmbedded returns a boolean if a field has been set.

### GetAuthorityHints

`func (o *Service) GetAuthorityHints() []string`

GetAuthorityHints returns the AuthorityHints field if non-nil, zero value otherwise.

### GetAuthorityHintsOk

`func (o *Service) GetAuthorityHintsOk() (*[]string, bool)`

GetAuthorityHintsOk returns a tuple with the AuthorityHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityHints

`func (o *Service) SetAuthorityHints(v []string)`

SetAuthorityHints sets AuthorityHints field to given value.

### HasAuthorityHints

`func (o *Service) HasAuthorityHints() bool`

HasAuthorityHints returns a boolean if a field has been set.

### GetFederationEnabled

`func (o *Service) GetFederationEnabled() bool`

GetFederationEnabled returns the FederationEnabled field if non-nil, zero value otherwise.

### GetFederationEnabledOk

`func (o *Service) GetFederationEnabledOk() (*bool, bool)`

GetFederationEnabledOk returns a tuple with the FederationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationEnabled

`func (o *Service) SetFederationEnabled(v bool)`

SetFederationEnabled sets FederationEnabled field to given value.

### HasFederationEnabled

`func (o *Service) HasFederationEnabled() bool`

HasFederationEnabled returns a boolean if a field has been set.

### GetFederationJwks

`func (o *Service) GetFederationJwks() string`

GetFederationJwks returns the FederationJwks field if non-nil, zero value otherwise.

### GetFederationJwksOk

`func (o *Service) GetFederationJwksOk() (*string, bool)`

GetFederationJwksOk returns a tuple with the FederationJwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationJwks

`func (o *Service) SetFederationJwks(v string)`

SetFederationJwks sets FederationJwks field to given value.

### HasFederationJwks

`func (o *Service) HasFederationJwks() bool`

HasFederationJwks returns a boolean if a field has been set.

### GetFederationSignatureKeyId

`func (o *Service) GetFederationSignatureKeyId() string`

GetFederationSignatureKeyId returns the FederationSignatureKeyId field if non-nil, zero value otherwise.

### GetFederationSignatureKeyIdOk

`func (o *Service) GetFederationSignatureKeyIdOk() (*string, bool)`

GetFederationSignatureKeyIdOk returns a tuple with the FederationSignatureKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationSignatureKeyId

`func (o *Service) SetFederationSignatureKeyId(v string)`

SetFederationSignatureKeyId sets FederationSignatureKeyId field to given value.

### HasFederationSignatureKeyId

`func (o *Service) HasFederationSignatureKeyId() bool`

HasFederationSignatureKeyId returns a boolean if a field has been set.

### GetFederationConfigurationDuration

`func (o *Service) GetFederationConfigurationDuration() int32`

GetFederationConfigurationDuration returns the FederationConfigurationDuration field if non-nil, zero value otherwise.

### GetFederationConfigurationDurationOk

`func (o *Service) GetFederationConfigurationDurationOk() (*int32, bool)`

GetFederationConfigurationDurationOk returns a tuple with the FederationConfigurationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationConfigurationDuration

`func (o *Service) SetFederationConfigurationDuration(v int32)`

SetFederationConfigurationDuration sets FederationConfigurationDuration field to given value.

### HasFederationConfigurationDuration

`func (o *Service) HasFederationConfigurationDuration() bool`

HasFederationConfigurationDuration returns a boolean if a field has been set.

### GetFederationRegistrationEndpoint

`func (o *Service) GetFederationRegistrationEndpoint() string`

GetFederationRegistrationEndpoint returns the FederationRegistrationEndpoint field if non-nil, zero value otherwise.

### GetFederationRegistrationEndpointOk

`func (o *Service) GetFederationRegistrationEndpointOk() (*string, bool)`

GetFederationRegistrationEndpointOk returns a tuple with the FederationRegistrationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationRegistrationEndpoint

`func (o *Service) SetFederationRegistrationEndpoint(v string)`

SetFederationRegistrationEndpoint sets FederationRegistrationEndpoint field to given value.

### HasFederationRegistrationEndpoint

`func (o *Service) HasFederationRegistrationEndpoint() bool`

HasFederationRegistrationEndpoint returns a boolean if a field has been set.

### GetOrganizationName

`func (o *Service) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *Service) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *Service) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *Service) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetPredefinedTransformedClaims

`func (o *Service) GetPredefinedTransformedClaims() string`

GetPredefinedTransformedClaims returns the PredefinedTransformedClaims field if non-nil, zero value otherwise.

### GetPredefinedTransformedClaimsOk

`func (o *Service) GetPredefinedTransformedClaimsOk() (*string, bool)`

GetPredefinedTransformedClaimsOk returns a tuple with the PredefinedTransformedClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedTransformedClaims

`func (o *Service) SetPredefinedTransformedClaims(v string)`

SetPredefinedTransformedClaims sets PredefinedTransformedClaims field to given value.

### HasPredefinedTransformedClaims

`func (o *Service) HasPredefinedTransformedClaims() bool`

HasPredefinedTransformedClaims returns a boolean if a field has been set.

### GetRefreshTokenIdempotent

`func (o *Service) GetRefreshTokenIdempotent() bool`

GetRefreshTokenIdempotent returns the RefreshTokenIdempotent field if non-nil, zero value otherwise.

### GetRefreshTokenIdempotentOk

`func (o *Service) GetRefreshTokenIdempotentOk() (*bool, bool)`

GetRefreshTokenIdempotentOk returns a tuple with the RefreshTokenIdempotent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenIdempotent

`func (o *Service) SetRefreshTokenIdempotent(v bool)`

SetRefreshTokenIdempotent sets RefreshTokenIdempotent field to given value.

### HasRefreshTokenIdempotent

`func (o *Service) HasRefreshTokenIdempotent() bool`

HasRefreshTokenIdempotent returns a boolean if a field has been set.

### GetSignedJwksUri

`func (o *Service) GetSignedJwksUri() string`

GetSignedJwksUri returns the SignedJwksUri field if non-nil, zero value otherwise.

### GetSignedJwksUriOk

`func (o *Service) GetSignedJwksUriOk() (*string, bool)`

GetSignedJwksUriOk returns a tuple with the SignedJwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedJwksUri

`func (o *Service) SetSignedJwksUri(v string)`

SetSignedJwksUri sets SignedJwksUri field to given value.

### HasSignedJwksUri

`func (o *Service) HasSignedJwksUri() bool`

HasSignedJwksUri returns a boolean if a field has been set.

### GetSupportedAttachments

`func (o *Service) GetSupportedAttachments() []AttachmentType`

GetSupportedAttachments returns the SupportedAttachments field if non-nil, zero value otherwise.

### GetSupportedAttachmentsOk

`func (o *Service) GetSupportedAttachmentsOk() (*[]AttachmentType, bool)`

GetSupportedAttachmentsOk returns a tuple with the SupportedAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAttachments

`func (o *Service) SetSupportedAttachments(v []AttachmentType)`

SetSupportedAttachments sets SupportedAttachments field to given value.

### HasSupportedAttachments

`func (o *Service) HasSupportedAttachments() bool`

HasSupportedAttachments returns a boolean if a field has been set.

### GetSupportedDigestAlgorithms

`func (o *Service) GetSupportedDigestAlgorithms() []string`

GetSupportedDigestAlgorithms returns the SupportedDigestAlgorithms field if non-nil, zero value otherwise.

### GetSupportedDigestAlgorithmsOk

`func (o *Service) GetSupportedDigestAlgorithmsOk() (*[]string, bool)`

GetSupportedDigestAlgorithmsOk returns a tuple with the SupportedDigestAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDigestAlgorithms

`func (o *Service) SetSupportedDigestAlgorithms(v []string)`

SetSupportedDigestAlgorithms sets SupportedDigestAlgorithms field to given value.

### HasSupportedDigestAlgorithms

`func (o *Service) HasSupportedDigestAlgorithms() bool`

HasSupportedDigestAlgorithms returns a boolean if a field has been set.

### GetSupportedDocuments

`func (o *Service) GetSupportedDocuments() []string`

GetSupportedDocuments returns the SupportedDocuments field if non-nil, zero value otherwise.

### GetSupportedDocumentsOk

`func (o *Service) GetSupportedDocumentsOk() (*[]string, bool)`

GetSupportedDocumentsOk returns a tuple with the SupportedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDocuments

`func (o *Service) SetSupportedDocuments(v []string)`

SetSupportedDocuments sets SupportedDocuments field to given value.

### HasSupportedDocuments

`func (o *Service) HasSupportedDocuments() bool`

HasSupportedDocuments returns a boolean if a field has been set.

### GetSupportedDocumentsMethods

`func (o *Service) GetSupportedDocumentsMethods() []string`

GetSupportedDocumentsMethods returns the SupportedDocumentsMethods field if non-nil, zero value otherwise.

### GetSupportedDocumentsMethodsOk

`func (o *Service) GetSupportedDocumentsMethodsOk() (*[]string, bool)`

GetSupportedDocumentsMethodsOk returns a tuple with the SupportedDocumentsMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDocumentsMethods

`func (o *Service) SetSupportedDocumentsMethods(v []string)`

SetSupportedDocumentsMethods sets SupportedDocumentsMethods field to given value.

### HasSupportedDocumentsMethods

`func (o *Service) HasSupportedDocumentsMethods() bool`

HasSupportedDocumentsMethods returns a boolean if a field has been set.

### GetSupportedDocumentsValidationMethods

`func (o *Service) GetSupportedDocumentsValidationMethods() []string`

GetSupportedDocumentsValidationMethods returns the SupportedDocumentsValidationMethods field if non-nil, zero value otherwise.

### GetSupportedDocumentsValidationMethodsOk

`func (o *Service) GetSupportedDocumentsValidationMethodsOk() (*[]string, bool)`

GetSupportedDocumentsValidationMethodsOk returns a tuple with the SupportedDocumentsValidationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDocumentsValidationMethods

`func (o *Service) SetSupportedDocumentsValidationMethods(v []string)`

SetSupportedDocumentsValidationMethods sets SupportedDocumentsValidationMethods field to given value.

### HasSupportedDocumentsValidationMethods

`func (o *Service) HasSupportedDocumentsValidationMethods() bool`

HasSupportedDocumentsValidationMethods returns a boolean if a field has been set.

### GetSupportedDocumentsVerificationMethods

`func (o *Service) GetSupportedDocumentsVerificationMethods() []string`

GetSupportedDocumentsVerificationMethods returns the SupportedDocumentsVerificationMethods field if non-nil, zero value otherwise.

### GetSupportedDocumentsVerificationMethodsOk

`func (o *Service) GetSupportedDocumentsVerificationMethodsOk() (*[]string, bool)`

GetSupportedDocumentsVerificationMethodsOk returns a tuple with the SupportedDocumentsVerificationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDocumentsVerificationMethods

`func (o *Service) SetSupportedDocumentsVerificationMethods(v []string)`

SetSupportedDocumentsVerificationMethods sets SupportedDocumentsVerificationMethods field to given value.

### HasSupportedDocumentsVerificationMethods

`func (o *Service) HasSupportedDocumentsVerificationMethods() bool`

HasSupportedDocumentsVerificationMethods returns a boolean if a field has been set.

### GetSupportedElectronicRecords

`func (o *Service) GetSupportedElectronicRecords() []string`

GetSupportedElectronicRecords returns the SupportedElectronicRecords field if non-nil, zero value otherwise.

### GetSupportedElectronicRecordsOk

`func (o *Service) GetSupportedElectronicRecordsOk() (*[]string, bool)`

GetSupportedElectronicRecordsOk returns a tuple with the SupportedElectronicRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedElectronicRecords

`func (o *Service) SetSupportedElectronicRecords(v []string)`

SetSupportedElectronicRecords sets SupportedElectronicRecords field to given value.

### HasSupportedElectronicRecords

`func (o *Service) HasSupportedElectronicRecords() bool`

HasSupportedElectronicRecords returns a boolean if a field has been set.

### GetSupportedClientRegistrationTypes

`func (o *Service) GetSupportedClientRegistrationTypes() []ClientRegistrationType`

GetSupportedClientRegistrationTypes returns the SupportedClientRegistrationTypes field if non-nil, zero value otherwise.

### GetSupportedClientRegistrationTypesOk

`func (o *Service) GetSupportedClientRegistrationTypesOk() (*[]ClientRegistrationType, bool)`

GetSupportedClientRegistrationTypesOk returns a tuple with the SupportedClientRegistrationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedClientRegistrationTypes

`func (o *Service) SetSupportedClientRegistrationTypes(v []ClientRegistrationType)`

SetSupportedClientRegistrationTypes sets SupportedClientRegistrationTypes field to given value.

### HasSupportedClientRegistrationTypes

`func (o *Service) HasSupportedClientRegistrationTypes() bool`

HasSupportedClientRegistrationTypes returns a boolean if a field has been set.

### GetTokenExchangeByIdentifiableClientsOnly

`func (o *Service) GetTokenExchangeByIdentifiableClientsOnly() bool`

GetTokenExchangeByIdentifiableClientsOnly returns the TokenExchangeByIdentifiableClientsOnly field if non-nil, zero value otherwise.

### GetTokenExchangeByIdentifiableClientsOnlyOk

`func (o *Service) GetTokenExchangeByIdentifiableClientsOnlyOk() (*bool, bool)`

GetTokenExchangeByIdentifiableClientsOnlyOk returns a tuple with the TokenExchangeByIdentifiableClientsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExchangeByIdentifiableClientsOnly

`func (o *Service) SetTokenExchangeByIdentifiableClientsOnly(v bool)`

SetTokenExchangeByIdentifiableClientsOnly sets TokenExchangeByIdentifiableClientsOnly field to given value.

### HasTokenExchangeByIdentifiableClientsOnly

`func (o *Service) HasTokenExchangeByIdentifiableClientsOnly() bool`

HasTokenExchangeByIdentifiableClientsOnly returns a boolean if a field has been set.

### GetTokenExchangeByConfidentialClientsOnly

`func (o *Service) GetTokenExchangeByConfidentialClientsOnly() bool`

GetTokenExchangeByConfidentialClientsOnly returns the TokenExchangeByConfidentialClientsOnly field if non-nil, zero value otherwise.

### GetTokenExchangeByConfidentialClientsOnlyOk

`func (o *Service) GetTokenExchangeByConfidentialClientsOnlyOk() (*bool, bool)`

GetTokenExchangeByConfidentialClientsOnlyOk returns a tuple with the TokenExchangeByConfidentialClientsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExchangeByConfidentialClientsOnly

`func (o *Service) SetTokenExchangeByConfidentialClientsOnly(v bool)`

SetTokenExchangeByConfidentialClientsOnly sets TokenExchangeByConfidentialClientsOnly field to given value.

### HasTokenExchangeByConfidentialClientsOnly

`func (o *Service) HasTokenExchangeByConfidentialClientsOnly() bool`

HasTokenExchangeByConfidentialClientsOnly returns a boolean if a field has been set.

### GetTokenExchangeByPermittedClientsOnly

`func (o *Service) GetTokenExchangeByPermittedClientsOnly() bool`

GetTokenExchangeByPermittedClientsOnly returns the TokenExchangeByPermittedClientsOnly field if non-nil, zero value otherwise.

### GetTokenExchangeByPermittedClientsOnlyOk

`func (o *Service) GetTokenExchangeByPermittedClientsOnlyOk() (*bool, bool)`

GetTokenExchangeByPermittedClientsOnlyOk returns a tuple with the TokenExchangeByPermittedClientsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExchangeByPermittedClientsOnly

`func (o *Service) SetTokenExchangeByPermittedClientsOnly(v bool)`

SetTokenExchangeByPermittedClientsOnly sets TokenExchangeByPermittedClientsOnly field to given value.

### HasTokenExchangeByPermittedClientsOnly

`func (o *Service) HasTokenExchangeByPermittedClientsOnly() bool`

HasTokenExchangeByPermittedClientsOnly returns a boolean if a field has been set.

### GetTokenExchangeEncryptedJwtRejected

`func (o *Service) GetTokenExchangeEncryptedJwtRejected() bool`

GetTokenExchangeEncryptedJwtRejected returns the TokenExchangeEncryptedJwtRejected field if non-nil, zero value otherwise.

### GetTokenExchangeEncryptedJwtRejectedOk

`func (o *Service) GetTokenExchangeEncryptedJwtRejectedOk() (*bool, bool)`

GetTokenExchangeEncryptedJwtRejectedOk returns a tuple with the TokenExchangeEncryptedJwtRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExchangeEncryptedJwtRejected

`func (o *Service) SetTokenExchangeEncryptedJwtRejected(v bool)`

SetTokenExchangeEncryptedJwtRejected sets TokenExchangeEncryptedJwtRejected field to given value.

### HasTokenExchangeEncryptedJwtRejected

`func (o *Service) HasTokenExchangeEncryptedJwtRejected() bool`

HasTokenExchangeEncryptedJwtRejected returns a boolean if a field has been set.

### GetTokenExchangeUnsignedJwtRejected

`func (o *Service) GetTokenExchangeUnsignedJwtRejected() bool`

GetTokenExchangeUnsignedJwtRejected returns the TokenExchangeUnsignedJwtRejected field if non-nil, zero value otherwise.

### GetTokenExchangeUnsignedJwtRejectedOk

`func (o *Service) GetTokenExchangeUnsignedJwtRejectedOk() (*bool, bool)`

GetTokenExchangeUnsignedJwtRejectedOk returns a tuple with the TokenExchangeUnsignedJwtRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExchangeUnsignedJwtRejected

`func (o *Service) SetTokenExchangeUnsignedJwtRejected(v bool)`

SetTokenExchangeUnsignedJwtRejected sets TokenExchangeUnsignedJwtRejected field to given value.

### HasTokenExchangeUnsignedJwtRejected

`func (o *Service) HasTokenExchangeUnsignedJwtRejected() bool`

HasTokenExchangeUnsignedJwtRejected returns a boolean if a field has been set.

### GetJwtGrantByIdentifiableClientsOnly

`func (o *Service) GetJwtGrantByIdentifiableClientsOnly() bool`

GetJwtGrantByIdentifiableClientsOnly returns the JwtGrantByIdentifiableClientsOnly field if non-nil, zero value otherwise.

### GetJwtGrantByIdentifiableClientsOnlyOk

`func (o *Service) GetJwtGrantByIdentifiableClientsOnlyOk() (*bool, bool)`

GetJwtGrantByIdentifiableClientsOnlyOk returns a tuple with the JwtGrantByIdentifiableClientsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtGrantByIdentifiableClientsOnly

`func (o *Service) SetJwtGrantByIdentifiableClientsOnly(v bool)`

SetJwtGrantByIdentifiableClientsOnly sets JwtGrantByIdentifiableClientsOnly field to given value.

### HasJwtGrantByIdentifiableClientsOnly

`func (o *Service) HasJwtGrantByIdentifiableClientsOnly() bool`

HasJwtGrantByIdentifiableClientsOnly returns a boolean if a field has been set.

### GetJwtGrantEncryptedJwtRejected

`func (o *Service) GetJwtGrantEncryptedJwtRejected() bool`

GetJwtGrantEncryptedJwtRejected returns the JwtGrantEncryptedJwtRejected field if non-nil, zero value otherwise.

### GetJwtGrantEncryptedJwtRejectedOk

`func (o *Service) GetJwtGrantEncryptedJwtRejectedOk() (*bool, bool)`

GetJwtGrantEncryptedJwtRejectedOk returns a tuple with the JwtGrantEncryptedJwtRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtGrantEncryptedJwtRejected

`func (o *Service) SetJwtGrantEncryptedJwtRejected(v bool)`

SetJwtGrantEncryptedJwtRejected sets JwtGrantEncryptedJwtRejected field to given value.

### HasJwtGrantEncryptedJwtRejected

`func (o *Service) HasJwtGrantEncryptedJwtRejected() bool`

HasJwtGrantEncryptedJwtRejected returns a boolean if a field has been set.

### GetJwtGrantUnsignedJwtRejected

`func (o *Service) GetJwtGrantUnsignedJwtRejected() bool`

GetJwtGrantUnsignedJwtRejected returns the JwtGrantUnsignedJwtRejected field if non-nil, zero value otherwise.

### GetJwtGrantUnsignedJwtRejectedOk

`func (o *Service) GetJwtGrantUnsignedJwtRejectedOk() (*bool, bool)`

GetJwtGrantUnsignedJwtRejectedOk returns a tuple with the JwtGrantUnsignedJwtRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtGrantUnsignedJwtRejected

`func (o *Service) SetJwtGrantUnsignedJwtRejected(v bool)`

SetJwtGrantUnsignedJwtRejected sets JwtGrantUnsignedJwtRejected field to given value.

### HasJwtGrantUnsignedJwtRejected

`func (o *Service) HasJwtGrantUnsignedJwtRejected() bool`

HasJwtGrantUnsignedJwtRejected returns a boolean if a field has been set.

### GetDcrDuplicateSoftwareIdBlocked

`func (o *Service) GetDcrDuplicateSoftwareIdBlocked() bool`

GetDcrDuplicateSoftwareIdBlocked returns the DcrDuplicateSoftwareIdBlocked field if non-nil, zero value otherwise.

### GetDcrDuplicateSoftwareIdBlockedOk

`func (o *Service) GetDcrDuplicateSoftwareIdBlockedOk() (*bool, bool)`

GetDcrDuplicateSoftwareIdBlockedOk returns a tuple with the DcrDuplicateSoftwareIdBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcrDuplicateSoftwareIdBlocked

`func (o *Service) SetDcrDuplicateSoftwareIdBlocked(v bool)`

SetDcrDuplicateSoftwareIdBlocked sets DcrDuplicateSoftwareIdBlocked field to given value.

### HasDcrDuplicateSoftwareIdBlocked

`func (o *Service) HasDcrDuplicateSoftwareIdBlocked() bool`

HasDcrDuplicateSoftwareIdBlocked returns a boolean if a field has been set.

### GetTrustAnchors

`func (o *Service) GetTrustAnchors() []TrustAnchor`

GetTrustAnchors returns the TrustAnchors field if non-nil, zero value otherwise.

### GetTrustAnchorsOk

`func (o *Service) GetTrustAnchorsOk() (*[]TrustAnchor, bool)`

GetTrustAnchorsOk returns a tuple with the TrustAnchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAnchors

`func (o *Service) SetTrustAnchors(v []TrustAnchor)`

SetTrustAnchors sets TrustAnchors field to given value.

### HasTrustAnchors

`func (o *Service) HasTrustAnchors() bool`

HasTrustAnchors returns a boolean if a field has been set.

### GetOpenidDroppedOnRefreshWithoutOfflineAccess

`func (o *Service) GetOpenidDroppedOnRefreshWithoutOfflineAccess() bool`

GetOpenidDroppedOnRefreshWithoutOfflineAccess returns the OpenidDroppedOnRefreshWithoutOfflineAccess field if non-nil, zero value otherwise.

### GetOpenidDroppedOnRefreshWithoutOfflineAccessOk

`func (o *Service) GetOpenidDroppedOnRefreshWithoutOfflineAccessOk() (*bool, bool)`

GetOpenidDroppedOnRefreshWithoutOfflineAccessOk returns a tuple with the OpenidDroppedOnRefreshWithoutOfflineAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenidDroppedOnRefreshWithoutOfflineAccess

`func (o *Service) SetOpenidDroppedOnRefreshWithoutOfflineAccess(v bool)`

SetOpenidDroppedOnRefreshWithoutOfflineAccess sets OpenidDroppedOnRefreshWithoutOfflineAccess field to given value.

### HasOpenidDroppedOnRefreshWithoutOfflineAccess

`func (o *Service) HasOpenidDroppedOnRefreshWithoutOfflineAccess() bool`

HasOpenidDroppedOnRefreshWithoutOfflineAccess returns a boolean if a field has been set.

### GetSupportedDocumentsCheckMethods

`func (o *Service) GetSupportedDocumentsCheckMethods() []string`

GetSupportedDocumentsCheckMethods returns the SupportedDocumentsCheckMethods field if non-nil, zero value otherwise.

### GetSupportedDocumentsCheckMethodsOk

`func (o *Service) GetSupportedDocumentsCheckMethodsOk() (*[]string, bool)`

GetSupportedDocumentsCheckMethodsOk returns a tuple with the SupportedDocumentsCheckMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDocumentsCheckMethods

`func (o *Service) SetSupportedDocumentsCheckMethods(v []string)`

SetSupportedDocumentsCheckMethods sets SupportedDocumentsCheckMethods field to given value.

### HasSupportedDocumentsCheckMethods

`func (o *Service) HasSupportedDocumentsCheckMethods() bool`

HasSupportedDocumentsCheckMethods returns a boolean if a field has been set.

### GetRsResponseSigned

`func (o *Service) GetRsResponseSigned() bool`

GetRsResponseSigned returns the RsResponseSigned field if non-nil, zero value otherwise.

### GetRsResponseSignedOk

`func (o *Service) GetRsResponseSignedOk() (*bool, bool)`

GetRsResponseSignedOk returns a tuple with the RsResponseSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsResponseSigned

`func (o *Service) SetRsResponseSigned(v bool)`

SetRsResponseSigned sets RsResponseSigned field to given value.

### HasRsResponseSigned

`func (o *Service) HasRsResponseSigned() bool`

HasRsResponseSigned returns a boolean if a field has been set.

### GetCnonceDuration

`func (o *Service) GetCnonceDuration() int64`

GetCnonceDuration returns the CnonceDuration field if non-nil, zero value otherwise.

### GetCnonceDurationOk

`func (o *Service) GetCnonceDurationOk() (*int64, bool)`

GetCnonceDurationOk returns a tuple with the CnonceDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnonceDuration

`func (o *Service) SetCnonceDuration(v int64)`

SetCnonceDuration sets CnonceDuration field to given value.

### HasCnonceDuration

`func (o *Service) HasCnonceDuration() bool`

HasCnonceDuration returns a boolean if a field has been set.

### GetDpopNonceRequired

`func (o *Service) GetDpopNonceRequired() bool`

GetDpopNonceRequired returns the DpopNonceRequired field if non-nil, zero value otherwise.

### GetDpopNonceRequiredOk

`func (o *Service) GetDpopNonceRequiredOk() (*bool, bool)`

GetDpopNonceRequiredOk returns a tuple with the DpopNonceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpopNonceRequired

`func (o *Service) SetDpopNonceRequired(v bool)`

SetDpopNonceRequired sets DpopNonceRequired field to given value.

### HasDpopNonceRequired

`func (o *Service) HasDpopNonceRequired() bool`

HasDpopNonceRequired returns a boolean if a field has been set.

### GetVerifiableCredentialsEnabled

`func (o *Service) GetVerifiableCredentialsEnabled() bool`

GetVerifiableCredentialsEnabled returns the VerifiableCredentialsEnabled field if non-nil, zero value otherwise.

### GetVerifiableCredentialsEnabledOk

`func (o *Service) GetVerifiableCredentialsEnabledOk() (*bool, bool)`

GetVerifiableCredentialsEnabledOk returns a tuple with the VerifiableCredentialsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiableCredentialsEnabled

`func (o *Service) SetVerifiableCredentialsEnabled(v bool)`

SetVerifiableCredentialsEnabled sets VerifiableCredentialsEnabled field to given value.

### HasVerifiableCredentialsEnabled

`func (o *Service) HasVerifiableCredentialsEnabled() bool`

HasVerifiableCredentialsEnabled returns a boolean if a field has been set.

### GetCredentialJwksUri

`func (o *Service) GetCredentialJwksUri() string`

GetCredentialJwksUri returns the CredentialJwksUri field if non-nil, zero value otherwise.

### GetCredentialJwksUriOk

`func (o *Service) GetCredentialJwksUriOk() (*string, bool)`

GetCredentialJwksUriOk returns a tuple with the CredentialJwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialJwksUri

`func (o *Service) SetCredentialJwksUri(v string)`

SetCredentialJwksUri sets CredentialJwksUri field to given value.

### HasCredentialJwksUri

`func (o *Service) HasCredentialJwksUri() bool`

HasCredentialJwksUri returns a boolean if a field has been set.

### GetCredentialOfferDuration

`func (o *Service) GetCredentialOfferDuration() int64`

GetCredentialOfferDuration returns the CredentialOfferDuration field if non-nil, zero value otherwise.

### GetCredentialOfferDurationOk

`func (o *Service) GetCredentialOfferDurationOk() (*int64, bool)`

GetCredentialOfferDurationOk returns a tuple with the CredentialOfferDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialOfferDuration

`func (o *Service) SetCredentialOfferDuration(v int64)`

SetCredentialOfferDuration sets CredentialOfferDuration field to given value.

### HasCredentialOfferDuration

`func (o *Service) HasCredentialOfferDuration() bool`

HasCredentialOfferDuration returns a boolean if a field has been set.

### GetDpopNonceDuration

`func (o *Service) GetDpopNonceDuration() int64`

GetDpopNonceDuration returns the DpopNonceDuration field if non-nil, zero value otherwise.

### GetDpopNonceDurationOk

`func (o *Service) GetDpopNonceDurationOk() (*int64, bool)`

GetDpopNonceDurationOk returns a tuple with the DpopNonceDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpopNonceDuration

`func (o *Service) SetDpopNonceDuration(v int64)`

SetDpopNonceDuration sets DpopNonceDuration field to given value.

### HasDpopNonceDuration

`func (o *Service) HasDpopNonceDuration() bool`

HasDpopNonceDuration returns a boolean if a field has been set.

### GetPreAuthorizedGrantAnonymousAccessSupported

`func (o *Service) GetPreAuthorizedGrantAnonymousAccessSupported() bool`

GetPreAuthorizedGrantAnonymousAccessSupported returns the PreAuthorizedGrantAnonymousAccessSupported field if non-nil, zero value otherwise.

### GetPreAuthorizedGrantAnonymousAccessSupportedOk

`func (o *Service) GetPreAuthorizedGrantAnonymousAccessSupportedOk() (*bool, bool)`

GetPreAuthorizedGrantAnonymousAccessSupportedOk returns a tuple with the PreAuthorizedGrantAnonymousAccessSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthorizedGrantAnonymousAccessSupported

`func (o *Service) SetPreAuthorizedGrantAnonymousAccessSupported(v bool)`

SetPreAuthorizedGrantAnonymousAccessSupported sets PreAuthorizedGrantAnonymousAccessSupported field to given value.

### HasPreAuthorizedGrantAnonymousAccessSupported

`func (o *Service) HasPreAuthorizedGrantAnonymousAccessSupported() bool`

HasPreAuthorizedGrantAnonymousAccessSupported returns a boolean if a field has been set.

### GetCredentialTransactionDuration

`func (o *Service) GetCredentialTransactionDuration() int64`

GetCredentialTransactionDuration returns the CredentialTransactionDuration field if non-nil, zero value otherwise.

### GetCredentialTransactionDurationOk

`func (o *Service) GetCredentialTransactionDurationOk() (*int64, bool)`

GetCredentialTransactionDurationOk returns a tuple with the CredentialTransactionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialTransactionDuration

`func (o *Service) SetCredentialTransactionDuration(v int64)`

SetCredentialTransactionDuration sets CredentialTransactionDuration field to given value.

### HasCredentialTransactionDuration

`func (o *Service) HasCredentialTransactionDuration() bool`

HasCredentialTransactionDuration returns a boolean if a field has been set.

### GetIntrospectionSignatureKeyId

`func (o *Service) GetIntrospectionSignatureKeyId() string`

GetIntrospectionSignatureKeyId returns the IntrospectionSignatureKeyId field if non-nil, zero value otherwise.

### GetIntrospectionSignatureKeyIdOk

`func (o *Service) GetIntrospectionSignatureKeyIdOk() (*string, bool)`

GetIntrospectionSignatureKeyIdOk returns a tuple with the IntrospectionSignatureKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionSignatureKeyId

`func (o *Service) SetIntrospectionSignatureKeyId(v string)`

SetIntrospectionSignatureKeyId sets IntrospectionSignatureKeyId field to given value.

### HasIntrospectionSignatureKeyId

`func (o *Service) HasIntrospectionSignatureKeyId() bool`

HasIntrospectionSignatureKeyId returns a boolean if a field has been set.

### GetResourceSignatureKeyId

`func (o *Service) GetResourceSignatureKeyId() string`

GetResourceSignatureKeyId returns the ResourceSignatureKeyId field if non-nil, zero value otherwise.

### GetResourceSignatureKeyIdOk

`func (o *Service) GetResourceSignatureKeyIdOk() (*string, bool)`

GetResourceSignatureKeyIdOk returns a tuple with the ResourceSignatureKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSignatureKeyId

`func (o *Service) SetResourceSignatureKeyId(v string)`

SetResourceSignatureKeyId sets ResourceSignatureKeyId field to given value.

### HasResourceSignatureKeyId

`func (o *Service) HasResourceSignatureKeyId() bool`

HasResourceSignatureKeyId returns a boolean if a field has been set.

### GetUserPinLength

`func (o *Service) GetUserPinLength() int32`

GetUserPinLength returns the UserPinLength field if non-nil, zero value otherwise.

### GetUserPinLengthOk

`func (o *Service) GetUserPinLengthOk() (*int32, bool)`

GetUserPinLengthOk returns a tuple with the UserPinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPinLength

`func (o *Service) SetUserPinLength(v int32)`

SetUserPinLength sets UserPinLength field to given value.

### HasUserPinLength

`func (o *Service) HasUserPinLength() bool`

HasUserPinLength returns a boolean if a field has been set.

### GetSupportedPromptValues

`func (o *Service) GetSupportedPromptValues() []Prompt`

GetSupportedPromptValues returns the SupportedPromptValues field if non-nil, zero value otherwise.

### GetSupportedPromptValuesOk

`func (o *Service) GetSupportedPromptValuesOk() (*[]Prompt, bool)`

GetSupportedPromptValuesOk returns a tuple with the SupportedPromptValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPromptValues

`func (o *Service) SetSupportedPromptValues(v []Prompt)`

SetSupportedPromptValues sets SupportedPromptValues field to given value.

### HasSupportedPromptValues

`func (o *Service) HasSupportedPromptValues() bool`

HasSupportedPromptValues returns a boolean if a field has been set.

### GetIdTokenReissuable

`func (o *Service) GetIdTokenReissuable() bool`

GetIdTokenReissuable returns the IdTokenReissuable field if non-nil, zero value otherwise.

### GetIdTokenReissuableOk

`func (o *Service) GetIdTokenReissuableOk() (*bool, bool)`

GetIdTokenReissuableOk returns a tuple with the IdTokenReissuable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenReissuable

`func (o *Service) SetIdTokenReissuable(v bool)`

SetIdTokenReissuable sets IdTokenReissuable field to given value.

### HasIdTokenReissuable

`func (o *Service) HasIdTokenReissuable() bool`

HasIdTokenReissuable returns a boolean if a field has been set.

### GetCredentialJwks

`func (o *Service) GetCredentialJwks() string`

GetCredentialJwks returns the CredentialJwks field if non-nil, zero value otherwise.

### GetCredentialJwksOk

`func (o *Service) GetCredentialJwksOk() (*string, bool)`

GetCredentialJwksOk returns a tuple with the CredentialJwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialJwks

`func (o *Service) SetCredentialJwks(v string)`

SetCredentialJwks sets CredentialJwks field to given value.

### HasCredentialJwks

`func (o *Service) HasCredentialJwks() bool`

HasCredentialJwks returns a boolean if a field has been set.

### GetFapiModes

`func (o *Service) GetFapiModes() []FapiMode`

GetFapiModes returns the FapiModes field if non-nil, zero value otherwise.

### GetFapiModesOk

`func (o *Service) GetFapiModesOk() (*[]FapiMode, bool)`

GetFapiModesOk returns a tuple with the FapiModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFapiModes

`func (o *Service) SetFapiModes(v []FapiMode)`

SetFapiModes sets FapiModes field to given value.

### HasFapiModes

`func (o *Service) HasFapiModes() bool`

HasFapiModes returns a boolean if a field has been set.

### GetCredentialDuration

`func (o *Service) GetCredentialDuration() int64`

GetCredentialDuration returns the CredentialDuration field if non-nil, zero value otherwise.

### GetCredentialDurationOk

`func (o *Service) GetCredentialDurationOk() (*int64, bool)`

GetCredentialDurationOk returns a tuple with the CredentialDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialDuration

`func (o *Service) SetCredentialDuration(v int64)`

SetCredentialDuration sets CredentialDuration field to given value.

### HasCredentialDuration

`func (o *Service) HasCredentialDuration() bool`

HasCredentialDuration returns a boolean if a field has been set.

### GetCredentialIssuerMetadata

`func (o *Service) GetCredentialIssuerMetadata() CredentialIssuerMetadata`

GetCredentialIssuerMetadata returns the CredentialIssuerMetadata field if non-nil, zero value otherwise.

### GetCredentialIssuerMetadataOk

`func (o *Service) GetCredentialIssuerMetadataOk() (*CredentialIssuerMetadata, bool)`

GetCredentialIssuerMetadataOk returns a tuple with the CredentialIssuerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialIssuerMetadata

`func (o *Service) SetCredentialIssuerMetadata(v CredentialIssuerMetadata)`

SetCredentialIssuerMetadata sets CredentialIssuerMetadata field to given value.

### HasCredentialIssuerMetadata

`func (o *Service) HasCredentialIssuerMetadata() bool`

HasCredentialIssuerMetadata returns a boolean if a field has been set.

### GetIdTokenAudType

`func (o *Service) GetIdTokenAudType() string`

GetIdTokenAudType returns the IdTokenAudType field if non-nil, zero value otherwise.

### GetIdTokenAudTypeOk

`func (o *Service) GetIdTokenAudTypeOk() (*string, bool)`

GetIdTokenAudTypeOk returns a tuple with the IdTokenAudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenAudType

`func (o *Service) SetIdTokenAudType(v string)`

SetIdTokenAudType sets IdTokenAudType field to given value.

### HasIdTokenAudType

`func (o *Service) HasIdTokenAudType() bool`

HasIdTokenAudType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


