/*
Synadia Control Plane

Testing SigKeyAPIService

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

func Test_api_SigKeyAPIService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SigKeyAPIService DeleteAccountSk", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var keyId string

		httpRes, err := apiClient.SigKeyAPI.DeleteAccountSk(context.Background(), keyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SigKeyAPIService GetAccountSk", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var keyId string

		resp, httpRes, err := apiClient.SigKeyAPI.GetAccountSk(context.Background(), keyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SigKeyAPIService UpdateAccountSk", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var keyId string

		resp, httpRes, err := apiClient.SigKeyAPI.UpdateAccountSk(context.Background(), keyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
