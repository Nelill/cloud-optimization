package gcp

import (
	"context"

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

var (
	projectID      string
	computeService *compute.Service
)

func Configure(projID, zone string) error {
	projectID = projID

	ctx := context.Background()
	service, err := compute.NewService(ctx, option.WithCredentialsFile("https://www.googleapis.com/auth/cloud-platform.read-only"))
	if err != nil {
		return err
	}

	computeService = service
	return nil
}

func GetInstances(zone string) ([]*compute.Instance, error) {
	instanceList, err := computeService.Instances.List(projectID, zone).Do()
	if err != nil {
		return nil, err
	}

	return instanceList.Items, nil
}
