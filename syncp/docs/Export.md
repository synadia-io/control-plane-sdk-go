# Export

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**InfoUrl** | Pointer to **string** |  | [optional] 
**AccountTokenPosition** | Pointer to **int32** |  | [optional] 
**Advertise** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ResponseThreshold** | Pointer to **int64** |  | [optional] 
**ResponseType** | Pointer to [**ResponseType**](ResponseType.md) |  | [optional] 
**Revocations** | Pointer to **map[string]int64** | RevocationList is used to store a mapping of public keys to unix timestamps | [optional] 
**ServiceLatency** | Pointer to [**NullableExportAllOfServiceLatency**](ExportAllOfServiceLatency.md) |  | [optional] 
**Subject** | Pointer to **string** | Subject is a string that represents a NATS subject | [optional] 
**TokenReq** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to [**ExportType**](ExportType.md) |  | [optional] 

## Methods

### NewExport

`func NewExport() *Export`

NewExport instantiates a new Export object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportWithDefaults

`func NewExportWithDefaults() *Export`

NewExportWithDefaults instantiates a new Export object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Export) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Export) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Export) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Export) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInfoUrl

`func (o *Export) GetInfoUrl() string`

GetInfoUrl returns the InfoUrl field if non-nil, zero value otherwise.

### GetInfoUrlOk

`func (o *Export) GetInfoUrlOk() (*string, bool)`

GetInfoUrlOk returns a tuple with the InfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoUrl

`func (o *Export) SetInfoUrl(v string)`

SetInfoUrl sets InfoUrl field to given value.

### HasInfoUrl

`func (o *Export) HasInfoUrl() bool`

HasInfoUrl returns a boolean if a field has been set.

### GetAccountTokenPosition

`func (o *Export) GetAccountTokenPosition() int32`

GetAccountTokenPosition returns the AccountTokenPosition field if non-nil, zero value otherwise.

### GetAccountTokenPositionOk

`func (o *Export) GetAccountTokenPositionOk() (*int32, bool)`

GetAccountTokenPositionOk returns a tuple with the AccountTokenPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTokenPosition

`func (o *Export) SetAccountTokenPosition(v int32)`

SetAccountTokenPosition sets AccountTokenPosition field to given value.

### HasAccountTokenPosition

`func (o *Export) HasAccountTokenPosition() bool`

HasAccountTokenPosition returns a boolean if a field has been set.

### GetAdvertise

`func (o *Export) GetAdvertise() bool`

GetAdvertise returns the Advertise field if non-nil, zero value otherwise.

### GetAdvertiseOk

`func (o *Export) GetAdvertiseOk() (*bool, bool)`

GetAdvertiseOk returns a tuple with the Advertise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertise

`func (o *Export) SetAdvertise(v bool)`

SetAdvertise sets Advertise field to given value.

### HasAdvertise

`func (o *Export) HasAdvertise() bool`

HasAdvertise returns a boolean if a field has been set.

### GetName

`func (o *Export) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Export) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Export) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Export) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResponseThreshold

`func (o *Export) GetResponseThreshold() int64`

GetResponseThreshold returns the ResponseThreshold field if non-nil, zero value otherwise.

### GetResponseThresholdOk

`func (o *Export) GetResponseThresholdOk() (*int64, bool)`

GetResponseThresholdOk returns a tuple with the ResponseThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseThreshold

`func (o *Export) SetResponseThreshold(v int64)`

SetResponseThreshold sets ResponseThreshold field to given value.

### HasResponseThreshold

`func (o *Export) HasResponseThreshold() bool`

HasResponseThreshold returns a boolean if a field has been set.

### GetResponseType

`func (o *Export) GetResponseType() ResponseType`

GetResponseType returns the ResponseType field if non-nil, zero value otherwise.

### GetResponseTypeOk

`func (o *Export) GetResponseTypeOk() (*ResponseType, bool)`

GetResponseTypeOk returns a tuple with the ResponseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseType

`func (o *Export) SetResponseType(v ResponseType)`

SetResponseType sets ResponseType field to given value.

### HasResponseType

`func (o *Export) HasResponseType() bool`

HasResponseType returns a boolean if a field has been set.

### GetRevocations

`func (o *Export) GetRevocations() map[string]int64`

GetRevocations returns the Revocations field if non-nil, zero value otherwise.

### GetRevocationsOk

`func (o *Export) GetRevocationsOk() (*map[string]int64, bool)`

GetRevocationsOk returns a tuple with the Revocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocations

`func (o *Export) SetRevocations(v map[string]int64)`

SetRevocations sets Revocations field to given value.

### HasRevocations

`func (o *Export) HasRevocations() bool`

HasRevocations returns a boolean if a field has been set.

### GetServiceLatency

`func (o *Export) GetServiceLatency() ExportAllOfServiceLatency`

GetServiceLatency returns the ServiceLatency field if non-nil, zero value otherwise.

### GetServiceLatencyOk

`func (o *Export) GetServiceLatencyOk() (*ExportAllOfServiceLatency, bool)`

GetServiceLatencyOk returns a tuple with the ServiceLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLatency

`func (o *Export) SetServiceLatency(v ExportAllOfServiceLatency)`

SetServiceLatency sets ServiceLatency field to given value.

### HasServiceLatency

`func (o *Export) HasServiceLatency() bool`

HasServiceLatency returns a boolean if a field has been set.

### SetServiceLatencyNil

`func (o *Export) SetServiceLatencyNil(b bool)`

 SetServiceLatencyNil sets the value for ServiceLatency to be an explicit nil

### UnsetServiceLatency
`func (o *Export) UnsetServiceLatency()`

UnsetServiceLatency ensures that no value is present for ServiceLatency, not even an explicit nil
### GetSubject

`func (o *Export) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Export) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Export) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Export) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTokenReq

`func (o *Export) GetTokenReq() bool`

GetTokenReq returns the TokenReq field if non-nil, zero value otherwise.

### GetTokenReqOk

`func (o *Export) GetTokenReqOk() (*bool, bool)`

GetTokenReqOk returns a tuple with the TokenReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenReq

`func (o *Export) SetTokenReq(v bool)`

SetTokenReq sets TokenReq field to given value.

### HasTokenReq

`func (o *Export) HasTokenReq() bool`

HasTokenReq returns a boolean if a field has been set.

### GetType

`func (o *Export) GetType() ExportType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Export) GetTypeOk() (*ExportType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Export) SetType(v ExportType)`

SetType sets Type field to given value.

### HasType

`func (o *Export) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


