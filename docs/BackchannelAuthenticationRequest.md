# BackchannelAuthenticationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | **string** | Parameters of a backchannel authentication request which are the request parameters that the backchannel authentication endpoint of the OpenID provider implementation received from the client application.  The value of &#x60;parameters&#x60; is the entire entity body (which is formatted in &#x60;application/x-www-form-urlencoded&#x60;) of the request from the client application.  | 
**ClientId** | Pointer to **string** | The client ID extracted from Authorization header of the backchannel authentication request from the client application.  If the backchannel authentication endpoint of the OpenID provider implementation supports Basic Authentication as a means of client authentication, and the request from the client application contained its client ID in Authorization header, the value should be extracted and set to this parameter.  | [optional] 
**ClientSecret** | Pointer to **string** | The client secret extracted from Authorization header of the backchannel authentication request from the client application.  If the backchannel authentication endpoint of the OpenID provider implementation supports Basic Authentication as a means of client authentication, and the request from the client application contained its client secret in Authorization header, the value should be extracted and set to this parameter.  | [optional] 
**ClientCertificate** | Pointer to **string** | The client certification used in the TLS connection between the client application and the backchannel authentication endpoint of the OpenID provider.  | [optional] 
**ClientCertificatePath** | Pointer to **string** | The client certificate path presented by the client during client authentication. Each element is a string in PEM format. | [optional] 

## Methods

### NewBackchannelAuthenticationRequest

`func NewBackchannelAuthenticationRequest(parameters string, ) *BackchannelAuthenticationRequest`

NewBackchannelAuthenticationRequest instantiates a new BackchannelAuthenticationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackchannelAuthenticationRequestWithDefaults

`func NewBackchannelAuthenticationRequestWithDefaults() *BackchannelAuthenticationRequest`

NewBackchannelAuthenticationRequestWithDefaults instantiates a new BackchannelAuthenticationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *BackchannelAuthenticationRequest) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *BackchannelAuthenticationRequest) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *BackchannelAuthenticationRequest) SetParameters(v string)`

SetParameters sets Parameters field to given value.


### GetClientId

`func (o *BackchannelAuthenticationRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BackchannelAuthenticationRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BackchannelAuthenticationRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BackchannelAuthenticationRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *BackchannelAuthenticationRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *BackchannelAuthenticationRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *BackchannelAuthenticationRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *BackchannelAuthenticationRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientCertificate

`func (o *BackchannelAuthenticationRequest) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *BackchannelAuthenticationRequest) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *BackchannelAuthenticationRequest) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *BackchannelAuthenticationRequest) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### GetClientCertificatePath

`func (o *BackchannelAuthenticationRequest) GetClientCertificatePath() string`

GetClientCertificatePath returns the ClientCertificatePath field if non-nil, zero value otherwise.

### GetClientCertificatePathOk

`func (o *BackchannelAuthenticationRequest) GetClientCertificatePathOk() (*string, bool)`

GetClientCertificatePathOk returns a tuple with the ClientCertificatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificatePath

`func (o *BackchannelAuthenticationRequest) SetClientCertificatePath(v string)`

SetClientCertificatePath sets ClientCertificatePath field to given value.

### HasClientCertificatePath

`func (o *BackchannelAuthenticationRequest) HasClientCertificatePath() bool`

HasClientCertificatePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


