/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the StreamExportSharedViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamExportSharedViewResponse{}

// StreamExportSharedViewResponse struct for StreamExportSharedViewResponse
type StreamExportSharedViewResponse struct {
	DeliverSubjectPrefix    string `json:"deliver_subject_prefix"`
	IsPublic                bool   `json:"is_public"`
	JsSubjectPrefix         string `json:"js_subject_prefix"`
	RemoteAccountNkeyPublic string `json:"remote_account_nkey_public"`
	StreamName              string `json:"stream_name"`
}

func (o StreamExportSharedViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deliver_subject_prefix"] = o.DeliverSubjectPrefix
	toSerialize["is_public"] = o.IsPublic
	toSerialize["js_subject_prefix"] = o.JsSubjectPrefix
	toSerialize["remote_account_nkey_public"] = o.RemoteAccountNkeyPublic
	toSerialize["stream_name"] = o.StreamName
	return toSerialize, nil
}
