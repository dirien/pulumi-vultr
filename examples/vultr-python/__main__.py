from pulumi import export
import ediri_vultr as vultr

#vke = vultr.Kubernetes('vke', region='fra', version='v1.25.4+1',
#                       label='pulumi-vke',
#                       node_pools=vultr.KubernetesNodePoolsArgs(node_quantity=1,
#                                                                plan='vc2-2c-4gb', label='pulumi-vultr-nodepool'))

#export('kubeConfig', vke.kube_config)

vm = vultr.Instance('vm', region='fra',
                   plan='vc2-1c-2gb',
                   label='pulumi-vultr-vm',
                   os_id='1743',
                   enable_ipv6=True)
export('vmIpAddress', vm.id)
