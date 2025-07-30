# Go API client for syncp

API for Synadia Control Plane / Synadia Cloud


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

All URIs are relative to *http://localhost/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountAPI* | [**AssignAccountTeamAppUser**](docs/AccountAPI.md#assignaccountteamappuser) | **Post** /core/beta/accounts/{accountId}/app-users/{teamAppUserId} | Assign Team App User to Account
*AccountAPI* | [**CreateAccountSkGroup**](docs/AccountAPI.md#createaccountskgroup) | **Post** /core/beta/accounts/{accountId}/account-sk-groups | Create Account Signing Key Group
*AccountAPI* | [**CreateAlertRule**](docs/AccountAPI.md#createalertrule) | **Post** /core/beta/accounts/{accountId}/alert-rules | Create Account Alert Rule
*AccountAPI* | [**CreateKvBucket**](docs/AccountAPI.md#createkvbucket) | **Post** /core/beta/accounts/{accountId}/jetstream/kv-buckets | Create KV Bucket
*AccountAPI* | [**CreateMirror**](docs/AccountAPI.md#createmirror) | **Post** /core/beta/accounts/{accountId}/jetstream/mirrors | Create Mirror
*AccountAPI* | [**CreateObjectBucket**](docs/AccountAPI.md#createobjectbucket) | **Post** /core/beta/accounts/{accountId}/jetstream/object-buckets | Create Object Bucket
*AccountAPI* | [**CreateOrUpdateNatsUserRevocation**](docs/AccountAPI.md#createorupdatenatsuserrevocation) | **Put** /core/beta/accounts/{accountId}/nats-user-revocations/{userNkeyPublic} | Create or Update Revocation for a NATS User NKey
*AccountAPI* | [**CreateStream**](docs/AccountAPI.md#createstream) | **Post** /core/beta/accounts/{accountId}/jetstream/streams | Create Stream
*AccountAPI* | [**CreateStreamExport**](docs/AccountAPI.md#createstreamexport) | **Post** /core/beta/accounts/{accountId}/stream-exports | Create Stream Export
*AccountAPI* | [**CreateStreamImport**](docs/AccountAPI.md#createstreamimport) | **Post** /core/beta/accounts/{accountId}/stream-imports | Create Stream Import
*AccountAPI* | [**CreateSubjectExport**](docs/AccountAPI.md#createsubjectexport) | **Post** /core/beta/accounts/{accountId}/subject-exports | Create Subject Export
*AccountAPI* | [**CreateSubjectImport**](docs/AccountAPI.md#createsubjectimport) | **Post** /core/beta/accounts/{accountId}/subject-imports | Create Subject Import
*AccountAPI* | [**CreateUser**](docs/AccountAPI.md#createuser) | **Post** /core/beta/accounts/{accountId}/nats-users | Create NATS User
*AccountAPI* | [**DeleteAccount**](docs/AccountAPI.md#deleteaccount) | **Delete** /core/beta/accounts/{accountId} | Delete Account
*AccountAPI* | [**DeleteAlertRule**](docs/AccountAPI.md#deletealertrule) | **Delete** /core/beta/accounts/{accountId}/alert-rules/{alertRuleId} | Delete Account Alert Rule
*AccountAPI* | [**DeleteNatsUserRevocation**](docs/AccountAPI.md#deletenatsuserrevocation) | **Delete** /core/beta/accounts/{accountId}/nats-user-revocations/{userNkeyPublic} | Delete a for a NATS User NKey
*AccountAPI* | [**GetAccount**](docs/AccountAPI.md#getaccount) | **Get** /core/beta/accounts/{accountId} | Get Account
*AccountAPI* | [**GetAccountExport**](docs/AccountAPI.md#getaccountexport) | **Get** /core/beta/accounts/{accountId}/export | Export Account Seeds
*AccountAPI* | [**GetAccountInfo**](docs/AccountAPI.md#getaccountinfo) | **Get** /core/beta/accounts/{accountId}/info | Get Account Info
*AccountAPI* | [**GetAccountMetrics**](docs/AccountAPI.md#getaccountmetrics) | **Get** /core/beta/accounts/{accountId}/metrics | Get Account Metrics
*AccountAPI* | [**GetAlertRule**](docs/AccountAPI.md#getalertrule) | **Get** /core/beta/accounts/{accountId}/alert-rules/{alertRuleId} | Get Account Alert Rule
*AccountAPI* | [**GetJetStreamPlacementOptions**](docs/AccountAPI.md#getjetstreamplacementoptions) | **Get** /core/beta/accounts/{accountId}/jetstream/placement-options | Get JetStream Placement Options
*AccountAPI* | [**GetNatsUserRevocation**](docs/AccountAPI.md#getnatsuserrevocation) | **Get** /core/beta/accounts/{accountId}/nats-user-revocations/{userNkeyPublic} | Get Revocation for a NATS User NKey
*AccountAPI* | [**ListAccountConnections**](docs/AccountAPI.md#listaccountconnections) | **Get** /core/beta/accounts/{accountId}/connections | List Account Connections
*AccountAPI* | [**ListAccountSkGroup**](docs/AccountAPI.md#listaccountskgroup) | **Get** /core/beta/accounts/{accountId}/account-sk-groups | List Account Signing Key Groups
*AccountAPI* | [**ListAccountTeamAppUsers**](docs/AccountAPI.md#listaccountteamappusers) | **Get** /core/beta/accounts/{accountId}/app-users | List Account Team App Users
*AccountAPI* | [**ListAlertRules**](docs/AccountAPI.md#listalertrules) | **Get** /core/beta/accounts/{accountId}/alert-rules | List Account Alert Rules
*AccountAPI* | [**ListJetStreamAssets**](docs/AccountAPI.md#listjetstreamassets) | **Get** /core/beta/accounts/{accountId}/jetstream | List JetStream Assets
*AccountAPI* | [**ListKvBuckets**](docs/AccountAPI.md#listkvbuckets) | **Get** /core/beta/accounts/{accountId}/jetstream/kv-buckets | List KV buckets
*AccountAPI* | [**ListMirrors**](docs/AccountAPI.md#listmirrors) | **Get** /core/beta/accounts/{accountId}/jetstream/mirrors | List Mirrors
*AccountAPI* | [**ListObjectBuckets**](docs/AccountAPI.md#listobjectbuckets) | **Get** /core/beta/accounts/{accountId}/jetstream/object-buckets | List Object buckets
*AccountAPI* | [**ListStreamExports**](docs/AccountAPI.md#liststreamexports) | **Get** /core/beta/accounts/{accountId}/stream-exports | List Stream Exports
*AccountAPI* | [**ListStreamExportsShared**](docs/AccountAPI.md#liststreamexportsshared) | **Get** /core/beta/accounts/{accountId}/stream-imports/shared | List Shared Stream Exports
*AccountAPI* | [**ListStreamImports**](docs/AccountAPI.md#liststreamimports) | **Get** /core/beta/accounts/{accountId}/stream-imports | List Stream Imports
*AccountAPI* | [**ListStreams**](docs/AccountAPI.md#liststreams) | **Get** /core/beta/accounts/{accountId}/jetstream/streams | List Streams
*AccountAPI* | [**ListSubjectExports**](docs/AccountAPI.md#listsubjectexports) | **Get** /core/beta/accounts/{accountId}/subject-exports | List Subject Exports
*AccountAPI* | [**ListSubjectExportsShared**](docs/AccountAPI.md#listsubjectexportsshared) | **Get** /core/beta/accounts/{accountId}/subject-imports/shared | List Shared Subject Exports
*AccountAPI* | [**ListSubjectImports**](docs/AccountAPI.md#listsubjectimports) | **Get** /core/beta/accounts/{accountId}/subject-imports | List Subject Imports
*AccountAPI* | [**ListUsers**](docs/AccountAPI.md#listusers) | **Get** /core/beta/accounts/{accountId}/nats-users | List NATS Users
*AccountAPI* | [**NatsCoreWebsocketViewer**](docs/AccountAPI.md#natscorewebsocketviewer) | **Get** /core/beta/accounts/{accountId}/nats-core/websocket/view | Subscribe to a NATS Core subject over websockets
*AccountAPI* | [**RunAlertRule**](docs/AccountAPI.md#runalertrule) | **Get** /core/beta/accounts/{accountId}/alert-rules/{alertRuleId}/run | Run Account Alert Rule
*AccountAPI* | [**UnAssignAccountTeamAppUser**](docs/AccountAPI.md#unassignaccountteamappuser) | **Delete** /core/beta/accounts/{accountId}/app-users/{teamAppUserId} | Unassign Team App User from Account
*AccountAPI* | [**UnmanageAccount**](docs/AccountAPI.md#unmanageaccount) | **Delete** /core/beta/accounts/{accountId}/unmanage | Unmanage Account
*AccountAPI* | [**UpdateAccount**](docs/AccountAPI.md#updateaccount) | **Patch** /core/beta/accounts/{accountId} | Update Account
*AccountAPI* | [**UpdateAlertRule**](docs/AccountAPI.md#updatealertrule) | **Patch** /core/beta/accounts/{accountId}/alert-rules/{alertRuleId} | Update Account Alert Rule
*AdminAPI* | [**CheckAppAdmin**](docs/AdminAPI.md#checkappadmin) | **Get** /core/beta/admin/app-user | Check for Admin User
*AdminAPI* | [**CreateAppAdmin**](docs/AdminAPI.md#createappadmin) | **Post** /core/beta/admin/app-user | Create an App Admin User
*AgentTokenAPI* | [**DeleteAgentToken**](docs/AgentTokenAPI.md#deleteagenttoken) | **Delete** /core/beta/agent-tokens/{tokenId} | Delete Agent Token
*AlertAPI* | [**AcknowledgeAlert**](docs/AlertAPI.md#acknowledgealert) | **Patch** /core/beta/alerts/{alertId}/acknowledge | Acknowledge Alert
*AlertAPI* | [**GetAlert**](docs/AlertAPI.md#getalert) | **Get** /core/beta/alerts/{alertId} | Get Alert
*AppServiceAccountAPI* | [**CreateAppServiceAccountToken**](docs/AppServiceAccountAPI.md#createappserviceaccounttoken) | **Post** /core/beta/service-accounts/app/{serviceAccountId}/tokens | Create Acess Token for App Service Account
*AppServiceAccountAPI* | [**DeleteAppServiceAccount**](docs/AppServiceAccountAPI.md#deleteappserviceaccount) | **Delete** /core/beta/service-accounts/app/{serviceAccountId} | Delete App Service Account
*AppServiceAccountAPI* | [**GetAppServiceAccount**](docs/AppServiceAccountAPI.md#getappserviceaccount) | **Get** /core/beta/service-accounts/app/{serviceAccountId} | Get App Service Account
*AppServiceAccountAPI* | [**ListAppServiceAccountTokens**](docs/AppServiceAccountAPI.md#listappserviceaccounttokens) | **Get** /core/beta/service-accounts/app/{serviceAccountId}/tokens | List Access Tokens for App Service Account
*AppServiceAccountAPI* | [**UpdateAppServiceAccount**](docs/AppServiceAccountAPI.md#updateappserviceaccount) | **Patch** /core/beta/service-accounts/app/{serviceAccountId} | Update App Service Account
*AppServiceAccountTokensAPI* | [**DeleteAppServiceAccountToken**](docs/AppServiceAccountTokensAPI.md#deleteappserviceaccounttoken) | **Delete** /core/beta/service-account-tokens/app/{tokenId} | Delete App Service Account Token
*AppServiceAccountTokensAPI* | [**GetAppServiceAccountToken**](docs/AppServiceAccountTokensAPI.md#getappserviceaccounttoken) | **Get** /core/beta/service-account-tokens/app/{tokenId} | Get App Service Account Token
*AppServiceAccountTokensAPI* | [**UpdateAppServiceAccountToken**](docs/AppServiceAccountTokensAPI.md#updateappserviceaccounttoken) | **Patch** /core/beta/service-account-tokens/app/{tokenId} | Update App Service Account Token
*AppUserAPI* | [**AssignTeamAppUser**](docs/AppUserAPI.md#assignteamappuser) | **Post** /core/beta/app-users/{appUserId}/teams/{teamId} | Assign App User to Team
*AppUserAPI* | [**DeleteAppUser**](docs/AppUserAPI.md#deleteappuser) | **Delete** /core/beta/app-users/{appUserId} | Delete App User
*AppUserAPI* | [**GetAppUser**](docs/AppUserAPI.md#getappuser) | **Get** /core/beta/app-users/{appUserId} | Get App User
*AppUserAPI* | [**ListAppUserRoles**](docs/AppUserAPI.md#listappuserroles) | **Get** /core/beta/app-users/{appUserId}/roles | List Roles
*AppUserAPI* | [**UpdateAppUser**](docs/AppUserAPI.md#updateappuser) | **Patch** /core/beta/app-users/{appUserId} | Update App User
*AuthCalloutAPI* | [**AddAuthCalloutTargetAccount**](docs/AuthCalloutAPI.md#addauthcallouttargetaccount) | **Post** /core/beta/auth-callout/{authCalloutId}/target-accounts | Configure Target Account
*AuthCalloutAPI* | [**AddAuthCalloutUser**](docs/AuthCalloutAPI.md#addauthcalloutuser) | **Post** /core/beta/auth-callout/{authCalloutId}/users | Create Auth Callout User
*AuthCalloutAPI* | [**DeleteAuthCallout**](docs/AuthCalloutAPI.md#deleteauthcallout) | **Delete** /core/beta/auth-callout/{authCalloutId} | Delete Auth Callout Config
*AuthCalloutAPI* | [**DeleteAuthCalloutTargetAccount**](docs/AuthCalloutAPI.md#deleteauthcallouttargetaccount) | **Delete** /core/beta/auth-callout/{authCalloutId}/target-accounts/{targetAccountId} | Delete Target Account
*AuthCalloutAPI* | [**DeleteAuthCalloutUser**](docs/AuthCalloutAPI.md#deleteauthcalloutuser) | **Delete** /core/beta/auth-callout/{authCalloutId}/users/{acUserId} | Delete Control Account User
*AuthCalloutAPI* | [**GetAuthCallout**](docs/AuthCalloutAPI.md#getauthcallout) | **Get** /core/beta/auth-callout/{authCalloutId} | Auth Callout Config
*AuthCalloutAPI* | [**ListAuthCalloutTargetAccounts**](docs/AuthCalloutAPI.md#listauthcallouttargetaccounts) | **Get** /core/beta/auth-callout/{authCalloutId}/target-accounts | Get Target Account List
*AuthCalloutAPI* | [**ListAuthCalloutUsers**](docs/AuthCalloutAPI.md#listauthcalloutusers) | **Get** /core/beta/auth-callout/{authCalloutId}/users | Get Target Account List
*AuthzAPI* | [**Check**](docs/AuthzAPI.md#check) | **Post** /core/beta/authz/check | Check Authz Decisions
*AuthzAPI* | [**ListPolicies**](docs/AuthzAPI.md#listpolicies) | **Get** /core/beta/authz/policies | Get Policy List
*AuthzAPI* | [**ListRoles**](docs/AuthzAPI.md#listroles) | **Get** /core/beta/authz/roles | Get Authz roles List
*ComponentVersionsAPI* | [**GetAvailableComponentVersions**](docs/ComponentVersionsAPI.md#getavailablecomponentversions) | **Get** /core/beta/component-versions | Get Available Component Versions
*ComponentVersionsAPI* | [**GetAvailableComponentVersionsByType**](docs/ComponentVersionsAPI.md#getavailablecomponentversionsbytype) | **Post** /core/beta/component-versions | Get Available Component Versions By Type
*IssuanceAPI* | [**GetNatsUserIssuance**](docs/IssuanceAPI.md#getnatsuserissuance) | **Get** /core/beta/nats-user-issuances/{issuanceId} | Get nats user issuance
*KvBucketAPI* | [**CreateKvPullConsumer**](docs/KvBucketAPI.md#createkvpullconsumer) | **Post** /core/beta/jetstream/kv-bucket/{streamId}/consumers/pull | Create Pull Consumer
*KvBucketAPI* | [**CreateKvPushConsumer**](docs/KvBucketAPI.md#createkvpushconsumer) | **Post** /core/beta/jetstream/kv-bucket/{streamId}/consumers/push | Create Push Consumer
*KvBucketAPI* | [**DeleteKvBucket**](docs/KvBucketAPI.md#deletekvbucket) | **Delete** /core/beta/jetstream/kv-bucket/{streamId} | Delete KV Bucket
*KvBucketAPI* | [**GetKvBucket**](docs/KvBucketAPI.md#getkvbucket) | **Get** /core/beta/jetstream/kv-bucket/{streamId} | Get KV Bucket
*KvBucketAPI* | [**ListKvConsumers**](docs/KvBucketAPI.md#listkvconsumers) | **Get** /core/beta/jetstream/kv-bucket/{streamId}/consumers | List Consumers
*KvBucketAPI* | [**PurgeKvBucket**](docs/KvBucketAPI.md#purgekvbucket) | **Delete** /core/beta/jetstream/kv-bucket/{streamId}/purge | Purge KV Bucket
*KvBucketAPI* | [**UpdateKvBucket**](docs/KvBucketAPI.md#updatekvbucket) | **Patch** /core/beta/jetstream/kv-bucket/{streamId} | Update KV Bucket
*MirrorAPI* | [**CreateMirrorPullConsumer**](docs/MirrorAPI.md#createmirrorpullconsumer) | **Post** /core/beta/jetstream/mirror/{streamId}/consumers/pull | Create Pull Consumer
*MirrorAPI* | [**CreateMirrorPushConsumer**](docs/MirrorAPI.md#createmirrorpushconsumer) | **Post** /core/beta/jetstream/mirror/{streamId}/consumers/push | Create Push consumer
*MirrorAPI* | [**DeleteMirror**](docs/MirrorAPI.md#deletemirror) | **Delete** /core/beta/jetstream/mirror/{streamId} | Delete Mirror
*MirrorAPI* | [**GetMirror**](docs/MirrorAPI.md#getmirror) | **Get** /core/beta/jetstream/mirror/{streamId} | Get Mirror
*MirrorAPI* | [**ListMirrorConsumers**](docs/MirrorAPI.md#listmirrorconsumers) | **Get** /core/beta/jetstream/mirror/{streamId}/consumers | List Consumers
*MirrorAPI* | [**PurgeMirror**](docs/MirrorAPI.md#purgemirror) | **Delete** /core/beta/jetstream/mirror/{streamId}/purge | Purge Mirror
*MirrorAPI* | [**UpdateMirror**](docs/MirrorAPI.md#updatemirror) | **Patch** /core/beta/jetstream/mirror/{streamId} | Update Mirror
*NatsUserAPI* | [**AssignNatsUserTeamAppUser**](docs/NatsUserAPI.md#assignnatsuserteamappuser) | **Post** /core/beta/nats-users/{userId}/app-users/{teamAppUserId} | Assign Team App User to NATS User
*NatsUserAPI* | [**CopyNatsUser**](docs/NatsUserAPI.md#copynatsuser) | **Post** /core/beta/nats-users/{userId}/copy | Copy nats user
*NatsUserAPI* | [**DeleteNatsUser**](docs/NatsUserAPI.md#deletenatsuser) | **Delete** /core/beta/nats-users/{userId} | Delete NATS User
*NatsUserAPI* | [**DownloadNatsUserBearerJwt**](docs/NatsUserAPI.md#downloadnatsuserbearerjwt) | **Post** /core/beta/nats-users/{userId}/bearer-jwt | Get Bearer JWT
*NatsUserAPI* | [**DownloadNatsUserCreds**](docs/NatsUserAPI.md#downloadnatsusercreds) | **Post** /core/beta/nats-users/{userId}/creds | Get Creds
*NatsUserAPI* | [**DownloadNatsUserHttpGwToken**](docs/NatsUserAPI.md#downloadnatsuserhttpgwtoken) | **Post** /core/beta/nats-users/{userId}/http-gw-token | Get HTTP Gateway Token
*NatsUserAPI* | [**GetNatsUser**](docs/NatsUserAPI.md#getnatsuser) | **Get** /core/beta/nats-users/{userId} | Get NATS User
*NatsUserAPI* | [**ListNatsUserConnections**](docs/NatsUserAPI.md#listnatsuserconnections) | **Get** /core/beta/nats-users/{userId}/connections | List NATs User Connections
*NatsUserAPI* | [**ListNatsUserIssuances**](docs/NatsUserAPI.md#listnatsuserissuances) | **Get** /core/beta/nats-users/{userId}/issuances | List nats user issuances
*NatsUserAPI* | [**ListNatsUserTeamAppUsers**](docs/NatsUserAPI.md#listnatsuserteamappusers) | **Get** /core/beta/nats-users/{userId}/app-users | List Team App Users
*NatsUserAPI* | [**RotateNatsUser**](docs/NatsUserAPI.md#rotatenatsuser) | **Post** /core/beta/nats-users/{userId}/rotate | Rotate nats user nkey and seed
*NatsUserAPI* | [**UnAssignNatsUserTeamAppUser**](docs/NatsUserAPI.md#unassignnatsuserteamappuser) | **Delete** /core/beta/nats-users/{userId}/app-users/{teamAppUserId} | Unassign Team App User from NATS User
*NatsUserAPI* | [**UpdateNatsUser**](docs/NatsUserAPI.md#updatenatsuser) | **Patch** /core/beta/nats-users/{userId} | Update NATS User
*ObjectBucketAPI* | [**CreateObjPullConsumer**](docs/ObjectBucketAPI.md#createobjpullconsumer) | **Post** /core/beta/jetstream/object-bucket/{streamId}/consumers/pull | Create Pull Consumer
*ObjectBucketAPI* | [**CreateObjPushConsumer**](docs/ObjectBucketAPI.md#createobjpushconsumer) | **Post** /core/beta/jetstream/object-bucket/{streamId}/consumers/push | Create Push Consumer
*ObjectBucketAPI* | [**DeleteObjectBucket**](docs/ObjectBucketAPI.md#deleteobjectbucket) | **Delete** /core/beta/jetstream/object-bucket/{streamId} | Delete Object Bucket
*ObjectBucketAPI* | [**GetObjectBucket**](docs/ObjectBucketAPI.md#getobjectbucket) | **Get** /core/beta/jetstream/object-bucket/{streamId} | Get Object Bucket
*ObjectBucketAPI* | [**ListObjConsumers**](docs/ObjectBucketAPI.md#listobjconsumers) | **Get** /core/beta/jetstream/object-bucket/{streamId}/consumers | List Consumers
*ObjectBucketAPI* | [**PurgeObjBucket**](docs/ObjectBucketAPI.md#purgeobjbucket) | **Delete** /core/beta/jetstream/object-bucket/{streamId}/purge | Purge Object Bucket
*ObjectBucketAPI* | [**UpdateObjectBucket**](docs/ObjectBucketAPI.md#updateobjectbucket) | **Patch** /core/beta/jetstream/object-bucket/{streamId} | Update Object Bucket
*OidcProviderAPI* | [**CreateOidcProvider**](docs/OidcProviderAPI.md#createoidcprovider) | **Post** /core/beta/oidc-providers | Create OIDC Provider
*OidcProviderAPI* | [**DeleteOidcProvider**](docs/OidcProviderAPI.md#deleteoidcprovider) | **Delete** /core/beta/oidc-providers/{oidcProviderId} | Delete OIDC Provider
*OidcProviderAPI* | [**GetOidcProvider**](docs/OidcProviderAPI.md#getoidcprovider) | **Get** /core/beta/oidc-providers/{oidcProviderId} | Get OIDC Provider
*OidcProviderAPI* | [**ListOidcDefaultProvider**](docs/OidcProviderAPI.md#listoidcdefaultprovider) | **Get** /core/beta/oidc-providers/defaults | List OIDC Provider Defaults
*OidcProviderAPI* | [**ListOidcProviders**](docs/OidcProviderAPI.md#listoidcproviders) | **Get** /core/beta/oidc-providers | List OIDC Providers
*OidcProviderAPI* | [**UpdateOidcProvider**](docs/OidcProviderAPI.md#updateoidcprovider) | **Patch** /core/beta/oidc-providers/{oidcProviderId} | Update OIDC Provider
*PersonalAccessTokenAPI* | [**DeletePersonalAccessToken**](docs/PersonalAccessTokenAPI.md#deletepersonalaccesstoken) | **Delete** /core/beta/personal-access-tokens/{tokenId} | Delete Personal Access Token
*PersonalAccessTokenAPI* | [**GetPersonalAccessToken**](docs/PersonalAccessTokenAPI.md#getpersonalaccesstoken) | **Get** /core/beta/personal-access-tokens/{tokenId} | Get Personal Access Token
*PersonalAccessTokenAPI* | [**UpdatePersonalAccessToken**](docs/PersonalAccessTokenAPI.md#updatepersonalaccesstoken) | **Patch** /core/beta/personal-access-tokens/{tokenId} | Update Personal Access Token
*PlatformComponentsAPI* | [**PlatformComponentConnect**](docs/PlatformComponentsAPI.md#platformcomponentconnect) | **Post** /core/beta/platform-components/connect | Connect a Platform Component
*PullConsumerAPI* | [**DeletePullConsumer**](docs/PullConsumerAPI.md#deletepullconsumer) | **Delete** /core/beta/consumers/pull/{consumerId} | Delete Pull Consumer
*PullConsumerAPI* | [**GetPullConsumerInfo**](docs/PullConsumerAPI.md#getpullconsumerinfo) | **Get** /core/beta/consumers/pull/{consumerId} | Get Pull Consumer
*PullConsumerAPI* | [**UpdatePullConsumer**](docs/PullConsumerAPI.md#updatepullconsumer) | **Patch** /core/beta/consumers/pull/{consumerId} | Update Pull Consumer
*PushConsumerAPI* | [**DeletePushConsumer**](docs/PushConsumerAPI.md#deletepushconsumer) | **Delete** /core/beta/consumers/push/{consumerId} | Delete Push Consumer
*PushConsumerAPI* | [**GetPushConsumerInfo**](docs/PushConsumerAPI.md#getpushconsumerinfo) | **Get** /core/beta/consumers/push/{consumerId} | Get Push Consumer
*PushConsumerAPI* | [**UpdatePushConsumer**](docs/PushConsumerAPI.md#updatepushconsumer) | **Patch** /core/beta/consumers/push/{consumerId} | Update Push Consumer
*SessionAPI* | [**AcceptTerms**](docs/SessionAPI.md#acceptterms) | **Post** /core/beta/terms/accept | Accept terms
*SessionAPI* | [**CreateAppServiceAccount**](docs/SessionAPI.md#createappserviceaccount) | **Post** /core/beta/service-accounts/app | Create App Service Account
*SessionAPI* | [**CreateAppUser**](docs/SessionAPI.md#createappuser) | **Post** /core/beta/app-users | Create App User
*SessionAPI* | [**CreateAppUserAccessToken**](docs/SessionAPI.md#createappuseraccesstoken) | **Post** /core/beta/personal-access-tokens | Create Personal Access Token
*SessionAPI* | [**CreateTeam**](docs/SessionAPI.md#createteam) | **Post** /core/beta/teams | Create Team
*SessionAPI* | [**DecideInvitation**](docs/SessionAPI.md#decideinvitation) | **Post** /core/beta/invitations/{teamId} | Accept or reject team invitation
*SessionAPI* | [**GetVersion**](docs/SessionAPI.md#getversion) | **Get** /core/beta/version | Get Version
*SessionAPI* | [**ListAlerts**](docs/SessionAPI.md#listalerts) | **Get** /core/beta/alerts | List Alerts
*SessionAPI* | [**ListAppServiceAccounts**](docs/SessionAPI.md#listappserviceaccounts) | **Get** /core/beta/service-accounts/app | List App Service Accounts
*SessionAPI* | [**ListAppUserAccessTokens**](docs/SessionAPI.md#listappuseraccesstokens) | **Get** /core/beta/personal-access-tokens | List Personal Access Tokens
*SessionAPI* | [**ListAppUsers**](docs/SessionAPI.md#listappusers) | **Get** /core/beta/app-users | List App Users
*SessionAPI* | [**ListInvitations**](docs/SessionAPI.md#listinvitations) | **Get** /core/beta/invitations | List of pending invitations
*SessionAPI* | [**ListTeams**](docs/SessionAPI.md#listteams) | **Get** /core/beta/teams | List Teams
*SigKeyAPI* | [**DeleteAccountSk**](docs/SigKeyAPI.md#deleteaccountsk) | **Delete** /core/beta/account-sks/{keyId} | Delete Account Signing
*SigKeyAPI* | [**GetAccountSk**](docs/SigKeyAPI.md#getaccountsk) | **Get** /core/beta/account-sks/{keyId} | Get Account Signing
*SigKeyAPI* | [**UpdateAccountSk**](docs/SigKeyAPI.md#updateaccountsk) | **Patch** /core/beta/account-sks/{keyId} | Update Account Signing
*SigKeyGroupAPI* | [**CopyAccountSkGroup**](docs/SigKeyGroupAPI.md#copyaccountskgroup) | **Post** /core/beta/account-sk-groups/{groupId}/copy | Copy Account SK Group
*SigKeyGroupAPI* | [**DeleteAccountSkGroup**](docs/SigKeyGroupAPI.md#deleteaccountskgroup) | **Delete** /core/beta/account-sk-groups/{groupId} | Delete Account Signing Key Group
*SigKeyGroupAPI* | [**GetAccountSkGroup**](docs/SigKeyGroupAPI.md#getaccountskgroup) | **Get** /core/beta/account-sk-groups/{groupId} | Get Account Signing Key Group
*SigKeyGroupAPI* | [**ListAccountSKGroupUsers**](docs/SigKeyGroupAPI.md#listaccountskgroupusers) | **Get** /core/beta/account-sk-groups/{groupId}/nats-users | List NATS Users
*SigKeyGroupAPI* | [**ListAccountSkGroupKeys**](docs/SigKeyGroupAPI.md#listaccountskgroupkeys) | **Get** /core/beta/account-sk-groups/{groupId}/account-sks | List Signing Keys
*SigKeyGroupAPI* | [**RotateAccountSk**](docs/SigKeyGroupAPI.md#rotateaccountsk) | **Post** /core/beta/account-sk-groups/{groupId}/rotate-sk | Roate Active Signing Key
*SigKeyGroupAPI* | [**UpdateAccountSkGroup**](docs/SigKeyGroupAPI.md#updateaccountskgroup) | **Patch** /core/beta/account-sk-groups/{groupId} | Update Account Signing Key Group
*StreamAPI* | [**CreatePullConsumer**](docs/StreamAPI.md#createpullconsumer) | **Post** /core/beta/jetstream/stream/{streamId}/consumers/pull | Create Pull Consumer
*StreamAPI* | [**CreatePushConsumer**](docs/StreamAPI.md#createpushconsumer) | **Post** /core/beta/jetstream/stream/{streamId}/consumers/push | Create Push Consumer
*StreamAPI* | [**DeleteStream**](docs/StreamAPI.md#deletestream) | **Delete** /core/beta/jetstream/stream/{streamId} | Delete Stream
*StreamAPI* | [**GetStreamInfo**](docs/StreamAPI.md#getstreaminfo) | **Get** /core/beta/jetstream/stream/{streamId} | Get Stream
*StreamAPI* | [**ListConsumers**](docs/StreamAPI.md#listconsumers) | **Get** /core/beta/jetstream/stream/{streamId}/consumers | List Consumers
*StreamAPI* | [**PurgeStream**](docs/StreamAPI.md#purgestream) | **Delete** /core/beta/jetstream/stream/{streamId}/purge | Purge Stream
*StreamAPI* | [**UpdateStream**](docs/StreamAPI.md#updatestream) | **Patch** /core/beta/jetstream/stream/{streamId} | Update Stream
*StreamExportAPI* | [**CreateStreamShares**](docs/StreamExportAPI.md#createstreamshares) | **Post** /core/beta/stream-exports/{streamExportId}/shares | Create Stream Shares
*StreamExportAPI* | [**DeleteStreamExport**](docs/StreamExportAPI.md#deletestreamexport) | **Delete** /core/beta/stream-exports/{streamExportId} | Delete Stream Export
*StreamExportAPI* | [**DeleteStreamShare**](docs/StreamExportAPI.md#deletestreamshare) | **Delete** /core/beta/stream-exports/{streamExportId}/shares/{targetAccountNKeyPublic} | Delete Stream Share
*StreamExportAPI* | [**GetStreamExport**](docs/StreamExportAPI.md#getstreamexport) | **Get** /core/beta/stream-exports/{streamExportId} | Get Stream Export
*StreamExportAPI* | [**ListStreamShares**](docs/StreamExportAPI.md#liststreamshares) | **Get** /core/beta/stream-exports/{streamExportId}/shares | List Stream Shares
*StreamImportAPI* | [**DeleteStreamImport**](docs/StreamImportAPI.md#deletestreamimport) | **Delete** /core/beta/stream-imports/{streamImportId} | Delete Stream Import
*StreamImportAPI* | [**GetStreamImport**](docs/StreamImportAPI.md#getstreamimport) | **Get** /core/beta/stream-imports/{streamImportId} | Get Stream Import
*SubjectExportAPI* | [**CreateSubjectShares**](docs/SubjectExportAPI.md#createsubjectshares) | **Post** /core/beta/subject-exports/{subjectExportId}/shares | Create Subject Shares
*SubjectExportAPI* | [**DeleteSubjectExport**](docs/SubjectExportAPI.md#deletesubjectexport) | **Delete** /core/beta/subject-exports/{subjectExportId} | Delete Subject Export
*SubjectExportAPI* | [**DeleteSubjectShare**](docs/SubjectExportAPI.md#deletesubjectshare) | **Delete** /core/beta/subject-exports/{subjectExportId}/shares/{targetAccountNKeyPublic} | Delete Subject Share
*SubjectExportAPI* | [**GetSubjectExport**](docs/SubjectExportAPI.md#getsubjectexport) | **Get** /core/beta/subject-exports/{subjectExportId} | Get Subject Export
*SubjectExportAPI* | [**ListSubjectShares**](docs/SubjectExportAPI.md#listsubjectshares) | **Get** /core/beta/subject-exports/{subjectExportId}/shares | List Subject Shares
*SubjectExportAPI* | [**UpdateSubjectExport**](docs/SubjectExportAPI.md#updatesubjectexport) | **Patch** /core/beta/subject-exports/{subjectExportId} | Update Subject Export
*SubjectImportAPI* | [**DeleteSubjectImport**](docs/SubjectImportAPI.md#deletesubjectimport) | **Delete** /core/beta/subject-imports/{subjectImportId} | Delete Subject Import
*SubjectImportAPI* | [**GetSubjectImport**](docs/SubjectImportAPI.md#getsubjectimport) | **Get** /core/beta/subject-imports/{subjectImportId} | Get Subject Import
*SubjectImportAPI* | [**UpdateSubjectImport**](docs/SubjectImportAPI.md#updatesubjectimport) | **Patch** /core/beta/subject-imports/{subjectImportId} | Update Subject Import
*SystemAPI* | [**AssignSystemTeamAppUser**](docs/SystemAPI.md#assignsystemteamappuser) | **Post** /core/beta/systems/{systemId}/app-users/{teamAppUserId} | Assign Team App User to System
*SystemAPI* | [**BulkShare**](docs/SystemAPI.md#bulkshare) | **Post** /core/beta/systems/{systemId}/bulk/shares | Share assets across accounts
*SystemAPI* | [**CreateAccount**](docs/SystemAPI.md#createaccount) | **Post** /core/beta/systems/{systemId}/accounts | Create Account
*SystemAPI* | [**CreateSystemAlertRule**](docs/SystemAPI.md#createsystemalertrule) | **Post** /core/beta/systems/{systemId}/alert-rules | Create System Alert Rule
*SystemAPI* | [**DeleteSystem**](docs/SystemAPI.md#deletesystem) | **Delete** /core/beta/systems/{systemId} | Delete System
*SystemAPI* | [**DeleteSystemAlertRule**](docs/SystemAPI.md#deletesystemalertrule) | **Delete** /core/beta/systems/{systemId}/alert-rules/{alertRuleId} | Delete System Alert Rule
*SystemAPI* | [**EnableAuthCallout**](docs/SystemAPI.md#enableauthcallout) | **Post** /core/beta/systems/{systemId}/auth-callout | Enable Auth Callout For System
*SystemAPI* | [**GetComponentToken**](docs/SystemAPI.md#getcomponenttoken) | **Get** /core/beta/systems/{systemId}/platform-components/{id}/tokens | Get a component access token
*SystemAPI* | [**GetCurrentAgentToken**](docs/SystemAPI.md#getcurrentagenttoken) | **Get** /core/beta/systems/{systemId}/agent-tokens/current | Get Current Agent Token
*SystemAPI* | [**GetSystem**](docs/SystemAPI.md#getsystem) | **Get** /core/beta/systems/{systemId} | Get System
*SystemAPI* | [**GetSystemAlertRule**](docs/SystemAPI.md#getsystemalertrule) | **Get** /core/beta/systems/{systemId}/alert-rules/{alertRuleId} | Get System Alert Rule
*SystemAPI* | [**GetSystemExport**](docs/SystemAPI.md#getsystemexport) | **Get** /core/beta/systems/{systemId}/export | Export System Seeds
*SystemAPI* | [**GetSystemLimits**](docs/SystemAPI.md#getsystemlimits) | **Get** /core/beta/systems/{systemId}/limits | Get System Limits
*SystemAPI* | [**GetSystemPrometheusMetrics**](docs/SystemAPI.md#getsystemprometheusmetrics) | **Get** /core/beta/systems/{systemId}/prometheus-metrics | Get Prometheus Metrics
*SystemAPI* | [**ImportAccount**](docs/SystemAPI.md#importaccount) | **Post** /core/beta/systems/{systemId}/import-account | Import Account
*SystemAPI* | [**ImportUser**](docs/SystemAPI.md#importuser) | **Post** /core/beta/systems/{systemId}/import-user | Import User
*SystemAPI* | [**ListAccounts**](docs/SystemAPI.md#listaccounts) | **Get** /core/beta/systems/{systemId}/accounts | List Accounts
*SystemAPI* | [**ListAccountsOverviewMetrics**](docs/SystemAPI.md#listaccountsoverviewmetrics) | **Get** /core/beta/systems/{systemId}/accounts-overview-metrics | List Accounts overview metrics
*SystemAPI* | [**ListAgentTokens**](docs/SystemAPI.md#listagenttokens) | **Get** /core/beta/systems/{systemId}/agent-tokens | List Agent Tokens
*SystemAPI* | [**ListAuthCalloutAuthenticators**](docs/SystemAPI.md#listauthcalloutauthenticators) | **Get** /core/beta/systems/{systemId}/auth-callout/authenticators | Get List of Available Authenticators
*SystemAPI* | [**ListAuthCalloutConfigs**](docs/SystemAPI.md#listauthcalloutconfigs) | **Get** /core/beta/systems/{systemId}/auth-callout | List Auth Callout Configs
*SystemAPI* | [**ListClusters**](docs/SystemAPI.md#listclusters) | **Get** /core/beta/systems/{systemId}/nats-clusters | List Clusters
*SystemAPI* | [**ListConnections**](docs/SystemAPI.md#listconnections) | **Get** /core/beta/systems/{systemId}/connections | List Connections
*SystemAPI* | [**ListServers**](docs/SystemAPI.md#listservers) | **Get** /core/beta/systems/{systemId}/servers | List Servers
*SystemAPI* | [**ListSystemAlertRules**](docs/SystemAPI.md#listsystemalertrules) | **Get** /core/beta/systems/{systemId}/alert-rules | List System Alert Rules
*SystemAPI* | [**ListSystemInfoAccounts**](docs/SystemAPI.md#listsysteminfoaccounts) | **Get** /core/beta/systems/{systemId}/info/accounts | List System Accounts Info
*SystemAPI* | [**ListSystemInfoServers**](docs/SystemAPI.md#listsysteminfoservers) | **Get** /core/beta/systems/{systemId}/info/servers | List System Servers info
*SystemAPI* | [**ListSystemTeamAppUsers**](docs/SystemAPI.md#listsystemteamappusers) | **Get** /core/beta/systems/{systemId}/app-users | List System Team App Users
*SystemAPI* | [**RotateAgentToken**](docs/SystemAPI.md#rotateagenttoken) | **Post** /core/beta/systems/{systemId}/agent-tokens | Rotate Agent Token
*SystemAPI* | [**RunSystemAlertRule**](docs/SystemAPI.md#runsystemalertrule) | **Get** /core/beta/systems/{systemId}/alert-rules/{alertRuleId}/run | Run System Alert Rule
*SystemAPI* | [**SystemJWTSync**](docs/SystemAPI.md#systemjwtsync) | **Post** /core/beta/systems/{systemId}/jwt-sync | Re-sync JWTs of all accounts in this system
*SystemAPI* | [**UnAssignSystemTeamAppUser**](docs/SystemAPI.md#unassignsystemteamappuser) | **Delete** /core/beta/systems/{systemId}/app-users/{teamAppUserId} | Unassign Team App User from System
*SystemAPI* | [**UnmanageSystem**](docs/SystemAPI.md#unmanagesystem) | **Delete** /core/beta/systems/{systemId}/unmanage | Unmanage System
*SystemAPI* | [**UpdatePlatformComponents**](docs/SystemAPI.md#updateplatformcomponents) | **Patch** /core/beta/systems/{systemId}/platform-components | Update Platform Components for System
*SystemAPI* | [**UpdateSystem**](docs/SystemAPI.md#updatesystem) | **Patch** /core/beta/systems/{systemId} | Update System
*SystemAPI* | [**UpdateSystemAlertRule**](docs/SystemAPI.md#updatesystemalertrule) | **Patch** /core/beta/systems/{systemId}/alert-rules/{alertRuleId} | Update System Alert Rules
*TeamAPI* | [**CreateSystem**](docs/TeamAPI.md#createsystem) | **Post** /core/beta/teams/{teamId}/systems | Create System
*TeamAPI* | [**CreateTeamServiceAccount**](docs/TeamAPI.md#createteamserviceaccount) | **Post** /core/beta/teams/{teamId}/service-accounts | Create Team Service Account
*TeamAPI* | [**DeleteTeam**](docs/TeamAPI.md#deleteteam) | **Delete** /core/beta/teams/{teamId} | Delete Team
*TeamAPI* | [**GetTeam**](docs/TeamAPI.md#getteam) | **Get** /core/beta/teams/{teamId} | Get Team
*TeamAPI* | [**GetTeamLimits**](docs/TeamAPI.md#getteamlimits) | **Get** /core/beta/teams/{teamId}/team-limits | Get Team Limits
*TeamAPI* | [**ImportSystem**](docs/TeamAPI.md#importsystem) | **Post** /core/beta/teams/{teamId}/import-system | Import a System
*TeamAPI* | [**InviteAppUser**](docs/TeamAPI.md#inviteappuser) | **Post** /core/beta/teams/{teamId}/app-users/invitations | Invite App Users
*TeamAPI* | [**LeaveTeam**](docs/TeamAPI.md#leaveteam) | **Post** /core/beta/teams/{teamId}/app-users/leave | Leave Team
*TeamAPI* | [**ListTeamAppUsers**](docs/TeamAPI.md#listteamappusers) | **Get** /core/beta/teams/{teamId}/app-users | List App Users
*TeamAPI* | [**ListTeamInfoAppUsers**](docs/TeamAPI.md#listteaminfoappusers) | **Get** /core/beta/teams/{teamId}/info/app-users | List info of App Users in Team
*TeamAPI* | [**ListTeamServiceAccounts**](docs/TeamAPI.md#listteamserviceaccounts) | **Get** /core/beta/teams/{teamId}/service-accounts | List Team Service Accounts
*TeamAPI* | [**ListTeamSystems**](docs/TeamAPI.md#listteamsystems) | **Get** /core/beta/teams/{teamId}/systems | List Systems
*TeamAPI* | [**UnAssignTeamAppUser**](docs/TeamAPI.md#unassignteamappuser) | **Delete** /core/beta/teams/{teamId}/app-users/{appUserId} | Unassign App User from Team
*TeamAPI* | [**UpdateTeam**](docs/TeamAPI.md#updateteam) | **Patch** /core/beta/teams/{teamId} | Update Team
*TeamAPI* | [**UpdateTeamAppUser**](docs/TeamAPI.md#updateteamappuser) | **Patch** /core/beta/teams/{teamId}/app-users/{appUserId} | Update App User Team Assignment
*TeamServiceAccountAPI* | [**CreateTeamServiceAccountToken**](docs/TeamServiceAccountAPI.md#createteamserviceaccounttoken) | **Post** /core/beta/service-accounts/team/{serviceAccountId}/tokens | Create Acess Token for Team Service Account
*TeamServiceAccountAPI* | [**DeleteTeamServiceAccount**](docs/TeamServiceAccountAPI.md#deleteteamserviceaccount) | **Delete** /core/beta/service-accounts/team/{serviceAccountId} | Delete Team Service Account
*TeamServiceAccountAPI* | [**GetTeamServiceAccount**](docs/TeamServiceAccountAPI.md#getteamserviceaccount) | **Get** /core/beta/service-accounts/team/{serviceAccountId} | Get Team Service Account
*TeamServiceAccountAPI* | [**ListTeamServiceAccountTokens**](docs/TeamServiceAccountAPI.md#listteamserviceaccounttokens) | **Get** /core/beta/service-accounts/team/{serviceAccountId}/tokens | List Access Tokens for Team Service Account
*TeamServiceAccountAPI* | [**UpdateTeamServiceAccount**](docs/TeamServiceAccountAPI.md#updateteamserviceaccount) | **Patch** /core/beta/service-accounts/team/{serviceAccountId} | Update Team Service Account
*TeamServiceAccountTokensAPI* | [**DeleteTeamServiceAccountToken**](docs/TeamServiceAccountTokensAPI.md#deleteteamserviceaccounttoken) | **Delete** /core/beta/service-account-tokens/team/{tokenId} | Delete Team Service Account Token
*TeamServiceAccountTokensAPI* | [**GetTeamServiceAccountToken**](docs/TeamServiceAccountTokensAPI.md#getteamserviceaccounttoken) | **Get** /core/beta/service-account-tokens/team/{tokenId} | Get Team Service Account Token
*TeamServiceAccountTokensAPI* | [**UpdateTeamServiceAccountToken**](docs/TeamServiceAccountTokensAPI.md#updateteamserviceaccounttoken) | **Patch** /core/beta/service-account-tokens/team/{tokenId} | Update Team Service Account Token
*WorkloadsAPI* | [**GetWorkloadLimits**](docs/WorkloadsAPI.md#getworkloadlimits) | **Get** /workloads/alpha/{accountId}/limits | Get Workloads limits


## Documentation For Models

 - [AcceptTermsResponse](docs/AcceptTermsResponse.md)
 - [Account](docs/Account.md)
 - [AccountChartType](docs/AccountChartType.md)
 - [AccountClaims](docs/AccountClaims.md)
 - [AccountClaimsInfo](docs/AccountClaimsInfo.md)
 - [AccountConnectionsListResponse](docs/AccountConnectionsListResponse.md)
 - [AccountCreateRequest](docs/AccountCreateRequest.md)
 - [AccountExportResponse](docs/AccountExportResponse.md)
 - [AccountInfo](docs/AccountInfo.md)
 - [AccountInfoListResponse](docs/AccountInfoListResponse.md)
 - [AccountJWTSettings](docs/AccountJWTSettings.md)
 - [AccountJWTSettingsPatch](docs/AccountJWTSettingsPatch.md)
 - [AccountLimits](docs/AccountLimits.md)
 - [AccountListResponse](docs/AccountListResponse.md)
 - [AccountMetrics](docs/AccountMetrics.md)
 - [AccountOverviewResponse](docs/AccountOverviewResponse.md)
 - [AccountOverviewResponseMetrics](docs/AccountOverviewResponseMetrics.md)
 - [AccountUpdateRequest](docs/AccountUpdateRequest.md)
 - [AccountViewResponse](docs/AccountViewResponse.md)
 - [AccountsOverviewListResponse](docs/AccountsOverviewListResponse.md)
 - [AckPolicy](docs/AckPolicy.md)
 - [Activation](docs/Activation.md)
 - [ActivationClaims](docs/ActivationClaims.md)
 - [AddRequest](docs/AddRequest.md)
 - [AddResponse](docs/AddResponse.md)
 - [AdminAppUserCreateRequest](docs/AdminAppUserCreateRequest.md)
 - [AdminAppUserCreateResponse](docs/AdminAppUserCreateResponse.md)
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
 - [AlertRuleThresholdMetric](docs/AlertRuleThresholdMetric.md)
 - [AlertRuleType](docs/AlertRuleType.md)
 - [AlertRuleUpdateRequest](docs/AlertRuleUpdateRequest.md)
 - [AlertRuleViewResponse](docs/AlertRuleViewResponse.md)
 - [AlertStatus](docs/AlertStatus.md)
 - [AlertViewResponse](docs/AlertViewResponse.md)
 - [ApexCoordinate](docs/ApexCoordinate.md)
 - [ApexSeries](docs/ApexSeries.md)
 - [AppPolicy](docs/AppPolicy.md)
 - [AppPolicyGroup](docs/AppPolicyGroup.md)
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
 - [AuthCalloutAddTargetAccountRequest](docs/AuthCalloutAddTargetAccountRequest.md)
 - [AuthCalloutAddUserRequest](docs/AuthCalloutAddUserRequest.md)
 - [AuthCalloutAuthenticatorComponentConfig](docs/AuthCalloutAuthenticatorComponentConfig.md)
 - [AuthCalloutAuthenticatorConfigViewResponse](docs/AuthCalloutAuthenticatorConfigViewResponse.md)
 - [AuthCalloutAuthenticatorUpdateRequest](docs/AuthCalloutAuthenticatorUpdateRequest.md)
 - [AuthCalloutAuthenticatorViewResponse](docs/AuthCalloutAuthenticatorViewResponse.md)
 - [AuthCalloutAuthenticatorsListResponse](docs/AuthCalloutAuthenticatorsListResponse.md)
 - [AuthCalloutComponentConnectConfig](docs/AuthCalloutComponentConnectConfig.md)
 - [AuthCalloutConfig](docs/AuthCalloutConfig.md)
 - [AuthCalloutConfigListResponse](docs/AuthCalloutConfigListResponse.md)
 - [AuthCalloutEnableRequest](docs/AuthCalloutEnableRequest.md)
 - [AuthCalloutTargetAccountListResponse](docs/AuthCalloutTargetAccountListResponse.md)
 - [AuthCalloutTargetAccountViewResponse](docs/AuthCalloutTargetAccountViewResponse.md)
 - [AuthCalloutUserViewResponse](docs/AuthCalloutUserViewResponse.md)
 - [AuthCalloutUsersListResponse](docs/AuthCalloutUsersListResponse.md)
 - [AuthCalloutViewResponse](docs/AuthCalloutViewResponse.md)
 - [AuthzRequest](docs/AuthzRequest.md)
 - [AuthzResponse](docs/AuthzResponse.md)
 - [AvailableChartVersion](docs/AvailableChartVersion.md)
 - [AvailableComponentVersionsRequest](docs/AvailableComponentVersionsRequest.md)
 - [BulkShareRequest](docs/BulkShareRequest.md)
 - [BulkShareTarget](docs/BulkShareTarget.md)
 - [BulkShareType](docs/BulkShareType.md)
 - [ClaimsData](docs/ClaimsData.md)
 - [ClusterChartType](docs/ClusterChartType.md)
 - [ClusterInfo](docs/ClusterInfo.md)
 - [ConnInfo](docs/ConnInfo.md)
 - [ConnectorComponentConstraint](docs/ConnectorComponentConstraint.md)
 - [ConnectorComponentField](docs/ConnectorComponentField.md)
 - [ConnectorComponentFieldKind](docs/ConnectorComponentFieldKind.md)
 - [ConnectorComponentFieldType](docs/ConnectorComponentFieldType.md)
 - [ConnectorComponentKind](docs/ConnectorComponentKind.md)
 - [ConnectorComponentRange](docs/ConnectorComponentRange.md)
 - [ConnectorComponentStatus](docs/ConnectorComponentStatus.md)
 - [ConnectorComponentSummary](docs/ConnectorComponentSummary.md)
 - [ConnectorComponentSummaryListResponse](docs/ConnectorComponentSummaryListResponse.md)
 - [ConnectorComponentViewResponse](docs/ConnectorComponentViewResponse.md)
 - [ConnectorConfig](docs/ConnectorConfig.md)
 - [ConnectorConsumer](docs/ConnectorConsumer.md)
 - [ConnectorConsumerCore](docs/ConnectorConsumerCore.md)
 - [ConnectorConsumerKv](docs/ConnectorConsumerKv.md)
 - [ConnectorConsumerPatch](docs/ConnectorConsumerPatch.md)
 - [ConnectorConsumerPatchCore](docs/ConnectorConsumerPatchCore.md)
 - [ConnectorConsumerPatchKv](docs/ConnectorConsumerPatchKv.md)
 - [ConnectorConsumerPatchStream](docs/ConnectorConsumerPatchStream.md)
 - [ConnectorConsumerStream](docs/ConnectorConsumerStream.md)
 - [ConnectorCreateRequest](docs/ConnectorCreateRequest.md)
 - [ConnectorDeployRequest](docs/ConnectorDeployRequest.md)
 - [ConnectorDeploymentViewResponse](docs/ConnectorDeploymentViewResponse.md)
 - [ConnectorExecutionLog](docs/ConnectorExecutionLog.md)
 - [ConnectorExecutionLogListResponse](docs/ConnectorExecutionLogListResponse.md)
 - [ConnectorInstanceEvent](docs/ConnectorInstanceEvent.md)
 - [ConnectorInstanceEventListResponse](docs/ConnectorInstanceEventListResponse.md)
 - [ConnectorInstanceEventType](docs/ConnectorInstanceEventType.md)
 - [ConnectorInstanceListResponse](docs/ConnectorInstanceListResponse.md)
 - [ConnectorInstanceViewResponse](docs/ConnectorInstanceViewResponse.md)
 - [ConnectorKind](docs/ConnectorKind.md)
 - [ConnectorListResponse](docs/ConnectorListResponse.md)
 - [ConnectorMetricChartResponse](docs/ConnectorMetricChartResponse.md)
 - [ConnectorMetricChartTypes](docs/ConnectorMetricChartTypes.md)
 - [ConnectorMetricEndpoint](docs/ConnectorMetricEndpoint.md)
 - [ConnectorMetricEndpointPatch](docs/ConnectorMetricEndpointPatch.md)
 - [ConnectorNatsConfig](docs/ConnectorNatsConfig.md)
 - [ConnectorNatsConfigPatch](docs/ConnectorNatsConfigPatch.md)
 - [ConnectorProducer](docs/ConnectorProducer.md)
 - [ConnectorProducerKv](docs/ConnectorProducerKv.md)
 - [ConnectorProducerPatch](docs/ConnectorProducerPatch.md)
 - [ConnectorSink](docs/ConnectorSink.md)
 - [ConnectorSinkPatch](docs/ConnectorSinkPatch.md)
 - [ConnectorSource](docs/ConnectorSource.md)
 - [ConnectorSourcePatch](docs/ConnectorSourcePatch.md)
 - [ConnectorSteps](docs/ConnectorSteps.md)
 - [ConnectorStepsPatch](docs/ConnectorStepsPatch.md)
 - [ConnectorSummary](docs/ConnectorSummary.md)
 - [ConnectorSummaryInstances](docs/ConnectorSummaryInstances.md)
 - [ConnectorTransformer](docs/ConnectorTransformer.md)
 - [ConnectorTransformerComposite](docs/ConnectorTransformerComposite.md)
 - [ConnectorTransformerKind](docs/ConnectorTransformerKind.md)
 - [ConnectorTransformerMapping](docs/ConnectorTransformerMapping.md)
 - [ConnectorTransformerPatch](docs/ConnectorTransformerPatch.md)
 - [ConnectorTransformerPatchComposite](docs/ConnectorTransformerPatchComposite.md)
 - [ConnectorTransformerPatchMapping](docs/ConnectorTransformerPatchMapping.md)
 - [ConnectorTransformerPatchService](docs/ConnectorTransformerPatchService.md)
 - [ConnectorTransformerService](docs/ConnectorTransformerService.md)
 - [ConnectorUpdateRequest](docs/ConnectorUpdateRequest.md)
 - [ConnectorViewResponse](docs/ConnectorViewResponse.md)
 - [ConnectorsPlatformComponentConfig](docs/ConnectorsPlatformComponentConfig.md)
 - [Connz](docs/Connz.md)
 - [ConsumerConfig](docs/ConsumerConfig.md)
 - [ConsumerInfo](docs/ConsumerInfo.md)
 - [DataLimits](docs/DataLimits.md)
 - [DataStats](docs/DataStats.md)
 - [DeliverPolicy](docs/DeliverPolicy.md)
 - [DiscardPolicy](docs/DiscardPolicy.md)
 - [Export](docs/Export.md)
 - [ExportPatch](docs/ExportPatch.md)
 - [ExportType](docs/ExportType.md)
 - [ExternalAuthorization](docs/ExternalAuthorization.md)
 - [ExternalAuthorizationPatch](docs/ExternalAuthorizationPatch.md)
 - [ExternalStream](docs/ExternalStream.md)
 - [GatewayStat](docs/GatewayStat.md)
 - [GenericFields](docs/GenericFields.md)
 - [GetResponse](docs/GetResponse.md)
 - [HttpGatewayPlatformComponentConfig](docs/HttpGatewayPlatformComponentConfig.md)
 - [Import](docs/Import.md)
 - [ImportPatch](docs/ImportPatch.md)
 - [Info](docs/Info.md)
 - [InvitationDecisionRequest](docs/InvitationDecisionRequest.md)
 - [InvitationListResponse](docs/InvitationListResponse.md)
 - [JSApiError](docs/JSApiError.md)
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
 - [JwtSyncStatus](docs/JwtSyncStatus.md)
 - [LeafInfo](docs/LeafInfo.md)
 - [Leafz](docs/Leafz.md)
 - [ListResponse](docs/ListResponse.md)
 - [LostStreamData](docs/LostStreamData.md)
 - [MetaClusterInfo](docs/MetaClusterInfo.md)
 - [NatsAlertingConfig](docs/NatsAlertingConfig.md)
 - [NatsAlertingConfigPatch](docs/NatsAlertingConfigPatch.md)
 - [NatsCluster](docs/NatsCluster.md)
 - [NatsClusterListResponse](docs/NatsClusterListResponse.md)
 - [NatsCreateUserJwtSettings](docs/NatsCreateUserJwtSettings.md)
 - [NatsLimits](docs/NatsLimits.md)
 - [NatsServer](docs/NatsServer.md)
 - [NatsServerInfoListResponse](docs/NatsServerInfoListResponse.md)
 - [NatsServerListResponse](docs/NatsServerListResponse.md)
 - [NatsUserConnectionsListResponse](docs/NatsUserConnectionsListResponse.md)
 - [NatsUserCopyRequest](docs/NatsUserCopyRequest.md)
 - [NatsUserCreateRequest](docs/NatsUserCreateRequest.md)
 - [NatsUserHTTPGWTokenCreateReply](docs/NatsUserHTTPGWTokenCreateReply.md)
 - [NatsUserInfo](docs/NatsUserInfo.md)
 - [NatsUserIssuanceEvent](docs/NatsUserIssuanceEvent.md)
 - [NatsUserIssuanceEventType](docs/NatsUserIssuanceEventType.md)
 - [NatsUserIssuanceInfo](docs/NatsUserIssuanceInfo.md)
 - [NatsUserIssuanceStatus](docs/NatsUserIssuanceStatus.md)
 - [NatsUserIssuanceViewResponse](docs/NatsUserIssuanceViewResponse.md)
 - [NatsUserIssuancesListResponse](docs/NatsUserIssuancesListResponse.md)
 - [NatsUserJwtSettings](docs/NatsUserJwtSettings.md)
 - [NatsUserJwtSettingsPatch](docs/NatsUserJwtSettingsPatch.md)
 - [NatsUserListResponse](docs/NatsUserListResponse.md)
 - [NatsUserRevocationRequest](docs/NatsUserRevocationRequest.md)
 - [NatsUserRevocationViewResponse](docs/NatsUserRevocationViewResponse.md)
 - [NatsUserUpdateRequest](docs/NatsUserUpdateRequest.md)
 - [NatsUserViewResponse](docs/NatsUserViewResponse.md)
 - [ObjectStoreConfig](docs/ObjectStoreConfig.md)
 - [OidcProviderAddRequest](docs/OidcProviderAddRequest.md)
 - [OidcProviderDefaultsListResponse](docs/OidcProviderDefaultsListResponse.md)
 - [OidcProviderDefaultsResponse](docs/OidcProviderDefaultsResponse.md)
 - [OidcProviderListResponse](docs/OidcProviderListResponse.md)
 - [OidcProviderUpdateRequest](docs/OidcProviderUpdateRequest.md)
 - [OidcProviderViewResponse](docs/OidcProviderViewResponse.md)
 - [Operator](docs/Operator.md)
 - [OperatorClaims](docs/OperatorClaims.md)
 - [OperatorExport](docs/OperatorExport.md)
 - [OperatorLimits](docs/OperatorLimits.md)
 - [OperatorLimitsPatch](docs/OperatorLimitsPatch.md)
 - [PeerInfo](docs/PeerInfo.md)
 - [Permission](docs/Permission.md)
 - [PermissionPatch](docs/PermissionPatch.md)
 - [Permissions](docs/Permissions.md)
 - [Placement](docs/Placement.md)
 - [PlatformComponent](docs/PlatformComponent.md)
 - [PlatformComponentConfigs](docs/PlatformComponentConfigs.md)
 - [PlatformComponentConnectConfig](docs/PlatformComponentConnectConfig.md)
 - [PlatformComponentTokenViewResponse](docs/PlatformComponentTokenViewResponse.md)
 - [PlatformComponentType](docs/PlatformComponentType.md)
 - [PlatformComponentsConfigSchema](docs/PlatformComponentsConfigSchema.md)
 - [PlatformComponentsConnectRequest](docs/PlatformComponentsConnectRequest.md)
 - [PlatformComponentsConnectViewResponse](docs/PlatformComponentsConnectViewResponse.md)
 - [PlatformComponentsUpdateRequest](docs/PlatformComponentsUpdateRequest.md)
 - [PlatformComponentsViewResponse](docs/PlatformComponentsViewResponse.md)
 - [PolicyListResponse](docs/PolicyListResponse.md)
 - [PromSampleValue](docs/PromSampleValue.md)
 - [RePublish](docs/RePublish.md)
 - [RemoveRequest](docs/RemoveRequest.md)
 - [ReplayPolicy](docs/ReplayPolicy.md)
 - [ResponsePermission](docs/ResponsePermission.md)
 - [ResponsePermissionPatch](docs/ResponsePermissionPatch.md)
 - [ResponseType](docs/ResponseType.md)
 - [RetentionPolicy](docs/RetentionPolicy.md)
 - [RoleListResponse](docs/RoleListResponse.md)
 - [RouteStat](docs/RouteStat.md)
 - [S2Compression](docs/S2Compression.md)
 - [SDKubeLimit](docs/SDKubeLimit.md)
 - [Schema](docs/Schema.md)
 - [SchemaRegistryCompatPolicy](docs/SchemaRegistryCompatPolicy.md)
 - [SchemaRegistryFormat](docs/SchemaRegistryFormat.md)
 - [SchemaRegistryPlatformComponentConfig](docs/SchemaRegistryPlatformComponentConfig.md)
 - [ScraperComponentConfig](docs/ScraperComponentConfig.md)
 - [SequenceInfo](docs/SequenceInfo.md)
 - [ServerInfo](docs/ServerInfo.md)
 - [ServerStats](docs/ServerStats.md)
 - [ServerStatsMsg](docs/ServerStatsMsg.md)
 - [ServiceAccountCreateRequest](docs/ServiceAccountCreateRequest.md)
 - [ServiceAccountListResponse](docs/ServiceAccountListResponse.md)
 - [ServiceAccountResourceScope](docs/ServiceAccountResourceScope.md)
 - [ServiceAccountScope](docs/ServiceAccountScope.md)
 - [ServiceAccountUpdateRequest](docs/ServiceAccountUpdateRequest.md)
 - [ServiceAccountViewResponse](docs/ServiceAccountViewResponse.md)
 - [ServiceLatency](docs/ServiceLatency.md)
 - [ServiceLatencyPatch](docs/ServiceLatencyPatch.md)
 - [SigningKeyExport](docs/SigningKeyExport.md)
 - [SigningKeyGroupCopyRequest](docs/SigningKeyGroupCopyRequest.md)
 - [SigningKeyGroupCreateRequest](docs/SigningKeyGroupCreateRequest.md)
 - [SigningKeyGroupCreateResponse](docs/SigningKeyGroupCreateResponse.md)
 - [SigningKeyGroupExport](docs/SigningKeyGroupExport.md)
 - [SigningKeyGroupListResponse](docs/SigningKeyGroupListResponse.md)
 - [SigningKeyGroupUpdateRequest](docs/SigningKeyGroupUpdateRequest.md)
 - [SigningKeyGroupViewResponse](docs/SigningKeyGroupViewResponse.md)
 - [SigningKeyListResponse](docs/SigningKeyListResponse.md)
 - [SigningKeyRotateResponse](docs/SigningKeyRotateResponse.md)
 - [SigningKeyScope](docs/SigningKeyScope.md)
 - [SigningKeyScopeType](docs/SigningKeyScopeType.md)
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
 - [SystemConnectionType](docs/SystemConnectionType.md)
 - [SystemConnectionsListResponse](docs/SystemConnectionsListResponse.md)
 - [SystemCreateRequest](docs/SystemCreateRequest.md)
 - [SystemDirectConnectionOpts](docs/SystemDirectConnectionOpts.md)
 - [SystemDirectConnectionOptsPatch](docs/SystemDirectConnectionOptsPatch.md)
 - [SystemExportResponse](docs/SystemExportResponse.md)
 - [SystemImportRequest](docs/SystemImportRequest.md)
 - [SystemInfo](docs/SystemInfo.md)
 - [SystemLimitsResponse](docs/SystemLimitsResponse.md)
 - [SystemListResponse](docs/SystemListResponse.md)
 - [SystemLogForwardingOpts](docs/SystemLogForwardingOpts.md)
 - [SystemLogForwardingOptsPatch](docs/SystemLogForwardingOptsPatch.md)
 - [SystemLogLevel](docs/SystemLogLevel.md)
 - [SystemManagedBy](docs/SystemManagedBy.md)
 - [SystemState](docs/SystemState.md)
 - [SystemUpdateRequest](docs/SystemUpdateRequest.md)
 - [SystemUserImportRequest](docs/SystemUserImportRequest.md)
 - [SystemViewResponse](docs/SystemViewResponse.md)
 - [TLSMode](docs/TLSMode.md)
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
 - [Update](docs/Update.md)
 - [UpdateResponse](docs/UpdateResponse.md)
 - [User](docs/User.md)
 - [UserClaims](docs/UserClaims.md)
 - [UserExport](docs/UserExport.md)
 - [UserLimits](docs/UserLimits.md)
 - [UserPermissionLimits](docs/UserPermissionLimits.md)
 - [UserPermissionLimitsPatch](docs/UserPermissionLimitsPatch.md)
 - [ValidateRequest](docs/ValidateRequest.md)
 - [VersionResponse](docs/VersionResponse.md)
 - [WeightedMapping](docs/WeightedMapping.md)
 - [WorkloadAuctionRequest](docs/WorkloadAuctionRequest.md)
 - [WorkloadAuctionResponse](docs/WorkloadAuctionResponse.md)
 - [WorkloadField](docs/WorkloadField.md)
 - [WorkloadFieldConstraint](docs/WorkloadFieldConstraint.md)
 - [WorkloadFieldContstraintRange](docs/WorkloadFieldContstraintRange.md)
 - [WorkloadFieldKind](docs/WorkloadFieldKind.md)
 - [WorkloadFieldType](docs/WorkloadFieldType.md)
 - [WorkloadLifecycle](docs/WorkloadLifecycle.md)
 - [WorkloadLimits](docs/WorkloadLimits.md)
 - [WorkloadLimitsResponse](docs/WorkloadLimitsResponse.md)
 - [WorkloadMetricChartResponse](docs/WorkloadMetricChartResponse.md)
 - [WorkloadMetricChartTypes](docs/WorkloadMetricChartTypes.md)
 - [WorkloadStartRequest](docs/WorkloadStartRequest.md)
 - [WorkloadStartResponse](docs/WorkloadStartResponse.md)
 - [WorkloadState](docs/WorkloadState.md)
 - [WorkloadStateLimit](docs/WorkloadStateLimit.md)
 - [WorkloadStateLimits](docs/WorkloadStateLimits.md)
 - [WorkloadStopResponse](docs/WorkloadStopResponse.md)
 - [WorkloadSummary](docs/WorkloadSummary.md)
 - [WorkloadTags](docs/WorkloadTags.md)
 - [WorkloadsListResponse](docs/WorkloadsListResponse.md)
 - [WorkloadsPlatformComponentConfig](docs/WorkloadsPlatformComponentConfig.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), syncp.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```

### sessionAuth

- **Type**: API key
- **API key parameter name**: control_plane_session
- **Location**: 

Note, each API key must be added to a map of `map[string]APIKey` where the key is: control_plane_session and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		syncp.ContextAPIKeys,
		map[string]syncp.APIKey{
			"control_plane_session": {Key: "API_KEY_STRING"},
		},
	)
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



