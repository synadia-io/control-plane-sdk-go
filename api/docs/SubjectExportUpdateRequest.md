# SubjectExportUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JwtSettings** | Pointer to [**Export**](Export.md) |  | [optional] 
**MetricsEnabled** | Pointer to **bool** |  | [optional] 
**MetricsSamplingPercentage** | Pointer to **int64** |  | [optional] 

## Methods

### NewSubjectExportUpdateRequest

`func NewSubjectExportUpdateRequest() *SubjectExportUpdateRequest`

NewSubjectExportUpdateRequest instantiates a new SubjectExportUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectExportUpdateRequestWithDefaults

`func NewSubjectExportUpdateRequestWithDefaults() *SubjectExportUpdateRequest`

NewSubjectExportUpdateRequestWithDefaults instantiates a new SubjectExportUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwtSettings

`func (o *SubjectExportUpdateRequest) GetJwtSettings() Export`

GetJwtSettings returns the JwtSettings field if non-nil, zero value otherwise.

### GetJwtSettingsOk

`func (o *SubjectExportUpdateRequest) GetJwtSettingsOk() (*Export, bool)`

GetJwtSettingsOk returns a tuple with the JwtSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSettings

`func (o *SubjectExportUpdateRequest) SetJwtSettings(v Export)`

SetJwtSettings sets JwtSettings field to given value.

### HasJwtSettings

`func (o *SubjectExportUpdateRequest) HasJwtSettings() bool`

HasJwtSettings returns a boolean if a field has been set.

### GetMetricsEnabled

`func (o *SubjectExportUpdateRequest) GetMetricsEnabled() bool`

GetMetricsEnabled returns the MetricsEnabled field if non-nil, zero value otherwise.

### GetMetricsEnabledOk

`func (o *SubjectExportUpdateRequest) GetMetricsEnabledOk() (*bool, bool)`

GetMetricsEnabledOk returns a tuple with the MetricsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsEnabled

`func (o *SubjectExportUpdateRequest) SetMetricsEnabled(v bool)`

SetMetricsEnabled sets MetricsEnabled field to given value.

### HasMetricsEnabled

`func (o *SubjectExportUpdateRequest) HasMetricsEnabled() bool`

HasMetricsEnabled returns a boolean if a field has been set.

### GetMetricsSamplingPercentage

`func (o *SubjectExportUpdateRequest) GetMetricsSamplingPercentage() int64`

GetMetricsSamplingPercentage returns the MetricsSamplingPercentage field if non-nil, zero value otherwise.

### GetMetricsSamplingPercentageOk

`func (o *SubjectExportUpdateRequest) GetMetricsSamplingPercentageOk() (*int64, bool)`

GetMetricsSamplingPercentageOk returns a tuple with the MetricsSamplingPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsSamplingPercentage

`func (o *SubjectExportUpdateRequest) SetMetricsSamplingPercentage(v int64)`

SetMetricsSamplingPercentage sets MetricsSamplingPercentage field to given value.

### HasMetricsSamplingPercentage

`func (o *SubjectExportUpdateRequest) HasMetricsSamplingPercentage() bool`

HasMetricsSamplingPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


