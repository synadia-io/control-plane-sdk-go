/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
	"time"
)

// checks if the AlertRuleViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertRuleViewResponse{}

// AlertRuleViewResponse struct for AlertRuleViewResponse
type AlertRuleViewResponse struct {
	AccountId           NullableString    `json:"account_id,omitempty"`
	ConsumerName        NullableString    `json:"consumer_name,omitempty"`
	Created             time.Time         `json:"created"`
	DurationInSecs      int32             `json:"duration_in_secs"`
	Id                  string            `json:"id"`
	IsEnabled           bool              `json:"is_enabled"`
	Message             string            `json:"message"`
	Metric              string            `json:"metric"`
	ParentRuleId        NullableString    `json:"parent_rule_id,omitempty"`
	RuleType            AlertRuleType     `json:"rule_type"`
	Severity            AlertRuleSeverity `json:"severity"`
	StreamName          NullableString    `json:"stream_name,omitempty"`
	SystemId            NullableString    `json:"system_id,omitempty"`
	ThresholdExpression NullableString    `json:"threshold_expression,omitempty"`
	ThresholdFixedValue NullableFloat32   `json:"threshold_fixed_value,omitempty"`
	ThresholdOperator   string            `json:"threshold_operator"`
}

// NewAlertRuleViewResponse instantiates a new AlertRuleViewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertRuleViewResponse(created time.Time, durationInSecs int32, id string, isEnabled bool, message string, metric string, ruleType AlertRuleType, severity AlertRuleSeverity, thresholdOperator string) *AlertRuleViewResponse {
	this := AlertRuleViewResponse{}
	this.Created = created
	this.DurationInSecs = durationInSecs
	this.Id = id
	this.IsEnabled = isEnabled
	this.Message = message
	this.Metric = metric
	this.RuleType = ruleType
	this.Severity = severity
	this.ThresholdOperator = thresholdOperator
	return &this
}

// NewAlertRuleViewResponseWithDefaults instantiates a new AlertRuleViewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertRuleViewResponseWithDefaults() *AlertRuleViewResponse {
	this := AlertRuleViewResponse{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertRuleViewResponse) GetAccountId() string {
	if o == nil || IsNil(o.AccountId.Get()) {
		var ret string
		return ret
	}
	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertRuleViewResponse) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlertRuleViewResponse) HasAccountId() bool {
	if o != nil && o.AccountId.IsSet() {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given NullableString and assigns it to the AccountId field.
func (o *AlertRuleViewResponse) SetAccountId(v string) {
	o.AccountId.Set(&v)
}

// SetAccountIdNil sets the value for AccountId to be an explicit nil
func (o *AlertRuleViewResponse) SetAccountIdNil() {
	o.AccountId.Set(nil)
}

// UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
func (o *AlertRuleViewResponse) UnsetAccountId() {
	o.AccountId.Unset()
}

// GetConsumerName returns the ConsumerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertRuleViewResponse) GetConsumerName() string {
	if o == nil || IsNil(o.ConsumerName.Get()) {
		var ret string
		return ret
	}
	return *o.ConsumerName.Get()
}

// GetConsumerNameOk returns a tuple with the ConsumerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertRuleViewResponse) GetConsumerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConsumerName.Get(), o.ConsumerName.IsSet()
}

// HasConsumerName returns a boolean if a field has been set.
func (o *AlertRuleViewResponse) HasConsumerName() bool {
	if o != nil && o.ConsumerName.IsSet() {
		return true
	}

	return false
}

// SetConsumerName gets a reference to the given NullableString and assigns it to the ConsumerName field.
func (o *AlertRuleViewResponse) SetConsumerName(v string) {
	o.ConsumerName.Set(&v)
}

// SetConsumerNameNil sets the value for ConsumerName to be an explicit nil
func (o *AlertRuleViewResponse) SetConsumerNameNil() {
	o.ConsumerName.Set(nil)
}

// UnsetConsumerName ensures that no value is present for ConsumerName, not even an explicit nil
func (o *AlertRuleViewResponse) UnsetConsumerName() {
	o.ConsumerName.Unset()
}

// GetCreated returns the Created field value
func (o *AlertRuleViewResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *AlertRuleViewResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *AlertRuleViewResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetDurationInSecs returns the DurationInSecs field value
func (o *AlertRuleViewResponse) GetDurationInSecs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DurationInSecs
}

// GetDurationInSecsOk returns a tuple with the DurationInSecs field value
// and a boolean to check if the value has been set.
func (o *AlertRuleViewResponse) GetDurationInSecsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationInSecs, true
}

// SetDurationInSecs sets field value
func (o *AlertRuleViewResponse) SetDurationInSecs(v int32) {
	o.DurationInSecs = v
}

// GetId returns the Id field value
func (o *AlertRuleViewResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AlertRuleViewResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AlertRuleViewResponse) SetId(v string) {
	o.Id = v
}

// GetIsEnabled returns the IsEnabled field value
func (o *AlertRuleViewResponse) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *AlertRuleViewResponse) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *AlertRuleViewResponse) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetMessage returns the Message field value
func (o *AlertRuleViewResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlertRuleViewResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AlertRuleViewResponse) SetMessage(v string) {
	o.Message = v
}

// GetMetric returns the Metric field value
func (o *AlertRuleViewResponse) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *AlertRuleViewResponse) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *AlertRuleViewResponse) SetMetric(v string) {
	o.Metric = v
}

// GetParentRuleId returns the ParentRuleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertRuleViewResponse) GetParentRuleId() string {
	if o == nil || IsNil(o.ParentRuleId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentRuleId.Get()
}

// GetParentRuleIdOk returns a tuple with the ParentRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertRuleViewResponse) GetParentRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentRuleId.Get(), o.ParentRuleId.IsSet()
}

// HasParentRuleId returns a boolean if a field has been set.
func (o *AlertRuleViewResponse) HasParentRuleId() bool {
	if o != nil && o.ParentRuleId.IsSet() {
		return true
	}

	return false
}

// SetParentRuleId gets a reference to the given NullableString and assigns it to the ParentRuleId field.
func (o *AlertRuleViewResponse) SetParentRuleId(v string) {
	o.ParentRuleId.Set(&v)
}

// SetParentRuleIdNil sets the value for ParentRuleId to be an explicit nil
func (o *AlertRuleViewResponse) SetParentRuleIdNil() {
	o.ParentRuleId.Set(nil)
}

// UnsetParentRuleId ensures that no value is present for ParentRuleId, not even an explicit nil
func (o *AlertRuleViewResponse) UnsetParentRuleId() {
	o.ParentRuleId.Unset()
}

// GetRuleType returns the RuleType field value
func (o *AlertRuleViewResponse) GetRuleType() AlertRuleType {
	if o == nil {
		var ret AlertRuleType
		return ret
	}

	return o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value
// and a boolean to check if the value has been set.
func (o *AlertRuleViewResponse) GetRuleTypeOk() (*AlertRuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleType, true
}

// SetRuleType sets field value
func (o *AlertRuleViewResponse) SetRuleType(v AlertRuleType) {
	o.RuleType = v
}

// GetSeverity returns the Severity field value
func (o *AlertRuleViewResponse) GetSeverity() AlertRuleSeverity {
	if o == nil {
		var ret AlertRuleSeverity
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *AlertRuleViewResponse) GetSeverityOk() (*AlertRuleSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *AlertRuleViewResponse) SetSeverity(v AlertRuleSeverity) {
	o.Severity = v
}

// GetStreamName returns the StreamName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertRuleViewResponse) GetStreamName() string {
	if o == nil || IsNil(o.StreamName.Get()) {
		var ret string
		return ret
	}
	return *o.StreamName.Get()
}

// GetStreamNameOk returns a tuple with the StreamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertRuleViewResponse) GetStreamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreamName.Get(), o.StreamName.IsSet()
}

// HasStreamName returns a boolean if a field has been set.
func (o *AlertRuleViewResponse) HasStreamName() bool {
	if o != nil && o.StreamName.IsSet() {
		return true
	}

	return false
}

// SetStreamName gets a reference to the given NullableString and assigns it to the StreamName field.
func (o *AlertRuleViewResponse) SetStreamName(v string) {
	o.StreamName.Set(&v)
}

// SetStreamNameNil sets the value for StreamName to be an explicit nil
func (o *AlertRuleViewResponse) SetStreamNameNil() {
	o.StreamName.Set(nil)
}

// UnsetStreamName ensures that no value is present for StreamName, not even an explicit nil
func (o *AlertRuleViewResponse) UnsetStreamName() {
	o.StreamName.Unset()
}

// GetSystemId returns the SystemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertRuleViewResponse) GetSystemId() string {
	if o == nil || IsNil(o.SystemId.Get()) {
		var ret string
		return ret
	}
	return *o.SystemId.Get()
}

// GetSystemIdOk returns a tuple with the SystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertRuleViewResponse) GetSystemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SystemId.Get(), o.SystemId.IsSet()
}

// HasSystemId returns a boolean if a field has been set.
func (o *AlertRuleViewResponse) HasSystemId() bool {
	if o != nil && o.SystemId.IsSet() {
		return true
	}

	return false
}

// SetSystemId gets a reference to the given NullableString and assigns it to the SystemId field.
func (o *AlertRuleViewResponse) SetSystemId(v string) {
	o.SystemId.Set(&v)
}

// SetSystemIdNil sets the value for SystemId to be an explicit nil
func (o *AlertRuleViewResponse) SetSystemIdNil() {
	o.SystemId.Set(nil)
}

// UnsetSystemId ensures that no value is present for SystemId, not even an explicit nil
func (o *AlertRuleViewResponse) UnsetSystemId() {
	o.SystemId.Unset()
}

// GetThresholdExpression returns the ThresholdExpression field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertRuleViewResponse) GetThresholdExpression() string {
	if o == nil || IsNil(o.ThresholdExpression.Get()) {
		var ret string
		return ret
	}
	return *o.ThresholdExpression.Get()
}

// GetThresholdExpressionOk returns a tuple with the ThresholdExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertRuleViewResponse) GetThresholdExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThresholdExpression.Get(), o.ThresholdExpression.IsSet()
}

// HasThresholdExpression returns a boolean if a field has been set.
func (o *AlertRuleViewResponse) HasThresholdExpression() bool {
	if o != nil && o.ThresholdExpression.IsSet() {
		return true
	}

	return false
}

// SetThresholdExpression gets a reference to the given NullableString and assigns it to the ThresholdExpression field.
func (o *AlertRuleViewResponse) SetThresholdExpression(v string) {
	o.ThresholdExpression.Set(&v)
}

// SetThresholdExpressionNil sets the value for ThresholdExpression to be an explicit nil
func (o *AlertRuleViewResponse) SetThresholdExpressionNil() {
	o.ThresholdExpression.Set(nil)
}

// UnsetThresholdExpression ensures that no value is present for ThresholdExpression, not even an explicit nil
func (o *AlertRuleViewResponse) UnsetThresholdExpression() {
	o.ThresholdExpression.Unset()
}

// GetThresholdFixedValue returns the ThresholdFixedValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertRuleViewResponse) GetThresholdFixedValue() float32 {
	if o == nil || IsNil(o.ThresholdFixedValue.Get()) {
		var ret float32
		return ret
	}
	return *o.ThresholdFixedValue.Get()
}

// GetThresholdFixedValueOk returns a tuple with the ThresholdFixedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertRuleViewResponse) GetThresholdFixedValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThresholdFixedValue.Get(), o.ThresholdFixedValue.IsSet()
}

// HasThresholdFixedValue returns a boolean if a field has been set.
func (o *AlertRuleViewResponse) HasThresholdFixedValue() bool {
	if o != nil && o.ThresholdFixedValue.IsSet() {
		return true
	}

	return false
}

// SetThresholdFixedValue gets a reference to the given NullableFloat32 and assigns it to the ThresholdFixedValue field.
func (o *AlertRuleViewResponse) SetThresholdFixedValue(v float32) {
	o.ThresholdFixedValue.Set(&v)
}

// SetThresholdFixedValueNil sets the value for ThresholdFixedValue to be an explicit nil
func (o *AlertRuleViewResponse) SetThresholdFixedValueNil() {
	o.ThresholdFixedValue.Set(nil)
}

// UnsetThresholdFixedValue ensures that no value is present for ThresholdFixedValue, not even an explicit nil
func (o *AlertRuleViewResponse) UnsetThresholdFixedValue() {
	o.ThresholdFixedValue.Unset()
}

// GetThresholdOperator returns the ThresholdOperator field value
func (o *AlertRuleViewResponse) GetThresholdOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThresholdOperator
}

// GetThresholdOperatorOk returns a tuple with the ThresholdOperator field value
// and a boolean to check if the value has been set.
func (o *AlertRuleViewResponse) GetThresholdOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThresholdOperator, true
}

// SetThresholdOperator sets field value
func (o *AlertRuleViewResponse) SetThresholdOperator(v string) {
	o.ThresholdOperator = v
}

func (o AlertRuleViewResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertRuleViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId.IsSet() {
		toSerialize["account_id"] = o.AccountId.Get()
	}
	if o.ConsumerName.IsSet() {
		toSerialize["consumer_name"] = o.ConsumerName.Get()
	}
	toSerialize["created"] = o.Created
	toSerialize["duration_in_secs"] = o.DurationInSecs
	toSerialize["id"] = o.Id
	toSerialize["is_enabled"] = o.IsEnabled
	toSerialize["message"] = o.Message
	toSerialize["metric"] = o.Metric
	if o.ParentRuleId.IsSet() {
		toSerialize["parent_rule_id"] = o.ParentRuleId.Get()
	}
	toSerialize["rule_type"] = o.RuleType
	toSerialize["severity"] = o.Severity
	if o.StreamName.IsSet() {
		toSerialize["stream_name"] = o.StreamName.Get()
	}
	if o.SystemId.IsSet() {
		toSerialize["system_id"] = o.SystemId.Get()
	}
	if o.ThresholdExpression.IsSet() {
		toSerialize["threshold_expression"] = o.ThresholdExpression.Get()
	}
	if o.ThresholdFixedValue.IsSet() {
		toSerialize["threshold_fixed_value"] = o.ThresholdFixedValue.Get()
	}
	toSerialize["threshold_operator"] = o.ThresholdOperator
	return toSerialize, nil
}

type NullableAlertRuleViewResponse struct {
	value *AlertRuleViewResponse
	isSet bool
}

func (v NullableAlertRuleViewResponse) Get() *AlertRuleViewResponse {
	return v.value
}

func (v *NullableAlertRuleViewResponse) Set(val *AlertRuleViewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertRuleViewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertRuleViewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertRuleViewResponse(val *AlertRuleViewResponse) *NullableAlertRuleViewResponse {
	return &NullableAlertRuleViewResponse{value: val, isSet: true}
}

func (v NullableAlertRuleViewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertRuleViewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}