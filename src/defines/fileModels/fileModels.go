/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/12/1 17:09
 */

package fileModels

// FileModule 模块
type FileModule struct {
	Module  string
	Route   string
	Groups  []FileGroup
	Notices []FileNotice
}

// FileGroup 分组
type FileGroup struct {
	Group string
	Heads []string
	Apis  []FileApi
}

// FileApi api
type FileApi struct {
	Api string
	//查询参数
	Query []string
	//json参数
	Json []string
	//头参数
	Head []string
	//返回值
	Return string
}
type FileModels struct {
	Models []FileModel
}

// FileModel 定义
type FileModel struct {
	Model string
	Props []string
}

type FileRoutes struct {
	Root   string
	Routes []string
}
type FileNotice struct {
	Notice string
	Params []string
}
