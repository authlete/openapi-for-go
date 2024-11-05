# ClientGetListApi200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **int32** | Start index (inclusive) of the result set of the query.  | [optional] 
**End** | Pointer to **int32** | End index (exclusive) of the result set of the query.  | [optional] 
**TotalCount** | Pointer to **int32** | Total number of clients that belong to the service. This doesn&#39;t mean the number of clients contained in the response.  | [optional] 
**Clients** | Pointer to [**[]ClientLimited**](ClientLimited.md) | An array of clients.  | [optional] 

## Methods

### NewClientGetListApi200Response

`func NewClientGetListApi200Response() *ClientGetListApi200Response`

NewClientGetListApi200Response instantiates a new ClientGetListApi200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientGetListApi200ResponseWithDefaults

`func NewClientGetListApi200ResponseWithDefaults() *ClientGetListApi200Response`

NewClientGetListApi200ResponseWithDefaults instantiates a new ClientGetListApi200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *ClientGetListApi200Response) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ClientGetListApi200Response) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ClientGetListApi200Response) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ClientGetListApi200Response) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *ClientGetListApi200Response) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ClientGetListApi200Response) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ClientGetListApi200Response) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ClientGetListApi200Response) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetTotalCount

`func (o *ClientGetListApi200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ClientGetListApi200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ClientGetListApi200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ClientGetListApi200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetClients

`func (o *ClientGetListApi200Response) GetClients() []ClientLimited`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ClientGetListApi200Response) GetClientsOk() (*[]ClientLimited, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ClientGetListApi200Response) SetClients(v []ClientLimited)`

SetClients sets Clients field to given value.

### HasClients

`func (o *ClientGetListApi200Response) HasClients() bool`

HasClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


