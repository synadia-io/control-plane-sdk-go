/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the InvitationDecisionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvitationDecisionRequest{}

// InvitationDecisionRequest struct for InvitationDecisionRequest
type InvitationDecisionRequest struct {
	Accept bool `json:"accept"`
}

func (o InvitationDecisionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accept"] = o.Accept
	return toSerialize, nil
}
