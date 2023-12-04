package yaml

import (
	"github.com/spf13/viper"
)

func Unmarshal[T any](path string, setting *T) error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return viper.Unmarshal(setting)
}
func UnmarshalArray[T any](path string, setting []T) error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return viper.Unmarshal(setting)
}
