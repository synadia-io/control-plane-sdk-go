# RouteStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Pending** | **int32** |  | 
**Received** | [**DataStats**](DataStats.md) |  | 
**Rid** | **int32** |  | 
**Sent** | [**DataStats**](DataStats.md) |  | 

## Methods

### NewRouteStat

`func NewRouteStat(pending int32, received DataStats, rid int32, sent DataStats, ) *RouteStat`

NewRouteStat instantiates a new RouteStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteStatWithDefaults

`func NewRouteStatWithDefaults() *RouteStat`

NewRouteStatWithDefaults instantiates a new RouteStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RouteStat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteStat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteStat) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RouteStat) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPending

`func (o *RouteStat) GetPending() int32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *RouteStat) GetPendingOk() (*int32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *RouteStat) SetPending(v int32)`

SetPending sets Pending field to given value.


### GetReceived

`func (o *RouteStat) GetReceived() DataStats`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *RouteStat) GetReceivedOk() (*DataStats, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *RouteStat) SetReceived(v DataStats)`

SetReceived sets Received field to given value.


### GetRid

`func (o *RouteStat) GetRid() int32`

GetRid returns the Rid field if non-nil, zero value otherwise.

### GetRidOk

`func (o *RouteStat) GetRidOk() (*int32, bool)`

GetRidOk returns a tuple with the Rid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRid

`func (o *RouteStat) SetRid(v int32)`

SetRid sets Rid field to given value.


### GetSent

`func (o *RouteStat) GetSent() DataStats`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *RouteStat) GetSentOk() (*DataStats, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *RouteStat) SetSent(v DataStats)`

SetSent sets Sent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


