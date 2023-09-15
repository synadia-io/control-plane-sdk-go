# AlertViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcknowledgedByAppUserID** | Pointer to **NullableString** |  | [optional] 
**Account** | Pointer to [**NullableAlertViewResponseAccount**](AlertViewResponseAccount.md) |  | [optional] 
**AccountId** | Pointer to **NullableString** |  | [optional] 
**AcknowledgedAt** | Pointer to **NullableString** |  | [optional] 
**AlertRule** | [**AlertRuleViewResponse**](AlertRuleViewResponse.md) |  | 
**AlertRuleId** | **string** |  | 
**ClosedAt** | Pointer to **NullableString** |  | [optional] 
**ConsumerName** | Pointer to **NullableString** |  | [optional] 
**Created** | **time.Time** |  | 
**Id** | **string** |  | 
**IsAcknowledged** | **bool** |  | 
**IsClosed** | **bool** |  | 
**IsFiring** | **bool** |  | 
**Message** | **string** |  | 
**MetricValue** | **float64** |  | 
**SampleValues** | [**[]PromSampleValue**](PromSampleValue.md) |  | 
**Severity** | [**AlertRuleSeverity**](AlertRuleSeverity.md) |  | 
**StreamName** | Pointer to **NullableString** |  | [optional] 
**System** | Pointer to [**NullableAlertViewResponseSystem**](AlertViewResponseSystem.md) |  | [optional] 
**SystemId** | Pointer to **NullableString** |  | [optional] 
**ThresholdValue** | **float64** |  | 

## Methods

### NewAlertViewResponse

`func NewAlertViewResponse(alertRule AlertRuleViewResponse, alertRuleId string, created time.Time, id string, isAcknowledged bool, isClosed bool, isFiring bool, message string, metricValue float64, sampleValues []PromSampleValue, severity AlertRuleSeverity, thresholdValue float64, ) *AlertViewResponse`

NewAlertViewResponse instantiates a new AlertViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertViewResponseWithDefaults

`func NewAlertViewResponseWithDefaults() *AlertViewResponse`

NewAlertViewResponseWithDefaults instantiates a new AlertViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledgedByAppUserID

`func (o *AlertViewResponse) GetAcknowledgedByAppUserID() string`

GetAcknowledgedByAppUserID returns the AcknowledgedByAppUserID field if non-nil, zero value otherwise.

### GetAcknowledgedByAppUserIDOk

`func (o *AlertViewResponse) GetAcknowledgedByAppUserIDOk() (*string, bool)`

GetAcknowledgedByAppUserIDOk returns a tuple with the AcknowledgedByAppUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedByAppUserID

`func (o *AlertViewResponse) SetAcknowledgedByAppUserID(v string)`

SetAcknowledgedByAppUserID sets AcknowledgedByAppUserID field to given value.

### HasAcknowledgedByAppUserID

`func (o *AlertViewResponse) HasAcknowledgedByAppUserID() bool`

HasAcknowledgedByAppUserID returns a boolean if a field has been set.

### SetAcknowledgedByAppUserIDNil

`func (o *AlertViewResponse) SetAcknowledgedByAppUserIDNil(b bool)`

 SetAcknowledgedByAppUserIDNil sets the value for AcknowledgedByAppUserID to be an explicit nil

### UnsetAcknowledgedByAppUserID
`func (o *AlertViewResponse) UnsetAcknowledgedByAppUserID()`

UnsetAcknowledgedByAppUserID ensures that no value is present for AcknowledgedByAppUserID, not even an explicit nil
### GetAccount

`func (o *AlertViewResponse) GetAccount() AlertViewResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AlertViewResponse) GetAccountOk() (*AlertViewResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AlertViewResponse) SetAccount(v AlertViewResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AlertViewResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *AlertViewResponse) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *AlertViewResponse) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetAccountId

`func (o *AlertViewResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AlertViewResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AlertViewResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AlertViewResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *AlertViewResponse) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *AlertViewResponse) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetAcknowledgedAt

`func (o *AlertViewResponse) GetAcknowledgedAt() string`

GetAcknowledgedAt returns the AcknowledgedAt field if non-nil, zero value otherwise.

### GetAcknowledgedAtOk

`func (o *AlertViewResponse) GetAcknowledgedAtOk() (*string, bool)`

GetAcknowledgedAtOk returns a tuple with the AcknowledgedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgedAt

`func (o *AlertViewResponse) SetAcknowledgedAt(v string)`

SetAcknowledgedAt sets AcknowledgedAt field to given value.

### HasAcknowledgedAt

`func (o *AlertViewResponse) HasAcknowledgedAt() bool`

HasAcknowledgedAt returns a boolean if a field has been set.

### SetAcknowledgedAtNil

`func (o *AlertViewResponse) SetAcknowledgedAtNil(b bool)`

 SetAcknowledgedAtNil sets the value for AcknowledgedAt to be an explicit nil

### UnsetAcknowledgedAt
`func (o *AlertViewResponse) UnsetAcknowledgedAt()`

UnsetAcknowledgedAt ensures that no value is present for AcknowledgedAt, not even an explicit nil
### GetAlertRule

`func (o *AlertViewResponse) GetAlertRule() AlertRuleViewResponse`

GetAlertRule returns the AlertRule field if non-nil, zero value otherwise.

### GetAlertRuleOk

`func (o *AlertViewResponse) GetAlertRuleOk() (*AlertRuleViewResponse, bool)`

GetAlertRuleOk returns a tuple with the AlertRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRule

`func (o *AlertViewResponse) SetAlertRule(v AlertRuleViewResponse)`

SetAlertRule sets AlertRule field to given value.


### GetAlertRuleId

`func (o *AlertViewResponse) GetAlertRuleId() string`

GetAlertRuleId returns the AlertRuleId field if non-nil, zero value otherwise.

### GetAlertRuleIdOk

`func (o *AlertViewResponse) GetAlertRuleIdOk() (*string, bool)`

GetAlertRuleIdOk returns a tuple with the AlertRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRuleId

`func (o *AlertViewResponse) SetAlertRuleId(v string)`

SetAlertRuleId sets AlertRuleId field to given value.


### GetClosedAt

`func (o *AlertViewResponse) GetClosedAt() string`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *AlertViewResponse) GetClosedAtOk() (*string, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *AlertViewResponse) SetClosedAt(v string)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *AlertViewResponse) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.

### SetClosedAtNil

`func (o *AlertViewResponse) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *AlertViewResponse) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetConsumerName

`func (o *AlertViewResponse) GetConsumerName() string`

GetConsumerName returns the ConsumerName field if non-nil, zero value otherwise.

### GetConsumerNameOk

`func (o *AlertViewResponse) GetConsumerNameOk() (*string, bool)`

GetConsumerNameOk returns a tuple with the ConsumerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerName

`func (o *AlertViewResponse) SetConsumerName(v string)`

SetConsumerName sets ConsumerName field to given value.

### HasConsumerName

`func (o *AlertViewResponse) HasConsumerName() bool`

HasConsumerName returns a boolean if a field has been set.

### SetConsumerNameNil

`func (o *AlertViewResponse) SetConsumerNameNil(b bool)`

 SetConsumerNameNil sets the value for ConsumerName to be an explicit nil

### UnsetConsumerName
`func (o *AlertViewResponse) UnsetConsumerName()`

UnsetConsumerName ensures that no value is present for ConsumerName, not even an explicit nil
### GetCreated

`func (o *AlertViewResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AlertViewResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AlertViewResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *AlertViewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertViewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertViewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsAcknowledged

`func (o *AlertViewResponse) GetIsAcknowledged() bool`

GetIsAcknowledged returns the IsAcknowledged field if non-nil, zero value otherwise.

### GetIsAcknowledgedOk

`func (o *AlertViewResponse) GetIsAcknowledgedOk() (*bool, bool)`

GetIsAcknowledgedOk returns a tuple with the IsAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAcknowledged

`func (o *AlertViewResponse) SetIsAcknowledged(v bool)`

SetIsAcknowledged sets IsAcknowledged field to given value.


### GetIsClosed

`func (o *AlertViewResponse) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *AlertViewResponse) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *AlertViewResponse) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.


### GetIsFiring

`func (o *AlertViewResponse) GetIsFiring() bool`

GetIsFiring returns the IsFiring field if non-nil, zero value otherwise.

### GetIsFiringOk

`func (o *AlertViewResponse) GetIsFiringOk() (*bool, bool)`

GetIsFiringOk returns a tuple with the IsFiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFiring

`func (o *AlertViewResponse) SetIsFiring(v bool)`

SetIsFiring sets IsFiring field to given value.


### GetMessage

`func (o *AlertViewResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertViewResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertViewResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMetricValue

`func (o *AlertViewResponse) GetMetricValue() float64`

GetMetricValue returns the MetricValue field if non-nil, zero value otherwise.

### GetMetricValueOk

`func (o *AlertViewResponse) GetMetricValueOk() (*float64, bool)`

GetMetricValueOk returns a tuple with the MetricValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValue

`func (o *AlertViewResponse) SetMetricValue(v float64)`

SetMetricValue sets MetricValue field to given value.


### GetSampleValues

`func (o *AlertViewResponse) GetSampleValues() []PromSampleValue`

GetSampleValues returns the SampleValues field if non-nil, zero value otherwise.

### GetSampleValuesOk

`func (o *AlertViewResponse) GetSampleValuesOk() (*[]PromSampleValue, bool)`

GetSampleValuesOk returns a tuple with the SampleValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleValues

`func (o *AlertViewResponse) SetSampleValues(v []PromSampleValue)`

SetSampleValues sets SampleValues field to given value.


### GetSeverity

`func (o *AlertViewResponse) GetSeverity() AlertRuleSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertViewResponse) GetSeverityOk() (*AlertRuleSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertViewResponse) SetSeverity(v AlertRuleSeverity)`

SetSeverity sets Severity field to given value.


### GetStreamName

`func (o *AlertViewResponse) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *AlertViewResponse) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *AlertViewResponse) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.

### HasStreamName

`func (o *AlertViewResponse) HasStreamName() bool`

HasStreamName returns a boolean if a field has been set.

### SetStreamNameNil

`func (o *AlertViewResponse) SetStreamNameNil(b bool)`

 SetStreamNameNil sets the value for StreamName to be an explicit nil

### UnsetStreamName
`func (o *AlertViewResponse) UnsetStreamName()`

UnsetStreamName ensures that no value is present for StreamName, not even an explicit nil
### GetSystem

`func (o *AlertViewResponse) GetSystem() AlertViewResponseSystem`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *AlertViewResponse) GetSystemOk() (*AlertViewResponseSystem, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *AlertViewResponse) SetSystem(v AlertViewResponseSystem)`

SetSystem sets System field to given value.

### HasSystem

`func (o *AlertViewResponse) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *AlertViewResponse) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *AlertViewResponse) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil
### GetSystemId

`func (o *AlertViewResponse) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *AlertViewResponse) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *AlertViewResponse) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.

### HasSystemId

`func (o *AlertViewResponse) HasSystemId() bool`

HasSystemId returns a boolean if a field has been set.

### SetSystemIdNil

`func (o *AlertViewResponse) SetSystemIdNil(b bool)`

 SetSystemIdNil sets the value for SystemId to be an explicit nil

### UnsetSystemId
`func (o *AlertViewResponse) UnsetSystemId()`

UnsetSystemId ensures that no value is present for SystemId, not even an explicit nil
### GetThresholdValue

`func (o *AlertViewResponse) GetThresholdValue() float64`

GetThresholdValue returns the ThresholdValue field if non-nil, zero value otherwise.

### GetThresholdValueOk

`func (o *AlertViewResponse) GetThresholdValueOk() (*float64, bool)`

GetThresholdValueOk returns a tuple with the ThresholdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValue

`func (o *AlertViewResponse) SetThresholdValue(v float64)`

SetThresholdValue sets ThresholdValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


