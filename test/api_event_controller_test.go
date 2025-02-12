/*
Api Documentation

Testing EventControllerAPIService

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

func Test_atapi_EventControllerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EventControllerAPIService FindEventsByEventTypeUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var eventType string

		resp, httpRes, err := apiClient.EventControllerAPI.FindEventsByEventTypeUsingGET(context.Background(), eventType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventControllerAPIService FindEventsByFullTextUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fullTextPart string

		resp, httpRes, err := apiClient.EventControllerAPI.FindEventsByFullTextUsingGET(context.Background(), fullTextPart).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventControllerAPIService FindEventsByRealmsUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EventControllerAPI.FindEventsByRealmsUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventControllerAPIService GetAllUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EventControllerAPI.GetAllUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventControllerAPIService GetMyEventsUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EventControllerAPI.GetMyEventsUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
