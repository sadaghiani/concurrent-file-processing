package config

import (
	"fmt"
)

type MustBindEnv string

type MustBindEnvs []MustBindEnv

func mustEnvs(mustSet MustBindEnvs) error {

	var errs []string
	for _, v := range mustSet {
		if !Config.IsSet(string(v)) {
			errs = append(errs, string(v))
		}

	}

	if len(errs) != 0 {
		return fmt.Errorf("%s must declare in os environment", errs)
	}
	return nil
}
