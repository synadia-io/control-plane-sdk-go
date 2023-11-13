# JSStreamConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDirect** | **bool** |  | 
**AllowRollupHdrs** | **bool** |  | 
**DenyDelete** | **bool** |  | 
**DenyPurge** | **bool** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Discard** | [**DiscardPolicy**](DiscardPolicy.md) |  | 
**DiscardNewPerSubject** | Pointer to **bool** |  | [optional] 
**DuplicateWindow** | Pointer to **int64** |  | [optional] 
**MaxAge** | **int64** |  | 
**MaxBytes** | **int64** |  | 
**MaxConsumers** | **int32** |  | 
**MaxMsgSize** | Pointer to **int32** |  | [optional] 
**MaxMsgs** | **int64** |  | 
**MaxMsgsPerSubject** | **int64** |  | 
**Name** | **string** |  | 
**NoAck** | Pointer to **bool** |  | [optional] 
**NumReplicas** | **int32** |  | 
**Placement** | Pointer to [**NullableJSCommonStreamConfigPlacement**](JSCommonStreamConfigPlacement.md) |  | [optional] 
**Republish** | Pointer to [**NullableJSCommonStreamConfigRepublish**](JSCommonStreamConfigRepublish.md) |  | [optional] 
**Retention** | [**RetentionPolicy**](RetentionPolicy.md) |  | 
**Sealed** | **bool** |  | 
**Sources** | Pointer to [**[]StreamSource**](StreamSource.md) |  | [optional] 
**Storage** | [**StorageType**](StorageType.md) |  | 
**TemplateOwner** | Pointer to **string** |  | [optional] 
**Subjects** | Pointer to **[]string** |  | [optional] 

## Methods

### NewJSStreamConfigRequest

`func NewJSStreamConfigRequest(allowDirect bool, allowRollupHdrs bool, denyDelete bool, denyPurge bool, discard DiscardPolicy, maxAge int64, maxBytes int64, maxConsumers int32, maxMsgs int64, maxMsgsPerSubject int64, name string, numReplicas int32, retention RetentionPolicy, sealed bool, storage StorageType, ) *JSStreamConfigRequest`

NewJSStreamConfigRequest instantiates a new JSStreamConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSStreamConfigRequestWithDefaults

`func NewJSStreamConfigRequestWithDefaults() *JSStreamConfigRequest`

NewJSStreamConfigRequestWithDefaults instantiates a new JSStreamConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowDirect

`func (o *JSStreamConfigRequest) GetAllowDirect() bool`

GetAllowDirect returns the AllowDirect field if non-nil, zero value otherwise.

### GetAllowDirectOk

`func (o *JSStreamConfigRequest) GetAllowDirectOk() (*bool, bool)`

GetAllowDirectOk returns a tuple with the AllowDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDirect

`func (o *JSStreamConfigRequest) SetAllowDirect(v bool)`

SetAllowDirect sets AllowDirect field to given value.


### GetAllowRollupHdrs

`func (o *JSStreamConfigRequest) GetAllowRollupHdrs() bool`

GetAllowRollupHdrs returns the AllowRollupHdrs field if non-nil, zero value otherwise.

### GetAllowRollupHdrsOk

`func (o *JSStreamConfigRequest) GetAllowRollupHdrsOk() (*bool, bool)`

GetAllowRollupHdrsOk returns a tuple with the AllowRollupHdrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRollupHdrs

`func (o *JSStreamConfigRequest) SetAllowRollupHdrs(v bool)`

SetAllowRollupHdrs sets AllowRollupHdrs field to given value.


### GetDenyDelete

`func (o *JSStreamConfigRequest) GetDenyDelete() bool`

GetDenyDelete returns the DenyDelete field if non-nil, zero value otherwise.

### GetDenyDeleteOk

`func (o *JSStreamConfigRequest) GetDenyDeleteOk() (*bool, bool)`

GetDenyDeleteOk returns a tuple with the DenyDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyDelete

`func (o *JSStreamConfigRequest) SetDenyDelete(v bool)`

SetDenyDelete sets DenyDelete field to given value.


### GetDenyPurge

`func (o *JSStreamConfigRequest) GetDenyPurge() bool`

GetDenyPurge returns the DenyPurge field if non-nil, zero value otherwise.

### GetDenyPurgeOk

`func (o *JSStreamConfigRequest) GetDenyPurgeOk() (*bool, bool)`

GetDenyPurgeOk returns a tuple with the DenyPurge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyPurge

`func (o *JSStreamConfigRequest) SetDenyPurge(v bool)`

SetDenyPurge sets DenyPurge field to given value.


### GetDescription

`func (o *JSStreamConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JSStreamConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JSStreamConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JSStreamConfigRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscard

`func (o *JSStreamConfigRequest) GetDiscard() DiscardPolicy`

GetDiscard returns the Discard field if non-nil, zero value otherwise.

### GetDiscardOk

`func (o *JSStreamConfigRequest) GetDiscardOk() (*DiscardPolicy, bool)`

GetDiscardOk returns a tuple with the Discard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscard

`func (o *JSStreamConfigRequest) SetDiscard(v DiscardPolicy)`

SetDiscard sets Discard field to given value.


### GetDiscardNewPerSubject

`func (o *JSStreamConfigRequest) GetDiscardNewPerSubject() bool`

GetDiscardNewPerSubject returns the DiscardNewPerSubject field if non-nil, zero value otherwise.

### GetDiscardNewPerSubjectOk

`func (o *JSStreamConfigRequest) GetDiscardNewPerSubjectOk() (*bool, bool)`

GetDiscardNewPerSubjectOk returns a tuple with the DiscardNewPerSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardNewPerSubject

`func (o *JSStreamConfigRequest) SetDiscardNewPerSubject(v bool)`

SetDiscardNewPerSubject sets DiscardNewPerSubject field to given value.

### HasDiscardNewPerSubject

`func (o *JSStreamConfigRequest) HasDiscardNewPerSubject() bool`

HasDiscardNewPerSubject returns a boolean if a field has been set.

### GetDuplicateWindow

`func (o *JSStreamConfigRequest) GetDuplicateWindow() int64`

GetDuplicateWindow returns the DuplicateWindow field if non-nil, zero value otherwise.

### GetDuplicateWindowOk

`func (o *JSStreamConfigRequest) GetDuplicateWindowOk() (*int64, bool)`

GetDuplicateWindowOk returns a tuple with the DuplicateWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateWindow

`func (o *JSStreamConfigRequest) SetDuplicateWindow(v int64)`

SetDuplicateWindow sets DuplicateWindow field to given value.

### HasDuplicateWindow

`func (o *JSStreamConfigRequest) HasDuplicateWindow() bool`

HasDuplicateWindow returns a boolean if a field has been set.

### GetMaxAge

`func (o *JSStreamConfigRequest) GetMaxAge() int64`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *JSStreamConfigRequest) GetMaxAgeOk() (*int64, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *JSStreamConfigRequest) SetMaxAge(v int64)`

SetMaxAge sets MaxAge field to given value.


### GetMaxBytes

`func (o *JSStreamConfigRequest) GetMaxBytes() int64`

GetMaxBytes returns the MaxBytes field if non-nil, zero value otherwise.

### GetMaxBytesOk

`func (o *JSStreamConfigRequest) GetMaxBytesOk() (*int64, bool)`

GetMaxBytesOk returns a tuple with the MaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytes

`func (o *JSStreamConfigRequest) SetMaxBytes(v int64)`

SetMaxBytes sets MaxBytes field to given value.


### GetMaxConsumers

`func (o *JSStreamConfigRequest) GetMaxConsumers() int32`

GetMaxConsumers returns the MaxConsumers field if non-nil, zero value otherwise.

### GetMaxConsumersOk

`func (o *JSStreamConfigRequest) GetMaxConsumersOk() (*int32, bool)`

GetMaxConsumersOk returns a tuple with the MaxConsumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConsumers

`func (o *JSStreamConfigRequest) SetMaxConsumers(v int32)`

SetMaxConsumers sets MaxConsumers field to given value.


### GetMaxMsgSize

`func (o *JSStreamConfigRequest) GetMaxMsgSize() int32`

GetMaxMsgSize returns the MaxMsgSize field if non-nil, zero value otherwise.

### GetMaxMsgSizeOk

`func (o *JSStreamConfigRequest) GetMaxMsgSizeOk() (*int32, bool)`

GetMaxMsgSizeOk returns a tuple with the MaxMsgSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgSize

`func (o *JSStreamConfigRequest) SetMaxMsgSize(v int32)`

SetMaxMsgSize sets MaxMsgSize field to given value.

### HasMaxMsgSize

`func (o *JSStreamConfigRequest) HasMaxMsgSize() bool`

HasMaxMsgSize returns a boolean if a field has been set.

### GetMaxMsgs

`func (o *JSStreamConfigRequest) GetMaxMsgs() int64`

GetMaxMsgs returns the MaxMsgs field if non-nil, zero value otherwise.

### GetMaxMsgsOk

`func (o *JSStreamConfigRequest) GetMaxMsgsOk() (*int64, bool)`

GetMaxMsgsOk returns a tuple with the MaxMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgs

`func (o *JSStreamConfigRequest) SetMaxMsgs(v int64)`

SetMaxMsgs sets MaxMsgs field to given value.


### GetMaxMsgsPerSubject

`func (o *JSStreamConfigRequest) GetMaxMsgsPerSubject() int64`

GetMaxMsgsPerSubject returns the MaxMsgsPerSubject field if non-nil, zero value otherwise.

### GetMaxMsgsPerSubjectOk

`func (o *JSStreamConfigRequest) GetMaxMsgsPerSubjectOk() (*int64, bool)`

GetMaxMsgsPerSubjectOk returns a tuple with the MaxMsgsPerSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgsPerSubject

`func (o *JSStreamConfigRequest) SetMaxMsgsPerSubject(v int64)`

SetMaxMsgsPerSubject sets MaxMsgsPerSubject field to given value.


### GetName

`func (o *JSStreamConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JSStreamConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JSStreamConfigRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNoAck

`func (o *JSStreamConfigRequest) GetNoAck() bool`

GetNoAck returns the NoAck field if non-nil, zero value otherwise.

### GetNoAckOk

`func (o *JSStreamConfigRequest) GetNoAckOk() (*bool, bool)`

GetNoAckOk returns a tuple with the NoAck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAck

`func (o *JSStreamConfigRequest) SetNoAck(v bool)`

SetNoAck sets NoAck field to given value.

### HasNoAck

`func (o *JSStreamConfigRequest) HasNoAck() bool`

HasNoAck returns a boolean if a field has been set.

### GetNumReplicas

`func (o *JSStreamConfigRequest) GetNumReplicas() int32`

GetNumReplicas returns the NumReplicas field if non-nil, zero value otherwise.

### GetNumReplicasOk

`func (o *JSStreamConfigRequest) GetNumReplicasOk() (*int32, bool)`

GetNumReplicasOk returns a tuple with the NumReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReplicas

`func (o *JSStreamConfigRequest) SetNumReplicas(v int32)`

SetNumReplicas sets NumReplicas field to given value.


### GetPlacement

`func (o *JSStreamConfigRequest) GetPlacement() JSCommonStreamConfigPlacement`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *JSStreamConfigRequest) GetPlacementOk() (*JSCommonStreamConfigPlacement, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *JSStreamConfigRequest) SetPlacement(v JSCommonStreamConfigPlacement)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *JSStreamConfigRequest) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *JSStreamConfigRequest) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *JSStreamConfigRequest) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetRepublish

`func (o *JSStreamConfigRequest) GetRepublish() JSCommonStreamConfigRepublish`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *JSStreamConfigRequest) GetRepublishOk() (*JSCommonStreamConfigRepublish, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *JSStreamConfigRequest) SetRepublish(v JSCommonStreamConfigRepublish)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *JSStreamConfigRequest) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### SetRepublishNil

`func (o *JSStreamConfigRequest) SetRepublishNil(b bool)`

 SetRepublishNil sets the value for Republish to be an explicit nil

### UnsetRepublish
`func (o *JSStreamConfigRequest) UnsetRepublish()`

UnsetRepublish ensures that no value is present for Republish, not even an explicit nil
### GetRetention

`func (o *JSStreamConfigRequest) GetRetention() RetentionPolicy`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *JSStreamConfigRequest) GetRetentionOk() (*RetentionPolicy, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *JSStreamConfigRequest) SetRetention(v RetentionPolicy)`

SetRetention sets Retention field to given value.


### GetSealed

`func (o *JSStreamConfigRequest) GetSealed() bool`

GetSealed returns the Sealed field if non-nil, zero value otherwise.

### GetSealedOk

`func (o *JSStreamConfigRequest) GetSealedOk() (*bool, bool)`

GetSealedOk returns a tuple with the Sealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealed

`func (o *JSStreamConfigRequest) SetSealed(v bool)`

SetSealed sets Sealed field to given value.


### GetSources

`func (o *JSStreamConfigRequest) GetSources() []StreamSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *JSStreamConfigRequest) GetSourcesOk() (*[]StreamSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *JSStreamConfigRequest) SetSources(v []StreamSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *JSStreamConfigRequest) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *JSStreamConfigRequest) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *JSStreamConfigRequest) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil
### GetStorage

`func (o *JSStreamConfigRequest) GetStorage() StorageType`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *JSStreamConfigRequest) GetStorageOk() (*StorageType, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *JSStreamConfigRequest) SetStorage(v StorageType)`

SetStorage sets Storage field to given value.


### GetTemplateOwner

`func (o *JSStreamConfigRequest) GetTemplateOwner() string`

GetTemplateOwner returns the TemplateOwner field if non-nil, zero value otherwise.

### GetTemplateOwnerOk

`func (o *JSStreamConfigRequest) GetTemplateOwnerOk() (*string, bool)`

GetTemplateOwnerOk returns a tuple with the TemplateOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateOwner

`func (o *JSStreamConfigRequest) SetTemplateOwner(v string)`

SetTemplateOwner sets TemplateOwner field to given value.

### HasTemplateOwner

`func (o *JSStreamConfigRequest) HasTemplateOwner() bool`

HasTemplateOwner returns a boolean if a field has been set.

### GetSubjects

`func (o *JSStreamConfigRequest) GetSubjects() []string`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *JSStreamConfigRequest) GetSubjectsOk() (*[]string, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *JSStreamConfigRequest) SetSubjects(v []string)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *JSStreamConfigRequest) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


