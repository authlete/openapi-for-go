# DeviceCompleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserCode** | **string** | A user code.  | 
**Result** | **string** | The result of the end-user authentication and authorization. One of the following. Details are described in the description.  | 
**Subject** | **string** | The subject (&#x3D; unique identifier) of the end-user.  | 
**Sub** | Pointer to **string** | The value of the sub claim that should be used in the ID token.  | [optional] 
**AuthTime** | Pointer to **string** | The time at which the end-user was authenticated. Its value is the number of seconds from &#x60;1970-01-01&#x60;.  | [optional] 
**Acr** | Pointer to **string** | The reference of the authentication context class which the end-user authentication satisfied.  | [optional] 
**Claims** | Pointer to **string** | Additional claims which will be embedded in the ID token.  | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) | The extra properties associated with the access token.  | [optional] 
**Scopes** | Pointer to **[]string** | Scopes to replace the scopes specified in the original device authorization request with. When nothing is specified for this parameter, replacement is not performed.  | [optional] 
**ErrorDescription** | Pointer to **string** | The description of the error. If this optional request parameter is given, its value is used as the value of the &#x60;error_description&#x60; property, but it is used only when the result is not &#x60;AUTHORIZED&#x60;. To comply with the specification strictly, the description must not include characters outside the set &#x60;%x20-21 / %x23-5B / %x5D-7E&#x60;.  | [optional] 
**ErrorUri** | Pointer to **string** | The URI of a document which describes the error in detail. This corresponds to the &#x60;error_uri&#x60; property in the response to the client.  | [optional] 
**IdHeaderParams** | Pointer to **string** | JSON that represents additional JWS header parameters for ID tokens. | [optional] 

## Methods

### NewDeviceCompleteRequest

`func NewDeviceCompleteRequest(userCode string, result string, subject string, ) *DeviceCompleteRequest`

NewDeviceCompleteRequest instantiates a new DeviceCompleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCompleteRequestWithDefaults

`func NewDeviceCompleteRequestWithDefaults() *DeviceCompleteRequest`

NewDeviceCompleteRequestWithDefaults instantiates a new DeviceCompleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserCode

`func (o *DeviceCompleteRequest) GetUserCode() string`

GetUserCode returns the UserCode field if non-nil, zero value otherwise.

### GetUserCodeOk

`func (o *DeviceCompleteRequest) GetUserCodeOk() (*string, bool)`

GetUserCodeOk returns a tuple with the UserCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCode

`func (o *DeviceCompleteRequest) SetUserCode(v string)`

SetUserCode sets UserCode field to given value.


### GetResult

`func (o *DeviceCompleteRequest) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DeviceCompleteRequest) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DeviceCompleteRequest) SetResult(v string)`

SetResult sets Result field to given value.


### GetSubject

`func (o *DeviceCompleteRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *DeviceCompleteRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *DeviceCompleteRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetSub

`func (o *DeviceCompleteRequest) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *DeviceCompleteRequest) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *DeviceCompleteRequest) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *DeviceCompleteRequest) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetAuthTime

`func (o *DeviceCompleteRequest) GetAuthTime() string`

GetAuthTime returns the AuthTime field if non-nil, zero value otherwise.

### GetAuthTimeOk

`func (o *DeviceCompleteRequest) GetAuthTimeOk() (*string, bool)`

GetAuthTimeOk returns a tuple with the AuthTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTime

`func (o *DeviceCompleteRequest) SetAuthTime(v string)`

SetAuthTime sets AuthTime field to given value.

### HasAuthTime

`func (o *DeviceCompleteRequest) HasAuthTime() bool`

HasAuthTime returns a boolean if a field has been set.

### GetAcr

`func (o *DeviceCompleteRequest) GetAcr() string`

GetAcr returns the Acr field if non-nil, zero value otherwise.

### GetAcrOk

`func (o *DeviceCompleteRequest) GetAcrOk() (*string, bool)`

GetAcrOk returns a tuple with the Acr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcr

`func (o *DeviceCompleteRequest) SetAcr(v string)`

SetAcr sets Acr field to given value.

### HasAcr

`func (o *DeviceCompleteRequest) HasAcr() bool`

HasAcr returns a boolean if a field has been set.

### GetClaims

`func (o *DeviceCompleteRequest) GetClaims() string`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *DeviceCompleteRequest) GetClaimsOk() (*string, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *DeviceCompleteRequest) SetClaims(v string)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *DeviceCompleteRequest) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetProperties

`func (o *DeviceCompleteRequest) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DeviceCompleteRequest) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DeviceCompleteRequest) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DeviceCompleteRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetScopes

`func (o *DeviceCompleteRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *DeviceCompleteRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *DeviceCompleteRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *DeviceCompleteRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetErrorDescription

`func (o *DeviceCompleteRequest) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *DeviceCompleteRequest) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *DeviceCompleteRequest) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *DeviceCompleteRequest) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorUri

`func (o *DeviceCompleteRequest) GetErrorUri() string`

GetErrorUri returns the ErrorUri field if non-nil, zero value otherwise.

### GetErrorUriOk

`func (o *DeviceCompleteRequest) GetErrorUriOk() (*string, bool)`

GetErrorUriOk returns a tuple with the ErrorUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUri

`func (o *DeviceCompleteRequest) SetErrorUri(v string)`

SetErrorUri sets ErrorUri field to given value.

### HasErrorUri

`func (o *DeviceCompleteRequest) HasErrorUri() bool`

HasErrorUri returns a boolean if a field has been set.

### GetIdHeaderParams

`func (o *DeviceCompleteRequest) GetIdHeaderParams() string`

GetIdHeaderParams returns the IdHeaderParams field if non-nil, zero value otherwise.

### GetIdHeaderParamsOk

`func (o *DeviceCompleteRequest) GetIdHeaderParamsOk() (*string, bool)`

GetIdHeaderParamsOk returns a tuple with the IdHeaderParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdHeaderParams

`func (o *DeviceCompleteRequest) SetIdHeaderParams(v string)`

SetIdHeaderParams sets IdHeaderParams field to given value.

### HasIdHeaderParams

`func (o *DeviceCompleteRequest) HasIdHeaderParams() bool`

HasIdHeaderParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


