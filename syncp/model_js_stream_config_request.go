/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the JSStreamConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSStreamConfigRequest{}

// JSStreamConfigRequest struct for JSStreamConfigRequest
type JSStreamConfigRequest struct {
	JSCommonStreamConfig
	Subjects []string `json:"subjects,omitempty"`
}

func (o JSStreamConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if len(o.Subjects) != 0 {
		toSerialize["subjects"] = o.Subjects
	}
	toSerialize["allow_direct"] = o.AllowDirect
	toSerialize["allow_rollup_hdrs"] = o.AllowRollupHdrs
	if o.Compression != nil {
		toSerialize["compression"] = o.Compression
	}
	toSerialize["deny_delete"] = o.DenyDelete
	toSerialize["deny_purge"] = o.DenyPurge
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["discard"] = o.Discard
	if o.DiscardNewPerSubject != nil {
		toSerialize["discard_new_per_subject"] = o.DiscardNewPerSubject
	}
	if o.DuplicateWindow != nil {
		toSerialize["duplicate_window"] = o.DuplicateWindow
	}
	if o.FirstSeq != nil {
		toSerialize["first_seq"] = o.FirstSeq
	}
	toSerialize["max_age"] = o.MaxAge
	toSerialize["max_bytes"] = o.MaxBytes
	toSerialize["max_consumers"] = o.MaxConsumers
	if o.MaxMsgSize != nil {
		toSerialize["max_msg_size"] = o.MaxMsgSize
	}
	toSerialize["max_msgs"] = o.MaxMsgs
	toSerialize["max_msgs_per_subject"] = o.MaxMsgsPerSubject
	toSerialize["name"] = o.Name
	if o.NoAck != nil {
		toSerialize["no_ack"] = o.NoAck
	}
	toSerialize["num_replicas"] = o.NumReplicas
	if o.Placement != nil {
		toSerialize["placement"] = o.Placement
	}
	if o.Republish != nil {
		toSerialize["republish"] = o.Republish
	}
	toSerialize["retention"] = o.Retention
	toSerialize["sealed"] = o.Sealed
	if len(o.Sources) != 0 {
		toSerialize["sources"] = o.Sources
	}
	toSerialize["storage"] = o.Storage
	if o.SubjectTransform != nil {
		toSerialize["subject_transform"] = o.SubjectTransform
	}
	if o.TemplateOwner != nil {
		toSerialize["template_owner"] = o.TemplateOwner
	}
	return toSerialize, nil
}
