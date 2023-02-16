# TokenRevokeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenIdentifier** | Pointer to **string** | The identifier of an access token to revoke.  | [optional] 
**ClientIdentifier** | Pointer to **string** | The identifier of a client.  | [optional] 
**RefreshTokenIdentifier** | Pointer to **string** | The identifier of a refresh token to revoke.  | [optional] 
**Subject** | Pointer to **string** | The subject of a resource owner. | [optional] 

## Methods

### NewTokenRevokeRequest

`func NewTokenRevokeRequest() *TokenRevokeRequest`

NewTokenRevokeRequest instantiates a new TokenRevokeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRevokeRequestWithDefaults

`func NewTokenRevokeRequestWithDefaults() *TokenRevokeRequest`

NewTokenRevokeRequestWithDefaults instantiates a new TokenRevokeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenIdentifier

`func (o *TokenRevokeRequest) GetAccessTokenIdentifier() string`

GetAccessTokenIdentifier returns the AccessTokenIdentifier field if non-nil, zero value otherwise.

### GetAccessTokenIdentifierOk

`func (o *TokenRevokeRequest) GetAccessTokenIdentifierOk() (*string, bool)`

GetAccessTokenIdentifierOk returns a tuple with the AccessTokenIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenIdentifier

`func (o *TokenRevokeRequest) SetAccessTokenIdentifier(v string)`

SetAccessTokenIdentifier sets AccessTokenIdentifier field to given value.

### HasAccessTokenIdentifier

`func (o *TokenRevokeRequest) HasAccessTokenIdentifier() bool`

HasAccessTokenIdentifier returns a boolean if a field has been set.

### GetClientIdentifier

`func (o *TokenRevokeRequest) GetClientIdentifier() string`

GetClientIdentifier returns the ClientIdentifier field if non-nil, zero value otherwise.

### GetClientIdentifierOk

`func (o *TokenRevokeRequest) GetClientIdentifierOk() (*string, bool)`

GetClientIdentifierOk returns a tuple with the ClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdentifier

`func (o *TokenRevokeRequest) SetClientIdentifier(v string)`

SetClientIdentifier sets ClientIdentifier field to given value.

### HasClientIdentifier

`func (o *TokenRevokeRequest) HasClientIdentifier() bool`

HasClientIdentifier returns a boolean if a field has been set.

### GetRefreshTokenIdentifier

`func (o *TokenRevokeRequest) GetRefreshTokenIdentifier() string`

GetRefreshTokenIdentifier returns the RefreshTokenIdentifier field if non-nil, zero value otherwise.

### GetRefreshTokenIdentifierOk

`func (o *TokenRevokeRequest) GetRefreshTokenIdentifierOk() (*string, bool)`

GetRefreshTokenIdentifierOk returns a tuple with the RefreshTokenIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenIdentifier

`func (o *TokenRevokeRequest) SetRefreshTokenIdentifier(v string)`

SetRefreshTokenIdentifier sets RefreshTokenIdentifier field to given value.

### HasRefreshTokenIdentifier

`func (o *TokenRevokeRequest) HasRefreshTokenIdentifier() bool`

HasRefreshTokenIdentifier returns a boolean if a field has been set.

### GetSubject

`func (o *TokenRevokeRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TokenRevokeRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TokenRevokeRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TokenRevokeRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


