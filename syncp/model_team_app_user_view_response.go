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

// checks if the TeamAppUserViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamAppUserViewResponse{}

// TeamAppUserViewResponse struct for TeamAppUserViewResponse
type TeamAppUserViewResponse struct {
	AppUser           AppUserInfo `json:"app_user"`
	Created           time.Time   `json:"created"`
	Id                string      `json:"id"`
	PendingInvitation bool        `json:"pending_invitation"`
	RoleId            string      `json:"role_id"`
	RoleName          string      `json:"role_name"`
	Team              TeamInfo    `json:"team"`
}

type _TeamAppUserViewResponse TeamAppUserViewResponse

// NewTeamAppUserViewResponse instantiates a new TeamAppUserViewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamAppUserViewResponse(appUser AppUserInfo, created time.Time, id string, pendingInvitation bool, roleId string, roleName string, team TeamInfo) *TeamAppUserViewResponse {
	this := TeamAppUserViewResponse{}
	this.AppUser = appUser
	this.Created = created
	this.Id = id
	this.PendingInvitation = pendingInvitation
	this.RoleId = roleId
	this.RoleName = roleName
	this.Team = team
	return &this
}

// NewTeamAppUserViewResponseWithDefaults instantiates a new TeamAppUserViewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamAppUserViewResponseWithDefaults() *TeamAppUserViewResponse {
	this := TeamAppUserViewResponse{}
	return &this
}

// GetAppUser returns the AppUser field value
func (o *TeamAppUserViewResponse) GetAppUser() AppUserInfo {
	if o == nil {
		var ret AppUserInfo
		return ret
	}

	return o.AppUser
}

// GetAppUserOk returns a tuple with the AppUser field value
// and a boolean to check if the value has been set.
func (o *TeamAppUserViewResponse) GetAppUserOk() (*AppUserInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppUser, true
}

// SetAppUser sets field value
func (o *TeamAppUserViewResponse) SetAppUser(v AppUserInfo) {
	o.AppUser = v
}

// GetCreated returns the Created field value
func (o *TeamAppUserViewResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TeamAppUserViewResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TeamAppUserViewResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetId returns the Id field value
func (o *TeamAppUserViewResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeamAppUserViewResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TeamAppUserViewResponse) SetId(v string) {
	o.Id = v
}

// GetPendingInvitation returns the PendingInvitation field value
func (o *TeamAppUserViewResponse) GetPendingInvitation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PendingInvitation
}

// GetPendingInvitationOk returns a tuple with the PendingInvitation field value
// and a boolean to check if the value has been set.
func (o *TeamAppUserViewResponse) GetPendingInvitationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingInvitation, true
}

// SetPendingInvitation sets field value
func (o *TeamAppUserViewResponse) SetPendingInvitation(v bool) {
	o.PendingInvitation = v
}

// GetRoleId returns the RoleId field value
func (o *TeamAppUserViewResponse) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *TeamAppUserViewResponse) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *TeamAppUserViewResponse) SetRoleId(v string) {
	o.RoleId = v
}

// GetRoleName returns the RoleName field value
func (o *TeamAppUserViewResponse) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *TeamAppUserViewResponse) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *TeamAppUserViewResponse) SetRoleName(v string) {
	o.RoleName = v
}

// GetTeam returns the Team field value
func (o *TeamAppUserViewResponse) GetTeam() TeamInfo {
	if o == nil {
		var ret TeamInfo
		return ret
	}

	return o.Team
}

// GetTeamOk returns a tuple with the Team field value
// and a boolean to check if the value has been set.
func (o *TeamAppUserViewResponse) GetTeamOk() (*TeamInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Team, true
}

// SetTeam sets field value
func (o *TeamAppUserViewResponse) SetTeam(v TeamInfo) {
	o.Team = v
}

func (o TeamAppUserViewResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamAppUserViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app_user"] = o.AppUser
	toSerialize["created"] = o.Created
	toSerialize["id"] = o.Id
	toSerialize["pending_invitation"] = o.PendingInvitation
	toSerialize["role_id"] = o.RoleId
	toSerialize["role_name"] = o.RoleName
	toSerialize["team"] = o.Team
	return toSerialize, nil
}

func (o *TeamAppUserViewResponse) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"app_user",
		"created",
		"id",
		"pending_invitation",
		"role_id",
		"role_name",
		"team",
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

	varTeamAppUserViewResponse := _TeamAppUserViewResponse{}

	err = json.Unmarshal(bytes, &varTeamAppUserViewResponse)

	if err != nil {
		return err
	}

	*o = TeamAppUserViewResponse(varTeamAppUserViewResponse)

	return err
}

type NullableTeamAppUserViewResponse struct {
	value *TeamAppUserViewResponse
	isSet bool
}

func (v NullableTeamAppUserViewResponse) Get() *TeamAppUserViewResponse {
	return v.value
}

func (v *NullableTeamAppUserViewResponse) Set(val *TeamAppUserViewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamAppUserViewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamAppUserViewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamAppUserViewResponse(val *TeamAppUserViewResponse) *NullableTeamAppUserViewResponse {
	return &NullableTeamAppUserViewResponse{value: val, isSet: true}
}

func (v NullableTeamAppUserViewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamAppUserViewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}