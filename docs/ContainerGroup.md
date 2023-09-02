# ContainerGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**DisplayName** | **string** |  | 
**Container** | [**Container**](Container.md) |  | 
**RestartPolicy** | [**ContainerRestartPolicy**](ContainerRestartPolicy.md) |  | 
**Replicas** | **int32** |  | 
**CurrentState** | [**ContainerGroupState**](ContainerGroupState.md) |  | 
**CountryCodes** | Pointer to [**[]CountryCode**](CountryCode.md) |  | [optional] 
**Networking** | Pointer to [**NullableContainerGroupNetworking**](ContainerGroupNetworking.md) |  | [optional] 
**LivenessProbe** | Pointer to [**NullableContainerGroupProbe**](ContainerGroupProbe.md) |  | [optional] 
**ReadinessProbe** | Pointer to [**NullableContainerGroupProbe**](ContainerGroupProbe.md) |  | [optional] 
**StartupProbe** | Pointer to [**NullableContainerGroupProbe**](ContainerGroupProbe.md) |  | [optional] 
**CreateTime** | **time.Time** |  | 
**UpdateTime** | **time.Time** |  | 

## Methods

### NewContainerGroup

`func NewContainerGroup(id string, name string, displayName string, container Container, restartPolicy ContainerRestartPolicy, replicas int32, currentState ContainerGroupState, createTime time.Time, updateTime time.Time, ) *ContainerGroup`

NewContainerGroup instantiates a new ContainerGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupWithDefaults

`func NewContainerGroupWithDefaults() *ContainerGroup`

NewContainerGroupWithDefaults instantiates a new ContainerGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContainerGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerGroup) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ContainerGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *ContainerGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ContainerGroup) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ContainerGroup) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetContainer

`func (o *ContainerGroup) GetContainer() Container`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ContainerGroup) GetContainerOk() (*Container, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ContainerGroup) SetContainer(v Container)`

SetContainer sets Container field to given value.


### GetRestartPolicy

`func (o *ContainerGroup) GetRestartPolicy() ContainerRestartPolicy`

GetRestartPolicy returns the RestartPolicy field if non-nil, zero value otherwise.

### GetRestartPolicyOk

`func (o *ContainerGroup) GetRestartPolicyOk() (*ContainerRestartPolicy, bool)`

GetRestartPolicyOk returns a tuple with the RestartPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartPolicy

`func (o *ContainerGroup) SetRestartPolicy(v ContainerRestartPolicy)`

SetRestartPolicy sets RestartPolicy field to given value.


### GetReplicas

`func (o *ContainerGroup) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *ContainerGroup) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *ContainerGroup) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetCurrentState

`func (o *ContainerGroup) GetCurrentState() ContainerGroupState`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *ContainerGroup) GetCurrentStateOk() (*ContainerGroupState, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *ContainerGroup) SetCurrentState(v ContainerGroupState)`

SetCurrentState sets CurrentState field to given value.


### GetCountryCodes

`func (o *ContainerGroup) GetCountryCodes() []CountryCode`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *ContainerGroup) GetCountryCodesOk() (*[]CountryCode, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *ContainerGroup) SetCountryCodes(v []CountryCode)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *ContainerGroup) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetNetworking

`func (o *ContainerGroup) GetNetworking() ContainerGroupNetworking`

GetNetworking returns the Networking field if non-nil, zero value otherwise.

### GetNetworkingOk

`func (o *ContainerGroup) GetNetworkingOk() (*ContainerGroupNetworking, bool)`

GetNetworkingOk returns a tuple with the Networking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworking

`func (o *ContainerGroup) SetNetworking(v ContainerGroupNetworking)`

SetNetworking sets Networking field to given value.

### HasNetworking

`func (o *ContainerGroup) HasNetworking() bool`

HasNetworking returns a boolean if a field has been set.

### SetNetworkingNil

`func (o *ContainerGroup) SetNetworkingNil(b bool)`

 SetNetworkingNil sets the value for Networking to be an explicit nil

### UnsetNetworking
`func (o *ContainerGroup) UnsetNetworking()`

UnsetNetworking ensures that no value is present for Networking, not even an explicit nil
### GetLivenessProbe

`func (o *ContainerGroup) GetLivenessProbe() ContainerGroupProbe`

GetLivenessProbe returns the LivenessProbe field if non-nil, zero value otherwise.

### GetLivenessProbeOk

`func (o *ContainerGroup) GetLivenessProbeOk() (*ContainerGroupProbe, bool)`

GetLivenessProbeOk returns a tuple with the LivenessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbe

`func (o *ContainerGroup) SetLivenessProbe(v ContainerGroupProbe)`

SetLivenessProbe sets LivenessProbe field to given value.

### HasLivenessProbe

`func (o *ContainerGroup) HasLivenessProbe() bool`

HasLivenessProbe returns a boolean if a field has been set.

### SetLivenessProbeNil

`func (o *ContainerGroup) SetLivenessProbeNil(b bool)`

 SetLivenessProbeNil sets the value for LivenessProbe to be an explicit nil

### UnsetLivenessProbe
`func (o *ContainerGroup) UnsetLivenessProbe()`

UnsetLivenessProbe ensures that no value is present for LivenessProbe, not even an explicit nil
### GetReadinessProbe

`func (o *ContainerGroup) GetReadinessProbe() ContainerGroupProbe`

GetReadinessProbe returns the ReadinessProbe field if non-nil, zero value otherwise.

### GetReadinessProbeOk

`func (o *ContainerGroup) GetReadinessProbeOk() (*ContainerGroupProbe, bool)`

GetReadinessProbeOk returns a tuple with the ReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbe

`func (o *ContainerGroup) SetReadinessProbe(v ContainerGroupProbe)`

SetReadinessProbe sets ReadinessProbe field to given value.

### HasReadinessProbe

`func (o *ContainerGroup) HasReadinessProbe() bool`

HasReadinessProbe returns a boolean if a field has been set.

### SetReadinessProbeNil

`func (o *ContainerGroup) SetReadinessProbeNil(b bool)`

 SetReadinessProbeNil sets the value for ReadinessProbe to be an explicit nil

### UnsetReadinessProbe
`func (o *ContainerGroup) UnsetReadinessProbe()`

UnsetReadinessProbe ensures that no value is present for ReadinessProbe, not even an explicit nil
### GetStartupProbe

`func (o *ContainerGroup) GetStartupProbe() ContainerGroupProbe`

GetStartupProbe returns the StartupProbe field if non-nil, zero value otherwise.

### GetStartupProbeOk

`func (o *ContainerGroup) GetStartupProbeOk() (*ContainerGroupProbe, bool)`

GetStartupProbeOk returns a tuple with the StartupProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupProbe

`func (o *ContainerGroup) SetStartupProbe(v ContainerGroupProbe)`

SetStartupProbe sets StartupProbe field to given value.

### HasStartupProbe

`func (o *ContainerGroup) HasStartupProbe() bool`

HasStartupProbe returns a boolean if a field has been set.

### SetStartupProbeNil

`func (o *ContainerGroup) SetStartupProbeNil(b bool)`

 SetStartupProbeNil sets the value for StartupProbe to be an explicit nil

### UnsetStartupProbe
`func (o *ContainerGroup) UnsetStartupProbe()`

UnsetStartupProbe ensures that no value is present for StartupProbe, not even an explicit nil
### GetCreateTime

`func (o *ContainerGroup) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ContainerGroup) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ContainerGroup) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetUpdateTime

`func (o *ContainerGroup) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *ContainerGroup) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *ContainerGroup) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


