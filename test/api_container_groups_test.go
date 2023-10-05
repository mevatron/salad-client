/*
SaladCloud Public API

Testing ContainerGroupsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_ContainerGroupsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContainerGroupsAPIService ContainerGroupInstanceReallocate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string
		var machineId string

		httpRes, err := apiClient.ContainerGroupsAPI.ContainerGroupInstanceReallocate(context.Background(), organizationName, projectName, containerGroupName, machineId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService ContainerGroupInstanceRecreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string
		var machineId string

		httpRes, err := apiClient.ContainerGroupsAPI.ContainerGroupInstanceRecreate(context.Background(), organizationName, projectName, containerGroupName, machineId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService ContainerGroupInstanceRestart", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string
		var machineId string

		httpRes, err := apiClient.ContainerGroupsAPI.ContainerGroupInstanceRestart(context.Background(), organizationName, projectName, containerGroupName, machineId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService CreateContainerGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string

		resp, httpRes, err := apiClient.ContainerGroupsAPI.CreateContainerGroup(context.Background(), organizationName, projectName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService DeleteContainerGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string

		httpRes, err := apiClient.ContainerGroupsAPI.DeleteContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService GetContainerGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string

		resp, httpRes, err := apiClient.ContainerGroupsAPI.GetContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService ListContainerGroupInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string

		resp, httpRes, err := apiClient.ContainerGroupsAPI.ListContainerGroupInstances(context.Background(), organizationName, projectName, containerGroupName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService ListContainerGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string

		resp, httpRes, err := apiClient.ContainerGroupsAPI.ListContainerGroups(context.Background(), organizationName, projectName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService RestartContainerGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string

		httpRes, err := apiClient.ContainerGroupsAPI.RestartContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService StartContainerGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string

		httpRes, err := apiClient.ContainerGroupsAPI.StartContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService StopContainerGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string

		httpRes, err := apiClient.ContainerGroupsAPI.StopContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContainerGroupsAPIService UpdateContainerGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organizationName string
		var projectName string
		var containerGroupName string

		resp, httpRes, err := apiClient.ContainerGroupsAPI.UpdateContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
