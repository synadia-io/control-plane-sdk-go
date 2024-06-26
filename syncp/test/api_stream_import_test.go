/*
Synadia Control Plane / Synadia Cloud

Testing StreamImportAPIService

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

func Test_syncp_StreamImportAPIService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StreamImportAPIService DeleteStreamImport", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var streamImportId string

		httpRes, err := apiClient.StreamImportAPI.DeleteStreamImport(context.Background(), streamImportId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test StreamImportAPIService GetStreamImport", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var streamImportId string

		resp, httpRes, err := apiClient.StreamImportAPI.GetStreamImport(context.Background(), streamImportId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
