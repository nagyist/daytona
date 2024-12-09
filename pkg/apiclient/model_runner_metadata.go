/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the RunnerMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunnerMetadata{}

// RunnerMetadata struct for RunnerMetadata
type RunnerMetadata struct {
	Providers   []TargetProviderInfo `json:"providers"`
	RunnerId    string               `json:"runnerId"`
	RunningJobs int32                `json:"runningJobs"`
	UpdatedAt   string               `json:"updatedAt"`
	Uptime      int32                `json:"uptime"`
}

type _RunnerMetadata RunnerMetadata

// NewRunnerMetadata instantiates a new RunnerMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunnerMetadata(providers []TargetProviderInfo, runnerId string, runningJobs int32, updatedAt string, uptime int32) *RunnerMetadata {
	this := RunnerMetadata{}
	this.Providers = providers
	this.RunnerId = runnerId
	this.RunningJobs = runningJobs
	this.UpdatedAt = updatedAt
	this.Uptime = uptime
	return &this
}

// NewRunnerMetadataWithDefaults instantiates a new RunnerMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunnerMetadataWithDefaults() *RunnerMetadata {
	this := RunnerMetadata{}
	return &this
}

// GetProviders returns the Providers field value
func (o *RunnerMetadata) GetProviders() []TargetProviderInfo {
	if o == nil {
		var ret []TargetProviderInfo
		return ret
	}

	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value
// and a boolean to check if the value has been set.
func (o *RunnerMetadata) GetProvidersOk() ([]TargetProviderInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Providers, true
}

// SetProviders sets field value
func (o *RunnerMetadata) SetProviders(v []TargetProviderInfo) {
	o.Providers = v
}

// GetRunnerId returns the RunnerId field value
func (o *RunnerMetadata) GetRunnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RunnerId
}

// GetRunnerIdOk returns a tuple with the RunnerId field value
// and a boolean to check if the value has been set.
func (o *RunnerMetadata) GetRunnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunnerId, true
}

// SetRunnerId sets field value
func (o *RunnerMetadata) SetRunnerId(v string) {
	o.RunnerId = v
}

// GetRunningJobs returns the RunningJobs field value
func (o *RunnerMetadata) GetRunningJobs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RunningJobs
}

// GetRunningJobsOk returns a tuple with the RunningJobs field value
// and a boolean to check if the value has been set.
func (o *RunnerMetadata) GetRunningJobsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunningJobs, true
}

// SetRunningJobs sets field value
func (o *RunnerMetadata) SetRunningJobs(v int32) {
	o.RunningJobs = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *RunnerMetadata) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *RunnerMetadata) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *RunnerMetadata) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetUptime returns the Uptime field value
func (o *RunnerMetadata) GetUptime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value
// and a boolean to check if the value has been set.
func (o *RunnerMetadata) GetUptimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uptime, true
}

// SetUptime sets field value
func (o *RunnerMetadata) SetUptime(v int32) {
	o.Uptime = v
}

func (o RunnerMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunnerMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["providers"] = o.Providers
	toSerialize["runnerId"] = o.RunnerId
	toSerialize["runningJobs"] = o.RunningJobs
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["uptime"] = o.Uptime
	return toSerialize, nil
}

func (o *RunnerMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"providers",
		"runnerId",
		"runningJobs",
		"updatedAt",
		"uptime",
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

	varRunnerMetadata := _RunnerMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunnerMetadata)

	if err != nil {
		return err
	}

	*o = RunnerMetadata(varRunnerMetadata)

	return err
}

type NullableRunnerMetadata struct {
	value *RunnerMetadata
	isSet bool
}

func (v NullableRunnerMetadata) Get() *RunnerMetadata {
	return v.value
}

func (v *NullableRunnerMetadata) Set(val *RunnerMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableRunnerMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableRunnerMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunnerMetadata(val *RunnerMetadata) *NullableRunnerMetadata {
	return &NullableRunnerMetadata{value: val, isSet: true}
}

func (v NullableRunnerMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunnerMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
