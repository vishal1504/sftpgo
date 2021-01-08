// +build linux

package config

import "github.com/spf13/viper"

// linux specific config search path
func setViperAdditionalConfigPaths() {
	viper.AddConfigPath("/app")
}
