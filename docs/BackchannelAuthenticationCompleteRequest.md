# BackchannelAuthenticationCompleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ticket** | **string** | The ticket issued by Authlete&#39;s &#x60;/backchannel/authentication&#x60; API.  | 
**Result** | **string** | The result of the end-user authentication and authorization. One of the following. Details are described in the description.  | 
**Subject** | **string** | The subject (&#x3D; unique identifier) of the end-user.  | 
**Sub** | Pointer to **string** | The value of the sub claim that should be used in the ID token.  | [optional] 
**AuthTime** | Pointer to **int64** | The time at which the end-user was authenticated. Its value is the number of seconds from &#x60;1970-01-01&#x60;.  | [optional] 
**Acr** | Pointer to **string** | The reference of the authentication context class which the end-user authentication satisfied.  | [optional] 
**Claims** | Pointer to **string** | Additional claims which will be embedded in the ID token.  | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) | The extra properties associated with the access token.  | [optional] 
**Scopes** | Pointer to **[]string** | Scopes to replace the scopes specified in the original backchannel authentication request with. When nothing is specified for this parameter, replacement is not performed.  | [optional] 
**IdtHeaderParams** | Pointer to **string** | JSON that represents additional JWS header parameters for ID tokens.  | [optional] 
**ErrorDescription** | Pointer to **string** | The description of the error. If this optional request parameter is given, its value is used as the value of the &#x60;error_description&#x60; property, but it is used only when the result is not &#x60;AUTHORIZED&#x60;. To comply with the specification strictly, the description must not include characters outside the set &#x60;%x20-21 / %x23-5B / %x5D-7E&#x60;.  | [optional] 
**ErrorUri** | Pointer to **string** | The URI of a document which describes the error in detail. This corresponds to the &#x60;error_uri&#x60; property in the response to the client.  | [optional] 
**ConsentedClaims** | Pointer to **[]string** | the claims that the user has consented for the client application to know.  | [optional] 
**JwtAtClaims** | Pointer to **string** | Additional claims that are added to the payload part of the JWT access token.  | [optional] 
**AccessToken** | Pointer to **string** | The representation of an access token that may be issued as a result of the Authlete API call.  | [optional] 

## Methods

### NewBackchannelAuthenticationCompleteRequest

`func NewBackchannelAuthenticationCompleteRequest(ticket string, result string, subject string, ) *BackchannelAuthenticationCompleteRequest`

NewBackchannelAuthenticationCompleteRequest instantiates a new BackchannelAuthenticationCompleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackchannelAuthenticationCompleteRequestWithDefaults

`func NewBackchannelAuthenticationCompleteRequestWithDefaults() *BackchannelAuthenticationCompleteRequest`

NewBackchannelAuthenticationCompleteRequestWithDefaults instantiates a new BackchannelAuthenticationCompleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicket

`func (o *BackchannelAuthenticationCompleteRequest) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *BackchannelAuthenticationCompleteRequest) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *BackchannelAuthenticationCompleteRequest) SetTicket(v string)`

SetTicket sets Ticket field to given value.


### GetResult

`func (o *BackchannelAuthenticationCompleteRequest) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BackchannelAuthenticationCompleteRequest) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BackchannelAuthenticationCompleteRequest) SetResult(v string)`

SetResult sets Result field to given value.


### GetSubject

`func (o *BackchannelAuthenticationCompleteRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *BackchannelAuthenticationCompleteRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *BackchannelAuthenticationCompleteRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetSub

`func (o *BackchannelAuthenticationCompleteRequest) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *BackchannelAuthenticationCompleteRequest) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *BackchannelAuthenticationCompleteRequest) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *BackchannelAuthenticationCompleteRequest) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetAuthTime

`func (o *BackchannelAuthenticationCompleteRequest) GetAuthTime() int64`

GetAuthTime returns the AuthTime field if non-nil, zero value otherwise.

### GetAuthTimeOk

`func (o *BackchannelAuthenticationCompleteRequest) GetAuthTimeOk() (*int64, bool)`

GetAuthTimeOk returns a tuple with the AuthTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTime

`func (o *BackchannelAuthenticationCompleteRequest) SetAuthTime(v int64)`

SetAuthTime sets AuthTime field to given value.

### HasAuthTime

`func (o *BackchannelAuthenticationCompleteRequest) HasAuthTime() bool`

HasAuthTime returns a boolean if a field has been set.

### GetAcr

`func (o *BackchannelAuthenticationCompleteRequest) GetAcr() string`

GetAcr returns the Acr field if non-nil, zero value otherwise.

### GetAcrOk

`func (o *BackchannelAuthenticationCompleteRequest) GetAcrOk() (*string, bool)`

GetAcrOk returns a tuple with the Acr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcr

`func (o *BackchannelAuthenticationCompleteRequest) SetAcr(v string)`

SetAcr sets Acr field to given value.

### HasAcr

`func (o *BackchannelAuthenticationCompleteRequest) HasAcr() bool`

HasAcr returns a boolean if a field has been set.

### GetClaims

`func (o *BackchannelAuthenticationCompleteRequest) GetClaims() string`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *BackchannelAuthenticationCompleteRequest) GetClaimsOk() (*string, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *BackchannelAuthenticationCompleteRequest) SetClaims(v string)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *BackchannelAuthenticationCompleteRequest) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetProperties

`func (o *BackchannelAuthenticationCompleteRequest) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BackchannelAuthenticationCompleteRequest) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BackchannelAuthenticationCompleteRequest) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BackchannelAuthenticationCompleteRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetScopes

`func (o *BackchannelAuthenticationCompleteRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *BackchannelAuthenticationCompleteRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *BackchannelAuthenticationCompleteRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *BackchannelAuthenticationCompleteRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetIdtHeaderParams

`func (o *BackchannelAuthenticationCompleteRequest) GetIdtHeaderParams() string`

GetIdtHeaderParams returns the IdtHeaderParams field if non-nil, zero value otherwise.

### GetIdtHeaderParamsOk

`func (o *BackchannelAuthenticationCompleteRequest) GetIdtHeaderParamsOk() (*string, bool)`

GetIdtHeaderParamsOk returns a tuple with the IdtHeaderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdtHeaderParams

`func (o *BackchannelAuthenticationCompleteRequest) SetIdtHeaderParams(v string)`

SetIdtHeaderParams sets IdtHeaderParams field to given value.

### HasIdtHeaderParams

`func (o *BackchannelAuthenticationCompleteRequest) HasIdtHeaderParams() bool`

HasIdtHeaderParams returns a boolean if a field has been set.

### GetErrorDescription

`func (o *BackchannelAuthenticationCompleteRequest) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *BackchannelAuthenticationCompleteRequest) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *BackchannelAuthenticationCompleteRequest) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *BackchannelAuthenticationCompleteRequest) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorUri

`func (o *BackchannelAuthenticationCompleteRequest) GetErrorUri() string`

GetErrorUri returns the ErrorUri field if non-nil, zero value otherwise.

### GetErrorUriOk

`func (o *BackchannelAuthenticationCompleteRequest) GetErrorUriOk() (*string, bool)`

GetErrorUriOk returns a tuple with the ErrorUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUri

`func (o *BackchannelAuthenticationCompleteRequest) SetErrorUri(v string)`

SetErrorUri sets ErrorUri field to given value.

### HasErrorUri

`func (o *BackchannelAuthenticationCompleteRequest) HasErrorUri() bool`

HasErrorUri returns a boolean if a field has been set.

### GetConsentedClaims

`func (o *BackchannelAuthenticationCompleteRequest) GetConsentedClaims() []string`

GetConsentedClaims returns the ConsentedClaims field if non-nil, zero value otherwise.

### GetConsentedClaimsOk

`func (o *BackchannelAuthenticationCompleteRequest) GetConsentedClaimsOk() (*[]string, bool)`

GetConsentedClaimsOk returns a tuple with the ConsentedClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentedClaims

`func (o *BackchannelAuthenticationCompleteRequest) SetConsentedClaims(v []string)`

SetConsentedClaims sets ConsentedClaims field to given value.

### HasConsentedClaims

`func (o *BackchannelAuthenticationCompleteRequest) HasConsentedClaims() bool`

HasConsentedClaims returns a boolean if a field has been set.

### GetJwtAtClaims

`func (o *BackchannelAuthenticationCompleteRequest) GetJwtAtClaims() string`

GetJwtAtClaims returns the JwtAtClaims field if non-nil, zero value otherwise.

### GetJwtAtClaimsOk

`func (o *BackchannelAuthenticationCompleteRequest) GetJwtAtClaimsOk() (*string, bool)`

GetJwtAtClaimsOk returns a tuple with the JwtAtClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtAtClaims

`func (o *BackchannelAuthenticationCompleteRequest) SetJwtAtClaims(v string)`

SetJwtAtClaims sets JwtAtClaims field to given value.

### HasJwtAtClaims

`func (o *BackchannelAuthenticationCompleteRequest) HasJwtAtClaims() bool`

HasJwtAtClaims returns a boolean if a field has been set.

### GetAccessToken

`func (o *BackchannelAuthenticationCompleteRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *BackchannelAuthenticationCompleteRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *BackchannelAuthenticationCompleteRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *BackchannelAuthenticationCompleteRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


