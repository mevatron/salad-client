# RecipeDeploymentInstancesInstancesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineId** | **string** | The organization-specific machine ID | 
**State** | **string** | The state of the recipe deployment instance | 
**UpdateTime** | **time.Time** | The UTC date &amp; time when the workload on this machine transitioned to the current state | 

## Methods

### NewRecipeDeploymentInstancesInstancesInner

`func NewRecipeDeploymentInstancesInstancesInner(machineId string, state string, updateTime time.Time, ) *RecipeDeploymentInstancesInstancesInner`

NewRecipeDeploymentInstancesInstancesInner instantiates a new RecipeDeploymentInstancesInstancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipeDeploymentInstancesInstancesInnerWithDefaults

`func NewRecipeDeploymentInstancesInstancesInnerWithDefaults() *RecipeDeploymentInstancesInstancesInner`

NewRecipeDeploymentInstancesInstancesInnerWithDefaults instantiates a new RecipeDeploymentInstancesInstancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineId

`func (o *RecipeDeploymentInstancesInstancesInner) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *RecipeDeploymentInstancesInstancesInner) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *RecipeDeploymentInstancesInstancesInner) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.


### GetState

`func (o *RecipeDeploymentInstancesInstancesInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RecipeDeploymentInstancesInstancesInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RecipeDeploymentInstancesInstancesInner) SetState(v string)`

SetState sets State field to given value.


### GetUpdateTime

`func (o *RecipeDeploymentInstancesInstancesInner) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *RecipeDeploymentInstancesInstancesInner) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *RecipeDeploymentInstancesInstancesInner) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


