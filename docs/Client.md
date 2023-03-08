# Client

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | The sequential number of the client. The value of this property is assigned by Authlete.  | [optional] [readonly] 
**ServiceNumber** | Pointer to **int32** | The sequential number of the service of the client application. The value of this property is assigned by Authlete.  | [optional] [readonly] 
**Developer** | Pointer to **string** | The developer of the client application.  | [optional] 
**ClientName** | Pointer to **string** | The name of the client application. This property corresponds to &#x60;client_name&#x60; in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**ClientNames** | Pointer to [**[]TaggedValue**](TaggedValue.md) | Client names with language tags. If the client application has different names for different languages, this property can be used to register the names.  | [optional] 
**Description** | Pointer to **string** | The description about the client application. | [optional] 
**Descriptions** | Pointer to [**[]TaggedValue**](TaggedValue.md) | Descriptions about the client application with language tags. If the client application has different descriptions for different languages, this property can be used to register the descriptions.  | [optional] 
**ClientId** | Pointer to **int64** | The client ID. The value of this property is assigned by Authlete. | [optional] [readonly] 
**ClientSecret** | Pointer to **string** | The client secret. A random 512-bit value encoded by base64url (86 letters). The value of this property is assigned by Authlete.  Note that Authlete issues a client secret even to a \&quot;public\&quot; client application, but the client application should not use the client secret unless it changes its client type to \&quot;confidential\&quot;. That is, a public client application should behave as if it had not been issued a client secret. To be specific, a token request from a public client of Authlete should not come along with a client secret although [RFC 6749, 3.2.1. Client Authentication](https://datatracker.ietf.org/doc/html/rfc6749#section-3.2.1) says as follows.  &gt; Confidential clients or other clients issued client credentials MUST authenticate with the authorization server as described in Section 2.3 when making requests to the token endpoint.  | [optional] [readonly] 
**ClientIdAlias** | Pointer to **string** | The alias of the client ID.  Note that the client ID alias is recognized only when this client&#39;s &#x60;clientIdAliasEnabled&#x60; property is set to &#x60;true&#x60; AND the service&#39;s &#x60;clientIdAliasEnabled&#x60; property is also set to &#x60;true&#x60;.  | [optional] 
**ClientIdAliasEnabled** | Pointer to **bool** | The flag to indicate whether the client ID alias is enabled or not.  Note that a service also has &#x60;clientIdAliasEnabled&#x60; property. If the service&#39;s &#x60;clientIdAliasEnabled&#x60; property is set to &#x60;false&#x60;, the client ID alias of this client is not recognized even if this client&#39;s &#x60;clientIdAliasEnabled&#x60; property is set to &#x60;true&#x60;.  | [optional] 
**ClientType** | Pointer to [**ClientType**](ClientType.md) |  | [optional] 
**ApplicationType** | Pointer to [**NullableApplicationType**](ApplicationType.md) |  | [optional] 
**LogoUri** | Pointer to **string** | The URL pointing to the logo image of the client application.  This property corresponds to &#x60;logo_uri&#x60; in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**LogoUris** | Pointer to [**[]TaggedValue**](TaggedValue.md) | Logo image URLs with language tags. If the client application has different logo images for different languages, this property can be used to register URLs of the images.  | [optional] 
**Contacts** | Pointer to **[]string** | An array of email addresses of people responsible for the client application.  This property corresponds to contacts in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**TlsClientCertificateBoundAccessTokens** | Pointer to **bool** | The flag to indicate whether this client use TLS client certificate bound access tokens.  | [optional] 
**DynamicallyRegistered** | Pointer to **bool** | The flag to indicate whether this client has been registered dynamically. For more details, see [RFC 7591](https://datatracker.ietf.org/doc/html/rfc7591).  | [optional] [readonly] 
**SoftwareId** | Pointer to **string** | The unique identifier string assigned by the client developer or software publisher used by registration endpoints to identify the client software to be dynamically registered.  This property corresponds to the &#x60;software_id metadata&#x60; defined in [2. Client Metadata](https://datatracker.ietf.org/doc/html/rfc7591#section-2) of [RFC 7591](https://datatracker.ietf.org/doc/html/rfc7591).  | [optional] 
**SoftwareVersion** | Pointer to **string** | The version identifier string for the client software identified by the software ID.  This property corresponds to the software_version metadata defined in [2. Client Metadata](https://datatracker.ietf.org/doc/html/rfc7591#section-2) of [RFC 7591](https://datatracker.ietf.org/doc/html/rfc7591).  | [optional] 
**RegistrationAccessTokenHash** | Pointer to **string** | The hash of the registration access token for this client.  | [optional] 
**CreatedAt** | Pointer to **int64** | The time at which this client was created. The value is represented as milliseconds since the UNIX epoch (1970-01-01). | [optional] [readonly] 
**ModifiedAt** | Pointer to **int64** | The time at which this client was last modified. The value is represented as milliseconds since the UNIX epoch (1970-01-01). | [optional] [readonly] 
**GrantTypes** | Pointer to [**[]GrantType**](GrantType.md) | A string array of grant types which the client application declares that it will restrict itself to using. This property corresponds to &#x60;grant_types&#x60; in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**ResponseTypes** | Pointer to [**[]ResponseType**](ResponseType.md) | A string array of response types which the client application declares that it will restrict itself to using. This property corresponds to &#x60;response_types&#x60; in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**RedirectUris** | Pointer to **[]string** | Redirect URIs that the client application uses to receive a response from the authorization endpoint. Requirements for a redirect URI are as follows.  **Requirements by RFC 6749** (From [RFC 6749, 3.1.2. Redirection Endpoint](https://datatracker.ietf.org/doc/html/rfc6749#section-3.1.2))  - Must be an absolute URI. - Must not have a fragment component.  **Requirements by OpenID Connect** (From \&quot;[OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata), application_type\&quot;)  - The scheme of the redirect URI used for Implicit Grant by a client application whose application is &#x60;web&#x60; must be &#x60;https&#x60;. This is checked at runtime by Authlete. - The hostname of the redirect URI used for Implicit Grant by a client application whose application type is &#x60;web&#x60; must not be &#x60;localhost&#x60;. This is checked at runtime by Authlete. - The scheme of the redirect URI used by a client application whose application type is &#x60;native&#x60; must be either (1) a custom scheme or (2) &#x60;http&#x60;, which is allowed only when the hostname part is &#x60;localhost&#x60;. This is checked at runtime by Authlete.  **Requirements by Authlete**  - Must consist of printable ASCII letters only. - Must not exceed 200 letters.  Note that Authlete allows the application type to be &#x60;null&#x60;. In other words, a client application does not have to choose &#x60;web&#x60; or &#x60;native&#x60; as its application type. If the application type is &#x60;null&#x60;, the requirements by OpenID Connect are not checked at runtime.  An authorization request from a client application which has not registered any redirect URI fails unless at least all the following conditions are satisfied.  - The client type of the client application is &#x60;confidential&#x60;. - The value of &#x60;response_type&#x60; request parameter is &#x60;code&#x60;. - The authorization request has the &#x60;redirect_uri&#x60; request parameter. - The value of &#x60;scope&#x60; request parameter does not contain &#x60;openid&#x60;.  RFC 6749 allows partial match of redirect URI under some conditions (see [RFC 6749, 3.1.2.2. Registration Requirements](https://datatracker.ietf.org/doc/html/rfc6749#section-3.1.2.2) for details), but OpenID Connect requires exact match.  | [optional] 
**AuthorizationSignAlg** | Pointer to [**JwsAlg**](JwsAlg.md) |  | [optional] 
**AuthorizationEncryptionAlg** | Pointer to [**JweAlg**](JweAlg.md) |  | [optional] 
**AuthorizationEncryptionEnc** | Pointer to [**JweEnc**](JweEnc.md) |  | [optional] 
**TokenAuthMethod** | Pointer to [**ClientAuthenticationMethod**](ClientAuthenticationMethod.md) |  | [optional] 
**TokenAuthSignAlg** | Pointer to [**JwsAlg**](JwsAlg.md) |  | [optional] 
**SelfSignedCertificateKeyId** | Pointer to **string** | The key ID of a JWK containing a self-signed certificate of this client.  | [optional] 
**TlsClientAuthSubjectDn** | Pointer to **string** | The string representation of the expected subject distinguished name of the certificate this client will use in mutual TLS authentication.  See &#x60;tls_client_auth_subject_dn&#x60; in \&quot;Mutual TLS Profiles for OAuth Clients, 2.3. Dynamic Client Registration\&quot; for details.  | [optional] 
**TlsClientAuthSanDns** | Pointer to **string** | The string representation of the expected DNS subject alternative name of the certificate this client will use in mutual TLS authentication.  See &#x60;tls_client_auth_san_dns&#x60; in \&quot;Mutual TLS Profiles for OAuth Clients, 2.3. Dynamic Client Registration\&quot; for details.  | [optional] 
**TlsClientAuthSanUri** | Pointer to **string** | The string representation of the expected URI subject alternative name of the certificate this client will use in mutual TLS authentication.  See &#x60;tls_client_auth_san_uri&#x60; in \&quot;Mutual TLS Profiles for OAuth Clients, 2.3. Dynamic Client Registration\&quot; for details.  | [optional] 
**TlsClientAuthSanIp** | Pointer to **string** | The string representation of the expected IP address subject alternative name of the certificate this client will use in mutual TLS authentication.  See &#x60;tls_client_auth_san_ip&#x60; in \&quot;Mutual TLS Profiles for OAuth Clients, 2.3. Dynamic Client Registration\&quot; for details.  | [optional] 
**TlsClientAuthSanEmail** | Pointer to **string** | The string representation of the expected email address subject alternative name of the certificate this client will use in mutual TLS authentication.  See &#x60;tls_client_auth_san_email&#x60; in \&quot;Mutual TLS Profiles for OAuth Clients, 2.3. Dynamic Client Registration\&quot; for details.  | [optional] 
**ParRequired** | Pointer to **bool** | The flag to indicate whether this client is required to use the pushed authorization request endpoint. This property corresponds to the &#x60;require_pushed_authorization_requests&#x60; client metadata defined in \&quot;OAuth 2.0 Pushed Authorization Requests\&quot;.  | [optional] 
**RequestObjectRequired** | Pointer to **bool** | The flag to indicate whether authorization requests from this client are always required to utilize a request object by using either &#x60;request&#x60; or &#x60;request_uri&#x60; request parameter.  If this flag is set to &#x60;true&#x60; and the service&#39;s &#x60;traditionalRequestObjectProcessingApplied&#x60; is set to &#x60;false&#x60;, authorization requests from this client are processed as if &#x60;require_signed_request_object&#x60; client metadata of this client is &#x60;true&#x60;. The metadata is defined in \&quot;JAR (JWT Secured Authorization Request)\&quot;.  | [optional] 
**RequestSignAlg** | Pointer to [**JwsAlg**](JwsAlg.md) |  | [optional] 
**RequestEncryptionAlg** | Pointer to [**JweAlg**](JweAlg.md) |  | [optional] 
**RequestEncryptionEnc** | Pointer to [**JweEnc**](JweEnc.md) |  | [optional] 
**RequestUris** | Pointer to **[]string** | An array of URLs each of which points to a request object.  Authlete requires that URLs used as values for &#x60;request_uri&#x60; request parameter be pre-registered. This property is used for the pre-registration. See [OpenID Connect Core 1.0, 6.2. Passing a Request Object by Reference](https://openid.net/specs/openid-connect-core-1_0.html#RequestUriParameter) for details.  | [optional] 
**DefaultMaxAge** | Pointer to **int32** | The default maximum authentication age in seconds. This value is used when an authorization request from the client application does not have &#x60;max_age&#x60; request parameter.  This property corresponds to &#x60;default_max_age&#x60; in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**DefaultAcrs** | Pointer to **[]string** | The default ACRs (Authentication Context Class References). This value is used when an authorization request from the client application has neither &#x60;acr_values&#x60; request parameter nor &#x60;acr&#x60; claim in claims request parameter.  | [optional] 
**IdTokenSignAlg** | Pointer to [**JwsAlg**](JwsAlg.md) |  | [optional] 
**IdTokenEncryptionAlg** | Pointer to [**JweAlg**](JweAlg.md) |  | [optional] 
**IdTokenEncryptionEnc** | Pointer to [**JweEnc**](JweEnc.md) |  | [optional] 
**AuthTimeRequired** | Pointer to **bool** | The flag to indicate whether this client requires &#x60;auth_time&#x60; claim to be embedded in the ID token.  This property corresponds to &#x60;require_auth_time&#x60; in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**SubjectType** | Pointer to [**SubjectType**](SubjectType.md) |  | [optional] 
**SectorIdentifierUri** | Pointer to **string** | The value of the sector identifier URI. This represents the &#x60;sector_identifier_uri&#x60; client metadata which is defined in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata)  | [optional] 
**DerivedSectorIdentifier** | Pointer to **string** | The sector identifier host component as derived from either the &#x60;sector_identifier_uri&#x60; or the registered redirect URI. If no &#x60;sector_identifier_uri&#x60; is registered and multiple redirect URIs are also registered, the value of this property is &#x60;null&#x60;.  | [optional] [readonly] 
**JwksUri** | Pointer to **string** | The URL pointing to the JWK Set of the client application. The content pointed to by the URL is JSON which complies with the format described in [JSON Web Key (JWK), 5. JWK Set Format](https://datatracker.ietf.org/doc/html/rfc7517#section-5). The JWK Set must not include private keys of the client application.  If the client application requests encryption for ID tokens (from the authorization/token/userinfo endpoints) and/or signs request objects, it must make available its JWK Set containing public keys for the encryption and/or the signature at the URL of &#x60;jwksUri&#x60;. The service (Authlete) fetches the JWK Set from the URL as necessary.  [OpenID Connect Dynamic Client Registration 1.0](https://openid.net/specs/openid-connect-registration-1_0.html) says that &#x60;jwks&#x60; must not be used when the client can use &#x60;jwks_uri&#x60;, but Authlete allows both properties to be registered at the same time. However, Authlete does not use the content of &#x60;jwks&#x60; when &#x60;jwksUri&#x60; is registered.  This property corresponds to &#x60;jwks_uri&#x60; in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**Jwks** | Pointer to **string** | The content of the JWK Set of the client application. The format is described in [JSON Web Key (JWK), 5. JWK Set Format](https://datatracker.ietf.org/doc/html/rfc7517#section-5). The JWK Set must not include private keys of the client application.  [OpenID Connect Dynamic Client Registration 1.0](https://openid.net/specs/openid-connect-registration-1_0.html) says that &#x60;jwks&#x60; must not be used when the client can use &#x60;jwks_uri&#x60;, but Authlete allows both properties to be registered at the same time. However, Authlete does not use the content of &#x60;jwks&#x60; when &#x60;jwksUri&#x60; is registered.  This property corresponds to &#x60;jwks_uri&#x60; in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**UserInfoSignAlg** | Pointer to [**JwsAlg**](JwsAlg.md) |  | [optional] 
**UserInfoEncryptionAlg** | Pointer to [**JweAlg**](JweAlg.md) |  | [optional] 
**UserInfoEncryptionEnc** | Pointer to [**JweEnc**](JweEnc.md) |  | [optional] 
**LoginUri** | Pointer to **string** | The URL which a third party can use to initiate a login by the client application.  This property corresponds to &#x60;initiate_login_uri&#x60; in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**TosUri** | Pointer to **string** | The URL pointing to the \&quot;Terms Of Service\&quot; page.  This property corresponds to &#x60;tos_uri&#x60; in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**TosUris** | Pointer to [**[]TaggedValue**](TaggedValue.md) | URLs of \&quot;Terms Of Service\&quot; pages with language tags.  If the client application has different \&quot;Terms Of Service\&quot; pages for different languages, this property can be used to register the URLs.  | [optional] 
**PolicyUri** | Pointer to **string** | The URL pointing to the page which describes the policy as to how end-user&#39;s profile data is used.  This property corresponds to &#x60;policy_uri&#x60; in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**PolicyUris** | Pointer to [**[]TaggedValue**](TaggedValue.md) | URLs of policy pages with language tags. If the client application has different policy pages for different languages, this property can be used to register the URLs.  | [optional] 
**ClientUri** | Pointer to **string** | The URL pointing to the home page of the client application.  This property corresponds to &#x60;client_uri&#x60; in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**ClientUris** | Pointer to [**[]TaggedValue**](TaggedValue.md) | Home page URLs with language tags. If the client application has different home pages for different languages, this property can be used to register the URLs.  | [optional] 
**BcDeliveryMode** | Pointer to **string** | The backchannel token delivery mode.  This property corresponds to the &#x60;backchannel_token_delivery_mode&#x60; metadata. The backchannel token delivery mode is defined in the specification of \&quot;CIBA (Client Initiated Backchannel Authentication)\&quot;.  | [optional] 
**BcNotificationEndpoint** | Pointer to **string** | The backchannel client notification endpoint.  This property corresponds to the &#x60;backchannel_client_notification_endpoint&#x60; metadata. The backchannel token delivery mode is defined in the specification of \&quot;CIBA (Client Initiated Backchannel Authentication)\&quot;.  | [optional] 
**BcRequestSignAlg** | Pointer to [**JwsAlg**](JwsAlg.md) |  | [optional] 
**BcUserCodeRequired** | Pointer to **bool** | The boolean flag to indicate whether a user code is required when this client makes a backchannel authentication request.  This property corresponds to the &#x60;backchannel_user_code_parameter&#x60; metadata.  | [optional] 
**Attributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of this client.  | [optional] 
**Extension** | Pointer to [**ClientExtension**](ClientExtension.md) |  | [optional] 
**AuthorizationDetailsTypes** | Pointer to **[]string** | The authorization details types that this client may use as values of the &#x60;type&#x60; field in &#x60;authorization_details&#x60;.  This property corresponds to the &#x60;authorization_details_types&#x60; metadata. See [OAuth 2.0 Rich Authorization Requests (RAR)](https://datatracker.ietf.org/doc/draft-ietf-oauth-rar/) for details.  Note that the property name was renamed from authorizationDataTypes to authorizationDetailsTypes to align with the change made by the 5th draft of the RAR specification.  | [optional] 
**CustomMetadata** | Pointer to **string** | The custom client metadata in JSON format.  Standard specifications define client metadata as necessary. The following are such examples.  * [OpenID Connect Dynamic Client Registration 1.0](https://openid.net/specs/openid-connect-registration-1_0.html) * [RFC 7591 OAuth 2.0 Dynamic Client Registration Protocol](https://www.rfc-editor.org/rfc/rfc7591.html) * [RFC 8705 OAuth 2.0 Mutual-TLS Client Authentication and Certificate-Bound Access Tokens](https://www.rfc-editor.org/rfc/rfc8705.html) * [OpenID Connect Client-Initiated Backchannel Authentication Flow - Core 1.0](https://openid.net/specs/openid-client-initiated-backchannel-authentication-core-1_0.html) * [The OAuth 2.0 Authorization Framework: JWT Secured Authorization Request (JAR)](https://datatracker.ietf.org/doc/draft-ietf-oauth-jwsreq/) * [Financial-grade API: JWT Secured Authorization Response Mode for OAuth 2.0 (JARM)](https://openid.net/specs/openid-financial-api-jarm.html) * [OAuth 2.0 Pushed Authorization Requests (PAR)](https://datatracker.ietf.org/doc/rfc9126/) * [OAuth 2.0 Rich Authorization Requests (RAR)](https://datatracker.ietf.org/doc/draft-ietf-oauth-rar/)  Standard client metadata included in Client Registration Request and Client Update Request (cf. [OIDC DynReg](https://openid.net/specs/openid-connect-registration-1_0.html), [RFC 7591](https://www.rfc-editor.org/rfc/rfc7591.html) and [RFC 7592](https://www.rfc-editor.org/rfc/rfc7592.html)) are, if supported by Authlete, set to corresponding properties of the client application. For example, the value of the &#x60;client_name&#x60; client metadata in Client Registration/Update Request is set to the clientName property. On the other hand, unrecognized client metadata are discarded.  By listing up custom client metadata in advance by using the &#x60;supportedCustomClientMetadata&#x60; property of Service, Authlete can recognize them and stores their values into the database. The stored custom client metadata values can be referenced by this property.  | [optional] 
**FrontChannelRequestObjectEncryptionRequired** | Pointer to **bool** | The flag indicating whether encryption of request object is required when the request object is passed through the front channel.  This flag does not affect the processing of request objects at the Pushed Authorization Request Endpoint, which is defined in [OAuth 2.0 Pushed Authorization Requests](https://datatracker.ietf.org/doc/rfc9126/). Unecrypted request objects are accepted at the endpoint even if this flag is &#x60;true&#x60;.  This flag does not indicate whether a request object is always required. There is a different flag, &#x60;requestObjectRequired&#x60;, for the purpose.  Even if this flag is &#x60;false&#x60;, encryption of request object is required if the &#x60;frontChannelRequestObjectEncryptionRequired&#x60; flag of the service is &#x60;true&#x60;.  | [optional] 
**RequestObjectEncryptionAlgMatchRequired** | Pointer to **bool** | The flag indicating whether the JWE alg of encrypted request object must match the &#x60;request_object_encryption_alg&#x60; client metadata.  The &#x60;request_object_encryption_alg&#x60; client metadata itself is defined in [OpenID Connect Dynamic Client Registration 1.0](https://openid.net/specs/openid-connect-registration-1_0.html) as follows.  &gt; request_object_encryption_alg &gt; &gt; OPTIONAL. JWE [JWE] alg algorithm [JWA] the RP is declaring that it may use for encrypting Request   Objects sent to the OP. This parameter SHOULD be included when symmetric encryption will be used,   since this signals to the OP that a client_secret value needs to be returned from which the   symmetric key will be derived, that might not otherwise be returned. The RP MAY still use other   supported encryption algorithms or send unencrypted Request Objects, even when this parameter   is present. If both signing and encryption are requested, the Request Object will be signed   then encrypted, with the result being a Nested JWT, as defined in [JWT]. The default, if omitted,   is that the RP is not declaring whether it might encrypt any Request Objects.  The point here is \&quot;The RP MAY still use other supported encryption algorithms or send unencrypted Request Objects, even when this parameter is present.\&quot;  The property that represents the client metadata is &#x60;requestEncryptionAlg&#x60;. See the description of &#x60;requestEncryptionAlg&#x60; for details.  Even if this flag is &#x60;false&#x60;, the match is required if the &#x60;requestObjectEncryptionAlgMatchRequired&#x60; flag of the service is &#x60;true&#x60;.  | [optional] 
**RequestObjectEncryptionEncMatchRequired** | Pointer to **bool** | The flag indicating whether the JWE enc of encrypted request object must match the &#x60;request_object_encryption_enc&#x60; client metadata.  The &#x60;request_object_encryption_enc&#x60; client metadata itself is defined in [OpenID Connect Dynamic Client Registration 1.0](https://openid.net/specs/openid-connect-registration-1_0.html) as follows.  &gt; request_object_encryption_enc &gt; &gt; OPTIONAL. JWE enc algorithm [JWA] the RP is declaring that it may use for encrypting Request   Objects sent to the OP. If request_object_encryption_alg is specified, the default for this   value is A128CBC-HS256. When request_object_encryption_enc is included, request_object_encryption_alg   MUST also be provided.  The property that represents the client metadata is &#x60;requestEncryptionEnc&#x60;. See the description of &#x60;requestEncryptionEnc&#x60;  for details.  Even if this flag is &#x60;false&#x60;, the match is required if the &#x60;requestObjectEncryptionEncMatchRequired&#x60; flag of the service is &#x60;true&#x60;.  | [optional] 
**DigestAlgorithm** | Pointer to **string** | The digest algorithm that this client requests the server to use when it computes digest values of &lt;a href&#x3D; \&quot;https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html#name-external-attachments\&quot; &gt;external attachments&lt;/a&gt;, which may be referenced from within ID tokens or userinfo responses (or any place that can have the &#x60;verified_claims&#x60; claim).  Possible values are listed in the &lt;a href&#x3D; \&quot;https://www.iana.org/assignments/named-information/named-information.xhtml#hash-alg\&quot; &gt;Hash Algorithm Registry&lt;/a&gt; of IANA (Internet Assigned Numbers Authority), but the server does not necessarily support all the values there. When this property is omitted, &#x60;sha-256&#x60; is used as the default algorithm.  This property corresponds to the &#x60;digest_algorithm&#x60; client metadata which was defined by the third implementer&#39;s draft of [OpenID Connect for Identity Assurance 1.0](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html).  | [optional] 
**SingleAccessTokenPerSubject** | Pointer to **bool** | If &#x60;Enabled&#x60; is selected, an attempt to issue a new access token invalidates existing access tokens that are associated with the same combination of subject and client.  Note that, however, attempts by Client Credentials Flow do not invalidate existing access tokens because access tokens issued by Client Credentials Flow are not associated with any end-user&#39;s subject.  Even if &#x60;Disabled&#x60; is selected here, single access token per subject is effective if &#x60;singleAccessTokenPerSubject&#x60; of the &#x60;Service&#x60; this client belongs to is Enabled.  | [optional] 
**PkceRequired** | Pointer to **bool** | The flag to indicate whether the use of Proof Key for Code Exchange (PKCE) is always required for authorization requests by Authorization Code Flow.  If &#x60;true&#x60;, &#x60;code_challenge&#x60; request parameter is always required for authorization requests using Authorization Code Flow.  See [RFC 7636](https://tools.ietf.org/html/rfc7636) (Proof Key for Code Exchange by OAuth Public Clients) for details about &#x60;code_challenge&#x60; request parameter.  | [optional] 
**PkceS256Required** | Pointer to **bool** | The flag to indicate whether &#x60;S256&#x60; is always required as the code challenge method whenever [PKCE (RFC 7636)](https://tools.ietf.org/html/rfc7636) is used.  If this flag is set to &#x60;true&#x60;, &#x60;code_challenge_method&#x3D;S256&#x60; must be included in the authorization request whenever it includes the &#x60;code_challenge&#x60; request parameter. Neither omission of the &#x60;code_challenge_method&#x60; request parameter nor use of plain (&#x60;code_challenge_method&#x3D;plain&#x60;) is allowed.  | [optional] 
**DpopRequired** | Pointer to **bool** | If the DPoP is required for this client  | [optional] 
**AutomaticallyRegistered** | Pointer to **bool** | The flag indicating whether this client was registered by the \&quot;automatic\&quot; client registration of OIDC Federation.  | [optional] 
**ExplicitlyRegistered** | Pointer to **bool** | The flag indicating whether this client was registered by the \&quot;explicit\&quot; client registration of OIDC Federation.  | [optional] 
**RsResponseSigned** | Pointer to **bool** | The flag indicating whether this service signs responses from the resource server.  | [optional] 
**RsSignedRequestKeyId** | Pointer to **string** | Get the key ID of a JWK containing the public key used by this client to sign requests to the resource server.  | [optional] 
**ClientRegistrationTypes** | Pointer to [**[]ClientRegistrationType**](ClientRegistrationType.md) | Get the client registration types that the client has declared it may use.  | [optional] 
**OrganizationName** | Pointer to **string** | Get the human-readable name representing the organization that manages this client. This property corresponds  to the organization_name client metadata that is defined in OpenID Connect Federation 1.0.  | [optional] 
**SignedJwksUri** | Pointer to **string** | Get the URI of the endpoint that returns this client&#39;s JWK Set document in the JWT format. This property  corresponds to the &#x60;signed_jwks_uri&#x60; client metadata defined in OpenID Connect Federation 1.0.  | [optional] 
**EntityId** | Pointer to **string** | the entity ID of this client.  | [optional] 
**TrustAnchorId** | Pointer to **string** | The entity ID of the trust anchor of the trust chain that was used when this client was registered or updated by  the mechanism defined in OpenID Connect Federation 1.0  | [optional] 
**TrustChain** | Pointer to **[]string** | The trust chain that was used when this client was registered or updated by the mechanism defined in OpenID Connect Federation 1.0  | [optional] 
**TrustChainExpiresAt** | Pointer to **int64** | the expiration time of the trust chain that was used when this client was registered or updated by the mechanism  defined in OpenID Connect Federation 1.0. The value is represented as milliseconds elapsed since the Unix epoch (1970-01-01).  | [optional] 
**TrustChainUpdatedAt** | Pointer to **int64** | the time at which the trust chain was updated by the mechanism defined in OpenID Connect Federation 1.0  | [optional] 

## Methods

### NewClient

`func NewClient() *Client`

NewClient instantiates a new Client object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientWithDefaults

`func NewClientWithDefaults() *Client`

NewClientWithDefaults instantiates a new Client object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *Client) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Client) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Client) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Client) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetServiceNumber

`func (o *Client) GetServiceNumber() int32`

GetServiceNumber returns the ServiceNumber field if non-nil, zero value otherwise.

### GetServiceNumberOk

`func (o *Client) GetServiceNumberOk() (*int32, bool)`

GetServiceNumberOk returns a tuple with the ServiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNumber

`func (o *Client) SetServiceNumber(v int32)`

SetServiceNumber sets ServiceNumber field to given value.

### HasServiceNumber

`func (o *Client) HasServiceNumber() bool`

HasServiceNumber returns a boolean if a field has been set.

### GetDeveloper

`func (o *Client) GetDeveloper() string`

GetDeveloper returns the Developer field if non-nil, zero value otherwise.

### GetDeveloperOk

`func (o *Client) GetDeveloperOk() (*string, bool)`

GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloper

`func (o *Client) SetDeveloper(v string)`

SetDeveloper sets Developer field to given value.

### HasDeveloper

`func (o *Client) HasDeveloper() bool`

HasDeveloper returns a boolean if a field has been set.

### GetClientName

`func (o *Client) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *Client) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *Client) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *Client) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientNames

`func (o *Client) GetClientNames() []TaggedValue`

GetClientNames returns the ClientNames field if non-nil, zero value otherwise.

### GetClientNamesOk

`func (o *Client) GetClientNamesOk() (*[]TaggedValue, bool)`

GetClientNamesOk returns a tuple with the ClientNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNames

`func (o *Client) SetClientNames(v []TaggedValue)`

SetClientNames sets ClientNames field to given value.

### HasClientNames

`func (o *Client) HasClientNames() bool`

HasClientNames returns a boolean if a field has been set.

### GetDescription

`func (o *Client) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Client) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Client) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Client) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptions

`func (o *Client) GetDescriptions() []TaggedValue`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *Client) GetDescriptionsOk() (*[]TaggedValue, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *Client) SetDescriptions(v []TaggedValue)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *Client) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetClientId

`func (o *Client) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Client) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Client) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Client) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *Client) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *Client) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *Client) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *Client) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientIdAlias

`func (o *Client) GetClientIdAlias() string`

GetClientIdAlias returns the ClientIdAlias field if non-nil, zero value otherwise.

### GetClientIdAliasOk

`func (o *Client) GetClientIdAliasOk() (*string, bool)`

GetClientIdAliasOk returns a tuple with the ClientIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAlias

`func (o *Client) SetClientIdAlias(v string)`

SetClientIdAlias sets ClientIdAlias field to given value.

### HasClientIdAlias

`func (o *Client) HasClientIdAlias() bool`

HasClientIdAlias returns a boolean if a field has been set.

### GetClientIdAliasEnabled

`func (o *Client) GetClientIdAliasEnabled() bool`

GetClientIdAliasEnabled returns the ClientIdAliasEnabled field if non-nil, zero value otherwise.

### GetClientIdAliasEnabledOk

`func (o *Client) GetClientIdAliasEnabledOk() (*bool, bool)`

GetClientIdAliasEnabledOk returns a tuple with the ClientIdAliasEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasEnabled

`func (o *Client) SetClientIdAliasEnabled(v bool)`

SetClientIdAliasEnabled sets ClientIdAliasEnabled field to given value.

### HasClientIdAliasEnabled

`func (o *Client) HasClientIdAliasEnabled() bool`

HasClientIdAliasEnabled returns a boolean if a field has been set.

### GetClientType

`func (o *Client) GetClientType() ClientType`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *Client) GetClientTypeOk() (*ClientType, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *Client) SetClientType(v ClientType)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *Client) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetApplicationType

`func (o *Client) GetApplicationType() ApplicationType`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *Client) GetApplicationTypeOk() (*ApplicationType, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *Client) SetApplicationType(v ApplicationType)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *Client) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### SetApplicationTypeNil

`func (o *Client) SetApplicationTypeNil(b bool)`

 SetApplicationTypeNil sets the value for ApplicationType to be an explicit nil

### UnsetApplicationType
`func (o *Client) UnsetApplicationType()`

UnsetApplicationType ensures that no value is present for ApplicationType, not even an explicit nil
### GetLogoUri

`func (o *Client) GetLogoUri() string`

GetLogoUri returns the LogoUri field if non-nil, zero value otherwise.

### GetLogoUriOk

`func (o *Client) GetLogoUriOk() (*string, bool)`

GetLogoUriOk returns a tuple with the LogoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUri

`func (o *Client) SetLogoUri(v string)`

SetLogoUri sets LogoUri field to given value.

### HasLogoUri

`func (o *Client) HasLogoUri() bool`

HasLogoUri returns a boolean if a field has been set.

### GetLogoUris

`func (o *Client) GetLogoUris() []TaggedValue`

GetLogoUris returns the LogoUris field if non-nil, zero value otherwise.

### GetLogoUrisOk

`func (o *Client) GetLogoUrisOk() (*[]TaggedValue, bool)`

GetLogoUrisOk returns a tuple with the LogoUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUris

`func (o *Client) SetLogoUris(v []TaggedValue)`

SetLogoUris sets LogoUris field to given value.

### HasLogoUris

`func (o *Client) HasLogoUris() bool`

HasLogoUris returns a boolean if a field has been set.

### GetContacts

`func (o *Client) GetContacts() []string`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *Client) GetContactsOk() (*[]string, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *Client) SetContacts(v []string)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *Client) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetTlsClientCertificateBoundAccessTokens

`func (o *Client) GetTlsClientCertificateBoundAccessTokens() bool`

GetTlsClientCertificateBoundAccessTokens returns the TlsClientCertificateBoundAccessTokens field if non-nil, zero value otherwise.

### GetTlsClientCertificateBoundAccessTokensOk

`func (o *Client) GetTlsClientCertificateBoundAccessTokensOk() (*bool, bool)`

GetTlsClientCertificateBoundAccessTokensOk returns a tuple with the TlsClientCertificateBoundAccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientCertificateBoundAccessTokens

`func (o *Client) SetTlsClientCertificateBoundAccessTokens(v bool)`

SetTlsClientCertificateBoundAccessTokens sets TlsClientCertificateBoundAccessTokens field to given value.

### HasTlsClientCertificateBoundAccessTokens

`func (o *Client) HasTlsClientCertificateBoundAccessTokens() bool`

HasTlsClientCertificateBoundAccessTokens returns a boolean if a field has been set.

### GetDynamicallyRegistered

`func (o *Client) GetDynamicallyRegistered() bool`

GetDynamicallyRegistered returns the DynamicallyRegistered field if non-nil, zero value otherwise.

### GetDynamicallyRegisteredOk

`func (o *Client) GetDynamicallyRegisteredOk() (*bool, bool)`

GetDynamicallyRegisteredOk returns a tuple with the DynamicallyRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicallyRegistered

`func (o *Client) SetDynamicallyRegistered(v bool)`

SetDynamicallyRegistered sets DynamicallyRegistered field to given value.

### HasDynamicallyRegistered

`func (o *Client) HasDynamicallyRegistered() bool`

HasDynamicallyRegistered returns a boolean if a field has been set.

### GetSoftwareId

`func (o *Client) GetSoftwareId() string`

GetSoftwareId returns the SoftwareId field if non-nil, zero value otherwise.

### GetSoftwareIdOk

`func (o *Client) GetSoftwareIdOk() (*string, bool)`

GetSoftwareIdOk returns a tuple with the SoftwareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareId

`func (o *Client) SetSoftwareId(v string)`

SetSoftwareId sets SoftwareId field to given value.

### HasSoftwareId

`func (o *Client) HasSoftwareId() bool`

HasSoftwareId returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *Client) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *Client) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *Client) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *Client) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetRegistrationAccessTokenHash

`func (o *Client) GetRegistrationAccessTokenHash() string`

GetRegistrationAccessTokenHash returns the RegistrationAccessTokenHash field if non-nil, zero value otherwise.

### GetRegistrationAccessTokenHashOk

`func (o *Client) GetRegistrationAccessTokenHashOk() (*string, bool)`

GetRegistrationAccessTokenHashOk returns a tuple with the RegistrationAccessTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationAccessTokenHash

`func (o *Client) SetRegistrationAccessTokenHash(v string)`

SetRegistrationAccessTokenHash sets RegistrationAccessTokenHash field to given value.

### HasRegistrationAccessTokenHash

`func (o *Client) HasRegistrationAccessTokenHash() bool`

HasRegistrationAccessTokenHash returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Client) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Client) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Client) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Client) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Client) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Client) GetModifiedAtOk() (*int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Client) SetModifiedAt(v int64)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Client) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetGrantTypes

`func (o *Client) GetGrantTypes() []GrantType`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *Client) GetGrantTypesOk() (*[]GrantType, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *Client) SetGrantTypes(v []GrantType)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *Client) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetResponseTypes

`func (o *Client) GetResponseTypes() []ResponseType`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *Client) GetResponseTypesOk() (*[]ResponseType, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *Client) SetResponseTypes(v []ResponseType)`

SetResponseTypes sets ResponseTypes field to given value.

### HasResponseTypes

`func (o *Client) HasResponseTypes() bool`

HasResponseTypes returns a boolean if a field has been set.

### GetRedirectUris

`func (o *Client) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *Client) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *Client) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *Client) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetAuthorizationSignAlg

`func (o *Client) GetAuthorizationSignAlg() JwsAlg`

GetAuthorizationSignAlg returns the AuthorizationSignAlg field if non-nil, zero value otherwise.

### GetAuthorizationSignAlgOk

`func (o *Client) GetAuthorizationSignAlgOk() (*JwsAlg, bool)`

GetAuthorizationSignAlgOk returns a tuple with the AuthorizationSignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationSignAlg

`func (o *Client) SetAuthorizationSignAlg(v JwsAlg)`

SetAuthorizationSignAlg sets AuthorizationSignAlg field to given value.

### HasAuthorizationSignAlg

`func (o *Client) HasAuthorizationSignAlg() bool`

HasAuthorizationSignAlg returns a boolean if a field has been set.

### GetAuthorizationEncryptionAlg

`func (o *Client) GetAuthorizationEncryptionAlg() JweAlg`

GetAuthorizationEncryptionAlg returns the AuthorizationEncryptionAlg field if non-nil, zero value otherwise.

### GetAuthorizationEncryptionAlgOk

`func (o *Client) GetAuthorizationEncryptionAlgOk() (*JweAlg, bool)`

GetAuthorizationEncryptionAlgOk returns a tuple with the AuthorizationEncryptionAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEncryptionAlg

`func (o *Client) SetAuthorizationEncryptionAlg(v JweAlg)`

SetAuthorizationEncryptionAlg sets AuthorizationEncryptionAlg field to given value.

### HasAuthorizationEncryptionAlg

`func (o *Client) HasAuthorizationEncryptionAlg() bool`

HasAuthorizationEncryptionAlg returns a boolean if a field has been set.

### GetAuthorizationEncryptionEnc

`func (o *Client) GetAuthorizationEncryptionEnc() JweEnc`

GetAuthorizationEncryptionEnc returns the AuthorizationEncryptionEnc field if non-nil, zero value otherwise.

### GetAuthorizationEncryptionEncOk

`func (o *Client) GetAuthorizationEncryptionEncOk() (*JweEnc, bool)`

GetAuthorizationEncryptionEncOk returns a tuple with the AuthorizationEncryptionEnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEncryptionEnc

`func (o *Client) SetAuthorizationEncryptionEnc(v JweEnc)`

SetAuthorizationEncryptionEnc sets AuthorizationEncryptionEnc field to given value.

### HasAuthorizationEncryptionEnc

`func (o *Client) HasAuthorizationEncryptionEnc() bool`

HasAuthorizationEncryptionEnc returns a boolean if a field has been set.

### GetTokenAuthMethod

`func (o *Client) GetTokenAuthMethod() ClientAuthenticationMethod`

GetTokenAuthMethod returns the TokenAuthMethod field if non-nil, zero value otherwise.

### GetTokenAuthMethodOk

`func (o *Client) GetTokenAuthMethodOk() (*ClientAuthenticationMethod, bool)`

GetTokenAuthMethodOk returns a tuple with the TokenAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAuthMethod

`func (o *Client) SetTokenAuthMethod(v ClientAuthenticationMethod)`

SetTokenAuthMethod sets TokenAuthMethod field to given value.

### HasTokenAuthMethod

`func (o *Client) HasTokenAuthMethod() bool`

HasTokenAuthMethod returns a boolean if a field has been set.

### GetTokenAuthSignAlg

`func (o *Client) GetTokenAuthSignAlg() JwsAlg`

GetTokenAuthSignAlg returns the TokenAuthSignAlg field if non-nil, zero value otherwise.

### GetTokenAuthSignAlgOk

`func (o *Client) GetTokenAuthSignAlgOk() (*JwsAlg, bool)`

GetTokenAuthSignAlgOk returns a tuple with the TokenAuthSignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAuthSignAlg

`func (o *Client) SetTokenAuthSignAlg(v JwsAlg)`

SetTokenAuthSignAlg sets TokenAuthSignAlg field to given value.

### HasTokenAuthSignAlg

`func (o *Client) HasTokenAuthSignAlg() bool`

HasTokenAuthSignAlg returns a boolean if a field has been set.

### GetSelfSignedCertificateKeyId

`func (o *Client) GetSelfSignedCertificateKeyId() string`

GetSelfSignedCertificateKeyId returns the SelfSignedCertificateKeyId field if non-nil, zero value otherwise.

### GetSelfSignedCertificateKeyIdOk

`func (o *Client) GetSelfSignedCertificateKeyIdOk() (*string, bool)`

GetSelfSignedCertificateKeyIdOk returns a tuple with the SelfSignedCertificateKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSignedCertificateKeyId

`func (o *Client) SetSelfSignedCertificateKeyId(v string)`

SetSelfSignedCertificateKeyId sets SelfSignedCertificateKeyId field to given value.

### HasSelfSignedCertificateKeyId

`func (o *Client) HasSelfSignedCertificateKeyId() bool`

HasSelfSignedCertificateKeyId returns a boolean if a field has been set.

### GetTlsClientAuthSubjectDn

`func (o *Client) GetTlsClientAuthSubjectDn() string`

GetTlsClientAuthSubjectDn returns the TlsClientAuthSubjectDn field if non-nil, zero value otherwise.

### GetTlsClientAuthSubjectDnOk

`func (o *Client) GetTlsClientAuthSubjectDnOk() (*string, bool)`

GetTlsClientAuthSubjectDnOk returns a tuple with the TlsClientAuthSubjectDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientAuthSubjectDn

`func (o *Client) SetTlsClientAuthSubjectDn(v string)`

SetTlsClientAuthSubjectDn sets TlsClientAuthSubjectDn field to given value.

### HasTlsClientAuthSubjectDn

`func (o *Client) HasTlsClientAuthSubjectDn() bool`

HasTlsClientAuthSubjectDn returns a boolean if a field has been set.

### GetTlsClientAuthSanDns

`func (o *Client) GetTlsClientAuthSanDns() string`

GetTlsClientAuthSanDns returns the TlsClientAuthSanDns field if non-nil, zero value otherwise.

### GetTlsClientAuthSanDnsOk

`func (o *Client) GetTlsClientAuthSanDnsOk() (*string, bool)`

GetTlsClientAuthSanDnsOk returns a tuple with the TlsClientAuthSanDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientAuthSanDns

`func (o *Client) SetTlsClientAuthSanDns(v string)`

SetTlsClientAuthSanDns sets TlsClientAuthSanDns field to given value.

### HasTlsClientAuthSanDns

`func (o *Client) HasTlsClientAuthSanDns() bool`

HasTlsClientAuthSanDns returns a boolean if a field has been set.

### GetTlsClientAuthSanUri

`func (o *Client) GetTlsClientAuthSanUri() string`

GetTlsClientAuthSanUri returns the TlsClientAuthSanUri field if non-nil, zero value otherwise.

### GetTlsClientAuthSanUriOk

`func (o *Client) GetTlsClientAuthSanUriOk() (*string, bool)`

GetTlsClientAuthSanUriOk returns a tuple with the TlsClientAuthSanUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientAuthSanUri

`func (o *Client) SetTlsClientAuthSanUri(v string)`

SetTlsClientAuthSanUri sets TlsClientAuthSanUri field to given value.

### HasTlsClientAuthSanUri

`func (o *Client) HasTlsClientAuthSanUri() bool`

HasTlsClientAuthSanUri returns a boolean if a field has been set.

### GetTlsClientAuthSanIp

`func (o *Client) GetTlsClientAuthSanIp() string`

GetTlsClientAuthSanIp returns the TlsClientAuthSanIp field if non-nil, zero value otherwise.

### GetTlsClientAuthSanIpOk

`func (o *Client) GetTlsClientAuthSanIpOk() (*string, bool)`

GetTlsClientAuthSanIpOk returns a tuple with the TlsClientAuthSanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientAuthSanIp

`func (o *Client) SetTlsClientAuthSanIp(v string)`

SetTlsClientAuthSanIp sets TlsClientAuthSanIp field to given value.

### HasTlsClientAuthSanIp

`func (o *Client) HasTlsClientAuthSanIp() bool`

HasTlsClientAuthSanIp returns a boolean if a field has been set.

### GetTlsClientAuthSanEmail

`func (o *Client) GetTlsClientAuthSanEmail() string`

GetTlsClientAuthSanEmail returns the TlsClientAuthSanEmail field if non-nil, zero value otherwise.

### GetTlsClientAuthSanEmailOk

`func (o *Client) GetTlsClientAuthSanEmailOk() (*string, bool)`

GetTlsClientAuthSanEmailOk returns a tuple with the TlsClientAuthSanEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsClientAuthSanEmail

`func (o *Client) SetTlsClientAuthSanEmail(v string)`

SetTlsClientAuthSanEmail sets TlsClientAuthSanEmail field to given value.

### HasTlsClientAuthSanEmail

`func (o *Client) HasTlsClientAuthSanEmail() bool`

HasTlsClientAuthSanEmail returns a boolean if a field has been set.

### GetParRequired

`func (o *Client) GetParRequired() bool`

GetParRequired returns the ParRequired field if non-nil, zero value otherwise.

### GetParRequiredOk

`func (o *Client) GetParRequiredOk() (*bool, bool)`

GetParRequiredOk returns a tuple with the ParRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParRequired

`func (o *Client) SetParRequired(v bool)`

SetParRequired sets ParRequired field to given value.

### HasParRequired

`func (o *Client) HasParRequired() bool`

HasParRequired returns a boolean if a field has been set.

### GetRequestObjectRequired

`func (o *Client) GetRequestObjectRequired() bool`

GetRequestObjectRequired returns the RequestObjectRequired field if non-nil, zero value otherwise.

### GetRequestObjectRequiredOk

`func (o *Client) GetRequestObjectRequiredOk() (*bool, bool)`

GetRequestObjectRequiredOk returns a tuple with the RequestObjectRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectRequired

`func (o *Client) SetRequestObjectRequired(v bool)`

SetRequestObjectRequired sets RequestObjectRequired field to given value.

### HasRequestObjectRequired

`func (o *Client) HasRequestObjectRequired() bool`

HasRequestObjectRequired returns a boolean if a field has been set.

### GetRequestSignAlg

`func (o *Client) GetRequestSignAlg() JwsAlg`

GetRequestSignAlg returns the RequestSignAlg field if non-nil, zero value otherwise.

### GetRequestSignAlgOk

`func (o *Client) GetRequestSignAlgOk() (*JwsAlg, bool)`

GetRequestSignAlgOk returns a tuple with the RequestSignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSignAlg

`func (o *Client) SetRequestSignAlg(v JwsAlg)`

SetRequestSignAlg sets RequestSignAlg field to given value.

### HasRequestSignAlg

`func (o *Client) HasRequestSignAlg() bool`

HasRequestSignAlg returns a boolean if a field has been set.

### GetRequestEncryptionAlg

`func (o *Client) GetRequestEncryptionAlg() JweAlg`

GetRequestEncryptionAlg returns the RequestEncryptionAlg field if non-nil, zero value otherwise.

### GetRequestEncryptionAlgOk

`func (o *Client) GetRequestEncryptionAlgOk() (*JweAlg, bool)`

GetRequestEncryptionAlgOk returns a tuple with the RequestEncryptionAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestEncryptionAlg

`func (o *Client) SetRequestEncryptionAlg(v JweAlg)`

SetRequestEncryptionAlg sets RequestEncryptionAlg field to given value.

### HasRequestEncryptionAlg

`func (o *Client) HasRequestEncryptionAlg() bool`

HasRequestEncryptionAlg returns a boolean if a field has been set.

### GetRequestEncryptionEnc

`func (o *Client) GetRequestEncryptionEnc() JweEnc`

GetRequestEncryptionEnc returns the RequestEncryptionEnc field if non-nil, zero value otherwise.

### GetRequestEncryptionEncOk

`func (o *Client) GetRequestEncryptionEncOk() (*JweEnc, bool)`

GetRequestEncryptionEncOk returns a tuple with the RequestEncryptionEnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestEncryptionEnc

`func (o *Client) SetRequestEncryptionEnc(v JweEnc)`

SetRequestEncryptionEnc sets RequestEncryptionEnc field to given value.

### HasRequestEncryptionEnc

`func (o *Client) HasRequestEncryptionEnc() bool`

HasRequestEncryptionEnc returns a boolean if a field has been set.

### GetRequestUris

`func (o *Client) GetRequestUris() []string`

GetRequestUris returns the RequestUris field if non-nil, zero value otherwise.

### GetRequestUrisOk

`func (o *Client) GetRequestUrisOk() (*[]string, bool)`

GetRequestUrisOk returns a tuple with the RequestUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUris

`func (o *Client) SetRequestUris(v []string)`

SetRequestUris sets RequestUris field to given value.

### HasRequestUris

`func (o *Client) HasRequestUris() bool`

HasRequestUris returns a boolean if a field has been set.

### GetDefaultMaxAge

`func (o *Client) GetDefaultMaxAge() int32`

GetDefaultMaxAge returns the DefaultMaxAge field if non-nil, zero value otherwise.

### GetDefaultMaxAgeOk

`func (o *Client) GetDefaultMaxAgeOk() (*int32, bool)`

GetDefaultMaxAgeOk returns a tuple with the DefaultMaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMaxAge

`func (o *Client) SetDefaultMaxAge(v int32)`

SetDefaultMaxAge sets DefaultMaxAge field to given value.

### HasDefaultMaxAge

`func (o *Client) HasDefaultMaxAge() bool`

HasDefaultMaxAge returns a boolean if a field has been set.

### GetDefaultAcrs

`func (o *Client) GetDefaultAcrs() []string`

GetDefaultAcrs returns the DefaultAcrs field if non-nil, zero value otherwise.

### GetDefaultAcrsOk

`func (o *Client) GetDefaultAcrsOk() (*[]string, bool)`

GetDefaultAcrsOk returns a tuple with the DefaultAcrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAcrs

`func (o *Client) SetDefaultAcrs(v []string)`

SetDefaultAcrs sets DefaultAcrs field to given value.

### HasDefaultAcrs

`func (o *Client) HasDefaultAcrs() bool`

HasDefaultAcrs returns a boolean if a field has been set.

### GetIdTokenSignAlg

`func (o *Client) GetIdTokenSignAlg() JwsAlg`

GetIdTokenSignAlg returns the IdTokenSignAlg field if non-nil, zero value otherwise.

### GetIdTokenSignAlgOk

`func (o *Client) GetIdTokenSignAlgOk() (*JwsAlg, bool)`

GetIdTokenSignAlgOk returns a tuple with the IdTokenSignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenSignAlg

`func (o *Client) SetIdTokenSignAlg(v JwsAlg)`

SetIdTokenSignAlg sets IdTokenSignAlg field to given value.

### HasIdTokenSignAlg

`func (o *Client) HasIdTokenSignAlg() bool`

HasIdTokenSignAlg returns a boolean if a field has been set.

### GetIdTokenEncryptionAlg

`func (o *Client) GetIdTokenEncryptionAlg() JweAlg`

GetIdTokenEncryptionAlg returns the IdTokenEncryptionAlg field if non-nil, zero value otherwise.

### GetIdTokenEncryptionAlgOk

`func (o *Client) GetIdTokenEncryptionAlgOk() (*JweAlg, bool)`

GetIdTokenEncryptionAlgOk returns a tuple with the IdTokenEncryptionAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenEncryptionAlg

`func (o *Client) SetIdTokenEncryptionAlg(v JweAlg)`

SetIdTokenEncryptionAlg sets IdTokenEncryptionAlg field to given value.

### HasIdTokenEncryptionAlg

`func (o *Client) HasIdTokenEncryptionAlg() bool`

HasIdTokenEncryptionAlg returns a boolean if a field has been set.

### GetIdTokenEncryptionEnc

`func (o *Client) GetIdTokenEncryptionEnc() JweEnc`

GetIdTokenEncryptionEnc returns the IdTokenEncryptionEnc field if non-nil, zero value otherwise.

### GetIdTokenEncryptionEncOk

`func (o *Client) GetIdTokenEncryptionEncOk() (*JweEnc, bool)`

GetIdTokenEncryptionEncOk returns a tuple with the IdTokenEncryptionEnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenEncryptionEnc

`func (o *Client) SetIdTokenEncryptionEnc(v JweEnc)`

SetIdTokenEncryptionEnc sets IdTokenEncryptionEnc field to given value.

### HasIdTokenEncryptionEnc

`func (o *Client) HasIdTokenEncryptionEnc() bool`

HasIdTokenEncryptionEnc returns a boolean if a field has been set.

### GetAuthTimeRequired

`func (o *Client) GetAuthTimeRequired() bool`

GetAuthTimeRequired returns the AuthTimeRequired field if non-nil, zero value otherwise.

### GetAuthTimeRequiredOk

`func (o *Client) GetAuthTimeRequiredOk() (*bool, bool)`

GetAuthTimeRequiredOk returns a tuple with the AuthTimeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTimeRequired

`func (o *Client) SetAuthTimeRequired(v bool)`

SetAuthTimeRequired sets AuthTimeRequired field to given value.

### HasAuthTimeRequired

`func (o *Client) HasAuthTimeRequired() bool`

HasAuthTimeRequired returns a boolean if a field has been set.

### GetSubjectType

`func (o *Client) GetSubjectType() SubjectType`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *Client) GetSubjectTypeOk() (*SubjectType, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *Client) SetSubjectType(v SubjectType)`

SetSubjectType sets SubjectType field to given value.

### HasSubjectType

`func (o *Client) HasSubjectType() bool`

HasSubjectType returns a boolean if a field has been set.

### GetSectorIdentifierUri

`func (o *Client) GetSectorIdentifierUri() string`

GetSectorIdentifierUri returns the SectorIdentifierUri field if non-nil, zero value otherwise.

### GetSectorIdentifierUriOk

`func (o *Client) GetSectorIdentifierUriOk() (*string, bool)`

GetSectorIdentifierUriOk returns a tuple with the SectorIdentifierUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorIdentifierUri

`func (o *Client) SetSectorIdentifierUri(v string)`

SetSectorIdentifierUri sets SectorIdentifierUri field to given value.

### HasSectorIdentifierUri

`func (o *Client) HasSectorIdentifierUri() bool`

HasSectorIdentifierUri returns a boolean if a field has been set.

### GetDerivedSectorIdentifier

`func (o *Client) GetDerivedSectorIdentifier() string`

GetDerivedSectorIdentifier returns the DerivedSectorIdentifier field if non-nil, zero value otherwise.

### GetDerivedSectorIdentifierOk

`func (o *Client) GetDerivedSectorIdentifierOk() (*string, bool)`

GetDerivedSectorIdentifierOk returns a tuple with the DerivedSectorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedSectorIdentifier

`func (o *Client) SetDerivedSectorIdentifier(v string)`

SetDerivedSectorIdentifier sets DerivedSectorIdentifier field to given value.

### HasDerivedSectorIdentifier

`func (o *Client) HasDerivedSectorIdentifier() bool`

HasDerivedSectorIdentifier returns a boolean if a field has been set.

### GetJwksUri

`func (o *Client) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *Client) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *Client) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *Client) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetJwks

`func (o *Client) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *Client) GetJwksOk() (*string, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *Client) SetJwks(v string)`

SetJwks sets Jwks field to given value.

### HasJwks

`func (o *Client) HasJwks() bool`

HasJwks returns a boolean if a field has been set.

### GetUserInfoSignAlg

`func (o *Client) GetUserInfoSignAlg() JwsAlg`

GetUserInfoSignAlg returns the UserInfoSignAlg field if non-nil, zero value otherwise.

### GetUserInfoSignAlgOk

`func (o *Client) GetUserInfoSignAlgOk() (*JwsAlg, bool)`

GetUserInfoSignAlgOk returns a tuple with the UserInfoSignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoSignAlg

`func (o *Client) SetUserInfoSignAlg(v JwsAlg)`

SetUserInfoSignAlg sets UserInfoSignAlg field to given value.

### HasUserInfoSignAlg

`func (o *Client) HasUserInfoSignAlg() bool`

HasUserInfoSignAlg returns a boolean if a field has been set.

### GetUserInfoEncryptionAlg

`func (o *Client) GetUserInfoEncryptionAlg() JweAlg`

GetUserInfoEncryptionAlg returns the UserInfoEncryptionAlg field if non-nil, zero value otherwise.

### GetUserInfoEncryptionAlgOk

`func (o *Client) GetUserInfoEncryptionAlgOk() (*JweAlg, bool)`

GetUserInfoEncryptionAlgOk returns a tuple with the UserInfoEncryptionAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoEncryptionAlg

`func (o *Client) SetUserInfoEncryptionAlg(v JweAlg)`

SetUserInfoEncryptionAlg sets UserInfoEncryptionAlg field to given value.

### HasUserInfoEncryptionAlg

`func (o *Client) HasUserInfoEncryptionAlg() bool`

HasUserInfoEncryptionAlg returns a boolean if a field has been set.

### GetUserInfoEncryptionEnc

`func (o *Client) GetUserInfoEncryptionEnc() JweEnc`

GetUserInfoEncryptionEnc returns the UserInfoEncryptionEnc field if non-nil, zero value otherwise.

### GetUserInfoEncryptionEncOk

`func (o *Client) GetUserInfoEncryptionEncOk() (*JweEnc, bool)`

GetUserInfoEncryptionEncOk returns a tuple with the UserInfoEncryptionEnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoEncryptionEnc

`func (o *Client) SetUserInfoEncryptionEnc(v JweEnc)`

SetUserInfoEncryptionEnc sets UserInfoEncryptionEnc field to given value.

### HasUserInfoEncryptionEnc

`func (o *Client) HasUserInfoEncryptionEnc() bool`

HasUserInfoEncryptionEnc returns a boolean if a field has been set.

### GetLoginUri

`func (o *Client) GetLoginUri() string`

GetLoginUri returns the LoginUri field if non-nil, zero value otherwise.

### GetLoginUriOk

`func (o *Client) GetLoginUriOk() (*string, bool)`

GetLoginUriOk returns a tuple with the LoginUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUri

`func (o *Client) SetLoginUri(v string)`

SetLoginUri sets LoginUri field to given value.

### HasLoginUri

`func (o *Client) HasLoginUri() bool`

HasLoginUri returns a boolean if a field has been set.

### GetTosUri

`func (o *Client) GetTosUri() string`

GetTosUri returns the TosUri field if non-nil, zero value otherwise.

### GetTosUriOk

`func (o *Client) GetTosUriOk() (*string, bool)`

GetTosUriOk returns a tuple with the TosUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosUri

`func (o *Client) SetTosUri(v string)`

SetTosUri sets TosUri field to given value.

### HasTosUri

`func (o *Client) HasTosUri() bool`

HasTosUri returns a boolean if a field has been set.

### GetTosUris

`func (o *Client) GetTosUris() []TaggedValue`

GetTosUris returns the TosUris field if non-nil, zero value otherwise.

### GetTosUrisOk

`func (o *Client) GetTosUrisOk() (*[]TaggedValue, bool)`

GetTosUrisOk returns a tuple with the TosUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosUris

`func (o *Client) SetTosUris(v []TaggedValue)`

SetTosUris sets TosUris field to given value.

### HasTosUris

`func (o *Client) HasTosUris() bool`

HasTosUris returns a boolean if a field has been set.

### GetPolicyUri

`func (o *Client) GetPolicyUri() string`

GetPolicyUri returns the PolicyUri field if non-nil, zero value otherwise.

### GetPolicyUriOk

`func (o *Client) GetPolicyUriOk() (*string, bool)`

GetPolicyUriOk returns a tuple with the PolicyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUri

`func (o *Client) SetPolicyUri(v string)`

SetPolicyUri sets PolicyUri field to given value.

### HasPolicyUri

`func (o *Client) HasPolicyUri() bool`

HasPolicyUri returns a boolean if a field has been set.

### GetPolicyUris

`func (o *Client) GetPolicyUris() []TaggedValue`

GetPolicyUris returns the PolicyUris field if non-nil, zero value otherwise.

### GetPolicyUrisOk

`func (o *Client) GetPolicyUrisOk() (*[]TaggedValue, bool)`

GetPolicyUrisOk returns a tuple with the PolicyUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUris

`func (o *Client) SetPolicyUris(v []TaggedValue)`

SetPolicyUris sets PolicyUris field to given value.

### HasPolicyUris

`func (o *Client) HasPolicyUris() bool`

HasPolicyUris returns a boolean if a field has been set.

### GetClientUri

`func (o *Client) GetClientUri() string`

GetClientUri returns the ClientUri field if non-nil, zero value otherwise.

### GetClientUriOk

`func (o *Client) GetClientUriOk() (*string, bool)`

GetClientUriOk returns a tuple with the ClientUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUri

`func (o *Client) SetClientUri(v string)`

SetClientUri sets ClientUri field to given value.

### HasClientUri

`func (o *Client) HasClientUri() bool`

HasClientUri returns a boolean if a field has been set.

### GetClientUris

`func (o *Client) GetClientUris() []TaggedValue`

GetClientUris returns the ClientUris field if non-nil, zero value otherwise.

### GetClientUrisOk

`func (o *Client) GetClientUrisOk() (*[]TaggedValue, bool)`

GetClientUrisOk returns a tuple with the ClientUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUris

`func (o *Client) SetClientUris(v []TaggedValue)`

SetClientUris sets ClientUris field to given value.

### HasClientUris

`func (o *Client) HasClientUris() bool`

HasClientUris returns a boolean if a field has been set.

### GetBcDeliveryMode

`func (o *Client) GetBcDeliveryMode() string`

GetBcDeliveryMode returns the BcDeliveryMode field if non-nil, zero value otherwise.

### GetBcDeliveryModeOk

`func (o *Client) GetBcDeliveryModeOk() (*string, bool)`

GetBcDeliveryModeOk returns a tuple with the BcDeliveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcDeliveryMode

`func (o *Client) SetBcDeliveryMode(v string)`

SetBcDeliveryMode sets BcDeliveryMode field to given value.

### HasBcDeliveryMode

`func (o *Client) HasBcDeliveryMode() bool`

HasBcDeliveryMode returns a boolean if a field has been set.

### GetBcNotificationEndpoint

`func (o *Client) GetBcNotificationEndpoint() string`

GetBcNotificationEndpoint returns the BcNotificationEndpoint field if non-nil, zero value otherwise.

### GetBcNotificationEndpointOk

`func (o *Client) GetBcNotificationEndpointOk() (*string, bool)`

GetBcNotificationEndpointOk returns a tuple with the BcNotificationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcNotificationEndpoint

`func (o *Client) SetBcNotificationEndpoint(v string)`

SetBcNotificationEndpoint sets BcNotificationEndpoint field to given value.

### HasBcNotificationEndpoint

`func (o *Client) HasBcNotificationEndpoint() bool`

HasBcNotificationEndpoint returns a boolean if a field has been set.

### GetBcRequestSignAlg

`func (o *Client) GetBcRequestSignAlg() JwsAlg`

GetBcRequestSignAlg returns the BcRequestSignAlg field if non-nil, zero value otherwise.

### GetBcRequestSignAlgOk

`func (o *Client) GetBcRequestSignAlgOk() (*JwsAlg, bool)`

GetBcRequestSignAlgOk returns a tuple with the BcRequestSignAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcRequestSignAlg

`func (o *Client) SetBcRequestSignAlg(v JwsAlg)`

SetBcRequestSignAlg sets BcRequestSignAlg field to given value.

### HasBcRequestSignAlg

`func (o *Client) HasBcRequestSignAlg() bool`

HasBcRequestSignAlg returns a boolean if a field has been set.

### GetBcUserCodeRequired

`func (o *Client) GetBcUserCodeRequired() bool`

GetBcUserCodeRequired returns the BcUserCodeRequired field if non-nil, zero value otherwise.

### GetBcUserCodeRequiredOk

`func (o *Client) GetBcUserCodeRequiredOk() (*bool, bool)`

GetBcUserCodeRequiredOk returns a tuple with the BcUserCodeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcUserCodeRequired

`func (o *Client) SetBcUserCodeRequired(v bool)`

SetBcUserCodeRequired sets BcUserCodeRequired field to given value.

### HasBcUserCodeRequired

`func (o *Client) HasBcUserCodeRequired() bool`

HasBcUserCodeRequired returns a boolean if a field has been set.

### GetAttributes

`func (o *Client) GetAttributes() []Pair`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Client) GetAttributesOk() (*[]Pair, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Client) SetAttributes(v []Pair)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Client) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetExtension

`func (o *Client) GetExtension() ClientExtension`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *Client) GetExtensionOk() (*ClientExtension, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *Client) SetExtension(v ClientExtension)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *Client) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetAuthorizationDetailsTypes

`func (o *Client) GetAuthorizationDetailsTypes() []string`

GetAuthorizationDetailsTypes returns the AuthorizationDetailsTypes field if non-nil, zero value otherwise.

### GetAuthorizationDetailsTypesOk

`func (o *Client) GetAuthorizationDetailsTypesOk() (*[]string, bool)`

GetAuthorizationDetailsTypesOk returns a tuple with the AuthorizationDetailsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetailsTypes

`func (o *Client) SetAuthorizationDetailsTypes(v []string)`

SetAuthorizationDetailsTypes sets AuthorizationDetailsTypes field to given value.

### HasAuthorizationDetailsTypes

`func (o *Client) HasAuthorizationDetailsTypes() bool`

HasAuthorizationDetailsTypes returns a boolean if a field has been set.

### GetCustomMetadata

`func (o *Client) GetCustomMetadata() string`

GetCustomMetadata returns the CustomMetadata field if non-nil, zero value otherwise.

### GetCustomMetadataOk

`func (o *Client) GetCustomMetadataOk() (*string, bool)`

GetCustomMetadataOk returns a tuple with the CustomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetadata

`func (o *Client) SetCustomMetadata(v string)`

SetCustomMetadata sets CustomMetadata field to given value.

### HasCustomMetadata

`func (o *Client) HasCustomMetadata() bool`

HasCustomMetadata returns a boolean if a field has been set.

### GetFrontChannelRequestObjectEncryptionRequired

`func (o *Client) GetFrontChannelRequestObjectEncryptionRequired() bool`

GetFrontChannelRequestObjectEncryptionRequired returns the FrontChannelRequestObjectEncryptionRequired field if non-nil, zero value otherwise.

### GetFrontChannelRequestObjectEncryptionRequiredOk

`func (o *Client) GetFrontChannelRequestObjectEncryptionRequiredOk() (*bool, bool)`

GetFrontChannelRequestObjectEncryptionRequiredOk returns a tuple with the FrontChannelRequestObjectEncryptionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontChannelRequestObjectEncryptionRequired

`func (o *Client) SetFrontChannelRequestObjectEncryptionRequired(v bool)`

SetFrontChannelRequestObjectEncryptionRequired sets FrontChannelRequestObjectEncryptionRequired field to given value.

### HasFrontChannelRequestObjectEncryptionRequired

`func (o *Client) HasFrontChannelRequestObjectEncryptionRequired() bool`

HasFrontChannelRequestObjectEncryptionRequired returns a boolean if a field has been set.

### GetRequestObjectEncryptionAlgMatchRequired

`func (o *Client) GetRequestObjectEncryptionAlgMatchRequired() bool`

GetRequestObjectEncryptionAlgMatchRequired returns the RequestObjectEncryptionAlgMatchRequired field if non-nil, zero value otherwise.

### GetRequestObjectEncryptionAlgMatchRequiredOk

`func (o *Client) GetRequestObjectEncryptionAlgMatchRequiredOk() (*bool, bool)`

GetRequestObjectEncryptionAlgMatchRequiredOk returns a tuple with the RequestObjectEncryptionAlgMatchRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectEncryptionAlgMatchRequired

`func (o *Client) SetRequestObjectEncryptionAlgMatchRequired(v bool)`

SetRequestObjectEncryptionAlgMatchRequired sets RequestObjectEncryptionAlgMatchRequired field to given value.

### HasRequestObjectEncryptionAlgMatchRequired

`func (o *Client) HasRequestObjectEncryptionAlgMatchRequired() bool`

HasRequestObjectEncryptionAlgMatchRequired returns a boolean if a field has been set.

### GetRequestObjectEncryptionEncMatchRequired

`func (o *Client) GetRequestObjectEncryptionEncMatchRequired() bool`

GetRequestObjectEncryptionEncMatchRequired returns the RequestObjectEncryptionEncMatchRequired field if non-nil, zero value otherwise.

### GetRequestObjectEncryptionEncMatchRequiredOk

`func (o *Client) GetRequestObjectEncryptionEncMatchRequiredOk() (*bool, bool)`

GetRequestObjectEncryptionEncMatchRequiredOk returns a tuple with the RequestObjectEncryptionEncMatchRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectEncryptionEncMatchRequired

`func (o *Client) SetRequestObjectEncryptionEncMatchRequired(v bool)`

SetRequestObjectEncryptionEncMatchRequired sets RequestObjectEncryptionEncMatchRequired field to given value.

### HasRequestObjectEncryptionEncMatchRequired

`func (o *Client) HasRequestObjectEncryptionEncMatchRequired() bool`

HasRequestObjectEncryptionEncMatchRequired returns a boolean if a field has been set.

### GetDigestAlgorithm

`func (o *Client) GetDigestAlgorithm() string`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *Client) GetDigestAlgorithmOk() (*string, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *Client) SetDigestAlgorithm(v string)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *Client) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetSingleAccessTokenPerSubject

`func (o *Client) GetSingleAccessTokenPerSubject() bool`

GetSingleAccessTokenPerSubject returns the SingleAccessTokenPerSubject field if non-nil, zero value otherwise.

### GetSingleAccessTokenPerSubjectOk

`func (o *Client) GetSingleAccessTokenPerSubjectOk() (*bool, bool)`

GetSingleAccessTokenPerSubjectOk returns a tuple with the SingleAccessTokenPerSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleAccessTokenPerSubject

`func (o *Client) SetSingleAccessTokenPerSubject(v bool)`

SetSingleAccessTokenPerSubject sets SingleAccessTokenPerSubject field to given value.

### HasSingleAccessTokenPerSubject

`func (o *Client) HasSingleAccessTokenPerSubject() bool`

HasSingleAccessTokenPerSubject returns a boolean if a field has been set.

### GetPkceRequired

`func (o *Client) GetPkceRequired() bool`

GetPkceRequired returns the PkceRequired field if non-nil, zero value otherwise.

### GetPkceRequiredOk

`func (o *Client) GetPkceRequiredOk() (*bool, bool)`

GetPkceRequiredOk returns a tuple with the PkceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceRequired

`func (o *Client) SetPkceRequired(v bool)`

SetPkceRequired sets PkceRequired field to given value.

### HasPkceRequired

`func (o *Client) HasPkceRequired() bool`

HasPkceRequired returns a boolean if a field has been set.

### GetPkceS256Required

`func (o *Client) GetPkceS256Required() bool`

GetPkceS256Required returns the PkceS256Required field if non-nil, zero value otherwise.

### GetPkceS256RequiredOk

`func (o *Client) GetPkceS256RequiredOk() (*bool, bool)`

GetPkceS256RequiredOk returns a tuple with the PkceS256Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceS256Required

`func (o *Client) SetPkceS256Required(v bool)`

SetPkceS256Required sets PkceS256Required field to given value.

### HasPkceS256Required

`func (o *Client) HasPkceS256Required() bool`

HasPkceS256Required returns a boolean if a field has been set.

### GetDpopRequired

`func (o *Client) GetDpopRequired() bool`

GetDpopRequired returns the DpopRequired field if non-nil, zero value otherwise.

### GetDpopRequiredOk

`func (o *Client) GetDpopRequiredOk() (*bool, bool)`

GetDpopRequiredOk returns a tuple with the DpopRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpopRequired

`func (o *Client) SetDpopRequired(v bool)`

SetDpopRequired sets DpopRequired field to given value.

### HasDpopRequired

`func (o *Client) HasDpopRequired() bool`

HasDpopRequired returns a boolean if a field has been set.

### GetAutomaticallyRegistered

`func (o *Client) GetAutomaticallyRegistered() bool`

GetAutomaticallyRegistered returns the AutomaticallyRegistered field if non-nil, zero value otherwise.

### GetAutomaticallyRegisteredOk

`func (o *Client) GetAutomaticallyRegisteredOk() (*bool, bool)`

GetAutomaticallyRegisteredOk returns a tuple with the AutomaticallyRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyRegistered

`func (o *Client) SetAutomaticallyRegistered(v bool)`

SetAutomaticallyRegistered sets AutomaticallyRegistered field to given value.

### HasAutomaticallyRegistered

`func (o *Client) HasAutomaticallyRegistered() bool`

HasAutomaticallyRegistered returns a boolean if a field has been set.

### GetExplicitlyRegistered

`func (o *Client) GetExplicitlyRegistered() bool`

GetExplicitlyRegistered returns the ExplicitlyRegistered field if non-nil, zero value otherwise.

### GetExplicitlyRegisteredOk

`func (o *Client) GetExplicitlyRegisteredOk() (*bool, bool)`

GetExplicitlyRegisteredOk returns a tuple with the ExplicitlyRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitlyRegistered

`func (o *Client) SetExplicitlyRegistered(v bool)`

SetExplicitlyRegistered sets ExplicitlyRegistered field to given value.

### HasExplicitlyRegistered

`func (o *Client) HasExplicitlyRegistered() bool`

HasExplicitlyRegistered returns a boolean if a field has been set.

### GetRsResponseSigned

`func (o *Client) GetRsResponseSigned() bool`

GetRsResponseSigned returns the RsResponseSigned field if non-nil, zero value otherwise.

### GetRsResponseSignedOk

`func (o *Client) GetRsResponseSignedOk() (*bool, bool)`

GetRsResponseSignedOk returns a tuple with the RsResponseSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsResponseSigned

`func (o *Client) SetRsResponseSigned(v bool)`

SetRsResponseSigned sets RsResponseSigned field to given value.

### HasRsResponseSigned

`func (o *Client) HasRsResponseSigned() bool`

HasRsResponseSigned returns a boolean if a field has been set.

### GetRsSignedRequestKeyId

`func (o *Client) GetRsSignedRequestKeyId() string`

GetRsSignedRequestKeyId returns the RsSignedRequestKeyId field if non-nil, zero value otherwise.

### GetRsSignedRequestKeyIdOk

`func (o *Client) GetRsSignedRequestKeyIdOk() (*string, bool)`

GetRsSignedRequestKeyIdOk returns a tuple with the RsSignedRequestKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsSignedRequestKeyId

`func (o *Client) SetRsSignedRequestKeyId(v string)`

SetRsSignedRequestKeyId sets RsSignedRequestKeyId field to given value.

### HasRsSignedRequestKeyId

`func (o *Client) HasRsSignedRequestKeyId() bool`

HasRsSignedRequestKeyId returns a boolean if a field has been set.

### GetClientRegistrationTypes

`func (o *Client) GetClientRegistrationTypes() []ClientRegistrationType`

GetClientRegistrationTypes returns the ClientRegistrationTypes field if non-nil, zero value otherwise.

### GetClientRegistrationTypesOk

`func (o *Client) GetClientRegistrationTypesOk() (*[]ClientRegistrationType, bool)`

GetClientRegistrationTypesOk returns a tuple with the ClientRegistrationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRegistrationTypes

`func (o *Client) SetClientRegistrationTypes(v []ClientRegistrationType)`

SetClientRegistrationTypes sets ClientRegistrationTypes field to given value.

### HasClientRegistrationTypes

`func (o *Client) HasClientRegistrationTypes() bool`

HasClientRegistrationTypes returns a boolean if a field has been set.

### GetOrganizationName

`func (o *Client) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *Client) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *Client) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *Client) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetSignedJwksUri

`func (o *Client) GetSignedJwksUri() string`

GetSignedJwksUri returns the SignedJwksUri field if non-nil, zero value otherwise.

### GetSignedJwksUriOk

`func (o *Client) GetSignedJwksUriOk() (*string, bool)`

GetSignedJwksUriOk returns a tuple with the SignedJwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedJwksUri

`func (o *Client) SetSignedJwksUri(v string)`

SetSignedJwksUri sets SignedJwksUri field to given value.

### HasSignedJwksUri

`func (o *Client) HasSignedJwksUri() bool`

HasSignedJwksUri returns a boolean if a field has been set.

### GetEntityId

`func (o *Client) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Client) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Client) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *Client) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetTrustAnchorId

`func (o *Client) GetTrustAnchorId() string`

GetTrustAnchorId returns the TrustAnchorId field if non-nil, zero value otherwise.

### GetTrustAnchorIdOk

`func (o *Client) GetTrustAnchorIdOk() (*string, bool)`

GetTrustAnchorIdOk returns a tuple with the TrustAnchorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAnchorId

`func (o *Client) SetTrustAnchorId(v string)`

SetTrustAnchorId sets TrustAnchorId field to given value.

### HasTrustAnchorId

`func (o *Client) HasTrustAnchorId() bool`

HasTrustAnchorId returns a boolean if a field has been set.

### GetTrustChain

`func (o *Client) GetTrustChain() []string`

GetTrustChain returns the TrustChain field if non-nil, zero value otherwise.

### GetTrustChainOk

`func (o *Client) GetTrustChainOk() (*[]string, bool)`

GetTrustChainOk returns a tuple with the TrustChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustChain

`func (o *Client) SetTrustChain(v []string)`

SetTrustChain sets TrustChain field to given value.

### HasTrustChain

`func (o *Client) HasTrustChain() bool`

HasTrustChain returns a boolean if a field has been set.

### GetTrustChainExpiresAt

`func (o *Client) GetTrustChainExpiresAt() int64`

GetTrustChainExpiresAt returns the TrustChainExpiresAt field if non-nil, zero value otherwise.

### GetTrustChainExpiresAtOk

`func (o *Client) GetTrustChainExpiresAtOk() (*int64, bool)`

GetTrustChainExpiresAtOk returns a tuple with the TrustChainExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustChainExpiresAt

`func (o *Client) SetTrustChainExpiresAt(v int64)`

SetTrustChainExpiresAt sets TrustChainExpiresAt field to given value.

### HasTrustChainExpiresAt

`func (o *Client) HasTrustChainExpiresAt() bool`

HasTrustChainExpiresAt returns a boolean if a field has been set.

### GetTrustChainUpdatedAt

`func (o *Client) GetTrustChainUpdatedAt() int64`

GetTrustChainUpdatedAt returns the TrustChainUpdatedAt field if non-nil, zero value otherwise.

### GetTrustChainUpdatedAtOk

`func (o *Client) GetTrustChainUpdatedAtOk() (*int64, bool)`

GetTrustChainUpdatedAtOk returns a tuple with the TrustChainUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustChainUpdatedAt

`func (o *Client) SetTrustChainUpdatedAt(v int64)`

SetTrustChainUpdatedAt sets TrustChainUpdatedAt field to given value.

### HasTrustChainUpdatedAt

`func (o *Client) HasTrustChainUpdatedAt() bool`

HasTrustChainUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


