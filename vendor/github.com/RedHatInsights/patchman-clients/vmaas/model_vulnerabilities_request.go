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

// VulnerabilitiesRequest struct for VulnerabilitiesRequest
type VulnerabilitiesRequest struct {
	PackageList []string `json:"package_list"`
	RepositoryList *[]string `json:"repository_list,omitempty"`
	ModulesList *[]UpdatesV3RequestModulesList `json:"modules_list,omitempty"`
	Releasever *string `json:"releasever,omitempty"`
	Basearch *string `json:"basearch,omitempty"`
	Oval *bool `json:"oval,omitempty"`
	OvalOnly *bool `json:"oval_only,omitempty"`
	// Include content from \"third party\" repositories into the response, disabled by default.
	ThirdParty *bool `json:"third_party,omitempty"`
}

// NewVulnerabilitiesRequest instantiates a new VulnerabilitiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVulnerabilitiesRequest(packageList []string, ) *VulnerabilitiesRequest {
	this := VulnerabilitiesRequest{}
	this.PackageList = packageList
	var thirdParty bool = false
	this.ThirdParty = &thirdParty
	return &this
}

// NewVulnerabilitiesRequestWithDefaults instantiates a new VulnerabilitiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVulnerabilitiesRequestWithDefaults() *VulnerabilitiesRequest {
	this := VulnerabilitiesRequest{}
	var thirdParty bool = false
	this.ThirdParty = &thirdParty
	return &this
}

// GetPackageList returns the PackageList field value
func (o *VulnerabilitiesRequest) GetPackageList() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.PackageList
}

// GetPackageListOk returns a tuple with the PackageList field value
// and a boolean to check if the value has been set.
func (o *VulnerabilitiesRequest) GetPackageListOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PackageList, true
}

// SetPackageList sets field value
func (o *VulnerabilitiesRequest) SetPackageList(v []string) {
	o.PackageList = v
}

// GetRepositoryList returns the RepositoryList field value if set, zero value otherwise.
func (o *VulnerabilitiesRequest) GetRepositoryList() []string {
	if o == nil || o.RepositoryList == nil {
		var ret []string
		return ret
	}
	return *o.RepositoryList
}

// GetRepositoryListOk returns a tuple with the RepositoryList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnerabilitiesRequest) GetRepositoryListOk() (*[]string, bool) {
	if o == nil || o.RepositoryList == nil {
		return nil, false
	}
	return o.RepositoryList, true
}

// HasRepositoryList returns a boolean if a field has been set.
func (o *VulnerabilitiesRequest) HasRepositoryList() bool {
	if o != nil && o.RepositoryList != nil {
		return true
	}

	return false
}

// SetRepositoryList gets a reference to the given []string and assigns it to the RepositoryList field.
func (o *VulnerabilitiesRequest) SetRepositoryList(v []string) {
	o.RepositoryList = &v
}

// GetModulesList returns the ModulesList field value if set, zero value otherwise.
func (o *VulnerabilitiesRequest) GetModulesList() []UpdatesV3RequestModulesList {
	if o == nil || o.ModulesList == nil {
		var ret []UpdatesV3RequestModulesList
		return ret
	}
	return *o.ModulesList
}

// GetModulesListOk returns a tuple with the ModulesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnerabilitiesRequest) GetModulesListOk() (*[]UpdatesV3RequestModulesList, bool) {
	if o == nil || o.ModulesList == nil {
		return nil, false
	}
	return o.ModulesList, true
}

// HasModulesList returns a boolean if a field has been set.
func (o *VulnerabilitiesRequest) HasModulesList() bool {
	if o != nil && o.ModulesList != nil {
		return true
	}

	return false
}

// SetModulesList gets a reference to the given []UpdatesV3RequestModulesList and assigns it to the ModulesList field.
func (o *VulnerabilitiesRequest) SetModulesList(v []UpdatesV3RequestModulesList) {
	o.ModulesList = &v
}

// GetReleasever returns the Releasever field value if set, zero value otherwise.
func (o *VulnerabilitiesRequest) GetReleasever() string {
	if o == nil || o.Releasever == nil {
		var ret string
		return ret
	}
	return *o.Releasever
}

// GetReleaseverOk returns a tuple with the Releasever field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnerabilitiesRequest) GetReleaseverOk() (*string, bool) {
	if o == nil || o.Releasever == nil {
		return nil, false
	}
	return o.Releasever, true
}

// HasReleasever returns a boolean if a field has been set.
func (o *VulnerabilitiesRequest) HasReleasever() bool {
	if o != nil && o.Releasever != nil {
		return true
	}

	return false
}

// SetReleasever gets a reference to the given string and assigns it to the Releasever field.
func (o *VulnerabilitiesRequest) SetReleasever(v string) {
	o.Releasever = &v
}

// GetBasearch returns the Basearch field value if set, zero value otherwise.
func (o *VulnerabilitiesRequest) GetBasearch() string {
	if o == nil || o.Basearch == nil {
		var ret string
		return ret
	}
	return *o.Basearch
}

// GetBasearchOk returns a tuple with the Basearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnerabilitiesRequest) GetBasearchOk() (*string, bool) {
	if o == nil || o.Basearch == nil {
		return nil, false
	}
	return o.Basearch, true
}

// HasBasearch returns a boolean if a field has been set.
func (o *VulnerabilitiesRequest) HasBasearch() bool {
	if o != nil && o.Basearch != nil {
		return true
	}

	return false
}

// SetBasearch gets a reference to the given string and assigns it to the Basearch field.
func (o *VulnerabilitiesRequest) SetBasearch(v string) {
	o.Basearch = &v
}

// GetOval returns the Oval field value if set, zero value otherwise.
func (o *VulnerabilitiesRequest) GetOval() bool {
	if o == nil || o.Oval == nil {
		var ret bool
		return ret
	}
	return *o.Oval
}

// GetOvalOk returns a tuple with the Oval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnerabilitiesRequest) GetOvalOk() (*bool, bool) {
	if o == nil || o.Oval == nil {
		return nil, false
	}
	return o.Oval, true
}

// HasOval returns a boolean if a field has been set.
func (o *VulnerabilitiesRequest) HasOval() bool {
	if o != nil && o.Oval != nil {
		return true
	}

	return false
}

// SetOval gets a reference to the given bool and assigns it to the Oval field.
func (o *VulnerabilitiesRequest) SetOval(v bool) {
	o.Oval = &v
}

// GetOvalOnly returns the OvalOnly field value if set, zero value otherwise.
func (o *VulnerabilitiesRequest) GetOvalOnly() bool {
	if o == nil || o.OvalOnly == nil {
		var ret bool
		return ret
	}
	return *o.OvalOnly
}

// GetOvalOnlyOk returns a tuple with the OvalOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnerabilitiesRequest) GetOvalOnlyOk() (*bool, bool) {
	if o == nil || o.OvalOnly == nil {
		return nil, false
	}
	return o.OvalOnly, true
}

// HasOvalOnly returns a boolean if a field has been set.
func (o *VulnerabilitiesRequest) HasOvalOnly() bool {
	if o != nil && o.OvalOnly != nil {
		return true
	}

	return false
}

// SetOvalOnly gets a reference to the given bool and assigns it to the OvalOnly field.
func (o *VulnerabilitiesRequest) SetOvalOnly(v bool) {
	o.OvalOnly = &v
}

// GetThirdParty returns the ThirdParty field value if set, zero value otherwise.
func (o *VulnerabilitiesRequest) GetThirdParty() bool {
	if o == nil || o.ThirdParty == nil {
		var ret bool
		return ret
	}
	return *o.ThirdParty
}

// GetThirdPartyOk returns a tuple with the ThirdParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnerabilitiesRequest) GetThirdPartyOk() (*bool, bool) {
	if o == nil || o.ThirdParty == nil {
		return nil, false
	}
	return o.ThirdParty, true
}

// HasThirdParty returns a boolean if a field has been set.
func (o *VulnerabilitiesRequest) HasThirdParty() bool {
	if o != nil && o.ThirdParty != nil {
		return true
	}

	return false
}

// SetThirdParty gets a reference to the given bool and assigns it to the ThirdParty field.
func (o *VulnerabilitiesRequest) SetThirdParty(v bool) {
	o.ThirdParty = &v
}

func (o VulnerabilitiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["package_list"] = o.PackageList
	}
	if o.RepositoryList != nil {
		toSerialize["repository_list"] = o.RepositoryList
	}
	if o.ModulesList != nil {
		toSerialize["modules_list"] = o.ModulesList
	}
	if o.Releasever != nil {
		toSerialize["releasever"] = o.Releasever
	}
	if o.Basearch != nil {
		toSerialize["basearch"] = o.Basearch
	}
	if o.Oval != nil {
		toSerialize["oval"] = o.Oval
	}
	if o.OvalOnly != nil {
		toSerialize["oval_only"] = o.OvalOnly
	}
	if o.ThirdParty != nil {
		toSerialize["third_party"] = o.ThirdParty
	}
	return json.Marshal(toSerialize)
}

type NullableVulnerabilitiesRequest struct {
	value *VulnerabilitiesRequest
	isSet bool
}

func (v NullableVulnerabilitiesRequest) Get() *VulnerabilitiesRequest {
	return v.value
}

func (v *NullableVulnerabilitiesRequest) Set(val *VulnerabilitiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVulnerabilitiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVulnerabilitiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVulnerabilitiesRequest(val *VulnerabilitiesRequest) *NullableVulnerabilitiesRequest {
	return &NullableVulnerabilitiesRequest{value: val, isSet: true}
}

func (v NullableVulnerabilitiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVulnerabilitiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

