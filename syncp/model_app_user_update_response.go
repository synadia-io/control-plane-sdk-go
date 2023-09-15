/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"encoding/json"
	"time"
)

// checks if the AppUserUpdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserUpdateResponse{}

// AppUserUpdateResponse struct for AppUserUpdateResponse
type AppUserUpdateResponse struct {
	ResetPasswordLink *string        `json:"reset_password_link,omitempty"`
	Created           time.Time      `json:"created"`
	Id                string         `json:"id"`
	Identifier        NullableString `json:"identifier"`
	Name              string         `json:"name"`
	RoleId            string         `json:"role_id"`
	RoleName          string         `json:"role_name"`
	Type              AppUserType    `json:"type"`
}

// NewAppUserUpdateResponse instantiates a new AppUserUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUserUpdateResponse(created time.Time, id string, identifier NullableString, name string, roleId string, roleName string, type_ AppUserType) *AppUserUpdateResponse {
	this := AppUserUpdateResponse{}
	this.Created = created
	this.Id = id
	this.Identifier = identifier
	this.Name = name
	this.RoleId = roleId
	this.RoleName = roleName
	this.Type = type_
	return &this
}

// NewAppUserUpdateResponseWithDefaults instantiates a new AppUserUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUserUpdateResponseWithDefaults() *AppUserUpdateResponse {
	this := AppUserUpdateResponse{}
	return &this
}

// GetResetPasswordLink returns the ResetPasswordLink field value if set, zero value otherwise.
func (o *AppUserUpdateResponse) GetResetPasswordLink() string {
	if o == nil || IsNil(o.ResetPasswordLink) {
		var ret string
		return ret
	}
	return *o.ResetPasswordLink
}

// GetResetPasswordLinkOk returns a tuple with the ResetPasswordLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserUpdateResponse) GetResetPasswordLinkOk() (*string, bool) {
	if o == nil || IsNil(o.ResetPasswordLink) {
		return nil, false
	}
	return o.ResetPasswordLink, true
}

// HasResetPasswordLink returns a boolean if a field has been set.
func (o *AppUserUpdateResponse) HasResetPasswordLink() bool {
	if o != nil && !IsNil(o.ResetPasswordLink) {
		return true
	}

	return false
}

// SetResetPasswordLink gets a reference to the given string and assigns it to the ResetPasswordLink field.
func (o *AppUserUpdateResponse) SetResetPasswordLink(v string) {
	o.ResetPasswordLink = &v
}

// GetCreated returns the Created field value
func (o *AppUserUpdateResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *AppUserUpdateResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *AppUserUpdateResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetId returns the Id field value
func (o *AppUserUpdateResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppUserUpdateResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppUserUpdateResponse) SetId(v string) {
	o.Id = v
}

// GetIdentifier returns the Identifier field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AppUserUpdateResponse) GetIdentifier() string {
	if o == nil || o.Identifier.Get() == nil {
		var ret string
		return ret
	}

	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUserUpdateResponse) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// SetIdentifier sets field value
func (o *AppUserUpdateResponse) SetIdentifier(v string) {
	o.Identifier.Set(&v)
}

// GetName returns the Name field value
func (o *AppUserUpdateResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppUserUpdateResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppUserUpdateResponse) SetName(v string) {
	o.Name = v
}

// GetRoleId returns the RoleId field value
func (o *AppUserUpdateResponse) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *AppUserUpdateResponse) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *AppUserUpdateResponse) SetRoleId(v string) {
	o.RoleId = v
}

// GetRoleName returns the RoleName field value
func (o *AppUserUpdateResponse) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *AppUserUpdateResponse) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *AppUserUpdateResponse) SetRoleName(v string) {
	o.RoleName = v
}

// GetType returns the Type field value
func (o *AppUserUpdateResponse) GetType() AppUserType {
	if o == nil {
		var ret AppUserType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppUserUpdateResponse) GetTypeOk() (*AppUserType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppUserUpdateResponse) SetType(v AppUserType) {
	o.Type = v
}

func (o AppUserUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUserUpdateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResetPasswordLink) {
		toSerialize["reset_password_link"] = o.ResetPasswordLink
	}
	toSerialize["created"] = o.Created
	toSerialize["id"] = o.Id
	toSerialize["identifier"] = o.Identifier.Get()
	toSerialize["name"] = o.Name
	toSerialize["role_id"] = o.RoleId
	toSerialize["role_name"] = o.RoleName
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAppUserUpdateResponse struct {
	value *AppUserUpdateResponse
	isSet bool
}

func (v NullableAppUserUpdateResponse) Get() *AppUserUpdateResponse {
	return v.value
}

func (v *NullableAppUserUpdateResponse) Set(val *AppUserUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUserUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUserUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUserUpdateResponse(val *AppUserUpdateResponse) *NullableAppUserUpdateResponse {
	return &NullableAppUserUpdateResponse{value: val, isSet: true}
}

func (v NullableAppUserUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUserUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
