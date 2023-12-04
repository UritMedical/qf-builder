/**
 * @Author: Joey
 * @Description: 对define中的文件进行解析
 * @Create Date: 2023/12/1 17:05
 */

package defines

import (
	_ "embed"
	"fmt"
	"os"
	"qf-builder/defines/decode"
	"qf-builder/defines/fileModels"
	"qf-builder/models"
	"qf-builder/util/io"
	"qf-builder/util/yaml"
)

// 这里在编译时将文件嵌入到源码
//
//go:embed models.demo.yaml
var demoModelsMod []byte

//go:embed module.demo.yaml
var demoModuleMod []byte

//go:embed module.demo_user.yaml
var demoUserModuleMod []byte

// Export 导出默认参数配置
func Export() int {
	var err error
	defer func() {
		if err != nil {
			fmt.Println("Error:", err)
		}
	}()
	err = os.MkdirAll("define", 0755)
	if err != nil {
		return models.ErrorCodeSystemError
	}
	err = os.WriteFile("define/models.demo.yaml", demoModelsMod, 0644)
	if err != nil {
		return models.ErrorCodeSystemError
	}
	err = os.WriteFile("define/module.demo.yaml", demoModuleMod, 0644)
	if err != nil {
		return models.ErrorCodeSystemError
	}
	err = os.WriteFile("define/module.demo_user.yaml", demoUserModuleMod, 0644)
	if err != nil {
		return models.ErrorCodeSystemError
	}
	fmt.Println("Export DemoDefine success")
	return 0
}

// LoadModels 加载指定路径下所有*models.yaml并将其转换为Models对象
func LoadModels(path string) ([]models.Model, error) {

	var modelsList []fileModels.FileModel

	// 定义正则表达式匹配文件名
	filePattern := `^[Mm]odels\S*.yaml$`
	files, err := io.FindFiles(path, filePattern)
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no *models.yaml files found in %s with %s", path, filePattern)
	}
	for _, file := range files {
		var node fileModels.FileModels
		err := yaml.Unmarshal(file, &node)
		if err != nil {
			return nil, fmt.Errorf("unmarshal yaml from %s error: %v", file, err)
		}
		modelsList = append(modelsList, node.Models...)
	}
	return decode.FileToModels(modelsList)
}

func LoadApis(path string) ([]models.Module, error) {
	var fileModules []fileModels.FileModule
	// 定义正则表达式匹配文件名
	filePattern := `^[Mm]odule\S*.yaml$`
	files, err := io.FindFiles(path, filePattern)
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no *module.yaml files found in %s with %s", path, filePattern)
	}
	for _, file := range files {
		var node fileModels.FileModule
		err := yaml.Unmarshal(file, &node)
		if err != nil {
			return nil, fmt.Errorf("unmarshal yaml from %s error: %v", file, err)
		}
		fileModules = append(fileModules, node)
	}
	return decode.FileToModules(fileModules)
}
