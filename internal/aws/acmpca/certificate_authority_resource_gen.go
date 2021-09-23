// Code generated by generators/resource/main.go; DO NOT EDIT.

package acmpca

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_acmpca_certificate_authority", certificateAuthorityResourceType)
}

// certificateAuthorityResourceType returns the Terraform awscc_acmpca_certificate_authority resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ACMPCA::CertificateAuthority resource type.
func certificateAuthorityResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"certificate_signing_request": {
			// Property: CertificateSigningRequest
			// CloudFormation resource type schema:
			// {
			//   "description": "The base64 PEM-encoded certificate signing request (CSR) for your certificate authority certificate.",
			//   "type": "string"
			// }
			Description: "The base64 PEM-encoded certificate signing request (CSR) for your certificate authority certificate.",
			Type:        types.StringType,
			Computed:    true,
		},
		"csr_extensions": {
			// Property: CsrExtensions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Structure that contains CSR pass though extensions information.",
			//   "properties": {
			//     "KeyUsage": {
			//       "additionalProperties": false,
			//       "description": "Structure that contains X.509 KeyUsage information.",
			//       "properties": {
			//         "CRLSign": {
			//           "default": false,
			//           "type": "boolean"
			//         },
			//         "DataEncipherment": {
			//           "default": false,
			//           "type": "boolean"
			//         },
			//         "DecipherOnly": {
			//           "default": false,
			//           "type": "boolean"
			//         },
			//         "DigitalSignature": {
			//           "default": false,
			//           "type": "boolean"
			//         },
			//         "EncipherOnly": {
			//           "default": false,
			//           "type": "boolean"
			//         },
			//         "KeyAgreement": {
			//           "default": false,
			//           "type": "boolean"
			//         },
			//         "KeyCertSign": {
			//           "default": false,
			//           "type": "boolean"
			//         },
			//         "KeyEncipherment": {
			//           "default": false,
			//           "type": "boolean"
			//         },
			//         "NonRepudiation": {
			//           "default": false,
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "SubjectInformationAccess": {
			//       "description": "Array of X.509 AccessDescription.",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "Structure that contains X.509 AccessDescription information.",
			//         "properties": {
			//           "AccessLocation": {
			//             "additionalProperties": false,
			//             "description": "Structure that contains X.509 GeneralName information. Assign one and ONLY one field.",
			//             "properties": {
			//               "DirectoryName": {
			//                 "additionalProperties": false,
			//                 "description": "Structure that contains X.500 distinguished name information for your CA.",
			//                 "properties": {
			//                   "CommonName": {
			//                     "type": "string"
			//                   },
			//                   "Country": {
			//                     "type": "string"
			//                   },
			//                   "DistinguishedNameQualifier": {
			//                     "type": "string"
			//                   },
			//                   "GenerationQualifier": {
			//                     "type": "string"
			//                   },
			//                   "GivenName": {
			//                     "type": "string"
			//                   },
			//                   "Initials": {
			//                     "type": "string"
			//                   },
			//                   "Locality": {
			//                     "type": "string"
			//                   },
			//                   "Organization": {
			//                     "type": "string"
			//                   },
			//                   "OrganizationalUnit": {
			//                     "type": "string"
			//                   },
			//                   "Pseudonym": {
			//                     "type": "string"
			//                   },
			//                   "SerialNumber": {
			//                     "type": "string"
			//                   },
			//                   "State": {
			//                     "type": "string"
			//                   },
			//                   "Surname": {
			//                     "type": "string"
			//                   },
			//                   "Title": {
			//                     "type": "string"
			//                   }
			//                 },
			//                 "type": "object"
			//               },
			//               "DnsName": {
			//                 "description": "String that contains X.509 DnsName information.",
			//                 "type": "string"
			//               },
			//               "EdiPartyName": {
			//                 "additionalProperties": false,
			//                 "description": "Structure that contains X.509 EdiPartyName information.",
			//                 "properties": {
			//                   "NameAssigner": {
			//                     "type": "string"
			//                   },
			//                   "PartyName": {
			//                     "type": "string"
			//                   }
			//                 },
			//                 "required": [
			//                   "PartyName",
			//                   "NameAssigner"
			//                 ],
			//                 "type": "object"
			//               },
			//               "IpAddress": {
			//                 "description": "String that contains X.509 IpAddress information.",
			//                 "type": "string"
			//               },
			//               "OtherName": {
			//                 "additionalProperties": false,
			//                 "description": "Structure that contains X.509 OtherName information.",
			//                 "properties": {
			//                   "TypeId": {
			//                     "description": "String that contains X.509 ObjectIdentifier information.",
			//                     "type": "string"
			//                   },
			//                   "Value": {
			//                     "type": "string"
			//                   }
			//                 },
			//                 "required": [
			//                   "TypeId",
			//                   "Value"
			//                 ],
			//                 "type": "object"
			//               },
			//               "RegisteredId": {
			//                 "description": "String that contains X.509 ObjectIdentifier information.",
			//                 "type": "string"
			//               },
			//               "Rfc822Name": {
			//                 "description": "String that contains X.509 Rfc822Name information.",
			//                 "type": "string"
			//               },
			//               "UniformResourceIdentifier": {
			//                 "description": "String that contains X.509 UniformResourceIdentifier information.",
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "AccessMethod": {
			//             "additionalProperties": false,
			//             "description": "Structure that contains X.509 AccessMethod information. Assign one and ONLY one field.",
			//             "properties": {
			//               "AccessMethodType": {
			//                 "description": "Pre-defined enum string for X.509 AccessMethod ObjectIdentifiers.",
			//                 "type": "string"
			//               },
			//               "CustomObjectIdentifier": {
			//                 "description": "String that contains X.509 ObjectIdentifier information.",
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "required": [
			//           "AccessMethod",
			//           "AccessLocation"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Structure that contains CSR pass though extensions information.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"key_usage": {
						// Property: KeyUsage
						Description: "Structure that contains X.509 KeyUsage information.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"crl_sign": {
									// Property: CRLSign
									Type:     types.BoolType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										DefaultValue(types.Bool{Value: false}),
									},
								},
								"data_encipherment": {
									// Property: DataEncipherment
									Type:     types.BoolType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										DefaultValue(types.Bool{Value: false}),
									},
								},
								"decipher_only": {
									// Property: DecipherOnly
									Type:     types.BoolType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										DefaultValue(types.Bool{Value: false}),
									},
								},
								"digital_signature": {
									// Property: DigitalSignature
									Type:     types.BoolType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										DefaultValue(types.Bool{Value: false}),
									},
								},
								"encipher_only": {
									// Property: EncipherOnly
									Type:     types.BoolType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										DefaultValue(types.Bool{Value: false}),
									},
								},
								"key_agreement": {
									// Property: KeyAgreement
									Type:     types.BoolType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										DefaultValue(types.Bool{Value: false}),
									},
								},
								"key_cert_sign": {
									// Property: KeyCertSign
									Type:     types.BoolType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										DefaultValue(types.Bool{Value: false}),
									},
								},
								"key_encipherment": {
									// Property: KeyEncipherment
									Type:     types.BoolType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										DefaultValue(types.Bool{Value: false}),
									},
								},
								"non_repudiation": {
									// Property: NonRepudiation
									Type:     types.BoolType,
									Optional: true,
									Computed: true,
									PlanModifiers: []tfsdk.AttributePlanModifier{
										DefaultValue(types.Bool{Value: false}),
									},
								},
							},
						),
						Optional: true,
					},
					"subject_information_access": {
						// Property: SubjectInformationAccess
						Description: "Array of X.509 AccessDescription.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"access_location": {
									// Property: AccessLocation
									Description: "Structure that contains X.509 GeneralName information. Assign one and ONLY one field.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"directory_name": {
												// Property: DirectoryName
												Description: "Structure that contains X.500 distinguished name information for your CA.",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"common_name": {
															// Property: CommonName
															Type:     types.StringType,
															Optional: true,
														},
														"country": {
															// Property: Country
															Type:     types.StringType,
															Optional: true,
														},
														"distinguished_name_qualifier": {
															// Property: DistinguishedNameQualifier
															Type:     types.StringType,
															Optional: true,
														},
														"generation_qualifier": {
															// Property: GenerationQualifier
															Type:     types.StringType,
															Optional: true,
														},
														"given_name": {
															// Property: GivenName
															Type:     types.StringType,
															Optional: true,
														},
														"initials": {
															// Property: Initials
															Type:     types.StringType,
															Optional: true,
														},
														"locality": {
															// Property: Locality
															Type:     types.StringType,
															Optional: true,
														},
														"organization": {
															// Property: Organization
															Type:     types.StringType,
															Optional: true,
														},
														"organizational_unit": {
															// Property: OrganizationalUnit
															Type:     types.StringType,
															Optional: true,
														},
														"pseudonym": {
															// Property: Pseudonym
															Type:     types.StringType,
															Optional: true,
														},
														"serial_number": {
															// Property: SerialNumber
															Type:     types.StringType,
															Optional: true,
														},
														"state": {
															// Property: State
															Type:     types.StringType,
															Optional: true,
														},
														"surname": {
															// Property: Surname
															Type:     types.StringType,
															Optional: true,
														},
														"title": {
															// Property: Title
															Type:     types.StringType,
															Optional: true,
														},
													},
												),
												Optional: true,
											},
											"dns_name": {
												// Property: DnsName
												Description: "String that contains X.509 DnsName information.",
												Type:        types.StringType,
												Optional:    true,
											},
											"edi_party_name": {
												// Property: EdiPartyName
												Description: "Structure that contains X.509 EdiPartyName information.",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"name_assigner": {
															// Property: NameAssigner
															Type:     types.StringType,
															Required: true,
														},
														"party_name": {
															// Property: PartyName
															Type:     types.StringType,
															Required: true,
														},
													},
												),
												Optional: true,
											},
											"ip_address": {
												// Property: IpAddress
												Description: "String that contains X.509 IpAddress information.",
												Type:        types.StringType,
												Optional:    true,
											},
											"other_name": {
												// Property: OtherName
												Description: "Structure that contains X.509 OtherName information.",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"type_id": {
															// Property: TypeId
															Description: "String that contains X.509 ObjectIdentifier information.",
															Type:        types.StringType,
															Required:    true,
														},
														"value": {
															// Property: Value
															Type:     types.StringType,
															Required: true,
														},
													},
												),
												Optional: true,
											},
											"registered_id": {
												// Property: RegisteredId
												Description: "String that contains X.509 ObjectIdentifier information.",
												Type:        types.StringType,
												Optional:    true,
											},
											"rfc_822_name": {
												// Property: Rfc822Name
												Description: "String that contains X.509 Rfc822Name information.",
												Type:        types.StringType,
												Optional:    true,
											},
											"uniform_resource_identifier": {
												// Property: UniformResourceIdentifier
												Description: "String that contains X.509 UniformResourceIdentifier information.",
												Type:        types.StringType,
												Optional:    true,
											},
										},
									),
									Required: true,
								},
								"access_method": {
									// Property: AccessMethod
									Description: "Structure that contains X.509 AccessMethod information. Assign one and ONLY one field.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"access_method_type": {
												// Property: AccessMethodType
												Description: "Pre-defined enum string for X.509 AccessMethod ObjectIdentifiers.",
												Type:        types.StringType,
												Optional:    true,
											},
											"custom_object_identifier": {
												// Property: CustomObjectIdentifier
												Description: "String that contains X.509 ObjectIdentifier information.",
												Type:        types.StringType,
												Optional:    true,
											},
										},
									),
									Required: true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"key_algorithm": {
			// Property: KeyAlgorithm
			// CloudFormation resource type schema:
			// {
			//   "description": "Public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.",
			//   "type": "string"
			// }
			Description: "Public key algorithm and size, in bits, of the key pair that your CA creates when it issues a certificate.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"key_storage_security_standard": {
			// Property: KeyStorageSecurityStandard
			// CloudFormation resource type schema:
			// {
			//   "description": "KeyStorageSecurityStadard defines a cryptographic key management compliance standard used for handling CA keys.",
			//   "type": "string"
			// }
			Description: "KeyStorageSecurityStadard defines a cryptographic key management compliance standard used for handling CA keys.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"revocation_configuration": {
			// Property: RevocationConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Certificate Authority revocation information.",
			//   "properties": {
			//     "CrlConfiguration": {
			//       "additionalProperties": false,
			//       "description": "Your certificate authority can create and maintain a certificate revocation list (CRL). A CRL contains information about certificates that have been revoked.",
			//       "properties": {
			//         "CustomCname": {
			//           "type": "string"
			//         },
			//         "Enabled": {
			//           "type": "boolean"
			//         },
			//         "ExpirationInDays": {
			//           "type": "integer"
			//         },
			//         "S3BucketName": {
			//           "type": "string"
			//         },
			//         "S3ObjectAcl": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "OcspConfiguration": {
			//       "additionalProperties": false,
			//       "description": "Helps to configure online certificate status protocol (OCSP) responder for your certificate authority",
			//       "properties": {
			//         "Enabled": {
			//           "type": "boolean"
			//         },
			//         "OcspCustomCname": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Certificate Authority revocation information.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"crl_configuration": {
						// Property: CrlConfiguration
						Description: "Your certificate authority can create and maintain a certificate revocation list (CRL). A CRL contains information about certificates that have been revoked.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"custom_cname": {
									// Property: CustomCname
									Type:     types.StringType,
									Optional: true,
								},
								"enabled": {
									// Property: Enabled
									Type:     types.BoolType,
									Optional: true,
								},
								"expiration_in_days": {
									// Property: ExpirationInDays
									Type:     types.NumberType,
									Optional: true,
								},
								"s3_bucket_name": {
									// Property: S3BucketName
									Type:     types.StringType,
									Optional: true,
								},
								"s3_object_acl": {
									// Property: S3ObjectAcl
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"ocsp_configuration": {
						// Property: OcspConfiguration
						Description: "Helps to configure online certificate status protocol (OCSP) responder for your certificate authority",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Type:     types.BoolType,
									Optional: true,
								},
								"ocsp_custom_cname": {
									// Property: OcspCustomCname
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"signing_algorithm": {
			// Property: SigningAlgorithm
			// CloudFormation resource type schema:
			// {
			//   "description": "Algorithm your CA uses to sign certificate requests.",
			//   "type": "string"
			// }
			Description: "Algorithm your CA uses to sign certificate requests.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"subject": {
			// Property: Subject
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Structure that contains X.500 distinguished name information for your CA.",
			//   "properties": {
			//     "CommonName": {
			//       "type": "string"
			//     },
			//     "Country": {
			//       "type": "string"
			//     },
			//     "DistinguishedNameQualifier": {
			//       "type": "string"
			//     },
			//     "GenerationQualifier": {
			//       "type": "string"
			//     },
			//     "GivenName": {
			//       "type": "string"
			//     },
			//     "Initials": {
			//       "type": "string"
			//     },
			//     "Locality": {
			//       "type": "string"
			//     },
			//     "Organization": {
			//       "type": "string"
			//     },
			//     "OrganizationalUnit": {
			//       "type": "string"
			//     },
			//     "Pseudonym": {
			//       "type": "string"
			//     },
			//     "SerialNumber": {
			//       "type": "string"
			//     },
			//     "State": {
			//       "type": "string"
			//     },
			//     "Surname": {
			//       "type": "string"
			//     },
			//     "Title": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Structure that contains X.500 distinguished name information for your CA.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"common_name": {
						// Property: CommonName
						Type:     types.StringType,
						Optional: true,
					},
					"country": {
						// Property: Country
						Type:     types.StringType,
						Optional: true,
					},
					"distinguished_name_qualifier": {
						// Property: DistinguishedNameQualifier
						Type:     types.StringType,
						Optional: true,
					},
					"generation_qualifier": {
						// Property: GenerationQualifier
						Type:     types.StringType,
						Optional: true,
					},
					"given_name": {
						// Property: GivenName
						Type:     types.StringType,
						Optional: true,
					},
					"initials": {
						// Property: Initials
						Type:     types.StringType,
						Optional: true,
					},
					"locality": {
						// Property: Locality
						Type:     types.StringType,
						Optional: true,
					},
					"organization": {
						// Property: Organization
						Type:     types.StringType,
						Optional: true,
					},
					"organizational_unit": {
						// Property: OrganizationalUnit
						Type:     types.StringType,
						Optional: true,
					},
					"pseudonym": {
						// Property: Pseudonym
						Type:     types.StringType,
						Optional: true,
					},
					"serial_number": {
						// Property: SerialNumber
						Type:     types.StringType,
						Optional: true,
					},
					"state": {
						// Property: State
						Type:     types.StringType,
						Optional: true,
					},
					"surname": {
						// Property: Surname
						Type:     types.StringType,
						Optional: true,
					},
					"title": {
						// Property: Title
						Type:     types.StringType,
						Optional: true,
					},
				},
			),
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
			// Subject is a write-only property.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of the certificate authority.",
			//   "type": "string"
			// }
			Description: "The type of the certificate authority.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Private certificate authority.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ACMPCA::CertificateAuthority").WithTerraformTypeName("awscc_acmpca_certificate_authority")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_location":               "AccessLocation",
		"access_method":                 "AccessMethod",
		"access_method_type":            "AccessMethodType",
		"arn":                           "Arn",
		"certificate_signing_request":   "CertificateSigningRequest",
		"common_name":                   "CommonName",
		"country":                       "Country",
		"crl_configuration":             "CrlConfiguration",
		"crl_sign":                      "CRLSign",
		"csr_extensions":                "CsrExtensions",
		"custom_cname":                  "CustomCname",
		"custom_object_identifier":      "CustomObjectIdentifier",
		"data_encipherment":             "DataEncipherment",
		"decipher_only":                 "DecipherOnly",
		"digital_signature":             "DigitalSignature",
		"directory_name":                "DirectoryName",
		"distinguished_name_qualifier":  "DistinguishedNameQualifier",
		"dns_name":                      "DnsName",
		"edi_party_name":                "EdiPartyName",
		"enabled":                       "Enabled",
		"encipher_only":                 "EncipherOnly",
		"expiration_in_days":            "ExpirationInDays",
		"generation_qualifier":          "GenerationQualifier",
		"given_name":                    "GivenName",
		"initials":                      "Initials",
		"ip_address":                    "IpAddress",
		"key":                           "Key",
		"key_agreement":                 "KeyAgreement",
		"key_algorithm":                 "KeyAlgorithm",
		"key_cert_sign":                 "KeyCertSign",
		"key_encipherment":              "KeyEncipherment",
		"key_storage_security_standard": "KeyStorageSecurityStandard",
		"key_usage":                     "KeyUsage",
		"locality":                      "Locality",
		"name_assigner":                 "NameAssigner",
		"non_repudiation":               "NonRepudiation",
		"ocsp_configuration":            "OcspConfiguration",
		"ocsp_custom_cname":             "OcspCustomCname",
		"organization":                  "Organization",
		"organizational_unit":           "OrganizationalUnit",
		"other_name":                    "OtherName",
		"party_name":                    "PartyName",
		"pseudonym":                     "Pseudonym",
		"registered_id":                 "RegisteredId",
		"revocation_configuration":      "RevocationConfiguration",
		"rfc_822_name":                  "Rfc822Name",
		"s3_bucket_name":                "S3BucketName",
		"s3_object_acl":                 "S3ObjectAcl",
		"serial_number":                 "SerialNumber",
		"signing_algorithm":             "SigningAlgorithm",
		"state":                         "State",
		"subject":                       "Subject",
		"subject_information_access":    "SubjectInformationAccess",
		"surname":                       "Surname",
		"tags":                          "Tags",
		"title":                         "Title",
		"type":                          "Type",
		"type_id":                       "TypeId",
		"uniform_resource_identifier":   "UniformResourceIdentifier",
		"value":                         "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Subject",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
