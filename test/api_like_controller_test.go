/*
Api Documentation

Testing LikeControllerAPIService

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

func Test_atapi_LikeControllerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LikeControllerAPIService DeleteLikeUsingDELETE", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var postId string

		resp, httpRes, err := apiClient.LikeControllerAPI.DeleteLikeUsingDELETE(context.Background(), postId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LikeControllerAPIService LikePostUsingPUT", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var postId string

		resp, httpRes, err := apiClient.LikeControllerAPI.LikePostUsingPUT(context.Background(), postId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LikeControllerAPIService ListLikesUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var postId string

		resp, httpRes, err := apiClient.LikeControllerAPI.ListLikesUsingGET(context.Background(), postId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
