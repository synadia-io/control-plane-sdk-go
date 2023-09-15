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

// checks if the AgentTokenCurrentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentTokenCurrentResponse{}

// AgentTokenCurrentResponse struct for AgentTokenCurrentResponse
type AgentTokenCurrentResponse struct {
	Token          *string    `json:"token,omitempty"`
	Url            *string    `json:"url,omitempty"`
	Created        time.Time  `json:"created"`
	Id             string     `json:"id"`
	IsCurrent      bool       `json:"is_current"`
	LastAccessedAt *time.Time `json:"last_accessed_at,omitempty"`
	NkeyPublic     string     `json:"nkey_public"`
	RotatedAt      *time.Time `json:"rotated_at,omitempty"`
}

// NewAgentTokenCurrentResponse instantiates a new AgentTokenCurrentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentTokenCurrentResponse(created time.Time, id string, isCurrent bool, nkeyPublic string) *AgentTokenCurrentResponse {
	this := AgentTokenCurrentResponse{}
	this.Created = created
	this.Id = id
	this.IsCurrent = isCurrent
	this.NkeyPublic = nkeyPublic
	return &this
}

// NewAgentTokenCurrentResponseWithDefaults instantiates a new AgentTokenCurrentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentTokenCurrentResponseWithDefaults() *AgentTokenCurrentResponse {
	this := AgentTokenCurrentResponse{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AgentTokenCurrentResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentTokenCurrentResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AgentTokenCurrentResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AgentTokenCurrentResponse) SetToken(v string) {
	o.Token = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AgentTokenCurrentResponse) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentTokenCurrentResponse) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AgentTokenCurrentResponse) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AgentTokenCurrentResponse) SetUrl(v string) {
	o.Url = &v
}

// GetCreated returns the Created field value
func (o *AgentTokenCurrentResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *AgentTokenCurrentResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *AgentTokenCurrentResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetId returns the Id field value
func (o *AgentTokenCurrentResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AgentTokenCurrentResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AgentTokenCurrentResponse) SetId(v string) {
	o.Id = v
}

// GetIsCurrent returns the IsCurrent field value
func (o *AgentTokenCurrentResponse) GetIsCurrent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCurrent
}

// GetIsCurrentOk returns a tuple with the IsCurrent field value
// and a boolean to check if the value has been set.
func (o *AgentTokenCurrentResponse) GetIsCurrentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCurrent, true
}

// SetIsCurrent sets field value
func (o *AgentTokenCurrentResponse) SetIsCurrent(v bool) {
	o.IsCurrent = v
}

// GetLastAccessedAt returns the LastAccessedAt field value if set, zero value otherwise.
func (o *AgentTokenCurrentResponse) GetLastAccessedAt() time.Time {
	if o == nil || IsNil(o.LastAccessedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastAccessedAt
}

// GetLastAccessedAtOk returns a tuple with the LastAccessedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentTokenCurrentResponse) GetLastAccessedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastAccessedAt) {
		return nil, false
	}
	return o.LastAccessedAt, true
}

// HasLastAccessedAt returns a boolean if a field has been set.
func (o *AgentTokenCurrentResponse) HasLastAccessedAt() bool {
	if o != nil && !IsNil(o.LastAccessedAt) {
		return true
	}

	return false
}

// SetLastAccessedAt gets a reference to the given time.Time and assigns it to the LastAccessedAt field.
func (o *AgentTokenCurrentResponse) SetLastAccessedAt(v time.Time) {
	o.LastAccessedAt = &v
}

// GetNkeyPublic returns the NkeyPublic field value
func (o *AgentTokenCurrentResponse) GetNkeyPublic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NkeyPublic
}

// GetNkeyPublicOk returns a tuple with the NkeyPublic field value
// and a boolean to check if the value has been set.
func (o *AgentTokenCurrentResponse) GetNkeyPublicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NkeyPublic, true
}

// SetNkeyPublic sets field value
func (o *AgentTokenCurrentResponse) SetNkeyPublic(v string) {
	o.NkeyPublic = v
}

// GetRotatedAt returns the RotatedAt field value if set, zero value otherwise.
func (o *AgentTokenCurrentResponse) GetRotatedAt() time.Time {
	if o == nil || IsNil(o.RotatedAt) {
		var ret time.Time
		return ret
	}
	return *o.RotatedAt
}

// GetRotatedAtOk returns a tuple with the RotatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentTokenCurrentResponse) GetRotatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RotatedAt) {
		return nil, false
	}
	return o.RotatedAt, true
}

// HasRotatedAt returns a boolean if a field has been set.
func (o *AgentTokenCurrentResponse) HasRotatedAt() bool {
	if o != nil && !IsNil(o.RotatedAt) {
		return true
	}

	return false
}

// SetRotatedAt gets a reference to the given time.Time and assigns it to the RotatedAt field.
func (o *AgentTokenCurrentResponse) SetRotatedAt(v time.Time) {
	o.RotatedAt = &v
}

func (o AgentTokenCurrentResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentTokenCurrentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
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

type NullableAgentTokenCurrentResponse struct {
	value *AgentTokenCurrentResponse
	isSet bool
}

func (v NullableAgentTokenCurrentResponse) Get() *AgentTokenCurrentResponse {
	return v.value
}

func (v *NullableAgentTokenCurrentResponse) Set(val *AgentTokenCurrentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentTokenCurrentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentTokenCurrentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentTokenCurrentResponse(val *AgentTokenCurrentResponse) *NullableAgentTokenCurrentResponse {
	return &NullableAgentTokenCurrentResponse{value: val, isSet: true}
}

func (v NullableAgentTokenCurrentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentTokenCurrentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
