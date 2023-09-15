# WeightedMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to **string** |  | [optional] 
**Subject** | **string** | Subject is a string that represents a NATS subject | 
**Weight** | Pointer to **int32** |  | [optional] 

## Methods

### NewWeightedMapping

`func NewWeightedMapping(subject string, ) *WeightedMapping`

NewWeightedMapping instantiates a new WeightedMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeightedMappingWithDefaults

`func NewWeightedMappingWithDefaults() *WeightedMapping`

NewWeightedMappingWithDefaults instantiates a new WeightedMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *WeightedMapping) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *WeightedMapping) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *WeightedMapping) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *WeightedMapping) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetSubject

`func (o *WeightedMapping) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *WeightedMapping) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *WeightedMapping) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetWeight

`func (o *WeightedMapping) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *WeightedMapping) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *WeightedMapping) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *WeightedMapping) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


