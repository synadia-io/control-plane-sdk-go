/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the TimeRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeRange{}

// TimeRange TimeRange is used to represent a start and end time
type TimeRange struct {
	End   *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
}

func (o TimeRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	return toSerialize, nil
}
