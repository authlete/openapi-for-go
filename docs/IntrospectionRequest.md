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
**Resources** | Pointer to **[]string** | The resources specified by the &#x60;resource&#x60; request parameters in the token request. See \&quot;Resource Indicators for OAuth 2.0\&quot; for details.  | [optional] 
**AcrValues** | Pointer to **[]string** | Authentication Context Class Reference values one of which the user authentication performed during the course  of issuing the access token must satisfy.  | [optional] 
**MaxAge** | Pointer to **int64** | The maximum authentication age which is the maximum allowable elapsed time since the user authentication  was performed during the course of issuing the access token.  | [optional] 
**RequiredComponents** | Pointer to **[]string** | HTTP Message Components required to be in the signature. If absent, defaults to [ \&quot;@method\&quot;, \&quot;@target-uri\&quot;, \&quot;authorization\&quot; ].  | [optional] 
**Uri** | Pointer to **string** | The full URL of the userinfo endpoint.  | [optional] 
**Message** | Pointer to **string** | The HTTP message body of the request, if present.  | [optional] 
**Headers** | Pointer to [**[]Pair**](Pair.md) | HTTP headers to be included in processing the signature. If this is a signed request, this must include the  Signature and Signature-Input headers, as well as any additional headers covered by the signature.  | [optional] 

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

### GetResources

`func (o *IntrospectionRequest) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *IntrospectionRequest) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *IntrospectionRequest) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *IntrospectionRequest) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetAcrValues

`func (o *IntrospectionRequest) GetAcrValues() []string`

GetAcrValues returns the AcrValues field if non-nil, zero value otherwise.

### GetAcrValuesOk

`func (o *IntrospectionRequest) GetAcrValuesOk() (*[]string, bool)`

GetAcrValuesOk returns a tuple with the AcrValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrValues

`func (o *IntrospectionRequest) SetAcrValues(v []string)`

SetAcrValues sets AcrValues field to given value.

### HasAcrValues

`func (o *IntrospectionRequest) HasAcrValues() bool`

HasAcrValues returns a boolean if a field has been set.

### GetMaxAge

`func (o *IntrospectionRequest) GetMaxAge() int64`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *IntrospectionRequest) GetMaxAgeOk() (*int64, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *IntrospectionRequest) SetMaxAge(v int64)`

SetMaxAge sets MaxAge field to given value.

### HasMaxAge

`func (o *IntrospectionRequest) HasMaxAge() bool`

HasMaxAge returns a boolean if a field has been set.

### GetRequiredComponents

`func (o *IntrospectionRequest) GetRequiredComponents() []string`

GetRequiredComponents returns the RequiredComponents field if non-nil, zero value otherwise.

### GetRequiredComponentsOk

`func (o *IntrospectionRequest) GetRequiredComponentsOk() (*[]string, bool)`

GetRequiredComponentsOk returns a tuple with the RequiredComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredComponents

`func (o *IntrospectionRequest) SetRequiredComponents(v []string)`

SetRequiredComponents sets RequiredComponents field to given value.

### HasRequiredComponents

`func (o *IntrospectionRequest) HasRequiredComponents() bool`

HasRequiredComponents returns a boolean if a field has been set.

### GetUri

`func (o *IntrospectionRequest) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *IntrospectionRequest) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *IntrospectionRequest) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *IntrospectionRequest) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetMessage

`func (o *IntrospectionRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IntrospectionRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IntrospectionRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IntrospectionRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetHeaders

`func (o *IntrospectionRequest) GetHeaders() []Pair`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *IntrospectionRequest) GetHeadersOk() (*[]Pair, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *IntrospectionRequest) SetHeaders(v []Pair)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *IntrospectionRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


