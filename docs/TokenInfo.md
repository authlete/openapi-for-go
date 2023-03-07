# TokenInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **int32** | The client id. | [optional] 
**ClientIdAlias** | Pointer to **string** | The alias of the client. | [optional] 
**ClientIdAliasUsed** | Pointer to **bool** | Flag specifying if the alias was used to identify the client | [optional] 
**Subject** | Pointer to **string** | the resource owner unique id | [optional] 
**Scopes** | Pointer to **[]string** | The scopes granted on the token | [optional] 
**ExpiresAt** | Pointer to **int32** | time which the token expires. | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) | Extra properties associated with the token | [optional] 
**Resources** | Pointer to **[]string** | The array of the resources of the token. | [optional] 
**AuthorizationDetails** | Pointer to [**AuthorizationDetailsElement**](AuthorizationDetailsElement.md) |  | [optional] 
**ClientEntityId** | Pointer to **string** | The entity ID of the client.  | [optional] 
**ClientEntityIdUsed** | Pointer to **bool** | Flag which indicates whether the entity ID of the client was used when the request for the access token was made. | [optional] 

## Methods

### NewTokenInfo

`func NewTokenInfo() *TokenInfo`

NewTokenInfo instantiates a new TokenInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenInfoWithDefaults

`func NewTokenInfoWithDefaults() *TokenInfo`

NewTokenInfoWithDefaults instantiates a new TokenInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *TokenInfo) GetClientId() int32`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TokenInfo) GetClientIdOk() (*int32, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TokenInfo) SetClientId(v int32)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TokenInfo) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIdAlias

`func (o *TokenInfo) GetClientIdAlias() string`

GetClientIdAlias returns the ClientIdAlias field if non-nil, zero value otherwise.

### GetClientIdAliasOk

`func (o *TokenInfo) GetClientIdAliasOk() (*string, bool)`

GetClientIdAliasOk returns a tuple with the ClientIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAlias

`func (o *TokenInfo) SetClientIdAlias(v string)`

SetClientIdAlias sets ClientIdAlias field to given value.

### HasClientIdAlias

`func (o *TokenInfo) HasClientIdAlias() bool`

HasClientIdAlias returns a boolean if a field has been set.

### GetClientIdAliasUsed

`func (o *TokenInfo) GetClientIdAliasUsed() bool`

GetClientIdAliasUsed returns the ClientIdAliasUsed field if non-nil, zero value otherwise.

### GetClientIdAliasUsedOk

`func (o *TokenInfo) GetClientIdAliasUsedOk() (*bool, bool)`

GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasUsed

`func (o *TokenInfo) SetClientIdAliasUsed(v bool)`

SetClientIdAliasUsed sets ClientIdAliasUsed field to given value.

### HasClientIdAliasUsed

`func (o *TokenInfo) HasClientIdAliasUsed() bool`

HasClientIdAliasUsed returns a boolean if a field has been set.

### GetSubject

`func (o *TokenInfo) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TokenInfo) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TokenInfo) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TokenInfo) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetScopes

`func (o *TokenInfo) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenInfo) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenInfo) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenInfo) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetExpiresAt

`func (o *TokenInfo) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TokenInfo) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TokenInfo) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TokenInfo) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetProperties

`func (o *TokenInfo) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TokenInfo) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TokenInfo) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TokenInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetResources

`func (o *TokenInfo) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *TokenInfo) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *TokenInfo) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *TokenInfo) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetAuthorizationDetails

`func (o *TokenInfo) GetAuthorizationDetails() AuthorizationDetailsElement`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *TokenInfo) GetAuthorizationDetailsOk() (*AuthorizationDetailsElement, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *TokenInfo) SetAuthorizationDetails(v AuthorizationDetailsElement)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *TokenInfo) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.

### GetClientEntityId

`func (o *TokenInfo) GetClientEntityId() string`

GetClientEntityId returns the ClientEntityId field if non-nil, zero value otherwise.

### GetClientEntityIdOk

`func (o *TokenInfo) GetClientEntityIdOk() (*string, bool)`

GetClientEntityIdOk returns a tuple with the ClientEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEntityId

`func (o *TokenInfo) SetClientEntityId(v string)`

SetClientEntityId sets ClientEntityId field to given value.

### HasClientEntityId

`func (o *TokenInfo) HasClientEntityId() bool`

HasClientEntityId returns a boolean if a field has been set.

### GetClientEntityIdUsed

`func (o *TokenInfo) GetClientEntityIdUsed() bool`

GetClientEntityIdUsed returns the ClientEntityIdUsed field if non-nil, zero value otherwise.

### GetClientEntityIdUsedOk

`func (o *TokenInfo) GetClientEntityIdUsedOk() (*bool, bool)`

GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEntityIdUsed

`func (o *TokenInfo) SetClientEntityIdUsed(v bool)`

SetClientEntityIdUsed sets ClientEntityIdUsed field to given value.

### HasClientEntityIdUsed

`func (o *TokenInfo) HasClientEntityIdUsed() bool`

HasClientEntityIdUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


