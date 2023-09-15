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

// checks if the Activation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Activation{}

// Activation Activation defines the custom parts of an activation claim
type Activation struct {
	GenericFields GenericFields `json:"GenericFields"`
	// IssuerAccount stores the public key for the account the issuer represents. When set, the claim was issued by a signing key.
	IssuerAccount *string     `json:"issuer_account,omitempty"`
	Kind          *ExportType `json:"kind,omitempty"`
	// Subject is a string that represents a NATS subject
	Subject *string `json:"subject,omitempty"`
}

// NewActivation instantiates a new Activation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivation(genericFields GenericFields) *Activation {
	this := Activation{}
	this.GenericFields = genericFields
	return &this
}

// NewActivationWithDefaults instantiates a new Activation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivationWithDefaults() *Activation {
	this := Activation{}
	return &this
}

// GetGenericFields returns the GenericFields field value
func (o *Activation) GetGenericFields() GenericFields {
	if o == nil {
		var ret GenericFields
		return ret
	}

	return o.GenericFields
}

// GetGenericFieldsOk returns a tuple with the GenericFields field value
// and a boolean to check if the value has been set.
func (o *Activation) GetGenericFieldsOk() (*GenericFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GenericFields, true
}

// SetGenericFields sets field value
func (o *Activation) SetGenericFields(v GenericFields) {
	o.GenericFields = v
}

// GetIssuerAccount returns the IssuerAccount field value if set, zero value otherwise.
func (o *Activation) GetIssuerAccount() string {
	if o == nil || IsNil(o.IssuerAccount) {
		var ret string
		return ret
	}
	return *o.IssuerAccount
}

// GetIssuerAccountOk returns a tuple with the IssuerAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activation) GetIssuerAccountOk() (*string, bool) {
	if o == nil || IsNil(o.IssuerAccount) {
		return nil, false
	}
	return o.IssuerAccount, true
}

// HasIssuerAccount returns a boolean if a field has been set.
func (o *Activation) HasIssuerAccount() bool {
	if o != nil && !IsNil(o.IssuerAccount) {
		return true
	}

	return false
}

// SetIssuerAccount gets a reference to the given string and assigns it to the IssuerAccount field.
func (o *Activation) SetIssuerAccount(v string) {
	o.IssuerAccount = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Activation) GetKind() ExportType {
	if o == nil || IsNil(o.Kind) {
		var ret ExportType
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activation) GetKindOk() (*ExportType, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Activation) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given ExportType and assigns it to the Kind field.
func (o *Activation) SetKind(v ExportType) {
	o.Kind = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *Activation) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activation) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *Activation) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *Activation) SetSubject(v string) {
	o.Subject = &v
}

func (o Activation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Activation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["GenericFields"] = o.GenericFields
	if !IsNil(o.IssuerAccount) {
		toSerialize["issuer_account"] = o.IssuerAccount
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	return toSerialize, nil
}

type NullableActivation struct {
	value *Activation
	isSet bool
}

func (v NullableActivation) Get() *Activation {
	return v.value
}

func (v *NullableActivation) Set(val *Activation) {
	v.value = val
	v.isSet = true
}

func (v NullableActivation) IsSet() bool {
	return v.isSet
}

func (v *NullableActivation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivation(val *Activation) *NullableActivation {
	return &NullableActivation{value: val, isSet: true}
}

func (v NullableActivation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
