// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// Create a new VKE cluster:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vultr/sdk/v2/go/vultr"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vultr.NewKubernetes(ctx, "k8", &vultr.KubernetesArgs{
//				Label: pulumi.String("tf-test"),
//				NodePools: &vultr.KubernetesNodePoolsTypeArgs{
//					AutoScaler:   pulumi.Bool(true),
//					Label:        pulumi.String("my-label"),
//					MaxNodes:     pulumi.Int(2),
//					MinNodes:     pulumi.Int(1),
//					NodeQuantity: pulumi.Int(1),
//					Plan:         pulumi.String("vc2-2c-4gb"),
//				},
//				Region:  pulumi.String("ewr"),
//				Version: pulumi.String("v1.23.5+1"),
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
// A default node pool is required when first creating the resource but it can be removed at a later point so long as there is a separate `KubernetesNodePools` resource attached. For example:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vultr/sdk/v2/go/vultr"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			k8, err := vultr.NewKubernetes(ctx, "k8", &vultr.KubernetesArgs{
//				Region:  pulumi.String("ewr"),
//				Label:   pulumi.String("tf-test"),
//				Version: pulumi.String("v1.23.5+1"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vultr.NewKubernetesNodePools(ctx, "np", &vultr.KubernetesNodePoolsArgs{
//				ClusterId:    k8.ID(),
//				NodeQuantity: pulumi.Int(1),
//				Plan:         pulumi.String("vc2-2c-4gb"),
//				Label:        pulumi.String("my-label"),
//				AutoScaler:   pulumi.Bool(true),
//				MinNodes:     pulumi.Int(1),
//				MaxNodes:     pulumi.Int(2),
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
// There is still a requirement that there be one node pool attached to the cluster but this should allow more flexibility about which node pool that is.
type Kubernetes struct {
	pulumi.CustomResourceState

	// IP range that your pods will run on in this cluster.
	ClusterSubnet pulumi.StringOutput `pulumi:"clusterSubnet"`
	// Date node was created.
	DateCreated pulumi.StringOutput `pulumi:"dateCreated"`
	// Domain for your Kubernetes clusters control plane.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// IP address of VKE cluster control plane.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Base64 encoded Kubeconfig for this VKE cluster.
	KubeConfig pulumi.StringOutput `pulumi:"kubeConfig"`
	// The VKE clusters label.
	Label pulumi.StringOutput `pulumi:"label"`
	// Contains the default node pool that was deployed.
	NodePools KubernetesNodePoolsTypePtrOutput `pulumi:"nodePools"`
	// The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
	Region pulumi.StringOutput `pulumi:"region"`
	// IP range that services will run on this cluster.
	ServiceSubnet pulumi.StringOutput `pulumi:"serviceSubnet"`
	// Status of node.
	Status pulumi.StringOutput `pulumi:"status"`
	// The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewKubernetes registers a new resource with the given unique name, arguments, and options.
func NewKubernetes(ctx *pulumi.Context,
	name string, args *KubernetesArgs, opts ...pulumi.ResourceOption) (*Kubernetes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Label == nil {
		return nil, errors.New("invalid value for required argument 'Label'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Kubernetes
	err := ctx.RegisterResource("vultr:index/kubernetes:Kubernetes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetes gets an existing Kubernetes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesState, opts ...pulumi.ResourceOption) (*Kubernetes, error) {
	var resource Kubernetes
	err := ctx.ReadResource("vultr:index/kubernetes:Kubernetes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Kubernetes resources.
type kubernetesState struct {
	// IP range that your pods will run on in this cluster.
	ClusterSubnet *string `pulumi:"clusterSubnet"`
	// Date node was created.
	DateCreated *string `pulumi:"dateCreated"`
	// Domain for your Kubernetes clusters control plane.
	Endpoint *string `pulumi:"endpoint"`
	// IP address of VKE cluster control plane.
	Ip *string `pulumi:"ip"`
	// Base64 encoded Kubeconfig for this VKE cluster.
	KubeConfig *string `pulumi:"kubeConfig"`
	// The VKE clusters label.
	Label *string `pulumi:"label"`
	// Contains the default node pool that was deployed.
	NodePools *KubernetesNodePoolsType `pulumi:"nodePools"`
	// The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
	Region *string `pulumi:"region"`
	// IP range that services will run on this cluster.
	ServiceSubnet *string `pulumi:"serviceSubnet"`
	// Status of node.
	Status *string `pulumi:"status"`
	// The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
	Version *string `pulumi:"version"`
}

type KubernetesState struct {
	// IP range that your pods will run on in this cluster.
	ClusterSubnet pulumi.StringPtrInput
	// Date node was created.
	DateCreated pulumi.StringPtrInput
	// Domain for your Kubernetes clusters control plane.
	Endpoint pulumi.StringPtrInput
	// IP address of VKE cluster control plane.
	Ip pulumi.StringPtrInput
	// Base64 encoded Kubeconfig for this VKE cluster.
	KubeConfig pulumi.StringPtrInput
	// The VKE clusters label.
	Label pulumi.StringPtrInput
	// Contains the default node pool that was deployed.
	NodePools KubernetesNodePoolsTypePtrInput
	// The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
	Region pulumi.StringPtrInput
	// IP range that services will run on this cluster.
	ServiceSubnet pulumi.StringPtrInput
	// Status of node.
	Status pulumi.StringPtrInput
	// The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
	Version pulumi.StringPtrInput
}

func (KubernetesState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesState)(nil)).Elem()
}

type kubernetesArgs struct {
	// The VKE clusters label.
	Label string `pulumi:"label"`
	// Contains the default node pool that was deployed.
	NodePools *KubernetesNodePoolsType `pulumi:"nodePools"`
	// The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
	Region string `pulumi:"region"`
	// The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a Kubernetes resource.
type KubernetesArgs struct {
	// The VKE clusters label.
	Label pulumi.StringInput
	// Contains the default node pool that was deployed.
	NodePools KubernetesNodePoolsTypePtrInput
	// The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
	Region pulumi.StringInput
	// The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
	Version pulumi.StringInput
}

func (KubernetesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesArgs)(nil)).Elem()
}

type KubernetesInput interface {
	pulumi.Input

	ToKubernetesOutput() KubernetesOutput
	ToKubernetesOutputWithContext(ctx context.Context) KubernetesOutput
}

func (*Kubernetes) ElementType() reflect.Type {
	return reflect.TypeOf((**Kubernetes)(nil)).Elem()
}

func (i *Kubernetes) ToKubernetesOutput() KubernetesOutput {
	return i.ToKubernetesOutputWithContext(context.Background())
}

func (i *Kubernetes) ToKubernetesOutputWithContext(ctx context.Context) KubernetesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesOutput)
}

// KubernetesArrayInput is an input type that accepts KubernetesArray and KubernetesArrayOutput values.
// You can construct a concrete instance of `KubernetesArrayInput` via:
//
//	KubernetesArray{ KubernetesArgs{...} }
type KubernetesArrayInput interface {
	pulumi.Input

	ToKubernetesArrayOutput() KubernetesArrayOutput
	ToKubernetesArrayOutputWithContext(context.Context) KubernetesArrayOutput
}

type KubernetesArray []KubernetesInput

func (KubernetesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Kubernetes)(nil)).Elem()
}

func (i KubernetesArray) ToKubernetesArrayOutput() KubernetesArrayOutput {
	return i.ToKubernetesArrayOutputWithContext(context.Background())
}

func (i KubernetesArray) ToKubernetesArrayOutputWithContext(ctx context.Context) KubernetesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesArrayOutput)
}

// KubernetesMapInput is an input type that accepts KubernetesMap and KubernetesMapOutput values.
// You can construct a concrete instance of `KubernetesMapInput` via:
//
//	KubernetesMap{ "key": KubernetesArgs{...} }
type KubernetesMapInput interface {
	pulumi.Input

	ToKubernetesMapOutput() KubernetesMapOutput
	ToKubernetesMapOutputWithContext(context.Context) KubernetesMapOutput
}

type KubernetesMap map[string]KubernetesInput

func (KubernetesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Kubernetes)(nil)).Elem()
}

func (i KubernetesMap) ToKubernetesMapOutput() KubernetesMapOutput {
	return i.ToKubernetesMapOutputWithContext(context.Background())
}

func (i KubernetesMap) ToKubernetesMapOutputWithContext(ctx context.Context) KubernetesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesMapOutput)
}

type KubernetesOutput struct{ *pulumi.OutputState }

func (KubernetesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Kubernetes)(nil)).Elem()
}

func (o KubernetesOutput) ToKubernetesOutput() KubernetesOutput {
	return o
}

func (o KubernetesOutput) ToKubernetesOutputWithContext(ctx context.Context) KubernetesOutput {
	return o
}

// IP range that your pods will run on in this cluster.
func (o KubernetesOutput) ClusterSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringOutput { return v.ClusterSubnet }).(pulumi.StringOutput)
}

// Date node was created.
func (o KubernetesOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringOutput { return v.DateCreated }).(pulumi.StringOutput)
}

// Domain for your Kubernetes clusters control plane.
func (o KubernetesOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// IP address of VKE cluster control plane.
func (o KubernetesOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// Base64 encoded Kubeconfig for this VKE cluster.
func (o KubernetesOutput) KubeConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringOutput { return v.KubeConfig }).(pulumi.StringOutput)
}

// The VKE clusters label.
func (o KubernetesOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringOutput { return v.Label }).(pulumi.StringOutput)
}

// Contains the default node pool that was deployed.
func (o KubernetesOutput) NodePools() KubernetesNodePoolsTypePtrOutput {
	return o.ApplyT(func(v *Kubernetes) KubernetesNodePoolsTypePtrOutput { return v.NodePools }).(KubernetesNodePoolsTypePtrOutput)
}

// The region your VKE cluster will be deployed in. Currently, supported values are `ewr` and `lax`
func (o KubernetesOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// IP range that services will run on this cluster.
func (o KubernetesOutput) ServiceSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringOutput { return v.ServiceSubnet }).(pulumi.StringOutput)
}

// Status of node.
func (o KubernetesOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The version your VKE cluster you want deployed. [See Available Version](https://www.vultr.com/api/#operation/get-kubernetes-versions)
func (o KubernetesOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type KubernetesArrayOutput struct{ *pulumi.OutputState }

func (KubernetesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Kubernetes)(nil)).Elem()
}

func (o KubernetesArrayOutput) ToKubernetesArrayOutput() KubernetesArrayOutput {
	return o
}

func (o KubernetesArrayOutput) ToKubernetesArrayOutputWithContext(ctx context.Context) KubernetesArrayOutput {
	return o
}

func (o KubernetesArrayOutput) Index(i pulumi.IntInput) KubernetesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Kubernetes {
		return vs[0].([]*Kubernetes)[vs[1].(int)]
	}).(KubernetesOutput)
}

type KubernetesMapOutput struct{ *pulumi.OutputState }

func (KubernetesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Kubernetes)(nil)).Elem()
}

func (o KubernetesMapOutput) ToKubernetesMapOutput() KubernetesMapOutput {
	return o
}

func (o KubernetesMapOutput) ToKubernetesMapOutputWithContext(ctx context.Context) KubernetesMapOutput {
	return o
}

func (o KubernetesMapOutput) MapIndex(k pulumi.StringInput) KubernetesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Kubernetes {
		return vs[0].(map[string]*Kubernetes)[vs[1].(string)]
	}).(KubernetesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesInput)(nil)).Elem(), &Kubernetes{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesArrayInput)(nil)).Elem(), KubernetesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesMapInput)(nil)).Elem(), KubernetesMap{})
	pulumi.RegisterOutputType(KubernetesOutput{})
	pulumi.RegisterOutputType(KubernetesArrayOutput{})
	pulumi.RegisterOutputType(KubernetesMapOutput{})
}
