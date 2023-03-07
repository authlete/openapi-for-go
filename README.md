# Go API client for authlete

Authlete API Document.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.3.0
- Package version: 2.3
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import authlete "github.com/authlete/openapi-for-go"
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
ctx := context.WithValue(context.Background(), authlete.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), authlete.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
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

All URIs are relative to *https://api.authlete.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthorizationEndpointApi* | [**AuthAuthorizationApi**](docs/AuthorizationEndpointApi.md#authauthorizationapi) | **Post** /api/auth/authorization | /api/auth/authorization API
*AuthorizationEndpointApi* | [**AuthAuthorizationFailApi**](docs/AuthorizationEndpointApi.md#authauthorizationfailapi) | **Post** /api/auth/authorization/fail | /api/auth/authorization/fail API
*AuthorizationEndpointApi* | [**AuthAuthorizationIssueApi**](docs/AuthorizationEndpointApi.md#authauthorizationissueapi) | **Post** /api/auth/authorization/issue | /api/auth/authorization/issue API
*CIBAApi* | [**BackchannelAuthenticationApi**](docs/CIBAApi.md#backchannelauthenticationapi) | **Post** /api/backchannel/authentication | /api/backchannel/authentication API
*CIBAApi* | [**BackchannelAuthenticationCompleteApi**](docs/CIBAApi.md#backchannelauthenticationcompleteapi) | **Post** /api/backchannel/authentication/complete | /api/backchannel/authentication/complete API
*CIBAApi* | [**BackchannelAuthenticationFailApi**](docs/CIBAApi.md#backchannelauthenticationfailapi) | **Post** /api/backchannel/authentication/fail | /api/backchannel/authentication/fail API
*CIBAApi* | [**BackchannelAuthenticationIssueApi**](docs/CIBAApi.md#backchannelauthenticationissueapi) | **Post** /api/backchannel/authentication/issue | /api/backchannel/authentication/issue API
*ClientManagementApi* | [**ClientAuthorizationDeleteApi**](docs/ClientManagementApi.md#clientauthorizationdeleteapi) | **Delete** /api/client/authorization/delete/{clientId}/{subject} | /api/client/authorization/delete/{clientId}/{subject} API
*ClientManagementApi* | [**ClientAuthorizationGetListApi**](docs/ClientManagementApi.md#clientauthorizationgetlistapi) | **Get** /api/client/authorization/get/list/{subject} | /api/client/authorization/get/list/{subject} API
*ClientManagementApi* | [**ClientAuthorizationUpdateApi**](docs/ClientManagementApi.md#clientauthorizationupdateapi) | **Post** /api/client/authorization/update/{clientId} | /api/client/authorization/update/{clientId} API
*ClientManagementApi* | [**ClientCreateApi**](docs/ClientManagementApi.md#clientcreateapi) | **Post** /api/client/create | /api/client/create API
*ClientManagementApi* | [**ClientDeleteApi**](docs/ClientManagementApi.md#clientdeleteapi) | **Delete** /api/client/delete/{clientId} | /api/client/delete/{clientId} API
*ClientManagementApi* | [**ClientFlagUpdateApi**](docs/ClientManagementApi.md#clientflagupdateapi) | **Post** /api/client/lock_flag/update/{clientIdentifier} | /api/client/lock_flag/update/{clientIdentifier} API
*ClientManagementApi* | [**ClientGetApi**](docs/ClientManagementApi.md#clientgetapi) | **Get** /api/client/get/{clientId} | /api/client/get/{clientId} API
*ClientManagementApi* | [**ClientGetListApi**](docs/ClientManagementApi.md#clientgetlistapi) | **Get** /api/client/get/list | /api/client/get/list API
*ClientManagementApi* | [**ClientGrantedScopesDeleteApi**](docs/ClientManagementApi.md#clientgrantedscopesdeleteapi) | **Delete** /api/client/granted_scopes/delete/{clientId}/{subject} | /api/client/granted_scopes/delete/{clientId}/{subject} API
*ClientManagementApi* | [**ClientGrantedScopesGetApi**](docs/ClientManagementApi.md#clientgrantedscopesgetapi) | **Get** /api/client/granted_scopes/get/{clientId}/{subject} | /api/client/granted_scopes/get/{clientId}/{subject} API
*ClientManagementApi* | [**ClientSecretRefreshApi**](docs/ClientManagementApi.md#clientsecretrefreshapi) | **Get** /api/client/secret/refresh/{clientIdentifier} | /api/client/secret/refresh API
*ClientManagementApi* | [**ClientSecretUpdateApi**](docs/ClientManagementApi.md#clientsecretupdateapi) | **Post** /api/client/secret/update/{clientIdentifier} | /api/client/secret/update API
*ClientManagementApi* | [**ClientUpdateApi**](docs/ClientManagementApi.md#clientupdateapi) | **Post** /api/client/update/{clientId} | /api/client/update/{clientId} API
*ConfigurationEndpointApi* | [**ServiceConfigurationApi**](docs/ConfigurationEndpointApi.md#serviceconfigurationapi) | **Get** /api/service/configuration | /api/service/configuration API
*DeviceFlowApi* | [**DeviceAuthorizationApi**](docs/DeviceFlowApi.md#deviceauthorizationapi) | **Post** /api/device/authorization | /api/device/authorization API
*DeviceFlowApi* | [**DeviceCompleteApi**](docs/DeviceFlowApi.md#devicecompleteapi) | **Post** /api/device/complete | /api/device/complete API
*DeviceFlowApi* | [**DeviceVerificationApi**](docs/DeviceFlowApi.md#deviceverificationapi) | **Post** /api/device/verification | /api/device/verification API
*DynamicClientRegistrationApi* | [**ClientRegistrationApi**](docs/DynamicClientRegistrationApi.md#clientregistrationapi) | **Post** /api/client/registration | /api/client/registration API
*DynamicClientRegistrationApi* | [**ClientRegistrationDeleteApi**](docs/DynamicClientRegistrationApi.md#clientregistrationdeleteapi) | **Post** /api/client/registration/delete | /api/client/registration/delete API
*DynamicClientRegistrationApi* | [**ClientRegistrationGetApi**](docs/DynamicClientRegistrationApi.md#clientregistrationgetapi) | **Post** /api/client/registration/get | /api/client/registration/get API
*DynamicClientRegistrationApi* | [**ClientRegistrationUpdateApi**](docs/DynamicClientRegistrationApi.md#clientregistrationupdateapi) | **Post** /api/client/registration/update | /api/client/registration/update API
*FederationEndpointApi* | [**FederationConfigurationApi**](docs/FederationEndpointApi.md#federationconfigurationapi) | **Post** /api/federation/configuration | /api/federation/configuration API
*GrantManagementEndpointApi* | [**GrantMApi**](docs/GrantManagementEndpointApi.md#grantmapi) | **Post** /api/gm | /api/gm API
*IntrospectionEndpointApi* | [**AuthIntrospectionApi**](docs/IntrospectionEndpointApi.md#authintrospectionapi) | **Post** /api/auth/introspection | /api/auth/introspection API
*IntrospectionEndpointApi* | [**AuthIntrospectionStandardApi**](docs/IntrospectionEndpointApi.md#authintrospectionstandardapi) | **Post** /api/auth/introspection/standard | /api/auth/intraspection/standard API
*JWKSetEndpointApi* | [**ServiceJwksGetApi**](docs/JWKSetEndpointApi.md#servicejwksgetapi) | **Get** /api/service/jwks/get | /api/service/jwks/get API
*JoseObjectApi* | [**JoseVerifyApi**](docs/JoseObjectApi.md#joseverifyapi) | **Post** /api/jose/verify | /api/jose/verify API
*PushedAuthorizationEndpointApi* | [**PushedAuthReqApi**](docs/PushedAuthorizationEndpointApi.md#pushedauthreqapi) | **Post** /api/pushed_auth_req | /api/pushed_auth_req API
*RevocationEndpointApi* | [**AuthRevocationApi**](docs/RevocationEndpointApi.md#authrevocationapi) | **Post** /api/auth/revocation | /api/auth/revocation API
*ServerMetadataApi* | [**InfoApi**](docs/ServerMetadataApi.md#infoapi) | **Get** /api/info | /api/info API
*ServiceManagementApi* | [**ServiceCreateApi**](docs/ServiceManagementApi.md#servicecreateapi) | **Post** /api/service/create | /api/service/create API
*ServiceManagementApi* | [**ServiceDeleteApi**](docs/ServiceManagementApi.md#servicedeleteapi) | **Delete** /api/service/delete/{serviceApiKey} | /api/service/delete/{serviceApiKey} API
*ServiceManagementApi* | [**ServiceGetApi**](docs/ServiceManagementApi.md#servicegetapi) | **Get** /api/service/get/{serviceApiKey} | /api/service/get/{serviceApiKey} API
*ServiceManagementApi* | [**ServiceGetListApi**](docs/ServiceManagementApi.md#servicegetlistapi) | **Get** /api/service/get/list | /api/service/get/list API
*ServiceManagementApi* | [**ServiceUpdateApi**](docs/ServiceManagementApi.md#serviceupdateapi) | **Post** /api/service/update/{serviceApiKey} | /api/service/update/{serviceApiKey} API
*TokenEndpointApi* | [**AuthTokenApi**](docs/TokenEndpointApi.md#authtokenapi) | **Post** /api/auth/token | /api/auth/token API
*TokenEndpointApi* | [**AuthTokenFailApi**](docs/TokenEndpointApi.md#authtokenfailapi) | **Post** /api/auth/token/fail | /api/auth/token/fail API
*TokenEndpointApi* | [**AuthTokenIssueApi**](docs/TokenEndpointApi.md#authtokenissueapi) | **Post** /api/auth/token/issue | /api/auth/token/issue API
*TokenOperationsApi* | [**AuthTokenCreateApi**](docs/TokenOperationsApi.md#authtokencreateapi) | **Post** /api/auth/token/create | /api/auth/token/create API
*TokenOperationsApi* | [**AuthTokenDeleteApi**](docs/TokenOperationsApi.md#authtokendeleteapi) | **Delete** /api/auth/token/delete/{accessTokenIdentifier} | /api/auth/token/delete API
*TokenOperationsApi* | [**AuthTokenGetListApi**](docs/TokenOperationsApi.md#authtokengetlistapi) | **Get** /api/auth/token/get/list | /api/auth/token/get/list API
*TokenOperationsApi* | [**AuthTokenRevokeApi**](docs/TokenOperationsApi.md#authtokenrevokeapi) | **Post** /api/auth/token/revoke | /api/auth/token/revoke API
*TokenOperationsApi* | [**AuthTokenUpdateApi**](docs/TokenOperationsApi.md#authtokenupdateapi) | **Post** /api/auth/token/update | /api/auth/token/update API
*UserInfoEndpointApi* | [**AuthUserinfoApi**](docs/UserInfoEndpointApi.md#authuserinfoapi) | **Post** /api/auth/userinfo | /api/auth/userinfo API
*UserInfoEndpointApi* | [**AuthUserinfoIssueApi**](docs/UserInfoEndpointApi.md#authuserinfoissueapi) | **Post** /api/auth/userinfo/issue | /api/auth/userinfo/issue API


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
 - [ClientAuthenticationMethod](docs/ClientAuthenticationMethod.md)
 - [ClientAuthorizationDeleteResponse](docs/ClientAuthorizationDeleteResponse.md)
 - [ClientAuthorizationGetListResponse](docs/ClientAuthorizationGetListResponse.md)
 - [ClientAuthorizationUpdateRequest](docs/ClientAuthorizationUpdateRequest.md)
 - [ClientAuthorizationUpdateResponse](docs/ClientAuthorizationUpdateResponse.md)
 - [ClientExtension](docs/ClientExtension.md)
 - [ClientFlagUpdateRequest](docs/ClientFlagUpdateRequest.md)
 - [ClientFlagUpdateResponse](docs/ClientFlagUpdateResponse.md)
 - [ClientGetListResponse](docs/ClientGetListResponse.md)
 - [ClientGrantedScopesDeleteResponse](docs/ClientGrantedScopesDeleteResponse.md)
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
 - [DeliveryMode](docs/DeliveryMode.md)
 - [DeviceAuthorizationRequest](docs/DeviceAuthorizationRequest.md)
 - [DeviceAuthorizationResponse](docs/DeviceAuthorizationResponse.md)
 - [DeviceCompleteRequest](docs/DeviceCompleteRequest.md)
 - [DeviceCompleteResponse](docs/DeviceCompleteResponse.md)
 - [DeviceVerificationRequest](docs/DeviceVerificationRequest.md)
 - [DeviceVerificationResponse](docs/DeviceVerificationResponse.md)
 - [Display](docs/Display.md)
 - [DynamicScope](docs/DynamicScope.md)
 - [FederationConfigurationResponse](docs/FederationConfigurationResponse.md)
 - [GMRequest](docs/GMRequest.md)
 - [GMResponse](docs/GMResponse.md)
 - [Grant](docs/Grant.md)
 - [GrantManagementAction](docs/GrantManagementAction.md)
 - [GrantScope](docs/GrantScope.md)
 - [GrantType](docs/GrantType.md)
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


## Documentation For Authorization



### ServiceCredentials

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


### ServiceOwnerCredentials

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
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



