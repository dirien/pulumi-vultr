// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get information about a Vultr Reverse IPv4.
//
// ## Example Usage
//
// Get the information for an IPv4 reverse DNS record by `reverse`:
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
//			_, err := vultr.LookupReverseIpv4(ctx, &vultr.LookupReverseIpv4Args{
//				Filters: []vultr.GetReverseIpv4Filter{
//					{
//						Name: "reverse",
//						Values: []string{
//							"host.example.com",
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
func LookupReverseIpv4(ctx *pulumi.Context, args *LookupReverseIpv4Args, opts ...pulumi.InvokeOption) (*LookupReverseIpv4Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupReverseIpv4Result
	err := ctx.Invoke("vultr:index/getReverseIpv4:getReverseIpv4", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getReverseIpv4.
type LookupReverseIpv4Args struct {
	// Query parameters for finding IPv4 reverse DNS records.
	Filters []GetReverseIpv4Filter `pulumi:"filters"`
}

// A collection of values returned by getReverseIpv4.
type LookupReverseIpv4Result struct {
	Filters []GetReverseIpv4Filter `pulumi:"filters"`
	// The gateway IP address.
	Gateway string `pulumi:"gateway"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ID of the instance the IPv4 reverse DNS record was set for.
	InstanceId string `pulumi:"instanceId"`
	// The IPv4 address in canonical format used in the reverse DNS record.
	Ip string `pulumi:"ip"`
	// The IPv4 netmask in dot-decimal notation.
	Netmask string `pulumi:"netmask"`
	// The hostname used in the IPv4 reverse DNS record.
	Reverse string `pulumi:"reverse"`
}

func LookupReverseIpv4Output(ctx *pulumi.Context, args LookupReverseIpv4OutputArgs, opts ...pulumi.InvokeOption) LookupReverseIpv4ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReverseIpv4Result, error) {
			args := v.(LookupReverseIpv4Args)
			r, err := LookupReverseIpv4(ctx, &args, opts...)
			var s LookupReverseIpv4Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReverseIpv4ResultOutput)
}

// A collection of arguments for invoking getReverseIpv4.
type LookupReverseIpv4OutputArgs struct {
	// Query parameters for finding IPv4 reverse DNS records.
	Filters GetReverseIpv4FilterArrayInput `pulumi:"filters"`
}

func (LookupReverseIpv4OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReverseIpv4Args)(nil)).Elem()
}

// A collection of values returned by getReverseIpv4.
type LookupReverseIpv4ResultOutput struct{ *pulumi.OutputState }

func (LookupReverseIpv4ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReverseIpv4Result)(nil)).Elem()
}

func (o LookupReverseIpv4ResultOutput) ToLookupReverseIpv4ResultOutput() LookupReverseIpv4ResultOutput {
	return o
}

func (o LookupReverseIpv4ResultOutput) ToLookupReverseIpv4ResultOutputWithContext(ctx context.Context) LookupReverseIpv4ResultOutput {
	return o
}

func (o LookupReverseIpv4ResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupReverseIpv4Result] {
	return pulumix.Output[LookupReverseIpv4Result]{
		OutputState: o.OutputState,
	}
}

func (o LookupReverseIpv4ResultOutput) Filters() GetReverseIpv4FilterArrayOutput {
	return o.ApplyT(func(v LookupReverseIpv4Result) []GetReverseIpv4Filter { return v.Filters }).(GetReverseIpv4FilterArrayOutput)
}

// The gateway IP address.
func (o LookupReverseIpv4ResultOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReverseIpv4Result) string { return v.Gateway }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupReverseIpv4ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReverseIpv4Result) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the instance the IPv4 reverse DNS record was set for.
func (o LookupReverseIpv4ResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReverseIpv4Result) string { return v.InstanceId }).(pulumi.StringOutput)
}

// The IPv4 address in canonical format used in the reverse DNS record.
func (o LookupReverseIpv4ResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReverseIpv4Result) string { return v.Ip }).(pulumi.StringOutput)
}

// The IPv4 netmask in dot-decimal notation.
func (o LookupReverseIpv4ResultOutput) Netmask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReverseIpv4Result) string { return v.Netmask }).(pulumi.StringOutput)
}

// The hostname used in the IPv4 reverse DNS record.
func (o LookupReverseIpv4ResultOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReverseIpv4Result) string { return v.Reverse }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReverseIpv4ResultOutput{})
}
