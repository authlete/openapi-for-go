# ClientExtensionRequestableScopesUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestableScopes** | Pointer to **[]string** | The set of scopes that the client application is allowed to request. This parameter will be one of the following. Details are described in the description.  - null - an empty set - a set with at least one element  If this parameter contains scopes that the service does not support, those scopes are just ignored. Also, if this parameter is &#x60;null&#x60; or is not included in the request, it is equivalent to calling &#x60;/client/extension/requestable_scopes/delete&#x60; API. | [optional] 

## Methods

### NewClientExtensionRequestableScopesUpdateRequest

`func NewClientExtensionRequestableScopesUpdateRequest() *ClientExtensionRequestableScopesUpdateRequest`

NewClientExtensionRequestableScopesUpdateRequest instantiates a new ClientExtensionRequestableScopesUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientExtensionRequestableScopesUpdateRequestWithDefaults

`func NewClientExtensionRequestableScopesUpdateRequestWithDefaults() *ClientExtensionRequestableScopesUpdateRequest`

NewClientExtensionRequestableScopesUpdateRequestWithDefaults instantiates a new ClientExtensionRequestableScopesUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestableScopes

`func (o *ClientExtensionRequestableScopesUpdateRequest) GetRequestableScopes() []string`

GetRequestableScopes returns the RequestableScopes field if non-nil, zero value otherwise.

### GetRequestableScopesOk

`func (o *ClientExtensionRequestableScopesUpdateRequest) GetRequestableScopesOk() (*[]string, bool)`

GetRequestableScopesOk returns a tuple with the RequestableScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestableScopes

`func (o *ClientExtensionRequestableScopesUpdateRequest) SetRequestableScopes(v []string)`

SetRequestableScopes sets RequestableScopes field to given value.

### HasRequestableScopes

`func (o *ClientExtensionRequestableScopesUpdateRequest) HasRequestableScopes() bool`

HasRequestableScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


