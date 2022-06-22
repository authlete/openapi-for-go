# ClientRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | **string** | Client metadata in JSON format that complies with [RFC 7591](https://datatracker.ietf.org/doc/html/rfc7591) (OAuth 2.0 Dynamic Client Registration Protocol).  | 
**Token** | Pointer to **string** | The client registration access token. Used only for GET, UPDATE, and DELETE requests.  | [optional] 
**ClientId** | Pointer to **string** | The client&#39;s identifier. Used for GET, UPDATE, and DELETE requests  | [optional] 

## Methods

### NewClientRegistrationRequest

`func NewClientRegistrationRequest(json string, ) *ClientRegistrationRequest`

NewClientRegistrationRequest instantiates a new ClientRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientRegistrationRequestWithDefaults

`func NewClientRegistrationRequestWithDefaults() *ClientRegistrationRequest`

NewClientRegistrationRequestWithDefaults instantiates a new ClientRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJson

`func (o *ClientRegistrationRequest) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *ClientRegistrationRequest) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *ClientRegistrationRequest) SetJson(v string)`

SetJson sets Json field to given value.


### GetToken

`func (o *ClientRegistrationRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ClientRegistrationRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ClientRegistrationRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ClientRegistrationRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetClientId

`func (o *ClientRegistrationRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientRegistrationRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientRegistrationRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClientRegistrationRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


