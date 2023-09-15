/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
)

// checks if the JSCommonStreamConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSCommonStreamConfig{}

// JSCommonStreamConfig struct for JSCommonStreamConfig
type JSCommonStreamConfig struct {
	AllowDirect          bool                                  `json:"allow_direct"`
	AllowRollupHdrs      bool                                  `json:"allow_rollup_hdrs"`
	DenyDelete           bool                                  `json:"deny_delete"`
	DenyPurge            bool                                  `json:"deny_purge"`
	Description          *string                               `json:"description,omitempty"`
	Discard              DiscardPolicy                         `json:"discard"`
	DiscardNewPerSubject *bool                                 `json:"discard_new_per_subject,omitempty"`
	DuplicateWindow      *int64                                `json:"duplicate_window,omitempty"`
	MaxAge               int64                                 `json:"max_age"`
	MaxBytes             int64                                 `json:"max_bytes"`
	MaxConsumers         int32                                 `json:"max_consumers"`
	MaxMsgSize           *int32                                `json:"max_msg_size,omitempty"`
	MaxMsgs              int64                                 `json:"max_msgs"`
	MaxMsgsPerSubject    int64                                 `json:"max_msgs_per_subject"`
	Name                 string                                `json:"name"`
	NoAck                *bool                                 `json:"no_ack,omitempty"`
	NumReplicas          int32                                 `json:"num_replicas"`
	Placement            NullableJSCommonStreamConfigPlacement `json:"placement,omitempty"`
	Republish            NullableJSCommonStreamConfigRepublish `json:"republish,omitempty"`
	Retention            RetentionPolicy                       `json:"retention"`
	Sealed               bool                                  `json:"sealed"`
	Sources              []StreamSource                        `json:"sources,omitempty"`
	Storage              StorageType                           `json:"storage"`
	TemplateOwner        *string                               `json:"template_owner,omitempty"`
}

// NewJSCommonStreamConfig instantiates a new JSCommonStreamConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSCommonStreamConfig(allowDirect bool, allowRollupHdrs bool, denyDelete bool, denyPurge bool, discard DiscardPolicy, maxAge int64, maxBytes int64, maxConsumers int32, maxMsgs int64, maxMsgsPerSubject int64, name string, numReplicas int32, retention RetentionPolicy, sealed bool, storage StorageType) *JSCommonStreamConfig {
	this := JSCommonStreamConfig{}
	this.AllowDirect = allowDirect
	this.AllowRollupHdrs = allowRollupHdrs
	this.DenyDelete = denyDelete
	this.DenyPurge = denyPurge
	this.Discard = discard
	this.MaxAge = maxAge
	this.MaxBytes = maxBytes
	this.MaxConsumers = maxConsumers
	this.MaxMsgs = maxMsgs
	this.MaxMsgsPerSubject = maxMsgsPerSubject
	this.Name = name
	this.NumReplicas = numReplicas
	this.Retention = retention
	this.Sealed = sealed
	this.Storage = storage
	return &this
}

// NewJSCommonStreamConfigWithDefaults instantiates a new JSCommonStreamConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSCommonStreamConfigWithDefaults() *JSCommonStreamConfig {
	this := JSCommonStreamConfig{}
	return &this
}

// GetAllowDirect returns the AllowDirect field value
func (o *JSCommonStreamConfig) GetAllowDirect() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowDirect
}

// GetAllowDirectOk returns a tuple with the AllowDirect field value
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetAllowDirectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowDirect, true
}

// SetAllowDirect sets field value
func (o *JSCommonStreamConfig) SetAllowDirect(v bool) {
	o.AllowDirect = v
}

// GetAllowRollupHdrs returns the AllowRollupHdrs field value
func (o *JSCommonStreamConfig) GetAllowRollupHdrs() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowRollupHdrs
}

// GetAllowRollupHdrsOk returns a tuple with the AllowRollupHdrs field value
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetAllowRollupHdrsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowRollupHdrs, true
}

// SetAllowRollupHdrs sets field value
func (o *JSCommonStreamConfig) SetAllowRollupHdrs(v bool) {
	o.AllowRollupHdrs = v
}

// GetDenyDelete returns the DenyDelete field value
func (o *JSCommonStreamConfig) GetDenyDelete() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DenyDelete
}

// GetDenyDeleteOk returns a tuple with the DenyDelete field value
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetDenyDeleteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DenyDelete, true
}

// SetDenyDelete sets field value
func (o *JSCommonStreamConfig) SetDenyDelete(v bool) {
	o.DenyDelete = v
}

// GetDenyPurge returns the DenyPurge field value
func (o *JSCommonStreamConfig) GetDenyPurge() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DenyPurge
}

// GetDenyPurgeOk returns a tuple with the DenyPurge field value
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetDenyPurgeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DenyPurge, true
}

// SetDenyPurge sets field value
func (o *JSCommonStreamConfig) SetDenyPurge(v bool) {
	o.DenyPurge = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JSCommonStreamConfig) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JSCommonStreamConfig) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JSCommonStreamConfig) SetDescription(v string) {
	o.Description = &v
}

// GetDiscard returns the Discard field value
func (o *JSCommonStreamConfig) GetDiscard() DiscardPolicy {
	if o == nil {
		var ret DiscardPolicy
		return ret
	}

	return o.Discard
}

// GetDiscardOk returns a tuple with the Discard field value
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetDiscardOk() (*DiscardPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Discard, true
}

// SetDiscard sets field value
func (o *JSCommonStreamConfig) SetDiscard(v DiscardPolicy) {
	o.Discard = v
}

// GetDiscardNewPerSubject returns the DiscardNewPerSubject field value if set, zero value otherwise.
func (o *JSCommonStreamConfig) GetDiscardNewPerSubject() bool {
	if o == nil || IsNil(o.DiscardNewPerSubject) {
		var ret bool
		return ret
	}
	return *o.DiscardNewPerSubject
}

// GetDiscardNewPerSubjectOk returns a tuple with the DiscardNewPerSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetDiscardNewPerSubjectOk() (*bool, bool) {
	if o == nil || IsNil(o.DiscardNewPerSubject) {
		return nil, false
	}
	return o.DiscardNewPerSubject, true
}

// HasDiscardNewPerSubject returns a boolean if a field has been set.
func (o *JSCommonStreamConfig) HasDiscardNewPerSubject() bool {
	if o != nil && !IsNil(o.DiscardNewPerSubject) {
		return true
	}

	return false
}

// SetDiscardNewPerSubject gets a reference to the given bool and assigns it to the DiscardNewPerSubject field.
func (o *JSCommonStreamConfig) SetDiscardNewPerSubject(v bool) {
	o.DiscardNewPerSubject = &v
}

// GetDuplicateWindow returns the DuplicateWindow field value if set, zero value otherwise.
func (o *JSCommonStreamConfig) GetDuplicateWindow() int64 {
	if o == nil || IsNil(o.DuplicateWindow) {
		var ret int64
		return ret
	}
	return *o.DuplicateWindow
}

// GetDuplicateWindowOk returns a tuple with the DuplicateWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetDuplicateWindowOk() (*int64, bool) {
	if o == nil || IsNil(o.DuplicateWindow) {
		return nil, false
	}
	return o.DuplicateWindow, true
}

// HasDuplicateWindow returns a boolean if a field has been set.
func (o *JSCommonStreamConfig) HasDuplicateWindow() bool {
	if o != nil && !IsNil(o.DuplicateWindow) {
		return true
	}

	return false
}

// SetDuplicateWindow gets a reference to the given int64 and assigns it to the DuplicateWindow field.
func (o *JSCommonStreamConfig) SetDuplicateWindow(v int64) {
	o.DuplicateWindow = &v
}

// GetMaxAge returns the MaxAge field value
func (o *JSCommonStreamConfig) GetMaxAge() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxAge
}

// GetMaxAgeOk returns a tuple with the MaxAge field value
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetMaxAgeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxAge, true
}

// SetMaxAge sets field value
func (o *JSCommonStreamConfig) SetMaxAge(v int64) {
	o.MaxAge = v
}

// GetMaxBytes returns the MaxBytes field value
func (o *JSCommonStreamConfig) GetMaxBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxBytes
}

// GetMaxBytesOk returns a tuple with the MaxBytes field value
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetMaxBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxBytes, true
}

// SetMaxBytes sets field value
func (o *JSCommonStreamConfig) SetMaxBytes(v int64) {
	o.MaxBytes = v
}

// GetMaxConsumers returns the MaxConsumers field value
func (o *JSCommonStreamConfig) GetMaxConsumers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxConsumers
}

// GetMaxConsumersOk returns a tuple with the MaxConsumers field value
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetMaxConsumersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxConsumers, true
}

// SetMaxConsumers sets field value
func (o *JSCommonStreamConfig) SetMaxConsumers(v int32) {
	o.MaxConsumers = v
}

// GetMaxMsgSize returns the MaxMsgSize field value if set, zero value otherwise.
func (o *JSCommonStreamConfig) GetMaxMsgSize() int32 {
	if o == nil || IsNil(o.MaxMsgSize) {
		var ret int32
		return ret
	}
	return *o.MaxMsgSize
}

// GetMaxMsgSizeOk returns a tuple with the MaxMsgSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetMaxMsgSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxMsgSize) {
		return nil, false
	}
	return o.MaxMsgSize, true
}

// HasMaxMsgSize returns a boolean if a field has been set.
func (o *JSCommonStreamConfig) HasMaxMsgSize() bool {
	if o != nil && !IsNil(o.MaxMsgSize) {
		return true
	}

	return false
}

// SetMaxMsgSize gets a reference to the given int32 and assigns it to the MaxMsgSize field.
func (o *JSCommonStreamConfig) SetMaxMsgSize(v int32) {
	o.MaxMsgSize = &v
}

// GetMaxMsgs returns the MaxMsgs field value
func (o *JSCommonStreamConfig) GetMaxMsgs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxMsgs
}

// GetMaxMsgsOk returns a tuple with the MaxMsgs field value
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetMaxMsgsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxMsgs, true
}

// SetMaxMsgs sets field value
func (o *JSCommonStreamConfig) SetMaxMsgs(v int64) {
	o.MaxMsgs = v
}

// GetMaxMsgsPerSubject returns the MaxMsgsPerSubject field value
func (o *JSCommonStreamConfig) GetMaxMsgsPerSubject() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxMsgsPerSubject
}

// GetMaxMsgsPerSubjectOk returns a tuple with the MaxMsgsPerSubject field value
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetMaxMsgsPerSubjectOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxMsgsPerSubject, true
}

// SetMaxMsgsPerSubject sets field value
func (o *JSCommonStreamConfig) SetMaxMsgsPerSubject(v int64) {
	o.MaxMsgsPerSubject = v
}

// GetName returns the Name field value
func (o *JSCommonStreamConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JSCommonStreamConfig) SetName(v string) {
	o.Name = v
}

// GetNoAck returns the NoAck field value if set, zero value otherwise.
func (o *JSCommonStreamConfig) GetNoAck() bool {
	if o == nil || IsNil(o.NoAck) {
		var ret bool
		return ret
	}
	return *o.NoAck
}

// GetNoAckOk returns a tuple with the NoAck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetNoAckOk() (*bool, bool) {
	if o == nil || IsNil(o.NoAck) {
		return nil, false
	}
	return o.NoAck, true
}

// HasNoAck returns a boolean if a field has been set.
func (o *JSCommonStreamConfig) HasNoAck() bool {
	if o != nil && !IsNil(o.NoAck) {
		return true
	}

	return false
}

// SetNoAck gets a reference to the given bool and assigns it to the NoAck field.
func (o *JSCommonStreamConfig) SetNoAck(v bool) {
	o.NoAck = &v
}

// GetNumReplicas returns the NumReplicas field value
func (o *JSCommonStreamConfig) GetNumReplicas() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumReplicas
}

// GetNumReplicasOk returns a tuple with the NumReplicas field value
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetNumReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumReplicas, true
}

// SetNumReplicas sets field value
func (o *JSCommonStreamConfig) SetNumReplicas(v int32) {
	o.NumReplicas = v
}

// GetPlacement returns the Placement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JSCommonStreamConfig) GetPlacement() JSCommonStreamConfigPlacement {
	if o == nil || IsNil(o.Placement.Get()) {
		var ret JSCommonStreamConfigPlacement
		return ret
	}
	return *o.Placement.Get()
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JSCommonStreamConfig) GetPlacementOk() (*JSCommonStreamConfigPlacement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Placement.Get(), o.Placement.IsSet()
}

// HasPlacement returns a boolean if a field has been set.
func (o *JSCommonStreamConfig) HasPlacement() bool {
	if o != nil && o.Placement.IsSet() {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given NullableJSCommonStreamConfigPlacement and assigns it to the Placement field.
func (o *JSCommonStreamConfig) SetPlacement(v JSCommonStreamConfigPlacement) {
	o.Placement.Set(&v)
}

// SetPlacementNil sets the value for Placement to be an explicit nil
func (o *JSCommonStreamConfig) SetPlacementNil() {
	o.Placement.Set(nil)
}

// UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
func (o *JSCommonStreamConfig) UnsetPlacement() {
	o.Placement.Unset()
}

// GetRepublish returns the Republish field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JSCommonStreamConfig) GetRepublish() JSCommonStreamConfigRepublish {
	if o == nil || IsNil(o.Republish.Get()) {
		var ret JSCommonStreamConfigRepublish
		return ret
	}
	return *o.Republish.Get()
}

// GetRepublishOk returns a tuple with the Republish field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JSCommonStreamConfig) GetRepublishOk() (*JSCommonStreamConfigRepublish, bool) {
	if o == nil {
		return nil, false
	}
	return o.Republish.Get(), o.Republish.IsSet()
}

// HasRepublish returns a boolean if a field has been set.
func (o *JSCommonStreamConfig) HasRepublish() bool {
	if o != nil && o.Republish.IsSet() {
		return true
	}

	return false
}

// SetRepublish gets a reference to the given NullableJSCommonStreamConfigRepublish and assigns it to the Republish field.
func (o *JSCommonStreamConfig) SetRepublish(v JSCommonStreamConfigRepublish) {
	o.Republish.Set(&v)
}

// SetRepublishNil sets the value for Republish to be an explicit nil
func (o *JSCommonStreamConfig) SetRepublishNil() {
	o.Republish.Set(nil)
}

// UnsetRepublish ensures that no value is present for Republish, not even an explicit nil
func (o *JSCommonStreamConfig) UnsetRepublish() {
	o.Republish.Unset()
}

// GetRetention returns the Retention field value
func (o *JSCommonStreamConfig) GetRetention() RetentionPolicy {
	if o == nil {
		var ret RetentionPolicy
		return ret
	}

	return o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetRetentionOk() (*RetentionPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retention, true
}

// SetRetention sets field value
func (o *JSCommonStreamConfig) SetRetention(v RetentionPolicy) {
	o.Retention = v
}

// GetSealed returns the Sealed field value
func (o *JSCommonStreamConfig) GetSealed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sealed
}

// GetSealedOk returns a tuple with the Sealed field value
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetSealedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sealed, true
}

// SetSealed sets field value
func (o *JSCommonStreamConfig) SetSealed(v bool) {
	o.Sealed = v
}

// GetSources returns the Sources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JSCommonStreamConfig) GetSources() []StreamSource {
	if o == nil {
		var ret []StreamSource
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JSCommonStreamConfig) GetSourcesOk() ([]StreamSource, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *JSCommonStreamConfig) HasSources() bool {
	if o != nil && IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []StreamSource and assigns it to the Sources field.
func (o *JSCommonStreamConfig) SetSources(v []StreamSource) {
	o.Sources = v
}

// GetStorage returns the Storage field value
func (o *JSCommonStreamConfig) GetStorage() StorageType {
	if o == nil {
		var ret StorageType
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetStorageOk() (*StorageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *JSCommonStreamConfig) SetStorage(v StorageType) {
	o.Storage = v
}

// GetTemplateOwner returns the TemplateOwner field value if set, zero value otherwise.
func (o *JSCommonStreamConfig) GetTemplateOwner() string {
	if o == nil || IsNil(o.TemplateOwner) {
		var ret string
		return ret
	}
	return *o.TemplateOwner
}

// GetTemplateOwnerOk returns a tuple with the TemplateOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSCommonStreamConfig) GetTemplateOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateOwner) {
		return nil, false
	}
	return o.TemplateOwner, true
}

// HasTemplateOwner returns a boolean if a field has been set.
func (o *JSCommonStreamConfig) HasTemplateOwner() bool {
	if o != nil && !IsNil(o.TemplateOwner) {
		return true
	}

	return false
}

// SetTemplateOwner gets a reference to the given string and assigns it to the TemplateOwner field.
func (o *JSCommonStreamConfig) SetTemplateOwner(v string) {
	o.TemplateOwner = &v
}

func (o JSCommonStreamConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JSCommonStreamConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allow_direct"] = o.AllowDirect
	toSerialize["allow_rollup_hdrs"] = o.AllowRollupHdrs
	toSerialize["deny_delete"] = o.DenyDelete
	toSerialize["deny_purge"] = o.DenyPurge
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["discard"] = o.Discard
	if !IsNil(o.DiscardNewPerSubject) {
		toSerialize["discard_new_per_subject"] = o.DiscardNewPerSubject
	}
	if !IsNil(o.DuplicateWindow) {
		toSerialize["duplicate_window"] = o.DuplicateWindow
	}
	toSerialize["max_age"] = o.MaxAge
	toSerialize["max_bytes"] = o.MaxBytes
	toSerialize["max_consumers"] = o.MaxConsumers
	if !IsNil(o.MaxMsgSize) {
		toSerialize["max_msg_size"] = o.MaxMsgSize
	}
	toSerialize["max_msgs"] = o.MaxMsgs
	toSerialize["max_msgs_per_subject"] = o.MaxMsgsPerSubject
	toSerialize["name"] = o.Name
	if !IsNil(o.NoAck) {
		toSerialize["no_ack"] = o.NoAck
	}
	toSerialize["num_replicas"] = o.NumReplicas
	if o.Placement.IsSet() {
		toSerialize["placement"] = o.Placement.Get()
	}
	if o.Republish.IsSet() {
		toSerialize["republish"] = o.Republish.Get()
	}
	toSerialize["retention"] = o.Retention
	toSerialize["sealed"] = o.Sealed
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	toSerialize["storage"] = o.Storage
	if !IsNil(o.TemplateOwner) {
		toSerialize["template_owner"] = o.TemplateOwner
	}
	return toSerialize, nil
}

type NullableJSCommonStreamConfig struct {
	value *JSCommonStreamConfig
	isSet bool
}

func (v NullableJSCommonStreamConfig) Get() *JSCommonStreamConfig {
	return v.value
}

func (v *NullableJSCommonStreamConfig) Set(val *JSCommonStreamConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableJSCommonStreamConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableJSCommonStreamConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSCommonStreamConfig(val *JSCommonStreamConfig) *NullableJSCommonStreamConfig {
	return &NullableJSCommonStreamConfig{value: val, isSet: true}
}

func (v NullableJSCommonStreamConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSCommonStreamConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
