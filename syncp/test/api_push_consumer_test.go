/*
Synadia Control Plane / Synadia Cloud

Testing PushConsumerAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package syncp

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	openapiclient "github.com/synadia-io/control-plane-sdk-go/syncp"
)

func Test_syncp_PushConsumerAPIService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PushConsumerAPIService DeletePushConsumer", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var consumerId string

		httpRes, err := apiClient.PushConsumerAPI.DeletePushConsumer(context.Background(), consumerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PushConsumerAPIService GetPushConsumerInfo", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var consumerId string

		resp, httpRes, err := apiClient.PushConsumerAPI.GetPushConsumerInfo(context.Background(), consumerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test PushConsumerAPIService UpdatePushConsumer", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var consumerId string

		resp, httpRes, err := apiClient.PushConsumerAPI.UpdatePushConsumer(context.Background(), consumerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
