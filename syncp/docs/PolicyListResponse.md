# PolicyListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]AppPolicy**](AppPolicy.md) |  | 

## Methods

### NewPolicyListResponse

`func NewPolicyListResponse(items []AppPolicy, ) *PolicyListResponse`

NewPolicyListResponse instantiates a new PolicyListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyListResponseWithDefaults

`func NewPolicyListResponseWithDefaults() *PolicyListResponse`

NewPolicyListResponseWithDefaults instantiates a new PolicyListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PolicyListResponse) GetItems() []AppPolicy`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PolicyListResponse) GetItemsOk() (*[]AppPolicy, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PolicyListResponse) SetItems(v []AppPolicy)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


