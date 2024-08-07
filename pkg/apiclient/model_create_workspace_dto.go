/*
Daytona Server API

Daytona Server API

API version: 0.24.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CreateWorkspaceDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWorkspaceDTO{}

// CreateWorkspaceDTO struct for CreateWorkspaceDTO
type CreateWorkspaceDTO struct {
	Id       *string                  `json:"id,omitempty"`
	Name     *string                  `json:"name,omitempty"`
	Projects []CreateProjectConfigDTO `json:"projects"`
	Target   *string                  `json:"target,omitempty"`
}

type _CreateWorkspaceDTO CreateWorkspaceDTO

// NewCreateWorkspaceDTO instantiates a new CreateWorkspaceDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkspaceDTO(projects []CreateProjectConfigDTO) *CreateWorkspaceDTO {
	this := CreateWorkspaceDTO{}
	this.Projects = projects
	return &this
}

// NewCreateWorkspaceDTOWithDefaults instantiates a new CreateWorkspaceDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkspaceDTOWithDefaults() *CreateWorkspaceDTO {
	this := CreateWorkspaceDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateWorkspaceDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateWorkspaceDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateWorkspaceDTO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateWorkspaceDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateWorkspaceDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateWorkspaceDTO) SetName(v string) {
	o.Name = &v
}

// GetProjects returns the Projects field value
func (o *CreateWorkspaceDTO) GetProjects() []CreateProjectConfigDTO {
	if o == nil {
		var ret []CreateProjectConfigDTO
		return ret
	}

	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceDTO) GetProjectsOk() ([]CreateProjectConfigDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Projects, true
}

// SetProjects sets field value
func (o *CreateWorkspaceDTO) SetProjects(v []CreateProjectConfigDTO) {
	o.Projects = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *CreateWorkspaceDTO) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceDTO) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *CreateWorkspaceDTO) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *CreateWorkspaceDTO) SetTarget(v string) {
	o.Target = &v
}

func (o CreateWorkspaceDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWorkspaceDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["projects"] = o.Projects
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return toSerialize, nil
}

func (o *CreateWorkspaceDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"projects",
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

	varCreateWorkspaceDTO := _CreateWorkspaceDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateWorkspaceDTO)

	if err != nil {
		return err
	}

	*o = CreateWorkspaceDTO(varCreateWorkspaceDTO)

	return err
}

type NullableCreateWorkspaceDTO struct {
	value *CreateWorkspaceDTO
	isSet bool
}

func (v NullableCreateWorkspaceDTO) Get() *CreateWorkspaceDTO {
	return v.value
}

func (v *NullableCreateWorkspaceDTO) Set(val *CreateWorkspaceDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkspaceDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkspaceDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkspaceDTO(val *CreateWorkspaceDTO) *NullableCreateWorkspaceDTO {
	return &NullableCreateWorkspaceDTO{value: val, isSet: true}
}

func (v NullableCreateWorkspaceDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkspaceDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
