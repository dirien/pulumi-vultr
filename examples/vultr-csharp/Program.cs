using System.Collections.Generic;
using Pulumi;
using ediri.Vultr; 

return await Deployment.RunAsync(() =>
{
   var vke = new Kubernetes("vke", new KubernetesArgs{
      Region = "fra",
      Version = "v1.25.4+1",
      Label = "pulumi-vultr",
      NodePools = new ediri.Vultr.Inputs.KubernetesNodePoolsArgs{
           NodeQuantity = 1,
           Plan = "vc2-2c-4gb",
           Label = "pulumi-vultr",
         },
      }
   );
   return new Dictionary<string, object?>
   {
      ["kubeConfig"] = vke.KubeConfig,
   };
});
