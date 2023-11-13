# AcceptTermsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedAt** | **time.Time** |  | 
**IpAddress** | **string** |  | 

## Methods

### NewAcceptTermsResponse

`func NewAcceptTermsResponse(acceptedAt time.Time, ipAddress string, ) *AcceptTermsResponse`

NewAcceptTermsResponse instantiates a new AcceptTermsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptTermsResponseWithDefaults

`func NewAcceptTermsResponseWithDefaults() *AcceptTermsResponse`

NewAcceptTermsResponseWithDefaults instantiates a new AcceptTermsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedAt

`func (o *AcceptTermsResponse) GetAcceptedAt() time.Time`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *AcceptTermsResponse) GetAcceptedAtOk() (*time.Time, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *AcceptTermsResponse) SetAcceptedAt(v time.Time)`

SetAcceptedAt sets AcceptedAt field to given value.


### GetIpAddress

`func (o *AcceptTermsResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *AcceptTermsResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *AcceptTermsResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


