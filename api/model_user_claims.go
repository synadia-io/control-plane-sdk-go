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

// checks if the UserClaims type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserClaims{}

// UserClaims UserClaims defines a user JWT
type UserClaims struct {
	Nats *User   `json:"nats,omitempty"`
	Aud  *string `json:"aud,omitempty"`
	Exp  *int64  `json:"exp,omitempty"`
	Iat  *int64  `json:"iat,omitempty"`
	Iss  *string `json:"iss,omitempty"`
	Jti  *string `json:"jti,omitempty"`
	Name *string `json:"name,omitempty"`
	Nbf  *int64  `json:"nbf,omitempty"`
	Sub  *string `json:"sub,omitempty"`
}

// NewUserClaims instantiates a new UserClaims object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserClaims() *UserClaims {
	this := UserClaims{}
	return &this
}

// NewUserClaimsWithDefaults instantiates a new UserClaims object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserClaimsWithDefaults() *UserClaims {
	this := UserClaims{}
	return &this
}

// GetNats returns the Nats field value if set, zero value otherwise.
func (o *UserClaims) GetNats() User {
	if o == nil || IsNil(o.Nats) {
		var ret User
		return ret
	}
	return *o.Nats
}

// GetNatsOk returns a tuple with the Nats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserClaims) GetNatsOk() (*User, bool) {
	if o == nil || IsNil(o.Nats) {
		return nil, false
	}
	return o.Nats, true
}

// HasNats returns a boolean if a field has been set.
func (o *UserClaims) HasNats() bool {
	if o != nil && !IsNil(o.Nats) {
		return true
	}

	return false
}

// SetNats gets a reference to the given User and assigns it to the Nats field.
func (o *UserClaims) SetNats(v User) {
	o.Nats = &v
}

// GetAud returns the Aud field value if set, zero value otherwise.
func (o *UserClaims) GetAud() string {
	if o == nil || IsNil(o.Aud) {
		var ret string
		return ret
	}
	return *o.Aud
}

// GetAudOk returns a tuple with the Aud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserClaims) GetAudOk() (*string, bool) {
	if o == nil || IsNil(o.Aud) {
		return nil, false
	}
	return o.Aud, true
}

// HasAud returns a boolean if a field has been set.
func (o *UserClaims) HasAud() bool {
	if o != nil && !IsNil(o.Aud) {
		return true
	}

	return false
}

// SetAud gets a reference to the given string and assigns it to the Aud field.
func (o *UserClaims) SetAud(v string) {
	o.Aud = &v
}

// GetExp returns the Exp field value if set, zero value otherwise.
func (o *UserClaims) GetExp() int64 {
	if o == nil || IsNil(o.Exp) {
		var ret int64
		return ret
	}
	return *o.Exp
}

// GetExpOk returns a tuple with the Exp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserClaims) GetExpOk() (*int64, bool) {
	if o == nil || IsNil(o.Exp) {
		return nil, false
	}
	return o.Exp, true
}

// HasExp returns a boolean if a field has been set.
func (o *UserClaims) HasExp() bool {
	if o != nil && !IsNil(o.Exp) {
		return true
	}

	return false
}

// SetExp gets a reference to the given int64 and assigns it to the Exp field.
func (o *UserClaims) SetExp(v int64) {
	o.Exp = &v
}

// GetIat returns the Iat field value if set, zero value otherwise.
func (o *UserClaims) GetIat() int64 {
	if o == nil || IsNil(o.Iat) {
		var ret int64
		return ret
	}
	return *o.Iat
}

// GetIatOk returns a tuple with the Iat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserClaims) GetIatOk() (*int64, bool) {
	if o == nil || IsNil(o.Iat) {
		return nil, false
	}
	return o.Iat, true
}

// HasIat returns a boolean if a field has been set.
func (o *UserClaims) HasIat() bool {
	if o != nil && !IsNil(o.Iat) {
		return true
	}

	return false
}

// SetIat gets a reference to the given int64 and assigns it to the Iat field.
func (o *UserClaims) SetIat(v int64) {
	o.Iat = &v
}

// GetIss returns the Iss field value if set, zero value otherwise.
func (o *UserClaims) GetIss() string {
	if o == nil || IsNil(o.Iss) {
		var ret string
		return ret
	}
	return *o.Iss
}

// GetIssOk returns a tuple with the Iss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserClaims) GetIssOk() (*string, bool) {
	if o == nil || IsNil(o.Iss) {
		return nil, false
	}
	return o.Iss, true
}

// HasIss returns a boolean if a field has been set.
func (o *UserClaims) HasIss() bool {
	if o != nil && !IsNil(o.Iss) {
		return true
	}

	return false
}

// SetIss gets a reference to the given string and assigns it to the Iss field.
func (o *UserClaims) SetIss(v string) {
	o.Iss = &v
}

// GetJti returns the Jti field value if set, zero value otherwise.
func (o *UserClaims) GetJti() string {
	if o == nil || IsNil(o.Jti) {
		var ret string
		return ret
	}
	return *o.Jti
}

// GetJtiOk returns a tuple with the Jti field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserClaims) GetJtiOk() (*string, bool) {
	if o == nil || IsNil(o.Jti) {
		return nil, false
	}
	return o.Jti, true
}

// HasJti returns a boolean if a field has been set.
func (o *UserClaims) HasJti() bool {
	if o != nil && !IsNil(o.Jti) {
		return true
	}

	return false
}

// SetJti gets a reference to the given string and assigns it to the Jti field.
func (o *UserClaims) SetJti(v string) {
	o.Jti = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserClaims) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserClaims) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserClaims) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserClaims) SetName(v string) {
	o.Name = &v
}

// GetNbf returns the Nbf field value if set, zero value otherwise.
func (o *UserClaims) GetNbf() int64 {
	if o == nil || IsNil(o.Nbf) {
		var ret int64
		return ret
	}
	return *o.Nbf
}

// GetNbfOk returns a tuple with the Nbf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserClaims) GetNbfOk() (*int64, bool) {
	if o == nil || IsNil(o.Nbf) {
		return nil, false
	}
	return o.Nbf, true
}

// HasNbf returns a boolean if a field has been set.
func (o *UserClaims) HasNbf() bool {
	if o != nil && !IsNil(o.Nbf) {
		return true
	}

	return false
}

// SetNbf gets a reference to the given int64 and assigns it to the Nbf field.
func (o *UserClaims) SetNbf(v int64) {
	o.Nbf = &v
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *UserClaims) GetSub() string {
	if o == nil || IsNil(o.Sub) {
		var ret string
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserClaims) GetSubOk() (*string, bool) {
	if o == nil || IsNil(o.Sub) {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *UserClaims) HasSub() bool {
	if o != nil && !IsNil(o.Sub) {
		return true
	}

	return false
}

// SetSub gets a reference to the given string and assigns it to the Sub field.
func (o *UserClaims) SetSub(v string) {
	o.Sub = &v
}

func (o UserClaims) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserClaims) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nats) {
		toSerialize["nats"] = o.Nats
	}
	if !IsNil(o.Aud) {
		toSerialize["aud"] = o.Aud
	}
	if !IsNil(o.Exp) {
		toSerialize["exp"] = o.Exp
	}
	if !IsNil(o.Iat) {
		toSerialize["iat"] = o.Iat
	}
	if !IsNil(o.Iss) {
		toSerialize["iss"] = o.Iss
	}
	if !IsNil(o.Jti) {
		toSerialize["jti"] = o.Jti
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Nbf) {
		toSerialize["nbf"] = o.Nbf
	}
	if !IsNil(o.Sub) {
		toSerialize["sub"] = o.Sub
	}
	return toSerialize, nil
}

type NullableUserClaims struct {
	value *UserClaims
	isSet bool
}

func (v NullableUserClaims) Get() *UserClaims {
	return v.value
}

func (v *NullableUserClaims) Set(val *UserClaims) {
	v.value = val
	v.isSet = true
}

func (v NullableUserClaims) IsSet() bool {
	return v.isSet
}

func (v *NullableUserClaims) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserClaims(val *UserClaims) *NullableUserClaims {
	return &NullableUserClaims{value: val, isSet: true}
}

func (v NullableUserClaims) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserClaims) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
