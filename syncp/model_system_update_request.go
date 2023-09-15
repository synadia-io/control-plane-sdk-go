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

// checks if the SystemUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemUpdateRequest{}

// SystemUpdateRequest struct for SystemUpdateRequest
type SystemUpdateRequest struct {
	JetstreamDomain  NullableString `json:"jetstream_domain,omitempty"`
	JetstreamEnabled *bool          `json:"jetstream_enabled,omitempty"`
	Name             string         `json:"name"`
	Url              NullableString `json:"url,omitempty"`
}

// NewSystemUpdateRequest instantiates a new SystemUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemUpdateRequest(name string) *SystemUpdateRequest {
	this := SystemUpdateRequest{}
	this.Name = name
	return &this
}

// NewSystemUpdateRequestWithDefaults instantiates a new SystemUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemUpdateRequestWithDefaults() *SystemUpdateRequest {
	this := SystemUpdateRequest{}
	return &this
}

// GetJetstreamDomain returns the JetstreamDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemUpdateRequest) GetJetstreamDomain() string {
	if o == nil || IsNil(o.JetstreamDomain.Get()) {
		var ret string
		return ret
	}
	return *o.JetstreamDomain.Get()
}

// GetJetstreamDomainOk returns a tuple with the JetstreamDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemUpdateRequest) GetJetstreamDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JetstreamDomain.Get(), o.JetstreamDomain.IsSet()
}

// HasJetstreamDomain returns a boolean if a field has been set.
func (o *SystemUpdateRequest) HasJetstreamDomain() bool {
	if o != nil && o.JetstreamDomain.IsSet() {
		return true
	}

	return false
}

// SetJetstreamDomain gets a reference to the given NullableString and assigns it to the JetstreamDomain field.
func (o *SystemUpdateRequest) SetJetstreamDomain(v string) {
	o.JetstreamDomain.Set(&v)
}

// SetJetstreamDomainNil sets the value for JetstreamDomain to be an explicit nil
func (o *SystemUpdateRequest) SetJetstreamDomainNil() {
	o.JetstreamDomain.Set(nil)
}

// UnsetJetstreamDomain ensures that no value is present for JetstreamDomain, not even an explicit nil
func (o *SystemUpdateRequest) UnsetJetstreamDomain() {
	o.JetstreamDomain.Unset()
}

// GetJetstreamEnabled returns the JetstreamEnabled field value if set, zero value otherwise.
func (o *SystemUpdateRequest) GetJetstreamEnabled() bool {
	if o == nil || IsNil(o.JetstreamEnabled) {
		var ret bool
		return ret
	}
	return *o.JetstreamEnabled
}

// GetJetstreamEnabledOk returns a tuple with the JetstreamEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemUpdateRequest) GetJetstreamEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.JetstreamEnabled) {
		return nil, false
	}
	return o.JetstreamEnabled, true
}

// HasJetstreamEnabled returns a boolean if a field has been set.
func (o *SystemUpdateRequest) HasJetstreamEnabled() bool {
	if o != nil && !IsNil(o.JetstreamEnabled) {
		return true
	}

	return false
}

// SetJetstreamEnabled gets a reference to the given bool and assigns it to the JetstreamEnabled field.
func (o *SystemUpdateRequest) SetJetstreamEnabled(v bool) {
	o.JetstreamEnabled = &v
}

// GetName returns the Name field value
func (o *SystemUpdateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SystemUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SystemUpdateRequest) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemUpdateRequest) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemUpdateRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *SystemUpdateRequest) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *SystemUpdateRequest) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *SystemUpdateRequest) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *SystemUpdateRequest) UnsetUrl() {
	o.Url.Unset()
}

func (o SystemUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.JetstreamDomain.IsSet() {
		toSerialize["jetstream_domain"] = o.JetstreamDomain.Get()
	}
	if !IsNil(o.JetstreamEnabled) {
		toSerialize["jetstream_enabled"] = o.JetstreamEnabled
	}
	toSerialize["name"] = o.Name
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return toSerialize, nil
}

type NullableSystemUpdateRequest struct {
	value *SystemUpdateRequest
	isSet bool
}

func (v NullableSystemUpdateRequest) Get() *SystemUpdateRequest {
	return v.value
}

func (v *NullableSystemUpdateRequest) Set(val *SystemUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemUpdateRequest(val *SystemUpdateRequest) *NullableSystemUpdateRequest {
	return &NullableSystemUpdateRequest{value: val, isSet: true}
}

func (v NullableSystemUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}