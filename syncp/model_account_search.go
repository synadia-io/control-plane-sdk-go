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

// checks if the AccountSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountSearch{}

// AccountSearch struct for AccountSearch
type AccountSearch struct {
	AccountPublicKey string `json:"account_public_key"`
	Id               string `json:"id"`
	IsScpAccount     bool   `json:"is_scp_account"`
	IsSystemAccount  bool   `json:"is_system_account"`
	Name             string `json:"name"`
}

type _AccountSearch AccountSearch

// NewAccountSearch instantiates a new AccountSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountSearch(accountPublicKey string, id string, isScpAccount bool, isSystemAccount bool, name string) *AccountSearch {
	this := AccountSearch{}
	this.AccountPublicKey = accountPublicKey
	this.Id = id
	this.IsScpAccount = isScpAccount
	this.IsSystemAccount = isSystemAccount
	this.Name = name
	return &this
}

// NewAccountSearchWithDefaults instantiates a new AccountSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountSearchWithDefaults() *AccountSearch {
	this := AccountSearch{}
	return &this
}

// GetAccountPublicKey returns the AccountPublicKey field value
func (o *AccountSearch) GetAccountPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountPublicKey
}

// GetAccountPublicKeyOk returns a tuple with the AccountPublicKey field value
// and a boolean to check if the value has been set.
func (o *AccountSearch) GetAccountPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountPublicKey, true
}

// SetAccountPublicKey sets field value
func (o *AccountSearch) SetAccountPublicKey(v string) {
	o.AccountPublicKey = v
}

// GetId returns the Id field value
func (o *AccountSearch) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountSearch) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountSearch) SetId(v string) {
	o.Id = v
}

// GetIsScpAccount returns the IsScpAccount field value
func (o *AccountSearch) GetIsScpAccount() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsScpAccount
}

// GetIsScpAccountOk returns a tuple with the IsScpAccount field value
// and a boolean to check if the value has been set.
func (o *AccountSearch) GetIsScpAccountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsScpAccount, true
}

// SetIsScpAccount sets field value
func (o *AccountSearch) SetIsScpAccount(v bool) {
	o.IsScpAccount = v
}

// GetIsSystemAccount returns the IsSystemAccount field value
func (o *AccountSearch) GetIsSystemAccount() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSystemAccount
}

// GetIsSystemAccountOk returns a tuple with the IsSystemAccount field value
// and a boolean to check if the value has been set.
func (o *AccountSearch) GetIsSystemAccountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSystemAccount, true
}

// SetIsSystemAccount sets field value
func (o *AccountSearch) SetIsSystemAccount(v bool) {
	o.IsSystemAccount = v
}

// GetName returns the Name field value
func (o *AccountSearch) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountSearch) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountSearch) SetName(v string) {
	o.Name = v
}

func (o AccountSearch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_public_key"] = o.AccountPublicKey
	toSerialize["id"] = o.Id
	toSerialize["is_scp_account"] = o.IsScpAccount
	toSerialize["is_system_account"] = o.IsSystemAccount
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *AccountSearch) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_public_key",
		"id",
		"is_scp_account",
		"is_system_account",
		"name",
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

	varAccountSearch := _AccountSearch{}

	err = json.Unmarshal(bytes, &varAccountSearch)

	if err != nil {
		return err
	}

	*o = AccountSearch(varAccountSearch)

	return err
}

type NullableAccountSearch struct {
	value *AccountSearch
	isSet bool
}

func (v NullableAccountSearch) Get() *AccountSearch {
	return v.value
}

func (v *NullableAccountSearch) Set(val *AccountSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSearch(val *AccountSearch) *NullableAccountSearch {
	return &NullableAccountSearch{value: val, isSet: true}
}

func (v NullableAccountSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
