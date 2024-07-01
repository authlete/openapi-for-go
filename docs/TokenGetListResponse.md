# TokenGetListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **int32** | Start index of search results (inclusive).  | [optional] 
**End** | Pointer to **int32** | End index of search results (exclusive).  | [optional] 
**TotalCount** | Pointer to **int32** | Unique ID of a client developer.  | [optional] 
**Client** | Pointer to [**ClientLimited**](ClientLimited.md) |  | [optional] 
**Subject** | Pointer to **string** | Unique user ID of an end-user.  | [optional] 
**AccessTokens** | Pointer to [**[]AccessToken**](AccessToken.md) | An array of access tokens.  | [optional] 

## Methods

### NewTokenGetListResponse

`func NewTokenGetListResponse() *TokenGetListResponse`

NewTokenGetListResponse instantiates a new TokenGetListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenGetListResponseWithDefaults

`func NewTokenGetListResponseWithDefaults() *TokenGetListResponse`

NewTokenGetListResponseWithDefaults instantiates a new TokenGetListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *TokenGetListResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TokenGetListResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TokenGetListResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *TokenGetListResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *TokenGetListResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TokenGetListResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TokenGetListResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *TokenGetListResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetTotalCount

`func (o *TokenGetListResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TokenGetListResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TokenGetListResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *TokenGetListResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetClient

`func (o *TokenGetListResponse) GetClient() ClientLimited`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *TokenGetListResponse) GetClientOk() (*ClientLimited, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *TokenGetListResponse) SetClient(v ClientLimited)`

SetClient sets Client field to given value.

### HasClient

`func (o *TokenGetListResponse) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetSubject

`func (o *TokenGetListResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TokenGetListResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TokenGetListResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TokenGetListResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetAccessTokens

`func (o *TokenGetListResponse) GetAccessTokens() []AccessToken`

GetAccessTokens returns the AccessTokens field if non-nil, zero value otherwise.

### GetAccessTokensOk

`func (o *TokenGetListResponse) GetAccessTokensOk() (*[]AccessToken, bool)`

GetAccessTokensOk returns a tuple with the AccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokens

`func (o *TokenGetListResponse) SetAccessTokens(v []AccessToken)`

SetAccessTokens sets AccessTokens field to given value.

### HasAccessTokens

`func (o *TokenGetListResponse) HasAccessTokens() bool`

HasAccessTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


