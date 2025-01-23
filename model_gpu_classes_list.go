/*
SaladCloud API

The SaladCloud REST API. Please refer to the [SaladCloud API Documentation](https://docs.salad.com/api-reference) for more details.

API version: 0.9.0-alpha.5
Contact: cloud@salad.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package saladclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GpuClassesList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GpuClassesList{}

// GpuClassesList Represents a list of GPU classes
type GpuClassesList struct {
	// The list of GPU classes
	Items []GpuClass `json:"items"`
}

type _GpuClassesList GpuClassesList

// NewGpuClassesList instantiates a new GpuClassesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGpuClassesList(items []GpuClass) *GpuClassesList {
	this := GpuClassesList{}
	this.Items = items
	return &this
}

// NewGpuClassesListWithDefaults instantiates a new GpuClassesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGpuClassesListWithDefaults() *GpuClassesList {
	this := GpuClassesList{}
	return &this
}

// GetItems returns the Items field value
func (o *GpuClassesList) GetItems() []GpuClass {
	if o == nil {
		var ret []GpuClass
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *GpuClassesList) GetItemsOk() ([]GpuClass, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *GpuClassesList) SetItems(v []GpuClass) {
	o.Items = v
}

func (o GpuClassesList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GpuClassesList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *GpuClassesList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGpuClassesList := _GpuClassesList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGpuClassesList)

	if err != nil {
		return err
	}

	*o = GpuClassesList(varGpuClassesList)

	return err
}

type NullableGpuClassesList struct {
	value *GpuClassesList
	isSet bool
}

func (v NullableGpuClassesList) Get() *GpuClassesList {
	return v.value
}

func (v *NullableGpuClassesList) Set(val *GpuClassesList) {
	v.value = val
	v.isSet = true
}

func (v NullableGpuClassesList) IsSet() bool {
	return v.isSet
}

func (v *NullableGpuClassesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGpuClassesList(val *GpuClassesList) *NullableGpuClassesList {
	return &NullableGpuClassesList{value: val, isSet: true}
}

func (v NullableGpuClassesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGpuClassesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
