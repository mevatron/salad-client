# \OrganizationDataAPI

All URIs are relative to *https://api.salad.com/api/public*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListGpuClasses**](OrganizationDataAPI.md#ListGpuClasses) | **Get** /organizations/{organization_name}/gpu-classes | List the GPU Classes



## ListGpuClasses

> GpuClassesList ListGpuClasses(ctx, organizationName).Execute()

List the GPU Classes



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
    resp, r, err := apiClient.OrganizationDataAPI.ListGpuClasses(context.Background(), organizationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationDataAPI.ListGpuClasses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGpuClasses`: GpuClassesList
    fmt.Fprintf(os.Stdout, "Response from `OrganizationDataAPI.ListGpuClasses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | The unique organization name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGpuClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GpuClassesList**](GpuClassesList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

