/*
Authlete API Explorer

Testing TokenOperationsAPIService

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

func Test_authlete_TokenOperationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TokenOperationsAPIService AuthTokenCreateApi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string

		resp, httpRes, err := apiClient.TokenOperationsAPI.AuthTokenCreateApi(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TokenOperationsAPIService AuthTokenDeleteApi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string
		var accessTokenIdentifier string

		httpRes, err := apiClient.TokenOperationsAPI.AuthTokenDeleteApi(context.Background(), serviceId, accessTokenIdentifier).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TokenOperationsAPIService AuthTokenGetListApi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string

		resp, httpRes, err := apiClient.TokenOperationsAPI.AuthTokenGetListApi(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TokenOperationsAPIService AuthTokenRevokeApi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string

		resp, httpRes, err := apiClient.TokenOperationsAPI.AuthTokenRevokeApi(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TokenOperationsAPIService AuthTokenUpdateApi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string

		resp, httpRes, err := apiClient.TokenOperationsAPI.AuthTokenUpdateApi(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
