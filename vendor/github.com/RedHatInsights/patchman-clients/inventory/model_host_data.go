/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package inventory

import (
	"encoding/json"
)

// HostData Data of a single host to be updated.
type HostData struct {
	// The ansible host name for remediations
	AnsibleHost *string `json:"ansible_host,omitempty"`
	// A host’s human-readable display name, e.g. in a form of a domain name.
	DisplayName *string `json:"display_name,omitempty"`
}

// NewHostData instantiates a new HostData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostData() *HostData {
	this := HostData{}
	return &this
}

// NewHostDataWithDefaults instantiates a new HostData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostDataWithDefaults() *HostData {
	this := HostData{}
	return &this
}

// GetAnsibleHost returns the AnsibleHost field value if set, zero value otherwise.
func (o *HostData) GetAnsibleHost() string {
	if o == nil || o.AnsibleHost == nil {
		var ret string
		return ret
	}
	return *o.AnsibleHost
}

// GetAnsibleHostOk returns a tuple with the AnsibleHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostData) GetAnsibleHostOk() (*string, bool) {
	if o == nil || o.AnsibleHost == nil {
		return nil, false
	}
	return o.AnsibleHost, true
}

// HasAnsibleHost returns a boolean if a field has been set.
func (o *HostData) HasAnsibleHost() bool {
	if o != nil && o.AnsibleHost != nil {
		return true
	}

	return false
}

// SetAnsibleHost gets a reference to the given string and assigns it to the AnsibleHost field.
func (o *HostData) SetAnsibleHost(v string) {
	o.AnsibleHost = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *HostData) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostData) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *HostData) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *HostData) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o HostData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnsibleHost != nil {
		toSerialize["ansible_host"] = o.AnsibleHost
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	return json.Marshal(toSerialize)
}

type NullableHostData struct {
	value *HostData
	isSet bool
}

func (v NullableHostData) Get() *HostData {
	return v.value
}

func (v *NullableHostData) Set(val *HostData) {
	v.value = val
	v.isSet = true
}

func (v NullableHostData) IsSet() bool {
	return v.isSet
}

func (v *NullableHostData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostData(val *HostData) *NullableHostData {
	return &NullableHostData{value: val, isSet: true}
}

func (v NullableHostData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

