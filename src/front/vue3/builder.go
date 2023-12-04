/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/12/2 15:37
 */

package vue3

import (
	"qf-builder/front/vue3/apis"
	modelsBuilder "qf-builder/front/vue3/models"
	"qf-builder/models"
	"strings"
)

var pathRoot = ""

// Build 生成vue3代码
func Build(setting *models.Setting) (errorCode int) {
	pathRoot = strings.Replace(setting.Output, "${project}", setting.Project, 1)
	//fmt.Println("pathRoot", pathRoot)
	errorCode = modelsBuilder.Build(setting, pathRoot)
	if errorCode != 0 {
		return
	}
	errorCode = apis.Build(setting, pathRoot)
	if errorCode != 0 {
		return
	}
	return
}
