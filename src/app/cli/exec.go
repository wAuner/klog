package cli

import (
	"errors"
	"fmt"
	"github.com/alecthomas/kong"
	"klog"
	"klog/app"
	"reflect"
)

type cli struct {
	Print    Print    `cmd help:"Show records in a file"`
	Evaluate Evaluate `cmd help:"Evaluate the times in records"`
	Widget   Widget   `cmd help:"Launch menu bar widget (MacOS only)"`
	Version  kong.VersionFlag
}

func Execute() int {
	ctx, err := app.NewContextFromEnv()
	if err != nil {
		fmt.Println("Failed to initialise application. Error:")
		fmt.Println(err)
		return -1
	}
	args := kong.Parse(
		&cli{},
		kong.Name("klog"),
		kong.Description("klog time tracking: command line app for interacting with `.klg` files."),
		kong.UsageOnError(),
		func() kong.Option {
			datePrototype, _ := klog.NewDate(1, 1, 1)
			return kong.TypeMapper(reflect.TypeOf(&datePrototype).Elem(), dateDecoder())
		}(),
		kong.Vars{
			"version": ctx.Version(),
		},
	)
	err = args.Run(ctx)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return 0
}

func dateDecoder() kong.MapperFunc {
	return func(ctx *kong.DecodeContext, target reflect.Value) error {
		var value string
		if err := ctx.Scan.PopValueInto("date", &value); err != nil {
			return err
		}
		if value == "" {
			return errors.New("please provide a valid date")
		}
		d, err := klog.NewDateFromString(value)
		if err != nil {
			return errors.New("`" + value + "` is not a valid date")
		}
		target.Set(reflect.ValueOf(d))
		return nil
	}
}
