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

// InlineResponse2002Results struct for InlineResponse2002Results
type InlineResponse2002Results struct {
	Count *int32 `json:"count,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewInlineResponse2002Results instantiates a new InlineResponse2002Results object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2002Results() *InlineResponse2002Results {
	this := InlineResponse2002Results{}
	return &this
}

// NewInlineResponse2002ResultsWithDefaults instantiates a new InlineResponse2002Results object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2002ResultsWithDefaults() *InlineResponse2002Results {
	this := InlineResponse2002Results{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineResponse2002Results) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Results) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineResponse2002Results) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineResponse2002Results) SetCount(v int32) {
	o.Count = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InlineResponse2002Results) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002Results) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineResponse2002Results) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InlineResponse2002Results) SetValue(v string) {
	o.Value = &v
}

func (o InlineResponse2002Results) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2002Results struct {
	value *InlineResponse2002Results
	isSet bool
}

func (v NullableInlineResponse2002Results) Get() *InlineResponse2002Results {
	return v.value
}

func (v *NullableInlineResponse2002Results) Set(val *InlineResponse2002Results) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2002Results) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2002Results) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2002Results(val *InlineResponse2002Results) *NullableInlineResponse2002Results {
	return &NullableInlineResponse2002Results{value: val, isSet: true}
}

func (v NullableInlineResponse2002Results) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2002Results) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


