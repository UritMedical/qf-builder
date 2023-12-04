/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/12/4 8:24
 */

package strHelper

import (
	"fmt"
	"github.com/gobeam/stringy"
	"regexp"
	"strings"
)

// GetMatchedString 从源字符串中返回匹配的字符串
func GetMatchedString(source, pattern string) (string, error) {

	reg, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}

	match := reg.FindStringSubmatch(source)
	if len(match) > 0 {
		return match[0], nil // 返回第一个捕获组的内容
	}
	return "", nil
}

// Clear 清理字符串
func Clear(source string, pattern string) string {
	reg, _ := regexp.Compile(pattern)
	return strings.TrimLeft(reg.ReplaceAllString(source, ""), "\r\n")

}
func Match(source string, pattern string) bool {
	reg, _ := regexp.Compile(pattern)
	return reg.MatchString(source)
}

func GetSection(source string, sectionName string) (string, error) {
	pattern := `(?s)\[` + sectionName + `\].*\[End` + sectionName + `\]`
	matched, err := GetMatchedString(source, pattern)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	clearPattern := `\[\S*\]`
	return Clear(matched, clearPattern), nil
}

// ConvertToTsType 转换后端数据类型到前端 TypeScript 数据类型
func ConvertToTsType(typeName string) string {
	typeMappings := map[string]string{
		`^[Uu]?[Ii]nt\S*`: "number",
		`^[Uu]?[Ss]hort`:  "number",
		`^[Ff]loat`:       "number",
		`^[Ss]ingle`:      "number",
		`^[Dd]ouble`:      "number",
		`^[Uu]?[Ll]ong`:   "number",
		`^byte`:           "number",
		`^[Dd]ecimal`:     "number",
		`^[Bb]ool\S*`:     "boolean",
		`^[Dd]ate\S*`:     "string",
		`^[Tt]ime\S*`:     "string",
		`^[Gg]uid`:        "string",
		`^[Ss]tring`:      "string",
	}
	for pattern, tsType := range typeMappings {
		if Match(typeName, pattern) {
			return tsType
		}
	}

	return typeName
}
func IsTsValueType(typeName string) bool {
	return typeName == "string" || typeName == "number" || typeName == "boolean"
}

// ToFileName 将字符串转换为文件名 蛇形
func ToFileName(name string, suffix string) string {
	snakeName := stringy.New(name).SnakeCase().ToLower()
	//if suffix == "" {
	//	return snakeName
	//}
	return strings.Join([]string{snakeName, suffix}, ".")
}

// ToSnake 将字符串转换为驼峰
func ToSnake(name string) string {
	name = strings.Trim(name, " ")
	return stringy.New(name).SnakeCase().ToLower()
}

func ToCamel(name string) string {
	name = strings.Trim(name, " ")
	return stringy.New(name).CamelCase()
}
