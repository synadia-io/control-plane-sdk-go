/*
Synadia Control Plane

Testing SystemAPIService

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

func Test_api_SystemAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SystemAPIService AssignSystemAppUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string
		var appUserId string

		resp, httpRes, err := apiClient.SystemAPI.AssignSystemAppUser(context.Background(), systemId, appUserId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService CreateAccount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemAPI.CreateAccount(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService CreateSystemAlertRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemAPI.CreateSystemAlertRule(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService DeleteSystem", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string

		httpRes, err := apiClient.SystemAPI.DeleteSystem(context.Background(), systemId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService DeleteSystemAlertRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string
		var alertRuleId string

		httpRes, err := apiClient.SystemAPI.DeleteSystemAlertRule(context.Background(), systemId, alertRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService GetCurrentAgentToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemAPI.GetCurrentAgentToken(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService GetSystem", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemAPI.GetSystem(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService GetSystemAlertRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string
		var alertRuleId string

		resp, httpRes, err := apiClient.SystemAPI.GetSystemAlertRule(context.Background(), systemId, alertRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService ImportAccount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string

		httpRes, err := apiClient.SystemAPI.ImportAccount(context.Background(), systemId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService ImportUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemAPI.ImportUser(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService ListAccounts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemAPI.ListAccounts(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService ListAgentTokens", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemAPI.ListAgentTokens(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService ListClusters", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemAPI.ListClusters(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService ListConnections", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemAPI.ListConnections(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService ListServers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemAPI.ListServers(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService ListSystemAlertRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemAPI.ListSystemAlertRules(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService ListSystemAppUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemAPI.ListSystemAppUsers(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService RotateAgentToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemAPI.RotateAgentToken(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService RunSystemAlertRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string
		var alertRuleId string

		resp, httpRes, err := apiClient.SystemAPI.RunSystemAlertRule(context.Background(), systemId, alertRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService UnAssignSystemAppUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string
		var appUserId string

		httpRes, err := apiClient.SystemAPI.UnAssignSystemAppUser(context.Background(), systemId, appUserId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService UpdateSystem", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string

		resp, httpRes, err := apiClient.SystemAPI.UpdateSystem(context.Background(), systemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SystemAPIService UpdateSystemAlertRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var systemId string
		var alertRuleId string

		resp, httpRes, err := apiClient.SystemAPI.UpdateSystemAlertRule(context.Background(), systemId, alertRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
