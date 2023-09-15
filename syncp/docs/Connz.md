# Connz

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | [**[]ConnzConnectionsInner**](ConnzConnectionsInner.md) |  | 
**Limit** | **int32** |  | 
**Now** | **time.Time** |  | 
**NumConnections** | **int32** |  | 
**Offset** | **int32** |  | 
**ServerId** | **string** |  | 
**Total** | **int32** |  | 

## Methods

### NewConnz

`func NewConnz(connections []ConnzConnectionsInner, limit int32, now time.Time, numConnections int32, offset int32, serverId string, total int32, ) *Connz`

NewConnz instantiates a new Connz object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnzWithDefaults

`func NewConnzWithDefaults() *Connz`

NewConnzWithDefaults instantiates a new Connz object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *Connz) GetConnections() []ConnzConnectionsInner`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *Connz) GetConnectionsOk() (*[]ConnzConnectionsInner, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *Connz) SetConnections(v []ConnzConnectionsInner)`

SetConnections sets Connections field to given value.


### GetLimit

`func (o *Connz) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Connz) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Connz) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNow

`func (o *Connz) GetNow() time.Time`

GetNow returns the Now field if non-nil, zero value otherwise.

### GetNowOk

`func (o *Connz) GetNowOk() (*time.Time, bool)`

GetNowOk returns a tuple with the Now field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNow

`func (o *Connz) SetNow(v time.Time)`

SetNow sets Now field to given value.


### GetNumConnections

`func (o *Connz) GetNumConnections() int32`

GetNumConnections returns the NumConnections field if non-nil, zero value otherwise.

### GetNumConnectionsOk

`func (o *Connz) GetNumConnectionsOk() (*int32, bool)`

GetNumConnectionsOk returns a tuple with the NumConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumConnections

`func (o *Connz) SetNumConnections(v int32)`

SetNumConnections sets NumConnections field to given value.


### GetOffset

`func (o *Connz) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Connz) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Connz) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetServerId

`func (o *Connz) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *Connz) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *Connz) SetServerId(v string)`

SetServerId sets ServerId field to given value.


### GetTotal

`func (o *Connz) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Connz) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Connz) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


