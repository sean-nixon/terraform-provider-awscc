// Code generated by generators/resource/main.go; DO NOT EDIT.

package athena_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSAthenaWorkGroup_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Athena::WorkGroup", "awscc_athena_work_group", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
