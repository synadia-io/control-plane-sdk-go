/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the JSAssetInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSAssetInfoResponse{}

// JSAssetInfoResponse struct for JSAssetInfoResponse
type JSAssetInfoResponse struct {
	JSCommonStreamInfo
	Config        JSCommonStreamConfig `json:"config"`
	Id            string               `json:"id"`
	JetstreamType JSType               `json:"jetstream_type"`
}

func (o JSAssetInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.Config
	toSerialize["id"] = o.Id
	toSerialize["jetstream_type"] = o.JetstreamType
	if len(o.Alternates) != 0 {
		toSerialize["alternates"] = o.Alternates
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	toSerialize["created"] = o.Created
	if len(o.Sources) != 0 {
		toSerialize["sources"] = o.Sources
	}
	toSerialize["state"] = o.State
	return toSerialize, nil
}
