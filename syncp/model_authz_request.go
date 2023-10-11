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

// checks if the AuthzRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthzRequest{}

// AuthzRequest struct for AuthzRequest
type AuthzRequest struct {
	ResourceId string `json:"resource_id"`
	Service    string `json:"service"`
}

// NewAuthzRequest instantiates a new AuthzRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthzRequest(resourceId string, service string) *AuthzRequest {
	this := AuthzRequest{}
	this.ResourceId = resourceId
	this.Service = service
	return &this
}

// NewAuthzRequestWithDefaults instantiates a new AuthzRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthzRequestWithDefaults() *AuthzRequest {
	this := AuthzRequest{}
	var resourceId string = ""
	this.ResourceId = resourceId
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *AuthzRequest) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *AuthzRequest) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *AuthzRequest) SetResourceId(v string) {
	o.ResourceId = v
}

// GetService returns the Service field value
func (o *AuthzRequest) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *AuthzRequest) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *AuthzRequest) SetService(v string) {
	o.Service = v
}

func (o AuthzRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthzRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["service"] = o.Service
	return toSerialize, nil
}

type NullableAuthzRequest struct {
	value *AuthzRequest
	isSet bool
}

func (v NullableAuthzRequest) Get() *AuthzRequest {
	return v.value
}

func (v *NullableAuthzRequest) Set(val *AuthzRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthzRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthzRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthzRequest(val *AuthzRequest) *NullableAuthzRequest {
	return &NullableAuthzRequest{value: val, isSet: true}
}

func (v NullableAuthzRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthzRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
