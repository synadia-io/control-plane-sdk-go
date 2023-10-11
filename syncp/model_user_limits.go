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

// checks if the UserLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserLimits{}

// UserLimits Src is a comma separated list of CIDR specifications
type UserLimits struct {
	Src           []string    `json:"src,omitempty"`
	Times         []TimeRange `json:"times,omitempty"`
	TimesLocation *string     `json:"times_location,omitempty"`
}

// NewUserLimits instantiates a new UserLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLimits() *UserLimits {
	this := UserLimits{}
	return &this
}

// NewUserLimitsWithDefaults instantiates a new UserLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLimitsWithDefaults() *UserLimits {
	this := UserLimits{}
	return &this
}

// GetSrc returns the Src field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserLimits) GetSrc() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Src
}

// GetSrcOk returns a tuple with the Src field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserLimits) GetSrcOk() ([]string, bool) {
	if o == nil || IsNil(o.Src) {
		return nil, false
	}
	return o.Src, true
}

// HasSrc returns a boolean if a field has been set.
func (o *UserLimits) HasSrc() bool {
	if o != nil && IsNil(o.Src) {
		return true
	}

	return false
}

// SetSrc gets a reference to the given []string and assigns it to the Src field.
func (o *UserLimits) SetSrc(v []string) {
	o.Src = v
}

// GetTimes returns the Times field value if set, zero value otherwise.
func (o *UserLimits) GetTimes() []TimeRange {
	if o == nil || IsNil(o.Times) {
		var ret []TimeRange
		return ret
	}
	return o.Times
}

// GetTimesOk returns a tuple with the Times field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLimits) GetTimesOk() ([]TimeRange, bool) {
	if o == nil || IsNil(o.Times) {
		return nil, false
	}
	return o.Times, true
}

// HasTimes returns a boolean if a field has been set.
func (o *UserLimits) HasTimes() bool {
	if o != nil && !IsNil(o.Times) {
		return true
	}

	return false
}

// SetTimes gets a reference to the given []TimeRange and assigns it to the Times field.
func (o *UserLimits) SetTimes(v []TimeRange) {
	o.Times = v
}

// GetTimesLocation returns the TimesLocation field value if set, zero value otherwise.
func (o *UserLimits) GetTimesLocation() string {
	if o == nil || IsNil(o.TimesLocation) {
		var ret string
		return ret
	}
	return *o.TimesLocation
}

// GetTimesLocationOk returns a tuple with the TimesLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLimits) GetTimesLocationOk() (*string, bool) {
	if o == nil || IsNil(o.TimesLocation) {
		return nil, false
	}
	return o.TimesLocation, true
}

// HasTimesLocation returns a boolean if a field has been set.
func (o *UserLimits) HasTimesLocation() bool {
	if o != nil && !IsNil(o.TimesLocation) {
		return true
	}

	return false
}

// SetTimesLocation gets a reference to the given string and assigns it to the TimesLocation field.
func (o *UserLimits) SetTimesLocation(v string) {
	o.TimesLocation = &v
}

func (o UserLimits) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Src != nil {
		toSerialize["src"] = o.Src
	}
	if !IsNil(o.Times) {
		toSerialize["times"] = o.Times
	}
	if !IsNil(o.TimesLocation) {
		toSerialize["times_location"] = o.TimesLocation
	}
	return toSerialize, nil
}

type NullableUserLimits struct {
	value *UserLimits
	isSet bool
}

func (v NullableUserLimits) Get() *UserLimits {
	return v.value
}

func (v *NullableUserLimits) Set(val *UserLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLimits(val *UserLimits) *NullableUserLimits {
	return &NullableUserLimits{value: val, isSet: true}
}

func (v NullableUserLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
