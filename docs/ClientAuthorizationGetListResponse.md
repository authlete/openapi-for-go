# ClientAuthorizationGetListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **int32** | Start index of search results (inclusive).  | [optional] 
**End** | Pointer to **int32** | End index of search results (exclusive).  | [optional] 
**Developer** | Pointer to **string** | Unique ID of a client developer.  | [optional] 
**Subject** | Pointer to **string** | Unique user ID of an end-user.  | [optional] 
**TotalCount** | Pointer to **int32** | Unique ID of a client developer.  | [optional] 
**Clients** | Pointer to [**[]Client**](Client.md) | An array of clients.  | [optional] 

## Methods

### NewClientAuthorizationGetListResponse

`func NewClientAuthorizationGetListResponse() *ClientAuthorizationGetListResponse`

NewClientAuthorizationGetListResponse instantiates a new ClientAuthorizationGetListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAuthorizationGetListResponseWithDefaults

`func NewClientAuthorizationGetListResponseWithDefaults() *ClientAuthorizationGetListResponse`

NewClientAuthorizationGetListResponseWithDefaults instantiates a new ClientAuthorizationGetListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *ClientAuthorizationGetListResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ClientAuthorizationGetListResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ClientAuthorizationGetListResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ClientAuthorizationGetListResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *ClientAuthorizationGetListResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ClientAuthorizationGetListResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ClientAuthorizationGetListResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ClientAuthorizationGetListResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetDeveloper

`func (o *ClientAuthorizationGetListResponse) GetDeveloper() string`

GetDeveloper returns the Developer field if non-nil, zero value otherwise.

### GetDeveloperOk

`func (o *ClientAuthorizationGetListResponse) GetDeveloperOk() (*string, bool)`

GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloper

`func (o *ClientAuthorizationGetListResponse) SetDeveloper(v string)`

SetDeveloper sets Developer field to given value.

### HasDeveloper

`func (o *ClientAuthorizationGetListResponse) HasDeveloper() bool`

HasDeveloper returns a boolean if a field has been set.

### GetSubject

`func (o *ClientAuthorizationGetListResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ClientAuthorizationGetListResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ClientAuthorizationGetListResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ClientAuthorizationGetListResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTotalCount

`func (o *ClientAuthorizationGetListResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ClientAuthorizationGetListResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ClientAuthorizationGetListResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ClientAuthorizationGetListResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetClients

`func (o *ClientAuthorizationGetListResponse) GetClients() []Client`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ClientAuthorizationGetListResponse) GetClientsOk() (*[]Client, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ClientAuthorizationGetListResponse) SetClients(v []Client)`

SetClients sets Clients field to given value.

### HasClients

`func (o *ClientAuthorizationGetListResponse) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


