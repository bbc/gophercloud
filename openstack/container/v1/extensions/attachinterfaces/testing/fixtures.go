package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/container/v1/extensions/attachinterfaces"
	th "github.com/gophercloud/gophercloud/testhelper"
	"github.com/gophercloud/gophercloud/testhelper/client"
)

// ListInterfacesExpected represents an expected repsonse from a ListInterfaces request.
var ListInterfacesExpected = []attachinterfaces.Interface{
	{
		FixedIP: attachinterfaces.FixedIP{
			SubnetID:  "d7906db4-a566-4546-b1f4-5c7fa70f0bf3",
			IPAddress: "10.0.0.7",
			Version:   4,
		},
		PortID:  "0dde1598-b374-474e-986f-5b8dd1df1d4e",
		NetID:   "8a5fe506-7e9f-4091-899b-96336909d93c",
	},
}

// GetInterfaceExpected represents an expected repsonse from a GetInterface request.
var GetInterfaceExpected = attachinterfaces.Interface{
	FixedIP: attachinterfaces.FixedIP{
		SubnetID:  "d7906db4-a566-4546-b1f4-5c7fa70f0bf3",
		IPAddress: "10.0.0.7",
		Version:   4,
	},
	PortID:  "0dde1598-b374-474e-986f-5b8dd1df1d4e",
	NetID:   "8a5fe506-7e9f-4091-899b-96336909d93c",
}

// HandleInterfaceListSuccessfully sets up the test server to respond to a ListInterfaces request.
func HandleInterfaceListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/containers/b07e7a3b-d951-4efc-a4f9-ac9f001afb7f/network_list", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"networks": [
				{
					"fixed_ips": {
						"subnet_id": "d7906db4-a566-4546-b1f4-5c7fa70f0bf3",
						"ip_address": "10.0.0.7",
						"version": 4
					},
					"port_id": "0dde1598-b374-474e-986f-5b8dd1df1d4e",
					"net_id": "8a5fe506-7e9f-4091-899b-96336909d93c"
				}
			]
		}`)
	})
}

// HandleInterfaceCreateSuccessfully sets up the test server to respond to a CreateInterface request.
func HandleInterfaceCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/containers/b07e7a3b-d951-4efc-a4f9-ac9f001afb7f/network_attach", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		network, ok := r.URL.Query()["network"]
		if !ok {
			t.Errorf("Key 'network' not found")
		}

		expected := "8a5fe506-7e9f-4091-899b-96336909d93c"
		if network[0] != expected {
			t.Errorf("Value of 'network' = %v, expected %v", network[0], expected)
		}

		w.WriteHeader(http.StatusAccepted)
	})
}

// HandleInterfaceDeleteSuccessfully sets up the test server to respond to a DeleteInterface request.
func HandleInterfaceDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/containers/b07e7a3b-d951-4efc-a4f9-ac9f001afb7f/network_detach", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		network, ok := r.URL.Query()["network"]
		if !ok {
			t.Errorf("Key 'network' not found")
		}

		expected := "8a5fe506-7e9f-4091-899b-96336909d93c"
		if network[0] != expected {
			t.Errorf("Value of 'network' = %v, expected %v", network[0], expected)
		}

		w.WriteHeader(http.StatusAccepted)
	})
}
