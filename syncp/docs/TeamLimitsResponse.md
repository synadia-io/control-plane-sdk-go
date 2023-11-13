# TeamLimitsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allocated** | [**TeamLimits**](TeamLimits.md) |  | 
**Available** | [**TeamLimits**](TeamLimits.md) |  | 
**Total** | [**TeamLimits**](TeamLimits.md) |  | 

## Methods

### NewTeamLimitsResponse

`func NewTeamLimitsResponse(allocated TeamLimits, available TeamLimits, total TeamLimits, ) *TeamLimitsResponse`

NewTeamLimitsResponse instantiates a new TeamLimitsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamLimitsResponseWithDefaults

`func NewTeamLimitsResponseWithDefaults() *TeamLimitsResponse`

NewTeamLimitsResponseWithDefaults instantiates a new TeamLimitsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocated

`func (o *TeamLimitsResponse) GetAllocated() TeamLimits`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *TeamLimitsResponse) GetAllocatedOk() (*TeamLimits, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *TeamLimitsResponse) SetAllocated(v TeamLimits)`

SetAllocated sets Allocated field to given value.


### GetAvailable

`func (o *TeamLimitsResponse) GetAvailable() TeamLimits`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *TeamLimitsResponse) GetAvailableOk() (*TeamLimits, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *TeamLimitsResponse) SetAvailable(v TeamLimits)`

SetAvailable sets Available field to given value.


### GetTotal

`func (o *TeamLimitsResponse) GetTotal() TeamLimits`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TeamLimitsResponse) GetTotalOk() (*TeamLimits, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TeamLimitsResponse) SetTotal(v TeamLimits)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


