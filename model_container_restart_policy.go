/*
SaladCloud API

The SaladCloud REST API. Please refer to the [SaladCloud API Documentation](https://docs.salad.com/api-reference) for more details.

API version: 0.9.0-alpha.5
Contact: cloud@salad.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package saladclient

import (
	"encoding/json"
	"fmt"
)

// ContainerRestartPolicy the model 'ContainerRestartPolicy'
type ContainerRestartPolicy string

// List of ContainerRestartPolicy
const (
	CONTAINERRESTARTPOLICY_ALWAYS     ContainerRestartPolicy = "always"
	CONTAINERRESTARTPOLICY_ON_FAILURE ContainerRestartPolicy = "on_failure"
	CONTAINERRESTARTPOLICY_NEVER      ContainerRestartPolicy = "never"
)

// All allowed values of ContainerRestartPolicy enum
var AllowedContainerRestartPolicyEnumValues = []ContainerRestartPolicy{
	"always",
	"on_failure",
	"never",
}

func (v *ContainerRestartPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContainerRestartPolicy(value)
	for _, existing := range AllowedContainerRestartPolicyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContainerRestartPolicy", value)
}

// NewContainerRestartPolicyFromValue returns a pointer to a valid ContainerRestartPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContainerRestartPolicyFromValue(v string) (*ContainerRestartPolicy, error) {
	ev := ContainerRestartPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContainerRestartPolicy: valid values are %v", v, AllowedContainerRestartPolicyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContainerRestartPolicy) IsValid() bool {
	for _, existing := range AllowedContainerRestartPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContainerRestartPolicy value
func (v ContainerRestartPolicy) Ptr() *ContainerRestartPolicy {
	return &v
}

type NullableContainerRestartPolicy struct {
	value *ContainerRestartPolicy
	isSet bool
}

func (v NullableContainerRestartPolicy) Get() *ContainerRestartPolicy {
	return v.value
}

func (v *NullableContainerRestartPolicy) Set(val *ContainerRestartPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerRestartPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerRestartPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerRestartPolicy(val *ContainerRestartPolicy) *NullableContainerRestartPolicy {
	return &NullableContainerRestartPolicy{value: val, isSet: true}
}

func (v NullableContainerRestartPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerRestartPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
