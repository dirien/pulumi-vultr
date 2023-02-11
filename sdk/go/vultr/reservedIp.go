// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Vultr reserved IP resource. This can be used to create, read, modify, and delete reserved IP addresses on your Vultr account.
//
// ## Example Usage
//
// Create a new reserved IP:
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
//			_, err := vultr.NewReservedIp(ctx, "myReservedIp", &vultr.ReservedIpArgs{
//				IpType: pulumi.String("v4"),
//				Label:  pulumi.String("my-reserved-ip"),
//				Region: pulumi.String("sea"),
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
// Attach a reserved IP to a instance:
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
//			_, err := vultr.NewReservedIp(ctx, "myReservedIp", &vultr.ReservedIpArgs{
//				InstanceId: pulumi.String("b9cc6fad-70b1-40ee-ab6a-4d622858962f"),
//				IpType:     pulumi.String("v4"),
//				Label:      pulumi.String("my-reserved-ip"),
//				Region:     pulumi.String("sea"),
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
// Reserved IPs can be imported using the reserved IP `ID`, e.g.
//
// ```sh
//
//	$ pulumi import vultr:index/reservedIp:ReservedIp my_reserved_ip b9cc6fad-70b1-40ee-ab6a-4d622858962f
//
// ```
type ReservedIp struct {
	pulumi.CustomResourceState

	// The VPS ID you want this reserved IP to be attached to.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The type of reserved IP that you want. Either "v4" or "v6".
	IpType pulumi.StringOutput `pulumi:"ipType"`
	// The label you want to give your reserved IP.
	Label pulumi.StringPtrOutput `pulumi:"label"`
	// The region ID that you want the reserved IP to be created in.
	Region pulumi.StringOutput `pulumi:"region"`
	// The reserved IP's subnet.
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// The reserved IP's subnet size.
	SubnetSize pulumi.IntOutput `pulumi:"subnetSize"`
}

// NewReservedIp registers a new resource with the given unique name, arguments, and options.
func NewReservedIp(ctx *pulumi.Context,
	name string, args *ReservedIpArgs, opts ...pulumi.ResourceOption) (*ReservedIp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpType == nil {
		return nil, errors.New("invalid value for required argument 'IpType'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ReservedIp
	err := ctx.RegisterResource("vultr:index/reservedIp:ReservedIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReservedIp gets an existing ReservedIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReservedIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReservedIpState, opts ...pulumi.ResourceOption) (*ReservedIp, error) {
	var resource ReservedIp
	err := ctx.ReadResource("vultr:index/reservedIp:ReservedIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReservedIp resources.
type reservedIpState struct {
	// The VPS ID you want this reserved IP to be attached to.
	InstanceId *string `pulumi:"instanceId"`
	// The type of reserved IP that you want. Either "v4" or "v6".
	IpType *string `pulumi:"ipType"`
	// The label you want to give your reserved IP.
	Label *string `pulumi:"label"`
	// The region ID that you want the reserved IP to be created in.
	Region *string `pulumi:"region"`
	// The reserved IP's subnet.
	Subnet *string `pulumi:"subnet"`
	// The reserved IP's subnet size.
	SubnetSize *int `pulumi:"subnetSize"`
}

type ReservedIpState struct {
	// The VPS ID you want this reserved IP to be attached to.
	InstanceId pulumi.StringPtrInput
	// The type of reserved IP that you want. Either "v4" or "v6".
	IpType pulumi.StringPtrInput
	// The label you want to give your reserved IP.
	Label pulumi.StringPtrInput
	// The region ID that you want the reserved IP to be created in.
	Region pulumi.StringPtrInput
	// The reserved IP's subnet.
	Subnet pulumi.StringPtrInput
	// The reserved IP's subnet size.
	SubnetSize pulumi.IntPtrInput
}

func (ReservedIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*reservedIpState)(nil)).Elem()
}

type reservedIpArgs struct {
	// The VPS ID you want this reserved IP to be attached to.
	InstanceId *string `pulumi:"instanceId"`
	// The type of reserved IP that you want. Either "v4" or "v6".
	IpType string `pulumi:"ipType"`
	// The label you want to give your reserved IP.
	Label *string `pulumi:"label"`
	// The region ID that you want the reserved IP to be created in.
	Region string `pulumi:"region"`
}

// The set of arguments for constructing a ReservedIp resource.
type ReservedIpArgs struct {
	// The VPS ID you want this reserved IP to be attached to.
	InstanceId pulumi.StringPtrInput
	// The type of reserved IP that you want. Either "v4" or "v6".
	IpType pulumi.StringInput
	// The label you want to give your reserved IP.
	Label pulumi.StringPtrInput
	// The region ID that you want the reserved IP to be created in.
	Region pulumi.StringInput
}

func (ReservedIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reservedIpArgs)(nil)).Elem()
}

type ReservedIpInput interface {
	pulumi.Input

	ToReservedIpOutput() ReservedIpOutput
	ToReservedIpOutputWithContext(ctx context.Context) ReservedIpOutput
}

func (*ReservedIp) ElementType() reflect.Type {
	return reflect.TypeOf((**ReservedIp)(nil)).Elem()
}

func (i *ReservedIp) ToReservedIpOutput() ReservedIpOutput {
	return i.ToReservedIpOutputWithContext(context.Background())
}

func (i *ReservedIp) ToReservedIpOutputWithContext(ctx context.Context) ReservedIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservedIpOutput)
}

// ReservedIpArrayInput is an input type that accepts ReservedIpArray and ReservedIpArrayOutput values.
// You can construct a concrete instance of `ReservedIpArrayInput` via:
//
//	ReservedIpArray{ ReservedIpArgs{...} }
type ReservedIpArrayInput interface {
	pulumi.Input

	ToReservedIpArrayOutput() ReservedIpArrayOutput
	ToReservedIpArrayOutputWithContext(context.Context) ReservedIpArrayOutput
}

type ReservedIpArray []ReservedIpInput

func (ReservedIpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReservedIp)(nil)).Elem()
}

func (i ReservedIpArray) ToReservedIpArrayOutput() ReservedIpArrayOutput {
	return i.ToReservedIpArrayOutputWithContext(context.Background())
}

func (i ReservedIpArray) ToReservedIpArrayOutputWithContext(ctx context.Context) ReservedIpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservedIpArrayOutput)
}

// ReservedIpMapInput is an input type that accepts ReservedIpMap and ReservedIpMapOutput values.
// You can construct a concrete instance of `ReservedIpMapInput` via:
//
//	ReservedIpMap{ "key": ReservedIpArgs{...} }
type ReservedIpMapInput interface {
	pulumi.Input

	ToReservedIpMapOutput() ReservedIpMapOutput
	ToReservedIpMapOutputWithContext(context.Context) ReservedIpMapOutput
}

type ReservedIpMap map[string]ReservedIpInput

func (ReservedIpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReservedIp)(nil)).Elem()
}

func (i ReservedIpMap) ToReservedIpMapOutput() ReservedIpMapOutput {
	return i.ToReservedIpMapOutputWithContext(context.Background())
}

func (i ReservedIpMap) ToReservedIpMapOutputWithContext(ctx context.Context) ReservedIpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReservedIpMapOutput)
}

type ReservedIpOutput struct{ *pulumi.OutputState }

func (ReservedIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReservedIp)(nil)).Elem()
}

func (o ReservedIpOutput) ToReservedIpOutput() ReservedIpOutput {
	return o
}

func (o ReservedIpOutput) ToReservedIpOutputWithContext(ctx context.Context) ReservedIpOutput {
	return o
}

// The VPS ID you want this reserved IP to be attached to.
func (o ReservedIpOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReservedIp) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The type of reserved IP that you want. Either "v4" or "v6".
func (o ReservedIpOutput) IpType() pulumi.StringOutput {
	return o.ApplyT(func(v *ReservedIp) pulumi.StringOutput { return v.IpType }).(pulumi.StringOutput)
}

// The label you want to give your reserved IP.
func (o ReservedIpOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReservedIp) pulumi.StringPtrOutput { return v.Label }).(pulumi.StringPtrOutput)
}

// The region ID that you want the reserved IP to be created in.
func (o ReservedIpOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ReservedIp) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The reserved IP's subnet.
func (o ReservedIpOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *ReservedIp) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

// The reserved IP's subnet size.
func (o ReservedIpOutput) SubnetSize() pulumi.IntOutput {
	return o.ApplyT(func(v *ReservedIp) pulumi.IntOutput { return v.SubnetSize }).(pulumi.IntOutput)
}

type ReservedIpArrayOutput struct{ *pulumi.OutputState }

func (ReservedIpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReservedIp)(nil)).Elem()
}

func (o ReservedIpArrayOutput) ToReservedIpArrayOutput() ReservedIpArrayOutput {
	return o
}

func (o ReservedIpArrayOutput) ToReservedIpArrayOutputWithContext(ctx context.Context) ReservedIpArrayOutput {
	return o
}

func (o ReservedIpArrayOutput) Index(i pulumi.IntInput) ReservedIpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReservedIp {
		return vs[0].([]*ReservedIp)[vs[1].(int)]
	}).(ReservedIpOutput)
}

type ReservedIpMapOutput struct{ *pulumi.OutputState }

func (ReservedIpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReservedIp)(nil)).Elem()
}

func (o ReservedIpMapOutput) ToReservedIpMapOutput() ReservedIpMapOutput {
	return o
}

func (o ReservedIpMapOutput) ToReservedIpMapOutputWithContext(ctx context.Context) ReservedIpMapOutput {
	return o
}

func (o ReservedIpMapOutput) MapIndex(k pulumi.StringInput) ReservedIpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReservedIp {
		return vs[0].(map[string]*ReservedIp)[vs[1].(string)]
	}).(ReservedIpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReservedIpInput)(nil)).Elem(), &ReservedIp{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReservedIpArrayInput)(nil)).Elem(), ReservedIpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReservedIpMapInput)(nil)).Elem(), ReservedIpMap{})
	pulumi.RegisterOutputType(ReservedIpOutput{})
	pulumi.RegisterOutputType(ReservedIpArrayOutput{})
	pulumi.RegisterOutputType(ReservedIpMapOutput{})
}
