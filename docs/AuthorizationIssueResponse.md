# AuthorizationIssueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultCode** | Pointer to **string** | The code which represents the result of the API call. | [optional] 
**ResultMessage** | Pointer to **string** | A short message which explains the result of the API call. | [optional] 
**Action** | Pointer to **string** | The next action that the authorization server implementation should take. | [optional] 
**ResponseContent** | Pointer to **string** | The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of &#x60;action&#x60; parameter.  | [optional] 
**AccessToken** | Pointer to **string** | The newly issued access token. Note that an access token is issued from an authorization endpoint only when &#x60;response_type&#x60; contains token.  | [optional] 
**AccessTokenExpiresAt** | Pointer to **int64** | The datetime at which the newly issued access token will expire. The value is represented in milliseconds since the Unix epoch (1970-01-01).  | [optional] 
**AccessTokenDuration** | Pointer to **int64** | The duration of the newly issued access token in seconds.  | [optional] 
**IdToken** | Pointer to **string** | The newly issued ID token. Note that an ID token is issued from an authorization endpoint only when &#x60;response_type&#x60; contains &#x60;id_token&#x60;.  | [optional] 
**AuthorizationCode** | Pointer to **string** | The newly issued authorization code. Note that an authorization code is issued only when &#x60;response_type&#x60; contains code.  | [optional] 
**JwtAccessToken** | Pointer to **string** | The newly issued access token in JWT format. If the service is not configured to issue JWT-based access tokens, this property is always set to &#x60;null&#x60;.  | [optional] 
**TicketInfo** | Pointer to **string** | The information about the ticket.  | [optional] 

## Methods

### NewAuthorizationIssueResponse

`func NewAuthorizationIssueResponse() *AuthorizationIssueResponse`

NewAuthorizationIssueResponse instantiates a new AuthorizationIssueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationIssueResponseWithDefaults

`func NewAuthorizationIssueResponseWithDefaults() *AuthorizationIssueResponse`

NewAuthorizationIssueResponseWithDefaults instantiates a new AuthorizationIssueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultCode

`func (o *AuthorizationIssueResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *AuthorizationIssueResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *AuthorizationIssueResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *AuthorizationIssueResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetResultMessage

`func (o *AuthorizationIssueResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *AuthorizationIssueResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *AuthorizationIssueResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *AuthorizationIssueResponse) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### GetAction

`func (o *AuthorizationIssueResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuthorizationIssueResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuthorizationIssueResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AuthorizationIssueResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResponseContent

`func (o *AuthorizationIssueResponse) GetResponseContent() string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *AuthorizationIssueResponse) GetResponseContentOk() (*string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *AuthorizationIssueResponse) SetResponseContent(v string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *AuthorizationIssueResponse) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetAccessToken

`func (o *AuthorizationIssueResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AuthorizationIssueResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AuthorizationIssueResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AuthorizationIssueResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAccessTokenExpiresAt

`func (o *AuthorizationIssueResponse) GetAccessTokenExpiresAt() int64`

GetAccessTokenExpiresAt returns the AccessTokenExpiresAt field if non-nil, zero value otherwise.

### GetAccessTokenExpiresAtOk

`func (o *AuthorizationIssueResponse) GetAccessTokenExpiresAtOk() (*int64, bool)`

GetAccessTokenExpiresAtOk returns a tuple with the AccessTokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpiresAt

`func (o *AuthorizationIssueResponse) SetAccessTokenExpiresAt(v int64)`

SetAccessTokenExpiresAt sets AccessTokenExpiresAt field to given value.

### HasAccessTokenExpiresAt

`func (o *AuthorizationIssueResponse) HasAccessTokenExpiresAt() bool`

HasAccessTokenExpiresAt returns a boolean if a field has been set.

### GetAccessTokenDuration

`func (o *AuthorizationIssueResponse) GetAccessTokenDuration() int64`

GetAccessTokenDuration returns the AccessTokenDuration field if non-nil, zero value otherwise.

### GetAccessTokenDurationOk

`func (o *AuthorizationIssueResponse) GetAccessTokenDurationOk() (*int64, bool)`

GetAccessTokenDurationOk returns a tuple with the AccessTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenDuration

`func (o *AuthorizationIssueResponse) SetAccessTokenDuration(v int64)`

SetAccessTokenDuration sets AccessTokenDuration field to given value.

### HasAccessTokenDuration

`func (o *AuthorizationIssueResponse) HasAccessTokenDuration() bool`

HasAccessTokenDuration returns a boolean if a field has been set.

### GetIdToken

`func (o *AuthorizationIssueResponse) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *AuthorizationIssueResponse) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *AuthorizationIssueResponse) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *AuthorizationIssueResponse) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.

### GetAuthorizationCode

`func (o *AuthorizationIssueResponse) GetAuthorizationCode() string`

GetAuthorizationCode returns the AuthorizationCode field if non-nil, zero value otherwise.

### GetAuthorizationCodeOk

`func (o *AuthorizationIssueResponse) GetAuthorizationCodeOk() (*string, bool)`

GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCode

`func (o *AuthorizationIssueResponse) SetAuthorizationCode(v string)`

SetAuthorizationCode sets AuthorizationCode field to given value.

### HasAuthorizationCode

`func (o *AuthorizationIssueResponse) HasAuthorizationCode() bool`

HasAuthorizationCode returns a boolean if a field has been set.

### GetJwtAccessToken

`func (o *AuthorizationIssueResponse) GetJwtAccessToken() string`

GetJwtAccessToken returns the JwtAccessToken field if non-nil, zero value otherwise.

### GetJwtAccessTokenOk

`func (o *AuthorizationIssueResponse) GetJwtAccessTokenOk() (*string, bool)`

GetJwtAccessTokenOk returns a tuple with the JwtAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtAccessToken

`func (o *AuthorizationIssueResponse) SetJwtAccessToken(v string)`

SetJwtAccessToken sets JwtAccessToken field to given value.

### HasJwtAccessToken

`func (o *AuthorizationIssueResponse) HasJwtAccessToken() bool`

HasJwtAccessToken returns a boolean if a field has been set.

### GetTicketInfo

`func (o *AuthorizationIssueResponse) GetTicketInfo() string`

GetTicketInfo returns the TicketInfo field if non-nil, zero value otherwise.

### GetTicketInfoOk

`func (o *AuthorizationIssueResponse) GetTicketInfoOk() (*string, bool)`

GetTicketInfoOk returns a tuple with the TicketInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketInfo

`func (o *AuthorizationIssueResponse) SetTicketInfo(v string)`

SetTicketInfo sets TicketInfo field to given value.

### HasTicketInfo

`func (o *AuthorizationIssueResponse) HasTicketInfo() bool`

HasTicketInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


