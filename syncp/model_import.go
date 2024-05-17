/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the Import type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Import{}

// Import Import describes a mapping from another account into this one
type Import struct {
	Account *string `json:"account,omitempty"`
	// Local subject used to subscribe (for streams) and publish (for services) to. This value only needs setting if you want to change the value of Subject. If the value of Subject ends in > then LocalSubject needs to end in > as well. LocalSubject can contain $<number> wildcard references where number references the nth wildcard in Subject. The sum of wildcard reference and * tokens needs to match the number of * token in Subject.
	LocalSubject *string `json:"local_subject,omitempty"`
	Name         *string `json:"name,omitempty"`
	Share        *bool   `json:"share,omitempty"`
	// Subject field in an import is always from the perspective of the initial publisher - in the case of a stream it is the account owning the stream (the exporter), and in the case of a service it is the account making the request (the importer).
	Subject *string `json:"subject,omitempty"`
	// Deprecated: use LocalSubject instead To field in an import is always from the perspective of the subscriber in the case of a stream it is the client of the stream (the importer), from the perspective of a service, it is the subscription waiting for requests (the exporter). If the field is empty, it will default to the value in the Subject field.
	To    *string     `json:"to,omitempty"`
	Token *string     `json:"token,omitempty"`
	Type  *ExportType `json:"type,omitempty"`
}

func (o Import) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.LocalSubject != nil {
		toSerialize["local_subject"] = o.LocalSubject
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Share != nil {
		toSerialize["share"] = o.Share
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}
