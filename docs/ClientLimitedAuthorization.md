# ClientLimitedAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | The sequential number of the client. The value of this property is assigned by Authlete.  | [optional] [readonly] 
**ClientName** | Pointer to **string** | The name of the client application. This property corresponds to &#x60;client_name&#x60; in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**ClientNames** | Pointer to [**[]TaggedValue**](TaggedValue.md) | Client names with language tags. If the client application has different names for different languages, this property can be used to register the names.  | [optional] 
**Description** | Pointer to **string** | The description about the client application. | [optional] 
**Descriptions** | Pointer to [**[]TaggedValue**](TaggedValue.md) | Descriptions about the client application with language tags. If the client application has different descriptions for different languages, this property can be used to register the descriptions.  | [optional] 
**ClientId** | Pointer to **int64** | The client identifier used in Authlete API calls. The value of this property is assigned by Authlete. | [optional] [readonly] 
**ClientIdAlias** | Pointer to **string** | The value of the client&#39;s &#x60;client_id&#x60; property used in OAuth and OpenID Connect calls. By default, this is a string version of the &#x60;clientId&#x60; property.  | [optional] 
**ClientIdAliasEnabled** | Pointer to **bool** | Deprecated. Always set to &#x60;true&#x60;. | [optional] 
**ClientType** | Pointer to [**ClientType**](ClientType.md) |  | [optional] 
**LogoUri** | Pointer to **string** | The URL pointing to the logo image of the client application.  This property corresponds to &#x60;logo_uri&#x60; in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**LogoUris** | Pointer to [**[]TaggedValue**](TaggedValue.md) | Logo image URLs with language tags. If the client application has different logo images for different languages, this property can be used to register URLs of the images.  | [optional] 
**TosUri** | Pointer to **string** | The URL pointing to the \&quot;Terms Of Service\&quot; page.  This property corresponds to &#x60;tos_uri&#x60; in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**TosUris** | Pointer to [**[]TaggedValue**](TaggedValue.md) | URLs of \&quot;Terms Of Service\&quot; pages with language tags.  If the client application has different \&quot;Terms Of Service\&quot; pages for different languages, this property can be used to register the URLs.  | [optional] 
**PolicyUri** | Pointer to **string** | The URL pointing to the page which describes the policy as to how end-user&#39;s profile data is used.  This property corresponds to &#x60;policy_uri&#x60; in [OpenID Connect Dynamic Client Registration 1.0, 2. Client Metadata](https://openid.net/specs/openid-connect-registration-1_0.html#ClientMetadata).  | [optional] 
**PolicyUris** | Pointer to [**[]TaggedValue**](TaggedValue.md) | URLs of policy pages with language tags. If the client application has different policy pages for different languages, this property can be used to register the URLs.  | [optional] 

## Methods

### NewClientLimitedAuthorization

`func NewClientLimitedAuthorization() *ClientLimitedAuthorization`

NewClientLimitedAuthorization instantiates a new ClientLimitedAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientLimitedAuthorizationWithDefaults

`func NewClientLimitedAuthorizationWithDefaults() *ClientLimitedAuthorization`

NewClientLimitedAuthorizationWithDefaults instantiates a new ClientLimitedAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *ClientLimitedAuthorization) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ClientLimitedAuthorization) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ClientLimitedAuthorization) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ClientLimitedAuthorization) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetClientName

`func (o *ClientLimitedAuthorization) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ClientLimitedAuthorization) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ClientLimitedAuthorization) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *ClientLimitedAuthorization) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientNames

`func (o *ClientLimitedAuthorization) GetClientNames() []TaggedValue`

GetClientNames returns the ClientNames field if non-nil, zero value otherwise.

### GetClientNamesOk

`func (o *ClientLimitedAuthorization) GetClientNamesOk() (*[]TaggedValue, bool)`

GetClientNamesOk returns a tuple with the ClientNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNames

`func (o *ClientLimitedAuthorization) SetClientNames(v []TaggedValue)`

SetClientNames sets ClientNames field to given value.

### HasClientNames

`func (o *ClientLimitedAuthorization) HasClientNames() bool`

HasClientNames returns a boolean if a field has been set.

### GetDescription

`func (o *ClientLimitedAuthorization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientLimitedAuthorization) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientLimitedAuthorization) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientLimitedAuthorization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptions

`func (o *ClientLimitedAuthorization) GetDescriptions() []TaggedValue`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *ClientLimitedAuthorization) GetDescriptionsOk() (*[]TaggedValue, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *ClientLimitedAuthorization) SetDescriptions(v []TaggedValue)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *ClientLimitedAuthorization) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetClientId

`func (o *ClientLimitedAuthorization) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientLimitedAuthorization) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientLimitedAuthorization) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClientLimitedAuthorization) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIdAlias

`func (o *ClientLimitedAuthorization) GetClientIdAlias() string`

GetClientIdAlias returns the ClientIdAlias field if non-nil, zero value otherwise.

### GetClientIdAliasOk

`func (o *ClientLimitedAuthorization) GetClientIdAliasOk() (*string, bool)`

GetClientIdAliasOk returns a tuple with the ClientIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAlias

`func (o *ClientLimitedAuthorization) SetClientIdAlias(v string)`

SetClientIdAlias sets ClientIdAlias field to given value.

### HasClientIdAlias

`func (o *ClientLimitedAuthorization) HasClientIdAlias() bool`

HasClientIdAlias returns a boolean if a field has been set.

### GetClientIdAliasEnabled

`func (o *ClientLimitedAuthorization) GetClientIdAliasEnabled() bool`

GetClientIdAliasEnabled returns the ClientIdAliasEnabled field if non-nil, zero value otherwise.

### GetClientIdAliasEnabledOk

`func (o *ClientLimitedAuthorization) GetClientIdAliasEnabledOk() (*bool, bool)`

GetClientIdAliasEnabledOk returns a tuple with the ClientIdAliasEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasEnabled

`func (o *ClientLimitedAuthorization) SetClientIdAliasEnabled(v bool)`

SetClientIdAliasEnabled sets ClientIdAliasEnabled field to given value.

### HasClientIdAliasEnabled

`func (o *ClientLimitedAuthorization) HasClientIdAliasEnabled() bool`

HasClientIdAliasEnabled returns a boolean if a field has been set.

### GetClientType

`func (o *ClientLimitedAuthorization) GetClientType() ClientType`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *ClientLimitedAuthorization) GetClientTypeOk() (*ClientType, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *ClientLimitedAuthorization) SetClientType(v ClientType)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *ClientLimitedAuthorization) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetLogoUri

`func (o *ClientLimitedAuthorization) GetLogoUri() string`

GetLogoUri returns the LogoUri field if non-nil, zero value otherwise.

### GetLogoUriOk

`func (o *ClientLimitedAuthorization) GetLogoUriOk() (*string, bool)`

GetLogoUriOk returns a tuple with the LogoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUri

`func (o *ClientLimitedAuthorization) SetLogoUri(v string)`

SetLogoUri sets LogoUri field to given value.

### HasLogoUri

`func (o *ClientLimitedAuthorization) HasLogoUri() bool`

HasLogoUri returns a boolean if a field has been set.

### GetLogoUris

`func (o *ClientLimitedAuthorization) GetLogoUris() []TaggedValue`

GetLogoUris returns the LogoUris field if non-nil, zero value otherwise.

### GetLogoUrisOk

`func (o *ClientLimitedAuthorization) GetLogoUrisOk() (*[]TaggedValue, bool)`

GetLogoUrisOk returns a tuple with the LogoUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUris

`func (o *ClientLimitedAuthorization) SetLogoUris(v []TaggedValue)`

SetLogoUris sets LogoUris field to given value.

### HasLogoUris

`func (o *ClientLimitedAuthorization) HasLogoUris() bool`

HasLogoUris returns a boolean if a field has been set.

### GetTosUri

`func (o *ClientLimitedAuthorization) GetTosUri() string`

GetTosUri returns the TosUri field if non-nil, zero value otherwise.

### GetTosUriOk

`func (o *ClientLimitedAuthorization) GetTosUriOk() (*string, bool)`

GetTosUriOk returns a tuple with the TosUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosUri

`func (o *ClientLimitedAuthorization) SetTosUri(v string)`

SetTosUri sets TosUri field to given value.

### HasTosUri

`func (o *ClientLimitedAuthorization) HasTosUri() bool`

HasTosUri returns a boolean if a field has been set.

### GetTosUris

`func (o *ClientLimitedAuthorization) GetTosUris() []TaggedValue`

GetTosUris returns the TosUris field if non-nil, zero value otherwise.

### GetTosUrisOk

`func (o *ClientLimitedAuthorization) GetTosUrisOk() (*[]TaggedValue, bool)`

GetTosUrisOk returns a tuple with the TosUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosUris

`func (o *ClientLimitedAuthorization) SetTosUris(v []TaggedValue)`

SetTosUris sets TosUris field to given value.

### HasTosUris

`func (o *ClientLimitedAuthorization) HasTosUris() bool`

HasTosUris returns a boolean if a field has been set.

### GetPolicyUri

`func (o *ClientLimitedAuthorization) GetPolicyUri() string`

GetPolicyUri returns the PolicyUri field if non-nil, zero value otherwise.

### GetPolicyUriOk

`func (o *ClientLimitedAuthorization) GetPolicyUriOk() (*string, bool)`

GetPolicyUriOk returns a tuple with the PolicyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUri

`func (o *ClientLimitedAuthorization) SetPolicyUri(v string)`

SetPolicyUri sets PolicyUri field to given value.

### HasPolicyUri

`func (o *ClientLimitedAuthorization) HasPolicyUri() bool`

HasPolicyUri returns a boolean if a field has been set.

### GetPolicyUris

`func (o *ClientLimitedAuthorization) GetPolicyUris() []TaggedValue`

GetPolicyUris returns the PolicyUris field if non-nil, zero value otherwise.

### GetPolicyUrisOk

`func (o *ClientLimitedAuthorization) GetPolicyUrisOk() (*[]TaggedValue, bool)`

GetPolicyUrisOk returns a tuple with the PolicyUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUris

`func (o *ClientLimitedAuthorization) SetPolicyUris(v []TaggedValue)`

SetPolicyUris sets PolicyUris field to given value.

### HasPolicyUris

`func (o *ClientLimitedAuthorization) HasPolicyUris() bool`

HasPolicyUris returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


