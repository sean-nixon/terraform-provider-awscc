// Code generated by generators/resource/main.go; DO NOT EDIT.

package ses

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ses_template", templateResource)
}

// templateResource returns the Terraform awscc_ses_template resource.
// This Terraform resource corresponds to the CloudFormation AWS::SES::Template resource.
func templateResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Template
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The content of the email, composed of a subject line, an HTML part, and a text-only part",
		//	  "properties": {
		//	    "HtmlPart": {
		//	      "description": "The HTML body of the email.",
		//	      "type": "string"
		//	    },
		//	    "SubjectPart": {
		//	      "description": "The subject line of the email.",
		//	      "type": "string"
		//	    },
		//	    "TemplateName": {
		//	      "description": "The name of the template.",
		//	      "maxLength": 64,
		//	      "minLength": 1,
		//	      "pattern": "^[a-zA-Z0-9_-]{1,64}$",
		//	      "type": "string"
		//	    },
		//	    "TextPart": {
		//	      "description": "The email body that is visible to recipients whose email clients do not display HTML content.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "SubjectPart"
		//	  ],
		//	  "type": "object"
		//	}
		"template": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: HtmlPart
				"html_part": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The HTML body of the email.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SubjectPart
				"subject_part": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The subject line of the email.",
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: TemplateName
				"template_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the template.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 64),
						stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]{1,64}$"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplace(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TextPart
				"text_part": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The email body that is visible to recipients whose email clients do not display HTML content.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The content of the email, composed of a subject line, an HTML part, and a text-only part",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::SES::Template",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SES::Template").WithTerraformTypeName("awscc_ses_template")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"html_part":     "HtmlPart",
		"id":            "Id",
		"subject_part":  "SubjectPart",
		"template":      "Template",
		"template_name": "TemplateName",
		"text_part":     "TextPart",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
