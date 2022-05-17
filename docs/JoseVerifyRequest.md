# JoseVerifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jose** | **string** | A JOSE object.  | 
**MandatoryClaims** | Pointer to **string** | Mandatory claims that are required to be included in the JOSE object.  | [optional] 
**ClockSkew** | Pointer to **int32** | Allowable clock skew in seconds.  | [optional] 
**ClientIdentifier** | Pointer to **string** | The identifier of the client application whose keys are required for verification of the JOSE object.  | [optional] 
**SignedByClient** | Pointer to **bool** | The flag which indicates whether the signature of the JOSE object has been signed by a client application with the client&#39;s private key or a shared symmetric key. | [optional] 

## Methods

### NewJoseVerifyRequest

`func NewJoseVerifyRequest(jose string, ) *JoseVerifyRequest`

NewJoseVerifyRequest instantiates a new JoseVerifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoseVerifyRequestWithDefaults

`func NewJoseVerifyRequestWithDefaults() *JoseVerifyRequest`

NewJoseVerifyRequestWithDefaults instantiates a new JoseVerifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJose

`func (o *JoseVerifyRequest) GetJose() string`

GetJose returns the Jose field if non-nil, zero value otherwise.

### GetJoseOk

`func (o *JoseVerifyRequest) GetJoseOk() (*string, bool)`

GetJoseOk returns a tuple with the Jose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJose

`func (o *JoseVerifyRequest) SetJose(v string)`

SetJose sets Jose field to given value.


### GetMandatoryClaims

`func (o *JoseVerifyRequest) GetMandatoryClaims() string`

GetMandatoryClaims returns the MandatoryClaims field if non-nil, zero value otherwise.

### GetMandatoryClaimsOk

`func (o *JoseVerifyRequest) GetMandatoryClaimsOk() (*string, bool)`

GetMandatoryClaimsOk returns a tuple with the MandatoryClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryClaims

`func (o *JoseVerifyRequest) SetMandatoryClaims(v string)`

SetMandatoryClaims sets MandatoryClaims field to given value.

### HasMandatoryClaims

`func (o *JoseVerifyRequest) HasMandatoryClaims() bool`

HasMandatoryClaims returns a boolean if a field has been set.

### GetClockSkew

`func (o *JoseVerifyRequest) GetClockSkew() int32`

GetClockSkew returns the ClockSkew field if non-nil, zero value otherwise.

### GetClockSkewOk

`func (o *JoseVerifyRequest) GetClockSkewOk() (*int32, bool)`

GetClockSkewOk returns a tuple with the ClockSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockSkew

`func (o *JoseVerifyRequest) SetClockSkew(v int32)`

SetClockSkew sets ClockSkew field to given value.

### HasClockSkew

`func (o *JoseVerifyRequest) HasClockSkew() bool`

HasClockSkew returns a boolean if a field has been set.

### GetClientIdentifier

`func (o *JoseVerifyRequest) GetClientIdentifier() string`

GetClientIdentifier returns the ClientIdentifier field if non-nil, zero value otherwise.

### GetClientIdentifierOk

`func (o *JoseVerifyRequest) GetClientIdentifierOk() (*string, bool)`

GetClientIdentifierOk returns a tuple with the ClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdentifier

`func (o *JoseVerifyRequest) SetClientIdentifier(v string)`

SetClientIdentifier sets ClientIdentifier field to given value.

### HasClientIdentifier

`func (o *JoseVerifyRequest) HasClientIdentifier() bool`

HasClientIdentifier returns a boolean if a field has been set.

### GetSignedByClient

`func (o *JoseVerifyRequest) GetSignedByClient() bool`

GetSignedByClient returns the SignedByClient field if non-nil, zero value otherwise.

### GetSignedByClientOk

`func (o *JoseVerifyRequest) GetSignedByClientOk() (*bool, bool)`

GetSignedByClientOk returns a tuple with the SignedByClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedByClient

`func (o *JoseVerifyRequest) SetSignedByClient(v bool)`

SetSignedByClient sets SignedByClient field to given value.

### HasSignedByClient

`func (o *JoseVerifyRequest) HasSignedByClient() bool`

HasSignedByClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


