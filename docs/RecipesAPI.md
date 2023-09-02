# \RecipesAPI

All URIs are relative to *https://api.salad.com/api/public*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRecipe**](RecipesAPI.md#GetRecipe) | **Get** /organizations/{organization_name}/recipes/{recipe_name} | Get a Recipe
[**ListRecipes**](RecipesAPI.md#ListRecipes) | **Get** /organizations/{organization_name}/recipes | List Recipes



## GetRecipe

> Recipe GetRecipe(ctx, organizationName, recipeName).Execute()

Get a Recipe



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
    recipeName := "recipeName_example" // string | The unique recipe name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecipesAPI.GetRecipe(context.Background(), organizationName, recipeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecipesAPI.GetRecipe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecipe`: Recipe
    fmt.Fprintf(os.Stdout, "Response from `RecipesAPI.GetRecipe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | The unique organization name | 
**recipeName** | **string** | The unique recipe name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecipeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Recipe**](Recipe.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecipes

> RecipeList ListRecipes(ctx, organizationName).Execute()

List Recipes



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecipesAPI.ListRecipes(context.Background(), organizationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecipesAPI.ListRecipes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRecipes`: RecipeList
    fmt.Fprintf(os.Stdout, "Response from `RecipesAPI.ListRecipes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | The unique organization name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRecipesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RecipeList**](RecipeList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

