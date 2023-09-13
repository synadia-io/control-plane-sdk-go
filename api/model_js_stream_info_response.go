/*
Synadia Control Plane

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the JSStreamInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSStreamInfoResponse{}

// JSStreamInfoResponse struct for JSStreamInfoResponse
type JSStreamInfoResponse struct {
	Config     *JSStreamConfigRequest `json:"config,omitempty"`
	Id         *string                `json:"id,omitempty"`
	Alternates []StreamAlternate      `json:"alternates,omitempty"`
	Cluster    *ClusterInfo           `json:"cluster,omitempty"`
	Created    time.Time              `json:"created"`
	Sources    []StreamSourceInfo     `json:"sources,omitempty"`
	State      StreamState            `json:"state"`
}

// NewJSStreamInfoResponse instantiates a new JSStreamInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSStreamInfoResponse(created time.Time, state StreamState) *JSStreamInfoResponse {
	this := JSStreamInfoResponse{}
	this.Created = created
	this.State = state
	return &this
}

// NewJSStreamInfoResponseWithDefaults instantiates a new JSStreamInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSStreamInfoResponseWithDefaults() *JSStreamInfoResponse {
	this := JSStreamInfoResponse{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *JSStreamInfoResponse) GetConfig() JSStreamConfigRequest {
	if o == nil || IsNil(o.Config) {
		var ret JSStreamConfigRequest
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSStreamInfoResponse) GetConfigOk() (*JSStreamConfigRequest, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *JSStreamInfoResponse) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given JSStreamConfigRequest and assigns it to the Config field.
func (o *JSStreamInfoResponse) SetConfig(v JSStreamConfigRequest) {
	o.Config = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JSStreamInfoResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSStreamInfoResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JSStreamInfoResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *JSStreamInfoResponse) SetId(v string) {
	o.Id = &v
}

// GetAlternates returns the Alternates field value if set, zero value otherwise.
func (o *JSStreamInfoResponse) GetAlternates() []StreamAlternate {
	if o == nil || IsNil(o.Alternates) {
		var ret []StreamAlternate
		return ret
	}
	return o.Alternates
}

// GetAlternatesOk returns a tuple with the Alternates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSStreamInfoResponse) GetAlternatesOk() ([]StreamAlternate, bool) {
	if o == nil || IsNil(o.Alternates) {
		return nil, false
	}
	return o.Alternates, true
}

// HasAlternates returns a boolean if a field has been set.
func (o *JSStreamInfoResponse) HasAlternates() bool {
	if o != nil && !IsNil(o.Alternates) {
		return true
	}

	return false
}

// SetAlternates gets a reference to the given []StreamAlternate and assigns it to the Alternates field.
func (o *JSStreamInfoResponse) SetAlternates(v []StreamAlternate) {
	o.Alternates = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *JSStreamInfoResponse) GetCluster() ClusterInfo {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterInfo
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSStreamInfoResponse) GetClusterOk() (*ClusterInfo, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *JSStreamInfoResponse) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterInfo and assigns it to the Cluster field.
func (o *JSStreamInfoResponse) SetCluster(v ClusterInfo) {
	o.Cluster = &v
}

// GetCreated returns the Created field value
func (o *JSStreamInfoResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *JSStreamInfoResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *JSStreamInfoResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *JSStreamInfoResponse) GetSources() []StreamSourceInfo {
	if o == nil || IsNil(o.Sources) {
		var ret []StreamSourceInfo
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSStreamInfoResponse) GetSourcesOk() ([]StreamSourceInfo, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *JSStreamInfoResponse) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []StreamSourceInfo and assigns it to the Sources field.
func (o *JSStreamInfoResponse) SetSources(v []StreamSourceInfo) {
	o.Sources = v
}

// GetState returns the State field value
func (o *JSStreamInfoResponse) GetState() StreamState {
	if o == nil {
		var ret StreamState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *JSStreamInfoResponse) GetStateOk() (*StreamState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *JSStreamInfoResponse) SetState(v StreamState) {
	o.State = v
}

func (o JSStreamInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JSStreamInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Alternates) {
		toSerialize["alternates"] = o.Alternates
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	toSerialize["created"] = o.Created
	if !IsNil(o.Sources) {
		toSerialize["sources"] = o.Sources
	}
	toSerialize["state"] = o.State
	return toSerialize, nil
}

type NullableJSStreamInfoResponse struct {
	value *JSStreamInfoResponse
	isSet bool
}

func (v NullableJSStreamInfoResponse) Get() *JSStreamInfoResponse {
	return v.value
}

func (v *NullableJSStreamInfoResponse) Set(val *JSStreamInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJSStreamInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJSStreamInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSStreamInfoResponse(val *JSStreamInfoResponse) *NullableJSStreamInfoResponse {
	return &NullableJSStreamInfoResponse{value: val, isSet: true}
}

func (v NullableJSStreamInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSStreamInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
