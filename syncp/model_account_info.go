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

// checks if the AccountInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountInfo{}

// AccountInfo struct for AccountInfo
type AccountInfo struct {
	AccountPublicKey     string `json:"account_public_key"`
	Id                   string `json:"id"`
	IsSystemAccount      bool   `json:"is_system_account"`
	Name                 string `json:"name"`
	UserJwtExpiresInSecs int64  `json:"user_jwt_expires_in_secs"`
}

// NewAccountInfo instantiates a new AccountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInfo(accountPublicKey string, id string, isSystemAccount bool, name string, userJwtExpiresInSecs int64) *AccountInfo {
	this := AccountInfo{}
	this.AccountPublicKey = accountPublicKey
	this.Id = id
	this.IsSystemAccount = isSystemAccount
	this.Name = name
	this.UserJwtExpiresInSecs = userJwtExpiresInSecs
	return &this
}

// NewAccountInfoWithDefaults instantiates a new AccountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInfoWithDefaults() *AccountInfo {
	this := AccountInfo{}
	return &this
}

// GetAccountPublicKey returns the AccountPublicKey field value
func (o *AccountInfo) GetAccountPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountPublicKey
}

// GetAccountPublicKeyOk returns a tuple with the AccountPublicKey field value
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetAccountPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountPublicKey, true
}

// SetAccountPublicKey sets field value
func (o *AccountInfo) SetAccountPublicKey(v string) {
	o.AccountPublicKey = v
}

// GetId returns the Id field value
func (o *AccountInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountInfo) SetId(v string) {
	o.Id = v
}

// GetIsSystemAccount returns the IsSystemAccount field value
func (o *AccountInfo) GetIsSystemAccount() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSystemAccount
}

// GetIsSystemAccountOk returns a tuple with the IsSystemAccount field value
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetIsSystemAccountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSystemAccount, true
}

// SetIsSystemAccount sets field value
func (o *AccountInfo) SetIsSystemAccount(v bool) {
	o.IsSystemAccount = v
}

// GetName returns the Name field value
func (o *AccountInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountInfo) SetName(v string) {
	o.Name = v
}

// GetUserJwtExpiresInSecs returns the UserJwtExpiresInSecs field value
func (o *AccountInfo) GetUserJwtExpiresInSecs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserJwtExpiresInSecs
}

// GetUserJwtExpiresInSecsOk returns a tuple with the UserJwtExpiresInSecs field value
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetUserJwtExpiresInSecsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserJwtExpiresInSecs, true
}

// SetUserJwtExpiresInSecs sets field value
func (o *AccountInfo) SetUserJwtExpiresInSecs(v int64) {
	o.UserJwtExpiresInSecs = v
}

func (o AccountInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_public_key"] = o.AccountPublicKey
	toSerialize["id"] = o.Id
	toSerialize["is_system_account"] = o.IsSystemAccount
	toSerialize["name"] = o.Name
	toSerialize["user_jwt_expires_in_secs"] = o.UserJwtExpiresInSecs
	return toSerialize, nil
}

type NullableAccountInfo struct {
	value *AccountInfo
	isSet bool
}

func (v NullableAccountInfo) Get() *AccountInfo {
	return v.value
}

func (v *NullableAccountInfo) Set(val *AccountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInfo(val *AccountInfo) *NullableAccountInfo {
	return &NullableAccountInfo{value: val, isSet: true}
}

func (v NullableAccountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
