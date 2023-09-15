# JSKVBucketCreateRequestMirror

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**External** | Pointer to [**NullableStreamSourceExternal**](StreamSourceExternal.md) |  | [optional] 
**FilterSubject** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**OptStartSeq** | Pointer to **int32** |  | [optional] 
**OptStartTime** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewJSKVBucketCreateRequestMirror

`func NewJSKVBucketCreateRequestMirror(name string, ) *JSKVBucketCreateRequestMirror`

NewJSKVBucketCreateRequestMirror instantiates a new JSKVBucketCreateRequestMirror object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSKVBucketCreateRequestMirrorWithDefaults

`func NewJSKVBucketCreateRequestMirrorWithDefaults() *JSKVBucketCreateRequestMirror`

NewJSKVBucketCreateRequestMirrorWithDefaults instantiates a new JSKVBucketCreateRequestMirror object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternal

`func (o *JSKVBucketCreateRequestMirror) GetExternal() StreamSourceExternal`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *JSKVBucketCreateRequestMirror) GetExternalOk() (*StreamSourceExternal, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *JSKVBucketCreateRequestMirror) SetExternal(v StreamSourceExternal)`

SetExternal sets External field to given value.

### HasExternal

`func (o *JSKVBucketCreateRequestMirror) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### SetExternalNil

`func (o *JSKVBucketCreateRequestMirror) SetExternalNil(b bool)`

 SetExternalNil sets the value for External to be an explicit nil

### UnsetExternal
`func (o *JSKVBucketCreateRequestMirror) UnsetExternal()`

UnsetExternal ensures that no value is present for External, not even an explicit nil
### GetFilterSubject

`func (o *JSKVBucketCreateRequestMirror) GetFilterSubject() string`

GetFilterSubject returns the FilterSubject field if non-nil, zero value otherwise.

### GetFilterSubjectOk

`func (o *JSKVBucketCreateRequestMirror) GetFilterSubjectOk() (*string, bool)`

GetFilterSubjectOk returns a tuple with the FilterSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSubject

`func (o *JSKVBucketCreateRequestMirror) SetFilterSubject(v string)`

SetFilterSubject sets FilterSubject field to given value.

### HasFilterSubject

`func (o *JSKVBucketCreateRequestMirror) HasFilterSubject() bool`

HasFilterSubject returns a boolean if a field has been set.

### GetName

`func (o *JSKVBucketCreateRequestMirror) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JSKVBucketCreateRequestMirror) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JSKVBucketCreateRequestMirror) SetName(v string)`

SetName sets Name field to given value.


### GetOptStartSeq

`func (o *JSKVBucketCreateRequestMirror) GetOptStartSeq() int32`

GetOptStartSeq returns the OptStartSeq field if non-nil, zero value otherwise.

### GetOptStartSeqOk

`func (o *JSKVBucketCreateRequestMirror) GetOptStartSeqOk() (*int32, bool)`

GetOptStartSeqOk returns a tuple with the OptStartSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptStartSeq

`func (o *JSKVBucketCreateRequestMirror) SetOptStartSeq(v int32)`

SetOptStartSeq sets OptStartSeq field to given value.

### HasOptStartSeq

`func (o *JSKVBucketCreateRequestMirror) HasOptStartSeq() bool`

HasOptStartSeq returns a boolean if a field has been set.

### GetOptStartTime

`func (o *JSKVBucketCreateRequestMirror) GetOptStartTime() string`

GetOptStartTime returns the OptStartTime field if non-nil, zero value otherwise.

### GetOptStartTimeOk

`func (o *JSKVBucketCreateRequestMirror) GetOptStartTimeOk() (*string, bool)`

GetOptStartTimeOk returns a tuple with the OptStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptStartTime

`func (o *JSKVBucketCreateRequestMirror) SetOptStartTime(v string)`

SetOptStartTime sets OptStartTime field to given value.

### HasOptStartTime

`func (o *JSKVBucketCreateRequestMirror) HasOptStartTime() bool`

HasOptStartTime returns a boolean if a field has been set.

### SetOptStartTimeNil

`func (o *JSKVBucketCreateRequestMirror) SetOptStartTimeNil(b bool)`

 SetOptStartTimeNil sets the value for OptStartTime to be an explicit nil

### UnsetOptStartTime
`func (o *JSKVBucketCreateRequestMirror) UnsetOptStartTime()`

UnsetOptStartTime ensures that no value is present for OptStartTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


