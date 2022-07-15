# TrustAnchor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** | the entity ID of the trust anchor  | [optional] 
**Jwks** | Pointer to **string** | the JWK Set document containing public keys of the trust anchor  | [optional] 

## Methods

### NewTrustAnchor

`func NewTrustAnchor() *TrustAnchor`

NewTrustAnchor instantiates a new TrustAnchor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustAnchorWithDefaults

`func NewTrustAnchorWithDefaults() *TrustAnchor`

NewTrustAnchorWithDefaults instantiates a new TrustAnchor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *TrustAnchor) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *TrustAnchor) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *TrustAnchor) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *TrustAnchor) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetJwks

`func (o *TrustAnchor) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *TrustAnchor) GetJwksOk() (*string, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *TrustAnchor) SetJwks(v string)`

SetJwks sets Jwks field to given value.

### HasJwks

`func (o *TrustAnchor) HasJwks() bool`

HasJwks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


