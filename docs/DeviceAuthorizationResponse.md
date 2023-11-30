# DeviceAuthorizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take. | [optional] 
**ResponseContent** | Pointer to **string** | The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of &#x60;action&#x60; parameter.  | [optional] 
**ClientId** | Pointer to **int64** | The client ID of the client application that has made the device authorization request.  | [optional] 
**ClientIdAlias** | Pointer to **string** | The client ID alias of the client application that has made the device authorization request.  | [optional] 
**ClientIdAliasUsed** | Pointer to **bool** | &#x60;true&#x60; if the value of the client_id request parameter included in the device authorization request is the client ID alias. &#x60;false&#x60; if the value is the original numeric client ID.  | [optional] 
**ClientName** | Pointer to **string** | The name of the client application which has made the device authorization request.  | [optional] 
**ClientAuthMethod** | Pointer to **string** | The client authentication method that should be performed at the device authorization endpoint.  | [optional] 
**Scopes** | Pointer to **[]string** | The scopes requested by the device authorization request.  Basically, this property holds the value of the scope request parameter in the device authorization request. However, because unregistered scopes are dropped on Authlete side, if the &#x60;scope&#x60; request parameter contains unknown scopes, the list returned by this property becomes different from the value of the &#x60;scope&#x60; request parameter.  Note that &#x60;description&#x60; property and &#x60;descriptions&#x60; property of each scope object in the array contained in this property is always &#x60;null&#x60; even if descriptions of the scopes are registered.  | [optional] 
**ClaimNames** | Pointer to **[]string** | The names of the claims which were requested indirectly via some special scopes. See [5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) in OpenID Connect Core 1.0 for details.  | [optional] 
**Acrs** | Pointer to **[]string** | The list of ACR values requested by the device authorization request.  Basically, this property holds the value of the &#x60;acr_values&#x60; request parameter in the device authorization request. However, because unsupported ACR values are dropped on Authlete side, if the &#x60;acr_values&#x60; request parameter contains unrecognized ACR values, the list returned by this property becomes different from the value of the &#x60;acr_values&#x60; request parameter.  | [optional] 
**DeviceCode** | Pointer to **string** | The device verification code. This corresponds to the &#x60;device_code&#x60; property in the response to the client.  | [optional] 
**UserCode** | Pointer to **string** | The end-user verification code. This corresponds to the &#x60;user_code&#x60; property in the response to the client.  | [optional] 
**VerificationUri** | Pointer to **string** | The end-user verification URI. This corresponds to the &#x60;verification_uri&#x60; property in the response to the client.  | [optional] 
**VerificationUriComplete** | Pointer to **string** | The end-user verification URI that includes the end-user verification code. This corresponds to the &#x60;verification_uri_complete&#x60; property in the response to the client.  | [optional] 
**ExpiresIn** | Pointer to **int32** | The duration of the device verification code in seconds. This corresponds to the &#x60;expires_in&#x60; property in the response to the client.  | [optional] 
**Interval** | Pointer to **int32** | The minimum amount of time in seconds that the client must wait for between polling requests to the token endpoint. This corresponds to the &#x60;interval&#x60; property in the response to the client.  | [optional] 
**Warnings** | Pointer to **[]string** | The warnings raised during processing the backchannel authentication request.  | [optional] 
**Resources** | Pointer to **[]string** | The resources specified by the &#x60;resource&#x60; request parameters. See \&quot;Resource Indicators for OAuth 2.0\&quot; for details.  | [optional] 
**AuthorizationDetails** | Pointer to [**AuthorizationDetails**](AuthorizationDetails.md) |  | [optional] 
**ServiceAttributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of this service that the client application belongs to.  | [optional] 
**ClientAttributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of the client.  | [optional] 
**DynamicScopes** | Pointer to [**[]DynamicScope**](DynamicScope.md) | The dynamic scopes which the client application requested by the scope request parameter.  | [optional] 
**GmAction** | Pointer to [**GrantManagementAction**](GrantManagementAction.md) |  | [optional] 
**GrantId** | Pointer to **string** | the value of the &#x60;grant_id&#x60; request parameter of the device authorization request.  The &#x60;grant_id&#x60; request parameter is defined in [Grant Management for OAuth 2.0](https://openid.net/specs/fapi-grant-management.html) , which is supported by Authlete 2.3 and newer versions.  | [optional] 
**Grant** | Pointer to [**Grant**](Grant.md) |  | [optional] 
**GrantSubject** | Pointer to **string** | The subject identifying the user who has given the grant identified by the &#x60;grant_id&#x60; request parameter of the device authorization request.  Authlete 2.3 and newer versions support &lt;a href&#x3D; \&quot;https://openid.net/specs/fapi-grant-management.html\&quot;&gt;Grant Management for OAuth 2.0&lt;/a&gt;. An authorization request may contain a &#x60;grant_id&#x60; request parameter which is defined in the specification. If the value of the request parameter is valid, {@link #getGrantSubject()} will return the subject of the user who has given the grant to the client application. Authorization server implementations may use the value returned from {@link #getGrantSubject()} in order to determine the user to authenticate.  The user your system will authenticate during the authorization process (or has already authenticated) may be different from the user of the grant. The first implementer&#39;s draft of \&quot;Grant Management for OAuth 2.0\&quot; does not mention anything about the case, so the behavior in the case is left to implementations. Authlete will not perform the grant management action when the &#x60;subject&#x60; passed to Authlete does not match the user of the grant.  | [optional] 
**ClientEntityId** | Pointer to **string** | The entity ID of the client.  | [optional] 
**ClientEntityIdUsed** | Pointer to **bool** | Flag which indicates whether the entity ID of the client was used when the request for the access token was made. | [optional] 

## Methods

### NewDeviceAuthorizationResponse

`func NewDeviceAuthorizationResponse() *DeviceAuthorizationResponse`

NewDeviceAuthorizationResponse instantiates a new DeviceAuthorizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthorizationResponseWithDefaults

`func NewDeviceAuthorizationResponseWithDefaults() *DeviceAuthorizationResponse`

NewDeviceAuthorizationResponseWithDefaults instantiates a new DeviceAuthorizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *DeviceAuthorizationResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *DeviceAuthorizationResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *DeviceAuthorizationResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *DeviceAuthorizationResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *DeviceAuthorizationResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *DeviceAuthorizationResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *DeviceAuthorizationResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *DeviceAuthorizationResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *DeviceAuthorizationResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeviceAuthorizationResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeviceAuthorizationResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DeviceAuthorizationResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResponseContent

`func (o *DeviceAuthorizationResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *DeviceAuthorizationResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *DeviceAuthorizationResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *DeviceAuthorizationResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetClientId

`func (o *DeviceAuthorizationResponse) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *DeviceAuthorizationResponse) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *DeviceAuthorizationResponse) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *DeviceAuthorizationResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIdAlias

`func (o *DeviceAuthorizationResponse) GetClientIdAlias() string`

GetClientIdAlias returns the ClientIdAlias field if non-nil, zero value otherwise.

### GetClientIdAliasOk

`func (o *DeviceAuthorizationResponse) GetClientIdAliasOk() (*string, bool)`

GetClientIdAliasOk returns a tuple with the ClientIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAlias

`func (o *DeviceAuthorizationResponse) SetClientIdAlias(v string)`

SetClientIdAlias sets ClientIdAlias field to given value.

### HasClientIdAlias

`func (o *DeviceAuthorizationResponse) HasClientIdAlias() bool`

HasClientIdAlias returns a boolean if a field has been set.

### GetClientIdAliasUsed

`func (o *DeviceAuthorizationResponse) GetClientIdAliasUsed() bool`

GetClientIdAliasUsed returns the ClientIdAliasUsed field if non-nil, zero value otherwise.

### GetClientIdAliasUsedOk

`func (o *DeviceAuthorizationResponse) GetClientIdAliasUsedOk() (*bool, bool)`

GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasUsed

`func (o *DeviceAuthorizationResponse) SetClientIdAliasUsed(v bool)`

SetClientIdAliasUsed sets ClientIdAliasUsed field to given value.

### HasClientIdAliasUsed

`func (o *DeviceAuthorizationResponse) HasClientIdAliasUsed() bool`

HasClientIdAliasUsed returns a boolean if a field has been set.

### GetClientName

`func (o *DeviceAuthorizationResponse) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *DeviceAuthorizationResponse) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *DeviceAuthorizationResponse) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *DeviceAuthorizationResponse) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientAuthMethod

`func (o *DeviceAuthorizationResponse) GetClientAuthMethod() string`

GetClientAuthMethod returns the ClientAuthMethod field if non-nil, zero value otherwise.

### GetClientAuthMethodOk

`func (o *DeviceAuthorizationResponse) GetClientAuthMethodOk() (*string, bool)`

GetClientAuthMethodOk returns a tuple with the ClientAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthMethod

`func (o *DeviceAuthorizationResponse) SetClientAuthMethod(v string)`

SetClientAuthMethod sets ClientAuthMethod field to given value.

### HasClientAuthMethod

`func (o *DeviceAuthorizationResponse) HasClientAuthMethod() bool`

HasClientAuthMethod returns a boolean if a field has been set.

### GetScopes

`func (o *DeviceAuthorizationResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *DeviceAuthorizationResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *DeviceAuthorizationResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *DeviceAuthorizationResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetClaimNames

`func (o *DeviceAuthorizationResponse) GetClaimNames() []string`

GetClaimNames returns the ClaimNames field if non-nil, zero value otherwise.

### GetClaimNamesOk

`func (o *DeviceAuthorizationResponse) GetClaimNamesOk() (*[]string, bool)`

GetClaimNamesOk returns a tuple with the ClaimNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimNames

`func (o *DeviceAuthorizationResponse) SetClaimNames(v []string)`

SetClaimNames sets ClaimNames field to given value.

### HasClaimNames

`func (o *DeviceAuthorizationResponse) HasClaimNames() bool`

HasClaimNames returns a boolean if a field has been set.

### GetAcrs

`func (o *DeviceAuthorizationResponse) GetAcrs() []string`

GetAcrs returns the Acrs field if non-nil, zero value otherwise.

### GetAcrsOk

`func (o *DeviceAuthorizationResponse) GetAcrsOk() (*[]string, bool)`

GetAcrsOk returns a tuple with the Acrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrs

`func (o *DeviceAuthorizationResponse) SetAcrs(v []string)`

SetAcrs sets Acrs field to given value.

### HasAcrs

`func (o *DeviceAuthorizationResponse) HasAcrs() bool`

HasAcrs returns a boolean if a field has been set.

### GetDeviceCode

`func (o *DeviceAuthorizationResponse) GetDeviceCode() string`

GetDeviceCode returns the DeviceCode field if non-nil, zero value otherwise.

### GetDeviceCodeOk

`func (o *DeviceAuthorizationResponse) GetDeviceCodeOk() (*string, bool)`

GetDeviceCodeOk returns a tuple with the DeviceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCode

`func (o *DeviceAuthorizationResponse) SetDeviceCode(v string)`

SetDeviceCode sets DeviceCode field to given value.

### HasDeviceCode

`func (o *DeviceAuthorizationResponse) HasDeviceCode() bool`

HasDeviceCode returns a boolean if a field has been set.

### GetUserCode

`func (o *DeviceAuthorizationResponse) GetUserCode() string`

GetUserCode returns the UserCode field if non-nil, zero value otherwise.

### GetUserCodeOk

`func (o *DeviceAuthorizationResponse) GetUserCodeOk() (*string, bool)`

GetUserCodeOk returns a tuple with the UserCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCode

`func (o *DeviceAuthorizationResponse) SetUserCode(v string)`

SetUserCode sets UserCode field to given value.

### HasUserCode

`func (o *DeviceAuthorizationResponse) HasUserCode() bool`

HasUserCode returns a boolean if a field has been set.

### GetVerificationUri

`func (o *DeviceAuthorizationResponse) GetVerificationUri() string`

GetVerificationUri returns the VerificationUri field if non-nil, zero value otherwise.

### GetVerificationUriOk

`func (o *DeviceAuthorizationResponse) GetVerificationUriOk() (*string, bool)`

GetVerificationUriOk returns a tuple with the VerificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationUri

`func (o *DeviceAuthorizationResponse) SetVerificationUri(v string)`

SetVerificationUri sets VerificationUri field to given value.

### HasVerificationUri

`func (o *DeviceAuthorizationResponse) HasVerificationUri() bool`

HasVerificationUri returns a boolean if a field has been set.

### GetVerificationUriComplete

`func (o *DeviceAuthorizationResponse) GetVerificationUriComplete() string`

GetVerificationUriComplete returns the VerificationUriComplete field if non-nil, zero value otherwise.

### GetVerificationUriCompleteOk

`func (o *DeviceAuthorizationResponse) GetVerificationUriCompleteOk() (*string, bool)`

GetVerificationUriCompleteOk returns a tuple with the VerificationUriComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationUriComplete

`func (o *DeviceAuthorizationResponse) SetVerificationUriComplete(v string)`

SetVerificationUriComplete sets VerificationUriComplete field to given value.

### HasVerificationUriComplete

`func (o *DeviceAuthorizationResponse) HasVerificationUriComplete() bool`

HasVerificationUriComplete returns a boolean if a field has been set.

### GetExpiresIn

`func (o *DeviceAuthorizationResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *DeviceAuthorizationResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *DeviceAuthorizationResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *DeviceAuthorizationResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetInterval

`func (o *DeviceAuthorizationResponse) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *DeviceAuthorizationResponse) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *DeviceAuthorizationResponse) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *DeviceAuthorizationResponse) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetWarnings

`func (o *DeviceAuthorizationResponse) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DeviceAuthorizationResponse) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DeviceAuthorizationResponse) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DeviceAuthorizationResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetResources

`func (o *DeviceAuthorizationResponse) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DeviceAuthorizationResponse) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DeviceAuthorizationResponse) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *DeviceAuthorizationResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetAuthorizationDetails

`func (o *DeviceAuthorizationResponse) GetAuthorizationDetails() AuthorizationDetails`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *DeviceAuthorizationResponse) GetAuthorizationDetailsOk() (*AuthorizationDetails, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *DeviceAuthorizationResponse) SetAuthorizationDetails(v AuthorizationDetails)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *DeviceAuthorizationResponse) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.

### GetServiceAttributes

`func (o *DeviceAuthorizationResponse) GetServiceAttributes() []Pair`

GetServiceAttributes returns the ServiceAttributes field if non-nil, zero value otherwise.

### GetServiceAttributesOk

`func (o *DeviceAuthorizationResponse) GetServiceAttributesOk() (*[]Pair, bool)`

GetServiceAttributesOk returns a tuple with the ServiceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAttributes

`func (o *DeviceAuthorizationResponse) SetServiceAttributes(v []Pair)`

SetServiceAttributes sets ServiceAttributes field to given value.

### HasServiceAttributes

`func (o *DeviceAuthorizationResponse) HasServiceAttributes() bool`

HasServiceAttributes returns a boolean if a field has been set.

### GetClientAttributes

`func (o *DeviceAuthorizationResponse) GetClientAttributes() []Pair`

GetClientAttributes returns the ClientAttributes field if non-nil, zero value otherwise.

### GetClientAttributesOk

`func (o *DeviceAuthorizationResponse) GetClientAttributesOk() (*[]Pair, bool)`

GetClientAttributesOk returns a tuple with the ClientAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAttributes

`func (o *DeviceAuthorizationResponse) SetClientAttributes(v []Pair)`

SetClientAttributes sets ClientAttributes field to given value.

### HasClientAttributes

`func (o *DeviceAuthorizationResponse) HasClientAttributes() bool`

HasClientAttributes returns a boolean if a field has been set.

### GetDynamicScopes

`func (o *DeviceAuthorizationResponse) GetDynamicScopes() []DynamicScope`

GetDynamicScopes returns the DynamicScopes field if non-nil, zero value otherwise.

### GetDynamicScopesOk

`func (o *DeviceAuthorizationResponse) GetDynamicScopesOk() (*[]DynamicScope, bool)`

GetDynamicScopesOk returns a tuple with the DynamicScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicScopes

`func (o *DeviceAuthorizationResponse) SetDynamicScopes(v []DynamicScope)`

SetDynamicScopes sets DynamicScopes field to given value.

### HasDynamicScopes

`func (o *DeviceAuthorizationResponse) HasDynamicScopes() bool`

HasDynamicScopes returns a boolean if a field has been set.

### GetGmAction

`func (o *DeviceAuthorizationResponse) GetGmAction() GrantManagementAction`

GetGmAction returns the GmAction field if non-nil, zero value otherwise.

### GetGmActionOk

`func (o *DeviceAuthorizationResponse) GetGmActionOk() (*GrantManagementAction, bool)`

GetGmActionOk returns a tuple with the GmAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmAction

`func (o *DeviceAuthorizationResponse) SetGmAction(v GrantManagementAction)`

SetGmAction sets GmAction field to given value.

### HasGmAction

`func (o *DeviceAuthorizationResponse) HasGmAction() bool`

HasGmAction returns a boolean if a field has been set.

### GetGrantId

`func (o *DeviceAuthorizationResponse) GetGrantId() string`

GetGrantId returns the GrantId field if non-nil, zero value otherwise.

### GetGrantIdOk

`func (o *DeviceAuthorizationResponse) GetGrantIdOk() (*string, bool)`

GetGrantIdOk returns a tuple with the GrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantId

`func (o *DeviceAuthorizationResponse) SetGrantId(v string)`

SetGrantId sets GrantId field to given value.

### HasGrantId

`func (o *DeviceAuthorizationResponse) HasGrantId() bool`

HasGrantId returns a boolean if a field has been set.

### GetGrant

`func (o *DeviceAuthorizationResponse) GetGrant() Grant`

GetGrant returns the Grant field if non-nil, zero value otherwise.

### GetGrantOk

`func (o *DeviceAuthorizationResponse) GetGrantOk() (*Grant, bool)`

GetGrantOk returns a tuple with the Grant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrant

`func (o *DeviceAuthorizationResponse) SetGrant(v Grant)`

SetGrant sets Grant field to given value.

### HasGrant

`func (o *DeviceAuthorizationResponse) HasGrant() bool`

HasGrant returns a boolean if a field has been set.

### GetGrantSubject

`func (o *DeviceAuthorizationResponse) GetGrantSubject() string`

GetGrantSubject returns the GrantSubject field if non-nil, zero value otherwise.

### GetGrantSubjectOk

`func (o *DeviceAuthorizationResponse) GetGrantSubjectOk() (*string, bool)`

GetGrantSubjectOk returns a tuple with the GrantSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantSubject

`func (o *DeviceAuthorizationResponse) SetGrantSubject(v string)`

SetGrantSubject sets GrantSubject field to given value.

### HasGrantSubject

`func (o *DeviceAuthorizationResponse) HasGrantSubject() bool`

HasGrantSubject returns a boolean if a field has been set.

### GetClientEntityId

`func (o *DeviceAuthorizationResponse) GetClientEntityId() string`

GetClientEntityId returns the ClientEntityId field if non-nil, zero value otherwise.

### GetClientEntityIdOk

`func (o *DeviceAuthorizationResponse) GetClientEntityIdOk() (*string, bool)`

GetClientEntityIdOk returns a tuple with the ClientEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEntityId

`func (o *DeviceAuthorizationResponse) SetClientEntityId(v string)`

SetClientEntityId sets ClientEntityId field to given value.

### HasClientEntityId

`func (o *DeviceAuthorizationResponse) HasClientEntityId() bool`

HasClientEntityId returns a boolean if a field has been set.

### GetClientEntityIdUsed

`func (o *DeviceAuthorizationResponse) GetClientEntityIdUsed() bool`

GetClientEntityIdUsed returns the ClientEntityIdUsed field if non-nil, zero value otherwise.

### GetClientEntityIdUsedOk

`func (o *DeviceAuthorizationResponse) GetClientEntityIdUsedOk() (*bool, bool)`

GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEntityIdUsed

`func (o *DeviceAuthorizationResponse) SetClientEntityIdUsed(v bool)`

SetClientEntityIdUsed sets ClientEntityIdUsed field to given value.

### HasClientEntityIdUsed

`func (o *DeviceAuthorizationResponse) HasClientEntityIdUsed() bool`

HasClientEntityIdUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


