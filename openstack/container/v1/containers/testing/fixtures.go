package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/container/v1/capsules"
	"github.com/gophercloud/gophercloud/openstack/container/v1/containers"
	th "github.com/gophercloud/gophercloud/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/testhelper/client"
)

const ContainerGetBody = `
{
    "addresses": {
        "b1295212-64e1-471d-aa01-25ff46f9818d":
            [
                {
                    "preserve_on_delete": false,
                    "addr": "172.24.4.11",
                    "port": "8439060f-381a-4386-a518-33d5a4058636",
                    "version": 4,
                    "subnet_id": "4a2bcd64-93ad-4436-9f48-3a7f9b267e0a"
                }
            ]
    },
    "auto_heal": false,
    "auto_remove": false,
    "command": [
        "testcmd"
    ],
    "cpu": 1,
    "cpu_policy": "shared",
    "disk": 0,
    "entrypoint": [],
    "environment": {
        "USER1": "test"
    },
    "healthcheck": {},
    "hostname": "test-hostname",
    "image": "test",
    "image_driver": "docker",
    "interactive": true,
    "labels": {
        "foo": "bar"
    },
    "links": [
        {
            "href": "https://example.com/v1/containers/1739e28a-d391-4fd9-93a5-3ba3f29a4c9b",
            "rel": "self"
        },
        {
            "href": "https://example.com/containers/1739e28a-d391-4fd9-93a5-3ba3f29a4c9b",
            "rel": "bookmark"
        }
    ],
    "memory": "1024M",
    "name": "test-demo-omicron-13",
    "ports": [ 80 ],
    "project_id": "6b8ffef2a0ac42ee87887b9cc98bdf68",
    "registry_id": null,
    "restart_policy": {
        "MaximumRetryCount": "0",
        "Name": "always"
    },
    "security_groups": [
        "default"
    ],
    "status": "Running",
    "status_detail": "Just created",
    "status_reason": "No reason",
    "task_state": "Creating",
    "tty": true,
    "user_id": "d33b18c384574fd2a3299447aac285f0",
    "uuid": "1739e28a-d391-4fd9-93a5-3ba3f29a4c9b",
    "workdir": "/root",
    "host": "test-host"
}`

const ContainerListBody = `
{
    "containers": [
        {
            "addresses": {
                "b1295212-64e1-471d-aa01-25ff46f9818d":
                    [
                        {
                            "preserve_on_delete": false,
                            "addr": "172.24.4.11",
                            "port": "8439060f-381a-4386-a518-33d5a4058636",
                            "version": 4,
                            "subnet_id": "4a2bcd64-93ad-4436-9f48-3a7f9b267e0a"
                        }
                    ]
            },
            "auto_heal": false,
            "auto_remove": false,
            "command": [
                "testcmd"
            ],
            "cpu": 1,
            "cpu_policy": "shared",
            "disk": 0,
            "entrypoint": [],
            "environment": {
                "USER1": "test"
            },
            "healthcheck": {},
            "hostname": "test-hostname",
            "image": "test",
            "image_driver": "docker",
            "interactive": true,
            "labels": {
                "foo": "bar"
            },
            "links": [
                {
                    "href": "https://example.com/v1/containers/1739e28a-d391-4fd9-93a5-3ba3f29a4c9b",
                    "rel": "self"
                },
                {
                    "href": "https://example.com/containers/1739e28a-d391-4fd9-93a5-3ba3f29a4c9b",
                    "rel": "bookmark"
                }
            ],
            "memory": "1024M",
            "name": "test-demo-omicron-13",
            "ports": [ 80 ],
            "project_id": "6b8ffef2a0ac42ee87887b9cc98bdf68",
            "registry_id": null,
            "restart_policy": {
                "MaximumRetryCount": "0",
                "Name": "always"
            },
            "security_groups": [
                "default"
            ],
            "status": "Running",
            "status_detail": "Just created",
            "status_reason": "No reason",
            "task_state": "Creating",
            "tty": true,
            "user_id": "d33b18c384574fd2a3299447aac285f0",
            "uuid": "1739e28a-d391-4fd9-93a5-3ba3f29a4c9b",
            "workdir": "/root",
            "host": "test-host"
        }
    ]
}`

var ExpectedContainer = containers.Container{
	Container: capsules.Container{
		Name:      "test-demo-omicron-13",
		UUID:      "1739e28a-d391-4fd9-93a5-3ba3f29a4c9b",
		UserID:    "d33b18c384574fd2a3299447aac285f0",
		ProjectID: "6b8ffef2a0ac42ee87887b9cc98bdf68",
		CPU:       float64(1),
		Memory:    "1024M",
		Host:      "test-host",
		Status:    "Running",
		Image:     "test",
		Labels: map[string]string{
			"foo": "bar",
		},
		WorkDir: "/root",
		Disk:    0,
		Command: []string{
			"testcmd",
		},
		Ports: []int{
			80,
		},
		SecurityGroups: []string{
			"default",
		},
		TaskState: "Creating",
		HostName:  "test-hostname",
		Environment: map[string]string{
			"USER1": "test",
		},
		StatusReason: "No reason",
		StatusDetail: "Just created",
		ImageDriver:  "docker",
		Interactive:  true,
		AutoRemove:   false,
		AutoHeal:     false,
		RestartPolicy: map[string]string{
			"MaximumRetryCount": "0",
			"Name":              "always",
		},
		Addresses: map[string][]capsules.Address{
			"b1295212-64e1-471d-aa01-25ff46f9818d": {
				{
					PreserveOnDelete: false,
					Addr:             "172.24.4.11",
					Port:             "8439060f-381a-4386-a518-33d5a4058636",
					Version:          float64(4),
					SubnetID:         "4a2bcd64-93ad-4436-9f48-3a7f9b267e0a",
				},
			},
		},
		TTY:         true,
		CPUPolicy:   "shared",
		Healthcheck: map[string]string{},
		Links: []interface{}{
			map[string]string{
				"href": "https://example.com/v1/containers/1739e28a-d391-4fd9-93a5-3ba3f29a4c9b",
				"rel":  "self",
			},
			map[string]string{
				"href": "https://example.com/containers/1739e28a-d391-4fd9-93a5-3ba3f29a4c9b",
				"rel":  "bookmark",
			},
		},
	},
}

// HandleContainerGetSuccessfully test setup
func HandleContainerGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/containers/1739e28a-d391-4fd9-93a5-3ba3f29a4c9b", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ContainerGetBody)
	})
}

// HandleContainerCreateSuccessfully creates an HTTP handler at `/container` on the test handler mux
// that responds with a `Create` response.
func HandleContainerCreateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/containers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, ContainerGetBody)
	})
}

// HandleContainerListSuccessfully test setup
func HandleContainerListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/containers/", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ContainerListBody)
	})
}

func HandleContainerDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/containers/963a239d-3946-452b-be5a-055eab65a421", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})
}
