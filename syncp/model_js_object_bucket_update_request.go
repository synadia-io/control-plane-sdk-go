/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the JSObjectBucketUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSObjectBucketUpdateRequest{}

// JSObjectBucketUpdateRequest struct for JSObjectBucketUpdateRequest
type JSObjectBucketUpdateRequest struct {
	Config       UpdatableObjectBucketConfig `json:"config"`
	ReallySealed *bool                       `json:"really_sealed,omitempty"`
	Sealed       *bool                       `json:"sealed,omitempty"`
}

func (o JSObjectBucketUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.Config
	if o.ReallySealed != nil {
		toSerialize["really_sealed"] = o.ReallySealed
	}
	if o.Sealed != nil {
		toSerialize["sealed"] = o.Sealed
	}
	return toSerialize, nil
}
