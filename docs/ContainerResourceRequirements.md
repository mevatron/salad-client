# ContainerResourceRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | **int32** |  | 
**Memory** | **int32** |  | 
**GpuClass** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewContainerResourceRequirements

`func NewContainerResourceRequirements(cpu int32, memory int32, ) *ContainerResourceRequirements`

NewContainerResourceRequirements instantiates a new ContainerResourceRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerResourceRequirementsWithDefaults

`func NewContainerResourceRequirementsWithDefaults() *ContainerResourceRequirements`

NewContainerResourceRequirementsWithDefaults instantiates a new ContainerResourceRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *ContainerResourceRequirements) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ContainerResourceRequirements) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ContainerResourceRequirements) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *ContainerResourceRequirements) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ContainerResourceRequirements) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ContainerResourceRequirements) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### GetGpuClass

`func (o *ContainerResourceRequirements) GetGpuClass() string`

GetGpuClass returns the GpuClass field if non-nil, zero value otherwise.

### GetGpuClassOk

`func (o *ContainerResourceRequirements) GetGpuClassOk() (*string, bool)`

GetGpuClassOk returns a tuple with the GpuClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuClass

`func (o *ContainerResourceRequirements) SetGpuClass(v string)`

SetGpuClass sets GpuClass field to given value.

### HasGpuClass

`func (o *ContainerResourceRequirements) HasGpuClass() bool`

HasGpuClass returns a boolean if a field has been set.

### SetGpuClassNil

`func (o *ContainerResourceRequirements) SetGpuClassNil(b bool)`

 SetGpuClassNil sets the value for GpuClass to be an explicit nil

### UnsetGpuClass
`func (o *ContainerResourceRequirements) UnsetGpuClass()`

UnsetGpuClass ensures that no value is present for GpuClass, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


