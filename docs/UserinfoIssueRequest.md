# UserinfoIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | The access token that has been passed to the userinfo endpoint by the client application. In other words, the access token which was contained in the userinfo request.  | 
**Claims** | Pointer to **string** | Claims in JSON format. As for the format, see [OpenID Connect Core 1.0, 5.1. Standard Claims](https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims).  | [optional] 
**Sub** | Pointer to **string** | The value of the &#x60;sub&#x60; claim. If the value of this request parameter is not empty, it is used as the value of the &#x60;sub&#x60; claim. Otherwise, the value of the subject associated with the access token is used.  | [optional] 
**ClaimsForTx** | Pointer to **string** | Claim key-value pairs that are used to compute transformed claims.  | [optional] 
**RequestSignature** | Pointer to **string** | The Signature header value from the request.  | [optional] 
**Headers** | Pointer to [**[]Pair**](Pair.md) | HTTP headers to be included in processing the signature. If this is a signed request, this must include the  Signature and Signature-Input headers, as well as any additional headers covered by the signature. | [optional] 

## Methods

### NewUserinfoIssueRequest

`func NewUserinfoIssueRequest(token string, ) *UserinfoIssueRequest`

NewUserinfoIssueRequest instantiates a new UserinfoIssueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserinfoIssueRequestWithDefaults

`func NewUserinfoIssueRequestWithDefaults() *UserinfoIssueRequest`

NewUserinfoIssueRequestWithDefaults instantiates a new UserinfoIssueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *UserinfoIssueRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UserinfoIssueRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UserinfoIssueRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetClaims

`func (o *UserinfoIssueRequest) GetClaims() string`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *UserinfoIssueRequest) GetClaimsOk() (*string, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *UserinfoIssueRequest) SetClaims(v string)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *UserinfoIssueRequest) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetSub

`func (o *UserinfoIssueRequest) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *UserinfoIssueRequest) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *UserinfoIssueRequest) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *UserinfoIssueRequest) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetClaimsForTx

`func (o *UserinfoIssueRequest) GetClaimsForTx() string`

GetClaimsForTx returns the ClaimsForTx field if non-nil, zero value otherwise.

### GetClaimsForTxOk

`func (o *UserinfoIssueRequest) GetClaimsForTxOk() (*string, bool)`

GetClaimsForTxOk returns a tuple with the ClaimsForTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsForTx

`func (o *UserinfoIssueRequest) SetClaimsForTx(v string)`

SetClaimsForTx sets ClaimsForTx field to given value.

### HasClaimsForTx

`func (o *UserinfoIssueRequest) HasClaimsForTx() bool`

HasClaimsForTx returns a boolean if a field has been set.

### GetRequestSignature

`func (o *UserinfoIssueRequest) GetRequestSignature() string`

GetRequestSignature returns the RequestSignature field if non-nil, zero value otherwise.

### GetRequestSignatureOk

`func (o *UserinfoIssueRequest) GetRequestSignatureOk() (*string, bool)`

GetRequestSignatureOk returns a tuple with the RequestSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSignature

`func (o *UserinfoIssueRequest) SetRequestSignature(v string)`

SetRequestSignature sets RequestSignature field to given value.

### HasRequestSignature

`func (o *UserinfoIssueRequest) HasRequestSignature() bool`

HasRequestSignature returns a boolean if a field has been set.

### GetHeaders

`func (o *UserinfoIssueRequest) GetHeaders() []Pair`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *UserinfoIssueRequest) GetHeadersOk() (*[]Pair, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *UserinfoIssueRequest) SetHeaders(v []Pair)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *UserinfoIssueRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


