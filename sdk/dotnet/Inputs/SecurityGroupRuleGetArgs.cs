// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GenesisCloud.Inputs
{

    public sealed class SecurityGroupRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The direction of the rule.
        ///   - The value must be one of: ["egress" "ingress"].
        /// </summary>
        [Input("direction", required: true)]
        public Input<string> Direction { get; set; } = null!;

        /// <summary>
        /// The maximum port number of the rule.
        ///   - The value must be between 1 and 65535.
        /// </summary>
        [Input("portRangeMax")]
        public Input<int>? PortRangeMax { get; set; }

        /// <summary>
        /// The minimum port number of the rule.
        ///   - The value must be between 1 and 65535.
        /// </summary>
        [Input("portRangeMin")]
        public Input<int>? PortRangeMin { get; set; }

        /// <summary>
        /// The protocol of the rule.
        ///   - The value must be one of: ["all" "icmp" "tcp" "udp"].
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        public SecurityGroupRuleGetArgs()
        {
        }
        public static new SecurityGroupRuleGetArgs Empty => new SecurityGroupRuleGetArgs();
    }
}
