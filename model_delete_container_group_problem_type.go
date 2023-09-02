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

// DeleteContainerGroupProblemType the model 'DeleteContainerGroupProblemType'
type DeleteContainerGroupProblemType string

// List of DeleteContainerGroupProblemType
const (
	DELETECONTAINERGROUPPROBLEMTYPE_NULL DeleteContainerGroupProblemType = "null"
)

// All allowed values of DeleteContainerGroupProblemType enum
var AllowedDeleteContainerGroupProblemTypeEnumValues = []DeleteContainerGroupProblemType{
	"null",
}

func (v *DeleteContainerGroupProblemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeleteContainerGroupProblemType(value)
	for _, existing := range AllowedDeleteContainerGroupProblemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeleteContainerGroupProblemType", value)
}

// NewDeleteContainerGroupProblemTypeFromValue returns a pointer to a valid DeleteContainerGroupProblemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeleteContainerGroupProblemTypeFromValue(v string) (*DeleteContainerGroupProblemType, error) {
	ev := DeleteContainerGroupProblemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeleteContainerGroupProblemType: valid values are %v", v, AllowedDeleteContainerGroupProblemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeleteContainerGroupProblemType) IsValid() bool {
	for _, existing := range AllowedDeleteContainerGroupProblemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeleteContainerGroupProblemType value
func (v DeleteContainerGroupProblemType) Ptr() *DeleteContainerGroupProblemType {
	return &v
}

type NullableDeleteContainerGroupProblemType struct {
	value *DeleteContainerGroupProblemType
	isSet bool
}

func (v NullableDeleteContainerGroupProblemType) Get() *DeleteContainerGroupProblemType {
	return v.value
}

func (v *NullableDeleteContainerGroupProblemType) Set(val *DeleteContainerGroupProblemType) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteContainerGroupProblemType) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteContainerGroupProblemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteContainerGroupProblemType(val *DeleteContainerGroupProblemType) *NullableDeleteContainerGroupProblemType {
	return &NullableDeleteContainerGroupProblemType{value: val, isSet: true}
}

func (v NullableDeleteContainerGroupProblemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteContainerGroupProblemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

