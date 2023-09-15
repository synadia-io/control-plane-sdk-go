# AlertRuleBaseCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationInSecs** | **int32** |  | 
**Message** | **string** |  | 
**Metric** | **string** |  | 
**Severity** | [**AlertRuleSeverity**](AlertRuleSeverity.md) |  | 
**ThresholdExpression** | Pointer to **NullableString** |  | [optional] 
**ThresholdFixedValue** | Pointer to **NullableFloat32** |  | [optional] 
**ThresholdOperator** | [**AlertRuleOperator**](AlertRuleOperator.md) |  | 

## Methods

### NewAlertRuleBaseCreateRequest

`func NewAlertRuleBaseCreateRequest(durationInSecs int32, message string, metric string, severity AlertRuleSeverity, thresholdOperator AlertRuleOperator, ) *AlertRuleBaseCreateRequest`

NewAlertRuleBaseCreateRequest instantiates a new AlertRuleBaseCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleBaseCreateRequestWithDefaults

`func NewAlertRuleBaseCreateRequestWithDefaults() *AlertRuleBaseCreateRequest`

NewAlertRuleBaseCreateRequestWithDefaults instantiates a new AlertRuleBaseCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationInSecs

`func (o *AlertRuleBaseCreateRequest) GetDurationInSecs() int32`

GetDurationInSecs returns the DurationInSecs field if non-nil, zero value otherwise.

### GetDurationInSecsOk

`func (o *AlertRuleBaseCreateRequest) GetDurationInSecsOk() (*int32, bool)`

GetDurationInSecsOk returns a tuple with the DurationInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInSecs

`func (o *AlertRuleBaseCreateRequest) SetDurationInSecs(v int32)`

SetDurationInSecs sets DurationInSecs field to given value.


### GetMessage

`func (o *AlertRuleBaseCreateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertRuleBaseCreateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertRuleBaseCreateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMetric

`func (o *AlertRuleBaseCreateRequest) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *AlertRuleBaseCreateRequest) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *AlertRuleBaseCreateRequest) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetSeverity

`func (o *AlertRuleBaseCreateRequest) GetSeverity() AlertRuleSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertRuleBaseCreateRequest) GetSeverityOk() (*AlertRuleSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertRuleBaseCreateRequest) SetSeverity(v AlertRuleSeverity)`

SetSeverity sets Severity field to given value.


### GetThresholdExpression

`func (o *AlertRuleBaseCreateRequest) GetThresholdExpression() string`

GetThresholdExpression returns the ThresholdExpression field if non-nil, zero value otherwise.

### GetThresholdExpressionOk

`func (o *AlertRuleBaseCreateRequest) GetThresholdExpressionOk() (*string, bool)`

GetThresholdExpressionOk returns a tuple with the ThresholdExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdExpression

`func (o *AlertRuleBaseCreateRequest) SetThresholdExpression(v string)`

SetThresholdExpression sets ThresholdExpression field to given value.

### HasThresholdExpression

`func (o *AlertRuleBaseCreateRequest) HasThresholdExpression() bool`

HasThresholdExpression returns a boolean if a field has been set.

### SetThresholdExpressionNil

`func (o *AlertRuleBaseCreateRequest) SetThresholdExpressionNil(b bool)`

 SetThresholdExpressionNil sets the value for ThresholdExpression to be an explicit nil

### UnsetThresholdExpression
`func (o *AlertRuleBaseCreateRequest) UnsetThresholdExpression()`

UnsetThresholdExpression ensures that no value is present for ThresholdExpression, not even an explicit nil
### GetThresholdFixedValue

`func (o *AlertRuleBaseCreateRequest) GetThresholdFixedValue() float32`

GetThresholdFixedValue returns the ThresholdFixedValue field if non-nil, zero value otherwise.

### GetThresholdFixedValueOk

`func (o *AlertRuleBaseCreateRequest) GetThresholdFixedValueOk() (*float32, bool)`

GetThresholdFixedValueOk returns a tuple with the ThresholdFixedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdFixedValue

`func (o *AlertRuleBaseCreateRequest) SetThresholdFixedValue(v float32)`

SetThresholdFixedValue sets ThresholdFixedValue field to given value.

### HasThresholdFixedValue

`func (o *AlertRuleBaseCreateRequest) HasThresholdFixedValue() bool`

HasThresholdFixedValue returns a boolean if a field has been set.

### SetThresholdFixedValueNil

`func (o *AlertRuleBaseCreateRequest) SetThresholdFixedValueNil(b bool)`

 SetThresholdFixedValueNil sets the value for ThresholdFixedValue to be an explicit nil

### UnsetThresholdFixedValue
`func (o *AlertRuleBaseCreateRequest) UnsetThresholdFixedValue()`

UnsetThresholdFixedValue ensures that no value is present for ThresholdFixedValue, not even an explicit nil
### GetThresholdOperator

`func (o *AlertRuleBaseCreateRequest) GetThresholdOperator() AlertRuleOperator`

GetThresholdOperator returns the ThresholdOperator field if non-nil, zero value otherwise.

### GetThresholdOperatorOk

`func (o *AlertRuleBaseCreateRequest) GetThresholdOperatorOk() (*AlertRuleOperator, bool)`

GetThresholdOperatorOk returns a tuple with the ThresholdOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdOperator

`func (o *AlertRuleBaseCreateRequest) SetThresholdOperator(v AlertRuleOperator)`

SetThresholdOperator sets ThresholdOperator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


