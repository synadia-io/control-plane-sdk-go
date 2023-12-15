# Operator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
 | [**GenericFields**](GenericFields.md) |   | Embedded Struct
**AccountServerUrl** | Pointer to **string** | AccountServerURL is a partial URL like \&quot;https://host.domain.org:&lt;port&gt;/jwt/v1\&quot; tools will use the prefix and build queries by appending /accounts/&lt;account_id&gt; or /operator to the path provided. Note this assumes that the account server can handle requests in a account-server compatible way. See https://github.com/nats-io/account-server. | [optional] 
**AssertServerVersion** | Pointer to **string** | Min Server version | [optional] 
**OperatorServiceUrls** | Pointer to Nullable[[]**string**] | StringList is a wrapper for an array of strings | [optional] 
**SigningKeys** | Pointer to Nullable[[]**string**] | StringList is a wrapper for an array of strings | [optional] 
**StrictSigningKeyUsage** | Pointer to **bool** | Signing of subordinate objects will require signing keys | [optional] 
**SystemAccount** | Pointer to **string** | Identity of the system account | [optional] 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


