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

// Access struct for Access
type Access struct {
	Permission string `json:"permission"`
	ResourceDefinitions []ResourceDefinition `json:"resourceDefinitions"`
}

// NewAccess instantiates a new Access object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccess(permission string, resourceDefinitions []ResourceDefinition, ) *Access {
	this := Access{}
	this.Permission = permission
	this.ResourceDefinitions = resourceDefinitions
	return &this
}

// NewAccessWithDefaults instantiates a new Access object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessWithDefaults() *Access {
	this := Access{}
	return &this
}

// GetPermission returns the Permission field value
func (o *Access) GetPermission() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *Access) GetPermissionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *Access) SetPermission(v string) {
	o.Permission = v
}

// GetResourceDefinitions returns the ResourceDefinitions field value
func (o *Access) GetResourceDefinitions() []ResourceDefinition {
	if o == nil  {
		var ret []ResourceDefinition
		return ret
	}

	return o.ResourceDefinitions
}

// GetResourceDefinitionsOk returns a tuple with the ResourceDefinitions field value
// and a boolean to check if the value has been set.
func (o *Access) GetResourceDefinitionsOk() (*[]ResourceDefinition, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResourceDefinitions, true
}

// SetResourceDefinitions sets field value
func (o *Access) SetResourceDefinitions(v []ResourceDefinition) {
	o.ResourceDefinitions = v
}

func (o Access) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["permission"] = o.Permission
	}
	if true {
		toSerialize["resourceDefinitions"] = o.ResourceDefinitions
	}
	return json.Marshal(toSerialize)
}

type NullableAccess struct {
	value *Access
	isSet bool
}

func (v NullableAccess) Get() *Access {
	return v.value
}

func (v *NullableAccess) Set(val *Access) {
	v.value = val
	v.isSet = true
}

func (v NullableAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccess(val *Access) *NullableAccess {
	return &NullableAccess{value: val, isSet: true}
}

func (v NullableAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

