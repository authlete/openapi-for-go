package testing

import (
	"strconv"
	"testing"

	authlete "github.com/authlete/openapi-for-go"
)

func TestClientLifecycle(t *testing.T) {
	t.SkipNow()
	authleteClient, srv, apiCtx, newCli := createTestClient(t)
	apiKey := strconv.FormatInt(srv.GetApiKey(), 10)
	defer ensureDeleteService(apiKey, authleteClient)

	client_id := strconv.FormatInt(newCli.GetClientId(), 10)
	client_secret := newCli.GetClientSecret()

	// update the client

	att1 := authlete.NewPair()
	att1.SetKey("key1")
	att1.SetValue("val1")
	newCli.SetAttributes([]authlete.Pair{*att1})
	newCli.SetDescription("Changed Description")

	updCli, _, errUp := authleteClient.ClientManagementApi.ClientUpdateApi(apiCtx, client_id).Client(*newCli).Execute()

	if errUp != nil {
		t.Errorf("An error occured when updating a testing client: %q", errUp.Error())
	}

	if len(updCli.GetAttributes()) == 0 {
		t.Errorf("Changes on attributes not reflected on returned client from update client api")
	}

	// get the client

	getCli, _, errGet := authleteClient.ClientManagementApi.ClientGetApi(apiCtx, client_id).Execute()

	if errGet != nil {
		t.Errorf("An error occured when retrieving the testing client: %q", errGet.Error())
	}

	if updCli.GetDescription() != getCli.GetDescription() {
		t.Errorf("The change to description is not reflected on retrieved client ")
	}

	// list the clients

	cliList, _, errList := authleteClient.ClientManagementApi.ClientGetListApi(apiCtx).Start(0).End(100).Execute()

	if errList != nil {
		t.Errorf("An error occured when retrieving the client list: %q", errList.Error())
	}

	found := false
	var cli authlete.Client
	for _, cli = range cliList.GetClients() {
		if cli.GetClientId() == newCli.GetClientId() {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Could not find the created client on client list")
	}
	if cli.GetClientSecret() != client_secret {
		t.Errorf("The client secret was changed on previous calls")
	}

	// delete the client

	_, errDelete := authleteClient.ClientManagementApi.ClientDeleteApi(apiCtx, client_id).Execute()

	if errDelete != nil {
		t.Errorf("An error occured when deleting the test client: %q", errDelete.Error())
	}

	// check if it was deleted

	cliList, _, errList = authleteClient.ClientManagementApi.ClientGetListApi(apiCtx).Start(0).End(100).Execute()

	if errList != nil {
		t.Errorf("An error occured when retrieving the client list: %q", errList.Error())
	}

	found = false
	for _, cli = range cliList.GetClients() {
		if cli.GetClientId() == newCli.GetClientId() {
			found = true
			break
		}
	}
	if found {
		t.Errorf("deleted client still on client list")
	}
}

func TestSecretManagement(t *testing.T) {
	t.SkipNow()
	authleteClient, srv, apiCtx, newCli := createTestClient(t)
	apiKey := strconv.FormatInt(srv.GetApiKey(), 10)
	defer ensureDeleteService(apiKey, authleteClient)

	client_id := strconv.FormatInt(newCli.GetClientId(), 10)

	client_secret_up := "jsadfkljhfadkjsadfj1"
	req := authlete.NewClientSecretUpdateRequest(client_secret_up)
	resp, _, err := authleteClient.ClientManagementApi.ClientSecretUpdateApi(apiCtx, client_id).ClientSecretUpdateRequest(*req).Execute()

	if err != nil {
		t.Errorf("An error occured when updating the secret: %q", err.Error())
	}

	if resp.GetNewClientSecret() != client_secret_up {
		t.Errorf("new secret not matching %q - expected %q", resp.GetNewClientSecret(), client_secret_up)
	}

	respRefresh, _, errRefresh := authleteClient.ClientManagementApi.ClientSecretRefreshApi(apiCtx, client_id).Execute()

	if errRefresh != nil {
		t.Errorf("An error occured when updating the secret: %q", errRefresh.Error())
	}

	if respRefresh.GetOldClientSecret() != client_secret_up {
		t.Errorf("old secret not matching %q - expected %q", respRefresh.GetOldClientSecret(), client_secret_up)
	}

	if respRefresh.GetOldClientSecret() == respRefresh.GetNewClientSecret() {
		t.Errorf("old secret and new secret are matching %q", respRefresh.GetOldClientSecret())
	}

}
