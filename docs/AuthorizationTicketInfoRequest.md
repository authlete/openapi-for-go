# AuthorizationTicketInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ticket** | **string** | The ticket that has been issued from the &#x60;/auth/authorization&#x60; API. | 

## Methods

### NewAuthorizationTicketInfoRequest

`func NewAuthorizationTicketInfoRequest(ticket string, ) *AuthorizationTicketInfoRequest`

NewAuthorizationTicketInfoRequest instantiates a new AuthorizationTicketInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationTicketInfoRequestWithDefaults

`func NewAuthorizationTicketInfoRequestWithDefaults() *AuthorizationTicketInfoRequest`

NewAuthorizationTicketInfoRequestWithDefaults instantiates a new AuthorizationTicketInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicket

`func (o *AuthorizationTicketInfoRequest) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *AuthorizationTicketInfoRequest) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *AuthorizationTicketInfoRequest) SetTicket(v string)`

SetTicket sets Ticket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


