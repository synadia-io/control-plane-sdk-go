/*
Synadia Control Plane

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the AppUserAssignResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserAssignResponse{}

// AppUserAssignResponse struct for AppUserAssignResponse
type AppUserAssignResponse struct {
	Account    *AccountInfo  `json:"account,omitempty"`
	AppUser    AppUserInfo   `json:"app_user"`
	Created    time.Time     `json:"created"`
	NatsUser   *NatsUserInfo `json:"nats_user,omitempty"`
	ResourceId string        `json:"resource_id"`
	RoleId     string        `json:"role_id"`
	RoleName   string        `json:"role_name"`
	Scope      AppRoleScope  `json:"scope"`
	System     *SystemInfo   `json:"system,omitempty"`
	Team       *TeamInfo     `json:"team,omitempty"`
}

// NewAppUserAssignResponse instantiates a new AppUserAssignResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUserAssignResponse(appUser AppUserInfo, created time.Time, resourceId string, roleId string, roleName string, scope AppRoleScope) *AppUserAssignResponse {
	this := AppUserAssignResponse{}
	this.AppUser = appUser
	this.Created = created
	this.ResourceId = resourceId
	this.RoleId = roleId
	this.RoleName = roleName
	this.Scope = scope
	return &this
}

// NewAppUserAssignResponseWithDefaults instantiates a new AppUserAssignResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUserAssignResponseWithDefaults() *AppUserAssignResponse {
	this := AppUserAssignResponse{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AppUserAssignResponse) GetAccount() AccountInfo {
	if o == nil || IsNil(o.Account) {
		var ret AccountInfo
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserAssignResponse) GetAccountOk() (*AccountInfo, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AppUserAssignResponse) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given AccountInfo and assigns it to the Account field.
func (o *AppUserAssignResponse) SetAccount(v AccountInfo) {
	o.Account = &v
}

// GetAppUser returns the AppUser field value
func (o *AppUserAssignResponse) GetAppUser() AppUserInfo {
	if o == nil {
		var ret AppUserInfo
		return ret
	}

	return o.AppUser
}

// GetAppUserOk returns a tuple with the AppUser field value
// and a boolean to check if the value has been set.
func (o *AppUserAssignResponse) GetAppUserOk() (*AppUserInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppUser, true
}

// SetAppUser sets field value
func (o *AppUserAssignResponse) SetAppUser(v AppUserInfo) {
	o.AppUser = v
}

// GetCreated returns the Created field value
func (o *AppUserAssignResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *AppUserAssignResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *AppUserAssignResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetNatsUser returns the NatsUser field value if set, zero value otherwise.
func (o *AppUserAssignResponse) GetNatsUser() NatsUserInfo {
	if o == nil || IsNil(o.NatsUser) {
		var ret NatsUserInfo
		return ret
	}
	return *o.NatsUser
}

// GetNatsUserOk returns a tuple with the NatsUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserAssignResponse) GetNatsUserOk() (*NatsUserInfo, bool) {
	if o == nil || IsNil(o.NatsUser) {
		return nil, false
	}
	return o.NatsUser, true
}

// HasNatsUser returns a boolean if a field has been set.
func (o *AppUserAssignResponse) HasNatsUser() bool {
	if o != nil && !IsNil(o.NatsUser) {
		return true
	}

	return false
}

// SetNatsUser gets a reference to the given NatsUserInfo and assigns it to the NatsUser field.
func (o *AppUserAssignResponse) SetNatsUser(v NatsUserInfo) {
	o.NatsUser = &v
}

// GetResourceId returns the ResourceId field value
func (o *AppUserAssignResponse) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *AppUserAssignResponse) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *AppUserAssignResponse) SetResourceId(v string) {
	o.ResourceId = v
}

// GetRoleId returns the RoleId field value
func (o *AppUserAssignResponse) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *AppUserAssignResponse) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *AppUserAssignResponse) SetRoleId(v string) {
	o.RoleId = v
}

// GetRoleName returns the RoleName field value
func (o *AppUserAssignResponse) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *AppUserAssignResponse) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *AppUserAssignResponse) SetRoleName(v string) {
	o.RoleName = v
}

// GetScope returns the Scope field value
func (o *AppUserAssignResponse) GetScope() AppRoleScope {
	if o == nil {
		var ret AppRoleScope
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *AppUserAssignResponse) GetScopeOk() (*AppRoleScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *AppUserAssignResponse) SetScope(v AppRoleScope) {
	o.Scope = v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *AppUserAssignResponse) GetSystem() SystemInfo {
	if o == nil || IsNil(o.System) {
		var ret SystemInfo
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserAssignResponse) GetSystemOk() (*SystemInfo, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *AppUserAssignResponse) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given SystemInfo and assigns it to the System field.
func (o *AppUserAssignResponse) SetSystem(v SystemInfo) {
	o.System = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *AppUserAssignResponse) GetTeam() TeamInfo {
	if o == nil || IsNil(o.Team) {
		var ret TeamInfo
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserAssignResponse) GetTeamOk() (*TeamInfo, bool) {
	if o == nil || IsNil(o.Team) {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *AppUserAssignResponse) HasTeam() bool {
	if o != nil && !IsNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given TeamInfo and assigns it to the Team field.
func (o *AppUserAssignResponse) SetTeam(v TeamInfo) {
	o.Team = &v
}

func (o AppUserAssignResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUserAssignResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	toSerialize["app_user"] = o.AppUser
	toSerialize["created"] = o.Created
	if !IsNil(o.NatsUser) {
		toSerialize["nats_user"] = o.NatsUser
	}
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["role_id"] = o.RoleId
	toSerialize["role_name"] = o.RoleName
	toSerialize["scope"] = o.Scope
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	if !IsNil(o.Team) {
		toSerialize["team"] = o.Team
	}
	return toSerialize, nil
}

type NullableAppUserAssignResponse struct {
	value *AppUserAssignResponse
	isSet bool
}

func (v NullableAppUserAssignResponse) Get() *AppUserAssignResponse {
	return v.value
}

func (v *NullableAppUserAssignResponse) Set(val *AppUserAssignResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUserAssignResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUserAssignResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUserAssignResponse(val *AppUserAssignResponse) *NullableAppUserAssignResponse {
	return &NullableAppUserAssignResponse{value: val, isSet: true}
}

func (v NullableAppUserAssignResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUserAssignResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
