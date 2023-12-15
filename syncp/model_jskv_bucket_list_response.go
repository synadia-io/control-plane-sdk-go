/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the JSKVBucketListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSKVBucketListResponse{}

// JSKVBucketListResponse struct for JSKVBucketListResponse
type JSKVBucketListResponse struct {
	Items []JSKVBucketViewResponse `json:"items"`
}

func (o JSKVBucketListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}
