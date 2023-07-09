package config

import (
	"flag"

	"github.com/spf13/pflag"
)

type MustBindFlag struct{ Name, Value, Usage string }

type MustBindFlags []MustBindFlag

func initFlag(flags MustBindFlags) error {

	for _, f := range flags {
		flag.String(f.Name, f.Value, f.Usage)
	}

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.Parse()

	if err := Config.BindPFlags(pflag.CommandLine); err != nil {
		return err
	}
	return nil
}
