/*
Api Documentation

Testing ListingControllerAPIService

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

func Test_atapi_ListingControllerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ListingControllerAPIService FindListingsByAsinPartUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var securityIdentifierPart string

		resp, httpRes, err := apiClient.ListingControllerAPI.FindListingsByAsinPartUsingGET(context.Background(), securityIdentifierPart).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListingControllerAPIService FindListingsByAsinUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ListingControllerAPI.FindListingsByAsinUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListingControllerAPIService GetListingProfileUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ListingControllerAPI.GetListingProfileUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListingControllerAPIService GetListingUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var securityIdentifier string

		resp, httpRes, err := apiClient.ListingControllerAPI.GetListingUsingGET(context.Background(), securityIdentifier).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListingControllerAPIService GetOutstandingSharesUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var securityIdentifier string

		resp, httpRes, err := apiClient.ListingControllerAPI.GetOutstandingSharesUsingGET(context.Background(), securityIdentifier).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListingControllerAPIService GetProfileV1UsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var securityIdentifier string

		resp, httpRes, err := apiClient.ListingControllerAPI.GetProfileV1UsingGET(context.Background(), securityIdentifier).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListingControllerAPIService ListHistorizedCompanyDataUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var securityIdentifier string

		resp, httpRes, err := apiClient.ListingControllerAPI.ListHistorizedCompanyDataUsingGET(context.Background(), securityIdentifier).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListingControllerAPIService ListHistorizedListingDataUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var securityIdentifier string

		resp, httpRes, err := apiClient.ListingControllerAPI.ListHistorizedListingDataUsingGET(context.Background(), securityIdentifier).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListingControllerAPIService ListListingsUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ListingControllerAPI.ListListingsUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListingControllerAPIService ListListingsV1UsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ListingControllerAPI.ListListingsV1UsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListingControllerAPIService SearchListingsUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var securityIdentifier string

		resp, httpRes, err := apiClient.ListingControllerAPI.SearchListingsUsingGET(context.Background(), securityIdentifier).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
