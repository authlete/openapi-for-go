# AuthzDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elements** | Pointer to [**[]AuthorizationDetailsElement**](AuthorizationDetailsElement.md) | Elements of this authorization details.  | [optional] 

## Methods

### NewAuthzDetails

`func NewAuthzDetails() *AuthzDetails`

NewAuthzDetails instantiates a new AuthzDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthzDetailsWithDefaults

`func NewAuthzDetailsWithDefaults() *AuthzDetails`

NewAuthzDetailsWithDefaults instantiates a new AuthzDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElements

`func (o *AuthzDetails) GetElements() []AuthorizationDetailsElement`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *AuthzDetails) GetElementsOk() (*[]AuthorizationDetailsElement, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *AuthzDetails) SetElements(v []AuthorizationDetailsElement)`

SetElements sets Elements field to given value.

### HasElements

`func (o *AuthzDetails) HasElements() bool`

HasElements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


