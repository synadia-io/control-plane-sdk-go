/*
Synadia Control Plane

API for Synadia Control Plane Server

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the ObjectStoreConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectStoreConfig{}

// ObjectStoreConfig ObjectStoreConfig is the config for the object store.
type ObjectStoreConfig struct {
	Bucket      string  `json:"bucket"`
	Description *string `json:"description,omitempty"`
	MaxAge      *int64  `json:"max_age,omitempty"`
	MaxBytes    *int64  `json:"max_bytes,omitempty"`
	// Bucket-specific metadata NOTE: Metadata requires nats-server v2.10.0+
	Metadata    *map[string]string `json:"metadata,omitempty"`
	NumReplicas *int32             `json:"num_replicas,omitempty"`
	Placement   *Placement         `json:"placement,omitempty"`
	Storage     *StorageType       `json:"storage,omitempty"`
}

func (o ObjectStoreConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucket"] = o.Bucket
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.MaxAge != nil {
		toSerialize["max_age"] = o.MaxAge
	}
	if o.MaxBytes != nil {
		toSerialize["max_bytes"] = o.MaxBytes
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.NumReplicas != nil {
		toSerialize["num_replicas"] = o.NumReplicas
	}
	if o.Placement != nil {
		toSerialize["placement"] = o.Placement
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	return toSerialize, nil
}
