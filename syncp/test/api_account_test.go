/*
Synadia Control Plane

Testing AccountAPIService

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

func Test_syncp_AccountAPIService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccountAPIService AssignAccountTeamAppUser", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string
		var teamAppUserId string

		resp, httpRes, err := apiClient.AccountAPI.AssignAccountTeamAppUser(context.Background(), accountId, teamAppUserId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService CreateAccountSkGroup", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.CreateAccountSkGroup(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService CreateAlertRule", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.CreateAlertRule(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService CreateKvBucket", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.CreateKvBucket(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService CreateMirror", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.CreateMirror(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService CreateStream", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.CreateStream(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService CreateStreamExport", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.CreateStreamExport(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService CreateStreamImport", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.CreateStreamImport(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService CreateSubjectExport", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.CreateSubjectExport(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService CreateSubjectImport", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.CreateSubjectImport(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService CreateUser", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.CreateUser(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService DeleteAccount", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		httpRes, err := apiClient.AccountAPI.DeleteAccount(context.Background(), accountId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService DeleteAlertRule", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string
		var alertRuleId string

		httpRes, err := apiClient.AccountAPI.DeleteAlertRule(context.Background(), accountId, alertRuleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService GetAccount", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.GetAccount(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService GetAccountInfo", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.GetAccountInfo(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService GetAccountMetrics", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.GetAccountMetrics(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService GetAlertRule", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string
		var alertRuleId string

		resp, httpRes, err := apiClient.AccountAPI.GetAlertRule(context.Background(), accountId, alertRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService GetJetStreamPlacementOptions", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.GetJetStreamPlacementOptions(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService ListAccountConnections", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.ListAccountConnections(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService ListAccountSkGroup", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.ListAccountSkGroup(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService ListAccountTeamAppUsers", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.ListAccountTeamAppUsers(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService ListAlertRules", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.ListAlertRules(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService ListJetStreamAssets", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.ListJetStreamAssets(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService ListKvBuckets", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.ListKvBuckets(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService ListMirrors", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.ListMirrors(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService ListStreamExports", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.ListStreamExports(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService ListStreamExportsShared", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.ListStreamExportsShared(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService ListStreamImports", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.ListStreamImports(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService ListStreams", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.ListStreams(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService ListSubjectExports", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.ListSubjectExports(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService ListSubjectExportsShared", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.ListSubjectExportsShared(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService ListSubjectImports", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.ListSubjectImports(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService ListUsers", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.ListUsers(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService RunAlertRule", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string
		var alertRuleId string

		resp, httpRes, err := apiClient.AccountAPI.RunAlertRule(context.Background(), accountId, alertRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService UnAssignAccountTeamAppUser", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string
		var teamAppUserId string

		httpRes, err := apiClient.AccountAPI.UnAssignAccountTeamAppUser(context.Background(), accountId, teamAppUserId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService UnmanageAccount", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		httpRes, err := apiClient.AccountAPI.UnmanageAccount(context.Background(), accountId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService UpdateAccount", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string

		resp, httpRes, err := apiClient.AccountAPI.UpdateAccount(context.Background(), accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test AccountAPIService UpdateAlertRule", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var accountId string
		var alertRuleId string

		resp, httpRes, err := apiClient.AccountAPI.UpdateAlertRule(context.Background(), accountId, alertRuleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
