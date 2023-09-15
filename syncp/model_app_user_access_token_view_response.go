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

// checks if the AppUserAccessTokenViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserAccessTokenViewResponse{}

// AppUserAccessTokenViewResponse struct for AppUserAccessTokenViewResponse
type AppUserAccessTokenViewResponse struct {
	Created time.Time      `json:"created"`
	Expires NullableString `json:"expires"`
	Id      string         `json:"id"`
	Name    string         `json:"name"`
}

// NewAppUserAccessTokenViewResponse instantiates a new AppUserAccessTokenViewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUserAccessTokenViewResponse(created time.Time, expires NullableString, id string, name string) *AppUserAccessTokenViewResponse {
	this := AppUserAccessTokenViewResponse{}
	this.Created = created
	this.Expires = expires
	this.Id = id
	this.Name = name
	return &this
}

// NewAppUserAccessTokenViewResponseWithDefaults instantiates a new AppUserAccessTokenViewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUserAccessTokenViewResponseWithDefaults() *AppUserAccessTokenViewResponse {
	this := AppUserAccessTokenViewResponse{}
	return &this
}

// GetCreated returns the Created field value
func (o *AppUserAccessTokenViewResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *AppUserAccessTokenViewResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *AppUserAccessTokenViewResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetExpires returns the Expires field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AppUserAccessTokenViewResponse) GetExpires() string {
	if o == nil || o.Expires.Get() == nil {
		var ret string
		return ret
	}

	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUserAccessTokenViewResponse) GetExpiresOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// SetExpires sets field value
func (o *AppUserAccessTokenViewResponse) SetExpires(v string) {
	o.Expires.Set(&v)
}

// GetId returns the Id field value
func (o *AppUserAccessTokenViewResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppUserAccessTokenViewResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppUserAccessTokenViewResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AppUserAccessTokenViewResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppUserAccessTokenViewResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppUserAccessTokenViewResponse) SetName(v string) {
	o.Name = v
}

func (o AppUserAccessTokenViewResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUserAccessTokenViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["expires"] = o.Expires.Get()
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableAppUserAccessTokenViewResponse struct {
	value *AppUserAccessTokenViewResponse
	isSet bool
}

func (v NullableAppUserAccessTokenViewResponse) Get() *AppUserAccessTokenViewResponse {
	return v.value
}

func (v *NullableAppUserAccessTokenViewResponse) Set(val *AppUserAccessTokenViewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUserAccessTokenViewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUserAccessTokenViewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUserAccessTokenViewResponse(val *AppUserAccessTokenViewResponse) *NullableAppUserAccessTokenViewResponse {
	return &NullableAppUserAccessTokenViewResponse{value: val, isSet: true}
}

func (v NullableAppUserAccessTokenViewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUserAccessTokenViewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}