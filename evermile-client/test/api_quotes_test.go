/*
Evermile Commercial Quotes and Booking API

Testing QuotesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package EvermileClient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_EvermileClient_QuotesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test QuotesApiService ProposalProposalIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var proposalId string

		resp, httpRes, err := apiClient.QuotesApi.ProposalProposalIdGet(context.Background(), proposalId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QuotesApiService QuotePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.QuotesApi.QuotePost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}