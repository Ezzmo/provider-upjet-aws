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
	common "github.com/upbound/provider-aws/config/common"
	lambda "github.com/upbound/provider-aws/config/common/apis/lambda"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *Authorizer) ResolveReferences( // ResolveReferences of this Authorizer.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("apigatewayv2.aws.upbound.io", "v1beta2", "API", "APIList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.APIIDRef,
			Selector:     mg.Spec.ForProvider.APIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta2", "Function", "FunctionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AuthorizerURI),
			Extract:      lambda.FunctionInvokeARN(),
			Reference:    mg.Spec.ForProvider.AuthorizerURIRef,
			Selector:     mg.Spec.ForProvider.AuthorizerURISelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AuthorizerURI")
	}
	mg.Spec.ForProvider.AuthorizerURI = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AuthorizerURIRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigatewayv2.aws.upbound.io", "v1beta2", "API", "APIList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.APIIDRef,
			Selector:     mg.Spec.InitProvider.APIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.APIID")
	}
	mg.Spec.InitProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.APIIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta2", "Function", "FunctionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AuthorizerURI),
			Extract:      lambda.FunctionInvokeARN(),
			Reference:    mg.Spec.InitProvider.AuthorizerURIRef,
			Selector:     mg.Spec.InitProvider.AuthorizerURISelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AuthorizerURI")
	}
	mg.Spec.InitProvider.AuthorizerURI = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AuthorizerURIRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DomainName.
func (mg *DomainName) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.DomainNameConfiguration != nil {
		{
			m, l, err = apisresolver.GetManagedResource("acm.aws.upbound.io", "v1beta2", "Certificate", "CertificateList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DomainNameConfiguration.CertificateArn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.ForProvider.DomainNameConfiguration.CertificateArnRef,
				Selector:     mg.Spec.ForProvider.DomainNameConfiguration.CertificateArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DomainNameConfiguration.CertificateArn")
		}
		mg.Spec.ForProvider.DomainNameConfiguration.CertificateArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DomainNameConfiguration.CertificateArnRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.DomainNameConfiguration != nil {
		{
			m, l, err = apisresolver.GetManagedResource("acm.aws.upbound.io", "v1beta2", "Certificate", "CertificateList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DomainNameConfiguration.CertificateArn),
				Extract:      common.ARNExtractor(),
				Reference:    mg.Spec.InitProvider.DomainNameConfiguration.CertificateArnRef,
				Selector:     mg.Spec.InitProvider.DomainNameConfiguration.CertificateArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DomainNameConfiguration.CertificateArn")
		}
		mg.Spec.InitProvider.DomainNameConfiguration.CertificateArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DomainNameConfiguration.CertificateArnRef = rsp.ResolvedReference

	}

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
		m, l, err = apisresolver.GetManagedResource("apigatewayv2.aws.upbound.io", "v1beta2", "API", "APIList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.APIIDRef,
			Selector:     mg.Spec.ForProvider.APIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigatewayv2.aws.upbound.io", "v1beta1", "VPCLink", "VPCLinkList")
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
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CredentialsArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.CredentialsArnRef,
			Selector:     mg.Spec.ForProvider.CredentialsArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CredentialsArn")
	}
	mg.Spec.ForProvider.CredentialsArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CredentialsArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta2", "Function", "FunctionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IntegrationURI),
			Extract:      resource.ExtractParamPath("invoke_arn", true),
			Reference:    mg.Spec.ForProvider.IntegrationURIRef,
			Selector:     mg.Spec.ForProvider.IntegrationURISelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IntegrationURI")
	}
	mg.Spec.ForProvider.IntegrationURI = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IntegrationURIRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigatewayv2.aws.upbound.io", "v1beta2", "API", "APIList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.APIIDRef,
			Selector:     mg.Spec.InitProvider.APIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.APIID")
	}
	mg.Spec.InitProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.APIIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigatewayv2.aws.upbound.io", "v1beta1", "VPCLink", "VPCLinkList")
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
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io", "v1beta1", "Role", "RoleList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CredentialsArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.CredentialsArnRef,
			Selector:     mg.Spec.InitProvider.CredentialsArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CredentialsArn")
	}
	mg.Spec.InitProvider.CredentialsArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CredentialsArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta2", "Function", "FunctionList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IntegrationURI),
			Extract:      resource.ExtractParamPath("invoke_arn", true),
			Reference:    mg.Spec.InitProvider.IntegrationURIRef,
			Selector:     mg.Spec.InitProvider.IntegrationURISelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IntegrationURI")
	}
	mg.Spec.InitProvider.IntegrationURI = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IntegrationURIRef = rsp.ResolvedReference

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
		m, l, err = apisresolver.GetManagedResource("apigatewayv2.aws.upbound.io", "v1beta2", "API", "APIList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.APIIDRef,
			Selector:     mg.Spec.ForProvider.APIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigatewayv2.aws.upbound.io", "v1beta1", "Deployment", "DeploymentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DeploymentID),
			Extract:      reference.ExternalName(),
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
		m, l, err = apisresolver.GetManagedResource("apigatewayv2.aws.upbound.io", "v1beta2", "API", "APIList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.APIIDRef,
			Selector:     mg.Spec.InitProvider.APIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.APIID")
	}
	mg.Spec.InitProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.APIIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("apigatewayv2.aws.upbound.io", "v1beta1", "Deployment", "DeploymentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DeploymentID),
			Extract:      reference.ExternalName(),
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

	return nil
}