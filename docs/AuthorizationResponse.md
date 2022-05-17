# AuthorizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take. | [optional] 
**Client** | Pointer to [**Client**](Client.md) |  | [optional] 
**Display** | Pointer to [**Display**](Display.md) |  | [optional] 
**MaxAge** | Pointer to **int32** | The maximum authentication age. This value comes from &#x60;max_age&#x60; request parameter, or &#x60;defaultMaxAge&#x60; configuration parameter of the client application when the authorization request does not contain &#x60;max_age&#x60; request parameter.  See \&quot;[OpenID Connect Core 1.0, 3.1.2.1. Authentication Request](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest), max_age\&quot; for &#x60;max_age&#x60; request parameter, and see \&quot;[OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata), default_max_age\&quot; for &#x60;defaultMaxAge&#x60; configuration parameter.  | [optional] 
**Service** | Pointer to [**Service**](Service.md) |  | [optional] 
**Scopes** | Pointer to [**[]Scope**](Scope.md) | The scopes that the client application requests. This value comes from &#x60;scope&#x60; request parameter. If the request does not contain &#x60;scope&#x60; parameter, this parameter is a list of scopes which are registered as default. If the authorization request does not have &#x60;scope&#x60; request parameter and the service has not registered any default scope, the value of this parameter is &#x60;null&#x60;.  It is ensured that scopes listed by this parameters are contained in the list of supported scopes which are specified by &#x60;supportedScopes&#x60; configuration parameter of the service. Unsupported scopes in the authorization request do not cause an error and are just ignored.  OpenID Connect defines some scope names which need to be treated specially. The table below lists the special scope names.  | Name | Description | | --- | --- | | &#x60;openid&#x60; | This scope must be contained in &#x60;scope&#x60; request parameter to promote an OAuth 2.0 authorization request to an OpenID Connect request. It is described in \&quot;[OpenID Connect Core 1.0, 3.1.2.1. Authentication Request](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest), scope\&quot;. | | &#x60;profile&#x60; | This scope is used to request some claims to be embedded in the ID token. The claims are &#x60;name&#x60;, &#x60;family_name&#x60;, &#x60;given_name&#x60;, &#x60;middle_name&#x60;, &#x60;nickname&#x60;, &#x60;preferred_username&#x60;, &#x60;profile&#x60;, &#x60;picture&#x60;, &#x60;website&#x60;, &#x60;gender&#x60;, &#x60;birthdate&#x60;, &#x60;zoneinfo&#x60;, &#x60;locale&#x60;, and &#x60;updated_at&#x60;. It is described in [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims). | | &#x60;email&#x60; | This scope is used to request some claims to be embedded in the ID token. The claims are &#x60;email&#x60; and &#x60;email_verified&#x60;. It is described in [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims). | | &#x60;address&#x60; |  This scope is used to request &#x60;address&#x60; claim to be embedded in the ID token. It is described in [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims).&lt;br&gt;&lt;br&gt; The format of &#x60;address&#x60; claim is not a simple string. It is described in [OpenID Connect Core 1.0, 5.1.1. Address Claim](https://openid.net/specs/openid-connect-core-1_0.html#AddressClaim).  | | &#x60;phone&#x60; | This scope is used to request some claims to be embedded in the ID token. The claims are &#x60;phone_number&#x60; and &#x60;phone_number_verified&#x60;. It is described in [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims).  | | &#x60;offline_access&#x60; | The following is an excerpt about this scope from [OpenID Connect Core 1.0, 11. Offline Access](https://openid.net/specs/openid-connect-core-1_0.html#OfflineAccess). &lt;blockquote&gt;This scope value requests that an OAuth 2.0 Refresh Token be issued that can be used to obtain an Access Token that grants access to the end-user&#39;s userinfo endpoint even when the end-user is not present (not logged in).&lt;/blockquote&gt;  |  Note that, if &#x60;response_type&#x60; request parameter does not contain code, &#x60;offline_acccess&#x60; scope is removed from this list even when scope request parameter contains &#x60;offline_access&#x60;. This behavior is a requirement written in [OpenID Connect Core 1.0, 11. Offline Access](https://openid.net/specs/openid-connect-core-1_0.html#OfflineAccess).  | [optional] 
**UiLocales** | Pointer to **[]string** | The locales that the client application presented as candidates to be used for UI. This value comes from &#x60;ui_locales&#x60; request parameter. The format of &#x60;ui_locales&#x60; is a space-separated list of language tag values defined in [RFC5646](https://datatracker.ietf.org/doc/html/rfc5646). See \&quot;[OpenID Connect Core 1.0, 3.1.2.1. Authentication Request](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest), ui_locales\&quot; for details.  It is ensured that locales listed by this parameters are contained in the list of supported UI locales which are specified by &#x60;supportedUiLocales&#x60; configuration parameter of the service. Unsupported UI locales in the authorization request do not cause an error and are just ignored.  | [optional] 
**ClaimsLocales** | Pointer to **[]string** | End-user&#39;s preferred languages and scripts for claims. This value comes from &#x60;claims_locales&#x60; request parameter. The format of &#x60;claims_locales&#x60; is a space-separated list of language tag values defined in [RFC5646](https://datatracker.ietf.org/doc/html/rfc5646). See \&quot;[OpenID Connect Core 1.0, 5.2. Claims Languages and Scripts](https://openid.net/specs/openid-connect-core-1_0.html#ClaimsLanguagesAndScripts)\&quot; for details.  It is ensured that locales listed by this parameters are contained in the list of supported claim locales which are specified by &#x60;supportedClaimsLocales&#x60; configuration parameter of the service. Unsupported claim locales in the authorization request do not cause an error and are just ignored.  | [optional] 
**Claims** | Pointer to **[]string** | The list of claims that the client application requests to be embedded in the ID token. The value comes from (1) &#x60;id_token&#x60; in &#x60;claims&#x60; request parameter [1] and/or (2) special scopes (&#x60;profile&#x60;, &#x60;email&#x60;, &#x60;address&#x60; and &#x60;phone&#x60;) which are expanded to claims.  See [OpenID Connect Core 1.0, 5.5. Requesting Claims using the \&quot;claims\&quot; Request Parameter](https://openid.net/specs/openid-connect-core-1_0.html#ClaimsParameter) for &#x60;claims&#x60; request parameter, and see [OpenID Connect Core 1.0, 5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) for the special scopes.  | [optional] 
**AcrEssential** | Pointer to **bool** | This boolean value indicates whether the authentication of the end-user must be one of the ACRs (Authentication Context Class References) listed in &#x60;acrs&#x60; parameter. This parameter becomes &#x60;true&#x60; only when (1) the authorization request contains &#x60;claims&#x60; request parameter and (2) &#x60;acr&#x60; claim is in it, and (3) &#x60;essential&#x60; property of the &#x60;acr&#x60; claim is &#x60;true&#x60;. See [OpenID Connect Core 1.0, 5.5.1.1. Requesting the \&quot;acr\&quot; Claim](https://openid.net/specs/openid-connect-core-1_0.html#acrSemantics) for details.  | [optional] 
**ClientIdAliasUsed** | Pointer to **bool** | &#x60;true&#x60; if the value of the &#x60;client_id&#x60; request parameter included in the authorization request is the client ID alias. &#x60;false&#x60; if the value is the original numeric client ID.  | [optional] 
**Acrs** | Pointer to **[]string** | The list of ACRs (Authentication Context Class References) one of which the client application requests to be satisfied for the authentication of the end-user. This value comes from &#x60;acr_values&#x60; request parameter or &#x60;defaultAcrs&#x60; configuration parameter of the client application.  See \&quot;[OpenID Connect Core 1.0, 3.1.2.1. Authentication Request](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest), acr_values\&quot; for &#x60;acr_values&#x60; request parameter, and see \&quot;[OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata), default_acr_values\&quot; for &#x60;defaultAcrs&#x60; configuration parameter.  | [optional] 
**Subject** | Pointer to **string** | The subject (&#x3D; unique user ID managed by the authorization server implementation) that the client application expects to grant authorization. The value comes from &#x60;sub&#x60; claim in &#x60;claims&#x60; request parameter.  | [optional] 
**LoginHint** | Pointer to **string** | A hint about the login identifier of the end-user. The value comes from &#x60;login_hint&#x60; request parameter. | [optional] 
**Prompts** | Pointer to [**[]Prompt**](Prompt.md) | The list of values of prompt request parameter. See \&quot;[OpenID Connect Core 1.0, 3.1.2.1. Authentication Request](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest), prompt\&quot; for prompt request parameter. | [optional] 
**LowestPrompt** | Pointer to [**Prompt**](Prompt.md) |  | [optional] 
**RequestObjectPayload** | Pointer to **string** | The payload part of the request object. The value of this proprty is &#x60;null&#x60; if the authorization request does not include a request object.  | [optional] 
**IdTokenClaims** | Pointer to **string** | The value of the &#x60;id_token&#x60; property in the claims request parameter or in the claims property in a request object.  A client application may request certain claims be embedded in an ID token or in a response from the userInfo endpoint. There are several ways. Including the &#x60;claims&#x60; request parameter and including the &#x60;claims&#x60; property in a request object are such examples. In both the cases, the value of the &#x60;claims&#x60; parameter/property is JSON. Its format is described in [5.5. Requesting Claims using the \&quot;claims\&quot; Request Parameter](https://openid.net/specs/openid-connect-core-1_0.html#ClaimsParameter).  The following is an excerpt from the specification. You can find &#x60;userinfo&#x60; and &#x60;id_token&#x60; are top-level properties.  &#x60;&#x60;&#x60;json {   \&quot;userinfo\&quot;:   {     \&quot;given_name\&quot;: { \&quot;essential\&quot;: true },     \&quot;nickname\&quot;: null,     \&quot;email\&quot;: { \&quot;essential\&quot;: true },     \&quot;email_verified\&quot;: { \&quot;essential\&quot;: true },     \&quot;picture\&quot;: null,     \&quot;http://example.info/claims/groups\&quot;: null   },   \&quot;id_token\&quot;:   {     \&quot;auth_time\&quot;: { \&quot;essential\&quot;: true },     \&quot;acr\&quot;: { \&quot;values\&quot;: [ \&quot;urn:mace:incommon:iap:silver\&quot; ] }   } } &#x60;&#x60;&#x60;  This value of this property is the value of the &#x60;id_token&#x60; property in JSON format. For example, if the JSON above is included in an authorization request, this property holds JSON equivalent to the following.  &#x60;&#x60;&#x60;json {   \&quot;auth_time\&quot;: { \&quot;essential\&quot;: true },   \&quot;acr\&quot;: { \&quot;values\&quot;: [ \&quot;urn:mace:incommon:iap:silver\&quot; ] } } &#x60;&#x60;&#x60;  Note that if a request object is given and it contains the &#x60;claims&#x60; property and if the &#x60;claims&#x60; request parameter is also given, this property holds the former value.  | [optional] 
**UserInfoClaims** | Pointer to **string** | The value of the &#x60;userinfo&#x60; property in the &#x60;claims&#x60; request parameter or in the &#x60;claims&#x60; property in a request object.  A client application may request certain claims be embedded in an ID token or in a response from the userInfo endpoint. There are several ways. Including the &#x60;claims&#x60; request parameter and including the &#x60;claims&#x60; property in a request object are such examples. In both the cases, the value of the &#x60;claims&#x60; parameter/property is JSON. Its format is described in [5.5. Requesting Claims using the \&quot;claims\&quot; Request Parameter](https://openid.net/specs/openid-connect-core-1_0.html#ClaimsParameter).  The following is an excerpt from the specification. You can find &#x60;userinfo&#x60; and &#x60;id_token&#x60; are top-level properties.  &#x60;&#x60;&#x60;json {   \&quot;userinfo\&quot;:   {     \&quot;given_name\&quot;: { \&quot;essential\&quot;: true },     \&quot;nickname\&quot;: null,     \&quot;email\&quot;: { \&quot;essential\&quot;: true },     \&quot;email_verified\&quot;: { \&quot;essential\&quot;: true },     \&quot;picture\&quot;: null,     \&quot;http://example.info/claims/groups\&quot;: null   },   \&quot;id_token\&quot;:   {     \&quot;auth_time\&quot;: { \&quot;essential\&quot;: true },     \&quot;acr\&quot;: { \&quot;values\&quot;: [ \&quot;urn:mace:incommon:iap:silver\&quot; ] }   } } &#x60;&#x60;&#x60;&#x60;  The value of this property is the value of the &#x60;userinfo&#x60; property in JSON format. For example, if the JSON above is included in an authorization request, this property holds JSON equivalent to the following.  &#x60;&#x60;&#x60;json {   \&quot;given_name\&quot;: { \&quot;essential\&quot;: true },   \&quot;nickname\&quot;: null,   \&quot;email\&quot;: { \&quot;essential\&quot;: true },   \&quot;email_verified\&quot;: { \&quot;essential\&quot;: true },   \&quot;picture\&quot;: null,   \&quot;http://example.info/claims/groups\&quot;: null } &#x60;&#x60;&#x60;  Note that if a request object is given and it contains the &#x60;claims&#x60; property and if the &#x60;claims&#x60; request parameter is also given, the value of this property holds the former value.  | [optional] 
**Resources** | Pointer to **string** | The resources specified by the &#x60;resource&#x60; request parameters or by the &#x60;resource&#x60; property in the request object. If both are given, the values in the request object should be set. See \&quot;Resource Indicators for OAuth 2.0\&quot; for details.  | [optional] 
**AuthorizationDetails** | Pointer to [**AuthorizationDetails**](AuthorizationDetails.md) |  | [optional] 
**Purpose** | Pointer to **string** | The &#x60;purpose&#x60; request parameter is defined in [9. Transaction-specific Purpose](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html#name-transaction-specific-purpos) of [OpenID Connect for Identity Assurance 1.0](https://openid.net/specs/openid-connect-4-identity-assurance-1_0.html) as follows:  &gt; purpose: OPTIONAL. String describing the purpose for obtaining certain user data from the OP. The purpose MUST NOT be shorter than 3 characters and MUST NOT be longer than 300 characters. If these rules are violated, the authentication request MUST fail and the OP returns an error invalid_request to the RP.  | [optional] 
**ResponseContent** | Pointer to **string** | The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of &#x60;action&#x60; parameter.  | [optional] 
**Ticket** | Pointer to **string** | A ticket issued by Authlete to the service implementation. This is needed when the service implementation calls either &#x60;/auth/authorization/fail&#x60; API or &#x60;/auth/authorization/issue&#x60; API.  | [optional] 
**DynamicScopes** | Pointer to [**[]DynamicScope**](DynamicScope.md) | The dynamic scopes which the client application requested by the scope request parameter.  | [optional] 

## Methods

### NewAuthorizationResponse

`func NewAuthorizationResponse() *AuthorizationResponse`

NewAuthorizationResponse instantiates a new AuthorizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationResponseWithDefaults

`func NewAuthorizationResponseWithDefaults() *AuthorizationResponse`

NewAuthorizationResponseWithDefaults instantiates a new AuthorizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *AuthorizationResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *AuthorizationResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *AuthorizationResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *AuthorizationResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *AuthorizationResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *AuthorizationResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *AuthorizationResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *AuthorizationResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *AuthorizationResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuthorizationResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuthorizationResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AuthorizationResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetClient

`func (o *AuthorizationResponse) GetClient() Client`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *AuthorizationResponse) GetClientOk() (*Client, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *AuthorizationResponse) SetClient(v Client)`

SetClient sets Client field to given value.

### HasClient

`func (o *AuthorizationResponse) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetDisplay

`func (o *AuthorizationResponse) GetDisplay() Display`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *AuthorizationResponse) GetDisplayOk() (*Display, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *AuthorizationResponse) SetDisplay(v Display)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *AuthorizationResponse) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetMaxAge

`func (o *AuthorizationResponse) GetMaxAge() int32`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *AuthorizationResponse) GetMaxAgeOk() (*int32, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *AuthorizationResponse) SetMaxAge(v int32)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *AuthorizationResponse) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetService

`func (o *AuthorizationResponse) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AuthorizationResponse) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AuthorizationResponse) SetService(v Service)`

SetService sets Service field to given value.

### HasService

`func (o *AuthorizationResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetScopes

`func (o *AuthorizationResponse) GetScopes() []Scope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AuthorizationResponse) GetScopesOk() (*[]Scope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AuthorizationResponse) SetScopes(v []Scope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AuthorizationResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetUiLocales

`func (o *AuthorizationResponse) GetUiLocales() []string`

GetUiLocales returns the UiLocales field if non-nil, zero value otherwise.

### GetUiLocalesOk

`func (o *AuthorizationResponse) GetUiLocalesOk() (*[]string, bool)`

GetUiLocalesOk returns a tuple with the UiLocales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiLocales

`func (o *AuthorizationResponse) SetUiLocales(v []string)`

SetUiLocales sets UiLocales field to given value.

### HasUiLocales

`func (o *AuthorizationResponse) HasUiLocales() bool`

HasUiLocales returns a boolean if a field has been set.

### GetClaimsLocales

`func (o *AuthorizationResponse) GetClaimsLocales() []string`

GetClaimsLocales returns the ClaimsLocales field if non-nil, zero value otherwise.

### GetClaimsLocalesOk

`func (o *AuthorizationResponse) GetClaimsLocalesOk() (*[]string, bool)`

GetClaimsLocalesOk returns a tuple with the ClaimsLocales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsLocales

`func (o *AuthorizationResponse) SetClaimsLocales(v []string)`

SetClaimsLocales sets ClaimsLocales field to given value.

### HasClaimsLocales

`func (o *AuthorizationResponse) HasClaimsLocales() bool`

HasClaimsLocales returns a boolean if a field has been set.

### GetClaims

`func (o *AuthorizationResponse) GetClaims() []string`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *AuthorizationResponse) GetClaimsOk() (*[]string, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *AuthorizationResponse) SetClaims(v []string)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *AuthorizationResponse) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetAcrEssential

`func (o *AuthorizationResponse) GetAcrEssential() bool`

GetAcrEssential returns the AcrEssential field if non-nil, zero value otherwise.

### GetAcrEssentialOk

`func (o *AuthorizationResponse) GetAcrEssentialOk() (*bool, bool)`

GetAcrEssentialOk returns a tuple with the AcrEssential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrEssential

`func (o *AuthorizationResponse) SetAcrEssential(v bool)`

SetAcrEssential sets AcrEssential field to given value.

### HasAcrEssential

`func (o *AuthorizationResponse) HasAcrEssential() bool`

HasAcrEssential returns a boolean if a field has been set.

### GetClientIdAliasUsed

`func (o *AuthorizationResponse) GetClientIdAliasUsed() bool`

GetClientIdAliasUsed returns the ClientIdAliasUsed field if non-nil, zero value otherwise.

### GetClientIdAliasUsedOk

`func (o *AuthorizationResponse) GetClientIdAliasUsedOk() (*bool, bool)`

GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasUsed

`func (o *AuthorizationResponse) SetClientIdAliasUsed(v bool)`

SetClientIdAliasUsed sets ClientIdAliasUsed field to given value.

### HasClientIdAliasUsed

`func (o *AuthorizationResponse) HasClientIdAliasUsed() bool`

HasClientIdAliasUsed returns a boolean if a field has been set.

### GetAcrs

`func (o *AuthorizationResponse) GetAcrs() []string`

GetAcrs returns the Acrs field if non-nil, zero value otherwise.

### GetAcrsOk

`func (o *AuthorizationResponse) GetAcrsOk() (*[]string, bool)`

GetAcrsOk returns a tuple with the Acrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrs

`func (o *AuthorizationResponse) SetAcrs(v []string)`

SetAcrs sets Acrs field to given value.

### HasAcrs

`func (o *AuthorizationResponse) HasAcrs() bool`

HasAcrs returns a boolean if a field has been set.

### GetSubject

`func (o *AuthorizationResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AuthorizationResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AuthorizationResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AuthorizationResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetLoginHint

`func (o *AuthorizationResponse) GetLoginHint() string`

GetLoginHint returns the LoginHint field if non-nil, zero value otherwise.

### GetLoginHintOk

`func (o *AuthorizationResponse) GetLoginHintOk() (*string, bool)`

GetLoginHintOk returns a tuple with the LoginHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginHint

`func (o *AuthorizationResponse) SetLoginHint(v string)`

SetLoginHint sets LoginHint field to given value.

### HasLoginHint

`func (o *AuthorizationResponse) HasLoginHint() bool`

HasLoginHint returns a boolean if a field has been set.

### GetPrompts

`func (o *AuthorizationResponse) GetPrompts() []Prompt`

GetPrompts returns the Prompts field if non-nil, zero value otherwise.

### GetPromptsOk

`func (o *AuthorizationResponse) GetPromptsOk() (*[]Prompt, bool)`

GetPromptsOk returns a tuple with the Prompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompts

`func (o *AuthorizationResponse) SetPrompts(v []Prompt)`

SetPrompts sets Prompts field to given value.

### HasPrompts

`func (o *AuthorizationResponse) HasPrompts() bool`

HasPrompts returns a boolean if a field has been set.

### GetLowestPrompt

`func (o *AuthorizationResponse) GetLowestPrompt() Prompt`

GetLowestPrompt returns the LowestPrompt field if non-nil, zero value otherwise.

### GetLowestPromptOk

`func (o *AuthorizationResponse) GetLowestPromptOk() (*Prompt, bool)`

GetLowestPromptOk returns a tuple with the LowestPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestPrompt

`func (o *AuthorizationResponse) SetLowestPrompt(v Prompt)`

SetLowestPrompt sets LowestPrompt field to given value.

### HasLowestPrompt

`func (o *AuthorizationResponse) HasLowestPrompt() bool`

HasLowestPrompt returns a boolean if a field has been set.

### GetRequestObjectPayload

`func (o *AuthorizationResponse) GetRequestObjectPayload() string`

GetRequestObjectPayload returns the RequestObjectPayload field if non-nil, zero value otherwise.

### GetRequestObjectPayloadOk

`func (o *AuthorizationResponse) GetRequestObjectPayloadOk() (*string, bool)`

GetRequestObjectPayloadOk returns a tuple with the RequestObjectPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectPayload

`func (o *AuthorizationResponse) SetRequestObjectPayload(v string)`

SetRequestObjectPayload sets RequestObjectPayload field to given value.

### HasRequestObjectPayload

`func (o *AuthorizationResponse) HasRequestObjectPayload() bool`

HasRequestObjectPayload returns a boolean if a field has been set.

### GetIdTokenClaims

`func (o *AuthorizationResponse) GetIdTokenClaims() string`

GetIdTokenClaims returns the IdTokenClaims field if non-nil, zero value otherwise.

### GetIdTokenClaimsOk

`func (o *AuthorizationResponse) GetIdTokenClaimsOk() (*string, bool)`

GetIdTokenClaimsOk returns a tuple with the IdTokenClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenClaims

`func (o *AuthorizationResponse) SetIdTokenClaims(v string)`

SetIdTokenClaims sets IdTokenClaims field to given value.

### HasIdTokenClaims

`func (o *AuthorizationResponse) HasIdTokenClaims() bool`

HasIdTokenClaims returns a boolean if a field has been set.

### GetUserInfoClaims

`func (o *AuthorizationResponse) GetUserInfoClaims() string`

GetUserInfoClaims returns the UserInfoClaims field if non-nil, zero value otherwise.

### GetUserInfoClaimsOk

`func (o *AuthorizationResponse) GetUserInfoClaimsOk() (*string, bool)`

GetUserInfoClaimsOk returns a tuple with the UserInfoClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoClaims

`func (o *AuthorizationResponse) SetUserInfoClaims(v string)`

SetUserInfoClaims sets UserInfoClaims field to given value.

### HasUserInfoClaims

`func (o *AuthorizationResponse) HasUserInfoClaims() bool`

HasUserInfoClaims returns a boolean if a field has been set.

### GetResources

`func (o *AuthorizationResponse) GetResources() string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AuthorizationResponse) GetResourcesOk() (*string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AuthorizationResponse) SetResources(v string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AuthorizationResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetAuthorizationDetails

`func (o *AuthorizationResponse) GetAuthorizationDetails() AuthorizationDetails`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *AuthorizationResponse) GetAuthorizationDetailsOk() (*AuthorizationDetails, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *AuthorizationResponse) SetAuthorizationDetails(v AuthorizationDetails)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *AuthorizationResponse) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.

### GetPurpose

`func (o *AuthorizationResponse) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *AuthorizationResponse) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *AuthorizationResponse) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *AuthorizationResponse) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetResponseContent

`func (o *AuthorizationResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *AuthorizationResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *AuthorizationResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *AuthorizationResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetTicket

`func (o *AuthorizationResponse) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *AuthorizationResponse) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *AuthorizationResponse) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *AuthorizationResponse) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetDynamicScopes

`func (o *AuthorizationResponse) GetDynamicScopes() []DynamicScope`

GetDynamicScopes returns the DynamicScopes field if non-nil, zero value otherwise.

### GetDynamicScopesOk

`func (o *AuthorizationResponse) GetDynamicScopesOk() (*[]DynamicScope, bool)`

GetDynamicScopesOk returns a tuple with the DynamicScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicScopes

`func (o *AuthorizationResponse) SetDynamicScopes(v []DynamicScope)`

SetDynamicScopes sets DynamicScopes field to given value.

### HasDynamicScopes

`func (o *AuthorizationResponse) HasDynamicScopes() bool`

HasDynamicScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


