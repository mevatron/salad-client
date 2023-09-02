# \ContainerGroupsAPI

All URIs are relative to *https://api.salad.com/api/public*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContainerGroupInstanceReallocate**](ContainerGroupsAPI.md#ContainerGroupInstanceReallocate) | **Post** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{machine_id}/reallocate | Reallocate container group instance to another node
[**ContainerGroupInstanceRecreate**](ContainerGroupsAPI.md#ContainerGroupInstanceRecreate) | **Post** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{machine_id}/recreate | Recreate container on a node
[**ContainerGroupInstanceRestart**](ContainerGroupsAPI.md#ContainerGroupInstanceRestart) | **Post** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{machine_id}/restart | Restart container on a node
[**CreateContainerGroup**](ContainerGroupsAPI.md#CreateContainerGroup) | **Post** /organizations/{organization_name}/projects/{project_name}/containers | Create a Container Group
[**DeleteContainerGroup**](ContainerGroupsAPI.md#DeleteContainerGroup) | **Delete** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name} | Delete a Container Group
[**GetContainerGroup**](ContainerGroupsAPI.md#GetContainerGroup) | **Get** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name} | Get a Container Group
[**ListContainerGroupInstances**](ContainerGroupsAPI.md#ListContainerGroupInstances) | **Get** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances | List Container Group Instances
[**ListContainerGroups**](ContainerGroupsAPI.md#ListContainerGroups) | **Get** /organizations/{organization_name}/projects/{project_name}/containers | List Container Groups
[**RestartContainerGroup**](ContainerGroupsAPI.md#RestartContainerGroup) | **Post** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/restart | Restart a Container Group
[**StartContainerGroup**](ContainerGroupsAPI.md#StartContainerGroup) | **Post** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/start | Start a Container Group
[**StopContainerGroup**](ContainerGroupsAPI.md#StopContainerGroup) | **Post** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/stop | Stop a Container Group
[**UpdateContainerGroup**](ContainerGroupsAPI.md#UpdateContainerGroup) | **Patch** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name} | Update a Container Group



## ContainerGroupInstanceReallocate

> ContainerGroupInstanceReallocate(ctx, organizationName, projectName, containerGroupName, machineId).Execute()

Reallocate container group instance to another node



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
    containerGroupName := "containerGroupName_example" // string | The unique container group name
    machineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique machine identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContainerGroupsAPI.ContainerGroupInstanceReallocate(context.Background(), organizationName, projectName, containerGroupName, machineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.ContainerGroupInstanceReallocate``: %v\n", err)
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
**containerGroupName** | **string** | The unique container group name | 
**machineId** | **string** | The unique machine identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerGroupInstanceReallocateRequest struct via the builder pattern


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


## ContainerGroupInstanceRecreate

> ContainerGroupInstanceRecreate(ctx, organizationName, projectName, containerGroupName, machineId).Execute()

Recreate container on a node



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
    containerGroupName := "containerGroupName_example" // string | The unique container group name
    machineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique machine identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContainerGroupsAPI.ContainerGroupInstanceRecreate(context.Background(), organizationName, projectName, containerGroupName, machineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.ContainerGroupInstanceRecreate``: %v\n", err)
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
**containerGroupName** | **string** | The unique container group name | 
**machineId** | **string** | The unique machine identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerGroupInstanceRecreateRequest struct via the builder pattern


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


## ContainerGroupInstanceRestart

> ContainerGroupInstanceRestart(ctx, organizationName, projectName, containerGroupName, machineId).Execute()

Restart container on a node



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
    containerGroupName := "containerGroupName_example" // string | The unique container group name
    machineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique machine identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContainerGroupsAPI.ContainerGroupInstanceRestart(context.Background(), organizationName, projectName, containerGroupName, machineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.ContainerGroupInstanceRestart``: %v\n", err)
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
**containerGroupName** | **string** | The unique container group name | 
**machineId** | **string** | The unique machine identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiContainerGroupInstanceRestartRequest struct via the builder pattern


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


## CreateContainerGroup

> ContainerGroup CreateContainerGroup(ctx, organizationName, projectName).CreateContainerGroup(createContainerGroup).Execute()

Create a Container Group



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
    createContainerGroup := *openapiclient.NewCreateContainerGroup("Name_example", *openapiclient.NewCreateContainer("Image_example", *openapiclient.NewContainerResourceRequirements(int32(123), int32(123))), openapiclient.ContainerRestartPolicy("always"), int32(123)) // CreateContainerGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerGroupsAPI.CreateContainerGroup(context.Background(), organizationName, projectName).CreateContainerGroup(createContainerGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.CreateContainerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContainerGroup`: ContainerGroup
    fmt.Fprintf(os.Stdout, "Response from `ContainerGroupsAPI.CreateContainerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | The unique organization name | 
**projectName** | **string** | The unique project name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createContainerGroup** | [**CreateContainerGroup**](CreateContainerGroup.md) |  | 

### Return type

[**ContainerGroup**](ContainerGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContainerGroup

> DeleteContainerGroup(ctx, organizationName, projectName, containerGroupName).Execute()

Delete a Container Group



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
    containerGroupName := "containerGroupName_example" // string | The unique container group name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContainerGroupsAPI.DeleteContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.DeleteContainerGroup``: %v\n", err)
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
**containerGroupName** | **string** | The unique container group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerGroupRequest struct via the builder pattern


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


## GetContainerGroup

> ContainerGroup GetContainerGroup(ctx, organizationName, projectName, containerGroupName).Execute()

Get a Container Group



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
    containerGroupName := "containerGroupName_example" // string | The unique container group name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerGroupsAPI.GetContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.GetContainerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerGroup`: ContainerGroup
    fmt.Fprintf(os.Stdout, "Response from `ContainerGroupsAPI.GetContainerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | The unique organization name | 
**projectName** | **string** | The unique project name | 
**containerGroupName** | **string** | The unique container group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ContainerGroup**](ContainerGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainerGroupInstances

> ContainerGroupInstances ListContainerGroupInstances(ctx, organizationName, projectName, containerGroupName).Execute()

List Container Group Instances



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
    containerGroupName := "containerGroupName_example" // string | The unique container group name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerGroupsAPI.ListContainerGroupInstances(context.Background(), organizationName, projectName, containerGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.ListContainerGroupInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContainerGroupInstances`: ContainerGroupInstances
    fmt.Fprintf(os.Stdout, "Response from `ContainerGroupsAPI.ListContainerGroupInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | The unique organization name | 
**projectName** | **string** | The unique project name | 
**containerGroupName** | **string** | The unique container group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerGroupInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ContainerGroupInstances**](ContainerGroupInstances.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainerGroups

> ContainerGroupList ListContainerGroups(ctx, organizationName, projectName).Execute()

List Container Groups



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
    resp, r, err := apiClient.ContainerGroupsAPI.ListContainerGroups(context.Background(), organizationName, projectName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.ListContainerGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContainerGroups`: ContainerGroupList
    fmt.Fprintf(os.Stdout, "Response from `ContainerGroupsAPI.ListContainerGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | The unique organization name | 
**projectName** | **string** | The unique project name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContainerGroupList**](ContainerGroupList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartContainerGroup

> RestartContainerGroup(ctx, organizationName, projectName, containerGroupName).Execute()

Restart a Container Group



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
    containerGroupName := "containerGroupName_example" // string | The unique container group name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContainerGroupsAPI.RestartContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.RestartContainerGroup``: %v\n", err)
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
**containerGroupName** | **string** | The unique container group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartContainerGroupRequest struct via the builder pattern


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


## StartContainerGroup

> StartContainerGroup(ctx, organizationName, projectName, containerGroupName).Execute()

Start a Container Group



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
    containerGroupName := "containerGroupName_example" // string | The unique container group name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContainerGroupsAPI.StartContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.StartContainerGroup``: %v\n", err)
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
**containerGroupName** | **string** | The unique container group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartContainerGroupRequest struct via the builder pattern


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


## StopContainerGroup

> StopContainerGroup(ctx, organizationName, projectName, containerGroupName).Execute()

Stop a Container Group



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
    containerGroupName := "containerGroupName_example" // string | The unique container group name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContainerGroupsAPI.StopContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.StopContainerGroup``: %v\n", err)
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
**containerGroupName** | **string** | The unique container group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopContainerGroupRequest struct via the builder pattern


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


## UpdateContainerGroup

> ContainerGroup UpdateContainerGroup(ctx, organizationName, projectName, containerGroupName).UpdateContainerGroup(updateContainerGroup).Execute()

Update a Container Group



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
    containerGroupName := "containerGroupName_example" // string | The unique container group name
    updateContainerGroup := *openapiclient.NewUpdateContainerGroup() // UpdateContainerGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContainerGroupsAPI.UpdateContainerGroup(context.Background(), organizationName, projectName, containerGroupName).UpdateContainerGroup(updateContainerGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.UpdateContainerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateContainerGroup`: ContainerGroup
    fmt.Fprintf(os.Stdout, "Response from `ContainerGroupsAPI.UpdateContainerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | The unique organization name | 
**projectName** | **string** | The unique project name | 
**containerGroupName** | **string** | The unique container group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContainerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateContainerGroup** | [**UpdateContainerGroup**](UpdateContainerGroup.md) |  | 

### Return type

[**ContainerGroup**](ContainerGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

