// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vultr

import (
	"fmt"
	"github.com/dirien/pulumi-vultr/provider/v2/pkg/version"
	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/vultr/terraform-provider-vultr/vultr"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "vultr"
	// modules:
	mainMod = "index" // the vultr module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(vultr.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "vultr",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Vultr",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "dirien",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/dirien/pulumi-vultr",
		Description:       "A Pulumi package for creating and managing Vultr cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "vultr", "category/utility"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/dirien/pulumi-vultr",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "vultr",
		Config:    map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags")},
			// 	},
			// },
			"vultr_bare_metal_server":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "BareMetalServer")},
			"vultr_block_storage":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "BlockStorage")},
			"vultr_dns_domain":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DnsDomain")},
			"vultr_dns_record":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DnsRecord")},
			"vultr_firewall_group":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FirewallGroup")},
			"vultr_firewall_rule":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FirewallRule")},
			"vultr_instance":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Instance")},
			"vultr_instance_ipv4":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "InstanceIpv4")},
			"vultr_iso_private":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IsoPrivate")},
			"vultr_kubernetes":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Kubernetes")},
			"vultr_kubernetes_node_pools": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "KubernetesNodePools")},
			"vultr_load_balancer":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "LoadBalancer")},
			"vultr_object_storage":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ObjectStorage")},
			"vultr_private_network":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "PrivateNetwork")},
			"vultr_reserved_ip":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ReservedIp")},
			"vultr_reverse_ipv4":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ReverseIpv4")},
			"vultr_reverse_ipv6":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ReverseIpv6")},
			"vultr_snapshot":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Snapshot")},
			"vultr_snapshot_from_url":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SnapshotFromUrl")},
			"vultr_ssh_key":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SshKey")},
			"vultr_startup_script":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "StartupScript")},
			"vultr_user":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "User")},
			"vultr_vpc":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Vpc")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi")},
			"vultr_account:":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAccount")},
			"vultr_application":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getApplication")},
			"vultr_backup:":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBackup")},
			"vultr_bare_metal_plan:":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBareMetalPlan")},
			"vultr_bare_metal_server:":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBareMetalServer")},
			"vultr_block_storage:":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getBlockStorage")},
			"vultr_dns_domain:":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getDnsDomain")},
			"vultr_firewall_group":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getFirewallGroup")},
			"vultr_instance":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getInstance")},
			"vultr_instance_ipv4":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getInstanceIpv4")},
			"vultr_iso_private":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIsoPrivate")},
			"vultr_iso_public":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIsoPublic")},
			"vultr_kubernetes":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getKubernetes")},
			"vultr_load_balancer":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLoadBalancer")},
			"vultr_object_storage":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getObjectStorage")},
			"vultr_object_storage_cluster": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getObjectStorageCluster")},
			"vultr_os":                     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getOs")},
			"vultr_plan":                   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getPlan")},
			"vultr_private_network":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getPrivateNetwork")},
			"vultr_region":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRegion")},
			"vultr_reserved_ip":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getReservedIp")},
			"vultr_reverse_ipv4":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getReverseIpv4")},
			"vultr_reverse_ipv6":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getReverseIpv6")},
			"vultr_snapshot":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSnapshot")},
			"vultr_ssh_key":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSshKey")},
			"vultr_startup_script":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getStartupScript")},
			"vultr_user":                   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUser")},
			"vultr_vpc":                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVpc")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@dirien/vultr",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "dirien_vultr",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumiverse/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "dirien",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "io.dirien",
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
