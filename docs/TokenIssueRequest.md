# TokenIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ticket** | **string** | The ticket issued from Authlete &#x60;/auth/token&#x60; API.  | 
**Subject** | **string** | The subject (&#x3D; unique identifier) of the authenticated user.  | 
**Properties** | Pointer to [**[]Property**](Property.md) | Extra properties to associate with a newly created access token. Note that properties parameter is accepted only when &#x60;Content-Type&#x60; of the request is &#x60;application/json&#x60;, so don&#39;t use &#x60;application/x-www-form-urlencoded&#x60; if you want to specify properties.  | [optional] 
**AccessToken** | Pointer to **string** | The representation of an access token that may be issued as a result of the Authlete API call. | [optional] 

## Methods

### NewTokenIssueRequest

`func NewTokenIssueRequest(ticket string, subject string, ) *TokenIssueRequest`

NewTokenIssueRequest instantiates a new TokenIssueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenIssueRequestWithDefaults

`func NewTokenIssueRequestWithDefaults() *TokenIssueRequest`

NewTokenIssueRequestWithDefaults instantiates a new TokenIssueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicket

`func (o *TokenIssueRequest) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *TokenIssueRequest) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *TokenIssueRequest) SetTicket(v string)`

SetTicket sets Ticket field to given value.


### GetSubject

`func (o *TokenIssueRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TokenIssueRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TokenIssueRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetProperties

`func (o *TokenIssueRequest) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TokenIssueRequest) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TokenIssueRequest) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TokenIssueRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetAccessToken

`func (o *TokenIssueRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TokenIssueRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TokenIssueRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *TokenIssueRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


