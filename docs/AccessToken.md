# AccessToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenHash** | Pointer to **string** | The hash of the access token. | [optional] 
**AccessTokenExpiresAt** | Pointer to **int64** | The timestamp at which the access token will expire. | [optional] 
**RefreshTokenHash** | Pointer to **string** | The hash of the refresh token. | [optional] 
**RefreshTokenExpiresAt** | Pointer to **int64** | The timestamp at which the refresh token will expire. | [optional] 
**CreatedAt** | Pointer to **int64** | The timestamp at which the access token was first created.  | [optional] 
**LastRefreshedAt** | Pointer to **int64** | The timestamp at which the access token was last refreshed using the refresh token.  | [optional] 
**ClientId** | Pointer to **int64** | The ID of the client associated with the access token.  | [optional] 
**Subject** | Pointer to **string** | The subject (&#x3D; unique user ID) associated with the access token.  | [optional] 
**GrantType** | Pointer to [**GrantType**](GrantType.md) |  | [optional] 
**Scopes** | Pointer to **[]string** | The scopes associated with the access token.  | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) | The properties associated with the access token.  | [optional] 

## Methods

### NewAccessToken

`func NewAccessToken() *AccessToken`

NewAccessToken instantiates a new AccessToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenWithDefaults

`func NewAccessTokenWithDefaults() *AccessToken`

NewAccessTokenWithDefaults instantiates a new AccessToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenHash

`func (o *AccessToken) GetAccessTokenHash() string`

GetAccessTokenHash returns the AccessTokenHash field if non-nil, zero value otherwise.

### GetAccessTokenHashOk

`func (o *AccessToken) GetAccessTokenHashOk() (*string, bool)`

GetAccessTokenHashOk returns a tuple with the AccessTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenHash

`func (o *AccessToken) SetAccessTokenHash(v string)`

SetAccessTokenHash sets AccessTokenHash field to given value.

### HasAccessTokenHash

`func (o *AccessToken) HasAccessTokenHash() bool`

HasAccessTokenHash returns a boolean if a field has been set.

### GetAccessTokenExpiresAt

`func (o *AccessToken) GetAccessTokenExpiresAt() int64`

GetAccessTokenExpiresAt returns the AccessTokenExpiresAt field if non-nil, zero value otherwise.

### GetAccessTokenExpiresAtOk

`func (o *AccessToken) GetAccessTokenExpiresAtOk() (*int64, bool)`

GetAccessTokenExpiresAtOk returns a tuple with the AccessTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpiresAt

`func (o *AccessToken) SetAccessTokenExpiresAt(v int64)`

SetAccessTokenExpiresAt sets AccessTokenExpiresAt field to given value.

### HasAccessTokenExpiresAt

`func (o *AccessToken) HasAccessTokenExpiresAt() bool`

HasAccessTokenExpiresAt returns a boolean if a field has been set.

### GetRefreshTokenHash

`func (o *AccessToken) GetRefreshTokenHash() string`

GetRefreshTokenHash returns the RefreshTokenHash field if non-nil, zero value otherwise.

### GetRefreshTokenHashOk

`func (o *AccessToken) GetRefreshTokenHashOk() (*string, bool)`

GetRefreshTokenHashOk returns a tuple with the RefreshTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenHash

`func (o *AccessToken) SetRefreshTokenHash(v string)`

SetRefreshTokenHash sets RefreshTokenHash field to given value.

### HasRefreshTokenHash

`func (o *AccessToken) HasRefreshTokenHash() bool`

HasRefreshTokenHash returns a boolean if a field has been set.

### GetRefreshTokenExpiresAt

`func (o *AccessToken) GetRefreshTokenExpiresAt() int64`

GetRefreshTokenExpiresAt returns the RefreshTokenExpiresAt field if non-nil, zero value otherwise.

### GetRefreshTokenExpiresAtOk

`func (o *AccessToken) GetRefreshTokenExpiresAtOk() (*int64, bool)`

GetRefreshTokenExpiresAtOk returns a tuple with the RefreshTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenExpiresAt

`func (o *AccessToken) SetRefreshTokenExpiresAt(v int64)`

SetRefreshTokenExpiresAt sets RefreshTokenExpiresAt field to given value.

### HasRefreshTokenExpiresAt

`func (o *AccessToken) HasRefreshTokenExpiresAt() bool`

HasRefreshTokenExpiresAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccessToken) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccessToken) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccessToken) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccessToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastRefreshedAt

`func (o *AccessToken) GetLastRefreshedAt() int64`

GetLastRefreshedAt returns the LastRefreshedAt field if non-nil, zero value otherwise.

### GetLastRefreshedAtOk

`func (o *AccessToken) GetLastRefreshedAtOk() (*int64, bool)`

GetLastRefreshedAtOk returns a tuple with the LastRefreshedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshedAt

`func (o *AccessToken) SetLastRefreshedAt(v int64)`

SetLastRefreshedAt sets LastRefreshedAt field to given value.

### HasLastRefreshedAt

`func (o *AccessToken) HasLastRefreshedAt() bool`

HasLastRefreshedAt returns a boolean if a field has been set.

### GetClientId

`func (o *AccessToken) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AccessToken) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AccessToken) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AccessToken) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSubject

`func (o *AccessToken) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AccessToken) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AccessToken) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AccessToken) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetGrantType

`func (o *AccessToken) GetGrantType() GrantType`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *AccessToken) GetGrantTypeOk() (*GrantType, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *AccessToken) SetGrantType(v GrantType)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *AccessToken) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetScopes

`func (o *AccessToken) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AccessToken) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AccessToken) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AccessToken) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetProperties

`func (o *AccessToken) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AccessToken) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AccessToken) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AccessToken) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


