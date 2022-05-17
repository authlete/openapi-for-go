# TokenIssueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take. | [optional] 
**ResponseContent** | Pointer to **string** | The content that the authorization server implementation is to return to the client application. Its format is JSON.  | [optional] 
**Username** | Pointer to **string** | The value of &#x60;username&#x60; request parameter in the token request. The client application must specify username when it uses [Resource Owner Password Grant](https://datatracker.ietf.org/doc/html/rfc6749#section-4.3). In other words, when the value of &#x60;grant_type&#x60; request parameter is &#x60;password&#x60;, &#x60;username&#x60; request parameter must come along.  This parameter has a value only if the value of &#x60;grant_type&#x60; request parameter is &#x60;password&#x60; and the token request is valid.  | [optional] 
**AccessToken** | Pointer to **string** | The newly issued access token. This parameter is a non-null value only when the value of &#x60;action&#x60; parameter is &#x60;OK&#x60;. | [optional] 
**AccessTokenExpiresAt** | Pointer to **int64** | The datetime at which the newly issued access token will expire. The value is represented in milliseconds since the Unix epoch (1970-01-01).  | [optional] 
**AccessTokenDuration** | Pointer to **int64** | The duration of the newly issued access token in seconds. | [optional] 
**RefreshToken** | Pointer to **string** | The refresh token. This parameter is a non-null value only when &#x60;action&#x60; is &#x60;OK&#x60; and the service supports the refresh token flow. If &#x60;refreshTokenKept&#x60; is set to &#x60;false&#x60;, a new refresh token is issued and the old refresh token used in the refresh token flow is invalidated. On the contrary, if &#x60;refreshTokenKept&#x60; is set to &#x60;true&#x60;, the refresh token itself is not refreshed.  | [optional] 
**RefreshTokenExpiresAt** | Pointer to **int64** | The datetime at which the newly issued refresh token will expire. The value is represented in milliseconds since the Unix epoch (1970-01-01).  | [optional] 
**RefreshTokenDuration** | Pointer to **int64** | The duration of the newly issued refresh token in seconds. | [optional] 
**ClientId** | Pointer to **int64** | The client ID. | [optional] 
**ClientIdAlias** | Pointer to **string** | The client ID alias. If the client did not have an alias, this parameter is &#x60;null&#x60;.  | [optional] 
**ClientIdAliasUsed** | Pointer to **bool** | The flag which indicates whether the client ID alias was used when the token request was made. &#x60;true&#x60; if the client ID alias was used when the token request was made.  | [optional] 
**Subject** | Pointer to **string** | The subject (&#x3D; resource owner&#39;s ID) of the access token. Even if an access token has been issued by calling &#x60;/api/auth/token&#x60; API, this parameter is &#x60;null&#x60; if the flow of the token request was [Client Credentials Flow](https://datatracker.ietf.org/doc/html/rfc6749#section-4.4) (&#x60;grant_type&#x3D;client_credentials&#x60;) because it means the access token is not associated with any specific end-user.  | [optional] 
**Scopes** | Pointer to **[]string** | The scopes covered by the access token. | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) | The extra properties associated with the access token. This parameter is &#x60;null&#x60; when no extra property is associated with the issued access token.  | [optional] 
**JwtAccessToken** | Pointer to **string** | The newly issued access token in JWT format. If the authorization server is configured to issue JWT-based access tokens (&#x3D; if the service&#39;s &#x60;accessTokenSignAlg&#x60; value is a non-null value), a JWT-based access token is issued along with the original random-string one.  | [optional] 
**Resources** | Pointer to **[]string** | The resources specified by the &#x60;resource&#x60; request parameters in the token request. See \&quot;Resource Indicators for OAuth 2.0\&quot; for details.  | [optional] 
**AccessTokenResources** | Pointer to **[]string** | The target resources of the access token being issued. See \&quot;Resource Indicators for OAuth 2.0\&quot; for details.  | [optional] 
**AuthorizationDetails** | Pointer to [**AuthorizationDetails**](AuthorizationDetails.md) |  | [optional] 
**ServiceAttributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of this service that the client application belongs to.  | [optional] 
**ClientAttributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of the client.  | [optional] 

## Methods

### NewTokenIssueResponse

`func NewTokenIssueResponse() *TokenIssueResponse`

NewTokenIssueResponse instantiates a new TokenIssueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenIssueResponseWithDefaults

`func NewTokenIssueResponseWithDefaults() *TokenIssueResponse`

NewTokenIssueResponseWithDefaults instantiates a new TokenIssueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *TokenIssueResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *TokenIssueResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *TokenIssueResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *TokenIssueResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *TokenIssueResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *TokenIssueResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *TokenIssueResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *TokenIssueResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *TokenIssueResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TokenIssueResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TokenIssueResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TokenIssueResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResponseContent

`func (o *TokenIssueResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *TokenIssueResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *TokenIssueResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *TokenIssueResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetUsername

`func (o *TokenIssueResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TokenIssueResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TokenIssueResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TokenIssueResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAccessToken

`func (o *TokenIssueResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TokenIssueResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TokenIssueResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *TokenIssueResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAccessTokenExpiresAt

`func (o *TokenIssueResponse) GetAccessTokenExpiresAt() int64`

GetAccessTokenExpiresAt returns the AccessTokenExpiresAt field if non-nil, zero value otherwise.

### GetAccessTokenExpiresAtOk

`func (o *TokenIssueResponse) GetAccessTokenExpiresAtOk() (*int64, bool)`

GetAccessTokenExpiresAtOk returns a tuple with the AccessTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpiresAt

`func (o *TokenIssueResponse) SetAccessTokenExpiresAt(v int64)`

SetAccessTokenExpiresAt sets AccessTokenExpiresAt field to given value.

### HasAccessTokenExpiresAt

`func (o *TokenIssueResponse) HasAccessTokenExpiresAt() bool`

HasAccessTokenExpiresAt returns a boolean if a field has been set.

### GetAccessTokenDuration

`func (o *TokenIssueResponse) GetAccessTokenDuration() int64`

GetAccessTokenDuration returns the AccessTokenDuration field if non-nil, zero value otherwise.

### GetAccessTokenDurationOk

`func (o *TokenIssueResponse) GetAccessTokenDurationOk() (*int64, bool)`

GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenDuration

`func (o *TokenIssueResponse) SetAccessTokenDuration(v int64)`

SetAccessTokenDuration sets AccessTokenDuration field to given value.

### HasAccessTokenDuration

`func (o *TokenIssueResponse) HasAccessTokenDuration() bool`

HasAccessTokenDuration returns a boolean if a field has been set.

### GetRefreshToken

`func (o *TokenIssueResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *TokenIssueResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *TokenIssueResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *TokenIssueResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetRefreshTokenExpiresAt

`func (o *TokenIssueResponse) GetRefreshTokenExpiresAt() int64`

GetRefreshTokenExpiresAt returns the RefreshTokenExpiresAt field if non-nil, zero value otherwise.

### GetRefreshTokenExpiresAtOk

`func (o *TokenIssueResponse) GetRefreshTokenExpiresAtOk() (*int64, bool)`

GetRefreshTokenExpiresAtOk returns a tuple with the RefreshTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenExpiresAt

`func (o *TokenIssueResponse) SetRefreshTokenExpiresAt(v int64)`

SetRefreshTokenExpiresAt sets RefreshTokenExpiresAt field to given value.

### HasRefreshTokenExpiresAt

`func (o *TokenIssueResponse) HasRefreshTokenExpiresAt() bool`

HasRefreshTokenExpiresAt returns a boolean if a field has been set.

### GetRefreshTokenDuration

`func (o *TokenIssueResponse) GetRefreshTokenDuration() int64`

GetRefreshTokenDuration returns the RefreshTokenDuration field if non-nil, zero value otherwise.

### GetRefreshTokenDurationOk

`func (o *TokenIssueResponse) GetRefreshTokenDurationOk() (*int64, bool)`

GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDuration

`func (o *TokenIssueResponse) SetRefreshTokenDuration(v int64)`

SetRefreshTokenDuration sets RefreshTokenDuration field to given value.

### HasRefreshTokenDuration

`func (o *TokenIssueResponse) HasRefreshTokenDuration() bool`

HasRefreshTokenDuration returns a boolean if a field has been set.

### GetClientId

`func (o *TokenIssueResponse) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TokenIssueResponse) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TokenIssueResponse) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TokenIssueResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIdAlias

`func (o *TokenIssueResponse) GetClientIdAlias() string`

GetClientIdAlias returns the ClientIdAlias field if non-nil, zero value otherwise.

### GetClientIdAliasOk

`func (o *TokenIssueResponse) GetClientIdAliasOk() (*string, bool)`

GetClientIdAliasOk returns a tuple with the ClientIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAlias

`func (o *TokenIssueResponse) SetClientIdAlias(v string)`

SetClientIdAlias sets ClientIdAlias field to given value.

### HasClientIdAlias

`func (o *TokenIssueResponse) HasClientIdAlias() bool`

HasClientIdAlias returns a boolean if a field has been set.

### GetClientIdAliasUsed

`func (o *TokenIssueResponse) GetClientIdAliasUsed() bool`

GetClientIdAliasUsed returns the ClientIdAliasUsed field if non-nil, zero value otherwise.

### GetClientIdAliasUsedOk

`func (o *TokenIssueResponse) GetClientIdAliasUsedOk() (*bool, bool)`

GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasUsed

`func (o *TokenIssueResponse) SetClientIdAliasUsed(v bool)`

SetClientIdAliasUsed sets ClientIdAliasUsed field to given value.

### HasClientIdAliasUsed

`func (o *TokenIssueResponse) HasClientIdAliasUsed() bool`

HasClientIdAliasUsed returns a boolean if a field has been set.

### GetSubject

`func (o *TokenIssueResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TokenIssueResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TokenIssueResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TokenIssueResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetScopes

`func (o *TokenIssueResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenIssueResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenIssueResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenIssueResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetProperties

`func (o *TokenIssueResponse) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TokenIssueResponse) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TokenIssueResponse) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TokenIssueResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetJwtAccessToken

`func (o *TokenIssueResponse) GetJwtAccessToken() string`

GetJwtAccessToken returns the JwtAccessToken field if non-nil, zero value otherwise.

### GetJwtAccessTokenOk

`func (o *TokenIssueResponse) GetJwtAccessTokenOk() (*string, bool)`

GetJwtAccessTokenOk returns a tuple with the JwtAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtAccessToken

`func (o *TokenIssueResponse) SetJwtAccessToken(v string)`

SetJwtAccessToken sets JwtAccessToken field to given value.

### HasJwtAccessToken

`func (o *TokenIssueResponse) HasJwtAccessToken() bool`

HasJwtAccessToken returns a boolean if a field has been set.

### GetResources

`func (o *TokenIssueResponse) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *TokenIssueResponse) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *TokenIssueResponse) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *TokenIssueResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetAccessTokenResources

`func (o *TokenIssueResponse) GetAccessTokenResources() []string`

GetAccessTokenResources returns the AccessTokenResources field if non-nil, zero value otherwise.

### GetAccessTokenResourcesOk

`func (o *TokenIssueResponse) GetAccessTokenResourcesOk() (*[]string, bool)`

GetAccessTokenResourcesOk returns a tuple with the AccessTokenResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenResources

`func (o *TokenIssueResponse) SetAccessTokenResources(v []string)`

SetAccessTokenResources sets AccessTokenResources field to given value.

### HasAccessTokenResources

`func (o *TokenIssueResponse) HasAccessTokenResources() bool`

HasAccessTokenResources returns a boolean if a field has been set.

### GetAuthorizationDetails

`func (o *TokenIssueResponse) GetAuthorizationDetails() AuthorizationDetails`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *TokenIssueResponse) GetAuthorizationDetailsOk() (*AuthorizationDetails, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *TokenIssueResponse) SetAuthorizationDetails(v AuthorizationDetails)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *TokenIssueResponse) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.

### GetServiceAttributes

`func (o *TokenIssueResponse) GetServiceAttributes() []Pair`

GetServiceAttributes returns the ServiceAttributes field if non-nil, zero value otherwise.

### GetServiceAttributesOk

`func (o *TokenIssueResponse) GetServiceAttributesOk() (*[]Pair, bool)`

GetServiceAttributesOk returns a tuple with the ServiceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAttributes

`func (o *TokenIssueResponse) SetServiceAttributes(v []Pair)`

SetServiceAttributes sets ServiceAttributes field to given value.

### HasServiceAttributes

`func (o *TokenIssueResponse) HasServiceAttributes() bool`

HasServiceAttributes returns a boolean if a field has been set.

### GetClientAttributes

`func (o *TokenIssueResponse) GetClientAttributes() []Pair`

GetClientAttributes returns the ClientAttributes field if non-nil, zero value otherwise.

### GetClientAttributesOk

`func (o *TokenIssueResponse) GetClientAttributesOk() (*[]Pair, bool)`

GetClientAttributesOk returns a tuple with the ClientAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAttributes

`func (o *TokenIssueResponse) SetClientAttributes(v []Pair)`

SetClientAttributes sets ClientAttributes field to given value.

### HasClientAttributes

`func (o *TokenIssueResponse) HasClientAttributes() bool`

HasClientAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


