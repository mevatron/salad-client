# Recipe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier | 
**Name** | **string** | The recipe name | 
**Readme** | **string** | A markdown file containing a brief summary of the recipe | 
**Resources** | Pointer to [**RecipeResources**](RecipeResources.md) |  | [optional] 
**Networking** | Pointer to [**NullableRecipeNetworking**](RecipeNetworking.md) |  | [optional] 

## Methods

### NewRecipe

`func NewRecipe(id string, name string, readme string, ) *Recipe`

NewRecipe instantiates a new Recipe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipeWithDefaults

`func NewRecipeWithDefaults() *Recipe`

NewRecipeWithDefaults instantiates a new Recipe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Recipe) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Recipe) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Recipe) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Recipe) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Recipe) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Recipe) SetName(v string)`

SetName sets Name field to given value.


### GetReadme

`func (o *Recipe) GetReadme() string`

GetReadme returns the Readme field if non-nil, zero value otherwise.

### GetReadmeOk

`func (o *Recipe) GetReadmeOk() (*string, bool)`

GetReadmeOk returns a tuple with the Readme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadme

`func (o *Recipe) SetReadme(v string)`

SetReadme sets Readme field to given value.


### GetResources

`func (o *Recipe) GetResources() RecipeResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Recipe) GetResourcesOk() (*RecipeResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Recipe) SetResources(v RecipeResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *Recipe) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetNetworking

`func (o *Recipe) GetNetworking() RecipeNetworking`

GetNetworking returns the Networking field if non-nil, zero value otherwise.

### GetNetworkingOk

`func (o *Recipe) GetNetworkingOk() (*RecipeNetworking, bool)`

GetNetworkingOk returns a tuple with the Networking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworking

`func (o *Recipe) SetNetworking(v RecipeNetworking)`

SetNetworking sets Networking field to given value.

### HasNetworking

`func (o *Recipe) HasNetworking() bool`

HasNetworking returns a boolean if a field has been set.

### SetNetworkingNil

`func (o *Recipe) SetNetworkingNil(b bool)`

 SetNetworkingNil sets the value for Networking to be an explicit nil

### UnsetNetworking
`func (o *Recipe) UnsetNetworking()`

UnsetNetworking ensures that no value is present for Networking, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


