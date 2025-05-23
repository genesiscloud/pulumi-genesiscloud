// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Snapshot resource
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import genesiscloud:index/snapshot:Snapshot example 18efeec8-94f0-4776-8ff2-5e9b49c74608
 * ```
 */
export class Snapshot extends pulumi.CustomResource {
    /**
     * Get an existing Snapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnapshotState, opts?: pulumi.CustomResourceOptions): Snapshot {
        return new Snapshot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'genesiscloud:index/snapshot:Snapshot';

    /**
     * Returns true if the given object is an instance of Snapshot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Snapshot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Snapshot.__pulumiType;
    }

    /**
     * The timestamp when this snapshot was created in RFC 3339.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The human-readable name for the snapshot.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region identifier. Should only be explicity specified when using the 'source_snapshot_id'.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Target region for snapshot replication. When specified, also creates a copy of the snapshot in the given region. If
     * omitted, the snapshot exists only in the current region.
     */
    public readonly replicatedRegion!: pulumi.Output<string | undefined>;
    /**
     * Flag to retain the snapshot when the resource is deleted. - Sets the default value "false" if the attribute is not set.
     */
    public readonly retainOnDelete!: pulumi.Output<boolean>;
    /**
     * The storage size of this snapshot given in GiB.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * The id of the source instance from which this snapshot was derived. - If the value of this attribute changes, the
     * resource will be replaced.
     */
    public readonly sourceInstanceId!: pulumi.Output<string | undefined>;
    /**
     * The id of the source snapshot from which this snapsot was derived. - If the value of this attribute changes, the
     * resource will be replaced.
     */
    public readonly sourceSnapshotId!: pulumi.Output<string | undefined>;
    /**
     * The snapshot status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly timeouts!: pulumi.Output<outputs.SnapshotTimeouts | undefined>;

    /**
     * Create a Snapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SnapshotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnapshotArgs | SnapshotState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnapshotState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["replicatedRegion"] = state ? state.replicatedRegion : undefined;
            resourceInputs["retainOnDelete"] = state ? state.retainOnDelete : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["sourceInstanceId"] = state ? state.sourceInstanceId : undefined;
            resourceInputs["sourceSnapshotId"] = state ? state.sourceSnapshotId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["timeouts"] = state ? state.timeouts : undefined;
        } else {
            const args = argsOrState as SnapshotArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["replicatedRegion"] = args ? args.replicatedRegion : undefined;
            resourceInputs["retainOnDelete"] = args ? args.retainOnDelete : undefined;
            resourceInputs["sourceInstanceId"] = args ? args.sourceInstanceId : undefined;
            resourceInputs["sourceSnapshotId"] = args ? args.sourceSnapshotId : undefined;
            resourceInputs["timeouts"] = args ? args.timeouts : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Snapshot.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Snapshot resources.
 */
export interface SnapshotState {
    /**
     * The timestamp when this snapshot was created in RFC 3339.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The human-readable name for the snapshot.
     */
    name?: pulumi.Input<string>;
    /**
     * The region identifier. Should only be explicity specified when using the 'source_snapshot_id'.
     */
    region?: pulumi.Input<string>;
    /**
     * Target region for snapshot replication. When specified, also creates a copy of the snapshot in the given region. If
     * omitted, the snapshot exists only in the current region.
     */
    replicatedRegion?: pulumi.Input<string>;
    /**
     * Flag to retain the snapshot when the resource is deleted. - Sets the default value "false" if the attribute is not set.
     */
    retainOnDelete?: pulumi.Input<boolean>;
    /**
     * The storage size of this snapshot given in GiB.
     */
    size?: pulumi.Input<number>;
    /**
     * The id of the source instance from which this snapshot was derived. - If the value of this attribute changes, the
     * resource will be replaced.
     */
    sourceInstanceId?: pulumi.Input<string>;
    /**
     * The id of the source snapshot from which this snapsot was derived. - If the value of this attribute changes, the
     * resource will be replaced.
     */
    sourceSnapshotId?: pulumi.Input<string>;
    /**
     * The snapshot status.
     */
    status?: pulumi.Input<string>;
    timeouts?: pulumi.Input<inputs.SnapshotTimeouts>;
}

/**
 * The set of arguments for constructing a Snapshot resource.
 */
export interface SnapshotArgs {
    /**
     * The human-readable name for the snapshot.
     */
    name?: pulumi.Input<string>;
    /**
     * The region identifier. Should only be explicity specified when using the 'source_snapshot_id'.
     */
    region?: pulumi.Input<string>;
    /**
     * Target region for snapshot replication. When specified, also creates a copy of the snapshot in the given region. If
     * omitted, the snapshot exists only in the current region.
     */
    replicatedRegion?: pulumi.Input<string>;
    /**
     * Flag to retain the snapshot when the resource is deleted. - Sets the default value "false" if the attribute is not set.
     */
    retainOnDelete?: pulumi.Input<boolean>;
    /**
     * The id of the source instance from which this snapshot was derived. - If the value of this attribute changes, the
     * resource will be replaced.
     */
    sourceInstanceId?: pulumi.Input<string>;
    /**
     * The id of the source snapshot from which this snapsot was derived. - If the value of this attribute changes, the
     * resource will be replaced.
     */
    sourceSnapshotId?: pulumi.Input<string>;
    timeouts?: pulumi.Input<inputs.SnapshotTimeouts>;
}
