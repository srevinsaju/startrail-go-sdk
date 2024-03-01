/*
startrail-service

Testing ServiceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_ServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServiceAPIService Activity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenant string
		var environment string
		var name string

		resp, httpRes, err := apiClient.ServiceAPI.Activity(context.Background(), tenant, environment, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAPIService Create", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ServiceAPI.Create(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAPIService Delete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenant string
		var environment string
		var name string

		resp, httpRes, err := apiClient.ServiceAPI.Delete(context.Background(), tenant, environment, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAPIService Get", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenant string
		var environment string
		var name string

		resp, httpRes, err := apiClient.ServiceAPI.Get(context.Background(), tenant, environment, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAPIService History", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenant string
		var environment string
		var name string

		resp, httpRes, err := apiClient.ServiceAPI.History(context.Background(), tenant, environment, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAPIService List", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tenant string

		resp, httpRes, err := apiClient.ServiceAPI.List(context.Background(), tenant).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
