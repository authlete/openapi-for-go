/*
Authlete API

Authlete API Document. 

API version: 2.2.25
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

// ClientGetListResponse struct for ClientGetListResponse
type ClientGetListResponse struct {
	// The developer of the client applications. If the request did not contain `developer` request parameter, this property is set to `null`. 
	Developer *string `json:"developer,omitempty"`
	// Start index (inclusive) of the result set of the query. 
	Start *int32 `json:"start,omitempty"`
	// End index (exclusive) of the result set of the query. 
	End *int32 `json:"end,omitempty"`
	// Total number of clients that belong to the service. This doesn't mean the number of clients contained in the response. 
	TotalCount *int32 `json:"totalCount,omitempty"`
	// An array of clients. 
	Clients []Client `json:"clients,omitempty"`
}

// NewClientGetListResponse instantiates a new ClientGetListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientGetListResponse() *ClientGetListResponse {
	this := ClientGetListResponse{}
	return &this
}

// NewClientGetListResponseWithDefaults instantiates a new ClientGetListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientGetListResponseWithDefaults() *ClientGetListResponse {
	this := ClientGetListResponse{}
	return &this
}

// GetDeveloper returns the Developer field value if set, zero value otherwise.
func (o *ClientGetListResponse) GetDeveloper() string {
	if o == nil || o.Developer == nil {
		var ret string
		return ret
	}
	return *o.Developer
}

// GetDeveloperOk returns a tuple with the Developer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGetListResponse) GetDeveloperOk() (*string, bool) {
	if o == nil || o.Developer == nil {
		return nil, false
	}
	return o.Developer, true
}

// HasDeveloper returns a boolean if a field has been set.
func (o *ClientGetListResponse) HasDeveloper() bool {
	if o != nil && o.Developer != nil {
		return true
	}

	return false
}

// SetDeveloper gets a reference to the given string and assigns it to the Developer field.
func (o *ClientGetListResponse) SetDeveloper(v string) {
	o.Developer = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ClientGetListResponse) GetStart() int32 {
	if o == nil || o.Start == nil {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGetListResponse) GetStartOk() (*int32, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ClientGetListResponse) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *ClientGetListResponse) SetStart(v int32) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ClientGetListResponse) GetEnd() int32 {
	if o == nil || o.End == nil {
		var ret int32
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGetListResponse) GetEndOk() (*int32, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ClientGetListResponse) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given int32 and assigns it to the End field.
func (o *ClientGetListResponse) SetEnd(v int32) {
	o.End = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ClientGetListResponse) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGetListResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ClientGetListResponse) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ClientGetListResponse) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *ClientGetListResponse) GetClients() []Client {
	if o == nil || o.Clients == nil {
		var ret []Client
		return ret
	}
	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGetListResponse) GetClientsOk() ([]Client, bool) {
	if o == nil || o.Clients == nil {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *ClientGetListResponse) HasClients() bool {
	if o != nil && o.Clients != nil {
		return true
	}

	return false
}

// SetClients gets a reference to the given []Client and assigns it to the Clients field.
func (o *ClientGetListResponse) SetClients(v []Client) {
	o.Clients = v
}

func (o ClientGetListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Developer != nil {
		toSerialize["developer"] = o.Developer
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.Clients != nil {
		toSerialize["clients"] = o.Clients
	}
	return json.Marshal(toSerialize)
}

type NullableClientGetListResponse struct {
	value *ClientGetListResponse
	isSet bool
}

func (v NullableClientGetListResponse) Get() *ClientGetListResponse {
	return v.value
}

func (v *NullableClientGetListResponse) Set(val *ClientGetListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClientGetListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClientGetListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientGetListResponse(val *ClientGetListResponse) *NullableClientGetListResponse {
	return &NullableClientGetListResponse{value: val, isSet: true}
}

func (v NullableClientGetListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientGetListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

