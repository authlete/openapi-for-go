# TokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | **string** | OAuth 2.0 token request parameters which are the request parameters that the OAuth 2.0 token endpoint of the authorization server implementation received from the client application.  The value of parameters is the entire entity body (which is formatted in &#x60;application/x-www-form-urlencoded&#x60;) of the request from the client application.  | 
**ClientId** | Pointer to **string** | The client ID extracted from &#x60;Authorization&#x60; header of the token request from the client application.  If the token endpoint of the authorization server implementation supports basic authentication as a means of client authentication, and the request from the client application contained its client ID in &#x60;Authorization&#x60; header, the value should be extracted and set to this parameter.  | [optional] 
**ClientSecret** | Pointer to **string** | The client secret extracted from &#x60;Authorization&#x60; header of the token request from the client application.  If the token endpoint of the authorization server implementation supports basic authentication as a means of client authentication, and the request from the client application contained its client secret in &#x60;Authorization&#x60; header, the value should be extracted and set to this parameter.  | [optional] 
**ClientCertificate** | Pointer to **string** | The client certificate from the MTLS of the token request from the client application. | [optional] 
**ClientCertificatePath** | Pointer to **string** | The certificate path presented by the client during client authentication. These certificates are strings in PEM format.  | [optional] 
**Properties** | Pointer to **string** | Extra properties to associate with an access token. See [Extra Properties](https://www.authlete.com/developers/definitive_guide/extra_properties/) for details.  | [optional] 
**Dpop** | Pointer to **string** | &#x60;DPoP&#x60; header presented by the client during the request to the token endpoint.  The header contains a signed JWT which includes the public key that is paired with the private key used to sign the JWT. See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details.  | [optional] 
**Htm** | Pointer to **string** | HTTP method of the token request. This field is used to validate the &#x60;DPoP&#x60; header.  In normal cases, the value is &#x60;POST&#x60;. When this parameter is omitted, &#x60;POST&#x60; is used as the default value. See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details.  | [optional] 
**Htu** | Pointer to **string** | URL of the token endpoint. This field is used to validate the &#x60;DPoP&#x60; header.  If this parameter is omitted, the &#x60;tokenEndpoint&#x60; property of the Service is used as the default value. See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details.  | [optional] 
**AccessToken** | Pointer to **string** | The representation of an access token that may be issued as a result of the Authlete API call.  | [optional] 
**JwtAtClaims** | Pointer to **string** | Additional claims that are added to the payload part of the JWT access token. | [optional] 

## Methods

### NewTokenRequest

`func NewTokenRequest(parameters string, ) *TokenRequest`

NewTokenRequest instantiates a new TokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRequestWithDefaults

`func NewTokenRequestWithDefaults() *TokenRequest`

NewTokenRequestWithDefaults instantiates a new TokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *TokenRequest) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TokenRequest) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TokenRequest) SetParameters(v string)`

SetParameters sets Parameters field to given value.


### GetClientId

`func (o *TokenRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TokenRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TokenRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TokenRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *TokenRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *TokenRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *TokenRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *TokenRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientCertificate

`func (o *TokenRequest) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *TokenRequest) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *TokenRequest) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *TokenRequest) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### GetClientCertificatePath

`func (o *TokenRequest) GetClientCertificatePath() string`

GetClientCertificatePath returns the ClientCertificatePath field if non-nil, zero value otherwise.

### GetClientCertificatePathOk

`func (o *TokenRequest) GetClientCertificatePathOk() (*string, bool)`

GetClientCertificatePathOk returns a tuple with the ClientCertificatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificatePath

`func (o *TokenRequest) SetClientCertificatePath(v string)`

SetClientCertificatePath sets ClientCertificatePath field to given value.

### HasClientCertificatePath

`func (o *TokenRequest) HasClientCertificatePath() bool`

HasClientCertificatePath returns a boolean if a field has been set.

### GetProperties

`func (o *TokenRequest) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TokenRequest) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TokenRequest) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TokenRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetDpop

`func (o *TokenRequest) GetDpop() string`

GetDpop returns the Dpop field if non-nil, zero value otherwise.

### GetDpopOk

`func (o *TokenRequest) GetDpopOk() (*string, bool)`

GetDpopOk returns a tuple with the Dpop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpop

`func (o *TokenRequest) SetDpop(v string)`

SetDpop sets Dpop field to given value.

### HasDpop

`func (o *TokenRequest) HasDpop() bool`

HasDpop returns a boolean if a field has been set.

### GetHtm

`func (o *TokenRequest) GetHtm() string`

GetHtm returns the Htm field if non-nil, zero value otherwise.

### GetHtmOk

`func (o *TokenRequest) GetHtmOk() (*string, bool)`

GetHtmOk returns a tuple with the Htm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtm

`func (o *TokenRequest) SetHtm(v string)`

SetHtm sets Htm field to given value.

### HasHtm

`func (o *TokenRequest) HasHtm() bool`

HasHtm returns a boolean if a field has been set.

### GetHtu

`func (o *TokenRequest) GetHtu() string`

GetHtu returns the Htu field if non-nil, zero value otherwise.

### GetHtuOk

`func (o *TokenRequest) GetHtuOk() (*string, bool)`

GetHtuOk returns a tuple with the Htu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtu

`func (o *TokenRequest) SetHtu(v string)`

SetHtu sets Htu field to given value.

### HasHtu

`func (o *TokenRequest) HasHtu() bool`

HasHtu returns a boolean if a field has been set.

### GetAccessToken

`func (o *TokenRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TokenRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TokenRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *TokenRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetJwtAtClaims

`func (o *TokenRequest) GetJwtAtClaims() string`

GetJwtAtClaims returns the JwtAtClaims field if non-nil, zero value otherwise.

### GetJwtAtClaimsOk

`func (o *TokenRequest) GetJwtAtClaimsOk() (*string, bool)`

GetJwtAtClaimsOk returns a tuple with the JwtAtClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtAtClaims

`func (o *TokenRequest) SetJwtAtClaims(v string)`

SetJwtAtClaims sets JwtAtClaims field to given value.

### HasJwtAtClaims

`func (o *TokenRequest) HasJwtAtClaims() bool`

HasJwtAtClaims returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


