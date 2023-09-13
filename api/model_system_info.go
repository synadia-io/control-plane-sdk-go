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

// checks if the SystemInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemInfo{}

// SystemInfo struct for SystemInfo
type SystemInfo struct {
	Id                   string         `json:"id"`
	JetstreamDomain      NullableString `json:"jetstream_domain,omitempty"`
	JetstreamEnabled     bool           `json:"jetstream_enabled"`
	Name                 string         `json:"name"`
	ServerUrls           *string        `json:"server_urls,omitempty"`
	State                SystemState    `json:"state"`
	UserJwtExpiresInSecs int64          `json:"user_jwt_expires_in_secs"`
}

// NewSystemInfo instantiates a new SystemInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemInfo(id string, jetstreamEnabled bool, name string, state SystemState, userJwtExpiresInSecs int64) *SystemInfo {
	this := SystemInfo{}
	this.Id = id
	this.JetstreamEnabled = jetstreamEnabled
	this.Name = name
	this.State = state
	this.UserJwtExpiresInSecs = userJwtExpiresInSecs
	return &this
}

// NewSystemInfoWithDefaults instantiates a new SystemInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemInfoWithDefaults() *SystemInfo {
	this := SystemInfo{}
	return &this
}

// GetId returns the Id field value
func (o *SystemInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SystemInfo) SetId(v string) {
	o.Id = v
}

// GetJetstreamDomain returns the JetstreamDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemInfo) GetJetstreamDomain() string {
	if o == nil || IsNil(o.JetstreamDomain.Get()) {
		var ret string
		return ret
	}
	return *o.JetstreamDomain.Get()
}

// GetJetstreamDomainOk returns a tuple with the JetstreamDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemInfo) GetJetstreamDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JetstreamDomain.Get(), o.JetstreamDomain.IsSet()
}

// HasJetstreamDomain returns a boolean if a field has been set.
func (o *SystemInfo) HasJetstreamDomain() bool {
	if o != nil && o.JetstreamDomain.IsSet() {
		return true
	}

	return false
}

// SetJetstreamDomain gets a reference to the given NullableString and assigns it to the JetstreamDomain field.
func (o *SystemInfo) SetJetstreamDomain(v string) {
	o.JetstreamDomain.Set(&v)
}

// SetJetstreamDomainNil sets the value for JetstreamDomain to be an explicit nil
func (o *SystemInfo) SetJetstreamDomainNil() {
	o.JetstreamDomain.Set(nil)
}

// UnsetJetstreamDomain ensures that no value is present for JetstreamDomain, not even an explicit nil
func (o *SystemInfo) UnsetJetstreamDomain() {
	o.JetstreamDomain.Unset()
}

// GetJetstreamEnabled returns the JetstreamEnabled field value
func (o *SystemInfo) GetJetstreamEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.JetstreamEnabled
}

// GetJetstreamEnabledOk returns a tuple with the JetstreamEnabled field value
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetJetstreamEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JetstreamEnabled, true
}

// SetJetstreamEnabled sets field value
func (o *SystemInfo) SetJetstreamEnabled(v bool) {
	o.JetstreamEnabled = v
}

// GetName returns the Name field value
func (o *SystemInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SystemInfo) SetName(v string) {
	o.Name = v
}

// GetServerUrls returns the ServerUrls field value if set, zero value otherwise.
func (o *SystemInfo) GetServerUrls() string {
	if o == nil || IsNil(o.ServerUrls) {
		var ret string
		return ret
	}
	return *o.ServerUrls
}

// GetServerUrlsOk returns a tuple with the ServerUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetServerUrlsOk() (*string, bool) {
	if o == nil || IsNil(o.ServerUrls) {
		return nil, false
	}
	return o.ServerUrls, true
}

// HasServerUrls returns a boolean if a field has been set.
func (o *SystemInfo) HasServerUrls() bool {
	if o != nil && !IsNil(o.ServerUrls) {
		return true
	}

	return false
}

// SetServerUrls gets a reference to the given string and assigns it to the ServerUrls field.
func (o *SystemInfo) SetServerUrls(v string) {
	o.ServerUrls = &v
}

// GetState returns the State field value
func (o *SystemInfo) GetState() SystemState {
	if o == nil {
		var ret SystemState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetStateOk() (*SystemState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SystemInfo) SetState(v SystemState) {
	o.State = v
}

// GetUserJwtExpiresInSecs returns the UserJwtExpiresInSecs field value
func (o *SystemInfo) GetUserJwtExpiresInSecs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserJwtExpiresInSecs
}

// GetUserJwtExpiresInSecsOk returns a tuple with the UserJwtExpiresInSecs field value
// and a boolean to check if the value has been set.
func (o *SystemInfo) GetUserJwtExpiresInSecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserJwtExpiresInSecs, true
}

// SetUserJwtExpiresInSecs sets field value
func (o *SystemInfo) SetUserJwtExpiresInSecs(v int64) {
	o.UserJwtExpiresInSecs = v
}

func (o SystemInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.JetstreamDomain.IsSet() {
		toSerialize["jetstream_domain"] = o.JetstreamDomain.Get()
	}
	toSerialize["jetstream_enabled"] = o.JetstreamEnabled
	toSerialize["name"] = o.Name
	if !IsNil(o.ServerUrls) {
		toSerialize["server_urls"] = o.ServerUrls
	}
	toSerialize["state"] = o.State
	toSerialize["user_jwt_expires_in_secs"] = o.UserJwtExpiresInSecs
	return toSerialize, nil
}

type NullableSystemInfo struct {
	value *SystemInfo
	isSet bool
}

func (v NullableSystemInfo) Get() *SystemInfo {
	return v.value
}

func (v *NullableSystemInfo) Set(val *SystemInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemInfo(val *SystemInfo) *NullableSystemInfo {
	return &NullableSystemInfo{value: val, isSet: true}
}

func (v NullableSystemInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
