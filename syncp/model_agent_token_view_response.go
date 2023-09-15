/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
	"time"
)

// checks if the AgentTokenViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentTokenViewResponse{}

// AgentTokenViewResponse struct for AgentTokenViewResponse
type AgentTokenViewResponse struct {
	Created        time.Time  `json:"created"`
	Id             string     `json:"id"`
	IsCurrent      bool       `json:"is_current"`
	LastAccessedAt *time.Time `json:"last_accessed_at,omitempty"`
	NkeyPublic     string     `json:"nkey_public"`
	RotatedAt      *time.Time `json:"rotated_at,omitempty"`
}

// NewAgentTokenViewResponse instantiates a new AgentTokenViewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentTokenViewResponse(created time.Time, id string, isCurrent bool, nkeyPublic string) *AgentTokenViewResponse {
	this := AgentTokenViewResponse{}
	this.Created = created
	this.Id = id
	this.IsCurrent = isCurrent
	this.NkeyPublic = nkeyPublic
	return &this
}

// NewAgentTokenViewResponseWithDefaults instantiates a new AgentTokenViewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentTokenViewResponseWithDefaults() *AgentTokenViewResponse {
	this := AgentTokenViewResponse{}
	return &this
}

// GetCreated returns the Created field value
func (o *AgentTokenViewResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *AgentTokenViewResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *AgentTokenViewResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetId returns the Id field value
func (o *AgentTokenViewResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AgentTokenViewResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AgentTokenViewResponse) SetId(v string) {
	o.Id = v
}

// GetIsCurrent returns the IsCurrent field value
func (o *AgentTokenViewResponse) GetIsCurrent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCurrent
}

// GetIsCurrentOk returns a tuple with the IsCurrent field value
// and a boolean to check if the value has been set.
func (o *AgentTokenViewResponse) GetIsCurrentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCurrent, true
}

// SetIsCurrent sets field value
func (o *AgentTokenViewResponse) SetIsCurrent(v bool) {
	o.IsCurrent = v
}

// GetLastAccessedAt returns the LastAccessedAt field value if set, zero value otherwise.
func (o *AgentTokenViewResponse) GetLastAccessedAt() time.Time {
	if o == nil || IsNil(o.LastAccessedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastAccessedAt
}

// GetLastAccessedAtOk returns a tuple with the LastAccessedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentTokenViewResponse) GetLastAccessedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastAccessedAt) {
		return nil, false
	}
	return o.LastAccessedAt, true
}

// HasLastAccessedAt returns a boolean if a field has been set.
func (o *AgentTokenViewResponse) HasLastAccessedAt() bool {
	if o != nil && !IsNil(o.LastAccessedAt) {
		return true
	}

	return false
}

// SetLastAccessedAt gets a reference to the given time.Time and assigns it to the LastAccessedAt field.
func (o *AgentTokenViewResponse) SetLastAccessedAt(v time.Time) {
	o.LastAccessedAt = &v
}

// GetNkeyPublic returns the NkeyPublic field value
func (o *AgentTokenViewResponse) GetNkeyPublic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NkeyPublic
}

// GetNkeyPublicOk returns a tuple with the NkeyPublic field value
// and a boolean to check if the value has been set.
func (o *AgentTokenViewResponse) GetNkeyPublicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NkeyPublic, true
}

// SetNkeyPublic sets field value
func (o *AgentTokenViewResponse) SetNkeyPublic(v string) {
	o.NkeyPublic = v
}

// GetRotatedAt returns the RotatedAt field value if set, zero value otherwise.
func (o *AgentTokenViewResponse) GetRotatedAt() time.Time {
	if o == nil || IsNil(o.RotatedAt) {
		var ret time.Time
		return ret
	}
	return *o.RotatedAt
}

// GetRotatedAtOk returns a tuple with the RotatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentTokenViewResponse) GetRotatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RotatedAt) {
		return nil, false
	}
	return o.RotatedAt, true
}

// HasRotatedAt returns a boolean if a field has been set.
func (o *AgentTokenViewResponse) HasRotatedAt() bool {
	if o != nil && !IsNil(o.RotatedAt) {
		return true
	}

	return false
}

// SetRotatedAt gets a reference to the given time.Time and assigns it to the RotatedAt field.
func (o *AgentTokenViewResponse) SetRotatedAt(v time.Time) {
	o.RotatedAt = &v
}

func (o AgentTokenViewResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentTokenViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["id"] = o.Id
	toSerialize["is_current"] = o.IsCurrent
	if !IsNil(o.LastAccessedAt) {
		toSerialize["last_accessed_at"] = o.LastAccessedAt
	}
	toSerialize["nkey_public"] = o.NkeyPublic
	if !IsNil(o.RotatedAt) {
		toSerialize["rotated_at"] = o.RotatedAt
	}
	return toSerialize, nil
}

type NullableAgentTokenViewResponse struct {
	value *AgentTokenViewResponse
	isSet bool
}

func (v NullableAgentTokenViewResponse) Get() *AgentTokenViewResponse {
	return v.value
}

func (v *NullableAgentTokenViewResponse) Set(val *AgentTokenViewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentTokenViewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentTokenViewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentTokenViewResponse(val *AgentTokenViewResponse) *NullableAgentTokenViewResponse {
	return &NullableAgentTokenViewResponse{value: val, isSet: true}
}

func (v NullableAgentTokenViewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentTokenViewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}