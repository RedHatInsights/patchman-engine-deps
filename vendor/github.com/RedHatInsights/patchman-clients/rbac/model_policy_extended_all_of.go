/*
 * Role Based Access Control
 *
 * The API for Role Based Access Control.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rbac

import (
	"encoding/json"
)

// PolicyExtendedAllOf struct for PolicyExtendedAllOf
type PolicyExtendedAllOf struct {
	Group GroupOut `json:"group"`
	Roles []RoleOut `json:"roles"`
}

// NewPolicyExtendedAllOf instantiates a new PolicyExtendedAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyExtendedAllOf(group GroupOut, roles []RoleOut, ) *PolicyExtendedAllOf {
	this := PolicyExtendedAllOf{}
	this.Group = group
	this.Roles = roles
	return &this
}

// NewPolicyExtendedAllOfWithDefaults instantiates a new PolicyExtendedAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyExtendedAllOfWithDefaults() *PolicyExtendedAllOf {
	this := PolicyExtendedAllOf{}
	return &this
}

// GetGroup returns the Group field value
func (o *PolicyExtendedAllOf) GetGroup() GroupOut {
	if o == nil  {
		var ret GroupOut
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *PolicyExtendedAllOf) GetGroupOk() (*GroupOut, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *PolicyExtendedAllOf) SetGroup(v GroupOut) {
	o.Group = v
}

// GetRoles returns the Roles field value
func (o *PolicyExtendedAllOf) GetRoles() []RoleOut {
	if o == nil  {
		var ret []RoleOut
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *PolicyExtendedAllOf) GetRolesOk() (*[]RoleOut, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *PolicyExtendedAllOf) SetRoles(v []RoleOut) {
	o.Roles = v
}

func (o PolicyExtendedAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyExtendedAllOf struct {
	value *PolicyExtendedAllOf
	isSet bool
}

func (v NullablePolicyExtendedAllOf) Get() *PolicyExtendedAllOf {
	return v.value
}

func (v *NullablePolicyExtendedAllOf) Set(val *PolicyExtendedAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyExtendedAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyExtendedAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyExtendedAllOf(val *PolicyExtendedAllOf) *NullablePolicyExtendedAllOf {
	return &NullablePolicyExtendedAllOf{value: val, isSet: true}
}

func (v NullablePolicyExtendedAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyExtendedAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

