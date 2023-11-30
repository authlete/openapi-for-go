# HskCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kty** | Pointer to **string** | The key type (EC or RSA)  | [optional] 
**Use** | Pointer to **string** | The key on the HSM.  When the key use is \&quot;sig\&quot; (signature), the private key on the HSM is used to sign data and the corresponding public key is used to verify the signature. When the key use is \&quot;enc\&quot; (encryption), the private key on the HSM is used to decrypt encrypted data which have been encrypted with the corresponding public key  | [optional] 
**Kid** | Pointer to **string** | Key ID for the key on the HSM.  | [optional] 
**HsmName** | Pointer to **string** | The name of the HSM. The identifier for the HSM that sits behind the Authlete server. For example, \&quot;google\&quot;.  | [optional] 
**Handle** | Pointer to **string** | The handle for the key on the HSM. A handle is a base64url-encoded 256-bit random value (43 letters) which is assigned by Authlete on the call of the /api/hsk/create API  | [optional] 
**PublicKey** | Pointer to **string** | The public key that corresponds to the key on the HSM. | [optional] 

## Methods

### NewHskCreateRequest

`func NewHskCreateRequest() *HskCreateRequest`

NewHskCreateRequest instantiates a new HskCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHskCreateRequestWithDefaults

`func NewHskCreateRequestWithDefaults() *HskCreateRequest`

NewHskCreateRequestWithDefaults instantiates a new HskCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKty

`func (o *HskCreateRequest) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *HskCreateRequest) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *HskCreateRequest) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *HskCreateRequest) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetUse

`func (o *HskCreateRequest) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *HskCreateRequest) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *HskCreateRequest) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *HskCreateRequest) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetKid

`func (o *HskCreateRequest) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *HskCreateRequest) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *HskCreateRequest) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *HskCreateRequest) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetHsmName

`func (o *HskCreateRequest) GetHsmName() string`

GetHsmName returns the HsmName field if non-nil, zero value otherwise.

### GetHsmNameOk

`func (o *HskCreateRequest) GetHsmNameOk() (*string, bool)`

GetHsmNameOk returns a tuple with the HsmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmName

`func (o *HskCreateRequest) SetHsmName(v string)`

SetHsmName sets HsmName field to given value.

### HasHsmName

`func (o *HskCreateRequest) HasHsmName() bool`

HasHsmName returns a boolean if a field has been set.

### GetHandle

`func (o *HskCreateRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *HskCreateRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *HskCreateRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *HskCreateRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetPublicKey

`func (o *HskCreateRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *HskCreateRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *HskCreateRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *HskCreateRequest) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


