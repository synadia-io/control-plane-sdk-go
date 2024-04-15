/*
Synadia Control Plane

Testing TeamServiceAccountAPIService

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

func Test_syncp_TeamServiceAccountAPIService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TeamServiceAccountAPIService CreateTeamServiceAccountToken", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var serviceAccountId string

		resp, httpRes, err := apiClient.TeamServiceAccountAPI.CreateTeamServiceAccountToken(context.Background(), serviceAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamServiceAccountAPIService DeleteTeamServiceAccount", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var serviceAccountId string

		httpRes, err := apiClient.TeamServiceAccountAPI.DeleteTeamServiceAccount(context.Background(), serviceAccountId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamServiceAccountAPIService GetTeamServiceAccount", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var serviceAccountId string

		resp, httpRes, err := apiClient.TeamServiceAccountAPI.GetTeamServiceAccount(context.Background(), serviceAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamServiceAccountAPIService ListTeamServiceAccountTokens", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var serviceAccountId string

		resp, httpRes, err := apiClient.TeamServiceAccountAPI.ListTeamServiceAccountTokens(context.Background(), serviceAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamServiceAccountAPIService UpdateTeamServiceAccount", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var serviceAccountId string

		resp, httpRes, err := apiClient.TeamServiceAccountAPI.UpdateTeamServiceAccount(context.Background(), serviceAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
