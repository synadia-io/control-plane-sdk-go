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

// checks if the AppPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPolicy{}

// AppPolicy struct for AppPolicy
type AppPolicy struct {
	Description *string                       `json:"description,omitempty"`
	Name        string                        `json:"name"`
	Statements  map[string]AppPolicyStatement `json:"statements"`
}

// NewAppPolicy instantiates a new AppPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPolicy(name string, statements map[string]AppPolicyStatement) *AppPolicy {
	this := AppPolicy{}
	this.Name = name
	this.Statements = statements
	return &this
}

// NewAppPolicyWithDefaults instantiates a new AppPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPolicyWithDefaults() *AppPolicy {
	this := AppPolicy{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AppPolicy) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AppPolicy) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AppPolicy) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *AppPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppPolicy) SetName(v string) {
	o.Name = v
}

// GetStatements returns the Statements field value
func (o *AppPolicy) GetStatements() map[string]AppPolicyStatement {
	if o == nil {
		var ret map[string]AppPolicyStatement
		return ret
	}

	return o.Statements
}

// GetStatementsOk returns a tuple with the Statements field value
// and a boolean to check if the value has been set.
func (o *AppPolicy) GetStatementsOk() (*map[string]AppPolicyStatement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Statements, true
}

// SetStatements sets field value
func (o *AppPolicy) SetStatements(v map[string]AppPolicyStatement) {
	o.Statements = v
}

func (o AppPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	toSerialize["statements"] = o.Statements
	return toSerialize, nil
}

type NullableAppPolicy struct {
	value *AppPolicy
	isSet bool
}

func (v NullableAppPolicy) Get() *AppPolicy {
	return v.value
}

func (v *NullableAppPolicy) Set(val *AppPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPolicy(val *AppPolicy) *NullableAppPolicy {
	return &NullableAppPolicy{value: val, isSet: true}
}

func (v NullableAppPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
