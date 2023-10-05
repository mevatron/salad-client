/*
SaladCloud Public API

The SaladCloud Public API.

API version: 1.0.0-alpha.56
Contact: support@salad.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the RecipeDeploymentInstancesInstancesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipeDeploymentInstancesInstancesInner{}

// RecipeDeploymentInstancesInstancesInner struct for RecipeDeploymentInstancesInstancesInner
type RecipeDeploymentInstancesInstancesInner struct {
	// The organization-specific machine ID
	MachineId string `json:"machine_id"`
	// The state of the recipe deployment instance
	State string `json:"state"`
	// The UTC date & time when the workload on this machine transitioned to the current state
	UpdateTime time.Time `json:"update_time"`
}

// NewRecipeDeploymentInstancesInstancesInner instantiates a new RecipeDeploymentInstancesInstancesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipeDeploymentInstancesInstancesInner(machineId string, state string, updateTime time.Time) *RecipeDeploymentInstancesInstancesInner {
	this := RecipeDeploymentInstancesInstancesInner{}
	this.MachineId = machineId
	this.State = state
	this.UpdateTime = updateTime
	return &this
}

// NewRecipeDeploymentInstancesInstancesInnerWithDefaults instantiates a new RecipeDeploymentInstancesInstancesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipeDeploymentInstancesInstancesInnerWithDefaults() *RecipeDeploymentInstancesInstancesInner {
	this := RecipeDeploymentInstancesInstancesInner{}
	return &this
}

// GetMachineId returns the MachineId field value
func (o *RecipeDeploymentInstancesInstancesInner) GetMachineId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MachineId
}

// GetMachineIdOk returns a tuple with the MachineId field value
// and a boolean to check if the value has been set.
func (o *RecipeDeploymentInstancesInstancesInner) GetMachineIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MachineId, true
}

// SetMachineId sets field value
func (o *RecipeDeploymentInstancesInstancesInner) SetMachineId(v string) {
	o.MachineId = v
}

// GetState returns the State field value
func (o *RecipeDeploymentInstancesInstancesInner) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *RecipeDeploymentInstancesInstancesInner) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *RecipeDeploymentInstancesInstancesInner) SetState(v string) {
	o.State = v
}

// GetUpdateTime returns the UpdateTime field value
func (o *RecipeDeploymentInstancesInstancesInner) GetUpdateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value
// and a boolean to check if the value has been set.
func (o *RecipeDeploymentInstancesInstancesInner) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateTime, true
}

// SetUpdateTime sets field value
func (o *RecipeDeploymentInstancesInstancesInner) SetUpdateTime(v time.Time) {
	o.UpdateTime = v
}

func (o RecipeDeploymentInstancesInstancesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipeDeploymentInstancesInstancesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["machine_id"] = o.MachineId
	toSerialize["state"] = o.State
	toSerialize["update_time"] = o.UpdateTime
	return toSerialize, nil
}

type NullableRecipeDeploymentInstancesInstancesInner struct {
	value *RecipeDeploymentInstancesInstancesInner
	isSet bool
}

func (v NullableRecipeDeploymentInstancesInstancesInner) Get() *RecipeDeploymentInstancesInstancesInner {
	return v.value
}

func (v *NullableRecipeDeploymentInstancesInstancesInner) Set(val *RecipeDeploymentInstancesInstancesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipeDeploymentInstancesInstancesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipeDeploymentInstancesInstancesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipeDeploymentInstancesInstancesInner(val *RecipeDeploymentInstancesInstancesInner) *NullableRecipeDeploymentInstancesInstancesInner {
	return &NullableRecipeDeploymentInstancesInstancesInner{value: val, isSet: true}
}

func (v NullableRecipeDeploymentInstancesInstancesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipeDeploymentInstancesInstancesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


