/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the Export type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Export{}

// Export Export represents a single export
type Export struct {
	Info
	AccountTokenPosition *int64        `json:"account_token_position,omitempty"`
	Advertise            *bool         `json:"advertise,omitempty"`
	Name                 *string       `json:"name,omitempty"`
	ResponseThreshold    *int64        `json:"response_threshold,omitempty"`
	ResponseType         *ResponseType `json:"response_type,omitempty"`
	// RevocationList is used to store a mapping of public keys to unix timestamps
	Revocations    map[string]int64 `json:"revocations,omitempty"`
	ServiceLatency *ServiceLatency  `json:"service_latency,omitempty"`
	Subject        *string          `json:"subject,omitempty"`
	TokenReq       *bool            `json:"token_req,omitempty"`
	Type           *ExportType      `json:"type,omitempty"`
}

func (o Export) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountTokenPosition != nil {
		toSerialize["account_token_position"] = o.AccountTokenPosition
	}
	if o.Advertise != nil {
		toSerialize["advertise"] = o.Advertise
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ResponseThreshold != nil {
		toSerialize["response_threshold"] = o.ResponseThreshold
	}
	if o.ResponseType != nil {
		toSerialize["response_type"] = o.ResponseType
	}
	if len(o.Revocations) != 0 {
		toSerialize["revocations"] = o.Revocations
	}
	if o.ServiceLatency != nil {
		toSerialize["service_latency"] = o.ServiceLatency
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.TokenReq != nil {
		toSerialize["token_req"] = o.TokenReq
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.InfoUrl != nil {
		toSerialize["info_url"] = o.InfoUrl
	}
	return toSerialize, nil
}
