/*
Synadia Control Plane

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UserPermissionLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPermissionLimits{}

// UserPermissionLimits struct for UserPermissionLimits
type UserPermissionLimits struct {
	// StringList is a wrapper for an array of strings
	AllowedConnectionTypes []string                `json:"allowed_connection_types,omitempty"`
	BearerToken            *bool                   `json:"bearer_token,omitempty"`
	Pub                    *Permission             `json:"pub,omitempty"`
	Resp                   NullablePermissionsResp `json:"resp,omitempty"`
	Sub                    *Permission             `json:"sub,omitempty"`
	Src                    []string                `json:"src,omitempty"`
	Times                  []TimeRange             `json:"times,omitempty"`
	TimesLocation          *string                 `json:"times_location,omitempty"`
	Data                   *int64                  `json:"data,omitempty"`
	Payload                *int64                  `json:"payload,omitempty"`
	Subs                   *int64                  `json:"subs,omitempty"`
}

// NewUserPermissionLimits instantiates a new UserPermissionLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPermissionLimits() *UserPermissionLimits {
	this := UserPermissionLimits{}
	return &this
}

// NewUserPermissionLimitsWithDefaults instantiates a new UserPermissionLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPermissionLimitsWithDefaults() *UserPermissionLimits {
	this := UserPermissionLimits{}
	return &this
}

// GetAllowedConnectionTypes returns the AllowedConnectionTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserPermissionLimits) GetAllowedConnectionTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedConnectionTypes
}

// GetAllowedConnectionTypesOk returns a tuple with the AllowedConnectionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPermissionLimits) GetAllowedConnectionTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedConnectionTypes) {
		return nil, false
	}
	return o.AllowedConnectionTypes, true
}

// HasAllowedConnectionTypes returns a boolean if a field has been set.
func (o *UserPermissionLimits) HasAllowedConnectionTypes() bool {
	if o != nil && IsNil(o.AllowedConnectionTypes) {
		return true
	}

	return false
}

// SetAllowedConnectionTypes gets a reference to the given []string and assigns it to the AllowedConnectionTypes field.
func (o *UserPermissionLimits) SetAllowedConnectionTypes(v []string) {
	o.AllowedConnectionTypes = v
}

// GetBearerToken returns the BearerToken field value if set, zero value otherwise.
func (o *UserPermissionLimits) GetBearerToken() bool {
	if o == nil || IsNil(o.BearerToken) {
		var ret bool
		return ret
	}
	return *o.BearerToken
}

// GetBearerTokenOk returns a tuple with the BearerToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPermissionLimits) GetBearerTokenOk() (*bool, bool) {
	if o == nil || IsNil(o.BearerToken) {
		return nil, false
	}
	return o.BearerToken, true
}

// HasBearerToken returns a boolean if a field has been set.
func (o *UserPermissionLimits) HasBearerToken() bool {
	if o != nil && !IsNil(o.BearerToken) {
		return true
	}

	return false
}

// SetBearerToken gets a reference to the given bool and assigns it to the BearerToken field.
func (o *UserPermissionLimits) SetBearerToken(v bool) {
	o.BearerToken = &v
}

// GetPub returns the Pub field value if set, zero value otherwise.
func (o *UserPermissionLimits) GetPub() Permission {
	if o == nil || IsNil(o.Pub) {
		var ret Permission
		return ret
	}
	return *o.Pub
}

// GetPubOk returns a tuple with the Pub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPermissionLimits) GetPubOk() (*Permission, bool) {
	if o == nil || IsNil(o.Pub) {
		return nil, false
	}
	return o.Pub, true
}

// HasPub returns a boolean if a field has been set.
func (o *UserPermissionLimits) HasPub() bool {
	if o != nil && !IsNil(o.Pub) {
		return true
	}

	return false
}

// SetPub gets a reference to the given Permission and assigns it to the Pub field.
func (o *UserPermissionLimits) SetPub(v Permission) {
	o.Pub = &v
}

// GetResp returns the Resp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserPermissionLimits) GetResp() PermissionsResp {
	if o == nil || IsNil(o.Resp.Get()) {
		var ret PermissionsResp
		return ret
	}
	return *o.Resp.Get()
}

// GetRespOk returns a tuple with the Resp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPermissionLimits) GetRespOk() (*PermissionsResp, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resp.Get(), o.Resp.IsSet()
}

// HasResp returns a boolean if a field has been set.
func (o *UserPermissionLimits) HasResp() bool {
	if o != nil && o.Resp.IsSet() {
		return true
	}

	return false
}

// SetResp gets a reference to the given NullablePermissionsResp and assigns it to the Resp field.
func (o *UserPermissionLimits) SetResp(v PermissionsResp) {
	o.Resp.Set(&v)
}

// SetRespNil sets the value for Resp to be an explicit nil
func (o *UserPermissionLimits) SetRespNil() {
	o.Resp.Set(nil)
}

// UnsetResp ensures that no value is present for Resp, not even an explicit nil
func (o *UserPermissionLimits) UnsetResp() {
	o.Resp.Unset()
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *UserPermissionLimits) GetSub() Permission {
	if o == nil || IsNil(o.Sub) {
		var ret Permission
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPermissionLimits) GetSubOk() (*Permission, bool) {
	if o == nil || IsNil(o.Sub) {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *UserPermissionLimits) HasSub() bool {
	if o != nil && !IsNil(o.Sub) {
		return true
	}

	return false
}

// SetSub gets a reference to the given Permission and assigns it to the Sub field.
func (o *UserPermissionLimits) SetSub(v Permission) {
	o.Sub = &v
}

// GetSrc returns the Src field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserPermissionLimits) GetSrc() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Src
}

// GetSrcOk returns a tuple with the Src field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserPermissionLimits) GetSrcOk() ([]string, bool) {
	if o == nil || IsNil(o.Src) {
		return nil, false
	}
	return o.Src, true
}

// HasSrc returns a boolean if a field has been set.
func (o *UserPermissionLimits) HasSrc() bool {
	if o != nil && IsNil(o.Src) {
		return true
	}

	return false
}

// SetSrc gets a reference to the given []string and assigns it to the Src field.
func (o *UserPermissionLimits) SetSrc(v []string) {
	o.Src = v
}

// GetTimes returns the Times field value if set, zero value otherwise.
func (o *UserPermissionLimits) GetTimes() []TimeRange {
	if o == nil || IsNil(o.Times) {
		var ret []TimeRange
		return ret
	}
	return o.Times
}

// GetTimesOk returns a tuple with the Times field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPermissionLimits) GetTimesOk() ([]TimeRange, bool) {
	if o == nil || IsNil(o.Times) {
		return nil, false
	}
	return o.Times, true
}

// HasTimes returns a boolean if a field has been set.
func (o *UserPermissionLimits) HasTimes() bool {
	if o != nil && !IsNil(o.Times) {
		return true
	}

	return false
}

// SetTimes gets a reference to the given []TimeRange and assigns it to the Times field.
func (o *UserPermissionLimits) SetTimes(v []TimeRange) {
	o.Times = v
}

// GetTimesLocation returns the TimesLocation field value if set, zero value otherwise.
func (o *UserPermissionLimits) GetTimesLocation() string {
	if o == nil || IsNil(o.TimesLocation) {
		var ret string
		return ret
	}
	return *o.TimesLocation
}

// GetTimesLocationOk returns a tuple with the TimesLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPermissionLimits) GetTimesLocationOk() (*string, bool) {
	if o == nil || IsNil(o.TimesLocation) {
		return nil, false
	}
	return o.TimesLocation, true
}

// HasTimesLocation returns a boolean if a field has been set.
func (o *UserPermissionLimits) HasTimesLocation() bool {
	if o != nil && !IsNil(o.TimesLocation) {
		return true
	}

	return false
}

// SetTimesLocation gets a reference to the given string and assigns it to the TimesLocation field.
func (o *UserPermissionLimits) SetTimesLocation(v string) {
	o.TimesLocation = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UserPermissionLimits) GetData() int64 {
	if o == nil || IsNil(o.Data) {
		var ret int64
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPermissionLimits) GetDataOk() (*int64, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UserPermissionLimits) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given int64 and assigns it to the Data field.
func (o *UserPermissionLimits) SetData(v int64) {
	o.Data = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *UserPermissionLimits) GetPayload() int64 {
	if o == nil || IsNil(o.Payload) {
		var ret int64
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPermissionLimits) GetPayloadOk() (*int64, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *UserPermissionLimits) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given int64 and assigns it to the Payload field.
func (o *UserPermissionLimits) SetPayload(v int64) {
	o.Payload = &v
}

// GetSubs returns the Subs field value if set, zero value otherwise.
func (o *UserPermissionLimits) GetSubs() int64 {
	if o == nil || IsNil(o.Subs) {
		var ret int64
		return ret
	}
	return *o.Subs
}

// GetSubsOk returns a tuple with the Subs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPermissionLimits) GetSubsOk() (*int64, bool) {
	if o == nil || IsNil(o.Subs) {
		return nil, false
	}
	return o.Subs, true
}

// HasSubs returns a boolean if a field has been set.
func (o *UserPermissionLimits) HasSubs() bool {
	if o != nil && !IsNil(o.Subs) {
		return true
	}

	return false
}

// SetSubs gets a reference to the given int64 and assigns it to the Subs field.
func (o *UserPermissionLimits) SetSubs(v int64) {
	o.Subs = &v
}

func (o UserPermissionLimits) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPermissionLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedConnectionTypes != nil {
		toSerialize["allowed_connection_types"] = o.AllowedConnectionTypes
	}
	if !IsNil(o.BearerToken) {
		toSerialize["bearer_token"] = o.BearerToken
	}
	if !IsNil(o.Pub) {
		toSerialize["pub"] = o.Pub
	}
	if o.Resp.IsSet() {
		toSerialize["resp"] = o.Resp.Get()
	}
	if !IsNil(o.Sub) {
		toSerialize["sub"] = o.Sub
	}
	if o.Src != nil {
		toSerialize["src"] = o.Src
	}
	if !IsNil(o.Times) {
		toSerialize["times"] = o.Times
	}
	if !IsNil(o.TimesLocation) {
		toSerialize["times_location"] = o.TimesLocation
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Subs) {
		toSerialize["subs"] = o.Subs
	}
	return toSerialize, nil
}

type NullableUserPermissionLimits struct {
	value *UserPermissionLimits
	isSet bool
}

func (v NullableUserPermissionLimits) Get() *UserPermissionLimits {
	return v.value
}

func (v *NullableUserPermissionLimits) Set(val *UserPermissionLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPermissionLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPermissionLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPermissionLimits(val *UserPermissionLimits) *NullableUserPermissionLimits {
	return &NullableUserPermissionLimits{value: val, isSet: true}
}

func (v NullableUserPermissionLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPermissionLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
