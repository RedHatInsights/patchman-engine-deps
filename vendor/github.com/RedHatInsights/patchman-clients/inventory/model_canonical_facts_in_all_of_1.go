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

// CanonicalFactsInAllOf1 struct for CanonicalFactsInAllOf1
type CanonicalFactsInAllOf1 struct {
	BiosUuid *string `json:"bios_uuid,omitempty"`
	ExternalId *string `json:"external_id,omitempty"`
	Fqdn *string `json:"fqdn,omitempty"`
	InsightsId *string `json:"insights_id,omitempty"`
	IpAddresses *[]string `json:"ip_addresses,omitempty"`
	MacAddresses *[]string `json:"mac_addresses,omitempty"`
	RhelMachineId *string `json:"rhel_machine_id,omitempty"`
	SatelliteId *string `json:"satellite_id,omitempty"`
	SubscriptionManagerId *string `json:"subscription_manager_id,omitempty"`
}

// NewCanonicalFactsInAllOf1 instantiates a new CanonicalFactsInAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCanonicalFactsInAllOf1() *CanonicalFactsInAllOf1 {
	this := CanonicalFactsInAllOf1{}
	return &this
}

// NewCanonicalFactsInAllOf1WithDefaults instantiates a new CanonicalFactsInAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCanonicalFactsInAllOf1WithDefaults() *CanonicalFactsInAllOf1 {
	this := CanonicalFactsInAllOf1{}
	return &this
}

// GetBiosUuid returns the BiosUuid field value if set, zero value otherwise.
func (o *CanonicalFactsInAllOf1) GetBiosUuid() string {
	if o == nil || o.BiosUuid == nil {
		var ret string
		return ret
	}
	return *o.BiosUuid
}

// GetBiosUuidOk returns a tuple with the BiosUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanonicalFactsInAllOf1) GetBiosUuidOk() (*string, bool) {
	if o == nil || o.BiosUuid == nil {
		return nil, false
	}
	return o.BiosUuid, true
}

// HasBiosUuid returns a boolean if a field has been set.
func (o *CanonicalFactsInAllOf1) HasBiosUuid() bool {
	if o != nil && o.BiosUuid != nil {
		return true
	}

	return false
}

// SetBiosUuid gets a reference to the given string and assigns it to the BiosUuid field.
func (o *CanonicalFactsInAllOf1) SetBiosUuid(v string) {
	o.BiosUuid = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *CanonicalFactsInAllOf1) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanonicalFactsInAllOf1) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *CanonicalFactsInAllOf1) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *CanonicalFactsInAllOf1) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *CanonicalFactsInAllOf1) GetFqdn() string {
	if o == nil || o.Fqdn == nil {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanonicalFactsInAllOf1) GetFqdnOk() (*string, bool) {
	if o == nil || o.Fqdn == nil {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *CanonicalFactsInAllOf1) HasFqdn() bool {
	if o != nil && o.Fqdn != nil {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *CanonicalFactsInAllOf1) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetInsightsId returns the InsightsId field value if set, zero value otherwise.
func (o *CanonicalFactsInAllOf1) GetInsightsId() string {
	if o == nil || o.InsightsId == nil {
		var ret string
		return ret
	}
	return *o.InsightsId
}

// GetInsightsIdOk returns a tuple with the InsightsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanonicalFactsInAllOf1) GetInsightsIdOk() (*string, bool) {
	if o == nil || o.InsightsId == nil {
		return nil, false
	}
	return o.InsightsId, true
}

// HasInsightsId returns a boolean if a field has been set.
func (o *CanonicalFactsInAllOf1) HasInsightsId() bool {
	if o != nil && o.InsightsId != nil {
		return true
	}

	return false
}

// SetInsightsId gets a reference to the given string and assigns it to the InsightsId field.
func (o *CanonicalFactsInAllOf1) SetInsightsId(v string) {
	o.InsightsId = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *CanonicalFactsInAllOf1) GetIpAddresses() []string {
	if o == nil || o.IpAddresses == nil {
		var ret []string
		return ret
	}
	return *o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanonicalFactsInAllOf1) GetIpAddressesOk() (*[]string, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *CanonicalFactsInAllOf1) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *CanonicalFactsInAllOf1) SetIpAddresses(v []string) {
	o.IpAddresses = &v
}

// GetMacAddresses returns the MacAddresses field value if set, zero value otherwise.
func (o *CanonicalFactsInAllOf1) GetMacAddresses() []string {
	if o == nil || o.MacAddresses == nil {
		var ret []string
		return ret
	}
	return *o.MacAddresses
}

// GetMacAddressesOk returns a tuple with the MacAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanonicalFactsInAllOf1) GetMacAddressesOk() (*[]string, bool) {
	if o == nil || o.MacAddresses == nil {
		return nil, false
	}
	return o.MacAddresses, true
}

// HasMacAddresses returns a boolean if a field has been set.
func (o *CanonicalFactsInAllOf1) HasMacAddresses() bool {
	if o != nil && o.MacAddresses != nil {
		return true
	}

	return false
}

// SetMacAddresses gets a reference to the given []string and assigns it to the MacAddresses field.
func (o *CanonicalFactsInAllOf1) SetMacAddresses(v []string) {
	o.MacAddresses = &v
}

// GetRhelMachineId returns the RhelMachineId field value if set, zero value otherwise.
func (o *CanonicalFactsInAllOf1) GetRhelMachineId() string {
	if o == nil || o.RhelMachineId == nil {
		var ret string
		return ret
	}
	return *o.RhelMachineId
}

// GetRhelMachineIdOk returns a tuple with the RhelMachineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanonicalFactsInAllOf1) GetRhelMachineIdOk() (*string, bool) {
	if o == nil || o.RhelMachineId == nil {
		return nil, false
	}
	return o.RhelMachineId, true
}

// HasRhelMachineId returns a boolean if a field has been set.
func (o *CanonicalFactsInAllOf1) HasRhelMachineId() bool {
	if o != nil && o.RhelMachineId != nil {
		return true
	}

	return false
}

// SetRhelMachineId gets a reference to the given string and assigns it to the RhelMachineId field.
func (o *CanonicalFactsInAllOf1) SetRhelMachineId(v string) {
	o.RhelMachineId = &v
}

// GetSatelliteId returns the SatelliteId field value if set, zero value otherwise.
func (o *CanonicalFactsInAllOf1) GetSatelliteId() string {
	if o == nil || o.SatelliteId == nil {
		var ret string
		return ret
	}
	return *o.SatelliteId
}

// GetSatelliteIdOk returns a tuple with the SatelliteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanonicalFactsInAllOf1) GetSatelliteIdOk() (*string, bool) {
	if o == nil || o.SatelliteId == nil {
		return nil, false
	}
	return o.SatelliteId, true
}

// HasSatelliteId returns a boolean if a field has been set.
func (o *CanonicalFactsInAllOf1) HasSatelliteId() bool {
	if o != nil && o.SatelliteId != nil {
		return true
	}

	return false
}

// SetSatelliteId gets a reference to the given string and assigns it to the SatelliteId field.
func (o *CanonicalFactsInAllOf1) SetSatelliteId(v string) {
	o.SatelliteId = &v
}

// GetSubscriptionManagerId returns the SubscriptionManagerId field value if set, zero value otherwise.
func (o *CanonicalFactsInAllOf1) GetSubscriptionManagerId() string {
	if o == nil || o.SubscriptionManagerId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionManagerId
}

// GetSubscriptionManagerIdOk returns a tuple with the SubscriptionManagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanonicalFactsInAllOf1) GetSubscriptionManagerIdOk() (*string, bool) {
	if o == nil || o.SubscriptionManagerId == nil {
		return nil, false
	}
	return o.SubscriptionManagerId, true
}

// HasSubscriptionManagerId returns a boolean if a field has been set.
func (o *CanonicalFactsInAllOf1) HasSubscriptionManagerId() bool {
	if o != nil && o.SubscriptionManagerId != nil {
		return true
	}

	return false
}

// SetSubscriptionManagerId gets a reference to the given string and assigns it to the SubscriptionManagerId field.
func (o *CanonicalFactsInAllOf1) SetSubscriptionManagerId(v string) {
	o.SubscriptionManagerId = &v
}

func (o CanonicalFactsInAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BiosUuid != nil {
		toSerialize["bios_uuid"] = o.BiosUuid
	}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.Fqdn != nil {
		toSerialize["fqdn"] = o.Fqdn
	}
	if o.InsightsId != nil {
		toSerialize["insights_id"] = o.InsightsId
	}
	if o.IpAddresses != nil {
		toSerialize["ip_addresses"] = o.IpAddresses
	}
	if o.MacAddresses != nil {
		toSerialize["mac_addresses"] = o.MacAddresses
	}
	if o.RhelMachineId != nil {
		toSerialize["rhel_machine_id"] = o.RhelMachineId
	}
	if o.SatelliteId != nil {
		toSerialize["satellite_id"] = o.SatelliteId
	}
	if o.SubscriptionManagerId != nil {
		toSerialize["subscription_manager_id"] = o.SubscriptionManagerId
	}
	return json.Marshal(toSerialize)
}

type NullableCanonicalFactsInAllOf1 struct {
	value *CanonicalFactsInAllOf1
	isSet bool
}

func (v NullableCanonicalFactsInAllOf1) Get() *CanonicalFactsInAllOf1 {
	return v.value
}

func (v *NullableCanonicalFactsInAllOf1) Set(val *CanonicalFactsInAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableCanonicalFactsInAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableCanonicalFactsInAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCanonicalFactsInAllOf1(val *CanonicalFactsInAllOf1) *NullableCanonicalFactsInAllOf1 {
	return &NullableCanonicalFactsInAllOf1{value: val, isSet: true}
}

func (v NullableCanonicalFactsInAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCanonicalFactsInAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


