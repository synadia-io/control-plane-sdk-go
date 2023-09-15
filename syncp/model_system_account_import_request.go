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

// checks if the SystemAccountImportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemAccountImportRequest{}

// SystemAccountImportRequest struct for SystemAccountImportRequest
type SystemAccountImportRequest struct {
	Jwt         string   `json:"jwt"`
	Seed        string   `json:"seed"`
	SigningKeys []string `json:"signing_keys,omitempty"`
}

// NewSystemAccountImportRequest instantiates a new SystemAccountImportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemAccountImportRequest(jwt string, seed string) *SystemAccountImportRequest {
	this := SystemAccountImportRequest{}
	this.Jwt = jwt
	this.Seed = seed
	return &this
}

// NewSystemAccountImportRequestWithDefaults instantiates a new SystemAccountImportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemAccountImportRequestWithDefaults() *SystemAccountImportRequest {
	this := SystemAccountImportRequest{}
	return &this
}

// GetJwt returns the Jwt field value
func (o *SystemAccountImportRequest) GetJwt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Jwt
}

// GetJwtOk returns a tuple with the Jwt field value
// and a boolean to check if the value has been set.
func (o *SystemAccountImportRequest) GetJwtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jwt, true
}

// SetJwt sets field value
func (o *SystemAccountImportRequest) SetJwt(v string) {
	o.Jwt = v
}

// GetSeed returns the Seed field value
func (o *SystemAccountImportRequest) GetSeed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Seed
}

// GetSeedOk returns a tuple with the Seed field value
// and a boolean to check if the value has been set.
func (o *SystemAccountImportRequest) GetSeedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seed, true
}

// SetSeed sets field value
func (o *SystemAccountImportRequest) SetSeed(v string) {
	o.Seed = v
}

// GetSigningKeys returns the SigningKeys field value if set, zero value otherwise.
func (o *SystemAccountImportRequest) GetSigningKeys() []string {
	if o == nil || IsNil(o.SigningKeys) {
		var ret []string
		return ret
	}
	return o.SigningKeys
}

// GetSigningKeysOk returns a tuple with the SigningKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemAccountImportRequest) GetSigningKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.SigningKeys) {
		return nil, false
	}
	return o.SigningKeys, true
}

// HasSigningKeys returns a boolean if a field has been set.
func (o *SystemAccountImportRequest) HasSigningKeys() bool {
	if o != nil && !IsNil(o.SigningKeys) {
		return true
	}

	return false
}

// SetSigningKeys gets a reference to the given []string and assigns it to the SigningKeys field.
func (o *SystemAccountImportRequest) SetSigningKeys(v []string) {
	o.SigningKeys = v
}

func (o SystemAccountImportRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemAccountImportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jwt"] = o.Jwt
	toSerialize["seed"] = o.Seed
	if !IsNil(o.SigningKeys) {
		toSerialize["signing_keys"] = o.SigningKeys
	}
	return toSerialize, nil
}

type NullableSystemAccountImportRequest struct {
	value *SystemAccountImportRequest
	isSet bool
}

func (v NullableSystemAccountImportRequest) Get() *SystemAccountImportRequest {
	return v.value
}

func (v *NullableSystemAccountImportRequest) Set(val *SystemAccountImportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemAccountImportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemAccountImportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemAccountImportRequest(val *SystemAccountImportRequest) *NullableSystemAccountImportRequest {
	return &NullableSystemAccountImportRequest{value: val, isSet: true}
}

func (v NullableSystemAccountImportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemAccountImportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}