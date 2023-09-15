# SubjectExportViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**AccountInfo**](AccountInfo.md) |  | 
**Created** | **time.Time** |  | 
**Id** | **string** |  | 
**IsPublic** | **bool** |  | 
**JwtSettings** | [**Export**](Export.md) |  | 
**MetricsEnabled** | **bool** |  | 
**MetricsSamplingPercentage** | **int64** |  | 
**Name** | **NullableString** |  | 
**ShareSkPublic** | **NullableString** |  | 
**StreamExportId** | **NullableString** |  | 
**Subject** | **string** |  | 

## Methods

### NewSubjectExportViewResponse

`func NewSubjectExportViewResponse(account AccountInfo, created time.Time, id string, isPublic bool, jwtSettings Export, metricsEnabled bool, metricsSamplingPercentage int64, name NullableString, shareSkPublic NullableString, streamExportId NullableString, subject string, ) *SubjectExportViewResponse`

NewSubjectExportViewResponse instantiates a new SubjectExportViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectExportViewResponseWithDefaults

`func NewSubjectExportViewResponseWithDefaults() *SubjectExportViewResponse`

NewSubjectExportViewResponseWithDefaults instantiates a new SubjectExportViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *SubjectExportViewResponse) GetAccount() AccountInfo`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SubjectExportViewResponse) GetAccountOk() (*AccountInfo, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SubjectExportViewResponse) SetAccount(v AccountInfo)`

SetAccount sets Account field to given value.


### GetCreated

`func (o *SubjectExportViewResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SubjectExportViewResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SubjectExportViewResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *SubjectExportViewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubjectExportViewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubjectExportViewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsPublic

`func (o *SubjectExportViewResponse) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *SubjectExportViewResponse) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *SubjectExportViewResponse) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetJwtSettings

`func (o *SubjectExportViewResponse) GetJwtSettings() Export`

GetJwtSettings returns the JwtSettings field if non-nil, zero value otherwise.

### GetJwtSettingsOk

`func (o *SubjectExportViewResponse) GetJwtSettingsOk() (*Export, bool)`

GetJwtSettingsOk returns a tuple with the JwtSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSettings

`func (o *SubjectExportViewResponse) SetJwtSettings(v Export)`

SetJwtSettings sets JwtSettings field to given value.


### GetMetricsEnabled

`func (o *SubjectExportViewResponse) GetMetricsEnabled() bool`

GetMetricsEnabled returns the MetricsEnabled field if non-nil, zero value otherwise.

### GetMetricsEnabledOk

`func (o *SubjectExportViewResponse) GetMetricsEnabledOk() (*bool, bool)`

GetMetricsEnabledOk returns a tuple with the MetricsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsEnabled

`func (o *SubjectExportViewResponse) SetMetricsEnabled(v bool)`

SetMetricsEnabled sets MetricsEnabled field to given value.


### GetMetricsSamplingPercentage

`func (o *SubjectExportViewResponse) GetMetricsSamplingPercentage() int64`

GetMetricsSamplingPercentage returns the MetricsSamplingPercentage field if non-nil, zero value otherwise.

### GetMetricsSamplingPercentageOk

`func (o *SubjectExportViewResponse) GetMetricsSamplingPercentageOk() (*int64, bool)`

GetMetricsSamplingPercentageOk returns a tuple with the MetricsSamplingPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsSamplingPercentage

`func (o *SubjectExportViewResponse) SetMetricsSamplingPercentage(v int64)`

SetMetricsSamplingPercentage sets MetricsSamplingPercentage field to given value.


### GetName

`func (o *SubjectExportViewResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubjectExportViewResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubjectExportViewResponse) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *SubjectExportViewResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SubjectExportViewResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetShareSkPublic

`func (o *SubjectExportViewResponse) GetShareSkPublic() string`

GetShareSkPublic returns the ShareSkPublic field if non-nil, zero value otherwise.

### GetShareSkPublicOk

`func (o *SubjectExportViewResponse) GetShareSkPublicOk() (*string, bool)`

GetShareSkPublicOk returns a tuple with the ShareSkPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareSkPublic

`func (o *SubjectExportViewResponse) SetShareSkPublic(v string)`

SetShareSkPublic sets ShareSkPublic field to given value.


### SetShareSkPublicNil

`func (o *SubjectExportViewResponse) SetShareSkPublicNil(b bool)`

 SetShareSkPublicNil sets the value for ShareSkPublic to be an explicit nil

### UnsetShareSkPublic
`func (o *SubjectExportViewResponse) UnsetShareSkPublic()`

UnsetShareSkPublic ensures that no value is present for ShareSkPublic, not even an explicit nil
### GetStreamExportId

`func (o *SubjectExportViewResponse) GetStreamExportId() string`

GetStreamExportId returns the StreamExportId field if non-nil, zero value otherwise.

### GetStreamExportIdOk

`func (o *SubjectExportViewResponse) GetStreamExportIdOk() (*string, bool)`

GetStreamExportIdOk returns a tuple with the StreamExportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamExportId

`func (o *SubjectExportViewResponse) SetStreamExportId(v string)`

SetStreamExportId sets StreamExportId field to given value.


### SetStreamExportIdNil

`func (o *SubjectExportViewResponse) SetStreamExportIdNil(b bool)`

 SetStreamExportIdNil sets the value for StreamExportId to be an explicit nil

### UnsetStreamExportId
`func (o *SubjectExportViewResponse) UnsetStreamExportId()`

UnsetStreamExportId ensures that no value is present for StreamExportId, not even an explicit nil
### GetSubject

`func (o *SubjectExportViewResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SubjectExportViewResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SubjectExportViewResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


