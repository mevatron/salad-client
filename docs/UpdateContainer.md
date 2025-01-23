# UpdateContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **NullableString** |  | [optional] 
**Resources** | Pointer to [**NullableUpdateContainerResources**](UpdateContainerResources.md) |  | [optional] 
**Command** | Pointer to **[]string** | Pass a command (and optional arguments) to override the ENTRYPOINT and CMD of a container image. | [optional] 
**Priority** | Pointer to [**NullableContainerGroupPriority**](ContainerGroupPriority.md) |  | [optional] 
**EnvironmentVariables** | Pointer to **map[string]string** |  | [optional] 
**Logging** | Pointer to [**NullableContainerLogging**](ContainerLogging.md) |  | [optional] 
**RegistryAuthentication** | Pointer to [**NullableCreateContainerRegistryAuthentication**](CreateContainerRegistryAuthentication.md) |  | [optional] 

## Methods

### NewUpdateContainer

`func NewUpdateContainer() *UpdateContainer`

NewUpdateContainer instantiates a new UpdateContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContainerWithDefaults

`func NewUpdateContainerWithDefaults() *UpdateContainer`

NewUpdateContainerWithDefaults instantiates a new UpdateContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *UpdateContainer) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *UpdateContainer) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *UpdateContainer) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *UpdateContainer) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *UpdateContainer) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *UpdateContainer) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetResources

`func (o *UpdateContainer) GetResources() UpdateContainerResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *UpdateContainer) GetResourcesOk() (*UpdateContainerResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *UpdateContainer) SetResources(v UpdateContainerResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *UpdateContainer) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *UpdateContainer) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *UpdateContainer) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetCommand

`func (o *UpdateContainer) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *UpdateContainer) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *UpdateContainer) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *UpdateContainer) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *UpdateContainer) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *UpdateContainer) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetPriority

`func (o *UpdateContainer) GetPriority() ContainerGroupPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateContainer) GetPriorityOk() (*ContainerGroupPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateContainer) SetPriority(v ContainerGroupPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UpdateContainer) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *UpdateContainer) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *UpdateContainer) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetEnvironmentVariables

`func (o *UpdateContainer) GetEnvironmentVariables() map[string]string`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *UpdateContainer) GetEnvironmentVariablesOk() (*map[string]string, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *UpdateContainer) SetEnvironmentVariables(v map[string]string)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *UpdateContainer) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetLogging

`func (o *UpdateContainer) GetLogging() ContainerLogging`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *UpdateContainer) GetLoggingOk() (*ContainerLogging, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *UpdateContainer) SetLogging(v ContainerLogging)`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *UpdateContainer) HasLogging() bool`

HasLogging returns a boolean if a field has been set.

### SetLoggingNil

`func (o *UpdateContainer) SetLoggingNil(b bool)`

 SetLoggingNil sets the value for Logging to be an explicit nil

### UnsetLogging
`func (o *UpdateContainer) UnsetLogging()`

UnsetLogging ensures that no value is present for Logging, not even an explicit nil
### GetRegistryAuthentication

`func (o *UpdateContainer) GetRegistryAuthentication() CreateContainerRegistryAuthentication`

GetRegistryAuthentication returns the RegistryAuthentication field if non-nil, zero value otherwise.

### GetRegistryAuthenticationOk

`func (o *UpdateContainer) GetRegistryAuthenticationOk() (*CreateContainerRegistryAuthentication, bool)`

GetRegistryAuthenticationOk returns a tuple with the RegistryAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryAuthentication

`func (o *UpdateContainer) SetRegistryAuthentication(v CreateContainerRegistryAuthentication)`

SetRegistryAuthentication sets RegistryAuthentication field to given value.

### HasRegistryAuthentication

`func (o *UpdateContainer) HasRegistryAuthentication() bool`

HasRegistryAuthentication returns a boolean if a field has been set.

### SetRegistryAuthenticationNil

`func (o *UpdateContainer) SetRegistryAuthenticationNil(b bool)`

 SetRegistryAuthenticationNil sets the value for RegistryAuthentication to be an explicit nil

### UnsetRegistryAuthentication
`func (o *UpdateContainer) UnsetRegistryAuthentication()`

UnsetRegistryAuthentication ensures that no value is present for RegistryAuthentication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


