# Go API client for syncp

API for Synadia Control Plane Server


## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import "github.com/synadia-io/control-plane-sdk-go/syncp"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), syncp.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), syncp.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), syncp.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), syncp.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost/api/core/beta*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountAPI* | [**AssignAccountAppUser**](docs/AccountAPI.md#assignaccountappuser) | **Post** /accounts/{accountId}/app-users/{appUserId} | Assign App User to Account
*AccountAPI* | [**CreateAccountSkGroup**](docs/AccountAPI.md#createaccountskgroup) | **Post** /accounts/{accountId}/account-sk-groups | Create Account Signing Key Group
*AccountAPI* | [**CreateAlertRule**](docs/AccountAPI.md#createalertrule) | **Post** /accounts/{accountId}/alert-rules | Create Account Alert Rule
*AccountAPI* | [**CreateKvBucket**](docs/AccountAPI.md#createkvbucket) | **Post** /accounts/{accountId}/jetstream/kv-buckets | Create KV Bucket
*AccountAPI* | [**CreateMirror**](docs/AccountAPI.md#createmirror) | **Post** /accounts/{accountId}/jetstream/mirrors | Create Mirror
*AccountAPI* | [**CreateStream**](docs/AccountAPI.md#createstream) | **Post** /accounts/{accountId}/jetstream/streams | Create Stream
*AccountAPI* | [**CreateStreamExport**](docs/AccountAPI.md#createstreamexport) | **Post** /accounts/{accountId}/stream-exports | Create Stream Export
*AccountAPI* | [**CreateStreamImport**](docs/AccountAPI.md#createstreamimport) | **Post** /accounts/{accountId}/stream-imports | Create Stream Import
*AccountAPI* | [**CreateSubjectExport**](docs/AccountAPI.md#createsubjectexport) | **Post** /accounts/{accountId}/subject-exports | Create Subject Export
*AccountAPI* | [**CreateSubjectImport**](docs/AccountAPI.md#createsubjectimport) | **Post** /accounts/{accountId}/subject-imports | Create Subject Import
*AccountAPI* | [**CreateUser**](docs/AccountAPI.md#createuser) | **Post** /accounts/{accountId}/nats-users | Create NATS User
*AccountAPI* | [**DeleteAccount**](docs/AccountAPI.md#deleteaccount) | **Delete** /accounts/{accountId} | Delete Account
*AccountAPI* | [**DeleteAlertRule**](docs/AccountAPI.md#deletealertrule) | **Delete** /accounts/{accountId}/alert-rules/{alertRuleId} | Delete Account Alert Rule
*AccountAPI* | [**GetAccount**](docs/AccountAPI.md#getaccount) | **Get** /accounts/{accountId} | Get Account
*AccountAPI* | [**GetAccountInfo**](docs/AccountAPI.md#getaccountinfo) | **Get** /accounts/{accountId}/info | Get Account Info
*AccountAPI* | [**GetAccountMetrics**](docs/AccountAPI.md#getaccountmetrics) | **Get** /accounts/{accountId}/metrics | Get Account Metrics
*AccountAPI* | [**GetAlertRule**](docs/AccountAPI.md#getalertrule) | **Get** /accounts/{accountId}/alert-rules/{alertRuleId} | Get Account Alert Rule
*AccountAPI* | [**ListAccountAppUsers**](docs/AccountAPI.md#listaccountappusers) | **Get** /accounts/{accountId}/app-users | List App Users
*AccountAPI* | [**ListAccountConnections**](docs/AccountAPI.md#listaccountconnections) | **Get** /accounts/{accountId}/connections | List Account Connections
*AccountAPI* | [**ListAccountSkGroup**](docs/AccountAPI.md#listaccountskgroup) | **Get** /accounts/{accountId}/account-sk-groups | List Account Signing Key Groups
*AccountAPI* | [**ListAlertRules**](docs/AccountAPI.md#listalertrules) | **Get** /accounts/{accountId}/alert-rules | List Account Alert Rules
*AccountAPI* | [**ListJetStreamAssets**](docs/AccountAPI.md#listjetstreamassets) | **Get** /accounts/{accountId}/jetstream | List JetStream Assets
*AccountAPI* | [**ListKvBuckets**](docs/AccountAPI.md#listkvbuckets) | **Get** /accounts/{accountId}/jetstream/kv-buckets | List KV buckets
*AccountAPI* | [**ListMirrors**](docs/AccountAPI.md#listmirrors) | **Get** /accounts/{accountId}/jetstream/mirrors | List Mirrors
*AccountAPI* | [**ListStreamExports**](docs/AccountAPI.md#liststreamexports) | **Get** /accounts/{accountId}/stream-exports | List Stream Exports
*AccountAPI* | [**ListStreamExportsShared**](docs/AccountAPI.md#liststreamexportsshared) | **Get** /accounts/{accountId}/stream-imports/shared | List Shared Stream Exports
*AccountAPI* | [**ListStreamImports**](docs/AccountAPI.md#liststreamimports) | **Get** /accounts/{accountId}/stream-imports | List Stream Imports
*AccountAPI* | [**ListStreams**](docs/AccountAPI.md#liststreams) | **Get** /accounts/{accountId}/jetstream/streams | List Streams
*AccountAPI* | [**ListSubjectExports**](docs/AccountAPI.md#listsubjectexports) | **Get** /accounts/{accountId}/subject-exports | List Subject Exports
*AccountAPI* | [**ListSubjectExportsShared**](docs/AccountAPI.md#listsubjectexportsshared) | **Get** /accounts/{accountId}/subject-imports/shared | List Shared Subject Exports
*AccountAPI* | [**ListSubjectImports**](docs/AccountAPI.md#listsubjectimports) | **Get** /accounts/{accountId}/subject-imports | List Subject Imports
*AccountAPI* | [**ListUsers**](docs/AccountAPI.md#listusers) | **Get** /accounts/{accountId}/nats-users | List NATS Users
*AccountAPI* | [**RunAlertRule**](docs/AccountAPI.md#runalertrule) | **Get** /accounts/{accountId}/alert-rules/{alertRuleId}/run | Run Account Alert Rule
*AccountAPI* | [**UnAssignAccountAppUser**](docs/AccountAPI.md#unassignaccountappuser) | **Delete** /accounts/{accountId}/app-users/{appUserId} | Unassign App User from Account
*AccountAPI* | [**UpdateAccount**](docs/AccountAPI.md#updateaccount) | **Patch** /accounts/{accountId} | Update Account
*AccountAPI* | [**UpdateAlertRule**](docs/AccountAPI.md#updatealertrule) | **Patch** /accounts/{accountId}/alert-rules/{alertRuleId} | Update Account Alert Rule
*AgentTokenAPI* | [**DeleteAgentToken**](docs/AgentTokenAPI.md#deleteagenttoken) | **Delete** /agent-tokens/{tokenId} | Delete Agent Token
*AlertAPI* | [**AcknowledgeAlert**](docs/AlertAPI.md#acknowledgealert) | **Patch** /alerts/{alertId}/acknowledge | Acknowledge Alert
*AlertAPI* | [**GetAlert**](docs/AlertAPI.md#getalert) | **Get** /alerts/{alertId} | Get Alert
*AppUserAPI* | [**DeleteAppUser**](docs/AppUserAPI.md#deleteappuser) | **Delete** /app-users/{appUserId} | Delete App User
*AppUserAPI* | [**GetAppUser**](docs/AppUserAPI.md#getappuser) | **Get** /app-users/{appUserId} | Get App User
*AppUserAPI* | [**ListAppUserRoles**](docs/AppUserAPI.md#listappuserroles) | **Get** /app-users/{appUserId}/roles | List Roles
*AppUserAPI* | [**UpdateAppUser**](docs/AppUserAPI.md#updateappuser) | **Patch** /app-users/{appUserId} | Update App User
*AuthzAPI* | [**Check**](docs/AuthzAPI.md#check) | **Post** /authz/check | Check Authz Decisions
*AuthzAPI* | [**ListPolicies**](docs/AuthzAPI.md#listpolicies) | **Get** /authz/policies | Get Policy List
*AuthzAPI* | [**ListRoles**](docs/AuthzAPI.md#listroles) | **Get** /authz/roles | Get Authz roles List
*KvBucketAPI* | [**CreateKvPullConsumer**](docs/KvBucketAPI.md#createkvpullconsumer) | **Post** /jetstream/kv-bucket/{streamId}/consumers/pull | Create Pull Consumer
*KvBucketAPI* | [**CreateKvPushConsumer**](docs/KvBucketAPI.md#createkvpushconsumer) | **Post** /jetstream/kv-bucket/{streamId}/consumers/push | Create Push Consumer
*KvBucketAPI* | [**DeleteBucket**](docs/KvBucketAPI.md#deletebucket) | **Delete** /jetstream/kv-bucket/{streamId} | Delete KV Bucket
*KvBucketAPI* | [**GetKvBucket**](docs/KvBucketAPI.md#getkvbucket) | **Get** /jetstream/kv-bucket/{streamId} | Get KV Bucket
*KvBucketAPI* | [**ListKvConsumers**](docs/KvBucketAPI.md#listkvconsumers) | **Get** /jetstream/kv-bucket/{streamId}/consumers | List Consumers
*KvBucketAPI* | [**UpdateJSAsset**](docs/KvBucketAPI.md#updatejsasset) | **Patch** /jetstream/kv-bucket/{streamId} | Update KV Bucket
*MirrorAPI* | [**CreateMirrorPullConsumer**](docs/MirrorAPI.md#createmirrorpullconsumer) | **Post** /jetstream/mirror/{streamId}/consumers/pull | Create Pull Consumer
*MirrorAPI* | [**CreateMirrorPushConsumer**](docs/MirrorAPI.md#createmirrorpushconsumer) | **Post** /jetstream/mirror/{streamId}/consumers/push | Create Push consumer
*MirrorAPI* | [**DeleteMirror**](docs/MirrorAPI.md#deletemirror) | **Delete** /jetstream/mirror/{streamId} | Delete Mirror
*MirrorAPI* | [**GetMirror**](docs/MirrorAPI.md#getmirror) | **Get** /jetstream/mirror/{streamId} | Get Mirror
*MirrorAPI* | [**ListMirrorConsumers**](docs/MirrorAPI.md#listmirrorconsumers) | **Get** /jetstream/mirror/{streamId}/consumers | List Consumers
*MirrorAPI* | [**UpdateMirror**](docs/MirrorAPI.md#updatemirror) | **Patch** /jetstream/mirror/{streamId} | Update Mirror
*NatsUserAPI* | [**AssignNatsUserAppUser**](docs/NatsUserAPI.md#assignnatsuserappuser) | **Post** /nats-users/{userId}/app-users/{appUserId} | Assign App User to NATS User
*NatsUserAPI* | [**DeleteNatsUser**](docs/NatsUserAPI.md#deletenatsuser) | **Delete** /nats-users/{userId} | Delete NATS User
*NatsUserAPI* | [**DownloadNatsUserCreds**](docs/NatsUserAPI.md#downloadnatsusercreds) | **Post** /nats-users/{userId}/creds | Get Creds
*NatsUserAPI* | [**GetNatsUser**](docs/NatsUserAPI.md#getnatsuser) | **Get** /nats-users/{userId} | Get NATS User
*NatsUserAPI* | [**ListNatsUserAppUsers**](docs/NatsUserAPI.md#listnatsuserappusers) | **Get** /nats-users/{userId}/app-users | List App Users
*NatsUserAPI* | [**ListNatsUserConnections**](docs/NatsUserAPI.md#listnatsuserconnections) | **Get** /nats-users/{userId}/connections | List NATs User Connections
*NatsUserAPI* | [**UnAssignNatsUserAppUser**](docs/NatsUserAPI.md#unassignnatsuserappuser) | **Delete** /nats-users/{userId}/app-users/{appUserId} | Unassign App User from NATS User
*NatsUserAPI* | [**UpdateNatsUser**](docs/NatsUserAPI.md#updatenatsuser) | **Patch** /nats-users/{userId} | Update NATS User
*PersonalAccessTokenAPI* | [**DeletePersonalAccessToken**](docs/PersonalAccessTokenAPI.md#deletepersonalaccesstoken) | **Delete** /personal-access-tokens/{tokenId} | Delete Personal Access Token
*PersonalAccessTokenAPI* | [**GetPersonalAccessToken**](docs/PersonalAccessTokenAPI.md#getpersonalaccesstoken) | **Get** /personal-access-tokens/{tokenId} | Get Personal Access Token
*PersonalAccessTokenAPI* | [**UpdatePersonalAccessToken**](docs/PersonalAccessTokenAPI.md#updatepersonalaccesstoken) | **Patch** /personal-access-tokens/{tokenId} | Update Personal Access Token
*PullConsumerAPI* | [**DeletePullConsumer**](docs/PullConsumerAPI.md#deletepullconsumer) | **Delete** /consumers/pull/{consumerId} | Delete Pull Consumer
*PullConsumerAPI* | [**GetPullConsumerInfo**](docs/PullConsumerAPI.md#getpullconsumerinfo) | **Get** /consumers/pull/{consumerId} | Get Pull Consumer
*PullConsumerAPI* | [**UpdatePullConsumer**](docs/PullConsumerAPI.md#updatepullconsumer) | **Patch** /consumers/pull/{consumerId} | Update Pull Consumer
*PushConsumerAPI* | [**DeletePushConsumer**](docs/PushConsumerAPI.md#deletepushconsumer) | **Delete** /consumers/push/{consumerId} | Delete Push Consumer
*PushConsumerAPI* | [**GetPushConsumerInfo**](docs/PushConsumerAPI.md#getpushconsumerinfo) | **Get** /consumers/push/{consumerId} | Get Push Consumer
*PushConsumerAPI* | [**UpdatePushConsumer**](docs/PushConsumerAPI.md#updatepushconsumer) | **Patch** /consumers/push/{consumerId} | Update Push Consumer
*SessionAPI* | [**CreateAppUser**](docs/SessionAPI.md#createappuser) | **Post** /app-users | Create App User
*SessionAPI* | [**CreatePersonalAccessToken**](docs/SessionAPI.md#createpersonalaccesstoken) | **Post** /personal-access-tokens | Create Personal Access Token
*SessionAPI* | [**CreateSystem**](docs/SessionAPI.md#createsystem) | **Post** /systems | Create System
*SessionAPI* | [**GetVersion**](docs/SessionAPI.md#getversion) | **Get** /version | Get Version
*SessionAPI* | [**ImportSystem**](docs/SessionAPI.md#importsystem) | **Post** /import-system | Import a System
*SessionAPI* | [**ListAlerts**](docs/SessionAPI.md#listalerts) | **Get** /alerts | List Alerts
*SessionAPI* | [**ListAppUsers**](docs/SessionAPI.md#listappusers) | **Get** /app-users | List App Users
*SessionAPI* | [**ListPersonalAccessTokens**](docs/SessionAPI.md#listpersonalaccesstokens) | **Get** /personal-access-tokens | List Personal Access Tokens
*SessionAPI* | [**ListSessionAccounts**](docs/SessionAPI.md#listsessionaccounts) | **Get** /accounts | List Accounts
*SessionAPI* | [**ListSessionNatsUsers**](docs/SessionAPI.md#listsessionnatsusers) | **Get** /nats-users | List NATS Users
*SessionAPI* | [**ListSystems**](docs/SessionAPI.md#listsystems) | **Get** /systems | List Systems
*SessionAPI* | [**SearchAppUsers**](docs/SessionAPI.md#searchappusers) | **Get** /search/app-users | Search App Users
*SessionAPI* | [**SearchSystemAccounts**](docs/SessionAPI.md#searchsystemaccounts) | **Get** /search/systems/{systemId}/accounts | Search System Accounts
*SessionAPI* | [**SearchSystemServers**](docs/SessionAPI.md#searchsystemservers) | **Get** /search/systems/{systemId}/servers | Search System Servers
*SigKeyAPI* | [**DeleteAccountSk**](docs/SigKeyAPI.md#deleteaccountsk) | **Delete** /account-sks/{keyId} | Delete Account Signing
*SigKeyAPI* | [**GetAccountSk**](docs/SigKeyAPI.md#getaccountsk) | **Get** /account-sks/{keyId} | Get Account Signing
*SigKeyAPI* | [**UpdateAccountSk**](docs/SigKeyAPI.md#updateaccountsk) | **Patch** /account-sks/{keyId} | Update Account Signing
*SigKeyGroupAPI* | [**DeleteAccountSkGroup**](docs/SigKeyGroupAPI.md#deleteaccountskgroup) | **Delete** /account-sk-groups/{groupId} | Delete Account Signing Key Group
*SigKeyGroupAPI* | [**GetAccountSkGroup**](docs/SigKeyGroupAPI.md#getaccountskgroup) | **Get** /account-sk-groups/{groupId} | Get Account Signing Key Group
*SigKeyGroupAPI* | [**ListAccountSkGroupKeys**](docs/SigKeyGroupAPI.md#listaccountskgroupkeys) | **Get** /account-sk-groups/{groupId}/account-sks | List Signing Keys
*SigKeyGroupAPI* | [**RotateAccountSk**](docs/SigKeyGroupAPI.md#rotateaccountsk) | **Post** /account-sk-groups/{groupId}/rotate-sk | Roate Active Signing Key
*SigKeyGroupAPI* | [**UpdateAccountSkGroup**](docs/SigKeyGroupAPI.md#updateaccountskgroup) | **Patch** /account-sk-groups/{groupId} | Update Account Signing Key Group
*StreamAPI* | [**CreatePullConsumer**](docs/StreamAPI.md#createpullconsumer) | **Post** /jetstream/stream/{streamId}/consumers/pull | Create Pull Consumer
*StreamAPI* | [**CreatePushConsumer**](docs/StreamAPI.md#createpushconsumer) | **Post** /jetstream/stream/{streamId}/consumers/push | Create Push Consumer
*StreamAPI* | [**DeleteStream**](docs/StreamAPI.md#deletestream) | **Delete** /jetstream/stream/{streamId} | Delete Stream
*StreamAPI* | [**GetStreamInfo**](docs/StreamAPI.md#getstreaminfo) | **Get** /jetstream/stream/{streamId} | Get Stream
*StreamAPI* | [**ListConsumers**](docs/StreamAPI.md#listconsumers) | **Get** /jetstream/stream/{streamId}/consumers | List Consumers
*StreamAPI* | [**UpdateStream**](docs/StreamAPI.md#updatestream) | **Patch** /jetstream/stream/{streamId} | Update Stream
*StreamExportAPI* | [**CreateStreamShares**](docs/StreamExportAPI.md#createstreamshares) | **Post** /stream-exports/{streamExportId}/shares | Create Stream Shares
*StreamExportAPI* | [**DeleteStreamExport**](docs/StreamExportAPI.md#deletestreamexport) | **Delete** /stream-exports/{streamExportId} | Delete Stream Export
*StreamExportAPI* | [**DeleteStreamShare**](docs/StreamExportAPI.md#deletestreamshare) | **Delete** /stream-exports/{streamExportId}/shares/{targetAccountNKeyPublic} | Delete Stream Share
*StreamExportAPI* | [**GetStreamExport**](docs/StreamExportAPI.md#getstreamexport) | **Get** /stream-exports/{streamExportId} | Get Stream Export
*StreamExportAPI* | [**ListStreamShares**](docs/StreamExportAPI.md#liststreamshares) | **Get** /stream-exports/{streamExportId}/shares | List Stream Shares
*StreamImportAPI* | [**DeleteStreamImport**](docs/StreamImportAPI.md#deletestreamimport) | **Delete** /stream-imports/{streamImportId} | Delete Stream Import
*StreamImportAPI* | [**GetStreamImport**](docs/StreamImportAPI.md#getstreamimport) | **Get** /stream-imports/{streamImportId} | Get Stream Import
*SubjectExportAPI* | [**CreateSubjectShares**](docs/SubjectExportAPI.md#createsubjectshares) | **Post** /subject-exports/{subjectExportId}/shares | Create Subject Shares
*SubjectExportAPI* | [**DeleteSubjectExport**](docs/SubjectExportAPI.md#deletesubjectexport) | **Delete** /subject-exports/{subjectExportId} | Delete Subject Export
*SubjectExportAPI* | [**DeleteSubjectShare**](docs/SubjectExportAPI.md#deletesubjectshare) | **Delete** /subject-exports/{subjectExportId}/shares/{targetAccountNKeyPublic} | Delete Subject Share
*SubjectExportAPI* | [**GetSubjectExport**](docs/SubjectExportAPI.md#getsubjectexport) | **Get** /subject-exports/{subjectExportId} | Get Subject Export
*SubjectExportAPI* | [**ListSubjectShares**](docs/SubjectExportAPI.md#listsubjectshares) | **Get** /subject-exports/{subjectExportId}/shares | List Subject Shares
*SubjectExportAPI* | [**UpdateSubjectExport**](docs/SubjectExportAPI.md#updatesubjectexport) | **Patch** /subject-exports/{subjectExportId} | Update Subject Export
*SubjectImportAPI* | [**DeleteSubjectImport**](docs/SubjectImportAPI.md#deletesubjectimport) | **Delete** /subject-imports/{subjectImportId} | Delete Subject Import
*SubjectImportAPI* | [**GetSubjectImport**](docs/SubjectImportAPI.md#getsubjectimport) | **Get** /subject-imports/{subjectImportId} | Get Subject Import
*SubjectImportAPI* | [**UpdateSubjectImport**](docs/SubjectImportAPI.md#updatesubjectimport) | **Patch** /subject-imports/{subjectImportId} | Update Subject Import
*SystemAPI* | [**AssignSystemAppUser**](docs/SystemAPI.md#assignsystemappuser) | **Post** /systems/{systemId}/app-users/{appUserId} | Assign App User to System
*SystemAPI* | [**CreateAccount**](docs/SystemAPI.md#createaccount) | **Post** /systems/{systemId}/accounts | Create Account
*SystemAPI* | [**CreateSystemAlertRule**](docs/SystemAPI.md#createsystemalertrule) | **Post** /systems/{systemId}/alert-rules | Create System Alert Rule
*SystemAPI* | [**DeleteSystem**](docs/SystemAPI.md#deletesystem) | **Delete** /systems/{systemId} | Delete System
*SystemAPI* | [**DeleteSystemAlertRule**](docs/SystemAPI.md#deletesystemalertrule) | **Delete** /systems/{systemId}/alert-rules/{alertRuleId} | Delete System Alert Rule
*SystemAPI* | [**GetCurrentAgentToken**](docs/SystemAPI.md#getcurrentagenttoken) | **Get** /systems/{systemId}/agent-tokens/current | Get Current Agent Token
*SystemAPI* | [**GetSystem**](docs/SystemAPI.md#getsystem) | **Get** /systems/{systemId} | Get System
*SystemAPI* | [**GetSystemAlertRule**](docs/SystemAPI.md#getsystemalertrule) | **Get** /systems/{systemId}/alert-rules/{alertRuleId} | Get System Alert Rule
*SystemAPI* | [**ImportAccount**](docs/SystemAPI.md#importaccount) | **Post** /systems/{systemId}/import-account | Import Account
*SystemAPI* | [**ImportUser**](docs/SystemAPI.md#importuser) | **Post** /systems/{systemId}/import-user | Import User
*SystemAPI* | [**ListAccounts**](docs/SystemAPI.md#listaccounts) | **Get** /systems/{systemId}/accounts | List Accounts
*SystemAPI* | [**ListAgentTokens**](docs/SystemAPI.md#listagenttokens) | **Get** /systems/{systemId}/agent-tokens | List Agent Tokens
*SystemAPI* | [**ListClusters**](docs/SystemAPI.md#listclusters) | **Get** /systems/{systemId}/nats-clusters | List Clusters
*SystemAPI* | [**ListConnections**](docs/SystemAPI.md#listconnections) | **Get** /systems/{systemId}/connections | List Connections
*SystemAPI* | [**ListServers**](docs/SystemAPI.md#listservers) | **Get** /systems/{systemId}/servers | List Servers
*SystemAPI* | [**ListSystemAlertRules**](docs/SystemAPI.md#listsystemalertrules) | **Get** /systems/{systemId}/alert-rules | List System Alert Rules
*SystemAPI* | [**ListSystemAppUsers**](docs/SystemAPI.md#listsystemappusers) | **Get** /systems/{systemId}/app-users | List App Users
*SystemAPI* | [**RotateAgentToken**](docs/SystemAPI.md#rotateagenttoken) | **Post** /systems/{systemId}/agent-tokens | Rotate Agent Token
*SystemAPI* | [**RunSystemAlertRule**](docs/SystemAPI.md#runsystemalertrule) | **Get** /systems/{systemId}/alert-rules/{alertRuleId}/run | Run System Alert Rule
*SystemAPI* | [**UnAssignSystemAppUser**](docs/SystemAPI.md#unassignsystemappuser) | **Delete** /systems/{systemId}/app-users/{appUserId} | Unassign App User from System
*SystemAPI* | [**UpdateSystem**](docs/SystemAPI.md#updatesystem) | **Patch** /systems/{systemId} | Update System
*SystemAPI* | [**UpdateSystemAlertRule**](docs/SystemAPI.md#updatesystemalertrule) | **Patch** /systems/{systemId}/alert-rules/{alertRuleId} | Update System Alert Rules


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountChartType](docs/AccountChartType.md)
 - [AccountClaims](docs/AccountClaims.md)
 - [AccountClaimsInfo](docs/AccountClaimsInfo.md)
 - [AccountConnectionsListResponse](docs/AccountConnectionsListResponse.md)
 - [AccountCreateRequest](docs/AccountCreateRequest.md)
 - [AccountCreateRequestJwtSettings](docs/AccountCreateRequestJwtSettings.md)
 - [AccountInfo](docs/AccountInfo.md)
 - [AccountJWTSettings](docs/AccountJWTSettings.md)
 - [AccountLimits](docs/AccountLimits.md)
 - [AccountListResponse](docs/AccountListResponse.md)
 - [AccountMetrics](docs/AccountMetrics.md)
 - [AccountSearch](docs/AccountSearch.md)
 - [AccountSearchListResponse](docs/AccountSearchListResponse.md)
 - [AccountUpdateRequest](docs/AccountUpdateRequest.md)
 - [AccountViewResponse](docs/AccountViewResponse.md)
 - [AckPolicy](docs/AckPolicy.md)
 - [Activation](docs/Activation.md)
 - [ActivationClaims](docs/ActivationClaims.md)
 - [AgentTokenCurrentResponse](docs/AgentTokenCurrentResponse.md)
 - [AgentTokenListResponse](docs/AgentTokenListResponse.md)
 - [AgentTokenRotateResponse](docs/AgentTokenRotateResponse.md)
 - [AgentTokenViewResponse](docs/AgentTokenViewResponse.md)
 - [AlertListResponse](docs/AlertListResponse.md)
 - [AlertRuleAccountCreateRequest](docs/AlertRuleAccountCreateRequest.md)
 - [AlertRuleBaseCreateRequest](docs/AlertRuleBaseCreateRequest.md)
 - [AlertRuleListResponse](docs/AlertRuleListResponse.md)
 - [AlertRuleOperator](docs/AlertRuleOperator.md)
 - [AlertRuleSeverity](docs/AlertRuleSeverity.md)
 - [AlertRuleType](docs/AlertRuleType.md)
 - [AlertRuleUpdateRequest](docs/AlertRuleUpdateRequest.md)
 - [AlertRuleViewResponse](docs/AlertRuleViewResponse.md)
 - [AlertStatus](docs/AlertStatus.md)
 - [AlertViewResponse](docs/AlertViewResponse.md)
 - [AlertViewResponseAccount](docs/AlertViewResponseAccount.md)
 - [AlertViewResponseSystem](docs/AlertViewResponseSystem.md)
 - [ApexCoordinate](docs/ApexCoordinate.md)
 - [ApexSeries](docs/ApexSeries.md)
 - [AppPolicy](docs/AppPolicy.md)
 - [AppPolicyStatement](docs/AppPolicyStatement.md)
 - [AppRole](docs/AppRole.md)
 - [AppRoleAction](docs/AppRoleAction.md)
 - [AppRoleEffect](docs/AppRoleEffect.md)
 - [AppRoleScope](docs/AppRoleScope.md)
 - [AppUserAccessTokenCreateRequest](docs/AppUserAccessTokenCreateRequest.md)
 - [AppUserAccessTokenCreateResponse](docs/AppUserAccessTokenCreateResponse.md)
 - [AppUserAccessTokenListResponse](docs/AppUserAccessTokenListResponse.md)
 - [AppUserAccessTokenUpdateRequest](docs/AppUserAccessTokenUpdateRequest.md)
 - [AppUserAccessTokenViewResponse](docs/AppUserAccessTokenViewResponse.md)
 - [AppUserAssignListResponse](docs/AppUserAssignListResponse.md)
 - [AppUserAssignRequest](docs/AppUserAssignRequest.md)
 - [AppUserAssignResponse](docs/AppUserAssignResponse.md)
 - [AppUserCreateRequest](docs/AppUserCreateRequest.md)
 - [AppUserCreateResponse](docs/AppUserCreateResponse.md)
 - [AppUserInfo](docs/AppUserInfo.md)
 - [AppUserInfoListResponse](docs/AppUserInfoListResponse.md)
 - [AppUserListResponse](docs/AppUserListResponse.md)
 - [AppUserType](docs/AppUserType.md)
 - [AppUserUpdateRequest](docs/AppUserUpdateRequest.md)
 - [AppUserUpdateResponse](docs/AppUserUpdateResponse.md)
 - [AppUserViewResponse](docs/AppUserViewResponse.md)
 - [Authorization](docs/Authorization.md)
 - [AuthzRequest](docs/AuthzRequest.md)
 - [AuthzResponse](docs/AuthzResponse.md)
 - [ClaimsData](docs/ClaimsData.md)
 - [ClusterChartType](docs/ClusterChartType.md)
 - [ClusterInfo](docs/ClusterInfo.md)
 - [ClusterInfoReplicasInner](docs/ClusterInfoReplicasInner.md)
 - [ConnInfo](docs/ConnInfo.md)
 - [ConnInfoTlsPeerCertsInner](docs/ConnInfoTlsPeerCertsInner.md)
 - [Connz](docs/Connz.md)
 - [ConnzConnectionsInner](docs/ConnzConnectionsInner.md)
 - [ConsumerConfig](docs/ConsumerConfig.md)
 - [ConsumerInfo](docs/ConsumerInfo.md)
 - [ConsumerInfoCluster](docs/ConsumerInfoCluster.md)
 - [DataStats](docs/DataStats.md)
 - [DeliverPolicy](docs/DeliverPolicy.md)
 - [DiscardPolicy](docs/DiscardPolicy.md)
 - [Export](docs/Export.md)
 - [ExportType](docs/ExportType.md)
 - [ExportsInner](docs/ExportsInner.md)
 - [ExternalStream](docs/ExternalStream.md)
 - [GatewayStat](docs/GatewayStat.md)
 - [GenericFields](docs/GenericFields.md)
 - [Import](docs/Import.md)
 - [ImportsInner](docs/ImportsInner.md)
 - [Info](docs/Info.md)
 - [JSApiError](docs/JSApiError.md)
 - [JSAssetInfoListResponse](docs/JSAssetInfoListResponse.md)
 - [JSAssetInfoResponse](docs/JSAssetInfoResponse.md)
 - [JSCommonConsumerConfigRequest](docs/JSCommonConsumerConfigRequest.md)
 - [JSCommonStreamConfig](docs/JSCommonStreamConfig.md)
 - [JSCommonStreamConfigPlacement](docs/JSCommonStreamConfigPlacement.md)
 - [JSCommonStreamConfigRepublish](docs/JSCommonStreamConfigRepublish.md)
 - [JSCommonStreamInfo](docs/JSCommonStreamInfo.md)
 - [JSConsumerInfoListResponse](docs/JSConsumerInfoListResponse.md)
 - [JSConsumerInfoResponse](docs/JSConsumerInfoResponse.md)
 - [JSConsumerType](docs/JSConsumerType.md)
 - [JSKVBucketCreateRequest](docs/JSKVBucketCreateRequest.md)
 - [JSKVBucketCreateRequestMirror](docs/JSKVBucketCreateRequestMirror.md)
 - [JSKVBucketListResponse](docs/JSKVBucketListResponse.md)
 - [JSKVBucketViewResponse](docs/JSKVBucketViewResponse.md)
 - [JSMirrorConfigRequest](docs/JSMirrorConfigRequest.md)
 - [JSMirrorInfoListResponse](docs/JSMirrorInfoListResponse.md)
 - [JSMirrorInfoResponse](docs/JSMirrorInfoResponse.md)
 - [JSPullConsumerConfigRequest](docs/JSPullConsumerConfigRequest.md)
 - [JSPullConsumerInfoResponse](docs/JSPullConsumerInfoResponse.md)
 - [JSPullConsumerUpdateRequest](docs/JSPullConsumerUpdateRequest.md)
 - [JSPushConsumerConfigRequest](docs/JSPushConsumerConfigRequest.md)
 - [JSPushConsumerInfoResponse](docs/JSPushConsumerInfoResponse.md)
 - [JSPushConsumerUpdateRequest](docs/JSPushConsumerUpdateRequest.md)
 - [JSStreamConfigRequest](docs/JSStreamConfigRequest.md)
 - [JSStreamInfoListResponse](docs/JSStreamInfoListResponse.md)
 - [JSStreamInfoResponse](docs/JSStreamInfoResponse.md)
 - [JSType](docs/JSType.md)
 - [JetStreamAPIStats](docs/JetStreamAPIStats.md)
 - [JetStreamAccountLimits](docs/JetStreamAccountLimits.md)
 - [JetStreamAccountStats](docs/JetStreamAccountStats.md)
 - [JetStreamConfig](docs/JetStreamConfig.md)
 - [JetStreamLimits](docs/JetStreamLimits.md)
 - [JetStreamStats](docs/JetStreamStats.md)
 - [JetStreamTier](docs/JetStreamTier.md)
 - [JetStreamVarz](docs/JetStreamVarz.md)
 - [JetStreamVarzConfig](docs/JetStreamVarzConfig.md)
 - [JetStreamVarzMeta](docs/JetStreamVarzMeta.md)
 - [JetStreamVarzStats](docs/JetStreamVarzStats.md)
 - [JwtGenericFieldsEditable](docs/JwtGenericFieldsEditable.md)
 - [LostStreamData](docs/LostStreamData.md)
 - [MetaClusterInfo](docs/MetaClusterInfo.md)
 - [NatsCluster](docs/NatsCluster.md)
 - [NatsClusterListResponse](docs/NatsClusterListResponse.md)
 - [NatsLimits](docs/NatsLimits.md)
 - [NatsServerInfoListResponse](docs/NatsServerInfoListResponse.md)
 - [NatsServerListResponse](docs/NatsServerListResponse.md)
 - [NatsUserClaimsRequest](docs/NatsUserClaimsRequest.md)
 - [NatsUserConnectionsListResponse](docs/NatsUserConnectionsListResponse.md)
 - [NatsUserCreateRequest](docs/NatsUserCreateRequest.md)
 - [NatsUserInfo](docs/NatsUserInfo.md)
 - [NatsUserListResponse](docs/NatsUserListResponse.md)
 - [NatsUserUpdateRequest](docs/NatsUserUpdateRequest.md)
 - [NatsUserUpdateRequestUserClaims](docs/NatsUserUpdateRequestUserClaims.md)
 - [NatsUserViewResponse](docs/NatsUserViewResponse.md)
 - [Operator](docs/Operator.md)
 - [OperatorClaims](docs/OperatorClaims.md)
 - [OperatorLimits](docs/OperatorLimits.md)
 - [PeerInfo](docs/PeerInfo.md)
 - [Permission](docs/Permission.md)
 - [Permissions](docs/Permissions.md)
 - [PermissionsResp](docs/PermissionsResp.md)
 - [Placement](docs/Placement.md)
 - [PolicyListResponse](docs/PolicyListResponse.md)
 - [PromSampleValue](docs/PromSampleValue.md)
 - [RePublish](docs/RePublish.md)
 - [ReplayPolicy](docs/ReplayPolicy.md)
 - [ResponsePermission](docs/ResponsePermission.md)
 - [ResponseType](docs/ResponseType.md)
 - [RetentionPolicy](docs/RetentionPolicy.md)
 - [RoleListResponse](docs/RoleListResponse.md)
 - [RouteStat](docs/RouteStat.md)
 - [SequenceInfo](docs/SequenceInfo.md)
 - [ServerInfo](docs/ServerInfo.md)
 - [ServerStats](docs/ServerStats.md)
 - [ServerStatsGatewaysInner](docs/ServerStatsGatewaysInner.md)
 - [ServerStatsJetstream](docs/ServerStatsJetstream.md)
 - [ServerStatsMsg](docs/ServerStatsMsg.md)
 - [ServerStatsRoutesInner](docs/ServerStatsRoutesInner.md)
 - [ServiceLatency](docs/ServiceLatency.md)
 - [SessionAccountListResponse](docs/SessionAccountListResponse.md)
 - [SessionAccountListResponseItemsInner](docs/SessionAccountListResponseItemsInner.md)
 - [SessionNatsUserListResponse](docs/SessionNatsUserListResponse.md)
 - [SessionNatsUserListResponseItemsInner](docs/SessionNatsUserListResponseItemsInner.md)
 - [SessionSystemListResponse](docs/SessionSystemListResponse.md)
 - [SessionSystemListResponseItemsInner](docs/SessionSystemListResponseItemsInner.md)
 - [SigningKeyGroupCreateRequest](docs/SigningKeyGroupCreateRequest.md)
 - [SigningKeyGroupListResponse](docs/SigningKeyGroupListResponse.md)
 - [SigningKeyGroupUpdateRequest](docs/SigningKeyGroupUpdateRequest.md)
 - [SigningKeyGroupViewResponse](docs/SigningKeyGroupViewResponse.md)
 - [SigningKeyListResponse](docs/SigningKeyListResponse.md)
 - [SigningKeyUpdateRequest](docs/SigningKeyUpdateRequest.md)
 - [SigningKeyViewResponse](docs/SigningKeyViewResponse.md)
 - [StorageType](docs/StorageType.md)
 - [StreamAlternate](docs/StreamAlternate.md)
 - [StreamExportCreateRequest](docs/StreamExportCreateRequest.md)
 - [StreamExportListResponse](docs/StreamExportListResponse.md)
 - [StreamExportSharedListResponse](docs/StreamExportSharedListResponse.md)
 - [StreamExportSharedViewResponse](docs/StreamExportSharedViewResponse.md)
 - [StreamExportViewResponse](docs/StreamExportViewResponse.md)
 - [StreamImportCreateRequest](docs/StreamImportCreateRequest.md)
 - [StreamImportListResponse](docs/StreamImportListResponse.md)
 - [StreamImportViewResponse](docs/StreamImportViewResponse.md)
 - [StreamShareCreateRequest](docs/StreamShareCreateRequest.md)
 - [StreamShareListResponse](docs/StreamShareListResponse.md)
 - [StreamShareViewResponse](docs/StreamShareViewResponse.md)
 - [StreamSource](docs/StreamSource.md)
 - [StreamSourceExternal](docs/StreamSourceExternal.md)
 - [StreamSourceInfo](docs/StreamSourceInfo.md)
 - [StreamSourceInfoError](docs/StreamSourceInfoError.md)
 - [StreamState](docs/StreamState.md)
 - [StreamStateLost](docs/StreamStateLost.md)
 - [SubDetail](docs/SubDetail.md)
 - [SubjectExportChartTypes](docs/SubjectExportChartTypes.md)
 - [SubjectExportCreateRequest](docs/SubjectExportCreateRequest.md)
 - [SubjectExportListResponse](docs/SubjectExportListResponse.md)
 - [SubjectExportSharedListResponse](docs/SubjectExportSharedListResponse.md)
 - [SubjectExportSharedViewResponse](docs/SubjectExportSharedViewResponse.md)
 - [SubjectExportUpdateRequest](docs/SubjectExportUpdateRequest.md)
 - [SubjectExportViewResponse](docs/SubjectExportViewResponse.md)
 - [SubjectImportCreateRequest](docs/SubjectImportCreateRequest.md)
 - [SubjectImportListResponse](docs/SubjectImportListResponse.md)
 - [SubjectImportUpdateRequest](docs/SubjectImportUpdateRequest.md)
 - [SubjectImportViewResponse](docs/SubjectImportViewResponse.md)
 - [SubjectShareCreateRequest](docs/SubjectShareCreateRequest.md)
 - [SubjectShareListResponse](docs/SubjectShareListResponse.md)
 - [SubjectShareViewResponse](docs/SubjectShareViewResponse.md)
 - [SystemAccountImportRequest](docs/SystemAccountImportRequest.md)
 - [SystemConnectionsListResponse](docs/SystemConnectionsListResponse.md)
 - [SystemCreateRequest](docs/SystemCreateRequest.md)
 - [SystemImportRequest](docs/SystemImportRequest.md)
 - [SystemInfo](docs/SystemInfo.md)
 - [SystemListResponse](docs/SystemListResponse.md)
 - [SystemState](docs/SystemState.md)
 - [SystemUpdateRequest](docs/SystemUpdateRequest.md)
 - [SystemUserImportRequest](docs/SystemUserImportRequest.md)
 - [SystemViewResponse](docs/SystemViewResponse.md)
 - [TLSPeerCert](docs/TLSPeerCert.md)
 - [TimeRange](docs/TimeRange.md)
 - [User](docs/User.md)
 - [UserClaims](docs/UserClaims.md)
 - [UserLimits](docs/UserLimits.md)
 - [UserPermissionLimits](docs/UserPermissionLimits.md)
 - [VersionResponse](docs/VersionResponse.md)
 - [WeightedMapping](docs/WeightedMapping.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


