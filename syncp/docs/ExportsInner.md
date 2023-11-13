# ExportsInner

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

### NewExportsInner

`func NewExportsInner() *ExportsInner`

NewExportsInner instantiates a new ExportsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportsInnerWithDefaults

`func NewExportsInnerWithDefaults() *ExportsInner`

NewExportsInnerWithDefaults instantiates a new ExportsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ExportsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExportsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExportsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExportsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInfoUrl

`func (o *ExportsInner) GetInfoUrl() string`

GetInfoUrl returns the InfoUrl field if non-nil, zero value otherwise.

### GetInfoUrlOk

`func (o *ExportsInner) GetInfoUrlOk() (*string, bool)`

GetInfoUrlOk returns a tuple with the InfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoUrl

`func (o *ExportsInner) SetInfoUrl(v string)`

SetInfoUrl sets InfoUrl field to given value.

### HasInfoUrl

`func (o *ExportsInner) HasInfoUrl() bool`

HasInfoUrl returns a boolean if a field has been set.

### GetAccountTokenPosition

`func (o *ExportsInner) GetAccountTokenPosition() int32`

GetAccountTokenPosition returns the AccountTokenPosition field if non-nil, zero value otherwise.

### GetAccountTokenPositionOk

`func (o *ExportsInner) GetAccountTokenPositionOk() (*int32, bool)`

GetAccountTokenPositionOk returns a tuple with the AccountTokenPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTokenPosition

`func (o *ExportsInner) SetAccountTokenPosition(v int32)`

SetAccountTokenPosition sets AccountTokenPosition field to given value.

### HasAccountTokenPosition

`func (o *ExportsInner) HasAccountTokenPosition() bool`

HasAccountTokenPosition returns a boolean if a field has been set.

### GetAdvertise

`func (o *ExportsInner) GetAdvertise() bool`

GetAdvertise returns the Advertise field if non-nil, zero value otherwise.

### GetAdvertiseOk

`func (o *ExportsInner) GetAdvertiseOk() (*bool, bool)`

GetAdvertiseOk returns a tuple with the Advertise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertise

`func (o *ExportsInner) SetAdvertise(v bool)`

SetAdvertise sets Advertise field to given value.

### HasAdvertise

`func (o *ExportsInner) HasAdvertise() bool`

HasAdvertise returns a boolean if a field has been set.

### GetName

`func (o *ExportsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExportsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExportsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExportsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResponseThreshold

`func (o *ExportsInner) GetResponseThreshold() int64`

GetResponseThreshold returns the ResponseThreshold field if non-nil, zero value otherwise.

### GetResponseThresholdOk

`func (o *ExportsInner) GetResponseThresholdOk() (*int64, bool)`

GetResponseThresholdOk returns a tuple with the ResponseThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseThreshold

`func (o *ExportsInner) SetResponseThreshold(v int64)`

SetResponseThreshold sets ResponseThreshold field to given value.

### HasResponseThreshold

`func (o *ExportsInner) HasResponseThreshold() bool`

HasResponseThreshold returns a boolean if a field has been set.

### GetResponseType

`func (o *ExportsInner) GetResponseType() ResponseType`

GetResponseType returns the ResponseType field if non-nil, zero value otherwise.

### GetResponseTypeOk

`func (o *ExportsInner) GetResponseTypeOk() (*ResponseType, bool)`

GetResponseTypeOk returns a tuple with the ResponseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseType

`func (o *ExportsInner) SetResponseType(v ResponseType)`

SetResponseType sets ResponseType field to given value.

### HasResponseType

`func (o *ExportsInner) HasResponseType() bool`

HasResponseType returns a boolean if a field has been set.

### GetRevocations

`func (o *ExportsInner) GetRevocations() map[string]int64`

GetRevocations returns the Revocations field if non-nil, zero value otherwise.

### GetRevocationsOk

`func (o *ExportsInner) GetRevocationsOk() (*map[string]int64, bool)`

GetRevocationsOk returns a tuple with the Revocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocations

`func (o *ExportsInner) SetRevocations(v map[string]int64)`

SetRevocations sets Revocations field to given value.

### HasRevocations

`func (o *ExportsInner) HasRevocations() bool`

HasRevocations returns a boolean if a field has been set.

### GetServiceLatency

`func (o *ExportsInner) GetServiceLatency() ExportAllOfServiceLatency`

GetServiceLatency returns the ServiceLatency field if non-nil, zero value otherwise.

### GetServiceLatencyOk

`func (o *ExportsInner) GetServiceLatencyOk() (*ExportAllOfServiceLatency, bool)`

GetServiceLatencyOk returns a tuple with the ServiceLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLatency

`func (o *ExportsInner) SetServiceLatency(v ExportAllOfServiceLatency)`

SetServiceLatency sets ServiceLatency field to given value.

### HasServiceLatency

`func (o *ExportsInner) HasServiceLatency() bool`

HasServiceLatency returns a boolean if a field has been set.

### SetServiceLatencyNil

`func (o *ExportsInner) SetServiceLatencyNil(b bool)`

 SetServiceLatencyNil sets the value for ServiceLatency to be an explicit nil

### UnsetServiceLatency
`func (o *ExportsInner) UnsetServiceLatency()`

UnsetServiceLatency ensures that no value is present for ServiceLatency, not even an explicit nil
### GetSubject

`func (o *ExportsInner) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ExportsInner) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ExportsInner) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ExportsInner) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTokenReq

`func (o *ExportsInner) GetTokenReq() bool`

GetTokenReq returns the TokenReq field if non-nil, zero value otherwise.

### GetTokenReqOk

`func (o *ExportsInner) GetTokenReqOk() (*bool, bool)`

GetTokenReqOk returns a tuple with the TokenReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenReq

`func (o *ExportsInner) SetTokenReq(v bool)`

SetTokenReq sets TokenReq field to given value.

### HasTokenReq

`func (o *ExportsInner) HasTokenReq() bool`

HasTokenReq returns a boolean if a field has been set.

### GetType

`func (o *ExportsInner) GetType() ExportType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportsInner) GetTypeOk() (*ExportType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportsInner) SetType(v ExportType)`

SetType sets Type field to given value.

### HasType

`func (o *ExportsInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


