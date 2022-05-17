# UserinfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | An access token.  | 
**ClientCertificate** | Pointer to **string** | Client certificate used in the TLS connection established between the client application and the userinfo endpoint.  The value of this request parameter is referred to when the access token given to the userinfo endpoint was bound to a client certificate when it was issued. See [OAuth 2.0 Mutual TLS Client Authentication and Certificate-Bound Access Tokens] (https://datatracker.ietf.org/doc/rfc8705/) for details about the specification of certificate-bound access tokens.  | [optional] 
**Dpop** | Pointer to **string** | &#x60;DPoP&#x60; header presented by the client during the request to the user info endpoint.  The header contains a signed JWT which includes the public key that is paired with the private key used to sign the JWT. See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details.  | [optional] 
**Htm** | Pointer to **string** | HTTP method of the user info request. This field is used to validate the DPoP header. In normal cases, the value is either &#x60;GET&#x60; or &#x60;POST&#x60;.  | [optional] 
**Htu** | Pointer to **string** | URL of the user info endpoint. This field is used to validate the DPoP header.  If this parameter is omitted, the &#x60;userInfoEndpoint&#x60; property of the service is used as the default value. See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details. | [optional] 

## Methods

### NewUserinfoRequest

`func NewUserinfoRequest(token string, ) *UserinfoRequest`

NewUserinfoRequest instantiates a new UserinfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserinfoRequestWithDefaults

`func NewUserinfoRequestWithDefaults() *UserinfoRequest`

NewUserinfoRequestWithDefaults instantiates a new UserinfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *UserinfoRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UserinfoRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UserinfoRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetClientCertificate

`func (o *UserinfoRequest) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *UserinfoRequest) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *UserinfoRequest) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *UserinfoRequest) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### GetDpop

`func (o *UserinfoRequest) GetDpop() string`

GetDpop returns the Dpop field if non-nil, zero value otherwise.

### GetDpopOk

`func (o *UserinfoRequest) GetDpopOk() (*string, bool)`

GetDpopOk returns a tuple with the Dpop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpop

`func (o *UserinfoRequest) SetDpop(v string)`

SetDpop sets Dpop field to given value.

### HasDpop

`func (o *UserinfoRequest) HasDpop() bool`

HasDpop returns a boolean if a field has been set.

### GetHtm

`func (o *UserinfoRequest) GetHtm() string`

GetHtm returns the Htm field if non-nil, zero value otherwise.

### GetHtmOk

`func (o *UserinfoRequest) GetHtmOk() (*string, bool)`

GetHtmOk returns a tuple with the Htm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtm

`func (o *UserinfoRequest) SetHtm(v string)`

SetHtm sets Htm field to given value.

### HasHtm

`func (o *UserinfoRequest) HasHtm() bool`

HasHtm returns a boolean if a field has been set.

### GetHtu

`func (o *UserinfoRequest) GetHtu() string`

GetHtu returns the Htu field if non-nil, zero value otherwise.

### GetHtuOk

`func (o *UserinfoRequest) GetHtuOk() (*string, bool)`

GetHtuOk returns a tuple with the Htu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtu

`func (o *UserinfoRequest) SetHtu(v string)`

SetHtu sets Htu field to given value.

### HasHtu

`func (o *UserinfoRequest) HasHtu() bool`

HasHtu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


