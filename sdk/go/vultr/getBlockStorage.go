// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Vultr block storage subscription.
//
// ## Example Usage
//
// Get the information for a block storage subscription by `label`:
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
//			_, err := vultr.LookupBlockStorage(ctx, &vultr.LookupBlockStorageArgs{
//				Filters: []vultr.GetBlockStorageFilter{
//					{
//						Name: "label",
//						Values: []string{
//							"my-block-storage-label",
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
func LookupBlockStorage(ctx *pulumi.Context, args *LookupBlockStorageArgs, opts ...pulumi.InvokeOption) (*LookupBlockStorageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBlockStorageResult
	err := ctx.Invoke("vultr:index/getBlockStorage:getBlockStorage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBlockStorage.
type LookupBlockStorageArgs struct {
	// Query parameters for finding block storage subscriptions.
	Filters []GetBlockStorageFilter `pulumi:"filters"`
}

// A collection of values returned by getBlockStorage.
type LookupBlockStorageResult struct {
	// The ID of the VPS the block storage subscription is attached to.
	AttachedToInstance string `pulumi:"attachedToInstance"`
	// The type of block storage volume.
	BlockType string `pulumi:"blockType"`
	// The cost per month of the block storage subscription in USD.
	Cost int `pulumi:"cost"`
	// The date the block storage subscription was added to your Vultr account.
	DateCreated string                  `pulumi:"dateCreated"`
	Filters     []GetBlockStorageFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The label of the block storage subscription.
	Label string `pulumi:"label"`
	// An ID associated with the instance, when mounted the ID can be found in /dev/disk/by-id prefixed with virtio.
	MountId string `pulumi:"mountId"`
	// The region ID of the block storage subscription.
	Region string `pulumi:"region"`
	// The size of the block storage subscription in GB.
	SizeGb int `pulumi:"sizeGb"`
	// The status of the block storage subscription.
	Status string `pulumi:"status"`
}

func LookupBlockStorageOutput(ctx *pulumi.Context, args LookupBlockStorageOutputArgs, opts ...pulumi.InvokeOption) LookupBlockStorageResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupBlockStorageResultOutput, error) {
			args := v.(LookupBlockStorageArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vultr:index/getBlockStorage:getBlockStorage", args, LookupBlockStorageResultOutput{}, options).(LookupBlockStorageResultOutput), nil
		}).(LookupBlockStorageResultOutput)
}

// A collection of arguments for invoking getBlockStorage.
type LookupBlockStorageOutputArgs struct {
	// Query parameters for finding block storage subscriptions.
	Filters GetBlockStorageFilterArrayInput `pulumi:"filters"`
}

func (LookupBlockStorageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlockStorageArgs)(nil)).Elem()
}

// A collection of values returned by getBlockStorage.
type LookupBlockStorageResultOutput struct{ *pulumi.OutputState }

func (LookupBlockStorageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlockStorageResult)(nil)).Elem()
}

func (o LookupBlockStorageResultOutput) ToLookupBlockStorageResultOutput() LookupBlockStorageResultOutput {
	return o
}

func (o LookupBlockStorageResultOutput) ToLookupBlockStorageResultOutputWithContext(ctx context.Context) LookupBlockStorageResultOutput {
	return o
}

// The ID of the VPS the block storage subscription is attached to.
func (o LookupBlockStorageResultOutput) AttachedToInstance() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockStorageResult) string { return v.AttachedToInstance }).(pulumi.StringOutput)
}

// The type of block storage volume.
func (o LookupBlockStorageResultOutput) BlockType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockStorageResult) string { return v.BlockType }).(pulumi.StringOutput)
}

// The cost per month of the block storage subscription in USD.
func (o LookupBlockStorageResultOutput) Cost() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBlockStorageResult) int { return v.Cost }).(pulumi.IntOutput)
}

// The date the block storage subscription was added to your Vultr account.
func (o LookupBlockStorageResultOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockStorageResult) string { return v.DateCreated }).(pulumi.StringOutput)
}

func (o LookupBlockStorageResultOutput) Filters() GetBlockStorageFilterArrayOutput {
	return o.ApplyT(func(v LookupBlockStorageResult) []GetBlockStorageFilter { return v.Filters }).(GetBlockStorageFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupBlockStorageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockStorageResult) string { return v.Id }).(pulumi.StringOutput)
}

// The label of the block storage subscription.
func (o LookupBlockStorageResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockStorageResult) string { return v.Label }).(pulumi.StringOutput)
}

// An ID associated with the instance, when mounted the ID can be found in /dev/disk/by-id prefixed with virtio.
func (o LookupBlockStorageResultOutput) MountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockStorageResult) string { return v.MountId }).(pulumi.StringOutput)
}

// The region ID of the block storage subscription.
func (o LookupBlockStorageResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockStorageResult) string { return v.Region }).(pulumi.StringOutput)
}

// The size of the block storage subscription in GB.
func (o LookupBlockStorageResultOutput) SizeGb() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBlockStorageResult) int { return v.SizeGb }).(pulumi.IntOutput)
}

// The status of the block storage subscription.
func (o LookupBlockStorageResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockStorageResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlockStorageResultOutput{})
}
