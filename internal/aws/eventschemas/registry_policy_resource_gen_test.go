// Code generated by generators/resource/main.go; DO NOT EDIT.

package eventschemas_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSEventSchemasRegistryPolicy_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::EventSchemas::RegistryPolicy", "awscc_eventschemas_registry_policy", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
