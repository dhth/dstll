package cmd

import (
	"github.com/BurntSushi/toml"
)

type dstllConfig struct {
	TUICfg TUIConfig `toml:"tui"`
}

type TUIConfig struct {
	ViewFileCmd *[]string `toml:"view_file_command"`
}

func readConfig(filePath string) (dstllConfig, error) {

	var config dstllConfig
	_, err := toml.DecodeFile(filePath, &config)
	if err != nil {
		return config, err
	}

	return config, nil

}
