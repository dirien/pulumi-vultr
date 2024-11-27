// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-vultr/sdk/v2/go/vultr/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a Vultr instance.
//
// ## Example Usage
//
// Get the information for a instance by `label`:
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
//			_, err := vultr.LookupInstance(ctx, &vultr.LookupInstanceArgs{
//				Filters: []vultr.GetInstanceFilter{
//					{
//						Name: "label",
//						Values: []string{
//							"my-instance-label",
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
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInstanceResult
	err := ctx.Invoke("vultr:index/getInstance:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstance.
type LookupInstanceArgs struct {
	// Query parameters for finding instances.
	Filters []GetInstanceFilter `pulumi:"filters"`
}

// A collection of values returned by getInstance.
type LookupInstanceResult struct {
	// The server's allowed bandwidth usage in GB.
	AllowedBandwidth int `pulumi:"allowedBandwidth"`
	// The server's application ID.
	AppId   int    `pulumi:"appId"`
	Backups string `pulumi:"backups"`
	// The current configuration for backups
	BackupsSchedule map[string]string `pulumi:"backupsSchedule"`
	// The date the server was added to your Vultr account.
	DateCreated string `pulumi:"dateCreated"`
	// The description of the disk(s) on the server.
	Disk int `pulumi:"disk"`
	// Array of which features are enabled.
	Features []string            `pulumi:"features"`
	Filters  []GetInstanceFilter `pulumi:"filters"`
	// The ID of the firewall group applied to this server.
	FirewallGroupId string `pulumi:"firewallGroupId"`
	// The server's IPv4 gateway.
	GatewayV4 string `pulumi:"gatewayV4"`
	// The hostname assigned to the server.
	Hostname string `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Marketplace ID for this application.
	ImageId string `pulumi:"imageId"`
	// The server's internal IP address.
	InternalIp string `pulumi:"internalIp"`
	// The server's current KVM URL. This URL will change periodically. It is not advised to cache this value.
	Kvm string `pulumi:"kvm"`
	// The server's label.
	Label    string `pulumi:"label"`
	Location string `pulumi:"location"`
	// The server's main IP address.
	MainIp string `pulumi:"mainIp"`
	// The server's IPv4 netmask.
	NetmaskV4 string `pulumi:"netmaskV4"`
	// The operating system of the instance.
	Os string `pulumi:"os"`
	// The server's operating system ID.
	OsId int `pulumi:"osId"`
	// The server's plan ID.
	Plan string `pulumi:"plan"`
	// Whether the server is powered on or not.
	PowerStatus string `pulumi:"powerStatus"`
	// The amount of memory available on the instance in MB.
	Ram int `pulumi:"ram"`
	// The region ID of the server.
	Region string `pulumi:"region"`
	// A more detailed server status (none, locked, installingbooting, isomounting, ok).
	ServerStatus string `pulumi:"serverStatus"`
	// The status of the server's subscription.
	Status string `pulumi:"status"`
	// A list of tags applied to the instance.
	Tags []string `pulumi:"tags"`
	// The scheme used for the default user (linux servers only).
	UserScheme string `pulumi:"userScheme"`
	// The main IPv6 network address.
	V6MainIp string `pulumi:"v6MainIp"`
	// The IPv6 subnet.
	V6Network string `pulumi:"v6Network"`
	// The IPv6 network size in bits.
	V6NetworkSize int `pulumi:"v6NetworkSize"`
	// The number of virtual CPUs available on the server.
	VcpuCount int `pulumi:"vcpuCount"`
	// A list of VPC 2.0 IDs attached to the server.
	Vpc2Ids []string `pulumi:"vpc2Ids"`
	VpcIds  []string `pulumi:"vpcIds"`
}

func LookupInstanceOutput(ctx *pulumi.Context, args LookupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceResultOutput, error) {
			args := v.(LookupInstanceArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupInstanceResult
			secret, err := ctx.InvokePackageRaw("vultr:index/getInstance:getInstance", args, &rv, "", opts...)
			if err != nil {
				return LookupInstanceResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupInstanceResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupInstanceResultOutput), nil
			}
			return output, nil
		}).(LookupInstanceResultOutput)
}

// A collection of arguments for invoking getInstance.
type LookupInstanceOutputArgs struct {
	// Query parameters for finding instances.
	Filters GetInstanceFilterArrayInput `pulumi:"filters"`
}

func (LookupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getInstance.
type LookupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceResult)(nil)).Elem()
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutput() LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutputWithContext(ctx context.Context) LookupInstanceResultOutput {
	return o
}

// The server's allowed bandwidth usage in GB.
func (o LookupInstanceResultOutput) AllowedBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.AllowedBandwidth }).(pulumi.IntOutput)
}

// The server's application ID.
func (o LookupInstanceResultOutput) AppId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.AppId }).(pulumi.IntOutput)
}

func (o LookupInstanceResultOutput) Backups() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Backups }).(pulumi.StringOutput)
}

// The current configuration for backups
func (o LookupInstanceResultOutput) BackupsSchedule() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceResult) map[string]string { return v.BackupsSchedule }).(pulumi.StringMapOutput)
}

// The date the server was added to your Vultr account.
func (o LookupInstanceResultOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.DateCreated }).(pulumi.StringOutput)
}

// The description of the disk(s) on the server.
func (o LookupInstanceResultOutput) Disk() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.Disk }).(pulumi.IntOutput)
}

// Array of which features are enabled.
func (o LookupInstanceResultOutput) Features() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []string { return v.Features }).(pulumi.StringArrayOutput)
}

func (o LookupInstanceResultOutput) Filters() GetInstanceFilterArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []GetInstanceFilter { return v.Filters }).(GetInstanceFilterArrayOutput)
}

// The ID of the firewall group applied to this server.
func (o LookupInstanceResultOutput) FirewallGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.FirewallGroupId }).(pulumi.StringOutput)
}

// The server's IPv4 gateway.
func (o LookupInstanceResultOutput) GatewayV4() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.GatewayV4 }).(pulumi.StringOutput)
}

// The hostname assigned to the server.
func (o LookupInstanceResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Marketplace ID for this application.
func (o LookupInstanceResultOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.ImageId }).(pulumi.StringOutput)
}

// The server's internal IP address.
func (o LookupInstanceResultOutput) InternalIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.InternalIp }).(pulumi.StringOutput)
}

// The server's current KVM URL. This URL will change periodically. It is not advised to cache this value.
func (o LookupInstanceResultOutput) Kvm() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Kvm }).(pulumi.StringOutput)
}

// The server's label.
func (o LookupInstanceResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Label }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The server's main IP address.
func (o LookupInstanceResultOutput) MainIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.MainIp }).(pulumi.StringOutput)
}

// The server's IPv4 netmask.
func (o LookupInstanceResultOutput) NetmaskV4() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.NetmaskV4 }).(pulumi.StringOutput)
}

// The operating system of the instance.
func (o LookupInstanceResultOutput) Os() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Os }).(pulumi.StringOutput)
}

// The server's operating system ID.
func (o LookupInstanceResultOutput) OsId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.OsId }).(pulumi.IntOutput)
}

// The server's plan ID.
func (o LookupInstanceResultOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Plan }).(pulumi.StringOutput)
}

// Whether the server is powered on or not.
func (o LookupInstanceResultOutput) PowerStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.PowerStatus }).(pulumi.StringOutput)
}

// The amount of memory available on the instance in MB.
func (o LookupInstanceResultOutput) Ram() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.Ram }).(pulumi.IntOutput)
}

// The region ID of the server.
func (o LookupInstanceResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Region }).(pulumi.StringOutput)
}

// A more detailed server status (none, locked, installingbooting, isomounting, ok).
func (o LookupInstanceResultOutput) ServerStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.ServerStatus }).(pulumi.StringOutput)
}

// The status of the server's subscription.
func (o LookupInstanceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Status }).(pulumi.StringOutput)
}

// A list of tags applied to the instance.
func (o LookupInstanceResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The scheme used for the default user (linux servers only).
func (o LookupInstanceResultOutput) UserScheme() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.UserScheme }).(pulumi.StringOutput)
}

// The main IPv6 network address.
func (o LookupInstanceResultOutput) V6MainIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.V6MainIp }).(pulumi.StringOutput)
}

// The IPv6 subnet.
func (o LookupInstanceResultOutput) V6Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.V6Network }).(pulumi.StringOutput)
}

// The IPv6 network size in bits.
func (o LookupInstanceResultOutput) V6NetworkSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.V6NetworkSize }).(pulumi.IntOutput)
}

// The number of virtual CPUs available on the server.
func (o LookupInstanceResultOutput) VcpuCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceResult) int { return v.VcpuCount }).(pulumi.IntOutput)
}

// A list of VPC 2.0 IDs attached to the server.
func (o LookupInstanceResultOutput) Vpc2Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []string { return v.Vpc2Ids }).(pulumi.StringArrayOutput)
}

func (o LookupInstanceResultOutput) VpcIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []string { return v.VpcIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}
