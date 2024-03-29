/*
Authlete API

Testing HskOperationsApiService

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

func Test_authlete_HskOperationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HskOperationsApiService HskCreateApi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.HskOperationsApi.HskCreateApi(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HskOperationsApiService HskDeleteApi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var handle string

		resp, httpRes, err := apiClient.HskOperationsApi.HskDeleteApi(context.Background(), handle).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HskOperationsApiService HskGetApi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var handle string

		resp, httpRes, err := apiClient.HskOperationsApi.HskGetApi(context.Background(), handle).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HskOperationsApiService HskGetListApi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.HskOperationsApi.HskGetListApi(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
