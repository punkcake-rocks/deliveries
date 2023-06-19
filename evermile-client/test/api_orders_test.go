/*
Evermile Commercial Quotes and Booking API

Testing OrdersApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package EvermileClient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_EvermileClient_OrdersApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrdersApiService OrderOrderIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orderId string

		httpRes, err := apiClient.OrdersApi.OrderOrderIdDelete(context.Background(), orderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersApiService OrderOrderIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orderId string

		resp, httpRes, err := apiClient.OrdersApi.OrderOrderIdGet(context.Background(), orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersApiService OrderOrderIdLabelGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orderId string

		resp, httpRes, err := apiClient.OrdersApi.OrderOrderIdLabelGet(context.Background(), orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersApiService OrderOrderIdLiveTrackingGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orderId string

		resp, httpRes, err := apiClient.OrdersApi.OrderOrderIdLiveTrackingGet(context.Background(), orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersApiService OrderOrderIdPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orderId string

		httpRes, err := apiClient.OrdersApi.OrderOrderIdPatch(context.Background(), orderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersApiService OrderPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrdersApi.OrderPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersApiService OrdersGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OrdersApi.OrdersGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}