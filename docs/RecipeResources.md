# RecipeResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | **int32** |  | 
**Ram** | **int32** |  | 
**GpuClass** | **string** |  | 

## Methods

### NewRecipeResources

`func NewRecipeResources(cpu int32, ram int32, gpuClass string, ) *RecipeResources`

NewRecipeResources instantiates a new RecipeResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipeResourcesWithDefaults

`func NewRecipeResourcesWithDefaults() *RecipeResources`

NewRecipeResourcesWithDefaults instantiates a new RecipeResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *RecipeResources) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *RecipeResources) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *RecipeResources) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetRam

`func (o *RecipeResources) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *RecipeResources) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *RecipeResources) SetRam(v int32)`

SetRam sets Ram field to given value.


### GetGpuClass

`func (o *RecipeResources) GetGpuClass() string`

GetGpuClass returns the GpuClass field if non-nil, zero value otherwise.

### GetGpuClassOk

`func (o *RecipeResources) GetGpuClassOk() (*string, bool)`

GetGpuClassOk returns a tuple with the GpuClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuClass

`func (o *RecipeResources) SetGpuClass(v string)`

SetGpuClass sets GpuClass field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


