/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
	"fmt"
)

// checks if the SubjectExportSharedViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubjectExportSharedViewResponse{}

// SubjectExportSharedViewResponse struct for SubjectExportSharedViewResponse
type SubjectExportSharedViewResponse struct {
	ExportId    string `json:"export_id"`
	JwtSettings Import `json:"jwt_settings"`
}

type _SubjectExportSharedViewResponse SubjectExportSharedViewResponse

// NewSubjectExportSharedViewResponse instantiates a new SubjectExportSharedViewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubjectExportSharedViewResponse(exportId string, jwtSettings Import) *SubjectExportSharedViewResponse {
	this := SubjectExportSharedViewResponse{}
	this.ExportId = exportId
	this.JwtSettings = jwtSettings
	return &this
}

// NewSubjectExportSharedViewResponseWithDefaults instantiates a new SubjectExportSharedViewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectExportSharedViewResponseWithDefaults() *SubjectExportSharedViewResponse {
	this := SubjectExportSharedViewResponse{}
	return &this
}

// GetExportId returns the ExportId field value
func (o *SubjectExportSharedViewResponse) GetExportId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExportId
}

// GetExportIdOk returns a tuple with the ExportId field value
// and a boolean to check if the value has been set.
func (o *SubjectExportSharedViewResponse) GetExportIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportId, true
}

// SetExportId sets field value
func (o *SubjectExportSharedViewResponse) SetExportId(v string) {
	o.ExportId = v
}

// GetJwtSettings returns the JwtSettings field value
func (o *SubjectExportSharedViewResponse) GetJwtSettings() Import {
	if o == nil {
		var ret Import
		return ret
	}

	return o.JwtSettings
}

// GetJwtSettingsOk returns a tuple with the JwtSettings field value
// and a boolean to check if the value has been set.
func (o *SubjectExportSharedViewResponse) GetJwtSettingsOk() (*Import, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JwtSettings, true
}

// SetJwtSettings sets field value
func (o *SubjectExportSharedViewResponse) SetJwtSettings(v Import) {
	o.JwtSettings = v
}

func (o SubjectExportSharedViewResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubjectExportSharedViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["export_id"] = o.ExportId
	toSerialize["jwt_settings"] = o.JwtSettings
	return toSerialize, nil
}

func (o *SubjectExportSharedViewResponse) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"export_id",
		"jwt_settings",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSubjectExportSharedViewResponse := _SubjectExportSharedViewResponse{}

	err = json.Unmarshal(bytes, &varSubjectExportSharedViewResponse)

	if err != nil {
		return err
	}

	*o = SubjectExportSharedViewResponse(varSubjectExportSharedViewResponse)

	return err
}

type NullableSubjectExportSharedViewResponse struct {
	value *SubjectExportSharedViewResponse
	isSet bool
}

func (v NullableSubjectExportSharedViewResponse) Get() *SubjectExportSharedViewResponse {
	return v.value
}

func (v *NullableSubjectExportSharedViewResponse) Set(val *SubjectExportSharedViewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectExportSharedViewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectExportSharedViewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectExportSharedViewResponse(val *SubjectExportSharedViewResponse) *NullableSubjectExportSharedViewResponse {
	return &NullableSubjectExportSharedViewResponse{value: val, isSet: true}
}

func (v NullableSubjectExportSharedViewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectExportSharedViewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
