/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the Placement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Placement{}

// Placement Placement describes stream placement requirements for a stream
type Placement struct {
	Cluster string    `json:"cluster"`
	Tags    *[]string `json:"tags,omitempty"`
}

func (o Placement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster"] = o.Cluster
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}
