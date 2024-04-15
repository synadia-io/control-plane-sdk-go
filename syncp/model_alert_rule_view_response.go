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
	AccountId                     *string           `json:"account_id,omitempty"`
	ConsumerName                  *string           `json:"consumer_name,omitempty"`
	Created                       time.Time         `json:"created"`
	DurationInSecs                int32             `json:"duration_in_secs"`
	Id                            string            `json:"id"`
	IsEnabled                     bool              `json:"is_enabled"`
	Message                       string            `json:"message"`
	Metric                        string            `json:"metric"`
	RuleType                      AlertRuleType     `json:"rule_type"`
	Severity                      AlertRuleSeverity `json:"severity"`
	StreamName                    *string           `json:"stream_name,omitempty"`
	SystemId                      *string           `json:"system_id,omitempty"`
	ThresholdExpressionMetric     *string           `json:"threshold_expression_metric,omitempty"`
	ThresholdExpressionMultiplier *float32          `json:"threshold_expression_multiplier,omitempty"`
	ThresholdFixedValue           *float32          `json:"threshold_fixed_value,omitempty"`
	ThresholdOperator             string            `json:"threshold_operator"`
}

func (o AlertRuleViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.ConsumerName != nil {
		toSerialize["consumer_name"] = o.ConsumerName
	}
	toSerialize["created"] = o.Created
	toSerialize["duration_in_secs"] = o.DurationInSecs
	toSerialize["id"] = o.Id
	toSerialize["is_enabled"] = o.IsEnabled
	toSerialize["message"] = o.Message
	toSerialize["metric"] = o.Metric
	toSerialize["rule_type"] = o.RuleType
	toSerialize["severity"] = o.Severity
	if o.StreamName != nil {
		toSerialize["stream_name"] = o.StreamName
	}
	if o.SystemId != nil {
		toSerialize["system_id"] = o.SystemId
	}
	if o.ThresholdExpressionMetric != nil {
		toSerialize["threshold_expression_metric"] = o.ThresholdExpressionMetric
	}
	if o.ThresholdExpressionMultiplier != nil {
		toSerialize["threshold_expression_multiplier"] = o.ThresholdExpressionMultiplier
	}
	if o.ThresholdFixedValue != nil {
		toSerialize["threshold_fixed_value"] = o.ThresholdFixedValue
	}
	toSerialize["threshold_operator"] = o.ThresholdOperator
	return toSerialize, nil
}
