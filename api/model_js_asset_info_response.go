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

// checks if the JSAssetInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSAssetInfoResponse{}

// JSAssetInfoResponse struct for JSAssetInfoResponse
type JSAssetInfoResponse struct {
	Config        *JSCommonStreamConfig `json:"config,omitempty"`
	Id            *string               `json:"id,omitempty"`
	JetstreamType *JSType               `json:"jetstream_type,omitempty"`
	Alternates    []StreamAlternate     `json:"alternates,omitempty"`
	Cluster       *ClusterInfo          `json:"cluster,omitempty"`
	Created       time.Time             `json:"created"`
	Sources       []StreamSourceInfo    `json:"sources,omitempty"`
	State         StreamState           `json:"state"`
}

// NewJSAssetInfoResponse instantiates a new JSAssetInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSAssetInfoResponse(created time.Time, state StreamState) *JSAssetInfoResponse {
	this := JSAssetInfoResponse{}
	this.Created = created
	this.State = state
	return &this
}

// NewJSAssetInfoResponseWithDefaults instantiates a new JSAssetInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSAssetInfoResponseWithDefaults() *JSAssetInfoResponse {
	this := JSAssetInfoResponse{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *JSAssetInfoResponse) GetConfig() JSCommonStreamConfig {
	if o == nil || IsNil(o.Config) {
		var ret JSCommonStreamConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSAssetInfoResponse) GetConfigOk() (*JSCommonStreamConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *JSAssetInfoResponse) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given JSCommonStreamConfig and assigns it to the Config field.
func (o *JSAssetInfoResponse) SetConfig(v JSCommonStreamConfig) {
	o.Config = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JSAssetInfoResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSAssetInfoResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JSAssetInfoResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *JSAssetInfoResponse) SetId(v string) {
	o.Id = &v
}

// GetJetstreamType returns the JetstreamType field value if set, zero value otherwise.
func (o *JSAssetInfoResponse) GetJetstreamType() JSType {
	if o == nil || IsNil(o.JetstreamType) {
		var ret JSType
		return ret
	}
	return *o.JetstreamType
}

// GetJetstreamTypeOk returns a tuple with the JetstreamType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSAssetInfoResponse) GetJetstreamTypeOk() (*JSType, bool) {
	if o == nil || IsNil(o.JetstreamType) {
		return nil, false
	}
	return o.JetstreamType, true
}

// HasJetstreamType returns a boolean if a field has been set.
func (o *JSAssetInfoResponse) HasJetstreamType() bool {
	if o != nil && !IsNil(o.JetstreamType) {
		return true
	}

	return false
}

// SetJetstreamType gets a reference to the given JSType and assigns it to the JetstreamType field.
func (o *JSAssetInfoResponse) SetJetstreamType(v JSType) {
	o.JetstreamType = &v
}

// GetAlternates returns the Alternates field value if set, zero value otherwise.
func (o *JSAssetInfoResponse) GetAlternates() []StreamAlternate {
	if o == nil || IsNil(o.Alternates) {
		var ret []StreamAlternate
		return ret
	}
	return o.Alternates
}

// GetAlternatesOk returns a tuple with the Alternates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSAssetInfoResponse) GetAlternatesOk() ([]StreamAlternate, bool) {
	if o == nil || IsNil(o.Alternates) {
		return nil, false
	}
	return o.Alternates, true
}

// HasAlternates returns a boolean if a field has been set.
func (o *JSAssetInfoResponse) HasAlternates() bool {
	if o != nil && !IsNil(o.Alternates) {
		return true
	}

	return false
}

// SetAlternates gets a reference to the given []StreamAlternate and assigns it to the Alternates field.
func (o *JSAssetInfoResponse) SetAlternates(v []StreamAlternate) {
	o.Alternates = v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *JSAssetInfoResponse) GetCluster() ClusterInfo {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterInfo
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSAssetInfoResponse) GetClusterOk() (*ClusterInfo, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *JSAssetInfoResponse) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterInfo and assigns it to the Cluster field.
func (o *JSAssetInfoResponse) SetCluster(v ClusterInfo) {
	o.Cluster = &v
}

// GetCreated returns the Created field value
func (o *JSAssetInfoResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *JSAssetInfoResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *JSAssetInfoResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *JSAssetInfoResponse) GetSources() []StreamSourceInfo {
	if o == nil || IsNil(o.Sources) {
		var ret []StreamSourceInfo
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSAssetInfoResponse) GetSourcesOk() ([]StreamSourceInfo, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *JSAssetInfoResponse) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []StreamSourceInfo and assigns it to the Sources field.
func (o *JSAssetInfoResponse) SetSources(v []StreamSourceInfo) {
	o.Sources = v
}

// GetState returns the State field value
func (o *JSAssetInfoResponse) GetState() StreamState {
	if o == nil {
		var ret StreamState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *JSAssetInfoResponse) GetStateOk() (*StreamState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *JSAssetInfoResponse) SetState(v StreamState) {
	o.State = v
}

func (o JSAssetInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JSAssetInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.JetstreamType) {
		toSerialize["jetstream_type"] = o.JetstreamType
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

type NullableJSAssetInfoResponse struct {
	value *JSAssetInfoResponse
	isSet bool
}

func (v NullableJSAssetInfoResponse) Get() *JSAssetInfoResponse {
	return v.value
}

func (v *NullableJSAssetInfoResponse) Set(val *JSAssetInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJSAssetInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJSAssetInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSAssetInfoResponse(val *JSAssetInfoResponse) *NullableJSAssetInfoResponse {
	return &NullableJSAssetInfoResponse{value: val, isSet: true}
}

func (v NullableJSAssetInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSAssetInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
