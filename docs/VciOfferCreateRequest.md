# VciOfferCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to **string** | The value of the &#x60;credentials&#x60; object in the JSON format. | [optional] 
**AuthorizationCodeGrantIncluded** | Pointer to **bool** | The flag indicating whether the &#x60;authorization_code&#x60; object is included in the &#x60;grants&#x60; object.  | [optional] 
**IssuerStateIncluded** | Pointer to **bool** | The flag indicating whether the &#x60;issuer_state&#x60; property is included in the &#x60;authorization_code&#x60; object in the &#x60;grants&#x60; object.  | [optional] 
**PreAuthorizedCodeGrantIncluded** | Pointer to **bool** | The flag to include the &#x60;urn:ietf:params:oauth:grant-type:pre-authorized_code&#x60; object in the &#x60;grants&#x60; object.  | [optional] 
**UserPinRequired** | Pointer to **bool** | The value of the &#x60;user_pin_required&#x60; property in the &#x60;urn:ietf:params:oauth:grant-type:pre-authorized_code&#x60; object in the &#x60;grants&#x60; object.  | [optional] 
**UserPinLength** | Pointer to **int32** | The length of the user PIN to generate. | [optional] 
**Subject** | Pointer to **string** | The subject associated with the credential offer. | [optional] 
**Duration** | Pointer to **int64** | The duration of the credential offer. | [optional] 
**Context** | Pointer to **string** | The general-purpose arbitrary string. | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) | Extra properties to associate with the credential offer. | [optional] 
**JwtAtClaims** | Pointer to **string** | Additional claims that are added to the payload part of the JWT access token.  | [optional] 
**AuthTime** | Pointer to **int64** | The time at which the user authentication was performed during the course of issuing the credential offer.  | [optional] 
**Acr** | Pointer to **string** | The Authentication Context Class Reference of the user authentication performed during the course of issuing the credential offer. | [optional] 

## Methods

### NewVciOfferCreateRequest

`func NewVciOfferCreateRequest() *VciOfferCreateRequest`

NewVciOfferCreateRequest instantiates a new VciOfferCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVciOfferCreateRequestWithDefaults

`func NewVciOfferCreateRequestWithDefaults() *VciOfferCreateRequest`

NewVciOfferCreateRequestWithDefaults instantiates a new VciOfferCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *VciOfferCreateRequest) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *VciOfferCreateRequest) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *VciOfferCreateRequest) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *VciOfferCreateRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetAuthorizationCodeGrantIncluded

`func (o *VciOfferCreateRequest) GetAuthorizationCodeGrantIncluded() bool`

GetAuthorizationCodeGrantIncluded returns the AuthorizationCodeGrantIncluded field if non-nil, zero value otherwise.

### GetAuthorizationCodeGrantIncludedOk

`func (o *VciOfferCreateRequest) GetAuthorizationCodeGrantIncludedOk() (*bool, bool)`

GetAuthorizationCodeGrantIncludedOk returns a tuple with the AuthorizationCodeGrantIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCodeGrantIncluded

`func (o *VciOfferCreateRequest) SetAuthorizationCodeGrantIncluded(v bool)`

SetAuthorizationCodeGrantIncluded sets AuthorizationCodeGrantIncluded field to given value.

### HasAuthorizationCodeGrantIncluded

`func (o *VciOfferCreateRequest) HasAuthorizationCodeGrantIncluded() bool`

HasAuthorizationCodeGrantIncluded returns a boolean if a field has been set.

### GetIssuerStateIncluded

`func (o *VciOfferCreateRequest) GetIssuerStateIncluded() bool`

GetIssuerStateIncluded returns the IssuerStateIncluded field if non-nil, zero value otherwise.

### GetIssuerStateIncludedOk

`func (o *VciOfferCreateRequest) GetIssuerStateIncludedOk() (*bool, bool)`

GetIssuerStateIncludedOk returns a tuple with the IssuerStateIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerStateIncluded

`func (o *VciOfferCreateRequest) SetIssuerStateIncluded(v bool)`

SetIssuerStateIncluded sets IssuerStateIncluded field to given value.

### HasIssuerStateIncluded

`func (o *VciOfferCreateRequest) HasIssuerStateIncluded() bool`

HasIssuerStateIncluded returns a boolean if a field has been set.

### GetPreAuthorizedCodeGrantIncluded

`func (o *VciOfferCreateRequest) GetPreAuthorizedCodeGrantIncluded() bool`

GetPreAuthorizedCodeGrantIncluded returns the PreAuthorizedCodeGrantIncluded field if non-nil, zero value otherwise.

### GetPreAuthorizedCodeGrantIncludedOk

`func (o *VciOfferCreateRequest) GetPreAuthorizedCodeGrantIncludedOk() (*bool, bool)`

GetPreAuthorizedCodeGrantIncludedOk returns a tuple with the PreAuthorizedCodeGrantIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthorizedCodeGrantIncluded

`func (o *VciOfferCreateRequest) SetPreAuthorizedCodeGrantIncluded(v bool)`

SetPreAuthorizedCodeGrantIncluded sets PreAuthorizedCodeGrantIncluded field to given value.

### HasPreAuthorizedCodeGrantIncluded

`func (o *VciOfferCreateRequest) HasPreAuthorizedCodeGrantIncluded() bool`

HasPreAuthorizedCodeGrantIncluded returns a boolean if a field has been set.

### GetUserPinRequired

`func (o *VciOfferCreateRequest) GetUserPinRequired() bool`

GetUserPinRequired returns the UserPinRequired field if non-nil, zero value otherwise.

### GetUserPinRequiredOk

`func (o *VciOfferCreateRequest) GetUserPinRequiredOk() (*bool, bool)`

GetUserPinRequiredOk returns a tuple with the UserPinRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPinRequired

`func (o *VciOfferCreateRequest) SetUserPinRequired(v bool)`

SetUserPinRequired sets UserPinRequired field to given value.

### HasUserPinRequired

`func (o *VciOfferCreateRequest) HasUserPinRequired() bool`

HasUserPinRequired returns a boolean if a field has been set.

### GetUserPinLength

`func (o *VciOfferCreateRequest) GetUserPinLength() int32`

GetUserPinLength returns the UserPinLength field if non-nil, zero value otherwise.

### GetUserPinLengthOk

`func (o *VciOfferCreateRequest) GetUserPinLengthOk() (*int32, bool)`

GetUserPinLengthOk returns a tuple with the UserPinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPinLength

`func (o *VciOfferCreateRequest) SetUserPinLength(v int32)`

SetUserPinLength sets UserPinLength field to given value.

### HasUserPinLength

`func (o *VciOfferCreateRequest) HasUserPinLength() bool`

HasUserPinLength returns a boolean if a field has been set.

### GetSubject

`func (o *VciOfferCreateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *VciOfferCreateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *VciOfferCreateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *VciOfferCreateRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetDuration

`func (o *VciOfferCreateRequest) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VciOfferCreateRequest) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VciOfferCreateRequest) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VciOfferCreateRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetContext

`func (o *VciOfferCreateRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *VciOfferCreateRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *VciOfferCreateRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *VciOfferCreateRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetProperties

`func (o *VciOfferCreateRequest) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *VciOfferCreateRequest) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *VciOfferCreateRequest) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *VciOfferCreateRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetJwtAtClaims

`func (o *VciOfferCreateRequest) GetJwtAtClaims() string`

GetJwtAtClaims returns the JwtAtClaims field if non-nil, zero value otherwise.

### GetJwtAtClaimsOk

`func (o *VciOfferCreateRequest) GetJwtAtClaimsOk() (*string, bool)`

GetJwtAtClaimsOk returns a tuple with the JwtAtClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtAtClaims

`func (o *VciOfferCreateRequest) SetJwtAtClaims(v string)`

SetJwtAtClaims sets JwtAtClaims field to given value.

### HasJwtAtClaims

`func (o *VciOfferCreateRequest) HasJwtAtClaims() bool`

HasJwtAtClaims returns a boolean if a field has been set.

### GetAuthTime

`func (o *VciOfferCreateRequest) GetAuthTime() int64`

GetAuthTime returns the AuthTime field if non-nil, zero value otherwise.

### GetAuthTimeOk

`func (o *VciOfferCreateRequest) GetAuthTimeOk() (*int64, bool)`

GetAuthTimeOk returns a tuple with the AuthTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTime

`func (o *VciOfferCreateRequest) SetAuthTime(v int64)`

SetAuthTime sets AuthTime field to given value.

### HasAuthTime

`func (o *VciOfferCreateRequest) HasAuthTime() bool`

HasAuthTime returns a boolean if a field has been set.

### GetAcr

`func (o *VciOfferCreateRequest) GetAcr() string`

GetAcr returns the Acr field if non-nil, zero value otherwise.

### GetAcrOk

`func (o *VciOfferCreateRequest) GetAcrOk() (*string, bool)`

GetAcrOk returns a tuple with the Acr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcr

`func (o *VciOfferCreateRequest) SetAcr(v string)`

SetAcr sets Acr field to given value.

### HasAcr

`func (o *VciOfferCreateRequest) HasAcr() bool`

HasAcr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


