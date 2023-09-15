/*
Synadia Control Plane

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
)

// checks if the MetaClusterInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetaClusterInfo{}

// MetaClusterInfo MetaClusterInfo shows information about the meta group.
type MetaClusterInfo struct {
	ClusterSize int32                      `json:"cluster_size"`
	Leader      *string                    `json:"leader,omitempty"`
	Name        *string                    `json:"name,omitempty"`
	Replicas    []ClusterInfoReplicasInner `json:"replicas,omitempty"`
}

// NewMetaClusterInfo instantiates a new MetaClusterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaClusterInfo(clusterSize int32) *MetaClusterInfo {
	this := MetaClusterInfo{}
	this.ClusterSize = clusterSize
	return &this
}

// NewMetaClusterInfoWithDefaults instantiates a new MetaClusterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaClusterInfoWithDefaults() *MetaClusterInfo {
	this := MetaClusterInfo{}
	return &this
}

// GetClusterSize returns the ClusterSize field value
func (o *MetaClusterInfo) GetClusterSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClusterSize
}

// GetClusterSizeOk returns a tuple with the ClusterSize field value
// and a boolean to check if the value has been set.
func (o *MetaClusterInfo) GetClusterSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterSize, true
}

// SetClusterSize sets field value
func (o *MetaClusterInfo) SetClusterSize(v int32) {
	o.ClusterSize = v
}

// GetLeader returns the Leader field value if set, zero value otherwise.
func (o *MetaClusterInfo) GetLeader() string {
	if o == nil || IsNil(o.Leader) {
		var ret string
		return ret
	}
	return *o.Leader
}

// GetLeaderOk returns a tuple with the Leader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaClusterInfo) GetLeaderOk() (*string, bool) {
	if o == nil || IsNil(o.Leader) {
		return nil, false
	}
	return o.Leader, true
}

// HasLeader returns a boolean if a field has been set.
func (o *MetaClusterInfo) HasLeader() bool {
	if o != nil && !IsNil(o.Leader) {
		return true
	}

	return false
}

// SetLeader gets a reference to the given string and assigns it to the Leader field.
func (o *MetaClusterInfo) SetLeader(v string) {
	o.Leader = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MetaClusterInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaClusterInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MetaClusterInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MetaClusterInfo) SetName(v string) {
	o.Name = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *MetaClusterInfo) GetReplicas() []ClusterInfoReplicasInner {
	if o == nil || IsNil(o.Replicas) {
		var ret []ClusterInfoReplicasInner
		return ret
	}
	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaClusterInfo) GetReplicasOk() ([]ClusterInfoReplicasInner, bool) {
	if o == nil || IsNil(o.Replicas) {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *MetaClusterInfo) HasReplicas() bool {
	if o != nil && !IsNil(o.Replicas) {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given []ClusterInfoReplicasInner and assigns it to the Replicas field.
func (o *MetaClusterInfo) SetReplicas(v []ClusterInfoReplicasInner) {
	o.Replicas = v
}

func (o MetaClusterInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaClusterInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster_size"] = o.ClusterSize
	if !IsNil(o.Leader) {
		toSerialize["leader"] = o.Leader
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Replicas) {
		toSerialize["replicas"] = o.Replicas
	}
	return toSerialize, nil
}

type NullableMetaClusterInfo struct {
	value *MetaClusterInfo
	isSet bool
}

func (v NullableMetaClusterInfo) Get() *MetaClusterInfo {
	return v.value
}

func (v *NullableMetaClusterInfo) Set(val *MetaClusterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaClusterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaClusterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaClusterInfo(val *MetaClusterInfo) *NullableMetaClusterInfo {
	return &NullableMetaClusterInfo{value: val, isSet: true}
}

func (v NullableMetaClusterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaClusterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
