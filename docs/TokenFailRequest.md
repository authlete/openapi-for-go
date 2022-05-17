# TokenFailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ticket** | **string** | The ticket issued from Authlete &#x60;/auth/token&#x60; API.  | 
**Reason** | **string** | The reason of the failure of the token request. | 

## Methods

### NewTokenFailRequest

`func NewTokenFailRequest(ticket string, reason string, ) *TokenFailRequest`

NewTokenFailRequest instantiates a new TokenFailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenFailRequestWithDefaults

`func NewTokenFailRequestWithDefaults() *TokenFailRequest`

NewTokenFailRequestWithDefaults instantiates a new TokenFailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicket

`func (o *TokenFailRequest) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *TokenFailRequest) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *TokenFailRequest) SetTicket(v string)`

SetTicket sets Ticket field to given value.


### GetReason

`func (o *TokenFailRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TokenFailRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TokenFailRequest) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


