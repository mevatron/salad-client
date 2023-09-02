# RecipeDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**DisplayName** | **string** |  | 
**Replicas** | **int32** |  | 
**CurrentState** | [**ContainerGroupState**](ContainerGroupState.md) |  | 
**Recipe** | [**Recipe**](Recipe.md) |  | 
**Networking** | Pointer to [**NullableRecipeNetworking**](RecipeNetworking.md) |  | [optional] 

## Methods

### NewRecipeDeployment

`func NewRecipeDeployment(id string, name string, displayName string, replicas int32, currentState ContainerGroupState, recipe Recipe, ) *RecipeDeployment`

NewRecipeDeployment instantiates a new RecipeDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipeDeploymentWithDefaults

`func NewRecipeDeploymentWithDefaults() *RecipeDeployment`

NewRecipeDeploymentWithDefaults instantiates a new RecipeDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RecipeDeployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecipeDeployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecipeDeployment) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RecipeDeployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecipeDeployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecipeDeployment) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *RecipeDeployment) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RecipeDeployment) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RecipeDeployment) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetReplicas

`func (o *RecipeDeployment) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *RecipeDeployment) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *RecipeDeployment) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetCurrentState

`func (o *RecipeDeployment) GetCurrentState() ContainerGroupState`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *RecipeDeployment) GetCurrentStateOk() (*ContainerGroupState, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *RecipeDeployment) SetCurrentState(v ContainerGroupState)`

SetCurrentState sets CurrentState field to given value.


### GetRecipe

`func (o *RecipeDeployment) GetRecipe() Recipe`

GetRecipe returns the Recipe field if non-nil, zero value otherwise.

### GetRecipeOk

`func (o *RecipeDeployment) GetRecipeOk() (*Recipe, bool)`

GetRecipeOk returns a tuple with the Recipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipe

`func (o *RecipeDeployment) SetRecipe(v Recipe)`

SetRecipe sets Recipe field to given value.


### GetNetworking

`func (o *RecipeDeployment) GetNetworking() RecipeNetworking`

GetNetworking returns the Networking field if non-nil, zero value otherwise.

### GetNetworkingOk

`func (o *RecipeDeployment) GetNetworkingOk() (*RecipeNetworking, bool)`

GetNetworkingOk returns a tuple with the Networking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworking

`func (o *RecipeDeployment) SetNetworking(v RecipeNetworking)`

SetNetworking sets Networking field to given value.

### HasNetworking

`func (o *RecipeDeployment) HasNetworking() bool`

HasNetworking returns a boolean if a field has been set.

### SetNetworkingNil

`func (o *RecipeDeployment) SetNetworkingNil(b bool)`

 SetNetworkingNil sets the value for Networking to be an explicit nil

### UnsetNetworking
`func (o *RecipeDeployment) UnsetNetworking()`

UnsetNetworking ensures that no value is present for Networking, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


