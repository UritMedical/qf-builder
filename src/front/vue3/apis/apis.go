package apis

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
	path = "src/defines/apis/"
)

// 嵌入文件 modelsTemp.yaml
//
//go:embed apisTemp.txt
var temp string
var (
	tmpIndex        string
	tmpIndexImport  string
	tmpIndexExport  string
	tmpIndexProp    string
	tmpModule       string
	tmpApi          string
	tmpApiComment   string
	tmpParamInput   string
	tmpParamSet     string
	tmpParamComment string
)

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
		{&tmpIndex, "Index"},
		{&tmpIndexImport, "IndexImport"},
		{&tmpIndexExport, "IndexExport"},
		{&tmpIndexProp, "IndexProp"},
		{&tmpModule, "Module"},
		{&tmpApi, "Api"},
		{&tmpApiComment, "ApiComment"},
		{&tmpParamInput, "ParamInput"},
		{&tmpParamComment, "ParamComment"},
		{&tmpParamSet, "ParamSet"},
	}
	for _, v := range fields {
		*v.target, err = strHelper.GetSection(temp, v.section)
		if err != nil {
			return
		}
	}

}

// Build 生成
func Build(setting *models.Setting, pathRoot string) (errorCode int) {
	var err error
	var modules []models.Module
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()
	modules, err = defines.LoadApis(setting.Define)
	if err != nil {
		return models.ErrorCodeDefineFormatError
	}
	err = buildIndex(pathRoot, modules)
	if err != nil {
		return models.ErrorCodeDefineFormatError
	}
	for _, m := range modules {
		err = buildModule(pathRoot, m)
		if err != nil {
			return models.ErrorCodeDefineFormatError
		}
	}
	return
}

// buildIndex 基于模板 生成index.ts文件
func buildIndex(pathRoot string, modules []models.Module) error {
	output := tmpIndex
	output = strings.Replace(output, "{{Imports}}", buildIndexImports(modules), 1)
	output = strings.Replace(output, "{{Exports}}", buildIndexExports(modules), 1)
	return io.SaveToFile(strings.Join([]string{pathRoot, path, "index.ts"}, "/"), []byte(output))
}
func buildIndexImports(modules []models.Module) string {
	var output []string
	for _, m := range modules {
		str := tmpIndexImport
		str = strings.Replace(str, "{{ModuleName}}", m.Name, -1)
		str = strings.Replace(str, "{{ModuleFile}}", strHelper.ToFileName(m.Name, ""), -1)
		output = append(output, str)
	}
	str := strings.Join(output, "")
	return strings.TrimRight(str, "\r\n,")
}
func buildIndexExports(modules []models.Module) string {
	var output []string
	for _, m := range modules {
		str := tmpIndexExport
		str = strings.Replace(str, "{{Name}}", m.Name, -1)
		str = strings.Replace(str, "{{Comment}}", m.Comment, -1)
		output = append(output, str)
	}
	str := strings.Join(output, "")
	return strings.TrimRight(str, "\r\n,")
}

func buildModule(pathRoot string, module models.Module) error {
	output := tmpModule
	groupPath := strings.ToLower(module.Route)
	output = strings.Replace(output, "{{Name}}", module.Name, 1)
	output = strings.Replace(output, "{{Route}}", groupPath, 1)
	var apis []string
	for _, group := range module.Groups {
		for _, api := range group.Apis {
			str := tmpApi
			str = strings.Replace(str, "{{Name}}", strHelper.ToCamel(api.Name), -1)
			str = strings.Replace(str, "{{Comment}}", buildApiComment(api), 1)
			str = strings.Replace(str, "{{Return}}", buildReturn(api.Return), 1)
			str = strings.Replace(str, "{{ParamsInputs}}", buildParamsInputs(api), 1)
			str = strings.Replace(str, "{{ParamsInputs}}", buildReturn(api.Return), 1)
			str = strings.Replace(str, "{{ReqType}}", api.ReqType, 1)
			str = strings.Replace(str, "{{Route}}", buildRoute(group.Name, api.Route), 1)
			str = strings.Replace(str, "{{ParamSets}}", buildParamSets(api), 1)
			apis = append(apis, str)
		}
	}
	apiStr := strings.TrimRight(strings.Join(apis, ""), ",\r\n")
	output = strings.Replace(output, "{{Apis}}", apiStr, 1)
	fileName := strHelper.ToFileName(module.Name, "ts")
	return io.SaveToFile(strings.Join([]string{pathRoot, path, fileName}, "/"), []byte(output))
}

// buildApiComment 基于模板 生成api注释
func buildApiComment(api models.Api) string {
	str := strings.TrimRight(tmpApiComment, " \r\n")
	str = strings.Replace(str, "{{Comment}}", api.Comment, 1)
	var paramComments []string
	for _, p := range api.Query {
		paramComments = append(paramComments, fmt.Sprintf("%s:%s ", p.Name, p.Comment))
	}
	for _, p := range api.Json {
		paramComments = append(paramComments, fmt.Sprintf("%s:%s ", p.Name, p.Comment))
	}
	str = strings.Replace(str, "{{ParamComments}}", strings.Join(paramComments, ","), 1)
	return str
}

// buildRoute 基于模板 生成路由
func buildRoute(group string, route string) string {
	return strings.Join([]string{strHelper.ToSnake(group), route}, "/")
}

// buildParamSets 基于模板 生成函数赋值表达式
func buildParamSets(api models.Api) string {

	var output []string
	for _, p := range api.Json {
		output = append(output, buildParamSetsSub(p))
	}
	for _, p := range api.Query {
		output = append(output, buildParamSetsSub(p))
	}

	return strings.TrimRight(strings.Join(output, " "), ",\r\n")
}

func buildParamSetsSub(p models.Prop) string {
	str := tmpParamSet
	str = strings.Replace(str, "{{Name}}", p.Name, -1)
	return str
}

// buildReturn 基于模板 生成函数返回
func buildReturn(prop models.Prop) string {
	name := strHelper.ConvertToTsType(prop.TypeName)
	if !strHelper.IsTsValueType(prop.TypeName) {
		name = "Models." + name
	}
	if prop.IsArray {
		return name + "[]"
	}
	return name
}

// buildParamSets 基于模板 生成函数形参
func buildParamsInputs(api models.Api) string {
	var output []string
	for _, p := range api.Json {
		output = append(output, buildParamInputSub(p))
	}
	for _, p := range api.Query {
		output = append(output, buildParamInputSub(p))
	}
	return strings.TrimRight(strings.Join(output, ","), ",\r\n")
}

func buildParamInputSub(p models.Prop) string {
	str := strings.Trim(tmpParamInput, " \r\n")
	str = strings.Replace(str, "{{Name}}", p.Name, 1)
	tName := strHelper.ConvertToTsType(p.TypeName)
	if !strHelper.IsTsValueType(p.TypeName) {
		tName = "Models." + tName
	}
	str = strings.Replace(str, "{{Type}}", tName, 1)
	return str
}
