# VciSingleIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | The access token that came along with the credential request. | [optional] 
**Order** | Pointer to [**CredentialIssuanceOrder**](CredentialIssuanceOrder.md) |  | [optional] 

## Methods

### NewVciSingleIssueRequest

`func NewVciSingleIssueRequest() *VciSingleIssueRequest`

NewVciSingleIssueRequest instantiates a new VciSingleIssueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVciSingleIssueRequestWithDefaults

`func NewVciSingleIssueRequestWithDefaults() *VciSingleIssueRequest`

NewVciSingleIssueRequestWithDefaults instantiates a new VciSingleIssueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *VciSingleIssueRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *VciSingleIssueRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *VciSingleIssueRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *VciSingleIssueRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetOrder

`func (o *VciSingleIssueRequest) GetOrder() CredentialIssuanceOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VciSingleIssueRequest) GetOrderOk() (*CredentialIssuanceOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VciSingleIssueRequest) SetOrder(v CredentialIssuanceOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VciSingleIssueRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


