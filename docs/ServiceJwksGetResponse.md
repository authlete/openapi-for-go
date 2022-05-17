# ServiceJwksGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to **[]map[string]interface{}** | An array of [JWK](https://datatracker.ietf.org/doc/html/rfc7517)s. | [optional] 

## Methods

### NewServiceJwksGetResponse

`func NewServiceJwksGetResponse() *ServiceJwksGetResponse`

NewServiceJwksGetResponse instantiates a new ServiceJwksGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceJwksGetResponseWithDefaults

`func NewServiceJwksGetResponseWithDefaults() *ServiceJwksGetResponse`

NewServiceJwksGetResponseWithDefaults instantiates a new ServiceJwksGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *ServiceJwksGetResponse) GetKeys() []map[string]interface{}`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ServiceJwksGetResponse) GetKeysOk() (*[]map[string]interface{}, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ServiceJwksGetResponse) SetKeys(v []map[string]interface{})`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ServiceJwksGetResponse) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


