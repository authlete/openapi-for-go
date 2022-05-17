# BackchannelAuthenticationFailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ticket** | **string** | The ticket which should be deleted on a call of Authlete&#39;s &#x60;/backchannel/authentication/fail&#x60; API. This request parameter is not mandatory but optional. If this request parameter is given and the ticket belongs to the service, the specified ticket is deleted from the database. Giving this parameter is recommended to clean up the storage area for the service.  | 
**Reason** | **string** | The reason of the failure of the backchannel authentication request. This request parameter is not mandatory but optional. However, giving this parameter is recommended. If omitted, &#x60;SERVER_ERROR&#x60; is used as a reason.  | 
**ErrorDescription** | Pointer to **string** | The description of the error. This corresponds to the &#x60;error_description&#x60; property in the response to the client.  | [optional] 
**ErrorUri** | Pointer to **string** | The URI of a document which describes the error in detail. If this optional request parameter is given, its value is used as the value of the &#x60;error_uri&#x60; property. | [optional] 

## Methods

### NewBackchannelAuthenticationFailRequest

`func NewBackchannelAuthenticationFailRequest(ticket string, reason string, ) *BackchannelAuthenticationFailRequest`

NewBackchannelAuthenticationFailRequest instantiates a new BackchannelAuthenticationFailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackchannelAuthenticationFailRequestWithDefaults

`func NewBackchannelAuthenticationFailRequestWithDefaults() *BackchannelAuthenticationFailRequest`

NewBackchannelAuthenticationFailRequestWithDefaults instantiates a new BackchannelAuthenticationFailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicket

`func (o *BackchannelAuthenticationFailRequest) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *BackchannelAuthenticationFailRequest) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *BackchannelAuthenticationFailRequest) SetTicket(v string)`

SetTicket sets Ticket field to given value.


### GetReason

`func (o *BackchannelAuthenticationFailRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BackchannelAuthenticationFailRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BackchannelAuthenticationFailRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetErrorDescription

`func (o *BackchannelAuthenticationFailRequest) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *BackchannelAuthenticationFailRequest) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *BackchannelAuthenticationFailRequest) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *BackchannelAuthenticationFailRequest) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorUri

`func (o *BackchannelAuthenticationFailRequest) GetErrorUri() string`

GetErrorUri returns the ErrorUri field if non-nil, zero value otherwise.

### GetErrorUriOk

`func (o *BackchannelAuthenticationFailRequest) GetErrorUriOk() (*string, bool)`

GetErrorUriOk returns a tuple with the ErrorUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUri

`func (o *BackchannelAuthenticationFailRequest) SetErrorUri(v string)`

SetErrorUri sets ErrorUri field to given value.

### HasErrorUri

`func (o *BackchannelAuthenticationFailRequest) HasErrorUri() bool`

HasErrorUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


