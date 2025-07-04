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

// Create and update a Vultr container registry.
//
// ## Example Usage
//
// Create a new container registry:
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
//			_, err := vultr.NewContainerRegistry(ctx, "vcr1", &vultr.ContainerRegistryArgs{
//				Plan:   pulumi.String("start_up"),
//				Public: pulumi.Bool(false),
//				Region: pulumi.String("sjc"),
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
// The `name` for container registries must be all lowercase and only contain alphanumeric characters.
type ContainerRegistry struct {
	pulumi.CustomResourceState

	// The URN of the container registry.
	ContainerRegistryURN pulumi.StringOutput `pulumi:"containerRegistryURN"`
	// A date-time of when the root user was created.
	DateCreated pulumi.StringOutput `pulumi:"dateCreated"`
	// The name for your container registry.  Must be lowercase and only alphanumeric characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// The billing plan for the container registry. [See available plans](https://www.vultr.com/api/#tag/Container-Registry/operation/list-registry-plans)
	Plan pulumi.StringOutput `pulumi:"plan"`
	// Boolean indicating if the container registry should be created with public visibility or if it should require credentials.
	Public pulumi.BoolOutput `pulumi:"public"`
	// The region where your container registry will be deployed. [See available regions](https://www.vultr.com/api/#tag/Container-Registry/operation/list-registry-regions)
	Region pulumi.StringOutput `pulumi:"region"`
	// The user associated with the container registry.
	RootUser pulumi.StringMapOutput `pulumi:"rootUser"`
	// A listing of current storage usage relevant to the container registry.
	Storage pulumi.StringMapOutput `pulumi:"storage"`
}

// NewContainerRegistry registers a new resource with the given unique name, arguments, and options.
func NewContainerRegistry(ctx *pulumi.Context,
	name string, args *ContainerRegistryArgs, opts ...pulumi.ResourceOption) (*ContainerRegistry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Plan == nil {
		return nil, errors.New("invalid value for required argument 'Plan'")
	}
	if args.Public == nil {
		return nil, errors.New("invalid value for required argument 'Public'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContainerRegistry
	err := ctx.RegisterResource("vultr:index/containerRegistry:ContainerRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerRegistry gets an existing ContainerRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerRegistryState, opts ...pulumi.ResourceOption) (*ContainerRegistry, error) {
	var resource ContainerRegistry
	err := ctx.ReadResource("vultr:index/containerRegistry:ContainerRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerRegistry resources.
type containerRegistryState struct {
	// The URN of the container registry.
	ContainerRegistryURN *string `pulumi:"containerRegistryURN"`
	// A date-time of when the root user was created.
	DateCreated *string `pulumi:"dateCreated"`
	// The name for your container registry.  Must be lowercase and only alphanumeric characters.
	Name *string `pulumi:"name"`
	// The billing plan for the container registry. [See available plans](https://www.vultr.com/api/#tag/Container-Registry/operation/list-registry-plans)
	Plan *string `pulumi:"plan"`
	// Boolean indicating if the container registry should be created with public visibility or if it should require credentials.
	Public *bool `pulumi:"public"`
	// The region where your container registry will be deployed. [See available regions](https://www.vultr.com/api/#tag/Container-Registry/operation/list-registry-regions)
	Region *string `pulumi:"region"`
	// The user associated with the container registry.
	RootUser map[string]string `pulumi:"rootUser"`
	// A listing of current storage usage relevant to the container registry.
	Storage map[string]string `pulumi:"storage"`
}

type ContainerRegistryState struct {
	// The URN of the container registry.
	ContainerRegistryURN pulumi.StringPtrInput
	// A date-time of when the root user was created.
	DateCreated pulumi.StringPtrInput
	// The name for your container registry.  Must be lowercase and only alphanumeric characters.
	Name pulumi.StringPtrInput
	// The billing plan for the container registry. [See available plans](https://www.vultr.com/api/#tag/Container-Registry/operation/list-registry-plans)
	Plan pulumi.StringPtrInput
	// Boolean indicating if the container registry should be created with public visibility or if it should require credentials.
	Public pulumi.BoolPtrInput
	// The region where your container registry will be deployed. [See available regions](https://www.vultr.com/api/#tag/Container-Registry/operation/list-registry-regions)
	Region pulumi.StringPtrInput
	// The user associated with the container registry.
	RootUser pulumi.StringMapInput
	// A listing of current storage usage relevant to the container registry.
	Storage pulumi.StringMapInput
}

func (ContainerRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryState)(nil)).Elem()
}

type containerRegistryArgs struct {
	// The name for your container registry.  Must be lowercase and only alphanumeric characters.
	Name *string `pulumi:"name"`
	// The billing plan for the container registry. [See available plans](https://www.vultr.com/api/#tag/Container-Registry/operation/list-registry-plans)
	Plan string `pulumi:"plan"`
	// Boolean indicating if the container registry should be created with public visibility or if it should require credentials.
	Public bool `pulumi:"public"`
	// The region where your container registry will be deployed. [See available regions](https://www.vultr.com/api/#tag/Container-Registry/operation/list-registry-regions)
	Region string `pulumi:"region"`
}

// The set of arguments for constructing a ContainerRegistry resource.
type ContainerRegistryArgs struct {
	// The name for your container registry.  Must be lowercase and only alphanumeric characters.
	Name pulumi.StringPtrInput
	// The billing plan for the container registry. [See available plans](https://www.vultr.com/api/#tag/Container-Registry/operation/list-registry-plans)
	Plan pulumi.StringInput
	// Boolean indicating if the container registry should be created with public visibility or if it should require credentials.
	Public pulumi.BoolInput
	// The region where your container registry will be deployed. [See available regions](https://www.vultr.com/api/#tag/Container-Registry/operation/list-registry-regions)
	Region pulumi.StringInput
}

func (ContainerRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryArgs)(nil)).Elem()
}

type ContainerRegistryInput interface {
	pulumi.Input

	ToContainerRegistryOutput() ContainerRegistryOutput
	ToContainerRegistryOutputWithContext(ctx context.Context) ContainerRegistryOutput
}

func (*ContainerRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistry)(nil)).Elem()
}

func (i *ContainerRegistry) ToContainerRegistryOutput() ContainerRegistryOutput {
	return i.ToContainerRegistryOutputWithContext(context.Background())
}

func (i *ContainerRegistry) ToContainerRegistryOutputWithContext(ctx context.Context) ContainerRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryOutput)
}

// ContainerRegistryArrayInput is an input type that accepts ContainerRegistryArray and ContainerRegistryArrayOutput values.
// You can construct a concrete instance of `ContainerRegistryArrayInput` via:
//
//	ContainerRegistryArray{ ContainerRegistryArgs{...} }
type ContainerRegistryArrayInput interface {
	pulumi.Input

	ToContainerRegistryArrayOutput() ContainerRegistryArrayOutput
	ToContainerRegistryArrayOutputWithContext(context.Context) ContainerRegistryArrayOutput
}

type ContainerRegistryArray []ContainerRegistryInput

func (ContainerRegistryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRegistry)(nil)).Elem()
}

func (i ContainerRegistryArray) ToContainerRegistryArrayOutput() ContainerRegistryArrayOutput {
	return i.ToContainerRegistryArrayOutputWithContext(context.Background())
}

func (i ContainerRegistryArray) ToContainerRegistryArrayOutputWithContext(ctx context.Context) ContainerRegistryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryArrayOutput)
}

// ContainerRegistryMapInput is an input type that accepts ContainerRegistryMap and ContainerRegistryMapOutput values.
// You can construct a concrete instance of `ContainerRegistryMapInput` via:
//
//	ContainerRegistryMap{ "key": ContainerRegistryArgs{...} }
type ContainerRegistryMapInput interface {
	pulumi.Input

	ToContainerRegistryMapOutput() ContainerRegistryMapOutput
	ToContainerRegistryMapOutputWithContext(context.Context) ContainerRegistryMapOutput
}

type ContainerRegistryMap map[string]ContainerRegistryInput

func (ContainerRegistryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRegistry)(nil)).Elem()
}

func (i ContainerRegistryMap) ToContainerRegistryMapOutput() ContainerRegistryMapOutput {
	return i.ToContainerRegistryMapOutputWithContext(context.Background())
}

func (i ContainerRegistryMap) ToContainerRegistryMapOutputWithContext(ctx context.Context) ContainerRegistryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryMapOutput)
}

type ContainerRegistryOutput struct{ *pulumi.OutputState }

func (ContainerRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistry)(nil)).Elem()
}

func (o ContainerRegistryOutput) ToContainerRegistryOutput() ContainerRegistryOutput {
	return o
}

func (o ContainerRegistryOutput) ToContainerRegistryOutputWithContext(ctx context.Context) ContainerRegistryOutput {
	return o
}

// The URN of the container registry.
func (o ContainerRegistryOutput) ContainerRegistryURN() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.ContainerRegistryURN }).(pulumi.StringOutput)
}

// A date-time of when the root user was created.
func (o ContainerRegistryOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.DateCreated }).(pulumi.StringOutput)
}

// The name for your container registry.  Must be lowercase and only alphanumeric characters.
func (o ContainerRegistryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The billing plan for the container registry. [See available plans](https://www.vultr.com/api/#tag/Container-Registry/operation/list-registry-plans)
func (o ContainerRegistryOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.Plan }).(pulumi.StringOutput)
}

// Boolean indicating if the container registry should be created with public visibility or if it should require credentials.
func (o ContainerRegistryOutput) Public() pulumi.BoolOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.BoolOutput { return v.Public }).(pulumi.BoolOutput)
}

// The region where your container registry will be deployed. [See available regions](https://www.vultr.com/api/#tag/Container-Registry/operation/list-registry-regions)
func (o ContainerRegistryOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The user associated with the container registry.
func (o ContainerRegistryOutput) RootUser() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringMapOutput { return v.RootUser }).(pulumi.StringMapOutput)
}

// A listing of current storage usage relevant to the container registry.
func (o ContainerRegistryOutput) Storage() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContainerRegistry) pulumi.StringMapOutput { return v.Storage }).(pulumi.StringMapOutput)
}

type ContainerRegistryArrayOutput struct{ *pulumi.OutputState }

func (ContainerRegistryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRegistry)(nil)).Elem()
}

func (o ContainerRegistryArrayOutput) ToContainerRegistryArrayOutput() ContainerRegistryArrayOutput {
	return o
}

func (o ContainerRegistryArrayOutput) ToContainerRegistryArrayOutputWithContext(ctx context.Context) ContainerRegistryArrayOutput {
	return o
}

func (o ContainerRegistryArrayOutput) Index(i pulumi.IntInput) ContainerRegistryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerRegistry {
		return vs[0].([]*ContainerRegistry)[vs[1].(int)]
	}).(ContainerRegistryOutput)
}

type ContainerRegistryMapOutput struct{ *pulumi.OutputState }

func (ContainerRegistryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRegistry)(nil)).Elem()
}

func (o ContainerRegistryMapOutput) ToContainerRegistryMapOutput() ContainerRegistryMapOutput {
	return o
}

func (o ContainerRegistryMapOutput) ToContainerRegistryMapOutputWithContext(ctx context.Context) ContainerRegistryMapOutput {
	return o
}

func (o ContainerRegistryMapOutput) MapIndex(k pulumi.StringInput) ContainerRegistryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerRegistry {
		return vs[0].(map[string]*ContainerRegistry)[vs[1].(string)]
	}).(ContainerRegistryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryInput)(nil)).Elem(), &ContainerRegistry{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryArrayInput)(nil)).Elem(), ContainerRegistryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryMapInput)(nil)).Elem(), ContainerRegistryMap{})
	pulumi.RegisterOutputType(ContainerRegistryOutput{})
	pulumi.RegisterOutputType(ContainerRegistryArrayOutput{})
	pulumi.RegisterOutputType(ContainerRegistryMapOutput{})
}
