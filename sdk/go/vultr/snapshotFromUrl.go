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

// Provides a Vultr Snapshots from URL resource. This can be used to create, read, modify, and delete Snapshots from URL.
//
// ## Example Usage
//
// # Create a new Snapshots from URL
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
//			_, err := vultr.NewSnapshotFromUrl(ctx, "mySnapshot", &vultr.SnapshotFromUrlArgs{
//				Url: pulumi.String("http://dl-cdn.alpinelinux.org/alpine/v3.9/releases/x86_64/alpine-virt-3.9.1-x86_64.iso"),
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
// Snapshots from Url can be imported using the Snapshot `ID`, e.g.
//
// ```sh
//
//	$ pulumi import vultr:index/snapshotFromUrl:SnapshotFromUrl my_snapshot e60dc0a2-9313-4bab-bffc-57ffe33d99f6
//
// ```
type SnapshotFromUrl struct {
	pulumi.CustomResourceState

	// The app id which the snapshot is associated with.
	AppId pulumi.IntOutput `pulumi:"appId"`
	// The date the snapshot was created.
	DateCreated pulumi.StringOutput `pulumi:"dateCreated"`
	// The description for the given snapshot.
	Description pulumi.StringOutput `pulumi:"description"`
	// The os id which the snapshot is associated with.
	OsId pulumi.IntOutput `pulumi:"osId"`
	// The size of the snapshot in Bytes.
	Size pulumi.IntOutput `pulumi:"size"`
	// The status for the given snapshot.
	Status pulumi.StringOutput `pulumi:"status"`
	// URL of the given resource you want to create a snapshot from.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewSnapshotFromUrl registers a new resource with the given unique name, arguments, and options.
func NewSnapshotFromUrl(ctx *pulumi.Context,
	name string, args *SnapshotFromUrlArgs, opts ...pulumi.ResourceOption) (*SnapshotFromUrl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SnapshotFromUrl
	err := ctx.RegisterResource("vultr:index/snapshotFromUrl:SnapshotFromUrl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotFromUrl gets an existing SnapshotFromUrl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotFromUrl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotFromUrlState, opts ...pulumi.ResourceOption) (*SnapshotFromUrl, error) {
	var resource SnapshotFromUrl
	err := ctx.ReadResource("vultr:index/snapshotFromUrl:SnapshotFromUrl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotFromUrl resources.
type snapshotFromUrlState struct {
	// The app id which the snapshot is associated with.
	AppId *int `pulumi:"appId"`
	// The date the snapshot was created.
	DateCreated *string `pulumi:"dateCreated"`
	// The description for the given snapshot.
	Description *string `pulumi:"description"`
	// The os id which the snapshot is associated with.
	OsId *int `pulumi:"osId"`
	// The size of the snapshot in Bytes.
	Size *int `pulumi:"size"`
	// The status for the given snapshot.
	Status *string `pulumi:"status"`
	// URL of the given resource you want to create a snapshot from.
	Url *string `pulumi:"url"`
}

type SnapshotFromUrlState struct {
	// The app id which the snapshot is associated with.
	AppId pulumi.IntPtrInput
	// The date the snapshot was created.
	DateCreated pulumi.StringPtrInput
	// The description for the given snapshot.
	Description pulumi.StringPtrInput
	// The os id which the snapshot is associated with.
	OsId pulumi.IntPtrInput
	// The size of the snapshot in Bytes.
	Size pulumi.IntPtrInput
	// The status for the given snapshot.
	Status pulumi.StringPtrInput
	// URL of the given resource you want to create a snapshot from.
	Url pulumi.StringPtrInput
}

func (SnapshotFromUrlState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotFromUrlState)(nil)).Elem()
}

type snapshotFromUrlArgs struct {
	// URL of the given resource you want to create a snapshot from.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a SnapshotFromUrl resource.
type SnapshotFromUrlArgs struct {
	// URL of the given resource you want to create a snapshot from.
	Url pulumi.StringInput
}

func (SnapshotFromUrlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotFromUrlArgs)(nil)).Elem()
}

type SnapshotFromUrlInput interface {
	pulumi.Input

	ToSnapshotFromUrlOutput() SnapshotFromUrlOutput
	ToSnapshotFromUrlOutputWithContext(ctx context.Context) SnapshotFromUrlOutput
}

func (*SnapshotFromUrl) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotFromUrl)(nil)).Elem()
}

func (i *SnapshotFromUrl) ToSnapshotFromUrlOutput() SnapshotFromUrlOutput {
	return i.ToSnapshotFromUrlOutputWithContext(context.Background())
}

func (i *SnapshotFromUrl) ToSnapshotFromUrlOutputWithContext(ctx context.Context) SnapshotFromUrlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotFromUrlOutput)
}

func (i *SnapshotFromUrl) ToOutput(ctx context.Context) pulumix.Output[*SnapshotFromUrl] {
	return pulumix.Output[*SnapshotFromUrl]{
		OutputState: i.ToSnapshotFromUrlOutputWithContext(ctx).OutputState,
	}
}

// SnapshotFromUrlArrayInput is an input type that accepts SnapshotFromUrlArray and SnapshotFromUrlArrayOutput values.
// You can construct a concrete instance of `SnapshotFromUrlArrayInput` via:
//
//	SnapshotFromUrlArray{ SnapshotFromUrlArgs{...} }
type SnapshotFromUrlArrayInput interface {
	pulumi.Input

	ToSnapshotFromUrlArrayOutput() SnapshotFromUrlArrayOutput
	ToSnapshotFromUrlArrayOutputWithContext(context.Context) SnapshotFromUrlArrayOutput
}

type SnapshotFromUrlArray []SnapshotFromUrlInput

func (SnapshotFromUrlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotFromUrl)(nil)).Elem()
}

func (i SnapshotFromUrlArray) ToSnapshotFromUrlArrayOutput() SnapshotFromUrlArrayOutput {
	return i.ToSnapshotFromUrlArrayOutputWithContext(context.Background())
}

func (i SnapshotFromUrlArray) ToSnapshotFromUrlArrayOutputWithContext(ctx context.Context) SnapshotFromUrlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotFromUrlArrayOutput)
}

func (i SnapshotFromUrlArray) ToOutput(ctx context.Context) pulumix.Output[[]*SnapshotFromUrl] {
	return pulumix.Output[[]*SnapshotFromUrl]{
		OutputState: i.ToSnapshotFromUrlArrayOutputWithContext(ctx).OutputState,
	}
}

// SnapshotFromUrlMapInput is an input type that accepts SnapshotFromUrlMap and SnapshotFromUrlMapOutput values.
// You can construct a concrete instance of `SnapshotFromUrlMapInput` via:
//
//	SnapshotFromUrlMap{ "key": SnapshotFromUrlArgs{...} }
type SnapshotFromUrlMapInput interface {
	pulumi.Input

	ToSnapshotFromUrlMapOutput() SnapshotFromUrlMapOutput
	ToSnapshotFromUrlMapOutputWithContext(context.Context) SnapshotFromUrlMapOutput
}

type SnapshotFromUrlMap map[string]SnapshotFromUrlInput

func (SnapshotFromUrlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotFromUrl)(nil)).Elem()
}

func (i SnapshotFromUrlMap) ToSnapshotFromUrlMapOutput() SnapshotFromUrlMapOutput {
	return i.ToSnapshotFromUrlMapOutputWithContext(context.Background())
}

func (i SnapshotFromUrlMap) ToSnapshotFromUrlMapOutputWithContext(ctx context.Context) SnapshotFromUrlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotFromUrlMapOutput)
}

func (i SnapshotFromUrlMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SnapshotFromUrl] {
	return pulumix.Output[map[string]*SnapshotFromUrl]{
		OutputState: i.ToSnapshotFromUrlMapOutputWithContext(ctx).OutputState,
	}
}

type SnapshotFromUrlOutput struct{ *pulumi.OutputState }

func (SnapshotFromUrlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotFromUrl)(nil)).Elem()
}

func (o SnapshotFromUrlOutput) ToSnapshotFromUrlOutput() SnapshotFromUrlOutput {
	return o
}

func (o SnapshotFromUrlOutput) ToSnapshotFromUrlOutputWithContext(ctx context.Context) SnapshotFromUrlOutput {
	return o
}

func (o SnapshotFromUrlOutput) ToOutput(ctx context.Context) pulumix.Output[*SnapshotFromUrl] {
	return pulumix.Output[*SnapshotFromUrl]{
		OutputState: o.OutputState,
	}
}

// The app id which the snapshot is associated with.
func (o SnapshotFromUrlOutput) AppId() pulumi.IntOutput {
	return o.ApplyT(func(v *SnapshotFromUrl) pulumi.IntOutput { return v.AppId }).(pulumi.IntOutput)
}

// The date the snapshot was created.
func (o SnapshotFromUrlOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotFromUrl) pulumi.StringOutput { return v.DateCreated }).(pulumi.StringOutput)
}

// The description for the given snapshot.
func (o SnapshotFromUrlOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotFromUrl) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The os id which the snapshot is associated with.
func (o SnapshotFromUrlOutput) OsId() pulumi.IntOutput {
	return o.ApplyT(func(v *SnapshotFromUrl) pulumi.IntOutput { return v.OsId }).(pulumi.IntOutput)
}

// The size of the snapshot in Bytes.
func (o SnapshotFromUrlOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *SnapshotFromUrl) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// The status for the given snapshot.
func (o SnapshotFromUrlOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotFromUrl) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// URL of the given resource you want to create a snapshot from.
func (o SnapshotFromUrlOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotFromUrl) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type SnapshotFromUrlArrayOutput struct{ *pulumi.OutputState }

func (SnapshotFromUrlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotFromUrl)(nil)).Elem()
}

func (o SnapshotFromUrlArrayOutput) ToSnapshotFromUrlArrayOutput() SnapshotFromUrlArrayOutput {
	return o
}

func (o SnapshotFromUrlArrayOutput) ToSnapshotFromUrlArrayOutputWithContext(ctx context.Context) SnapshotFromUrlArrayOutput {
	return o
}

func (o SnapshotFromUrlArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SnapshotFromUrl] {
	return pulumix.Output[[]*SnapshotFromUrl]{
		OutputState: o.OutputState,
	}
}

func (o SnapshotFromUrlArrayOutput) Index(i pulumi.IntInput) SnapshotFromUrlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SnapshotFromUrl {
		return vs[0].([]*SnapshotFromUrl)[vs[1].(int)]
	}).(SnapshotFromUrlOutput)
}

type SnapshotFromUrlMapOutput struct{ *pulumi.OutputState }

func (SnapshotFromUrlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotFromUrl)(nil)).Elem()
}

func (o SnapshotFromUrlMapOutput) ToSnapshotFromUrlMapOutput() SnapshotFromUrlMapOutput {
	return o
}

func (o SnapshotFromUrlMapOutput) ToSnapshotFromUrlMapOutputWithContext(ctx context.Context) SnapshotFromUrlMapOutput {
	return o
}

func (o SnapshotFromUrlMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SnapshotFromUrl] {
	return pulumix.Output[map[string]*SnapshotFromUrl]{
		OutputState: o.OutputState,
	}
}

func (o SnapshotFromUrlMapOutput) MapIndex(k pulumi.StringInput) SnapshotFromUrlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SnapshotFromUrl {
		return vs[0].(map[string]*SnapshotFromUrl)[vs[1].(string)]
	}).(SnapshotFromUrlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotFromUrlInput)(nil)).Elem(), &SnapshotFromUrl{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotFromUrlArrayInput)(nil)).Elem(), SnapshotFromUrlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotFromUrlMapInput)(nil)).Elem(), SnapshotFromUrlMap{})
	pulumi.RegisterOutputType(SnapshotFromUrlOutput{})
	pulumi.RegisterOutputType(SnapshotFromUrlArrayOutput{})
	pulumi.RegisterOutputType(SnapshotFromUrlMapOutput{})
}
