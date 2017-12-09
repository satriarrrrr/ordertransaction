package infrastructures

import "github.com/spf13/viper"

// ReadConfig read configuration from file given its path and filename
func ReadConfig(path, filename string) error {
	viper.SetConfigName(filename)
	viper.AddConfigPath(path)
	return viper.ReadInConfig()
}
