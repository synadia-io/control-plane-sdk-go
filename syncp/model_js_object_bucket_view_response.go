/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the JSObjectBucketViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSObjectBucketViewResponse{}

// JSObjectBucketViewResponse struct for JSObjectBucketViewResponse
type JSObjectBucketViewResponse struct {
	Bytes      uint64               `json:"bytes"`
	Config     JSObjectBucketConfig `json:"config"`
	Id         string               `json:"id"`
	StreamName string               `json:"stream_name"`
}

func (o JSObjectBucketViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bytes"] = o.Bytes
	toSerialize["config"] = o.Config
	toSerialize["id"] = o.Id
	toSerialize["stream_name"] = o.StreamName
	return toSerialize, nil
}
