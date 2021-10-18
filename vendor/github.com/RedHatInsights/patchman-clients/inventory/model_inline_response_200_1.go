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

// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	// A number of entries on the current page.
	Count *int32 `json:"count,omitempty"`
	// A current page number.
	Page *int32 `json:"page,omitempty"`
	// A page size – a number of entries per single page.
	PerPage *int32 `json:"per_page,omitempty"`
	// The list of tags on the systems
	Results *map[string]int32 `json:"results,omitempty"`
	// Total number of items in the \"data\" list.
	Total *int32 `json:"total,omitempty"`
}

// NewInlineResponse2001 instantiates a new InlineResponse2001 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001WithDefaults() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineResponse2001) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineResponse2001) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineResponse2001) SetCount(v int32) {
	o.Count = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *InlineResponse2001) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *InlineResponse2001) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *InlineResponse2001) SetPage(v int32) {
	o.Page = &v
}

// GetPerPage returns the PerPage field value if set, zero value otherwise.
func (o *InlineResponse2001) GetPerPage() int32 {
	if o == nil || o.PerPage == nil {
		var ret int32
		return ret
	}
	return *o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetPerPageOk() (*int32, bool) {
	if o == nil || o.PerPage == nil {
		return nil, false
	}
	return o.PerPage, true
}

// HasPerPage returns a boolean if a field has been set.
func (o *InlineResponse2001) HasPerPage() bool {
	if o != nil && o.PerPage != nil {
		return true
	}

	return false
}

// SetPerPage gets a reference to the given int32 and assigns it to the PerPage field.
func (o *InlineResponse2001) SetPerPage(v int32) {
	o.PerPage = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *InlineResponse2001) GetResults() map[string]int32 {
	if o == nil || o.Results == nil {
		var ret map[string]int32
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetResultsOk() (*map[string]int32, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *InlineResponse2001) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given map[string]int32 and assigns it to the Results field.
func (o *InlineResponse2001) SetResults(v map[string]int32) {
	o.Results = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse2001) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse2001) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse2001) SetTotal(v int32) {
	o.Total = &v
}

func (o InlineResponse2001) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.PerPage != nil {
		toSerialize["per_page"] = o.PerPage
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2001 struct {
	value *InlineResponse2001
	isSet bool
}

func (v NullableInlineResponse2001) Get() *InlineResponse2001 {
	return v.value
}

func (v *NullableInlineResponse2001) Set(val *InlineResponse2001) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001(val *InlineResponse2001) *NullableInlineResponse2001 {
	return &NullableInlineResponse2001{value: val, isSet: true}
}

func (v NullableInlineResponse2001) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

