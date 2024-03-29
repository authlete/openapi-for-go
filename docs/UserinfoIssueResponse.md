# UserinfoIssueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take. | [optional] 
**ResponseContent** | Pointer to **string** | The content that the authorization server implementation can use as the value of &#x60;WWW-Authenticate&#x60; header on errors.  | [optional] 
**Signature** | Pointer to **string** | The signature header of the response message.  | [optional] 
**SignatureInput** | Pointer to **string** | The signature-input header of the response message  | [optional] 
**ContentDigest** | Pointer to **string** | The content-digest header of the response message  | [optional] 

## Methods

### NewUserinfoIssueResponse

`func NewUserinfoIssueResponse() *UserinfoIssueResponse`

NewUserinfoIssueResponse instantiates a new UserinfoIssueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserinfoIssueResponseWithDefaults

`func NewUserinfoIssueResponseWithDefaults() *UserinfoIssueResponse`

NewUserinfoIssueResponseWithDefaults instantiates a new UserinfoIssueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *UserinfoIssueResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *UserinfoIssueResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *UserinfoIssueResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *UserinfoIssueResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *UserinfoIssueResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *UserinfoIssueResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *UserinfoIssueResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *UserinfoIssueResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *UserinfoIssueResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UserinfoIssueResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UserinfoIssueResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *UserinfoIssueResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResponseContent

`func (o *UserinfoIssueResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *UserinfoIssueResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *UserinfoIssueResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *UserinfoIssueResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetSignature

`func (o *UserinfoIssueResponse) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *UserinfoIssueResponse) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *UserinfoIssueResponse) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *UserinfoIssueResponse) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSignatureInput

`func (o *UserinfoIssueResponse) GetSignatureInput() string`

GetSignatureInput returns the SignatureInput field if non-nil, zero value otherwise.

### GetSignatureInputOk

`func (o *UserinfoIssueResponse) GetSignatureInputOk() (*string, bool)`

GetSignatureInputOk returns a tuple with the SignatureInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureInput

`func (o *UserinfoIssueResponse) SetSignatureInput(v string)`

SetSignatureInput sets SignatureInput field to given value.

### HasSignatureInput

`func (o *UserinfoIssueResponse) HasSignatureInput() bool`

HasSignatureInput returns a boolean if a field has been set.

### GetContentDigest

`func (o *UserinfoIssueResponse) GetContentDigest() string`

GetContentDigest returns the ContentDigest field if non-nil, zero value otherwise.

### GetContentDigestOk

`func (o *UserinfoIssueResponse) GetContentDigestOk() (*string, bool)`

GetContentDigestOk returns a tuple with the ContentDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentDigest

`func (o *UserinfoIssueResponse) SetContentDigest(v string)`

SetContentDigest sets ContentDigest field to given value.

### HasContentDigest

`func (o *UserinfoIssueResponse) HasContentDigest() bool`

HasContentDigest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


