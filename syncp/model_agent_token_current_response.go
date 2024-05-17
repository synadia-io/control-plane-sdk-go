/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the AgentTokenCurrentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentTokenCurrentResponse{}

// AgentTokenCurrentResponse struct for AgentTokenCurrentResponse
type AgentTokenCurrentResponse struct {
	AgentTokenViewResponse
	Token string `json:"token"`
	Url   string `json:"url"`
}

func (o AgentTokenCurrentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	toSerialize["url"] = o.Url
	toSerialize["created"] = o.Created
	toSerialize["id"] = o.Id
	toSerialize["is_current"] = o.IsCurrent
	if o.LastAccessedAt != nil {
		toSerialize["last_accessed_at"] = o.LastAccessedAt
	}
	toSerialize["nkey_public"] = o.NkeyPublic
	if o.RotatedAt != nil {
		toSerialize["rotated_at"] = o.RotatedAt
	}
	return toSerialize, nil
}
