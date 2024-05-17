/*
Synadia Control Plane / Synadia Cloud

Testing ObjectBucketAPIService

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

func Test_syncp_ObjectBucketAPIService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectBucketAPIService CreateObjPullConsumer", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var streamId string

		resp, httpRes, err := apiClient.ObjectBucketAPI.CreateObjPullConsumer(context.Background(), streamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ObjectBucketAPIService CreateObjPushConsumer", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var streamId string

		resp, httpRes, err := apiClient.ObjectBucketAPI.CreateObjPushConsumer(context.Background(), streamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ObjectBucketAPIService DeleteObjectBucket", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var streamId string

		httpRes, err := apiClient.ObjectBucketAPI.DeleteObjectBucket(context.Background(), streamId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ObjectBucketAPIService GetObjectBucket", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var streamId string

		resp, httpRes, err := apiClient.ObjectBucketAPI.GetObjectBucket(context.Background(), streamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ObjectBucketAPIService ListObjConsumers", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var streamId string

		resp, httpRes, err := apiClient.ObjectBucketAPI.ListObjConsumers(context.Background(), streamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test ObjectBucketAPIService UpdateObjectBucket", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var streamId string

		resp, httpRes, err := apiClient.ObjectBucketAPI.UpdateObjectBucket(context.Background(), streamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
