/*
Authlete API

Authlete API Document. 

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

// checks if the VciBatchIssueRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VciBatchIssueRequest{}

// VciBatchIssueRequest struct for VciBatchIssueRequest
type VciBatchIssueRequest struct {
	// The access token that came along with the credential request.
	AccessToken *string `json:"accessToken,omitempty"`
	// The instructions for issuance of credentials and/or transaction IDs.
	Orders []CredentialIssuanceOrder `json:"orders,omitempty"`
}

// NewVciBatchIssueRequest instantiates a new VciBatchIssueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVciBatchIssueRequest() *VciBatchIssueRequest {
	this := VciBatchIssueRequest{}
	return &this
}

// NewVciBatchIssueRequestWithDefaults instantiates a new VciBatchIssueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVciBatchIssueRequestWithDefaults() *VciBatchIssueRequest {
	this := VciBatchIssueRequest{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *VciBatchIssueRequest) GetAccessToken() string {
	if o == nil || isNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciBatchIssueRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil || isNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *VciBatchIssueRequest) HasAccessToken() bool {
	if o != nil && !isNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *VciBatchIssueRequest) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *VciBatchIssueRequest) GetOrders() []CredentialIssuanceOrder {
	if o == nil || isNil(o.Orders) {
		var ret []CredentialIssuanceOrder
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VciBatchIssueRequest) GetOrdersOk() ([]CredentialIssuanceOrder, bool) {
	if o == nil || isNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *VciBatchIssueRequest) HasOrders() bool {
	if o != nil && !isNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []CredentialIssuanceOrder and assigns it to the Orders field.
func (o *VciBatchIssueRequest) SetOrders(v []CredentialIssuanceOrder) {
	o.Orders = v
}

func (o VciBatchIssueRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VciBatchIssueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if !isNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}
	return toSerialize, nil
}

type NullableVciBatchIssueRequest struct {
	value *VciBatchIssueRequest
	isSet bool
}

func (v NullableVciBatchIssueRequest) Get() *VciBatchIssueRequest {
	return v.value
}

func (v *NullableVciBatchIssueRequest) Set(val *VciBatchIssueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVciBatchIssueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVciBatchIssueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVciBatchIssueRequest(val *VciBatchIssueRequest) *NullableVciBatchIssueRequest {
	return &NullableVciBatchIssueRequest{value: val, isSet: true}
}

func (v NullableVciBatchIssueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVciBatchIssueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


