# TokenUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take. | [optional] 
**AccessToken** | Pointer to **string** | The access token which has been specified by the request. | [optional] 
**AccessTokenExpiresAt** | Pointer to **int64** | The date at which the access token will expire.  | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) | The extra properties associated with the access token.  | [optional] 
**Scopes** | Pointer to **[]string** | The scopes associated with the access token.  | [optional] 
**AuthorizationDetails** | Pointer to [**AuthzDetails**](AuthzDetails.md) |  | [optional] 
**TokenType** | Pointer to **string** | The token type associated with the access token.  | [optional] 
**ForExternalAttachment** | Pointer to **bool** | the flag which indicates whether the access token is for an external attachment.  | [optional] 
**RefreshTokenExpiresAt** | Pointer to **int64** | The date at which the refresh token will expire. | [optional] 

## Methods

### NewTokenUpdateResponse

`func NewTokenUpdateResponse() *TokenUpdateResponse`

NewTokenUpdateResponse instantiates a new TokenUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenUpdateResponseWithDefaults

`func NewTokenUpdateResponseWithDefaults() *TokenUpdateResponse`

NewTokenUpdateResponseWithDefaults instantiates a new TokenUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *TokenUpdateResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *TokenUpdateResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *TokenUpdateResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *TokenUpdateResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *TokenUpdateResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *TokenUpdateResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *TokenUpdateResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *TokenUpdateResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *TokenUpdateResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TokenUpdateResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TokenUpdateResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TokenUpdateResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAccessToken

`func (o *TokenUpdateResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TokenUpdateResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TokenUpdateResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *TokenUpdateResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAccessTokenExpiresAt

`func (o *TokenUpdateResponse) GetAccessTokenExpiresAt() int64`

GetAccessTokenExpiresAt returns the AccessTokenExpiresAt field if non-nil, zero value otherwise.

### GetAccessTokenExpiresAtOk

`func (o *TokenUpdateResponse) GetAccessTokenExpiresAtOk() (*int64, bool)`

GetAccessTokenExpiresAtOk returns a tuple with the AccessTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpiresAt

`func (o *TokenUpdateResponse) SetAccessTokenExpiresAt(v int64)`

SetAccessTokenExpiresAt sets AccessTokenExpiresAt field to given value.

### HasAccessTokenExpiresAt

`func (o *TokenUpdateResponse) HasAccessTokenExpiresAt() bool`

HasAccessTokenExpiresAt returns a boolean if a field has been set.

### GetProperties

`func (o *TokenUpdateResponse) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TokenUpdateResponse) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TokenUpdateResponse) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TokenUpdateResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetScopes

`func (o *TokenUpdateResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenUpdateResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenUpdateResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenUpdateResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetAuthorizationDetails

`func (o *TokenUpdateResponse) GetAuthorizationDetails() AuthzDetails`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *TokenUpdateResponse) GetAuthorizationDetailsOk() (*AuthzDetails, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *TokenUpdateResponse) SetAuthorizationDetails(v AuthzDetails)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *TokenUpdateResponse) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.

### GetTokenType

`func (o *TokenUpdateResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TokenUpdateResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TokenUpdateResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *TokenUpdateResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetForExternalAttachment

`func (o *TokenUpdateResponse) GetForExternalAttachment() bool`

GetForExternalAttachment returns the ForExternalAttachment field if non-nil, zero value otherwise.

### GetForExternalAttachmentOk

`func (o *TokenUpdateResponse) GetForExternalAttachmentOk() (*bool, bool)`

GetForExternalAttachmentOk returns a tuple with the ForExternalAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForExternalAttachment

`func (o *TokenUpdateResponse) SetForExternalAttachment(v bool)`

SetForExternalAttachment sets ForExternalAttachment field to given value.

### HasForExternalAttachment

`func (o *TokenUpdateResponse) HasForExternalAttachment() bool`

HasForExternalAttachment returns a boolean if a field has been set.

### GetRefreshTokenExpiresAt

`func (o *TokenUpdateResponse) GetRefreshTokenExpiresAt() int64`

GetRefreshTokenExpiresAt returns the RefreshTokenExpiresAt field if non-nil, zero value otherwise.

### GetRefreshTokenExpiresAtOk

`func (o *TokenUpdateResponse) GetRefreshTokenExpiresAtOk() (*int64, bool)`

GetRefreshTokenExpiresAtOk returns a tuple with the RefreshTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenExpiresAt

`func (o *TokenUpdateResponse) SetRefreshTokenExpiresAt(v int64)`

SetRefreshTokenExpiresAt sets RefreshTokenExpiresAt field to given value.

### HasRefreshTokenExpiresAt

`func (o *TokenUpdateResponse) HasRefreshTokenExpiresAt() bool`

HasRefreshTokenExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


