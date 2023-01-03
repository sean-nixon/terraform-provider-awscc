// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_mediaconnect_flow_source", flowSourceResource)
}

// flowSourceResource returns the Terraform awscc_mediaconnect_flow_source resource.
// This Terraform resource corresponds to the CloudFormation AWS::MediaConnect::FlowSource resource.
func flowSourceResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Decryption
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The type of encryption that is used on the content ingested from this source.",
		//	  "properties": {
		//	    "Algorithm": {
		//	      "description": "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
		//	      "enum": [
		//	        "aes128",
		//	        "aes192",
		//	        "aes256"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "ConstantInitializationVector": {
		//	      "description": "A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.",
		//	      "type": "string"
		//	    },
		//	    "DeviceId": {
		//	      "description": "The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
		//	      "type": "string"
		//	    },
		//	    "KeyType": {
		//	      "default": "static-key",
		//	      "description": "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
		//	      "enum": [
		//	        "speke",
		//	        "static-key"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Region": {
		//	      "description": "The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
		//	      "type": "string"
		//	    },
		//	    "ResourceId": {
		//	      "description": "An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
		//	      "type": "string"
		//	    },
		//	    "RoleArn": {
		//	      "description": "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
		//	      "type": "string"
		//	    },
		//	    "SecretArn": {
		//	      "description": " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
		//	      "type": "string"
		//	    },
		//	    "Url": {
		//	      "description": "The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Algorithm",
		//	    "RoleArn"
		//	  ],
		//	  "type": "object"
		//	}
		"decryption": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Algorithm
				"algorithm": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"aes128",
							"aes192",
							"aes256",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: ConstantInitializationVector
				"constant_initialization_vector": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DeviceId
				"device_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: KeyType
				"key_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"speke",
							"static-key",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						StringDefaultValue("static-key"),
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Region
				"region": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ResourceId
				"resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RoleArn
				"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: SecretArn
				"secret_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Url
				"url": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The type of encryption that is used on the content ingested from this source.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: EntitlementArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.",
		//	  "type": "string"
		//	}
		"entitlement_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FlowArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the flow.",
		//	  "type": "string"
		//	}
		"flow_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the flow.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IngestIp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IP address that the flow will be listening on for incoming content.",
		//	  "type": "string"
		//	}
		"ingest_ip": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IP address that the flow will be listening on for incoming content.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IngestPort
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The port that the flow will be listening on for incoming content.",
		//	  "type": "integer"
		//	}
		"ingest_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The port that the flow will be listening on for incoming content.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaxBitrate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.",
		//	  "type": "integer"
		//	}
		"max_bitrate": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaxLatency
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 2000,
		//	  "description": "The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.",
		//	  "type": "integer"
		//	}
		"max_latency": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				Int64Value(2000),
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the source.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the source.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Protocol
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The protocol that is used by the source.",
		//	  "enum": [
		//	    "zixi-push",
		//	    "rtp-fec",
		//	    "rtp",
		//	    "rist"
		//	  ],
		//	  "type": "string"
		//	}
		"protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The protocol that is used by the source.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"zixi-push",
					"rtp-fec",
					"rtp",
					"rist",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the source.",
		//	  "type": "string"
		//	}
		"source_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the source.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceIngestPort
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The port that the flow will be listening on for incoming content.(ReadOnly)",
		//	  "type": "string"
		//	}
		"source_ingest_port": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The port that the flow will be listening on for incoming content.(ReadOnly)",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StreamId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.",
		//	  "type": "string"
		//	}
		"stream_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcInterfaceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the VPC Interface this Source is configured with.",
		//	  "type": "string"
		//	}
		"vpc_interface_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the VPC Interface this Source is configured with.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: WhitelistCidr
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
		//	  "type": "string"
		//	}
		"whitelist_cidr": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
			Optional:    true,
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
		Description: "Resource schema for AWS::MediaConnect::FlowSource",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaConnect::FlowSource").WithTerraformTypeName("awscc_mediaconnect_flow_source")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"algorithm":                      "Algorithm",
		"constant_initialization_vector": "ConstantInitializationVector",
		"decryption":                     "Decryption",
		"description":                    "Description",
		"device_id":                      "DeviceId",
		"entitlement_arn":                "EntitlementArn",
		"flow_arn":                       "FlowArn",
		"ingest_ip":                      "IngestIp",
		"ingest_port":                    "IngestPort",
		"key_type":                       "KeyType",
		"max_bitrate":                    "MaxBitrate",
		"max_latency":                    "MaxLatency",
		"name":                           "Name",
		"protocol":                       "Protocol",
		"region":                         "Region",
		"resource_id":                    "ResourceId",
		"role_arn":                       "RoleArn",
		"secret_arn":                     "SecretArn",
		"source_arn":                     "SourceArn",
		"source_ingest_port":             "SourceIngestPort",
		"stream_id":                      "StreamId",
		"url":                            "Url",
		"vpc_interface_name":             "VpcInterfaceName",
		"whitelist_cidr":                 "WhitelistCidr",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
