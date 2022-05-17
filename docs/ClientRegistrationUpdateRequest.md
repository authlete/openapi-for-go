# ClientRegistrationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Client ID.  | 
**Token** | **string** | Client registration access token.  | 
**Json** | **string** | Client metadata in JSON format that complies with [RFC 7591](https://datatracker.ietf.org/doc/html/rfc7591) (OAuth 2.0 Dynamic Client Registration Protocol). | 

## Methods

### NewClientRegistrationUpdateRequest

`func NewClientRegistrationUpdateRequest(clientId string, token string, json string, ) *ClientRegistrationUpdateRequest`

NewClientRegistrationUpdateRequest instantiates a new ClientRegistrationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientRegistrationUpdateRequestWithDefaults

`func NewClientRegistrationUpdateRequestWithDefaults() *ClientRegistrationUpdateRequest`

NewClientRegistrationUpdateRequestWithDefaults instantiates a new ClientRegistrationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ClientRegistrationUpdateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientRegistrationUpdateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientRegistrationUpdateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetToken

`func (o *ClientRegistrationUpdateRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ClientRegistrationUpdateRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ClientRegistrationUpdateRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetJson

`func (o *ClientRegistrationUpdateRequest) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *ClientRegistrationUpdateRequest) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *ClientRegistrationUpdateRequest) SetJson(v string)`

SetJson sets Json field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


