# CredentialIssuanceOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestIdentifier** | Pointer to **string** | The identifier of a credential request. | [optional] 
**CredentialPayload** | Pointer to **string** | The additional payload that will be added into a credential to be issued. | [optional] 
**IssuanceDeferred** | Pointer to **bool** | The flag indicating whether to defer credential issuance. | [optional] 
**CredentialDuration** | Pointer to **int64** | The duration of a credential to be issued. | [optional] 
**SigningKeyId** | Pointer to **string** | The key ID of a private key that should be used for signing a credential to be issued. | [optional] 

## Methods

### NewCredentialIssuanceOrder

`func NewCredentialIssuanceOrder() *CredentialIssuanceOrder`

NewCredentialIssuanceOrder instantiates a new CredentialIssuanceOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialIssuanceOrderWithDefaults

`func NewCredentialIssuanceOrderWithDefaults() *CredentialIssuanceOrder`

NewCredentialIssuanceOrderWithDefaults instantiates a new CredentialIssuanceOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestIdentifier

`func (o *CredentialIssuanceOrder) GetRequestIdentifier() string`

GetRequestIdentifier returns the RequestIdentifier field if non-nil, zero value otherwise.

### GetRequestIdentifierOk

`func (o *CredentialIssuanceOrder) GetRequestIdentifierOk() (*string, bool)`

GetRequestIdentifierOk returns a tuple with the RequestIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestIdentifier

`func (o *CredentialIssuanceOrder) SetRequestIdentifier(v string)`

SetRequestIdentifier sets RequestIdentifier field to given value.

### HasRequestIdentifier

`func (o *CredentialIssuanceOrder) HasRequestIdentifier() bool`

HasRequestIdentifier returns a boolean if a field has been set.

### GetCredentialPayload

`func (o *CredentialIssuanceOrder) GetCredentialPayload() string`

GetCredentialPayload returns the CredentialPayload field if non-nil, zero value otherwise.

### GetCredentialPayloadOk

`func (o *CredentialIssuanceOrder) GetCredentialPayloadOk() (*string, bool)`

GetCredentialPayloadOk returns a tuple with the CredentialPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialPayload

`func (o *CredentialIssuanceOrder) SetCredentialPayload(v string)`

SetCredentialPayload sets CredentialPayload field to given value.

### HasCredentialPayload

`func (o *CredentialIssuanceOrder) HasCredentialPayload() bool`

HasCredentialPayload returns a boolean if a field has been set.

### GetIssuanceDeferred

`func (o *CredentialIssuanceOrder) GetIssuanceDeferred() bool`

GetIssuanceDeferred returns the IssuanceDeferred field if non-nil, zero value otherwise.

### GetIssuanceDeferredOk

`func (o *CredentialIssuanceOrder) GetIssuanceDeferredOk() (*bool, bool)`

GetIssuanceDeferredOk returns a tuple with the IssuanceDeferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceDeferred

`func (o *CredentialIssuanceOrder) SetIssuanceDeferred(v bool)`

SetIssuanceDeferred sets IssuanceDeferred field to given value.

### HasIssuanceDeferred

`func (o *CredentialIssuanceOrder) HasIssuanceDeferred() bool`

HasIssuanceDeferred returns a boolean if a field has been set.

### GetCredentialDuration

`func (o *CredentialIssuanceOrder) GetCredentialDuration() int64`

GetCredentialDuration returns the CredentialDuration field if non-nil, zero value otherwise.

### GetCredentialDurationOk

`func (o *CredentialIssuanceOrder) GetCredentialDurationOk() (*int64, bool)`

GetCredentialDurationOk returns a tuple with the CredentialDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialDuration

`func (o *CredentialIssuanceOrder) SetCredentialDuration(v int64)`

SetCredentialDuration sets CredentialDuration field to given value.

### HasCredentialDuration

`func (o *CredentialIssuanceOrder) HasCredentialDuration() bool`

HasCredentialDuration returns a boolean if a field has been set.

### GetSigningKeyId

`func (o *CredentialIssuanceOrder) GetSigningKeyId() string`

GetSigningKeyId returns the SigningKeyId field if non-nil, zero value otherwise.

### GetSigningKeyIdOk

`func (o *CredentialIssuanceOrder) GetSigningKeyIdOk() (*string, bool)`

GetSigningKeyIdOk returns a tuple with the SigningKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKeyId

`func (o *CredentialIssuanceOrder) SetSigningKeyId(v string)`

SetSigningKeyId sets SigningKeyId field to given value.

### HasSigningKeyId

`func (o *CredentialIssuanceOrder) HasSigningKeyId() bool`

HasSigningKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


