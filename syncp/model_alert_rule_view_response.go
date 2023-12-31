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

// checks if the AlertRuleViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertRuleViewResponse{}

// AlertRuleViewResponse struct for AlertRuleViewResponse
type AlertRuleViewResponse struct {
	AccountId           *Nullable[string]  `json:"account_id,omitempty"`
	ConsumerName        *Nullable[string]  `json:"consumer_name,omitempty"`
	Created             time.Time          `json:"created"`
	DurationInSecs      int32              `json:"duration_in_secs"`
	Id                  string             `json:"id"`
	IsEnabled           bool               `json:"is_enabled"`
	Message             string             `json:"message"`
	Metric              string             `json:"metric"`
	ParentRuleId        *Nullable[string]  `json:"parent_rule_id,omitempty"`
	RuleType            AlertRuleType      `json:"rule_type"`
	Severity            AlertRuleSeverity  `json:"severity"`
	StreamName          *Nullable[string]  `json:"stream_name,omitempty"`
	SystemId            *Nullable[string]  `json:"system_id,omitempty"`
	ThresholdExpression *Nullable[string]  `json:"threshold_expression,omitempty"`
	ThresholdFixedValue *Nullable[float32] `json:"threshold_fixed_value,omitempty"`
	ThresholdOperator   string             `json:"threshold_operator"`
}

func (o AlertRuleViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil && !o.AccountId.IsNull() {
		toSerialize["account_id"] = o.AccountId.Val
	}
	if o.ConsumerName != nil && !o.ConsumerName.IsNull() {
		toSerialize["consumer_name"] = o.ConsumerName.Val
	}
	toSerialize["created"] = o.Created
	toSerialize["duration_in_secs"] = o.DurationInSecs
	toSerialize["id"] = o.Id
	toSerialize["is_enabled"] = o.IsEnabled
	toSerialize["message"] = o.Message
	toSerialize["metric"] = o.Metric
	if o.ParentRuleId != nil && !o.ParentRuleId.IsNull() {
		toSerialize["parent_rule_id"] = o.ParentRuleId.Val
	}
	toSerialize["rule_type"] = o.RuleType
	toSerialize["severity"] = o.Severity
	if o.StreamName != nil && !o.StreamName.IsNull() {
		toSerialize["stream_name"] = o.StreamName.Val
	}
	if o.SystemId != nil && !o.SystemId.IsNull() {
		toSerialize["system_id"] = o.SystemId.Val
	}
	if o.ThresholdExpression != nil && !o.ThresholdExpression.IsNull() {
		toSerialize["threshold_expression"] = o.ThresholdExpression.Val
	}
	if o.ThresholdFixedValue != nil && !o.ThresholdFixedValue.IsNull() {
		toSerialize["threshold_fixed_value"] = o.ThresholdFixedValue.Val
	}
	toSerialize["threshold_operator"] = o.ThresholdOperator
	return toSerialize, nil
}
