# ClientAuthorizationDeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**ServiceApiKey** | Pointer to **int64** | A short message which explains the result of the API call. | [optional] 
**ClientId** | Pointer to **int64** | Get the client ID. | [optional] 
**Subject** | Pointer to **string** | Get the subject (&#x3D; unique identifier) of the user who has granted authorization to the client.  | [optional] 
**LatestGrantedScopes** | Pointer to **[]string** | Get the scopes granted to the client application by the last authorization process by the user (who is identified by the subject).  | [optional] 
**MergedGrantedScopes** | Pointer to **[]string** | Get the scopes granted to the client application by all the past authorization processes. Note that revoked scopes are not included.  | [optional] 
**ModifiedAt** | Pointer to **int64** | Get the timestamp in milliseconds since Unix epoch at which this record was modified. | [optional] 

## Methods

### NewClientAuthorizationDeleteResponse

`func NewClientAuthorizationDeleteResponse() *ClientAuthorizationDeleteResponse`

NewClientAuthorizationDeleteResponse instantiates a new ClientAuthorizationDeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAuthorizationDeleteResponseWithDefaults

`func NewClientAuthorizationDeleteResponseWithDefaults() *ClientAuthorizationDeleteResponse`

NewClientAuthorizationDeleteResponseWithDefaults instantiates a new ClientAuthorizationDeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *ClientAuthorizationDeleteResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *ClientAuthorizationDeleteResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *ClientAuthorizationDeleteResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *ClientAuthorizationDeleteResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *ClientAuthorizationDeleteResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *ClientAuthorizationDeleteResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *ClientAuthorizationDeleteResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *ClientAuthorizationDeleteResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetServiceApiKey

`func (o *ClientAuthorizationDeleteResponse) GetServiceApiKey() int64`

GetServiceApiKey returns the ServiceApiKey field if non-nil, zero value otherwise.

### GetServiceApiKeyOk

`func (o *ClientAuthorizationDeleteResponse) GetServiceApiKeyOk() (*int64, bool)`

GetServiceApiKeyOk returns a tuple with the ServiceApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceApiKey

`func (o *ClientAuthorizationDeleteResponse) SetServiceApiKey(v int64)`

SetServiceApiKey sets ServiceApiKey field to given value.

### HasServiceApiKey

`func (o *ClientAuthorizationDeleteResponse) HasServiceApiKey() bool`

HasServiceApiKey returns a boolean if a field has been set.

### GetClientId

`func (o *ClientAuthorizationDeleteResponse) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientAuthorizationDeleteResponse) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientAuthorizationDeleteResponse) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClientAuthorizationDeleteResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSubject

`func (o *ClientAuthorizationDeleteResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ClientAuthorizationDeleteResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ClientAuthorizationDeleteResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ClientAuthorizationDeleteResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetLatestGrantedScopes

`func (o *ClientAuthorizationDeleteResponse) GetLatestGrantedScopes() []string`

GetLatestGrantedScopes returns the LatestGrantedScopes field if non-nil, zero value otherwise.

### GetLatestGrantedScopesOk

`func (o *ClientAuthorizationDeleteResponse) GetLatestGrantedScopesOk() (*[]string, bool)`

GetLatestGrantedScopesOk returns a tuple with the LatestGrantedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestGrantedScopes

`func (o *ClientAuthorizationDeleteResponse) SetLatestGrantedScopes(v []string)`

SetLatestGrantedScopes sets LatestGrantedScopes field to given value.

### HasLatestGrantedScopes

`func (o *ClientAuthorizationDeleteResponse) HasLatestGrantedScopes() bool`

HasLatestGrantedScopes returns a boolean if a field has been set.

### GetMergedGrantedScopes

`func (o *ClientAuthorizationDeleteResponse) GetMergedGrantedScopes() []string`

GetMergedGrantedScopes returns the MergedGrantedScopes field if non-nil, zero value otherwise.

### GetMergedGrantedScopesOk

`func (o *ClientAuthorizationDeleteResponse) GetMergedGrantedScopesOk() (*[]string, bool)`

GetMergedGrantedScopesOk returns a tuple with the MergedGrantedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedGrantedScopes

`func (o *ClientAuthorizationDeleteResponse) SetMergedGrantedScopes(v []string)`

SetMergedGrantedScopes sets MergedGrantedScopes field to given value.

### HasMergedGrantedScopes

`func (o *ClientAuthorizationDeleteResponse) HasMergedGrantedScopes() bool`

HasMergedGrantedScopes returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ClientAuthorizationDeleteResponse) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ClientAuthorizationDeleteResponse) GetModifiedAtOk() (*int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ClientAuthorizationDeleteResponse) SetModifiedAt(v int64)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ClientAuthorizationDeleteResponse) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


