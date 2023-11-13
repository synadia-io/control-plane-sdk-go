# ExportAllOfServiceLatency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | **string** | Subject is a string that represents a NATS subject | 
**Sampling** | **map[string]interface{}** |  | 

## Methods

### NewExportAllOfServiceLatency

`func NewExportAllOfServiceLatency(results string, sampling map[string]interface{}, ) *ExportAllOfServiceLatency`

NewExportAllOfServiceLatency instantiates a new ExportAllOfServiceLatency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportAllOfServiceLatencyWithDefaults

`func NewExportAllOfServiceLatencyWithDefaults() *ExportAllOfServiceLatency`

NewExportAllOfServiceLatencyWithDefaults instantiates a new ExportAllOfServiceLatency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ExportAllOfServiceLatency) GetResults() string`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ExportAllOfServiceLatency) GetResultsOk() (*string, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ExportAllOfServiceLatency) SetResults(v string)`

SetResults sets Results field to given value.


### GetSampling

`func (o *ExportAllOfServiceLatency) GetSampling() map[string]interface{}`

GetSampling returns the Sampling field if non-nil, zero value otherwise.

### GetSamplingOk

`func (o *ExportAllOfServiceLatency) GetSamplingOk() (*map[string]interface{}, bool)`

GetSamplingOk returns a tuple with the Sampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampling

`func (o *ExportAllOfServiceLatency) SetSampling(v map[string]interface{})`

SetSampling sets Sampling field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


