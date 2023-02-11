// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Vultr DNS Domain resource. This can be used to create, read, modify, and delete DNS Domains.
//
// ## Example Usage
//
// # Create a new DNS Domain
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
//			_, err := vultr.NewDnsDomain(ctx, "myDomain", &vultr.DnsDomainArgs{
//				Domain: pulumi.String("domain.com"),
//				Ip:     pulumi.String("66.42.94.227"),
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
// DNS Domains can be imported using the Dns Domain `domain`, e.g.
//
// ```sh
//
//	$ pulumi import vultr:index/dnsDomain:DnsDomain name domain.com
//
// ```
type DnsDomain struct {
	pulumi.CustomResourceState

	// The date the domain was added to your account.
	DateCreated pulumi.StringOutput `pulumi:"dateCreated"`
	// The Domain's DNSSEC status. Valid options are `enabled` or `disabled`. Note `disabled` is default
	DnsSec pulumi.StringPtrOutput `pulumi:"dnsSec"`
	// Name of domain.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Instance IP you want associated to domain. If omitted this will create a domain with no records.
	Ip pulumi.StringPtrOutput `pulumi:"ip"`
}

// NewDnsDomain registers a new resource with the given unique name, arguments, and options.
func NewDnsDomain(ctx *pulumi.Context,
	name string, args *DnsDomainArgs, opts ...pulumi.ResourceOption) (*DnsDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DnsDomain
	err := ctx.RegisterResource("vultr:index/dnsDomain:DnsDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsDomain gets an existing DnsDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsDomainState, opts ...pulumi.ResourceOption) (*DnsDomain, error) {
	var resource DnsDomain
	err := ctx.ReadResource("vultr:index/dnsDomain:DnsDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsDomain resources.
type dnsDomainState struct {
	// The date the domain was added to your account.
	DateCreated *string `pulumi:"dateCreated"`
	// The Domain's DNSSEC status. Valid options are `enabled` or `disabled`. Note `disabled` is default
	DnsSec *string `pulumi:"dnsSec"`
	// Name of domain.
	Domain *string `pulumi:"domain"`
	// Instance IP you want associated to domain. If omitted this will create a domain with no records.
	Ip *string `pulumi:"ip"`
}

type DnsDomainState struct {
	// The date the domain was added to your account.
	DateCreated pulumi.StringPtrInput
	// The Domain's DNSSEC status. Valid options are `enabled` or `disabled`. Note `disabled` is default
	DnsSec pulumi.StringPtrInput
	// Name of domain.
	Domain pulumi.StringPtrInput
	// Instance IP you want associated to domain. If omitted this will create a domain with no records.
	Ip pulumi.StringPtrInput
}

func (DnsDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsDomainState)(nil)).Elem()
}

type dnsDomainArgs struct {
	// The Domain's DNSSEC status. Valid options are `enabled` or `disabled`. Note `disabled` is default
	DnsSec *string `pulumi:"dnsSec"`
	// Name of domain.
	Domain string `pulumi:"domain"`
	// Instance IP you want associated to domain. If omitted this will create a domain with no records.
	Ip *string `pulumi:"ip"`
}

// The set of arguments for constructing a DnsDomain resource.
type DnsDomainArgs struct {
	// The Domain's DNSSEC status. Valid options are `enabled` or `disabled`. Note `disabled` is default
	DnsSec pulumi.StringPtrInput
	// Name of domain.
	Domain pulumi.StringInput
	// Instance IP you want associated to domain. If omitted this will create a domain with no records.
	Ip pulumi.StringPtrInput
}

func (DnsDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsDomainArgs)(nil)).Elem()
}

type DnsDomainInput interface {
	pulumi.Input

	ToDnsDomainOutput() DnsDomainOutput
	ToDnsDomainOutputWithContext(ctx context.Context) DnsDomainOutput
}

func (*DnsDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsDomain)(nil)).Elem()
}

func (i *DnsDomain) ToDnsDomainOutput() DnsDomainOutput {
	return i.ToDnsDomainOutputWithContext(context.Background())
}

func (i *DnsDomain) ToDnsDomainOutputWithContext(ctx context.Context) DnsDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsDomainOutput)
}

// DnsDomainArrayInput is an input type that accepts DnsDomainArray and DnsDomainArrayOutput values.
// You can construct a concrete instance of `DnsDomainArrayInput` via:
//
//	DnsDomainArray{ DnsDomainArgs{...} }
type DnsDomainArrayInput interface {
	pulumi.Input

	ToDnsDomainArrayOutput() DnsDomainArrayOutput
	ToDnsDomainArrayOutputWithContext(context.Context) DnsDomainArrayOutput
}

type DnsDomainArray []DnsDomainInput

func (DnsDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsDomain)(nil)).Elem()
}

func (i DnsDomainArray) ToDnsDomainArrayOutput() DnsDomainArrayOutput {
	return i.ToDnsDomainArrayOutputWithContext(context.Background())
}

func (i DnsDomainArray) ToDnsDomainArrayOutputWithContext(ctx context.Context) DnsDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsDomainArrayOutput)
}

// DnsDomainMapInput is an input type that accepts DnsDomainMap and DnsDomainMapOutput values.
// You can construct a concrete instance of `DnsDomainMapInput` via:
//
//	DnsDomainMap{ "key": DnsDomainArgs{...} }
type DnsDomainMapInput interface {
	pulumi.Input

	ToDnsDomainMapOutput() DnsDomainMapOutput
	ToDnsDomainMapOutputWithContext(context.Context) DnsDomainMapOutput
}

type DnsDomainMap map[string]DnsDomainInput

func (DnsDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsDomain)(nil)).Elem()
}

func (i DnsDomainMap) ToDnsDomainMapOutput() DnsDomainMapOutput {
	return i.ToDnsDomainMapOutputWithContext(context.Background())
}

func (i DnsDomainMap) ToDnsDomainMapOutputWithContext(ctx context.Context) DnsDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsDomainMapOutput)
}

type DnsDomainOutput struct{ *pulumi.OutputState }

func (DnsDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsDomain)(nil)).Elem()
}

func (o DnsDomainOutput) ToDnsDomainOutput() DnsDomainOutput {
	return o
}

func (o DnsDomainOutput) ToDnsDomainOutputWithContext(ctx context.Context) DnsDomainOutput {
	return o
}

// The date the domain was added to your account.
func (o DnsDomainOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsDomain) pulumi.StringOutput { return v.DateCreated }).(pulumi.StringOutput)
}

// The Domain's DNSSEC status. Valid options are `enabled` or `disabled`. Note `disabled` is default
func (o DnsDomainOutput) DnsSec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsDomain) pulumi.StringPtrOutput { return v.DnsSec }).(pulumi.StringPtrOutput)
}

// Name of domain.
func (o DnsDomainOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsDomain) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Instance IP you want associated to domain. If omitted this will create a domain with no records.
func (o DnsDomainOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsDomain) pulumi.StringPtrOutput { return v.Ip }).(pulumi.StringPtrOutput)
}

type DnsDomainArrayOutput struct{ *pulumi.OutputState }

func (DnsDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsDomain)(nil)).Elem()
}

func (o DnsDomainArrayOutput) ToDnsDomainArrayOutput() DnsDomainArrayOutput {
	return o
}

func (o DnsDomainArrayOutput) ToDnsDomainArrayOutputWithContext(ctx context.Context) DnsDomainArrayOutput {
	return o
}

func (o DnsDomainArrayOutput) Index(i pulumi.IntInput) DnsDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsDomain {
		return vs[0].([]*DnsDomain)[vs[1].(int)]
	}).(DnsDomainOutput)
}

type DnsDomainMapOutput struct{ *pulumi.OutputState }

func (DnsDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsDomain)(nil)).Elem()
}

func (o DnsDomainMapOutput) ToDnsDomainMapOutput() DnsDomainMapOutput {
	return o
}

func (o DnsDomainMapOutput) ToDnsDomainMapOutputWithContext(ctx context.Context) DnsDomainMapOutput {
	return o
}

func (o DnsDomainMapOutput) MapIndex(k pulumi.StringInput) DnsDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsDomain {
		return vs[0].(map[string]*DnsDomain)[vs[1].(string)]
	}).(DnsDomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsDomainInput)(nil)).Elem(), &DnsDomain{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsDomainArrayInput)(nil)).Elem(), DnsDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsDomainMapInput)(nil)).Elem(), DnsDomainMap{})
	pulumi.RegisterOutputType(DnsDomainOutput{})
	pulumi.RegisterOutputType(DnsDomainArrayOutput{})
	pulumi.RegisterOutputType(DnsDomainMapOutput{})
}
