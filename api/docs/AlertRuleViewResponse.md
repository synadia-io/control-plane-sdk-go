# AlertRuleViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **NullableString** |  | [optional] 
**ConsumerName** | Pointer to **NullableString** |  | [optional] 
**Created** | **time.Time** |  | 
**DurationInSecs** | **int32** |  | 
**Id** | **string** |  | 
**IsEnabled** | **bool** |  | 
**Message** | **string** |  | 
**Metric** | **string** |  | 
**ParentRuleId** | Pointer to **NullableString** |  | [optional] 
**RuleType** | [**AlertRuleType**](AlertRuleType.md) |  | 
**Severity** | [**AlertRuleSeverity**](AlertRuleSeverity.md) |  | 
**StreamName** | Pointer to **NullableString** |  | [optional] 
**SystemId** | Pointer to **NullableString** |  | [optional] 
**ThresholdExpression** | Pointer to **NullableString** |  | [optional] 
**ThresholdFixedValue** | Pointer to **NullableFloat32** |  | [optional] 
**ThresholdOperator** | **string** |  | 

## Methods

### NewAlertRuleViewResponse

`func NewAlertRuleViewResponse(created time.Time, durationInSecs int32, id string, isEnabled bool, message string, metric string, ruleType AlertRuleType, severity AlertRuleSeverity, thresholdOperator string, ) *AlertRuleViewResponse`

NewAlertRuleViewResponse instantiates a new AlertRuleViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleViewResponseWithDefaults

`func NewAlertRuleViewResponseWithDefaults() *AlertRuleViewResponse`

NewAlertRuleViewResponseWithDefaults instantiates a new AlertRuleViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AlertRuleViewResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AlertRuleViewResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AlertRuleViewResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AlertRuleViewResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *AlertRuleViewResponse) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *AlertRuleViewResponse) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetConsumerName

`func (o *AlertRuleViewResponse) GetConsumerName() string`

GetConsumerName returns the ConsumerName field if non-nil, zero value otherwise.

### GetConsumerNameOk

`func (o *AlertRuleViewResponse) GetConsumerNameOk() (*string, bool)`

GetConsumerNameOk returns a tuple with the ConsumerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerName

`func (o *AlertRuleViewResponse) SetConsumerName(v string)`

SetConsumerName sets ConsumerName field to given value.

### HasConsumerName

`func (o *AlertRuleViewResponse) HasConsumerName() bool`

HasConsumerName returns a boolean if a field has been set.

### SetConsumerNameNil

`func (o *AlertRuleViewResponse) SetConsumerNameNil(b bool)`

 SetConsumerNameNil sets the value for ConsumerName to be an explicit nil

### UnsetConsumerName
`func (o *AlertRuleViewResponse) UnsetConsumerName()`

UnsetConsumerName ensures that no value is present for ConsumerName, not even an explicit nil
### GetCreated

`func (o *AlertRuleViewResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AlertRuleViewResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AlertRuleViewResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDurationInSecs

`func (o *AlertRuleViewResponse) GetDurationInSecs() int32`

GetDurationInSecs returns the DurationInSecs field if non-nil, zero value otherwise.

### GetDurationInSecsOk

`func (o *AlertRuleViewResponse) GetDurationInSecsOk() (*int32, bool)`

GetDurationInSecsOk returns a tuple with the DurationInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInSecs

`func (o *AlertRuleViewResponse) SetDurationInSecs(v int32)`

SetDurationInSecs sets DurationInSecs field to given value.


### GetId

`func (o *AlertRuleViewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertRuleViewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertRuleViewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsEnabled

`func (o *AlertRuleViewResponse) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AlertRuleViewResponse) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AlertRuleViewResponse) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetMessage

`func (o *AlertRuleViewResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertRuleViewResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertRuleViewResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMetric

`func (o *AlertRuleViewResponse) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *AlertRuleViewResponse) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *AlertRuleViewResponse) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetParentRuleId

`func (o *AlertRuleViewResponse) GetParentRuleId() string`

GetParentRuleId returns the ParentRuleId field if non-nil, zero value otherwise.

### GetParentRuleIdOk

`func (o *AlertRuleViewResponse) GetParentRuleIdOk() (*string, bool)`

GetParentRuleIdOk returns a tuple with the ParentRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRuleId

`func (o *AlertRuleViewResponse) SetParentRuleId(v string)`

SetParentRuleId sets ParentRuleId field to given value.

### HasParentRuleId

`func (o *AlertRuleViewResponse) HasParentRuleId() bool`

HasParentRuleId returns a boolean if a field has been set.

### SetParentRuleIdNil

`func (o *AlertRuleViewResponse) SetParentRuleIdNil(b bool)`

 SetParentRuleIdNil sets the value for ParentRuleId to be an explicit nil

### UnsetParentRuleId
`func (o *AlertRuleViewResponse) UnsetParentRuleId()`

UnsetParentRuleId ensures that no value is present for ParentRuleId, not even an explicit nil
### GetRuleType

`func (o *AlertRuleViewResponse) GetRuleType() AlertRuleType`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *AlertRuleViewResponse) GetRuleTypeOk() (*AlertRuleType, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *AlertRuleViewResponse) SetRuleType(v AlertRuleType)`

SetRuleType sets RuleType field to given value.


### GetSeverity

`func (o *AlertRuleViewResponse) GetSeverity() AlertRuleSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertRuleViewResponse) GetSeverityOk() (*AlertRuleSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertRuleViewResponse) SetSeverity(v AlertRuleSeverity)`

SetSeverity sets Severity field to given value.


### GetStreamName

`func (o *AlertRuleViewResponse) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *AlertRuleViewResponse) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *AlertRuleViewResponse) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.

### HasStreamName

`func (o *AlertRuleViewResponse) HasStreamName() bool`

HasStreamName returns a boolean if a field has been set.

### SetStreamNameNil

`func (o *AlertRuleViewResponse) SetStreamNameNil(b bool)`

 SetStreamNameNil sets the value for StreamName to be an explicit nil

### UnsetStreamName
`func (o *AlertRuleViewResponse) UnsetStreamName()`

UnsetStreamName ensures that no value is present for StreamName, not even an explicit nil
### GetSystemId

`func (o *AlertRuleViewResponse) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *AlertRuleViewResponse) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *AlertRuleViewResponse) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.

### HasSystemId

`func (o *AlertRuleViewResponse) HasSystemId() bool`

HasSystemId returns a boolean if a field has been set.

### SetSystemIdNil

`func (o *AlertRuleViewResponse) SetSystemIdNil(b bool)`

 SetSystemIdNil sets the value for SystemId to be an explicit nil

### UnsetSystemId
`func (o *AlertRuleViewResponse) UnsetSystemId()`

UnsetSystemId ensures that no value is present for SystemId, not even an explicit nil
### GetThresholdExpression

`func (o *AlertRuleViewResponse) GetThresholdExpression() string`

GetThresholdExpression returns the ThresholdExpression field if non-nil, zero value otherwise.

### GetThresholdExpressionOk

`func (o *AlertRuleViewResponse) GetThresholdExpressionOk() (*string, bool)`

GetThresholdExpressionOk returns a tuple with the ThresholdExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdExpression

`func (o *AlertRuleViewResponse) SetThresholdExpression(v string)`

SetThresholdExpression sets ThresholdExpression field to given value.

### HasThresholdExpression

`func (o *AlertRuleViewResponse) HasThresholdExpression() bool`

HasThresholdExpression returns a boolean if a field has been set.

### SetThresholdExpressionNil

`func (o *AlertRuleViewResponse) SetThresholdExpressionNil(b bool)`

 SetThresholdExpressionNil sets the value for ThresholdExpression to be an explicit nil

### UnsetThresholdExpression
`func (o *AlertRuleViewResponse) UnsetThresholdExpression()`

UnsetThresholdExpression ensures that no value is present for ThresholdExpression, not even an explicit nil
### GetThresholdFixedValue

`func (o *AlertRuleViewResponse) GetThresholdFixedValue() float32`

GetThresholdFixedValue returns the ThresholdFixedValue field if non-nil, zero value otherwise.

### GetThresholdFixedValueOk

`func (o *AlertRuleViewResponse) GetThresholdFixedValueOk() (*float32, bool)`

GetThresholdFixedValueOk returns a tuple with the ThresholdFixedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdFixedValue

`func (o *AlertRuleViewResponse) SetThresholdFixedValue(v float32)`

SetThresholdFixedValue sets ThresholdFixedValue field to given value.

### HasThresholdFixedValue

`func (o *AlertRuleViewResponse) HasThresholdFixedValue() bool`

HasThresholdFixedValue returns a boolean if a field has been set.

### SetThresholdFixedValueNil

`func (o *AlertRuleViewResponse) SetThresholdFixedValueNil(b bool)`

 SetThresholdFixedValueNil sets the value for ThresholdFixedValue to be an explicit nil

### UnsetThresholdFixedValue
`func (o *AlertRuleViewResponse) UnsetThresholdFixedValue()`

UnsetThresholdFixedValue ensures that no value is present for ThresholdFixedValue, not even an explicit nil
### GetThresholdOperator

`func (o *AlertRuleViewResponse) GetThresholdOperator() string`

GetThresholdOperator returns the ThresholdOperator field if non-nil, zero value otherwise.

### GetThresholdOperatorOk

`func (o *AlertRuleViewResponse) GetThresholdOperatorOk() (*string, bool)`

GetThresholdOperatorOk returns a tuple with the ThresholdOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdOperator

`func (o *AlertRuleViewResponse) SetThresholdOperator(v string)`

SetThresholdOperator sets ThresholdOperator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


