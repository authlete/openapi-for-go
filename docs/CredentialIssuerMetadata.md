# CredentialIssuerMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialIssuer** | Pointer to **string** | The identifier of a credential request. | [optional] 
**AuthorizationServer** | Pointer to **string** | The identifier of the authorization server that the credential issuer relies on for authorization.  | [optional] 
**CredentialEndpoint** | Pointer to **bool** | The URL of the credential endpoint of the credential issuer. | [optional] 
**BatchCredentialEndpoint** | Pointer to **int64** | The URL of the batch credential endpoint of the credential issuer. | [optional] 
**DeferredCredentialEndpoint** | Pointer to **string** | The URL of the deferred credential endpoint of the credential issuer. | [optional] 
**CredentialsSupported** | Pointer to **bool** | A JSON array describing supported credentials. | [optional] 

## Methods

### NewCredentialIssuerMetadata

`func NewCredentialIssuerMetadata() *CredentialIssuerMetadata`

NewCredentialIssuerMetadata instantiates a new CredentialIssuerMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialIssuerMetadataWithDefaults

`func NewCredentialIssuerMetadataWithDefaults() *CredentialIssuerMetadata`

NewCredentialIssuerMetadataWithDefaults instantiates a new CredentialIssuerMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialIssuer

`func (o *CredentialIssuerMetadata) GetCredentialIssuer() string`

GetCredentialIssuer returns the CredentialIssuer field if non-nil, zero value otherwise.

### GetCredentialIssuerOk

`func (o *CredentialIssuerMetadata) GetCredentialIssuerOk() (*string, bool)`

GetCredentialIssuerOk returns a tuple with the CredentialIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialIssuer

`func (o *CredentialIssuerMetadata) SetCredentialIssuer(v string)`

SetCredentialIssuer sets CredentialIssuer field to given value.

### HasCredentialIssuer

`func (o *CredentialIssuerMetadata) HasCredentialIssuer() bool`

HasCredentialIssuer returns a boolean if a field has been set.

### GetAuthorizationServer

`func (o *CredentialIssuerMetadata) GetAuthorizationServer() string`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *CredentialIssuerMetadata) GetAuthorizationServerOk() (*string, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *CredentialIssuerMetadata) SetAuthorizationServer(v string)`

SetAuthorizationServer sets AuthorizationServer field to given value.

### HasAuthorizationServer

`func (o *CredentialIssuerMetadata) HasAuthorizationServer() bool`

HasAuthorizationServer returns a boolean if a field has been set.

### GetCredentialEndpoint

`func (o *CredentialIssuerMetadata) GetCredentialEndpoint() bool`

GetCredentialEndpoint returns the CredentialEndpoint field if non-nil, zero value otherwise.

### GetCredentialEndpointOk

`func (o *CredentialIssuerMetadata) GetCredentialEndpointOk() (*bool, bool)`

GetCredentialEndpointOk returns a tuple with the CredentialEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialEndpoint

`func (o *CredentialIssuerMetadata) SetCredentialEndpoint(v bool)`

SetCredentialEndpoint sets CredentialEndpoint field to given value.

### HasCredentialEndpoint

`func (o *CredentialIssuerMetadata) HasCredentialEndpoint() bool`

HasCredentialEndpoint returns a boolean if a field has been set.

### GetBatchCredentialEndpoint

`func (o *CredentialIssuerMetadata) GetBatchCredentialEndpoint() int64`

GetBatchCredentialEndpoint returns the BatchCredentialEndpoint field if non-nil, zero value otherwise.

### GetBatchCredentialEndpointOk

`func (o *CredentialIssuerMetadata) GetBatchCredentialEndpointOk() (*int64, bool)`

GetBatchCredentialEndpointOk returns a tuple with the BatchCredentialEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchCredentialEndpoint

`func (o *CredentialIssuerMetadata) SetBatchCredentialEndpoint(v int64)`

SetBatchCredentialEndpoint sets BatchCredentialEndpoint field to given value.

### HasBatchCredentialEndpoint

`func (o *CredentialIssuerMetadata) HasBatchCredentialEndpoint() bool`

HasBatchCredentialEndpoint returns a boolean if a field has been set.

### GetDeferredCredentialEndpoint

`func (o *CredentialIssuerMetadata) GetDeferredCredentialEndpoint() string`

GetDeferredCredentialEndpoint returns the DeferredCredentialEndpoint field if non-nil, zero value otherwise.

### GetDeferredCredentialEndpointOk

`func (o *CredentialIssuerMetadata) GetDeferredCredentialEndpointOk() (*string, bool)`

GetDeferredCredentialEndpointOk returns a tuple with the DeferredCredentialEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredCredentialEndpoint

`func (o *CredentialIssuerMetadata) SetDeferredCredentialEndpoint(v string)`

SetDeferredCredentialEndpoint sets DeferredCredentialEndpoint field to given value.

### HasDeferredCredentialEndpoint

`func (o *CredentialIssuerMetadata) HasDeferredCredentialEndpoint() bool`

HasDeferredCredentialEndpoint returns a boolean if a field has been set.

### GetCredentialsSupported

`func (o *CredentialIssuerMetadata) GetCredentialsSupported() bool`

GetCredentialsSupported returns the CredentialsSupported field if non-nil, zero value otherwise.

### GetCredentialsSupportedOk

`func (o *CredentialIssuerMetadata) GetCredentialsSupportedOk() (*bool, bool)`

GetCredentialsSupportedOk returns a tuple with the CredentialsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsSupported

`func (o *CredentialIssuerMetadata) SetCredentialsSupported(v bool)`

SetCredentialsSupported sets CredentialsSupported field to given value.

### HasCredentialsSupported

`func (o *CredentialIssuerMetadata) HasCredentialsSupported() bool`

HasCredentialsSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


