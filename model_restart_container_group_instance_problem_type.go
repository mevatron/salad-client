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

// RestartContainerGroupInstanceProblemType the model 'RestartContainerGroupInstanceProblemType'
type RestartContainerGroupInstanceProblemType string

// List of RestartContainerGroupInstanceProblemType
const (
	RESTARTCONTAINERGROUPINSTANCEPROBLEMTYPE_ORGANIZATION_NOT_FOUND RestartContainerGroupInstanceProblemType = "organization_not_found"
	RESTARTCONTAINERGROUPINSTANCEPROBLEMTYPE_CONTAINER_GROUP_NOT_FOUND RestartContainerGroupInstanceProblemType = "container_group_not_found"
	RESTARTCONTAINERGROUPINSTANCEPROBLEMTYPE_NODE_NOT_SCHEDULED RestartContainerGroupInstanceProblemType = "node_not_scheduled"
	RESTARTCONTAINERGROUPINSTANCEPROBLEMTYPE_CONTAINER_GROUP_INSTANCE_NOT_FOUND RestartContainerGroupInstanceProblemType = "container_group_instance_not_found"
	RESTARTCONTAINERGROUPINSTANCEPROBLEMTYPE_ORGANIZATION_QUOTAS_NOT_DEFINED RestartContainerGroupInstanceProblemType = "organization_quotas_not_defined"
	RESTARTCONTAINERGROUPINSTANCEPROBLEMTYPE_TOO_MANY_REQUESTS RestartContainerGroupInstanceProblemType = "too_many_requests"
	RESTARTCONTAINERGROUPINSTANCEPROBLEMTYPE_UNEXPECTED_ERROR RestartContainerGroupInstanceProblemType = "unexpected_error"
	RESTARTCONTAINERGROUPINSTANCEPROBLEMTYPE_NULL RestartContainerGroupInstanceProblemType = "null"
)

// All allowed values of RestartContainerGroupInstanceProblemType enum
var AllowedRestartContainerGroupInstanceProblemTypeEnumValues = []RestartContainerGroupInstanceProblemType{
	"organization_not_found",
	"container_group_not_found",
	"node_not_scheduled",
	"container_group_instance_not_found",
	"organization_quotas_not_defined",
	"too_many_requests",
	"unexpected_error",
	"null",
}

func (v *RestartContainerGroupInstanceProblemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RestartContainerGroupInstanceProblemType(value)
	for _, existing := range AllowedRestartContainerGroupInstanceProblemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RestartContainerGroupInstanceProblemType", value)
}

// NewRestartContainerGroupInstanceProblemTypeFromValue returns a pointer to a valid RestartContainerGroupInstanceProblemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRestartContainerGroupInstanceProblemTypeFromValue(v string) (*RestartContainerGroupInstanceProblemType, error) {
	ev := RestartContainerGroupInstanceProblemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RestartContainerGroupInstanceProblemType: valid values are %v", v, AllowedRestartContainerGroupInstanceProblemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RestartContainerGroupInstanceProblemType) IsValid() bool {
	for _, existing := range AllowedRestartContainerGroupInstanceProblemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RestartContainerGroupInstanceProblemType value
func (v RestartContainerGroupInstanceProblemType) Ptr() *RestartContainerGroupInstanceProblemType {
	return &v
}

type NullableRestartContainerGroupInstanceProblemType struct {
	value *RestartContainerGroupInstanceProblemType
	isSet bool
}

func (v NullableRestartContainerGroupInstanceProblemType) Get() *RestartContainerGroupInstanceProblemType {
	return v.value
}

func (v *NullableRestartContainerGroupInstanceProblemType) Set(val *RestartContainerGroupInstanceProblemType) {
	v.value = val
	v.isSet = true
}

func (v NullableRestartContainerGroupInstanceProblemType) IsSet() bool {
	return v.isSet
}

func (v *NullableRestartContainerGroupInstanceProblemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestartContainerGroupInstanceProblemType(val *RestartContainerGroupInstanceProblemType) *NullableRestartContainerGroupInstanceProblemType {
	return &NullableRestartContainerGroupInstanceProblemType{value: val, isSet: true}
}

func (v NullableRestartContainerGroupInstanceProblemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestartContainerGroupInstanceProblemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

