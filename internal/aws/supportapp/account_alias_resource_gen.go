// Code generated by generators/resource/main.go; DO NOT EDIT.

package supportapp

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_supportapp_account_alias", accountAliasResource)
}

// accountAliasResource returns the Terraform awscc_supportapp_account_alias resource.
// This Terraform resource corresponds to the CloudFormation AWS::SupportApp::AccountAlias resource.
func accountAliasResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccountAlias
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An account alias associated with a customer's account.",
		//	  "maxLength": 30,
		//	  "minLength": 1,
		//	  "pattern": "^[\\w\\- ]+$",
		//	  "type": "string"
		//	}
		"account_alias": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An account alias associated with a customer's account.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 30),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\w\\- ]+$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: AccountAliasResourceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique identifier representing an alias tied to an account",
		//	  "maxLength": 29,
		//	  "minLength": 29,
		//	  "pattern": "^[\\w\\- ]+$",
		//	  "type": "string"
		//	}
		"account_alias_resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique identifier representing an alias tied to an account",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "An AWS Support App resource that creates, updates, reads, and deletes a customer's account alias.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SupportApp::AccountAlias").WithTerraformTypeName("awscc_supportapp_account_alias")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_alias":             "AccountAlias",
		"account_alias_resource_id": "AccountAliasResourceId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
