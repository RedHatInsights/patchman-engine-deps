/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.27.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vmaas

import (
	"encoding/json"
)

// DBChangeResponseDbchange struct for DBChangeResponseDbchange
type DBChangeResponseDbchange struct {
	ErrataChanges *string `json:"errata_changes,omitempty"`
	CveChanges *string `json:"cve_changes,omitempty"`
	RepositoryChanges *string `json:"repository_changes,omitempty"`
	LastChange *string `json:"last_change,omitempty"`
	Exported *string `json:"exported,omitempty"`
}

// NewDBChangeResponseDbchange instantiates a new DBChangeResponseDbchange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDBChangeResponseDbchange() *DBChangeResponseDbchange {
	this := DBChangeResponseDbchange{}
	return &this
}

// NewDBChangeResponseDbchangeWithDefaults instantiates a new DBChangeResponseDbchange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDBChangeResponseDbchangeWithDefaults() *DBChangeResponseDbchange {
	this := DBChangeResponseDbchange{}
	return &this
}

// GetErrataChanges returns the ErrataChanges field value if set, zero value otherwise.
func (o *DBChangeResponseDbchange) GetErrataChanges() string {
	if o == nil || o.ErrataChanges == nil {
		var ret string
		return ret
	}
	return *o.ErrataChanges
}

// GetErrataChangesOk returns a tuple with the ErrataChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBChangeResponseDbchange) GetErrataChangesOk() (*string, bool) {
	if o == nil || o.ErrataChanges == nil {
		return nil, false
	}
	return o.ErrataChanges, true
}

// HasErrataChanges returns a boolean if a field has been set.
func (o *DBChangeResponseDbchange) HasErrataChanges() bool {
	if o != nil && o.ErrataChanges != nil {
		return true
	}

	return false
}

// SetErrataChanges gets a reference to the given string and assigns it to the ErrataChanges field.
func (o *DBChangeResponseDbchange) SetErrataChanges(v string) {
	o.ErrataChanges = &v
}

// GetCveChanges returns the CveChanges field value if set, zero value otherwise.
func (o *DBChangeResponseDbchange) GetCveChanges() string {
	if o == nil || o.CveChanges == nil {
		var ret string
		return ret
	}
	return *o.CveChanges
}

// GetCveChangesOk returns a tuple with the CveChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBChangeResponseDbchange) GetCveChangesOk() (*string, bool) {
	if o == nil || o.CveChanges == nil {
		return nil, false
	}
	return o.CveChanges, true
}

// HasCveChanges returns a boolean if a field has been set.
func (o *DBChangeResponseDbchange) HasCveChanges() bool {
	if o != nil && o.CveChanges != nil {
		return true
	}

	return false
}

// SetCveChanges gets a reference to the given string and assigns it to the CveChanges field.
func (o *DBChangeResponseDbchange) SetCveChanges(v string) {
	o.CveChanges = &v
}

// GetRepositoryChanges returns the RepositoryChanges field value if set, zero value otherwise.
func (o *DBChangeResponseDbchange) GetRepositoryChanges() string {
	if o == nil || o.RepositoryChanges == nil {
		var ret string
		return ret
	}
	return *o.RepositoryChanges
}

// GetRepositoryChangesOk returns a tuple with the RepositoryChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBChangeResponseDbchange) GetRepositoryChangesOk() (*string, bool) {
	if o == nil || o.RepositoryChanges == nil {
		return nil, false
	}
	return o.RepositoryChanges, true
}

// HasRepositoryChanges returns a boolean if a field has been set.
func (o *DBChangeResponseDbchange) HasRepositoryChanges() bool {
	if o != nil && o.RepositoryChanges != nil {
		return true
	}

	return false
}

// SetRepositoryChanges gets a reference to the given string and assigns it to the RepositoryChanges field.
func (o *DBChangeResponseDbchange) SetRepositoryChanges(v string) {
	o.RepositoryChanges = &v
}

// GetLastChange returns the LastChange field value if set, zero value otherwise.
func (o *DBChangeResponseDbchange) GetLastChange() string {
	if o == nil || o.LastChange == nil {
		var ret string
		return ret
	}
	return *o.LastChange
}

// GetLastChangeOk returns a tuple with the LastChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBChangeResponseDbchange) GetLastChangeOk() (*string, bool) {
	if o == nil || o.LastChange == nil {
		return nil, false
	}
	return o.LastChange, true
}

// HasLastChange returns a boolean if a field has been set.
func (o *DBChangeResponseDbchange) HasLastChange() bool {
	if o != nil && o.LastChange != nil {
		return true
	}

	return false
}

// SetLastChange gets a reference to the given string and assigns it to the LastChange field.
func (o *DBChangeResponseDbchange) SetLastChange(v string) {
	o.LastChange = &v
}

// GetExported returns the Exported field value if set, zero value otherwise.
func (o *DBChangeResponseDbchange) GetExported() string {
	if o == nil || o.Exported == nil {
		var ret string
		return ret
	}
	return *o.Exported
}

// GetExportedOk returns a tuple with the Exported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DBChangeResponseDbchange) GetExportedOk() (*string, bool) {
	if o == nil || o.Exported == nil {
		return nil, false
	}
	return o.Exported, true
}

// HasExported returns a boolean if a field has been set.
func (o *DBChangeResponseDbchange) HasExported() bool {
	if o != nil && o.Exported != nil {
		return true
	}

	return false
}

// SetExported gets a reference to the given string and assigns it to the Exported field.
func (o *DBChangeResponseDbchange) SetExported(v string) {
	o.Exported = &v
}

func (o DBChangeResponseDbchange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrataChanges != nil {
		toSerialize["errata_changes"] = o.ErrataChanges
	}
	if o.CveChanges != nil {
		toSerialize["cve_changes"] = o.CveChanges
	}
	if o.RepositoryChanges != nil {
		toSerialize["repository_changes"] = o.RepositoryChanges
	}
	if o.LastChange != nil {
		toSerialize["last_change"] = o.LastChange
	}
	if o.Exported != nil {
		toSerialize["exported"] = o.Exported
	}
	return json.Marshal(toSerialize)
}

type NullableDBChangeResponseDbchange struct {
	value *DBChangeResponseDbchange
	isSet bool
}

func (v NullableDBChangeResponseDbchange) Get() *DBChangeResponseDbchange {
	return v.value
}

func (v *NullableDBChangeResponseDbchange) Set(val *DBChangeResponseDbchange) {
	v.value = val
	v.isSet = true
}

func (v NullableDBChangeResponseDbchange) IsSet() bool {
	return v.isSet
}

func (v *NullableDBChangeResponseDbchange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDBChangeResponseDbchange(val *DBChangeResponseDbchange) *NullableDBChangeResponseDbchange {
	return &NullableDBChangeResponseDbchange{value: val, isSet: true}
}

func (v NullableDBChangeResponseDbchange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDBChangeResponseDbchange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

