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

// CanonicalFactsOut struct for CanonicalFactsOut
type CanonicalFactsOut struct {
	// A UUID of the host machine BIOS.  This field is considered to be a canonical fact.
	BiosUuid NullableString `json:"bios_uuid,omitempty"`
	// Host’s reference in the external source e.g. AWS EC2, Azure, OpenStack, etc. This field is considered to be a canonical fact.
	ExternalId NullableString `json:"external_id,omitempty"`
	// A host’s Fully Qualified Domain Name.  This field is considered to be a canonical fact.
	Fqdn NullableString `json:"fqdn,omitempty"`
	// An ID defined in /etc/insights-client/machine-id. This field is considered a canonical fact.
	InsightsId NullableString `json:"insights_id,omitempty"`
	// Host’s network IP addresses.  This field is considered to be a canonical fact.
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Host’s network interfaces MAC addresses.  This field is considered to be a canonical fact.
	MacAddresses []string `json:"mac_addresses,omitempty"`
	// A Machine ID of a RHEL host.  This field is considered to be a canonical fact.
	RhelMachineId NullableString `json:"rhel_machine_id,omitempty"`
	// A Red Hat Satellite ID of a RHEL host.  This field is considered to be a canonical fact.
	SatelliteId NullableString `json:"satellite_id,omitempty"`
	// A Red Hat Subcription Manager ID of a RHEL host.  This field is considered to be a canonical fact.
	SubscriptionManagerId NullableString `json:"subscription_manager_id,omitempty"`
}

// NewCanonicalFactsOut instantiates a new CanonicalFactsOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCanonicalFactsOut() *CanonicalFactsOut {
	this := CanonicalFactsOut{}
	return &this
}

// NewCanonicalFactsOutWithDefaults instantiates a new CanonicalFactsOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCanonicalFactsOutWithDefaults() *CanonicalFactsOut {
	this := CanonicalFactsOut{}
	return &this
}

// GetBiosUuid returns the BiosUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CanonicalFactsOut) GetBiosUuid() string {
	if o == nil || o.BiosUuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.BiosUuid.Get()
}

// GetBiosUuidOk returns a tuple with the BiosUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CanonicalFactsOut) GetBiosUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BiosUuid.Get(), o.BiosUuid.IsSet()
}

// HasBiosUuid returns a boolean if a field has been set.
func (o *CanonicalFactsOut) HasBiosUuid() bool {
	if o != nil && o.BiosUuid.IsSet() {
		return true
	}

	return false
}

// SetBiosUuid gets a reference to the given NullableString and assigns it to the BiosUuid field.
func (o *CanonicalFactsOut) SetBiosUuid(v string) {
	o.BiosUuid.Set(&v)
}
// SetBiosUuidNil sets the value for BiosUuid to be an explicit nil
func (o *CanonicalFactsOut) SetBiosUuidNil() {
	o.BiosUuid.Set(nil)
}

// UnsetBiosUuid ensures that no value is present for BiosUuid, not even an explicit nil
func (o *CanonicalFactsOut) UnsetBiosUuid() {
	o.BiosUuid.Unset()
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CanonicalFactsOut) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CanonicalFactsOut) GetExternalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *CanonicalFactsOut) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *CanonicalFactsOut) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *CanonicalFactsOut) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *CanonicalFactsOut) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CanonicalFactsOut) GetFqdn() string {
	if o == nil || o.Fqdn.Get() == nil {
		var ret string
		return ret
	}
	return *o.Fqdn.Get()
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CanonicalFactsOut) GetFqdnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Fqdn.Get(), o.Fqdn.IsSet()
}

// HasFqdn returns a boolean if a field has been set.
func (o *CanonicalFactsOut) HasFqdn() bool {
	if o != nil && o.Fqdn.IsSet() {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given NullableString and assigns it to the Fqdn field.
func (o *CanonicalFactsOut) SetFqdn(v string) {
	o.Fqdn.Set(&v)
}
// SetFqdnNil sets the value for Fqdn to be an explicit nil
func (o *CanonicalFactsOut) SetFqdnNil() {
	o.Fqdn.Set(nil)
}

// UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
func (o *CanonicalFactsOut) UnsetFqdn() {
	o.Fqdn.Unset()
}

// GetInsightsId returns the InsightsId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CanonicalFactsOut) GetInsightsId() string {
	if o == nil || o.InsightsId.Get() == nil {
		var ret string
		return ret
	}
	return *o.InsightsId.Get()
}

// GetInsightsIdOk returns a tuple with the InsightsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CanonicalFactsOut) GetInsightsIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InsightsId.Get(), o.InsightsId.IsSet()
}

// HasInsightsId returns a boolean if a field has been set.
func (o *CanonicalFactsOut) HasInsightsId() bool {
	if o != nil && o.InsightsId.IsSet() {
		return true
	}

	return false
}

// SetInsightsId gets a reference to the given NullableString and assigns it to the InsightsId field.
func (o *CanonicalFactsOut) SetInsightsId(v string) {
	o.InsightsId.Set(&v)
}
// SetInsightsIdNil sets the value for InsightsId to be an explicit nil
func (o *CanonicalFactsOut) SetInsightsIdNil() {
	o.InsightsId.Set(nil)
}

// UnsetInsightsId ensures that no value is present for InsightsId, not even an explicit nil
func (o *CanonicalFactsOut) UnsetInsightsId() {
	o.InsightsId.Unset()
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CanonicalFactsOut) GetIpAddresses() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CanonicalFactsOut) GetIpAddressesOk() (*[]string, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return &o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *CanonicalFactsOut) HasIpAddresses() bool {
	if o != nil && o.IpAddresses != nil {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *CanonicalFactsOut) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetMacAddresses returns the MacAddresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CanonicalFactsOut) GetMacAddresses() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.MacAddresses
}

// GetMacAddressesOk returns a tuple with the MacAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CanonicalFactsOut) GetMacAddressesOk() (*[]string, bool) {
	if o == nil || o.MacAddresses == nil {
		return nil, false
	}
	return &o.MacAddresses, true
}

// HasMacAddresses returns a boolean if a field has been set.
func (o *CanonicalFactsOut) HasMacAddresses() bool {
	if o != nil && o.MacAddresses != nil {
		return true
	}

	return false
}

// SetMacAddresses gets a reference to the given []string and assigns it to the MacAddresses field.
func (o *CanonicalFactsOut) SetMacAddresses(v []string) {
	o.MacAddresses = v
}

// GetRhelMachineId returns the RhelMachineId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CanonicalFactsOut) GetRhelMachineId() string {
	if o == nil || o.RhelMachineId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RhelMachineId.Get()
}

// GetRhelMachineIdOk returns a tuple with the RhelMachineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CanonicalFactsOut) GetRhelMachineIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RhelMachineId.Get(), o.RhelMachineId.IsSet()
}

// HasRhelMachineId returns a boolean if a field has been set.
func (o *CanonicalFactsOut) HasRhelMachineId() bool {
	if o != nil && o.RhelMachineId.IsSet() {
		return true
	}

	return false
}

// SetRhelMachineId gets a reference to the given NullableString and assigns it to the RhelMachineId field.
func (o *CanonicalFactsOut) SetRhelMachineId(v string) {
	o.RhelMachineId.Set(&v)
}
// SetRhelMachineIdNil sets the value for RhelMachineId to be an explicit nil
func (o *CanonicalFactsOut) SetRhelMachineIdNil() {
	o.RhelMachineId.Set(nil)
}

// UnsetRhelMachineId ensures that no value is present for RhelMachineId, not even an explicit nil
func (o *CanonicalFactsOut) UnsetRhelMachineId() {
	o.RhelMachineId.Unset()
}

// GetSatelliteId returns the SatelliteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CanonicalFactsOut) GetSatelliteId() string {
	if o == nil || o.SatelliteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SatelliteId.Get()
}

// GetSatelliteIdOk returns a tuple with the SatelliteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CanonicalFactsOut) GetSatelliteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SatelliteId.Get(), o.SatelliteId.IsSet()
}

// HasSatelliteId returns a boolean if a field has been set.
func (o *CanonicalFactsOut) HasSatelliteId() bool {
	if o != nil && o.SatelliteId.IsSet() {
		return true
	}

	return false
}

// SetSatelliteId gets a reference to the given NullableString and assigns it to the SatelliteId field.
func (o *CanonicalFactsOut) SetSatelliteId(v string) {
	o.SatelliteId.Set(&v)
}
// SetSatelliteIdNil sets the value for SatelliteId to be an explicit nil
func (o *CanonicalFactsOut) SetSatelliteIdNil() {
	o.SatelliteId.Set(nil)
}

// UnsetSatelliteId ensures that no value is present for SatelliteId, not even an explicit nil
func (o *CanonicalFactsOut) UnsetSatelliteId() {
	o.SatelliteId.Unset()
}

// GetSubscriptionManagerId returns the SubscriptionManagerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CanonicalFactsOut) GetSubscriptionManagerId() string {
	if o == nil || o.SubscriptionManagerId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionManagerId.Get()
}

// GetSubscriptionManagerIdOk returns a tuple with the SubscriptionManagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CanonicalFactsOut) GetSubscriptionManagerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubscriptionManagerId.Get(), o.SubscriptionManagerId.IsSet()
}

// HasSubscriptionManagerId returns a boolean if a field has been set.
func (o *CanonicalFactsOut) HasSubscriptionManagerId() bool {
	if o != nil && o.SubscriptionManagerId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionManagerId gets a reference to the given NullableString and assigns it to the SubscriptionManagerId field.
func (o *CanonicalFactsOut) SetSubscriptionManagerId(v string) {
	o.SubscriptionManagerId.Set(&v)
}
// SetSubscriptionManagerIdNil sets the value for SubscriptionManagerId to be an explicit nil
func (o *CanonicalFactsOut) SetSubscriptionManagerIdNil() {
	o.SubscriptionManagerId.Set(nil)
}

// UnsetSubscriptionManagerId ensures that no value is present for SubscriptionManagerId, not even an explicit nil
func (o *CanonicalFactsOut) UnsetSubscriptionManagerId() {
	o.SubscriptionManagerId.Unset()
}

func (o CanonicalFactsOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BiosUuid.IsSet() {
		toSerialize["bios_uuid"] = o.BiosUuid.Get()
	}
	if o.ExternalId.IsSet() {
		toSerialize["external_id"] = o.ExternalId.Get()
	}
	if o.Fqdn.IsSet() {
		toSerialize["fqdn"] = o.Fqdn.Get()
	}
	if o.InsightsId.IsSet() {
		toSerialize["insights_id"] = o.InsightsId.Get()
	}
	if o.IpAddresses != nil {
		toSerialize["ip_addresses"] = o.IpAddresses
	}
	if o.MacAddresses != nil {
		toSerialize["mac_addresses"] = o.MacAddresses
	}
	if o.RhelMachineId.IsSet() {
		toSerialize["rhel_machine_id"] = o.RhelMachineId.Get()
	}
	if o.SatelliteId.IsSet() {
		toSerialize["satellite_id"] = o.SatelliteId.Get()
	}
	if o.SubscriptionManagerId.IsSet() {
		toSerialize["subscription_manager_id"] = o.SubscriptionManagerId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCanonicalFactsOut struct {
	value *CanonicalFactsOut
	isSet bool
}

func (v NullableCanonicalFactsOut) Get() *CanonicalFactsOut {
	return v.value
}

func (v *NullableCanonicalFactsOut) Set(val *CanonicalFactsOut) {
	v.value = val
	v.isSet = true
}

func (v NullableCanonicalFactsOut) IsSet() bool {
	return v.isSet
}

func (v *NullableCanonicalFactsOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCanonicalFactsOut(val *CanonicalFactsOut) *NullableCanonicalFactsOut {
	return &NullableCanonicalFactsOut{value: val, isSet: true}
}

func (v NullableCanonicalFactsOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCanonicalFactsOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


