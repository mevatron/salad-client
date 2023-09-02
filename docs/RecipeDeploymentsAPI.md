# \RecipeDeploymentsAPI

All URIs are relative to *https://api.salad.com/api/public*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRecipeDeployment**](RecipeDeploymentsAPI.md#CreateRecipeDeployment) | **Post** /organizations/{organization_name}/projects/{project_name}/recipe-deployments | Create a Recipe Deployment
[**DeleteRecipeDeployment**](RecipeDeploymentsAPI.md#DeleteRecipeDeployment) | **Delete** /organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name} | Delete a Recipe Deployment
[**GetRecipeDeployment**](RecipeDeploymentsAPI.md#GetRecipeDeployment) | **Get** /organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name} | Get a Recipe Deployment
[**ListRecipeDeploymentInstances**](RecipeDeploymentsAPI.md#ListRecipeDeploymentInstances) | **Get** /organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}/instances | List Recipe Deployment Instances
[**ListRecipeDeployments**](RecipeDeploymentsAPI.md#ListRecipeDeployments) | **Get** /organizations/{organization_name}/projects/{project_name}/recipe-deployments | List Recipe Deployments
[**RestartDeployedRecipe**](RecipeDeploymentsAPI.md#RestartDeployedRecipe) | **Post** /organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}/restart | Restart a Deployed Recipe
[**StartDeployedRecipe**](RecipeDeploymentsAPI.md#StartDeployedRecipe) | **Post** /organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}/start | Start a Deployed Recipe
[**StopDeployedRecipe**](RecipeDeploymentsAPI.md#StopDeployedRecipe) | **Post** /organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name}/stop | Stop a Deployed Recipe
[**UpdateRecipeDeployment**](RecipeDeploymentsAPI.md#UpdateRecipeDeployment) | **Patch** /organizations/{organization_name}/projects/{project_name}/recipe-deployments/{recipe_deployment_name} | Update a Recipe Deployment



## CreateRecipeDeployment

> RecipeDeployment CreateRecipeDeployment(ctx, organizationName, projectName).CreateRecipeDeployment(createRecipeDeployment).Execute()

Create a Recipe Deployment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    organizationName := "organizationName_example" // string | The unique organization name
    projectName := "projectName_example" // string | The unique project name
    createRecipeDeployment := *openapiclient.NewCreateRecipeDeployment("Name_example", "DisplayName_example", int32(123), "RecipeName_example") // CreateRecipeDeployment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecipeDeploymentsAPI.CreateRecipeDeployment(context.Background(), organizationName, projectName).CreateRecipeDeployment(createRecipeDeployment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecipeDeploymentsAPI.CreateRecipeDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRecipeDeployment`: RecipeDeployment
    fmt.Fprintf(os.Stdout, "Response from `RecipeDeploymentsAPI.CreateRecipeDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | The unique organization name | 
**projectName** | **string** | The unique project name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRecipeDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createRecipeDeployment** | [**CreateRecipeDeployment**](CreateRecipeDeployment.md) |  | 

### Return type

[**RecipeDeployment**](RecipeDeployment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecipeDeployment

> DeleteRecipeDeployment(ctx, organizationName, projectName, recipeDeploymentName).Execute()

Delete a Recipe Deployment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    organizationName := "organizationName_example" // string | The unique organization name
    projectName := "projectName_example" // string | The unique project name
    recipeDeploymentName := "recipeDeploymentName_example" // string | The unique recipe deployment name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RecipeDeploymentsAPI.DeleteRecipeDeployment(context.Background(), organizationName, projectName, recipeDeploymentName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecipeDeploymentsAPI.DeleteRecipeDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | The unique organization name | 
**projectName** | **string** | The unique project name | 
**recipeDeploymentName** | **string** | The unique recipe deployment name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecipeDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecipeDeployment

> RecipeDeployment GetRecipeDeployment(ctx, organizationName, projectName, recipeDeploymentName).Execute()

Get a Recipe Deployment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    organizationName := "organizationName_example" // string | The unique organization name
    projectName := "projectName_example" // string | The unique project name
    recipeDeploymentName := "recipeDeploymentName_example" // string | The unique recipe deployment name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecipeDeploymentsAPI.GetRecipeDeployment(context.Background(), organizationName, projectName, recipeDeploymentName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecipeDeploymentsAPI.GetRecipeDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecipeDeployment`: RecipeDeployment
    fmt.Fprintf(os.Stdout, "Response from `RecipeDeploymentsAPI.GetRecipeDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | The unique organization name | 
**projectName** | **string** | The unique project name | 
**recipeDeploymentName** | **string** | The unique recipe deployment name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecipeDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RecipeDeployment**](RecipeDeployment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecipeDeploymentInstances

> RecipeDeploymentInstances ListRecipeDeploymentInstances(ctx, organizationName, projectName, recipeDeploymentName).Execute()

List Recipe Deployment Instances



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    organizationName := "organizationName_example" // string | The unique organization name
    projectName := "projectName_example" // string | The unique project name
    recipeDeploymentName := "recipeDeploymentName_example" // string | The unique recipe deployment name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecipeDeploymentsAPI.ListRecipeDeploymentInstances(context.Background(), organizationName, projectName, recipeDeploymentName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecipeDeploymentsAPI.ListRecipeDeploymentInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRecipeDeploymentInstances`: RecipeDeploymentInstances
    fmt.Fprintf(os.Stdout, "Response from `RecipeDeploymentsAPI.ListRecipeDeploymentInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | The unique organization name | 
**projectName** | **string** | The unique project name | 
**recipeDeploymentName** | **string** | The unique recipe deployment name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRecipeDeploymentInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RecipeDeploymentInstances**](RecipeDeploymentInstances.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecipeDeployments

> RecipeDeploymentList ListRecipeDeployments(ctx, organizationName, projectName).Execute()

List Recipe Deployments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    organizationName := "organizationName_example" // string | The unique organization name
    projectName := "projectName_example" // string | The unique project name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecipeDeploymentsAPI.ListRecipeDeployments(context.Background(), organizationName, projectName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecipeDeploymentsAPI.ListRecipeDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRecipeDeployments`: RecipeDeploymentList
    fmt.Fprintf(os.Stdout, "Response from `RecipeDeploymentsAPI.ListRecipeDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | The unique organization name | 
**projectName** | **string** | The unique project name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRecipeDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RecipeDeploymentList**](RecipeDeploymentList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartDeployedRecipe

> RestartDeployedRecipe(ctx, organizationName, projectName, recipeDeploymentName).Execute()

Restart a Deployed Recipe



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    organizationName := "organizationName_example" // string | The unique organization name
    projectName := "projectName_example" // string | The unique project name
    recipeDeploymentName := "recipeDeploymentName_example" // string | The unique recipe deployment name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RecipeDeploymentsAPI.RestartDeployedRecipe(context.Background(), organizationName, projectName, recipeDeploymentName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecipeDeploymentsAPI.RestartDeployedRecipe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | The unique organization name | 
**projectName** | **string** | The unique project name | 
**recipeDeploymentName** | **string** | The unique recipe deployment name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartDeployedRecipeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartDeployedRecipe

> StartDeployedRecipe(ctx, organizationName, projectName, recipeDeploymentName).Execute()

Start a Deployed Recipe



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    organizationName := "organizationName_example" // string | The unique organization name
    projectName := "projectName_example" // string | The unique project name
    recipeDeploymentName := "recipeDeploymentName_example" // string | The unique recipe deployment name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RecipeDeploymentsAPI.StartDeployedRecipe(context.Background(), organizationName, projectName, recipeDeploymentName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecipeDeploymentsAPI.StartDeployedRecipe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | The unique organization name | 
**projectName** | **string** | The unique project name | 
**recipeDeploymentName** | **string** | The unique recipe deployment name | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartDeployedRecipeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopDeployedRecipe

> StopDeployedRecipe(ctx, organizationName, projectName, recipeDeploymentName).Execute()

Stop a Deployed Recipe



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    organizationName := "organizationName_example" // string | The unique organization name
    projectName := "projectName_example" // string | The unique project name
    recipeDeploymentName := "recipeDeploymentName_example" // string | The unique recipe deployment name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RecipeDeploymentsAPI.StopDeployedRecipe(context.Background(), organizationName, projectName, recipeDeploymentName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecipeDeploymentsAPI.StopDeployedRecipe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | The unique organization name | 
**projectName** | **string** | The unique project name | 
**recipeDeploymentName** | **string** | The unique recipe deployment name | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopDeployedRecipeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRecipeDeployment

> RecipeDeployment UpdateRecipeDeployment(ctx, organizationName, projectName, recipeDeploymentName).UpdateRecipeDeployment(updateRecipeDeployment).Execute()

Update a Recipe Deployment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    organizationName := "organizationName_example" // string | The unique organization name
    projectName := "projectName_example" // string | The unique project name
    recipeDeploymentName := "recipeDeploymentName_example" // string | The unique recipe deployment name
    updateRecipeDeployment := *openapiclient.NewUpdateRecipeDeployment() // UpdateRecipeDeployment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecipeDeploymentsAPI.UpdateRecipeDeployment(context.Background(), organizationName, projectName, recipeDeploymentName).UpdateRecipeDeployment(updateRecipeDeployment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecipeDeploymentsAPI.UpdateRecipeDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRecipeDeployment`: RecipeDeployment
    fmt.Fprintf(os.Stdout, "Response from `RecipeDeploymentsAPI.UpdateRecipeDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | The unique organization name | 
**projectName** | **string** | The unique project name | 
**recipeDeploymentName** | **string** | The unique recipe deployment name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRecipeDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateRecipeDeployment** | [**UpdateRecipeDeployment**](UpdateRecipeDeployment.md) |  | 

### Return type

[**RecipeDeployment**](RecipeDeployment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

