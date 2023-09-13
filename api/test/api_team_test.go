/*
Synadia Control Plane

Testing TeamAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	openapiclient "github.com/synadia-io/control-plane-sdk-go/api"
)

func Test_api_TeamAPIService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TeamAPIService AssignTeamAppUser", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var teamId string
		var appUserId string

		resp, httpRes, err := apiClient.TeamAPI.AssignTeamAppUser(context.Background(), teamId, appUserId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamAPIService CreateSystem", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var teamId string

		resp, httpRes, err := apiClient.TeamAPI.CreateSystem(context.Background(), teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamAPIService DeleteTeam", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var teamId string

		httpRes, err := apiClient.TeamAPI.DeleteTeam(context.Background(), teamId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamAPIService GetTeam", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var teamId string

		resp, httpRes, err := apiClient.TeamAPI.GetTeam(context.Background(), teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamAPIService ImportSystem", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var teamId string

		resp, httpRes, err := apiClient.TeamAPI.ImportSystem(context.Background(), teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamAPIService ListSystems", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var teamId string

		resp, httpRes, err := apiClient.TeamAPI.ListSystems(context.Background(), teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamAPIService ListTeamAccounts", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var teamId string

		resp, httpRes, err := apiClient.TeamAPI.ListTeamAccounts(context.Background(), teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamAPIService ListTeamAppUsers", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var teamId string

		resp, httpRes, err := apiClient.TeamAPI.ListTeamAppUsers(context.Background(), teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamAPIService ListTeamNatsUsers", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var teamId string

		resp, httpRes, err := apiClient.TeamAPI.ListTeamNatsUsers(context.Background(), teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamAPIService UnAssignTeamAppUser", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var teamId string
		var appUserId string

		httpRes, err := apiClient.TeamAPI.UnAssignTeamAppUser(context.Background(), teamId, appUserId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TeamAPIService UpdateTeam", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var teamId string

		resp, httpRes, err := apiClient.TeamAPI.UpdateTeam(context.Background(), teamId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
