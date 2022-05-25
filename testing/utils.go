package testing

import (
	"context"
	"os"

	authlete "github.com/authlete/openapi-for-go"
)

func createClient() *authlete.APIClient {
	api_server := os.Getenv("AUTHLETE_API_SERVER")
	if len(api_server) == 0 {
		api_server = "https://api.authlete.com/api"
	}
	cfg := authlete.NewConfiguration()
	cfg.UserAgent = "authlete-ci"
	cfg.Servers[0].URL = api_server

	return authlete.NewAPIClient(cfg)
}

func createSOContext() context.Context {
	so_key := os.Getenv("AUTHLETE_SO_KEY")
	so_secret := os.Getenv("AUTHLETE_SO_SECRET")

	return context.WithValue(context.Background(), authlete.ContextBasicAuth, authlete.BasicAuth{
		UserName: so_key,
		Password: so_secret,
	})
}

func ensureDeleteService(apiKey string, cli *authlete.APIClient) {

	ctx := createSOContext()

	cli.ServiceManagementApi.ServiceDeleteApi(ctx, apiKey).Execute()

}

func authorizationCodeDTO() *authlete.Service {
	var testService *authlete.Service
	testService = authlete.NewService()
	testService.SetServiceName("Test Service for client testing")
	testService.SetIssuer("https://test.com")
	testService.SetSupportedGrantTypes([]authlete.GrantType{
		authlete.GRANTTYPE_AUTHORIZATION_CODE,
		authlete.GRANTTYPE_REFRESH_TOKEN})
	testService.SetSupportedResponseTypes(
		[]authlete.ResponseType{authlete.RESPONSETYPE_CODE})

	openid := authlete.NewScope()
	openid.SetName("openid")
	openid.SetDescription("A permission to request an OpenID Provider to issue an ID Token. See OpenID Connect Core 1.0, 3.1.2.1. for details.")
	openid.SetDefaultEntry(false)

	profile := authlete.NewScope()
	profile.SetName("profile")
	profile.SetDescription("A permission to request an OpenID Provider to include claims related to end-user's profile in an ID Token. See OpenID Connect Core 1.0, 5.4. for details.")
	profile.SetDefaultEntry(false)

	testService.SupportedScopes = []authlete.Scope{
		*openid,
		*profile,
	}
	return testService
}
