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
	"time"
)

// checks if the AppUserAccessTokenCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserAccessTokenCreateResponse{}

// AppUserAccessTokenCreateResponse struct for AppUserAccessTokenCreateResponse
type AppUserAccessTokenCreateResponse struct {
	Created time.Time      `json:"created"`
	Expires NullableString `json:"expires"`
	Id      string         `json:"id"`
	Name    string         `json:"name"`
	Token   string         `json:"token"`
}

type _AppUserAccessTokenCreateResponse AppUserAccessTokenCreateResponse

// NewAppUserAccessTokenCreateResponse instantiates a new AppUserAccessTokenCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUserAccessTokenCreateResponse(created time.Time, expires NullableString, id string, name string, token string) *AppUserAccessTokenCreateResponse {
	this := AppUserAccessTokenCreateResponse{}
	this.Created = created
	this.Expires = expires
	this.Id = id
	this.Name = name
	this.Token = token
	return &this
}

// NewAppUserAccessTokenCreateResponseWithDefaults instantiates a new AppUserAccessTokenCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUserAccessTokenCreateResponseWithDefaults() *AppUserAccessTokenCreateResponse {
	this := AppUserAccessTokenCreateResponse{}
	return &this
}

// GetCreated returns the Created field value
func (o *AppUserAccessTokenCreateResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *AppUserAccessTokenCreateResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *AppUserAccessTokenCreateResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetExpires returns the Expires field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AppUserAccessTokenCreateResponse) GetExpires() string {
	if o == nil || o.Expires.Get() == nil {
		var ret string
		return ret
	}

	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUserAccessTokenCreateResponse) GetExpiresOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// SetExpires sets field value
func (o *AppUserAccessTokenCreateResponse) SetExpires(v string) {
	o.Expires.Set(&v)
}

// GetId returns the Id field value
func (o *AppUserAccessTokenCreateResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppUserAccessTokenCreateResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppUserAccessTokenCreateResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AppUserAccessTokenCreateResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppUserAccessTokenCreateResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppUserAccessTokenCreateResponse) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value
func (o *AppUserAccessTokenCreateResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *AppUserAccessTokenCreateResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *AppUserAccessTokenCreateResponse) SetToken(v string) {
	o.Token = v
}

func (o AppUserAccessTokenCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUserAccessTokenCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["expires"] = o.Expires.Get()
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *AppUserAccessTokenCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created",
		"expires",
		"id",
		"name",
		"token",
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

	varAppUserAccessTokenCreateResponse := _AppUserAccessTokenCreateResponse{}

	err = json.Unmarshal(bytes, &varAppUserAccessTokenCreateResponse)

	if err != nil {
		return err
	}

	*o = AppUserAccessTokenCreateResponse(varAppUserAccessTokenCreateResponse)

	return err
}

type NullableAppUserAccessTokenCreateResponse struct {
	value *AppUserAccessTokenCreateResponse
	isSet bool
}

func (v NullableAppUserAccessTokenCreateResponse) Get() *AppUserAccessTokenCreateResponse {
	return v.value
}

func (v *NullableAppUserAccessTokenCreateResponse) Set(val *AppUserAccessTokenCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUserAccessTokenCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUserAccessTokenCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUserAccessTokenCreateResponse(val *AppUserAccessTokenCreateResponse) *NullableAppUserAccessTokenCreateResponse {
	return &NullableAppUserAccessTokenCreateResponse{value: val, isSet: true}
}

func (v NullableAppUserAccessTokenCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUserAccessTokenCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
