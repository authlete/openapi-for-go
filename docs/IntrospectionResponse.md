# IntrospectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take. | [optional] 
**ResponseContent** | Pointer to **string** | The content that the authorization server implementation can use as the value of &#x60;WWW-Authenticate&#x60; header on errors.  | [optional] 
**ClientId** | Pointer to **int64** | The client ID. | [optional] 
**ClientIdAlias** | Pointer to **string** | The client ID alias when the token request was made. If the client did not have an alias, this parameter is &#x60;null&#x60;. Also, if the token request was invalid and it failed to identify a client, this parameter is &#x60;null&#x60;.  | [optional] 
**ClientIdAliasUsed** | Pointer to **bool** | The flag which indicates whether the client ID alias was used when the token request was made. &#x60;true&#x60; if the client ID alias was used when the token request was made.  | [optional] 
**ExpiresAt** | Pointer to **int64** | The time at which the access token expires. The value is represented in milliseconds since the Unix epoch (1970-01-01).  | [optional] 
**Subject** | Pointer to **string** | The subject who is associated with the access token. The value of this property is &#x60;null&#x60; if the access token was issued using the flow of [Client Credentials Grant](tools.ietf.org/html/rfc6749#section-4.4).  | [optional] 
**Scopes** | Pointer to **[]string** | The scopes covered by the access token.  | [optional] 
**Existent** | Pointer to **bool** | &#x60;true&#x60; if the access token exists.  | [optional] 
**Usable** | Pointer to **bool** | true&#x60; if the access token is usable (&#x3D; exists and has not expired).  | [optional] 
**Sufficient** | Pointer to **bool** | &#x60;true&#x60; if the access token exists.  | [optional] 
**Refreshable** | Pointer to **bool** | &#x60;true&#x60; if the access token can be refreshed using the associated refresh token which had been issued along with the access token. &#x60;false&#x60; if the refresh token for the access token has expired or the access token has no associated refresh token.  | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) | The extra properties associated with the access token. | [optional] 
**CertificateThumbprint** | Pointer to **string** | The client certificate thumbprint used to validate the access token.  | [optional] 
**Resources** | Pointer to **[]string** | The target resources. This represents the resources specified by the &#x60;resource&#x60; request parameters or by the &#x60;resource&#x60; property in the request object.  See \&quot;Resource Indicators for OAuth 2.0\&quot; for details.  | [optional] 
**AccessTokenResources** | Pointer to **[]string** | The target resources this proeprty holds may be the same as or different from the ones &#x60;resource&#x60; property holds.  In some flows, the initial request and the subsequent token request are sent to different endpoints. Example flows are the Authorization Code Flow, the Refresh Token Flow, the CIBA Ping Mode, the CIBA Poll Mode and the Device Flow. In these flows, not only the initial request but also the subsequent token request can include the &#x60;resource&#x60; request parameters. The purpose of the &#x60;resource&#x60; request parameters in the token request is to narrow the range of the target resources from the original set of target resources requested by the preceding initial request. If narrowing down is performed, the target resources holded by the &#x60;resource&#x60; proeprty and the ones holded by this property are different. This property holds the narrowed set of target resources.  See \&quot;Resource Indicators for OAuth 2.0\&quot; for details.  | [optional] 
**AuthorizationDetails** | Pointer to [**AuthorizationDetails**](AuthorizationDetails.md) |  | [optional] 
**ServiceAttributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of this service that the client application belongs to.  | [optional] 
**ClientAttributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of the client.  | [optional] 
**ScopeDetails** | Pointer to [**[]Scope**](Scope.md) | The scopes property of this class is a list of scope names. The property does not hold information about scope attributes. This scopeDetails property was newly created to convey information about scope attributes.  | [optional] 

## Methods

### NewIntrospectionResponse

`func NewIntrospectionResponse() *IntrospectionResponse`

NewIntrospectionResponse instantiates a new IntrospectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntrospectionResponseWithDefaults

`func NewIntrospectionResponseWithDefaults() *IntrospectionResponse`

NewIntrospectionResponseWithDefaults instantiates a new IntrospectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *IntrospectionResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *IntrospectionResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *IntrospectionResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *IntrospectionResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *IntrospectionResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *IntrospectionResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *IntrospectionResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *IntrospectionResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *IntrospectionResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IntrospectionResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IntrospectionResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *IntrospectionResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResponseContent

`func (o *IntrospectionResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *IntrospectionResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *IntrospectionResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *IntrospectionResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetClientId

`func (o *IntrospectionResponse) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IntrospectionResponse) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IntrospectionResponse) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IntrospectionResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIdAlias

`func (o *IntrospectionResponse) GetClientIdAlias() string`

GetClientIdAlias returns the ClientIdAlias field if non-nil, zero value otherwise.

### GetClientIdAliasOk

`func (o *IntrospectionResponse) GetClientIdAliasOk() (*string, bool)`

GetClientIdAliasOk returns a tuple with the ClientIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAlias

`func (o *IntrospectionResponse) SetClientIdAlias(v string)`

SetClientIdAlias sets ClientIdAlias field to given value.

### HasClientIdAlias

`func (o *IntrospectionResponse) HasClientIdAlias() bool`

HasClientIdAlias returns a boolean if a field has been set.

### GetClientIdAliasUsed

`func (o *IntrospectionResponse) GetClientIdAliasUsed() bool`

GetClientIdAliasUsed returns the ClientIdAliasUsed field if non-nil, zero value otherwise.

### GetClientIdAliasUsedOk

`func (o *IntrospectionResponse) GetClientIdAliasUsedOk() (*bool, bool)`

GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasUsed

`func (o *IntrospectionResponse) SetClientIdAliasUsed(v bool)`

SetClientIdAliasUsed sets ClientIdAliasUsed field to given value.

### HasClientIdAliasUsed

`func (o *IntrospectionResponse) HasClientIdAliasUsed() bool`

HasClientIdAliasUsed returns a boolean if a field has been set.

### GetExpiresAt

`func (o *IntrospectionResponse) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *IntrospectionResponse) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *IntrospectionResponse) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *IntrospectionResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetSubject

`func (o *IntrospectionResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *IntrospectionResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *IntrospectionResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *IntrospectionResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetScopes

`func (o *IntrospectionResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IntrospectionResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IntrospectionResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *IntrospectionResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetExistent

`func (o *IntrospectionResponse) GetExistent() bool`

GetExistent returns the Existent field if non-nil, zero value otherwise.

### GetExistentOk

`func (o *IntrospectionResponse) GetExistentOk() (*bool, bool)`

GetExistentOk returns a tuple with the Existent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistent

`func (o *IntrospectionResponse) SetExistent(v bool)`

SetExistent sets Existent field to given value.

### HasExistent

`func (o *IntrospectionResponse) HasExistent() bool`

HasExistent returns a boolean if a field has been set.

### GetUsable

`func (o *IntrospectionResponse) GetUsable() bool`

GetUsable returns the Usable field if non-nil, zero value otherwise.

### GetUsableOk

`func (o *IntrospectionResponse) GetUsableOk() (*bool, bool)`

GetUsableOk returns a tuple with the Usable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsable

`func (o *IntrospectionResponse) SetUsable(v bool)`

SetUsable sets Usable field to given value.

### HasUsable

`func (o *IntrospectionResponse) HasUsable() bool`

HasUsable returns a boolean if a field has been set.

### GetSufficient

`func (o *IntrospectionResponse) GetSufficient() bool`

GetSufficient returns the Sufficient field if non-nil, zero value otherwise.

### GetSufficientOk

`func (o *IntrospectionResponse) GetSufficientOk() (*bool, bool)`

GetSufficientOk returns a tuple with the Sufficient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSufficient

`func (o *IntrospectionResponse) SetSufficient(v bool)`

SetSufficient sets Sufficient field to given value.

### HasSufficient

`func (o *IntrospectionResponse) HasSufficient() bool`

HasSufficient returns a boolean if a field has been set.

### GetRefreshable

`func (o *IntrospectionResponse) GetRefreshable() bool`

GetRefreshable returns the Refreshable field if non-nil, zero value otherwise.

### GetRefreshableOk

`func (o *IntrospectionResponse) GetRefreshableOk() (*bool, bool)`

GetRefreshableOk returns a tuple with the Refreshable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshable

`func (o *IntrospectionResponse) SetRefreshable(v bool)`

SetRefreshable sets Refreshable field to given value.

### HasRefreshable

`func (o *IntrospectionResponse) HasRefreshable() bool`

HasRefreshable returns a boolean if a field has been set.

### GetProperties

`func (o *IntrospectionResponse) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *IntrospectionResponse) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *IntrospectionResponse) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *IntrospectionResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetCertificateThumbprint

`func (o *IntrospectionResponse) GetCertificateThumbprint() string`

GetCertificateThumbprint returns the CertificateThumbprint field if non-nil, zero value otherwise.

### GetCertificateThumbprintOk

`func (o *IntrospectionResponse) GetCertificateThumbprintOk() (*string, bool)`

GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateThumbprint

`func (o *IntrospectionResponse) SetCertificateThumbprint(v string)`

SetCertificateThumbprint sets CertificateThumbprint field to given value.

### HasCertificateThumbprint

`func (o *IntrospectionResponse) HasCertificateThumbprint() bool`

HasCertificateThumbprint returns a boolean if a field has been set.

### GetResources

`func (o *IntrospectionResponse) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *IntrospectionResponse) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *IntrospectionResponse) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *IntrospectionResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetAccessTokenResources

`func (o *IntrospectionResponse) GetAccessTokenResources() []string`

GetAccessTokenResources returns the AccessTokenResources field if non-nil, zero value otherwise.

### GetAccessTokenResourcesOk

`func (o *IntrospectionResponse) GetAccessTokenResourcesOk() (*[]string, bool)`

GetAccessTokenResourcesOk returns a tuple with the AccessTokenResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenResources

`func (o *IntrospectionResponse) SetAccessTokenResources(v []string)`

SetAccessTokenResources sets AccessTokenResources field to given value.

### HasAccessTokenResources

`func (o *IntrospectionResponse) HasAccessTokenResources() bool`

HasAccessTokenResources returns a boolean if a field has been set.

### GetAuthorizationDetails

`func (o *IntrospectionResponse) GetAuthorizationDetails() AuthorizationDetails`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *IntrospectionResponse) GetAuthorizationDetailsOk() (*AuthorizationDetails, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *IntrospectionResponse) SetAuthorizationDetails(v AuthorizationDetails)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *IntrospectionResponse) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.

### GetServiceAttributes

`func (o *IntrospectionResponse) GetServiceAttributes() []Pair`

GetServiceAttributes returns the ServiceAttributes field if non-nil, zero value otherwise.

### GetServiceAttributesOk

`func (o *IntrospectionResponse) GetServiceAttributesOk() (*[]Pair, bool)`

GetServiceAttributesOk returns a tuple with the ServiceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAttributes

`func (o *IntrospectionResponse) SetServiceAttributes(v []Pair)`

SetServiceAttributes sets ServiceAttributes field to given value.

### HasServiceAttributes

`func (o *IntrospectionResponse) HasServiceAttributes() bool`

HasServiceAttributes returns a boolean if a field has been set.

### GetClientAttributes

`func (o *IntrospectionResponse) GetClientAttributes() []Pair`

GetClientAttributes returns the ClientAttributes field if non-nil, zero value otherwise.

### GetClientAttributesOk

`func (o *IntrospectionResponse) GetClientAttributesOk() (*[]Pair, bool)`

GetClientAttributesOk returns a tuple with the ClientAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAttributes

`func (o *IntrospectionResponse) SetClientAttributes(v []Pair)`

SetClientAttributes sets ClientAttributes field to given value.

### HasClientAttributes

`func (o *IntrospectionResponse) HasClientAttributes() bool`

HasClientAttributes returns a boolean if a field has been set.

### GetScopeDetails

`func (o *IntrospectionResponse) GetScopeDetails() []Scope`

GetScopeDetails returns the ScopeDetails field if non-nil, zero value otherwise.

### GetScopeDetailsOk

`func (o *IntrospectionResponse) GetScopeDetailsOk() (*[]Scope, bool)`

GetScopeDetailsOk returns a tuple with the ScopeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeDetails

`func (o *IntrospectionResponse) SetScopeDetails(v []Scope)`

SetScopeDetails sets ScopeDetails field to given value.

### HasScopeDetails

`func (o *IntrospectionResponse) HasScopeDetails() bool`

HasScopeDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


