# ClientAuthorizationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** | The subject (&#x3D; unique identifier) of the end-user who has granted authorization to the client application.  | 
**Scopes** | Pointer to **[]string** | An array of new scopes. Optional. If a non-null value is given, the new scopes are set to all existing access tokens. If an API call is made using &#x60;\&quot;Content-Type: application/x-www-form-urlencoded\&quot;&#x60;, scope names listed in this request parameter should be delimited by spaces (after form encoding, spaces are converted to &#x60;+&#x60;). | [optional] 

## Methods

### NewClientAuthorizationUpdateRequest

`func NewClientAuthorizationUpdateRequest(subject string, ) *ClientAuthorizationUpdateRequest`

NewClientAuthorizationUpdateRequest instantiates a new ClientAuthorizationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAuthorizationUpdateRequestWithDefaults

`func NewClientAuthorizationUpdateRequestWithDefaults() *ClientAuthorizationUpdateRequest`

NewClientAuthorizationUpdateRequestWithDefaults instantiates a new ClientAuthorizationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *ClientAuthorizationUpdateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ClientAuthorizationUpdateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ClientAuthorizationUpdateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetScopes

`func (o *ClientAuthorizationUpdateRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ClientAuthorizationUpdateRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ClientAuthorizationUpdateRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ClientAuthorizationUpdateRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


