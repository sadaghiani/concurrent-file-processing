package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	Config    *viper.Viper
	onceInit  sync.Once
	configErr error
)

func Init(envs MustBindEnvs, flags MustBindFlags) error {

	onceInit.Do(func() {
		Config = viper.New()
		Config.AutomaticEnv()
		if configErr = mustEnvs(envs); configErr != nil {
			return
		}
		if configErr = initFlag(flags); configErr != nil {
			return
		}
	})

	return configErr

}

func NewConfig(envs MustBindEnvs, flags MustBindFlags) {

	if err := Init(envs, flags); err != nil {
		log.Fatalln("failed to initialize config: ", err)
	}
}
