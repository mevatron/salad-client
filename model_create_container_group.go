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

// checks if the CreateContainerGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateContainerGroup{}

// CreateContainerGroup Represents a request to create a container group
type CreateContainerGroup struct {
	Name string `json:"name"`
	DisplayName NullableString `json:"display_name,omitempty"`
	Container CreateContainer `json:"container"`
	RestartPolicy ContainerRestartPolicy `json:"restart_policy"`
	Replicas int32 `json:"replicas"`
	CountryCodes []CountryCode `json:"country_codes,omitempty"`
	Networking NullableCreateContainerGroupNetworking `json:"networking,omitempty"`
	LivenessProbe NullableContainerGroupProbe `json:"liveness_probe,omitempty"`
	ReadinessProbe NullableContainerGroupProbe `json:"readiness_probe,omitempty"`
	StartupProbe NullableContainerGroupProbe `json:"startup_probe,omitempty"`
}

// NewCreateContainerGroup instantiates a new CreateContainerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContainerGroup(name string, container CreateContainer, restartPolicy ContainerRestartPolicy, replicas int32) *CreateContainerGroup {
	this := CreateContainerGroup{}
	this.Name = name
	this.Container = container
	this.RestartPolicy = restartPolicy
	this.Replicas = replicas
	return &this
}

// NewCreateContainerGroupWithDefaults instantiates a new CreateContainerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContainerGroupWithDefaults() *CreateContainerGroup {
	this := CreateContainerGroup{}
	return &this
}

// GetName returns the Name field value
func (o *CreateContainerGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateContainerGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateContainerGroup) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContainerGroup) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContainerGroup) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateContainerGroup) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *CreateContainerGroup) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *CreateContainerGroup) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *CreateContainerGroup) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetContainer returns the Container field value
func (o *CreateContainerGroup) GetContainer() CreateContainer {
	if o == nil {
		var ret CreateContainer
		return ret
	}

	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *CreateContainerGroup) GetContainerOk() (*CreateContainer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Container, true
}

// SetContainer sets field value
func (o *CreateContainerGroup) SetContainer(v CreateContainer) {
	o.Container = v
}

// GetRestartPolicy returns the RestartPolicy field value
func (o *CreateContainerGroup) GetRestartPolicy() ContainerRestartPolicy {
	if o == nil {
		var ret ContainerRestartPolicy
		return ret
	}

	return o.RestartPolicy
}

// GetRestartPolicyOk returns a tuple with the RestartPolicy field value
// and a boolean to check if the value has been set.
func (o *CreateContainerGroup) GetRestartPolicyOk() (*ContainerRestartPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestartPolicy, true
}

// SetRestartPolicy sets field value
func (o *CreateContainerGroup) SetRestartPolicy(v ContainerRestartPolicy) {
	o.RestartPolicy = v
}

// GetReplicas returns the Replicas field value
func (o *CreateContainerGroup) GetReplicas() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *CreateContainerGroup) GetReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replicas, true
}

// SetReplicas sets field value
func (o *CreateContainerGroup) SetReplicas(v int32) {
	o.Replicas = v
}

// GetCountryCodes returns the CountryCodes field value if set, zero value otherwise.
func (o *CreateContainerGroup) GetCountryCodes() []CountryCode {
	if o == nil || IsNil(o.CountryCodes) {
		var ret []CountryCode
		return ret
	}
	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainerGroup) GetCountryCodesOk() ([]CountryCode, bool) {
	if o == nil || IsNil(o.CountryCodes) {
		return nil, false
	}
	return o.CountryCodes, true
}

// HasCountryCodes returns a boolean if a field has been set.
func (o *CreateContainerGroup) HasCountryCodes() bool {
	if o != nil && !IsNil(o.CountryCodes) {
		return true
	}

	return false
}

// SetCountryCodes gets a reference to the given []CountryCode and assigns it to the CountryCodes field.
func (o *CreateContainerGroup) SetCountryCodes(v []CountryCode) {
	o.CountryCodes = v
}

// GetNetworking returns the Networking field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContainerGroup) GetNetworking() CreateContainerGroupNetworking {
	if o == nil || IsNil(o.Networking.Get()) {
		var ret CreateContainerGroupNetworking
		return ret
	}
	return *o.Networking.Get()
}

// GetNetworkingOk returns a tuple with the Networking field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContainerGroup) GetNetworkingOk() (*CreateContainerGroupNetworking, bool) {
	if o == nil {
		return nil, false
	}
	return o.Networking.Get(), o.Networking.IsSet()
}

// HasNetworking returns a boolean if a field has been set.
func (o *CreateContainerGroup) HasNetworking() bool {
	if o != nil && o.Networking.IsSet() {
		return true
	}

	return false
}

// SetNetworking gets a reference to the given NullableCreateContainerGroupNetworking and assigns it to the Networking field.
func (o *CreateContainerGroup) SetNetworking(v CreateContainerGroupNetworking) {
	o.Networking.Set(&v)
}
// SetNetworkingNil sets the value for Networking to be an explicit nil
func (o *CreateContainerGroup) SetNetworkingNil() {
	o.Networking.Set(nil)
}

// UnsetNetworking ensures that no value is present for Networking, not even an explicit nil
func (o *CreateContainerGroup) UnsetNetworking() {
	o.Networking.Unset()
}

// GetLivenessProbe returns the LivenessProbe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContainerGroup) GetLivenessProbe() ContainerGroupProbe {
	if o == nil || IsNil(o.LivenessProbe.Get()) {
		var ret ContainerGroupProbe
		return ret
	}
	return *o.LivenessProbe.Get()
}

// GetLivenessProbeOk returns a tuple with the LivenessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContainerGroup) GetLivenessProbeOk() (*ContainerGroupProbe, bool) {
	if o == nil {
		return nil, false
	}
	return o.LivenessProbe.Get(), o.LivenessProbe.IsSet()
}

// HasLivenessProbe returns a boolean if a field has been set.
func (o *CreateContainerGroup) HasLivenessProbe() bool {
	if o != nil && o.LivenessProbe.IsSet() {
		return true
	}

	return false
}

// SetLivenessProbe gets a reference to the given NullableContainerGroupProbe and assigns it to the LivenessProbe field.
func (o *CreateContainerGroup) SetLivenessProbe(v ContainerGroupProbe) {
	o.LivenessProbe.Set(&v)
}
// SetLivenessProbeNil sets the value for LivenessProbe to be an explicit nil
func (o *CreateContainerGroup) SetLivenessProbeNil() {
	o.LivenessProbe.Set(nil)
}

// UnsetLivenessProbe ensures that no value is present for LivenessProbe, not even an explicit nil
func (o *CreateContainerGroup) UnsetLivenessProbe() {
	o.LivenessProbe.Unset()
}

// GetReadinessProbe returns the ReadinessProbe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContainerGroup) GetReadinessProbe() ContainerGroupProbe {
	if o == nil || IsNil(o.ReadinessProbe.Get()) {
		var ret ContainerGroupProbe
		return ret
	}
	return *o.ReadinessProbe.Get()
}

// GetReadinessProbeOk returns a tuple with the ReadinessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContainerGroup) GetReadinessProbeOk() (*ContainerGroupProbe, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReadinessProbe.Get(), o.ReadinessProbe.IsSet()
}

// HasReadinessProbe returns a boolean if a field has been set.
func (o *CreateContainerGroup) HasReadinessProbe() bool {
	if o != nil && o.ReadinessProbe.IsSet() {
		return true
	}

	return false
}

// SetReadinessProbe gets a reference to the given NullableContainerGroupProbe and assigns it to the ReadinessProbe field.
func (o *CreateContainerGroup) SetReadinessProbe(v ContainerGroupProbe) {
	o.ReadinessProbe.Set(&v)
}
// SetReadinessProbeNil sets the value for ReadinessProbe to be an explicit nil
func (o *CreateContainerGroup) SetReadinessProbeNil() {
	o.ReadinessProbe.Set(nil)
}

// UnsetReadinessProbe ensures that no value is present for ReadinessProbe, not even an explicit nil
func (o *CreateContainerGroup) UnsetReadinessProbe() {
	o.ReadinessProbe.Unset()
}

// GetStartupProbe returns the StartupProbe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContainerGroup) GetStartupProbe() ContainerGroupProbe {
	if o == nil || IsNil(o.StartupProbe.Get()) {
		var ret ContainerGroupProbe
		return ret
	}
	return *o.StartupProbe.Get()
}

// GetStartupProbeOk returns a tuple with the StartupProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContainerGroup) GetStartupProbeOk() (*ContainerGroupProbe, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartupProbe.Get(), o.StartupProbe.IsSet()
}

// HasStartupProbe returns a boolean if a field has been set.
func (o *CreateContainerGroup) HasStartupProbe() bool {
	if o != nil && o.StartupProbe.IsSet() {
		return true
	}

	return false
}

// SetStartupProbe gets a reference to the given NullableContainerGroupProbe and assigns it to the StartupProbe field.
func (o *CreateContainerGroup) SetStartupProbe(v ContainerGroupProbe) {
	o.StartupProbe.Set(&v)
}
// SetStartupProbeNil sets the value for StartupProbe to be an explicit nil
func (o *CreateContainerGroup) SetStartupProbeNil() {
	o.StartupProbe.Set(nil)
}

// UnsetStartupProbe ensures that no value is present for StartupProbe, not even an explicit nil
func (o *CreateContainerGroup) UnsetStartupProbe() {
	o.StartupProbe.Unset()
}

func (o CreateContainerGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateContainerGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.DisplayName.IsSet() {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	toSerialize["container"] = o.Container
	toSerialize["restart_policy"] = o.RestartPolicy
	toSerialize["replicas"] = o.Replicas
	if !IsNil(o.CountryCodes) {
		toSerialize["country_codes"] = o.CountryCodes
	}
	if o.Networking.IsSet() {
		toSerialize["networking"] = o.Networking.Get()
	}
	if o.LivenessProbe.IsSet() {
		toSerialize["liveness_probe"] = o.LivenessProbe.Get()
	}
	if o.ReadinessProbe.IsSet() {
		toSerialize["readiness_probe"] = o.ReadinessProbe.Get()
	}
	if o.StartupProbe.IsSet() {
		toSerialize["startup_probe"] = o.StartupProbe.Get()
	}
	return toSerialize, nil
}

type NullableCreateContainerGroup struct {
	value *CreateContainerGroup
	isSet bool
}

func (v NullableCreateContainerGroup) Get() *CreateContainerGroup {
	return v.value
}

func (v *NullableCreateContainerGroup) Set(val *CreateContainerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContainerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContainerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContainerGroup(val *CreateContainerGroup) *NullableCreateContainerGroup {
	return &NullableCreateContainerGroup{value: val, isSet: true}
}

func (v NullableCreateContainerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContainerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


