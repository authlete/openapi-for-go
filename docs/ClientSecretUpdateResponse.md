# ClientSecretUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**NewClientSecret** | Pointer to **string** | The new client secret.  | [optional] 
**OldClientSecret** | Pointer to **string** | The old client secret.  | [optional] 

## Methods

### NewClientSecretUpdateResponse

`func NewClientSecretUpdateResponse() *ClientSecretUpdateResponse`

NewClientSecretUpdateResponse instantiates a new ClientSecretUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientSecretUpdateResponseWithDefaults

`func NewClientSecretUpdateResponseWithDefaults() *ClientSecretUpdateResponse`

NewClientSecretUpdateResponseWithDefaults instantiates a new ClientSecretUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *ClientSecretUpdateResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *ClientSecretUpdateResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *ClientSecretUpdateResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *ClientSecretUpdateResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *ClientSecretUpdateResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *ClientSecretUpdateResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *ClientSecretUpdateResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *ClientSecretUpdateResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetNewClientSecret

`func (o *ClientSecretUpdateResponse) GetNewClientSecret() string`

GetNewClientSecret returns the NewClientSecret field if non-nil, zero value otherwise.

### GetNewClientSecretOk

`func (o *ClientSecretUpdateResponse) GetNewClientSecretOk() (*string, bool)`

GetNewClientSecretOk returns a tuple with the NewClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewClientSecret

`func (o *ClientSecretUpdateResponse) SetNewClientSecret(v string)`

SetNewClientSecret sets NewClientSecret field to given value.

### HasNewClientSecret

`func (o *ClientSecretUpdateResponse) HasNewClientSecret() bool`

HasNewClientSecret returns a boolean if a field has been set.

### GetOldClientSecret

`func (o *ClientSecretUpdateResponse) GetOldClientSecret() string`

GetOldClientSecret returns the OldClientSecret field if non-nil, zero value otherwise.

### GetOldClientSecretOk

`func (o *ClientSecretUpdateResponse) GetOldClientSecretOk() (*string, bool)`

GetOldClientSecretOk returns a tuple with the OldClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldClientSecret

`func (o *ClientSecretUpdateResponse) SetOldClientSecret(v string)`

SetOldClientSecret sets OldClientSecret field to given value.

### HasOldClientSecret

`func (o *ClientSecretUpdateResponse) HasOldClientSecret() bool`

HasOldClientSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


