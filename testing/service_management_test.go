package testing

import (
	"strconv"
	"testing"

	authlete "github.com/authlete/openapi-for-go"
)

func TestServiceLifecycle(t *testing.T) {
	t.SkipNow()
	authleteClient := createClient()
	testservice := authorizationCodeDTO()
	auth := createSOContext()
	srv, _, err := authleteClient.ServiceManagementApi.ServiceCreateApi(auth).Service(*testservice).Execute()

	if err != nil {
		t.Errorf("An error occured when create a service: %q", err.Error())
	}

	if srv.GetApiKey() == 0 {
		t.Error("Created service has no apikey")
	}

	key := strconv.FormatInt(srv.GetApiKey(), 10)
	defer ensureDeleteService(key, authleteClient)

	srv.SetDescription("Changed Description")
	if srv.GetDescription() == testservice.GetDescription() {
		t.Error("Create service method is updating in place the return")
	}

	authleteClient.ServiceManagementApi.ServiceUpdateApi(auth, key).Service(*srv).Execute()

	getSrv, _, errGet := authleteClient.ServiceManagementApi.ServiceGetApi(auth, key).Execute()
	if errGet != nil {
		t.Errorf("An error occured when create a service: %q", errGet.Error())
	}

	if getSrv.GetDescription() != "Changed Description" {
		t.Error("Update service method did not succeed")
	}

	founded := false
	var srvFromList authlete.Service
	srvList, _, errList := authleteClient.ServiceManagementApi.ServiceGetListApi(auth).Start(0).End(1000).Execute()
	if errList != nil {
		t.Errorf("An error occured when getting the service list: %q", errList.Error())
	}
	for _, srvFromList = range srvList.GetServices() {
		if srvFromList.GetApiKey() == srv.GetApiKey() {
			founded = true
			break
		}
	}
	if !founded {
		t.Error("Created service not found on the service list")
	}

	authleteClient.ServiceManagementApi.ServiceDeleteApi(auth, key).Execute()

	srvList, _, errList = authleteClient.ServiceManagementApi.ServiceGetListApi(auth).Start(0).End(1000).Execute()

	if errList != nil {
		t.Errorf("An error occured when getting the service list: %q", errList.Error())
	}
	founded = false
	for _, srvFromList = range srvList.GetServices() {
		if srvFromList.GetApiKey() == srv.GetApiKey() {
			founded = true
			break
		}
	}
	if founded {
		t.Error("Delete service not working, as it was found on the service list")
	}

}
