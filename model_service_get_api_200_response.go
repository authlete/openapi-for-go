/*
Authlete API Explorer

<div class=\"min-h-screen bg-gray-100 dark:bg-gray-900 text-gray-900 dark:text-gray-100 p-6\">   <div class=\"flex justify-end mb-4\">     <label for=\"theme-toggle\" class=\"flex items-center cursor-pointer\">       <div class=\"relative\">Dark mode:         <input type=\"checkbox\" id=\"theme-toggle\" class=\"sr-only\" onchange=\"toggleTheme()\">         <div class=\"block bg-gray-600 w-14 h-8 rounded-full\"></div>         <div class=\"dot absolute left-1 top-1 bg-white w-6 h-6 rounded-full transition\"></div>       </div>     </label>   </div>   <header class=\"bg-green-500 dark:bg-green-700 p-4 rounded-lg text-white text-center\">     <p>       Welcome to the <strong>Authlete API documentation</strong>. Authlete is an <strong>API-first service</strong>       where every aspect of the platform is configurable via API. This explorer provides a convenient way to       authenticate and interact with the API, allowing you to see Authlete in action quickly. 🚀     </p>     <p>       At a high level, the Authlete API is grouped into two categories:     </p>     <ul class=\"list-disc list-inside\">       <li><strong>Management APIs</strong>: Enable you to manage services and clients. 🔧</li>       <li><strong>Runtime APIs</strong>: Allow you to build your own Authorization Servers or Verifiable Credential (VC)         issuers. 🔐</li>     </ul>     <p>All API endpoints are secured using access tokens issued by Authlete's Identity Provider (IdP). If you already       have an Authlete account, simply use the <em>Get Token</em> option on the Authentication page to log in and obtain       an access token for API usage. If you don't have an account yet, <a href=\"https://console.authlete.com/register\">sign up         here</a> to get started.</p>   </header>   <main>     <section id=\"api-servers\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🌐 API Servers</h2>       <p>Authlete is a global service with clusters available in multiple regions across the world.</p>       <p>Currently, our service is available in the following regions:</p>       <div class=\"grid grid-cols-2 gap-4\">         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇺🇸 US</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇯🇵 JP</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇪🇺 EU</p>         </div>         <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">           <p class=\"text-center font-semibold\">🇧🇷 Brazil</p>         </div>       </div>       <p>Our customers can host their data in the region that best meets their requirements.</p>       <a href=\"#servers\" class=\"block mt-4 text-green-500 dark:text-green-300 hover:underline text-center\">Select your         preferred server</a>     </section>     <section id=\"authentication\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🔑 Authentication</h2>       <p>The API Explorer requires an access token to call the API.</p>       <p>You can create the access token from the <a href=\"https://console.authlete.com\">Authlete Management Console</a> and set it in the HTTP Bearer section of Authentication page.</p>       <p>Alternatively, if you have an Authlete account, the API Explorer can log you in with your Authlete account and         automatically acquire the required access token.</p>       <div class=\"theme-admonition theme-admonition-warning admonition_o5H7 alert alert--warning\">         <div class=\"admonitionContent_Knsx\">           <p>⚠️ <strong>Important Note:</strong> When the API Explorer acquires the token after login, the access tokens             will have the same permissions as the user who logs in as part of this flow.</p>         </div>       </div>       <a href=\"#auth\" class=\"block mt-4 text-green-500 dark:text-green-300 hover:underline text-center\">Setup your         access token</a>     </section>     <section id=\"tutorials\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🎓 Tutorials</h2>       <p>If you have successfully tested the API from the API Console and want to take the next step of integrating the         API into your application, or if you want to see a sample using Authlete APIs, follow the links below. These         resources will help you understand key concepts and how to integrate Authlete API into your applications.</p>       <div class=\"mt-4\">         <a href=\"https://www.authlete.com/developers/getting_started/\"           class=\"block text-green-500 dark:text-green-300 font-bold hover:underline mb-2\">🚀 Getting Started with           Authlete</a>           </br>         <a href=\"https://www.authlete.com/developers/tutorial/signup/\"           class=\"block text-green-500 dark:text-green-300 font-bold hover:underline\">🔑 From Sign-Up to the First API           Request</a>       </div>     </section>     <section id=\"support\" class=\"mb-10\">       <h2 class=\"text-2xl font-semibold mb-4\">🛠 Contact Us</h2>       <p>If you have any questions or need assistance, our team is here to help.</p>       <a href=\"https://www.authlete.com/contact/\"         class=\"block mt-4 text-green-500 dark:text-green-300 font-bold hover:underline\">Contact Page</a>     </section>   </main> </div>

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
	Service        *Service
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
func (obj *ServiceGetApi200Response) GetActualInstance() interface{} {
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
