// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package genesiscloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/genesiscloud/pulumi-genesiscloud/sdk/go/genesiscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Snapshot resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/genesiscloud/pulumi-genesiscloud/sdk/go/genesiscloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			target, err := genesiscloud.NewInstance(ctx, "target", nil)
//			if err != nil {
//				return err
//			}
//			_, err = genesiscloud.NewSnapshot(ctx, "example", &genesiscloud.SnapshotArgs{
//				InstanceId:     target.ID(),
//				RetainOnDelete: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
//
//	$ pulumi import genesiscloud:index/snapshot:Snapshot example 18efeec8-94f0-4776-8ff2-5e9b49c74608
//
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	// The timestamp when this snapshot was created in RFC 3339.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The id of the instance to snapshot. - If the value of this attribute changes, Terraform will destroy and recreate the
	// resource.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The human-readable name for the snapshot.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region identifier.
	Region pulumi.StringOutput `pulumi:"region"`
	// Flag to retain the snapshot when the resource is deleted. - Sets the default value "false" if the attribute is not set.
	RetainOnDelete pulumi.BoolOutput `pulumi:"retainOnDelete"`
	// The storage size of this snapshot given in bytes.
	Size pulumi.IntOutput `pulumi:"size"`
	// The snapshot status.
	Status   pulumi.StringOutput       `pulumi:"status"`
	Timeouts SnapshotTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Snapshot
	err := ctx.RegisterResource("genesiscloud:index/snapshot:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("genesiscloud:index/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// The timestamp when this snapshot was created in RFC 3339.
	CreatedAt *string `pulumi:"createdAt"`
	// The id of the instance to snapshot. - If the value of this attribute changes, Terraform will destroy and recreate the
	// resource.
	InstanceId *string `pulumi:"instanceId"`
	// The human-readable name for the snapshot.
	Name *string `pulumi:"name"`
	// The region identifier.
	Region *string `pulumi:"region"`
	// Flag to retain the snapshot when the resource is deleted. - Sets the default value "false" if the attribute is not set.
	RetainOnDelete *bool `pulumi:"retainOnDelete"`
	// The storage size of this snapshot given in bytes.
	Size *int `pulumi:"size"`
	// The snapshot status.
	Status   *string           `pulumi:"status"`
	Timeouts *SnapshotTimeouts `pulumi:"timeouts"`
}

type SnapshotState struct {
	// The timestamp when this snapshot was created in RFC 3339.
	CreatedAt pulumi.StringPtrInput
	// The id of the instance to snapshot. - If the value of this attribute changes, Terraform will destroy and recreate the
	// resource.
	InstanceId pulumi.StringPtrInput
	// The human-readable name for the snapshot.
	Name pulumi.StringPtrInput
	// The region identifier.
	Region pulumi.StringPtrInput
	// Flag to retain the snapshot when the resource is deleted. - Sets the default value "false" if the attribute is not set.
	RetainOnDelete pulumi.BoolPtrInput
	// The storage size of this snapshot given in bytes.
	Size pulumi.IntPtrInput
	// The snapshot status.
	Status   pulumi.StringPtrInput
	Timeouts SnapshotTimeoutsPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// The id of the instance to snapshot. - If the value of this attribute changes, Terraform will destroy and recreate the
	// resource.
	InstanceId string `pulumi:"instanceId"`
	// The human-readable name for the snapshot.
	Name *string `pulumi:"name"`
	// Flag to retain the snapshot when the resource is deleted. - Sets the default value "false" if the attribute is not set.
	RetainOnDelete *bool             `pulumi:"retainOnDelete"`
	Timeouts       *SnapshotTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// The id of the instance to snapshot. - If the value of this attribute changes, Terraform will destroy and recreate the
	// resource.
	InstanceId pulumi.StringInput
	// The human-readable name for the snapshot.
	Name pulumi.StringPtrInput
	// Flag to retain the snapshot when the resource is deleted. - Sets the default value "false" if the attribute is not set.
	RetainOnDelete pulumi.BoolPtrInput
	Timeouts       SnapshotTimeoutsPtrInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

// SnapshotArrayInput is an input type that accepts SnapshotArray and SnapshotArrayOutput values.
// You can construct a concrete instance of `SnapshotArrayInput` via:
//
//	SnapshotArray{ SnapshotArgs{...} }
type SnapshotArrayInput interface {
	pulumi.Input

	ToSnapshotArrayOutput() SnapshotArrayOutput
	ToSnapshotArrayOutputWithContext(context.Context) SnapshotArrayOutput
}

type SnapshotArray []SnapshotInput

func (SnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (i SnapshotArray) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return i.ToSnapshotArrayOutputWithContext(context.Background())
}

func (i SnapshotArray) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotArrayOutput)
}

// SnapshotMapInput is an input type that accepts SnapshotMap and SnapshotMapOutput values.
// You can construct a concrete instance of `SnapshotMapInput` via:
//
//	SnapshotMap{ "key": SnapshotArgs{...} }
type SnapshotMapInput interface {
	pulumi.Input

	ToSnapshotMapOutput() SnapshotMapOutput
	ToSnapshotMapOutputWithContext(context.Context) SnapshotMapOutput
}

type SnapshotMap map[string]SnapshotInput

func (SnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (i SnapshotMap) ToSnapshotMapOutput() SnapshotMapOutput {
	return i.ToSnapshotMapOutputWithContext(context.Background())
}

func (i SnapshotMap) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotMapOutput)
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

// The timestamp when this snapshot was created in RFC 3339.
func (o SnapshotOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The id of the instance to snapshot. - If the value of this attribute changes, Terraform will destroy and recreate the
// resource.
func (o SnapshotOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The human-readable name for the snapshot.
func (o SnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region identifier.
func (o SnapshotOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Flag to retain the snapshot when the resource is deleted. - Sets the default value "false" if the attribute is not set.
func (o SnapshotOutput) RetainOnDelete() pulumi.BoolOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.BoolOutput { return v.RetainOnDelete }).(pulumi.BoolOutput)
}

// The storage size of this snapshot given in bytes.
func (o SnapshotOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// The snapshot status.
func (o SnapshotOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o SnapshotOutput) Timeouts() SnapshotTimeoutsPtrOutput {
	return o.ApplyT(func(v *Snapshot) SnapshotTimeoutsPtrOutput { return v.Timeouts }).(SnapshotTimeoutsPtrOutput)
}

type SnapshotArrayOutput struct{ *pulumi.OutputState }

func (SnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) Index(i pulumi.IntInput) SnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].([]*Snapshot)[vs[1].(int)]
	}).(SnapshotOutput)
}

type SnapshotMapOutput struct{ *pulumi.OutputState }

func (SnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (o SnapshotMapOutput) ToSnapshotMapOutput() SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) MapIndex(k pulumi.StringInput) SnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].(map[string]*Snapshot)[vs[1].(string)]
	}).(SnapshotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotInput)(nil)).Elem(), &Snapshot{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotArrayInput)(nil)).Elem(), SnapshotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotMapInput)(nil)).Elem(), SnapshotMap{})
	pulumi.RegisterOutputType(SnapshotOutput{})
	pulumi.RegisterOutputType(SnapshotArrayOutput{})
	pulumi.RegisterOutputType(SnapshotMapOutput{})
}
