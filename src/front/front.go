/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/12/1 12:05
 */

package front

import (
	_ "embed"
	"fmt"
	"os"
	"qf-builder/front/vue3"
	"qf-builder/models"
)

// Build 生成
func Build(setting *models.Setting) (errorCode int) {
	fillSetting(setting)
	switch setting.Lang {
	case models.ELangVue3:
		fallthrough
	default:

		vue3.Build(setting)

	}
	return
}

// 这里在编译时将文件嵌入到源码
//
//go:embed default.yaml
var defaultYaml []byte

// Export 导出默认参数配置
func Export() int {

	path := "qf_front_default.yaml"
	err := os.WriteFile(path, defaultYaml, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return models.ErrorCodeSystemError
	}
	fmt.Println(path, "Export  success")
	return 0
}

var defaultSetting = models.Setting{
	Type:   models.ETypeFront,
	Lang:   models.ELangVue3,
	Define: "./define",
	Output: "./${project}_front_vue3",
	Src: []string{
		"front/project/desktop/std@latest",
		"front/style/desktop/std@latest",
	},
	Style: "std",
	Init:  "yarn upgrade",
}

// 使用默认值对setting进行填充
func fillSetting(t *models.Setting) {
	if t.Type == "" {
		t.Type = defaultSetting.Type
	}
	if t.Lang == "" {
		t.Lang = defaultSetting.Lang
	}
	if t.Define == "" {
		t.Define = defaultSetting.Define
	}
	if t.Output == "" {
		t.Output = defaultSetting.Output
	}
	if t.Src == nil {
		t.Src = defaultSetting.Src
	}
	if t.Style == "" {
		t.Style = defaultSetting.Style
	}
	if t.Init == "" {
		t.Init = defaultSetting.Init
	}
}
