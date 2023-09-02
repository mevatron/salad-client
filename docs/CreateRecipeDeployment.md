# CreateRecipeDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DisplayName** | **NullableString** |  | 
**Replicas** | **int32** |  | 
**RecipeName** | **string** |  | 
**Networking** | Pointer to [**NullableCreateRecipeNetworking**](CreateRecipeNetworking.md) |  | [optional] 

## Methods

### NewCreateRecipeDeployment

`func NewCreateRecipeDeployment(name string, displayName NullableString, replicas int32, recipeName string, ) *CreateRecipeDeployment`

NewCreateRecipeDeployment instantiates a new CreateRecipeDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRecipeDeploymentWithDefaults

`func NewCreateRecipeDeploymentWithDefaults() *CreateRecipeDeployment`

NewCreateRecipeDeploymentWithDefaults instantiates a new CreateRecipeDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRecipeDeployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRecipeDeployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRecipeDeployment) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *CreateRecipeDeployment) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateRecipeDeployment) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateRecipeDeployment) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### SetDisplayNameNil

`func (o *CreateRecipeDeployment) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CreateRecipeDeployment) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetReplicas

`func (o *CreateRecipeDeployment) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *CreateRecipeDeployment) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *CreateRecipeDeployment) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetRecipeName

`func (o *CreateRecipeDeployment) GetRecipeName() string`

GetRecipeName returns the RecipeName field if non-nil, zero value otherwise.

### GetRecipeNameOk

`func (o *CreateRecipeDeployment) GetRecipeNameOk() (*string, bool)`

GetRecipeNameOk returns a tuple with the RecipeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipeName

`func (o *CreateRecipeDeployment) SetRecipeName(v string)`

SetRecipeName sets RecipeName field to given value.


### GetNetworking

`func (o *CreateRecipeDeployment) GetNetworking() CreateRecipeNetworking`

GetNetworking returns the Networking field if non-nil, zero value otherwise.

### GetNetworkingOk

`func (o *CreateRecipeDeployment) GetNetworkingOk() (*CreateRecipeNetworking, bool)`

GetNetworkingOk returns a tuple with the Networking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworking

`func (o *CreateRecipeDeployment) SetNetworking(v CreateRecipeNetworking)`

SetNetworking sets Networking field to given value.

### HasNetworking

`func (o *CreateRecipeDeployment) HasNetworking() bool`

HasNetworking returns a boolean if a field has been set.

### SetNetworkingNil

`func (o *CreateRecipeDeployment) SetNetworkingNil(b bool)`

 SetNetworkingNil sets the value for Networking to be an explicit nil

### UnsetNetworking
`func (o *CreateRecipeDeployment) UnsetNetworking()`

UnsetNetworking ensures that no value is present for Networking, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


