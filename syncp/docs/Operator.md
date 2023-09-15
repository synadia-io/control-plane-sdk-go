# Operator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenericFields** | [**GenericFields**](GenericFields.md) |  | 
**AccountServerUrl** | Pointer to **string** | AccountServerURL is a partial URL like \&quot;https://host.domain.org:&lt;port&gt;/jwt/v1\&quot; tools will use the prefix and build queries by appending /accounts/&lt;account_id&gt; or /operator to the path provided. Note this assumes that the account server can handle requests in a account-server compatible way. See https://github.com/nats-io/account-server. | [optional] 
**AssertServerVersion** | Pointer to **string** | Min Server version | [optional] 
**OperatorServiceUrls** | Pointer to **[]string** | StringList is a wrapper for an array of strings | [optional] 
**SigningKeys** | Pointer to **[]string** | StringList is a wrapper for an array of strings | [optional] 
**StrictSigningKeyUsage** | Pointer to **bool** | Signing of subordinate objects will require signing keys | [optional] 
**SystemAccount** | Pointer to **string** | Identity of the system account | [optional] 

## Methods

### NewOperator

`func NewOperator(genericFields GenericFields, ) *Operator`

NewOperator instantiates a new Operator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorWithDefaults

`func NewOperatorWithDefaults() *Operator`

NewOperatorWithDefaults instantiates a new Operator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenericFields

`func (o *Operator) GetGenericFields() GenericFields`

GetGenericFields returns the GenericFields field if non-nil, zero value otherwise.

### GetGenericFieldsOk

`func (o *Operator) GetGenericFieldsOk() (*GenericFields, bool)`

GetGenericFieldsOk returns a tuple with the GenericFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericFields

`func (o *Operator) SetGenericFields(v GenericFields)`

SetGenericFields sets GenericFields field to given value.


### GetAccountServerUrl

`func (o *Operator) GetAccountServerUrl() string`

GetAccountServerUrl returns the AccountServerUrl field if non-nil, zero value otherwise.

### GetAccountServerUrlOk

`func (o *Operator) GetAccountServerUrlOk() (*string, bool)`

GetAccountServerUrlOk returns a tuple with the AccountServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountServerUrl

`func (o *Operator) SetAccountServerUrl(v string)`

SetAccountServerUrl sets AccountServerUrl field to given value.

### HasAccountServerUrl

`func (o *Operator) HasAccountServerUrl() bool`

HasAccountServerUrl returns a boolean if a field has been set.

### GetAssertServerVersion

`func (o *Operator) GetAssertServerVersion() string`

GetAssertServerVersion returns the AssertServerVersion field if non-nil, zero value otherwise.

### GetAssertServerVersionOk

`func (o *Operator) GetAssertServerVersionOk() (*string, bool)`

GetAssertServerVersionOk returns a tuple with the AssertServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertServerVersion

`func (o *Operator) SetAssertServerVersion(v string)`

SetAssertServerVersion sets AssertServerVersion field to given value.

### HasAssertServerVersion

`func (o *Operator) HasAssertServerVersion() bool`

HasAssertServerVersion returns a boolean if a field has been set.

### GetOperatorServiceUrls

`func (o *Operator) GetOperatorServiceUrls() []string`

GetOperatorServiceUrls returns the OperatorServiceUrls field if non-nil, zero value otherwise.

### GetOperatorServiceUrlsOk

`func (o *Operator) GetOperatorServiceUrlsOk() (*[]string, bool)`

GetOperatorServiceUrlsOk returns a tuple with the OperatorServiceUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorServiceUrls

`func (o *Operator) SetOperatorServiceUrls(v []string)`

SetOperatorServiceUrls sets OperatorServiceUrls field to given value.

### HasOperatorServiceUrls

`func (o *Operator) HasOperatorServiceUrls() bool`

HasOperatorServiceUrls returns a boolean if a field has been set.

### SetOperatorServiceUrlsNil

`func (o *Operator) SetOperatorServiceUrlsNil(b bool)`

 SetOperatorServiceUrlsNil sets the value for OperatorServiceUrls to be an explicit nil

### UnsetOperatorServiceUrls
`func (o *Operator) UnsetOperatorServiceUrls()`

UnsetOperatorServiceUrls ensures that no value is present for OperatorServiceUrls, not even an explicit nil
### GetSigningKeys

`func (o *Operator) GetSigningKeys() []string`

GetSigningKeys returns the SigningKeys field if non-nil, zero value otherwise.

### GetSigningKeysOk

`func (o *Operator) GetSigningKeysOk() (*[]string, bool)`

GetSigningKeysOk returns a tuple with the SigningKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKeys

`func (o *Operator) SetSigningKeys(v []string)`

SetSigningKeys sets SigningKeys field to given value.

### HasSigningKeys

`func (o *Operator) HasSigningKeys() bool`

HasSigningKeys returns a boolean if a field has been set.

### SetSigningKeysNil

`func (o *Operator) SetSigningKeysNil(b bool)`

 SetSigningKeysNil sets the value for SigningKeys to be an explicit nil

### UnsetSigningKeys
`func (o *Operator) UnsetSigningKeys()`

UnsetSigningKeys ensures that no value is present for SigningKeys, not even an explicit nil
### GetStrictSigningKeyUsage

`func (o *Operator) GetStrictSigningKeyUsage() bool`

GetStrictSigningKeyUsage returns the StrictSigningKeyUsage field if non-nil, zero value otherwise.

### GetStrictSigningKeyUsageOk

`func (o *Operator) GetStrictSigningKeyUsageOk() (*bool, bool)`

GetStrictSigningKeyUsageOk returns a tuple with the StrictSigningKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictSigningKeyUsage

`func (o *Operator) SetStrictSigningKeyUsage(v bool)`

SetStrictSigningKeyUsage sets StrictSigningKeyUsage field to given value.

### HasStrictSigningKeyUsage

`func (o *Operator) HasStrictSigningKeyUsage() bool`

HasStrictSigningKeyUsage returns a boolean if a field has been set.

### GetSystemAccount

`func (o *Operator) GetSystemAccount() string`

GetSystemAccount returns the SystemAccount field if non-nil, zero value otherwise.

### GetSystemAccountOk

`func (o *Operator) GetSystemAccountOk() (*string, bool)`

GetSystemAccountOk returns a tuple with the SystemAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAccount

`func (o *Operator) SetSystemAccount(v string)`

SetSystemAccount sets SystemAccount field to given value.

### HasSystemAccount

`func (o *Operator) HasSystemAccount() bool`

HasSystemAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


