# PushedAuthReqRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | **string** | The pushed authorization request body received from the client application.  The value of parameters is the entire entity body (which is formatted in &#x60;application/x-www-form-urlencoded&#x60;) of the request from the client application.  | 
**ClientId** | Pointer to **string** | The client ID extracted from &#x60;Authorization&#x60; header of the pushed request from the client application.  | [optional] 
**ClientSecret** | Pointer to **string** | The client secret extracted from &#x60;Authorization&#x60; header of the pushed authorization request from the client application.  | [optional] 
**ClientCertificate** | Pointer to **string** | The client certificate from the MTLS connection to pushed authorization endpoint from the client application. | [optional] 
**ClientCertificatePath** | Pointer to **string** | The certificate path presented by the client during client authentication. These certificates are strings in PEM format.  | [optional] 

## Methods

### NewPushedAuthReqRequest

`func NewPushedAuthReqRequest(parameters string, ) *PushedAuthReqRequest`

NewPushedAuthReqRequest instantiates a new PushedAuthReqRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushedAuthReqRequestWithDefaults

`func NewPushedAuthReqRequestWithDefaults() *PushedAuthReqRequest`

NewPushedAuthReqRequestWithDefaults instantiates a new PushedAuthReqRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *PushedAuthReqRequest) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PushedAuthReqRequest) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PushedAuthReqRequest) SetParameters(v string)`

SetParameters sets Parameters field to given value.


### GetClientId

`func (o *PushedAuthReqRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PushedAuthReqRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PushedAuthReqRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PushedAuthReqRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *PushedAuthReqRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PushedAuthReqRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PushedAuthReqRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PushedAuthReqRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientCertificate

`func (o *PushedAuthReqRequest) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *PushedAuthReqRequest) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *PushedAuthReqRequest) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *PushedAuthReqRequest) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### GetClientCertificatePath

`func (o *PushedAuthReqRequest) GetClientCertificatePath() string`

GetClientCertificatePath returns the ClientCertificatePath field if non-nil, zero value otherwise.

### GetClientCertificatePathOk

`func (o *PushedAuthReqRequest) GetClientCertificatePathOk() (*string, bool)`

GetClientCertificatePathOk returns a tuple with the ClientCertificatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificatePath

`func (o *PushedAuthReqRequest) SetClientCertificatePath(v string)`

SetClientCertificatePath sets ClientCertificatePath field to given value.

### HasClientCertificatePath

`func (o *PushedAuthReqRequest) HasClientCertificatePath() bool`

HasClientCertificatePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


