# Container

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** |  | 
**Resources** | [**ContainerResourceRequirements**](ContainerResourceRequirements.md) |  | 
**Command** | **[]string** |  | 
**EnvironmentVariables** | Pointer to **map[string]string** |  | [optional] 
**Logging** | Pointer to [**NullableContainerLogging**](ContainerLogging.md) |  | [optional] 

## Methods

### NewContainer

`func NewContainer(image string, resources ContainerResourceRequirements, command []string, ) *Container`

NewContainer instantiates a new Container object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerWithDefaults

`func NewContainerWithDefaults() *Container`

NewContainerWithDefaults instantiates a new Container object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *Container) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Container) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Container) SetImage(v string)`

SetImage sets Image field to given value.


### GetResources

`func (o *Container) GetResources() ContainerResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Container) GetResourcesOk() (*ContainerResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Container) SetResources(v ContainerResourceRequirements)`

SetResources sets Resources field to given value.


### GetCommand

`func (o *Container) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Container) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Container) SetCommand(v []string)`

SetCommand sets Command field to given value.


### GetEnvironmentVariables

`func (o *Container) GetEnvironmentVariables() map[string]string`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *Container) GetEnvironmentVariablesOk() (*map[string]string, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *Container) SetEnvironmentVariables(v map[string]string)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *Container) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetLogging

`func (o *Container) GetLogging() ContainerLogging`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *Container) GetLoggingOk() (*ContainerLogging, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *Container) SetLogging(v ContainerLogging)`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *Container) HasLogging() bool`

HasLogging returns a boolean if a field has been set.

### SetLoggingNil

`func (o *Container) SetLoggingNil(b bool)`

 SetLoggingNil sets the value for Logging to be an explicit nil

### UnsetLogging
`func (o *Container) UnsetLogging()`

UnsetLogging ensures that no value is present for Logging, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


