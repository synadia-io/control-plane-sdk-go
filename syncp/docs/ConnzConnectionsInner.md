# ConnzConnectionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**AuthorizedUser** | Pointer to **string** |  | [optional] 
**Cid** | **int32** |  | 
**Idle** | **string** |  | 
**InBytes** | **int64** |  | 
**InMsgs** | **int64** |  | 
**Ip** | **string** |  | 
**IssuerKey** | Pointer to **string** |  | [optional] 
**Jwt** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Lang** | Pointer to **string** |  | [optional] 
**LastActivity** | **time.Time** |  | 
**MqttClient** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NameTag** | Pointer to **string** |  | [optional] 
**OutBytes** | **int64** |  | 
**OutMsgs** | **int64** |  | 
**PendingBytes** | **int32** |  | 
**Port** | **int32** |  | 
**Reason** | Pointer to **string** |  | [optional] 
**Rtt** | Pointer to **string** |  | [optional] 
**Start** | **time.Time** |  | 
**Stop** | Pointer to **NullableString** |  | [optional] 
**Subscriptions** | **int32** |  | 
**SubscriptionsList** | Pointer to **[]string** |  | [optional] 
**SubscriptionsListDetail** | Pointer to [**[]SubDetail**](SubDetail.md) |  | [optional] 
**Tags** | Pointer to **[]string** | TagList is a unique array of lower case strings All tag list methods lower case the strings in the arguments | [optional] 
**TlsCipherSuite** | Pointer to **string** |  | [optional] 
**TlsPeerCerts** | Pointer to [**[]ConnInfoTlsPeerCertsInner**](ConnInfoTlsPeerCertsInner.md) |  | [optional] 
**TlsVersion** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Uptime** | **string** |  | 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewConnzConnectionsInner

`func NewConnzConnectionsInner(cid int32, idle string, inBytes int64, inMsgs int64, ip string, lastActivity time.Time, outBytes int64, outMsgs int64, pendingBytes int32, port int32, start time.Time, subscriptions int32, uptime string, ) *ConnzConnectionsInner`

NewConnzConnectionsInner instantiates a new ConnzConnectionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnzConnectionsInnerWithDefaults

`func NewConnzConnectionsInnerWithDefaults() *ConnzConnectionsInner`

NewConnzConnectionsInnerWithDefaults instantiates a new ConnzConnectionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ConnzConnectionsInner) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ConnzConnectionsInner) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ConnzConnectionsInner) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ConnzConnectionsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAuthorizedUser

`func (o *ConnzConnectionsInner) GetAuthorizedUser() string`

GetAuthorizedUser returns the AuthorizedUser field if non-nil, zero value otherwise.

### GetAuthorizedUserOk

`func (o *ConnzConnectionsInner) GetAuthorizedUserOk() (*string, bool)`

GetAuthorizedUserOk returns a tuple with the AuthorizedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedUser

`func (o *ConnzConnectionsInner) SetAuthorizedUser(v string)`

SetAuthorizedUser sets AuthorizedUser field to given value.

### HasAuthorizedUser

`func (o *ConnzConnectionsInner) HasAuthorizedUser() bool`

HasAuthorizedUser returns a boolean if a field has been set.

### GetCid

`func (o *ConnzConnectionsInner) GetCid() int32`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *ConnzConnectionsInner) GetCidOk() (*int32, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *ConnzConnectionsInner) SetCid(v int32)`

SetCid sets Cid field to given value.


### GetIdle

`func (o *ConnzConnectionsInner) GetIdle() string`

GetIdle returns the Idle field if non-nil, zero value otherwise.

### GetIdleOk

`func (o *ConnzConnectionsInner) GetIdleOk() (*string, bool)`

GetIdleOk returns a tuple with the Idle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdle

`func (o *ConnzConnectionsInner) SetIdle(v string)`

SetIdle sets Idle field to given value.


### GetInBytes

`func (o *ConnzConnectionsInner) GetInBytes() int64`

GetInBytes returns the InBytes field if non-nil, zero value otherwise.

### GetInBytesOk

`func (o *ConnzConnectionsInner) GetInBytesOk() (*int64, bool)`

GetInBytesOk returns a tuple with the InBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInBytes

`func (o *ConnzConnectionsInner) SetInBytes(v int64)`

SetInBytes sets InBytes field to given value.


### GetInMsgs

`func (o *ConnzConnectionsInner) GetInMsgs() int64`

GetInMsgs returns the InMsgs field if non-nil, zero value otherwise.

### GetInMsgsOk

`func (o *ConnzConnectionsInner) GetInMsgsOk() (*int64, bool)`

GetInMsgsOk returns a tuple with the InMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMsgs

`func (o *ConnzConnectionsInner) SetInMsgs(v int64)`

SetInMsgs sets InMsgs field to given value.


### GetIp

`func (o *ConnzConnectionsInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ConnzConnectionsInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ConnzConnectionsInner) SetIp(v string)`

SetIp sets Ip field to given value.


### GetIssuerKey

`func (o *ConnzConnectionsInner) GetIssuerKey() string`

GetIssuerKey returns the IssuerKey field if non-nil, zero value otherwise.

### GetIssuerKeyOk

`func (o *ConnzConnectionsInner) GetIssuerKeyOk() (*string, bool)`

GetIssuerKeyOk returns a tuple with the IssuerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerKey

`func (o *ConnzConnectionsInner) SetIssuerKey(v string)`

SetIssuerKey sets IssuerKey field to given value.

### HasIssuerKey

`func (o *ConnzConnectionsInner) HasIssuerKey() bool`

HasIssuerKey returns a boolean if a field has been set.

### GetJwt

`func (o *ConnzConnectionsInner) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *ConnzConnectionsInner) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *ConnzConnectionsInner) SetJwt(v string)`

SetJwt sets Jwt field to given value.

### HasJwt

`func (o *ConnzConnectionsInner) HasJwt() bool`

HasJwt returns a boolean if a field has been set.

### GetKind

`func (o *ConnzConnectionsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnzConnectionsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnzConnectionsInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ConnzConnectionsInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLang

`func (o *ConnzConnectionsInner) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *ConnzConnectionsInner) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *ConnzConnectionsInner) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *ConnzConnectionsInner) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetLastActivity

`func (o *ConnzConnectionsInner) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *ConnzConnectionsInner) GetLastActivityOk() (*time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *ConnzConnectionsInner) SetLastActivity(v time.Time)`

SetLastActivity sets LastActivity field to given value.


### GetMqttClient

`func (o *ConnzConnectionsInner) GetMqttClient() string`

GetMqttClient returns the MqttClient field if non-nil, zero value otherwise.

### GetMqttClientOk

`func (o *ConnzConnectionsInner) GetMqttClientOk() (*string, bool)`

GetMqttClientOk returns a tuple with the MqttClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttClient

`func (o *ConnzConnectionsInner) SetMqttClient(v string)`

SetMqttClient sets MqttClient field to given value.

### HasMqttClient

`func (o *ConnzConnectionsInner) HasMqttClient() bool`

HasMqttClient returns a boolean if a field has been set.

### GetName

`func (o *ConnzConnectionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnzConnectionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnzConnectionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnzConnectionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameTag

`func (o *ConnzConnectionsInner) GetNameTag() string`

GetNameTag returns the NameTag field if non-nil, zero value otherwise.

### GetNameTagOk

`func (o *ConnzConnectionsInner) GetNameTagOk() (*string, bool)`

GetNameTagOk returns a tuple with the NameTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameTag

`func (o *ConnzConnectionsInner) SetNameTag(v string)`

SetNameTag sets NameTag field to given value.

### HasNameTag

`func (o *ConnzConnectionsInner) HasNameTag() bool`

HasNameTag returns a boolean if a field has been set.

### GetOutBytes

`func (o *ConnzConnectionsInner) GetOutBytes() int64`

GetOutBytes returns the OutBytes field if non-nil, zero value otherwise.

### GetOutBytesOk

`func (o *ConnzConnectionsInner) GetOutBytesOk() (*int64, bool)`

GetOutBytesOk returns a tuple with the OutBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutBytes

`func (o *ConnzConnectionsInner) SetOutBytes(v int64)`

SetOutBytes sets OutBytes field to given value.


### GetOutMsgs

`func (o *ConnzConnectionsInner) GetOutMsgs() int64`

GetOutMsgs returns the OutMsgs field if non-nil, zero value otherwise.

### GetOutMsgsOk

`func (o *ConnzConnectionsInner) GetOutMsgsOk() (*int64, bool)`

GetOutMsgsOk returns a tuple with the OutMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutMsgs

`func (o *ConnzConnectionsInner) SetOutMsgs(v int64)`

SetOutMsgs sets OutMsgs field to given value.


### GetPendingBytes

`func (o *ConnzConnectionsInner) GetPendingBytes() int32`

GetPendingBytes returns the PendingBytes field if non-nil, zero value otherwise.

### GetPendingBytesOk

`func (o *ConnzConnectionsInner) GetPendingBytesOk() (*int32, bool)`

GetPendingBytesOk returns a tuple with the PendingBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingBytes

`func (o *ConnzConnectionsInner) SetPendingBytes(v int32)`

SetPendingBytes sets PendingBytes field to given value.


### GetPort

`func (o *ConnzConnectionsInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ConnzConnectionsInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ConnzConnectionsInner) SetPort(v int32)`

SetPort sets Port field to given value.


### GetReason

`func (o *ConnzConnectionsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ConnzConnectionsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ConnzConnectionsInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ConnzConnectionsInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRtt

`func (o *ConnzConnectionsInner) GetRtt() string`

GetRtt returns the Rtt field if non-nil, zero value otherwise.

### GetRttOk

`func (o *ConnzConnectionsInner) GetRttOk() (*string, bool)`

GetRttOk returns a tuple with the Rtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtt

`func (o *ConnzConnectionsInner) SetRtt(v string)`

SetRtt sets Rtt field to given value.

### HasRtt

`func (o *ConnzConnectionsInner) HasRtt() bool`

HasRtt returns a boolean if a field has been set.

### GetStart

`func (o *ConnzConnectionsInner) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ConnzConnectionsInner) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ConnzConnectionsInner) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetStop

`func (o *ConnzConnectionsInner) GetStop() string`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *ConnzConnectionsInner) GetStopOk() (*string, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *ConnzConnectionsInner) SetStop(v string)`

SetStop sets Stop field to given value.

### HasStop

`func (o *ConnzConnectionsInner) HasStop() bool`

HasStop returns a boolean if a field has been set.

### SetStopNil

`func (o *ConnzConnectionsInner) SetStopNil(b bool)`

 SetStopNil sets the value for Stop to be an explicit nil

### UnsetStop
`func (o *ConnzConnectionsInner) UnsetStop()`

UnsetStop ensures that no value is present for Stop, not even an explicit nil
### GetSubscriptions

`func (o *ConnzConnectionsInner) GetSubscriptions() int32`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *ConnzConnectionsInner) GetSubscriptionsOk() (*int32, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *ConnzConnectionsInner) SetSubscriptions(v int32)`

SetSubscriptions sets Subscriptions field to given value.


### GetSubscriptionsList

`func (o *ConnzConnectionsInner) GetSubscriptionsList() []string`

GetSubscriptionsList returns the SubscriptionsList field if non-nil, zero value otherwise.

### GetSubscriptionsListOk

`func (o *ConnzConnectionsInner) GetSubscriptionsListOk() (*[]string, bool)`

GetSubscriptionsListOk returns a tuple with the SubscriptionsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionsList

`func (o *ConnzConnectionsInner) SetSubscriptionsList(v []string)`

SetSubscriptionsList sets SubscriptionsList field to given value.

### HasSubscriptionsList

`func (o *ConnzConnectionsInner) HasSubscriptionsList() bool`

HasSubscriptionsList returns a boolean if a field has been set.

### GetSubscriptionsListDetail

`func (o *ConnzConnectionsInner) GetSubscriptionsListDetail() []SubDetail`

GetSubscriptionsListDetail returns the SubscriptionsListDetail field if non-nil, zero value otherwise.

### GetSubscriptionsListDetailOk

`func (o *ConnzConnectionsInner) GetSubscriptionsListDetailOk() (*[]SubDetail, bool)`

GetSubscriptionsListDetailOk returns a tuple with the SubscriptionsListDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionsListDetail

`func (o *ConnzConnectionsInner) SetSubscriptionsListDetail(v []SubDetail)`

SetSubscriptionsListDetail sets SubscriptionsListDetail field to given value.

### HasSubscriptionsListDetail

`func (o *ConnzConnectionsInner) HasSubscriptionsListDetail() bool`

HasSubscriptionsListDetail returns a boolean if a field has been set.

### GetTags

`func (o *ConnzConnectionsInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConnzConnectionsInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConnzConnectionsInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConnzConnectionsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ConnzConnectionsInner) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ConnzConnectionsInner) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTlsCipherSuite

`func (o *ConnzConnectionsInner) GetTlsCipherSuite() string`

GetTlsCipherSuite returns the TlsCipherSuite field if non-nil, zero value otherwise.

### GetTlsCipherSuiteOk

`func (o *ConnzConnectionsInner) GetTlsCipherSuiteOk() (*string, bool)`

GetTlsCipherSuiteOk returns a tuple with the TlsCipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCipherSuite

`func (o *ConnzConnectionsInner) SetTlsCipherSuite(v string)`

SetTlsCipherSuite sets TlsCipherSuite field to given value.

### HasTlsCipherSuite

`func (o *ConnzConnectionsInner) HasTlsCipherSuite() bool`

HasTlsCipherSuite returns a boolean if a field has been set.

### GetTlsPeerCerts

`func (o *ConnzConnectionsInner) GetTlsPeerCerts() []ConnInfoTlsPeerCertsInner`

GetTlsPeerCerts returns the TlsPeerCerts field if non-nil, zero value otherwise.

### GetTlsPeerCertsOk

`func (o *ConnzConnectionsInner) GetTlsPeerCertsOk() (*[]ConnInfoTlsPeerCertsInner, bool)`

GetTlsPeerCertsOk returns a tuple with the TlsPeerCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsPeerCerts

`func (o *ConnzConnectionsInner) SetTlsPeerCerts(v []ConnInfoTlsPeerCertsInner)`

SetTlsPeerCerts sets TlsPeerCerts field to given value.

### HasTlsPeerCerts

`func (o *ConnzConnectionsInner) HasTlsPeerCerts() bool`

HasTlsPeerCerts returns a boolean if a field has been set.

### GetTlsVersion

`func (o *ConnzConnectionsInner) GetTlsVersion() string`

GetTlsVersion returns the TlsVersion field if non-nil, zero value otherwise.

### GetTlsVersionOk

`func (o *ConnzConnectionsInner) GetTlsVersionOk() (*string, bool)`

GetTlsVersionOk returns a tuple with the TlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVersion

`func (o *ConnzConnectionsInner) SetTlsVersion(v string)`

SetTlsVersion sets TlsVersion field to given value.

### HasTlsVersion

`func (o *ConnzConnectionsInner) HasTlsVersion() bool`

HasTlsVersion returns a boolean if a field has been set.

### GetType

`func (o *ConnzConnectionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnzConnectionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnzConnectionsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConnzConnectionsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUptime

`func (o *ConnzConnectionsInner) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *ConnzConnectionsInner) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *ConnzConnectionsInner) SetUptime(v string)`

SetUptime sets Uptime field to given value.


### GetVersion

`func (o *ConnzConnectionsInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConnzConnectionsInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConnzConnectionsInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConnzConnectionsInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


