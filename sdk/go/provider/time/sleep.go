// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package time

import (
	"context"
	"reflect"

	"errors"
	"github.com/jhoeflaken/pulumi-azure-ext/provider/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sleep for a specified amount of time.
type Sleep struct {
	pulumi.CustomResourceState

	Duration pulumi.StringOutput `pulumi:"duration"`
}

// NewSleep registers a new resource with the given unique name, arguments, and options.
func NewSleep(ctx *pulumi.Context,
	name string, args *SleepArgs, opts ...pulumi.ResourceOption) (*Sleep, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Duration == nil {
		return nil, errors.New("invalid value for required argument 'Duration'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Sleep
	err := ctx.RegisterResource("azure-ext:time:Sleep", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSleep gets an existing Sleep resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSleep(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SleepState, opts ...pulumi.ResourceOption) (*Sleep, error) {
	var resource Sleep
	err := ctx.ReadResource("azure-ext:time:Sleep", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Sleep resources.
type sleepState struct {
}

type SleepState struct {
}

func (SleepState) ElementType() reflect.Type {
	return reflect.TypeOf((*sleepState)(nil)).Elem()
}

type sleepArgs struct {
	Duration string `pulumi:"duration"`
}

// The set of arguments for constructing a Sleep resource.
type SleepArgs struct {
	Duration pulumi.StringInput
}

func (SleepArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sleepArgs)(nil)).Elem()
}

type SleepInput interface {
	pulumi.Input

	ToSleepOutput() SleepOutput
	ToSleepOutputWithContext(ctx context.Context) SleepOutput
}

func (*Sleep) ElementType() reflect.Type {
	return reflect.TypeOf((**Sleep)(nil)).Elem()
}

func (i *Sleep) ToSleepOutput() SleepOutput {
	return i.ToSleepOutputWithContext(context.Background())
}

func (i *Sleep) ToSleepOutputWithContext(ctx context.Context) SleepOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SleepOutput)
}

// SleepArrayInput is an input type that accepts SleepArray and SleepArrayOutput values.
// You can construct a concrete instance of `SleepArrayInput` via:
//
//	SleepArray{ SleepArgs{...} }
type SleepArrayInput interface {
	pulumi.Input

	ToSleepArrayOutput() SleepArrayOutput
	ToSleepArrayOutputWithContext(context.Context) SleepArrayOutput
}

type SleepArray []SleepInput

func (SleepArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sleep)(nil)).Elem()
}

func (i SleepArray) ToSleepArrayOutput() SleepArrayOutput {
	return i.ToSleepArrayOutputWithContext(context.Background())
}

func (i SleepArray) ToSleepArrayOutputWithContext(ctx context.Context) SleepArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SleepArrayOutput)
}

// SleepMapInput is an input type that accepts SleepMap and SleepMapOutput values.
// You can construct a concrete instance of `SleepMapInput` via:
//
//	SleepMap{ "key": SleepArgs{...} }
type SleepMapInput interface {
	pulumi.Input

	ToSleepMapOutput() SleepMapOutput
	ToSleepMapOutputWithContext(context.Context) SleepMapOutput
}

type SleepMap map[string]SleepInput

func (SleepMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sleep)(nil)).Elem()
}

func (i SleepMap) ToSleepMapOutput() SleepMapOutput {
	return i.ToSleepMapOutputWithContext(context.Background())
}

func (i SleepMap) ToSleepMapOutputWithContext(ctx context.Context) SleepMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SleepMapOutput)
}

type SleepOutput struct{ *pulumi.OutputState }

func (SleepOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sleep)(nil)).Elem()
}

func (o SleepOutput) ToSleepOutput() SleepOutput {
	return o
}

func (o SleepOutput) ToSleepOutputWithContext(ctx context.Context) SleepOutput {
	return o
}

func (o SleepOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v *Sleep) pulumi.StringOutput { return v.Duration }).(pulumi.StringOutput)
}

type SleepArrayOutput struct{ *pulumi.OutputState }

func (SleepArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sleep)(nil)).Elem()
}

func (o SleepArrayOutput) ToSleepArrayOutput() SleepArrayOutput {
	return o
}

func (o SleepArrayOutput) ToSleepArrayOutputWithContext(ctx context.Context) SleepArrayOutput {
	return o
}

func (o SleepArrayOutput) Index(i pulumi.IntInput) SleepOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Sleep {
		return vs[0].([]*Sleep)[vs[1].(int)]
	}).(SleepOutput)
}

type SleepMapOutput struct{ *pulumi.OutputState }

func (SleepMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sleep)(nil)).Elem()
}

func (o SleepMapOutput) ToSleepMapOutput() SleepMapOutput {
	return o
}

func (o SleepMapOutput) ToSleepMapOutputWithContext(ctx context.Context) SleepMapOutput {
	return o
}

func (o SleepMapOutput) MapIndex(k pulumi.StringInput) SleepOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Sleep {
		return vs[0].(map[string]*Sleep)[vs[1].(string)]
	}).(SleepOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SleepInput)(nil)).Elem(), &Sleep{})
	pulumi.RegisterInputType(reflect.TypeOf((*SleepArrayInput)(nil)).Elem(), SleepArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SleepMapInput)(nil)).Elem(), SleepMap{})
	pulumi.RegisterOutputType(SleepOutput{})
	pulumi.RegisterOutputType(SleepArrayOutput{})
	pulumi.RegisterOutputType(SleepMapOutput{})
}
