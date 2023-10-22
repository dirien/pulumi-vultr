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

// Provides a Vultr Block Storage resource. This can be used to create, read, modify, and delete Block Storage.
//
// ## Example Usage
//
// # Create a new Block Storage
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
//			_, err := vultr.NewBlockStorage(ctx, "myBlockstorage", &vultr.BlockStorageArgs{
//				Label:  pulumi.String("vultr-block-storage"),
//				Region: pulumi.String("ewr"),
//				SizeGb: pulumi.Int(10),
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
// Block Storage can be imported using the Block Storage `ID`, e.g.
//
// ```sh
//
//	$ pulumi import vultr:index/blockStorage:BlockStorage my_blockstorage e315835e-d466-4e89-9b4c-dfd8788d7685
//
// ```
type BlockStorage struct {
	pulumi.CustomResourceState

	// VPS ID that you want to have this block storage attached to.
	AttachedToInstance pulumi.StringPtrOutput `pulumi:"attachedToInstance"`
	// Determines on the type of block storage volume that will be created. Soon to become a required parameter. Options are `highPerf` or `storageOpt`.
	BlockType pulumi.StringOutput `pulumi:"blockType"`
	// The monthly cost of this block storage.
	Cost pulumi.Float64Output `pulumi:"cost"`
	// The date this block storage was created.
	DateCreated pulumi.StringOutput `pulumi:"dateCreated"`
	// Label that is given to your block storage.
	Label pulumi.StringPtrOutput `pulumi:"label"`
	// Boolean value that will allow attachment of the volume to an instance without a restart. Default is false.
	Live pulumi.BoolPtrOutput `pulumi:"live"`
	// An ID associated with the instance, when mounted the ID can be found in /dev/disk/by-id prefixed with virtio.
	MountId pulumi.StringOutput `pulumi:"mountId"`
	// Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")
	Region pulumi.StringOutput `pulumi:"region"`
	// The size of the given block storage.
	SizeGb pulumi.IntOutput `pulumi:"sizeGb"`
	// Current status of your block storage.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewBlockStorage registers a new resource with the given unique name, arguments, and options.
func NewBlockStorage(ctx *pulumi.Context,
	name string, args *BlockStorageArgs, opts ...pulumi.ResourceOption) (*BlockStorage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.SizeGb == nil {
		return nil, errors.New("invalid value for required argument 'SizeGb'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BlockStorage
	err := ctx.RegisterResource("vultr:index/blockStorage:BlockStorage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlockStorage gets an existing BlockStorage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlockStorage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlockStorageState, opts ...pulumi.ResourceOption) (*BlockStorage, error) {
	var resource BlockStorage
	err := ctx.ReadResource("vultr:index/blockStorage:BlockStorage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BlockStorage resources.
type blockStorageState struct {
	// VPS ID that you want to have this block storage attached to.
	AttachedToInstance *string `pulumi:"attachedToInstance"`
	// Determines on the type of block storage volume that will be created. Soon to become a required parameter. Options are `highPerf` or `storageOpt`.
	BlockType *string `pulumi:"blockType"`
	// The monthly cost of this block storage.
	Cost *float64 `pulumi:"cost"`
	// The date this block storage was created.
	DateCreated *string `pulumi:"dateCreated"`
	// Label that is given to your block storage.
	Label *string `pulumi:"label"`
	// Boolean value that will allow attachment of the volume to an instance without a restart. Default is false.
	Live *bool `pulumi:"live"`
	// An ID associated with the instance, when mounted the ID can be found in /dev/disk/by-id prefixed with virtio.
	MountId *string `pulumi:"mountId"`
	// Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")
	Region *string `pulumi:"region"`
	// The size of the given block storage.
	SizeGb *int `pulumi:"sizeGb"`
	// Current status of your block storage.
	Status *string `pulumi:"status"`
}

type BlockStorageState struct {
	// VPS ID that you want to have this block storage attached to.
	AttachedToInstance pulumi.StringPtrInput
	// Determines on the type of block storage volume that will be created. Soon to become a required parameter. Options are `highPerf` or `storageOpt`.
	BlockType pulumi.StringPtrInput
	// The monthly cost of this block storage.
	Cost pulumi.Float64PtrInput
	// The date this block storage was created.
	DateCreated pulumi.StringPtrInput
	// Label that is given to your block storage.
	Label pulumi.StringPtrInput
	// Boolean value that will allow attachment of the volume to an instance without a restart. Default is false.
	Live pulumi.BoolPtrInput
	// An ID associated with the instance, when mounted the ID can be found in /dev/disk/by-id prefixed with virtio.
	MountId pulumi.StringPtrInput
	// Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")
	Region pulumi.StringPtrInput
	// The size of the given block storage.
	SizeGb pulumi.IntPtrInput
	// Current status of your block storage.
	Status pulumi.StringPtrInput
}

func (BlockStorageState) ElementType() reflect.Type {
	return reflect.TypeOf((*blockStorageState)(nil)).Elem()
}

type blockStorageArgs struct {
	// VPS ID that you want to have this block storage attached to.
	AttachedToInstance *string `pulumi:"attachedToInstance"`
	// Determines on the type of block storage volume that will be created. Soon to become a required parameter. Options are `highPerf` or `storageOpt`.
	BlockType *string `pulumi:"blockType"`
	// Label that is given to your block storage.
	Label *string `pulumi:"label"`
	// Boolean value that will allow attachment of the volume to an instance without a restart. Default is false.
	Live *bool `pulumi:"live"`
	// Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")
	Region string `pulumi:"region"`
	// The size of the given block storage.
	SizeGb int `pulumi:"sizeGb"`
}

// The set of arguments for constructing a BlockStorage resource.
type BlockStorageArgs struct {
	// VPS ID that you want to have this block storage attached to.
	AttachedToInstance pulumi.StringPtrInput
	// Determines on the type of block storage volume that will be created. Soon to become a required parameter. Options are `highPerf` or `storageOpt`.
	BlockType pulumi.StringPtrInput
	// Label that is given to your block storage.
	Label pulumi.StringPtrInput
	// Boolean value that will allow attachment of the volume to an instance without a restart. Default is false.
	Live pulumi.BoolPtrInput
	// Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")
	Region pulumi.StringInput
	// The size of the given block storage.
	SizeGb pulumi.IntInput
}

func (BlockStorageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blockStorageArgs)(nil)).Elem()
}

type BlockStorageInput interface {
	pulumi.Input

	ToBlockStorageOutput() BlockStorageOutput
	ToBlockStorageOutputWithContext(ctx context.Context) BlockStorageOutput
}

func (*BlockStorage) ElementType() reflect.Type {
	return reflect.TypeOf((**BlockStorage)(nil)).Elem()
}

func (i *BlockStorage) ToBlockStorageOutput() BlockStorageOutput {
	return i.ToBlockStorageOutputWithContext(context.Background())
}

func (i *BlockStorage) ToBlockStorageOutputWithContext(ctx context.Context) BlockStorageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockStorageOutput)
}

func (i *BlockStorage) ToOutput(ctx context.Context) pulumix.Output[*BlockStorage] {
	return pulumix.Output[*BlockStorage]{
		OutputState: i.ToBlockStorageOutputWithContext(ctx).OutputState,
	}
}

// BlockStorageArrayInput is an input type that accepts BlockStorageArray and BlockStorageArrayOutput values.
// You can construct a concrete instance of `BlockStorageArrayInput` via:
//
//	BlockStorageArray{ BlockStorageArgs{...} }
type BlockStorageArrayInput interface {
	pulumi.Input

	ToBlockStorageArrayOutput() BlockStorageArrayOutput
	ToBlockStorageArrayOutputWithContext(context.Context) BlockStorageArrayOutput
}

type BlockStorageArray []BlockStorageInput

func (BlockStorageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BlockStorage)(nil)).Elem()
}

func (i BlockStorageArray) ToBlockStorageArrayOutput() BlockStorageArrayOutput {
	return i.ToBlockStorageArrayOutputWithContext(context.Background())
}

func (i BlockStorageArray) ToBlockStorageArrayOutputWithContext(ctx context.Context) BlockStorageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockStorageArrayOutput)
}

func (i BlockStorageArray) ToOutput(ctx context.Context) pulumix.Output[[]*BlockStorage] {
	return pulumix.Output[[]*BlockStorage]{
		OutputState: i.ToBlockStorageArrayOutputWithContext(ctx).OutputState,
	}
}

// BlockStorageMapInput is an input type that accepts BlockStorageMap and BlockStorageMapOutput values.
// You can construct a concrete instance of `BlockStorageMapInput` via:
//
//	BlockStorageMap{ "key": BlockStorageArgs{...} }
type BlockStorageMapInput interface {
	pulumi.Input

	ToBlockStorageMapOutput() BlockStorageMapOutput
	ToBlockStorageMapOutputWithContext(context.Context) BlockStorageMapOutput
}

type BlockStorageMap map[string]BlockStorageInput

func (BlockStorageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BlockStorage)(nil)).Elem()
}

func (i BlockStorageMap) ToBlockStorageMapOutput() BlockStorageMapOutput {
	return i.ToBlockStorageMapOutputWithContext(context.Background())
}

func (i BlockStorageMap) ToBlockStorageMapOutputWithContext(ctx context.Context) BlockStorageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockStorageMapOutput)
}

func (i BlockStorageMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*BlockStorage] {
	return pulumix.Output[map[string]*BlockStorage]{
		OutputState: i.ToBlockStorageMapOutputWithContext(ctx).OutputState,
	}
}

type BlockStorageOutput struct{ *pulumi.OutputState }

func (BlockStorageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlockStorage)(nil)).Elem()
}

func (o BlockStorageOutput) ToBlockStorageOutput() BlockStorageOutput {
	return o
}

func (o BlockStorageOutput) ToBlockStorageOutputWithContext(ctx context.Context) BlockStorageOutput {
	return o
}

func (o BlockStorageOutput) ToOutput(ctx context.Context) pulumix.Output[*BlockStorage] {
	return pulumix.Output[*BlockStorage]{
		OutputState: o.OutputState,
	}
}

// VPS ID that you want to have this block storage attached to.
func (o BlockStorageOutput) AttachedToInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlockStorage) pulumi.StringPtrOutput { return v.AttachedToInstance }).(pulumi.StringPtrOutput)
}

// Determines on the type of block storage volume that will be created. Soon to become a required parameter. Options are `highPerf` or `storageOpt`.
func (o BlockStorageOutput) BlockType() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockStorage) pulumi.StringOutput { return v.BlockType }).(pulumi.StringOutput)
}

// The monthly cost of this block storage.
func (o BlockStorageOutput) Cost() pulumi.Float64Output {
	return o.ApplyT(func(v *BlockStorage) pulumi.Float64Output { return v.Cost }).(pulumi.Float64Output)
}

// The date this block storage was created.
func (o BlockStorageOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockStorage) pulumi.StringOutput { return v.DateCreated }).(pulumi.StringOutput)
}

// Label that is given to your block storage.
func (o BlockStorageOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlockStorage) pulumi.StringPtrOutput { return v.Label }).(pulumi.StringPtrOutput)
}

// Boolean value that will allow attachment of the volume to an instance without a restart. Default is false.
func (o BlockStorageOutput) Live() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BlockStorage) pulumi.BoolPtrOutput { return v.Live }).(pulumi.BoolPtrOutput)
}

// An ID associated with the instance, when mounted the ID can be found in /dev/disk/by-id prefixed with virtio.
func (o BlockStorageOutput) MountId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockStorage) pulumi.StringOutput { return v.MountId }).(pulumi.StringOutput)
}

// Region in which this block storage will reside in. (Currently only NJ/NY supported region "ewr")
func (o BlockStorageOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockStorage) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The size of the given block storage.
func (o BlockStorageOutput) SizeGb() pulumi.IntOutput {
	return o.ApplyT(func(v *BlockStorage) pulumi.IntOutput { return v.SizeGb }).(pulumi.IntOutput)
}

// Current status of your block storage.
func (o BlockStorageOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BlockStorage) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type BlockStorageArrayOutput struct{ *pulumi.OutputState }

func (BlockStorageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BlockStorage)(nil)).Elem()
}

func (o BlockStorageArrayOutput) ToBlockStorageArrayOutput() BlockStorageArrayOutput {
	return o
}

func (o BlockStorageArrayOutput) ToBlockStorageArrayOutputWithContext(ctx context.Context) BlockStorageArrayOutput {
	return o
}

func (o BlockStorageArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*BlockStorage] {
	return pulumix.Output[[]*BlockStorage]{
		OutputState: o.OutputState,
	}
}

func (o BlockStorageArrayOutput) Index(i pulumi.IntInput) BlockStorageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BlockStorage {
		return vs[0].([]*BlockStorage)[vs[1].(int)]
	}).(BlockStorageOutput)
}

type BlockStorageMapOutput struct{ *pulumi.OutputState }

func (BlockStorageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BlockStorage)(nil)).Elem()
}

func (o BlockStorageMapOutput) ToBlockStorageMapOutput() BlockStorageMapOutput {
	return o
}

func (o BlockStorageMapOutput) ToBlockStorageMapOutputWithContext(ctx context.Context) BlockStorageMapOutput {
	return o
}

func (o BlockStorageMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*BlockStorage] {
	return pulumix.Output[map[string]*BlockStorage]{
		OutputState: o.OutputState,
	}
}

func (o BlockStorageMapOutput) MapIndex(k pulumi.StringInput) BlockStorageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BlockStorage {
		return vs[0].(map[string]*BlockStorage)[vs[1].(string)]
	}).(BlockStorageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BlockStorageInput)(nil)).Elem(), &BlockStorage{})
	pulumi.RegisterInputType(reflect.TypeOf((*BlockStorageArrayInput)(nil)).Elem(), BlockStorageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BlockStorageMapInput)(nil)).Elem(), BlockStorageMap{})
	pulumi.RegisterOutputType(BlockStorageOutput{})
	pulumi.RegisterOutputType(BlockStorageArrayOutput{})
	pulumi.RegisterOutputType(BlockStorageMapOutput{})
}
