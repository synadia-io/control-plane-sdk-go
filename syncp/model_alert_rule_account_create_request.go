/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the AlertRuleAccountCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertRuleAccountCreateRequest{}

// AlertRuleAccountCreateRequest struct for AlertRuleAccountCreateRequest
type AlertRuleAccountCreateRequest struct {
	AlertRuleBaseCreateRequest
	ConsumerName *Nullable[string] `json:"consumer_name,omitempty"`
	RuleType     string            `json:"rule_type"`
	StreamName   *Nullable[string] `json:"stream_name,omitempty"`
}

func (o AlertRuleAccountCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ConsumerName != nil && !o.ConsumerName.IsNull() {
		toSerialize["consumer_name"] = o.ConsumerName.Val
	}
	toSerialize["rule_type"] = o.RuleType
	if o.StreamName != nil && !o.StreamName.IsNull() {
		toSerialize["stream_name"] = o.StreamName.Val
	}
	toSerialize["duration_in_secs"] = o.DurationInSecs
	toSerialize["message"] = o.Message
	toSerialize["metric"] = o.Metric
	toSerialize["severity"] = o.Severity
	if o.ThresholdExpression != nil && !o.ThresholdExpression.IsNull() {
		toSerialize["threshold_expression"] = o.ThresholdExpression.Val
	}
	if o.ThresholdFixedValue != nil && !o.ThresholdFixedValue.IsNull() {
		toSerialize["threshold_fixed_value"] = o.ThresholdFixedValue.Val
	}
	toSerialize["threshold_operator"] = o.ThresholdOperator
	return toSerialize, nil
}
