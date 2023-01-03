// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53resolver_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRoute53ResolverFirewallRuleGroupAssociation_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53Resolver::FirewallRuleGroupAssociation", "awscc_route53resolver_firewall_rule_group_association", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
