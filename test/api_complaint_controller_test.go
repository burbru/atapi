/*
Api Documentation

Testing ComplaintControllerAPIService

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

func Test_atapi_ComplaintControllerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ComplaintControllerAPIService CreateComplaintUsingPOST", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ComplaintControllerAPI.CreateComplaintUsingPOST(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ComplaintControllerAPIService GetComplaintUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var complaintId string

		resp, httpRes, err := apiClient.ComplaintControllerAPI.GetComplaintUsingGET(context.Background(), complaintId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
