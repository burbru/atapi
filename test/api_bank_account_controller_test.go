/*
Api Documentation

Testing BankAccountControllerAPIService

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

func Test_atapi_BankAccountControllerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BankAccountControllerAPIService GetBankAccountUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bankAccountId string

		resp, httpRes, err := apiClient.BankAccountControllerAPI.GetBankAccountUsingGET(context.Background(), bankAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BankAccountControllerAPIService GetMyBankAccountUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BankAccountControllerAPI.GetMyBankAccountUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BankAccountControllerAPIService GetMyBankAccountV1UsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BankAccountControllerAPI.GetMyBankAccountV1UsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BankAccountControllerAPIService SendMoneyUsingPUT", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var senderBankAccountId string

		resp, httpRes, err := apiClient.BankAccountControllerAPI.SendMoneyUsingPUT(context.Background(), senderBankAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
