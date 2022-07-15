# Grant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | Pointer to [**[]GrantScope**](GrantScope.md) |  | [optional] 
**Claims** | Pointer to **[]string** | The claims associated with the Grant.  | [optional] 
**AuthorizationDetails** | Pointer to [**AuthorizationDetails**](AuthorizationDetails.md) |  | [optional] 

## Methods

### NewGrant

`func NewGrant() *Grant`

NewGrant instantiates a new Grant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantWithDefaults

`func NewGrantWithDefaults() *Grant`

NewGrantWithDefaults instantiates a new Grant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopes

`func (o *Grant) GetScopes() []GrantScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *Grant) GetScopesOk() (*[]GrantScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *Grant) SetScopes(v []GrantScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *Grant) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetClaims

`func (o *Grant) GetClaims() []string`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *Grant) GetClaimsOk() (*[]string, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *Grant) SetClaims(v []string)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *Grant) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetAuthorizationDetails

`func (o *Grant) GetAuthorizationDetails() AuthorizationDetails`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *Grant) GetAuthorizationDetailsOk() (*AuthorizationDetails, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *Grant) SetAuthorizationDetails(v AuthorizationDetails)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *Grant) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


