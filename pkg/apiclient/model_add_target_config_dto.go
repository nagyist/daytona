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

// checks if the AddTargetConfigDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddTargetConfigDTO{}

// AddTargetConfigDTO struct for AddTargetConfigDTO
type AddTargetConfigDTO struct {
	Name         string       `json:"name"`
	Options      string       `json:"options"`
	ProviderInfo ProviderInfo `json:"providerInfo"`
}

type _AddTargetConfigDTO AddTargetConfigDTO

// NewAddTargetConfigDTO instantiates a new AddTargetConfigDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddTargetConfigDTO(name string, options string, providerInfo ProviderInfo) *AddTargetConfigDTO {
	this := AddTargetConfigDTO{}
	this.Name = name
	this.Options = options
	this.ProviderInfo = providerInfo
	return &this
}

// NewAddTargetConfigDTOWithDefaults instantiates a new AddTargetConfigDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddTargetConfigDTOWithDefaults() *AddTargetConfigDTO {
	this := AddTargetConfigDTO{}
	return &this
}

// GetName returns the Name field value
func (o *AddTargetConfigDTO) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddTargetConfigDTO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddTargetConfigDTO) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *AddTargetConfigDTO) GetOptions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *AddTargetConfigDTO) GetOptionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *AddTargetConfigDTO) SetOptions(v string) {
	o.Options = v
}

// GetProviderInfo returns the ProviderInfo field value
func (o *AddTargetConfigDTO) GetProviderInfo() ProviderInfo {
	if o == nil {
		var ret ProviderInfo
		return ret
	}

	return o.ProviderInfo
}

// GetProviderInfoOk returns a tuple with the ProviderInfo field value
// and a boolean to check if the value has been set.
func (o *AddTargetConfigDTO) GetProviderInfoOk() (*ProviderInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderInfo, true
}

// SetProviderInfo sets field value
func (o *AddTargetConfigDTO) SetProviderInfo(v ProviderInfo) {
	o.ProviderInfo = v
}

func (o AddTargetConfigDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddTargetConfigDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["options"] = o.Options
	toSerialize["providerInfo"] = o.ProviderInfo
	return toSerialize, nil
}

func (o *AddTargetConfigDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"options",
		"providerInfo",
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

	varAddTargetConfigDTO := _AddTargetConfigDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddTargetConfigDTO)

	if err != nil {
		return err
	}

	*o = AddTargetConfigDTO(varAddTargetConfigDTO)

	return err
}

type NullableAddTargetConfigDTO struct {
	value *AddTargetConfigDTO
	isSet bool
}

func (v NullableAddTargetConfigDTO) Get() *AddTargetConfigDTO {
	return v.value
}

func (v *NullableAddTargetConfigDTO) Set(val *AddTargetConfigDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAddTargetConfigDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAddTargetConfigDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddTargetConfigDTO(val *AddTargetConfigDTO) *NullableAddTargetConfigDTO {
	return &NullableAddTargetConfigDTO{value: val, isSet: true}
}

func (v NullableAddTargetConfigDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddTargetConfigDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
