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
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *SafetyRule) ResolveReferences( // ResolveReferences of this SafetyRule.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("route53recoverycontrolconfig.aws.upbound.io", "v1beta1", "RoutingControl", "RoutingControlList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.AssertedControls),
			Extract:       common.TerraformID(),
			References:    mg.Spec.ForProvider.AssertedControlsRefs,
			Selector:      mg.Spec.ForProvider.AssertedControlsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AssertedControls")
	}
	mg.Spec.ForProvider.AssertedControls = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.AssertedControlsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("route53recoverycontrolconfig.aws.upbound.io", "v1beta1", "ControlPanel", "ControlPanelList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ControlPanelArn),
			Extract:      common.TerraformID(),
			Reference:    mg.Spec.ForProvider.ControlPanelArnRef,
			Selector:     mg.Spec.ForProvider.ControlPanelArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ControlPanelArn")
	}
	mg.Spec.ForProvider.ControlPanelArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ControlPanelArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("route53recoverycontrolconfig.aws.upbound.io", "v1beta1", "RoutingControl", "RoutingControlList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.GatingControls),
			Extract:       resource.ExtractParamPath("arn", true),
			References:    mg.Spec.ForProvider.GatingControlsRefs,
			Selector:      mg.Spec.ForProvider.GatingControlsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GatingControls")
	}
	mg.Spec.ForProvider.GatingControls = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.GatingControlsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("route53recoverycontrolconfig.aws.upbound.io", "v1beta1", "RoutingControl", "RoutingControlList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.TargetControls),
			Extract:       resource.ExtractParamPath("arn", true),
			References:    mg.Spec.ForProvider.TargetControlsRefs,
			Selector:      mg.Spec.ForProvider.TargetControlsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TargetControls")
	}
	mg.Spec.ForProvider.TargetControls = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.TargetControlsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("route53recoverycontrolconfig.aws.upbound.io", "v1beta1", "RoutingControl", "RoutingControlList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.AssertedControls),
			Extract:       common.TerraformID(),
			References:    mg.Spec.InitProvider.AssertedControlsRefs,
			Selector:      mg.Spec.InitProvider.AssertedControlsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AssertedControls")
	}
	mg.Spec.InitProvider.AssertedControls = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.AssertedControlsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("route53recoverycontrolconfig.aws.upbound.io", "v1beta1", "ControlPanel", "ControlPanelList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ControlPanelArn),
			Extract:      common.TerraformID(),
			Reference:    mg.Spec.InitProvider.ControlPanelArnRef,
			Selector:     mg.Spec.InitProvider.ControlPanelArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ControlPanelArn")
	}
	mg.Spec.InitProvider.ControlPanelArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ControlPanelArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("route53recoverycontrolconfig.aws.upbound.io", "v1beta1", "RoutingControl", "RoutingControlList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.GatingControls),
			Extract:       resource.ExtractParamPath("arn", true),
			References:    mg.Spec.InitProvider.GatingControlsRefs,
			Selector:      mg.Spec.InitProvider.GatingControlsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.GatingControls")
	}
	mg.Spec.InitProvider.GatingControls = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.GatingControlsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("route53recoverycontrolconfig.aws.upbound.io", "v1beta1", "RoutingControl", "RoutingControlList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.TargetControls),
			Extract:       resource.ExtractParamPath("arn", true),
			References:    mg.Spec.InitProvider.TargetControlsRefs,
			Selector:      mg.Spec.InitProvider.TargetControlsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TargetControls")
	}
	mg.Spec.InitProvider.TargetControls = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.TargetControlsRefs = mrsp.ResolvedReferences

	return nil
}
