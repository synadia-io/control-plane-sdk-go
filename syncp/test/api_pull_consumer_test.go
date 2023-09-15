/*
Synadia Control Plane

Testing PullConsumerAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package syncp

import (
	"context"
	openapiclient "github.com/synadia-io/control-plane-sdk-go/syncp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_syncp_PullConsumerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PullConsumerAPIService DeletePullConsumer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var consumerId string

		httpRes, err := apiClient.PullConsumerAPI.DeletePullConsumer(context.Background(), consumerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PullConsumerAPIService GetPullConsumerInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var consumerId string

		resp, httpRes, err := apiClient.PullConsumerAPI.GetPullConsumerInfo(context.Background(), consumerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PullConsumerAPIService UpdatePullConsumer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var consumerId string

		resp, httpRes, err := apiClient.PullConsumerAPI.UpdatePullConsumer(context.Background(), consumerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
