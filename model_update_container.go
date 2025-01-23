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
)

// checks if the UpdateContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateContainer{}

// UpdateContainer Represents an update container object
type UpdateContainer struct {
	Image     NullableString                   `json:"image,omitempty"`
	Resources NullableUpdateContainerResources `json:"resources,omitempty"`
	// Pass a command (and optional arguments) to override the ENTRYPOINT and CMD of a container image.
	Command                []string                                      `json:"command,omitempty"`
	Priority               NullableContainerGroupPriority                `json:"priority,omitempty"`
	EnvironmentVariables   *map[string]string                            `json:"environment_variables,omitempty"`
	Logging                NullableContainerLogging                      `json:"logging,omitempty"`
	RegistryAuthentication NullableCreateContainerRegistryAuthentication `json:"registry_authentication,omitempty"`
}

// NewUpdateContainer instantiates a new UpdateContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateContainer() *UpdateContainer {
	this := UpdateContainer{}
	return &this
}

// NewUpdateContainerWithDefaults instantiates a new UpdateContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateContainerWithDefaults() *UpdateContainer {
	this := UpdateContainer{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateContainer) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateContainer) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *UpdateContainer) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *UpdateContainer) SetImage(v string) {
	o.Image.Set(&v)
}

// SetImageNil sets the value for Image to be an explicit nil
func (o *UpdateContainer) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *UpdateContainer) UnsetImage() {
	o.Image.Unset()
}

// GetResources returns the Resources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateContainer) GetResources() UpdateContainerResources {
	if o == nil || IsNil(o.Resources.Get()) {
		var ret UpdateContainerResources
		return ret
	}
	return *o.Resources.Get()
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateContainer) GetResourcesOk() (*UpdateContainerResources, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources.Get(), o.Resources.IsSet()
}

// HasResources returns a boolean if a field has been set.
func (o *UpdateContainer) HasResources() bool {
	if o != nil && o.Resources.IsSet() {
		return true
	}

	return false
}

// SetResources gets a reference to the given NullableUpdateContainerResources and assigns it to the Resources field.
func (o *UpdateContainer) SetResources(v UpdateContainerResources) {
	o.Resources.Set(&v)
}

// SetResourcesNil sets the value for Resources to be an explicit nil
func (o *UpdateContainer) SetResourcesNil() {
	o.Resources.Set(nil)
}

// UnsetResources ensures that no value is present for Resources, not even an explicit nil
func (o *UpdateContainer) UnsetResources() {
	o.Resources.Unset()
}

// GetCommand returns the Command field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateContainer) GetCommand() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateContainer) GetCommandOk() ([]string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *UpdateContainer) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given []string and assigns it to the Command field.
func (o *UpdateContainer) SetCommand(v []string) {
	o.Command = v
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateContainer) GetPriority() ContainerGroupPriority {
	if o == nil || IsNil(o.Priority.Get()) {
		var ret ContainerGroupPriority
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateContainer) GetPriorityOk() (*ContainerGroupPriority, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *UpdateContainer) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableContainerGroupPriority and assigns it to the Priority field.
func (o *UpdateContainer) SetPriority(v ContainerGroupPriority) {
	o.Priority.Set(&v)
}

// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *UpdateContainer) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *UpdateContainer) UnsetPriority() {
	o.Priority.Unset()
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise.
func (o *UpdateContainer) GetEnvironmentVariables() map[string]string {
	if o == nil || IsNil(o.EnvironmentVariables) {
		var ret map[string]string
		return ret
	}
	return *o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContainer) GetEnvironmentVariablesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.EnvironmentVariables) {
		return nil, false
	}
	return o.EnvironmentVariables, true
}

// HasEnvironmentVariables returns a boolean if a field has been set.
func (o *UpdateContainer) HasEnvironmentVariables() bool {
	if o != nil && !IsNil(o.EnvironmentVariables) {
		return true
	}

	return false
}

// SetEnvironmentVariables gets a reference to the given map[string]string and assigns it to the EnvironmentVariables field.
func (o *UpdateContainer) SetEnvironmentVariables(v map[string]string) {
	o.EnvironmentVariables = &v
}

// GetLogging returns the Logging field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateContainer) GetLogging() ContainerLogging {
	if o == nil || IsNil(o.Logging.Get()) {
		var ret ContainerLogging
		return ret
	}
	return *o.Logging.Get()
}

// GetLoggingOk returns a tuple with the Logging field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateContainer) GetLoggingOk() (*ContainerLogging, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logging.Get(), o.Logging.IsSet()
}

// HasLogging returns a boolean if a field has been set.
func (o *UpdateContainer) HasLogging() bool {
	if o != nil && o.Logging.IsSet() {
		return true
	}

	return false
}

// SetLogging gets a reference to the given NullableContainerLogging and assigns it to the Logging field.
func (o *UpdateContainer) SetLogging(v ContainerLogging) {
	o.Logging.Set(&v)
}

// SetLoggingNil sets the value for Logging to be an explicit nil
func (o *UpdateContainer) SetLoggingNil() {
	o.Logging.Set(nil)
}

// UnsetLogging ensures that no value is present for Logging, not even an explicit nil
func (o *UpdateContainer) UnsetLogging() {
	o.Logging.Unset()
}

// GetRegistryAuthentication returns the RegistryAuthentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateContainer) GetRegistryAuthentication() CreateContainerRegistryAuthentication {
	if o == nil || IsNil(o.RegistryAuthentication.Get()) {
		var ret CreateContainerRegistryAuthentication
		return ret
	}
	return *o.RegistryAuthentication.Get()
}

// GetRegistryAuthenticationOk returns a tuple with the RegistryAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateContainer) GetRegistryAuthenticationOk() (*CreateContainerRegistryAuthentication, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegistryAuthentication.Get(), o.RegistryAuthentication.IsSet()
}

// HasRegistryAuthentication returns a boolean if a field has been set.
func (o *UpdateContainer) HasRegistryAuthentication() bool {
	if o != nil && o.RegistryAuthentication.IsSet() {
		return true
	}

	return false
}

// SetRegistryAuthentication gets a reference to the given NullableCreateContainerRegistryAuthentication and assigns it to the RegistryAuthentication field.
func (o *UpdateContainer) SetRegistryAuthentication(v CreateContainerRegistryAuthentication) {
	o.RegistryAuthentication.Set(&v)
}

// SetRegistryAuthenticationNil sets the value for RegistryAuthentication to be an explicit nil
func (o *UpdateContainer) SetRegistryAuthenticationNil() {
	o.RegistryAuthentication.Set(nil)
}

// UnsetRegistryAuthentication ensures that no value is present for RegistryAuthentication, not even an explicit nil
func (o *UpdateContainer) UnsetRegistryAuthentication() {
	o.RegistryAuthentication.Unset()
}

func (o UpdateContainer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if o.Resources.IsSet() {
		toSerialize["resources"] = o.Resources.Get()
	}
	if o.Command != nil {
		toSerialize["command"] = o.Command
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	if !IsNil(o.EnvironmentVariables) {
		toSerialize["environment_variables"] = o.EnvironmentVariables
	}
	if o.Logging.IsSet() {
		toSerialize["logging"] = o.Logging.Get()
	}
	if o.RegistryAuthentication.IsSet() {
		toSerialize["registry_authentication"] = o.RegistryAuthentication.Get()
	}
	return toSerialize, nil
}

type NullableUpdateContainer struct {
	value *UpdateContainer
	isSet bool
}

func (v NullableUpdateContainer) Get() *UpdateContainer {
	return v.value
}

func (v *NullableUpdateContainer) Set(val *UpdateContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateContainer(val *UpdateContainer) *NullableUpdateContainer {
	return &NullableUpdateContainer{value: val, isSet: true}
}

func (v NullableUpdateContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
