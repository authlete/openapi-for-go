# ClientRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Json** | **string** | Client metadata in JSON format that complies with [RFC 7591](https://datatracker.ietf.org/doc/html/rfc7591) (OAuth 2.0 Dynamic Client Registration Protocol). | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


