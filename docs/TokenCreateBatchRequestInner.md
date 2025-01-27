# TokenCreateBatchRequestInner

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
**ClientIdAliasUsed** | Pointer to **bool** | A boolean request parameter which indicates whether to emulate that the client ID alias is used instead of the original numeric client ID when a new access token is created.  This has an effect only on the value of the aud claim in a response from [UserInfo endpoint](https://openid.net/specs/openid-connect-core-1_0.html#UserInfo). When you access the UserInfo endpoint (which is expected to be implemented using Authlete&#39;s &#x60;/api/auth/userinfo&#x60; API and &#x60;/api/auth/userinfo/issue&#x60; API) with an access token which has been created using Authlete&#39;s &#x60;/api/auth/token/create/batch&#x60; API with this property (&#x60;clientIdAliasUsed&#x60;) &#x60;true&#x60;, the client ID alias is used as the value of the aud claim in a response from the UserInfo endpoint.  Note that if a client ID alias is not assigned to the client when Authlete&#39;s &#x60;/api/auth/token/create/batch&#x60; API is called, this property (&#x60;clientIdAliasUsed&#x60;) has no effect (it is always regarded as &#x60;false&#x60;).  | [optional] 
**AccessToken** | **string** | The value of the new access token. you must specify the value of a newly created access token by  passing a non-null value as the value of accessToken request parameter.  | 
**RefreshToken** | Pointer to **string** | The value of the new refresh token. This is required under some conditions such as when the service  supports the refresh token flow.  | [optional] 
**AccessTokenPersistent** | Pointer to **bool** | Get whether the access token expires or not. By default, all access tokens expire after a period of time determined by their service.  If this request parameter is &#x60;true&#x60;, then the access token will not automatically expire and must be revoked or deleted manually at the service. If this request parameter is true, the &#x60;accessTokenDuration&#x60; request parameter is ignored.  | [optional] 
**CertificateThumbprint** | Pointer to **string** | The thumbprint of the MTLS certificate bound to this token. If this property is set, a certificate with the corresponding value MUST be presented with the access token when it is used by a client. The value of this property must be a SHA256 certificate thumbprint, base64url encoded.  | [optional] 
**DpopKeyThumbprint** | Pointer to **string** | The thumbprint of the public key used for DPoP presentation of this token. If this property is set, a DPoP proof signed with the corresponding private key MUST be presented with the access token when it is used by a client. Additionally, the token&#39;s &#x60;token_type&#x60; will be set to &#39;DPoP&#39;.  | [optional] 
**AuthorizationDetails** | Pointer to [**AuthzDetails**](AuthzDetails.md) |  | [optional] 
**Resources** | Pointer to **[]string** | The value of the resources to associate with the token. This property represents the value of one or more &#x60;resource&#x60; request parameters which is defined in \&quot;RFC8707 Resource Indicators for OAuth 2.0\&quot;.  | [optional] 
**ForExternalAttachment** | Pointer to **bool** | the flag which indicates whether the access token is for an external attachment.  | [optional] 
**JwtAtClaims** | Pointer to **string** | Additional claims that are added to the payload part of the JWT access token.  | [optional] 
**Acr** | Pointer to **string** | The Authentication Context Class Reference of the user authentication that the authorization server performed  during the course of issuing the access token.  | [optional] 
**AuthTime** | Pointer to **int64** | The time when the user authentication was performed during the course of issuing the access token.  | [optional] 
**ClientEntityIdUsed** | Pointer to **bool** | Flag which indicates whether the entity ID of the client was used when the request for the access token was made.  | [optional] 

## Methods

### NewTokenCreateBatchRequestInner

`func NewTokenCreateBatchRequestInner(grantType GrantType, clientId int64, accessToken string, ) *TokenCreateBatchRequestInner`

NewTokenCreateBatchRequestInner instantiates a new TokenCreateBatchRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCreateBatchRequestInnerWithDefaults

`func NewTokenCreateBatchRequestInnerWithDefaults() *TokenCreateBatchRequestInner`

NewTokenCreateBatchRequestInnerWithDefaults instantiates a new TokenCreateBatchRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *TokenCreateBatchRequestInner) GetGrantType() GrantType`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *TokenCreateBatchRequestInner) GetGrantTypeOk() (*GrantType, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *TokenCreateBatchRequestInner) SetGrantType(v GrantType)`

SetGrantType sets GrantType field to given value.


### GetClientId

`func (o *TokenCreateBatchRequestInner) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TokenCreateBatchRequestInner) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TokenCreateBatchRequestInner) SetClientId(v int64)`

SetClientId sets ClientId field to given value.


### GetSubject

`func (o *TokenCreateBatchRequestInner) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TokenCreateBatchRequestInner) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TokenCreateBatchRequestInner) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TokenCreateBatchRequestInner) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetScopes

`func (o *TokenCreateBatchRequestInner) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenCreateBatchRequestInner) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenCreateBatchRequestInner) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenCreateBatchRequestInner) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetAccessTokenDuration

`func (o *TokenCreateBatchRequestInner) GetAccessTokenDuration() int64`

GetAccessTokenDuration returns the AccessTokenDuration field if non-nil, zero value otherwise.

### GetAccessTokenDurationOk

`func (o *TokenCreateBatchRequestInner) GetAccessTokenDurationOk() (*int64, bool)`

GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenDuration

`func (o *TokenCreateBatchRequestInner) SetAccessTokenDuration(v int64)`

SetAccessTokenDuration sets AccessTokenDuration field to given value.

### HasAccessTokenDuration

`func (o *TokenCreateBatchRequestInner) HasAccessTokenDuration() bool`

HasAccessTokenDuration returns a boolean if a field has been set.

### GetRefreshTokenDuration

`func (o *TokenCreateBatchRequestInner) GetRefreshTokenDuration() int64`

GetRefreshTokenDuration returns the RefreshTokenDuration field if non-nil, zero value otherwise.

### GetRefreshTokenDurationOk

`func (o *TokenCreateBatchRequestInner) GetRefreshTokenDurationOk() (*int64, bool)`

GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDuration

`func (o *TokenCreateBatchRequestInner) SetRefreshTokenDuration(v int64)`

SetRefreshTokenDuration sets RefreshTokenDuration field to given value.

### HasRefreshTokenDuration

`func (o *TokenCreateBatchRequestInner) HasRefreshTokenDuration() bool`

HasRefreshTokenDuration returns a boolean if a field has been set.

### GetProperties

`func (o *TokenCreateBatchRequestInner) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TokenCreateBatchRequestInner) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TokenCreateBatchRequestInner) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TokenCreateBatchRequestInner) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetClientIdAliasUsed

`func (o *TokenCreateBatchRequestInner) GetClientIdAliasUsed() bool`

GetClientIdAliasUsed returns the ClientIdAliasUsed field if non-nil, zero value otherwise.

### GetClientIdAliasUsedOk

`func (o *TokenCreateBatchRequestInner) GetClientIdAliasUsedOk() (*bool, bool)`

GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasUsed

`func (o *TokenCreateBatchRequestInner) SetClientIdAliasUsed(v bool)`

SetClientIdAliasUsed sets ClientIdAliasUsed field to given value.

### HasClientIdAliasUsed

`func (o *TokenCreateBatchRequestInner) HasClientIdAliasUsed() bool`

HasClientIdAliasUsed returns a boolean if a field has been set.

### GetAccessToken

`func (o *TokenCreateBatchRequestInner) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TokenCreateBatchRequestInner) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TokenCreateBatchRequestInner) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetRefreshToken

`func (o *TokenCreateBatchRequestInner) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *TokenCreateBatchRequestInner) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *TokenCreateBatchRequestInner) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *TokenCreateBatchRequestInner) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetAccessTokenPersistent

`func (o *TokenCreateBatchRequestInner) GetAccessTokenPersistent() bool`

GetAccessTokenPersistent returns the AccessTokenPersistent field if non-nil, zero value otherwise.

### GetAccessTokenPersistentOk

`func (o *TokenCreateBatchRequestInner) GetAccessTokenPersistentOk() (*bool, bool)`

GetAccessTokenPersistentOk returns a tuple with the AccessTokenPersistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenPersistent

`func (o *TokenCreateBatchRequestInner) SetAccessTokenPersistent(v bool)`

SetAccessTokenPersistent sets AccessTokenPersistent field to given value.

### HasAccessTokenPersistent

`func (o *TokenCreateBatchRequestInner) HasAccessTokenPersistent() bool`

HasAccessTokenPersistent returns a boolean if a field has been set.

### GetCertificateThumbprint

`func (o *TokenCreateBatchRequestInner) GetCertificateThumbprint() string`

GetCertificateThumbprint returns the CertificateThumbprint field if non-nil, zero value otherwise.

### GetCertificateThumbprintOk

`func (o *TokenCreateBatchRequestInner) GetCertificateThumbprintOk() (*string, bool)`

GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateThumbprint

`func (o *TokenCreateBatchRequestInner) SetCertificateThumbprint(v string)`

SetCertificateThumbprint sets CertificateThumbprint field to given value.

### HasCertificateThumbprint

`func (o *TokenCreateBatchRequestInner) HasCertificateThumbprint() bool`

HasCertificateThumbprint returns a boolean if a field has been set.

### GetDpopKeyThumbprint

`func (o *TokenCreateBatchRequestInner) GetDpopKeyThumbprint() string`

GetDpopKeyThumbprint returns the DpopKeyThumbprint field if non-nil, zero value otherwise.

### GetDpopKeyThumbprintOk

`func (o *TokenCreateBatchRequestInner) GetDpopKeyThumbprintOk() (*string, bool)`

GetDpopKeyThumbprintOk returns a tuple with the DpopKeyThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpopKeyThumbprint

`func (o *TokenCreateBatchRequestInner) SetDpopKeyThumbprint(v string)`

SetDpopKeyThumbprint sets DpopKeyThumbprint field to given value.

### HasDpopKeyThumbprint

`func (o *TokenCreateBatchRequestInner) HasDpopKeyThumbprint() bool`

HasDpopKeyThumbprint returns a boolean if a field has been set.

### GetAuthorizationDetails

`func (o *TokenCreateBatchRequestInner) GetAuthorizationDetails() AuthzDetails`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *TokenCreateBatchRequestInner) GetAuthorizationDetailsOk() (*AuthzDetails, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *TokenCreateBatchRequestInner) SetAuthorizationDetails(v AuthzDetails)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *TokenCreateBatchRequestInner) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.

### GetResources

`func (o *TokenCreateBatchRequestInner) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *TokenCreateBatchRequestInner) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *TokenCreateBatchRequestInner) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *TokenCreateBatchRequestInner) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetForExternalAttachment

`func (o *TokenCreateBatchRequestInner) GetForExternalAttachment() bool`

GetForExternalAttachment returns the ForExternalAttachment field if non-nil, zero value otherwise.

### GetForExternalAttachmentOk

`func (o *TokenCreateBatchRequestInner) GetForExternalAttachmentOk() (*bool, bool)`

GetForExternalAttachmentOk returns a tuple with the ForExternalAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForExternalAttachment

`func (o *TokenCreateBatchRequestInner) SetForExternalAttachment(v bool)`

SetForExternalAttachment sets ForExternalAttachment field to given value.

### HasForExternalAttachment

`func (o *TokenCreateBatchRequestInner) HasForExternalAttachment() bool`

HasForExternalAttachment returns a boolean if a field has been set.

### GetJwtAtClaims

`func (o *TokenCreateBatchRequestInner) GetJwtAtClaims() string`

GetJwtAtClaims returns the JwtAtClaims field if non-nil, zero value otherwise.

### GetJwtAtClaimsOk

`func (o *TokenCreateBatchRequestInner) GetJwtAtClaimsOk() (*string, bool)`

GetJwtAtClaimsOk returns a tuple with the JwtAtClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtAtClaims

`func (o *TokenCreateBatchRequestInner) SetJwtAtClaims(v string)`

SetJwtAtClaims sets JwtAtClaims field to given value.

### HasJwtAtClaims

`func (o *TokenCreateBatchRequestInner) HasJwtAtClaims() bool`

HasJwtAtClaims returns a boolean if a field has been set.

### GetAcr

`func (o *TokenCreateBatchRequestInner) GetAcr() string`

GetAcr returns the Acr field if non-nil, zero value otherwise.

### GetAcrOk

`func (o *TokenCreateBatchRequestInner) GetAcrOk() (*string, bool)`

GetAcrOk returns a tuple with the Acr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcr

`func (o *TokenCreateBatchRequestInner) SetAcr(v string)`

SetAcr sets Acr field to given value.

### HasAcr

`func (o *TokenCreateBatchRequestInner) HasAcr() bool`

HasAcr returns a boolean if a field has been set.

### GetAuthTime

`func (o *TokenCreateBatchRequestInner) GetAuthTime() int64`

GetAuthTime returns the AuthTime field if non-nil, zero value otherwise.

### GetAuthTimeOk

`func (o *TokenCreateBatchRequestInner) GetAuthTimeOk() (*int64, bool)`

GetAuthTimeOk returns a tuple with the AuthTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTime

`func (o *TokenCreateBatchRequestInner) SetAuthTime(v int64)`

SetAuthTime sets AuthTime field to given value.

### HasAuthTime

`func (o *TokenCreateBatchRequestInner) HasAuthTime() bool`

HasAuthTime returns a boolean if a field has been set.

### GetClientEntityIdUsed

`func (o *TokenCreateBatchRequestInner) GetClientEntityIdUsed() bool`

GetClientEntityIdUsed returns the ClientEntityIdUsed field if non-nil, zero value otherwise.

### GetClientEntityIdUsedOk

`func (o *TokenCreateBatchRequestInner) GetClientEntityIdUsedOk() (*bool, bool)`

GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEntityIdUsed

`func (o *TokenCreateBatchRequestInner) SetClientEntityIdUsed(v bool)`

SetClientEntityIdUsed sets ClientEntityIdUsed field to given value.

### HasClientEntityIdUsed

`func (o *TokenCreateBatchRequestInner) HasClientEntityIdUsed() bool`

HasClientEntityIdUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


