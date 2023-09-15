# PeerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **int64** |  | 
**Current** | **bool** |  | 
**Lag** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Offline** | Pointer to **bool** |  | [optional] 

## Methods

### NewPeerInfo

`func NewPeerInfo(active int64, current bool, name string, ) *PeerInfo`

NewPeerInfo instantiates a new PeerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeerInfoWithDefaults

`func NewPeerInfoWithDefaults() *PeerInfo`

NewPeerInfoWithDefaults instantiates a new PeerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *PeerInfo) GetActive() int64`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PeerInfo) GetActiveOk() (*int64, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PeerInfo) SetActive(v int64)`

SetActive sets Active field to given value.


### GetCurrent

`func (o *PeerInfo) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *PeerInfo) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *PeerInfo) SetCurrent(v bool)`

SetCurrent sets Current field to given value.


### GetLag

`func (o *PeerInfo) GetLag() int32`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *PeerInfo) GetLagOk() (*int32, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *PeerInfo) SetLag(v int32)`

SetLag sets Lag field to given value.

### HasLag

`func (o *PeerInfo) HasLag() bool`

HasLag returns a boolean if a field has been set.

### GetName

`func (o *PeerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PeerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PeerInfo) SetName(v string)`

SetName sets Name field to given value.


### GetOffline

`func (o *PeerInfo) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *PeerInfo) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *PeerInfo) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *PeerInfo) HasOffline() bool`

HasOffline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


