/*
Api Documentation

Testing PostControllerAPIService

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

func Test_atapi_PostControllerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PostControllerAPIService CreateCommentUsingPOST", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var postId string

		resp, httpRes, err := apiClient.PostControllerAPI.CreateCommentUsingPOST(context.Background(), postId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostControllerAPIService CreatePostUsingPOST", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PostControllerAPI.CreatePostUsingPOST(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostControllerAPIService DeletePostUsingDELETE", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var postId string

		resp, httpRes, err := apiClient.PostControllerAPI.DeletePostUsingDELETE(context.Background(), postId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostControllerAPIService EditPostUsingPUT", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var postId string

		resp, httpRes, err := apiClient.PostControllerAPI.EditPostUsingPUT(context.Background(), postId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostControllerAPIService GetListingDescriptionUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PostControllerAPI.GetListingDescriptionUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostControllerAPIService GetPersonalInterestInPostUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var postId string

		resp, httpRes, err := apiClient.PostControllerAPI.GetPersonalInterestInPostUsingGET(context.Background(), postId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostControllerAPIService GetPostUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var postId string

		resp, httpRes, err := apiClient.PostControllerAPI.GetPostUsingGET(context.Background(), postId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostControllerAPIService ListCommentsUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var postId string

		resp, httpRes, err := apiClient.PostControllerAPI.ListCommentsUsingGET(context.Background(), postId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostControllerAPIService ListListingPostsUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PostControllerAPI.ListListingPostsUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
