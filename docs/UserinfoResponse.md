# UserinfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take. | [optional] 
**Claims** | Pointer to **[]string** | The list of claims that the client application requests to be embedded in the ID token.  | [optional] 
**ClientId** | Pointer to **string** | The ID of the client application which is associated with the access token.  | [optional] 
**ClientIdAlias** | Pointer to **string** | The client ID alias when the authorization request for the access token was made.  | [optional] 
**ClientIdAliasUsed** | Pointer to **bool** | The flag which indicates whether the client ID alias was used when the authorization request for the access token was made.  | [optional] 
**ResponseContent** | Pointer to **string** | The content that the authorization server implementation can use as the value of &#x60;WWW-Authenticate&#x60; header on errors.  | [optional] 
**Scopes** | Pointer to **[]string** | The scopes covered by the access token.  | [optional] 
**Subject** | Pointer to **string** | The subject (&#x3D; resource owner&#39;s ID).  | [optional] 
**Token** | Pointer to **string** | The access token that came along with the userinfo request.  | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) | The extra properties associated with the access token.  | [optional] 
**UserInfoClaims** | Pointer to **string** | The value of the &#x60;userinfo&#x60; property in the &#x60;claims&#x60; request parameter or in the &#x60;claims&#x60; property in an authorization request object.  A client application may request certain claims be embedded in an ID token or in a response from the userInfo endpoint. There are several ways. Including the &#x60;claims&#x60; request parameter and including the &#x60;claims&#x60; property in a request object are such examples. In both cases, the value of the &#x60;claims&#x60; parameter/property is JSON. Its format is described in [5.5. Requesting Claims using the \&quot;claims\&quot; Request Parameter](https://openid.net/specs/openid-connect-core-1_0.html#ClaimsParameter).  The following is an excerpt from the specification. You can find &#x60;userinfo&#x60; and &#x60;id_token&#x60; are top-level properties.  &#x60;&#x60;&#x60;json {   \&quot;userinfo\&quot;:   {     \&quot;given_name\&quot;: { \&quot;essential\&quot;: true },     \&quot;nickname\&quot;: null,     \&quot;email\&quot;: { \&quot;essential\&quot;: true },     \&quot;email_verified\&quot;: { \&quot;essential\&quot;: true },     \&quot;picture\&quot;: null,     \&quot;http://example.info/claims/groups\&quot;: null   },   \&quot;id_token\&quot;:   {     \&quot;auth_time\&quot;: { \&quot;essential\&quot;: true },     \&quot;acr\&quot;: { \&quot;values\&quot;: [ \&quot;urn:mace:incommon:iap:silver\&quot; ] }   } } &#x60;&#x60;&#x60;&#x60;  The value of this property is the value of the &#x60;userinfo&#x60; property in JSON format. For example, if the JSON above is included in an authorization request, this property holds JSON equivalent to the following.  &#x60;&#x60;&#x60;json {   \&quot;given_name\&quot;: { \&quot;essential\&quot;: true },   \&quot;nickname\&quot;: null,   \&quot;email\&quot;: { \&quot;essential\&quot;: true },   \&quot;email_verified\&quot;: { \&quot;essential\&quot;: true },   \&quot;picture\&quot;: null,   \&quot;http://example.info/claims/groups\&quot;: null } &#x60;&#x60;&#x60;  Note that if a request object is given and it contains the &#x60;claims&#x60; property and if the &#x60;claims&#x60; request parameter is also given, the value of this property holds the former value.  | [optional] 
**ServiceAttributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of this service that the client application belongs to.  | [optional] 
**ClientAttributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of the client.  | [optional] 

## Methods

### NewUserinfoResponse

`func NewUserinfoResponse() *UserinfoResponse`

NewUserinfoResponse instantiates a new UserinfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserinfoResponseWithDefaults

`func NewUserinfoResponseWithDefaults() *UserinfoResponse`

NewUserinfoResponseWithDefaults instantiates a new UserinfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *UserinfoResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *UserinfoResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *UserinfoResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *UserinfoResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *UserinfoResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *UserinfoResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *UserinfoResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *UserinfoResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *UserinfoResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UserinfoResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UserinfoResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *UserinfoResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetClaims

`func (o *UserinfoResponse) GetClaims() []string`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *UserinfoResponse) GetClaimsOk() (*[]string, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *UserinfoResponse) SetClaims(v []string)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *UserinfoResponse) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetClientId

`func (o *UserinfoResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *UserinfoResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *UserinfoResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *UserinfoResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIdAlias

`func (o *UserinfoResponse) GetClientIdAlias() string`

GetClientIdAlias returns the ClientIdAlias field if non-nil, zero value otherwise.

### GetClientIdAliasOk

`func (o *UserinfoResponse) GetClientIdAliasOk() (*string, bool)`

GetClientIdAliasOk returns a tuple with the ClientIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAlias

`func (o *UserinfoResponse) SetClientIdAlias(v string)`

SetClientIdAlias sets ClientIdAlias field to given value.

### HasClientIdAlias

`func (o *UserinfoResponse) HasClientIdAlias() bool`

HasClientIdAlias returns a boolean if a field has been set.

### GetClientIdAliasUsed

`func (o *UserinfoResponse) GetClientIdAliasUsed() bool`

GetClientIdAliasUsed returns the ClientIdAliasUsed field if non-nil, zero value otherwise.

### GetClientIdAliasUsedOk

`func (o *UserinfoResponse) GetClientIdAliasUsedOk() (*bool, bool)`

GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasUsed

`func (o *UserinfoResponse) SetClientIdAliasUsed(v bool)`

SetClientIdAliasUsed sets ClientIdAliasUsed field to given value.

### HasClientIdAliasUsed

`func (o *UserinfoResponse) HasClientIdAliasUsed() bool`

HasClientIdAliasUsed returns a boolean if a field has been set.

### GetResponseContent

`func (o *UserinfoResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *UserinfoResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *UserinfoResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *UserinfoResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetScopes

`func (o *UserinfoResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *UserinfoResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *UserinfoResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *UserinfoResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSubject

`func (o *UserinfoResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *UserinfoResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *UserinfoResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *UserinfoResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetToken

`func (o *UserinfoResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UserinfoResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UserinfoResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UserinfoResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetProperties

`func (o *UserinfoResponse) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UserinfoResponse) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UserinfoResponse) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *UserinfoResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetUserInfoClaims

`func (o *UserinfoResponse) GetUserInfoClaims() string`

GetUserInfoClaims returns the UserInfoClaims field if non-nil, zero value otherwise.

### GetUserInfoClaimsOk

`func (o *UserinfoResponse) GetUserInfoClaimsOk() (*string, bool)`

GetUserInfoClaimsOk returns a tuple with the UserInfoClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoClaims

`func (o *UserinfoResponse) SetUserInfoClaims(v string)`

SetUserInfoClaims sets UserInfoClaims field to given value.

### HasUserInfoClaims

`func (o *UserinfoResponse) HasUserInfoClaims() bool`

HasUserInfoClaims returns a boolean if a field has been set.

### GetServiceAttributes

`func (o *UserinfoResponse) GetServiceAttributes() []Pair`

GetServiceAttributes returns the ServiceAttributes field if non-nil, zero value otherwise.

### GetServiceAttributesOk

`func (o *UserinfoResponse) GetServiceAttributesOk() (*[]Pair, bool)`

GetServiceAttributesOk returns a tuple with the ServiceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAttributes

`func (o *UserinfoResponse) SetServiceAttributes(v []Pair)`

SetServiceAttributes sets ServiceAttributes field to given value.

### HasServiceAttributes

`func (o *UserinfoResponse) HasServiceAttributes() bool`

HasServiceAttributes returns a boolean if a field has been set.

### GetClientAttributes

`func (o *UserinfoResponse) GetClientAttributes() []Pair`

GetClientAttributes returns the ClientAttributes field if non-nil, zero value otherwise.

### GetClientAttributesOk

`func (o *UserinfoResponse) GetClientAttributesOk() (*[]Pair, bool)`

GetClientAttributesOk returns a tuple with the ClientAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAttributes

`func (o *UserinfoResponse) SetClientAttributes(v []Pair)`

SetClientAttributes sets ClientAttributes field to given value.

### HasClientAttributes

`func (o *UserinfoResponse) HasClientAttributes() bool`

HasClientAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


