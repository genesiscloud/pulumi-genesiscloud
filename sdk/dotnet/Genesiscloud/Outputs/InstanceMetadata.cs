// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace GenesisCloud.PulumiPackage.Genesiscloud.Outputs
{

    [OutputType]
    public sealed class InstanceMetadata
    {
        public readonly string? StartupScript;

        [OutputConstructor]
        private InstanceMetadata(string? startupScript)
        {
            StartupScript = startupScript;
        }
    }
}