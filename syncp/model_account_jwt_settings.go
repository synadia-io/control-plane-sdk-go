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

// checks if the AccountJWTSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountJWTSettings{}

// AccountJWTSettings struct for AccountJWTSettings
type AccountJWTSettings struct {
	Authorization *Authorization                `json:"authorization,omitempty"`
	Limits        *OperatorLimits               `json:"limits,omitempty"`
	Mappings      *map[string][]WeightedMapping `json:"mappings,omitempty"`
	// RevocationList is used to store a mapping of public keys to unix timestamps
	Revocations *map[string]int64 `json:"revocations,omitempty"`
	// TagList is a unique array of lower case strings All tag list methods lower case the strings in the arguments
	Tags        []string `json:"tags,omitempty"`
	Description *string  `json:"description,omitempty"`
	InfoUrl     *string  `json:"info_url,omitempty"`
}

// NewAccountJWTSettings instantiates a new AccountJWTSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountJWTSettings() *AccountJWTSettings {
	this := AccountJWTSettings{}
	return &this
}

// NewAccountJWTSettingsWithDefaults instantiates a new AccountJWTSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountJWTSettingsWithDefaults() *AccountJWTSettings {
	this := AccountJWTSettings{}
	return &this
}

// GetAuthorization returns the Authorization field value if set, zero value otherwise.
func (o *AccountJWTSettings) GetAuthorization() Authorization {
	if o == nil || IsNil(o.Authorization) {
		var ret Authorization
		return ret
	}
	return *o.Authorization
}

// GetAuthorizationOk returns a tuple with the Authorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountJWTSettings) GetAuthorizationOk() (*Authorization, bool) {
	if o == nil || IsNil(o.Authorization) {
		return nil, false
	}
	return o.Authorization, true
}

// HasAuthorization returns a boolean if a field has been set.
func (o *AccountJWTSettings) HasAuthorization() bool {
	if o != nil && !IsNil(o.Authorization) {
		return true
	}

	return false
}

// SetAuthorization gets a reference to the given Authorization and assigns it to the Authorization field.
func (o *AccountJWTSettings) SetAuthorization(v Authorization) {
	o.Authorization = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *AccountJWTSettings) GetLimits() OperatorLimits {
	if o == nil || IsNil(o.Limits) {
		var ret OperatorLimits
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountJWTSettings) GetLimitsOk() (*OperatorLimits, bool) {
	if o == nil || IsNil(o.Limits) {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *AccountJWTSettings) HasLimits() bool {
	if o != nil && !IsNil(o.Limits) {
		return true
	}

	return false
}

// SetLimits gets a reference to the given OperatorLimits and assigns it to the Limits field.
func (o *AccountJWTSettings) SetLimits(v OperatorLimits) {
	o.Limits = &v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *AccountJWTSettings) GetMappings() map[string][]WeightedMapping {
	if o == nil || IsNil(o.Mappings) {
		var ret map[string][]WeightedMapping
		return ret
	}
	return *o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountJWTSettings) GetMappingsOk() (*map[string][]WeightedMapping, bool) {
	if o == nil || IsNil(o.Mappings) {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *AccountJWTSettings) HasMappings() bool {
	if o != nil && !IsNil(o.Mappings) {
		return true
	}

	return false
}

// SetMappings gets a reference to the given map[string][]WeightedMapping and assigns it to the Mappings field.
func (o *AccountJWTSettings) SetMappings(v map[string][]WeightedMapping) {
	o.Mappings = &v
}

// GetRevocations returns the Revocations field value if set, zero value otherwise.
func (o *AccountJWTSettings) GetRevocations() map[string]int64 {
	if o == nil || IsNil(o.Revocations) {
		var ret map[string]int64
		return ret
	}
	return *o.Revocations
}

// GetRevocationsOk returns a tuple with the Revocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountJWTSettings) GetRevocationsOk() (*map[string]int64, bool) {
	if o == nil || IsNil(o.Revocations) {
		return nil, false
	}
	return o.Revocations, true
}

// HasRevocations returns a boolean if a field has been set.
func (o *AccountJWTSettings) HasRevocations() bool {
	if o != nil && !IsNil(o.Revocations) {
		return true
	}

	return false
}

// SetRevocations gets a reference to the given map[string]int64 and assigns it to the Revocations field.
func (o *AccountJWTSettings) SetRevocations(v map[string]int64) {
	o.Revocations = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountJWTSettings) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountJWTSettings) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AccountJWTSettings) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AccountJWTSettings) SetTags(v []string) {
	o.Tags = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccountJWTSettings) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountJWTSettings) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccountJWTSettings) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccountJWTSettings) SetDescription(v string) {
	o.Description = &v
}

// GetInfoUrl returns the InfoUrl field value if set, zero value otherwise.
func (o *AccountJWTSettings) GetInfoUrl() string {
	if o == nil || IsNil(o.InfoUrl) {
		var ret string
		return ret
	}
	return *o.InfoUrl
}

// GetInfoUrlOk returns a tuple with the InfoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountJWTSettings) GetInfoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.InfoUrl) {
		return nil, false
	}
	return o.InfoUrl, true
}

// HasInfoUrl returns a boolean if a field has been set.
func (o *AccountJWTSettings) HasInfoUrl() bool {
	if o != nil && !IsNil(o.InfoUrl) {
		return true
	}

	return false
}

// SetInfoUrl gets a reference to the given string and assigns it to the InfoUrl field.
func (o *AccountJWTSettings) SetInfoUrl(v string) {
	o.InfoUrl = &v
}

func (o AccountJWTSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountJWTSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authorization) {
		toSerialize["authorization"] = o.Authorization
	}
	if !IsNil(o.Limits) {
		toSerialize["limits"] = o.Limits
	}
	if !IsNil(o.Mappings) {
		toSerialize["mappings"] = o.Mappings
	}
	if !IsNil(o.Revocations) {
		toSerialize["revocations"] = o.Revocations
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.InfoUrl) {
		toSerialize["info_url"] = o.InfoUrl
	}
	return toSerialize, nil
}

type NullableAccountJWTSettings struct {
	value *AccountJWTSettings
	isSet bool
}

func (v NullableAccountJWTSettings) Get() *AccountJWTSettings {
	return v.value
}

func (v *NullableAccountJWTSettings) Set(val *AccountJWTSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountJWTSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountJWTSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountJWTSettings(val *AccountJWTSettings) *NullableAccountJWTSettings {
	return &NullableAccountJWTSettings{value: val, isSet: true}
}

func (v NullableAccountJWTSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountJWTSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
