/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/12/4 14:54
 */

package util

import (
	"errors"
	"qf-builder/models"
	"qf-builder/util/strHelper"
	"strings"
)

func Split(src string) []string {
	return strings.Split(src, "|")
}
func Trim(src string) string {
	return strings.Trim(src, " ")
}

// CheckType 检查类型 是否为数组 是否为指针
func CheckType(src string) (isArray, isPoint bool, typeName string) {
	if strings.Contains(src, "%") || strings.Contains(src, "[]") {
		isArray = true
	}
	if strings.Contains(src, "*") {
		isPoint = true
	}
	return isArray, isPoint, strings.Trim(src, "%[]* ")
}

// DecStdDefine 标准定义行解码
func DecStdDefine(define string) (name, comment string, err error) {
	array := Split(define)
	if len(array) < 1 {
		err = errors.New("invalid type define: " + define)
		return
	}
	name = strHelper.ToCamel(array[0])
	if len(array) > 1 {
		comment = Trim(array[1])
	}
	return
}

// DecProp Prop Param 行解码 example: id|uint64|用户id
func DecProp(prop string) (models.Prop, error) {
	return decP(prop, true)
}
func DecParam(param string) (models.Prop, error) {
	return decP(param, false)
}
func decP(param string, toCamel bool) (models.Prop, error) {
	array := Split(param)
	if len(array) < 2 {
		return models.Prop{}, errors.New("invalid prop define: " + param)
	}
	output := models.Prop{}
	if toCamel {
		output.Name = strHelper.ToCamel(array[0])
	} else {
		output.Name = Trim(array[0])
	}
	if len(array) > 2 {
		output.Comment = Trim(array[2])
	}
	tName := array[1]
	output.IsArray, output.IsPoint, output.TypeName = CheckType(tName)
	return output, nil
}
