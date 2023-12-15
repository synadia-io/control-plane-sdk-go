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

For using other server than the one defined on index 0 set context value `syncp.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), syncp.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `syncp.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), syncp.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `syncp.ContextOperationServerIndices` and `syncp.ContextOperationServerVariables` context maps.

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
*AccountAPI* | [**AssignAccountTeamAppUser**](docs/AccountAPI.md#assignaccountteamappuser) | **Post** /accounts/{accountId}/app-users/{teamAppUserId} | Assign Team App User to Account
*AccountAPI* | [**CreateAccountSkGroup**](docs/AccountAPI.md#createaccountskgroup) | **Post** /accounts/{accountId}/account-sk-groups | Create Account Signing Key Group
*AccountAPI* | [**CreateAlertRule**](docs/AccountAPI.md#createalertrule) | **Post** /accounts/{accountId}/alert-rules | Create Account Alert Rule
*AccountAPI* | [**CreateKvBucket**](docs/AccountAPI.md#createkvbucket) | **Post** /accounts/{accountId}/jetstream/kv-buckets | Create KV Bucket
*AccountAPI* | [**CreateMirror**](docs/AccountAPI.md#createmirror) | **Post** /accounts/{accountId}/jetstream/mirrors | Create Mirror
*AccountAPI* | [**CreateObjectBucket**](docs/AccountAPI.md#createobjectbucket) | **Post** /accounts/{accountId}/jetstream/object-buckets | Create Object Bucket
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
*AccountAPI* | [**GetJetStreamPlacementOptions**](docs/AccountAPI.md#getjetstreamplacementoptions) | **Get** /accounts/{accountId}/jetstream/placement-options | Get JetStream Placement Options
*AccountAPI* | [**ListAccountConnections**](docs/AccountAPI.md#listaccountconnections) | **Get** /accounts/{accountId}/connections | List Account Connections
*AccountAPI* | [**ListAccountSkGroup**](docs/AccountAPI.md#listaccountskgroup) | **Get** /accounts/{accountId}/account-sk-groups | List Account Signing Key Groups
*AccountAPI* | [**ListAccountTeamAppUsers**](docs/AccountAPI.md#listaccountteamappusers) | **Get** /accounts/{accountId}/app-users | List Account Team App Users
*AccountAPI* | [**ListAlertRules**](docs/AccountAPI.md#listalertrules) | **Get** /accounts/{accountId}/alert-rules | List Account Alert Rules
*AccountAPI* | [**ListJetStreamAssets**](docs/AccountAPI.md#listjetstreamassets) | **Get** /accounts/{accountId}/jetstream | List JetStream Assets
*AccountAPI* | [**ListKvBuckets**](docs/AccountAPI.md#listkvbuckets) | **Get** /accounts/{accountId}/jetstream/kv-buckets | List KV buckets
*AccountAPI* | [**ListMirrors**](docs/AccountAPI.md#listmirrors) | **Get** /accounts/{accountId}/jetstream/mirrors | List Mirrors
*AccountAPI* | [**ListObjectBuckets**](docs/AccountAPI.md#listobjectbuckets) | **Get** /accounts/{accountId}/jetstream/object-buckets | List Object buckets
*AccountAPI* | [**ListStreamExports**](docs/AccountAPI.md#liststreamexports) | **Get** /accounts/{accountId}/stream-exports | List Stream Exports
*AccountAPI* | [**ListStreamExportsShared**](docs/AccountAPI.md#liststreamexportsshared) | **Get** /accounts/{accountId}/stream-imports/shared | List Shared Stream Exports
*AccountAPI* | [**ListStreamImports**](docs/AccountAPI.md#liststreamimports) | **Get** /accounts/{accountId}/stream-imports | List Stream Imports
*AccountAPI* | [**ListStreams**](docs/AccountAPI.md#liststreams) | **Get** /accounts/{accountId}/jetstream/streams | List Streams
*AccountAPI* | [**ListSubjectExports**](docs/AccountAPI.md#listsubjectexports) | **Get** /accounts/{accountId}/subject-exports | List Subject Exports
*AccountAPI* | [**ListSubjectExportsShared**](docs/AccountAPI.md#listsubjectexportsshared) | **Get** /accounts/{accountId}/subject-imports/shared | List Shared Subject Exports
*AccountAPI* | [**ListSubjectImports**](docs/AccountAPI.md#listsubjectimports) | **Get** /accounts/{accountId}/subject-imports | List Subject Imports
*AccountAPI* | [**ListUsers**](docs/AccountAPI.md#listusers) | **Get** /accounts/{accountId}/nats-users | List NATS Users
*AccountAPI* | [**RunAlertRule**](docs/AccountAPI.md#runalertrule) | **Get** /accounts/{accountId}/alert-rules/{alertRuleId}/run | Run Account Alert Rule
*AccountAPI* | [**UnAssignAccountTeamAppUser**](docs/AccountAPI.md#unassignaccountteamappuser) | **Delete** /accounts/{accountId}/app-users/{teamAppUserId} | Unassign Team App User from Account
*AccountAPI* | [**UnmanageAccount**](docs/AccountAPI.md#unmanageaccount) | **Delete** /accounts/{accountId}/unmanage | Unmanage Account
*AccountAPI* | [**UpdateAccount**](docs/AccountAPI.md#updateaccount) | **Patch** /accounts/{accountId} | Update Account
*AccountAPI* | [**UpdateAlertRule**](docs/AccountAPI.md#updatealertrule) | **Patch** /accounts/{accountId}/alert-rules/{alertRuleId} | Update Account Alert Rule
*AgentTokenAPI* | [**DeleteAgentToken**](docs/AgentTokenAPI.md#deleteagenttoken) | **Delete** /agent-tokens/{tokenId} | Delete Agent Token
*AlertAPI* | [**AcknowledgeAlert**](docs/AlertAPI.md#acknowledgealert) | **Patch** /alerts/{alertId}/acknowledge | Acknowledge Alert
*AlertAPI* | [**GetAlert**](docs/AlertAPI.md#getalert) | **Get** /alerts/{alertId} | Get Alert
*AppUserAPI* | [**AssignTeamAppUser**](docs/AppUserAPI.md#assignteamappuser) | **Post** /app-users/{appUserId}/teams/{teamId} | Assign App User to Team
*AppUserAPI* | [**DeleteAppUser**](docs/AppUserAPI.md#deleteappuser) | **Delete** /app-users/{appUserId} | Delete App User
*AppUserAPI* | [**GetAppUser**](docs/AppUserAPI.md#getappuser) | **Get** /app-users/{appUserId} | Get App User
*AppUserAPI* | [**ListAppUserRoles**](docs/AppUserAPI.md#listappuserroles) | **Get** /app-users/{appUserId}/roles | List Roles
*AppUserAPI* | [**UpdateAppUser**](docs/AppUserAPI.md#updateappuser) | **Patch** /app-users/{appUserId} | Update App User
*AuthzAPI* | [**Check**](docs/AuthzAPI.md#check) | **Post** /authz/check | Check Authz Decisions
*AuthzAPI* | [**ListPolicies**](docs/AuthzAPI.md#listpolicies) | **Get** /authz/policies | Get Policy List
*AuthzAPI* | [**ListRoles**](docs/AuthzAPI.md#listroles) | **Get** /authz/roles | Get Authz roles List
*KvBucketAPI* | [**CreateKvPullConsumer**](docs/KvBucketAPI.md#createkvpullconsumer) | **Post** /jetstream/kv-bucket/{streamId}/consumers/pull | Create Pull Consumer
*KvBucketAPI* | [**CreateKvPushConsumer**](docs/KvBucketAPI.md#createkvpushconsumer) | **Post** /jetstream/kv-bucket/{streamId}/consumers/push | Create Push Consumer
*KvBucketAPI* | [**DeleteKvBucket**](docs/KvBucketAPI.md#deletekvbucket) | **Delete** /jetstream/kv-bucket/{streamId} | Delete KV Bucket
*KvBucketAPI* | [**GetKvBucket**](docs/KvBucketAPI.md#getkvbucket) | **Get** /jetstream/kv-bucket/{streamId} | Get KV Bucket
*KvBucketAPI* | [**ListKvConsumers**](docs/KvBucketAPI.md#listkvconsumers) | **Get** /jetstream/kv-bucket/{streamId}/consumers | List Consumers
*KvBucketAPI* | [**UpdateKvBucket**](docs/KvBucketAPI.md#updatekvbucket) | **Patch** /jetstream/kv-bucket/{streamId} | Update KV Bucket
*MirrorAPI* | [**CreateMirrorPullConsumer**](docs/MirrorAPI.md#createmirrorpullconsumer) | **Post** /jetstream/mirror/{streamId}/consumers/pull | Create Pull Consumer
*MirrorAPI* | [**CreateMirrorPushConsumer**](docs/MirrorAPI.md#createmirrorpushconsumer) | **Post** /jetstream/mirror/{streamId}/consumers/push | Create Push consumer
*MirrorAPI* | [**DeleteMirror**](docs/MirrorAPI.md#deletemirror) | **Delete** /jetstream/mirror/{streamId} | Delete Mirror
*MirrorAPI* | [**GetMirror**](docs/MirrorAPI.md#getmirror) | **Get** /jetstream/mirror/{streamId} | Get Mirror
*MirrorAPI* | [**ListMirrorConsumers**](docs/MirrorAPI.md#listmirrorconsumers) | **Get** /jetstream/mirror/{streamId}/consumers | List Consumers
*MirrorAPI* | [**UpdateMirror**](docs/MirrorAPI.md#updatemirror) | **Patch** /jetstream/mirror/{streamId} | Update Mirror
*NatsUserAPI* | [**AssignNatsUserTeamAppUser**](docs/NatsUserAPI.md#assignnatsuserteamappuser) | **Post** /nats-users/{userId}/app-users/{teamAppUserId} | Assign Team App User to NATS User
*NatsUserAPI* | [**CopyNatsUser**](docs/NatsUserAPI.md#copynatsuser) | **Post** /nats-users/{userId}/copy | Copy nats user
*NatsUserAPI* | [**DeleteNatsUser**](docs/NatsUserAPI.md#deletenatsuser) | **Delete** /nats-users/{userId} | Delete NATS User
*NatsUserAPI* | [**DownloadNatsUserCreds**](docs/NatsUserAPI.md#downloadnatsusercreds) | **Post** /nats-users/{userId}/creds | Get Creds
*NatsUserAPI* | [**GetNatsUser**](docs/NatsUserAPI.md#getnatsuser) | **Get** /nats-users/{userId} | Get NATS User
*NatsUserAPI* | [**GetNatsUserIssuance**](docs/NatsUserAPI.md#getnatsuserissuance) | **Get** /nats-user-issuances/{issuanceId} | Get nats user issuance
*NatsUserAPI* | [**ListNatsUserConnections**](docs/NatsUserAPI.md#listnatsuserconnections) | **Get** /nats-users/{userId}/connections | List NATs User Connections
*NatsUserAPI* | [**ListNatsUserIssuances**](docs/NatsUserAPI.md#listnatsuserissuances) | **Get** /nats-users/{userId}/issuances | List nats user issuances
*NatsUserAPI* | [**ListNatsUserTeamAppUsers**](docs/NatsUserAPI.md#listnatsuserteamappusers) | **Get** /nats-users/{userId}/app-users | List Team App Users
*NatsUserAPI* | [**RotateNatsUser**](docs/NatsUserAPI.md#rotatenatsuser) | **Post** /nats-users/{userId}/rotate | Rotate nats user nkey and seed
*NatsUserAPI* | [**UnAssignNatsUserTeamAppUser**](docs/NatsUserAPI.md#unassignnatsuserteamappuser) | **Delete** /nats-users/{userId}/app-users/{teamAppUserId} | Unassign Team App User from NATS User
*NatsUserAPI* | [**UpdateNatsUser**](docs/NatsUserAPI.md#updatenatsuser) | **Patch** /nats-users/{userId} | Update NATS User
*ObjectBucketAPI* | [**CreateObjPullConsumer**](docs/ObjectBucketAPI.md#createobjpullconsumer) | **Post** /jetstream/object-bucket/{streamId}/consumers/pull | Create Pull Consumer
*ObjectBucketAPI* | [**CreateObjPushConsumer**](docs/ObjectBucketAPI.md#createobjpushconsumer) | **Post** /jetstream/object-bucket/{streamId}/consumers/push | Create Push Consumer
*ObjectBucketAPI* | [**DeleteObjectBucket**](docs/ObjectBucketAPI.md#deleteobjectbucket) | **Delete** /jetstream/object-bucket/{streamId} | Delete Object Bucket
*ObjectBucketAPI* | [**GetObjectBucket**](docs/ObjectBucketAPI.md#getobjectbucket) | **Get** /jetstream/object-bucket/{streamId} | Get Object Bucket
*ObjectBucketAPI* | [**ListObjConsumers**](docs/ObjectBucketAPI.md#listobjconsumers) | **Get** /jetstream/object-bucket/{streamId}/consumers | List Consumers
*ObjectBucketAPI* | [**UpdateObjectBucket**](docs/ObjectBucketAPI.md#updateobjectbucket) | **Patch** /jetstream/object-bucket/{streamId} | Update Object Bucket
*PersonalAccessTokenAPI* | [**DeletePersonalAccessToken**](docs/PersonalAccessTokenAPI.md#deletepersonalaccesstoken) | **Delete** /personal-access-tokens/{tokenId} | Delete Personal Access Token
*PersonalAccessTokenAPI* | [**GetPersonalAccessToken**](docs/PersonalAccessTokenAPI.md#getpersonalaccesstoken) | **Get** /personal-access-tokens/{tokenId} | Get Personal Access Token
*PersonalAccessTokenAPI* | [**UpdatePersonalAccessToken**](docs/PersonalAccessTokenAPI.md#updatepersonalaccesstoken) | **Patch** /personal-access-tokens/{tokenId} | Update Personal Access Token
*PullConsumerAPI* | [**DeletePullConsumer**](docs/PullConsumerAPI.md#deletepullconsumer) | **Delete** /consumers/pull/{consumerId} | Delete Pull Consumer
*PullConsumerAPI* | [**GetPullConsumerInfo**](docs/PullConsumerAPI.md#getpullconsumerinfo) | **Get** /consumers/pull/{consumerId} | Get Pull Consumer
*PullConsumerAPI* | [**UpdatePullConsumer**](docs/PullConsumerAPI.md#updatepullconsumer) | **Patch** /consumers/pull/{consumerId} | Update Pull Consumer
*PushConsumerAPI* | [**DeletePushConsumer**](docs/PushConsumerAPI.md#deletepushconsumer) | **Delete** /consumers/push/{consumerId} | Delete Push Consumer
*PushConsumerAPI* | [**GetPushConsumerInfo**](docs/PushConsumerAPI.md#getpushconsumerinfo) | **Get** /consumers/push/{consumerId} | Get Push Consumer
*PushConsumerAPI* | [**UpdatePushConsumer**](docs/PushConsumerAPI.md#updatepushconsumer) | **Patch** /consumers/push/{consumerId} | Update Push Consumer
*SessionAPI* | [**AcceptTerms**](docs/SessionAPI.md#acceptterms) | **Post** /terms/accept | Accept terms
*SessionAPI* | [**CreateAppUser**](docs/SessionAPI.md#createappuser) | **Post** /app-users | Create App User
*SessionAPI* | [**CreatePersonalAccessToken**](docs/SessionAPI.md#createpersonalaccesstoken) | **Post** /personal-access-tokens | Create Personal Access Token
*SessionAPI* | [**CreateTeam**](docs/SessionAPI.md#createteam) | **Post** /teams | Create Team
*SessionAPI* | [**DecideInvitation**](docs/SessionAPI.md#decideinvitation) | **Post** /invitations/{teamId} | Accept or reject team invitation
*SessionAPI* | [**GetVersion**](docs/SessionAPI.md#getversion) | **Get** /version | Get Version
*SessionAPI* | [**ListAlerts**](docs/SessionAPI.md#listalerts) | **Get** /alerts | List Alerts
*SessionAPI* | [**ListAppUsers**](docs/SessionAPI.md#listappusers) | **Get** /app-users | List App Users
*SessionAPI* | [**ListInvitations**](docs/SessionAPI.md#listinvitations) | **Get** /invitations | List of pending invitations
*SessionAPI* | [**ListPersonalAccessTokens**](docs/SessionAPI.md#listpersonalaccesstokens) | **Get** /personal-access-tokens | List Personal Access Tokens
*SessionAPI* | [**ListTeams**](docs/SessionAPI.md#listteams) | **Get** /teams | List Teams
*SessionAPI* | [**SearchSystemAccounts**](docs/SessionAPI.md#searchsystemaccounts) | **Get** /search/systems/{systemId}/accounts | Search System Accounts
*SessionAPI* | [**SearchSystemServers**](docs/SessionAPI.md#searchsystemservers) | **Get** /search/systems/{systemId}/servers | Search System Servers
*SessionAPI* | [**SearchTeamAppUsers**](docs/SessionAPI.md#searchteamappusers) | **Get** /search/teams/{teamId}/app-users | Search App Users in Team
*SigKeyAPI* | [**DeleteAccountSk**](docs/SigKeyAPI.md#deleteaccountsk) | **Delete** /account-sks/{keyId} | Delete Account Signing
*SigKeyAPI* | [**GetAccountSk**](docs/SigKeyAPI.md#getaccountsk) | **Get** /account-sks/{keyId} | Get Account Signing
*SigKeyAPI* | [**UpdateAccountSk**](docs/SigKeyAPI.md#updateaccountsk) | **Patch** /account-sks/{keyId} | Update Account Signing
*SigKeyGroupAPI* | [**CopyAccountSKGroup**](docs/SigKeyGroupAPI.md#copyaccountskgroup) | **Post** /account-sk-groups/{groupId}/copy | Copy Account SK Group
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
*SystemAPI* | [**AssignSystemTeamAppUser**](docs/SystemAPI.md#assignsystemteamappuser) | **Post** /systems/{systemId}/app-users/{teamAppUserId} | Assign Team App User to System
*SystemAPI* | [**CreateAccount**](docs/SystemAPI.md#createaccount) | **Post** /systems/{systemId}/accounts | Create Account
*SystemAPI* | [**CreateSystemAlertRule**](docs/SystemAPI.md#createsystemalertrule) | **Post** /systems/{systemId}/alert-rules | Create System Alert Rule
*SystemAPI* | [**DeleteSystem**](docs/SystemAPI.md#deletesystem) | **Delete** /systems/{systemId} | Delete System
*SystemAPI* | [**DeleteSystemAlertRule**](docs/SystemAPI.md#deletesystemalertrule) | **Delete** /systems/{systemId}/alert-rules/{alertRuleId} | Delete System Alert Rule
*SystemAPI* | [**DownloadSystemLogs**](docs/SystemAPI.md#downloadsystemlogs) | **Post** /systems/{systemId}/logs | Download Logs
*SystemAPI* | [**GetCurrentAgentToken**](docs/SystemAPI.md#getcurrentagenttoken) | **Get** /systems/{systemId}/agent-tokens/current | Get Current Agent Token
*SystemAPI* | [**GetSystem**](docs/SystemAPI.md#getsystem) | **Get** /systems/{systemId} | Get System
*SystemAPI* | [**GetSystemAlertRule**](docs/SystemAPI.md#getsystemalertrule) | **Get** /systems/{systemId}/alert-rules/{alertRuleId} | Get System Alert Rule
*SystemAPI* | [**GetSystemLimits**](docs/SystemAPI.md#getsystemlimits) | **Get** /systems/{systemId}/limits | Get System Limits
*SystemAPI* | [**ImportAccount**](docs/SystemAPI.md#importaccount) | **Post** /systems/{systemId}/import-account | Import Account
*SystemAPI* | [**ImportUser**](docs/SystemAPI.md#importuser) | **Post** /systems/{systemId}/import-user | Import User
*SystemAPI* | [**ListAccounts**](docs/SystemAPI.md#listaccounts) | **Get** /systems/{systemId}/accounts | List Accounts
*SystemAPI* | [**ListAccountsOverviewMetrics**](docs/SystemAPI.md#listaccountsoverviewmetrics) | **Get** /systems/{systemId}/accounts-overview-metrics | List Accounts overview metrics
*SystemAPI* | [**ListAgentTokens**](docs/SystemAPI.md#listagenttokens) | **Get** /systems/{systemId}/agent-tokens | List Agent Tokens
*SystemAPI* | [**ListClusters**](docs/SystemAPI.md#listclusters) | **Get** /systems/{systemId}/nats-clusters | List Clusters
*SystemAPI* | [**ListConnections**](docs/SystemAPI.md#listconnections) | **Get** /systems/{systemId}/connections | List Connections
*SystemAPI* | [**ListServers**](docs/SystemAPI.md#listservers) | **Get** /systems/{systemId}/servers | List Servers
*SystemAPI* | [**ListSystemAlertRules**](docs/SystemAPI.md#listsystemalertrules) | **Get** /systems/{systemId}/alert-rules | List System Alert Rules
*SystemAPI* | [**ListSystemTeamAppUsers**](docs/SystemAPI.md#listsystemteamappusers) | **Get** /systems/{systemId}/app-users | List System Team App Users
*SystemAPI* | [**RotateAgentToken**](docs/SystemAPI.md#rotateagenttoken) | **Post** /systems/{systemId}/agent-tokens | Rotate Agent Token
*SystemAPI* | [**RunSystemAlertRule**](docs/SystemAPI.md#runsystemalertrule) | **Get** /systems/{systemId}/alert-rules/{alertRuleId}/run | Run System Alert Rule
*SystemAPI* | [**UnAssignSystemTeamAppUser**](docs/SystemAPI.md#unassignsystemteamappuser) | **Delete** /systems/{systemId}/app-users/{teamAppUserId} | Unassign Team App User from System
*SystemAPI* | [**UpdateSystem**](docs/SystemAPI.md#updatesystem) | **Patch** /systems/{systemId} | Update System
*SystemAPI* | [**UpdateSystemAlertRule**](docs/SystemAPI.md#updatesystemalertrule) | **Patch** /systems/{systemId}/alert-rules/{alertRuleId} | Update System Alert Rules
*TeamAPI* | [**CreateSystem**](docs/TeamAPI.md#createsystem) | **Post** /teams/{teamId}/systems | Create System
*TeamAPI* | [**DeleteTeam**](docs/TeamAPI.md#deleteteam) | **Delete** /teams/{teamId} | Delete Team
*TeamAPI* | [**GetTeam**](docs/TeamAPI.md#getteam) | **Get** /teams/{teamId} | Get Team
*TeamAPI* | [**GetTeamLimits**](docs/TeamAPI.md#getteamlimits) | **Get** /teams/{teamId}/team-limits | Get Team Limits
*TeamAPI* | [**ImportSystem**](docs/TeamAPI.md#importsystem) | **Post** /teams/{teamId}/import-system | Import a System
*TeamAPI* | [**InviteAppUser**](docs/TeamAPI.md#inviteappuser) | **Post** /teams/{teamId}/app-users/invitations | Invite App Users
*TeamAPI* | [**LeaveTeam**](docs/TeamAPI.md#leaveteam) | **Post** /teams/{teamId}/app-users/leave | Leave Team
*TeamAPI* | [**ListTeamAccounts**](docs/TeamAPI.md#listteamaccounts) | **Get** /teams/{teamId}/accounts | List Accounts
*TeamAPI* | [**ListTeamAppUsers**](docs/TeamAPI.md#listteamappusers) | **Get** /teams/{teamId}/app-users | List App Users
*TeamAPI* | [**ListTeamNatsUsers**](docs/TeamAPI.md#listteamnatsusers) | **Get** /teams/{teamId}/nats-users | List NATS Users
*TeamAPI* | [**ListTeamSystems**](docs/TeamAPI.md#listteamsystems) | **Get** /teams/{teamId}/systems | List Systems
*TeamAPI* | [**UnAssignTeamAppUser**](docs/TeamAPI.md#unassignteamappuser) | **Delete** /teams/{teamId}/app-users/{appUserId} | Unassign App User from Team
*TeamAPI* | [**UpdateTeam**](docs/TeamAPI.md#updateteam) | **Patch** /teams/{teamId} | Update Team
*TeamAPI* | [**UpdateTeamAppUser**](docs/TeamAPI.md#updateteamappuser) | **Patch** /teams/{teamId}/app-users/{appUserId} | Update App User Team Assignment


## Documentation For Models

 - [AcceptTermsResponse](docs/AcceptTermsResponse.md)
 - [Account](docs/Account.md)
 - [AccountChartType](docs/AccountChartType.md)
 - [AccountClaims](docs/AccountClaims.md)
 - [AccountClaimsInfo](docs/AccountClaimsInfo.md)
 - [AccountConnectionsListResponse](docs/AccountConnectionsListResponse.md)
 - [AccountCreateRequest](docs/AccountCreateRequest.md)
 - [AccountInfo](docs/AccountInfo.md)
 - [AccountJWTSettings](docs/AccountJWTSettings.md)
 - [AccountLimits](docs/AccountLimits.md)
 - [AccountListResponse](docs/AccountListResponse.md)
 - [AccountMetrics](docs/AccountMetrics.md)
 - [AccountOverviewResponse](docs/AccountOverviewResponse.md)
 - [AccountOverviewResponseMetrics](docs/AccountOverviewResponseMetrics.md)
 - [AccountSearch](docs/AccountSearch.md)
 - [AccountSearchListResponse](docs/AccountSearchListResponse.md)
 - [AccountUpdateRequest](docs/AccountUpdateRequest.md)
 - [AccountViewResponse](docs/AccountViewResponse.md)
 - [AccountsOverviewListResponse](docs/AccountsOverviewListResponse.md)
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
 - [ApexCoordinate](docs/ApexCoordinate.md)
 - [ApexSeries](docs/ApexSeries.md)
 - [AppPolicy](docs/AppPolicy.md)
 - [AppPolicyStatement](docs/AppPolicyStatement.md)
 - [AppRole](docs/AppRole.md)
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
 - [AppUserInvitationRequest](docs/AppUserInvitationRequest.md)
 - [AppUserInvitationResponse](docs/AppUserInvitationResponse.md)
 - [AppUserListResponse](docs/AppUserListResponse.md)
 - [AppUserType](docs/AppUserType.md)
 - [AppUserUpdateRequest](docs/AppUserUpdateRequest.md)
 - [AppUserUpdateResponse](docs/AppUserUpdateResponse.md)
 - [AppUserViewResponse](docs/AppUserViewResponse.md)
 - [AuthzRequest](docs/AuthzRequest.md)
 - [AuthzResponse](docs/AuthzResponse.md)
 - [ClaimsData](docs/ClaimsData.md)
 - [ClusterChartType](docs/ClusterChartType.md)
 - [ClusterInfo](docs/ClusterInfo.md)
 - [ConnInfo](docs/ConnInfo.md)
 - [Connz](docs/Connz.md)
 - [ConsumerConfig](docs/ConsumerConfig.md)
 - [ConsumerInfo](docs/ConsumerInfo.md)
 - [DataStats](docs/DataStats.md)
 - [DeliverPolicy](docs/DeliverPolicy.md)
 - [DiscardPolicy](docs/DiscardPolicy.md)
 - [Export](docs/Export.md)
 - [ExportType](docs/ExportType.md)
 - [ExternalAuthorization](docs/ExternalAuthorization.md)
 - [ExternalStream](docs/ExternalStream.md)
 - [GatewayStat](docs/GatewayStat.md)
 - [GenericFields](docs/GenericFields.md)
 - [Import](docs/Import.md)
 - [Info](docs/Info.md)
 - [InvitationDecisionRequest](docs/InvitationDecisionRequest.md)
 - [InvitationListResponse](docs/InvitationListResponse.md)
 - [JSApiError](docs/JSApiError.md)
 - [JSAssetInfoListResponse](docs/JSAssetInfoListResponse.md)
 - [JSAssetInfoResponse](docs/JSAssetInfoResponse.md)
 - [JSCommonConsumerConfigRequest](docs/JSCommonConsumerConfigRequest.md)
 - [JSCommonStreamConfig](docs/JSCommonStreamConfig.md)
 - [JSCommonStreamInfo](docs/JSCommonStreamInfo.md)
 - [JSConsumerInfoListResponse](docs/JSConsumerInfoListResponse.md)
 - [JSConsumerInfoResponse](docs/JSConsumerInfoResponse.md)
 - [JSConsumerType](docs/JSConsumerType.md)
 - [JSKVBucketConfig](docs/JSKVBucketConfig.md)
 - [JSKVBucketListResponse](docs/JSKVBucketListResponse.md)
 - [JSKVBucketUpdateRequest](docs/JSKVBucketUpdateRequest.md)
 - [JSKVBucketViewResponse](docs/JSKVBucketViewResponse.md)
 - [JSMirrorConfigRequest](docs/JSMirrorConfigRequest.md)
 - [JSMirrorInfoListResponse](docs/JSMirrorInfoListResponse.md)
 - [JSMirrorInfoResponse](docs/JSMirrorInfoResponse.md)
 - [JSObjectBucketConfig](docs/JSObjectBucketConfig.md)
 - [JSObjectBucketListResponse](docs/JSObjectBucketListResponse.md)
 - [JSObjectBucketUpdateRequest](docs/JSObjectBucketUpdateRequest.md)
 - [JSObjectBucketViewResponse](docs/JSObjectBucketViewResponse.md)
 - [JSPlacementOptionsResponse](docs/JSPlacementOptionsResponse.md)
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
 - [JwtGenericFieldsEditable](docs/JwtGenericFieldsEditable.md)
 - [LostStreamData](docs/LostStreamData.md)
 - [MetaClusterInfo](docs/MetaClusterInfo.md)
 - [NatsCluster](docs/NatsCluster.md)
 - [NatsClusterListResponse](docs/NatsClusterListResponse.md)
 - [NatsLimits](docs/NatsLimits.md)
 - [NatsServerInfoListResponse](docs/NatsServerInfoListResponse.md)
 - [NatsServerListResponse](docs/NatsServerListResponse.md)
 - [NatsUserConnectionsListResponse](docs/NatsUserConnectionsListResponse.md)
 - [NatsUserCopyRequest](docs/NatsUserCopyRequest.md)
 - [NatsUserCreateRequest](docs/NatsUserCreateRequest.md)
 - [NatsUserInfo](docs/NatsUserInfo.md)
 - [NatsUserIssuanceEvent](docs/NatsUserIssuanceEvent.md)
 - [NatsUserIssuanceEventType](docs/NatsUserIssuanceEventType.md)
 - [NatsUserIssuanceInfo](docs/NatsUserIssuanceInfo.md)
 - [NatsUserIssuanceViewResponse](docs/NatsUserIssuanceViewResponse.md)
 - [NatsUserIssuancesListResponse](docs/NatsUserIssuancesListResponse.md)
 - [NatsUserJwtSettings](docs/NatsUserJwtSettings.md)
 - [NatsUserListResponse](docs/NatsUserListResponse.md)
 - [NatsUserUpdateRequest](docs/NatsUserUpdateRequest.md)
 - [NatsUserViewResponse](docs/NatsUserViewResponse.md)
 - [ObjectStoreConfig](docs/ObjectStoreConfig.md)
 - [Operator](docs/Operator.md)
 - [OperatorClaims](docs/OperatorClaims.md)
 - [OperatorLimits](docs/OperatorLimits.md)
 - [PeerInfo](docs/PeerInfo.md)
 - [Permission](docs/Permission.md)
 - [Permissions](docs/Permissions.md)
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
 - [S2Compression](docs/S2Compression.md)
 - [SequenceInfo](docs/SequenceInfo.md)
 - [ServerInfo](docs/ServerInfo.md)
 - [ServerStats](docs/ServerStats.md)
 - [ServerStatsMsg](docs/ServerStatsMsg.md)
 - [ServiceLatency](docs/ServiceLatency.md)
 - [SigningKeyGroupCopyRequest](docs/SigningKeyGroupCopyRequest.md)
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
 - [StreamSourceInfo](docs/StreamSourceInfo.md)
 - [StreamState](docs/StreamState.md)
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
 - [SubjectTransformConfig](docs/SubjectTransformConfig.md)
 - [SystemAccountImportRequest](docs/SystemAccountImportRequest.md)
 - [SystemConnectionsListResponse](docs/SystemConnectionsListResponse.md)
 - [SystemCreateRequest](docs/SystemCreateRequest.md)
 - [SystemImportRequest](docs/SystemImportRequest.md)
 - [SystemInfo](docs/SystemInfo.md)
 - [SystemLimitsResponse](docs/SystemLimitsResponse.md)
 - [SystemListResponse](docs/SystemListResponse.md)
 - [SystemState](docs/SystemState.md)
 - [SystemUpdateRequest](docs/SystemUpdateRequest.md)
 - [SystemUserImportRequest](docs/SystemUserImportRequest.md)
 - [SystemViewResponse](docs/SystemViewResponse.md)
 - [TLSPeerCert](docs/TLSPeerCert.md)
 - [TeamAppUserAssociation](docs/TeamAppUserAssociation.md)
 - [TeamAppUserInfo](docs/TeamAppUserInfo.md)
 - [TeamAppUserListResponse](docs/TeamAppUserListResponse.md)
 - [TeamAppUserResourceType](docs/TeamAppUserResourceType.md)
 - [TeamAppUserViewResponse](docs/TeamAppUserViewResponse.md)
 - [TeamCreateRequest](docs/TeamCreateRequest.md)
 - [TeamInfo](docs/TeamInfo.md)
 - [TeamLimits](docs/TeamLimits.md)
 - [TeamLimitsResponse](docs/TeamLimitsResponse.md)
 - [TeamListResponse](docs/TeamListResponse.md)
 - [TeamUpdateRequest](docs/TeamUpdateRequest.md)
 - [TeamViewResponse](docs/TeamViewResponse.md)
 - [TenantLimits](docs/TenantLimits.md)
 - [TimeRange](docs/TimeRange.md)
 - [UpdatableKVBucketConfig](docs/UpdatableKVBucketConfig.md)
 - [UpdatableObjectBucketConfig](docs/UpdatableObjectBucketConfig.md)
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
auth := context.WithValue(context.Background(), syncp.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions and generic types to easily obtain pointers to values of basic types.
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



