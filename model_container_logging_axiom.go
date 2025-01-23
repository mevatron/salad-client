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

// checks if the ContainerLoggingAxiom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerLoggingAxiom{}

// ContainerLoggingAxiom struct for ContainerLoggingAxiom
type ContainerLoggingAxiom struct {
	Host     string `json:"host"`
	ApiToken string `json:"api_token"`
	Dataset  string `json:"dataset"`
}

type _ContainerLoggingAxiom ContainerLoggingAxiom

// NewContainerLoggingAxiom instantiates a new ContainerLoggingAxiom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerLoggingAxiom(host string, apiToken string, dataset string) *ContainerLoggingAxiom {
	this := ContainerLoggingAxiom{}
	this.Host = host
	this.ApiToken = apiToken
	this.Dataset = dataset
	return &this
}

// NewContainerLoggingAxiomWithDefaults instantiates a new ContainerLoggingAxiom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerLoggingAxiomWithDefaults() *ContainerLoggingAxiom {
	this := ContainerLoggingAxiom{}
	return &this
}

// GetHost returns the Host field value
func (o *ContainerLoggingAxiom) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *ContainerLoggingAxiom) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *ContainerLoggingAxiom) SetHost(v string) {
	o.Host = v
}

// GetApiToken returns the ApiToken field value
func (o *ContainerLoggingAxiom) GetApiToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value
// and a boolean to check if the value has been set.
func (o *ContainerLoggingAxiom) GetApiTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiToken, true
}

// SetApiToken sets field value
func (o *ContainerLoggingAxiom) SetApiToken(v string) {
	o.ApiToken = v
}

// GetDataset returns the Dataset field value
func (o *ContainerLoggingAxiom) GetDataset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value
// and a boolean to check if the value has been set.
func (o *ContainerLoggingAxiom) GetDatasetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dataset, true
}

// SetDataset sets field value
func (o *ContainerLoggingAxiom) SetDataset(v string) {
	o.Dataset = v
}

func (o ContainerLoggingAxiom) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerLoggingAxiom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["api_token"] = o.ApiToken
	toSerialize["dataset"] = o.Dataset
	return toSerialize, nil
}

func (o *ContainerLoggingAxiom) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host",
		"api_token",
		"dataset",
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

	varContainerLoggingAxiom := _ContainerLoggingAxiom{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerLoggingAxiom)

	if err != nil {
		return err
	}

	*o = ContainerLoggingAxiom(varContainerLoggingAxiom)

	return err
}

type NullableContainerLoggingAxiom struct {
	value *ContainerLoggingAxiom
	isSet bool
}

func (v NullableContainerLoggingAxiom) Get() *ContainerLoggingAxiom {
	return v.value
}

func (v *NullableContainerLoggingAxiom) Set(val *ContainerLoggingAxiom) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerLoggingAxiom) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerLoggingAxiom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerLoggingAxiom(val *ContainerLoggingAxiom) *NullableContainerLoggingAxiom {
	return &NullableContainerLoggingAxiom{value: val, isSet: true}
}

func (v NullableContainerLoggingAxiom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerLoggingAxiom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
