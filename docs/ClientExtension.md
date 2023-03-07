# ClientExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestableScopes** | Pointer to **[]string** | The set of scopes that the client application is allowed to request. This paramter will be one of the following.    - &#x60;null&#x60;   - an empty set   - a set with at least one element  When the value of this parameter is &#x60;null&#x60;, it means that the set of scopes that the client application is allowed to request is the set of the scopes that the service supports. When the value of this parameter is an empty set, it means that the client application is not allowed to request any scopes. When the value of this parameter is a set with at least one element, it means that the set is the set of scopes that the client application is allowed to request.  | [optional] 
**RequestableScopesEnabled** | Pointer to **bool** | The flag to indicate whether \&quot;Requestable Scopes per Client\&quot; is enabled or not. If &#x60;true&#x60;, you can define the set of scopes which this client application can request. If &#x60;false&#x60;, this client application can request any scope which is supported by the authorization server.  | [optional] 
**AccessTokenDuration** | Pointer to **int64** | The value of the duration of access tokens per client in seconds. In normal cases, the value of the service&#39;s &#x60;accessTokenDuration&#x60; property is used as the duration of access tokens issued by the service. However, if this &#x60;accessTokenDuration&#x60; property holds a non-zero positive number and its value is less than the duration configured by the service, the value is used as the duration of access tokens issued to the client application.  Note that the duration of access tokens can be controlled by the scope attribute &#x60;access_token.duration&#x60;, too. Authlete chooses the minimum value among the candidates.  | [optional] 
**RefreshTokenDuration** | Pointer to **int64** | The value of the duration of refresh tokens per client in seconds. In normal cases, the value of the service&#39;s &#x60;refreshTokenDuration&#x60; property is used as the duration of refresh tokens issued by the service. However, if this &#x60;refreshTokenDuration&#x60; property holds a non-zero positive number and its value is less than the duration configured by the service, the value is used as the duration of refresh tokens issued to the client application.  Note that the duration of refresh tokens can be controlled by the scope attribute &#x60;refresh_token.duration&#x60;, too. Authlete chooses the minimum value among the candidates.  | [optional] 
**TokenExchangePermitted** | Pointer to **bool** | Get the flag indicating whether the client is explicitly given a permission to make token exchange requests ([RFC 8693][https://www.rfc-editor.org/rfc/rfc8693.html]) | [optional] 

## Methods

### NewClientExtension

`func NewClientExtension() *ClientExtension`

NewClientExtension instantiates a new ClientExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientExtensionWithDefaults

`func NewClientExtensionWithDefaults() *ClientExtension`

NewClientExtensionWithDefaults instantiates a new ClientExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestableScopes

`func (o *ClientExtension) GetRequestableScopes() []string`

GetRequestableScopes returns the RequestableScopes field if non-nil, zero value otherwise.

### GetRequestableScopesOk

`func (o *ClientExtension) GetRequestableScopesOk() (*[]string, bool)`

GetRequestableScopesOk returns a tuple with the RequestableScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestableScopes

`func (o *ClientExtension) SetRequestableScopes(v []string)`

SetRequestableScopes sets RequestableScopes field to given value.

### HasRequestableScopes

`func (o *ClientExtension) HasRequestableScopes() bool`

HasRequestableScopes returns a boolean if a field has been set.

### GetRequestableScopesEnabled

`func (o *ClientExtension) GetRequestableScopesEnabled() bool`

GetRequestableScopesEnabled returns the RequestableScopesEnabled field if non-nil, zero value otherwise.

### GetRequestableScopesEnabledOk

`func (o *ClientExtension) GetRequestableScopesEnabledOk() (*bool, bool)`

GetRequestableScopesEnabledOk returns a tuple with the RequestableScopesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestableScopesEnabled

`func (o *ClientExtension) SetRequestableScopesEnabled(v bool)`

SetRequestableScopesEnabled sets RequestableScopesEnabled field to given value.

### HasRequestableScopesEnabled

`func (o *ClientExtension) HasRequestableScopesEnabled() bool`

HasRequestableScopesEnabled returns a boolean if a field has been set.

### GetAccessTokenDuration

`func (o *ClientExtension) GetAccessTokenDuration() int64`

GetAccessTokenDuration returns the AccessTokenDuration field if non-nil, zero value otherwise.

### GetAccessTokenDurationOk

`func (o *ClientExtension) GetAccessTokenDurationOk() (*int64, bool)`

GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenDuration

`func (o *ClientExtension) SetAccessTokenDuration(v int64)`

SetAccessTokenDuration sets AccessTokenDuration field to given value.

### HasAccessTokenDuration

`func (o *ClientExtension) HasAccessTokenDuration() bool`

HasAccessTokenDuration returns a boolean if a field has been set.

### GetRefreshTokenDuration

`func (o *ClientExtension) GetRefreshTokenDuration() int64`

GetRefreshTokenDuration returns the RefreshTokenDuration field if non-nil, zero value otherwise.

### GetRefreshTokenDurationOk

`func (o *ClientExtension) GetRefreshTokenDurationOk() (*int64, bool)`

GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDuration

`func (o *ClientExtension) SetRefreshTokenDuration(v int64)`

SetRefreshTokenDuration sets RefreshTokenDuration field to given value.

### HasRefreshTokenDuration

`func (o *ClientExtension) HasRefreshTokenDuration() bool`

HasRefreshTokenDuration returns a boolean if a field has been set.

### GetTokenExchangePermitted

`func (o *ClientExtension) GetTokenExchangePermitted() bool`

GetTokenExchangePermitted returns the TokenExchangePermitted field if non-nil, zero value otherwise.

### GetTokenExchangePermittedOk

`func (o *ClientExtension) GetTokenExchangePermittedOk() (*bool, bool)`

GetTokenExchangePermittedOk returns a tuple with the TokenExchangePermitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExchangePermitted

`func (o *ClientExtension) SetTokenExchangePermitted(v bool)`

SetTokenExchangePermitted sets TokenExchangePermitted field to given value.

### HasTokenExchangePermitted

`func (o *ClientExtension) HasTokenExchangePermitted() bool`

HasTokenExchangePermitted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


