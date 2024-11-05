# VciBatchIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | The access token that came along with the credential request. | [optional] 
**Orders** | Pointer to [**[]CredentialIssuanceOrder**](CredentialIssuanceOrder.md) | The instructions for issuance of credentials and/or transaction IDs. | [optional] 

## Methods

### NewVciBatchIssueRequest

`func NewVciBatchIssueRequest() *VciBatchIssueRequest`

NewVciBatchIssueRequest instantiates a new VciBatchIssueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVciBatchIssueRequestWithDefaults

`func NewVciBatchIssueRequestWithDefaults() *VciBatchIssueRequest`

NewVciBatchIssueRequestWithDefaults instantiates a new VciBatchIssueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *VciBatchIssueRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *VciBatchIssueRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *VciBatchIssueRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *VciBatchIssueRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetOrders

`func (o *VciBatchIssueRequest) GetOrders() []CredentialIssuanceOrder`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *VciBatchIssueRequest) GetOrdersOk() (*[]CredentialIssuanceOrder, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *VciBatchIssueRequest) SetOrders(v []CredentialIssuanceOrder)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *VciBatchIssueRequest) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


