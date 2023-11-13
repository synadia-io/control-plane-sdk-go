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

// checks if the StreamExportViewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamExportViewResponse{}

// StreamExportViewResponse struct for StreamExportViewResponse
type StreamExportViewResponse struct {
	AccountId            string    `json:"account_id"`
	Created              time.Time `json:"created"`
	DeliverSubjectPrefix string    `json:"deliver_subject_prefix"`
	Id                   string    `json:"id"`
	IsPublic             bool      `json:"is_public"`
	JsSubjectPrefix      string    `json:"js_subject_prefix"`
	StreamName           string    `json:"stream_name"`
}

type _StreamExportViewResponse StreamExportViewResponse

// NewStreamExportViewResponse instantiates a new StreamExportViewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamExportViewResponse(accountId string, created time.Time, deliverSubjectPrefix string, id string, isPublic bool, jsSubjectPrefix string, streamName string) *StreamExportViewResponse {
	this := StreamExportViewResponse{}
	this.AccountId = accountId
	this.Created = created
	this.DeliverSubjectPrefix = deliverSubjectPrefix
	this.Id = id
	this.IsPublic = isPublic
	this.JsSubjectPrefix = jsSubjectPrefix
	this.StreamName = streamName
	return &this
}

// NewStreamExportViewResponseWithDefaults instantiates a new StreamExportViewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamExportViewResponseWithDefaults() *StreamExportViewResponse {
	this := StreamExportViewResponse{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *StreamExportViewResponse) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *StreamExportViewResponse) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *StreamExportViewResponse) SetAccountId(v string) {
	o.AccountId = v
}

// GetCreated returns the Created field value
func (o *StreamExportViewResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *StreamExportViewResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *StreamExportViewResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetDeliverSubjectPrefix returns the DeliverSubjectPrefix field value
func (o *StreamExportViewResponse) GetDeliverSubjectPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliverSubjectPrefix
}

// GetDeliverSubjectPrefixOk returns a tuple with the DeliverSubjectPrefix field value
// and a boolean to check if the value has been set.
func (o *StreamExportViewResponse) GetDeliverSubjectPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliverSubjectPrefix, true
}

// SetDeliverSubjectPrefix sets field value
func (o *StreamExportViewResponse) SetDeliverSubjectPrefix(v string) {
	o.DeliverSubjectPrefix = v
}

// GetId returns the Id field value
func (o *StreamExportViewResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StreamExportViewResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StreamExportViewResponse) SetId(v string) {
	o.Id = v
}

// GetIsPublic returns the IsPublic field value
func (o *StreamExportViewResponse) GetIsPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value
// and a boolean to check if the value has been set.
func (o *StreamExportViewResponse) GetIsPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublic, true
}

// SetIsPublic sets field value
func (o *StreamExportViewResponse) SetIsPublic(v bool) {
	o.IsPublic = v
}

// GetJsSubjectPrefix returns the JsSubjectPrefix field value
func (o *StreamExportViewResponse) GetJsSubjectPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsSubjectPrefix
}

// GetJsSubjectPrefixOk returns a tuple with the JsSubjectPrefix field value
// and a boolean to check if the value has been set.
func (o *StreamExportViewResponse) GetJsSubjectPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsSubjectPrefix, true
}

// SetJsSubjectPrefix sets field value
func (o *StreamExportViewResponse) SetJsSubjectPrefix(v string) {
	o.JsSubjectPrefix = v
}

// GetStreamName returns the StreamName field value
func (o *StreamExportViewResponse) GetStreamName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StreamName
}

// GetStreamNameOk returns a tuple with the StreamName field value
// and a boolean to check if the value has been set.
func (o *StreamExportViewResponse) GetStreamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreamName, true
}

// SetStreamName sets field value
func (o *StreamExportViewResponse) SetStreamName(v string) {
	o.StreamName = v
}

func (o StreamExportViewResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamExportViewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["created"] = o.Created
	toSerialize["deliver_subject_prefix"] = o.DeliverSubjectPrefix
	toSerialize["id"] = o.Id
	toSerialize["is_public"] = o.IsPublic
	toSerialize["js_subject_prefix"] = o.JsSubjectPrefix
	toSerialize["stream_name"] = o.StreamName
	return toSerialize, nil
}

func (o *StreamExportViewResponse) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_id",
		"created",
		"deliver_subject_prefix",
		"id",
		"is_public",
		"js_subject_prefix",
		"stream_name",
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

	varStreamExportViewResponse := _StreamExportViewResponse{}

	err = json.Unmarshal(bytes, &varStreamExportViewResponse)

	if err != nil {
		return err
	}

	*o = StreamExportViewResponse(varStreamExportViewResponse)

	return err
}

type NullableStreamExportViewResponse struct {
	value *StreamExportViewResponse
	isSet bool
}

func (v NullableStreamExportViewResponse) Get() *StreamExportViewResponse {
	return v.value
}

func (v *NullableStreamExportViewResponse) Set(val *StreamExportViewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamExportViewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamExportViewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamExportViewResponse(val *StreamExportViewResponse) *NullableStreamExportViewResponse {
	return &NullableStreamExportViewResponse{value: val, isSet: true}
}

func (v NullableStreamExportViewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamExportViewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
