# GatewayStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gwid** | **int32** |  | 
**InboundConnections** | **int32** |  | 
**Name** | **string** |  | 
**Received** | [**DataStats**](DataStats.md) |  | 
**Sent** | [**DataStats**](DataStats.md) |  | 

## Methods

### NewGatewayStat

`func NewGatewayStat(gwid int32, inboundConnections int32, name string, received DataStats, sent DataStats, ) *GatewayStat`

NewGatewayStat instantiates a new GatewayStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayStatWithDefaults

`func NewGatewayStatWithDefaults() *GatewayStat`

NewGatewayStatWithDefaults instantiates a new GatewayStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGwid

`func (o *GatewayStat) GetGwid() int32`

GetGwid returns the Gwid field if non-nil, zero value otherwise.

### GetGwidOk

`func (o *GatewayStat) GetGwidOk() (*int32, bool)`

GetGwidOk returns a tuple with the Gwid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwid

`func (o *GatewayStat) SetGwid(v int32)`

SetGwid sets Gwid field to given value.


### GetInboundConnections

`func (o *GatewayStat) GetInboundConnections() int32`

GetInboundConnections returns the InboundConnections field if non-nil, zero value otherwise.

### GetInboundConnectionsOk

`func (o *GatewayStat) GetInboundConnectionsOk() (*int32, bool)`

GetInboundConnectionsOk returns a tuple with the InboundConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundConnections

`func (o *GatewayStat) SetInboundConnections(v int32)`

SetInboundConnections sets InboundConnections field to given value.


### GetName

`func (o *GatewayStat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayStat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayStat) SetName(v string)`

SetName sets Name field to given value.


### GetReceived

`func (o *GatewayStat) GetReceived() DataStats`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *GatewayStat) GetReceivedOk() (*DataStats, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *GatewayStat) SetReceived(v DataStats)`

SetReceived sets Received field to given value.


### GetSent

`func (o *GatewayStat) GetSent() DataStats`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *GatewayStat) GetSentOk() (*DataStats, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *GatewayStat) SetSent(v DataStats)`

SetSent sets Sent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


