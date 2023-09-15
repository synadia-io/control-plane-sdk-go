/*
Synadia Control Plane

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
	"time"
)

// checks if the SigningKeyGroupViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SigningKeyGroupViewResponse{}

// SigningKeyGroupViewResponse struct for SigningKeyGroupViewResponse
type SigningKeyGroupViewResponse struct {
	Created    time.Time             `json:"created"`
	Disabled   bool                  `json:"disabled"`
	DisabledAt *time.Time            `json:"disabled_at,omitempty"`
	Id         string                `json:"id"`
	IsScoped   bool                  `json:"is_scoped"`
	Name       string                `json:"name"`
	Scope      *UserPermissionLimits `json:"scope,omitempty"`
}

// NewSigningKeyGroupViewResponse instantiates a new SigningKeyGroupViewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSigningKeyGroupViewResponse(created time.Time, disabled bool, id string, isScoped bool, name string) *SigningKeyGroupViewResponse {
	this := SigningKeyGroupViewResponse{}
	this.Created = created
	this.Disabled = disabled
	this.Id = id
	this.IsScoped = isScoped
	this.Name = name
	return &this
}

// NewSigningKeyGroupViewResponseWithDefaults instantiates a new SigningKeyGroupViewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSigningKeyGroupViewResponseWithDefaults() *SigningKeyGroupViewResponse {
	this := SigningKeyGroupViewResponse{}
	return &this
}

// GetCreated returns the Created field value
func (o *SigningKeyGroupViewResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *SigningKeyGroupViewResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *SigningKeyGroupViewResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetDisabled returns the Disabled field value
func (o *SigningKeyGroupViewResponse) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *SigningKeyGroupViewResponse) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *SigningKeyGroupViewResponse) SetDisabled(v bool) {
	o.Disabled = v
}

// GetDisabledAt returns the DisabledAt field value if set, zero value otherwise.
func (o *SigningKeyGroupViewResponse) GetDisabledAt() time.Time {
	if o == nil || IsNil(o.DisabledAt) {
		var ret time.Time
		return ret
	}
	return *o.DisabledAt
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKeyGroupViewResponse) GetDisabledAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DisabledAt) {
		return nil, false
	}
	return o.DisabledAt, true
}

// HasDisabledAt returns a boolean if a field has been set.
func (o *SigningKeyGroupViewResponse) HasDisabledAt() bool {
	if o != nil && !IsNil(o.DisabledAt) {
		return true
	}

	return false
}

// SetDisabledAt gets a reference to the given time.Time and assigns it to the DisabledAt field.
func (o *SigningKeyGroupViewResponse) SetDisabledAt(v time.Time) {
	o.DisabledAt = &v
}

// GetId returns the Id field value
func (o *SigningKeyGroupViewResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SigningKeyGroupViewResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SigningKeyGroupViewResponse) SetId(v string) {
	o.Id = v
}

// GetIsScoped returns the IsScoped field value
func (o *SigningKeyGroupViewResponse) GetIsScoped() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsScoped
}

// GetIsScopedOk returns a tuple with the IsScoped field value
// and a boolean to check if the value has been set.
func (o *SigningKeyGroupViewResponse) GetIsScopedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsScoped, true
}

// SetIsScoped sets field value
func (o *SigningKeyGroupViewResponse) SetIsScoped(v bool) {
	o.IsScoped = v
}

// GetName returns the Name field value
func (o *SigningKeyGroupViewResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SigningKeyGroupViewResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SigningKeyGroupViewResponse) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *SigningKeyGroupViewResponse) GetScope() UserPermissionLimits {
	if o == nil || IsNil(o.Scope) {
		var ret UserPermissionLimits
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKeyGroupViewResponse) GetScopeOk() (*UserPermissionLimits, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *SigningKeyGroupViewResponse) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given UserPermissionLimits and assigns it to the Scope field.
func (o *SigningKeyGroupViewResponse) SetScope(v UserPermissionLimits) {
	o.Scope = &v
}

func (o SigningKeyGroupViewResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SigningKeyGroupViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["disabled"] = o.Disabled
	if !IsNil(o.DisabledAt) {
		toSerialize["disabled_at"] = o.DisabledAt
	}
	toSerialize["id"] = o.Id
	toSerialize["is_scoped"] = o.IsScoped
	toSerialize["name"] = o.Name
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	return toSerialize, nil
}

type NullableSigningKeyGroupViewResponse struct {
	value *SigningKeyGroupViewResponse
	isSet bool
}

func (v NullableSigningKeyGroupViewResponse) Get() *SigningKeyGroupViewResponse {
	return v.value
}

func (v *NullableSigningKeyGroupViewResponse) Set(val *SigningKeyGroupViewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSigningKeyGroupViewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSigningKeyGroupViewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSigningKeyGroupViewResponse(val *SigningKeyGroupViewResponse) *NullableSigningKeyGroupViewResponse {
	return &NullableSigningKeyGroupViewResponse{value: val, isSet: true}
}

func (v NullableSigningKeyGroupViewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSigningKeyGroupViewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
