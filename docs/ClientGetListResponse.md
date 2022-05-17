# ClientGetListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Developer** | Pointer to **string** | The developer of the client applications. If the request did not contain &#x60;developer&#x60; request parameter, this property is set to &#x60;null&#x60;.  | [optional] 
**Start** | Pointer to **int32** | Start index (inclusive) of the result set of the query.  | [optional] 
**End** | Pointer to **int32** | End index (exclusive) of the result set of the query.  | [optional] 
**TotalCount** | Pointer to **int32** | Total number of clients that belong to the service. This doesn&#39;t mean the number of clients contained in the response.  | [optional] 
**Clients** | Pointer to [**[]Client**](Client.md) | An array of clients.  | [optional] 

## Methods

### NewClientGetListResponse

`func NewClientGetListResponse() *ClientGetListResponse`

NewClientGetListResponse instantiates a new ClientGetListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientGetListResponseWithDefaults

`func NewClientGetListResponseWithDefaults() *ClientGetListResponse`

NewClientGetListResponseWithDefaults instantiates a new ClientGetListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeveloper

`func (o *ClientGetListResponse) GetDeveloper() string`

GetDeveloper returns the Developer field if non-nil, zero value otherwise.

### GetDeveloperOk

`func (o *ClientGetListResponse) GetDeveloperOk() (*string, bool)`

GetDeveloperOk returns a tuple with the Developer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloper

`func (o *ClientGetListResponse) SetDeveloper(v string)`

SetDeveloper sets Developer field to given value.

### HasDeveloper

`func (o *ClientGetListResponse) HasDeveloper() bool`

HasDeveloper returns a boolean if a field has been set.

### GetStart

`func (o *ClientGetListResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ClientGetListResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ClientGetListResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ClientGetListResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *ClientGetListResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ClientGetListResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ClientGetListResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ClientGetListResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetTotalCount

`func (o *ClientGetListResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ClientGetListResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ClientGetListResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ClientGetListResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetClients

`func (o *ClientGetListResponse) GetClients() []Client`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ClientGetListResponse) GetClientsOk() (*[]Client, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ClientGetListResponse) SetClients(v []Client)`

SetClients sets Clients field to given value.

### HasClients

`func (o *ClientGetListResponse) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


