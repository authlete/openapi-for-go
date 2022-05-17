# DeviceVerificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take. | [optional] 
**ResponseContent** | Pointer to **string** | The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of &#x60;action&#x60; parameter.  | [optional] 
**ClientId** | Pointer to **int64** | The client ID of the client application to which the user code has been issued.  | [optional] 
**ClientIdAlias** | Pointer to **string** | The client ID alias of the client application to which the user code has been issued.  | [optional] 
**ClientIdAliasUsed** | Pointer to **bool** | &#x60;true&#x60; if the value of the &#x60;client_id&#x60; request parameter included in the device authorization request is the client ID alias. &#x60;false&#x60; if the value is the original numeric client ID.  | [optional] 
**ClientName** | Pointer to **string** | The name of the client application to which the user code has been issued.  | [optional] 
**Scopes** | Pointer to [**[]Scope**](Scope.md) | The scopes requested by the device authorization request.  Note that &#x60;description&#x60; property and &#x60;descriptions&#x60; property of each scope object in the array contained in this property is always null even if descriptions of the scopes are registered.  | [optional] 
**ClaimNames** | Pointer to **[]string** | The names of the claims which were requested indirectly via some special scopes. See [5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) in OpenID Connect Core 1.0 for details.  This property is always &#x60;null&#x60; if the &#x60;scope&#x60; request parameter of the device authorization request does not include the &#x60;openid&#x60; scope even if special scopes (such as &#x60;profile&#x60;) are included in the request (unless the openid scope is included in the default set of scopes which is used when the &#x60;scope&#x60; request parameter is omitted).  | [optional] 
**Acrs** | Pointer to **[]string** | The list of ACR values requested by the device authorization request.  | [optional] 
**Resources** | Pointer to **[]string** | The resources specified by the &#x60;resource&#x60; request parameters or by the &#x60;resource&#x60; property in the request object. If both are given, the values in the request object should be set. See \&quot;Resource Indicators for OAuth 2.0\&quot; for details.  | [optional] 
**AuthorizationDetails** | Pointer to [**AuthorizationDetails**](AuthorizationDetails.md) |  | [optional] 
**ServiceAttributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of this service that the client application belongs to.  | [optional] 
**ClientAttributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of the client.  | [optional] 
**DynamicScopes** | Pointer to [**[]DynamicScope**](DynamicScope.md) | The dynamic scopes which the client application requested by the scope request parameter.  | [optional] 
**ExpiresAt** | Pointer to **int64** | Get the date in milliseconds since the Unix epoch (1970-01-01) at which the user code will expire.  | [optional] 

## Methods

### NewDeviceVerificationResponse

`func NewDeviceVerificationResponse() *DeviceVerificationResponse`

NewDeviceVerificationResponse instantiates a new DeviceVerificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceVerificationResponseWithDefaults

`func NewDeviceVerificationResponseWithDefaults() *DeviceVerificationResponse`

NewDeviceVerificationResponseWithDefaults instantiates a new DeviceVerificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *DeviceVerificationResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *DeviceVerificationResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *DeviceVerificationResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *DeviceVerificationResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *DeviceVerificationResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *DeviceVerificationResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *DeviceVerificationResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *DeviceVerificationResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *DeviceVerificationResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeviceVerificationResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeviceVerificationResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DeviceVerificationResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResponseContent

`func (o *DeviceVerificationResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *DeviceVerificationResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *DeviceVerificationResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *DeviceVerificationResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetClientId

`func (o *DeviceVerificationResponse) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *DeviceVerificationResponse) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *DeviceVerificationResponse) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *DeviceVerificationResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIdAlias

`func (o *DeviceVerificationResponse) GetClientIdAlias() string`

GetClientIdAlias returns the ClientIdAlias field if non-nil, zero value otherwise.

### GetClientIdAliasOk

`func (o *DeviceVerificationResponse) GetClientIdAliasOk() (*string, bool)`

GetClientIdAliasOk returns a tuple with the ClientIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAlias

`func (o *DeviceVerificationResponse) SetClientIdAlias(v string)`

SetClientIdAlias sets ClientIdAlias field to given value.

### HasClientIdAlias

`func (o *DeviceVerificationResponse) HasClientIdAlias() bool`

HasClientIdAlias returns a boolean if a field has been set.

### GetClientIdAliasUsed

`func (o *DeviceVerificationResponse) GetClientIdAliasUsed() bool`

GetClientIdAliasUsed returns the ClientIdAliasUsed field if non-nil, zero value otherwise.

### GetClientIdAliasUsedOk

`func (o *DeviceVerificationResponse) GetClientIdAliasUsedOk() (*bool, bool)`

GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasUsed

`func (o *DeviceVerificationResponse) SetClientIdAliasUsed(v bool)`

SetClientIdAliasUsed sets ClientIdAliasUsed field to given value.

### HasClientIdAliasUsed

`func (o *DeviceVerificationResponse) HasClientIdAliasUsed() bool`

HasClientIdAliasUsed returns a boolean if a field has been set.

### GetClientName

`func (o *DeviceVerificationResponse) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *DeviceVerificationResponse) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *DeviceVerificationResponse) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *DeviceVerificationResponse) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetScopes

`func (o *DeviceVerificationResponse) GetScopes() []Scope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *DeviceVerificationResponse) GetScopesOk() (*[]Scope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *DeviceVerificationResponse) SetScopes(v []Scope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *DeviceVerificationResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetClaimNames

`func (o *DeviceVerificationResponse) GetClaimNames() []string`

GetClaimNames returns the ClaimNames field if non-nil, zero value otherwise.

### GetClaimNamesOk

`func (o *DeviceVerificationResponse) GetClaimNamesOk() (*[]string, bool)`

GetClaimNamesOk returns a tuple with the ClaimNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimNames

`func (o *DeviceVerificationResponse) SetClaimNames(v []string)`

SetClaimNames sets ClaimNames field to given value.

### HasClaimNames

`func (o *DeviceVerificationResponse) HasClaimNames() bool`

HasClaimNames returns a boolean if a field has been set.

### GetAcrs

`func (o *DeviceVerificationResponse) GetAcrs() []string`

GetAcrs returns the Acrs field if non-nil, zero value otherwise.

### GetAcrsOk

`func (o *DeviceVerificationResponse) GetAcrsOk() (*[]string, bool)`

GetAcrsOk returns a tuple with the Acrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrs

`func (o *DeviceVerificationResponse) SetAcrs(v []string)`

SetAcrs sets Acrs field to given value.

### HasAcrs

`func (o *DeviceVerificationResponse) HasAcrs() bool`

HasAcrs returns a boolean if a field has been set.

### GetResources

`func (o *DeviceVerificationResponse) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DeviceVerificationResponse) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DeviceVerificationResponse) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *DeviceVerificationResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetAuthorizationDetails

`func (o *DeviceVerificationResponse) GetAuthorizationDetails() AuthorizationDetails`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *DeviceVerificationResponse) GetAuthorizationDetailsOk() (*AuthorizationDetails, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *DeviceVerificationResponse) SetAuthorizationDetails(v AuthorizationDetails)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *DeviceVerificationResponse) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.

### GetServiceAttributes

`func (o *DeviceVerificationResponse) GetServiceAttributes() []Pair`

GetServiceAttributes returns the ServiceAttributes field if non-nil, zero value otherwise.

### GetServiceAttributesOk

`func (o *DeviceVerificationResponse) GetServiceAttributesOk() (*[]Pair, bool)`

GetServiceAttributesOk returns a tuple with the ServiceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAttributes

`func (o *DeviceVerificationResponse) SetServiceAttributes(v []Pair)`

SetServiceAttributes sets ServiceAttributes field to given value.

### HasServiceAttributes

`func (o *DeviceVerificationResponse) HasServiceAttributes() bool`

HasServiceAttributes returns a boolean if a field has been set.

### GetClientAttributes

`func (o *DeviceVerificationResponse) GetClientAttributes() []Pair`

GetClientAttributes returns the ClientAttributes field if non-nil, zero value otherwise.

### GetClientAttributesOk

`func (o *DeviceVerificationResponse) GetClientAttributesOk() (*[]Pair, bool)`

GetClientAttributesOk returns a tuple with the ClientAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAttributes

`func (o *DeviceVerificationResponse) SetClientAttributes(v []Pair)`

SetClientAttributes sets ClientAttributes field to given value.

### HasClientAttributes

`func (o *DeviceVerificationResponse) HasClientAttributes() bool`

HasClientAttributes returns a boolean if a field has been set.

### GetDynamicScopes

`func (o *DeviceVerificationResponse) GetDynamicScopes() []DynamicScope`

GetDynamicScopes returns the DynamicScopes field if non-nil, zero value otherwise.

### GetDynamicScopesOk

`func (o *DeviceVerificationResponse) GetDynamicScopesOk() (*[]DynamicScope, bool)`

GetDynamicScopesOk returns a tuple with the DynamicScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicScopes

`func (o *DeviceVerificationResponse) SetDynamicScopes(v []DynamicScope)`

SetDynamicScopes sets DynamicScopes field to given value.

### HasDynamicScopes

`func (o *DeviceVerificationResponse) HasDynamicScopes() bool`

HasDynamicScopes returns a boolean if a field has been set.

### GetExpiresAt

`func (o *DeviceVerificationResponse) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DeviceVerificationResponse) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DeviceVerificationResponse) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DeviceVerificationResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


