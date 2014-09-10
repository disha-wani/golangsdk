// +build acceptance

package v3

import (
	"testing"

	"github.com/rackspace/gophercloud"
	services3 "github.com/rackspace/gophercloud/openstack/identity/v3/services"
)

func TestListServices(t *testing.T) {
	// Create a service client.
	serviceClient := createAuthenticatedClient(t)
	if serviceClient == nil {
		return
	}

	// Use the client to list all available services.
	results, err := services3.List(serviceClient, services3.ListOpts{})
	if err != nil {
		t.Fatalf("Unable to list services: %v", err)
	}

	err = gophercloud.EachPage(results, func(page gophercloud.Collection) bool {
		t.Logf("--- Page ---")
		for _, service := range services3.AsServices(page) {
			t.Logf("Service: %32s %15s %10s %s", service.ID, service.Type, service.Name, *service.Description)
		}
		return true
	})
	if err != nil {
		t.Errorf("Unexpected error traversing pages: %v", err)
	}
}
