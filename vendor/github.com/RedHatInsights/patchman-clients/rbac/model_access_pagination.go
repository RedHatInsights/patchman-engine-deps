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

// AccessPagination struct for AccessPagination
type AccessPagination struct {
	Meta *PaginationMeta `json:"meta,omitempty"`
	Links *PaginationLinks `json:"links,omitempty"`
	Data []Access `json:"data"`
}

// NewAccessPagination instantiates a new AccessPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPagination(data []Access, ) *AccessPagination {
	this := AccessPagination{}
	this.Data = data
	return &this
}

// NewAccessPaginationWithDefaults instantiates a new AccessPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPaginationWithDefaults() *AccessPagination {
	this := AccessPagination{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AccessPagination) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPagination) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AccessPagination) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *AccessPagination) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AccessPagination) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPagination) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AccessPagination) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *AccessPagination) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetData returns the Data field value
func (o *AccessPagination) GetData() []Access {
	if o == nil  {
		var ret []Access
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AccessPagination) GetDataOk() (*[]Access, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AccessPagination) SetData(v []Access) {
	o.Data = v
}

func (o AccessPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAccessPagination struct {
	value *AccessPagination
	isSet bool
}

func (v NullableAccessPagination) Get() *AccessPagination {
	return v.value
}

func (v *NullableAccessPagination) Set(val *AccessPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPagination(val *AccessPagination) *NullableAccessPagination {
	return &NullableAccessPagination{value: val, isSet: true}
}

func (v NullableAccessPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

