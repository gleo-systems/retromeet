package env

import (
	"github.com/spf13/viper"
)

func GetVar(name string) string {
	return viper.GetString(name)
}

func GetIntVar(name string) int {
	return viper.GetInt(name)
}

// GetBoolVar returns false in case of invalid or missing value.
func GetBoolVar(name string) bool {
	return viper.GetBool(name)
}
