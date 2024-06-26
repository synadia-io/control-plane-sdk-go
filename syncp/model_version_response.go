/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the VersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionResponse{}

// VersionResponse struct for VersionResponse
type VersionResponse struct {
	Commit  string `json:"commit"`
	Date    string `json:"date"`
	Version string `json:"version"`
}

func (o VersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["commit"] = o.Commit
	toSerialize["date"] = o.Date
	toSerialize["version"] = o.Version
	return toSerialize, nil
}
