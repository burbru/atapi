/*
Api Documentation

Testing CorporateActionControllerAPIService

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

func Test_atapi_CorporateActionControllerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CorporateActionControllerAPIService GeMergersUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CorporateActionControllerAPI.GeMergersUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorporateActionControllerAPIService GetCapitalIncreasesUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CorporateActionControllerAPI.GetCapitalIncreasesUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorporateActionControllerAPIService GetCapitalReductionsUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CorporateActionControllerAPI.GetCapitalReductionsUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorporateActionControllerAPIService GetCompanyCapitalIncreaseUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var companyId string

		resp, httpRes, err := apiClient.CorporateActionControllerAPI.GetCompanyCapitalIncreaseUsingGET(context.Background(), companyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorporateActionControllerAPIService GetCompanyCapitalReductionUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var companyId string

		resp, httpRes, err := apiClient.CorporateActionControllerAPI.GetCompanyCapitalReductionUsingGET(context.Background(), companyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorporateActionControllerAPIService GetCompanyDividendPaymentUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var companyId string

		resp, httpRes, err := apiClient.CorporateActionControllerAPI.GetCompanyDividendPaymentUsingGET(context.Background(), companyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorporateActionControllerAPIService GetDividendPaymentsUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CorporateActionControllerAPI.GetDividendPaymentsUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CorporateActionControllerAPIService GetMergerUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var companyId string

		resp, httpRes, err := apiClient.CorporateActionControllerAPI.GetMergerUsingGET(context.Background(), companyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
