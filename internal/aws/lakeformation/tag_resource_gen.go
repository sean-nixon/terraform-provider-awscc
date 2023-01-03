// Code generated by generators/resource/main.go; DO NOT EDIT.

package lakeformation

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_lakeformation_tag", tagResource)
}

// tagResource returns the Terraform awscc_lakeformation_tag resource.
// This Terraform resource corresponds to the CloudFormation AWS::LakeFormation::Tag resource.
func tagResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CatalogId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.",
		//	  "maxLength": 12,
		//	  "minLength": 12,
		//	  "type": "string"
		//	}
		"catalog_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(12, 12),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TagKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The key-name for the LF-tag.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^([{a-zA-Z}{\\s}{0-9}_.:\\/=+\\-@%]*)$",
		//	  "type": "string"
		//	}
		"tag_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The key-name for the LF-tag.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("^([{a-zA-Z}{\\s}{0-9}_.:\\/=+\\-@%]*)$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TagValues
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of possible values an attribute can take.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 256,
		//	    "minLength": 0,
		//	    "pattern": "^([{a-zA-Z}{\\s}{0-9}_.:\\*\\/=+\\-@%]*)$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 50,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"tag_values": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of possible values an attribute can take.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 50),
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(0, 256),
					stringvalidator.RegexMatches(regexp.MustCompile("^([{a-zA-Z}{\\s}{0-9}_.:\\*\\/=+\\-@%]*)$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				Multiset(),
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
		Description: "A resource schema representing a Lake Formation Tag.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::LakeFormation::Tag").WithTerraformTypeName("awscc_lakeformation_tag")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"catalog_id": "CatalogId",
		"tag_key":    "TagKey",
		"tag_values": "TagValues",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
