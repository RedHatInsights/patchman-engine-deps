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

// SRPMPkgNamesRequest struct for SRPMPkgNamesRequest
type SRPMPkgNamesRequest struct {
	SrpmNameList []string `json:"srpm_name_list"`
	ContentSetList *[]string `json:"content_set_list,omitempty"`
	// Include content from \"third party\" repositories into the response, disabled by default.
	ThirdParty *bool `json:"third_party,omitempty"`
}

// NewSRPMPkgNamesRequest instantiates a new SRPMPkgNamesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSRPMPkgNamesRequest(srpmNameList []string, ) *SRPMPkgNamesRequest {
	this := SRPMPkgNamesRequest{}
	this.SrpmNameList = srpmNameList
	var thirdParty bool = false
	this.ThirdParty = &thirdParty
	return &this
}

// NewSRPMPkgNamesRequestWithDefaults instantiates a new SRPMPkgNamesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSRPMPkgNamesRequestWithDefaults() *SRPMPkgNamesRequest {
	this := SRPMPkgNamesRequest{}
	var thirdParty bool = false
	this.ThirdParty = &thirdParty
	return &this
}

// GetSrpmNameList returns the SrpmNameList field value
func (o *SRPMPkgNamesRequest) GetSrpmNameList() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.SrpmNameList
}

// GetSrpmNameListOk returns a tuple with the SrpmNameList field value
// and a boolean to check if the value has been set.
func (o *SRPMPkgNamesRequest) GetSrpmNameListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SrpmNameList, true
}

// SetSrpmNameList sets field value
func (o *SRPMPkgNamesRequest) SetSrpmNameList(v []string) {
	o.SrpmNameList = v
}

// GetContentSetList returns the ContentSetList field value if set, zero value otherwise.
func (o *SRPMPkgNamesRequest) GetContentSetList() []string {
	if o == nil || o.ContentSetList == nil {
		var ret []string
		return ret
	}
	return *o.ContentSetList
}

// GetContentSetListOk returns a tuple with the ContentSetList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SRPMPkgNamesRequest) GetContentSetListOk() (*[]string, bool) {
	if o == nil || o.ContentSetList == nil {
		return nil, false
	}
	return o.ContentSetList, true
}

// HasContentSetList returns a boolean if a field has been set.
func (o *SRPMPkgNamesRequest) HasContentSetList() bool {
	if o != nil && o.ContentSetList != nil {
		return true
	}

	return false
}

// SetContentSetList gets a reference to the given []string and assigns it to the ContentSetList field.
func (o *SRPMPkgNamesRequest) SetContentSetList(v []string) {
	o.ContentSetList = &v
}

// GetThirdParty returns the ThirdParty field value if set, zero value otherwise.
func (o *SRPMPkgNamesRequest) GetThirdParty() bool {
	if o == nil || o.ThirdParty == nil {
		var ret bool
		return ret
	}
	return *o.ThirdParty
}

// GetThirdPartyOk returns a tuple with the ThirdParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SRPMPkgNamesRequest) GetThirdPartyOk() (*bool, bool) {
	if o == nil || o.ThirdParty == nil {
		return nil, false
	}
	return o.ThirdParty, true
}

// HasThirdParty returns a boolean if a field has been set.
func (o *SRPMPkgNamesRequest) HasThirdParty() bool {
	if o != nil && o.ThirdParty != nil {
		return true
	}

	return false
}

// SetThirdParty gets a reference to the given bool and assigns it to the ThirdParty field.
func (o *SRPMPkgNamesRequest) SetThirdParty(v bool) {
	o.ThirdParty = &v
}

func (o SRPMPkgNamesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["srpm_name_list"] = o.SrpmNameList
	}
	if o.ContentSetList != nil {
		toSerialize["content_set_list"] = o.ContentSetList
	}
	if o.ThirdParty != nil {
		toSerialize["third_party"] = o.ThirdParty
	}
	return json.Marshal(toSerialize)
}

type NullableSRPMPkgNamesRequest struct {
	value *SRPMPkgNamesRequest
	isSet bool
}

func (v NullableSRPMPkgNamesRequest) Get() *SRPMPkgNamesRequest {
	return v.value
}

func (v *NullableSRPMPkgNamesRequest) Set(val *SRPMPkgNamesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSRPMPkgNamesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSRPMPkgNamesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSRPMPkgNamesRequest(val *SRPMPkgNamesRequest) *NullableSRPMPkgNamesRequest {
	return &NullableSRPMPkgNamesRequest{value: val, isSet: true}
}

func (v NullableSRPMPkgNamesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSRPMPkgNamesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


