// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecs_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSECSClusterCapacityProviderAssociations_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ECS::ClusterCapacityProviderAssociations", "awscc_ecs_cluster_capacity_provider_associations", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
