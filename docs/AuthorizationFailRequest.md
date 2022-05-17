# AuthorizationFailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ticket** | **string** | The ticket issued from Authlete &#x60;/auth/authorization&#x60; API.  | 
**Reason** | **string** | The reason of the failure of the authorization request. For more details, see [NO_INTERACTION] in the description of &#x60;/auth/authorization&#x60; API.  | 
**Description** | Pointer to **string** | The custom description about the authorization failure. | [optional] 

## Methods

### NewAuthorizationFailRequest

`func NewAuthorizationFailRequest(ticket string, reason string, ) *AuthorizationFailRequest`

NewAuthorizationFailRequest instantiates a new AuthorizationFailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationFailRequestWithDefaults

`func NewAuthorizationFailRequestWithDefaults() *AuthorizationFailRequest`

NewAuthorizationFailRequestWithDefaults instantiates a new AuthorizationFailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicket

`func (o *AuthorizationFailRequest) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *AuthorizationFailRequest) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *AuthorizationFailRequest) SetTicket(v string)`

SetTicket sets Ticket field to given value.


### GetReason

`func (o *AuthorizationFailRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AuthorizationFailRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AuthorizationFailRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetDescription

`func (o *AuthorizationFailRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthorizationFailRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthorizationFailRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthorizationFailRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


