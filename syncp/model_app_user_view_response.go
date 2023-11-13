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

// checks if the AppUserViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserViewResponse{}

// AppUserViewResponse struct for AppUserViewResponse
type AppUserViewResponse struct {
	Created              time.Time      `json:"created"`
	Id                   string         `json:"id"`
	Identifier           NullableString `json:"identifier"`
	Name                 string         `json:"name"`
	OryId                NullableString `json:"ory_id"`
	RoleId               string         `json:"role_id"`
	RoleName             string         `json:"role_name"`
	TermsAcceptedUpdated NullableString `json:"terms_accepted_updated"`
	Type                 AppUserType    `json:"type"`
}

type _AppUserViewResponse AppUserViewResponse

// NewAppUserViewResponse instantiates a new AppUserViewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUserViewResponse(created time.Time, id string, identifier NullableString, name string, oryId NullableString, roleId string, roleName string, termsAcceptedUpdated NullableString, type_ AppUserType) *AppUserViewResponse {
	this := AppUserViewResponse{}
	this.Created = created
	this.Id = id
	this.Identifier = identifier
	this.Name = name
	this.OryId = oryId
	this.RoleId = roleId
	this.RoleName = roleName
	this.TermsAcceptedUpdated = termsAcceptedUpdated
	this.Type = type_
	return &this
}

// NewAppUserViewResponseWithDefaults instantiates a new AppUserViewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUserViewResponseWithDefaults() *AppUserViewResponse {
	this := AppUserViewResponse{}
	return &this
}

// GetCreated returns the Created field value
func (o *AppUserViewResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *AppUserViewResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *AppUserViewResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetId returns the Id field value
func (o *AppUserViewResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppUserViewResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppUserViewResponse) SetId(v string) {
	o.Id = v
}

// GetIdentifier returns the Identifier field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AppUserViewResponse) GetIdentifier() string {
	if o == nil || o.Identifier.Get() == nil {
		var ret string
		return ret
	}

	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUserViewResponse) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// SetIdentifier sets field value
func (o *AppUserViewResponse) SetIdentifier(v string) {
	o.Identifier.Set(&v)
}

// GetName returns the Name field value
func (o *AppUserViewResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppUserViewResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppUserViewResponse) SetName(v string) {
	o.Name = v
}

// GetOryId returns the OryId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AppUserViewResponse) GetOryId() string {
	if o == nil || o.OryId.Get() == nil {
		var ret string
		return ret
	}

	return *o.OryId.Get()
}

// GetOryIdOk returns a tuple with the OryId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUserViewResponse) GetOryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OryId.Get(), o.OryId.IsSet()
}

// SetOryId sets field value
func (o *AppUserViewResponse) SetOryId(v string) {
	o.OryId.Set(&v)
}

// GetRoleId returns the RoleId field value
func (o *AppUserViewResponse) GetRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value
// and a boolean to check if the value has been set.
func (o *AppUserViewResponse) GetRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleId, true
}

// SetRoleId sets field value
func (o *AppUserViewResponse) SetRoleId(v string) {
	o.RoleId = v
}

// GetRoleName returns the RoleName field value
func (o *AppUserViewResponse) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *AppUserViewResponse) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *AppUserViewResponse) SetRoleName(v string) {
	o.RoleName = v
}

// GetTermsAcceptedUpdated returns the TermsAcceptedUpdated field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AppUserViewResponse) GetTermsAcceptedUpdated() string {
	if o == nil || o.TermsAcceptedUpdated.Get() == nil {
		var ret string
		return ret
	}

	return *o.TermsAcceptedUpdated.Get()
}

// GetTermsAcceptedUpdatedOk returns a tuple with the TermsAcceptedUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUserViewResponse) GetTermsAcceptedUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TermsAcceptedUpdated.Get(), o.TermsAcceptedUpdated.IsSet()
}

// SetTermsAcceptedUpdated sets field value
func (o *AppUserViewResponse) SetTermsAcceptedUpdated(v string) {
	o.TermsAcceptedUpdated.Set(&v)
}

// GetType returns the Type field value
func (o *AppUserViewResponse) GetType() AppUserType {
	if o == nil {
		var ret AppUserType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppUserViewResponse) GetTypeOk() (*AppUserType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppUserViewResponse) SetType(v AppUserType) {
	o.Type = v
}

func (o AppUserViewResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUserViewResponse) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

func (o *AppUserViewResponse) UnmarshalJSON(bytes []byte) (err error) {
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

	varAppUserViewResponse := _AppUserViewResponse{}

	err = json.Unmarshal(bytes, &varAppUserViewResponse)

	if err != nil {
		return err
	}

	*o = AppUserViewResponse(varAppUserViewResponse)

	return err
}

type NullableAppUserViewResponse struct {
	value *AppUserViewResponse
	isSet bool
}

func (v NullableAppUserViewResponse) Get() *AppUserViewResponse {
	return v.value
}

func (v *NullableAppUserViewResponse) Set(val *AppUserViewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUserViewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUserViewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUserViewResponse(val *AppUserViewResponse) *NullableAppUserViewResponse {
	return &NullableAppUserViewResponse{value: val, isSet: true}
}

func (v NullableAppUserViewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUserViewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
