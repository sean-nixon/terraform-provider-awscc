// Code generated by generators/resource/main.go; DO NOT EDIT.

package cloudformation_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSCloudFormationHookDefaultVersion_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudFormation::HookDefaultVersion", "awscc_cloudformation_hook_default_version", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}

func TestAccAWSCloudFormationHookDefaultVersion_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::CloudFormation::HookDefaultVersion", "awscc_cloudformation_hook_default_version", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				td.DeleteResource(),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}
