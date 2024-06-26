/*
Synadia Control Plane / Synadia Cloud

Testing AgentTokenAPIService

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

func Test_syncp_AgentTokenAPIService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AgentTokenAPIService DeleteAgentToken", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var tokenId string

		httpRes, err := apiClient.AgentTokenAPI.DeleteAgentToken(context.Background(), tokenId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
