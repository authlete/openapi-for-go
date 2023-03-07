# GMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | An access token to introspect. | [optional] 
**Scopes** | Pointer to **[]string** | A string array listing names of scopes which the caller (&#x3D; a protected resource endpoint of the service) requires. When the content type of the request from the service is &#x60;application/x-www-form-urlencoded&#x60;, the format of &#x60;scopes&#x60; is a space-separated list of scope names.  If this parameter is a non-empty array and if it contains a scope which is not covered by the access token,&#x60;action&#x3D;FORBIDDEN&#x60; with &#x60;error&#x3D;insufficient_scope&#x60; is returned from Authlete.  | [optional] 
**Subject** | Pointer to **string** | A subject (&#x3D; a user account managed by the service) whom the caller (&#x3D; a protected resource endpoint of the service) requires.  If this parameter is not &#x60;null&#x60; and if the value does not match the subject who is associated with the access token, &#x60;action&#x3D;FORBIDDEN&#x60; with &#x60;error&#x3D;invalid_request&#x60; is returned from Authlete.  | [optional] 
**ClientCertificate** | Pointer to **string** | Client certificate in PEM format, used to validate binding against access tokens using the TLS client certificate confirmation method.  | [optional] 
**Dpop** | Pointer to **string** | &#x60;DPoP&#x60; header presented by the client during the request to the resource server.  The header contains a signed JWT which includes the public key that is paired with the private key used to sign the JWT. See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details.  | [optional] 
**Htm** | Pointer to **string** | HTTP method of the request from the client to the protected resource endpoint. This field is used to validate the &#x60;DPoP&#x60; header.  See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details.  | [optional] 
**Htu** | Pointer to **string** | URL of the protected resource endpoint. This field is used to validate the &#x60;DPoP&#x60; header.  See [OAuth 2.0 Demonstration of Proof-of-Possession at the Application Layer (DPoP)](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-dpop) for details.  | [optional] 
**Resources** | Pointer to **[]string** | The resources specified by the &#x60;resource&#x60; request parameters in the token request. See \&quot;Resource Indicators for OAuth 2.0\&quot; for details.  | [optional] 
**GmAction** | Pointer to [**GrantManagementAction**](GrantManagementAction.md) |  | [optional] 
**GrantId** | Pointer to **string** | The value of the &#x60;grant_id&#x60; request parameter of the device authorization request.  The &#x60;grant_id&#x60; request parameter is defined in [Grant Management for OAuth 2.0](https://openid.net/specs/fapi-grant-management.html) , which is supported by Authlete 2.3 and newer versions.  | [optional] 

## Methods

### NewGMRequest

`func NewGMRequest() *GMRequest`

NewGMRequest instantiates a new GMRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGMRequestWithDefaults

`func NewGMRequestWithDefaults() *GMRequest`

NewGMRequestWithDefaults instantiates a new GMRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *GMRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *GMRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *GMRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *GMRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetScopes

`func (o *GMRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GMRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GMRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *GMRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSubject

`func (o *GMRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GMRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GMRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GMRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetClientCertificate

`func (o *GMRequest) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *GMRequest) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *GMRequest) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *GMRequest) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### GetDpop

`func (o *GMRequest) GetDpop() string`

GetDpop returns the Dpop field if non-nil, zero value otherwise.

### GetDpopOk

`func (o *GMRequest) GetDpopOk() (*string, bool)`

GetDpopOk returns a tuple with the Dpop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpop

`func (o *GMRequest) SetDpop(v string)`

SetDpop sets Dpop field to given value.

### HasDpop

`func (o *GMRequest) HasDpop() bool`

HasDpop returns a boolean if a field has been set.

### GetHtm

`func (o *GMRequest) GetHtm() string`

GetHtm returns the Htm field if non-nil, zero value otherwise.

### GetHtmOk

`func (o *GMRequest) GetHtmOk() (*string, bool)`

GetHtmOk returns a tuple with the Htm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtm

`func (o *GMRequest) SetHtm(v string)`

SetHtm sets Htm field to given value.

### HasHtm

`func (o *GMRequest) HasHtm() bool`

HasHtm returns a boolean if a field has been set.

### GetHtu

`func (o *GMRequest) GetHtu() string`

GetHtu returns the Htu field if non-nil, zero value otherwise.

### GetHtuOk

`func (o *GMRequest) GetHtuOk() (*string, bool)`

GetHtuOk returns a tuple with the Htu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtu

`func (o *GMRequest) SetHtu(v string)`

SetHtu sets Htu field to given value.

### HasHtu

`func (o *GMRequest) HasHtu() bool`

HasHtu returns a boolean if a field has been set.

### GetResources

`func (o *GMRequest) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GMRequest) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GMRequest) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *GMRequest) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetGmAction

`func (o *GMRequest) GetGmAction() GrantManagementAction`

GetGmAction returns the GmAction field if non-nil, zero value otherwise.

### GetGmActionOk

`func (o *GMRequest) GetGmActionOk() (*GrantManagementAction, bool)`

GetGmActionOk returns a tuple with the GmAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmAction

`func (o *GMRequest) SetGmAction(v GrantManagementAction)`

SetGmAction sets GmAction field to given value.

### HasGmAction

`func (o *GMRequest) HasGmAction() bool`

HasGmAction returns a boolean if a field has been set.

### GetGrantId

`func (o *GMRequest) GetGrantId() string`

GetGrantId returns the GrantId field if non-nil, zero value otherwise.

### GetGrantIdOk

`func (o *GMRequest) GetGrantIdOk() (*string, bool)`

GetGrantIdOk returns a tuple with the GrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantId

`func (o *GMRequest) SetGrantId(v string)`

SetGrantId sets GrantId field to given value.

### HasGrantId

`func (o *GMRequest) HasGrantId() bool`

HasGrantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


