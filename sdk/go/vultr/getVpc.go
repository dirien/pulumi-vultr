// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Vultr VPC.
//
// ## Example Usage
//
// Get the information for a VPC by `description`:
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
//			_, err := vultr.LookupVpc(ctx, &vultr.LookupVpcArgs{
//				Filters: []vultr.GetVpcFilter{
//					{
//						Name: "description",
//						Values: []string{
//							"my-vpc-description",
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
func LookupVpc(ctx *pulumi.Context, args *LookupVpcArgs, opts ...pulumi.InvokeOption) (*LookupVpcResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcResult
	err := ctx.Invoke("vultr:index/getVpc:getVpc", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpc.
type LookupVpcArgs struct {
	// Query parameters for finding VPCs.
	Filters []GetVpcFilter `pulumi:"filters"`
}

// A collection of values returned by getVpc.
type LookupVpcResult struct {
	// The date the VPC was added to your Vultr account.
	DateCreated string `pulumi:"dateCreated"`
	// The VPC's description.
	Description string         `pulumi:"description"`
	Filters     []GetVpcFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ID of the region that the VPC is in.
	Region string `pulumi:"region"`
	// The IPv4 network address. For example: 10.1.1.0.
	V4Subnet string `pulumi:"v4Subnet"`
	// The number of bits for the netmask in CIDR notation. Example: 20
	V4SubnetMask int `pulumi:"v4SubnetMask"`
}

func LookupVpcOutput(ctx *pulumi.Context, args LookupVpcOutputArgs, opts ...pulumi.InvokeOption) LookupVpcResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcResult, error) {
			args := v.(LookupVpcArgs)
			r, err := LookupVpc(ctx, &args, opts...)
			var s LookupVpcResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcResultOutput)
}

// A collection of arguments for invoking getVpc.
type LookupVpcOutputArgs struct {
	// Query parameters for finding VPCs.
	Filters GetVpcFilterArrayInput `pulumi:"filters"`
}

func (LookupVpcOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcArgs)(nil)).Elem()
}

// A collection of values returned by getVpc.
type LookupVpcResultOutput struct{ *pulumi.OutputState }

func (LookupVpcResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcResult)(nil)).Elem()
}

func (o LookupVpcResultOutput) ToLookupVpcResultOutput() LookupVpcResultOutput {
	return o
}

func (o LookupVpcResultOutput) ToLookupVpcResultOutputWithContext(ctx context.Context) LookupVpcResultOutput {
	return o
}

// The date the VPC was added to your Vultr account.
func (o LookupVpcResultOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.DateCreated }).(pulumi.StringOutput)
}

// The VPC's description.
func (o LookupVpcResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupVpcResultOutput) Filters() GetVpcFilterArrayOutput {
	return o.ApplyT(func(v LookupVpcResult) []GetVpcFilter { return v.Filters }).(GetVpcFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVpcResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the region that the VPC is in.
func (o LookupVpcResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.Region }).(pulumi.StringOutput)
}

// The IPv4 network address. For example: 10.1.1.0.
func (o LookupVpcResultOutput) V4Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.V4Subnet }).(pulumi.StringOutput)
}

// The number of bits for the netmask in CIDR notation. Example: 20
func (o LookupVpcResultOutput) V4SubnetMask() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpcResult) int { return v.V4SubnetMask }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcResultOutput{})
}
