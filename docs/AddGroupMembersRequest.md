# AddGroupMembersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumbers** | Pointer to **[]string** | List of phone numbers to be added as group members | [optional] 

## Methods

### NewAddGroupMembersRequest

`func NewAddGroupMembersRequest() *AddGroupMembersRequest`

NewAddGroupMembersRequest instantiates a new AddGroupMembersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGroupMembersRequestWithDefaults

`func NewAddGroupMembersRequestWithDefaults() *AddGroupMembersRequest`

NewAddGroupMembersRequestWithDefaults instantiates a new AddGroupMembersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumbers

`func (o *AddGroupMembersRequest) GetPhoneNumbers() []string`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *AddGroupMembersRequest) GetPhoneNumbersOk() (*[]string, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *AddGroupMembersRequest) SetPhoneNumbers(v []string)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *AddGroupMembersRequest) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


