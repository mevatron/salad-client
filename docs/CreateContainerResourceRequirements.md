# CreateContainerResourceRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | **int32** | The number of CPU cores required by the container. Must be between 1 and 16. | 
**Memory** | **int32** | The amount of memory (in MB) required by the container. Must be between 1024 MB and 61440 MB. | 
**GpuClasses** | Pointer to **[]string** | A list of GPU class UUIDs required by the container. Can be null if no GPU is required. | [optional] 
**StorageAmount** | Pointer to **int64** | The amount of storage (in bytes) required by the container. Must be between 1 GB (1073741824 bytes) and 250 GB (268435456000 bytes). | [optional] 

## Methods

### NewCreateContainerResourceRequirements

`func NewCreateContainerResourceRequirements(cpu int32, memory int32, ) *CreateContainerResourceRequirements`

NewCreateContainerResourceRequirements instantiates a new CreateContainerResourceRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerResourceRequirementsWithDefaults

`func NewCreateContainerResourceRequirementsWithDefaults() *CreateContainerResourceRequirements`

NewCreateContainerResourceRequirementsWithDefaults instantiates a new CreateContainerResourceRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *CreateContainerResourceRequirements) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *CreateContainerResourceRequirements) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *CreateContainerResourceRequirements) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *CreateContainerResourceRequirements) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CreateContainerResourceRequirements) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CreateContainerResourceRequirements) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### GetGpuClasses

`func (o *CreateContainerResourceRequirements) GetGpuClasses() []string`

GetGpuClasses returns the GpuClasses field if non-nil, zero value otherwise.

### GetGpuClassesOk

`func (o *CreateContainerResourceRequirements) GetGpuClassesOk() (*[]string, bool)`

GetGpuClassesOk returns a tuple with the GpuClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuClasses

`func (o *CreateContainerResourceRequirements) SetGpuClasses(v []string)`

SetGpuClasses sets GpuClasses field to given value.

### HasGpuClasses

`func (o *CreateContainerResourceRequirements) HasGpuClasses() bool`

HasGpuClasses returns a boolean if a field has been set.

### GetStorageAmount

`func (o *CreateContainerResourceRequirements) GetStorageAmount() int64`

GetStorageAmount returns the StorageAmount field if non-nil, zero value otherwise.

### GetStorageAmountOk

`func (o *CreateContainerResourceRequirements) GetStorageAmountOk() (*int64, bool)`

GetStorageAmountOk returns a tuple with the StorageAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAmount

`func (o *CreateContainerResourceRequirements) SetStorageAmount(v int64)`

SetStorageAmount sets StorageAmount field to given value.

### HasStorageAmount

`func (o *CreateContainerResourceRequirements) HasStorageAmount() bool`

HasStorageAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


