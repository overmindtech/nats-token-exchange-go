/*
NATS Token Exchange

Exchanges OAuth tokens for NATS tokens

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ResponsePermission ResponsePermission can be used to allow responses to any reply subject that is received on a valid subscription.
type ResponsePermission struct {
	Max *int64 `json:"max,omitempty"`
	// A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years.
	Ttl *int64 `json:"ttl,omitempty"`
}

// NewResponsePermission instantiates a new ResponsePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsePermission() *ResponsePermission {
	this := ResponsePermission{}
	return &this
}

// NewResponsePermissionWithDefaults instantiates a new ResponsePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsePermissionWithDefaults() *ResponsePermission {
	this := ResponsePermission{}
	return &this
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *ResponsePermission) GetMax() int64 {
	if o == nil || o.Max == nil {
		var ret int64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePermission) GetMaxOk() (*int64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *ResponsePermission) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given int64 and assigns it to the Max field.
func (o *ResponsePermission) SetMax(v int64) {
	o.Max = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *ResponsePermission) GetTtl() int64 {
	if o == nil || o.Ttl == nil {
		var ret int64
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsePermission) GetTtlOk() (*int64, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *ResponsePermission) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int64 and assigns it to the Ttl field.
func (o *ResponsePermission) SetTtl(v int64) {
	o.Ttl = &v
}

func (o ResponsePermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	return json.Marshal(toSerialize)
}

type NullableResponsePermission struct {
	value *ResponsePermission
	isSet bool
}

func (v NullableResponsePermission) Get() *ResponsePermission {
	return v.value
}

func (v *NullableResponsePermission) Set(val *ResponsePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsePermission(val *ResponsePermission) *NullableResponsePermission {
	return &NullableResponsePermission{value: val, isSet: true}
}

func (v NullableResponsePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


