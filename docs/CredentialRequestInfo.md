# CredentialRequestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The identifier of the credential offer. | [optional] 
**Format** | Pointer to **string** | The value of the format parameter in the credential request. | [optional] 
**BindingKey** | Pointer to **string** | The binding key specified by the proof in the credential request. | [optional] 
**Details** | Pointer to **string** | The details about the credential request. | [optional] 

## Methods

### NewCredentialRequestInfo

`func NewCredentialRequestInfo() *CredentialRequestInfo`

NewCredentialRequestInfo instantiates a new CredentialRequestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialRequestInfoWithDefaults

`func NewCredentialRequestInfoWithDefaults() *CredentialRequestInfo`

NewCredentialRequestInfoWithDefaults instantiates a new CredentialRequestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *CredentialRequestInfo) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CredentialRequestInfo) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CredentialRequestInfo) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CredentialRequestInfo) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetFormat

`func (o *CredentialRequestInfo) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CredentialRequestInfo) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CredentialRequestInfo) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CredentialRequestInfo) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetBindingKey

`func (o *CredentialRequestInfo) GetBindingKey() string`

GetBindingKey returns the BindingKey field if non-nil, zero value otherwise.

### GetBindingKeyOk

`func (o *CredentialRequestInfo) GetBindingKeyOk() (*string, bool)`

GetBindingKeyOk returns a tuple with the BindingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingKey

`func (o *CredentialRequestInfo) SetBindingKey(v string)`

SetBindingKey sets BindingKey field to given value.

### HasBindingKey

`func (o *CredentialRequestInfo) HasBindingKey() bool`

HasBindingKey returns a boolean if a field has been set.

### GetDetails

`func (o *CredentialRequestInfo) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CredentialRequestInfo) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CredentialRequestInfo) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CredentialRequestInfo) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


