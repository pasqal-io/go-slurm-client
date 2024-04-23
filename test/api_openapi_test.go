/*
Slurm Rest API

Testing OpenapiAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package slurmclient

import (
	"context"
	"testing"

	openapiclient "github.com/pasqal-io/go-slurm-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_slurmclient_OpenapiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OpenapiAPIService OpenapiGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.OpenapiAPI.OpenapiGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OpenapiAPIService OpenapiJsonGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.OpenapiAPI.OpenapiJsonGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OpenapiAPIService OpenapiV3Get", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.OpenapiAPI.OpenapiV3Get(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OpenapiAPIService OpenapiYamlGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.OpenapiAPI.OpenapiYamlGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
