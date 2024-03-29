/*
Authlete API

Testing AuthorizationEndpointApiService

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

func Test_authlete_AuthorizationEndpointApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AuthorizationEndpointApiService AuthAuthorizationApi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AuthorizationEndpointApi.AuthAuthorizationApi(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthorizationEndpointApiService AuthAuthorizationFailApi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AuthorizationEndpointApi.AuthAuthorizationFailApi(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthorizationEndpointApiService AuthAuthorizationIssueApi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AuthorizationEndpointApi.AuthAuthorizationIssueApi(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
