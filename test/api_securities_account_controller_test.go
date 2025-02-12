/*
Api Documentation

Testing SecuritiesAccountControllerAPIService

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

func Test_atapi_SecuritiesAccountControllerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SecuritiesAccountControllerAPIService FindSecuritiesAccountDetailsUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SecuritiesAccountControllerAPI.FindSecuritiesAccountDetailsUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecuritiesAccountControllerAPIService GetSecuritiesAccountDetailsUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var securitiesAccountId string

		resp, httpRes, err := apiClient.SecuritiesAccountControllerAPI.GetSecuritiesAccountDetailsUsingGET(context.Background(), securitiesAccountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecuritiesAccountControllerAPIService GetSecuritiesAccountUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SecuritiesAccountControllerAPI.GetSecuritiesAccountUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
