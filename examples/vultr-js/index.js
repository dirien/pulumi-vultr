"use strict";
const vultr = require("@ediri/vultr");


const vke = new vultr.Kubernetes("vke", {
    region: "fra",
    version: "v1.25.4+1",
    label: "pulumi-vultr",
    nodePools: {
        nodeQuantity: 1,
        plan: "vc2-2c-4gb",
        label: "pulumi-vultr-nodepool",
    },
})

exports.kubeConfig = vke.kubeConfig
