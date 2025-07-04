// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Vultr bare metal server plan.
//
// ## Example Usage
//
// Get the information for a plan by `id`:
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
//			_, err := vultr.GetBareMetalPlan(ctx, &vultr.GetBareMetalPlanArgs{
//				Filters: []vultr.GetBareMetalPlanFilter{
//					{
//						Name: "id",
//						Values: []string{
//							"vbm-4c-32gb",
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
func GetBareMetalPlan(ctx *pulumi.Context, args *GetBareMetalPlanArgs, opts ...pulumi.InvokeOption) (*GetBareMetalPlanResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBareMetalPlanResult
	err := ctx.Invoke("vultr:index/getBareMetalPlan:getBareMetalPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBareMetalPlan.
type GetBareMetalPlanArgs struct {
	// Query parameters for finding plans.
	Filters []GetBareMetalPlanFilter `pulumi:"filters"`
}

// A collection of values returned by getBareMetalPlan.
type GetBareMetalPlanResult struct {
	// The bandwidth available on the plan.
	Bandwidth int `pulumi:"bandwidth"`
	// The number of CPUs available on the plan.
	CpuCount int `pulumi:"cpuCount"`
	// The CPU model of the plan.
	CpuModel string `pulumi:"cpuModel"`
	// The number of CPU threads.
	CpuThreads int `pulumi:"cpuThreads"`
	// The description of the disk(s) on the plan.
	Disk int `pulumi:"disk"`
	// The number of disks that this plan offers.
	DiskCount int                      `pulumi:"diskCount"`
	Filters   []GetBareMetalPlanFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Locations []string `pulumi:"locations"`
	// The price per month of the plan in USD.
	MonthlyCost int `pulumi:"monthlyCost"`
	// The amount of memory available on the plan in MB.
	Ram int `pulumi:"ram"`
	// The type of plan it is.
	Type string `pulumi:"type"`
}

func GetBareMetalPlanOutput(ctx *pulumi.Context, args GetBareMetalPlanOutputArgs, opts ...pulumi.InvokeOption) GetBareMetalPlanResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetBareMetalPlanResultOutput, error) {
			args := v.(GetBareMetalPlanArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vultr:index/getBareMetalPlan:getBareMetalPlan", args, GetBareMetalPlanResultOutput{}, options).(GetBareMetalPlanResultOutput), nil
		}).(GetBareMetalPlanResultOutput)
}

// A collection of arguments for invoking getBareMetalPlan.
type GetBareMetalPlanOutputArgs struct {
	// Query parameters for finding plans.
	Filters GetBareMetalPlanFilterArrayInput `pulumi:"filters"`
}

func (GetBareMetalPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBareMetalPlanArgs)(nil)).Elem()
}

// A collection of values returned by getBareMetalPlan.
type GetBareMetalPlanResultOutput struct{ *pulumi.OutputState }

func (GetBareMetalPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBareMetalPlanResult)(nil)).Elem()
}

func (o GetBareMetalPlanResultOutput) ToGetBareMetalPlanResultOutput() GetBareMetalPlanResultOutput {
	return o
}

func (o GetBareMetalPlanResultOutput) ToGetBareMetalPlanResultOutputWithContext(ctx context.Context) GetBareMetalPlanResultOutput {
	return o
}

// The bandwidth available on the plan.
func (o GetBareMetalPlanResultOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v GetBareMetalPlanResult) int { return v.Bandwidth }).(pulumi.IntOutput)
}

// The number of CPUs available on the plan.
func (o GetBareMetalPlanResultOutput) CpuCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetBareMetalPlanResult) int { return v.CpuCount }).(pulumi.IntOutput)
}

// The CPU model of the plan.
func (o GetBareMetalPlanResultOutput) CpuModel() pulumi.StringOutput {
	return o.ApplyT(func(v GetBareMetalPlanResult) string { return v.CpuModel }).(pulumi.StringOutput)
}

// The number of CPU threads.
func (o GetBareMetalPlanResultOutput) CpuThreads() pulumi.IntOutput {
	return o.ApplyT(func(v GetBareMetalPlanResult) int { return v.CpuThreads }).(pulumi.IntOutput)
}

// The description of the disk(s) on the plan.
func (o GetBareMetalPlanResultOutput) Disk() pulumi.IntOutput {
	return o.ApplyT(func(v GetBareMetalPlanResult) int { return v.Disk }).(pulumi.IntOutput)
}

// The number of disks that this plan offers.
func (o GetBareMetalPlanResultOutput) DiskCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetBareMetalPlanResult) int { return v.DiskCount }).(pulumi.IntOutput)
}

func (o GetBareMetalPlanResultOutput) Filters() GetBareMetalPlanFilterArrayOutput {
	return o.ApplyT(func(v GetBareMetalPlanResult) []GetBareMetalPlanFilter { return v.Filters }).(GetBareMetalPlanFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBareMetalPlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBareMetalPlanResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBareMetalPlanResultOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBareMetalPlanResult) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

// The price per month of the plan in USD.
func (o GetBareMetalPlanResultOutput) MonthlyCost() pulumi.IntOutput {
	return o.ApplyT(func(v GetBareMetalPlanResult) int { return v.MonthlyCost }).(pulumi.IntOutput)
}

// The amount of memory available on the plan in MB.
func (o GetBareMetalPlanResultOutput) Ram() pulumi.IntOutput {
	return o.ApplyT(func(v GetBareMetalPlanResult) int { return v.Ram }).(pulumi.IntOutput)
}

// The type of plan it is.
func (o GetBareMetalPlanResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetBareMetalPlanResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBareMetalPlanResultOutput{})
}
