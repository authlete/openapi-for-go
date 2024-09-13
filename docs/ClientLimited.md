# ClientLimited

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

## Methods

### NewClientLimited

`func NewClientLimited() *ClientLimited`

NewClientLimited instantiates a new ClientLimited object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientLimitedWithDefaults

`func NewClientLimitedWithDefaults() *ClientLimited`

NewClientLimitedWithDefaults instantiates a new ClientLimited object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *ClientLimited) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ClientLimited) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ClientLimited) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ClientLimited) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetClientName

`func (o *ClientLimited) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ClientLimited) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ClientLimited) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *ClientLimited) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientNames

`func (o *ClientLimited) GetClientNames() []TaggedValue`

GetClientNames returns the ClientNames field if non-nil, zero value otherwise.

### GetClientNamesOk

`func (o *ClientLimited) GetClientNamesOk() (*[]TaggedValue, bool)`

GetClientNamesOk returns a tuple with the ClientNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNames

`func (o *ClientLimited) SetClientNames(v []TaggedValue)`

SetClientNames sets ClientNames field to given value.

### HasClientNames

`func (o *ClientLimited) HasClientNames() bool`

HasClientNames returns a boolean if a field has been set.

### GetDescription

`func (o *ClientLimited) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientLimited) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientLimited) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientLimited) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptions

`func (o *ClientLimited) GetDescriptions() []TaggedValue`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *ClientLimited) GetDescriptionsOk() (*[]TaggedValue, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *ClientLimited) SetDescriptions(v []TaggedValue)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *ClientLimited) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetClientId

`func (o *ClientLimited) GetClientId() int64`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientLimited) GetClientIdOk() (*int64, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientLimited) SetClientId(v int64)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClientLimited) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientIdAlias

`func (o *ClientLimited) GetClientIdAlias() string`

GetClientIdAlias returns the ClientIdAlias field if non-nil, zero value otherwise.

### GetClientIdAliasOk

`func (o *ClientLimited) GetClientIdAliasOk() (*string, bool)`

GetClientIdAliasOk returns a tuple with the ClientIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAlias

`func (o *ClientLimited) SetClientIdAlias(v string)`

SetClientIdAlias sets ClientIdAlias field to given value.

### HasClientIdAlias

`func (o *ClientLimited) HasClientIdAlias() bool`

HasClientIdAlias returns a boolean if a field has been set.

### GetClientIdAliasEnabled

`func (o *ClientLimited) GetClientIdAliasEnabled() bool`

GetClientIdAliasEnabled returns the ClientIdAliasEnabled field if non-nil, zero value otherwise.

### GetClientIdAliasEnabledOk

`func (o *ClientLimited) GetClientIdAliasEnabledOk() (*bool, bool)`

GetClientIdAliasEnabledOk returns a tuple with the ClientIdAliasEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdAliasEnabled

`func (o *ClientLimited) SetClientIdAliasEnabled(v bool)`

SetClientIdAliasEnabled sets ClientIdAliasEnabled field to given value.

### HasClientIdAliasEnabled

`func (o *ClientLimited) HasClientIdAliasEnabled() bool`

HasClientIdAliasEnabled returns a boolean if a field has been set.

### GetClientType

`func (o *ClientLimited) GetClientType() ClientType`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *ClientLimited) GetClientTypeOk() (*ClientType, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *ClientLimited) SetClientType(v ClientType)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *ClientLimited) HasClientType() bool`

HasClientType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


