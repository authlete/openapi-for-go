# PushedAuthorizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | **string** | The pushed authorization request body received from the client application.  The value of parameters is the entire entity body (which is formatted in &#x60;application/x-www-form-urlencoded&#x60;) of the request from the client application.  | 
**ClientId** | Pointer to **string** | The client ID extracted from &#x60;Authorization&#x60; header of the pushed request from the client application.  | [optional] 
**ClientSecret** | Pointer to **string** | The client secret extracted from &#x60;Authorization&#x60; header of the pushed authorization request from the client application.  | [optional] 
**ClientCertificate** | Pointer to **string** | The client certificate from the MTLS connection to pushed authorization endpoint from the client application. | [optional] 
**ClientCertificatePath** | Pointer to **string** | The certificate path presented by the client during client authentication. These certificates are strings in PEM format.  | [optional] 
**Dpop** | Pointer to **string** | DPoP Header  | [optional] 
**Htm** | Pointer to **string** | HTTP Method (for DPoP validation).  | [optional] 
**Htu** | Pointer to **string** | HTTP URL base (for DPoP validation). | [optional] 

## Methods

### NewPushedAuthorizationRequest

`func NewPushedAuthorizationRequest(parameters string, ) *PushedAuthorizationRequest`

NewPushedAuthorizationRequest instantiates a new PushedAuthorizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushedAuthorizationRequestWithDefaults

`func NewPushedAuthorizationRequestWithDefaults() *PushedAuthorizationRequest`

NewPushedAuthorizationRequestWithDefaults instantiates a new PushedAuthorizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *PushedAuthorizationRequest) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PushedAuthorizationRequest) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PushedAuthorizationRequest) SetParameters(v string)`

SetParameters sets Parameters field to given value.


### GetClientId

`func (o *PushedAuthorizationRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PushedAuthorizationRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PushedAuthorizationRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PushedAuthorizationRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *PushedAuthorizationRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PushedAuthorizationRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PushedAuthorizationRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PushedAuthorizationRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientCertificate

`func (o *PushedAuthorizationRequest) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *PushedAuthorizationRequest) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *PushedAuthorizationRequest) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *PushedAuthorizationRequest) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### GetClientCertificatePath

`func (o *PushedAuthorizationRequest) GetClientCertificatePath() string`

GetClientCertificatePath returns the ClientCertificatePath field if non-nil, zero value otherwise.

### GetClientCertificatePathOk

`func (o *PushedAuthorizationRequest) GetClientCertificatePathOk() (*string, bool)`

GetClientCertificatePathOk returns a tuple with the ClientCertificatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificatePath

`func (o *PushedAuthorizationRequest) SetClientCertificatePath(v string)`

SetClientCertificatePath sets ClientCertificatePath field to given value.

### HasClientCertificatePath

`func (o *PushedAuthorizationRequest) HasClientCertificatePath() bool`

HasClientCertificatePath returns a boolean if a field has been set.

### GetDpop

`func (o *PushedAuthorizationRequest) GetDpop() string`

GetDpop returns the Dpop field if non-nil, zero value otherwise.

### GetDpopOk

`func (o *PushedAuthorizationRequest) GetDpopOk() (*string, bool)`

GetDpopOk returns a tuple with the Dpop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpop

`func (o *PushedAuthorizationRequest) SetDpop(v string)`

SetDpop sets Dpop field to given value.

### HasDpop

`func (o *PushedAuthorizationRequest) HasDpop() bool`

HasDpop returns a boolean if a field has been set.

### GetHtm

`func (o *PushedAuthorizationRequest) GetHtm() string`

GetHtm returns the Htm field if non-nil, zero value otherwise.

### GetHtmOk

`func (o *PushedAuthorizationRequest) GetHtmOk() (*string, bool)`

GetHtmOk returns a tuple with the Htm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtm

`func (o *PushedAuthorizationRequest) SetHtm(v string)`

SetHtm sets Htm field to given value.

### HasHtm

`func (o *PushedAuthorizationRequest) HasHtm() bool`

HasHtm returns a boolean if a field has been set.

### GetHtu

`func (o *PushedAuthorizationRequest) GetHtu() string`

GetHtu returns the Htu field if non-nil, zero value otherwise.

### GetHtuOk

`func (o *PushedAuthorizationRequest) GetHtuOk() (*string, bool)`

GetHtuOk returns a tuple with the Htu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtu

`func (o *PushedAuthorizationRequest) SetHtu(v string)`

SetHtu sets Htu field to given value.

### HasHtu

`func (o *PushedAuthorizationRequest) HasHtu() bool`

HasHtu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


