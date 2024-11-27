// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Vultr Reverse IPv6.
//
// ## Example Usage
//
// Get the information for an IPv6 reverse DNS record by `reverse`:
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
//			_, err := vultr.LookupReverseIpv6(ctx, &vultr.LookupReverseIpv6Args{
//				Filters: []vultr.GetReverseIpv6Filter{
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
func LookupReverseIpv6(ctx *pulumi.Context, args *LookupReverseIpv6Args, opts ...pulumi.InvokeOption) (*LookupReverseIpv6Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupReverseIpv6Result
	err := ctx.Invoke("vultr:index/getReverseIpv6:getReverseIpv6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getReverseIpv6.
type LookupReverseIpv6Args struct {
	// Query parameters for finding IPv6 reverse DNS records.
	Filters []GetReverseIpv6Filter `pulumi:"filters"`
}

// A collection of values returned by getReverseIpv6.
type LookupReverseIpv6Result struct {
	Filters []GetReverseIpv6Filter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ID of the instance the IPv6 reverse DNS record was set for.
	InstanceId string `pulumi:"instanceId"`
	// The IPv6 address in canonical format used in the reverse DNS record.
	Ip string `pulumi:"ip"`
	// The hostname used in the IPv6 reverse DNS record.
	Reverse string `pulumi:"reverse"`
}

func LookupReverseIpv6Output(ctx *pulumi.Context, args LookupReverseIpv6OutputArgs, opts ...pulumi.InvokeOption) LookupReverseIpv6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReverseIpv6ResultOutput, error) {
			args := v.(LookupReverseIpv6Args)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupReverseIpv6Result
			secret, err := ctx.InvokePackageRaw("vultr:index/getReverseIpv6:getReverseIpv6", args, &rv, "", opts...)
			if err != nil {
				return LookupReverseIpv6ResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupReverseIpv6ResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupReverseIpv6ResultOutput), nil
			}
			return output, nil
		}).(LookupReverseIpv6ResultOutput)
}

// A collection of arguments for invoking getReverseIpv6.
type LookupReverseIpv6OutputArgs struct {
	// Query parameters for finding IPv6 reverse DNS records.
	Filters GetReverseIpv6FilterArrayInput `pulumi:"filters"`
}

func (LookupReverseIpv6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReverseIpv6Args)(nil)).Elem()
}

// A collection of values returned by getReverseIpv6.
type LookupReverseIpv6ResultOutput struct{ *pulumi.OutputState }

func (LookupReverseIpv6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReverseIpv6Result)(nil)).Elem()
}

func (o LookupReverseIpv6ResultOutput) ToLookupReverseIpv6ResultOutput() LookupReverseIpv6ResultOutput {
	return o
}

func (o LookupReverseIpv6ResultOutput) ToLookupReverseIpv6ResultOutputWithContext(ctx context.Context) LookupReverseIpv6ResultOutput {
	return o
}

func (o LookupReverseIpv6ResultOutput) Filters() GetReverseIpv6FilterArrayOutput {
	return o.ApplyT(func(v LookupReverseIpv6Result) []GetReverseIpv6Filter { return v.Filters }).(GetReverseIpv6FilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupReverseIpv6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReverseIpv6Result) string { return v.Id }).(pulumi.StringOutput)
}

// The ID of the instance the IPv6 reverse DNS record was set for.
func (o LookupReverseIpv6ResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReverseIpv6Result) string { return v.InstanceId }).(pulumi.StringOutput)
}

// The IPv6 address in canonical format used in the reverse DNS record.
func (o LookupReverseIpv6ResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReverseIpv6Result) string { return v.Ip }).(pulumi.StringOutput)
}

// The hostname used in the IPv6 reverse DNS record.
func (o LookupReverseIpv6ResultOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReverseIpv6Result) string { return v.Reverse }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReverseIpv6ResultOutput{})
}
