# ContainerLogging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewRelic** | Pointer to [**NullableContainerLoggingNewRelic**](ContainerLoggingNewRelic.md) |  | [optional] 
**Splunk** | Pointer to [**NullableContainerLoggingSplunk**](ContainerLoggingSplunk.md) |  | [optional] 
**Tcp** | Pointer to [**NullableContainerLoggingTcp**](ContainerLoggingTcp.md) |  | [optional] 

## Methods

### NewContainerLogging

`func NewContainerLogging() *ContainerLogging`

NewContainerLogging instantiates a new ContainerLogging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerLoggingWithDefaults

`func NewContainerLoggingWithDefaults() *ContainerLogging`

NewContainerLoggingWithDefaults instantiates a new ContainerLogging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewRelic

`func (o *ContainerLogging) GetNewRelic() ContainerLoggingNewRelic`

GetNewRelic returns the NewRelic field if non-nil, zero value otherwise.

### GetNewRelicOk

`func (o *ContainerLogging) GetNewRelicOk() (*ContainerLoggingNewRelic, bool)`

GetNewRelicOk returns a tuple with the NewRelic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRelic

`func (o *ContainerLogging) SetNewRelic(v ContainerLoggingNewRelic)`

SetNewRelic sets NewRelic field to given value.

### HasNewRelic

`func (o *ContainerLogging) HasNewRelic() bool`

HasNewRelic returns a boolean if a field has been set.

### SetNewRelicNil

`func (o *ContainerLogging) SetNewRelicNil(b bool)`

 SetNewRelicNil sets the value for NewRelic to be an explicit nil

### UnsetNewRelic
`func (o *ContainerLogging) UnsetNewRelic()`

UnsetNewRelic ensures that no value is present for NewRelic, not even an explicit nil
### GetSplunk

`func (o *ContainerLogging) GetSplunk() ContainerLoggingSplunk`

GetSplunk returns the Splunk field if non-nil, zero value otherwise.

### GetSplunkOk

`func (o *ContainerLogging) GetSplunkOk() (*ContainerLoggingSplunk, bool)`

GetSplunkOk returns a tuple with the Splunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplunk

`func (o *ContainerLogging) SetSplunk(v ContainerLoggingSplunk)`

SetSplunk sets Splunk field to given value.

### HasSplunk

`func (o *ContainerLogging) HasSplunk() bool`

HasSplunk returns a boolean if a field has been set.

### SetSplunkNil

`func (o *ContainerLogging) SetSplunkNil(b bool)`

 SetSplunkNil sets the value for Splunk to be an explicit nil

### UnsetSplunk
`func (o *ContainerLogging) UnsetSplunk()`

UnsetSplunk ensures that no value is present for Splunk, not even an explicit nil
### GetTcp

`func (o *ContainerLogging) GetTcp() ContainerLoggingTcp`

GetTcp returns the Tcp field if non-nil, zero value otherwise.

### GetTcpOk

`func (o *ContainerLogging) GetTcpOk() (*ContainerLoggingTcp, bool)`

GetTcpOk returns a tuple with the Tcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcp

`func (o *ContainerLogging) SetTcp(v ContainerLoggingTcp)`

SetTcp sets Tcp field to given value.

### HasTcp

`func (o *ContainerLogging) HasTcp() bool`

HasTcp returns a boolean if a field has been set.

### SetTcpNil

`func (o *ContainerLogging) SetTcpNil(b bool)`

 SetTcpNil sets the value for Tcp to be an explicit nil

### UnsetTcp
`func (o *ContainerLogging) UnsetTcp()`

UnsetTcp ensures that no value is present for Tcp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


