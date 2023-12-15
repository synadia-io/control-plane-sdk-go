/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the StreamShareListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamShareListResponse{}

// StreamShareListResponse struct for StreamShareListResponse
type StreamShareListResponse struct {
	Items []StreamShareViewResponse `json:"items"`
}

func (o StreamShareListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}
