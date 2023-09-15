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

// checks if the RouteStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteStat{}

// RouteStat RouteStat holds route statistics.
type RouteStat struct {
	Name     *string   `json:"name,omitempty"`
	Pending  int32     `json:"pending"`
	Received DataStats `json:"received"`
	Rid      int32     `json:"rid"`
	Sent     DataStats `json:"sent"`
}

// NewRouteStat instantiates a new RouteStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteStat(pending int32, received DataStats, rid int32, sent DataStats) *RouteStat {
	this := RouteStat{}
	this.Pending = pending
	this.Received = received
	this.Rid = rid
	this.Sent = sent
	return &this
}

// NewRouteStatWithDefaults instantiates a new RouteStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteStatWithDefaults() *RouteStat {
	this := RouteStat{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RouteStat) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteStat) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RouteStat) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RouteStat) SetName(v string) {
	o.Name = &v
}

// GetPending returns the Pending field value
func (o *RouteStat) GetPending() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value
// and a boolean to check if the value has been set.
func (o *RouteStat) GetPendingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pending, true
}

// SetPending sets field value
func (o *RouteStat) SetPending(v int32) {
	o.Pending = v
}

// GetReceived returns the Received field value
func (o *RouteStat) GetReceived() DataStats {
	if o == nil {
		var ret DataStats
		return ret
	}

	return o.Received
}

// GetReceivedOk returns a tuple with the Received field value
// and a boolean to check if the value has been set.
func (o *RouteStat) GetReceivedOk() (*DataStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Received, true
}

// SetReceived sets field value
func (o *RouteStat) SetReceived(v DataStats) {
	o.Received = v
}

// GetRid returns the Rid field value
func (o *RouteStat) GetRid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rid
}

// GetRidOk returns a tuple with the Rid field value
// and a boolean to check if the value has been set.
func (o *RouteStat) GetRidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rid, true
}

// SetRid sets field value
func (o *RouteStat) SetRid(v int32) {
	o.Rid = v
}

// GetSent returns the Sent field value
func (o *RouteStat) GetSent() DataStats {
	if o == nil {
		var ret DataStats
		return ret
	}

	return o.Sent
}

// GetSentOk returns a tuple with the Sent field value
// and a boolean to check if the value has been set.
func (o *RouteStat) GetSentOk() (*DataStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sent, true
}

// SetSent sets field value
func (o *RouteStat) SetSent(v DataStats) {
	o.Sent = v
}

func (o RouteStat) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["pending"] = o.Pending
	toSerialize["received"] = o.Received
	toSerialize["rid"] = o.Rid
	toSerialize["sent"] = o.Sent
	return toSerialize, nil
}

type NullableRouteStat struct {
	value *RouteStat
	isSet bool
}

func (v NullableRouteStat) Get() *RouteStat {
	return v.value
}

func (v *NullableRouteStat) Set(val *RouteStat) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteStat) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteStat(val *RouteStat) *NullableRouteStat {
	return &NullableRouteStat{value: val, isSet: true}
}

func (v NullableRouteStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}