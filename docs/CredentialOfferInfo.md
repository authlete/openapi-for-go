# CredentialOfferInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The identifier of the credential offer. | [optional] 
**CredentialOffer** | Pointer to **string** | The credential offer in the JSON format. | [optional] 
**CredentialIssuer** | Pointer to **string** | The identifier of the credential issuer. | [optional] 
**Credentials** | Pointer to **string** | The value of the &#x60;credentials&#x60; object in the JSON format. | [optional] 
**AuthorizationCodeGrantIncluded** | Pointer to **bool** | The flag indicating whether the &#x60;authorization_code&#x60; object is included in the &#x60;grants&#x60; object.  | [optional] 
**IssuerStateIncluded** | Pointer to **bool** | The flag indicating whether the &#x60;issuer_state&#x60; property is included in the &#x60;authorization_code&#x60; object in the &#x60;grants&#x60; object.  | [optional] 
**IssuerState** | Pointer to **string** | The value of the &#x60;issuer_state&#x60; property in the &#x60;authorization_code&#x60; object in the &#x60;grants&#x60; object.  | [optional] 
**PreAuthorizedCodeGrantIncluded** | Pointer to **bool** | The flag indicating whether the &#x60;urn:ietf:params:oauth:grant-type:pre-authorized_code&#x60; object is included in the &#x60;grants&#x60; object.  | [optional] 
**PreAuthorizedCode** | Pointer to **string** | The value of the &#x60;pre-authorized_code&#x60; property in the &#x60;urn:ietf:params:oauth:grant-type:pre-authorized_code&#x60; object in the &#x60;grants&#x60; object.  | [optional] 
**UserPinRequired** | Pointer to **bool** | The value of the &#x60;user_pin_required&#x60; property in the &#x60;urn:ietf:params:oauth:grant-type:pre-authorized_code&#x60; object in the &#x60;grants&#x60; object.  | [optional] 
**UserPin** | Pointer to **string** | The value of the user PIN associated with the credential offer. | [optional] 
**Subject** | Pointer to **string** | The subject associated with the credential offer. | [optional] 
**ExpiresAt** | Pointer to **int64** | The time at which the credential offer will expire. | [optional] 
**Context** | Pointer to **string** | The general-purpose arbitrary string. | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) | Extra properties to associate with the credential offer. | [optional] 
**JwtAtClaims** | Pointer to **string** | Additional claims that are added to the payload part of the JWT access token.  | [optional] 
**AuthTime** | Pointer to **int64** | The time at which the user authentication was performed during the course of issuing the credential offer.  | [optional] 
**Acr** | Pointer to **string** | The Authentication Context Class Reference of the user authentication performed during the course of issuing the credential offer. | [optional] 

## Methods

### NewCredentialOfferInfo

`func NewCredentialOfferInfo() *CredentialOfferInfo`

NewCredentialOfferInfo instantiates a new CredentialOfferInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialOfferInfoWithDefaults

`func NewCredentialOfferInfoWithDefaults() *CredentialOfferInfo`

NewCredentialOfferInfoWithDefaults instantiates a new CredentialOfferInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *CredentialOfferInfo) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CredentialOfferInfo) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CredentialOfferInfo) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CredentialOfferInfo) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetCredentialOffer

`func (o *CredentialOfferInfo) GetCredentialOffer() string`

GetCredentialOffer returns the CredentialOffer field if non-nil, zero value otherwise.

### GetCredentialOfferOk

`func (o *CredentialOfferInfo) GetCredentialOfferOk() (*string, bool)`

GetCredentialOfferOk returns a tuple with the CredentialOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialOffer

`func (o *CredentialOfferInfo) SetCredentialOffer(v string)`

SetCredentialOffer sets CredentialOffer field to given value.

### HasCredentialOffer

`func (o *CredentialOfferInfo) HasCredentialOffer() bool`

HasCredentialOffer returns a boolean if a field has been set.

### GetCredentialIssuer

`func (o *CredentialOfferInfo) GetCredentialIssuer() string`

GetCredentialIssuer returns the CredentialIssuer field if non-nil, zero value otherwise.

### GetCredentialIssuerOk

`func (o *CredentialOfferInfo) GetCredentialIssuerOk() (*string, bool)`

GetCredentialIssuerOk returns a tuple with the CredentialIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialIssuer

`func (o *CredentialOfferInfo) SetCredentialIssuer(v string)`

SetCredentialIssuer sets CredentialIssuer field to given value.

### HasCredentialIssuer

`func (o *CredentialOfferInfo) HasCredentialIssuer() bool`

HasCredentialIssuer returns a boolean if a field has been set.

### GetCredentials

`func (o *CredentialOfferInfo) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CredentialOfferInfo) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CredentialOfferInfo) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CredentialOfferInfo) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetAuthorizationCodeGrantIncluded

`func (o *CredentialOfferInfo) GetAuthorizationCodeGrantIncluded() bool`

GetAuthorizationCodeGrantIncluded returns the AuthorizationCodeGrantIncluded field if non-nil, zero value otherwise.

### GetAuthorizationCodeGrantIncludedOk

`func (o *CredentialOfferInfo) GetAuthorizationCodeGrantIncludedOk() (*bool, bool)`

GetAuthorizationCodeGrantIncludedOk returns a tuple with the AuthorizationCodeGrantIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCodeGrantIncluded

`func (o *CredentialOfferInfo) SetAuthorizationCodeGrantIncluded(v bool)`

SetAuthorizationCodeGrantIncluded sets AuthorizationCodeGrantIncluded field to given value.

### HasAuthorizationCodeGrantIncluded

`func (o *CredentialOfferInfo) HasAuthorizationCodeGrantIncluded() bool`

HasAuthorizationCodeGrantIncluded returns a boolean if a field has been set.

### GetIssuerStateIncluded

`func (o *CredentialOfferInfo) GetIssuerStateIncluded() bool`

GetIssuerStateIncluded returns the IssuerStateIncluded field if non-nil, zero value otherwise.

### GetIssuerStateIncludedOk

`func (o *CredentialOfferInfo) GetIssuerStateIncludedOk() (*bool, bool)`

GetIssuerStateIncludedOk returns a tuple with the IssuerStateIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerStateIncluded

`func (o *CredentialOfferInfo) SetIssuerStateIncluded(v bool)`

SetIssuerStateIncluded sets IssuerStateIncluded field to given value.

### HasIssuerStateIncluded

`func (o *CredentialOfferInfo) HasIssuerStateIncluded() bool`

HasIssuerStateIncluded returns a boolean if a field has been set.

### GetIssuerState

`func (o *CredentialOfferInfo) GetIssuerState() string`

GetIssuerState returns the IssuerState field if non-nil, zero value otherwise.

### GetIssuerStateOk

`func (o *CredentialOfferInfo) GetIssuerStateOk() (*string, bool)`

GetIssuerStateOk returns a tuple with the IssuerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerState

`func (o *CredentialOfferInfo) SetIssuerState(v string)`

SetIssuerState sets IssuerState field to given value.

### HasIssuerState

`func (o *CredentialOfferInfo) HasIssuerState() bool`

HasIssuerState returns a boolean if a field has been set.

### GetPreAuthorizedCodeGrantIncluded

`func (o *CredentialOfferInfo) GetPreAuthorizedCodeGrantIncluded() bool`

GetPreAuthorizedCodeGrantIncluded returns the PreAuthorizedCodeGrantIncluded field if non-nil, zero value otherwise.

### GetPreAuthorizedCodeGrantIncludedOk

`func (o *CredentialOfferInfo) GetPreAuthorizedCodeGrantIncludedOk() (*bool, bool)`

GetPreAuthorizedCodeGrantIncludedOk returns a tuple with the PreAuthorizedCodeGrantIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthorizedCodeGrantIncluded

`func (o *CredentialOfferInfo) SetPreAuthorizedCodeGrantIncluded(v bool)`

SetPreAuthorizedCodeGrantIncluded sets PreAuthorizedCodeGrantIncluded field to given value.

### HasPreAuthorizedCodeGrantIncluded

`func (o *CredentialOfferInfo) HasPreAuthorizedCodeGrantIncluded() bool`

HasPreAuthorizedCodeGrantIncluded returns a boolean if a field has been set.

### GetPreAuthorizedCode

`func (o *CredentialOfferInfo) GetPreAuthorizedCode() string`

GetPreAuthorizedCode returns the PreAuthorizedCode field if non-nil, zero value otherwise.

### GetPreAuthorizedCodeOk

`func (o *CredentialOfferInfo) GetPreAuthorizedCodeOk() (*string, bool)`

GetPreAuthorizedCodeOk returns a tuple with the PreAuthorizedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthorizedCode

`func (o *CredentialOfferInfo) SetPreAuthorizedCode(v string)`

SetPreAuthorizedCode sets PreAuthorizedCode field to given value.

### HasPreAuthorizedCode

`func (o *CredentialOfferInfo) HasPreAuthorizedCode() bool`

HasPreAuthorizedCode returns a boolean if a field has been set.

### GetUserPinRequired

`func (o *CredentialOfferInfo) GetUserPinRequired() bool`

GetUserPinRequired returns the UserPinRequired field if non-nil, zero value otherwise.

### GetUserPinRequiredOk

`func (o *CredentialOfferInfo) GetUserPinRequiredOk() (*bool, bool)`

GetUserPinRequiredOk returns a tuple with the UserPinRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPinRequired

`func (o *CredentialOfferInfo) SetUserPinRequired(v bool)`

SetUserPinRequired sets UserPinRequired field to given value.

### HasUserPinRequired

`func (o *CredentialOfferInfo) HasUserPinRequired() bool`

HasUserPinRequired returns a boolean if a field has been set.

### GetUserPin

`func (o *CredentialOfferInfo) GetUserPin() string`

GetUserPin returns the UserPin field if non-nil, zero value otherwise.

### GetUserPinOk

`func (o *CredentialOfferInfo) GetUserPinOk() (*string, bool)`

GetUserPinOk returns a tuple with the UserPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPin

`func (o *CredentialOfferInfo) SetUserPin(v string)`

SetUserPin sets UserPin field to given value.

### HasUserPin

`func (o *CredentialOfferInfo) HasUserPin() bool`

HasUserPin returns a boolean if a field has been set.

### GetSubject

`func (o *CredentialOfferInfo) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CredentialOfferInfo) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CredentialOfferInfo) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CredentialOfferInfo) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CredentialOfferInfo) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CredentialOfferInfo) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CredentialOfferInfo) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CredentialOfferInfo) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetContext

`func (o *CredentialOfferInfo) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CredentialOfferInfo) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CredentialOfferInfo) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *CredentialOfferInfo) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetProperties

`func (o *CredentialOfferInfo) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CredentialOfferInfo) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CredentialOfferInfo) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CredentialOfferInfo) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetJwtAtClaims

`func (o *CredentialOfferInfo) GetJwtAtClaims() string`

GetJwtAtClaims returns the JwtAtClaims field if non-nil, zero value otherwise.

### GetJwtAtClaimsOk

`func (o *CredentialOfferInfo) GetJwtAtClaimsOk() (*string, bool)`

GetJwtAtClaimsOk returns a tuple with the JwtAtClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtAtClaims

`func (o *CredentialOfferInfo) SetJwtAtClaims(v string)`

SetJwtAtClaims sets JwtAtClaims field to given value.

### HasJwtAtClaims

`func (o *CredentialOfferInfo) HasJwtAtClaims() bool`

HasJwtAtClaims returns a boolean if a field has been set.

### GetAuthTime

`func (o *CredentialOfferInfo) GetAuthTime() int64`

GetAuthTime returns the AuthTime field if non-nil, zero value otherwise.

### GetAuthTimeOk

`func (o *CredentialOfferInfo) GetAuthTimeOk() (*int64, bool)`

GetAuthTimeOk returns a tuple with the AuthTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTime

`func (o *CredentialOfferInfo) SetAuthTime(v int64)`

SetAuthTime sets AuthTime field to given value.

### HasAuthTime

`func (o *CredentialOfferInfo) HasAuthTime() bool`

HasAuthTime returns a boolean if a field has been set.

### GetAcr

`func (o *CredentialOfferInfo) GetAcr() string`

GetAcr returns the Acr field if non-nil, zero value otherwise.

### GetAcrOk

`func (o *CredentialOfferInfo) GetAcrOk() (*string, bool)`

GetAcrOk returns a tuple with the Acr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcr

`func (o *CredentialOfferInfo) SetAcr(v string)`

SetAcr sets Acr field to given value.

### HasAcr

`func (o *CredentialOfferInfo) HasAcr() bool`

HasAcr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


