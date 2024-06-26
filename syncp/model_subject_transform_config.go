/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the SubjectTransformConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubjectTransformConfig{}

// SubjectTransformConfig SubjectTransformConfig is for applying a subject transform (to matching messages) before doing anything else when a new message is received
type SubjectTransformConfig struct {
	Dest string `json:"dest"`
	Src  string `json:"src"`
}

func (o SubjectTransformConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dest"] = o.Dest
	toSerialize["src"] = o.Src
	return toSerialize, nil
}
