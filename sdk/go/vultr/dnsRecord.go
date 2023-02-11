// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Vultr DNS Record resource. This can be used to create, read, modify, and delete DNS Records.
//
// ## Example Usage
//
// # Create a new DNS Record
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
//			myDomain, err := vultr.NewDnsDomain(ctx, "myDomain", &vultr.DnsDomainArgs{
//				Domain: pulumi.String("domain.com"),
//				Ip:     pulumi.String("66.42.94.227"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vultr.NewDnsRecord(ctx, "myRecord", &vultr.DnsRecordArgs{
//				Data:   pulumi.String("66.42.94.227"),
//				Domain: myDomain.ID(),
//				Type:   pulumi.String("A"),
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
// DNS Records can be imported using the Dns Domain `domain` and DNS Record `ID` e.g.
//
// ```sh
//
//	$ pulumi import vultr:index/dnsRecord:DnsRecord rec domain.com,1a0019bd-7645-4310-81bd-03bc5906940f
//
// ```
type DnsRecord struct {
	pulumi.CustomResourceState

	// IP Address of the instance the domain is associated with.
	Data pulumi.StringOutput `pulumi:"data"`
	// Name of the DNS Domain this record will belong to.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Name (subdomain) for this record.
	Name pulumi.StringOutput `pulumi:"name"`
	// Priority of this record (only required for MX and SRV).
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The time to live of this record.
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
	// Type of record.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDnsRecord registers a new resource with the given unique name, arguments, and options.
func NewDnsRecord(ctx *pulumi.Context,
	name string, args *DnsRecordArgs, opts ...pulumi.ResourceOption) (*DnsRecord, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Data == nil {
		return nil, errors.New("invalid value for required argument 'Data'")
	}
	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DnsRecord
	err := ctx.RegisterResource("vultr:index/dnsRecord:DnsRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsRecord gets an existing DnsRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsRecordState, opts ...pulumi.ResourceOption) (*DnsRecord, error) {
	var resource DnsRecord
	err := ctx.ReadResource("vultr:index/dnsRecord:DnsRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsRecord resources.
type dnsRecordState struct {
	// IP Address of the instance the domain is associated with.
	Data *string `pulumi:"data"`
	// Name of the DNS Domain this record will belong to.
	Domain *string `pulumi:"domain"`
	// Name (subdomain) for this record.
	Name *string `pulumi:"name"`
	// Priority of this record (only required for MX and SRV).
	Priority *int `pulumi:"priority"`
	// The time to live of this record.
	Ttl *int `pulumi:"ttl"`
	// Type of record.
	Type *string `pulumi:"type"`
}

type DnsRecordState struct {
	// IP Address of the instance the domain is associated with.
	Data pulumi.StringPtrInput
	// Name of the DNS Domain this record will belong to.
	Domain pulumi.StringPtrInput
	// Name (subdomain) for this record.
	Name pulumi.StringPtrInput
	// Priority of this record (only required for MX and SRV).
	Priority pulumi.IntPtrInput
	// The time to live of this record.
	Ttl pulumi.IntPtrInput
	// Type of record.
	Type pulumi.StringPtrInput
}

func (DnsRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsRecordState)(nil)).Elem()
}

type dnsRecordArgs struct {
	// IP Address of the instance the domain is associated with.
	Data string `pulumi:"data"`
	// Name of the DNS Domain this record will belong to.
	Domain string `pulumi:"domain"`
	// Name (subdomain) for this record.
	Name *string `pulumi:"name"`
	// Priority of this record (only required for MX and SRV).
	Priority *int `pulumi:"priority"`
	// The time to live of this record.
	Ttl *int `pulumi:"ttl"`
	// Type of record.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a DnsRecord resource.
type DnsRecordArgs struct {
	// IP Address of the instance the domain is associated with.
	Data pulumi.StringInput
	// Name of the DNS Domain this record will belong to.
	Domain pulumi.StringInput
	// Name (subdomain) for this record.
	Name pulumi.StringPtrInput
	// Priority of this record (only required for MX and SRV).
	Priority pulumi.IntPtrInput
	// The time to live of this record.
	Ttl pulumi.IntPtrInput
	// Type of record.
	Type pulumi.StringInput
}

func (DnsRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsRecordArgs)(nil)).Elem()
}

type DnsRecordInput interface {
	pulumi.Input

	ToDnsRecordOutput() DnsRecordOutput
	ToDnsRecordOutputWithContext(ctx context.Context) DnsRecordOutput
}

func (*DnsRecord) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsRecord)(nil)).Elem()
}

func (i *DnsRecord) ToDnsRecordOutput() DnsRecordOutput {
	return i.ToDnsRecordOutputWithContext(context.Background())
}

func (i *DnsRecord) ToDnsRecordOutputWithContext(ctx context.Context) DnsRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRecordOutput)
}

// DnsRecordArrayInput is an input type that accepts DnsRecordArray and DnsRecordArrayOutput values.
// You can construct a concrete instance of `DnsRecordArrayInput` via:
//
//	DnsRecordArray{ DnsRecordArgs{...} }
type DnsRecordArrayInput interface {
	pulumi.Input

	ToDnsRecordArrayOutput() DnsRecordArrayOutput
	ToDnsRecordArrayOutputWithContext(context.Context) DnsRecordArrayOutput
}

type DnsRecordArray []DnsRecordInput

func (DnsRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsRecord)(nil)).Elem()
}

func (i DnsRecordArray) ToDnsRecordArrayOutput() DnsRecordArrayOutput {
	return i.ToDnsRecordArrayOutputWithContext(context.Background())
}

func (i DnsRecordArray) ToDnsRecordArrayOutputWithContext(ctx context.Context) DnsRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRecordArrayOutput)
}

// DnsRecordMapInput is an input type that accepts DnsRecordMap and DnsRecordMapOutput values.
// You can construct a concrete instance of `DnsRecordMapInput` via:
//
//	DnsRecordMap{ "key": DnsRecordArgs{...} }
type DnsRecordMapInput interface {
	pulumi.Input

	ToDnsRecordMapOutput() DnsRecordMapOutput
	ToDnsRecordMapOutputWithContext(context.Context) DnsRecordMapOutput
}

type DnsRecordMap map[string]DnsRecordInput

func (DnsRecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsRecord)(nil)).Elem()
}

func (i DnsRecordMap) ToDnsRecordMapOutput() DnsRecordMapOutput {
	return i.ToDnsRecordMapOutputWithContext(context.Background())
}

func (i DnsRecordMap) ToDnsRecordMapOutputWithContext(ctx context.Context) DnsRecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsRecordMapOutput)
}

type DnsRecordOutput struct{ *pulumi.OutputState }

func (DnsRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsRecord)(nil)).Elem()
}

func (o DnsRecordOutput) ToDnsRecordOutput() DnsRecordOutput {
	return o
}

func (o DnsRecordOutput) ToDnsRecordOutputWithContext(ctx context.Context) DnsRecordOutput {
	return o
}

// IP Address of the instance the domain is associated with.
func (o DnsRecordOutput) Data() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringOutput { return v.Data }).(pulumi.StringOutput)
}

// Name of the DNS Domain this record will belong to.
func (o DnsRecordOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Name (subdomain) for this record.
func (o DnsRecordOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Priority of this record (only required for MX and SRV).
func (o DnsRecordOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// The time to live of this record.
func (o DnsRecordOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.IntPtrOutput { return v.Ttl }).(pulumi.IntPtrOutput)
}

// Type of record.
func (o DnsRecordOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsRecord) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type DnsRecordArrayOutput struct{ *pulumi.OutputState }

func (DnsRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsRecord)(nil)).Elem()
}

func (o DnsRecordArrayOutput) ToDnsRecordArrayOutput() DnsRecordArrayOutput {
	return o
}

func (o DnsRecordArrayOutput) ToDnsRecordArrayOutputWithContext(ctx context.Context) DnsRecordArrayOutput {
	return o
}

func (o DnsRecordArrayOutput) Index(i pulumi.IntInput) DnsRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsRecord {
		return vs[0].([]*DnsRecord)[vs[1].(int)]
	}).(DnsRecordOutput)
}

type DnsRecordMapOutput struct{ *pulumi.OutputState }

func (DnsRecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsRecord)(nil)).Elem()
}

func (o DnsRecordMapOutput) ToDnsRecordMapOutput() DnsRecordMapOutput {
	return o
}

func (o DnsRecordMapOutput) ToDnsRecordMapOutputWithContext(ctx context.Context) DnsRecordMapOutput {
	return o
}

func (o DnsRecordMapOutput) MapIndex(k pulumi.StringInput) DnsRecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsRecord {
		return vs[0].(map[string]*DnsRecord)[vs[1].(string)]
	}).(DnsRecordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRecordInput)(nil)).Elem(), &DnsRecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRecordArrayInput)(nil)).Elem(), DnsRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsRecordMapInput)(nil)).Elem(), DnsRecordMap{})
	pulumi.RegisterOutputType(DnsRecordOutput{})
	pulumi.RegisterOutputType(DnsRecordArrayOutput{})
	pulumi.RegisterOutputType(DnsRecordMapOutput{})
}
