// Code generated by generators/resource/main.go; DO NOT EDIT.

package opensearchserverless_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSOpenSearchServerlessCollection_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::OpenSearchServerless::Collection", "awscc_opensearchserverless_collection", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
