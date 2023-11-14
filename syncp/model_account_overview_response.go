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

// checks if the AccountOverviewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountOverviewResponse{}

// AccountOverviewResponse struct for AccountOverviewResponse
type AccountOverviewResponse struct {
	AccountPublicKey string                         `json:"account_public_key"`
	Metrics          AccountOverviewResponseMetrics `json:"metrics"`
}

type _AccountOverviewResponse AccountOverviewResponse

// NewAccountOverviewResponse instantiates a new AccountOverviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountOverviewResponse(accountPublicKey string, metrics AccountOverviewResponseMetrics) *AccountOverviewResponse {
	this := AccountOverviewResponse{}
	this.AccountPublicKey = accountPublicKey
	this.Metrics = metrics
	return &this
}

// NewAccountOverviewResponseWithDefaults instantiates a new AccountOverviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountOverviewResponseWithDefaults() *AccountOverviewResponse {
	this := AccountOverviewResponse{}
	return &this
}

// GetAccountPublicKey returns the AccountPublicKey field value
func (o *AccountOverviewResponse) GetAccountPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountPublicKey
}

// GetAccountPublicKeyOk returns a tuple with the AccountPublicKey field value
// and a boolean to check if the value has been set.
func (o *AccountOverviewResponse) GetAccountPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountPublicKey, true
}

// SetAccountPublicKey sets field value
func (o *AccountOverviewResponse) SetAccountPublicKey(v string) {
	o.AccountPublicKey = v
}

// GetMetrics returns the Metrics field value
func (o *AccountOverviewResponse) GetMetrics() AccountOverviewResponseMetrics {
	if o == nil {
		var ret AccountOverviewResponseMetrics
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *AccountOverviewResponse) GetMetricsOk() (*AccountOverviewResponseMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value
func (o *AccountOverviewResponse) SetMetrics(v AccountOverviewResponseMetrics) {
	o.Metrics = v
}

func (o AccountOverviewResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountOverviewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_public_key"] = o.AccountPublicKey
	toSerialize["metrics"] = o.Metrics
	return toSerialize, nil
}

func (o *AccountOverviewResponse) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_public_key",
		"metrics",
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

	varAccountOverviewResponse := _AccountOverviewResponse{}

	err = json.Unmarshal(bytes, &varAccountOverviewResponse)

	if err != nil {
		return err
	}

	*o = AccountOverviewResponse(varAccountOverviewResponse)

	return err
}

type NullableAccountOverviewResponse struct {
	value *AccountOverviewResponse
	isSet bool
}

func (v NullableAccountOverviewResponse) Get() *AccountOverviewResponse {
	return v.value
}

func (v *NullableAccountOverviewResponse) Set(val *AccountOverviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountOverviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountOverviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountOverviewResponse(val *AccountOverviewResponse) *NullableAccountOverviewResponse {
	return &NullableAccountOverviewResponse{value: val, isSet: true}
}

func (v NullableAccountOverviewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountOverviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}