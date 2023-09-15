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

// checks if the SubDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubDetail{}

// SubDetail SubDetail is for verbose information for subscriptions.
type SubDetail struct {
	Account *string `json:"account,omitempty"`
	Cid     int32   `json:"cid"`
	Max     *int64  `json:"max,omitempty"`
	Msgs    int64   `json:"msgs"`
	Qgroup  *string `json:"qgroup,omitempty"`
	Sid     string  `json:"sid"`
	Subject string  `json:"subject"`
}

// NewSubDetail instantiates a new SubDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubDetail(cid int32, msgs int64, sid string, subject string) *SubDetail {
	this := SubDetail{}
	this.Cid = cid
	this.Msgs = msgs
	this.Sid = sid
	this.Subject = subject
	return &this
}

// NewSubDetailWithDefaults instantiates a new SubDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubDetailWithDefaults() *SubDetail {
	this := SubDetail{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *SubDetail) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubDetail) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *SubDetail) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *SubDetail) SetAccount(v string) {
	o.Account = &v
}

// GetCid returns the Cid field value
func (o *SubDetail) GetCid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *SubDetail) GetCidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *SubDetail) SetCid(v int32) {
	o.Cid = v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *SubDetail) GetMax() int64 {
	if o == nil || IsNil(o.Max) {
		var ret int64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubDetail) GetMaxOk() (*int64, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *SubDetail) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int64 and assigns it to the Max field.
func (o *SubDetail) SetMax(v int64) {
	o.Max = &v
}

// GetMsgs returns the Msgs field value
func (o *SubDetail) GetMsgs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Msgs
}

// GetMsgsOk returns a tuple with the Msgs field value
// and a boolean to check if the value has been set.
func (o *SubDetail) GetMsgsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msgs, true
}

// SetMsgs sets field value
func (o *SubDetail) SetMsgs(v int64) {
	o.Msgs = v
}

// GetQgroup returns the Qgroup field value if set, zero value otherwise.
func (o *SubDetail) GetQgroup() string {
	if o == nil || IsNil(o.Qgroup) {
		var ret string
		return ret
	}
	return *o.Qgroup
}

// GetQgroupOk returns a tuple with the Qgroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubDetail) GetQgroupOk() (*string, bool) {
	if o == nil || IsNil(o.Qgroup) {
		return nil, false
	}
	return o.Qgroup, true
}

// HasQgroup returns a boolean if a field has been set.
func (o *SubDetail) HasQgroup() bool {
	if o != nil && !IsNil(o.Qgroup) {
		return true
	}

	return false
}

// SetQgroup gets a reference to the given string and assigns it to the Qgroup field.
func (o *SubDetail) SetQgroup(v string) {
	o.Qgroup = &v
}

// GetSid returns the Sid field value
func (o *SubDetail) GetSid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sid
}

// GetSidOk returns a tuple with the Sid field value
// and a boolean to check if the value has been set.
func (o *SubDetail) GetSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sid, true
}

// SetSid sets field value
func (o *SubDetail) SetSid(v string) {
	o.Sid = v
}

// GetSubject returns the Subject field value
func (o *SubDetail) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *SubDetail) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *SubDetail) SetSubject(v string) {
	o.Subject = v
}

func (o SubDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	toSerialize["cid"] = o.Cid
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	toSerialize["msgs"] = o.Msgs
	if !IsNil(o.Qgroup) {
		toSerialize["qgroup"] = o.Qgroup
	}
	toSerialize["sid"] = o.Sid
	toSerialize["subject"] = o.Subject
	return toSerialize, nil
}

type NullableSubDetail struct {
	value *SubDetail
	isSet bool
}

func (v NullableSubDetail) Get() *SubDetail {
	return v.value
}

func (v *NullableSubDetail) Set(val *SubDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableSubDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableSubDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubDetail(val *SubDetail) *NullableSubDetail {
	return &NullableSubDetail{value: val, isSet: true}
}

func (v NullableSubDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
