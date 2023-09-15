# ApexSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ApexCoordinate**](ApexCoordinate.md) |  | 
**Name** | **string** |  | 

## Methods

### NewApexSeries

`func NewApexSeries(data []ApexCoordinate, name string, ) *ApexSeries`

NewApexSeries instantiates a new ApexSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApexSeriesWithDefaults

`func NewApexSeriesWithDefaults() *ApexSeries`

NewApexSeriesWithDefaults instantiates a new ApexSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApexSeries) GetData() []ApexCoordinate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApexSeries) GetDataOk() (*[]ApexCoordinate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApexSeries) SetData(v []ApexCoordinate)`

SetData sets Data field to given value.


### GetName

`func (o *ApexSeries) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApexSeries) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApexSeries) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


