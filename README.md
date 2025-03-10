# Go API client for authlete

<div class=\"min-h-screen bg-gray-100 dark:bg-gray-900 text-gray-900 dark:text-gray-100 p-6\">
  <div class=\"flex justify-end mb-4\">
    <label for=\"theme-toggle\" class=\"flex items-center cursor-pointer\">
      <div class=\"relative\">Dark mode:
        <input type=\"checkbox\" id=\"theme-toggle\" class=\"sr-only\" onchange=\"toggleTheme()\">
        <div class=\"block bg-gray-600 w-14 h-8 rounded-full\"></div>
        <div class=\"dot absolute left-1 top-1 bg-white w-6 h-6 rounded-full transition\"></div>
      </div>
    </label>
  </div>
  <header class=\"bg-green-500 dark:bg-green-700 p-4 rounded-lg text-white text-center\">
    <p>
      Welcome to the <strong>Authlete API documentation</strong>. Authlete is an <strong>API-first service</strong>
      where every aspect of the platform is configurable via API. This explorer provides a convenient way to
      authenticate and interact with the API, allowing you to see Authlete in action quickly. 🚀
    </p>
    <p>
      At a high level, the Authlete API is grouped into two categories:
    </p>
    <ul class=\"list-disc list-inside\">
      <li><strong>Management APIs</strong>: Enable you to manage services and clients. 🔧</li>
      <li><strong>Runtime APIs</strong>: Allow you to build your own Authorization Servers or Verifiable Credential (VC)
        issuers. 🔐</li>
    </ul>
    <p>All API endpoints are secured using access tokens issued by Authlete's Identity Provider (IdP). If you already
      have an Authlete account, simply use the <em>Get Token</em> option on the Authentication page to log in and obtain
      an access token for API usage. If you don't have an account yet, <a href=\"https://console.authlete.com/register\">sign up
        here</a> to get started.</p>
  </header>
  <main>
    <section id=\"api-servers\" class=\"mb-10\">
      <h2 class=\"text-2xl font-semibold mb-4\">🌐 API Servers</h2>
      <p>Authlete is a global service with clusters available in multiple regions across the world.</p>
      <p>Currently, our service is available in the following regions:</p>
      <div class=\"grid grid-cols-2 gap-4\">
        <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">
          <p class=\"text-center font-semibold\">🇺🇸 US</p>
        </div>
        <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">
          <p class=\"text-center font-semibold\">🇯🇵 JP</p>
        </div>
        <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">
          <p class=\"text-center font-semibold\">🇪🇺 EU</p>
        </div>
        <div class=\"p-4 bg-white dark:bg-gray-800 rounded-lg shadow\">
          <p class=\"text-center font-semibold\">🇧🇷 Brazil</p>
        </div>
      </div>
      <p>Our customers can host their data in the region that best meets their requirements.</p>
      <a href=\"#servers\" class=\"block mt-4 text-green-500 dark:text-green-300 hover:underline text-center\">Select your
        preferred server</a>
    </section>
    <section id=\"authentication\" class=\"mb-10\">
      <h2 class=\"text-2xl font-semibold mb-4\">🔑 Authentication</h2>
      <p>The API Explorer requires an access token to call the API.</p>
      <p>You can create the access token from the <a href=\"https://console.authlete.com\">Authlete Management Console</a> and set it in the HTTP Bearer section of Authentication page.</p>
      <p>Alternatively, if you have an Authlete account, the API Explorer can log you in with your Authlete account and
        automatically acquire the required access token.</p>
      <div class=\"theme-admonition theme-admonition-warning admonition_o5H7 alert alert--warning\">
        <div class=\"admonitionContent_Knsx\">
          <p>⚠️ <strong>Important Note:</strong> When the API Explorer acquires the token after login, the access tokens
            will have the same permissions as the user who logs in as part of this flow.</p>
        </div>
      </div>
      <a href=\"#auth\" class=\"block mt-4 text-green-500 dark:text-green-300 hover:underline text-center\">Setup your
        access token</a>
    </section>
    <section id=\"tutorials\" class=\"mb-10\">
      <h2 class=\"text-2xl font-semibold mb-4\">🎓 Tutorials</h2>
      <p>If you have successfully tested the API from the API Console and want to take the next step of integrating the
        API into your application, or if you want to see a sample using Authlete APIs, follow the links below. These
        resources will help you understand key concepts and how to integrate Authlete API into your applications.</p>
      <div class=\"mt-4\">
        <a href=\"https://www.authlete.com/developers/getting_started/\"
          class=\"block text-green-500 dark:text-green-300 font-bold hover:underline mb-2\">🚀 Getting Started with
          Authlete</a>
          </br>
        <a href=\"https://www.authlete.com/developers/tutorial/signup/\"
          class=\"block text-green-500 dark:text-green-300 font-bold hover:underline\">🔑 From Sign-Up to the First API
          Request</a>
      </div>
    </section>
    <section id=\"support\" class=\"mb-10\">
      <h2 class=\"text-2xl font-semibold mb-4\">🛠 Contact Us</h2>
      <p>If you have any questions or need assistance, our team is here to help.</p>
      <a href=\"https://www.authlete.com/contact/\"
        class=\"block mt-4 text-green-500 dark:text-green-300 font-bold hover:underline\">Contact Page</a>
    </section>
  </main>
</div>


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.0
- Package version: 3.0
- Generator version: 7.11.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import authlete "github.com/authlete/openapi-for-go/v3"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `authlete.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), authlete.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `authlete.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), authlete.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `authlete.ContextOperationServerIndices` and `authlete.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), authlete.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), authlete.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://us.authlete.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthorizationEndpointAPI* | [**ApiServiceIdAuthAuthorizationTicketInfoGet**](docs/AuthorizationEndpointAPI.md#apiserviceidauthauthorizationticketinfoget) | **Get** /api/{serviceId}/auth/authorization/ticket/info | Get Ticket Information
*AuthorizationEndpointAPI* | [**ApiServiceIdAuthAuthorizationTicketUpdatePost**](docs/AuthorizationEndpointAPI.md#apiserviceidauthauthorizationticketupdatepost) | **Post** /api/{serviceId}/auth/authorization/ticket/update | Update Ticket Information
*AuthorizationEndpointAPI* | [**AuthAuthorizationApi**](docs/AuthorizationEndpointAPI.md#authauthorizationapi) | **Post** /api/{serviceId}/auth/authorization | Process Authorization Request
*AuthorizationEndpointAPI* | [**AuthAuthorizationFailApi**](docs/AuthorizationEndpointAPI.md#authauthorizationfailapi) | **Post** /api/{serviceId}/auth/authorization/fail | Fail Authorization Request
*AuthorizationEndpointAPI* | [**AuthAuthorizationIssueApi**](docs/AuthorizationEndpointAPI.md#authauthorizationissueapi) | **Post** /api/{serviceId}/auth/authorization/issue | Issue Authorization Response
*CIBAAPI* | [**BackchannelAuthenticationApi**](docs/CIBAAPI.md#backchannelauthenticationapi) | **Post** /api/{serviceId}/backchannel/authentication | Process Backchannel Authentication Request
*CIBAAPI* | [**BackchannelAuthenticationCompleteApi**](docs/CIBAAPI.md#backchannelauthenticationcompleteapi) | **Post** /api/{serviceId}/backchannel/authentication/complete | Complete Backchannel Authentication
*CIBAAPI* | [**BackchannelAuthenticationFailApi**](docs/CIBAAPI.md#backchannelauthenticationfailapi) | **Post** /api/{serviceId}/backchannel/authentication/fail | Fail Backchannel Authentication Request
*CIBAAPI* | [**BackchannelAuthenticationIssueApi**](docs/CIBAAPI.md#backchannelauthenticationissueapi) | **Post** /api/{serviceId}/backchannel/authentication/issue | Issue Backchannel Authentication Response
*ClientManagementAPI* | [**ClientAuthorizationDeleteApi**](docs/ClientManagementAPI.md#clientauthorizationdeleteapi) | **Delete** /api/{serviceId}/client/authorization/delete/{clientId} | Delete Client Tokens
*ClientManagementAPI* | [**ClientAuthorizationGetListApi**](docs/ClientManagementAPI.md#clientauthorizationgetlistapi) | **Get** /api/{serviceId}/client/authorization/get/list | Get Authorized Applications
*ClientManagementAPI* | [**ClientAuthorizationUpdateApi**](docs/ClientManagementAPI.md#clientauthorizationupdateapi) | **Post** /api/{serviceId}/client/authorization/update/{clientId} | Update Client Tokens
*ClientManagementAPI* | [**ClientCreateApi**](docs/ClientManagementAPI.md#clientcreateapi) | **Post** /api/{serviceId}/client/create | Create Client
*ClientManagementAPI* | [**ClientDeleteApi**](docs/ClientManagementAPI.md#clientdeleteapi) | **Delete** /api/{serviceId}/client/delete/{clientId} | Delete Client ⚡
*ClientManagementAPI* | [**ClientExtensionRequestablesScopesDeleteApi**](docs/ClientManagementAPI.md#clientextensionrequestablesscopesdeleteapi) | **Delete** /api/{serviceId}/client/extension/requestable_scopes/delete/{clientId} | Delete Requestable Scopes
*ClientManagementAPI* | [**ClientExtensionRequestablesScopesGetApi**](docs/ClientManagementAPI.md#clientextensionrequestablesscopesgetapi) | **Get** /api/{serviceId}/client/extension/requestable_scopes/get/{clientId} | Get Requestable Scopes
*ClientManagementAPI* | [**ClientExtensionRequestablesScopesUpdateApi**](docs/ClientManagementAPI.md#clientextensionrequestablesscopesupdateapi) | **Put** /api/{serviceId}/client/extension/requestable_scopes/update/{clientId} | Update Requestable Scopes
*ClientManagementAPI* | [**ClientFlagUpdateApi**](docs/ClientManagementAPI.md#clientflagupdateapi) | **Post** /api/{serviceId}/client/lock_flag/update/{clientIdentifier} | Update Client Lock
*ClientManagementAPI* | [**ClientGetApi**](docs/ClientManagementAPI.md#clientgetapi) | **Get** /api/{serviceId}/client/get/{clientId} | Get Client
*ClientManagementAPI* | [**ClientGetListApi**](docs/ClientManagementAPI.md#clientgetlistapi) | **Get** /api/{serviceId}/client/get/list | List Clients
*ClientManagementAPI* | [**ClientGrantedScopesDeleteApi**](docs/ClientManagementAPI.md#clientgrantedscopesdeleteapi) | **Delete** /api/{serviceId}/client/granted_scopes/delete/{clientId} | Delete Granted Scopes
*ClientManagementAPI* | [**ClientGrantedScopesGetApi**](docs/ClientManagementAPI.md#clientgrantedscopesgetapi) | **Get** /api/{serviceId}/client/granted_scopes/get/{clientId} | Get Granted Scopes
*ClientManagementAPI* | [**ClientSecretRefreshApi**](docs/ClientManagementAPI.md#clientsecretrefreshapi) | **Get** /api/{serviceId}/client/secret/refresh/{clientIdentifier} | Rotate Client Secret
*ClientManagementAPI* | [**ClientSecretUpdateApi**](docs/ClientManagementAPI.md#clientsecretupdateapi) | **Post** /api/{serviceId}/client/secret/update/{clientIdentifier} | Update Client Secret
*ClientManagementAPI* | [**ClientUpdateApi**](docs/ClientManagementAPI.md#clientupdateapi) | **Post** /api/{serviceId}/client/update/{clientId} | Update Client
*DeviceFlowAPI* | [**DeviceAuthorizationApi**](docs/DeviceFlowAPI.md#deviceauthorizationapi) | **Post** /api/{serviceId}/device/authorization | Process Device Authorization Request
*DeviceFlowAPI* | [**DeviceCompleteApi**](docs/DeviceFlowAPI.md#devicecompleteapi) | **Post** /api/{serviceId}/device/complete | Complete Device Authorization
*DeviceFlowAPI* | [**DeviceVerificationApi**](docs/DeviceFlowAPI.md#deviceverificationapi) | **Post** /api/{serviceId}/device/verification | Process Device Verification Request
*DynamicClientRegistrationAPI* | [**ClientRegistrationApi**](docs/DynamicClientRegistrationAPI.md#clientregistrationapi) | **Post** /api/{serviceId}/client/registration | Register Client
*DynamicClientRegistrationAPI* | [**ClientRegistrationDeleteApi**](docs/DynamicClientRegistrationAPI.md#clientregistrationdeleteapi) | **Post** /api/{serviceId}/client/registration/delete | Delete Client
*DynamicClientRegistrationAPI* | [**ClientRegistrationGetApi**](docs/DynamicClientRegistrationAPI.md#clientregistrationgetapi) | **Post** /api/{serviceId}/client/registration/get | Get Client
*DynamicClientRegistrationAPI* | [**ClientRegistrationUpdateApi**](docs/DynamicClientRegistrationAPI.md#clientregistrationupdateapi) | **Post** /api/{serviceId}/client/registration/update | Update Client
*FederationEndpointAPI* | [**FederationConfigurationApi**](docs/FederationEndpointAPI.md#federationconfigurationapi) | **Post** /api/{serviceId}/federation/configuration | Process Entity Configuration Request
*FederationEndpointAPI* | [**FederationRegistrationApi**](docs/FederationEndpointAPI.md#federationregistrationapi) | **Post** /api/{serviceId}/federation/registration | Process Federation Registration Request
*GrantManagementEndpointAPI* | [**GrantMApi**](docs/GrantManagementEndpointAPI.md#grantmapi) | **Post** /api/{serviceId}/gm | Process Grant Management Request
*HardwareSecurityKeyAPI* | [**HskCreateApi**](docs/HardwareSecurityKeyAPI.md#hskcreateapi) | **Post** /api/{serviceId}/hsk/create | Create Security Key
*HardwareSecurityKeyAPI* | [**HskDeleteApi**](docs/HardwareSecurityKeyAPI.md#hskdeleteapi) | **Delete** /api/{serviceId}/hsk/delete/{handle} | Delete Security Key
*HardwareSecurityKeyAPI* | [**HskGetApi**](docs/HardwareSecurityKeyAPI.md#hskgetapi) | **Get** /api/{serviceId}/hsk/get/{handle} | Get Security Key
*HardwareSecurityKeyAPI* | [**HskGetListApi**](docs/HardwareSecurityKeyAPI.md#hskgetlistapi) | **Get** /api/{serviceId}/hsk/get/list | List Security Keys
*IntrospectionEndpointAPI* | [**AuthIntrospectionApi**](docs/IntrospectionEndpointAPI.md#authintrospectionapi) | **Post** /api/{serviceId}/auth/introspection | Process Introspection Request
*IntrospectionEndpointAPI* | [**AuthIntrospectionStandardApi**](docs/IntrospectionEndpointAPI.md#authintrospectionstandardapi) | **Post** /api/{serviceId}/auth/introspection/standard | Process OAuth 2.0 Introspection Request
*JWKSetEndpointAPI* | [**ServiceJwksGetApi**](docs/JWKSetEndpointAPI.md#servicejwksgetapi) | **Get** /api/{serviceId}/service/jwks/get | Get JWK Set
*JoseObjectAPI* | [**JoseVerifyApi**](docs/JoseObjectAPI.md#joseverifyapi) | **Post** /api/{serviceId}/jose/verify | Verify JOSE
*PushedAuthorizationEndpointAPI* | [**PushedAuthReqApi**](docs/PushedAuthorizationEndpointAPI.md#pushedauthreqapi) | **Post** /api/{serviceId}/pushed_auth_req | Process Pushed Authorization Request
*RevocationEndpointAPI* | [**AuthRevocationApi**](docs/RevocationEndpointAPI.md#authrevocationapi) | **Post** /api/{serviceId}/auth/revocation | Process Revocation Request
*ServiceManagementAPI* | [**ServiceConfigurationApi**](docs/ServiceManagementAPI.md#serviceconfigurationapi) | **Get** /api/{serviceId}/service/configuration | Get Service Configuration
*ServiceManagementAPI* | [**ServiceCreateApi**](docs/ServiceManagementAPI.md#servicecreateapi) | **Post** /api/service/create | Create Service
*ServiceManagementAPI* | [**ServiceDeleteApi**](docs/ServiceManagementAPI.md#servicedeleteapi) | **Delete** /api/{serviceId}/service/delete | Delete Service ⚡
*ServiceManagementAPI* | [**ServiceGetApi**](docs/ServiceManagementAPI.md#servicegetapi) | **Get** /api/{serviceId}/service/get | Get Service
*ServiceManagementAPI* | [**ServiceGetListApi**](docs/ServiceManagementAPI.md#servicegetlistapi) | **Get** /api/service/get/list | List Services
*ServiceManagementAPI* | [**ServiceUpdateApi**](docs/ServiceManagementAPI.md#serviceupdateapi) | **Post** /api/{serviceId}/service/update | Update Service
*TokenEndpointAPI* | [**AuthTokenApi**](docs/TokenEndpointAPI.md#authtokenapi) | **Post** /api/{serviceId}/auth/token | Process Token Request
*TokenEndpointAPI* | [**AuthTokenFailApi**](docs/TokenEndpointAPI.md#authtokenfailapi) | **Post** /api/{serviceId}/auth/token/fail | Fail Token Request
*TokenEndpointAPI* | [**AuthTokenIssueApi**](docs/TokenEndpointAPI.md#authtokenissueapi) | **Post** /api/{serviceId}/auth/token/issue | Issue Token Response
*TokenEndpointAPI* | [**IdtokenReissueApi**](docs/TokenEndpointAPI.md#idtokenreissueapi) | **Post** /api/{serviceId}/idtoken/reissue | Reissue ID Token
*TokenOperationsAPI* | [**AuthTokenCreateApi**](docs/TokenOperationsAPI.md#authtokencreateapi) | **Post** /api/{serviceId}/auth/token/create | Create Access Token
*TokenOperationsAPI* | [**AuthTokenDeleteApi**](docs/TokenOperationsAPI.md#authtokendeleteapi) | **Delete** /api/{serviceId}/auth/token/delete/{accessTokenIdentifier} | Delete Access Token
*TokenOperationsAPI* | [**AuthTokenGetListApi**](docs/TokenOperationsAPI.md#authtokengetlistapi) | **Get** /api/{serviceId}/auth/token/get/list | List Issued Tokens
*TokenOperationsAPI* | [**AuthTokenRevokeApi**](docs/TokenOperationsAPI.md#authtokenrevokeapi) | **Post** /api/{serviceId}/auth/token/revoke | Revoke Access Token
*TokenOperationsAPI* | [**AuthTokenUpdateApi**](docs/TokenOperationsAPI.md#authtokenupdateapi) | **Post** /api/{serviceId}/auth/token/update | Update Access Token
*UserInfoEndpointAPI* | [**AuthUserinfoApi**](docs/UserInfoEndpointAPI.md#authuserinfoapi) | **Post** /api/{serviceId}/auth/userinfo | Process UserInfo Request
*UserInfoEndpointAPI* | [**AuthUserinfoIssueApi**](docs/UserInfoEndpointAPI.md#authuserinfoissueapi) | **Post** /api/{serviceId}/auth/userinfo/issue | Issue UserInfo Response
*UtilityEndpointsAPI* | [**InfoApi**](docs/UtilityEndpointsAPI.md#infoapi) | **Get** /api/info | Get Server Metadata
*UtilityEndpointsAPI* | [**MiscEchoApi**](docs/UtilityEndpointsAPI.md#miscechoapi) | **Get** /api/misc/echo | Echo
*VerifiableCredentialIssuerAPI* | [**VciBatchIssueApi**](docs/VerifiableCredentialIssuerAPI.md#vcibatchissueapi) | **Post** /api/{serviceId}/vci/batch/issue | /api/{serviceId}/vci/batch/issue API
*VerifiableCredentialIssuerAPI* | [**VciBatchParseApi**](docs/VerifiableCredentialIssuerAPI.md#vcibatchparseapi) | **Post** /api/{serviceId}/vci/batch/parse | /api/{serviceId}/vci/batch/parse API
*VerifiableCredentialIssuerAPI* | [**VciDeferredIssueApi**](docs/VerifiableCredentialIssuerAPI.md#vcideferredissueapi) | **Post** /api/{serviceId}/vci/deferred/issue | /api/{serviceId}/vci/deferred/issue API
*VerifiableCredentialIssuerAPI* | [**VciDeferredParseApi**](docs/VerifiableCredentialIssuerAPI.md#vcideferredparseapi) | **Post** /api/{serviceId}/vci/deferred/parse | /api/{serviceId}/vci/deferred/parse API
*VerifiableCredentialIssuerAPI* | [**VciJwksApi**](docs/VerifiableCredentialIssuerAPI.md#vcijwksapi) | **Post** /api/{serviceId}/vci/jwks | /api/{serviceId}/vci/jwks API
*VerifiableCredentialIssuerAPI* | [**VciJwtissuerApi**](docs/VerifiableCredentialIssuerAPI.md#vcijwtissuerapi) | **Post** /api/{serviceId}/vci/jwtissuer | /api/{serviceId}/vci/jwtissuer API
*VerifiableCredentialIssuerAPI* | [**VciMetadataApi**](docs/VerifiableCredentialIssuerAPI.md#vcimetadataapi) | **Post** /api/{serviceId}/vci/metadata | /api/{serviceId}/vci/metadata API
*VerifiableCredentialIssuerAPI* | [**VciOfferCreateApi**](docs/VerifiableCredentialIssuerAPI.md#vcioffercreateapi) | **Post** /api/{serviceId}/vci/offer/create | /api/{serviceId}/vci/offer/create API
*VerifiableCredentialIssuerAPI* | [**VciOfferInfoApi**](docs/VerifiableCredentialIssuerAPI.md#vciofferinfoapi) | **Post** /api/{serviceId}/vci/offer/info | /api/{serviceId}/vci/offer/info API
*VerifiableCredentialIssuerAPI* | [**VciSingleIssueApi**](docs/VerifiableCredentialIssuerAPI.md#vcisingleissueapi) | **Post** /api/{serviceId}/vci/single/issue | /api/{serviceId}/vci/single/issue API
*VerifiableCredentialIssuerAPI* | [**VciSingleParseApi**](docs/VerifiableCredentialIssuerAPI.md#vcisingleparseapi) | **Post** /api/{serviceId}/vci/single/parse | /api/{serviceId}/vci/single/parse API


## Documentation For Models

 - [AccessToken](docs/AccessToken.md)
 - [ApplicationType](docs/ApplicationType.md)
 - [AttachmentType](docs/AttachmentType.md)
 - [AuthorizationDetails](docs/AuthorizationDetails.md)
 - [AuthorizationDetailsElement](docs/AuthorizationDetailsElement.md)
 - [AuthorizationFailRequest](docs/AuthorizationFailRequest.md)
 - [AuthorizationFailResponse](docs/AuthorizationFailResponse.md)
 - [AuthorizationIssueRequest](docs/AuthorizationIssueRequest.md)
 - [AuthorizationIssueResponse](docs/AuthorizationIssueResponse.md)
 - [AuthorizationRequest](docs/AuthorizationRequest.md)
 - [AuthorizationResponse](docs/AuthorizationResponse.md)
 - [AuthorizationTicketInfoRequest](docs/AuthorizationTicketInfoRequest.md)
 - [AuthorizationTicketInfoResponse](docs/AuthorizationTicketInfoResponse.md)
 - [AuthorizationTicketUpdateRequest](docs/AuthorizationTicketUpdateRequest.md)
 - [AuthorizationTicketUpdateResponse](docs/AuthorizationTicketUpdateResponse.md)
 - [AuthzDetails](docs/AuthzDetails.md)
 - [BackchannelAuthenticationCompleteRequest](docs/BackchannelAuthenticationCompleteRequest.md)
 - [BackchannelAuthenticationCompleteResponse](docs/BackchannelAuthenticationCompleteResponse.md)
 - [BackchannelAuthenticationFailRequest](docs/BackchannelAuthenticationFailRequest.md)
 - [BackchannelAuthenticationFailResponse](docs/BackchannelAuthenticationFailResponse.md)
 - [BackchannelAuthenticationIssueRequest](docs/BackchannelAuthenticationIssueRequest.md)
 - [BackchannelAuthenticationIssueResponse](docs/BackchannelAuthenticationIssueResponse.md)
 - [BackchannelAuthenticationRequest](docs/BackchannelAuthenticationRequest.md)
 - [BackchannelAuthenticationResponse](docs/BackchannelAuthenticationResponse.md)
 - [ClaimType](docs/ClaimType.md)
 - [Client](docs/Client.md)
 - [ClientAuthMethod](docs/ClientAuthMethod.md)
 - [ClientAuthenticationMethod](docs/ClientAuthenticationMethod.md)
 - [ClientAuthorizationDeleteResponse](docs/ClientAuthorizationDeleteResponse.md)
 - [ClientAuthorizationGetListResponse](docs/ClientAuthorizationGetListResponse.md)
 - [ClientAuthorizationUpdateRequest](docs/ClientAuthorizationUpdateRequest.md)
 - [ClientAuthorizationUpdateResponse](docs/ClientAuthorizationUpdateResponse.md)
 - [ClientExtension](docs/ClientExtension.md)
 - [ClientExtensionRequestableScopesGetResponse](docs/ClientExtensionRequestableScopesGetResponse.md)
 - [ClientExtensionRequestableScopesUpdateRequest](docs/ClientExtensionRequestableScopesUpdateRequest.md)
 - [ClientExtensionRequestableScopesUpdateResponse](docs/ClientExtensionRequestableScopesUpdateResponse.md)
 - [ClientFlagUpdateRequest](docs/ClientFlagUpdateRequest.md)
 - [ClientFlagUpdateResponse](docs/ClientFlagUpdateResponse.md)
 - [ClientGetListResponse](docs/ClientGetListResponse.md)
 - [ClientGrantedScopesDeleteResponse](docs/ClientGrantedScopesDeleteResponse.md)
 - [ClientLimited](docs/ClientLimited.md)
 - [ClientLimitedAuthorization](docs/ClientLimitedAuthorization.md)
 - [ClientRegistrationDeleteRequest](docs/ClientRegistrationDeleteRequest.md)
 - [ClientRegistrationDeleteResponse](docs/ClientRegistrationDeleteResponse.md)
 - [ClientRegistrationRequest](docs/ClientRegistrationRequest.md)
 - [ClientRegistrationResponse](docs/ClientRegistrationResponse.md)
 - [ClientRegistrationType](docs/ClientRegistrationType.md)
 - [ClientRegistrationUpdateRequest](docs/ClientRegistrationUpdateRequest.md)
 - [ClientRegistrationUpdateResponse](docs/ClientRegistrationUpdateResponse.md)
 - [ClientSecretRefreshResponse](docs/ClientSecretRefreshResponse.md)
 - [ClientSecretUpdateRequest](docs/ClientSecretUpdateRequest.md)
 - [ClientSecretUpdateResponse](docs/ClientSecretUpdateResponse.md)
 - [ClientType](docs/ClientType.md)
 - [CredentialIssuanceOrder](docs/CredentialIssuanceOrder.md)
 - [CredentialIssuerMetadata](docs/CredentialIssuerMetadata.md)
 - [CredentialOfferInfo](docs/CredentialOfferInfo.md)
 - [CredentialRequestInfo](docs/CredentialRequestInfo.md)
 - [DeliveryMode](docs/DeliveryMode.md)
 - [DeviceAuthorizationRequest](docs/DeviceAuthorizationRequest.md)
 - [DeviceAuthorizationResponse](docs/DeviceAuthorizationResponse.md)
 - [DeviceCompleteRequest](docs/DeviceCompleteRequest.md)
 - [DeviceCompleteResponse](docs/DeviceCompleteResponse.md)
 - [DeviceVerificationRequest](docs/DeviceVerificationRequest.md)
 - [DeviceVerificationResponse](docs/DeviceVerificationResponse.md)
 - [Display](docs/Display.md)
 - [DynamicScope](docs/DynamicScope.md)
 - [FapiMode](docs/FapiMode.md)
 - [FederationConfigurationResponse](docs/FederationConfigurationResponse.md)
 - [FederationRegistrationRequest](docs/FederationRegistrationRequest.md)
 - [FederationRegistrationResponse](docs/FederationRegistrationResponse.md)
 - [GMRequest](docs/GMRequest.md)
 - [GMResponse](docs/GMResponse.md)
 - [Grant](docs/Grant.md)
 - [GrantManagementAction](docs/GrantManagementAction.md)
 - [GrantScope](docs/GrantScope.md)
 - [GrantType](docs/GrantType.md)
 - [Hsk](docs/Hsk.md)
 - [HskCreateRequest](docs/HskCreateRequest.md)
 - [HskCreateResponse](docs/HskCreateResponse.md)
 - [HskDeleteResponse](docs/HskDeleteResponse.md)
 - [HskGetListResponse](docs/HskGetListResponse.md)
 - [HskGetResponse](docs/HskGetResponse.md)
 - [IdtokenReissueRequest](docs/IdtokenReissueRequest.md)
 - [IdtokenReissueResponse](docs/IdtokenReissueResponse.md)
 - [InfoResponse](docs/InfoResponse.md)
 - [IntrospectionRequest](docs/IntrospectionRequest.md)
 - [IntrospectionResponse](docs/IntrospectionResponse.md)
 - [JoseVerifyRequest](docs/JoseVerifyRequest.md)
 - [JoseVerifyResponse](docs/JoseVerifyResponse.md)
 - [JweAlg](docs/JweAlg.md)
 - [JweEnc](docs/JweEnc.md)
 - [JwsAlg](docs/JwsAlg.md)
 - [NamedUri](docs/NamedUri.md)
 - [Pair](docs/Pair.md)
 - [Prompt](docs/Prompt.md)
 - [Property](docs/Property.md)
 - [PushedAuthorizationRequest](docs/PushedAuthorizationRequest.md)
 - [PushedAuthorizationResponse](docs/PushedAuthorizationResponse.md)
 - [ResponseType](docs/ResponseType.md)
 - [Result](docs/Result.md)
 - [RevocationRequest](docs/RevocationRequest.md)
 - [RevocationResponse](docs/RevocationResponse.md)
 - [Scope](docs/Scope.md)
 - [Service](docs/Service.md)
 - [ServiceGetListResponse](docs/ServiceGetListResponse.md)
 - [ServiceJwksGetResponse](docs/ServiceJwksGetResponse.md)
 - [ServiceProfile](docs/ServiceProfile.md)
 - [Sns](docs/Sns.md)
 - [SnsCredentials](docs/SnsCredentials.md)
 - [StandardIntrospectionRequest](docs/StandardIntrospectionRequest.md)
 - [StandardIntrospectionResponse](docs/StandardIntrospectionResponse.md)
 - [SubjectType](docs/SubjectType.md)
 - [TaggedValue](docs/TaggedValue.md)
 - [TokenCreateRequest](docs/TokenCreateRequest.md)
 - [TokenCreateResponse](docs/TokenCreateResponse.md)
 - [TokenFailRequest](docs/TokenFailRequest.md)
 - [TokenFailResponse](docs/TokenFailResponse.md)
 - [TokenGetListResponse](docs/TokenGetListResponse.md)
 - [TokenInfo](docs/TokenInfo.md)
 - [TokenIssueRequest](docs/TokenIssueRequest.md)
 - [TokenIssueResponse](docs/TokenIssueResponse.md)
 - [TokenRequest](docs/TokenRequest.md)
 - [TokenResponse](docs/TokenResponse.md)
 - [TokenRevokeRequest](docs/TokenRevokeRequest.md)
 - [TokenRevokeResponse](docs/TokenRevokeResponse.md)
 - [TokenType](docs/TokenType.md)
 - [TokenUpdateRequest](docs/TokenUpdateRequest.md)
 - [TokenUpdateResponse](docs/TokenUpdateResponse.md)
 - [TrustAnchor](docs/TrustAnchor.md)
 - [UserCodeCharset](docs/UserCodeCharset.md)
 - [UserinfoIssueRequest](docs/UserinfoIssueRequest.md)
 - [UserinfoIssueResponse](docs/UserinfoIssueResponse.md)
 - [UserinfoRequest](docs/UserinfoRequest.md)
 - [UserinfoResponse](docs/UserinfoResponse.md)
 - [VciBatchIssueRequest](docs/VciBatchIssueRequest.md)
 - [VciBatchIssueResponse](docs/VciBatchIssueResponse.md)
 - [VciBatchParseRequest](docs/VciBatchParseRequest.md)
 - [VciBatchParseResponse](docs/VciBatchParseResponse.md)
 - [VciDeferredIssueRequest](docs/VciDeferredIssueRequest.md)
 - [VciDeferredIssueResponse](docs/VciDeferredIssueResponse.md)
 - [VciDeferredParseRequest](docs/VciDeferredParseRequest.md)
 - [VciDeferredParseResponse](docs/VciDeferredParseResponse.md)
 - [VciJwksRequest](docs/VciJwksRequest.md)
 - [VciJwksResponse](docs/VciJwksResponse.md)
 - [VciJwtissuerRequest](docs/VciJwtissuerRequest.md)
 - [VciJwtissuerResponse](docs/VciJwtissuerResponse.md)
 - [VciMetadataRequest](docs/VciMetadataRequest.md)
 - [VciMetadataResponse](docs/VciMetadataResponse.md)
 - [VciOfferCreateRequest](docs/VciOfferCreateRequest.md)
 - [VciOfferCreateResponse](docs/VciOfferCreateResponse.md)
 - [VciOfferInfoRequest](docs/VciOfferInfoRequest.md)
 - [VciOfferInfoResponse](docs/VciOfferInfoResponse.md)
 - [VciSingleIssueRequest](docs/VciSingleIssueRequest.md)
 - [VciSingleIssueResponse](docs/VciSingleIssueResponse.md)
 - [VciSingleParseRequest](docs/VciSingleParseRequest.md)
 - [VciSingleParseResponse](docs/VciSingleParseResponse.md)
 - [VerifiedClaimsValidationSchema](docs/VerifiedClaimsValidationSchema.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### authlete


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://login.authlete.com/authorize
- **Scopes**: 
 - **authlete**: Inherit Authlete Account Permissions

Example

```go
auth := context.WithValue(context.Background(), authlete.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, authlete.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

### bearer

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), authlete.ContextAccessToken, "BEARER_TOKEN_STRING")
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



