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

// CvesRequest struct for CvesRequest
type CvesRequest struct {
	Page *float32 `json:"page,omitempty"`
	PageSize *float32 `json:"page_size,omitempty"`
	CveList []string `json:"cve_list"`
	ModifiedSince *string `json:"modified_since,omitempty"`
	PublishedSince *string `json:"published_since,omitempty"`
	RhOnly *bool `json:"rh_only,omitempty"`
	// Return only those CVEs which are associated with at least one errata. Defaults to false.
	ErrataAssociated *bool `json:"errata_associated,omitempty"`
	// Include content from \"third party\" repositories into the response, disabled by default.
	ThirdParty *bool `json:"third_party,omitempty"`
}

// NewCvesRequest instantiates a new CvesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCvesRequest(cveList []string, ) *CvesRequest {
	this := CvesRequest{}
	this.CveList = cveList
	var thirdParty bool = false
	this.ThirdParty = &thirdParty
	return &this
}

// NewCvesRequestWithDefaults instantiates a new CvesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCvesRequestWithDefaults() *CvesRequest {
	this := CvesRequest{}
	var thirdParty bool = false
	this.ThirdParty = &thirdParty
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *CvesRequest) GetPage() float32 {
	if o == nil || o.Page == nil {
		var ret float32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesRequest) GetPageOk() (*float32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *CvesRequest) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given float32 and assigns it to the Page field.
func (o *CvesRequest) SetPage(v float32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *CvesRequest) GetPageSize() float32 {
	if o == nil || o.PageSize == nil {
		var ret float32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesRequest) GetPageSizeOk() (*float32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *CvesRequest) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given float32 and assigns it to the PageSize field.
func (o *CvesRequest) SetPageSize(v float32) {
	o.PageSize = &v
}

// GetCveList returns the CveList field value
func (o *CvesRequest) GetCveList() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.CveList
}

// GetCveListOk returns a tuple with the CveList field value
// and a boolean to check if the value has been set.
func (o *CvesRequest) GetCveListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CveList, true
}

// SetCveList sets field value
func (o *CvesRequest) SetCveList(v []string) {
	o.CveList = v
}

// GetModifiedSince returns the ModifiedSince field value if set, zero value otherwise.
func (o *CvesRequest) GetModifiedSince() string {
	if o == nil || o.ModifiedSince == nil {
		var ret string
		return ret
	}
	return *o.ModifiedSince
}

// GetModifiedSinceOk returns a tuple with the ModifiedSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesRequest) GetModifiedSinceOk() (*string, bool) {
	if o == nil || o.ModifiedSince == nil {
		return nil, false
	}
	return o.ModifiedSince, true
}

// HasModifiedSince returns a boolean if a field has been set.
func (o *CvesRequest) HasModifiedSince() bool {
	if o != nil && o.ModifiedSince != nil {
		return true
	}

	return false
}

// SetModifiedSince gets a reference to the given string and assigns it to the ModifiedSince field.
func (o *CvesRequest) SetModifiedSince(v string) {
	o.ModifiedSince = &v
}

// GetPublishedSince returns the PublishedSince field value if set, zero value otherwise.
func (o *CvesRequest) GetPublishedSince() string {
	if o == nil || o.PublishedSince == nil {
		var ret string
		return ret
	}
	return *o.PublishedSince
}

// GetPublishedSinceOk returns a tuple with the PublishedSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesRequest) GetPublishedSinceOk() (*string, bool) {
	if o == nil || o.PublishedSince == nil {
		return nil, false
	}
	return o.PublishedSince, true
}

// HasPublishedSince returns a boolean if a field has been set.
func (o *CvesRequest) HasPublishedSince() bool {
	if o != nil && o.PublishedSince != nil {
		return true
	}

	return false
}

// SetPublishedSince gets a reference to the given string and assigns it to the PublishedSince field.
func (o *CvesRequest) SetPublishedSince(v string) {
	o.PublishedSince = &v
}

// GetRhOnly returns the RhOnly field value if set, zero value otherwise.
func (o *CvesRequest) GetRhOnly() bool {
	if o == nil || o.RhOnly == nil {
		var ret bool
		return ret
	}
	return *o.RhOnly
}

// GetRhOnlyOk returns a tuple with the RhOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesRequest) GetRhOnlyOk() (*bool, bool) {
	if o == nil || o.RhOnly == nil {
		return nil, false
	}
	return o.RhOnly, true
}

// HasRhOnly returns a boolean if a field has been set.
func (o *CvesRequest) HasRhOnly() bool {
	if o != nil && o.RhOnly != nil {
		return true
	}

	return false
}

// SetRhOnly gets a reference to the given bool and assigns it to the RhOnly field.
func (o *CvesRequest) SetRhOnly(v bool) {
	o.RhOnly = &v
}

// GetErrataAssociated returns the ErrataAssociated field value if set, zero value otherwise.
func (o *CvesRequest) GetErrataAssociated() bool {
	if o == nil || o.ErrataAssociated == nil {
		var ret bool
		return ret
	}
	return *o.ErrataAssociated
}

// GetErrataAssociatedOk returns a tuple with the ErrataAssociated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesRequest) GetErrataAssociatedOk() (*bool, bool) {
	if o == nil || o.ErrataAssociated == nil {
		return nil, false
	}
	return o.ErrataAssociated, true
}

// HasErrataAssociated returns a boolean if a field has been set.
func (o *CvesRequest) HasErrataAssociated() bool {
	if o != nil && o.ErrataAssociated != nil {
		return true
	}

	return false
}

// SetErrataAssociated gets a reference to the given bool and assigns it to the ErrataAssociated field.
func (o *CvesRequest) SetErrataAssociated(v bool) {
	o.ErrataAssociated = &v
}

// GetThirdParty returns the ThirdParty field value if set, zero value otherwise.
func (o *CvesRequest) GetThirdParty() bool {
	if o == nil || o.ThirdParty == nil {
		var ret bool
		return ret
	}
	return *o.ThirdParty
}

// GetThirdPartyOk returns a tuple with the ThirdParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesRequest) GetThirdPartyOk() (*bool, bool) {
	if o == nil || o.ThirdParty == nil {
		return nil, false
	}
	return o.ThirdParty, true
}

// HasThirdParty returns a boolean if a field has been set.
func (o *CvesRequest) HasThirdParty() bool {
	if o != nil && o.ThirdParty != nil {
		return true
	}

	return false
}

// SetThirdParty gets a reference to the given bool and assigns it to the ThirdParty field.
func (o *CvesRequest) SetThirdParty(v bool) {
	o.ThirdParty = &v
}

func (o CvesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if true {
		toSerialize["cve_list"] = o.CveList
	}
	if o.ModifiedSince != nil {
		toSerialize["modified_since"] = o.ModifiedSince
	}
	if o.PublishedSince != nil {
		toSerialize["published_since"] = o.PublishedSince
	}
	if o.RhOnly != nil {
		toSerialize["rh_only"] = o.RhOnly
	}
	if o.ErrataAssociated != nil {
		toSerialize["errata_associated"] = o.ErrataAssociated
	}
	if o.ThirdParty != nil {
		toSerialize["third_party"] = o.ThirdParty
	}
	return json.Marshal(toSerialize)
}

type NullableCvesRequest struct {
	value *CvesRequest
	isSet bool
}

func (v NullableCvesRequest) Get() *CvesRequest {
	return v.value
}

func (v *NullableCvesRequest) Set(val *CvesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCvesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCvesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCvesRequest(val *CvesRequest) *NullableCvesRequest {
	return &NullableCvesRequest{value: val, isSet: true}
}

func (v NullableCvesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCvesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


