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

// CrossAccountRequest struct for CrossAccountRequest
type CrossAccountRequest struct {
	RequestId *string `json:"request_id,omitempty"`
	TargetAccount *string `json:"target_account,omitempty"`
	Status *string `json:"status,omitempty"`
	Created *string `json:"created,omitempty"`
	StartDate interface{} `json:"start_date,omitempty"`
	EndDate interface{} `json:"end_date,omitempty"`
}

// NewCrossAccountRequest instantiates a new CrossAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrossAccountRequest() *CrossAccountRequest {
	this := CrossAccountRequest{}
	return &this
}

// NewCrossAccountRequestWithDefaults instantiates a new CrossAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrossAccountRequestWithDefaults() *CrossAccountRequest {
	this := CrossAccountRequest{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CrossAccountRequest) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequest) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CrossAccountRequest) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CrossAccountRequest) SetRequestId(v string) {
	o.RequestId = &v
}

// GetTargetAccount returns the TargetAccount field value if set, zero value otherwise.
func (o *CrossAccountRequest) GetTargetAccount() string {
	if o == nil || o.TargetAccount == nil {
		var ret string
		return ret
	}
	return *o.TargetAccount
}

// GetTargetAccountOk returns a tuple with the TargetAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequest) GetTargetAccountOk() (*string, bool) {
	if o == nil || o.TargetAccount == nil {
		return nil, false
	}
	return o.TargetAccount, true
}

// HasTargetAccount returns a boolean if a field has been set.
func (o *CrossAccountRequest) HasTargetAccount() bool {
	if o != nil && o.TargetAccount != nil {
		return true
	}

	return false
}

// SetTargetAccount gets a reference to the given string and assigns it to the TargetAccount field.
func (o *CrossAccountRequest) SetTargetAccount(v string) {
	o.TargetAccount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CrossAccountRequest) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequest) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CrossAccountRequest) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CrossAccountRequest) SetStatus(v string) {
	o.Status = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CrossAccountRequest) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequest) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CrossAccountRequest) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *CrossAccountRequest) SetCreated(v string) {
	o.Created = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CrossAccountRequest) GetStartDate() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CrossAccountRequest) GetStartDateOk() (*interface{}, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CrossAccountRequest) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given interface{} and assigns it to the StartDate field.
func (o *CrossAccountRequest) SetStartDate(v interface{}) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CrossAccountRequest) GetEndDate() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CrossAccountRequest) GetEndDateOk() (*interface{}, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CrossAccountRequest) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given interface{} and assigns it to the EndDate field.
func (o *CrossAccountRequest) SetEndDate(v interface{}) {
	o.EndDate = v
}

func (o CrossAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	if o.TargetAccount != nil {
		toSerialize["target_account"] = o.TargetAccount
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	return json.Marshal(toSerialize)
}

type NullableCrossAccountRequest struct {
	value *CrossAccountRequest
	isSet bool
}

func (v NullableCrossAccountRequest) Get() *CrossAccountRequest {
	return v.value
}

func (v *NullableCrossAccountRequest) Set(val *CrossAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCrossAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCrossAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrossAccountRequest(val *CrossAccountRequest) *NullableCrossAccountRequest {
	return &NullableCrossAccountRequest{value: val, isSet: true}
}

func (v NullableCrossAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrossAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

