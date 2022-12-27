import * as vultr from "@ediri/vultr";

let vke = new vultr.Kubernetes("vke", {
    region: "fra",
    version: "v1.25.4+1",
    label: "pulumi-vultr",
    nodePools: {
        nodeQuantity: 1,
        plan: "vc2-2c-4gb",
        label: "pulumi-vultr-nodepool",
    },
})

export const kubeconfig = vke.kubeConfig;
