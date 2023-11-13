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

// checks if the ApexCoordinate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApexCoordinate{}

// ApexCoordinate struct for ApexCoordinate
type ApexCoordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type _ApexCoordinate ApexCoordinate

// NewApexCoordinate instantiates a new ApexCoordinate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApexCoordinate(x float64, y float64) *ApexCoordinate {
	this := ApexCoordinate{}
	this.X = x
	this.Y = y
	return &this
}

// NewApexCoordinateWithDefaults instantiates a new ApexCoordinate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApexCoordinateWithDefaults() *ApexCoordinate {
	this := ApexCoordinate{}
	return &this
}

// GetX returns the X field value
func (o *ApexCoordinate) GetX() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *ApexCoordinate) GetXOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *ApexCoordinate) SetX(v float64) {
	o.X = v
}

// GetY returns the Y field value
func (o *ApexCoordinate) GetY() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *ApexCoordinate) GetYOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *ApexCoordinate) SetY(v float64) {
	o.Y = v
}

func (o ApexCoordinate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApexCoordinate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	return toSerialize, nil
}

func (o *ApexCoordinate) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"x",
		"y",
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

	varApexCoordinate := _ApexCoordinate{}

	err = json.Unmarshal(bytes, &varApexCoordinate)

	if err != nil {
		return err
	}

	*o = ApexCoordinate(varApexCoordinate)

	return err
}

type NullableApexCoordinate struct {
	value *ApexCoordinate
	isSet bool
}

func (v NullableApexCoordinate) Get() *ApexCoordinate {
	return v.value
}

func (v *NullableApexCoordinate) Set(val *ApexCoordinate) {
	v.value = val
	v.isSet = true
}

func (v NullableApexCoordinate) IsSet() bool {
	return v.isSet
}

func (v *NullableApexCoordinate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApexCoordinate(val *ApexCoordinate) *NullableApexCoordinate {
	return &NullableApexCoordinate{value: val, isSet: true}
}

func (v NullableApexCoordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApexCoordinate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
