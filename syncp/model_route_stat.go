/*
Synadia Control Plane / Synadia Cloud

API for Synadia Control Plane / Synadia Cloud

API version: beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package syncp

// checks if the RouteStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteStat{}

// RouteStat RouteStat holds route statistics.
type RouteStat struct {
	Name     *string   `json:"name,omitempty"`
	Pending  int64     `json:"pending"`
	Received DataStats `json:"received"`
	Rid      uint64    `json:"rid"`
	Sent     DataStats `json:"sent"`
}

func (o RouteStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["pending"] = o.Pending
	toSerialize["received"] = o.Received
	toSerialize["rid"] = o.Rid
	toSerialize["sent"] = o.Sent
	return toSerialize, nil
}
