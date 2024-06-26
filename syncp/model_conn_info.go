/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

import (
	"time"
)

// checks if the ConnInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnInfo{}

// ConnInfo ConnInfo has detailed information on a per connection basis.
type ConnInfo struct {
	Account                 *string           `json:"account,omitempty"`
	AuthorizedUser          *string           `json:"authorized_user,omitempty"`
	Cid                     uint64            `json:"cid"`
	Idle                    string            `json:"idle"`
	InBytes                 int64             `json:"in_bytes"`
	InMsgs                  int64             `json:"in_msgs"`
	Ip                      string            `json:"ip"`
	IssuerKey               *string           `json:"issuer_key,omitempty"`
	Jwt                     *string           `json:"jwt,omitempty"`
	Kind                    *string           `json:"kind,omitempty"`
	Lang                    *string           `json:"lang,omitempty"`
	LastActivity            time.Time         `json:"last_activity"`
	MqttClient              *string           `json:"mqtt_client,omitempty"`
	Name                    *string           `json:"name,omitempty"`
	NameTag                 *string           `json:"name_tag,omitempty"`
	OutBytes                int64             `json:"out_bytes"`
	OutMsgs                 int64             `json:"out_msgs"`
	PendingBytes            int64             `json:"pending_bytes"`
	Port                    int64             `json:"port"`
	Reason                  *string           `json:"reason,omitempty"`
	Rtt                     *string           `json:"rtt,omitempty"`
	Start                   time.Time         `json:"start"`
	Stop                    *Nullable[string] `json:"stop,omitempty"`
	Subscriptions           uint32            `json:"subscriptions"`
	SubscriptionsList       []string          `json:"subscriptions_list,omitempty"`
	SubscriptionsListDetail []SubDetail       `json:"subscriptions_list_detail,omitempty"`
	Tags                    []string          `json:"tags,omitempty"`
	TlsCipherSuite          *string           `json:"tls_cipher_suite,omitempty"`
	TlsPeerCerts            []TLSPeerCert     `json:"tls_peer_certs,omitempty"`
	TlsVersion              *string           `json:"tls_version,omitempty"`
	Type                    *string           `json:"type,omitempty"`
	Uptime                  string            `json:"uptime"`
	Version                 *string           `json:"version,omitempty"`
}

func (o ConnInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.AuthorizedUser != nil {
		toSerialize["authorized_user"] = o.AuthorizedUser
	}
	toSerialize["cid"] = o.Cid
	toSerialize["idle"] = o.Idle
	toSerialize["in_bytes"] = o.InBytes
	toSerialize["in_msgs"] = o.InMsgs
	toSerialize["ip"] = o.Ip
	if o.IssuerKey != nil {
		toSerialize["issuer_key"] = o.IssuerKey
	}
	if o.Jwt != nil {
		toSerialize["jwt"] = o.Jwt
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Lang != nil {
		toSerialize["lang"] = o.Lang
	}
	toSerialize["last_activity"] = o.LastActivity
	if o.MqttClient != nil {
		toSerialize["mqtt_client"] = o.MqttClient
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NameTag != nil {
		toSerialize["name_tag"] = o.NameTag
	}
	toSerialize["out_bytes"] = o.OutBytes
	toSerialize["out_msgs"] = o.OutMsgs
	toSerialize["pending_bytes"] = o.PendingBytes
	toSerialize["port"] = o.Port
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Rtt != nil {
		toSerialize["rtt"] = o.Rtt
	}
	toSerialize["start"] = o.Start
	if o.Stop != nil && !o.Stop.IsNull() {
		toSerialize["stop"] = o.Stop.Val
	}
	toSerialize["subscriptions"] = o.Subscriptions
	if len(o.SubscriptionsList) != 0 {
		toSerialize["subscriptions_list"] = o.SubscriptionsList
	}
	if len(o.SubscriptionsListDetail) != 0 {
		toSerialize["subscriptions_list_detail"] = o.SubscriptionsListDetail
	}
	if len(o.Tags) != 0 {
		toSerialize["tags"] = o.Tags
	}
	if o.TlsCipherSuite != nil {
		toSerialize["tls_cipher_suite"] = o.TlsCipherSuite
	}
	if len(o.TlsPeerCerts) != 0 {
		toSerialize["tls_peer_certs"] = o.TlsPeerCerts
	}
	if o.TlsVersion != nil {
		toSerialize["tls_version"] = o.TlsVersion
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["uptime"] = o.Uptime
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}
