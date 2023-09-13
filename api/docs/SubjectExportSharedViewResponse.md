# SubjectExportSharedViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportId** | **string** |  | 
**JwtSettings** | [**Import**](Import.md) |  | 

## Methods

### NewSubjectExportSharedViewResponse

`func NewSubjectExportSharedViewResponse(exportId string, jwtSettings Import, ) *SubjectExportSharedViewResponse`

NewSubjectExportSharedViewResponse instantiates a new SubjectExportSharedViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectExportSharedViewResponseWithDefaults

`func NewSubjectExportSharedViewResponseWithDefaults() *SubjectExportSharedViewResponse`

NewSubjectExportSharedViewResponseWithDefaults instantiates a new SubjectExportSharedViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportId

`func (o *SubjectExportSharedViewResponse) GetExportId() string`

GetExportId returns the ExportId field if non-nil, zero value otherwise.

### GetExportIdOk

`func (o *SubjectExportSharedViewResponse) GetExportIdOk() (*string, bool)`

GetExportIdOk returns a tuple with the ExportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportId

`func (o *SubjectExportSharedViewResponse) SetExportId(v string)`

SetExportId sets ExportId field to given value.


### GetJwtSettings

`func (o *SubjectExportSharedViewResponse) GetJwtSettings() Import`

GetJwtSettings returns the JwtSettings field if non-nil, zero value otherwise.

### GetJwtSettingsOk

`func (o *SubjectExportSharedViewResponse) GetJwtSettingsOk() (*Import, bool)`

GetJwtSettingsOk returns a tuple with the JwtSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSettings

`func (o *SubjectExportSharedViewResponse) SetJwtSettings(v Import)`

SetJwtSettings sets JwtSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


