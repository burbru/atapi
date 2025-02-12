/*
Api Documentation

Testing BankingLicenseControllerAPIService

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

func Test_atapi_BankingLicenseControllerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BankingLicenseControllerAPIService CreateBankingLicenseUsingPOST", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BankingLicenseControllerAPI.CreateBankingLicenseUsingPOST(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BankingLicenseControllerAPIService GetBankingLicenseUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bankingLicenseId string

		resp, httpRes, err := apiClient.BankingLicenseControllerAPI.GetBankingLicenseUsingGET(context.Background(), bankingLicenseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BankingLicenseControllerAPIService GetCompanyBankingLicenseUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BankingLicenseControllerAPI.GetCompanyBankingLicenseUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
