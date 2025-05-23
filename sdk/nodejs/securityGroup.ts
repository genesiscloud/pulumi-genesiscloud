// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Security group resource
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import genesiscloud:index/securityGroup:SecurityGroup example 18efeec8-94f0-4776-8ff2-5e9b49c74608
 * ```
 */
export class SecurityGroup extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityGroupState, opts?: pulumi.CustomResourceOptions): SecurityGroup {
        return new SecurityGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'genesiscloud:index/securityGroup:SecurityGroup';

    /**
     * Returns true if the given object is an instance of SecurityGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityGroup.__pulumiType;
    }

    /**
     * The timestamp when this security group was created in RFC 3339.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The human-readable description for the security group. - Sets the default value "" if the attribute is not set.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The human-readable name for the security group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region identifier. - If the value of this attribute changes, the resource will be replaced. - The value must be one
     * of: ["EUC-DE-MUC-1" "EUW-GB-MNC-1" "EUW-NL-AMS-1" "NA-CA-FTS-1" "NA-CA-MNZ-1" "NA-CA-PRG-1" "NORD-NO-KRS-1"].
     */
    public readonly region!: pulumi.Output<string>;
    public readonly rules!: pulumi.Output<outputs.SecurityGroupRule[]>;
    /**
     * The security group status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly timeouts!: pulumi.Output<outputs.SecurityGroupTimeouts | undefined>;

    /**
     * Create a SecurityGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityGroupArgs | SecurityGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityGroupState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["timeouts"] = state ? state.timeouts : undefined;
        } else {
            const args = argsOrState as SecurityGroupArgs | undefined;
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.rules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rules'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["timeouts"] = args ? args.timeouts : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroup resources.
 */
export interface SecurityGroupState {
    /**
     * The timestamp when this security group was created in RFC 3339.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The human-readable description for the security group. - Sets the default value "" if the attribute is not set.
     */
    description?: pulumi.Input<string>;
    /**
     * The human-readable name for the security group.
     */
    name?: pulumi.Input<string>;
    /**
     * The region identifier. - If the value of this attribute changes, the resource will be replaced. - The value must be one
     * of: ["EUC-DE-MUC-1" "EUW-GB-MNC-1" "EUW-NL-AMS-1" "NA-CA-FTS-1" "NA-CA-MNZ-1" "NA-CA-PRG-1" "NORD-NO-KRS-1"].
     */
    region?: pulumi.Input<string>;
    rules?: pulumi.Input<pulumi.Input<inputs.SecurityGroupRule>[]>;
    /**
     * The security group status.
     */
    status?: pulumi.Input<string>;
    timeouts?: pulumi.Input<inputs.SecurityGroupTimeouts>;
}

/**
 * The set of arguments for constructing a SecurityGroup resource.
 */
export interface SecurityGroupArgs {
    /**
     * The human-readable description for the security group. - Sets the default value "" if the attribute is not set.
     */
    description?: pulumi.Input<string>;
    /**
     * The human-readable name for the security group.
     */
    name?: pulumi.Input<string>;
    /**
     * The region identifier. - If the value of this attribute changes, the resource will be replaced. - The value must be one
     * of: ["EUC-DE-MUC-1" "EUW-GB-MNC-1" "EUW-NL-AMS-1" "NA-CA-FTS-1" "NA-CA-MNZ-1" "NA-CA-PRG-1" "NORD-NO-KRS-1"].
     */
    region: pulumi.Input<string>;
    rules: pulumi.Input<pulumi.Input<inputs.SecurityGroupRule>[]>;
    timeouts?: pulumi.Input<inputs.SecurityGroupTimeouts>;
}
