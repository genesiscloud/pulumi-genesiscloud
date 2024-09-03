// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace GenesisCloud.PulumiPackage.Genesiscloud.Inputs
{

    public sealed class InstanceMetadataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A plain text bash script or "cloud-config" file that will be executed after the first instance boot. It is limited to 64 KiB in size. You can use it to configure your instance, e.g. installing the NVIDIA GPU driver. Learn more about [startup scripts and installing the GPU driver](https://support.genesiscloud.com/support/solutions/articles/47001122478).
        ///   - If the value of this attribute changes, the resource will be replaced.
        /// </summary>
        [Input("startupScript")]
        public Input<string>? StartupScript { get; set; }

        public InstanceMetadataArgs()
        {
        }
        public static new InstanceMetadataArgs Empty => new InstanceMetadataArgs();
    }
}
