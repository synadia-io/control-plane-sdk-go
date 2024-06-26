/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"time"
)

// checks if the SubjectExportViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubjectExportViewResponse{}

// SubjectExportViewResponse struct for SubjectExportViewResponse
type SubjectExportViewResponse struct {
	Account                   AccountInfo      `json:"account"`
	Created                   time.Time        `json:"created"`
	Id                        string           `json:"id"`
	IsPublic                  bool             `json:"is_public"`
	JwtSettings               Export           `json:"jwt_settings"`
	MetricsEnabled            bool             `json:"metrics_enabled"`
	MetricsSamplingPercentage int64            `json:"metrics_sampling_percentage"`
	Name                      Nullable[string] `json:"name"`
	ShareSkPublic             Nullable[string] `json:"share_sk_public"`
	StreamExportId            Nullable[string] `json:"stream_export_id"`
	Subject                   string           `json:"subject"`
}

func (o SubjectExportViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account"] = o.Account
	toSerialize["created"] = o.Created
	toSerialize["id"] = o.Id
	toSerialize["is_public"] = o.IsPublic
	toSerialize["jwt_settings"] = o.JwtSettings
	toSerialize["metrics_enabled"] = o.MetricsEnabled
	toSerialize["metrics_sampling_percentage"] = o.MetricsSamplingPercentage
	if !o.Name.IsNull() {
		toSerialize["name"] = o.Name.Val
	}
	if !o.ShareSkPublic.IsNull() {
		toSerialize["share_sk_public"] = o.ShareSkPublic.Val
	}
	if !o.StreamExportId.IsNull() {
		toSerialize["stream_export_id"] = o.StreamExportId.Val
	}
	toSerialize["subject"] = o.Subject
	return toSerialize, nil
}
