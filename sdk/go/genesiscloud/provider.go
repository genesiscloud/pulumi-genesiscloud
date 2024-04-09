// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package genesiscloud

import (
	"context"
	"reflect"

	"github.com/genesiscloud/pulumi-genesiscloud/sdk/go/genesiscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the genesiscloud package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// Genesis Cloud API endpoint. May also be provided via `GENESISCLOUD_ENDPOINT` environment variable. If neither is
	// provided, defaults to `https://api.genesiscloud.com/compute/v1`.
	Endpoint pulumi.StringPtrOutput `pulumi:"endpoint"`
	// The polling interval. - The string must be a positive [time duration](https://pkg.go.dev/time#ParseDuration), for
	// example "10s".
	PollingInterval pulumi.StringPtrOutput `pulumi:"pollingInterval"`
	// Genesis Cloud API token. May also be provided via `GENESISCLOUD_TOKEN` environment variable.
	Token pulumi.StringPtrOutput `pulumi:"token"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.Endpoint == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "GENESISCLOUD_ENDPOINT"); d != nil {
			args.Endpoint = pulumi.StringPtr(d.(string))
		}
	}
	if args.Token == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "GENESISCLOUD_TOKEN"); d != nil {
			args.Token = pulumi.StringPtr(d.(string))
		}
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:genesiscloud", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// Genesis Cloud API endpoint. May also be provided via `GENESISCLOUD_ENDPOINT` environment variable. If neither is
	// provided, defaults to `https://api.genesiscloud.com/compute/v1`.
	Endpoint *string `pulumi:"endpoint"`
	// The polling interval. - The string must be a positive [time duration](https://pkg.go.dev/time#ParseDuration), for
	// example "10s".
	PollingInterval *string `pulumi:"pollingInterval"`
	// Genesis Cloud API token. May also be provided via `GENESISCLOUD_TOKEN` environment variable.
	Token *string `pulumi:"token"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// Genesis Cloud API endpoint. May also be provided via `GENESISCLOUD_ENDPOINT` environment variable. If neither is
	// provided, defaults to `https://api.genesiscloud.com/compute/v1`.
	Endpoint pulumi.StringPtrInput
	// The polling interval. - The string must be a positive [time duration](https://pkg.go.dev/time#ParseDuration), for
	// example "10s".
	PollingInterval pulumi.StringPtrInput
	// Genesis Cloud API token. May also be provided via `GENESISCLOUD_TOKEN` environment variable.
	Token pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// Genesis Cloud API endpoint. May also be provided via `GENESISCLOUD_ENDPOINT` environment variable. If neither is
// provided, defaults to `https://api.genesiscloud.com/compute/v1`.
func (o ProviderOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Endpoint }).(pulumi.StringPtrOutput)
}

// The polling interval. - The string must be a positive [time duration](https://pkg.go.dev/time#ParseDuration), for
// example "10s".
func (o ProviderOutput) PollingInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.PollingInterval }).(pulumi.StringPtrOutput)
}

// Genesis Cloud API token. May also be provided via `GENESISCLOUD_TOKEN` environment variable.
func (o ProviderOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Token }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
