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

// InlineObject1 struct for InlineObject1
type InlineObject1 struct {
	Name *string `json:"name,omitempty"`
}

// NewInlineObject1 instantiates a new InlineObject1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1() *InlineObject1 {
	this := InlineObject1{}
	return &this
}

// NewInlineObject1WithDefaults instantiates a new InlineObject1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1WithDefaults() *InlineObject1 {
	this := InlineObject1{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject1) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject1) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject1) SetName(v string) {
	o.Name = &v
}

func (o InlineObject1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1 struct {
	value *InlineObject1
	isSet bool
}

func (v NullableInlineObject1) Get() *InlineObject1 {
	return v.value
}

func (v *NullableInlineObject1) Set(val *InlineObject1) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1(val *InlineObject1) *NullableInlineObject1 {
	return &NullableInlineObject1{value: val, isSet: true}
}

func (v NullableInlineObject1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


