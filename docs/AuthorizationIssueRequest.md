# AuthorizationIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ticket** | **string** | The ticket issued from Authlete &#x60;/auth/authorization&#x60; API.  | 
**Subject** | **string** | The subject (&#x3D; a user account managed by the service) who has granted authorization to the client application.  | 
**AuthTime** | Pointer to **string** | The time when the authentication of the end-user occurred. Its value is the number of seconds from &#x60;1970-01-01&#x60;.  | [optional] 
**Acr** | Pointer to **string** | The Authentication Context Class Reference performed for the end-user authentication. | [optional] 
**Claims** | Pointer to **string** | The claims of the end-user (&#x3D; pieces of information about the end-user) in JSON format. See [OpenID Connect Core 1.0, 5.1. Standard Claims](https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims) for details about the format.  | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) | Extra properties to associate with an access token and/or an authorization code. | [optional] 
**Scopes** | Pointer to **[]string** | Scopes to associate with an access token and/or an authorization code. If a non-empty string array is given, it replaces the scopes specified by the original authorization request.  | [optional] 
**Sub** | Pointer to **string** | The value of the &#x60;sub&#x60; claim to embed in an ID token. If this request parameter is &#x60;null&#x60; or empty, the value of the &#x60;subject&#x60; request parameter is used as the value of the &#x60;sub&#x60; claim.  | [optional] 
**IdtHeaderParams** | Pointer to **string** | JSON that represents additional JWS header parameters for ID tokens that may be issued based on the authorization request.  | [optional] 
**AuthorizationDetails** | Pointer to [**AuthorizationDetails**](AuthorizationDetails.md) |  | [optional] 
**AccessToken** | Pointer to **string** | The representation of an access token that may be issued as a result of the Authlete API call. | [optional] 

## Methods

### NewAuthorizationIssueRequest

`func NewAuthorizationIssueRequest(ticket string, subject string, ) *AuthorizationIssueRequest`

NewAuthorizationIssueRequest instantiates a new AuthorizationIssueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationIssueRequestWithDefaults

`func NewAuthorizationIssueRequestWithDefaults() *AuthorizationIssueRequest`

NewAuthorizationIssueRequestWithDefaults instantiates a new AuthorizationIssueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicket

`func (o *AuthorizationIssueRequest) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *AuthorizationIssueRequest) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *AuthorizationIssueRequest) SetTicket(v string)`

SetTicket sets Ticket field to given value.


### GetSubject

`func (o *AuthorizationIssueRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AuthorizationIssueRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AuthorizationIssueRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetAuthTime

`func (o *AuthorizationIssueRequest) GetAuthTime() string`

GetAuthTime returns the AuthTime field if non-nil, zero value otherwise.

### GetAuthTimeOk

`func (o *AuthorizationIssueRequest) GetAuthTimeOk() (*string, bool)`

GetAuthTimeOk returns a tuple with the AuthTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTime

`func (o *AuthorizationIssueRequest) SetAuthTime(v string)`

SetAuthTime sets AuthTime field to given value.

### HasAuthTime

`func (o *AuthorizationIssueRequest) HasAuthTime() bool`

HasAuthTime returns a boolean if a field has been set.

### GetAcr

`func (o *AuthorizationIssueRequest) GetAcr() string`

GetAcr returns the Acr field if non-nil, zero value otherwise.

### GetAcrOk

`func (o *AuthorizationIssueRequest) GetAcrOk() (*string, bool)`

GetAcrOk returns a tuple with the Acr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcr

`func (o *AuthorizationIssueRequest) SetAcr(v string)`

SetAcr sets Acr field to given value.

### HasAcr

`func (o *AuthorizationIssueRequest) HasAcr() bool`

HasAcr returns a boolean if a field has been set.

### GetClaims

`func (o *AuthorizationIssueRequest) GetClaims() string`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *AuthorizationIssueRequest) GetClaimsOk() (*string, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *AuthorizationIssueRequest) SetClaims(v string)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *AuthorizationIssueRequest) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetProperties

`func (o *AuthorizationIssueRequest) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AuthorizationIssueRequest) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AuthorizationIssueRequest) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AuthorizationIssueRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetScopes

`func (o *AuthorizationIssueRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AuthorizationIssueRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AuthorizationIssueRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AuthorizationIssueRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSub

`func (o *AuthorizationIssueRequest) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *AuthorizationIssueRequest) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *AuthorizationIssueRequest) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *AuthorizationIssueRequest) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetIdtHeaderParams

`func (o *AuthorizationIssueRequest) GetIdtHeaderParams() string`

GetIdtHeaderParams returns the IdtHeaderParams field if non-nil, zero value otherwise.

### GetIdtHeaderParamsOk

`func (o *AuthorizationIssueRequest) GetIdtHeaderParamsOk() (*string, bool)`

GetIdtHeaderParamsOk returns a tuple with the IdtHeaderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdtHeaderParams

`func (o *AuthorizationIssueRequest) SetIdtHeaderParams(v string)`

SetIdtHeaderParams sets IdtHeaderParams field to given value.

### HasIdtHeaderParams

`func (o *AuthorizationIssueRequest) HasIdtHeaderParams() bool`

HasIdtHeaderParams returns a boolean if a field has been set.

### GetAuthorizationDetails

`func (o *AuthorizationIssueRequest) GetAuthorizationDetails() AuthorizationDetails`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *AuthorizationIssueRequest) GetAuthorizationDetailsOk() (*AuthorizationDetails, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *AuthorizationIssueRequest) SetAuthorizationDetails(v AuthorizationDetails)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *AuthorizationIssueRequest) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.

### GetAccessToken

`func (o *AuthorizationIssueRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AuthorizationIssueRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AuthorizationIssueRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AuthorizationIssueRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


