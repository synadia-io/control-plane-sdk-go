/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the Permissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Permissions{}

// Permissions Permissions are used to restrict subject access, either on a user or for everyone on a server by default
type Permissions struct {
	Pub  *Permission                   `json:"pub,omitempty"`
	Resp *Nullable[ResponsePermission] `json:"resp,omitempty"`
	Sub  *Permission                   `json:"sub,omitempty"`
}

func (o Permissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Pub != nil {
		toSerialize["pub"] = o.Pub
	}
	if o.Resp != nil && !o.Resp.IsNull() {
		toSerialize["resp"] = o.Resp.Val
	}
	if o.Sub != nil {
		toSerialize["sub"] = o.Sub
	}
	return toSerialize, nil
}
