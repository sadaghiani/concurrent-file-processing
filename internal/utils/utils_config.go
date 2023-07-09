package utils

import (
	"github.com/sadaghiani/concurrent-file-processing/pkg/config"
	"github.com/sadaghiani/concurrent-file-processing/pkg/database"
)

const (
	APP_LOG_LEVEL       config.MustBindEnv = "APP_LOG_LEVEL"
	APP_BATCH_SIZE      config.MustBindEnv = "APP_BATCH_SIZE"
	APP_NUMBER_WORKERS  config.MustBindEnv = "APP_NUMBER_WORKERS"
	DATABASE_URI        config.MustBindEnv = "DATABASE_URI"
	DATABASE_PORT       config.MustBindEnv = "DATABASE_PORT"
	DATABASE_NAME       config.MustBindEnv = "DATABASE_NAME"
	DATABASE_COLLECTION config.MustBindEnv = "DATABASE_COLLECTION"
	DATABASE_USER       config.MustBindEnv = "DATABASE_USER"
	DATABASE_PASSWORD   config.MustBindEnv = "DATABASE_PASSWORD"
)

func CreateMustBindEnvs() config.MustBindEnvs {
	return config.MustBindEnvs{
		APP_LOG_LEVEL,
		APP_BATCH_SIZE,
		APP_NUMBER_WORKERS,
		DATABASE_URI,
		DATABASE_PORT,
		DATABASE_NAME,
		DATABASE_COLLECTION,
		DATABASE_USER,
		DATABASE_PASSWORD,
	}
}

func CreateMustBindFlags() config.MustBindFlags {
	return config.MustBindFlags{
		config.MustBindFlag{
			Name:  "port",
			Value: "80",
			Usage: "http listener port",
		},
		config.MustBindFlag{
			Name:  "host",
			Value: "",
			Usage: "http listener address",
		},
	}
}

func CreateMongoDataStoreConfig() database.MongoDataStoreConfig {
	return database.MongoDataStoreConfig{
		DatabaseURI:        config.Config.GetString(MustBindEnvToString(DATABASE_URI)),
		DatabasePort:       config.Config.GetString(MustBindEnvToString(DATABASE_PORT)),
		DatabaseName:       config.Config.GetString(MustBindEnvToString(DATABASE_NAME)),
		DatabaseCollection: config.Config.GetString(MustBindEnvToString(DATABASE_COLLECTION)),
		DatabaseUser:       config.Config.GetString(MustBindEnvToString(DATABASE_USER)),
		DatabasePassword:   config.Config.GetString(MustBindEnvToString(DATABASE_PASSWORD)),
	}
}

func MustBindEnvToString(env config.MustBindEnv) string {
	return string(env)
}
