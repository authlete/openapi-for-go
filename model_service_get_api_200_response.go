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

// ServiceGetApi200Response - struct for ServiceGetApi200Response
type ServiceGetApi200Response struct {
	Service *Service
	ServiceLimited *ServiceLimited
}

// ServiceAsServiceGetApi200Response is a convenience function that returns Service wrapped in ServiceGetApi200Response
func ServiceAsServiceGetApi200Response(v *Service) ServiceGetApi200Response {
	return ServiceGetApi200Response{
		Service: v,
	}
}

// ServiceLimitedAsServiceGetApi200Response is a convenience function that returns ServiceLimited wrapped in ServiceGetApi200Response
func ServiceLimitedAsServiceGetApi200Response(v *ServiceLimited) ServiceGetApi200Response {
	return ServiceGetApi200Response{
		ServiceLimited: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ServiceGetApi200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Service
	err = json.Unmarshal(data, &dst.Service)
	if err == nil {
		jsonService, _ := json.Marshal(dst.Service)
		if string(jsonService) == "{}" { // empty struct
			dst.Service = nil
		} else {
			match++
		}
	} else {
		dst.Service = nil
	}

	// try to unmarshal data into ServiceLimited
	err = json.Unmarshal(data, &dst.ServiceLimited)
	if err == nil {
		jsonServiceLimited, _ := json.Marshal(dst.ServiceLimited)
		if string(jsonServiceLimited) == "{}" { // empty struct
			dst.ServiceLimited = nil
		} else {
			match++
		}
	} else {
		dst.ServiceLimited = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Service = nil
		dst.ServiceLimited = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ServiceGetApi200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ServiceGetApi200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ServiceGetApi200Response) MarshalJSON() ([]byte, error) {
	if src.Service != nil {
		return json.Marshal(&src.Service)
	}

	if src.ServiceLimited != nil {
		return json.Marshal(&src.ServiceLimited)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ServiceGetApi200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Service != nil {
		return obj.Service
	}

	if obj.ServiceLimited != nil {
		return obj.ServiceLimited
	}

	// all schemas are nil
	return nil
}

type NullableServiceGetApi200Response struct {
	value *ServiceGetApi200Response
	isSet bool
}

func (v NullableServiceGetApi200Response) Get() *ServiceGetApi200Response {
	return v.value
}

func (v *NullableServiceGetApi200Response) Set(val *ServiceGetApi200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceGetApi200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceGetApi200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceGetApi200Response(val *ServiceGetApi200Response) *NullableServiceGetApi200Response {
	return &NullableServiceGetApi200Response{value: val, isSet: true}
}

func (v NullableServiceGetApi200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceGetApi200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

