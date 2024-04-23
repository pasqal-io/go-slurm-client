/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.37
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slurmclient

import (
	"encoding/json"
)

// checks if the V0037Licenses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0037Licenses{}

// V0037Licenses struct for V0037Licenses
type V0037Licenses struct {
	// slurm errors
	Errors []V0037Error `json:"errors,omitempty"`
	// licenses info
	Licenses []V0037License `json:"licenses,omitempty"`
}

// NewV0037Licenses instantiates a new V0037Licenses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0037Licenses() *V0037Licenses {
	this := V0037Licenses{}
	return &this
}

// NewV0037LicensesWithDefaults instantiates a new V0037Licenses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0037LicensesWithDefaults() *V0037Licenses {
	this := V0037Licenses{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0037Licenses) GetErrors() []V0037Error {
	if o == nil || IsNil(o.Errors) {
		var ret []V0037Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Licenses) GetErrorsOk() ([]V0037Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0037Licenses) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0037Error and assigns it to the Errors field.
func (o *V0037Licenses) SetErrors(v []V0037Error) {
	o.Errors = v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *V0037Licenses) GetLicenses() []V0037License {
	if o == nil || IsNil(o.Licenses) {
		var ret []V0037License
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Licenses) GetLicensesOk() ([]V0037License, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *V0037Licenses) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []V0037License and assigns it to the Licenses field.
func (o *V0037Licenses) SetLicenses(v []V0037License) {
	o.Licenses = v
}

func (o V0037Licenses) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0037Licenses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	return toSerialize, nil
}

type NullableV0037Licenses struct {
	value *V0037Licenses
	isSet bool
}

func (v NullableV0037Licenses) Get() *V0037Licenses {
	return v.value
}

func (v *NullableV0037Licenses) Set(val *V0037Licenses) {
	v.value = val
	v.isSet = true
}

func (v NullableV0037Licenses) IsSet() bool {
	return v.isSet
}

func (v *NullableV0037Licenses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0037Licenses(val *V0037Licenses) *NullableV0037Licenses {
	return &NullableV0037Licenses{value: val, isSet: true}
}

func (v NullableV0037Licenses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0037Licenses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


