# StandardIntrospectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | **string** | Request parameters which comply with the introspection request defined in \&quot;[2.1. Introspection Request](https://datatracker.ietf.org/doc/html/rfc7662#section-2.1)\&quot; in RFC 7662.  The implementation of the introspection endpoint of your authorization server will receive an HTTP POST [[RFC 7231](https://datatracker.ietf.org/doc/html/rfc7231)] request with parameters in the &#x60;application/x-www-form-urlencoded&#x60; format. It is the entity body of the request that Authlete&#39;s  &#x60;/api/auth/introspection/standard&#x60; API expects as the value of &#x60;parameters&#x60;.  | 
**WithHiddenProperties** | Pointer to **string** | Flag indicating whether to include hidden properties in the output.  Authlete has a mechanism whereby to associate arbitrary key-value pairs with an access token. Each key-value pair has a hidden attribute. By default, key-value pairs whose hidden attribute is set to &#x60;true&#x60; are not embedded in the standard introspection output.  If the &#x60;withHiddenProperties&#x60; request parameter is given and its value is &#x60;true&#x60;, &#x60;/api/auth/introspection/standard API includes all the associated key-value pairs into the output regardless of the value of the hidden attribute. | [optional] 

## Methods

### NewStandardIntrospectionRequest

`func NewStandardIntrospectionRequest(parameters string, ) *StandardIntrospectionRequest`

NewStandardIntrospectionRequest instantiates a new StandardIntrospectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardIntrospectionRequestWithDefaults

`func NewStandardIntrospectionRequestWithDefaults() *StandardIntrospectionRequest`

NewStandardIntrospectionRequestWithDefaults instantiates a new StandardIntrospectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *StandardIntrospectionRequest) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *StandardIntrospectionRequest) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *StandardIntrospectionRequest) SetParameters(v string)`

SetParameters sets Parameters field to given value.


### GetWithHiddenProperties

`func (o *StandardIntrospectionRequest) GetWithHiddenProperties() string`

GetWithHiddenProperties returns the WithHiddenProperties field if non-nil, zero value otherwise.

### GetWithHiddenPropertiesOk

`func (o *StandardIntrospectionRequest) GetWithHiddenPropertiesOk() (*string, bool)`

GetWithHiddenPropertiesOk returns a tuple with the WithHiddenProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithHiddenProperties

`func (o *StandardIntrospectionRequest) SetWithHiddenProperties(v string)`

SetWithHiddenProperties sets WithHiddenProperties field to given value.

### HasWithHiddenProperties

`func (o *StandardIntrospectionRequest) HasWithHiddenProperties() bool`

HasWithHiddenProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


