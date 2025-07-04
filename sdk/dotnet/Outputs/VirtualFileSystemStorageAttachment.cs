// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Vultr.Outputs
{

    [OutputType]
    public sealed class VirtualFileSystemStorageAttachment
    {
        public readonly string? InstanceId;
        public readonly int? Mount;
        public readonly string? State;

        [OutputConstructor]
        private VirtualFileSystemStorageAttachment(
            string? instanceId,

            int? mount,

            string? state)
        {
            InstanceId = instanceId;
            Mount = mount;
            State = state;
        }
    }
}
