# JoseVerifyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Valid** | Pointer to **bool** | The result of the verification on the JOSE object.  | [optional] 
**SignatureValid** | Pointer to **bool** | The result of the signature verification.  | [optional] 
**MissingClaims** | Pointer to **[]string** | The list of missing claims.  | [optional] 
**InvalidClaims** | Pointer to **[]string** | The list of invalid claims.  | [optional] 
**ErrorDescriptions** | Pointer to **[]string** | The list of error messages.  | [optional] 

## Methods

### NewJoseVerifyResponse

`func NewJoseVerifyResponse() *JoseVerifyResponse`

NewJoseVerifyResponse instantiates a new JoseVerifyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoseVerifyResponseWithDefaults

`func NewJoseVerifyResponseWithDefaults() *JoseVerifyResponse`

NewJoseVerifyResponseWithDefaults instantiates a new JoseVerifyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *JoseVerifyResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *JoseVerifyResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *JoseVerifyResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *JoseVerifyResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *JoseVerifyResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *JoseVerifyResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *JoseVerifyResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *JoseVerifyResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetValid

`func (o *JoseVerifyResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *JoseVerifyResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *JoseVerifyResponse) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *JoseVerifyResponse) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetSignatureValid

`func (o *JoseVerifyResponse) GetSignatureValid() bool`

GetSignatureValid returns the SignatureValid field if non-nil, zero value otherwise.

### GetSignatureValidOk

`func (o *JoseVerifyResponse) GetSignatureValidOk() (*bool, bool)`

GetSignatureValidOk returns a tuple with the SignatureValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureValid

`func (o *JoseVerifyResponse) SetSignatureValid(v bool)`

SetSignatureValid sets SignatureValid field to given value.

### HasSignatureValid

`func (o *JoseVerifyResponse) HasSignatureValid() bool`

HasSignatureValid returns a boolean if a field has been set.

### GetMissingClaims

`func (o *JoseVerifyResponse) GetMissingClaims() []string`

GetMissingClaims returns the MissingClaims field if non-nil, zero value otherwise.

### GetMissingClaimsOk

`func (o *JoseVerifyResponse) GetMissingClaimsOk() (*[]string, bool)`

GetMissingClaimsOk returns a tuple with the MissingClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingClaims

`func (o *JoseVerifyResponse) SetMissingClaims(v []string)`

SetMissingClaims sets MissingClaims field to given value.

### HasMissingClaims

`func (o *JoseVerifyResponse) HasMissingClaims() bool`

HasMissingClaims returns a boolean if a field has been set.

### GetInvalidClaims

`func (o *JoseVerifyResponse) GetInvalidClaims() []string`

GetInvalidClaims returns the InvalidClaims field if non-nil, zero value otherwise.

### GetInvalidClaimsOk

`func (o *JoseVerifyResponse) GetInvalidClaimsOk() (*[]string, bool)`

GetInvalidClaimsOk returns a tuple with the InvalidClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidClaims

`func (o *JoseVerifyResponse) SetInvalidClaims(v []string)`

SetInvalidClaims sets InvalidClaims field to given value.

### HasInvalidClaims

`func (o *JoseVerifyResponse) HasInvalidClaims() bool`

HasInvalidClaims returns a boolean if a field has been set.

### GetErrorDescriptions

`func (o *JoseVerifyResponse) GetErrorDescriptions() []string`

GetErrorDescriptions returns the ErrorDescriptions field if non-nil, zero value otherwise.

### GetErrorDescriptionsOk

`func (o *JoseVerifyResponse) GetErrorDescriptionsOk() (*[]string, bool)`

GetErrorDescriptionsOk returns a tuple with the ErrorDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescriptions

`func (o *JoseVerifyResponse) SetErrorDescriptions(v []string)`

SetErrorDescriptions sets ErrorDescriptions field to given value.

### HasErrorDescriptions

`func (o *JoseVerifyResponse) HasErrorDescriptions() bool`

HasErrorDescriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


