/*
Authlete API

Authlete API Document. 

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
	"fmt"
)

// ClientGetListApi200Response - struct for ClientGetListApi200Response
type ClientGetListApi200Response struct {
	ClientGetListLimitedResponse *ClientGetListLimitedResponse
	ClientGetListResponse *ClientGetListResponse
}

// ClientGetListLimitedResponseAsClientGetListApi200Response is a convenience function that returns ClientGetListLimitedResponse wrapped in ClientGetListApi200Response
func ClientGetListLimitedResponseAsClientGetListApi200Response(v *ClientGetListLimitedResponse) ClientGetListApi200Response {
	return ClientGetListApi200Response{
		ClientGetListLimitedResponse: v,
	}
}

// ClientGetListResponseAsClientGetListApi200Response is a convenience function that returns ClientGetListResponse wrapped in ClientGetListApi200Response
func ClientGetListResponseAsClientGetListApi200Response(v *ClientGetListResponse) ClientGetListApi200Response {
	return ClientGetListApi200Response{
		ClientGetListResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ClientGetListApi200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ClientGetListLimitedResponse
	err = json.Unmarshal(data, &dst.ClientGetListLimitedResponse)
	if err == nil {
		jsonClientGetListLimitedResponse, _ := json.Marshal(dst.ClientGetListLimitedResponse)
		if string(jsonClientGetListLimitedResponse) == "{}" { // empty struct
			dst.ClientGetListLimitedResponse = nil
		} else {
			match++
		}
	} else {
		dst.ClientGetListLimitedResponse = nil
	}

	// try to unmarshal data into ClientGetListResponse
	err = json.Unmarshal(data, &dst.ClientGetListResponse)
	if err == nil {
		jsonClientGetListResponse, _ := json.Marshal(dst.ClientGetListResponse)
		if string(jsonClientGetListResponse) == "{}" { // empty struct
			dst.ClientGetListResponse = nil
		} else {
			match++
		}
	} else {
		dst.ClientGetListResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ClientGetListLimitedResponse = nil
		dst.ClientGetListResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ClientGetListApi200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ClientGetListApi200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ClientGetListApi200Response) MarshalJSON() ([]byte, error) {
	if src.ClientGetListLimitedResponse != nil {
		return json.Marshal(&src.ClientGetListLimitedResponse)
	}

	if src.ClientGetListResponse != nil {
		return json.Marshal(&src.ClientGetListResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ClientGetListApi200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ClientGetListLimitedResponse != nil {
		return obj.ClientGetListLimitedResponse
	}

	if obj.ClientGetListResponse != nil {
		return obj.ClientGetListResponse
	}

	// all schemas are nil
	return nil
}

type NullableClientGetListApi200Response struct {
	value *ClientGetListApi200Response
	isSet bool
}

func (v NullableClientGetListApi200Response) Get() *ClientGetListApi200Response {
	return v.value
}

func (v *NullableClientGetListApi200Response) Set(val *ClientGetListApi200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableClientGetListApi200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableClientGetListApi200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientGetListApi200Response(val *ClientGetListApi200Response) *NullableClientGetListApi200Response {
	return &NullableClientGetListApi200Response{value: val, isSet: true}
}

func (v NullableClientGetListApi200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientGetListApi200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


