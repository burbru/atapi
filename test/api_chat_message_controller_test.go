/*
Api Documentation

Testing ChatMessageControllerAPIService

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

func Test_atapi_ChatMessageControllerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ChatMessageControllerAPIService DeleteMessageUsingDELETE", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var messageId string

		resp, httpRes, err := apiClient.ChatMessageControllerAPI.DeleteMessageUsingDELETE(context.Background(), messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatMessageControllerAPIService DeleteMessageV1UsingDELETE", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var messageId string

		resp, httpRes, err := apiClient.ChatMessageControllerAPI.DeleteMessageV1UsingDELETE(context.Background(), messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatMessageControllerAPIService GetMessageUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var messageId string

		resp, httpRes, err := apiClient.ChatMessageControllerAPI.GetMessageUsingGET(context.Background(), messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatMessageControllerAPIService GetMessageV1UsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var messageId string

		resp, httpRes, err := apiClient.ChatMessageControllerAPI.GetMessageV1UsingGET(context.Background(), messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatMessageControllerAPIService GetMessagesOfChatUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var chatId string

		resp, httpRes, err := apiClient.ChatMessageControllerAPI.GetMessagesOfChatUsingGET(context.Background(), chatId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatMessageControllerAPIService GetMessagesOfChatV1UsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var chatId string

		resp, httpRes, err := apiClient.ChatMessageControllerAPI.GetMessagesOfChatV1UsingGET(context.Background(), chatId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatMessageControllerAPIService GetUnreadMessagesUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatMessageControllerAPI.GetUnreadMessagesUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatMessageControllerAPIService GetUnreadMessagesV1UsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatMessageControllerAPI.GetUnreadMessagesV1UsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatMessageControllerAPIService SendMessageUsingPOST", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatMessageControllerAPI.SendMessageUsingPOST(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatMessageControllerAPIService SendMessageV1UsingPOST", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatMessageControllerAPI.SendMessageV1UsingPOST(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatMessageControllerAPIService UpdateMessageReadUsingPUT", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatMessageControllerAPI.UpdateMessageReadUsingPUT(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ChatMessageControllerAPIService UpdateMessageReadV1UsingPUT", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ChatMessageControllerAPI.UpdateMessageReadV1UsingPUT(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
