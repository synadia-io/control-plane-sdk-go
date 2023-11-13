# OperatorLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **int64** |  | [optional] 
**Payload** | Pointer to **int64** |  | [optional] 
**Subs** | Pointer to **int64** |  | [optional] 
**Conn** | Pointer to **int64** |  | [optional] 
**DisallowBearer** | Pointer to **bool** |  | [optional] 
**Exports** | Pointer to **int64** |  | [optional] 
**Imports** | Pointer to **int64** |  | [optional] 
**Leaf** | Pointer to **int64** |  | [optional] 
**Wildcards** | Pointer to **bool** |  | [optional] 
**Consumer** | Pointer to **int64** |  | [optional] 
**DiskMaxStreamBytes** | Pointer to **int64** |  | [optional] 
**DiskStorage** | Pointer to **int64** |  | [optional] 
**MaxAckPending** | Pointer to **int64** |  | [optional] 
**MaxBytesRequired** | Pointer to **bool** |  | [optional] 
**MemMaxStreamBytes** | Pointer to **int64** |  | [optional] 
**MemStorage** | Pointer to **int64** |  | [optional] 
**Streams** | Pointer to **int64** |  | [optional] 
**TieredLimits** | Pointer to [**map[string]JetStreamLimits**](JetStreamLimits.md) |  | [optional] 

## Methods

### NewOperatorLimits

`func NewOperatorLimits() *OperatorLimits`

NewOperatorLimits instantiates a new OperatorLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorLimitsWithDefaults

`func NewOperatorLimitsWithDefaults() *OperatorLimits`

NewOperatorLimitsWithDefaults instantiates a new OperatorLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OperatorLimits) GetData() int64`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OperatorLimits) GetDataOk() (*int64, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OperatorLimits) SetData(v int64)`

SetData sets Data field to given value.

### HasData

`func (o *OperatorLimits) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPayload

`func (o *OperatorLimits) GetPayload() int64`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *OperatorLimits) GetPayloadOk() (*int64, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *OperatorLimits) SetPayload(v int64)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *OperatorLimits) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSubs

`func (o *OperatorLimits) GetSubs() int64`

GetSubs returns the Subs field if non-nil, zero value otherwise.

### GetSubsOk

`func (o *OperatorLimits) GetSubsOk() (*int64, bool)`

GetSubsOk returns a tuple with the Subs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubs

`func (o *OperatorLimits) SetSubs(v int64)`

SetSubs sets Subs field to given value.

### HasSubs

`func (o *OperatorLimits) HasSubs() bool`

HasSubs returns a boolean if a field has been set.

### GetConn

`func (o *OperatorLimits) GetConn() int64`

GetConn returns the Conn field if non-nil, zero value otherwise.

### GetConnOk

`func (o *OperatorLimits) GetConnOk() (*int64, bool)`

GetConnOk returns a tuple with the Conn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConn

`func (o *OperatorLimits) SetConn(v int64)`

SetConn sets Conn field to given value.

### HasConn

`func (o *OperatorLimits) HasConn() bool`

HasConn returns a boolean if a field has been set.

### GetDisallowBearer

`func (o *OperatorLimits) GetDisallowBearer() bool`

GetDisallowBearer returns the DisallowBearer field if non-nil, zero value otherwise.

### GetDisallowBearerOk

`func (o *OperatorLimits) GetDisallowBearerOk() (*bool, bool)`

GetDisallowBearerOk returns a tuple with the DisallowBearer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowBearer

`func (o *OperatorLimits) SetDisallowBearer(v bool)`

SetDisallowBearer sets DisallowBearer field to given value.

### HasDisallowBearer

`func (o *OperatorLimits) HasDisallowBearer() bool`

HasDisallowBearer returns a boolean if a field has been set.

### GetExports

`func (o *OperatorLimits) GetExports() int64`

GetExports returns the Exports field if non-nil, zero value otherwise.

### GetExportsOk

`func (o *OperatorLimits) GetExportsOk() (*int64, bool)`

GetExportsOk returns a tuple with the Exports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExports

`func (o *OperatorLimits) SetExports(v int64)`

SetExports sets Exports field to given value.

### HasExports

`func (o *OperatorLimits) HasExports() bool`

HasExports returns a boolean if a field has been set.

### GetImports

`func (o *OperatorLimits) GetImports() int64`

GetImports returns the Imports field if non-nil, zero value otherwise.

### GetImportsOk

`func (o *OperatorLimits) GetImportsOk() (*int64, bool)`

GetImportsOk returns a tuple with the Imports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImports

`func (o *OperatorLimits) SetImports(v int64)`

SetImports sets Imports field to given value.

### HasImports

`func (o *OperatorLimits) HasImports() bool`

HasImports returns a boolean if a field has been set.

### GetLeaf

`func (o *OperatorLimits) GetLeaf() int64`

GetLeaf returns the Leaf field if non-nil, zero value otherwise.

### GetLeafOk

`func (o *OperatorLimits) GetLeafOk() (*int64, bool)`

GetLeafOk returns a tuple with the Leaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaf

`func (o *OperatorLimits) SetLeaf(v int64)`

SetLeaf sets Leaf field to given value.

### HasLeaf

`func (o *OperatorLimits) HasLeaf() bool`

HasLeaf returns a boolean if a field has been set.

### GetWildcards

`func (o *OperatorLimits) GetWildcards() bool`

GetWildcards returns the Wildcards field if non-nil, zero value otherwise.

### GetWildcardsOk

`func (o *OperatorLimits) GetWildcardsOk() (*bool, bool)`

GetWildcardsOk returns a tuple with the Wildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcards

`func (o *OperatorLimits) SetWildcards(v bool)`

SetWildcards sets Wildcards field to given value.

### HasWildcards

`func (o *OperatorLimits) HasWildcards() bool`

HasWildcards returns a boolean if a field has been set.

### GetConsumer

`func (o *OperatorLimits) GetConsumer() int64`

GetConsumer returns the Consumer field if non-nil, zero value otherwise.

### GetConsumerOk

`func (o *OperatorLimits) GetConsumerOk() (*int64, bool)`

GetConsumerOk returns a tuple with the Consumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumer

`func (o *OperatorLimits) SetConsumer(v int64)`

SetConsumer sets Consumer field to given value.

### HasConsumer

`func (o *OperatorLimits) HasConsumer() bool`

HasConsumer returns a boolean if a field has been set.

### GetDiskMaxStreamBytes

`func (o *OperatorLimits) GetDiskMaxStreamBytes() int64`

GetDiskMaxStreamBytes returns the DiskMaxStreamBytes field if non-nil, zero value otherwise.

### GetDiskMaxStreamBytesOk

`func (o *OperatorLimits) GetDiskMaxStreamBytesOk() (*int64, bool)`

GetDiskMaxStreamBytesOk returns a tuple with the DiskMaxStreamBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskMaxStreamBytes

`func (o *OperatorLimits) SetDiskMaxStreamBytes(v int64)`

SetDiskMaxStreamBytes sets DiskMaxStreamBytes field to given value.

### HasDiskMaxStreamBytes

`func (o *OperatorLimits) HasDiskMaxStreamBytes() bool`

HasDiskMaxStreamBytes returns a boolean if a field has been set.

### GetDiskStorage

`func (o *OperatorLimits) GetDiskStorage() int64`

GetDiskStorage returns the DiskStorage field if non-nil, zero value otherwise.

### GetDiskStorageOk

`func (o *OperatorLimits) GetDiskStorageOk() (*int64, bool)`

GetDiskStorageOk returns a tuple with the DiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStorage

`func (o *OperatorLimits) SetDiskStorage(v int64)`

SetDiskStorage sets DiskStorage field to given value.

### HasDiskStorage

`func (o *OperatorLimits) HasDiskStorage() bool`

HasDiskStorage returns a boolean if a field has been set.

### GetMaxAckPending

`func (o *OperatorLimits) GetMaxAckPending() int64`

GetMaxAckPending returns the MaxAckPending field if non-nil, zero value otherwise.

### GetMaxAckPendingOk

`func (o *OperatorLimits) GetMaxAckPendingOk() (*int64, bool)`

GetMaxAckPendingOk returns a tuple with the MaxAckPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAckPending

`func (o *OperatorLimits) SetMaxAckPending(v int64)`

SetMaxAckPending sets MaxAckPending field to given value.

### HasMaxAckPending

`func (o *OperatorLimits) HasMaxAckPending() bool`

HasMaxAckPending returns a boolean if a field has been set.

### GetMaxBytesRequired

`func (o *OperatorLimits) GetMaxBytesRequired() bool`

GetMaxBytesRequired returns the MaxBytesRequired field if non-nil, zero value otherwise.

### GetMaxBytesRequiredOk

`func (o *OperatorLimits) GetMaxBytesRequiredOk() (*bool, bool)`

GetMaxBytesRequiredOk returns a tuple with the MaxBytesRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytesRequired

`func (o *OperatorLimits) SetMaxBytesRequired(v bool)`

SetMaxBytesRequired sets MaxBytesRequired field to given value.

### HasMaxBytesRequired

`func (o *OperatorLimits) HasMaxBytesRequired() bool`

HasMaxBytesRequired returns a boolean if a field has been set.

### GetMemMaxStreamBytes

`func (o *OperatorLimits) GetMemMaxStreamBytes() int64`

GetMemMaxStreamBytes returns the MemMaxStreamBytes field if non-nil, zero value otherwise.

### GetMemMaxStreamBytesOk

`func (o *OperatorLimits) GetMemMaxStreamBytesOk() (*int64, bool)`

GetMemMaxStreamBytesOk returns a tuple with the MemMaxStreamBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemMaxStreamBytes

`func (o *OperatorLimits) SetMemMaxStreamBytes(v int64)`

SetMemMaxStreamBytes sets MemMaxStreamBytes field to given value.

### HasMemMaxStreamBytes

`func (o *OperatorLimits) HasMemMaxStreamBytes() bool`

HasMemMaxStreamBytes returns a boolean if a field has been set.

### GetMemStorage

`func (o *OperatorLimits) GetMemStorage() int64`

GetMemStorage returns the MemStorage field if non-nil, zero value otherwise.

### GetMemStorageOk

`func (o *OperatorLimits) GetMemStorageOk() (*int64, bool)`

GetMemStorageOk returns a tuple with the MemStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemStorage

`func (o *OperatorLimits) SetMemStorage(v int64)`

SetMemStorage sets MemStorage field to given value.

### HasMemStorage

`func (o *OperatorLimits) HasMemStorage() bool`

HasMemStorage returns a boolean if a field has been set.

### GetStreams

`func (o *OperatorLimits) GetStreams() int64`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *OperatorLimits) GetStreamsOk() (*int64, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *OperatorLimits) SetStreams(v int64)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *OperatorLimits) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetTieredLimits

`func (o *OperatorLimits) GetTieredLimits() map[string]JetStreamLimits`

GetTieredLimits returns the TieredLimits field if non-nil, zero value otherwise.

### GetTieredLimitsOk

`func (o *OperatorLimits) GetTieredLimitsOk() (*map[string]JetStreamLimits, bool)`

GetTieredLimitsOk returns a tuple with the TieredLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTieredLimits

`func (o *OperatorLimits) SetTieredLimits(v map[string]JetStreamLimits)`

SetTieredLimits sets TieredLimits field to given value.

### HasTieredLimits

`func (o *OperatorLimits) HasTieredLimits() bool`

HasTieredLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


