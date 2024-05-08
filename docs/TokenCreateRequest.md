# TokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | [**GrantType**](GrantType.md) |  | 
**ClientId** | **int64** | The ID of the client application which will be associated with a newly created access token.  | 
**Subject** | Pointer to **string** | The subject (&#x3D; unique identifier) of the user who will be associated with a newly created access token. This parameter is required unless the grant type is &#x60;CLIENT_CREDENTIALS&#x60;. The value must consist of only ASCII characters and its length must not exceed 100.  | [optional] 
**Scopes** | Pointer to **[]string** | The scopes which will be associated with a newly created access token. Scopes that are not supported by the service cannot be specified and requesting them will cause an error.  | [optional] 
**AccessTokenDuration** | Pointer to **int64** | The duration of a newly created access token in seconds. If the value is 0, the duration is determined according to the settings of the service.  | [optional] 
**RefreshTokenDuration** | Pointer to **int64** | The duration of a newly created refresh token in seconds. If the value is 0, the duration is determined according to the settings of the service.  A refresh token is not created (1) if the service does not support &#x60;REFRESH_TOKEN&#x60;, or (2) if the specified grant type is either &#x60;IMPLICIT&#x60;or &#x60;CLIENT_CREDENTIALS&#x60;.  | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) | Extra properties to associate with a newly created access token. Note that properties parameter is accepted only when the HTTP method of the request is POST and Content-Type of the request is &#x60;application/json&#x60;, so don&#39;t use &#x60;GET&#x60; method or &#x60;application/x-www-form-urlencoded&#x60; if you want to specify properties.  | [optional] 
**ClientIdAliasUsed** | Pointer to **bool** | A boolean request parameter which indicates whether to emulate that the client ID alias is used instead of the original numeric client ID when a new access token is created.  This has an effect only on the value of the aud claim in a response from [UserInfo endpoint](https://openid.net/specs/openid-connect-core-1_0.html#UserInfo). When you access the UserInfo endpoint (which is expected to be implemented using Authlete&#39;s &#x60;/api/auth/userinfo&#x60; API and &#x60;/api/auth/userinfo/issue&#x60; API) with an access token which has been created using Authlete&#39;s &#x60;/api/auth/token/create&#x60; API with this property (&#x60;clientIdAliasUsed&#x60;) &#x60;true&#x60;, the client ID alias is used as the value of the aud claim in a response from the UserInfo endpoint.  Note that if a client ID alias is not assigned to the client when Authlete&#39;s &#x60;/api/auth/token/create&#x60; API is called, this property (&#x60;clientIdAliasUsed&#x60;) has no effect (it is always regarded as &#x60;false&#x60;).  | [optional] 
**AccessToken** | Pointer to **string** | The value of the new access token.  The &#x60;/api/auth/token/create&#x60; API generates an access token. Therefore, callers of the API do not have to specify values of newly created access tokens. However, in some cases, for example, if you want to migrate existing access tokens from an old system to Authlete, you may want to specify values of access tokens. In such a case, you can specify the value of a newly created access token by passing a non-null value as the value of accessToken request parameter. The implementation of the &#x60;/api/auth/token/create&#x60; uses the value of the accessToken request parameter instead of generating a new value when the request parameter holds a non-null value.  Note that if the hash value of the specified access token already exists in Authlete&#39;s database, the access token cannot be inserted and the &#x60;/api/auth/token/create&#x60; API will report an error.  | [optional] 
**RefreshToken** | Pointer to **string** | The value of the new refresh token.  The &#x60;/api/auth/token/create&#x60; API may generate a refresh token. Therefore, callers of the API do not have to specify values of newly created refresh tokens. However, in some cases, for example, if you want to migrate existing refresh tokens from an old system to Authlete, you may want to specify values of refresh tokens. In such a case, you can specify the value of a newly created refresh token by passing a non-null value as the value of refreshToken request parameter. The implementation of the &#x60;/api/auth/token/create&#x60; uses the value of the refreshToken request parameter instead of generating a new value when the request parameter holds a non-null value.  Note that if the hash value of the specified refresh token already exists in Authlete&#39;s database, the refresh token cannot be inserted and the &#x60;/api/auth/token/create&#x60; API will report an error.  | [optional] 
**AccessTokenPersistent** | Pointer to **bool** | Get whether the access token expires or not. By default, all access tokens expire after a period of time determined by their service.  If this request parameter is &#x60;true&#x60;, then the access token will not automatically expire and must be revoked or deleted manually at the service. If this request parameter is true, the &#x60;accessTokenDuration&#x60; request parameter is ignored.  | [optional] 
**CertificateThumbprint** | Pointer to **string** | The thumbprint of the MTLS certificate bound to this token. If this property is set, a certificate with the corresponding value MUST be presented with the access token when it is used by a client. The value of this property must be a SHA256 certificate thumbprint, base64url encoded.  | [optional] 
**DpopKeyThumbprint** | Pointer to **string** | The thumbprint of the public key used for DPoP presentation of this token. If this property is set, a DPoP proof signed with the corresponding private key MUST be presented with the access token when it is used by a client. Additionally, the token&#39;s &#x60;token_type&#x60; will be set to &#39;DPoP&#39;.  | [optional] 
**AuthorizationDetails** | Pointer to [**AuthzDetails**](AuthzDetails.md) |  | [optional] 
**Resources** | Pointer to **[]string** | The value of the resources to associate with the token. This property represents the value of one or more &#x60;resource&#x60; request parameters which is defined in \&quot;RFC8707 Resource Indicators for OAuth 2.0\&quot;.  | [optional] 
**ForExternalAttachment** | Pointer to **bool** | the flag which indicates whether the access token is for an external attachment.  | [optional] 
**JwtAtClaims** | Pointer to **string** | Additional claims that are added to the payload part of the JWT access token.  | [optional] 
**Acr** | Pointer to **string** | The Authentication Context Class Reference of the user authentication that the authorization server performed  during the course of issuing the access token.  | [optional] 
**AuthTime** | Pointer to **int64** | The time when the user authentication was performed during the course of issuing the access token.  | [optional] 
**ClientEntityIdUsed** | Pointer to **bool** | Flag which indicates whether the entity ID of the client was used when the request for the access token was made. | [optional] 

## Methods

### NewTokenCreateRequest

`func NewTokenCreateRequest(grantType GrantType, clientId int64, ) *TokenCreateRequest`

NewTokenCreateRequest instantiates a new TokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCreateRequestWithDefaults

`func NewTokenCreateRequestWithDefaults() *TokenCreateRequest`

NewTokenCreateRequestWithDefaults instantiates a new TokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *TokenCreateRequest) GetGrantType() GrantType`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *TokenCreateRequest) GetGrantTypeOk() (*GrantType, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *TokenCreateRequest) SetGrantType(v GrantType)`

SetGrantType sets GrantType field to given value.


### GetClientId

`func (o *TokenCreateRequest) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TokenCreateRequest) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TokenCreateRequest) SetClientId(v int64)`

SetClientId sets ClientId field to given value.


### GetSubject

`func (o *TokenCreateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TokenCreateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TokenCreateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TokenCreateRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetScopes

`func (o *TokenCreateRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenCreateRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenCreateRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenCreateRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetAccessTokenDuration

`func (o *TokenCreateRequest) GetAccessTokenDuration() int64`

GetAccessTokenDuration returns the AccessTokenDuration field if non-nil, zero value otherwise.

### GetAccessTokenDurationOk

`func (o *TokenCreateRequest) GetAccessTokenDurationOk() (*int64, bool)`

GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenDuration

`func (o *TokenCreateRequest) SetAccessTokenDuration(v int64)`

SetAccessTokenDuration sets AccessTokenDuration field to given value.

### HasAccessTokenDuration

`func (o *TokenCreateRequest) HasAccessTokenDuration() bool`

HasAccessTokenDuration returns a boolean if a field has been set.

### GetRefreshTokenDuration

`func (o *TokenCreateRequest) GetRefreshTokenDuration() int64`

GetRefreshTokenDuration returns the RefreshTokenDuration field if non-nil, zero value otherwise.

### GetRefreshTokenDurationOk

`func (o *TokenCreateRequest) GetRefreshTokenDurationOk() (*int64, bool)`

GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDuration

`func (o *TokenCreateRequest) SetRefreshTokenDuration(v int64)`

SetRefreshTokenDuration sets RefreshTokenDuration field to given value.

### HasRefreshTokenDuration

`func (o *TokenCreateRequest) HasRefreshTokenDuration() bool`

HasRefreshTokenDuration returns a boolean if a field has been set.

### GetProperties

`func (o *TokenCreateRequest) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TokenCreateRequest) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TokenCreateRequest) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TokenCreateRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetClientIdAliasUsed

`func (o *TokenCreateRequest) GetClientIdAliasUsed() bool`

GetClientIdAliasUsed returns the ClientIdAliasUsed field if non-nil, zero value otherwise.

### GetClientIdAliasUsedOk

`func (o *TokenCreateRequest) GetClientIdAliasUsedOk() (*bool, bool)`

GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasUsed

`func (o *TokenCreateRequest) SetClientIdAliasUsed(v bool)`

SetClientIdAliasUsed sets ClientIdAliasUsed field to given value.

### HasClientIdAliasUsed

`func (o *TokenCreateRequest) HasClientIdAliasUsed() bool`

HasClientIdAliasUsed returns a boolean if a field has been set.

### GetAccessToken

`func (o *TokenCreateRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TokenCreateRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TokenCreateRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *TokenCreateRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *TokenCreateRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *TokenCreateRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *TokenCreateRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *TokenCreateRequest) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetAccessTokenPersistent

`func (o *TokenCreateRequest) GetAccessTokenPersistent() bool`

GetAccessTokenPersistent returns the AccessTokenPersistent field if non-nil, zero value otherwise.

### GetAccessTokenPersistentOk

`func (o *TokenCreateRequest) GetAccessTokenPersistentOk() (*bool, bool)`

GetAccessTokenPersistentOk returns a tuple with the AccessTokenPersistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenPersistent

`func (o *TokenCreateRequest) SetAccessTokenPersistent(v bool)`

SetAccessTokenPersistent sets AccessTokenPersistent field to given value.

### HasAccessTokenPersistent

`func (o *TokenCreateRequest) HasAccessTokenPersistent() bool`

HasAccessTokenPersistent returns a boolean if a field has been set.

### GetCertificateThumbprint

`func (o *TokenCreateRequest) GetCertificateThumbprint() string`

GetCertificateThumbprint returns the CertificateThumbprint field if non-nil, zero value otherwise.

### GetCertificateThumbprintOk

`func (o *TokenCreateRequest) GetCertificateThumbprintOk() (*string, bool)`

GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateThumbprint

`func (o *TokenCreateRequest) SetCertificateThumbprint(v string)`

SetCertificateThumbprint sets CertificateThumbprint field to given value.

### HasCertificateThumbprint

`func (o *TokenCreateRequest) HasCertificateThumbprint() bool`

HasCertificateThumbprint returns a boolean if a field has been set.

### GetDpopKeyThumbprint

`func (o *TokenCreateRequest) GetDpopKeyThumbprint() string`

GetDpopKeyThumbprint returns the DpopKeyThumbprint field if non-nil, zero value otherwise.

### GetDpopKeyThumbprintOk

`func (o *TokenCreateRequest) GetDpopKeyThumbprintOk() (*string, bool)`

GetDpopKeyThumbprintOk returns a tuple with the DpopKeyThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpopKeyThumbprint

`func (o *TokenCreateRequest) SetDpopKeyThumbprint(v string)`

SetDpopKeyThumbprint sets DpopKeyThumbprint field to given value.

### HasDpopKeyThumbprint

`func (o *TokenCreateRequest) HasDpopKeyThumbprint() bool`

HasDpopKeyThumbprint returns a boolean if a field has been set.

### GetAuthorizationDetails

`func (o *TokenCreateRequest) GetAuthorizationDetails() AuthzDetails`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *TokenCreateRequest) GetAuthorizationDetailsOk() (*AuthzDetails, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *TokenCreateRequest) SetAuthorizationDetails(v AuthzDetails)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *TokenCreateRequest) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.

### GetResources

`func (o *TokenCreateRequest) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *TokenCreateRequest) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *TokenCreateRequest) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *TokenCreateRequest) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetForExternalAttachment

`func (o *TokenCreateRequest) GetForExternalAttachment() bool`

GetForExternalAttachment returns the ForExternalAttachment field if non-nil, zero value otherwise.

### GetForExternalAttachmentOk

`func (o *TokenCreateRequest) GetForExternalAttachmentOk() (*bool, bool)`

GetForExternalAttachmentOk returns a tuple with the ForExternalAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForExternalAttachment

`func (o *TokenCreateRequest) SetForExternalAttachment(v bool)`

SetForExternalAttachment sets ForExternalAttachment field to given value.

### HasForExternalAttachment

`func (o *TokenCreateRequest) HasForExternalAttachment() bool`

HasForExternalAttachment returns a boolean if a field has been set.

### GetJwtAtClaims

`func (o *TokenCreateRequest) GetJwtAtClaims() string`

GetJwtAtClaims returns the JwtAtClaims field if non-nil, zero value otherwise.

### GetJwtAtClaimsOk

`func (o *TokenCreateRequest) GetJwtAtClaimsOk() (*string, bool)`

GetJwtAtClaimsOk returns a tuple with the JwtAtClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtAtClaims

`func (o *TokenCreateRequest) SetJwtAtClaims(v string)`

SetJwtAtClaims sets JwtAtClaims field to given value.

### HasJwtAtClaims

`func (o *TokenCreateRequest) HasJwtAtClaims() bool`

HasJwtAtClaims returns a boolean if a field has been set.

### GetAcr

`func (o *TokenCreateRequest) GetAcr() string`

GetAcr returns the Acr field if non-nil, zero value otherwise.

### GetAcrOk

`func (o *TokenCreateRequest) GetAcrOk() (*string, bool)`

GetAcrOk returns a tuple with the Acr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcr

`func (o *TokenCreateRequest) SetAcr(v string)`

SetAcr sets Acr field to given value.

### HasAcr

`func (o *TokenCreateRequest) HasAcr() bool`

HasAcr returns a boolean if a field has been set.

### GetAuthTime

`func (o *TokenCreateRequest) GetAuthTime() int64`

GetAuthTime returns the AuthTime field if non-nil, zero value otherwise.

### GetAuthTimeOk

`func (o *TokenCreateRequest) GetAuthTimeOk() (*int64, bool)`

GetAuthTimeOk returns a tuple with the AuthTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTime

`func (o *TokenCreateRequest) SetAuthTime(v int64)`

SetAuthTime sets AuthTime field to given value.

### HasAuthTime

`func (o *TokenCreateRequest) HasAuthTime() bool`

HasAuthTime returns a boolean if a field has been set.

### GetClientEntityIdUsed

`func (o *TokenCreateRequest) GetClientEntityIdUsed() bool`

GetClientEntityIdUsed returns the ClientEntityIdUsed field if non-nil, zero value otherwise.

### GetClientEntityIdUsedOk

`func (o *TokenCreateRequest) GetClientEntityIdUsedOk() (*bool, bool)`

GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEntityIdUsed

`func (o *TokenCreateRequest) SetClientEntityIdUsed(v bool)`

SetClientEntityIdUsed sets ClientEntityIdUsed field to given value.

### HasClientEntityIdUsed

`func (o *TokenCreateRequest) HasClientEntityIdUsed() bool`

HasClientEntityIdUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


