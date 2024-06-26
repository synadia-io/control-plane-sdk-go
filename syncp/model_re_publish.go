/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the RePublish type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RePublish{}

// RePublish RePublish allows a source subject to be mapped to a destination subject for republishing.
type RePublish struct {
	Dest        string  `json:"dest"`
	HeadersOnly *bool   `json:"headers_only,omitempty"`
	Src         *string `json:"src,omitempty"`
}

func (o RePublish) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dest"] = o.Dest
	if o.HeadersOnly != nil {
		toSerialize["headers_only"] = o.HeadersOnly
	}
	if o.Src != nil {
		toSerialize["src"] = o.Src
	}
	return toSerialize, nil
}
