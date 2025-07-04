// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"errors"
	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Vultr VPC resource. This can be used to create, read, and delete VPCs on your Vultr account.
//
// ## Example Usage
//
// Create a new VPC with an automatically generated CIDR block:
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
//			_, err := vultr.NewVpc(ctx, "myVpc", &vultr.VpcArgs{
//				Description: pulumi.String("my vpc"),
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
// Create a new VPC with a specified CIDR block:
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
//			_, err := vultr.NewVpc(ctx, "myVpc", &vultr.VpcArgs{
//				Description:  pulumi.String("my private vpc"),
//				Region:       pulumi.String("ewr"),
//				V4Subnet:     pulumi.String("10.0.0.0"),
//				V4SubnetMask: pulumi.Int(24),
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
// VPCs can be imported using the VPC `ID`, e.g.
//
// ```sh
// $ pulumi import vultr:index/vpc:Vpc my_vpc 0e04f918-575e-41cb-86f6-d729b354a5a1
// ```
type Vpc struct {
	pulumi.CustomResourceState

	// The date that the VPC was added to your Vultr account.
	DateCreated pulumi.StringOutput `pulumi:"dateCreated"`
	// The description you want to give your VPC.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The region ID that you want the VPC to be created in.
	Region pulumi.StringOutput `pulumi:"region"`
	// The IPv4 subnet to be used when attaching instances to this VPC.
	V4Subnet pulumi.StringOutput `pulumi:"v4Subnet"`
	// The number of bits for the netmask in CIDR notation. Example: 32
	V4SubnetMask pulumi.IntOutput `pulumi:"v4SubnetMask"`
}

// NewVpc registers a new resource with the given unique name, arguments, and options.
func NewVpc(ctx *pulumi.Context,
	name string, args *VpcArgs, opts ...pulumi.ResourceOption) (*Vpc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Vpc
	err := ctx.RegisterResource("vultr:index/vpc:Vpc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpc gets an existing Vpc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcState, opts ...pulumi.ResourceOption) (*Vpc, error) {
	var resource Vpc
	err := ctx.ReadResource("vultr:index/vpc:Vpc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vpc resources.
type vpcState struct {
	// The date that the VPC was added to your Vultr account.
	DateCreated *string `pulumi:"dateCreated"`
	// The description you want to give your VPC.
	Description *string `pulumi:"description"`
	// The region ID that you want the VPC to be created in.
	Region *string `pulumi:"region"`
	// The IPv4 subnet to be used when attaching instances to this VPC.
	V4Subnet *string `pulumi:"v4Subnet"`
	// The number of bits for the netmask in CIDR notation. Example: 32
	V4SubnetMask *int `pulumi:"v4SubnetMask"`
}

type VpcState struct {
	// The date that the VPC was added to your Vultr account.
	DateCreated pulumi.StringPtrInput
	// The description you want to give your VPC.
	Description pulumi.StringPtrInput
	// The region ID that you want the VPC to be created in.
	Region pulumi.StringPtrInput
	// The IPv4 subnet to be used when attaching instances to this VPC.
	V4Subnet pulumi.StringPtrInput
	// The number of bits for the netmask in CIDR notation. Example: 32
	V4SubnetMask pulumi.IntPtrInput
}

func (VpcState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcState)(nil)).Elem()
}

type vpcArgs struct {
	// The description you want to give your VPC.
	Description *string `pulumi:"description"`
	// The region ID that you want the VPC to be created in.
	Region string `pulumi:"region"`
	// The IPv4 subnet to be used when attaching instances to this VPC.
	V4Subnet *string `pulumi:"v4Subnet"`
	// The number of bits for the netmask in CIDR notation. Example: 32
	V4SubnetMask *int `pulumi:"v4SubnetMask"`
}

// The set of arguments for constructing a Vpc resource.
type VpcArgs struct {
	// The description you want to give your VPC.
	Description pulumi.StringPtrInput
	// The region ID that you want the VPC to be created in.
	Region pulumi.StringInput
	// The IPv4 subnet to be used when attaching instances to this VPC.
	V4Subnet pulumi.StringPtrInput
	// The number of bits for the netmask in CIDR notation. Example: 32
	V4SubnetMask pulumi.IntPtrInput
}

func (VpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcArgs)(nil)).Elem()
}

type VpcInput interface {
	pulumi.Input

	ToVpcOutput() VpcOutput
	ToVpcOutputWithContext(ctx context.Context) VpcOutput
}

func (*Vpc) ElementType() reflect.Type {
	return reflect.TypeOf((**Vpc)(nil)).Elem()
}

func (i *Vpc) ToVpcOutput() VpcOutput {
	return i.ToVpcOutputWithContext(context.Background())
}

func (i *Vpc) ToVpcOutputWithContext(ctx context.Context) VpcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcOutput)
}

// VpcArrayInput is an input type that accepts VpcArray and VpcArrayOutput values.
// You can construct a concrete instance of `VpcArrayInput` via:
//
//	VpcArray{ VpcArgs{...} }
type VpcArrayInput interface {
	pulumi.Input

	ToVpcArrayOutput() VpcArrayOutput
	ToVpcArrayOutputWithContext(context.Context) VpcArrayOutput
}

type VpcArray []VpcInput

func (VpcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vpc)(nil)).Elem()
}

func (i VpcArray) ToVpcArrayOutput() VpcArrayOutput {
	return i.ToVpcArrayOutputWithContext(context.Background())
}

func (i VpcArray) ToVpcArrayOutputWithContext(ctx context.Context) VpcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcArrayOutput)
}

// VpcMapInput is an input type that accepts VpcMap and VpcMapOutput values.
// You can construct a concrete instance of `VpcMapInput` via:
//
//	VpcMap{ "key": VpcArgs{...} }
type VpcMapInput interface {
	pulumi.Input

	ToVpcMapOutput() VpcMapOutput
	ToVpcMapOutputWithContext(context.Context) VpcMapOutput
}

type VpcMap map[string]VpcInput

func (VpcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vpc)(nil)).Elem()
}

func (i VpcMap) ToVpcMapOutput() VpcMapOutput {
	return i.ToVpcMapOutputWithContext(context.Background())
}

func (i VpcMap) ToVpcMapOutputWithContext(ctx context.Context) VpcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcMapOutput)
}

type VpcOutput struct{ *pulumi.OutputState }

func (VpcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vpc)(nil)).Elem()
}

func (o VpcOutput) ToVpcOutput() VpcOutput {
	return o
}

func (o VpcOutput) ToVpcOutputWithContext(ctx context.Context) VpcOutput {
	return o
}

// The date that the VPC was added to your Vultr account.
func (o VpcOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.DateCreated }).(pulumi.StringOutput)
}

// The description you want to give your VPC.
func (o VpcOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The region ID that you want the VPC to be created in.
func (o VpcOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The IPv4 subnet to be used when attaching instances to this VPC.
func (o VpcOutput) V4Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.V4Subnet }).(pulumi.StringOutput)
}

// The number of bits for the netmask in CIDR notation. Example: 32
func (o VpcOutput) V4SubnetMask() pulumi.IntOutput {
	return o.ApplyT(func(v *Vpc) pulumi.IntOutput { return v.V4SubnetMask }).(pulumi.IntOutput)
}

type VpcArrayOutput struct{ *pulumi.OutputState }

func (VpcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vpc)(nil)).Elem()
}

func (o VpcArrayOutput) ToVpcArrayOutput() VpcArrayOutput {
	return o
}

func (o VpcArrayOutput) ToVpcArrayOutputWithContext(ctx context.Context) VpcArrayOutput {
	return o
}

func (o VpcArrayOutput) Index(i pulumi.IntInput) VpcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vpc {
		return vs[0].([]*Vpc)[vs[1].(int)]
	}).(VpcOutput)
}

type VpcMapOutput struct{ *pulumi.OutputState }

func (VpcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vpc)(nil)).Elem()
}

func (o VpcMapOutput) ToVpcMapOutput() VpcMapOutput {
	return o
}

func (o VpcMapOutput) ToVpcMapOutputWithContext(ctx context.Context) VpcMapOutput {
	return o
}

func (o VpcMapOutput) MapIndex(k pulumi.StringInput) VpcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vpc {
		return vs[0].(map[string]*Vpc)[vs[1].(string)]
	}).(VpcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcInput)(nil)).Elem(), &Vpc{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcArrayInput)(nil)).Elem(), VpcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcMapInput)(nil)).Elem(), VpcMap{})
	pulumi.RegisterOutputType(VpcOutput{})
	pulumi.RegisterOutputType(VpcArrayOutput{})
	pulumi.RegisterOutputType(VpcMapOutput{})
}
