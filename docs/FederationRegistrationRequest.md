# FederationRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityConfiguration** | Pointer to **string** | The entity configuration of a relying party.  | [optional] 
**TrustChain** | Pointer to **string** | The trust chain of a relying party.  | [optional] 

## Methods

### NewFederationRegistrationRequest

`func NewFederationRegistrationRequest() *FederationRegistrationRequest`

NewFederationRegistrationRequest instantiates a new FederationRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationRegistrationRequestWithDefaults

`func NewFederationRegistrationRequestWithDefaults() *FederationRegistrationRequest`

NewFederationRegistrationRequestWithDefaults instantiates a new FederationRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityConfiguration

`func (o *FederationRegistrationRequest) GetEntityConfiguration() string`

GetEntityConfiguration returns the EntityConfiguration field if non-nil, zero value otherwise.

### GetEntityConfigurationOk

`func (o *FederationRegistrationRequest) GetEntityConfigurationOk() (*string, bool)`

GetEntityConfigurationOk returns a tuple with the EntityConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityConfiguration

`func (o *FederationRegistrationRequest) SetEntityConfiguration(v string)`

SetEntityConfiguration sets EntityConfiguration field to given value.

### HasEntityConfiguration

`func (o *FederationRegistrationRequest) HasEntityConfiguration() bool`

HasEntityConfiguration returns a boolean if a field has been set.

### GetTrustChain

`func (o *FederationRegistrationRequest) GetTrustChain() string`

GetTrustChain returns the TrustChain field if non-nil, zero value otherwise.

### GetTrustChainOk

`func (o *FederationRegistrationRequest) GetTrustChainOk() (*string, bool)`

GetTrustChainOk returns a tuple with the TrustChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustChain

`func (o *FederationRegistrationRequest) SetTrustChain(v string)`

SetTrustChain sets TrustChain field to given value.

### HasTrustChain

`func (o *FederationRegistrationRequest) HasTrustChain() bool`

HasTrustChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


