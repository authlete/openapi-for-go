package testing

import (
	"context"
	"os"
	"strconv"
	"testing"

	authlete "github.com/authlete/openapi-for-go"
)

func createClient() *authlete.APIClient {
	api_server := os.Getenv("AUTHLETE_API_SERVER")

	cfg := authlete.NewConfiguration()
	cfg.UserAgent = "authlete-ci"
	if len(api_server) > 0 {
		cfg.Servers[0].URL = api_server
	}

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

func createAPIContext(apiKey string, apiSecret string) context.Context {

	return context.WithValue(context.Background(), authlete.ContextBasicAuth, authlete.BasicAuth{
		UserName: apiKey,
		Password: apiSecret,
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

func createTestClient(t *testing.T) (*authlete.APIClient, *authlete.Service, context.Context, *authlete.Client) {
	authleteClient := createClient()
	testservice := authorizationCodeDTO()
	auth := createSOContext()
	srv, _, err := authleteClient.ServiceManagementApi.ServiceCreateApi(auth).Service(*testservice).Execute()

	if err != nil {
		t.Errorf("An error occured when create a testing service: %q", err.Error())
	}

	apiKey := strconv.FormatInt(srv.GetApiKey(), 10)
	apiSecret := srv.GetApiSecret()

	oauthClient := authlete.NewClient()
	oauthClient.SetClientIdAlias("oauthclient1")
	oauthClient.SetDescription("this is a oauthclient")
	oauthClient.SetApplicationType(authlete.APPLICATIONTYPE_WEB)
	oauthClient.SetClientName("oauthclient")
	oauthClient.SetDeveloper("gotest")
	oauthClient.SetResponseTypes([]authlete.ResponseType{authlete.RESPONSETYPE_CODE})
	oauthClient.SetGrantTypes([]authlete.GrantType{authlete.GRANTTYPE_AUTHORIZATION_CODE, authlete.GRANTTYPE_REFRESH_TOKEN})

	apiCtx := createAPIContext(apiKey, apiSecret)

	newCli, _, errCli := authleteClient.ClientManagementApi.ClientCreateApi(apiCtx).Client(*oauthClient).Execute()

	if errCli != nil {
		defer ensureDeleteService(apiKey, authleteClient)
		t.Errorf("An error occured when create a testing client: %q", errCli.Error())
	}

	if newCli.GetClientId() == 0 || newCli.GetClientSecret() == "" {
		defer ensureDeleteService(apiKey, authleteClient)
		t.Errorf("Testing client has no key or secret")
	}
	return authleteClient, srv, apiCtx, newCli
}
