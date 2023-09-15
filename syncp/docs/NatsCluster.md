# NatsCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | **int32** |  | 
**IncomingGateways** | **int32** |  | 
**Name** | **string** |  | 
**NodeCount** | **int32** |  | 
**OutgoingGateways** | **int32** |  | 

## Methods

### NewNatsCluster

`func NewNatsCluster(connections int32, incomingGateways int32, name string, nodeCount int32, outgoingGateways int32, ) *NatsCluster`

NewNatsCluster instantiates a new NatsCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatsClusterWithDefaults

`func NewNatsClusterWithDefaults() *NatsCluster`

NewNatsClusterWithDefaults instantiates a new NatsCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *NatsCluster) GetConnections() int32`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *NatsCluster) GetConnectionsOk() (*int32, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *NatsCluster) SetConnections(v int32)`

SetConnections sets Connections field to given value.


### GetIncomingGateways

`func (o *NatsCluster) GetIncomingGateways() int32`

GetIncomingGateways returns the IncomingGateways field if non-nil, zero value otherwise.

### GetIncomingGatewaysOk

`func (o *NatsCluster) GetIncomingGatewaysOk() (*int32, bool)`

GetIncomingGatewaysOk returns a tuple with the IncomingGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingGateways

`func (o *NatsCluster) SetIncomingGateways(v int32)`

SetIncomingGateways sets IncomingGateways field to given value.


### GetName

`func (o *NatsCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NatsCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NatsCluster) SetName(v string)`

SetName sets Name field to given value.


### GetNodeCount

`func (o *NatsCluster) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *NatsCluster) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *NatsCluster) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### GetOutgoingGateways

`func (o *NatsCluster) GetOutgoingGateways() int32`

GetOutgoingGateways returns the OutgoingGateways field if non-nil, zero value otherwise.

### GetOutgoingGatewaysOk

`func (o *NatsCluster) GetOutgoingGatewaysOk() (*int32, bool)`

GetOutgoingGatewaysOk returns a tuple with the OutgoingGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingGateways

`func (o *NatsCluster) SetOutgoingGateways(v int32)`

SetOutgoingGateways sets OutgoingGateways field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


