/*
WhatsApp API

API for managing WhatsApp devices and interactions

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetGroupMembers200ResponseInner struct for GetGroupMembers200ResponseInner
type GetGroupMembers200ResponseInner struct {
	// ID of the group member
	MemberId *string `json:"memberId,omitempty"`
	// Name of the group member
	MemberName *string `json:"memberName,omitempty"`
}

// NewGetGroupMembers200ResponseInner instantiates a new GetGroupMembers200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGroupMembers200ResponseInner() *GetGroupMembers200ResponseInner {
	this := GetGroupMembers200ResponseInner{}
	return &this
}

// NewGetGroupMembers200ResponseInnerWithDefaults instantiates a new GetGroupMembers200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGroupMembers200ResponseInnerWithDefaults() *GetGroupMembers200ResponseInner {
	this := GetGroupMembers200ResponseInner{}
	return &this
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
func (o *GetGroupMembers200ResponseInner) GetMemberId() string {
	if o == nil || o.MemberId == nil {
		var ret string
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGroupMembers200ResponseInner) GetMemberIdOk() (*string, bool) {
	if o == nil || o.MemberId == nil {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *GetGroupMembers200ResponseInner) HasMemberId() bool {
	if o != nil && o.MemberId != nil {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given string and assigns it to the MemberId field.
func (o *GetGroupMembers200ResponseInner) SetMemberId(v string) {
	o.MemberId = &v
}

// GetMemberName returns the MemberName field value if set, zero value otherwise.
func (o *GetGroupMembers200ResponseInner) GetMemberName() string {
	if o == nil || o.MemberName == nil {
		var ret string
		return ret
	}
	return *o.MemberName
}

// GetMemberNameOk returns a tuple with the MemberName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGroupMembers200ResponseInner) GetMemberNameOk() (*string, bool) {
	if o == nil || o.MemberName == nil {
		return nil, false
	}
	return o.MemberName, true
}

// HasMemberName returns a boolean if a field has been set.
func (o *GetGroupMembers200ResponseInner) HasMemberName() bool {
	if o != nil && o.MemberName != nil {
		return true
	}

	return false
}

// SetMemberName gets a reference to the given string and assigns it to the MemberName field.
func (o *GetGroupMembers200ResponseInner) SetMemberName(v string) {
	o.MemberName = &v
}

func (o GetGroupMembers200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MemberId != nil {
		toSerialize["memberId"] = o.MemberId
	}
	if o.MemberName != nil {
		toSerialize["memberName"] = o.MemberName
	}
	return json.Marshal(toSerialize)
}

type NullableGetGroupMembers200ResponseInner struct {
	value *GetGroupMembers200ResponseInner
	isSet bool
}

func (v NullableGetGroupMembers200ResponseInner) Get() *GetGroupMembers200ResponseInner {
	return v.value
}

func (v *NullableGetGroupMembers200ResponseInner) Set(val *GetGroupMembers200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGroupMembers200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGroupMembers200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGroupMembers200ResponseInner(val *GetGroupMembers200ResponseInner) *NullableGetGroupMembers200ResponseInner {
	return &NullableGetGroupMembers200ResponseInner{value: val, isSet: true}
}

func (v NullableGetGroupMembers200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGroupMembers200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


