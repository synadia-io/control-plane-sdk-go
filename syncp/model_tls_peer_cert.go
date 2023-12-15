/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the TLSPeerCert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TLSPeerCert{}

// TLSPeerCert TLSPeerCert contains basic information about a TLS peer certificate
type TLSPeerCert struct {
	CertSha256 *string `json:"cert_sha256,omitempty"`
	SpkiSha256 *string `json:"spki_sha256,omitempty"`
	Subject    *string `json:"subject,omitempty"`
}

func (o TLSPeerCert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CertSha256 != nil {
		toSerialize["cert_sha256"] = o.CertSha256
	}
	if o.SpkiSha256 != nil {
		toSerialize["spki_sha256"] = o.SpkiSha256
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	return toSerialize, nil
}
