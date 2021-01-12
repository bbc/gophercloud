package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/container/v1/containers"
	"github.com/gophercloud/gophercloud/pagination"
	th "github.com/gophercloud/gophercloud/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestGetContainer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleContainerGetSuccessfully(t)

	actualContainer, err := containers.Get(fakeclient.ServiceClient(), ExpectedContainer.UUID).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, &ExpectedContainer, actualContainer)
}

func TestCreateContainer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleContainerCreateSuccessfully(t)

	createOpts := containers.CreateOpts{
		Name:             ExpectedContainer.Name,
		Image:            ExpectedContainer.Image,
		Command:          ExpectedContainer.Command,
		CPU:              ExpectedContainer.CPU,
		Memory:           ExpectedContainer.Memory,
		Workdir:          ExpectedContainer.WorkDir,
		Labels:           ExpectedContainer.Labels,
		Environment:      ExpectedContainer.Environment,
		RestartPolicy:    ExpectedContainer.RestartPolicy,
		Interactive:      &ExpectedContainer.Interactive,
		TTY:              &ExpectedContainer.TTY,
		ImageDriver:      ExpectedContainer.ImageDriver,
		SecurityGroups:   ExpectedContainer.SecurityGroups,
		Nets:             []map[string]string{},
		Runtime:          ExpectedContainer.Runtime,
		Hostname:         ExpectedContainer.HostName,
		AutoRemove:       &ExpectedContainer.AutoRemove,
		AutoHeal:         &ExpectedContainer.AutoHeal,
		AvailabilityZone: "AvailabilityZone1",
		Hints: map[string]string{
			"foo": "bar",
		},
		Mounts: []map[string]string{
			{
				"source":      "myvol",
				"destination": "/data",
			},
		},
		Privileged:   &ExpectedContainer.Privileged,
		Healthcheck:  ExpectedContainer.Healthcheck,
		ExposedPorts: nil,
		Host:         ExpectedContainer.Host,
		Entrypoint:   ExpectedContainer.Entrypoint,
	}

	actualContainer, err := containers.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, &ExpectedContainer, actualContainer)
}

func TestListContainer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleContainerListSuccessfully(t)

	expected := []containers.Container{ExpectedContainer}

	count := 0
	results := containers.List(fakeclient.ServiceClient(), nil)
	err := results.EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := containers.ExtractContainers(page)
		if err != nil {
			t.Errorf("Failed to extract containers: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, expected, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleContainerDeleteSuccessfully(t)

	res := containers.Delete(fakeclient.ServiceClient(), "963a239d-3946-452b-be5a-055eab65a421")
	th.AssertNoErr(t, res.Err)
}
