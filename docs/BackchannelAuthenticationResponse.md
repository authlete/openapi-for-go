# BackchannelAuthenticationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take. | [optional] 
**ResponseContent** | Pointer to **string** | The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of &#x60;action&#x60; parameter.  | [optional] 
**ClientId** | Pointer to **int64** | The client ID of the client application that has made the backchannel authentication request.  | [optional] 
**ClientIdAlias** | Pointer to **string** | The client ID alias of the client application that has made the backchannel authentication request.  | [optional] 
**ClientIdAliasUsed** | Pointer to **bool** | &#x60;true&#x60; if the value of the client_id request parameter included in the backchannel authentication request is the client ID alias. &#x60;false&#x60; if the value is the original numeric client ID.  | [optional] 
**ClientName** | Pointer to **string** | The name of the client application which has made the backchannel authentication request.  | [optional] 
**Scopes** | Pointer to [**[]Scope**](Scope.md) | The scopes requested by the backchannel authentication request.  Basically, this property holds the value of the &#x60;scope&#x60; request parameter in the backchannel authentication request. However, because unregistered scopes are dropped on Authlete side, if the &#x60;scope&#x60; request parameter contains unknown scopes, the list returned by this property becomes different from the value of the &#x60;scope&#x60; request parameter.  Note that &#x60;description&#x60; property and &#x60;descriptions&#x60; property of each &#x60;scope&#x60; object in the array contained in this property is always null even if descriptions of the scopes are registered.  | [optional] 
**ClaimNames** | Pointer to **[]string** | The names of the claims which were requested indirectly via some special scopes. See [5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) in OpenID Connect Core 1.0 for details.  | [optional] 
**ClientNotificationToken** | Pointer to **string** | The client notification token included in the backchannel authentication request.  | [optional] 
**Acrs** | Pointer to **[]string** | The list of ACR values requested by the backchannel authentication request.  Basically, this property holds the value of the &#x60;acr_values&#x60; request parameter in the backchannel authentication request. However, because unsupported ACR values are dropped on Authlete side, if the &#x60;acr_values&#x60; request parameter contains unrecognized ACR values, the list returned by this property becomes different from the value of the &#x60;acr_values&#x60; request parameter.  | [optional] 
**HintType** | Pointer to **string** | The type of the hint for end-user identification which was included in the backchannel authentication request.  | [optional] 
**Hint** | Pointer to **string** | The value of the hint for end-user identification.  | [optional] 
**Sub** | Pointer to **string** | The value of the &#x60;sub&#x60; claim contained in the ID token hint included in the backchannel authentication request.  | [optional] 
**BindingMessage** | Pointer to **string** | The binding message included in the backchannel authentication request.  | [optional] 
**UserCode** | Pointer to **string** | The binding message included in the backchannel authentication request.  | [optional] 
**UserCodeRequired** | Pointer to **bool** | The flag which indicates whether a user code is required.  &#x60;true&#x60; when both the &#x60;backchannel_user_code_parameter&#x60; metadata of the client (&#x3D; Client&#39;s &#x60;bcUserCodeRequired&#x60; property) and the &#x60;backchannel_user_code_parameter_supported&#x60; metadata of the service (&#x3D; Service&#39;s &#x60;backchannelUserCodeParameterSupported&#x60; property) are &#x60;true&#x60;.  | [optional] 
**RequestedExpiry** | Pointer to **int32** | The requested expiry for the authentication request ID (&#x60;auth_req_id&#x60;).  | [optional] 
**RequestContext** | Pointer to **string** | The request context of the backchannel authentication request.  It is the value of the request_context claim in the signed authentication request and its format is JSON. request_context is a new claim added by the FAPI-CIBA profile.  | [optional] 
**Warnings** | Pointer to **[]string** | The warnings raised during processing the backchannel authentication request.  | [optional] 
**Ticket** | Pointer to **string** | The ticket which is necessary to call Authlete&#39;s &#x60;/auth/token/fail&#x60; API or &#x60;/auth/token/issue&#x60; API.  This parameter has a value only if the value of &#x60;grant_type&#x60; request parameter is &#x60;password&#x60; and the token request is valid.  | [optional] 
**Resources** | Pointer to **[]string** | The resources specified by the &#x60;resource&#x60; request parameters or by the &#x60;resource&#x60; property in the request object. If both are given, the values in the request object should be set. See \&quot;Resource Indicators for OAuth 2.0\&quot; for details.  | [optional] 
**AuthorizationDetails** | Pointer to [**AuthzDetails**](AuthzDetails.md) |  | [optional] 
**ServiceAttributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of this service that the client application belongs to.  | [optional] 
**ClientAttributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of the client.  | [optional] 
**DynamicScopes** | Pointer to [**[]DynamicScope**](DynamicScope.md) | The dynamic scopes which the client application requested by the scope request parameter.  | [optional] 
**DeliveryMode** | Pointer to [**DeliveryMode**](DeliveryMode.md) |  | [optional] 
**GmAction** | Pointer to [**GrantManagementAction**](GrantManagementAction.md) |  | [optional] 
**GrantId** | Pointer to **string** | the value of the &#x60;grant_id&#x60; request parameter of the device authorization request.  The &#x60;grant_id&#x60; request parameter is defined in [Grant Management for OAuth 2.0](https://openid.net/specs/fapi-grant-management.html) , which is supported by Authlete 2.3 and newer versions.  | [optional] 
**Grant** | Pointer to [**Grant**](Grant.md) |  | [optional] 
**GrantSubject** | Pointer to **string** | The subject identifying the user who has given the grant identified by the &#x60;grant_id&#x60; request parameter of the device authorization request.  Authlete 2.3 and newer versions support &lt;a href&#x3D; \&quot;https://openid.net/specs/fapi-grant-management.html\&quot;&gt;Grant Management for OAuth 2.0&lt;/a&gt;. An authorization request may contain a &#x60;grant_id&#x60; request parameter which is defined in the specification. If the value of the request parameter is valid, {@link #getGrantSubject()} will return the subject of the user who has given the grant to the client application. Authorization server implementations may use the value returned from {@link #getGrantSubject()} in order to determine the user to authenticate.  The user your system will authenticate during the authorization process (or has already authenticated) may be different from the user of the grant. The first implementer&#39;s draft of \&quot;Grant Management for OAuth 2.0\&quot; does not mention anything about the case, so the behavior in the case is left to implementations. Authlete will not perform the grant management action when the &#x60;subject&#x60; passed to Authlete does not match the user of the grant.  | [optional] 
**ClientEntityId** | Pointer to **string** | The entity ID of the client.  | [optional] 
**ClientEntityIdUsed** | Pointer to **bool** | Flag which indicates whether the entity ID of the client was used when the request for the access token was made. | [optional] 

## Methods

### NewBackchannelAuthenticationResponse

`func NewBackchannelAuthenticationResponse() *BackchannelAuthenticationResponse`

NewBackchannelAuthenticationResponse instantiates a new BackchannelAuthenticationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackchannelAuthenticationResponseWithDefaults

`func NewBackchannelAuthenticationResponseWithDefaults() *BackchannelAuthenticationResponse`

NewBackchannelAuthenticationResponseWithDefaults instantiates a new BackchannelAuthenticationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *BackchannelAuthenticationResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *BackchannelAuthenticationResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *BackchannelAuthenticationResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *BackchannelAuthenticationResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *BackchannelAuthenticationResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *BackchannelAuthenticationResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *BackchannelAuthenticationResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *BackchannelAuthenticationResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *BackchannelAuthenticationResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BackchannelAuthenticationResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BackchannelAuthenticationResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BackchannelAuthenticationResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResponseContent

`func (o *BackchannelAuthenticationResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *BackchannelAuthenticationResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *BackchannelAuthenticationResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *BackchannelAuthenticationResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetClientId

`func (o *BackchannelAuthenticationResponse) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BackchannelAuthenticationResponse) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BackchannelAuthenticationResponse) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BackchannelAuthenticationResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIdAlias

`func (o *BackchannelAuthenticationResponse) GetClientIdAlias() string`

GetClientIdAlias returns the ClientIdAlias field if non-nil, zero value otherwise.

### GetClientIdAliasOk

`func (o *BackchannelAuthenticationResponse) GetClientIdAliasOk() (*string, bool)`

GetClientIdAliasOk returns a tuple with the ClientIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAlias

`func (o *BackchannelAuthenticationResponse) SetClientIdAlias(v string)`

SetClientIdAlias sets ClientIdAlias field to given value.

### HasClientIdAlias

`func (o *BackchannelAuthenticationResponse) HasClientIdAlias() bool`

HasClientIdAlias returns a boolean if a field has been set.

### GetClientIdAliasUsed

`func (o *BackchannelAuthenticationResponse) GetClientIdAliasUsed() bool`

GetClientIdAliasUsed returns the ClientIdAliasUsed field if non-nil, zero value otherwise.

### GetClientIdAliasUsedOk

`func (o *BackchannelAuthenticationResponse) GetClientIdAliasUsedOk() (*bool, bool)`

GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasUsed

`func (o *BackchannelAuthenticationResponse) SetClientIdAliasUsed(v bool)`

SetClientIdAliasUsed sets ClientIdAliasUsed field to given value.

### HasClientIdAliasUsed

`func (o *BackchannelAuthenticationResponse) HasClientIdAliasUsed() bool`

HasClientIdAliasUsed returns a boolean if a field has been set.

### GetClientName

`func (o *BackchannelAuthenticationResponse) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *BackchannelAuthenticationResponse) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *BackchannelAuthenticationResponse) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *BackchannelAuthenticationResponse) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetScopes

`func (o *BackchannelAuthenticationResponse) GetScopes() []Scope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *BackchannelAuthenticationResponse) GetScopesOk() (*[]Scope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *BackchannelAuthenticationResponse) SetScopes(v []Scope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *BackchannelAuthenticationResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetClaimNames

`func (o *BackchannelAuthenticationResponse) GetClaimNames() []string`

GetClaimNames returns the ClaimNames field if non-nil, zero value otherwise.

### GetClaimNamesOk

`func (o *BackchannelAuthenticationResponse) GetClaimNamesOk() (*[]string, bool)`

GetClaimNamesOk returns a tuple with the ClaimNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimNames

`func (o *BackchannelAuthenticationResponse) SetClaimNames(v []string)`

SetClaimNames sets ClaimNames field to given value.

### HasClaimNames

`func (o *BackchannelAuthenticationResponse) HasClaimNames() bool`

HasClaimNames returns a boolean if a field has been set.

### GetClientNotificationToken

`func (o *BackchannelAuthenticationResponse) GetClientNotificationToken() string`

GetClientNotificationToken returns the ClientNotificationToken field if non-nil, zero value otherwise.

### GetClientNotificationTokenOk

`func (o *BackchannelAuthenticationResponse) GetClientNotificationTokenOk() (*string, bool)`

GetClientNotificationTokenOk returns a tuple with the ClientNotificationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNotificationToken

`func (o *BackchannelAuthenticationResponse) SetClientNotificationToken(v string)`

SetClientNotificationToken sets ClientNotificationToken field to given value.

### HasClientNotificationToken

`func (o *BackchannelAuthenticationResponse) HasClientNotificationToken() bool`

HasClientNotificationToken returns a boolean if a field has been set.

### GetAcrs

`func (o *BackchannelAuthenticationResponse) GetAcrs() []string`

GetAcrs returns the Acrs field if non-nil, zero value otherwise.

### GetAcrsOk

`func (o *BackchannelAuthenticationResponse) GetAcrsOk() (*[]string, bool)`

GetAcrsOk returns a tuple with the Acrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrs

`func (o *BackchannelAuthenticationResponse) SetAcrs(v []string)`

SetAcrs sets Acrs field to given value.

### HasAcrs

`func (o *BackchannelAuthenticationResponse) HasAcrs() bool`

HasAcrs returns a boolean if a field has been set.

### GetHintType

`func (o *BackchannelAuthenticationResponse) GetHintType() string`

GetHintType returns the HintType field if non-nil, zero value otherwise.

### GetHintTypeOk

`func (o *BackchannelAuthenticationResponse) GetHintTypeOk() (*string, bool)`

GetHintTypeOk returns a tuple with the HintType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHintType

`func (o *BackchannelAuthenticationResponse) SetHintType(v string)`

SetHintType sets HintType field to given value.

### HasHintType

`func (o *BackchannelAuthenticationResponse) HasHintType() bool`

HasHintType returns a boolean if a field has been set.

### GetHint

`func (o *BackchannelAuthenticationResponse) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *BackchannelAuthenticationResponse) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *BackchannelAuthenticationResponse) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *BackchannelAuthenticationResponse) HasHint() bool`

HasHint returns a boolean if a field has been set.

### GetSub

`func (o *BackchannelAuthenticationResponse) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *BackchannelAuthenticationResponse) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *BackchannelAuthenticationResponse) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *BackchannelAuthenticationResponse) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetBindingMessage

`func (o *BackchannelAuthenticationResponse) GetBindingMessage() string`

GetBindingMessage returns the BindingMessage field if non-nil, zero value otherwise.

### GetBindingMessageOk

`func (o *BackchannelAuthenticationResponse) GetBindingMessageOk() (*string, bool)`

GetBindingMessageOk returns a tuple with the BindingMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingMessage

`func (o *BackchannelAuthenticationResponse) SetBindingMessage(v string)`

SetBindingMessage sets BindingMessage field to given value.

### HasBindingMessage

`func (o *BackchannelAuthenticationResponse) HasBindingMessage() bool`

HasBindingMessage returns a boolean if a field has been set.

### GetUserCode

`func (o *BackchannelAuthenticationResponse) GetUserCode() string`

GetUserCode returns the UserCode field if non-nil, zero value otherwise.

### GetUserCodeOk

`func (o *BackchannelAuthenticationResponse) GetUserCodeOk() (*string, bool)`

GetUserCodeOk returns a tuple with the UserCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCode

`func (o *BackchannelAuthenticationResponse) SetUserCode(v string)`

SetUserCode sets UserCode field to given value.

### HasUserCode

`func (o *BackchannelAuthenticationResponse) HasUserCode() bool`

HasUserCode returns a boolean if a field has been set.

### GetUserCodeRequired

`func (o *BackchannelAuthenticationResponse) GetUserCodeRequired() bool`

GetUserCodeRequired returns the UserCodeRequired field if non-nil, zero value otherwise.

### GetUserCodeRequiredOk

`func (o *BackchannelAuthenticationResponse) GetUserCodeRequiredOk() (*bool, bool)`

GetUserCodeRequiredOk returns a tuple with the UserCodeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCodeRequired

`func (o *BackchannelAuthenticationResponse) SetUserCodeRequired(v bool)`

SetUserCodeRequired sets UserCodeRequired field to given value.

### HasUserCodeRequired

`func (o *BackchannelAuthenticationResponse) HasUserCodeRequired() bool`

HasUserCodeRequired returns a boolean if a field has been set.

### GetRequestedExpiry

`func (o *BackchannelAuthenticationResponse) GetRequestedExpiry() int32`

GetRequestedExpiry returns the RequestedExpiry field if non-nil, zero value otherwise.

### GetRequestedExpiryOk

`func (o *BackchannelAuthenticationResponse) GetRequestedExpiryOk() (*int32, bool)`

GetRequestedExpiryOk returns a tuple with the RequestedExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedExpiry

`func (o *BackchannelAuthenticationResponse) SetRequestedExpiry(v int32)`

SetRequestedExpiry sets RequestedExpiry field to given value.

### HasRequestedExpiry

`func (o *BackchannelAuthenticationResponse) HasRequestedExpiry() bool`

HasRequestedExpiry returns a boolean if a field has been set.

### GetRequestContext

`func (o *BackchannelAuthenticationResponse) GetRequestContext() string`

GetRequestContext returns the RequestContext field if non-nil, zero value otherwise.

### GetRequestContextOk

`func (o *BackchannelAuthenticationResponse) GetRequestContextOk() (*string, bool)`

GetRequestContextOk returns a tuple with the RequestContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContext

`func (o *BackchannelAuthenticationResponse) SetRequestContext(v string)`

SetRequestContext sets RequestContext field to given value.

### HasRequestContext

`func (o *BackchannelAuthenticationResponse) HasRequestContext() bool`

HasRequestContext returns a boolean if a field has been set.

### GetWarnings

`func (o *BackchannelAuthenticationResponse) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *BackchannelAuthenticationResponse) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *BackchannelAuthenticationResponse) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *BackchannelAuthenticationResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetTicket

`func (o *BackchannelAuthenticationResponse) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *BackchannelAuthenticationResponse) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *BackchannelAuthenticationResponse) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *BackchannelAuthenticationResponse) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetResources

`func (o *BackchannelAuthenticationResponse) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *BackchannelAuthenticationResponse) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *BackchannelAuthenticationResponse) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *BackchannelAuthenticationResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetAuthorizationDetails

`func (o *BackchannelAuthenticationResponse) GetAuthorizationDetails() AuthzDetails`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *BackchannelAuthenticationResponse) GetAuthorizationDetailsOk() (*AuthzDetails, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *BackchannelAuthenticationResponse) SetAuthorizationDetails(v AuthzDetails)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *BackchannelAuthenticationResponse) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.

### GetServiceAttributes

`func (o *BackchannelAuthenticationResponse) GetServiceAttributes() []Pair`

GetServiceAttributes returns the ServiceAttributes field if non-nil, zero value otherwise.

### GetServiceAttributesOk

`func (o *BackchannelAuthenticationResponse) GetServiceAttributesOk() (*[]Pair, bool)`

GetServiceAttributesOk returns a tuple with the ServiceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAttributes

`func (o *BackchannelAuthenticationResponse) SetServiceAttributes(v []Pair)`

SetServiceAttributes sets ServiceAttributes field to given value.

### HasServiceAttributes

`func (o *BackchannelAuthenticationResponse) HasServiceAttributes() bool`

HasServiceAttributes returns a boolean if a field has been set.

### GetClientAttributes

`func (o *BackchannelAuthenticationResponse) GetClientAttributes() []Pair`

GetClientAttributes returns the ClientAttributes field if non-nil, zero value otherwise.

### GetClientAttributesOk

`func (o *BackchannelAuthenticationResponse) GetClientAttributesOk() (*[]Pair, bool)`

GetClientAttributesOk returns a tuple with the ClientAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAttributes

`func (o *BackchannelAuthenticationResponse) SetClientAttributes(v []Pair)`

SetClientAttributes sets ClientAttributes field to given value.

### HasClientAttributes

`func (o *BackchannelAuthenticationResponse) HasClientAttributes() bool`

HasClientAttributes returns a boolean if a field has been set.

### GetDynamicScopes

`func (o *BackchannelAuthenticationResponse) GetDynamicScopes() []DynamicScope`

GetDynamicScopes returns the DynamicScopes field if non-nil, zero value otherwise.

### GetDynamicScopesOk

`func (o *BackchannelAuthenticationResponse) GetDynamicScopesOk() (*[]DynamicScope, bool)`

GetDynamicScopesOk returns a tuple with the DynamicScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicScopes

`func (o *BackchannelAuthenticationResponse) SetDynamicScopes(v []DynamicScope)`

SetDynamicScopes sets DynamicScopes field to given value.

### HasDynamicScopes

`func (o *BackchannelAuthenticationResponse) HasDynamicScopes() bool`

HasDynamicScopes returns a boolean if a field has been set.

### GetDeliveryMode

`func (o *BackchannelAuthenticationResponse) GetDeliveryMode() DeliveryMode`

GetDeliveryMode returns the DeliveryMode field if non-nil, zero value otherwise.

### GetDeliveryModeOk

`func (o *BackchannelAuthenticationResponse) GetDeliveryModeOk() (*DeliveryMode, bool)`

GetDeliveryModeOk returns a tuple with the DeliveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMode

`func (o *BackchannelAuthenticationResponse) SetDeliveryMode(v DeliveryMode)`

SetDeliveryMode sets DeliveryMode field to given value.

### HasDeliveryMode

`func (o *BackchannelAuthenticationResponse) HasDeliveryMode() bool`

HasDeliveryMode returns a boolean if a field has been set.

### GetGmAction

`func (o *BackchannelAuthenticationResponse) GetGmAction() GrantManagementAction`

GetGmAction returns the GmAction field if non-nil, zero value otherwise.

### GetGmActionOk

`func (o *BackchannelAuthenticationResponse) GetGmActionOk() (*GrantManagementAction, bool)`

GetGmActionOk returns a tuple with the GmAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmAction

`func (o *BackchannelAuthenticationResponse) SetGmAction(v GrantManagementAction)`

SetGmAction sets GmAction field to given value.

### HasGmAction

`func (o *BackchannelAuthenticationResponse) HasGmAction() bool`

HasGmAction returns a boolean if a field has been set.

### GetGrantId

`func (o *BackchannelAuthenticationResponse) GetGrantId() string`

GetGrantId returns the GrantId field if non-nil, zero value otherwise.

### GetGrantIdOk

`func (o *BackchannelAuthenticationResponse) GetGrantIdOk() (*string, bool)`

GetGrantIdOk returns a tuple with the GrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantId

`func (o *BackchannelAuthenticationResponse) SetGrantId(v string)`

SetGrantId sets GrantId field to given value.

### HasGrantId

`func (o *BackchannelAuthenticationResponse) HasGrantId() bool`

HasGrantId returns a boolean if a field has been set.

### GetGrant

`func (o *BackchannelAuthenticationResponse) GetGrant() Grant`

GetGrant returns the Grant field if non-nil, zero value otherwise.

### GetGrantOk

`func (o *BackchannelAuthenticationResponse) GetGrantOk() (*Grant, bool)`

GetGrantOk returns a tuple with the Grant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrant

`func (o *BackchannelAuthenticationResponse) SetGrant(v Grant)`

SetGrant sets Grant field to given value.

### HasGrant

`func (o *BackchannelAuthenticationResponse) HasGrant() bool`

HasGrant returns a boolean if a field has been set.

### GetGrantSubject

`func (o *BackchannelAuthenticationResponse) GetGrantSubject() string`

GetGrantSubject returns the GrantSubject field if non-nil, zero value otherwise.

### GetGrantSubjectOk

`func (o *BackchannelAuthenticationResponse) GetGrantSubjectOk() (*string, bool)`

GetGrantSubjectOk returns a tuple with the GrantSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantSubject

`func (o *BackchannelAuthenticationResponse) SetGrantSubject(v string)`

SetGrantSubject sets GrantSubject field to given value.

### HasGrantSubject

`func (o *BackchannelAuthenticationResponse) HasGrantSubject() bool`

HasGrantSubject returns a boolean if a field has been set.

### GetClientEntityId

`func (o *BackchannelAuthenticationResponse) GetClientEntityId() string`

GetClientEntityId returns the ClientEntityId field if non-nil, zero value otherwise.

### GetClientEntityIdOk

`func (o *BackchannelAuthenticationResponse) GetClientEntityIdOk() (*string, bool)`

GetClientEntityIdOk returns a tuple with the ClientEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEntityId

`func (o *BackchannelAuthenticationResponse) SetClientEntityId(v string)`

SetClientEntityId sets ClientEntityId field to given value.

### HasClientEntityId

`func (o *BackchannelAuthenticationResponse) HasClientEntityId() bool`

HasClientEntityId returns a boolean if a field has been set.

### GetClientEntityIdUsed

`func (o *BackchannelAuthenticationResponse) GetClientEntityIdUsed() bool`

GetClientEntityIdUsed returns the ClientEntityIdUsed field if non-nil, zero value otherwise.

### GetClientEntityIdUsedOk

`func (o *BackchannelAuthenticationResponse) GetClientEntityIdUsedOk() (*bool, bool)`

GetClientEntityIdUsedOk returns a tuple with the ClientEntityIdUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEntityIdUsed

`func (o *BackchannelAuthenticationResponse) SetClientEntityIdUsed(v bool)`

SetClientEntityIdUsed sets ClientEntityIdUsed field to given value.

### HasClientEntityIdUsed

`func (o *BackchannelAuthenticationResponse) HasClientEntityIdUsed() bool`

HasClientEntityIdUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


