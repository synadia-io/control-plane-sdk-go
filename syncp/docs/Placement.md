# Placement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPlacement

`func NewPlacement(cluster string, ) *Placement`

NewPlacement instantiates a new Placement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementWithDefaults

`func NewPlacementWithDefaults() *Placement`

NewPlacementWithDefaults instantiates a new Placement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *Placement) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Placement) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Placement) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetTags

`func (o *Placement) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Placement) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Placement) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Placement) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

