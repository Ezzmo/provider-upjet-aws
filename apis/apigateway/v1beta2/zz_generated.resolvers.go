// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *DocumentationPart) ResolveReferences( // ResolveReferences of this DocumentationPart.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta2", "RestAPI", "RestAPIList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RestAPIID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.RestAPIIDRef,
			Selector:     mg.Spec.ForProvider.RestAPIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RestAPIID")
	}
	mg.Spec.ForProvider.RestAPIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RestAPIIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta2", "RestAPI", "RestAPIList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RestAPIID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.RestAPIIDRef,
			Selector:     mg.Spec.InitProvider.RestAPIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RestAPIID")
	}
	mg.Spec.InitProvider.RestAPIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RestAPIIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DomainName.
func (mg *DomainName) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("acm.aws.upbound.io", "v1beta1", "CertificateValidation", "CertificateValidationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CertificateArn),
			Extract:      resource.ExtractParamPath("certificate_arn", false),
			Reference:    mg.Spec.ForProvider.CertificateArnRef,
			Selector:     mg.Spec.ForProvider.CertificateArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CertificateArn")
	}
	mg.Spec.ForProvider.CertificateArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CertificateArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("acm.aws.upbound.io", "v1beta1", "CertificateValidation", "CertificateValidationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RegionalCertificateArn),
			Extract:      resource.ExtractParamPath("certificate_arn", false),
			Reference:    mg.Spec.ForProvider.RegionalCertificateArnRef,
			Selector:     mg.Spec.ForProvider.RegionalCertificateArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RegionalCertificateArn")
	}
	mg.Spec.ForProvider.RegionalCertificateArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RegionalCertificateArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("acm.aws.upbound.io", "v1beta1", "CertificateValidation", "CertificateValidationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CertificateArn),
			Extract:      resource.ExtractParamPath("certificate_arn", false),
			Reference:    mg.Spec.InitProvider.CertificateArnRef,
			Selector:     mg.Spec.InitProvider.CertificateArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CertificateArn")
	}
	mg.Spec.InitProvider.CertificateArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CertificateArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("acm.aws.upbound.io", "v1beta1", "CertificateValidation", "CertificateValidationList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RegionalCertificateArn),
			Extract:      resource.ExtractParamPath("certificate_arn", false),
			Reference:    mg.Spec.InitProvider.RegionalCertificateArnRef,
			Selector:     mg.Spec.InitProvider.RegionalCertificateArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RegionalCertificateArn")
	}
	mg.Spec.InitProvider.RegionalCertificateArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RegionalCertificateArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Integration.
func (mg *Integration) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta1", "VPCLink", "VPCLinkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ConnectionID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ConnectionIDRef,
			Selector:     mg.Spec.ForProvider.ConnectionIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ConnectionID")
	}
	mg.Spec.ForProvider.ConnectionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ConnectionIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta1", "Method", "MethodList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.HTTPMethod),
			Extract:      resource.ExtractParamPath("http_method", false),
			Reference:    mg.Spec.ForProvider.HTTPMethodRef,
			Selector:     mg.Spec.ForProvider.HTTPMethodSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.HTTPMethod")
	}
	mg.Spec.ForProvider.HTTPMethod = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.HTTPMethodRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta1", "Resource", "ResourceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ResourceIDRef,
			Selector:     mg.Spec.ForProvider.ResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceID")
	}
	mg.Spec.ForProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta2", "RestAPI", "RestAPIList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RestAPIID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.RestAPIIDRef,
			Selector:     mg.Spec.ForProvider.RestAPIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RestAPIID")
	}
	mg.Spec.ForProvider.RestAPIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RestAPIIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta2", "Function", "FunctionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.URI),
			Extract:      resource.ExtractParamPath("invoke_arn", true),
			Reference:    mg.Spec.ForProvider.URIRef,
			Selector:     mg.Spec.ForProvider.URISelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.URI")
	}
	mg.Spec.ForProvider.URI = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.URIRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta1", "VPCLink", "VPCLinkList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ConnectionID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ConnectionIDRef,
			Selector:     mg.Spec.InitProvider.ConnectionIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ConnectionID")
	}
	mg.Spec.InitProvider.ConnectionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ConnectionIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta1", "Method", "MethodList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.HTTPMethod),
			Extract:      resource.ExtractParamPath("http_method", false),
			Reference:    mg.Spec.InitProvider.HTTPMethodRef,
			Selector:     mg.Spec.InitProvider.HTTPMethodSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.HTTPMethod")
	}
	mg.Spec.InitProvider.HTTPMethod = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.HTTPMethodRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta1", "Resource", "ResourceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ResourceIDRef,
			Selector:     mg.Spec.InitProvider.ResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceID")
	}
	mg.Spec.InitProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta2", "RestAPI", "RestAPIList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RestAPIID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.RestAPIIDRef,
			Selector:     mg.Spec.InitProvider.RestAPIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RestAPIID")
	}
	mg.Spec.InitProvider.RestAPIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RestAPIIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta2", "Function", "FunctionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.URI),
			Extract:      resource.ExtractParamPath("invoke_arn", true),
			Reference:    mg.Spec.InitProvider.URIRef,
			Selector:     mg.Spec.InitProvider.URISelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.URI")
	}
	mg.Spec.InitProvider.URI = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.URIRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MethodSettings.
func (mg *MethodSettings) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta2", "RestAPI", "RestAPIList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RestAPIID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.RestAPIIDRef,
			Selector:     mg.Spec.ForProvider.RestAPIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RestAPIID")
	}
	mg.Spec.ForProvider.RestAPIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RestAPIIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta2", "Stage", "StageList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StageName),
			Extract:      resource.ExtractParamPath("stage_name", false),
			Reference:    mg.Spec.ForProvider.StageNameRef,
			Selector:     mg.Spec.ForProvider.StageNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StageName")
	}
	mg.Spec.ForProvider.StageName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StageNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta2", "RestAPI", "RestAPIList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RestAPIID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.RestAPIIDRef,
			Selector:     mg.Spec.InitProvider.RestAPIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RestAPIID")
	}
	mg.Spec.InitProvider.RestAPIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RestAPIIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta2", "Stage", "StageList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StageName),
			Extract:      resource.ExtractParamPath("stage_name", false),
			Reference:    mg.Spec.InitProvider.StageNameRef,
			Selector:     mg.Spec.InitProvider.StageNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.StageName")
	}
	mg.Spec.InitProvider.StageName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.StageNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RestAPI.
func (mg *RestAPI) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	if mg.Spec.ForProvider.EndpointConfiguration != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta2", "VPCEndpoint", "VPCEndpointList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.EndpointConfiguration.VPCEndpointIds),
				Extract:       resource.ExtractResourceID(),
				References:    mg.Spec.ForProvider.EndpointConfiguration.VPCEndpointIdsRefs,
				Selector:      mg.Spec.ForProvider.EndpointConfiguration.VPCEndpointIdsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EndpointConfiguration.VPCEndpointIds")
		}
		mg.Spec.ForProvider.EndpointConfiguration.VPCEndpointIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.EndpointConfiguration.VPCEndpointIdsRefs = mrsp.ResolvedReferences

	}
	if mg.Spec.InitProvider.EndpointConfiguration != nil {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta2", "VPCEndpoint", "VPCEndpointList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.EndpointConfiguration.VPCEndpointIds),
				Extract:       resource.ExtractResourceID(),
				References:    mg.Spec.InitProvider.EndpointConfiguration.VPCEndpointIdsRefs,
				Selector:      mg.Spec.InitProvider.EndpointConfiguration.VPCEndpointIdsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EndpointConfiguration.VPCEndpointIds")
		}
		mg.Spec.InitProvider.EndpointConfiguration.VPCEndpointIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.EndpointConfiguration.VPCEndpointIdsRefs = mrsp.ResolvedReferences

	}

	return nil
}

// ResolveReferences of this Stage.
func (mg *Stage) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta1", "Deployment", "DeploymentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DeploymentID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DeploymentIDRef,
			Selector:     mg.Spec.ForProvider.DeploymentIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DeploymentID")
	}
	mg.Spec.ForProvider.DeploymentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DeploymentIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta2", "RestAPI", "RestAPIList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RestAPIID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.RestAPIIDRef,
			Selector:     mg.Spec.ForProvider.RestAPIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RestAPIID")
	}
	mg.Spec.ForProvider.RestAPIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RestAPIIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta1", "Deployment", "DeploymentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DeploymentID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.DeploymentIDRef,
			Selector:     mg.Spec.InitProvider.DeploymentIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DeploymentID")
	}
	mg.Spec.InitProvider.DeploymentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DeploymentIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta2", "RestAPI", "RestAPIList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RestAPIID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.RestAPIIDRef,
			Selector:     mg.Spec.InitProvider.RestAPIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RestAPIID")
	}
	mg.Spec.InitProvider.RestAPIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RestAPIIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UsagePlan.
func (mg *UsagePlan) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.APIStages); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta2", "RestAPI", "RestAPIList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIStages[i3].APIID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.APIStages[i3].APIIDRef,
				Selector:     mg.Spec.ForProvider.APIStages[i3].APIIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.APIStages[i3].APIID")
		}
		mg.Spec.ForProvider.APIStages[i3].APIID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.APIStages[i3].APIIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.APIStages); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta2", "Stage", "StageList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIStages[i3].Stage),
				Extract:      resource.ExtractParamPath("stage_name", false),
				Reference:    mg.Spec.ForProvider.APIStages[i3].StageRef,
				Selector:     mg.Spec.ForProvider.APIStages[i3].StageSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.APIStages[i3].Stage")
		}
		mg.Spec.ForProvider.APIStages[i3].Stage = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.APIStages[i3].StageRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.APIStages); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta2", "RestAPI", "RestAPIList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIStages[i3].APIID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.APIStages[i3].APIIDRef,
				Selector:     mg.Spec.InitProvider.APIStages[i3].APIIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.APIStages[i3].APIID")
		}
		mg.Spec.InitProvider.APIStages[i3].APIID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.APIStages[i3].APIIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.APIStages); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("apigateway.aws.upbound.io", "v1beta2", "Stage", "StageList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIStages[i3].Stage),
				Extract:      resource.ExtractParamPath("stage_name", false),
				Reference:    mg.Spec.InitProvider.APIStages[i3].StageRef,
				Selector:     mg.Spec.InitProvider.APIStages[i3].StageSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.APIStages[i3].Stage")
		}
		mg.Spec.InitProvider.APIStages[i3].Stage = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.APIStages[i3].StageRef = rsp.ResolvedReference

	}

	return nil
}
