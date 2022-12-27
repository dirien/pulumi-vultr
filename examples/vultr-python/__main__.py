from pulumi import export
import ediri_vultr as vultr

vke = vultr.Kubernetes('vke', region='fra', version='v1.25.4+1',
                       label='pulumi-vke',
                       node_pools=vultr.KubernetesNodePoolsArgs(node_quantity=1,
                                                                plan='vc2-2c-4gb', label='pulumi-vultr-nodepool'))

export('kubeConfig', vke.kube_config)
