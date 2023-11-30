# TokenCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take. | [optional] 
**AccessToken** | Pointer to **string** | The newly issued access token. | [optional] 
**ClientId** | Pointer to **int64** | The ID of the client application associated with the access token.  | [optional] 
**ExpiresAt** | Pointer to **int64** | The time at which the access token expires.  | [optional] 
**ExpiresIn** | Pointer to **int64** | The duration of the newly issued access token in seconds.  | [optional] 
**GrantType** | Pointer to **string** | The grant type for the newly issued access token.  | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) | The extra properties associated with the access token.  | [optional] 
**RefreshToken** | Pointer to **string** | The newly issued refresh token.  | [optional] 
**Scopes** | Pointer to **[]string** | Scopes which are associated with the access token.  | [optional] 
**Subject** | Pointer to **string** | The subject (&#x3D; unique identifier) of the user associated with the newly issued access token.  | [optional] 
**TokenType** | Pointer to **string** | The token type of the access token.  | [optional] 
**JwtAccessToken** | Pointer to **string** | If the authorization server is configured to issue JWT-based access tokens (&#x3D; if &#x60;Service.accessTokenSignAlg&#x60; is set to a &#x60;non-null&#x60; value), a JWT-based access token is issued along with the original random-string one.  | [optional] 
**AuthorizationDetails** | Pointer to [**AuthorizationDetails**](AuthorizationDetails.md) |  | [optional] 
**ForExternalAttachment** | Pointer to **bool** | the flag which indicates whether the access token is for an external attachment. | [optional] 

## Methods

### NewTokenCreateResponse

`func NewTokenCreateResponse() *TokenCreateResponse`

NewTokenCreateResponse instantiates a new TokenCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCreateResponseWithDefaults

`func NewTokenCreateResponseWithDefaults() *TokenCreateResponse`

NewTokenCreateResponseWithDefaults instantiates a new TokenCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *TokenCreateResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *TokenCreateResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *TokenCreateResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *TokenCreateResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *TokenCreateResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *TokenCreateResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *TokenCreateResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *TokenCreateResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *TokenCreateResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TokenCreateResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TokenCreateResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TokenCreateResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAccessToken

`func (o *TokenCreateResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TokenCreateResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TokenCreateResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *TokenCreateResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetClientId

`func (o *TokenCreateResponse) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TokenCreateResponse) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TokenCreateResponse) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TokenCreateResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *TokenCreateResponse) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TokenCreateResponse) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TokenCreateResponse) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TokenCreateResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetExpiresIn

`func (o *TokenCreateResponse) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *TokenCreateResponse) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *TokenCreateResponse) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *TokenCreateResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetGrantType

`func (o *TokenCreateResponse) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *TokenCreateResponse) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *TokenCreateResponse) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *TokenCreateResponse) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetProperties

`func (o *TokenCreateResponse) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TokenCreateResponse) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TokenCreateResponse) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TokenCreateResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRefreshToken

`func (o *TokenCreateResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *TokenCreateResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *TokenCreateResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *TokenCreateResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetScopes

`func (o *TokenCreateResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenCreateResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenCreateResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenCreateResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSubject

`func (o *TokenCreateResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TokenCreateResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TokenCreateResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TokenCreateResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTokenType

`func (o *TokenCreateResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TokenCreateResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TokenCreateResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *TokenCreateResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetJwtAccessToken

`func (o *TokenCreateResponse) GetJwtAccessToken() string`

GetJwtAccessToken returns the JwtAccessToken field if non-nil, zero value otherwise.

### GetJwtAccessTokenOk

`func (o *TokenCreateResponse) GetJwtAccessTokenOk() (*string, bool)`

GetJwtAccessTokenOk returns a tuple with the JwtAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtAccessToken

`func (o *TokenCreateResponse) SetJwtAccessToken(v string)`

SetJwtAccessToken sets JwtAccessToken field to given value.

### HasJwtAccessToken

`func (o *TokenCreateResponse) HasJwtAccessToken() bool`

HasJwtAccessToken returns a boolean if a field has been set.

### GetAuthorizationDetails

`func (o *TokenCreateResponse) GetAuthorizationDetails() AuthorizationDetails`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *TokenCreateResponse) GetAuthorizationDetailsOk() (*AuthorizationDetails, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *TokenCreateResponse) SetAuthorizationDetails(v AuthorizationDetails)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *TokenCreateResponse) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.

### GetForExternalAttachment

`func (o *TokenCreateResponse) GetForExternalAttachment() bool`

GetForExternalAttachment returns the ForExternalAttachment field if non-nil, zero value otherwise.

### GetForExternalAttachmentOk

`func (o *TokenCreateResponse) GetForExternalAttachmentOk() (*bool, bool)`

GetForExternalAttachmentOk returns a tuple with the ForExternalAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForExternalAttachment

`func (o *TokenCreateResponse) SetForExternalAttachment(v bool)`

SetForExternalAttachment sets ForExternalAttachment field to given value.

### HasForExternalAttachment

`func (o *TokenCreateResponse) HasForExternalAttachment() bool`

HasForExternalAttachment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


