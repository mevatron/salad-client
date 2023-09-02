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
	"fmt"
)

// ContainerGroupStatus the model 'ContainerGroupStatus'
type ContainerGroupStatus string

// List of ContainerGroupStatus
const (
	CONTAINERGROUPSTATUS_PENDING ContainerGroupStatus = "pending"
	CONTAINERGROUPSTATUS_RUNNING ContainerGroupStatus = "running"
	CONTAINERGROUPSTATUS_STOPPED ContainerGroupStatus = "stopped"
	CONTAINERGROUPSTATUS_SUCCEEDED ContainerGroupStatus = "succeeded"
	CONTAINERGROUPSTATUS_FAILED ContainerGroupStatus = "failed"
	CONTAINERGROUPSTATUS_DEPLOYING ContainerGroupStatus = "deploying"
)

// All allowed values of ContainerGroupStatus enum
var AllowedContainerGroupStatusEnumValues = []ContainerGroupStatus{
	"pending",
	"running",
	"stopped",
	"succeeded",
	"failed",
	"deploying",
}

func (v *ContainerGroupStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContainerGroupStatus(value)
	for _, existing := range AllowedContainerGroupStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContainerGroupStatus", value)
}

// NewContainerGroupStatusFromValue returns a pointer to a valid ContainerGroupStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContainerGroupStatusFromValue(v string) (*ContainerGroupStatus, error) {
	ev := ContainerGroupStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContainerGroupStatus: valid values are %v", v, AllowedContainerGroupStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContainerGroupStatus) IsValid() bool {
	for _, existing := range AllowedContainerGroupStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContainerGroupStatus value
func (v ContainerGroupStatus) Ptr() *ContainerGroupStatus {
	return &v
}

type NullableContainerGroupStatus struct {
	value *ContainerGroupStatus
	isSet bool
}

func (v NullableContainerGroupStatus) Get() *ContainerGroupStatus {
	return v.value
}

func (v *NullableContainerGroupStatus) Set(val *ContainerGroupStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerGroupStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerGroupStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerGroupStatus(val *ContainerGroupStatus) *NullableContainerGroupStatus {
	return &NullableContainerGroupStatus{value: val, isSet: true}
}

func (v NullableContainerGroupStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerGroupStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

