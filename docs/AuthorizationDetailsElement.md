# AuthorizationDetailsElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of this element.  From _\&quot;OAuth 2.0 Rich Authorization Requests\&quot;_: _\&quot;The type of authorization data as a string. This field MAY define which other elements are allowed in the request. This element is REQUIRED.\&quot;_  This property is always NOT &#x60;null&#x60;.  | 
**Locations** | Pointer to **[]string** | The resources and/or resource servers. This property may be &#x60;null&#x60;.  From _\&quot;OAuth 2.0 Rich Authorization Requests\&quot;_: _\&quot;An array of strings representing the location of the resource or resource server. This is typically composed of URIs.\&quot;_  This property may be &#x60;null&#x60;.  | [optional] 
**Actions** | Pointer to **[]string** | The actions.  From _\&quot;OAuth 2.0 Rich Authorization Requests\&quot;_: _\&quot;An array of strings representing the kinds of actions to be taken at the resource. The values of the strings are determined by the API being protected.\&quot;_  This property may be &#x60;null&#x60;.  | [optional] 
**DataTypes** | Pointer to **[]string** | From _\&quot;OAuth 2.0 Rich Authorization Requests\&quot;_: _\&quot;An array of strings representing the kinds of data being requested from the resource.\&quot;_  This property may be &#x60;null&#x60;.  | [optional] 
**Identifier** | Pointer to **string** | The identifier of a specific resource. From _\&quot;OAuth 2.0 Rich Authorization Requests\&quot;_: _\&quot;A string identifier indicating a specific resource available at the API.\&quot;_  This property may be &#x60;null&#x60;.  | [optional] 
**Privileges** | Pointer to **[]string** | The types or levels of privilege. From \&quot;OAuth 2.0 Rich Authorization Requests\&quot;: _\&quot;An array of strings representing the types or levels of privilege being requested at the resource.\&quot;_  This property may be &#x60;null&#x60;.  | [optional] 

## Methods

### NewAuthorizationDetailsElement

`func NewAuthorizationDetailsElement(type_ string, ) *AuthorizationDetailsElement`

NewAuthorizationDetailsElement instantiates a new AuthorizationDetailsElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationDetailsElementWithDefaults

`func NewAuthorizationDetailsElementWithDefaults() *AuthorizationDetailsElement`

NewAuthorizationDetailsElementWithDefaults instantiates a new AuthorizationDetailsElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthorizationDetailsElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizationDetailsElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizationDetailsElement) SetType(v string)`

SetType sets Type field to given value.


### GetLocations

`func (o *AuthorizationDetailsElement) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *AuthorizationDetailsElement) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *AuthorizationDetailsElement) SetLocations(v []string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *AuthorizationDetailsElement) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetActions

`func (o *AuthorizationDetailsElement) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AuthorizationDetailsElement) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AuthorizationDetailsElement) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AuthorizationDetailsElement) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetDataTypes

`func (o *AuthorizationDetailsElement) GetDataTypes() []string`

GetDataTypes returns the DataTypes field if non-nil, zero value otherwise.

### GetDataTypesOk

`func (o *AuthorizationDetailsElement) GetDataTypesOk() (*[]string, bool)`

GetDataTypesOk returns a tuple with the DataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTypes

`func (o *AuthorizationDetailsElement) SetDataTypes(v []string)`

SetDataTypes sets DataTypes field to given value.

### HasDataTypes

`func (o *AuthorizationDetailsElement) HasDataTypes() bool`

HasDataTypes returns a boolean if a field has been set.

### GetIdentifier

`func (o *AuthorizationDetailsElement) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AuthorizationDetailsElement) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AuthorizationDetailsElement) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *AuthorizationDetailsElement) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetPrivileges

`func (o *AuthorizationDetailsElement) GetPrivileges() []string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *AuthorizationDetailsElement) GetPrivilegesOk() (*[]string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *AuthorizationDetailsElement) SetPrivileges(v []string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *AuthorizationDetailsElement) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


