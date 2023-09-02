# RecipeNetworking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | [**RecipeNetworkingProtocol**](RecipeNetworkingProtocol.md) |  | 
**Port** | **int32** |  | 
**Auth** | **bool** |  | 
**Dns** | **string** |  | 

## Methods

### NewRecipeNetworking

`func NewRecipeNetworking(protocol RecipeNetworkingProtocol, port int32, auth bool, dns string, ) *RecipeNetworking`

NewRecipeNetworking instantiates a new RecipeNetworking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipeNetworkingWithDefaults

`func NewRecipeNetworkingWithDefaults() *RecipeNetworking`

NewRecipeNetworkingWithDefaults instantiates a new RecipeNetworking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *RecipeNetworking) GetProtocol() RecipeNetworkingProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *RecipeNetworking) GetProtocolOk() (*RecipeNetworkingProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *RecipeNetworking) SetProtocol(v RecipeNetworkingProtocol)`

SetProtocol sets Protocol field to given value.


### GetPort

`func (o *RecipeNetworking) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RecipeNetworking) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RecipeNetworking) SetPort(v int32)`

SetPort sets Port field to given value.


### GetAuth

`func (o *RecipeNetworking) GetAuth() bool`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *RecipeNetworking) GetAuthOk() (*bool, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *RecipeNetworking) SetAuth(v bool)`

SetAuth sets Auth field to given value.


### GetDns

`func (o *RecipeNetworking) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *RecipeNetworking) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *RecipeNetworking) SetDns(v string)`

SetDns sets Dns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


