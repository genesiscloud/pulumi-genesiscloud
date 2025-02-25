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

// Volume resource
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
//			_, err := genesiscloud.NewVolume(ctx, "example", &genesiscloud.VolumeArgs{
//				Region: pulumi.String("NORD-NO-KRS-1"),
//				Size:   pulumi.Int(50),
//				Type:   pulumi.String("hdd"),
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
// $ pulumi import genesiscloud:index/volume:Volume example 18efeec8-94f0-4776-8ff2-5e9b49c74608
// ```
type Volume struct {
	pulumi.CustomResourceState

	// The timestamp when this volume was created in RFC 3339.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The human-readable description for the volume. - Sets the default value "" if the attribute is not set.
	Description pulumi.StringOutput `pulumi:"description"`
	// The human-readable name for the volume.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region identifier. - If the value of this attribute changes, the resource will be replaced. - The value must be one
	// of: ["EUC-DE-MUC-1" "NORD-NO-KRS-1"].
	Region pulumi.StringOutput `pulumi:"region"`
	// Flag to retain the volume when the resource is deleted - Sets the default value "false" if the attribute is not set.
	RetainOnDelete pulumi.BoolOutput `pulumi:"retainOnDelete"`
	// The storage size of this volume given in GiB. - The value must be at least 1.
	Size pulumi.IntOutput `pulumi:"size"`
	// The volume status.
	Status   pulumi.StringOutput     `pulumi:"status"`
	Timeouts VolumeTimeoutsPtrOutput `pulumi:"timeouts"`
	// The storage type of the volume. - If the value of this attribute changes, the resource will be replaced. - The value
	// must be one of: ["hdd" "ssd"].
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVolume registers a new resource with the given unique name, arguments, and options.
func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Volume
	err := ctx.RegisterResource("genesiscloud:index/volume:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolume gets an existing Volume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("genesiscloud:index/volume:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Volume resources.
type volumeState struct {
	// The timestamp when this volume was created in RFC 3339.
	CreatedAt *string `pulumi:"createdAt"`
	// The human-readable description for the volume. - Sets the default value "" if the attribute is not set.
	Description *string `pulumi:"description"`
	// The human-readable name for the volume.
	Name *string `pulumi:"name"`
	// The region identifier. - If the value of this attribute changes, the resource will be replaced. - The value must be one
	// of: ["EUC-DE-MUC-1" "NORD-NO-KRS-1"].
	Region *string `pulumi:"region"`
	// Flag to retain the volume when the resource is deleted - Sets the default value "false" if the attribute is not set.
	RetainOnDelete *bool `pulumi:"retainOnDelete"`
	// The storage size of this volume given in GiB. - The value must be at least 1.
	Size *int `pulumi:"size"`
	// The volume status.
	Status   *string         `pulumi:"status"`
	Timeouts *VolumeTimeouts `pulumi:"timeouts"`
	// The storage type of the volume. - If the value of this attribute changes, the resource will be replaced. - The value
	// must be one of: ["hdd" "ssd"].
	Type *string `pulumi:"type"`
}

type VolumeState struct {
	// The timestamp when this volume was created in RFC 3339.
	CreatedAt pulumi.StringPtrInput
	// The human-readable description for the volume. - Sets the default value "" if the attribute is not set.
	Description pulumi.StringPtrInput
	// The human-readable name for the volume.
	Name pulumi.StringPtrInput
	// The region identifier. - If the value of this attribute changes, the resource will be replaced. - The value must be one
	// of: ["EUC-DE-MUC-1" "NORD-NO-KRS-1"].
	Region pulumi.StringPtrInput
	// Flag to retain the volume when the resource is deleted - Sets the default value "false" if the attribute is not set.
	RetainOnDelete pulumi.BoolPtrInput
	// The storage size of this volume given in GiB. - The value must be at least 1.
	Size pulumi.IntPtrInput
	// The volume status.
	Status   pulumi.StringPtrInput
	Timeouts VolumeTimeoutsPtrInput
	// The storage type of the volume. - If the value of this attribute changes, the resource will be replaced. - The value
	// must be one of: ["hdd" "ssd"].
	Type pulumi.StringPtrInput
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	// The human-readable description for the volume. - Sets the default value "" if the attribute is not set.
	Description *string `pulumi:"description"`
	// The human-readable name for the volume.
	Name *string `pulumi:"name"`
	// The region identifier. - If the value of this attribute changes, the resource will be replaced. - The value must be one
	// of: ["EUC-DE-MUC-1" "NORD-NO-KRS-1"].
	Region string `pulumi:"region"`
	// Flag to retain the volume when the resource is deleted - Sets the default value "false" if the attribute is not set.
	RetainOnDelete *bool `pulumi:"retainOnDelete"`
	// The storage size of this volume given in GiB. - The value must be at least 1.
	Size     int             `pulumi:"size"`
	Timeouts *VolumeTimeouts `pulumi:"timeouts"`
	// The storage type of the volume. - If the value of this attribute changes, the resource will be replaced. - The value
	// must be one of: ["hdd" "ssd"].
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Volume resource.
type VolumeArgs struct {
	// The human-readable description for the volume. - Sets the default value "" if the attribute is not set.
	Description pulumi.StringPtrInput
	// The human-readable name for the volume.
	Name pulumi.StringPtrInput
	// The region identifier. - If the value of this attribute changes, the resource will be replaced. - The value must be one
	// of: ["EUC-DE-MUC-1" "NORD-NO-KRS-1"].
	Region pulumi.StringInput
	// Flag to retain the volume when the resource is deleted - Sets the default value "false" if the attribute is not set.
	RetainOnDelete pulumi.BoolPtrInput
	// The storage size of this volume given in GiB. - The value must be at least 1.
	Size     pulumi.IntInput
	Timeouts VolumeTimeoutsPtrInput
	// The storage type of the volume. - If the value of this attribute changes, the resource will be replaced. - The value
	// must be one of: ["hdd" "ssd"].
	Type pulumi.StringInput
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeArgs)(nil)).Elem()
}

type VolumeInput interface {
	pulumi.Input

	ToVolumeOutput() VolumeOutput
	ToVolumeOutputWithContext(ctx context.Context) VolumeOutput
}

func (*Volume) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (i *Volume) ToVolumeOutput() VolumeOutput {
	return i.ToVolumeOutputWithContext(context.Background())
}

func (i *Volume) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeOutput)
}

// VolumeArrayInput is an input type that accepts VolumeArray and VolumeArrayOutput values.
// You can construct a concrete instance of `VolumeArrayInput` via:
//
//	VolumeArray{ VolumeArgs{...} }
type VolumeArrayInput interface {
	pulumi.Input

	ToVolumeArrayOutput() VolumeArrayOutput
	ToVolumeArrayOutputWithContext(context.Context) VolumeArrayOutput
}

type VolumeArray []VolumeInput

func (VolumeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Volume)(nil)).Elem()
}

func (i VolumeArray) ToVolumeArrayOutput() VolumeArrayOutput {
	return i.ToVolumeArrayOutputWithContext(context.Background())
}

func (i VolumeArray) ToVolumeArrayOutputWithContext(ctx context.Context) VolumeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeArrayOutput)
}

// VolumeMapInput is an input type that accepts VolumeMap and VolumeMapOutput values.
// You can construct a concrete instance of `VolumeMapInput` via:
//
//	VolumeMap{ "key": VolumeArgs{...} }
type VolumeMapInput interface {
	pulumi.Input

	ToVolumeMapOutput() VolumeMapOutput
	ToVolumeMapOutputWithContext(context.Context) VolumeMapOutput
}

type VolumeMap map[string]VolumeInput

func (VolumeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Volume)(nil)).Elem()
}

func (i VolumeMap) ToVolumeMapOutput() VolumeMapOutput {
	return i.ToVolumeMapOutputWithContext(context.Background())
}

func (i VolumeMap) ToVolumeMapOutputWithContext(ctx context.Context) VolumeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeMapOutput)
}

type VolumeOutput struct{ *pulumi.OutputState }

func (VolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (o VolumeOutput) ToVolumeOutput() VolumeOutput {
	return o
}

func (o VolumeOutput) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return o
}

// The timestamp when this volume was created in RFC 3339.
func (o VolumeOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The human-readable description for the volume. - Sets the default value "" if the attribute is not set.
func (o VolumeOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The human-readable name for the volume.
func (o VolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region identifier. - If the value of this attribute changes, the resource will be replaced. - The value must be one
// of: ["EUC-DE-MUC-1" "NORD-NO-KRS-1"].
func (o VolumeOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Flag to retain the volume when the resource is deleted - Sets the default value "false" if the attribute is not set.
func (o VolumeOutput) RetainOnDelete() pulumi.BoolOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolOutput { return v.RetainOnDelete }).(pulumi.BoolOutput)
}

// The storage size of this volume given in GiB. - The value must be at least 1.
func (o VolumeOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *Volume) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// The volume status.
func (o VolumeOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o VolumeOutput) Timeouts() VolumeTimeoutsPtrOutput {
	return o.ApplyT(func(v *Volume) VolumeTimeoutsPtrOutput { return v.Timeouts }).(VolumeTimeoutsPtrOutput)
}

// The storage type of the volume. - If the value of this attribute changes, the resource will be replaced. - The value
// must be one of: ["hdd" "ssd"].
func (o VolumeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type VolumeArrayOutput struct{ *pulumi.OutputState }

func (VolumeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Volume)(nil)).Elem()
}

func (o VolumeArrayOutput) ToVolumeArrayOutput() VolumeArrayOutput {
	return o
}

func (o VolumeArrayOutput) ToVolumeArrayOutputWithContext(ctx context.Context) VolumeArrayOutput {
	return o
}

func (o VolumeArrayOutput) Index(i pulumi.IntInput) VolumeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Volume {
		return vs[0].([]*Volume)[vs[1].(int)]
	}).(VolumeOutput)
}

type VolumeMapOutput struct{ *pulumi.OutputState }

func (VolumeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Volume)(nil)).Elem()
}

func (o VolumeMapOutput) ToVolumeMapOutput() VolumeMapOutput {
	return o
}

func (o VolumeMapOutput) ToVolumeMapOutputWithContext(ctx context.Context) VolumeMapOutput {
	return o
}

func (o VolumeMapOutput) MapIndex(k pulumi.StringInput) VolumeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Volume {
		return vs[0].(map[string]*Volume)[vs[1].(string)]
	}).(VolumeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeInput)(nil)).Elem(), &Volume{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeArrayInput)(nil)).Elem(), VolumeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeMapInput)(nil)).Elem(), VolumeMap{})
	pulumi.RegisterOutputType(VolumeOutput{})
	pulumi.RegisterOutputType(VolumeArrayOutput{})
	pulumi.RegisterOutputType(VolumeMapOutput{})
}
