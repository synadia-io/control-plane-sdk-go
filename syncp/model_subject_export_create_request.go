/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the SubjectExportCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubjectExportCreateRequest{}

// SubjectExportCreateRequest struct for SubjectExportCreateRequest
type SubjectExportCreateRequest struct {
	JwtSettings               Export `json:"jwt_settings"`
	MetricsEnabled            bool   `json:"metrics_enabled"`
	MetricsSamplingPercentage int64  `json:"metrics_sampling_percentage"`
}

func (o SubjectExportCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jwt_settings"] = o.JwtSettings
	toSerialize["metrics_enabled"] = o.MetricsEnabled
	toSerialize["metrics_sampling_percentage"] = o.MetricsSamplingPercentage
	return toSerialize, nil
}
