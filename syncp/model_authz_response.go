/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
	"fmt"
)

// checks if the AuthzResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthzResponse{}

// AuthzResponse struct for AuthzResponse
type AuthzResponse struct {
	Operations []string `json:"operations"`
	ResourceId string   `json:"resource_id"`
	Service    string   `json:"service"`
}

type _AuthzResponse AuthzResponse

// NewAuthzResponse instantiates a new AuthzResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthzResponse(operations []string, resourceId string, service string) *AuthzResponse {
	this := AuthzResponse{}
	this.Operations = operations
	this.ResourceId = resourceId
	this.Service = service
	return &this
}

// NewAuthzResponseWithDefaults instantiates a new AuthzResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthzResponseWithDefaults() *AuthzResponse {
	this := AuthzResponse{}
	return &this
}

// GetOperations returns the Operations field value
func (o *AuthzResponse) GetOperations() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value
// and a boolean to check if the value has been set.
func (o *AuthzResponse) GetOperationsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operations, true
}

// SetOperations sets field value
func (o *AuthzResponse) SetOperations(v []string) {
	o.Operations = v
}

// GetResourceId returns the ResourceId field value
func (o *AuthzResponse) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *AuthzResponse) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *AuthzResponse) SetResourceId(v string) {
	o.ResourceId = v
}

// GetService returns the Service field value
func (o *AuthzResponse) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *AuthzResponse) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *AuthzResponse) SetService(v string) {
	o.Service = v
}

func (o AuthzResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthzResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operations"] = o.Operations
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["service"] = o.Service
	return toSerialize, nil
}

func (o *AuthzResponse) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operations",
		"resource_id",
		"service",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAuthzResponse := _AuthzResponse{}

	err = json.Unmarshal(bytes, &varAuthzResponse)

	if err != nil {
		return err
	}

	*o = AuthzResponse(varAuthzResponse)

	return err
}

type NullableAuthzResponse struct {
	value *AuthzResponse
	isSet bool
}

func (v NullableAuthzResponse) Get() *AuthzResponse {
	return v.value
}

func (v *NullableAuthzResponse) Set(val *AuthzResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthzResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthzResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthzResponse(val *AuthzResponse) *NullableAuthzResponse {
	return &NullableAuthzResponse{value: val, isSet: true}
}

func (v NullableAuthzResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthzResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
