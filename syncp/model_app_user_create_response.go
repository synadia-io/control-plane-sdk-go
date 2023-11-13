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

// checks if the AppUserCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserCreateResponse{}

// AppUserCreateResponse struct for AppUserCreateResponse
type AppUserCreateResponse struct {
	Created              time.Time      `json:"created"`
	Id                   string         `json:"id"`
	Identifier           NullableString `json:"identifier"`
	Name                 string         `json:"name"`
	OryId                NullableString `json:"ory_id"`
	RoleId               string         `json:"role_id"`
	RoleName             string         `json:"role_name"`
	TermsAcceptedUpdated NullableString `json:"terms_accepted_updated"`
	Type                 AppUserType    `json:"type"`
	InviteLink           string         `json:"invite_link"`
}

type _AppUserCreateResponse AppUserCreateResponse

// NewAppUserCreateResponse instantiates a new AppUserCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUserCreateResponse(created time.Time, id string, identifier NullableString, name string, oryId NullableString, roleId string, roleName string, termsAcceptedUpdated NullableString, type_ AppUserType, inviteLink string) *AppUserCreateResponse {
	this := AppUserCreateResponse{}
	this.Created = created
	this.Id = id
	this.Identifier = identifier
	this.Name = name
	this.OryId = oryId
	this.RoleId = roleId
	this.RoleName = roleName
	this.TermsAcceptedUpdated = termsAcceptedUpdated
	this.Type = type_
	this.InviteLink = inviteLink
	return &this
}

// NewAppUserCreateResponseWithDefaults instantiates a new AppUserCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUserCreateResponseWithDefaults() *AppUserCreateResponse {
	this := AppUserCreateResponse{}
	return &this
}

// GetCreated returns the Created field value
func (o *AppUserCreateResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *AppUserCreateResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *AppUserCreateResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetId returns the Id field value
func (o *AppUserCreateResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppUserCreateResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppUserCreateResponse) SetId(v string) {
	o.Id = v
}

// GetIdentifier returns the Identifier field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AppUserCreateResponse) GetIdentifier() string {
	if o == nil || o.Identifier.Get() == nil {
		var ret string
		return ret
	}

	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUserCreateResponse) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// SetIdentifier sets field value
func (o *AppUserCreateResponse) SetIdentifier(v string) {
	o.Identifier.Set(&v)
}

// GetName returns the Name field value
func (o *AppUserCreateResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppUserCreateResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppUserCreateResponse) SetName(v string) {
	o.Name = v
}

// GetOryId returns the OryId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AppUserCreateResponse) GetOryId() string {
	if o == nil || o.OryId.Get() == nil {
		var ret string
		return ret
	}

	return *o.OryId.Get()
}

// GetOryIdOk returns a tuple with the OryId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUserCreateResponse) GetOryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OryId.Get(), o.OryId.IsSet()
}

// SetOryId sets field value
func (o *AppUserCreateResponse) SetOryId(v string) {
	o.OryId.Set(&v)
}

// GetRoleId returns the RoleId field value
func (o *AppUserCreateResponse) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *AppUserCreateResponse) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *AppUserCreateResponse) SetRoleId(v string) {
	o.RoleId = v
}

// GetRoleName returns the RoleName field value
func (o *AppUserCreateResponse) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *AppUserCreateResponse) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *AppUserCreateResponse) SetRoleName(v string) {
	o.RoleName = v
}

// GetTermsAcceptedUpdated returns the TermsAcceptedUpdated field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AppUserCreateResponse) GetTermsAcceptedUpdated() string {
	if o == nil || o.TermsAcceptedUpdated.Get() == nil {
		var ret string
		return ret
	}

	return *o.TermsAcceptedUpdated.Get()
}

// GetTermsAcceptedUpdatedOk returns a tuple with the TermsAcceptedUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUserCreateResponse) GetTermsAcceptedUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TermsAcceptedUpdated.Get(), o.TermsAcceptedUpdated.IsSet()
}

// SetTermsAcceptedUpdated sets field value
func (o *AppUserCreateResponse) SetTermsAcceptedUpdated(v string) {
	o.TermsAcceptedUpdated.Set(&v)
}

// GetType returns the Type field value
func (o *AppUserCreateResponse) GetType() AppUserType {
	if o == nil {
		var ret AppUserType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppUserCreateResponse) GetTypeOk() (*AppUserType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppUserCreateResponse) SetType(v AppUserType) {
	o.Type = v
}

// GetInviteLink returns the InviteLink field value
func (o *AppUserCreateResponse) GetInviteLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InviteLink
}

// GetInviteLinkOk returns a tuple with the InviteLink field value
// and a boolean to check if the value has been set.
func (o *AppUserCreateResponse) GetInviteLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InviteLink, true
}

// SetInviteLink sets field value
func (o *AppUserCreateResponse) SetInviteLink(v string) {
	o.InviteLink = v
}

func (o AppUserCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUserCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["id"] = o.Id
	toSerialize["identifier"] = o.Identifier.Get()
	toSerialize["name"] = o.Name
	toSerialize["ory_id"] = o.OryId.Get()
	toSerialize["role_id"] = o.RoleId
	toSerialize["role_name"] = o.RoleName
	toSerialize["terms_accepted_updated"] = o.TermsAcceptedUpdated.Get()
	toSerialize["type"] = o.Type
	toSerialize["invite_link"] = o.InviteLink
	return toSerialize, nil
}

func (o *AppUserCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created",
		"id",
		"identifier",
		"name",
		"ory_id",
		"role_id",
		"role_name",
		"terms_accepted_updated",
		"type",
		"invite_link",
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

	varAppUserCreateResponse := _AppUserCreateResponse{}

	err = json.Unmarshal(bytes, &varAppUserCreateResponse)

	if err != nil {
		return err
	}

	*o = AppUserCreateResponse(varAppUserCreateResponse)

	return err
}

type NullableAppUserCreateResponse struct {
	value *AppUserCreateResponse
	isSet bool
}

func (v NullableAppUserCreateResponse) Get() *AppUserCreateResponse {
	return v.value
}

func (v *NullableAppUserCreateResponse) Set(val *AppUserCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUserCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUserCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUserCreateResponse(val *AppUserCreateResponse) *NullableAppUserCreateResponse {
	return &NullableAppUserCreateResponse{value: val, isSet: true}
}

func (v NullableAppUserCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUserCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
