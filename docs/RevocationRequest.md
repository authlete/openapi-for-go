# RevocationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | **string** | OAuth 2.0 token revocation request parameters which are the request parameters that the OAuth 2.0 token revocation endpoint ([RFC 7009](https://datatracker.ietf.org/doc/html/rfc7009)) of the authorization server implementation received from the client application.  The value of parameters is the entire entity body (which is formatted in &#x60;application/x-www-form-urlencoded&#x60;) of the request from the client application.  | 
**ClientId** | Pointer to **string** | The client ID extracted from &#x60;Authorization&#x60; header of the revocation request from the client application.  If the revocation endpoint of the authorization server implementation supports Basic Authentication as a means of client authentication, and the request from the client application contains its client ID in &#x60;Authorization&#x60; header, the value should be extracted and set to this parameter.  | [optional] 
**ClientSecret** | Pointer to **string** | The client secret extracted from &#x60;Authorization&#x60; header of the revocation request from the client application.  If the revocation endpoint of the authorization server implementation supports basic authentication as a means of client authentication, and the request from the client application contained its client secret in &#x60;Authorization&#x60; header, the value should be extracted and set to this parameter.  | [optional] 
**ClientCertificate** | Pointer to **string** | The client certificate used in the TLS connection between the client application and the revocation endpoint.  | [optional] 
**ClientCertificatePath** | Pointer to **string** | The certificate path presented by the client during client authentication. | [optional] 

## Methods

### NewRevocationRequest

`func NewRevocationRequest(parameters string, ) *RevocationRequest`

NewRevocationRequest instantiates a new RevocationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevocationRequestWithDefaults

`func NewRevocationRequestWithDefaults() *RevocationRequest`

NewRevocationRequestWithDefaults instantiates a new RevocationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *RevocationRequest) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *RevocationRequest) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *RevocationRequest) SetParameters(v string)`

SetParameters sets Parameters field to given value.


### GetClientId

`func (o *RevocationRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *RevocationRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *RevocationRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *RevocationRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *RevocationRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *RevocationRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *RevocationRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *RevocationRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientCertificate

`func (o *RevocationRequest) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *RevocationRequest) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *RevocationRequest) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *RevocationRequest) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### GetClientCertificatePath

`func (o *RevocationRequest) GetClientCertificatePath() string`

GetClientCertificatePath returns the ClientCertificatePath field if non-nil, zero value otherwise.

### GetClientCertificatePathOk

`func (o *RevocationRequest) GetClientCertificatePathOk() (*string, bool)`

GetClientCertificatePathOk returns a tuple with the ClientCertificatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificatePath

`func (o *RevocationRequest) SetClientCertificatePath(v string)`

SetClientCertificatePath sets ClientCertificatePath field to given value.

### HasClientCertificatePath

`func (o *RevocationRequest) HasClientCertificatePath() bool`

HasClientCertificatePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


