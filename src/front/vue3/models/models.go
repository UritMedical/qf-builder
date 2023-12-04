package models

import (
	_ "embed"
	"fmt"
	"qf-builder/defines"
	"qf-builder/models"
	"qf-builder/util/io"
	"qf-builder/util/strHelper"
	"strings"
)

const (
	pathModels = "src/defines/models/index.ts"
)

var (
	rootTemp  string
	modelTemp string
	propTemp  string
)

// 嵌入文件 modelsTemp.yaml
//
//go:embed modelsTemp.txt
var temp string

func init() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()
	fields := []struct {
		target  *string
		section string
	}{
		{&rootTemp, "Root"},
		{&modelTemp, "Model"},
		{&propTemp, "Prop"},
	}
	for _, v := range fields {
		*v.target, err = strHelper.GetSection(temp, v.section)
		if err != nil {
			return
		}
	}

}

// Build 生成models
func Build(setting *models.Setting, pathRoot string) (errorCode int) {
	var err error
	var ms []models.Model
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()
	ms, err = defines.LoadModels(setting.Define)
	if err != nil {
		return models.ErrorCodeDefineFormatError
	}
	str := rootTemp
	modelStr, e := buildModels(ms)
	if e != nil {
		return models.ErrorCodeDefineFormatError
	}
	str = strings.Replace(str, "{{Models}}", modelStr, 1)
	err = io.SaveToFile(strings.Join([]string{pathRoot, pathModels}, "/"), []byte(str))
	if err != nil {
		return 0
	}
	return
}
func buildModels(model []models.Model) (string, error) {
	output := ""
	for _, t := range model {
		str := modelTemp
		str = strings.Replace(str, "{{Name}}", t.Name, 1)
		str = strings.Replace(str, "{{Comment}}", t.Comment, 1)
		props, err := buildProps(t.Props)
		if err != nil {
			return "", err
		}
		str = strings.Replace(str, "{{Props}}", props, 1)
		output += str
	}
	return strings.TrimRight(output, "\r\n"), nil
}

func buildProps(props []models.Prop) (string, error) {
	output := ""
	for _, p := range props {
		str := propTemp
		str = strings.Replace(str, "{{Name}}", p.Name, 1)
		str = strings.Replace(str, "{{Model}}", strHelper.ConvertToTsType(p.TypeName), 1)
		str = strings.Replace(str, "{{Comment}}", p.Comment, 1)
		output += str
	}
	return strings.TrimRight(output, "\r\n"), nil

}
