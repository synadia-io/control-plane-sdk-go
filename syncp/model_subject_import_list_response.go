/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the SubjectImportListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubjectImportListResponse{}

// SubjectImportListResponse struct for SubjectImportListResponse
type SubjectImportListResponse struct {
	Items []SubjectImportViewResponse `json:"items"`
}

func (o SubjectImportListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}
