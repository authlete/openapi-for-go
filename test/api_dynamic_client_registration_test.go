/*
Authlete API Explorer

Testing DynamicClientRegistrationAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package authlete

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/authlete/openapi-for-go"
)

func Test_authlete_DynamicClientRegistrationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DynamicClientRegistrationAPIService ClientRegistrationApi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string

		resp, httpRes, err := apiClient.DynamicClientRegistrationAPI.ClientRegistrationApi(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DynamicClientRegistrationAPIService ClientRegistrationDeleteApi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string

		resp, httpRes, err := apiClient.DynamicClientRegistrationAPI.ClientRegistrationDeleteApi(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DynamicClientRegistrationAPIService ClientRegistrationGetApi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string

		resp, httpRes, err := apiClient.DynamicClientRegistrationAPI.ClientRegistrationGetApi(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DynamicClientRegistrationAPIService ClientRegistrationUpdateApi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string

		resp, httpRes, err := apiClient.DynamicClientRegistrationAPI.ClientRegistrationUpdateApi(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
