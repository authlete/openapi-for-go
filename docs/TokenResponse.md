# TokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take. | [optional] 
**ResponseContent** | Pointer to **string** | The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of &#x60;action&#x60; parameter.  | [optional] 
**Username** | Pointer to **string** | The value of &#x60;username&#x60; request parameter in the token request. The client application must specify username when it uses [Resource Owner Password Grant](https://datatracker.ietf.org/doc/html/rfc6749#section-4.3). In other words, when the value of &#x60;grant_type&#x60; request parameter is &#x60;password&#x60;, &#x60;username&#x60; request parameter must come along.  This parameter has a value only if the value of &#x60;grant_type&#x60; request parameter is &#x60;password&#x60; and the token request is valid.  | [optional] 
**Password** | Pointer to **string** | The value of &#x60;password&#x60; request parameter in the token request. The client application must specify password when it uses [Resource Owner Password Grant](https://datatracker.ietf.org/doc/html/rfc6749#section-4.3). In other words, when the value of &#x60;grant_type&#x60; request parameter is &#x60;password&#x60;, &#x60;password&#x60; request parameter must come along.  This parameter has a value only if the value of &#x60;grant_type&#x60; request parameter is &#x60;password&#x60; and the token request is valid.  | [optional] 
**Ticket** | Pointer to **string** | The ticket which is necessary to call Authlete&#39;s &#x60;/auth/token/fail&#x60; API or &#x60;/auth/token/issue&#x60; API.  This parameter has a value only if the value of &#x60;grant_type&#x60; request parameter is &#x60;password&#x60; and the token request is valid.  | [optional] 
**AccessToken** | Pointer to **string** | The newly issued access token. | [optional] 
**AccessTokenExpiresAt** | Pointer to **int64** | The datetime at which the newly issued access token will expire. The value is represented in milliseconds since the Unix epoch (1970-01-01).  | [optional] 
**AccessTokenDuration** | Pointer to **int64** | The duration of the newly issued access token in seconds. | [optional] 
**RefreshToken** | Pointer to **string** | The newly issued refresh token. | [optional] 
**RefreshTokenExpiresAt** | Pointer to **int64** | The datetime at which the newly issued refresh token will expire. The value is represented in milliseconds since the Unix epoch (1970-01-01).  | [optional] 
**RefreshTokenDuration** | Pointer to **int64** | The duration of the newly issued refresh token in seconds. | [optional] 
**IdToken** | Pointer to **string** | The newly issued ID token. Note that an ID token is issued from a token endpoint only when the &#x60;response_type&#x60; request parameter of the authorization request to an authorization endpoint has contained &#x60;code&#x60; and the &#x60;scope&#x60; request parameter has contained &#x60;openid&#x60;.  | [optional] 
**GrantType** | Pointer to **string** | The grant type of the token request. | [optional] 
**ClientId** | Pointer to **int64** | The client ID. | [optional] 
**ClientIdAlias** | Pointer to **string** | The client ID alias when the token request was made. If the client did not have an alias, this parameter is &#x60;null&#x60;. Also, if the token request was invalid and it failed to identify a client, this parameter is &#x60;null&#x60;.  | [optional] 
**ClientIdAliasUsed** | Pointer to **bool** | The flag which indicates whether the client ID alias was used when the token request was made. &#x60;true&#x60; if the client ID alias was used when the token request was made.  | [optional] 
**Subject** | Pointer to **string** | The subject (&#x3D; resource owner&#39;s ID) of the access token. Even if an access token has been issued by the call of &#x60;/api/auth/token&#x60; API, this parameter is &#x60;null&#x60; if the flow of the token request was [Client Credentials Flow](https://datatracker.ietf.org/doc/html/rfc6749#section-4.4) (&#x60;grant_type&#x3D;client_credentials&#x60;) because it means the access token is not associated with any specific end-user.  | [optional] 
**Scopes** | Pointer to **[]string** | The scopes covered by the access token. | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) | The extra properties associated with the access token. This parameter is &#x60;null&#x60; when no extra property is associated with the issued access token.  | [optional] 
**JwtAccessToken** | Pointer to **string** | The newly issued access token in JWT format. If the authorization server is configured to issue JWT-based access tokens (&#x3D; if the service&#39;s &#x60;accessTokenSignAlg&#x60; value is a non-null value), a JWT-based access token is issued along with the original random-string one.  | [optional] 
**Resources** | Pointer to **[]string** | The resources specified by the &#x60;resource&#x60; request parameters in the token request. See \&quot;Resource Indicators for OAuth 2.0\&quot; for details.  | [optional] 
**AccessTokenResources** | Pointer to **[]string** | The target resources of the access token being issued. See \&quot;Resource Indicators for OAuth 2.0\&quot; for details.  | [optional] 
**AuthorizationDetails** | Pointer to [**AuthzDetails**](AuthzDetails.md) |  | [optional] 
**ServiceAttributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of this service that the client application belongs to.  | [optional] 
**ClientAttributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of the client.  | [optional] 
**ClientAuthMethod** | Pointer to **string** | The client authentication method that was performed at the token endpoint.  | [optional] 
**GrantId** | Pointer to **string** | the value of the &#x60;grant_id&#x60; request parameter of the device authorization request.  The &#x60;grant_id&#x60; request parameter is defined in [Grant Management for OAuth 2.0](https://openid.net/specs/fapi-grant-management.html) , which is supported by Authlete 2.3 and newer versions.  | [optional] 
**Audiences** | Pointer to **[]string** | The audiences on the token exchange request  | [optional] 
**RequestedTokenType** | Pointer to [**TokenType**](TokenType.md) |  | [optional] 
**SubjectToken** | Pointer to **string** |  | [optional] 
**SubjectTokenType** | Pointer to [**TokenType**](TokenType.md) |  | [optional] 
**SubjectTokenInfo** | Pointer to [**TokenInfo**](TokenInfo.md) |  | [optional] 
**ActorToken** | Pointer to **string** |  | [optional] 
**ActorTokenType** | Pointer to [**TokenType**](TokenType.md) |  | [optional] 
**ActorTokenInfo** | Pointer to [**TokenInfo**](TokenInfo.md) |  | [optional] 
**Assertion** | Pointer to **string** | For RFC 7523 JSON Web Token (JWT) Profile for OAuth 2.0 Client Authentication and Authorization Grants  | [optional] 
**PreviousRefreshTokenUsed** | Pointer to **bool** | Indicate whether the previous refresh token that had been kept in the database for a short time was used  | [optional] 
**ClientEntityId** | Pointer to **string** | The entity ID of the client.  | [optional] 
**ClientEntityIdUsed** | Pointer to **bool** | Flag which indicates whether the entity ID of the client was used when the request for the access token was made.  | [optional] 
**CnonceDuration** | Pointer to **int64** | Duration of the &#x60;c_nonce&#x60; in seconds.  | [optional] 
**DpopNonce** | Pointer to **string** | Get the expected nonce value for DPoP proof JWT, which should be used as the value of the &#x60;DPoP-Nonce&#x60; HTTP header.  | [optional] 
**Cnonce** | Pointer to **string** | Get the &#x60;c_nonce&#x60;.  | [optional] 
**CnonceExpiresAt** | Pointer to **int64** | Get the time at which the &#x60;c_nonce&#x60; expires in milliseconds since the Unix epoch (1970-01-01).  | [optional] 
**RequestedIdTokenClaims** | Pointer to **[]string** | Get the names of the claims that the authorization request (which resulted in generation of the access token) requested to be embedded in ID tokens.  | [optional] 
**RefreshTokenScopes** | Pointer to **[]string** | Scopes associated with the refresh token. | [optional] 

## Methods

### NewTokenResponse

`func NewTokenResponse() *TokenResponse`

NewTokenResponse instantiates a new TokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenResponseWithDefaults

`func NewTokenResponseWithDefaults() *TokenResponse`

NewTokenResponseWithDefaults instantiates a new TokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *TokenResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *TokenResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *TokenResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *TokenResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *TokenResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *TokenResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *TokenResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *TokenResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *TokenResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TokenResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TokenResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TokenResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResponseContent

`func (o *TokenResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *TokenResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *TokenResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *TokenResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetUsername

`func (o *TokenResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TokenResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TokenResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TokenResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *TokenResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TokenResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TokenResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TokenResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTicket

`func (o *TokenResponse) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *TokenResponse) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *TokenResponse) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *TokenResponse) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetAccessToken

`func (o *TokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *TokenResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAccessTokenExpiresAt

`func (o *TokenResponse) GetAccessTokenExpiresAt() int64`

GetAccessTokenExpiresAt returns the AccessTokenExpiresAt field if non-nil, zero value otherwise.

### GetAccessTokenExpiresAtOk

`func (o *TokenResponse) GetAccessTokenExpiresAtOk() (*int64, bool)`

GetAccessTokenExpiresAtOk returns a tuple with the AccessTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpiresAt

`func (o *TokenResponse) SetAccessTokenExpiresAt(v int64)`

SetAccessTokenExpiresAt sets AccessTokenExpiresAt field to given value.

### HasAccessTokenExpiresAt

`func (o *TokenResponse) HasAccessTokenExpiresAt() bool`

HasAccessTokenExpiresAt returns a boolean if a field has been set.

### GetAccessTokenDuration

`func (o *TokenResponse) GetAccessTokenDuration() int64`

GetAccessTokenDuration returns the AccessTokenDuration field if non-nil, zero value otherwise.

### GetAccessTokenDurationOk

`func (o *TokenResponse) GetAccessTokenDurationOk() (*int64, bool)`

GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenDuration

`func (o *TokenResponse) SetAccessTokenDuration(v int64)`

SetAccessTokenDuration sets AccessTokenDuration field to given value.

### HasAccessTokenDuration

`func (o *TokenResponse) HasAccessTokenDuration() bool`

HasAccessTokenDuration returns a boolean if a field has been set.

### GetRefreshToken

`func (o *TokenResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *TokenResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *TokenResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *TokenResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetRefreshTokenExpiresAt

`func (o *TokenResponse) GetRefreshTokenExpiresAt() int64`

GetRefreshTokenExpiresAt returns the RefreshTokenExpiresAt field if non-nil, zero value otherwise.

### GetRefreshTokenExpiresAtOk

`func (o *TokenResponse) GetRefreshTokenExpiresAtOk() (*int64, bool)`

GetRefreshTokenExpiresAtOk returns a tuple with the RefreshTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenExpiresAt

`func (o *TokenResponse) SetRefreshTokenExpiresAt(v int64)`

SetRefreshTokenExpiresAt sets RefreshTokenExpiresAt field to given value.

### HasRefreshTokenExpiresAt

`func (o *TokenResponse) HasRefreshTokenExpiresAt() bool`

HasRefreshTokenExpiresAt returns a boolean if a field has been set.

### GetRefreshTokenDuration

`func (o *TokenResponse) GetRefreshTokenDuration() int64`

GetRefreshTokenDuration returns the RefreshTokenDuration field if non-nil, zero value otherwise.

### GetRefreshTokenDurationOk

`func (o *TokenResponse) GetRefreshTokenDurationOk() (*int64, bool)`

GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDuration

`func (o *TokenResponse) SetRefreshTokenDuration(v int64)`

SetRefreshTokenDuration sets RefreshTokenDuration field to given value.

### HasRefreshTokenDuration

`func (o *TokenResponse) HasRefreshTokenDuration() bool`

HasRefreshTokenDuration returns a boolean if a field has been set.

### GetIdToken

`func (o *TokenResponse) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *TokenResponse) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *TokenResponse) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *TokenResponse) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.

### GetGrantType

`func (o *TokenResponse) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *TokenResponse) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *TokenResponse) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *TokenResponse) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetClientId

`func (o *TokenResponse) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TokenResponse) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TokenResponse) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TokenResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIdAlias

`func (o *TokenResponse) GetClientIdAlias() string`

GetClientIdAlias returns the ClientIdAlias field if non-nil, zero value otherwise.

### GetClientIdAliasOk

`func (o *TokenResponse) GetClientIdAliasOk() (*string, bool)`

GetClientIdAliasOk returns a tuple with the ClientIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAlias

`func (o *TokenResponse) SetClientIdAlias(v string)`

SetClientIdAlias sets ClientIdAlias field to given value.

### HasClientIdAlias

`func (o *TokenResponse) HasClientIdAlias() bool`

HasClientIdAlias returns a boolean if a field has been set.

### GetClientIdAliasUsed

`func (o *TokenResponse) GetClientIdAliasUsed() bool`

GetClientIdAliasUsed returns the ClientIdAliasUsed field if non-nil, zero value otherwise.

### GetClientIdAliasUsedOk

`func (o *TokenResponse) GetClientIdAliasUsedOk() (*bool, bool)`

GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasUsed

`func (o *TokenResponse) SetClientIdAliasUsed(v bool)`

SetClientIdAliasUsed sets ClientIdAliasUsed field to given value.

### HasClientIdAliasUsed

`func (o *TokenResponse) HasClientIdAliasUsed() bool`

HasClientIdAliasUsed returns a boolean if a field has been set.

### GetSubject

`func (o *TokenResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TokenResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TokenResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TokenResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetScopes

`func (o *TokenResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetProperties

`func (o *TokenResponse) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TokenResponse) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TokenResponse) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TokenResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetJwtAccessToken

`func (o *TokenResponse) GetJwtAccessToken() string`

GetJwtAccessToken returns the JwtAccessToken field if non-nil, zero value otherwise.

### GetJwtAccessTokenOk

`func (o *TokenResponse) GetJwtAccessTokenOk() (*string, bool)`

GetJwtAccessTokenOk returns a tuple with the JwtAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtAccessToken

`func (o *TokenResponse) SetJwtAccessToken(v string)`

SetJwtAccessToken sets JwtAccessToken field to given value.

### HasJwtAccessToken

`func (o *TokenResponse) HasJwtAccessToken() bool`

HasJwtAccessToken returns a boolean if a field has been set.

### GetResources

`func (o *TokenResponse) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *TokenResponse) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *TokenResponse) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *TokenResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetAccessTokenResources

`func (o *TokenResponse) GetAccessTokenResources() []string`

GetAccessTokenResources returns the AccessTokenResources field if non-nil, zero value otherwise.

### GetAccessTokenResourcesOk

`func (o *TokenResponse) GetAccessTokenResourcesOk() (*[]string, bool)`

GetAccessTokenResourcesOk returns a tuple with the AccessTokenResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenResources

`func (o *TokenResponse) SetAccessTokenResources(v []string)`

SetAccessTokenResources sets AccessTokenResources field to given value.

### HasAccessTokenResources

`func (o *TokenResponse) HasAccessTokenResources() bool`

HasAccessTokenResources returns a boolean if a field has been set.

### GetAuthorizationDetails

`func (o *TokenResponse) GetAuthorizationDetails() AuthzDetails`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *TokenResponse) GetAuthorizationDetailsOk() (*AuthzDetails, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *TokenResponse) SetAuthorizationDetails(v AuthzDetails)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *TokenResponse) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.

### GetServiceAttributes

`func (o *TokenResponse) GetServiceAttributes() []Pair`

GetServiceAttributes returns the ServiceAttributes field if non-nil, zero value otherwise.

### GetServiceAttributesOk

`func (o *TokenResponse) GetServiceAttributesOk() (*[]Pair, bool)`

GetServiceAttributesOk returns a tuple with the ServiceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAttributes

`func (o *TokenResponse) SetServiceAttributes(v []Pair)`

SetServiceAttributes sets ServiceAttributes field to given value.

### HasServiceAttributes

`func (o *TokenResponse) HasServiceAttributes() bool`

HasServiceAttributes returns a boolean if a field has been set.

### GetClientAttributes

`func (o *TokenResponse) GetClientAttributes() []Pair`

GetClientAttributes returns the ClientAttributes field if non-nil, zero value otherwise.

### GetClientAttributesOk

`func (o *TokenResponse) GetClientAttributesOk() (*[]Pair, bool)`

GetClientAttributesOk returns a tuple with the ClientAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAttributes

`func (o *TokenResponse) SetClientAttributes(v []Pair)`

SetClientAttributes sets ClientAttributes field to given value.

### HasClientAttributes

`func (o *TokenResponse) HasClientAttributes() bool`

HasClientAttributes returns a boolean if a field has been set.

### GetClientAuthMethod

`func (o *TokenResponse) GetClientAuthMethod() string`

GetClientAuthMethod returns the ClientAuthMethod field if non-nil, zero value otherwise.

### GetClientAuthMethodOk

`func (o *TokenResponse) GetClientAuthMethodOk() (*string, bool)`

GetClientAuthMethodOk returns a tuple with the ClientAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthMethod

`func (o *TokenResponse) SetClientAuthMethod(v string)`

SetClientAuthMethod sets ClientAuthMethod field to given value.

### HasClientAuthMethod

`func (o *TokenResponse) HasClientAuthMethod() bool`

HasClientAuthMethod returns a boolean if a field has been set.

### GetGrantId

`func (o *TokenResponse) GetGrantId() string`

GetGrantId returns the GrantId field if non-nil, zero value otherwise.

### GetGrantIdOk

`func (o *TokenResponse) GetGrantIdOk() (*string, bool)`

GetGrantIdOk returns a tuple with the GrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantId

`func (o *TokenResponse) SetGrantId(v string)`

SetGrantId sets GrantId field to given value.

### HasGrantId

`func (o *TokenResponse) HasGrantId() bool`

HasGrantId returns a boolean if a field has been set.

### GetAudiences

`func (o *TokenResponse) GetAudiences() []string`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *TokenResponse) GetAudiencesOk() (*[]string, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *TokenResponse) SetAudiences(v []string)`

SetAudiences sets Audiences field to given value.

### HasAudiences

`func (o *TokenResponse) HasAudiences() bool`

HasAudiences returns a boolean if a field has been set.

### GetRequestedTokenType

`func (o *TokenResponse) GetRequestedTokenType() TokenType`

GetRequestedTokenType returns the RequestedTokenType field if non-nil, zero value otherwise.

### GetRequestedTokenTypeOk

`func (o *TokenResponse) GetRequestedTokenTypeOk() (*TokenType, bool)`

GetRequestedTokenTypeOk returns a tuple with the RequestedTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTokenType

`func (o *TokenResponse) SetRequestedTokenType(v TokenType)`

SetRequestedTokenType sets RequestedTokenType field to given value.

### HasRequestedTokenType

`func (o *TokenResponse) HasRequestedTokenType() bool`

HasRequestedTokenType returns a boolean if a field has been set.

### GetSubjectToken

`func (o *TokenResponse) GetSubjectToken() string`

GetSubjectToken returns the SubjectToken field if non-nil, zero value otherwise.

### GetSubjectTokenOk

`func (o *TokenResponse) GetSubjectTokenOk() (*string, bool)`

GetSubjectTokenOk returns a tuple with the SubjectToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectToken

`func (o *TokenResponse) SetSubjectToken(v string)`

SetSubjectToken sets SubjectToken field to given value.

### HasSubjectToken

`func (o *TokenResponse) HasSubjectToken() bool`

HasSubjectToken returns a boolean if a field has been set.

### GetSubjectTokenType

`func (o *TokenResponse) GetSubjectTokenType() TokenType`

GetSubjectTokenType returns the SubjectTokenType field if non-nil, zero value otherwise.

### GetSubjectTokenTypeOk

`func (o *TokenResponse) GetSubjectTokenTypeOk() (*TokenType, bool)`

GetSubjectTokenTypeOk returns a tuple with the SubjectTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectTokenType

`func (o *TokenResponse) SetSubjectTokenType(v TokenType)`

SetSubjectTokenType sets SubjectTokenType field to given value.

### HasSubjectTokenType

`func (o *TokenResponse) HasSubjectTokenType() bool`

HasSubjectTokenType returns a boolean if a field has been set.

### GetSubjectTokenInfo

`func (o *TokenResponse) GetSubjectTokenInfo() TokenInfo`

GetSubjectTokenInfo returns the SubjectTokenInfo field if non-nil, zero value otherwise.

### GetSubjectTokenInfoOk

`func (o *TokenResponse) GetSubjectTokenInfoOk() (*TokenInfo, bool)`

GetSubjectTokenInfoOk returns a tuple with the SubjectTokenInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectTokenInfo

`func (o *TokenResponse) SetSubjectTokenInfo(v TokenInfo)`

SetSubjectTokenInfo sets SubjectTokenInfo field to given value.

### HasSubjectTokenInfo

`func (o *TokenResponse) HasSubjectTokenInfo() bool`

HasSubjectTokenInfo returns a boolean if a field has been set.

### GetActorToken

`func (o *TokenResponse) GetActorToken() string`

GetActorToken returns the ActorToken field if non-nil, zero value otherwise.

### GetActorTokenOk

`func (o *TokenResponse) GetActorTokenOk() (*string, bool)`

GetActorTokenOk returns a tuple with the ActorToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorToken

`func (o *TokenResponse) SetActorToken(v string)`

SetActorToken sets ActorToken field to given value.

### HasActorToken

`func (o *TokenResponse) HasActorToken() bool`

HasActorToken returns a boolean if a field has been set.

### GetActorTokenType

`func (o *TokenResponse) GetActorTokenType() TokenType`

GetActorTokenType returns the ActorTokenType field if non-nil, zero value otherwise.

### GetActorTokenTypeOk

`func (o *TokenResponse) GetActorTokenTypeOk() (*TokenType, bool)`

GetActorTokenTypeOk returns a tuple with the ActorTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorTokenType

`func (o *TokenResponse) SetActorTokenType(v TokenType)`

SetActorTokenType sets ActorTokenType field to given value.

### HasActorTokenType

`func (o *TokenResponse) HasActorTokenType() bool`

HasActorTokenType returns a boolean if a field has been set.

### GetActorTokenInfo

`func (o *TokenResponse) GetActorTokenInfo() TokenInfo`

GetActorTokenInfo returns the ActorTokenInfo field if non-nil, zero value otherwise.

### GetActorTokenInfoOk

`func (o *TokenResponse) GetActorTokenInfoOk() (*TokenInfo, bool)`

GetActorTokenInfoOk returns a tuple with the ActorTokenInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorTokenInfo

`func (o *TokenResponse) SetActorTokenInfo(v TokenInfo)`

SetActorTokenInfo sets ActorTokenInfo field to given value.

### HasActorTokenInfo

`func (o *TokenResponse) HasActorTokenInfo() bool`

HasActorTokenInfo returns a boolean if a field has been set.

### GetAssertion

`func (o *TokenResponse) GetAssertion() string`

GetAssertion returns the Assertion field if non-nil, zero value otherwise.

### GetAssertionOk

`func (o *TokenResponse) GetAssertionOk() (*string, bool)`

GetAssertionOk returns a tuple with the Assertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertion

`func (o *TokenResponse) SetAssertion(v string)`

SetAssertion sets Assertion field to given value.

### HasAssertion

`func (o *TokenResponse) HasAssertion() bool`

HasAssertion returns a boolean if a field has been set.

### GetPreviousRefreshTokenUsed

`func (o *TokenResponse) GetPreviousRefreshTokenUsed() bool`

GetPreviousRefreshTokenUsed returns the PreviousRefreshTokenUsed field if non-nil, zero value otherwise.

### GetPreviousRefreshTokenUsedOk

`func (o *TokenResponse) GetPreviousRefreshTokenUsedOk() (*bool, bool)`

GetPreviousRefreshTokenUsedOk returns a tuple with the PreviousRefreshTokenUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousRefreshTokenUsed

`func (o *TokenResponse) SetPreviousRefreshTokenUsed(v bool)`

SetPreviousRefreshTokenUsed sets PreviousRefreshTokenUsed field to given value.

### HasPreviousRefreshTokenUsed

`func (o *TokenResponse) HasPreviousRefreshTokenUsed() bool`

HasPreviousRefreshTokenUsed returns a boolean if a field has been set.

### GetClientEntityId

`func (o *TokenResponse) GetClientEntityId() string`

GetClientEntityId returns the ClientEntityId field if non-nil, zero value otherwise.

### GetClientEntityIdOk

`func (o *TokenResponse) GetClientEntityIdOk() (*string, bool)`

GetClientEntityIdOk returns a tuple with the ClientEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEntityId

`func (o *TokenResponse) SetClientEntityId(v string)`

SetClientEntityId sets ClientEntityId field to given value.

### HasClientEntityId

`func (o *TokenResponse) HasClientEntityId() bool`

HasClientEntityId returns a boolean if a field has been set.

### GetClientEntityIdUsed

`func (o *TokenResponse) GetClientEntityIdUsed() bool`

GetClientEntityIdUsed returns the ClientEntityIdUsed field if non-nil, zero value otherwise.

### GetClientEntityIdUsedOk

`func (o *TokenResponse) GetClientEntityIdUsedOk() (*bool, bool)`

GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEntityIdUsed

`func (o *TokenResponse) SetClientEntityIdUsed(v bool)`

SetClientEntityIdUsed sets ClientEntityIdUsed field to given value.

### HasClientEntityIdUsed

`func (o *TokenResponse) HasClientEntityIdUsed() bool`

HasClientEntityIdUsed returns a boolean if a field has been set.

### GetCnonceDuration

`func (o *TokenResponse) GetCnonceDuration() int64`

GetCnonceDuration returns the CnonceDuration field if non-nil, zero value otherwise.

### GetCnonceDurationOk

`func (o *TokenResponse) GetCnonceDurationOk() (*int64, bool)`

GetCnonceDurationOk returns a tuple with the CnonceDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnonceDuration

`func (o *TokenResponse) SetCnonceDuration(v int64)`

SetCnonceDuration sets CnonceDuration field to given value.

### HasCnonceDuration

`func (o *TokenResponse) HasCnonceDuration() bool`

HasCnonceDuration returns a boolean if a field has been set.

### GetDpopNonce

`func (o *TokenResponse) GetDpopNonce() string`

GetDpopNonce returns the DpopNonce field if non-nil, zero value otherwise.

### GetDpopNonceOk

`func (o *TokenResponse) GetDpopNonceOk() (*string, bool)`

GetDpopNonceOk returns a tuple with the DpopNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpopNonce

`func (o *TokenResponse) SetDpopNonce(v string)`

SetDpopNonce sets DpopNonce field to given value.

### HasDpopNonce

`func (o *TokenResponse) HasDpopNonce() bool`

HasDpopNonce returns a boolean if a field has been set.

### GetCnonce

`func (o *TokenResponse) GetCnonce() string`

GetCnonce returns the Cnonce field if non-nil, zero value otherwise.

### GetCnonceOk

`func (o *TokenResponse) GetCnonceOk() (*string, bool)`

GetCnonceOk returns a tuple with the Cnonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnonce

`func (o *TokenResponse) SetCnonce(v string)`

SetCnonce sets Cnonce field to given value.

### HasCnonce

`func (o *TokenResponse) HasCnonce() bool`

HasCnonce returns a boolean if a field has been set.

### GetCnonceExpiresAt

`func (o *TokenResponse) GetCnonceExpiresAt() int64`

GetCnonceExpiresAt returns the CnonceExpiresAt field if non-nil, zero value otherwise.

### GetCnonceExpiresAtOk

`func (o *TokenResponse) GetCnonceExpiresAtOk() (*int64, bool)`

GetCnonceExpiresAtOk returns a tuple with the CnonceExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnonceExpiresAt

`func (o *TokenResponse) SetCnonceExpiresAt(v int64)`

SetCnonceExpiresAt sets CnonceExpiresAt field to given value.

### HasCnonceExpiresAt

`func (o *TokenResponse) HasCnonceExpiresAt() bool`

HasCnonceExpiresAt returns a boolean if a field has been set.

### GetRequestedIdTokenClaims

`func (o *TokenResponse) GetRequestedIdTokenClaims() []string`

GetRequestedIdTokenClaims returns the RequestedIdTokenClaims field if non-nil, zero value otherwise.

### GetRequestedIdTokenClaimsOk

`func (o *TokenResponse) GetRequestedIdTokenClaimsOk() (*[]string, bool)`

GetRequestedIdTokenClaimsOk returns a tuple with the RequestedIdTokenClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedIdTokenClaims

`func (o *TokenResponse) SetRequestedIdTokenClaims(v []string)`

SetRequestedIdTokenClaims sets RequestedIdTokenClaims field to given value.

### HasRequestedIdTokenClaims

`func (o *TokenResponse) HasRequestedIdTokenClaims() bool`

HasRequestedIdTokenClaims returns a boolean if a field has been set.

### GetRefreshTokenScopes

`func (o *TokenResponse) GetRefreshTokenScopes() []string`

GetRefreshTokenScopes returns the RefreshTokenScopes field if non-nil, zero value otherwise.

### GetRefreshTokenScopesOk

`func (o *TokenResponse) GetRefreshTokenScopesOk() (*[]string, bool)`

GetRefreshTokenScopesOk returns a tuple with the RefreshTokenScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenScopes

`func (o *TokenResponse) SetRefreshTokenScopes(v []string)`

SetRefreshTokenScopes sets RefreshTokenScopes field to given value.

### HasRefreshTokenScopes

`func (o *TokenResponse) HasRefreshTokenScopes() bool`

HasRefreshTokenScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


