# ClientSecretUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSecret** | **string** | The new value of the client secret. Valid characters for a client secret are &#x60;A-Z&#x60;, &#x60;a-z&#x60;, &#x60;0-9&#x60;, &#x60;-&#x60;, and &#x60;_&#x60;. The maximum length of a client secret is 86. | 

## Methods

### NewClientSecretUpdateRequest

`func NewClientSecretUpdateRequest(clientSecret string, ) *ClientSecretUpdateRequest`

NewClientSecretUpdateRequest instantiates a new ClientSecretUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientSecretUpdateRequestWithDefaults

`func NewClientSecretUpdateRequestWithDefaults() *ClientSecretUpdateRequest`

NewClientSecretUpdateRequestWithDefaults instantiates a new ClientSecretUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSecret

`func (o *ClientSecretUpdateRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ClientSecretUpdateRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ClientSecretUpdateRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


