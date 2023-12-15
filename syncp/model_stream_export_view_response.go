/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"time"
)

// checks if the StreamExportViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamExportViewResponse{}

// StreamExportViewResponse struct for StreamExportViewResponse
type StreamExportViewResponse struct {
	AccountId            string    `json:"account_id"`
	Created              time.Time `json:"created"`
	DeliverSubjectPrefix string    `json:"deliver_subject_prefix"`
	Id                   string    `json:"id"`
	IsPublic             bool      `json:"is_public"`
	JsSubjectPrefix      string    `json:"js_subject_prefix"`
	StreamName           string    `json:"stream_name"`
}

func (o StreamExportViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["created"] = o.Created
	toSerialize["deliver_subject_prefix"] = o.DeliverSubjectPrefix
	toSerialize["id"] = o.Id
	toSerialize["is_public"] = o.IsPublic
	toSerialize["js_subject_prefix"] = o.JsSubjectPrefix
	toSerialize["stream_name"] = o.StreamName
	return toSerialize, nil
}
