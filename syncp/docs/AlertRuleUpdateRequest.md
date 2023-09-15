# AlertRuleUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationInSecs** | **int32** |  | 
**IsEnabled** | **bool** |  | 
**Message** | **string** |  | 
**Metric** | **string** |  | 
**Severity** | [**AlertRuleSeverity**](AlertRuleSeverity.md) |  | 
**ThresholdExpression** | Pointer to **NullableString** |  | [optional] 
**ThresholdFixedValue** | Pointer to **NullableFloat32** |  | [optional] 
**ThresholdOperator** | [**AlertRuleOperator**](AlertRuleOperator.md) |  | 

## Methods

### NewAlertRuleUpdateRequest

`func NewAlertRuleUpdateRequest(durationInSecs int32, isEnabled bool, message string, metric string, severity AlertRuleSeverity, thresholdOperator AlertRuleOperator, ) *AlertRuleUpdateRequest`

NewAlertRuleUpdateRequest instantiates a new AlertRuleUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleUpdateRequestWithDefaults

`func NewAlertRuleUpdateRequestWithDefaults() *AlertRuleUpdateRequest`

NewAlertRuleUpdateRequestWithDefaults instantiates a new AlertRuleUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationInSecs

`func (o *AlertRuleUpdateRequest) GetDurationInSecs() int32`

GetDurationInSecs returns the DurationInSecs field if non-nil, zero value otherwise.

### GetDurationInSecsOk

`func (o *AlertRuleUpdateRequest) GetDurationInSecsOk() (*int32, bool)`

GetDurationInSecsOk returns a tuple with the DurationInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInSecs

`func (o *AlertRuleUpdateRequest) SetDurationInSecs(v int32)`

SetDurationInSecs sets DurationInSecs field to given value.


### GetIsEnabled

`func (o *AlertRuleUpdateRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AlertRuleUpdateRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AlertRuleUpdateRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetMessage

`func (o *AlertRuleUpdateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertRuleUpdateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertRuleUpdateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMetric

`func (o *AlertRuleUpdateRequest) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *AlertRuleUpdateRequest) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *AlertRuleUpdateRequest) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetSeverity

`func (o *AlertRuleUpdateRequest) GetSeverity() AlertRuleSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertRuleUpdateRequest) GetSeverityOk() (*AlertRuleSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertRuleUpdateRequest) SetSeverity(v AlertRuleSeverity)`

SetSeverity sets Severity field to given value.


### GetThresholdExpression

`func (o *AlertRuleUpdateRequest) GetThresholdExpression() string`

GetThresholdExpression returns the ThresholdExpression field if non-nil, zero value otherwise.

### GetThresholdExpressionOk

`func (o *AlertRuleUpdateRequest) GetThresholdExpressionOk() (*string, bool)`

GetThresholdExpressionOk returns a tuple with the ThresholdExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdExpression

`func (o *AlertRuleUpdateRequest) SetThresholdExpression(v string)`

SetThresholdExpression sets ThresholdExpression field to given value.

### HasThresholdExpression

`func (o *AlertRuleUpdateRequest) HasThresholdExpression() bool`

HasThresholdExpression returns a boolean if a field has been set.

### SetThresholdExpressionNil

`func (o *AlertRuleUpdateRequest) SetThresholdExpressionNil(b bool)`

 SetThresholdExpressionNil sets the value for ThresholdExpression to be an explicit nil

### UnsetThresholdExpression
`func (o *AlertRuleUpdateRequest) UnsetThresholdExpression()`

UnsetThresholdExpression ensures that no value is present for ThresholdExpression, not even an explicit nil
### GetThresholdFixedValue

`func (o *AlertRuleUpdateRequest) GetThresholdFixedValue() float32`

GetThresholdFixedValue returns the ThresholdFixedValue field if non-nil, zero value otherwise.

### GetThresholdFixedValueOk

`func (o *AlertRuleUpdateRequest) GetThresholdFixedValueOk() (*float32, bool)`

GetThresholdFixedValueOk returns a tuple with the ThresholdFixedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdFixedValue

`func (o *AlertRuleUpdateRequest) SetThresholdFixedValue(v float32)`

SetThresholdFixedValue sets ThresholdFixedValue field to given value.

### HasThresholdFixedValue

`func (o *AlertRuleUpdateRequest) HasThresholdFixedValue() bool`

HasThresholdFixedValue returns a boolean if a field has been set.

### SetThresholdFixedValueNil

`func (o *AlertRuleUpdateRequest) SetThresholdFixedValueNil(b bool)`

 SetThresholdFixedValueNil sets the value for ThresholdFixedValue to be an explicit nil

### UnsetThresholdFixedValue
`func (o *AlertRuleUpdateRequest) UnsetThresholdFixedValue()`

UnsetThresholdFixedValue ensures that no value is present for ThresholdFixedValue, not even an explicit nil
### GetThresholdOperator

`func (o *AlertRuleUpdateRequest) GetThresholdOperator() AlertRuleOperator`

GetThresholdOperator returns the ThresholdOperator field if non-nil, zero value otherwise.

### GetThresholdOperatorOk

`func (o *AlertRuleUpdateRequest) GetThresholdOperatorOk() (*AlertRuleOperator, bool)`

GetThresholdOperatorOk returns a tuple with the ThresholdOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdOperator

`func (o *AlertRuleUpdateRequest) SetThresholdOperator(v AlertRuleOperator)`

SetThresholdOperator sets ThresholdOperator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


