# AuthorizationTicketUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ticket** | **string** | The ticket. | 
**Info** | **string** | The information about the ticket. | 

## Methods

### NewAuthorizationTicketUpdateRequest

`func NewAuthorizationTicketUpdateRequest(ticket string, info string, ) *AuthorizationTicketUpdateRequest`

NewAuthorizationTicketUpdateRequest instantiates a new AuthorizationTicketUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationTicketUpdateRequestWithDefaults

`func NewAuthorizationTicketUpdateRequestWithDefaults() *AuthorizationTicketUpdateRequest`

NewAuthorizationTicketUpdateRequestWithDefaults instantiates a new AuthorizationTicketUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicket

`func (o *AuthorizationTicketUpdateRequest) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *AuthorizationTicketUpdateRequest) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *AuthorizationTicketUpdateRequest) SetTicket(v string)`

SetTicket sets Ticket field to given value.


### GetInfo

`func (o *AuthorizationTicketUpdateRequest) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AuthorizationTicketUpdateRequest) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AuthorizationTicketUpdateRequest) SetInfo(v string)`

SetInfo sets Info field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


