/*
Synadia Control Plane

Testing TeamServiceAccountTokensAPIService

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

func Test_syncp_TeamServiceAccountTokensAPIService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TeamServiceAccountTokensAPIService DeleteTeamServiceAccountToken", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var tokenId string

		httpRes, err := apiClient.TeamServiceAccountTokensAPI.DeleteTeamServiceAccountToken(context.Background(), tokenId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamServiceAccountTokensAPIService GetTeamServiceAccountToken", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var tokenId string

		resp, httpRes, err := apiClient.TeamServiceAccountTokensAPI.GetTeamServiceAccountToken(context.Background(), tokenId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamServiceAccountTokensAPIService UpdateTeamServiceAccountToken", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var tokenId string

		resp, httpRes, err := apiClient.TeamServiceAccountTokensAPI.UpdateTeamServiceAccountToken(context.Background(), tokenId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
