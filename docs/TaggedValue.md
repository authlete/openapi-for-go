# TaggedValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to **string** | The language tag part. | [optional] 
**Value** | Pointer to **string** | The value part. | [optional] 

## Methods

### NewTaggedValue

`func NewTaggedValue() *TaggedValue`

NewTaggedValue instantiates a new TaggedValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaggedValueWithDefaults

`func NewTaggedValueWithDefaults() *TaggedValue`

NewTaggedValueWithDefaults instantiates a new TaggedValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *TaggedValue) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TaggedValue) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TaggedValue) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *TaggedValue) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetValue

`func (o *TaggedValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TaggedValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TaggedValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TaggedValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


