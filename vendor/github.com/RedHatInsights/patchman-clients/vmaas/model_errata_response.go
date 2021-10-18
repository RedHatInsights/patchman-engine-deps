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

// ErrataResponse struct for ErrataResponse
type ErrataResponse struct {
	Page *float32 `json:"page,omitempty"`
	PageSize *float32 `json:"page_size,omitempty"`
	Pages *float32 `json:"pages,omitempty"`
	ErrataList *map[string]ErrataResponseErrataList `json:"errata_list,omitempty"`
	Type *[]string `json:"type,omitempty"`
	Severity *[]string `json:"severity,omitempty"`
	LastChange *string `json:"last_change,omitempty"`
}

// NewErrataResponse instantiates a new ErrataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrataResponse() *ErrataResponse {
	this := ErrataResponse{}
	return &this
}

// NewErrataResponseWithDefaults instantiates a new ErrataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrataResponseWithDefaults() *ErrataResponse {
	this := ErrataResponse{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ErrataResponse) GetPage() float32 {
	if o == nil || o.Page == nil {
		var ret float32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrataResponse) GetPageOk() (*float32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ErrataResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given float32 and assigns it to the Page field.
func (o *ErrataResponse) SetPage(v float32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ErrataResponse) GetPageSize() float32 {
	if o == nil || o.PageSize == nil {
		var ret float32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrataResponse) GetPageSizeOk() (*float32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ErrataResponse) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given float32 and assigns it to the PageSize field.
func (o *ErrataResponse) SetPageSize(v float32) {
	o.PageSize = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *ErrataResponse) GetPages() float32 {
	if o == nil || o.Pages == nil {
		var ret float32
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrataResponse) GetPagesOk() (*float32, bool) {
	if o == nil || o.Pages == nil {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *ErrataResponse) HasPages() bool {
	if o != nil && o.Pages != nil {
		return true
	}

	return false
}

// SetPages gets a reference to the given float32 and assigns it to the Pages field.
func (o *ErrataResponse) SetPages(v float32) {
	o.Pages = &v
}

// GetErrataList returns the ErrataList field value if set, zero value otherwise.
func (o *ErrataResponse) GetErrataList() map[string]ErrataResponseErrataList {
	if o == nil || o.ErrataList == nil {
		var ret map[string]ErrataResponseErrataList
		return ret
	}
	return *o.ErrataList
}

// GetErrataListOk returns a tuple with the ErrataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrataResponse) GetErrataListOk() (*map[string]ErrataResponseErrataList, bool) {
	if o == nil || o.ErrataList == nil {
		return nil, false
	}
	return o.ErrataList, true
}

// HasErrataList returns a boolean if a field has been set.
func (o *ErrataResponse) HasErrataList() bool {
	if o != nil && o.ErrataList != nil {
		return true
	}

	return false
}

// SetErrataList gets a reference to the given map[string]ErrataResponseErrataList and assigns it to the ErrataList field.
func (o *ErrataResponse) SetErrataList(v map[string]ErrataResponseErrataList) {
	o.ErrataList = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ErrataResponse) GetType() []string {
	if o == nil || o.Type == nil {
		var ret []string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrataResponse) GetTypeOk() (*[]string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ErrataResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given []string and assigns it to the Type field.
func (o *ErrataResponse) SetType(v []string) {
	o.Type = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ErrataResponse) GetSeverity() []string {
	if o == nil || o.Severity == nil {
		var ret []string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrataResponse) GetSeverityOk() (*[]string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ErrataResponse) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given []string and assigns it to the Severity field.
func (o *ErrataResponse) SetSeverity(v []string) {
	o.Severity = &v
}

// GetLastChange returns the LastChange field value if set, zero value otherwise.
func (o *ErrataResponse) GetLastChange() string {
	if o == nil || o.LastChange == nil {
		var ret string
		return ret
	}
	return *o.LastChange
}

// GetLastChangeOk returns a tuple with the LastChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrataResponse) GetLastChangeOk() (*string, bool) {
	if o == nil || o.LastChange == nil {
		return nil, false
	}
	return o.LastChange, true
}

// HasLastChange returns a boolean if a field has been set.
func (o *ErrataResponse) HasLastChange() bool {
	if o != nil && o.LastChange != nil {
		return true
	}

	return false
}

// SetLastChange gets a reference to the given string and assigns it to the LastChange field.
func (o *ErrataResponse) SetLastChange(v string) {
	o.LastChange = &v
}

func (o ErrataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if o.Pages != nil {
		toSerialize["pages"] = o.Pages
	}
	if o.ErrataList != nil {
		toSerialize["errata_list"] = o.ErrataList
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.LastChange != nil {
		toSerialize["last_change"] = o.LastChange
	}
	return json.Marshal(toSerialize)
}

type NullableErrataResponse struct {
	value *ErrataResponse
	isSet bool
}

func (v NullableErrataResponse) Get() *ErrataResponse {
	return v.value
}

func (v *NullableErrataResponse) Set(val *ErrataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrataResponse(val *ErrataResponse) *NullableErrataResponse {
	return &NullableErrataResponse{value: val, isSet: true}
}

func (v NullableErrataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

