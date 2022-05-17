# ClientRegistrationGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Client ID.  | 
**Token** | **string** | Client registration access token. | 

## Methods

### NewClientRegistrationGetRequest

`func NewClientRegistrationGetRequest(clientId string, token string, ) *ClientRegistrationGetRequest`

NewClientRegistrationGetRequest instantiates a new ClientRegistrationGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientRegistrationGetRequestWithDefaults

`func NewClientRegistrationGetRequestWithDefaults() *ClientRegistrationGetRequest`

NewClientRegistrationGetRequestWithDefaults instantiates a new ClientRegistrationGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ClientRegistrationGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientRegistrationGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientRegistrationGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetToken

`func (o *ClientRegistrationGetRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ClientRegistrationGetRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ClientRegistrationGetRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


