// +build acceptance metric archivepolicies

package v1

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/acdc-cloud/utils/acceptance/clients"
	"github.com/acdc-cloud/utils/gnocchi/metric/v1/archivepolicies"
)

func TestArchivePoliciesList(t *testing.T) {
	client, err := clients.NewGnocchiV1Client()
	if err != nil {
		t.Fatalf("Unable to create a Gnocchi client: %v", err)
	}

	allPages, err := archivepolicies.List(client).AllPages()
	if err != nil {
		t.Fatalf("Unable to list archive policies: %v", err)
	}

	allArchivePolicies, err := archivepolicies.ExtractArchivePolicies(allPages)
	if err != nil {
		t.Fatalf("Unable to extract archive policies: %v", err)
	}

	for _, archivePolicy := range allArchivePolicies {
		tools.PrintResource(t, archivePolicy)
	}
}
