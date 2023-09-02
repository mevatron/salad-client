/*
SaladCloud Public API

The SaladCloud Public API.

API version: 1.0.0-alpha.56
Contact: support@salad.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package saladclient

import (
	"encoding/json"
)

// checks if the ContainerLoggingTcp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerLoggingTcp{}

// ContainerLoggingTcp struct for ContainerLoggingTcp
type ContainerLoggingTcp struct {
	Host string `json:"host"`
	Port int32 `json:"port"`
}

// NewContainerLoggingTcp instantiates a new ContainerLoggingTcp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerLoggingTcp(host string, port int32) *ContainerLoggingTcp {
	this := ContainerLoggingTcp{}
	this.Host = host
	this.Port = port
	return &this
}

// NewContainerLoggingTcpWithDefaults instantiates a new ContainerLoggingTcp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerLoggingTcpWithDefaults() *ContainerLoggingTcp {
	this := ContainerLoggingTcp{}
	return &this
}

// GetHost returns the Host field value
func (o *ContainerLoggingTcp) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *ContainerLoggingTcp) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *ContainerLoggingTcp) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *ContainerLoggingTcp) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ContainerLoggingTcp) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ContainerLoggingTcp) SetPort(v int32) {
	o.Port = v
}

func (o ContainerLoggingTcp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerLoggingTcp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["port"] = o.Port
	return toSerialize, nil
}

type NullableContainerLoggingTcp struct {
	value *ContainerLoggingTcp
	isSet bool
}

func (v NullableContainerLoggingTcp) Get() *ContainerLoggingTcp {
	return v.value
}

func (v *NullableContainerLoggingTcp) Set(val *ContainerLoggingTcp) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerLoggingTcp) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerLoggingTcp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerLoggingTcp(val *ContainerLoggingTcp) *NullableContainerLoggingTcp {
	return &NullableContainerLoggingTcp{value: val, isSet: true}
}

func (v NullableContainerLoggingTcp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerLoggingTcp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


