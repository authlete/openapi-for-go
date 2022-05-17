# BackchannelAuthenticationIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ticket** | **string** | The ticket issued from Authlete&#39;s &#x60;/backchannel/authentication&#x60; API. | 

## Methods

### NewBackchannelAuthenticationIssueRequest

`func NewBackchannelAuthenticationIssueRequest(ticket string, ) *BackchannelAuthenticationIssueRequest`

NewBackchannelAuthenticationIssueRequest instantiates a new BackchannelAuthenticationIssueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackchannelAuthenticationIssueRequestWithDefaults

`func NewBackchannelAuthenticationIssueRequestWithDefaults() *BackchannelAuthenticationIssueRequest`

NewBackchannelAuthenticationIssueRequestWithDefaults instantiates a new BackchannelAuthenticationIssueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicket

`func (o *BackchannelAuthenticationIssueRequest) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *BackchannelAuthenticationIssueRequest) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *BackchannelAuthenticationIssueRequest) SetTicket(v string)`

SetTicket sets Ticket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


