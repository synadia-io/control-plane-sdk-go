# NatsServerInfoListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ServerInfo**](ServerInfo.md) |  | 

## Methods

### NewNatsServerInfoListResponse

`func NewNatsServerInfoListResponse(items []ServerInfo, ) *NatsServerInfoListResponse`

NewNatsServerInfoListResponse instantiates a new NatsServerInfoListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatsServerInfoListResponseWithDefaults

`func NewNatsServerInfoListResponseWithDefaults() *NatsServerInfoListResponse`

NewNatsServerInfoListResponseWithDefaults instantiates a new NatsServerInfoListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *NatsServerInfoListResponse) GetItems() []ServerInfo`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NatsServerInfoListResponse) GetItemsOk() (*[]ServerInfo, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NatsServerInfoListResponse) SetItems(v []ServerInfo)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


