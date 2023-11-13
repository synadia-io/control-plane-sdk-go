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

// checks if the Account type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Account{}

// Account struct for Account
type Account struct {
	Description *string `json:"description,omitempty"`
	InfoUrl     *string `json:"info_url,omitempty"`
	// TagList is a unique array of lower case strings All tag list methods lower case the strings in the arguments
	Tags []string `json:"tags,omitempty"`
	// ClaimType is used to indicate the type of JWT being stored in a Claim
	Type               *string        `json:"type,omitempty"`
	Version            *int32         `json:"version,omitempty"`
	Authorization      *Authorization `json:"authorization,omitempty"`
	DefaultPermissions *Permissions   `json:"default_permissions,omitempty"`
	// Exports is a slice of exports
	Exports []ExportsInner `json:"exports,omitempty"`
	// Imports is a list of import structs
	Imports  []ImportsInner                `json:"imports,omitempty"`
	Limits   *OperatorLimits               `json:"limits,omitempty"`
	Mappings *map[string][]WeightedMapping `json:"mappings,omitempty"`
	// RevocationList is used to store a mapping of public keys to unix timestamps
	Revocations *map[string]int64 `json:"revocations,omitempty"`
	// SigningKeys is a map keyed by a public account key
	SigningKeys interface{} `json:"signing_keys,omitempty"`
}

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount() *Account {
	this := Account{}
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Account) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Account) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Account) SetDescription(v string) {
	o.Description = &v
}

// GetInfoUrl returns the InfoUrl field value if set, zero value otherwise.
func (o *Account) GetInfoUrl() string {
	if o == nil || IsNil(o.InfoUrl) {
		var ret string
		return ret
	}
	return *o.InfoUrl
}

// GetInfoUrlOk returns a tuple with the InfoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetInfoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.InfoUrl) {
		return nil, false
	}
	return o.InfoUrl, true
}

// HasInfoUrl returns a boolean if a field has been set.
func (o *Account) HasInfoUrl() bool {
	if o != nil && !IsNil(o.InfoUrl) {
		return true
	}

	return false
}

// SetInfoUrl gets a reference to the given string and assigns it to the InfoUrl field.
func (o *Account) SetInfoUrl(v string) {
	o.InfoUrl = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Account) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Account) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Account) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Account) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Account) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Account) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Account) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Account) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *Account) SetVersion(v int32) {
	o.Version = &v
}

// GetAuthorization returns the Authorization field value if set, zero value otherwise.
func (o *Account) GetAuthorization() Authorization {
	if o == nil || IsNil(o.Authorization) {
		var ret Authorization
		return ret
	}
	return *o.Authorization
}

// GetAuthorizationOk returns a tuple with the Authorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAuthorizationOk() (*Authorization, bool) {
	if o == nil || IsNil(o.Authorization) {
		return nil, false
	}
	return o.Authorization, true
}

// HasAuthorization returns a boolean if a field has been set.
func (o *Account) HasAuthorization() bool {
	if o != nil && !IsNil(o.Authorization) {
		return true
	}

	return false
}

// SetAuthorization gets a reference to the given Authorization and assigns it to the Authorization field.
func (o *Account) SetAuthorization(v Authorization) {
	o.Authorization = &v
}

// GetDefaultPermissions returns the DefaultPermissions field value if set, zero value otherwise.
func (o *Account) GetDefaultPermissions() Permissions {
	if o == nil || IsNil(o.DefaultPermissions) {
		var ret Permissions
		return ret
	}
	return *o.DefaultPermissions
}

// GetDefaultPermissionsOk returns a tuple with the DefaultPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetDefaultPermissionsOk() (*Permissions, bool) {
	if o == nil || IsNil(o.DefaultPermissions) {
		return nil, false
	}
	return o.DefaultPermissions, true
}

// HasDefaultPermissions returns a boolean if a field has been set.
func (o *Account) HasDefaultPermissions() bool {
	if o != nil && !IsNil(o.DefaultPermissions) {
		return true
	}

	return false
}

// SetDefaultPermissions gets a reference to the given Permissions and assigns it to the DefaultPermissions field.
func (o *Account) SetDefaultPermissions(v Permissions) {
	o.DefaultPermissions = &v
}

// GetExports returns the Exports field value if set, zero value otherwise.
func (o *Account) GetExports() []ExportsInner {
	if o == nil || IsNil(o.Exports) {
		var ret []ExportsInner
		return ret
	}
	return o.Exports
}

// GetExportsOk returns a tuple with the Exports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetExportsOk() ([]ExportsInner, bool) {
	if o == nil || IsNil(o.Exports) {
		return nil, false
	}
	return o.Exports, true
}

// HasExports returns a boolean if a field has been set.
func (o *Account) HasExports() bool {
	if o != nil && !IsNil(o.Exports) {
		return true
	}

	return false
}

// SetExports gets a reference to the given []ExportsInner and assigns it to the Exports field.
func (o *Account) SetExports(v []ExportsInner) {
	o.Exports = v
}

// GetImports returns the Imports field value if set, zero value otherwise.
func (o *Account) GetImports() []ImportsInner {
	if o == nil || IsNil(o.Imports) {
		var ret []ImportsInner
		return ret
	}
	return o.Imports
}

// GetImportsOk returns a tuple with the Imports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetImportsOk() ([]ImportsInner, bool) {
	if o == nil || IsNil(o.Imports) {
		return nil, false
	}
	return o.Imports, true
}

// HasImports returns a boolean if a field has been set.
func (o *Account) HasImports() bool {
	if o != nil && !IsNil(o.Imports) {
		return true
	}

	return false
}

// SetImports gets a reference to the given []ImportsInner and assigns it to the Imports field.
func (o *Account) SetImports(v []ImportsInner) {
	o.Imports = v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *Account) GetLimits() OperatorLimits {
	if o == nil || IsNil(o.Limits) {
		var ret OperatorLimits
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetLimitsOk() (*OperatorLimits, bool) {
	if o == nil || IsNil(o.Limits) {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *Account) HasLimits() bool {
	if o != nil && !IsNil(o.Limits) {
		return true
	}

	return false
}

// SetLimits gets a reference to the given OperatorLimits and assigns it to the Limits field.
func (o *Account) SetLimits(v OperatorLimits) {
	o.Limits = &v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *Account) GetMappings() map[string][]WeightedMapping {
	if o == nil || IsNil(o.Mappings) {
		var ret map[string][]WeightedMapping
		return ret
	}
	return *o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetMappingsOk() (*map[string][]WeightedMapping, bool) {
	if o == nil || IsNil(o.Mappings) {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *Account) HasMappings() bool {
	if o != nil && !IsNil(o.Mappings) {
		return true
	}

	return false
}

// SetMappings gets a reference to the given map[string][]WeightedMapping and assigns it to the Mappings field.
func (o *Account) SetMappings(v map[string][]WeightedMapping) {
	o.Mappings = &v
}

// GetRevocations returns the Revocations field value if set, zero value otherwise.
func (o *Account) GetRevocations() map[string]int64 {
	if o == nil || IsNil(o.Revocations) {
		var ret map[string]int64
		return ret
	}
	return *o.Revocations
}

// GetRevocationsOk returns a tuple with the Revocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetRevocationsOk() (*map[string]int64, bool) {
	if o == nil || IsNil(o.Revocations) {
		return nil, false
	}
	return o.Revocations, true
}

// HasRevocations returns a boolean if a field has been set.
func (o *Account) HasRevocations() bool {
	if o != nil && !IsNil(o.Revocations) {
		return true
	}

	return false
}

// SetRevocations gets a reference to the given map[string]int64 and assigns it to the Revocations field.
func (o *Account) SetRevocations(v map[string]int64) {
	o.Revocations = &v
}

// GetSigningKeys returns the SigningKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Account) GetSigningKeys() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SigningKeys
}

// GetSigningKeysOk returns a tuple with the SigningKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Account) GetSigningKeysOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SigningKeys) {
		return nil, false
	}
	return &o.SigningKeys, true
}

// HasSigningKeys returns a boolean if a field has been set.
func (o *Account) HasSigningKeys() bool {
	if o != nil && IsNil(o.SigningKeys) {
		return true
	}

	return false
}

// SetSigningKeys gets a reference to the given interface{} and assigns it to the SigningKeys field.
func (o *Account) SetSigningKeys(v interface{}) {
	o.SigningKeys = v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Account) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.InfoUrl) {
		toSerialize["info_url"] = o.InfoUrl
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Authorization) {
		toSerialize["authorization"] = o.Authorization
	}
	if !IsNil(o.DefaultPermissions) {
		toSerialize["default_permissions"] = o.DefaultPermissions
	}
	if !IsNil(o.Exports) {
		toSerialize["exports"] = o.Exports
	}
	if !IsNil(o.Imports) {
		toSerialize["imports"] = o.Imports
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
	if o.SigningKeys != nil {
		toSerialize["signing_keys"] = o.SigningKeys
	}
	return toSerialize, nil
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
