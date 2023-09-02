# CreateContainerGroupNetworking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | [**ContainerNetworkingProtocol**](ContainerNetworkingProtocol.md) |  | 
**Port** | **int32** |  | 
**Auth** | **bool** |  | 

## Methods

### NewCreateContainerGroupNetworking

`func NewCreateContainerGroupNetworking(protocol ContainerNetworkingProtocol, port int32, auth bool, ) *CreateContainerGroupNetworking`

NewCreateContainerGroupNetworking instantiates a new CreateContainerGroupNetworking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerGroupNetworkingWithDefaults

`func NewCreateContainerGroupNetworkingWithDefaults() *CreateContainerGroupNetworking`

NewCreateContainerGroupNetworkingWithDefaults instantiates a new CreateContainerGroupNetworking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *CreateContainerGroupNetworking) GetProtocol() ContainerNetworkingProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateContainerGroupNetworking) GetProtocolOk() (*ContainerNetworkingProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateContainerGroupNetworking) SetProtocol(v ContainerNetworkingProtocol)`

SetProtocol sets Protocol field to given value.


### GetPort

`func (o *CreateContainerGroupNetworking) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateContainerGroupNetworking) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateContainerGroupNetworking) SetPort(v int32)`

SetPort sets Port field to given value.


### GetAuth

`func (o *CreateContainerGroupNetworking) GetAuth() bool`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *CreateContainerGroupNetworking) GetAuthOk() (*bool, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *CreateContainerGroupNetworking) SetAuth(v bool)`

SetAuth sets Auth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


