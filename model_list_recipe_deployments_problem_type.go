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

// ListRecipeDeploymentsProblemType the model 'ListRecipeDeploymentsProblemType'
type ListRecipeDeploymentsProblemType string

// List of ListRecipeDeploymentsProblemType
const (
	LISTRECIPEDEPLOYMENTSPROBLEMTYPE_NULL ListRecipeDeploymentsProblemType = "null"
)

// All allowed values of ListRecipeDeploymentsProblemType enum
var AllowedListRecipeDeploymentsProblemTypeEnumValues = []ListRecipeDeploymentsProblemType{
	"null",
}

func (v *ListRecipeDeploymentsProblemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ListRecipeDeploymentsProblemType(value)
	for _, existing := range AllowedListRecipeDeploymentsProblemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ListRecipeDeploymentsProblemType", value)
}

// NewListRecipeDeploymentsProblemTypeFromValue returns a pointer to a valid ListRecipeDeploymentsProblemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListRecipeDeploymentsProblemTypeFromValue(v string) (*ListRecipeDeploymentsProblemType, error) {
	ev := ListRecipeDeploymentsProblemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListRecipeDeploymentsProblemType: valid values are %v", v, AllowedListRecipeDeploymentsProblemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListRecipeDeploymentsProblemType) IsValid() bool {
	for _, existing := range AllowedListRecipeDeploymentsProblemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListRecipeDeploymentsProblemType value
func (v ListRecipeDeploymentsProblemType) Ptr() *ListRecipeDeploymentsProblemType {
	return &v
}

type NullableListRecipeDeploymentsProblemType struct {
	value *ListRecipeDeploymentsProblemType
	isSet bool
}

func (v NullableListRecipeDeploymentsProblemType) Get() *ListRecipeDeploymentsProblemType {
	return v.value
}

func (v *NullableListRecipeDeploymentsProblemType) Set(val *ListRecipeDeploymentsProblemType) {
	v.value = val
	v.isSet = true
}

func (v NullableListRecipeDeploymentsProblemType) IsSet() bool {
	return v.isSet
}

func (v *NullableListRecipeDeploymentsProblemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRecipeDeploymentsProblemType(val *ListRecipeDeploymentsProblemType) *NullableListRecipeDeploymentsProblemType {
	return &NullableListRecipeDeploymentsProblemType{value: val, isSet: true}
}

func (v NullableListRecipeDeploymentsProblemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRecipeDeploymentsProblemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

