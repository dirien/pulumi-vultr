// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vultr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Vultr bare metal server resource. This can be used to create, read, modify, and delete bare metal servers on your Vultr account.
//
// ## Example Usage
//
// Create a new bare metal server:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vultr/sdk/go/vultr"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vultr.NewBareMetalServer(ctx, "myServer", &vultr.BareMetalServerArgs{
//				OsId:   pulumi.Int(270),
//				Plan:   pulumi.String("vbm-4c-32gb"),
//				Region: pulumi.String("ewr"),
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
// Create a new bare metal server with options:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-vultr/sdk/go/vultr"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vultr.NewBareMetalServer(ctx, "myServer", &vultr.BareMetalServerArgs{
//				ActivationEmail: pulumi.Bool(false),
//				EnableIpv6:      pulumi.Bool(true),
//				Hostname:        pulumi.String("my-server-hostname"),
//				Label:           pulumi.String("my-server-label"),
//				OsId:            pulumi.Int(270),
//				Plan:            pulumi.String("vbm-4c-32gb"),
//				Region:          pulumi.String("ewr"),
//				Tags: pulumi.StringArray{
//					pulumi.String("my-server-tag"),
//				},
//				UserData: pulumi.String("this is my user data"),
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
// Bare Metal Servers can be imported using the server `ID`, e.g.
//
// ```sh
//
//	$ pulumi import vultr:index/bareMetalServer:BareMetalServer my_server b6a859c5-b299-49dd-8888-b1abbc517d08
//
// ```
type BareMetalServer struct {
	pulumi.CustomResourceState

	// Whether an activation email will be sent when the server is ready.
	ActivationEmail pulumi.BoolPtrOutput `pulumi:"activationEmail"`
	// The ID of the Vultr application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications)
	AppId pulumi.IntOutput `pulumi:"appId"`
	// The number of CPUs available on the server.
	CpuCount pulumi.IntOutput `pulumi:"cpuCount"`
	// The date the server was added to your Vultr account.
	DateCreated pulumi.StringOutput `pulumi:"dateCreated"`
	// The server's default password.
	DefaultPassword pulumi.StringOutput `pulumi:"defaultPassword"`
	// The description of the disk(s) on the server.
	Disk pulumi.StringOutput `pulumi:"disk"`
	// Whether the server has IPv6 networking activated.
	EnableIpv6 pulumi.BoolPtrOutput `pulumi:"enableIpv6"`
	// The server's IPv4 gateway.
	GatewayV4 pulumi.StringOutput `pulumi:"gatewayV4"`
	// The hostname to assign to the server.
	Hostname pulumi.StringPtrOutput `pulumi:"hostname"`
	// The ID of the Vultr marketplace application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications) Note marketplace applications are denoted by type: `marketplace` and you must use the `imageId` not the id.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// A label for the server.
	Label pulumi.StringPtrOutput `pulumi:"label"`
	// The MAC address associated with the server.
	MacAddress pulumi.IntOutput `pulumi:"macAddress"`
	// The server's main IP address.
	MainIp pulumi.StringOutput `pulumi:"mainIp"`
	// The server's IPv4 netmask.
	NetmaskV4 pulumi.StringOutput `pulumi:"netmaskV4"`
	// The string description of the operating system installed on the server.
	Os pulumi.StringOutput `pulumi:"os"`
	// The ID of the operating system to be installed on the server. [See List OS](https://www.vultr.com/api/#operation/list-os)
	OsId pulumi.IntOutput `pulumi:"osId"`
	// The ID of the plan that you want the server to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)
	Plan pulumi.StringOutput `pulumi:"plan"`
	// The amount of memory available on the server in MB.
	Ram pulumi.StringOutput `pulumi:"ram"`
	// The ID of the region that the server is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
	Region pulumi.StringOutput `pulumi:"region"`
	// The ID of the floating IP to use as the main IP of this server. [See Reserved IPs](https://www.vultr.com/api/#operation/list-reserved-ips)
	ReservedIpv4 pulumi.StringOutput `pulumi:"reservedIpv4"`
	// The ID of the startup script you want added to the server.
	ScriptId pulumi.StringPtrOutput `pulumi:"scriptId"`
	// The ID of the Vultr snapshot that the server will restore for the initial installation. [See List Snapshots](https://www.vultr.com/api/#operation/list-snapshots)
	SnapshotId pulumi.StringPtrOutput `pulumi:"snapshotId"`
	// A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).
	SshKeyIds pulumi.StringArrayOutput `pulumi:"sshKeyIds"`
	// The status of the server's subscription.
	Status pulumi.StringOutput `pulumi:"status"`
	// A list of tags to apply to the servier.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.
	UserData pulumi.StringOutput `pulumi:"userData"`
	// The main IPv6 network address.
	V6MainIp pulumi.StringOutput `pulumi:"v6MainIp"`
	// The IPv6 subnet.
	V6Network pulumi.StringOutput `pulumi:"v6Network"`
	// The IPv6 network size in bits.
	V6NetworkSize pulumi.IntOutput `pulumi:"v6NetworkSize"`
}

// NewBareMetalServer registers a new resource with the given unique name, arguments, and options.
func NewBareMetalServer(ctx *pulumi.Context,
	name string, args *BareMetalServerArgs, opts ...pulumi.ResourceOption) (*BareMetalServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Plan == nil {
		return nil, errors.New("invalid value for required argument 'Plan'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"defaultPassword",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource BareMetalServer
	err := ctx.RegisterResource("vultr:index/bareMetalServer:BareMetalServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBareMetalServer gets an existing BareMetalServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBareMetalServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BareMetalServerState, opts ...pulumi.ResourceOption) (*BareMetalServer, error) {
	var resource BareMetalServer
	err := ctx.ReadResource("vultr:index/bareMetalServer:BareMetalServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BareMetalServer resources.
type bareMetalServerState struct {
	// Whether an activation email will be sent when the server is ready.
	ActivationEmail *bool `pulumi:"activationEmail"`
	// The ID of the Vultr application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications)
	AppId *int `pulumi:"appId"`
	// The number of CPUs available on the server.
	CpuCount *int `pulumi:"cpuCount"`
	// The date the server was added to your Vultr account.
	DateCreated *string `pulumi:"dateCreated"`
	// The server's default password.
	DefaultPassword *string `pulumi:"defaultPassword"`
	// The description of the disk(s) on the server.
	Disk *string `pulumi:"disk"`
	// Whether the server has IPv6 networking activated.
	EnableIpv6 *bool `pulumi:"enableIpv6"`
	// The server's IPv4 gateway.
	GatewayV4 *string `pulumi:"gatewayV4"`
	// The hostname to assign to the server.
	Hostname *string `pulumi:"hostname"`
	// The ID of the Vultr marketplace application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications) Note marketplace applications are denoted by type: `marketplace` and you must use the `imageId` not the id.
	ImageId *string `pulumi:"imageId"`
	// A label for the server.
	Label *string `pulumi:"label"`
	// The MAC address associated with the server.
	MacAddress *int `pulumi:"macAddress"`
	// The server's main IP address.
	MainIp *string `pulumi:"mainIp"`
	// The server's IPv4 netmask.
	NetmaskV4 *string `pulumi:"netmaskV4"`
	// The string description of the operating system installed on the server.
	Os *string `pulumi:"os"`
	// The ID of the operating system to be installed on the server. [See List OS](https://www.vultr.com/api/#operation/list-os)
	OsId *int `pulumi:"osId"`
	// The ID of the plan that you want the server to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)
	Plan *string `pulumi:"plan"`
	// The amount of memory available on the server in MB.
	Ram *string `pulumi:"ram"`
	// The ID of the region that the server is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
	Region *string `pulumi:"region"`
	// The ID of the floating IP to use as the main IP of this server. [See Reserved IPs](https://www.vultr.com/api/#operation/list-reserved-ips)
	ReservedIpv4 *string `pulumi:"reservedIpv4"`
	// The ID of the startup script you want added to the server.
	ScriptId *string `pulumi:"scriptId"`
	// The ID of the Vultr snapshot that the server will restore for the initial installation. [See List Snapshots](https://www.vultr.com/api/#operation/list-snapshots)
	SnapshotId *string `pulumi:"snapshotId"`
	// A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).
	SshKeyIds []string `pulumi:"sshKeyIds"`
	// The status of the server's subscription.
	Status *string `pulumi:"status"`
	// A list of tags to apply to the servier.
	Tags []string `pulumi:"tags"`
	// Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.
	UserData *string `pulumi:"userData"`
	// The main IPv6 network address.
	V6MainIp *string `pulumi:"v6MainIp"`
	// The IPv6 subnet.
	V6Network *string `pulumi:"v6Network"`
	// The IPv6 network size in bits.
	V6NetworkSize *int `pulumi:"v6NetworkSize"`
}

type BareMetalServerState struct {
	// Whether an activation email will be sent when the server is ready.
	ActivationEmail pulumi.BoolPtrInput
	// The ID of the Vultr application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications)
	AppId pulumi.IntPtrInput
	// The number of CPUs available on the server.
	CpuCount pulumi.IntPtrInput
	// The date the server was added to your Vultr account.
	DateCreated pulumi.StringPtrInput
	// The server's default password.
	DefaultPassword pulumi.StringPtrInput
	// The description of the disk(s) on the server.
	Disk pulumi.StringPtrInput
	// Whether the server has IPv6 networking activated.
	EnableIpv6 pulumi.BoolPtrInput
	// The server's IPv4 gateway.
	GatewayV4 pulumi.StringPtrInput
	// The hostname to assign to the server.
	Hostname pulumi.StringPtrInput
	// The ID of the Vultr marketplace application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications) Note marketplace applications are denoted by type: `marketplace` and you must use the `imageId` not the id.
	ImageId pulumi.StringPtrInput
	// A label for the server.
	Label pulumi.StringPtrInput
	// The MAC address associated with the server.
	MacAddress pulumi.IntPtrInput
	// The server's main IP address.
	MainIp pulumi.StringPtrInput
	// The server's IPv4 netmask.
	NetmaskV4 pulumi.StringPtrInput
	// The string description of the operating system installed on the server.
	Os pulumi.StringPtrInput
	// The ID of the operating system to be installed on the server. [See List OS](https://www.vultr.com/api/#operation/list-os)
	OsId pulumi.IntPtrInput
	// The ID of the plan that you want the server to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)
	Plan pulumi.StringPtrInput
	// The amount of memory available on the server in MB.
	Ram pulumi.StringPtrInput
	// The ID of the region that the server is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
	Region pulumi.StringPtrInput
	// The ID of the floating IP to use as the main IP of this server. [See Reserved IPs](https://www.vultr.com/api/#operation/list-reserved-ips)
	ReservedIpv4 pulumi.StringPtrInput
	// The ID of the startup script you want added to the server.
	ScriptId pulumi.StringPtrInput
	// The ID of the Vultr snapshot that the server will restore for the initial installation. [See List Snapshots](https://www.vultr.com/api/#operation/list-snapshots)
	SnapshotId pulumi.StringPtrInput
	// A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).
	SshKeyIds pulumi.StringArrayInput
	// The status of the server's subscription.
	Status pulumi.StringPtrInput
	// A list of tags to apply to the servier.
	Tags pulumi.StringArrayInput
	// Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.
	UserData pulumi.StringPtrInput
	// The main IPv6 network address.
	V6MainIp pulumi.StringPtrInput
	// The IPv6 subnet.
	V6Network pulumi.StringPtrInput
	// The IPv6 network size in bits.
	V6NetworkSize pulumi.IntPtrInput
}

func (BareMetalServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*bareMetalServerState)(nil)).Elem()
}

type bareMetalServerArgs struct {
	// Whether an activation email will be sent when the server is ready.
	ActivationEmail *bool `pulumi:"activationEmail"`
	// The ID of the Vultr application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications)
	AppId *int `pulumi:"appId"`
	// Whether the server has IPv6 networking activated.
	EnableIpv6 *bool `pulumi:"enableIpv6"`
	// The hostname to assign to the server.
	Hostname *string `pulumi:"hostname"`
	// The ID of the Vultr marketplace application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications) Note marketplace applications are denoted by type: `marketplace` and you must use the `imageId` not the id.
	ImageId *string `pulumi:"imageId"`
	// A label for the server.
	Label *string `pulumi:"label"`
	// The ID of the operating system to be installed on the server. [See List OS](https://www.vultr.com/api/#operation/list-os)
	OsId *int `pulumi:"osId"`
	// The ID of the plan that you want the server to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)
	Plan string `pulumi:"plan"`
	// The ID of the region that the server is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
	Region string `pulumi:"region"`
	// The ID of the floating IP to use as the main IP of this server. [See Reserved IPs](https://www.vultr.com/api/#operation/list-reserved-ips)
	ReservedIpv4 *string `pulumi:"reservedIpv4"`
	// The ID of the startup script you want added to the server.
	ScriptId *string `pulumi:"scriptId"`
	// The ID of the Vultr snapshot that the server will restore for the initial installation. [See List Snapshots](https://www.vultr.com/api/#operation/list-snapshots)
	SnapshotId *string `pulumi:"snapshotId"`
	// A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).
	SshKeyIds []string `pulumi:"sshKeyIds"`
	// A list of tags to apply to the servier.
	Tags []string `pulumi:"tags"`
	// Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.
	UserData *string `pulumi:"userData"`
}

// The set of arguments for constructing a BareMetalServer resource.
type BareMetalServerArgs struct {
	// Whether an activation email will be sent when the server is ready.
	ActivationEmail pulumi.BoolPtrInput
	// The ID of the Vultr application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications)
	AppId pulumi.IntPtrInput
	// Whether the server has IPv6 networking activated.
	EnableIpv6 pulumi.BoolPtrInput
	// The hostname to assign to the server.
	Hostname pulumi.StringPtrInput
	// The ID of the Vultr marketplace application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications) Note marketplace applications are denoted by type: `marketplace` and you must use the `imageId` not the id.
	ImageId pulumi.StringPtrInput
	// A label for the server.
	Label pulumi.StringPtrInput
	// The ID of the operating system to be installed on the server. [See List OS](https://www.vultr.com/api/#operation/list-os)
	OsId pulumi.IntPtrInput
	// The ID of the plan that you want the server to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)
	Plan pulumi.StringInput
	// The ID of the region that the server is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
	Region pulumi.StringInput
	// The ID of the floating IP to use as the main IP of this server. [See Reserved IPs](https://www.vultr.com/api/#operation/list-reserved-ips)
	ReservedIpv4 pulumi.StringPtrInput
	// The ID of the startup script you want added to the server.
	ScriptId pulumi.StringPtrInput
	// The ID of the Vultr snapshot that the server will restore for the initial installation. [See List Snapshots](https://www.vultr.com/api/#operation/list-snapshots)
	SnapshotId pulumi.StringPtrInput
	// A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).
	SshKeyIds pulumi.StringArrayInput
	// A list of tags to apply to the servier.
	Tags pulumi.StringArrayInput
	// Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.
	UserData pulumi.StringPtrInput
}

func (BareMetalServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bareMetalServerArgs)(nil)).Elem()
}

type BareMetalServerInput interface {
	pulumi.Input

	ToBareMetalServerOutput() BareMetalServerOutput
	ToBareMetalServerOutputWithContext(ctx context.Context) BareMetalServerOutput
}

func (*BareMetalServer) ElementType() reflect.Type {
	return reflect.TypeOf((**BareMetalServer)(nil)).Elem()
}

func (i *BareMetalServer) ToBareMetalServerOutput() BareMetalServerOutput {
	return i.ToBareMetalServerOutputWithContext(context.Background())
}

func (i *BareMetalServer) ToBareMetalServerOutputWithContext(ctx context.Context) BareMetalServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BareMetalServerOutput)
}

// BareMetalServerArrayInput is an input type that accepts BareMetalServerArray and BareMetalServerArrayOutput values.
// You can construct a concrete instance of `BareMetalServerArrayInput` via:
//
//	BareMetalServerArray{ BareMetalServerArgs{...} }
type BareMetalServerArrayInput interface {
	pulumi.Input

	ToBareMetalServerArrayOutput() BareMetalServerArrayOutput
	ToBareMetalServerArrayOutputWithContext(context.Context) BareMetalServerArrayOutput
}

type BareMetalServerArray []BareMetalServerInput

func (BareMetalServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BareMetalServer)(nil)).Elem()
}

func (i BareMetalServerArray) ToBareMetalServerArrayOutput() BareMetalServerArrayOutput {
	return i.ToBareMetalServerArrayOutputWithContext(context.Background())
}

func (i BareMetalServerArray) ToBareMetalServerArrayOutputWithContext(ctx context.Context) BareMetalServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BareMetalServerArrayOutput)
}

// BareMetalServerMapInput is an input type that accepts BareMetalServerMap and BareMetalServerMapOutput values.
// You can construct a concrete instance of `BareMetalServerMapInput` via:
//
//	BareMetalServerMap{ "key": BareMetalServerArgs{...} }
type BareMetalServerMapInput interface {
	pulumi.Input

	ToBareMetalServerMapOutput() BareMetalServerMapOutput
	ToBareMetalServerMapOutputWithContext(context.Context) BareMetalServerMapOutput
}

type BareMetalServerMap map[string]BareMetalServerInput

func (BareMetalServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BareMetalServer)(nil)).Elem()
}

func (i BareMetalServerMap) ToBareMetalServerMapOutput() BareMetalServerMapOutput {
	return i.ToBareMetalServerMapOutputWithContext(context.Background())
}

func (i BareMetalServerMap) ToBareMetalServerMapOutputWithContext(ctx context.Context) BareMetalServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BareMetalServerMapOutput)
}

type BareMetalServerOutput struct{ *pulumi.OutputState }

func (BareMetalServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BareMetalServer)(nil)).Elem()
}

func (o BareMetalServerOutput) ToBareMetalServerOutput() BareMetalServerOutput {
	return o
}

func (o BareMetalServerOutput) ToBareMetalServerOutputWithContext(ctx context.Context) BareMetalServerOutput {
	return o
}

// Whether an activation email will be sent when the server is ready.
func (o BareMetalServerOutput) ActivationEmail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.BoolPtrOutput { return v.ActivationEmail }).(pulumi.BoolPtrOutput)
}

// The ID of the Vultr application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications)
func (o BareMetalServerOutput) AppId() pulumi.IntOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.IntOutput { return v.AppId }).(pulumi.IntOutput)
}

// The number of CPUs available on the server.
func (o BareMetalServerOutput) CpuCount() pulumi.IntOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.IntOutput { return v.CpuCount }).(pulumi.IntOutput)
}

// The date the server was added to your Vultr account.
func (o BareMetalServerOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringOutput { return v.DateCreated }).(pulumi.StringOutput)
}

// The server's default password.
func (o BareMetalServerOutput) DefaultPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringOutput { return v.DefaultPassword }).(pulumi.StringOutput)
}

// The description of the disk(s) on the server.
func (o BareMetalServerOutput) Disk() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringOutput { return v.Disk }).(pulumi.StringOutput)
}

// Whether the server has IPv6 networking activated.
func (o BareMetalServerOutput) EnableIpv6() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.BoolPtrOutput { return v.EnableIpv6 }).(pulumi.BoolPtrOutput)
}

// The server's IPv4 gateway.
func (o BareMetalServerOutput) GatewayV4() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringOutput { return v.GatewayV4 }).(pulumi.StringOutput)
}

// The hostname to assign to the server.
func (o BareMetalServerOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringPtrOutput { return v.Hostname }).(pulumi.StringPtrOutput)
}

// The ID of the Vultr marketplace application to be installed on the server. [See List Applications](https://www.vultr.com/api/#operation/list-applications) Note marketplace applications are denoted by type: `marketplace` and you must use the `imageId` not the id.
func (o BareMetalServerOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

// A label for the server.
func (o BareMetalServerOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringPtrOutput { return v.Label }).(pulumi.StringPtrOutput)
}

// The MAC address associated with the server.
func (o BareMetalServerOutput) MacAddress() pulumi.IntOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.IntOutput { return v.MacAddress }).(pulumi.IntOutput)
}

// The server's main IP address.
func (o BareMetalServerOutput) MainIp() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringOutput { return v.MainIp }).(pulumi.StringOutput)
}

// The server's IPv4 netmask.
func (o BareMetalServerOutput) NetmaskV4() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringOutput { return v.NetmaskV4 }).(pulumi.StringOutput)
}

// The string description of the operating system installed on the server.
func (o BareMetalServerOutput) Os() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringOutput { return v.Os }).(pulumi.StringOutput)
}

// The ID of the operating system to be installed on the server. [See List OS](https://www.vultr.com/api/#operation/list-os)
func (o BareMetalServerOutput) OsId() pulumi.IntOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.IntOutput { return v.OsId }).(pulumi.IntOutput)
}

// The ID of the plan that you want the server to subscribe to. [See List Plans](https://www.vultr.com/api/#tag/plans)
func (o BareMetalServerOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringOutput { return v.Plan }).(pulumi.StringOutput)
}

// The amount of memory available on the server in MB.
func (o BareMetalServerOutput) Ram() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringOutput { return v.Ram }).(pulumi.StringOutput)
}

// The ID of the region that the server is to be created in. [See List Regions](https://www.vultr.com/api/#operation/list-regions)
func (o BareMetalServerOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The ID of the floating IP to use as the main IP of this server. [See Reserved IPs](https://www.vultr.com/api/#operation/list-reserved-ips)
func (o BareMetalServerOutput) ReservedIpv4() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringOutput { return v.ReservedIpv4 }).(pulumi.StringOutput)
}

// The ID of the startup script you want added to the server.
func (o BareMetalServerOutput) ScriptId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringPtrOutput { return v.ScriptId }).(pulumi.StringPtrOutput)
}

// The ID of the Vultr snapshot that the server will restore for the initial installation. [See List Snapshots](https://www.vultr.com/api/#operation/list-snapshots)
func (o BareMetalServerOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringPtrOutput { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

// A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).
func (o BareMetalServerOutput) SshKeyIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringArrayOutput { return v.SshKeyIds }).(pulumi.StringArrayOutput)
}

// The status of the server's subscription.
func (o BareMetalServerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A list of tags to apply to the servier.
func (o BareMetalServerOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.
func (o BareMetalServerOutput) UserData() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringOutput { return v.UserData }).(pulumi.StringOutput)
}

// The main IPv6 network address.
func (o BareMetalServerOutput) V6MainIp() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringOutput { return v.V6MainIp }).(pulumi.StringOutput)
}

// The IPv6 subnet.
func (o BareMetalServerOutput) V6Network() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.StringOutput { return v.V6Network }).(pulumi.StringOutput)
}

// The IPv6 network size in bits.
func (o BareMetalServerOutput) V6NetworkSize() pulumi.IntOutput {
	return o.ApplyT(func(v *BareMetalServer) pulumi.IntOutput { return v.V6NetworkSize }).(pulumi.IntOutput)
}

type BareMetalServerArrayOutput struct{ *pulumi.OutputState }

func (BareMetalServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BareMetalServer)(nil)).Elem()
}

func (o BareMetalServerArrayOutput) ToBareMetalServerArrayOutput() BareMetalServerArrayOutput {
	return o
}

func (o BareMetalServerArrayOutput) ToBareMetalServerArrayOutputWithContext(ctx context.Context) BareMetalServerArrayOutput {
	return o
}

func (o BareMetalServerArrayOutput) Index(i pulumi.IntInput) BareMetalServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BareMetalServer {
		return vs[0].([]*BareMetalServer)[vs[1].(int)]
	}).(BareMetalServerOutput)
}

type BareMetalServerMapOutput struct{ *pulumi.OutputState }

func (BareMetalServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BareMetalServer)(nil)).Elem()
}

func (o BareMetalServerMapOutput) ToBareMetalServerMapOutput() BareMetalServerMapOutput {
	return o
}

func (o BareMetalServerMapOutput) ToBareMetalServerMapOutputWithContext(ctx context.Context) BareMetalServerMapOutput {
	return o
}

func (o BareMetalServerMapOutput) MapIndex(k pulumi.StringInput) BareMetalServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BareMetalServer {
		return vs[0].(map[string]*BareMetalServer)[vs[1].(string)]
	}).(BareMetalServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BareMetalServerInput)(nil)).Elem(), &BareMetalServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*BareMetalServerArrayInput)(nil)).Elem(), BareMetalServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BareMetalServerMapInput)(nil)).Elem(), BareMetalServerMap{})
	pulumi.RegisterOutputType(BareMetalServerOutput{})
	pulumi.RegisterOutputType(BareMetalServerArrayOutput{})
	pulumi.RegisterOutputType(BareMetalServerMapOutput{})
}
