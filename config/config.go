package config

import "github.com/spf13/viper"

type Config struct {
	//environment
	Environment int `mapstructure:"ENVIRONMENT"` //XXX 0 local, 1 dev, 2 test, 3 prod

	//database
	DbDriver string `mapstructure:"DB_DRIVER"` //DB_DRIVER database driver name e.g. pgx, postgres, ...
	DbUrl    string `mapstructure:"DB_URL"`    //DB_URL	   database connection url

	//logger
	LogConsoleEnabled bool   `mapstructure:"LOG_CONSOLE_ENABLED"` //LOG_CONSOLE_ENABLED console log
	LogEncodeAsJson   bool   `mapstructure:"LOG_ENCODE_AS_JSON"`  //LOG_ENCODE_AS_JSON  encode logs as json object
	LogFileEnabled    bool   `mapstructure:"LOG_FILE_ENABLED"`    //LOG_FILE_ENABLED 	disable in production
	LogDirectory      string `mapstructure:"LOG_DIRECTORY"`       //LOG_DIRECTORY 		an absolute path of log directory
	LogFilename       string `mapstructure:"LOG_FILENAME"`        //LOG_FILENAME 		log filename with extension
	LogMaxSize        int    `mapstructure:"LOG_MAX_SIZE"`        //LOG_MAX_SIZE 		max size in MB
	LogMaxBackups     int    `mapstructure:"LOG_MAX_BACKUPS"`     //LOG_MAX_BACKUPS 	number of backup files
	LogMaxAge         int    `mapstructure:"LOG_MAX_AGE"`         //LOG_MAX_AGE 		max age in days

	//log policy
	LogLevel      int  `mapstructure:"LOG_LEVEL"`       //LOG_LEVEL (-1 trace, 0 debug, 1 info, 2 warn, 3 error, 4 fatal, 5 panic, 6 no level)
	LogLevelMin   int  `mapstructure:"LOG_LEVEL_MIN"`   //LOG_LEVEL_MIN (-1 trace, 0 debug, 1 info, 2 warn, 3 error, 4 fatal, 5 panic, 6 no level)
	LogErrorStack bool `mapstructure:"LOG_ERROR_STACK"` //LOG_ERROR_STACK
}

func LoadConfig() (config Config, err error) {
	//set file name and working dir of the
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("config")

	//check if environment variables match any of the existing keys
	viper.AutomaticEnv()

	//read the config file
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	//unmarshal the config into the Config struct
	err = viper.Unmarshal(&config)
	return
}
