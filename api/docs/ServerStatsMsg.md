# ServerStatsMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | [**ServerInfo**](ServerInfo.md) |  | 
**Statsz** | [**ServerStats**](ServerStats.md) |  | 

## Methods

### NewServerStatsMsg

`func NewServerStatsMsg(server ServerInfo, statsz ServerStats, ) *ServerStatsMsg`

NewServerStatsMsg instantiates a new ServerStatsMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerStatsMsgWithDefaults

`func NewServerStatsMsgWithDefaults() *ServerStatsMsg`

NewServerStatsMsgWithDefaults instantiates a new ServerStatsMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *ServerStatsMsg) GetServer() ServerInfo`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ServerStatsMsg) GetServerOk() (*ServerInfo, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ServerStatsMsg) SetServer(v ServerInfo)`

SetServer sets Server field to given value.


### GetStatsz

`func (o *ServerStatsMsg) GetStatsz() ServerStats`

GetStatsz returns the Statsz field if non-nil, zero value otherwise.

### GetStatszOk

`func (o *ServerStatsMsg) GetStatszOk() (*ServerStats, bool)`

GetStatszOk returns a tuple with the Statsz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsz

`func (o *ServerStatsMsg) SetStatsz(v ServerStats)`

SetStatsz sets Statsz field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


