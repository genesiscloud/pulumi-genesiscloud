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

    public sealed class ImagesFilterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filter by the region identifier.
        ///   - The value must be one of: ["ARC-IS-HAF-1" "EUC-DE-MUC-1" "NORD-NO-KRS-1"].
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// Filter by the kind of image.
        ///   - The value must be one of: ["base-os" "cloud-image" "preconfigured" "snapshot"].
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public ImagesFilterArgs()
        {
        }
        public static new ImagesFilterArgs Empty => new ImagesFilterArgs();
    }
}