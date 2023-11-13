# AlertRuleAccountCreateRequest

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
**ConsumerName** | Pointer to **NullableString** |  | [optional] 
**RuleType** | **string** |  | 
**StreamName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAlertRuleAccountCreateRequest

`func NewAlertRuleAccountCreateRequest(durationInSecs int32, message string, metric string, severity AlertRuleSeverity, thresholdOperator AlertRuleOperator, ruleType string, ) *AlertRuleAccountCreateRequest`

NewAlertRuleAccountCreateRequest instantiates a new AlertRuleAccountCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleAccountCreateRequestWithDefaults

`func NewAlertRuleAccountCreateRequestWithDefaults() *AlertRuleAccountCreateRequest`

NewAlertRuleAccountCreateRequestWithDefaults instantiates a new AlertRuleAccountCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationInSecs

`func (o *AlertRuleAccountCreateRequest) GetDurationInSecs() int32`

GetDurationInSecs returns the DurationInSecs field if non-nil, zero value otherwise.

### GetDurationInSecsOk

`func (o *AlertRuleAccountCreateRequest) GetDurationInSecsOk() (*int32, bool)`

GetDurationInSecsOk returns a tuple with the DurationInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInSecs

`func (o *AlertRuleAccountCreateRequest) SetDurationInSecs(v int32)`

SetDurationInSecs sets DurationInSecs field to given value.


### GetMessage

`func (o *AlertRuleAccountCreateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertRuleAccountCreateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertRuleAccountCreateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMetric

`func (o *AlertRuleAccountCreateRequest) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *AlertRuleAccountCreateRequest) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *AlertRuleAccountCreateRequest) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetSeverity

`func (o *AlertRuleAccountCreateRequest) GetSeverity() AlertRuleSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertRuleAccountCreateRequest) GetSeverityOk() (*AlertRuleSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertRuleAccountCreateRequest) SetSeverity(v AlertRuleSeverity)`

SetSeverity sets Severity field to given value.


### GetThresholdExpression

`func (o *AlertRuleAccountCreateRequest) GetThresholdExpression() string`

GetThresholdExpression returns the ThresholdExpression field if non-nil, zero value otherwise.

### GetThresholdExpressionOk

`func (o *AlertRuleAccountCreateRequest) GetThresholdExpressionOk() (*string, bool)`

GetThresholdExpressionOk returns a tuple with the ThresholdExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdExpression

`func (o *AlertRuleAccountCreateRequest) SetThresholdExpression(v string)`

SetThresholdExpression sets ThresholdExpression field to given value.

### HasThresholdExpression

`func (o *AlertRuleAccountCreateRequest) HasThresholdExpression() bool`

HasThresholdExpression returns a boolean if a field has been set.

### SetThresholdExpressionNil

`func (o *AlertRuleAccountCreateRequest) SetThresholdExpressionNil(b bool)`

 SetThresholdExpressionNil sets the value for ThresholdExpression to be an explicit nil

### UnsetThresholdExpression
`func (o *AlertRuleAccountCreateRequest) UnsetThresholdExpression()`

UnsetThresholdExpression ensures that no value is present for ThresholdExpression, not even an explicit nil
### GetThresholdFixedValue

`func (o *AlertRuleAccountCreateRequest) GetThresholdFixedValue() float32`

GetThresholdFixedValue returns the ThresholdFixedValue field if non-nil, zero value otherwise.

### GetThresholdFixedValueOk

`func (o *AlertRuleAccountCreateRequest) GetThresholdFixedValueOk() (*float32, bool)`

GetThresholdFixedValueOk returns a tuple with the ThresholdFixedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdFixedValue

`func (o *AlertRuleAccountCreateRequest) SetThresholdFixedValue(v float32)`

SetThresholdFixedValue sets ThresholdFixedValue field to given value.

### HasThresholdFixedValue

`func (o *AlertRuleAccountCreateRequest) HasThresholdFixedValue() bool`

HasThresholdFixedValue returns a boolean if a field has been set.

### SetThresholdFixedValueNil

`func (o *AlertRuleAccountCreateRequest) SetThresholdFixedValueNil(b bool)`

 SetThresholdFixedValueNil sets the value for ThresholdFixedValue to be an explicit nil

### UnsetThresholdFixedValue
`func (o *AlertRuleAccountCreateRequest) UnsetThresholdFixedValue()`

UnsetThresholdFixedValue ensures that no value is present for ThresholdFixedValue, not even an explicit nil
### GetThresholdOperator

`func (o *AlertRuleAccountCreateRequest) GetThresholdOperator() AlertRuleOperator`

GetThresholdOperator returns the ThresholdOperator field if non-nil, zero value otherwise.

### GetThresholdOperatorOk

`func (o *AlertRuleAccountCreateRequest) GetThresholdOperatorOk() (*AlertRuleOperator, bool)`

GetThresholdOperatorOk returns a tuple with the ThresholdOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdOperator

`func (o *AlertRuleAccountCreateRequest) SetThresholdOperator(v AlertRuleOperator)`

SetThresholdOperator sets ThresholdOperator field to given value.


### GetConsumerName

`func (o *AlertRuleAccountCreateRequest) GetConsumerName() string`

GetConsumerName returns the ConsumerName field if non-nil, zero value otherwise.

### GetConsumerNameOk

`func (o *AlertRuleAccountCreateRequest) GetConsumerNameOk() (*string, bool)`

GetConsumerNameOk returns a tuple with the ConsumerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerName

`func (o *AlertRuleAccountCreateRequest) SetConsumerName(v string)`

SetConsumerName sets ConsumerName field to given value.

### HasConsumerName

`func (o *AlertRuleAccountCreateRequest) HasConsumerName() bool`

HasConsumerName returns a boolean if a field has been set.

### SetConsumerNameNil

`func (o *AlertRuleAccountCreateRequest) SetConsumerNameNil(b bool)`

 SetConsumerNameNil sets the value for ConsumerName to be an explicit nil

### UnsetConsumerName
`func (o *AlertRuleAccountCreateRequest) UnsetConsumerName()`

UnsetConsumerName ensures that no value is present for ConsumerName, not even an explicit nil
### GetRuleType

`func (o *AlertRuleAccountCreateRequest) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *AlertRuleAccountCreateRequest) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *AlertRuleAccountCreateRequest) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.


### GetStreamName

`func (o *AlertRuleAccountCreateRequest) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *AlertRuleAccountCreateRequest) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *AlertRuleAccountCreateRequest) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.

### HasStreamName

`func (o *AlertRuleAccountCreateRequest) HasStreamName() bool`

HasStreamName returns a boolean if a field has been set.

### SetStreamNameNil

`func (o *AlertRuleAccountCreateRequest) SetStreamNameNil(b bool)`

 SetStreamNameNil sets the value for StreamName to be an explicit nil

### UnsetStreamName
`func (o *AlertRuleAccountCreateRequest) UnsetStreamName()`

UnsetStreamName ensures that no value is present for StreamName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


