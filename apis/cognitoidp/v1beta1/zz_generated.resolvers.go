/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-aws/apis/acm/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this IdentityProvider.
func (mg *IdentityProvider) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserPoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserPoolID")
	}
	mg.Spec.ForProvider.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserPoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ResourceServer.
func (mg *ResourceServer) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserPoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserPoolID")
	}
	mg.Spec.ForProvider.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserPoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RiskConfiguration.
func (mg *RiskConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserPoolID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserPoolID")
	}
	mg.Spec.ForProvider.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserPoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this User.
func (mg *User) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserPoolID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserPoolID")
	}
	mg.Spec.ForProvider.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserPoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserGroup.
func (mg *UserGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserPoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserPoolID")
	}
	mg.Spec.ForProvider.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserPoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserInGroup.
func (mg *UserInGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GroupName),
		Extract:      resource.ExtractParamPath("name", false),
		Reference:    mg.Spec.ForProvider.GroupNameRef,
		Selector:     mg.Spec.ForProvider.GroupNameSelector,
		To: reference.To{
			List:    &UserGroupList{},
			Managed: &UserGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupName")
	}
	mg.Spec.ForProvider.GroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserPoolID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserPoolID")
	}
	mg.Spec.ForProvider.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserPoolIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Username),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UsernameRef,
		Selector:     mg.Spec.ForProvider.UsernameSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Username")
	}
	mg.Spec.ForProvider.Username = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UsernameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserPool.
func (mg *UserPool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SMSConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SMSConfiguration[i3].SnsCallerArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.SMSConfiguration[i3].SnsCallerArnRef,
			Selector:     mg.Spec.ForProvider.SMSConfiguration[i3].SnsCallerArnSelector,
			To: reference.To{
				List:    &v1beta1.RoleList{},
				Managed: &v1beta1.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SMSConfiguration[i3].SnsCallerArn")
		}
		mg.Spec.ForProvider.SMSConfiguration[i3].SnsCallerArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SMSConfiguration[i3].SnsCallerArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this UserPoolDomain.
func (mg *UserPoolDomain) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CertificateArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.CertificateArnRef,
		Selector:     mg.Spec.ForProvider.CertificateArnSelector,
		To: reference.To{
			List:    &v1beta11.CertificateList{},
			Managed: &v1beta11.Certificate{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CertificateArn")
	}
	mg.Spec.ForProvider.CertificateArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CertificateArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserPoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserPoolID")
	}
	mg.Spec.ForProvider.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserPoolIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserPoolUICustomization.
func (mg *UserPoolUICustomization) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserPoolID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserPoolIDRef,
		Selector:     mg.Spec.ForProvider.UserPoolIDSelector,
		To: reference.To{
			List:    &UserPoolList{},
			Managed: &UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserPoolID")
	}
	mg.Spec.ForProvider.UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserPoolIDRef = rsp.ResolvedReference

	return nil
}
