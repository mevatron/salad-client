# ContainerGroupsQuotas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxCreatedContainerGroups** | **int32** |  | 
**ContainerInstanceQuota** | **int32** |  | 
**MaxContainerGroupReallocationsPerMinute** | Pointer to **int32** |  | [optional] [default to 10]
**MaxContainerGroupRecreatesPerMinute** | Pointer to **int32** |  | [optional] [default to 10]
**MaxContainerGroupRestartsPerMinute** | Pointer to **int32** |  | [optional] [default to 10]

## Methods

### NewContainerGroupsQuotas

`func NewContainerGroupsQuotas(maxCreatedContainerGroups int32, containerInstanceQuota int32, ) *ContainerGroupsQuotas`

NewContainerGroupsQuotas instantiates a new ContainerGroupsQuotas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupsQuotasWithDefaults

`func NewContainerGroupsQuotasWithDefaults() *ContainerGroupsQuotas`

NewContainerGroupsQuotasWithDefaults instantiates a new ContainerGroupsQuotas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxCreatedContainerGroups

`func (o *ContainerGroupsQuotas) GetMaxCreatedContainerGroups() int32`

GetMaxCreatedContainerGroups returns the MaxCreatedContainerGroups field if non-nil, zero value otherwise.

### GetMaxCreatedContainerGroupsOk

`func (o *ContainerGroupsQuotas) GetMaxCreatedContainerGroupsOk() (*int32, bool)`

GetMaxCreatedContainerGroupsOk returns a tuple with the MaxCreatedContainerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCreatedContainerGroups

`func (o *ContainerGroupsQuotas) SetMaxCreatedContainerGroups(v int32)`

SetMaxCreatedContainerGroups sets MaxCreatedContainerGroups field to given value.


### GetContainerInstanceQuota

`func (o *ContainerGroupsQuotas) GetContainerInstanceQuota() int32`

GetContainerInstanceQuota returns the ContainerInstanceQuota field if non-nil, zero value otherwise.

### GetContainerInstanceQuotaOk

`func (o *ContainerGroupsQuotas) GetContainerInstanceQuotaOk() (*int32, bool)`

GetContainerInstanceQuotaOk returns a tuple with the ContainerInstanceQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerInstanceQuota

`func (o *ContainerGroupsQuotas) SetContainerInstanceQuota(v int32)`

SetContainerInstanceQuota sets ContainerInstanceQuota field to given value.


### GetMaxContainerGroupReallocationsPerMinute

`func (o *ContainerGroupsQuotas) GetMaxContainerGroupReallocationsPerMinute() int32`

GetMaxContainerGroupReallocationsPerMinute returns the MaxContainerGroupReallocationsPerMinute field if non-nil, zero value otherwise.

### GetMaxContainerGroupReallocationsPerMinuteOk

`func (o *ContainerGroupsQuotas) GetMaxContainerGroupReallocationsPerMinuteOk() (*int32, bool)`

GetMaxContainerGroupReallocationsPerMinuteOk returns a tuple with the MaxContainerGroupReallocationsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContainerGroupReallocationsPerMinute

`func (o *ContainerGroupsQuotas) SetMaxContainerGroupReallocationsPerMinute(v int32)`

SetMaxContainerGroupReallocationsPerMinute sets MaxContainerGroupReallocationsPerMinute field to given value.

### HasMaxContainerGroupReallocationsPerMinute

`func (o *ContainerGroupsQuotas) HasMaxContainerGroupReallocationsPerMinute() bool`

HasMaxContainerGroupReallocationsPerMinute returns a boolean if a field has been set.

### GetMaxContainerGroupRecreatesPerMinute

`func (o *ContainerGroupsQuotas) GetMaxContainerGroupRecreatesPerMinute() int32`

GetMaxContainerGroupRecreatesPerMinute returns the MaxContainerGroupRecreatesPerMinute field if non-nil, zero value otherwise.

### GetMaxContainerGroupRecreatesPerMinuteOk

`func (o *ContainerGroupsQuotas) GetMaxContainerGroupRecreatesPerMinuteOk() (*int32, bool)`

GetMaxContainerGroupRecreatesPerMinuteOk returns a tuple with the MaxContainerGroupRecreatesPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContainerGroupRecreatesPerMinute

`func (o *ContainerGroupsQuotas) SetMaxContainerGroupRecreatesPerMinute(v int32)`

SetMaxContainerGroupRecreatesPerMinute sets MaxContainerGroupRecreatesPerMinute field to given value.

### HasMaxContainerGroupRecreatesPerMinute

`func (o *ContainerGroupsQuotas) HasMaxContainerGroupRecreatesPerMinute() bool`

HasMaxContainerGroupRecreatesPerMinute returns a boolean if a field has been set.

### GetMaxContainerGroupRestartsPerMinute

`func (o *ContainerGroupsQuotas) GetMaxContainerGroupRestartsPerMinute() int32`

GetMaxContainerGroupRestartsPerMinute returns the MaxContainerGroupRestartsPerMinute field if non-nil, zero value otherwise.

### GetMaxContainerGroupRestartsPerMinuteOk

`func (o *ContainerGroupsQuotas) GetMaxContainerGroupRestartsPerMinuteOk() (*int32, bool)`

GetMaxContainerGroupRestartsPerMinuteOk returns a tuple with the MaxContainerGroupRestartsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContainerGroupRestartsPerMinute

`func (o *ContainerGroupsQuotas) SetMaxContainerGroupRestartsPerMinute(v int32)`

SetMaxContainerGroupRestartsPerMinute sets MaxContainerGroupRestartsPerMinute field to given value.

### HasMaxContainerGroupRestartsPerMinute

`func (o *ContainerGroupsQuotas) HasMaxContainerGroupRestartsPerMinute() bool`

HasMaxContainerGroupRestartsPerMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


