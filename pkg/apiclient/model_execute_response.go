/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the ExecuteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecuteResponse{}

// ExecuteResponse struct for ExecuteResponse
type ExecuteResponse struct {
	Code   *int32  `json:"code,omitempty"`
	Result *string `json:"result,omitempty"`
}

// NewExecuteResponse instantiates a new ExecuteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecuteResponse() *ExecuteResponse {
	this := ExecuteResponse{}
	return &this
}

// NewExecuteResponseWithDefaults instantiates a new ExecuteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecuteResponseWithDefaults() *ExecuteResponse {
	this := ExecuteResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ExecuteResponse) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteResponse) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ExecuteResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ExecuteResponse) SetCode(v int32) {
	o.Code = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ExecuteResponse) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteResponse) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ExecuteResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *ExecuteResponse) SetResult(v string) {
	o.Result = &v
}

func (o ExecuteResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecuteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableExecuteResponse struct {
	value *ExecuteResponse
	isSet bool
}

func (v NullableExecuteResponse) Get() *ExecuteResponse {
	return v.value
}

func (v *NullableExecuteResponse) Set(val *ExecuteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteResponse(val *ExecuteResponse) *NullableExecuteResponse {
	return &NullableExecuteResponse{value: val, isSet: true}
}

func (v NullableExecuteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
