# FederationRegistrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take. | [optional] 
**ResponseContent** | Pointer to **string** | The content that the authorization server implementation can use as the value of &#x60;WWW-Authenticate&#x60; header on errors.  | [optional] 
**Client** | Pointer to [**Client**](Client.md) |  | [optional] 

## Methods

### NewFederationRegistrationResponse

`func NewFederationRegistrationResponse() *FederationRegistrationResponse`

NewFederationRegistrationResponse instantiates a new FederationRegistrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationRegistrationResponseWithDefaults

`func NewFederationRegistrationResponseWithDefaults() *FederationRegistrationResponse`

NewFederationRegistrationResponseWithDefaults instantiates a new FederationRegistrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *FederationRegistrationResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *FederationRegistrationResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *FederationRegistrationResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *FederationRegistrationResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *FederationRegistrationResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *FederationRegistrationResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *FederationRegistrationResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *FederationRegistrationResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *FederationRegistrationResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FederationRegistrationResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FederationRegistrationResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *FederationRegistrationResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResponseContent

`func (o *FederationRegistrationResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *FederationRegistrationResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *FederationRegistrationResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *FederationRegistrationResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetClient

`func (o *FederationRegistrationResponse) GetClient() Client`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *FederationRegistrationResponse) GetClientOk() (*Client, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *FederationRegistrationResponse) SetClient(v Client)`

SetClient sets Client field to given value.

### HasClient

`func (o *FederationRegistrationResponse) HasClient() bool`

HasClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


