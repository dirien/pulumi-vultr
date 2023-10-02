// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"errors"
	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a Vultr VPC 2.0 resource. This can be used to create, read, and delete VPCs 2.0 on your Vultr account.
//
// ## Example Usage
//
// Create a new VPC 2.0 with an automatically generated CIDR block:
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
//			_, err := vultr.NewVpc2(ctx, "myVpc2", &vultr.Vpc2Args{
//				Description: pulumi.String("my vpc2"),
//				Region:      pulumi.String("ewr"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Create a new VPC 2.0 with a specified CIDR block:
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
//			_, err := vultr.NewVpc2(ctx, "myVpc2", &vultr.Vpc2Args{
//				Description:  pulumi.String("my private vpc2"),
//				IpBlock:      pulumi.String("10.0.0.0"),
//				PrefixLength: pulumi.Int(24),
//				Region:       pulumi.String("ewr"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// VPCs 2.0 can be imported using the VPC 2.0 `ID`, e.g.
//
// ```sh
//
//	$ pulumi import vultr:index/vpc2:Vpc2 my_vpc2 0e04f918-575e-41cb-86f6-d729b354a5a1
//
// ```
type Vpc2 struct {
	pulumi.CustomResourceState

	// The date that the VPC 2.0 was added to your Vultr account.
	DateCreated pulumi.StringOutput `pulumi:"dateCreated"`
	// The description you want to give your VPC 2.0.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The IPv4 subnet to be used when attaching instances to this VPC 2.0.
	IpBlock pulumi.StringPtrOutput `pulumi:"ipBlock"`
	// Accepted values: `v4`.
	IpType pulumi.StringPtrOutput `pulumi:"ipType"`
	// The number of bits for the netmask in CIDR notation. Example: 32
	PrefixLength pulumi.IntPtrOutput `pulumi:"prefixLength"`
	// The region ID that you want the VPC 2.0 to be created in.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewVpc2 registers a new resource with the given unique name, arguments, and options.
func NewVpc2(ctx *pulumi.Context,
	name string, args *Vpc2Args, opts ...pulumi.ResourceOption) (*Vpc2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Vpc2
	err := ctx.RegisterResource("vultr:index/vpc2:Vpc2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpc2 gets an existing Vpc2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpc2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Vpc2State, opts ...pulumi.ResourceOption) (*Vpc2, error) {
	var resource Vpc2
	err := ctx.ReadResource("vultr:index/vpc2:Vpc2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vpc2 resources.
type vpc2State struct {
	// The date that the VPC 2.0 was added to your Vultr account.
	DateCreated *string `pulumi:"dateCreated"`
	// The description you want to give your VPC 2.0.
	Description *string `pulumi:"description"`
	// The IPv4 subnet to be used when attaching instances to this VPC 2.0.
	IpBlock *string `pulumi:"ipBlock"`
	// Accepted values: `v4`.
	IpType *string `pulumi:"ipType"`
	// The number of bits for the netmask in CIDR notation. Example: 32
	PrefixLength *int `pulumi:"prefixLength"`
	// The region ID that you want the VPC 2.0 to be created in.
	Region *string `pulumi:"region"`
}

type Vpc2State struct {
	// The date that the VPC 2.0 was added to your Vultr account.
	DateCreated pulumi.StringPtrInput
	// The description you want to give your VPC 2.0.
	Description pulumi.StringPtrInput
	// The IPv4 subnet to be used when attaching instances to this VPC 2.0.
	IpBlock pulumi.StringPtrInput
	// Accepted values: `v4`.
	IpType pulumi.StringPtrInput
	// The number of bits for the netmask in CIDR notation. Example: 32
	PrefixLength pulumi.IntPtrInput
	// The region ID that you want the VPC 2.0 to be created in.
	Region pulumi.StringPtrInput
}

func (Vpc2State) ElementType() reflect.Type {
	return reflect.TypeOf((*vpc2State)(nil)).Elem()
}

type vpc2Args struct {
	// The description you want to give your VPC 2.0.
	Description *string `pulumi:"description"`
	// The IPv4 subnet to be used when attaching instances to this VPC 2.0.
	IpBlock *string `pulumi:"ipBlock"`
	// Accepted values: `v4`.
	IpType *string `pulumi:"ipType"`
	// The number of bits for the netmask in CIDR notation. Example: 32
	PrefixLength *int `pulumi:"prefixLength"`
	// The region ID that you want the VPC 2.0 to be created in.
	Region string `pulumi:"region"`
}

// The set of arguments for constructing a Vpc2 resource.
type Vpc2Args struct {
	// The description you want to give your VPC 2.0.
	Description pulumi.StringPtrInput
	// The IPv4 subnet to be used when attaching instances to this VPC 2.0.
	IpBlock pulumi.StringPtrInput
	// Accepted values: `v4`.
	IpType pulumi.StringPtrInput
	// The number of bits for the netmask in CIDR notation. Example: 32
	PrefixLength pulumi.IntPtrInput
	// The region ID that you want the VPC 2.0 to be created in.
	Region pulumi.StringInput
}

func (Vpc2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*vpc2Args)(nil)).Elem()
}

type Vpc2Input interface {
	pulumi.Input

	ToVpc2Output() Vpc2Output
	ToVpc2OutputWithContext(ctx context.Context) Vpc2Output
}

func (*Vpc2) ElementType() reflect.Type {
	return reflect.TypeOf((**Vpc2)(nil)).Elem()
}

func (i *Vpc2) ToVpc2Output() Vpc2Output {
	return i.ToVpc2OutputWithContext(context.Background())
}

func (i *Vpc2) ToVpc2OutputWithContext(ctx context.Context) Vpc2Output {
	return pulumi.ToOutputWithContext(ctx, i).(Vpc2Output)
}

func (i *Vpc2) ToOutput(ctx context.Context) pulumix.Output[*Vpc2] {
	return pulumix.Output[*Vpc2]{
		OutputState: i.ToVpc2OutputWithContext(ctx).OutputState,
	}
}

// Vpc2ArrayInput is an input type that accepts Vpc2Array and Vpc2ArrayOutput values.
// You can construct a concrete instance of `Vpc2ArrayInput` via:
//
//	Vpc2Array{ Vpc2Args{...} }
type Vpc2ArrayInput interface {
	pulumi.Input

	ToVpc2ArrayOutput() Vpc2ArrayOutput
	ToVpc2ArrayOutputWithContext(context.Context) Vpc2ArrayOutput
}

type Vpc2Array []Vpc2Input

func (Vpc2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vpc2)(nil)).Elem()
}

func (i Vpc2Array) ToVpc2ArrayOutput() Vpc2ArrayOutput {
	return i.ToVpc2ArrayOutputWithContext(context.Background())
}

func (i Vpc2Array) ToVpc2ArrayOutputWithContext(ctx context.Context) Vpc2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Vpc2ArrayOutput)
}

func (i Vpc2Array) ToOutput(ctx context.Context) pulumix.Output[[]*Vpc2] {
	return pulumix.Output[[]*Vpc2]{
		OutputState: i.ToVpc2ArrayOutputWithContext(ctx).OutputState,
	}
}

// Vpc2MapInput is an input type that accepts Vpc2Map and Vpc2MapOutput values.
// You can construct a concrete instance of `Vpc2MapInput` via:
//
//	Vpc2Map{ "key": Vpc2Args{...} }
type Vpc2MapInput interface {
	pulumi.Input

	ToVpc2MapOutput() Vpc2MapOutput
	ToVpc2MapOutputWithContext(context.Context) Vpc2MapOutput
}

type Vpc2Map map[string]Vpc2Input

func (Vpc2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vpc2)(nil)).Elem()
}

func (i Vpc2Map) ToVpc2MapOutput() Vpc2MapOutput {
	return i.ToVpc2MapOutputWithContext(context.Background())
}

func (i Vpc2Map) ToVpc2MapOutputWithContext(ctx context.Context) Vpc2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Vpc2MapOutput)
}

func (i Vpc2Map) ToOutput(ctx context.Context) pulumix.Output[map[string]*Vpc2] {
	return pulumix.Output[map[string]*Vpc2]{
		OutputState: i.ToVpc2MapOutputWithContext(ctx).OutputState,
	}
}

type Vpc2Output struct{ *pulumi.OutputState }

func (Vpc2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Vpc2)(nil)).Elem()
}

func (o Vpc2Output) ToVpc2Output() Vpc2Output {
	return o
}

func (o Vpc2Output) ToVpc2OutputWithContext(ctx context.Context) Vpc2Output {
	return o
}

func (o Vpc2Output) ToOutput(ctx context.Context) pulumix.Output[*Vpc2] {
	return pulumix.Output[*Vpc2]{
		OutputState: o.OutputState,
	}
}

// The date that the VPC 2.0 was added to your Vultr account.
func (o Vpc2Output) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc2) pulumi.StringOutput { return v.DateCreated }).(pulumi.StringOutput)
}

// The description you want to give your VPC 2.0.
func (o Vpc2Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vpc2) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The IPv4 subnet to be used when attaching instances to this VPC 2.0.
func (o Vpc2Output) IpBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vpc2) pulumi.StringPtrOutput { return v.IpBlock }).(pulumi.StringPtrOutput)
}

// Accepted values: `v4`.
func (o Vpc2Output) IpType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vpc2) pulumi.StringPtrOutput { return v.IpType }).(pulumi.StringPtrOutput)
}

// The number of bits for the netmask in CIDR notation. Example: 32
func (o Vpc2Output) PrefixLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Vpc2) pulumi.IntPtrOutput { return v.PrefixLength }).(pulumi.IntPtrOutput)
}

// The region ID that you want the VPC 2.0 to be created in.
func (o Vpc2Output) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc2) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type Vpc2ArrayOutput struct{ *pulumi.OutputState }

func (Vpc2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vpc2)(nil)).Elem()
}

func (o Vpc2ArrayOutput) ToVpc2ArrayOutput() Vpc2ArrayOutput {
	return o
}

func (o Vpc2ArrayOutput) ToVpc2ArrayOutputWithContext(ctx context.Context) Vpc2ArrayOutput {
	return o
}

func (o Vpc2ArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Vpc2] {
	return pulumix.Output[[]*Vpc2]{
		OutputState: o.OutputState,
	}
}

func (o Vpc2ArrayOutput) Index(i pulumi.IntInput) Vpc2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vpc2 {
		return vs[0].([]*Vpc2)[vs[1].(int)]
	}).(Vpc2Output)
}

type Vpc2MapOutput struct{ *pulumi.OutputState }

func (Vpc2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vpc2)(nil)).Elem()
}

func (o Vpc2MapOutput) ToVpc2MapOutput() Vpc2MapOutput {
	return o
}

func (o Vpc2MapOutput) ToVpc2MapOutputWithContext(ctx context.Context) Vpc2MapOutput {
	return o
}

func (o Vpc2MapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Vpc2] {
	return pulumix.Output[map[string]*Vpc2]{
		OutputState: o.OutputState,
	}
}

func (o Vpc2MapOutput) MapIndex(k pulumi.StringInput) Vpc2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vpc2 {
		return vs[0].(map[string]*Vpc2)[vs[1].(string)]
	}).(Vpc2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Vpc2Input)(nil)).Elem(), &Vpc2{})
	pulumi.RegisterInputType(reflect.TypeOf((*Vpc2ArrayInput)(nil)).Elem(), Vpc2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*Vpc2MapInput)(nil)).Elem(), Vpc2Map{})
	pulumi.RegisterOutputType(Vpc2Output{})
	pulumi.RegisterOutputType(Vpc2ArrayOutput{})
	pulumi.RegisterOutputType(Vpc2MapOutput{})
}
