/*
Api Documentation

Testing HistoryControllerAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package atapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_atapi_HistoryControllerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HistoryControllerAPIService GetHistoryEntriesUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.HistoryControllerAPI.GetHistoryEntriesUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
