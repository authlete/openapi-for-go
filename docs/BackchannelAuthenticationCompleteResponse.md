# BackchannelAuthenticationCompleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take.  | [optional] 
**ResponseContent** | Pointer to **string** | The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of &#x60;action&#x60; parameter.  | [optional] 
**ClientId** | Pointer to **int64** | The client ID of the client application that has made the backchannel authentication request.  | [optional] 
**ClientIdAlias** | Pointer to **string** | The client ID alias of the client application that has made the backchannel authentication request.  | [optional] 
**ClientIdAliasUsed** | Pointer to **bool** | &#x60;true&#x60; if the value of the client_id request parameter included in the backchannel authentication request is the client ID alias. &#x60;false&#x60; if the value is the original numeric client ID.  | [optional] 
**ClientName** | Pointer to **string** | The name of the client application which has made the backchannel authentication request.  | [optional] 
**DeliveryMode** | Pointer to [**DeliveryMode**](DeliveryMode.md) |  | [optional] 
**ClientNotificationEndpoint** | Pointer to **string** | The client notification endpoint to which a notification needs to be sent. This corresponds to the &#x60;client_notification_endpoint&#x60; metadata of the client application.  | [optional] 
**ClientNotificationToken** | Pointer to **string** | The client notification token which needs to be embedded as a Bearer token in the Authorization header in the notification. This is the value of the &#x60;client_notification_token&#x60; request parameter included in the backchannel authentication request.  | [optional] 
**AuthReqId** | Pointer to **string** | The newly issued authentication request ID.  | [optional] 
**AccessToken** | Pointer to **string** | The issued access token.  | [optional] 
**RefreshToken** | Pointer to **string** | The issued refresh token.  | [optional] 
**IdToken** | Pointer to **string** | The issued ID token.  | [optional] 
**AccessTokenDuration** | Pointer to **int64** | The duration of the access token in seconds.  | [optional] 
**RefreshTokenDuration** | Pointer to **int64** | The duration of the refresh token in seconds.  | [optional] 
**IdTokenDuration** | Pointer to **int64** | The duration of the ID token in seconds.  | [optional] 
**JwtAccessToken** | Pointer to **string** | The issued access token in JWT format.  | [optional] 
**Resources** | Pointer to **[]string** | The resources specified by the &#x60;resource&#x60; request parameters or by the &#x60;resource&#x60; property in the request object. If both are given, the values in the request object should be set. See \&quot;Resource Indicators for OAuth 2.0\&quot; for details.  | [optional] 
**AuthorizationDetails** | Pointer to [**AuthorizationDetails**](AuthorizationDetails.md) |  | [optional] 
**ServiceAttributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of this service that the client application belongs to.  | [optional] 
**ClientAttributes** | Pointer to [**[]Pair**](Pair.md) | The attributes of the client.  | [optional] 

## Methods

### NewBackchannelAuthenticationCompleteResponse

`func NewBackchannelAuthenticationCompleteResponse() *BackchannelAuthenticationCompleteResponse`

NewBackchannelAuthenticationCompleteResponse instantiates a new BackchannelAuthenticationCompleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackchannelAuthenticationCompleteResponseWithDefaults

`func NewBackchannelAuthenticationCompleteResponseWithDefaults() *BackchannelAuthenticationCompleteResponse`

NewBackchannelAuthenticationCompleteResponseWithDefaults instantiates a new BackchannelAuthenticationCompleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *BackchannelAuthenticationCompleteResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *BackchannelAuthenticationCompleteResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *BackchannelAuthenticationCompleteResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *BackchannelAuthenticationCompleteResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *BackchannelAuthenticationCompleteResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *BackchannelAuthenticationCompleteResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *BackchannelAuthenticationCompleteResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *BackchannelAuthenticationCompleteResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *BackchannelAuthenticationCompleteResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BackchannelAuthenticationCompleteResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BackchannelAuthenticationCompleteResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BackchannelAuthenticationCompleteResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResponseContent

`func (o *BackchannelAuthenticationCompleteResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *BackchannelAuthenticationCompleteResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *BackchannelAuthenticationCompleteResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *BackchannelAuthenticationCompleteResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetClientId

`func (o *BackchannelAuthenticationCompleteResponse) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BackchannelAuthenticationCompleteResponse) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BackchannelAuthenticationCompleteResponse) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BackchannelAuthenticationCompleteResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIdAlias

`func (o *BackchannelAuthenticationCompleteResponse) GetClientIdAlias() string`

GetClientIdAlias returns the ClientIdAlias field if non-nil, zero value otherwise.

### GetClientIdAliasOk

`func (o *BackchannelAuthenticationCompleteResponse) GetClientIdAliasOk() (*string, bool)`

GetClientIdAliasOk returns a tuple with the ClientIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAlias

`func (o *BackchannelAuthenticationCompleteResponse) SetClientIdAlias(v string)`

SetClientIdAlias sets ClientIdAlias field to given value.

### HasClientIdAlias

`func (o *BackchannelAuthenticationCompleteResponse) HasClientIdAlias() bool`

HasClientIdAlias returns a boolean if a field has been set.

### GetClientIdAliasUsed

`func (o *BackchannelAuthenticationCompleteResponse) GetClientIdAliasUsed() bool`

GetClientIdAliasUsed returns the ClientIdAliasUsed field if non-nil, zero value otherwise.

### GetClientIdAliasUsedOk

`func (o *BackchannelAuthenticationCompleteResponse) GetClientIdAliasUsedOk() (*bool, bool)`

GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasUsed

`func (o *BackchannelAuthenticationCompleteResponse) SetClientIdAliasUsed(v bool)`

SetClientIdAliasUsed sets ClientIdAliasUsed field to given value.

### HasClientIdAliasUsed

`func (o *BackchannelAuthenticationCompleteResponse) HasClientIdAliasUsed() bool`

HasClientIdAliasUsed returns a boolean if a field has been set.

### GetClientName

`func (o *BackchannelAuthenticationCompleteResponse) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *BackchannelAuthenticationCompleteResponse) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *BackchannelAuthenticationCompleteResponse) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *BackchannelAuthenticationCompleteResponse) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetDeliveryMode

`func (o *BackchannelAuthenticationCompleteResponse) GetDeliveryMode() DeliveryMode`

GetDeliveryMode returns the DeliveryMode field if non-nil, zero value otherwise.

### GetDeliveryModeOk

`func (o *BackchannelAuthenticationCompleteResponse) GetDeliveryModeOk() (*DeliveryMode, bool)`

GetDeliveryModeOk returns a tuple with the DeliveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMode

`func (o *BackchannelAuthenticationCompleteResponse) SetDeliveryMode(v DeliveryMode)`

SetDeliveryMode sets DeliveryMode field to given value.

### HasDeliveryMode

`func (o *BackchannelAuthenticationCompleteResponse) HasDeliveryMode() bool`

HasDeliveryMode returns a boolean if a field has been set.

### GetClientNotificationEndpoint

`func (o *BackchannelAuthenticationCompleteResponse) GetClientNotificationEndpoint() string`

GetClientNotificationEndpoint returns the ClientNotificationEndpoint field if non-nil, zero value otherwise.

### GetClientNotificationEndpointOk

`func (o *BackchannelAuthenticationCompleteResponse) GetClientNotificationEndpointOk() (*string, bool)`

GetClientNotificationEndpointOk returns a tuple with the ClientNotificationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNotificationEndpoint

`func (o *BackchannelAuthenticationCompleteResponse) SetClientNotificationEndpoint(v string)`

SetClientNotificationEndpoint sets ClientNotificationEndpoint field to given value.

### HasClientNotificationEndpoint

`func (o *BackchannelAuthenticationCompleteResponse) HasClientNotificationEndpoint() bool`

HasClientNotificationEndpoint returns a boolean if a field has been set.

### GetClientNotificationToken

`func (o *BackchannelAuthenticationCompleteResponse) GetClientNotificationToken() string`

GetClientNotificationToken returns the ClientNotificationToken field if non-nil, zero value otherwise.

### GetClientNotificationTokenOk

`func (o *BackchannelAuthenticationCompleteResponse) GetClientNotificationTokenOk() (*string, bool)`

GetClientNotificationTokenOk returns a tuple with the ClientNotificationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNotificationToken

`func (o *BackchannelAuthenticationCompleteResponse) SetClientNotificationToken(v string)`

SetClientNotificationToken sets ClientNotificationToken field to given value.

### HasClientNotificationToken

`func (o *BackchannelAuthenticationCompleteResponse) HasClientNotificationToken() bool`

HasClientNotificationToken returns a boolean if a field has been set.

### GetAuthReqId

`func (o *BackchannelAuthenticationCompleteResponse) GetAuthReqId() string`

GetAuthReqId returns the AuthReqId field if non-nil, zero value otherwise.

### GetAuthReqIdOk

`func (o *BackchannelAuthenticationCompleteResponse) GetAuthReqIdOk() (*string, bool)`

GetAuthReqIdOk returns a tuple with the AuthReqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthReqId

`func (o *BackchannelAuthenticationCompleteResponse) SetAuthReqId(v string)`

SetAuthReqId sets AuthReqId field to given value.

### HasAuthReqId

`func (o *BackchannelAuthenticationCompleteResponse) HasAuthReqId() bool`

HasAuthReqId returns a boolean if a field has been set.

### GetAccessToken

`func (o *BackchannelAuthenticationCompleteResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *BackchannelAuthenticationCompleteResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *BackchannelAuthenticationCompleteResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *BackchannelAuthenticationCompleteResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *BackchannelAuthenticationCompleteResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *BackchannelAuthenticationCompleteResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *BackchannelAuthenticationCompleteResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *BackchannelAuthenticationCompleteResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetIdToken

`func (o *BackchannelAuthenticationCompleteResponse) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *BackchannelAuthenticationCompleteResponse) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *BackchannelAuthenticationCompleteResponse) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *BackchannelAuthenticationCompleteResponse) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.

### GetAccessTokenDuration

`func (o *BackchannelAuthenticationCompleteResponse) GetAccessTokenDuration() int64`

GetAccessTokenDuration returns the AccessTokenDuration field if non-nil, zero value otherwise.

### GetAccessTokenDurationOk

`func (o *BackchannelAuthenticationCompleteResponse) GetAccessTokenDurationOk() (*int64, bool)`

GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenDuration

`func (o *BackchannelAuthenticationCompleteResponse) SetAccessTokenDuration(v int64)`

SetAccessTokenDuration sets AccessTokenDuration field to given value.

### HasAccessTokenDuration

`func (o *BackchannelAuthenticationCompleteResponse) HasAccessTokenDuration() bool`

HasAccessTokenDuration returns a boolean if a field has been set.

### GetRefreshTokenDuration

`func (o *BackchannelAuthenticationCompleteResponse) GetRefreshTokenDuration() int64`

GetRefreshTokenDuration returns the RefreshTokenDuration field if non-nil, zero value otherwise.

### GetRefreshTokenDurationOk

`func (o *BackchannelAuthenticationCompleteResponse) GetRefreshTokenDurationOk() (*int64, bool)`

GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDuration

`func (o *BackchannelAuthenticationCompleteResponse) SetRefreshTokenDuration(v int64)`

SetRefreshTokenDuration sets RefreshTokenDuration field to given value.

### HasRefreshTokenDuration

`func (o *BackchannelAuthenticationCompleteResponse) HasRefreshTokenDuration() bool`

HasRefreshTokenDuration returns a boolean if a field has been set.

### GetIdTokenDuration

`func (o *BackchannelAuthenticationCompleteResponse) GetIdTokenDuration() int64`

GetIdTokenDuration returns the IdTokenDuration field if non-nil, zero value otherwise.

### GetIdTokenDurationOk

`func (o *BackchannelAuthenticationCompleteResponse) GetIdTokenDurationOk() (*int64, bool)`

GetIdTokenDurationOk returns a tuple with the IdTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenDuration

`func (o *BackchannelAuthenticationCompleteResponse) SetIdTokenDuration(v int64)`

SetIdTokenDuration sets IdTokenDuration field to given value.

### HasIdTokenDuration

`func (o *BackchannelAuthenticationCompleteResponse) HasIdTokenDuration() bool`

HasIdTokenDuration returns a boolean if a field has been set.

### GetJwtAccessToken

`func (o *BackchannelAuthenticationCompleteResponse) GetJwtAccessToken() string`

GetJwtAccessToken returns the JwtAccessToken field if non-nil, zero value otherwise.

### GetJwtAccessTokenOk

`func (o *BackchannelAuthenticationCompleteResponse) GetJwtAccessTokenOk() (*string, bool)`

GetJwtAccessTokenOk returns a tuple with the JwtAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtAccessToken

`func (o *BackchannelAuthenticationCompleteResponse) SetJwtAccessToken(v string)`

SetJwtAccessToken sets JwtAccessToken field to given value.

### HasJwtAccessToken

`func (o *BackchannelAuthenticationCompleteResponse) HasJwtAccessToken() bool`

HasJwtAccessToken returns a boolean if a field has been set.

### GetResources

`func (o *BackchannelAuthenticationCompleteResponse) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *BackchannelAuthenticationCompleteResponse) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *BackchannelAuthenticationCompleteResponse) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *BackchannelAuthenticationCompleteResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetAuthorizationDetails

`func (o *BackchannelAuthenticationCompleteResponse) GetAuthorizationDetails() AuthorizationDetails`

GetAuthorizationDetails returns the AuthorizationDetails field if non-nil, zero value otherwise.

### GetAuthorizationDetailsOk

`func (o *BackchannelAuthenticationCompleteResponse) GetAuthorizationDetailsOk() (*AuthorizationDetails, bool)`

GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetails

`func (o *BackchannelAuthenticationCompleteResponse) SetAuthorizationDetails(v AuthorizationDetails)`

SetAuthorizationDetails sets AuthorizationDetails field to given value.

### HasAuthorizationDetails

`func (o *BackchannelAuthenticationCompleteResponse) HasAuthorizationDetails() bool`

HasAuthorizationDetails returns a boolean if a field has been set.

### GetServiceAttributes

`func (o *BackchannelAuthenticationCompleteResponse) GetServiceAttributes() []Pair`

GetServiceAttributes returns the ServiceAttributes field if non-nil, zero value otherwise.

### GetServiceAttributesOk

`func (o *BackchannelAuthenticationCompleteResponse) GetServiceAttributesOk() (*[]Pair, bool)`

GetServiceAttributesOk returns a tuple with the ServiceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAttributes

`func (o *BackchannelAuthenticationCompleteResponse) SetServiceAttributes(v []Pair)`

SetServiceAttributes sets ServiceAttributes field to given value.

### HasServiceAttributes

`func (o *BackchannelAuthenticationCompleteResponse) HasServiceAttributes() bool`

HasServiceAttributes returns a boolean if a field has been set.

### GetClientAttributes

`func (o *BackchannelAuthenticationCompleteResponse) GetClientAttributes() []Pair`

GetClientAttributes returns the ClientAttributes field if non-nil, zero value otherwise.

### GetClientAttributesOk

`func (o *BackchannelAuthenticationCompleteResponse) GetClientAttributesOk() (*[]Pair, bool)`

GetClientAttributesOk returns a tuple with the ClientAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAttributes

`func (o *BackchannelAuthenticationCompleteResponse) SetClientAttributes(v []Pair)`

SetClientAttributes sets ClientAttributes field to given value.

### HasClientAttributes

`func (o *BackchannelAuthenticationCompleteResponse) HasClientAttributes() bool`

HasClientAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


