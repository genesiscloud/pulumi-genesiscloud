// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("genesiscloud");

/**
 * Genesis Cloud API endpoint. May also be provided via `GENESISCLOUD_ENDPOINT` environment variable. If neither is
 * provided, defaults to `https://api.genesiscloud.com/compute/v1`.
 */
export declare const endpoint: string | undefined;
Object.defineProperty(exports, "endpoint", {
    get() {
        return __config.get("endpoint") ?? utilities.getEnv("GENESISCLOUD_ENDPOINT");
    },
    enumerable: true,
});

/**
 * The polling interval. - The string must be a positive [time duration](https://pkg.go.dev/time#ParseDuration), for
 * example "10s".
 */
export declare const pollingInterval: string | undefined;
Object.defineProperty(exports, "pollingInterval", {
    get() {
        return __config.get("pollingInterval");
    },
    enumerable: true,
});

/**
 * Genesis Cloud API token. May also be provided via `GENESISCLOUD_TOKEN` environment variable.
 */
export declare const token: string | undefined;
Object.defineProperty(exports, "token", {
    get() {
        return __config.get("token") ?? utilities.getEnv("GENESISCLOUD_TOKEN");
    },
    enumerable: true,
});

