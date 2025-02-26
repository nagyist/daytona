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

// checks if the WorkspaceTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkspaceTemplate{}

// WorkspaceTemplate struct for WorkspaceTemplate
type WorkspaceTemplate struct {
	BuildConfig         *BuildConfig      `json:"buildConfig,omitempty"`
	Default             bool              `json:"default"`
	EnvVars             map[string]string `json:"envVars"`
	GitProviderConfigId *string           `json:"gitProviderConfigId,omitempty"`
	Image               string            `json:"image"`
	Labels              map[string]string `json:"labels"`
	Name                string            `json:"name"`
	Prebuilds           []PrebuildConfig  `json:"prebuilds,omitempty"`
	RepositoryUrl       string            `json:"repositoryUrl"`
	User                string            `json:"user"`
}

type _WorkspaceTemplate WorkspaceTemplate

// NewWorkspaceTemplate instantiates a new WorkspaceTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceTemplate(default_ bool, envVars map[string]string, image string, labels map[string]string, name string, repositoryUrl string, user string) *WorkspaceTemplate {
	this := WorkspaceTemplate{}
	this.Default = default_
	this.EnvVars = envVars
	this.Image = image
	this.Labels = labels
	this.Name = name
	this.RepositoryUrl = repositoryUrl
	this.User = user
	return &this
}

// NewWorkspaceTemplateWithDefaults instantiates a new WorkspaceTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceTemplateWithDefaults() *WorkspaceTemplate {
	this := WorkspaceTemplate{}
	return &this
}

// GetBuildConfig returns the BuildConfig field value if set, zero value otherwise.
func (o *WorkspaceTemplate) GetBuildConfig() BuildConfig {
	if o == nil || IsNil(o.BuildConfig) {
		var ret BuildConfig
		return ret
	}
	return *o.BuildConfig
}

// GetBuildConfigOk returns a tuple with the BuildConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceTemplate) GetBuildConfigOk() (*BuildConfig, bool) {
	if o == nil || IsNil(o.BuildConfig) {
		return nil, false
	}
	return o.BuildConfig, true
}

// HasBuildConfig returns a boolean if a field has been set.
func (o *WorkspaceTemplate) HasBuildConfig() bool {
	if o != nil && !IsNil(o.BuildConfig) {
		return true
	}

	return false
}

// SetBuildConfig gets a reference to the given BuildConfig and assigns it to the BuildConfig field.
func (o *WorkspaceTemplate) SetBuildConfig(v BuildConfig) {
	o.BuildConfig = &v
}

// GetDefault returns the Default field value
func (o *WorkspaceTemplate) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *WorkspaceTemplate) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *WorkspaceTemplate) SetDefault(v bool) {
	o.Default = v
}

// GetEnvVars returns the EnvVars field value
func (o *WorkspaceTemplate) GetEnvVars() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.EnvVars
}

// GetEnvVarsOk returns a tuple with the EnvVars field value
// and a boolean to check if the value has been set.
func (o *WorkspaceTemplate) GetEnvVarsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvVars, true
}

// SetEnvVars sets field value
func (o *WorkspaceTemplate) SetEnvVars(v map[string]string) {
	o.EnvVars = v
}

// GetGitProviderConfigId returns the GitProviderConfigId field value if set, zero value otherwise.
func (o *WorkspaceTemplate) GetGitProviderConfigId() string {
	if o == nil || IsNil(o.GitProviderConfigId) {
		var ret string
		return ret
	}
	return *o.GitProviderConfigId
}

// GetGitProviderConfigIdOk returns a tuple with the GitProviderConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceTemplate) GetGitProviderConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.GitProviderConfigId) {
		return nil, false
	}
	return o.GitProviderConfigId, true
}

// HasGitProviderConfigId returns a boolean if a field has been set.
func (o *WorkspaceTemplate) HasGitProviderConfigId() bool {
	if o != nil && !IsNil(o.GitProviderConfigId) {
		return true
	}

	return false
}

// SetGitProviderConfigId gets a reference to the given string and assigns it to the GitProviderConfigId field.
func (o *WorkspaceTemplate) SetGitProviderConfigId(v string) {
	o.GitProviderConfigId = &v
}

// GetImage returns the Image field value
func (o *WorkspaceTemplate) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *WorkspaceTemplate) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *WorkspaceTemplate) SetImage(v string) {
	o.Image = v
}

// GetLabels returns the Labels field value
func (o *WorkspaceTemplate) GetLabels() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *WorkspaceTemplate) GetLabelsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Labels, true
}

// SetLabels sets field value
func (o *WorkspaceTemplate) SetLabels(v map[string]string) {
	o.Labels = v
}

// GetName returns the Name field value
func (o *WorkspaceTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkspaceTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkspaceTemplate) SetName(v string) {
	o.Name = v
}

// GetPrebuilds returns the Prebuilds field value if set, zero value otherwise.
func (o *WorkspaceTemplate) GetPrebuilds() []PrebuildConfig {
	if o == nil || IsNil(o.Prebuilds) {
		var ret []PrebuildConfig
		return ret
	}
	return o.Prebuilds
}

// GetPrebuildsOk returns a tuple with the Prebuilds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceTemplate) GetPrebuildsOk() ([]PrebuildConfig, bool) {
	if o == nil || IsNil(o.Prebuilds) {
		return nil, false
	}
	return o.Prebuilds, true
}

// HasPrebuilds returns a boolean if a field has been set.
func (o *WorkspaceTemplate) HasPrebuilds() bool {
	if o != nil && !IsNil(o.Prebuilds) {
		return true
	}

	return false
}

// SetPrebuilds gets a reference to the given []PrebuildConfig and assigns it to the Prebuilds field.
func (o *WorkspaceTemplate) SetPrebuilds(v []PrebuildConfig) {
	o.Prebuilds = v
}

// GetRepositoryUrl returns the RepositoryUrl field value
func (o *WorkspaceTemplate) GetRepositoryUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value
// and a boolean to check if the value has been set.
func (o *WorkspaceTemplate) GetRepositoryUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryUrl, true
}

// SetRepositoryUrl sets field value
func (o *WorkspaceTemplate) SetRepositoryUrl(v string) {
	o.RepositoryUrl = v
}

// GetUser returns the User field value
func (o *WorkspaceTemplate) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *WorkspaceTemplate) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *WorkspaceTemplate) SetUser(v string) {
	o.User = v
}

func (o WorkspaceTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkspaceTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuildConfig) {
		toSerialize["buildConfig"] = o.BuildConfig
	}
	toSerialize["default"] = o.Default
	toSerialize["envVars"] = o.EnvVars
	if !IsNil(o.GitProviderConfigId) {
		toSerialize["gitProviderConfigId"] = o.GitProviderConfigId
	}
	toSerialize["image"] = o.Image
	toSerialize["labels"] = o.Labels
	toSerialize["name"] = o.Name
	if !IsNil(o.Prebuilds) {
		toSerialize["prebuilds"] = o.Prebuilds
	}
	toSerialize["repositoryUrl"] = o.RepositoryUrl
	toSerialize["user"] = o.User
	return toSerialize, nil
}

func (o *WorkspaceTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"default",
		"envVars",
		"image",
		"labels",
		"name",
		"repositoryUrl",
		"user",
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

	varWorkspaceTemplate := _WorkspaceTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkspaceTemplate)

	if err != nil {
		return err
	}

	*o = WorkspaceTemplate(varWorkspaceTemplate)

	return err
}

type NullableWorkspaceTemplate struct {
	value *WorkspaceTemplate
	isSet bool
}

func (v NullableWorkspaceTemplate) Get() *WorkspaceTemplate {
	return v.value
}

func (v *NullableWorkspaceTemplate) Set(val *WorkspaceTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceTemplate(val *WorkspaceTemplate) *NullableWorkspaceTemplate {
	return &NullableWorkspaceTemplate{value: val, isSet: true}
}

func (v NullableWorkspaceTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
