# IntrospectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | An access token to introspect. | 
**Scopes** | Pointer to **[]string** | A string array listing names of scopes which the caller (&#x3D; a protected resource endpoint of the service) requires. When the content type of the request from the service is &#x60;application/x-www-form-urlencoded&#x60;, the format of &#x60;scopes&#x60; is a space-separated list of scope names.  If this parameter is a non-empty array and if it contains a scope which is not covered by the access token,&#x60;action&#x3D;FORBIDDEN&#x60; with &#x60;error&#x3D;insufficient_scope&#x60; is returned from Authlete.  | [optional] 
**Subject** | Pointer to **string** | A subject (&#x3D; a user account managed by the service) whom the caller (&#x3D; a protected resource endpoint of the service) requires.  If this parameter is not &#x60;null&#x60; and if the value does not match the subject who is associated with the access token, &#x60;action&#x3D;FORBIDDEN&#x60; with &#x60;error&#x3D;invalid_request&#x60; is returned from Authlete.  | [optional] 
**ClientCertificate** | Pointer to **string** | Client certificate in PEM format, used to validate binding against access tokens using the TLS client certificate confirmation method.  | [optional] 
**Dpop** | Pointer to **string** | &#x60;DPoP&#x60; header presented by the client during the request to the resource server.  The header contains a signed JWT which includes the public key that is paired with the private key used to sign the JWT. See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details.  | [optional] 
**Htm** | Pointer to **string** | HTTP method of the request from the client to the protected resource endpoint. This field is used to validate the &#x60;DPoP&#x60; header.  See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details.  | [optional] 
**Htu** | Pointer to **string** | URL of the protected resource endpoint. This field is used to validate the &#x60;DPoP&#x60; header.  See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details.  | [optional] 

## Methods

### NewIntrospectionRequest

`func NewIntrospectionRequest(token string, ) *IntrospectionRequest`

NewIntrospectionRequest instantiates a new IntrospectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntrospectionRequestWithDefaults

`func NewIntrospectionRequestWithDefaults() *IntrospectionRequest`

NewIntrospectionRequestWithDefaults instantiates a new IntrospectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *IntrospectionRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *IntrospectionRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *IntrospectionRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetScopes

`func (o *IntrospectionRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IntrospectionRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IntrospectionRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *IntrospectionRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSubject

`func (o *IntrospectionRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *IntrospectionRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *IntrospectionRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *IntrospectionRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetClientCertificate

`func (o *IntrospectionRequest) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *IntrospectionRequest) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *IntrospectionRequest) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *IntrospectionRequest) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### GetDpop

`func (o *IntrospectionRequest) GetDpop() string`

GetDpop returns the Dpop field if non-nil, zero value otherwise.

### GetDpopOk

`func (o *IntrospectionRequest) GetDpopOk() (*string, bool)`

GetDpopOk returns a tuple with the Dpop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpop

`func (o *IntrospectionRequest) SetDpop(v string)`

SetDpop sets Dpop field to given value.

### HasDpop

`func (o *IntrospectionRequest) HasDpop() bool`

HasDpop returns a boolean if a field has been set.

### GetHtm

`func (o *IntrospectionRequest) GetHtm() string`

GetHtm returns the Htm field if non-nil, zero value otherwise.

### GetHtmOk

`func (o *IntrospectionRequest) GetHtmOk() (*string, bool)`

GetHtmOk returns a tuple with the Htm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtm

`func (o *IntrospectionRequest) SetHtm(v string)`

SetHtm sets Htm field to given value.

### HasHtm

`func (o *IntrospectionRequest) HasHtm() bool`

HasHtm returns a boolean if a field has been set.

### GetHtu

`func (o *IntrospectionRequest) GetHtu() string`

GetHtu returns the Htu field if non-nil, zero value otherwise.

### GetHtuOk

`func (o *IntrospectionRequest) GetHtuOk() (*string, bool)`

GetHtuOk returns a tuple with the Htu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtu

`func (o *IntrospectionRequest) SetHtu(v string)`

SetHtu sets Htu field to given value.

### HasHtu

`func (o *IntrospectionRequest) HasHtu() bool`

HasHtu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


