/*
SaladCloud Public API

Testing RecipeDeploymentsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package saladclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_saladclient_RecipeDeploymentsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RecipeDeploymentsAPIService CreateRecipeDeployment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string

		resp, httpRes, err := apiClient.RecipeDeploymentsAPI.CreateRecipeDeployment(context.Background(), organizationName, projectName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RecipeDeploymentsAPIService DeleteRecipeDeployment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var recipeDeploymentName string

		httpRes, err := apiClient.RecipeDeploymentsAPI.DeleteRecipeDeployment(context.Background(), organizationName, projectName, recipeDeploymentName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RecipeDeploymentsAPIService GetRecipeDeployment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var recipeDeploymentName string

		resp, httpRes, err := apiClient.RecipeDeploymentsAPI.GetRecipeDeployment(context.Background(), organizationName, projectName, recipeDeploymentName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RecipeDeploymentsAPIService ListRecipeDeploymentInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var recipeDeploymentName string

		resp, httpRes, err := apiClient.RecipeDeploymentsAPI.ListRecipeDeploymentInstances(context.Background(), organizationName, projectName, recipeDeploymentName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RecipeDeploymentsAPIService ListRecipeDeployments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string

		resp, httpRes, err := apiClient.RecipeDeploymentsAPI.ListRecipeDeployments(context.Background(), organizationName, projectName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RecipeDeploymentsAPIService RestartDeployedRecipe", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var recipeDeploymentName string

		httpRes, err := apiClient.RecipeDeploymentsAPI.RestartDeployedRecipe(context.Background(), organizationName, projectName, recipeDeploymentName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RecipeDeploymentsAPIService StartDeployedRecipe", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var recipeDeploymentName string

		httpRes, err := apiClient.RecipeDeploymentsAPI.StartDeployedRecipe(context.Background(), organizationName, projectName, recipeDeploymentName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RecipeDeploymentsAPIService StopDeployedRecipe", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var recipeDeploymentName string

		httpRes, err := apiClient.RecipeDeploymentsAPI.StopDeployedRecipe(context.Background(), organizationName, projectName, recipeDeploymentName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RecipeDeploymentsAPIService UpdateRecipeDeployment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var recipeDeploymentName string

		resp, httpRes, err := apiClient.RecipeDeploymentsAPI.UpdateRecipeDeployment(context.Background(), organizationName, projectName, recipeDeploymentName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
