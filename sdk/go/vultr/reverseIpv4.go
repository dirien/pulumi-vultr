// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"errors"
	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Vultr Reverse IPv4 resource. This can be used to create, read, and
// modify reverse DNS records for IPv4 addresses. Upon success, DNS
// changes may take 6-12 hours to become active.
//
// ## Example Usage
//
// Create a new reverse DNS record for an IPv4 address:
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
//			myInstance, err := vultr.NewInstance(ctx, "myInstance", &vultr.InstanceArgs{
//				EnableIpv6: pulumi.Bool(true),
//				OsId:       pulumi.Int(477),
//				Plan:       pulumi.String("vc2-1c-2gb"),
//				Region:     pulumi.String("ewr"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vultr.NewReverseIpv4(ctx, "myReverseIpv4", &vultr.ReverseIpv4Args{
//				InstanceId: myInstance.ID(),
//				Ip:         myInstance.MainIp,
//				Reverse:    pulumi.String("host.example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ReverseIpv4 struct {
	pulumi.CustomResourceState

	// The gateway IP address.
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// The ID of the instance you want to set an IPv4
	// reverse DNS record for.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The IPv4 address used in the reverse DNS record.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// The IPv4 netmask in dot-decimal notation.
	Netmask pulumi.StringOutput `pulumi:"netmask"`
	// The hostname used in the IPv4 reverse DNS record.
	Reverse pulumi.StringOutput `pulumi:"reverse"`
}

// NewReverseIpv4 registers a new resource with the given unique name, arguments, and options.
func NewReverseIpv4(ctx *pulumi.Context,
	name string, args *ReverseIpv4Args, opts ...pulumi.ResourceOption) (*ReverseIpv4, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	if args.Reverse == nil {
		return nil, errors.New("invalid value for required argument 'Reverse'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ReverseIpv4
	err := ctx.RegisterResource("vultr:index/reverseIpv4:ReverseIpv4", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReverseIpv4 gets an existing ReverseIpv4 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReverseIpv4(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReverseIpv4State, opts ...pulumi.ResourceOption) (*ReverseIpv4, error) {
	var resource ReverseIpv4
	err := ctx.ReadResource("vultr:index/reverseIpv4:ReverseIpv4", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReverseIpv4 resources.
type reverseIpv4State struct {
	// The gateway IP address.
	Gateway *string `pulumi:"gateway"`
	// The ID of the instance you want to set an IPv4
	// reverse DNS record for.
	InstanceId *string `pulumi:"instanceId"`
	// The IPv4 address used in the reverse DNS record.
	Ip *string `pulumi:"ip"`
	// The IPv4 netmask in dot-decimal notation.
	Netmask *string `pulumi:"netmask"`
	// The hostname used in the IPv4 reverse DNS record.
	Reverse *string `pulumi:"reverse"`
}

type ReverseIpv4State struct {
	// The gateway IP address.
	Gateway pulumi.StringPtrInput
	// The ID of the instance you want to set an IPv4
	// reverse DNS record for.
	InstanceId pulumi.StringPtrInput
	// The IPv4 address used in the reverse DNS record.
	Ip pulumi.StringPtrInput
	// The IPv4 netmask in dot-decimal notation.
	Netmask pulumi.StringPtrInput
	// The hostname used in the IPv4 reverse DNS record.
	Reverse pulumi.StringPtrInput
}

func (ReverseIpv4State) ElementType() reflect.Type {
	return reflect.TypeOf((*reverseIpv4State)(nil)).Elem()
}

type reverseIpv4Args struct {
	// The ID of the instance you want to set an IPv4
	// reverse DNS record for.
	InstanceId string `pulumi:"instanceId"`
	// The IPv4 address used in the reverse DNS record.
	Ip string `pulumi:"ip"`
	// The hostname used in the IPv4 reverse DNS record.
	Reverse string `pulumi:"reverse"`
}

// The set of arguments for constructing a ReverseIpv4 resource.
type ReverseIpv4Args struct {
	// The ID of the instance you want to set an IPv4
	// reverse DNS record for.
	InstanceId pulumi.StringInput
	// The IPv4 address used in the reverse DNS record.
	Ip pulumi.StringInput
	// The hostname used in the IPv4 reverse DNS record.
	Reverse pulumi.StringInput
}

func (ReverseIpv4Args) ElementType() reflect.Type {
	return reflect.TypeOf((*reverseIpv4Args)(nil)).Elem()
}

type ReverseIpv4Input interface {
	pulumi.Input

	ToReverseIpv4Output() ReverseIpv4Output
	ToReverseIpv4OutputWithContext(ctx context.Context) ReverseIpv4Output
}

func (*ReverseIpv4) ElementType() reflect.Type {
	return reflect.TypeOf((**ReverseIpv4)(nil)).Elem()
}

func (i *ReverseIpv4) ToReverseIpv4Output() ReverseIpv4Output {
	return i.ToReverseIpv4OutputWithContext(context.Background())
}

func (i *ReverseIpv4) ToReverseIpv4OutputWithContext(ctx context.Context) ReverseIpv4Output {
	return pulumi.ToOutputWithContext(ctx, i).(ReverseIpv4Output)
}

// ReverseIpv4ArrayInput is an input type that accepts ReverseIpv4Array and ReverseIpv4ArrayOutput values.
// You can construct a concrete instance of `ReverseIpv4ArrayInput` via:
//
//	ReverseIpv4Array{ ReverseIpv4Args{...} }
type ReverseIpv4ArrayInput interface {
	pulumi.Input

	ToReverseIpv4ArrayOutput() ReverseIpv4ArrayOutput
	ToReverseIpv4ArrayOutputWithContext(context.Context) ReverseIpv4ArrayOutput
}

type ReverseIpv4Array []ReverseIpv4Input

func (ReverseIpv4Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReverseIpv4)(nil)).Elem()
}

func (i ReverseIpv4Array) ToReverseIpv4ArrayOutput() ReverseIpv4ArrayOutput {
	return i.ToReverseIpv4ArrayOutputWithContext(context.Background())
}

func (i ReverseIpv4Array) ToReverseIpv4ArrayOutputWithContext(ctx context.Context) ReverseIpv4ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReverseIpv4ArrayOutput)
}

// ReverseIpv4MapInput is an input type that accepts ReverseIpv4Map and ReverseIpv4MapOutput values.
// You can construct a concrete instance of `ReverseIpv4MapInput` via:
//
//	ReverseIpv4Map{ "key": ReverseIpv4Args{...} }
type ReverseIpv4MapInput interface {
	pulumi.Input

	ToReverseIpv4MapOutput() ReverseIpv4MapOutput
	ToReverseIpv4MapOutputWithContext(context.Context) ReverseIpv4MapOutput
}

type ReverseIpv4Map map[string]ReverseIpv4Input

func (ReverseIpv4Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReverseIpv4)(nil)).Elem()
}

func (i ReverseIpv4Map) ToReverseIpv4MapOutput() ReverseIpv4MapOutput {
	return i.ToReverseIpv4MapOutputWithContext(context.Background())
}

func (i ReverseIpv4Map) ToReverseIpv4MapOutputWithContext(ctx context.Context) ReverseIpv4MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReverseIpv4MapOutput)
}

type ReverseIpv4Output struct{ *pulumi.OutputState }

func (ReverseIpv4Output) ElementType() reflect.Type {
	return reflect.TypeOf((**ReverseIpv4)(nil)).Elem()
}

func (o ReverseIpv4Output) ToReverseIpv4Output() ReverseIpv4Output {
	return o
}

func (o ReverseIpv4Output) ToReverseIpv4OutputWithContext(ctx context.Context) ReverseIpv4Output {
	return o
}

// The gateway IP address.
func (o ReverseIpv4Output) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *ReverseIpv4) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// The ID of the instance you want to set an IPv4
// reverse DNS record for.
func (o ReverseIpv4Output) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReverseIpv4) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The IPv4 address used in the reverse DNS record.
func (o ReverseIpv4Output) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *ReverseIpv4) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// The IPv4 netmask in dot-decimal notation.
func (o ReverseIpv4Output) Netmask() pulumi.StringOutput {
	return o.ApplyT(func(v *ReverseIpv4) pulumi.StringOutput { return v.Netmask }).(pulumi.StringOutput)
}

// The hostname used in the IPv4 reverse DNS record.
func (o ReverseIpv4Output) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v *ReverseIpv4) pulumi.StringOutput { return v.Reverse }).(pulumi.StringOutput)
}

type ReverseIpv4ArrayOutput struct{ *pulumi.OutputState }

func (ReverseIpv4ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReverseIpv4)(nil)).Elem()
}

func (o ReverseIpv4ArrayOutput) ToReverseIpv4ArrayOutput() ReverseIpv4ArrayOutput {
	return o
}

func (o ReverseIpv4ArrayOutput) ToReverseIpv4ArrayOutputWithContext(ctx context.Context) ReverseIpv4ArrayOutput {
	return o
}

func (o ReverseIpv4ArrayOutput) Index(i pulumi.IntInput) ReverseIpv4Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReverseIpv4 {
		return vs[0].([]*ReverseIpv4)[vs[1].(int)]
	}).(ReverseIpv4Output)
}

type ReverseIpv4MapOutput struct{ *pulumi.OutputState }

func (ReverseIpv4MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReverseIpv4)(nil)).Elem()
}

func (o ReverseIpv4MapOutput) ToReverseIpv4MapOutput() ReverseIpv4MapOutput {
	return o
}

func (o ReverseIpv4MapOutput) ToReverseIpv4MapOutputWithContext(ctx context.Context) ReverseIpv4MapOutput {
	return o
}

func (o ReverseIpv4MapOutput) MapIndex(k pulumi.StringInput) ReverseIpv4Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReverseIpv4 {
		return vs[0].(map[string]*ReverseIpv4)[vs[1].(string)]
	}).(ReverseIpv4Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReverseIpv4Input)(nil)).Elem(), &ReverseIpv4{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReverseIpv4ArrayInput)(nil)).Elem(), ReverseIpv4Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReverseIpv4MapInput)(nil)).Elem(), ReverseIpv4Map{})
	pulumi.RegisterOutputType(ReverseIpv4Output{})
	pulumi.RegisterOutputType(ReverseIpv4ArrayOutput{})
	pulumi.RegisterOutputType(ReverseIpv4MapOutput{})
}
