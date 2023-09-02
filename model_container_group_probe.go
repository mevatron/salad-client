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

// checks if the ContainerGroupProbe type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerGroupProbe{}

// ContainerGroupProbe Represents container group probe
type ContainerGroupProbe struct {
	Http *ContainerGroupProbeHttp `json:"http,omitempty"`
	Tcp *ContainerGroupProbeTcp `json:"tcp,omitempty"`
	Grpc *ContainerGroupProbeGrpc `json:"grpc,omitempty"`
	Exec *ContainerGroupProbeExec `json:"exec,omitempty"`
	InitialDelaySeconds int32 `json:"initial_delay_seconds"`
	PeriodSeconds int32 `json:"period_seconds"`
	TimeoutSeconds int32 `json:"timeout_seconds"`
	SuccessThreshold int32 `json:"success_threshold"`
	FailureThreshold int32 `json:"failure_threshold"`
}

// NewContainerGroupProbe instantiates a new ContainerGroupProbe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerGroupProbe(initialDelaySeconds int32, periodSeconds int32, timeoutSeconds int32, successThreshold int32, failureThreshold int32) *ContainerGroupProbe {
	this := ContainerGroupProbe{}
	this.InitialDelaySeconds = initialDelaySeconds
	this.PeriodSeconds = periodSeconds
	this.TimeoutSeconds = timeoutSeconds
	this.SuccessThreshold = successThreshold
	this.FailureThreshold = failureThreshold
	return &this
}

// NewContainerGroupProbeWithDefaults instantiates a new ContainerGroupProbe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerGroupProbeWithDefaults() *ContainerGroupProbe {
	this := ContainerGroupProbe{}
	return &this
}

// GetHttp returns the Http field value if set, zero value otherwise.
func (o *ContainerGroupProbe) GetHttp() ContainerGroupProbeHttp {
	if o == nil || IsNil(o.Http) {
		var ret ContainerGroupProbeHttp
		return ret
	}
	return *o.Http
}

// GetHttpOk returns a tuple with the Http field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbe) GetHttpOk() (*ContainerGroupProbeHttp, bool) {
	if o == nil || IsNil(o.Http) {
		return nil, false
	}
	return o.Http, true
}

// HasHttp returns a boolean if a field has been set.
func (o *ContainerGroupProbe) HasHttp() bool {
	if o != nil && !IsNil(o.Http) {
		return true
	}

	return false
}

// SetHttp gets a reference to the given ContainerGroupProbeHttp and assigns it to the Http field.
func (o *ContainerGroupProbe) SetHttp(v ContainerGroupProbeHttp) {
	o.Http = &v
}

// GetTcp returns the Tcp field value if set, zero value otherwise.
func (o *ContainerGroupProbe) GetTcp() ContainerGroupProbeTcp {
	if o == nil || IsNil(o.Tcp) {
		var ret ContainerGroupProbeTcp
		return ret
	}
	return *o.Tcp
}

// GetTcpOk returns a tuple with the Tcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbe) GetTcpOk() (*ContainerGroupProbeTcp, bool) {
	if o == nil || IsNil(o.Tcp) {
		return nil, false
	}
	return o.Tcp, true
}

// HasTcp returns a boolean if a field has been set.
func (o *ContainerGroupProbe) HasTcp() bool {
	if o != nil && !IsNil(o.Tcp) {
		return true
	}

	return false
}

// SetTcp gets a reference to the given ContainerGroupProbeTcp and assigns it to the Tcp field.
func (o *ContainerGroupProbe) SetTcp(v ContainerGroupProbeTcp) {
	o.Tcp = &v
}

// GetGrpc returns the Grpc field value if set, zero value otherwise.
func (o *ContainerGroupProbe) GetGrpc() ContainerGroupProbeGrpc {
	if o == nil || IsNil(o.Grpc) {
		var ret ContainerGroupProbeGrpc
		return ret
	}
	return *o.Grpc
}

// GetGrpcOk returns a tuple with the Grpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbe) GetGrpcOk() (*ContainerGroupProbeGrpc, bool) {
	if o == nil || IsNil(o.Grpc) {
		return nil, false
	}
	return o.Grpc, true
}

// HasGrpc returns a boolean if a field has been set.
func (o *ContainerGroupProbe) HasGrpc() bool {
	if o != nil && !IsNil(o.Grpc) {
		return true
	}

	return false
}

// SetGrpc gets a reference to the given ContainerGroupProbeGrpc and assigns it to the Grpc field.
func (o *ContainerGroupProbe) SetGrpc(v ContainerGroupProbeGrpc) {
	o.Grpc = &v
}

// GetExec returns the Exec field value if set, zero value otherwise.
func (o *ContainerGroupProbe) GetExec() ContainerGroupProbeExec {
	if o == nil || IsNil(o.Exec) {
		var ret ContainerGroupProbeExec
		return ret
	}
	return *o.Exec
}

// GetExecOk returns a tuple with the Exec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbe) GetExecOk() (*ContainerGroupProbeExec, bool) {
	if o == nil || IsNil(o.Exec) {
		return nil, false
	}
	return o.Exec, true
}

// HasExec returns a boolean if a field has been set.
func (o *ContainerGroupProbe) HasExec() bool {
	if o != nil && !IsNil(o.Exec) {
		return true
	}

	return false
}

// SetExec gets a reference to the given ContainerGroupProbeExec and assigns it to the Exec field.
func (o *ContainerGroupProbe) SetExec(v ContainerGroupProbeExec) {
	o.Exec = &v
}

// GetInitialDelaySeconds returns the InitialDelaySeconds field value
func (o *ContainerGroupProbe) GetInitialDelaySeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InitialDelaySeconds
}

// GetInitialDelaySecondsOk returns a tuple with the InitialDelaySeconds field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbe) GetInitialDelaySecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialDelaySeconds, true
}

// SetInitialDelaySeconds sets field value
func (o *ContainerGroupProbe) SetInitialDelaySeconds(v int32) {
	o.InitialDelaySeconds = v
}

// GetPeriodSeconds returns the PeriodSeconds field value
func (o *ContainerGroupProbe) GetPeriodSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PeriodSeconds
}

// GetPeriodSecondsOk returns a tuple with the PeriodSeconds field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbe) GetPeriodSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodSeconds, true
}

// SetPeriodSeconds sets field value
func (o *ContainerGroupProbe) SetPeriodSeconds(v int32) {
	o.PeriodSeconds = v
}

// GetTimeoutSeconds returns the TimeoutSeconds field value
func (o *ContainerGroupProbe) GetTimeoutSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbe) GetTimeoutSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeoutSeconds, true
}

// SetTimeoutSeconds sets field value
func (o *ContainerGroupProbe) SetTimeoutSeconds(v int32) {
	o.TimeoutSeconds = v
}

// GetSuccessThreshold returns the SuccessThreshold field value
func (o *ContainerGroupProbe) GetSuccessThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SuccessThreshold
}

// GetSuccessThresholdOk returns a tuple with the SuccessThreshold field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbe) GetSuccessThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessThreshold, true
}

// SetSuccessThreshold sets field value
func (o *ContainerGroupProbe) SetSuccessThreshold(v int32) {
	o.SuccessThreshold = v
}

// GetFailureThreshold returns the FailureThreshold field value
func (o *ContainerGroupProbe) GetFailureThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FailureThreshold
}

// GetFailureThresholdOk returns a tuple with the FailureThreshold field value
// and a boolean to check if the value has been set.
func (o *ContainerGroupProbe) GetFailureThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailureThreshold, true
}

// SetFailureThreshold sets field value
func (o *ContainerGroupProbe) SetFailureThreshold(v int32) {
	o.FailureThreshold = v
}

func (o ContainerGroupProbe) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerGroupProbe) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Http) {
		toSerialize["http"] = o.Http
	}
	if !IsNil(o.Tcp) {
		toSerialize["tcp"] = o.Tcp
	}
	if !IsNil(o.Grpc) {
		toSerialize["grpc"] = o.Grpc
	}
	if !IsNil(o.Exec) {
		toSerialize["exec"] = o.Exec
	}
	toSerialize["initial_delay_seconds"] = o.InitialDelaySeconds
	toSerialize["period_seconds"] = o.PeriodSeconds
	toSerialize["timeout_seconds"] = o.TimeoutSeconds
	toSerialize["success_threshold"] = o.SuccessThreshold
	toSerialize["failure_threshold"] = o.FailureThreshold
	return toSerialize, nil
}

type NullableContainerGroupProbe struct {
	value *ContainerGroupProbe
	isSet bool
}

func (v NullableContainerGroupProbe) Get() *ContainerGroupProbe {
	return v.value
}

func (v *NullableContainerGroupProbe) Set(val *ContainerGroupProbe) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerGroupProbe) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerGroupProbe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerGroupProbe(val *ContainerGroupProbe) *NullableContainerGroupProbe {
	return &NullableContainerGroupProbe{value: val, isSet: true}
}

func (v NullableContainerGroupProbe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerGroupProbe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


