# IdtokenReissueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | &lt;p&gt; The value of this parameter should be (a) the value of the \&quot;&#x60;jwtAccessToken&#x60;\&quot; parameter in a response from the &#x60;/auth/token&#x60; API when the value is available, or (b) the value of the \&quot;&#x60;accessToken&#x60;\&quot; parameter in the response from the &#x60;/auth/token&#x60; API when the value of the \&quot;&#x60;jwtAccessToken&#x60;\&quot; parameter is not available. &lt;/p&gt;  | 
**RefreshToken** | **string** | &lt;p&gt; The value of this parameter should be the value of the \&quot;&#x60;refreshToken&#x60;\&quot; parameter in a response from the &#x60;/auth/token&#x60; API. &lt;/p&gt;  | 
**Sub** | Pointer to **string** | The value that should be used as the value of the \&quot;&#x60;sub&#x60;\&quot; claim of the ID token.  &lt;p&gt; This parameter is optional. When omitted, the value of the subject associated with the access token is used. &lt;/p&gt;  | [optional] 
**Claims** | Pointer to **string** | Additional claims that should be embedded in the payload part of the ID token. The format is a JSON object.  &lt;p&gt; This parameter is optional. &lt;/p&gt;  | [optional] 
**IdtHeaderParams** | Pointer to **string** | Additional parameters that should be embedded in the JWS header of the ID token. The format is a JSON object.  &lt;p&gt; This parameter is optional. &lt;/p&gt;  | [optional] 
**IdTokenAudType** | Pointer to **string** | The type of the \&quot;&#x60;aud&#x60;\&quot; claim of the ID token being issued.  &lt;p&gt; Valid values of this parameter are as follows. &lt;/p&gt;  &lt;blockquote&gt; &lt;table border&#x3D;\&quot;1\&quot; cellpadding&#x3D;\&quot;5\&quot; style&#x3D;\&quot;border-collapse: collapse;\&quot;&gt;   &lt;tr bgcolor&#x3D;\&quot;orange\&quot;&gt;     &lt;th&gt;Value&lt;/th&gt;     &lt;th&gt;Description&lt;/th&gt;   &lt;/tr&gt;   &lt;tr&gt;     &lt;td&gt;\&quot;&#x60;array&#x60;\&quot;&lt;/td&gt;     &lt;td&gt;The type of the &#x60;aud&#x60; claim becomes an array of strings.&lt;/td&gt;   &lt;/tr&gt;   &lt;tr&gt;     &lt;td&gt;\&quot;&#x60;string&#x60;\&quot;&lt;/td&gt;     &lt;td&gt;The type of the &#x60;aud&#x60; claim becomes a single string.&lt;/td&gt;   &lt;/tr&gt; &lt;/table&gt; &lt;/blockquote&gt;  &lt;p&gt; This parameter is optional, and the default value on omission is \&quot;&#x60;array&#x60;\&quot;. &lt;/p&gt;  &lt;p&gt; This parameter takes precedence over the &#x60;idTokenAudType&#x60; property of {@link Service} (cf. {@link Service#getIdTokenAudType()}). &lt;/p&gt; | [optional] 

## Methods

### NewIdtokenReissueRequest

`func NewIdtokenReissueRequest(accessToken string, refreshToken string, ) *IdtokenReissueRequest`

NewIdtokenReissueRequest instantiates a new IdtokenReissueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdtokenReissueRequestWithDefaults

`func NewIdtokenReissueRequestWithDefaults() *IdtokenReissueRequest`

NewIdtokenReissueRequestWithDefaults instantiates a new IdtokenReissueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *IdtokenReissueRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *IdtokenReissueRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *IdtokenReissueRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetRefreshToken

`func (o *IdtokenReissueRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *IdtokenReissueRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *IdtokenReissueRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.


### GetSub

`func (o *IdtokenReissueRequest) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *IdtokenReissueRequest) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *IdtokenReissueRequest) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *IdtokenReissueRequest) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetClaims

`func (o *IdtokenReissueRequest) GetClaims() string`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *IdtokenReissueRequest) GetClaimsOk() (*string, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *IdtokenReissueRequest) SetClaims(v string)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *IdtokenReissueRequest) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetIdtHeaderParams

`func (o *IdtokenReissueRequest) GetIdtHeaderParams() string`

GetIdtHeaderParams returns the IdtHeaderParams field if non-nil, zero value otherwise.

### GetIdtHeaderParamsOk

`func (o *IdtokenReissueRequest) GetIdtHeaderParamsOk() (*string, bool)`

GetIdtHeaderParamsOk returns a tuple with the IdtHeaderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdtHeaderParams

`func (o *IdtokenReissueRequest) SetIdtHeaderParams(v string)`

SetIdtHeaderParams sets IdtHeaderParams field to given value.

### HasIdtHeaderParams

`func (o *IdtokenReissueRequest) HasIdtHeaderParams() bool`

HasIdtHeaderParams returns a boolean if a field has been set.

### GetIdTokenAudType

`func (o *IdtokenReissueRequest) GetIdTokenAudType() string`

GetIdTokenAudType returns the IdTokenAudType field if non-nil, zero value otherwise.

### GetIdTokenAudTypeOk

`func (o *IdtokenReissueRequest) GetIdTokenAudTypeOk() (*string, bool)`

GetIdTokenAudTypeOk returns a tuple with the IdTokenAudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenAudType

`func (o *IdtokenReissueRequest) SetIdTokenAudType(v string)`

SetIdTokenAudType sets IdTokenAudType field to given value.

### HasIdTokenAudType

`func (o *IdtokenReissueRequest) HasIdTokenAudType() bool`

HasIdTokenAudType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


