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

// AHostInventoryQueryResult A paginated host search query result with host entries and their Inventory metadata.
type AHostInventoryQueryResult struct {
	// A number of entries on the current page.
	Count int32 `json:"count"`
	// A current page number.
	Page int32 `json:"page"`
	// A page size – a number of entries per single page.
	PerPage int32 `json:"per_page"`
	// Actual host search query result entries.
	Results []map[string]interface{} `json:"results"`
	// A total count of the found entries.
	Total int32 `json:"total"`
}

// NewAHostInventoryQueryResult instantiates a new AHostInventoryQueryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAHostInventoryQueryResult(count int32, page int32, perPage int32, results []map[string]interface{}, total int32, ) *AHostInventoryQueryResult {
	this := AHostInventoryQueryResult{}
	this.Count = count
	this.Page = page
	this.PerPage = perPage
	this.Results = results
	this.Total = total
	return &this
}

// NewAHostInventoryQueryResultWithDefaults instantiates a new AHostInventoryQueryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAHostInventoryQueryResultWithDefaults() *AHostInventoryQueryResult {
	this := AHostInventoryQueryResult{}
	return &this
}

// GetCount returns the Count field value
func (o *AHostInventoryQueryResult) GetCount() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *AHostInventoryQueryResult) GetCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *AHostInventoryQueryResult) SetCount(v int32) {
	o.Count = v
}

// GetPage returns the Page field value
func (o *AHostInventoryQueryResult) GetPage() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *AHostInventoryQueryResult) GetPageOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *AHostInventoryQueryResult) SetPage(v int32) {
	o.Page = v
}

// GetPerPage returns the PerPage field value
func (o *AHostInventoryQueryResult) GetPerPage() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value
// and a boolean to check if the value has been set.
func (o *AHostInventoryQueryResult) GetPerPageOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PerPage, true
}

// SetPerPage sets field value
func (o *AHostInventoryQueryResult) SetPerPage(v int32) {
	o.PerPage = v
}

// GetResults returns the Results field value
func (o *AHostInventoryQueryResult) GetResults() []map[string]interface{} {
	if o == nil  {
		var ret []map[string]interface{}
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *AHostInventoryQueryResult) GetResultsOk() (*[]map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *AHostInventoryQueryResult) SetResults(v []map[string]interface{}) {
	o.Results = v
}

// GetTotal returns the Total field value
func (o *AHostInventoryQueryResult) GetTotal() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *AHostInventoryQueryResult) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *AHostInventoryQueryResult) SetTotal(v int32) {
	o.Total = v
}

func (o AHostInventoryQueryResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["per_page"] = o.PerPage
	}
	if true {
		toSerialize["results"] = o.Results
	}
	if true {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableAHostInventoryQueryResult struct {
	value *AHostInventoryQueryResult
	isSet bool
}

func (v NullableAHostInventoryQueryResult) Get() *AHostInventoryQueryResult {
	return v.value
}

func (v *NullableAHostInventoryQueryResult) Set(val *AHostInventoryQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAHostInventoryQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAHostInventoryQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAHostInventoryQueryResult(val *AHostInventoryQueryResult) *NullableAHostInventoryQueryResult {
	return &NullableAHostInventoryQueryResult{value: val, isSet: true}
}

func (v NullableAHostInventoryQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAHostInventoryQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


