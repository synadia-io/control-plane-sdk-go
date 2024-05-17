/*
Synadia Control Plane / Synadia Cloud

Testing SubjectImportAPIService

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

func Test_syncp_SubjectImportAPIService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SubjectImportAPIService DeleteSubjectImport", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var subjectImportId string

		httpRes, err := apiClient.SubjectImportAPI.DeleteSubjectImport(context.Background(), subjectImportId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SubjectImportAPIService GetSubjectImport", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var subjectImportId string

		resp, httpRes, err := apiClient.SubjectImportAPI.GetSubjectImport(context.Background(), subjectImportId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test SubjectImportAPIService UpdateSubjectImport", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var subjectImportId string

		resp, httpRes, err := apiClient.SubjectImportAPI.UpdateSubjectImport(context.Background(), subjectImportId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
