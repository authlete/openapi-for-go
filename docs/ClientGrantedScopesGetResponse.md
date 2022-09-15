# ClientGrantedScopesGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**ServiceApiKey** | Pointer to **float32** | The API key of the service. | [optional] 
**ClientId** | Pointer to **float32** | The client ID. | [optional] 
**Subject** | Pointer to **string** | The subject (&#x3D; unique identifier) of the user who has granted authorization to the client. | [optional] 
**LatestGrantedScopes** | Pointer to **[]string** | The scopes granted to the client application by the last authorization process by the user (who is identified by the subject). &#x60;null&#x60; means that there is no record about granted scopes. An empty array means that there exists a record about granted scopes but no scope has been granted to the client application. If the returned array holds some elements, they are the scopes granted to the client application by the last authorization process.  | [optional] 
**MergedGrantedScopes** | Pointer to **[]string** | The scopes granted to the client application by all the past authorization processes. Note that revoked scopes are not included.  | [optional] 
**ModifiedAt** | Pointer to **float32** | The timestamp in milliseconds since Unix epoch at which this record was modified. | [optional] 

## Methods

### NewClientGrantedScopesGetResponse

`func NewClientGrantedScopesGetResponse() *ClientGrantedScopesGetResponse`

NewClientGrantedScopesGetResponse instantiates a new ClientGrantedScopesGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientGrantedScopesGetResponseWithDefaults

`func NewClientGrantedScopesGetResponseWithDefaults() *ClientGrantedScopesGetResponse`

NewClientGrantedScopesGetResponseWithDefaults instantiates a new ClientGrantedScopesGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *ClientGrantedScopesGetResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *ClientGrantedScopesGetResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *ClientGrantedScopesGetResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *ClientGrantedScopesGetResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *ClientGrantedScopesGetResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *ClientGrantedScopesGetResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *ClientGrantedScopesGetResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *ClientGrantedScopesGetResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetServiceApiKey

`func (o *ClientGrantedScopesGetResponse) GetServiceApiKey() float32`

GetServiceApiKey returns the ServiceApiKey field if non-nil, zero value otherwise.

### GetServiceApiKeyOk

`func (o *ClientGrantedScopesGetResponse) GetServiceApiKeyOk() (*float32, bool)`

GetServiceApiKeyOk returns a tuple with the ServiceApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceApiKey

`func (o *ClientGrantedScopesGetResponse) SetServiceApiKey(v float32)`

SetServiceApiKey sets ServiceApiKey field to given value.

### HasServiceApiKey

`func (o *ClientGrantedScopesGetResponse) HasServiceApiKey() bool`

HasServiceApiKey returns a boolean if a field has been set.

### GetClientId

`func (o *ClientGrantedScopesGetResponse) GetClientId() float32`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientGrantedScopesGetResponse) GetClientIdOk() (*float32, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientGrantedScopesGetResponse) SetClientId(v float32)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClientGrantedScopesGetResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSubject

`func (o *ClientGrantedScopesGetResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ClientGrantedScopesGetResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ClientGrantedScopesGetResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ClientGrantedScopesGetResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetLatestGrantedScopes

`func (o *ClientGrantedScopesGetResponse) GetLatestGrantedScopes() []string`

GetLatestGrantedScopes returns the LatestGrantedScopes field if non-nil, zero value otherwise.

### GetLatestGrantedScopesOk

`func (o *ClientGrantedScopesGetResponse) GetLatestGrantedScopesOk() (*[]string, bool)`

GetLatestGrantedScopesOk returns a tuple with the LatestGrantedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestGrantedScopes

`func (o *ClientGrantedScopesGetResponse) SetLatestGrantedScopes(v []string)`

SetLatestGrantedScopes sets LatestGrantedScopes field to given value.

### HasLatestGrantedScopes

`func (o *ClientGrantedScopesGetResponse) HasLatestGrantedScopes() bool`

HasLatestGrantedScopes returns a boolean if a field has been set.

### GetMergedGrantedScopes

`func (o *ClientGrantedScopesGetResponse) GetMergedGrantedScopes() []string`

GetMergedGrantedScopes returns the MergedGrantedScopes field if non-nil, zero value otherwise.

### GetMergedGrantedScopesOk

`func (o *ClientGrantedScopesGetResponse) GetMergedGrantedScopesOk() (*[]string, bool)`

GetMergedGrantedScopesOk returns a tuple with the MergedGrantedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedGrantedScopes

`func (o *ClientGrantedScopesGetResponse) SetMergedGrantedScopes(v []string)`

SetMergedGrantedScopes sets MergedGrantedScopes field to given value.

### HasMergedGrantedScopes

`func (o *ClientGrantedScopesGetResponse) HasMergedGrantedScopes() bool`

HasMergedGrantedScopes returns a boolean if a field has been set.

### GetModifiedAt

`func (o *ClientGrantedScopesGetResponse) GetModifiedAt() float32`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ClientGrantedScopesGetResponse) GetModifiedAtOk() (*float32, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ClientGrantedScopesGetResponse) SetModifiedAt(v float32)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ClientGrantedScopesGetResponse) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


