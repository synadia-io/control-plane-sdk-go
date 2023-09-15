# SubjectExportCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JwtSettings** | [**Export**](Export.md) |  | 
**MetricsEnabled** | **bool** |  | 
**MetricsSamplingPercentage** | **int64** |  | 

## Methods

### NewSubjectExportCreateRequest

`func NewSubjectExportCreateRequest(jwtSettings Export, metricsEnabled bool, metricsSamplingPercentage int64, ) *SubjectExportCreateRequest`

NewSubjectExportCreateRequest instantiates a new SubjectExportCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectExportCreateRequestWithDefaults

`func NewSubjectExportCreateRequestWithDefaults() *SubjectExportCreateRequest`

NewSubjectExportCreateRequestWithDefaults instantiates a new SubjectExportCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwtSettings

`func (o *SubjectExportCreateRequest) GetJwtSettings() Export`

GetJwtSettings returns the JwtSettings field if non-nil, zero value otherwise.

### GetJwtSettingsOk

`func (o *SubjectExportCreateRequest) GetJwtSettingsOk() (*Export, bool)`

GetJwtSettingsOk returns a tuple with the JwtSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSettings

`func (o *SubjectExportCreateRequest) SetJwtSettings(v Export)`

SetJwtSettings sets JwtSettings field to given value.


### GetMetricsEnabled

`func (o *SubjectExportCreateRequest) GetMetricsEnabled() bool`

GetMetricsEnabled returns the MetricsEnabled field if non-nil, zero value otherwise.

### GetMetricsEnabledOk

`func (o *SubjectExportCreateRequest) GetMetricsEnabledOk() (*bool, bool)`

GetMetricsEnabledOk returns a tuple with the MetricsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsEnabled

`func (o *SubjectExportCreateRequest) SetMetricsEnabled(v bool)`

SetMetricsEnabled sets MetricsEnabled field to given value.


### GetMetricsSamplingPercentage

`func (o *SubjectExportCreateRequest) GetMetricsSamplingPercentage() int64`

GetMetricsSamplingPercentage returns the MetricsSamplingPercentage field if non-nil, zero value otherwise.

### GetMetricsSamplingPercentageOk

`func (o *SubjectExportCreateRequest) GetMetricsSamplingPercentageOk() (*int64, bool)`

GetMetricsSamplingPercentageOk returns a tuple with the MetricsSamplingPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsSamplingPercentage

`func (o *SubjectExportCreateRequest) SetMetricsSamplingPercentage(v int64)`

SetMetricsSamplingPercentage sets MetricsSamplingPercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


