package main

import (
	"fmt"
	"os"
	"qf-builder/defines"
	"qf-builder/front"
	"qf-builder/models"
	"qf-builder/server"
	"qf-builder/simulator"
	"qf-builder/util/yaml"
)

var version string

func main() {
	errCode := 0
	args := os.Args
	defer func() {
		if errCode != 0 {
			showHelp()
		}
		os.Exit(errCode)
	}()
	length := len(args)
	if length == 1 {
		showHelp()
		errCode = models.ErrorCodeParamsError
		return
	}
	action := args[1]
	switch action {
	case "-v":
		fallthrough
	case "version":
		showVersion()
		return
	case "-b":
		fallthrough
	case "build":
		{ //代码生成 qf build ./xxx.yaml
			if length == 3 {
				settingFilePath := args[2]
				setting, err := loadSetting(settingFilePath)
				if err != nil {
					fmt.Println(err)
					errCode = models.ErrorCodeParamsError
					return
				}
				switch setting.Type {
				case models.ETypeFront:
					errCode = front.Build(setting)
					return
				case models.ETypeServ:
					errCode = server.Build(setting)
					return
				case models.ETypeSimulator:
					errCode = simulator.Build(setting)
					return
				}
			}
		}
	case "-h":
		fallthrough
	case "help":
		showHelp()
		return
	case "-e":
		fallthrough
	case "export":
		if length == 3 {
			expType := args[2]
			switch expType {
			case models.ETypeFront:
				errCode = front.Export()
				return
			case models.ETypeDefine:
				errCode = defines.Export()
				return
			default:
				fmt.Println("export type error")
				errCode = models.ErrorCodeParamsError
				return
			}
		}
		errCode = models.ErrorCodeParamsError
	}
}

// LoadSetting reads and parses a YAML file to load settings
func loadSetting(path string) (*models.Setting, error) {
	var setting models.Setting
	// Read YAML file

	// Unmarshal YAML into the Setting struct
	err := yaml.Unmarshal(path, &setting)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %v", err)
	}
	return &setting, nil
}
