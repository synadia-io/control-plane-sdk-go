# AccountOverviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountPublicKey** | **string** |  | 
**Metrics** | [**AccountOverviewResponseMetrics**](AccountOverviewResponseMetrics.md) |  | 

## Methods

### NewAccountOverviewResponse

`func NewAccountOverviewResponse(accountPublicKey string, metrics AccountOverviewResponseMetrics, ) *AccountOverviewResponse`

NewAccountOverviewResponse instantiates a new AccountOverviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountOverviewResponseWithDefaults

`func NewAccountOverviewResponseWithDefaults() *AccountOverviewResponse`

NewAccountOverviewResponseWithDefaults instantiates a new AccountOverviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountPublicKey

`func (o *AccountOverviewResponse) GetAccountPublicKey() string`

GetAccountPublicKey returns the AccountPublicKey field if non-nil, zero value otherwise.

### GetAccountPublicKeyOk

`func (o *AccountOverviewResponse) GetAccountPublicKeyOk() (*string, bool)`

GetAccountPublicKeyOk returns a tuple with the AccountPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPublicKey

`func (o *AccountOverviewResponse) SetAccountPublicKey(v string)`

SetAccountPublicKey sets AccountPublicKey field to given value.


### GetMetrics

`func (o *AccountOverviewResponse) GetMetrics() AccountOverviewResponseMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *AccountOverviewResponse) GetMetricsOk() (*AccountOverviewResponseMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *AccountOverviewResponse) SetMetrics(v AccountOverviewResponseMetrics)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


