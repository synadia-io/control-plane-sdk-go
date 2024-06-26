/*
Synadia Control Plane / Synadia Cloud

Testing SigKeyGroupAPIService

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

func Test_syncp_SigKeyGroupAPIService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SigKeyGroupAPIService CopyAccountSkGroup", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.SigKeyGroupAPI.CopyAccountSkGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SigKeyGroupAPIService DeleteAccountSkGroup", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var groupId string

		httpRes, err := apiClient.SigKeyGroupAPI.DeleteAccountSkGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SigKeyGroupAPIService GetAccountSkGroup", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.SigKeyGroupAPI.GetAccountSkGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SigKeyGroupAPIService ListAccountSkGroupKeys", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.SigKeyGroupAPI.ListAccountSkGroupKeys(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SigKeyGroupAPIService RotateAccountSk", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.SigKeyGroupAPI.RotateAccountSk(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SigKeyGroupAPIService UpdateAccountSkGroup", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.SigKeyGroupAPI.UpdateAccountSkGroup(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
