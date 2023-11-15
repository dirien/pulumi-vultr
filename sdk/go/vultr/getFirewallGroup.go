// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a firewall group on your Vultr account.
//
// ## Example Usage
//
// Get the information for a firewall group by `description`:
//
// ```go
// package main
//
// import (
//
//	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vultr.LookupFirewallGroup(ctx, &vultr.LookupFirewallGroupArgs{
//				Filters: []vultr.GetFirewallGroupFilter{
//					{
//						Name: "description",
//						Values: []string{
//							"fwg-description",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFirewallGroup(ctx *pulumi.Context, args *LookupFirewallGroupArgs, opts ...pulumi.InvokeOption) (*LookupFirewallGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFirewallGroupResult
	err := ctx.Invoke("vultr:index/getFirewallGroup:getFirewallGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewallGroup.
type LookupFirewallGroupArgs struct {
	// Query parameters for finding firewall groups.
	Filters []GetFirewallGroupFilter `pulumi:"filters"`
}

// A collection of values returned by getFirewallGroup.
type LookupFirewallGroupResult struct {
	// The date the firewall group was added to your Vultr account.
	DateCreated string `pulumi:"dateCreated"`
	// The date the firewall group was last modified.
	DateModified string `pulumi:"dateModified"`
	// The description of the firewall group.
	Description string                   `pulumi:"description"`
	Filters     []GetFirewallGroupFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The number of instances this firewall group is applied to.
	InstanceCount int `pulumi:"instanceCount"`
	// The maximum number of rules this firewall group can have.
	MaxRuleCount int `pulumi:"maxRuleCount"`
	// The number of rules added to this firewall group.
	RuleCount int `pulumi:"ruleCount"`
}

func LookupFirewallGroupOutput(ctx *pulumi.Context, args LookupFirewallGroupOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallGroupResult, error) {
			args := v.(LookupFirewallGroupArgs)
			r, err := LookupFirewallGroup(ctx, &args, opts...)
			var s LookupFirewallGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallGroupResultOutput)
}

// A collection of arguments for invoking getFirewallGroup.
type LookupFirewallGroupOutputArgs struct {
	// Query parameters for finding firewall groups.
	Filters GetFirewallGroupFilterArrayInput `pulumi:"filters"`
}

func (LookupFirewallGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallGroupArgs)(nil)).Elem()
}

// A collection of values returned by getFirewallGroup.
type LookupFirewallGroupResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallGroupResult)(nil)).Elem()
}

func (o LookupFirewallGroupResultOutput) ToLookupFirewallGroupResultOutput() LookupFirewallGroupResultOutput {
	return o
}

func (o LookupFirewallGroupResultOutput) ToLookupFirewallGroupResultOutputWithContext(ctx context.Context) LookupFirewallGroupResultOutput {
	return o
}

// The date the firewall group was added to your Vultr account.
func (o LookupFirewallGroupResultOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallGroupResult) string { return v.DateCreated }).(pulumi.StringOutput)
}

// The date the firewall group was last modified.
func (o LookupFirewallGroupResultOutput) DateModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallGroupResult) string { return v.DateModified }).(pulumi.StringOutput)
}

// The description of the firewall group.
func (o LookupFirewallGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupFirewallGroupResultOutput) Filters() GetFirewallGroupFilterArrayOutput {
	return o.ApplyT(func(v LookupFirewallGroupResult) []GetFirewallGroupFilter { return v.Filters }).(GetFirewallGroupFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFirewallGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The number of instances this firewall group is applied to.
func (o LookupFirewallGroupResultOutput) InstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallGroupResult) int { return v.InstanceCount }).(pulumi.IntOutput)
}

// The maximum number of rules this firewall group can have.
func (o LookupFirewallGroupResultOutput) MaxRuleCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallGroupResult) int { return v.MaxRuleCount }).(pulumi.IntOutput)
}

// The number of rules added to this firewall group.
func (o LookupFirewallGroupResultOutput) RuleCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallGroupResult) int { return v.RuleCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallGroupResultOutput{})
}
