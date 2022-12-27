// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Vultr Kubernetes Engine (VKE) Cluster.
//
// ## Example Usage
//
// Create a new VKE cluster:
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
//			_, err := vultr.LookupKubernetes(ctx, &vultr.LookupKubernetesArgs{
//				Filters: []vultr.GetKubernetesFilter{
//					{
//						Name: "label",
//						Values: []string{
//							"my-lb-label",
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
func LookupKubernetes(ctx *pulumi.Context, args *LookupKubernetesArgs, opts ...pulumi.InvokeOption) (*LookupKubernetesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupKubernetesResult
	err := ctx.Invoke("vultr:index/getKubernetes:getKubernetes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKubernetes.
type LookupKubernetesArgs struct {
	// Query parameters for finding VKE.
	Filters []GetKubernetesFilter `pulumi:"filters"`
}

// A collection of values returned by getKubernetes.
type LookupKubernetesResult struct {
	// IP range that your pods will run on in this cluster.
	ClusterSubnet string `pulumi:"clusterSubnet"`
	// Date node was created.
	DateCreated string `pulumi:"dateCreated"`
	// Domain for your Kubernetes clusters control plane.
	Endpoint string                `pulumi:"endpoint"`
	Filters  []GetKubernetesFilter `pulumi:"filters"`
	// ID of node.
	Id string `pulumi:"id"`
	// IP address of VKE cluster control plane.
	Ip string `pulumi:"ip"`
	// Base64 encoded Kubeconfig for this VKE cluster.
	KubeConfig string `pulumi:"kubeConfig"`
	// Label of node.
	Label string `pulumi:"label"`
	// Contains the default node pool that was deployed.
	NodePools []GetKubernetesNodePool `pulumi:"nodePools"`
	// The region your VKE cluster is deployed in.
	Region string `pulumi:"region"`
	// IP range that services will run on this cluster.
	ServiceSubnet string `pulumi:"serviceSubnet"`
	// Status of node.
	Status string `pulumi:"status"`
	// The current kubernetes version your VKE cluster is running on.
	Version string `pulumi:"version"`
}

func LookupKubernetesOutput(ctx *pulumi.Context, args LookupKubernetesOutputArgs, opts ...pulumi.InvokeOption) LookupKubernetesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKubernetesResult, error) {
			args := v.(LookupKubernetesArgs)
			r, err := LookupKubernetes(ctx, &args, opts...)
			var s LookupKubernetesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKubernetesResultOutput)
}

// A collection of arguments for invoking getKubernetes.
type LookupKubernetesOutputArgs struct {
	// Query parameters for finding VKE.
	Filters GetKubernetesFilterArrayInput `pulumi:"filters"`
}

func (LookupKubernetesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubernetesArgs)(nil)).Elem()
}

// A collection of values returned by getKubernetes.
type LookupKubernetesResultOutput struct{ *pulumi.OutputState }

func (LookupKubernetesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubernetesResult)(nil)).Elem()
}

func (o LookupKubernetesResultOutput) ToLookupKubernetesResultOutput() LookupKubernetesResultOutput {
	return o
}

func (o LookupKubernetesResultOutput) ToLookupKubernetesResultOutputWithContext(ctx context.Context) LookupKubernetesResultOutput {
	return o
}

// IP range that your pods will run on in this cluster.
func (o LookupKubernetesResultOutput) ClusterSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesResult) string { return v.ClusterSubnet }).(pulumi.StringOutput)
}

// Date node was created.
func (o LookupKubernetesResultOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesResult) string { return v.DateCreated }).(pulumi.StringOutput)
}

// Domain for your Kubernetes clusters control plane.
func (o LookupKubernetesResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o LookupKubernetesResultOutput) Filters() GetKubernetesFilterArrayOutput {
	return o.ApplyT(func(v LookupKubernetesResult) []GetKubernetesFilter { return v.Filters }).(GetKubernetesFilterArrayOutput)
}

// ID of node.
func (o LookupKubernetesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesResult) string { return v.Id }).(pulumi.StringOutput)
}

// IP address of VKE cluster control plane.
func (o LookupKubernetesResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesResult) string { return v.Ip }).(pulumi.StringOutput)
}

// Base64 encoded Kubeconfig for this VKE cluster.
func (o LookupKubernetesResultOutput) KubeConfig() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesResult) string { return v.KubeConfig }).(pulumi.StringOutput)
}

// Label of node.
func (o LookupKubernetesResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesResult) string { return v.Label }).(pulumi.StringOutput)
}

// Contains the default node pool that was deployed.
func (o LookupKubernetesResultOutput) NodePools() GetKubernetesNodePoolArrayOutput {
	return o.ApplyT(func(v LookupKubernetesResult) []GetKubernetesNodePool { return v.NodePools }).(GetKubernetesNodePoolArrayOutput)
}

// The region your VKE cluster is deployed in.
func (o LookupKubernetesResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesResult) string { return v.Region }).(pulumi.StringOutput)
}

// IP range that services will run on this cluster.
func (o LookupKubernetesResultOutput) ServiceSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesResult) string { return v.ServiceSubnet }).(pulumi.StringOutput)
}

// Status of node.
func (o LookupKubernetesResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesResult) string { return v.Status }).(pulumi.StringOutput)
}

// The current kubernetes version your VKE cluster is running on.
func (o LookupKubernetesResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKubernetesResultOutput{})
}
