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

// checks if the V0037Error type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0037Error{}

// V0037Error struct for V0037Error
type V0037Error struct {
	// error message
	Error *string `json:"error,omitempty"`
	// error number
	Errno *int32 `json:"errno,omitempty"`
}

// NewV0037Error instantiates a new V0037Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0037Error() *V0037Error {
	this := V0037Error{}
	return &this
}

// NewV0037ErrorWithDefaults instantiates a new V0037Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0037ErrorWithDefaults() *V0037Error {
	this := V0037Error{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *V0037Error) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Error) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *V0037Error) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *V0037Error) SetError(v string) {
	o.Error = &v
}

// GetErrno returns the Errno field value if set, zero value otherwise.
func (o *V0037Error) GetErrno() int32 {
	if o == nil || IsNil(o.Errno) {
		var ret int32
		return ret
	}
	return *o.Errno
}

// GetErrnoOk returns a tuple with the Errno field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Error) GetErrnoOk() (*int32, bool) {
	if o == nil || IsNil(o.Errno) {
		return nil, false
	}
	return o.Errno, true
}

// HasErrno returns a boolean if a field has been set.
func (o *V0037Error) HasErrno() bool {
	if o != nil && !IsNil(o.Errno) {
		return true
	}

	return false
}

// SetErrno gets a reference to the given int32 and assigns it to the Errno field.
func (o *V0037Error) SetErrno(v int32) {
	o.Errno = &v
}

func (o V0037Error) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0037Error) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Errno) {
		toSerialize["errno"] = o.Errno
	}
	return toSerialize, nil
}

type NullableV0037Error struct {
	value *V0037Error
	isSet bool
}

func (v NullableV0037Error) Get() *V0037Error {
	return v.value
}

func (v *NullableV0037Error) Set(val *V0037Error) {
	v.value = val
	v.isSet = true
}

func (v NullableV0037Error) IsSet() bool {
	return v.isSet
}

func (v *NullableV0037Error) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0037Error(val *V0037Error) *NullableV0037Error {
	return &NullableV0037Error{value: val, isSet: true}
}

func (v NullableV0037Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0037Error) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


