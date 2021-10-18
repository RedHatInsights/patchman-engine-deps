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

// ResourceDefinitionFilter struct for ResourceDefinitionFilter
type ResourceDefinitionFilter struct {
	Key string `json:"key"`
	Operation string `json:"operation"`
	Value string `json:"value"`
}

// NewResourceDefinitionFilter instantiates a new ResourceDefinitionFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceDefinitionFilter(key string, operation string, value string, ) *ResourceDefinitionFilter {
	this := ResourceDefinitionFilter{}
	this.Key = key
	this.Operation = operation
	this.Value = value
	return &this
}

// NewResourceDefinitionFilterWithDefaults instantiates a new ResourceDefinitionFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceDefinitionFilterWithDefaults() *ResourceDefinitionFilter {
	this := ResourceDefinitionFilter{}
	return &this
}

// GetKey returns the Key field value
func (o *ResourceDefinitionFilter) GetKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ResourceDefinitionFilter) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ResourceDefinitionFilter) SetKey(v string) {
	o.Key = v
}

// GetOperation returns the Operation field value
func (o *ResourceDefinitionFilter) GetOperation() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *ResourceDefinitionFilter) GetOperationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *ResourceDefinitionFilter) SetOperation(v string) {
	o.Operation = v
}

// GetValue returns the Value field value
func (o *ResourceDefinitionFilter) GetValue() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ResourceDefinitionFilter) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ResourceDefinitionFilter) SetValue(v string) {
	o.Value = v
}

func (o ResourceDefinitionFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["operation"] = o.Operation
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableResourceDefinitionFilter struct {
	value *ResourceDefinitionFilter
	isSet bool
}

func (v NullableResourceDefinitionFilter) Get() *ResourceDefinitionFilter {
	return v.value
}

func (v *NullableResourceDefinitionFilter) Set(val *ResourceDefinitionFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceDefinitionFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceDefinitionFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceDefinitionFilter(val *ResourceDefinitionFilter) *NullableResourceDefinitionFilter {
	return &NullableResourceDefinitionFilter{value: val, isSet: true}
}

func (v NullableResourceDefinitionFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceDefinitionFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

