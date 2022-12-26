package main

import (
	"github.com/dirien/pulumi-vultr/v2/go/vultr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		kubernetes, err := vultr.NewKubernetes(ctx, "kubernetes", &vultr.KubernetesArgs{
			Region:  pulumi.String("fra"),
			Version: pulumi.String("v1.25.4+1"),
			Label:   pulumi.String("pulumi-vultr"),
			NodePools: vultr.KubernetesNodePoolsTypeArgs{
				NodeQuantity: pulumi.Int(1),
				Plan:         pulumi.String("vc2-2c-4gb"),
				Label:        pulumi.String("pulumi-vultr-nodepool"),
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("kubeconfig", kubernetes.KubeConfig)
		return nil
	})
}
