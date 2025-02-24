// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * floating IP resource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as genesiscloud from "@genesiscloud/pulumi-genesiscloud";
 *
 * const floatingIp = new genesiscloud.FloatingIp("floatingIp", {
 *     region: "NORD-NO-KRS-1",
 *     version: "ipv4",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import genesiscloud:index/floatingIp:FloatingIp example 18efeec8-94f0-4776-8ff2-5e9b49c74608
 * ```
 */
export class FloatingIp extends pulumi.CustomResource {
    /**
     * Get an existing FloatingIp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FloatingIpState, opts?: pulumi.CustomResourceOptions): FloatingIp {
        return new FloatingIp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'genesiscloud:index/floatingIp:FloatingIp';

    /**
     * Returns true if the given object is an instance of FloatingIp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FloatingIp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FloatingIp.__pulumiType;
    }

    /**
     * The timestamp when this floating IP was created in RFC 3339.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The human-readable description set for the floating IP. - Sets the default value "" if the attribute is not set.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The IP address of the floating IP.
     */
    public /*out*/ readonly ipAddress!: pulumi.Output<string>;
    /**
     * Whether the floating IP is public or private. - Sets the default value "true" if the attribute is not set.
     */
    public /*out*/ readonly isPublic!: pulumi.Output<boolean>;
    /**
     * The human-readable name for the floating IP.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region identifier. - If the value of this attribute changes, the resource will be replaced. - The value must be one
     * of: ["EUC-DE-MUC-1" "NORD-NO-KRS-1"].
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The floating IP status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly timeouts!: pulumi.Output<outputs.FloatingIpTimeouts | undefined>;
    /**
     * The timestamp when this image was last updated in RFC 3339.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The version of the floating IP. - If the value of this attribute changes, the resource will be replaced. - The value
     * must be one of: ["ipv4"].
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a FloatingIp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FloatingIpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FloatingIpArgs | FloatingIpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FloatingIpState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["isPublic"] = state ? state.isPublic : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["timeouts"] = state ? state.timeouts : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as FloatingIpArgs | undefined;
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.version === undefined) && !opts.urn) {
                throw new Error("Missing required property 'version'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["timeouts"] = args ? args.timeouts : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["ipAddress"] = undefined /*out*/;
            resourceInputs["isPublic"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FloatingIp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FloatingIp resources.
 */
export interface FloatingIpState {
    /**
     * The timestamp when this floating IP was created in RFC 3339.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The human-readable description set for the floating IP. - Sets the default value "" if the attribute is not set.
     */
    description?: pulumi.Input<string>;
    /**
     * The IP address of the floating IP.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * Whether the floating IP is public or private. - Sets the default value "true" if the attribute is not set.
     */
    isPublic?: pulumi.Input<boolean>;
    /**
     * The human-readable name for the floating IP.
     */
    name?: pulumi.Input<string>;
    /**
     * The region identifier. - If the value of this attribute changes, the resource will be replaced. - The value must be one
     * of: ["EUC-DE-MUC-1" "NORD-NO-KRS-1"].
     */
    region?: pulumi.Input<string>;
    /**
     * The floating IP status.
     */
    status?: pulumi.Input<string>;
    timeouts?: pulumi.Input<inputs.FloatingIpTimeouts>;
    /**
     * The timestamp when this image was last updated in RFC 3339.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The version of the floating IP. - If the value of this attribute changes, the resource will be replaced. - The value
     * must be one of: ["ipv4"].
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FloatingIp resource.
 */
export interface FloatingIpArgs {
    /**
     * The human-readable description set for the floating IP. - Sets the default value "" if the attribute is not set.
     */
    description?: pulumi.Input<string>;
    /**
     * The human-readable name for the floating IP.
     */
    name?: pulumi.Input<string>;
    /**
     * The region identifier. - If the value of this attribute changes, the resource will be replaced. - The value must be one
     * of: ["EUC-DE-MUC-1" "NORD-NO-KRS-1"].
     */
    region: pulumi.Input<string>;
    timeouts?: pulumi.Input<inputs.FloatingIpTimeouts>;
    /**
     * The version of the floating IP. - If the value of this attribute changes, the resource will be replaced. - The value
     * must be one of: ["ipv4"].
     */
    version: pulumi.Input<string>;
}
