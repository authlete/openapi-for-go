# DeviceVerificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take. | [optional] 
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
**GmAction** | Pointer to [**GrantManagementAction**](GrantManagementAction.md) |  | [optional] 
**GrantId** | Pointer to **string** | the value of the &#x60;grant_id&#x60; request parameter of the device authorization request.  The &#x60;grant_id&#x60; request parameter is defined in [Grant Management for OAuth 2.0](https://openid.net/specs/fapi-grant-management.html) , which is supported by Authlete 2.3 and newer versions.  | [optional] 
**Grant** | Pointer to [**Grant**](Grant.md) |  | [optional] 
**GrantSubject** | Pointer to **string** | The subject identifying the user who has given the grant identified by the &#x60;grant_id&#x60; request parameter of the device authorization request.  Authlete 2.3 and newer versions support &lt;a href&#x3D; \&quot;https://openid.net/specs/fapi-grant-management.html\&quot;&gt;Grant Management for OAuth 2.0&lt;/a&gt;. An authorization request may contain a {@code grant_id} request parameter which is defined in the specification. If the value of the request parameter is valid, {@link #getGrantSubject()} will return the subject of the user who has given the grant to the client application. Authorization server implementations may use the value returned from {@link #getGrantSubject()} in order to determine the user to authenticate.  The user your system will authenticate during the authorization process (or has already authenticated) may be different from the user of the grant. The first implementer&#39;s draft of \&quot;Grant Management for OAuth 2.0\&quot; does not mention anything about the case, so the behavior in the case is left to implementations. Authlete will not perform the grant management action when the {@code subject} passed to Authlete does not match the user of the grant.  | [optional] 
**ClientEntityId** | Pointer to **string** | The entity ID of the client.  | [optional] 
**ClientEntityIdUsed** | Pointer to **bool** | Flag which indicates whether the entity ID of the client was used when the request for the access token was made. | [optional] 

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

### GetGmAction

`func (o *DeviceVerificationResponse) GetGmAction() GrantManagementAction`

GetGmAction returns the GmAction field if non-nil, zero value otherwise.

### GetGmActionOk

`func (o *DeviceVerificationResponse) GetGmActionOk() (*GrantManagementAction, bool)`

GetGmActionOk returns a tuple with the GmAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmAction

`func (o *DeviceVerificationResponse) SetGmAction(v GrantManagementAction)`

SetGmAction sets GmAction field to given value.

### HasGmAction

`func (o *DeviceVerificationResponse) HasGmAction() bool`

HasGmAction returns a boolean if a field has been set.

### GetGrantId

`func (o *DeviceVerificationResponse) GetGrantId() string`

GetGrantId returns the GrantId field if non-nil, zero value otherwise.

### GetGrantIdOk

`func (o *DeviceVerificationResponse) GetGrantIdOk() (*string, bool)`

GetGrantIdOk returns a tuple with the GrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantId

`func (o *DeviceVerificationResponse) SetGrantId(v string)`

SetGrantId sets GrantId field to given value.

### HasGrantId

`func (o *DeviceVerificationResponse) HasGrantId() bool`

HasGrantId returns a boolean if a field has been set.

### GetGrant

`func (o *DeviceVerificationResponse) GetGrant() Grant`

GetGrant returns the Grant field if non-nil, zero value otherwise.

### GetGrantOk

`func (o *DeviceVerificationResponse) GetGrantOk() (*Grant, bool)`

GetGrantOk returns a tuple with the Grant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrant

`func (o *DeviceVerificationResponse) SetGrant(v Grant)`

SetGrant sets Grant field to given value.

### HasGrant

`func (o *DeviceVerificationResponse) HasGrant() bool`

HasGrant returns a boolean if a field has been set.

### GetGrantSubject

`func (o *DeviceVerificationResponse) GetGrantSubject() string`

GetGrantSubject returns the GrantSubject field if non-nil, zero value otherwise.

### GetGrantSubjectOk

`func (o *DeviceVerificationResponse) GetGrantSubjectOk() (*string, bool)`

GetGrantSubjectOk returns a tuple with the GrantSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantSubject

`func (o *DeviceVerificationResponse) SetGrantSubject(v string)`

SetGrantSubject sets GrantSubject field to given value.

### HasGrantSubject

`func (o *DeviceVerificationResponse) HasGrantSubject() bool`

HasGrantSubject returns a boolean if a field has been set.

### GetClientEntityId

`func (o *DeviceVerificationResponse) GetClientEntityId() string`

GetClientEntityId returns the ClientEntityId field if non-nil, zero value otherwise.

### GetClientEntityIdOk

`func (o *DeviceVerificationResponse) GetClientEntityIdOk() (*string, bool)`

GetClientEntityIdOk returns a tuple with the ClientEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEntityId

`func (o *DeviceVerificationResponse) SetClientEntityId(v string)`

SetClientEntityId sets ClientEntityId field to given value.

### HasClientEntityId

`func (o *DeviceVerificationResponse) HasClientEntityId() bool`

HasClientEntityId returns a boolean if a field has been set.

### GetClientEntityIdUsed

`func (o *DeviceVerificationResponse) GetClientEntityIdUsed() bool`

GetClientEntityIdUsed returns the ClientEntityIdUsed field if non-nil, zero value otherwise.

### GetClientEntityIdUsedOk

`func (o *DeviceVerificationResponse) GetClientEntityIdUsedOk() (*bool, bool)`

GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEntityIdUsed

`func (o *DeviceVerificationResponse) SetClientEntityIdUsed(v bool)`

SetClientEntityIdUsed sets ClientEntityIdUsed field to given value.

### HasClientEntityIdUsed

`func (o *DeviceVerificationResponse) HasClientEntityIdUsed() bool`

HasClientEntityIdUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


